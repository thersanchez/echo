package main

import (
	"fmt"
	"os"
)

func main() {
	// we (sep)arate each string with a " ", but not the first one.
	sep := ""
	for i := 1; i < len(os.Args); i++ {
		fmt.Print(sep)
		fmt.Print(os.Args[i])
		sep = " "
	}
	fmt.Println()
}
