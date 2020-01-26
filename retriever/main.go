package main

/**
 * 组合接口以及继承
 */

import (
	"fmt"
	"learn/retriever/mock"
	"learn/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string] string) string
}

const url = "https://www.imooc.com"

func post(poster Poster){
	poster.Post(url, map[string]string{
		"name":"handsome",
		"course":"golang",
	})
}

// 组合接口
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents":"another faked imooc.com",
	})

	return s.Get(url)
}

func download(r Retriever) string {
	return r.Get(url)
}

func main(){
	var r Retriever
	retriever := mock.Retriever{"this is fake imooc.com"}
	r = &retriever
	inspect(r)

	fmt.Println()
	r = &real.Retriever{
		UserAgent: "Mozilala/5.0",
		TimeOut: time.Minute,
	}
	inspect(r)

	// Type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	//fmt.Println(download(r))
	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}

func inspect(r Retriever){
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	switch v:= r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:",v.Content)
	case *real.Retriever:
		fmt.Println("UserAgent:",v.UserAgent)
	}
}



