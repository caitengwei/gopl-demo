package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func counter(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "visited = %d\n", count)
}

func handler2(writer http.ResponseWriter, request *http.Request) {
	incr()
	fmt.Fprintf(writer, "url.path = %q\n", request.URL.Path)
}


func incr() {
	mu.Lock()
	count++
	mu.Unlock()
}
