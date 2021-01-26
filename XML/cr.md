package main

import (
	"encoding/xml"
	"log"
	"os"
)

type Character struct {
	XMLName     struct{} `xml:"character"`
	Name        string   `xml:"name"`
	Surname     string   `xml:"surname"`
	Job         string   `xml:"details>job,omitempty"`
	YearOfBirth int      `xml:"year_of_birth,attr,omitempty"`
	IgnoreMe    string   `xml:"-"`
}

func main() {
	e := xml.NewEncoder(os.Stdout)
	e.Indent("", "\t")
	c := Character{
		Name:        "Henry",
		Surname:     "Wentworth Akeley",
		Job:         "farmer",
		YearOfBirth: 1871,
		IgnoreMe:    "somevalue",
	}
	if err := e.Encode(c); err != nil {
		log.Fatalln(err)
	}
}
