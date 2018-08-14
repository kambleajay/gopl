//Exe1.3 benchmarks loop version and join version of echo, two ways are used to benchmark
//first is to use methods from time package, the second way (in the accompanying test) is to use benchmarking from testing package
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echoWithLoop() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echoWithJoin() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	start := time.Now()
	echoWithLoop()
	fmt.Println("Time taken/loop", time.Since(start).Nanoseconds())

	start = time.Now()
	echoWithJoin()
	fmt.Println("Time taken/join", time.Since(start).Nanoseconds())
}

//Conclusion: on my setup, loop took roughly 35000 nanoseconds and join took roughly 2000 nanoseconds

