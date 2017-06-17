package main

import (
	"fmt"
	"os"
)

func main() {
	sep := "" // we (sep)arate each string with a space, but not the first time.
	for i := 1; i < len(os.Args); i++ {
		fmt.Print(sep)
		fmt.Print(os.Args[i])
		sep = " "
	}
	fmt.Println()
}
