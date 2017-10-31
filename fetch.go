// fetch
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		//fetch1(url)
		//fetch2(url)
		//fetch3(url)
		fetch4(url)
	}
}

//basic fetch
func fetch1(url string) {
	rsp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(rsp.Body)
	closeup(url, rsp, err)
	fmt.Printf("%s", b)
}

//replace ReadAll with Copy
func fetch2(url string) {
	rsp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	dst, err := os.Create("temp")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating temp file")
		os.Exit(1)
	}
	w, err := io.Copy(dst, rsp.Body)
	fmt.Println("bytes received: ", w)
	closeup(url, rsp, err)
}

//add http prefix
func fetch3(url string) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	fetch1(url)
}

//output response status
func fetch4(url string) {
	rsp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(rsp.Body)
	s := rsp.Status
	closeup(url, rsp, err)
	fmt.Printf("%s", b)
	fmt.Println("status: ", s)
}

func closeup(url string, rsp *http.Response, err error) {
	rsp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
}
