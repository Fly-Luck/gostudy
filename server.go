// 1. simple use of http handler and detailed request/response
// 2. simple use of Mutex lock
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
	//http.HandleFunc("/", handler1)
	// "handler" is a reference to a function
	// all requests under root path will be handled by "handler"
	//log.Fatal(http.ListenAndServe("localhost:8080", nil))

	//	http.HandleFunc("/", handler2)
	//	http.HandleFunc("/count", counter) // only handles /count, calling /count doesn't count
	//	log.Fatal(http.ListenAndServe("localhost:8080", nil))

	http.HandleFunc("/", handler3)
	lissajousHandler := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w) // output gif to response
	}
	// anonymous func alternative:
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	lissajous(w)
	// })
	http.HandleFunc("/lissajous", lissajousHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

// handler request/response
func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path %q\n", r.URL.Path)
}

// handler request/response with mutex lock
func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock() //lock the mutex lock
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// handle request/response with detailed info
func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v) // write to ResponseWriter
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil { // a simple sentence before the "if" judge condition
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

}

// show counts
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
