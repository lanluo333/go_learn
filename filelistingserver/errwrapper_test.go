package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

type testingUserError string

func (e testingUserError) Error() string  {
	return e.Message()
}

func (e testingUserError) Message() string  {
	return string(e)
}

func errPanic(writer http.ResponseWriter, request *http.Request) error{
	panic(123)
}

func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknown error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrExist
}

func noError(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

var tests = []struct{
	h appHandler
	code int
	message string
}{
	{errPanic, 500, ""},
	{errNotFound, 404, "Not Found"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

func TestErrWrapper(t *testing.T){
	for _,tt := range tests{
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
		f(response, request)

		verifyResponse(response.Result(), tt.code, tt.message, t)
	}

}

func verifyResponse(resp *http.Response, expectedCode int, expectMsg string, t *testing.T){
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode || body != expectMsg {
		t.Errorf("expect (%d, %s); "+ "got (%d, %s)", expectedCode, expectMsg, resp.StatusCode, body)
	}
}

func TestErrWrapperInServer(t *testing.T){
	for _,tt := range tests{
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)

		verifyResponse(resp, tt.code, tt.message, t)
	}
}


