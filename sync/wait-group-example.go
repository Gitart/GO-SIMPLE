package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	filesInHomeDir, err := ioutil.ReadDir(homeDir)
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	wg.Add(len(filesInHomeDir))

	fmt.Println("Printing files in", homeDir)

	for _, file := range filesInHomeDir {
		go func(f os.FileInfo) {
			defer wg.Done()
			fmt.Println(f.Name())
		}(file)
	}

	wg.Wait()
	fmt.Println("finished....")
}
