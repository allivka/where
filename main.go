package main

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

func sliceToMap[T comparable](slice []T) map[T]struct{} {
	result := make(map[T]struct{}, len(slice))

	for _, element := range slice {
		result[element] = struct{}{}
	}

	return result
}

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

	if prev < len(path)-1 {
		paths = append(paths, path[prev:len(path)-1])
	}

	return paths
}

func validatePaths(paths []string) (result []string) {

	var (
		info os.FileInfo
		err  error
	)

	for i, path := range paths {
		info, err = os.Stat(path)

		if err != nil {
			slog.Error(fmt.Sprintf("Path number %v '%v': %v", i+1, path, err))
			continue
		}

		if !info.IsDir() {
			slog.Error(fmt.Sprintf("Path number %v '%v' is not a directory", i+1, path))
			continue
		}

		result = append(result, path)
	}

	return result
}

func getPathFiles(path string) (names []string, err error) {

	dir, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}

	stat, err := dir.Stat()
	if err != nil {
		return []string{}, err
	}

	if !stat.IsDir() {
		return []string{}, fmt.Errorf("File '%v' is not a directory", path)
	}

	entries, err := dir.Readdir(0)

	if err != nil {
		slog.Error(fmt.Sprintf("Failed reading inside of directory '%v': %v", path, err))
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			names = append(names, entry.Name())
		}
	}

	return names, nil
}

func main() {
	args := os.Args

	if len(args) < 2 || args[1] == "-h" || args[1] == "--help" {
		help()
	}

	paths := parsePath(os.Getenv("PATH"))

	paths = validatePaths(paths)

	files := make(map[string]map[string]struct{}, len(paths))

	for i, path := range paths {
		names, err := getPathFiles(path)

		if err != nil {
			slog.Error(fmt.Sprintf("Path number %v '%v': %v", i+1, path, err))
		}

		files[path] = sliceToMap(names)
	}

	target := args[1]

	for _, path := range paths {
		if _, ok := files[path][target]; ok {
			fmt.Println(filepath.Join(path, target))
		}
	}

}
