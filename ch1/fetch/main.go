//Fetch prints the content found at a url
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		// Check for http prefix
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}

		// http.Get makes an HTTP request and returns the result
		// if there's no error, as the response struct 'resp'
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// The 'Body' part of resp contains the server response as a
		// readable stream. Next, Copy copies the body,
		// storing it to b
		b, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s \n Status code: %v", b, resp.Status)
	}
}
