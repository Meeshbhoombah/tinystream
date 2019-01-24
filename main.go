package main

import (
	"fmt"
	"os"
)

const (
	VERSION = "0.0.0"
	// USAGE = ""
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please include seed to start using `tinystream`.")
	}
}
