// A simple cat(1) implementation in Go.
//
// Build Instructions:
//
//   $ go build -o cat
//
// Usage:
//
//  Basically, the usage is same as cat(1).
//
//   $ ./cat [-n] [file ...]
//
// Options:
//
// -n: Number output lines.
// -h: Display usage.
//
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var showNumLines = flag.Bool("n", false, "flat to output the line numbers")

func readLines(r io.Reader) {
	rd := bufio.NewReader(r)
	lineNum := 1
	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if *showNumLines {
			fmt.Printf("\r%6d  %s", lineNum, line)
		} else {
			fmt.Printf("%s", line)
		}
		lineNum++
	}
}

func cat(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	readLines(file)
}

func main() {
	flag.Parse()
	for i := 0; i < flag.NArg(); i++ {
		cat(flag.Arg(i))
	}
}
