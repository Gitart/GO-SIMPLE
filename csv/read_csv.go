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





package main

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Open the file
	csvfile, err := os.Open("input.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Question: %s Answer %s\n", record[0], record[1])
	}
}





