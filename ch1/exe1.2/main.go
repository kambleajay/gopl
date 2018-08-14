//Exe1.2 prints index and value of each command-line argument
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[:] {
		fmt.Println(index, arg)
	}
}
