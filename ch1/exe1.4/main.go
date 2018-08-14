package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileNames := make(map[string][]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, fileNames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileNames)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, fileNames[line])
		}
	}
}

func contains(a []string, elem string) bool {
	for _, n := range a {
		if n == elem {
			return true
		}
	}
	return false
}

func appendIfNotExists(fileNames map[string][]string, txt string, fileName string) []string {
	if contains(fileNames[txt], fileName) {
		return fileNames[txt]
	} else {
		return append(fileNames[txt], fileName)
	}
}

func countLines(f *os.File, counts map[string]int, fileNames map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		txt := input.Text()
		counts[txt]++
		fileNames[txt] = appendIfNotExists(fileNames, txt, f.Name())
	}
}
