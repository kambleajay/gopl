//Fetch prints the content found at a URL
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const httpPrefix = "http://"

func addHttpIfMissing(url string) string {
	if strings.HasPrefix(url, httpPrefix) {
		return url
	}
	return httpPrefix + url
}

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(addHttpIfMissing(url))
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
