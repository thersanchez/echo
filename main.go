package main

import (
	"fmt"
	"os"
)

func main() {
	sep := "" // we (sep)arate each string with a space, but not the first time.
	for _, v := range os.Args[1:] {
		fmt.Print(sep)
		fmt.Print(v)
		sep = " "
	}
	fmt.Println()
}
