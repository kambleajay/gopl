package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[:] {
		fmt.Printf("%d %s", index, arg)
		fmt.Println()
	}
}
