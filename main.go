package main

import (
	"fmt"
	"os"
)

func main() {
    args := os.Args[1:]

	if (len(args) > 1) {
		fmt.Println("Usage: cher [script]")
		os.Exit(0)
 	} else if (len(args) == 1) {
		fmt.Println("Run: cher [file]")
  	} else {
		fmt.Println("Run: cher [prompt]")
  	}
}
