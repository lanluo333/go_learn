package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type userError string

func (e userError) Error() string  {
	return e.Message()
}

func (e userError) Message() string  {
	return string(e)
}

const prefix  = "/list/"
func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	// strings.Index(s string, str string) int：判断str在s中首次出现的位置，如果没有出现，则返回-1
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("path must start with " + prefix)
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	all ,err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)

	return nil
}

