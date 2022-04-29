package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

//Define a regex list for identifying interesting filenames
var regexes = []*regexp.Regexp{
	regexp.MustCompile(`(?i)user`),
	regexp.MustCompile(`(?i)password`),
	regexp.MustCompile(`(?i)secret`),
	regexp.MustCompile(`(?i)login`),
}

//Create the function which accepts a file path and parametres
func walkFn(path string, file os.FileInfo, err error) error {
	//Create a loop checks  regular expression list for matches
	for _, r := range regexes {
		if r.MatchString(path) {
			fmt.Printf("[!]FOUND: %s[!]\n", path)
		}
	}
	return nil
}

func main() {
	root := os.Args[1]
	if err := filepath.Walk(root, walkFn); err != nil {
		log.Panic(err)
	}
}
