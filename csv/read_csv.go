package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
)

var posts = []byte(
	`#id,title,text
          1,hello world,"This is a ""blog""."
          2,second time,"Mysecondentry."
          3,ddd time,"Mysecondentry."

`)

func main() {
	// r can be any io.Reader, including a file.
	r := bytes.NewReader(posts)
	csvReader := csv.NewReader(r)
	// Set comment character to '#'.
	csvReader.Comment = '#'
	for {
		post, err := csvReader.Read()
		if err != nil {
			log.Println(err)
			// Will break on EOF.
			break
		}
		fmt.Printf("post with id %s is titled %q: %q\n", post[0], post[1], post[2])
	}
}
