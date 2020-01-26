package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, "https://www.imooc.com",nil)
	request.Header.Add("User-Agent","Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Mobile Safari/537.36")
	client := http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		fmt.Println("Redirect:",req)
		return nil
	}}
	//resp, err := http.Get("https://www.baidu.com")
	resp, err := client.Do(request)
	if err != nil {
		 panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", s)
}


