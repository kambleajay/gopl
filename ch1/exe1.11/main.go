//Fetchall fetches URLs in parallel and reports their times and sizes
//this exercise is to attempt fetching a long list of URLs and examine behavior
//also examine behavior if a website doesn't respond
package main

import (
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
		go fetch(url, ch) //start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)  //receive from channel ch
	}
	fmt.Printf("%.2fs elasped\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%2fs %7d %s", secs, nbytes, url)
}

//a long list from alexa global top sites:
// https://www.google.com https://www.youtube.com https://www.facebook.com https://www.baidu.com https://www.wikipedia.org https://www.yahoo.com https://im.qq.com https://www.taobao.com https://www.tmall.com https://www.amazon.com https://kambleajay.com https://www.twitter.com https://www.instagram.com https://www.google.co.in https://www.sohu.com https://www.jd.com

//Observation: all the fetches start in parallel, total time taken is the time taken by the longest fetch
