package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type FileInfo struct {
	Name string
}

func main() {
	// Get files present in the "tmp" directory
	entries, err := ioutil.ReadDir("tmp/")

	if err != nil {
		fmt.Println(err)
	}

	// Display in a loop, on name file
	for _, entry := range entries {
		f := FileInfo{
			Name: entry.Name(),
		}
		fmt.Println(". " + f.Name)
	}

	var filename string
	fmt.Println("Enter filename ")
	fmt.Scanf("%s", &filename)

	file := "tmp/" + filename

	// Get extension
	extension := filepath.Ext(file)

	// Accepted extensions
	if extension == ".html" || extension == ".txt" {
		// Read file
		file, err := ioutil.ReadFile(file)

		if err != nil {
			fmt.Println(err)
		}

		// Display file content
		fmt.Println(string(file)) // convert byte to string
	} else {
		fmt.Println("Extension '" + extension + "', not accepted")
	}
}
