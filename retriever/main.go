package main

import (
	"fmt"
	"golearning/retriever/mock"
	"golearning/retriever/real"
	"time"
)

// 实现是隐式的，不需要实现接口，只需要实现接口里的所有方法
type Retriever interface {
	Get(url string) string
}

const url = "http://www.baidu.com"

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func post(poster Poster) {
	poster.Post("http://www.baidu.com", map[string]string{
		"name":   "kbh",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another mock",
	})

	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a mock"}
	inspect(r)
	fmt.Printf("%T %v\n", r, r)
	r = &real.Retriever{
		UserAgent: "Mozilla 5.0",
		Timeout:   time.Minute,
	}
	fmt.Printf("%T %v\n", r, r)
	// fmt.Println(download(r))
	inspect(r)

	// type assertion
	// realRetriever := r.(*real.Retriever)
	// fmt.Println(realRetriever.Timeout)
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a Session")
	fmt.Println(session(&retriever))
}

func inspect(r Retriever) {
	fmt.Println("Inspecting,")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("User Agent", v.UserAgent)
	}
	fmt.Println()
}
