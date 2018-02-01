// 1. use of gorountine and channel
package basic

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	//"os"
	"time"
)

/*
func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //launch a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) //receive from channel, pass value of channel: <-ch
	}
	fmt.Printf("%.2fs elpsed\n", time.Since(start).Seconds())
}
*/

// http get and return info to channel
func fetch(url string, ch chan<- string) {
	start := time.Now()
	rsp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, rsp.Body)
	rsp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint("while reading %s: %v", url, err) //pass expression to channel: ch<-[expression]
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
