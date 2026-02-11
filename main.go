package main

import (
	"fmt"
	"log/slog"
	"os"
)

func help() {
	fmt.Printf("Usage:\n\t%v [-h|--help] <executables: name1 name2 ... nameN>\n", os.Args[0])
	os.Exit(0)
}

func parsePath(path string) (paths []string) {
	
	var prev int
	
	for i, r := range []rune(path) {
		if r == ':' {
			paths = append(paths, path[prev:i])
			prev = i + 1
			
		}
	}
	
	if prev < len(path) - 1 {
		paths = append(paths, path[prev:len(path) - 1])
	}
	
	return paths
}

func validatePaths(paths []string) (result []string) {
	
	var (
		info os.FileInfo
		err error
	)
	
	for i, path := range paths {
		info, err = os.Stat(path)
		
		if err != nil {
			slog.Error(fmt.Sprintf("Path number %v '%v': %v", i + 1, path, err))
			continue
		}
		
		if !info.IsDir() {
			slog.Error(fmt.Sprintf("Path number %v '%v' is not a directory", i + 1, path))
			continue
		}
		
		result = append(result, path)
	}
	
	return result
}

func main() {
	args := os.Args
	
	if len(args) < 2 || args[1] == "-h" || args[1] == "--help" {
		help()
	}
	
	paths := parsePath(os.Getenv("PATH"))
	
	paths = validatePaths(paths)
	
	fmt.Println(paths)
	
}