//Fetchall fetches URLs in parallel and reports their times and sizes
//their is code to examine cache behavior, the output is saved to file to examine
package main

import (
	"fmt"
	"io"
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

	fileName := fmt.Sprintf("fetch-%d.txt", time.Now().UnixNano() / 1000000)
	file, _ := os.Create(fileName)
	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%2fs %7d %s", secs, nbytes, url)
}

//conclusion: there is some caching at work, for different sites the first request takes longer, but all subsequent requests take less time, also the content is same
//Here is one outcome:
// ajay@Ajays-MacBook-Pro ~/l/g/c/exe1.10 (master)> ./main https://projecteuler.net/
// 4.699022s    4648 https://projecteuler.net/
// 4.70s elasped
// ajay@Ajays-MacBook-Pro ~/l/g/c/exe1.10 (master)> ./main https://projecteuler.net/
// 0.842593s    4648 https://projecteuler.net/
// 0.84s elasped
// ajay@Ajays-MacBook-Pro ~/l/g/c/exe1.10 (master)> ./main https://projecteuler.net/
// 0.765554s    4648 https://projecteuler.net/
// 0.77s elasped
