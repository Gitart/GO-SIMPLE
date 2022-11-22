// simulates "ls -l"

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func ls(path string) {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range entries {
		fmt.Printf("%v\t%d\t%v\t%s\n", f.Mode(), f.Size(), f.ModTime(), f.Name())
	}
}

func main() {
	flag.Parse()
	if flag.NArg() >= 1 {
		path := flag.Arg(0)
		ls(path)
	}
}
