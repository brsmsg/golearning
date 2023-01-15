package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc(
		"/list/",
		func(writer http.ResponseWriter, request *http.Request) {
			path := request.URL.Path // /list/fib.txt
			file, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			all, err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}

			writer.Write(all)
		},
	)
}
