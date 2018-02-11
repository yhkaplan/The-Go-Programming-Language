// Fetches URLs in parallel and reports their times and sizes
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		// Starts goroutine
		go fetch(url, ch)
	}

	f, _ := os.Create("results.txt")
	//Handle error

	// Writing defer f.Close here is a Go convention
	defer f.Close()

	// Creates a buffered writer from the file
	bufferedWriter := bufio.NewWriter(f)

	for range os.Args[1:] {
		// Receives from channel 'ch'
		bufferedWriter.WriteString(<-ch)
		//f.WriteString(<-ch)
	}
	o := fmt.Sprintf("%.2fs elapsed", time.Since(start).Seconds())
	bufferedWriter.WriteString(o)
	// Flush memory to disk
	bufferedWriter.Flush()
	//f.WriteString(o)
	//f.Sync()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		// Sends to channel ch
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	// Prevents resource leak
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v \n", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}
