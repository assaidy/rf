package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 4 {
		fmt.Fprintf(os.Stderr, "ERROR: missing arguments\n")
		usage()
		os.Exit(1)
	}
	args = args[1:]

	var (
		path        = args[0]
		pattern     = args[1]
		replacement = args[2]
	)

	if !checkFileExists(path) {
		fmt.Fprintf(os.Stderr, "ERROR: file doesn't exists\n")
		os.Exit(1)
	}

	fileName := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	newName := replaceWithRe(fileName, pattern, replacement)
	newPath := filepath.Join(filepath.Dir(path), newName+filepath.Ext(path))

	if err := renameFile(path, newPath); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: couldn't rename file\n")
		os.Exit(1)
	}
}

func checkFileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func replaceWithRe(src, pattern, repl string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(src, repl)
}

func renameFile(oldPath, newPath string) error {
	return os.Rename(oldPath, newPath)
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println("fr <FILE_PATH> <REGEX> <REPLACEMENT>")
	fmt.Println()
	fmt.Println("    FILE_PATH    the file path you want rename")
	fmt.Println("    REGEX        search pattern")
	fmt.Println("    REPLACEMENT  string replacement for regex matches")
}
