package main

import (
	"fmt"
	"os"
)

func help() {
	fmt.Printf("Usage:\n\t%v [-h|--help] <executables: name1 name2 ... nameN>\n", os.Args[0])
	os.Exit(0)
}

func main() {
	args := os.Args
	
	if len(args) < 2 || args[1] == "-h" || args[1] == "--help" {
		help()
	}
	
}