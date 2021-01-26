package main

import (
	"encoding/xml"
	"log"
	"os"
	"strings"
)

type Character struct {
	XMLName     struct{} `xml:"character"`
	ID          ID       `xml:"id,attr"`
	Name        string   `xml:"name"`
	Surname     string   `xml:"surname"`
	Job         string   `xml:"job,omitempty"`
	YearOfBirth int      `xml:"year_of_birth,omitempty"`
}

type ID string

func (i ID) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  xml.Name{Local: "codename"},
		Value: strings.ToUpper(string(i)),
	}, nil
}

func main() {
	e := xml.NewEncoder(os.Stdout)
	e.Indent("", "\t")
	c := Character{
		ID:          "aa",
		Name:        "Abdul",
		Surname:     "Alhazred",
		Job:         "poet",
		YearOfBirth: 700,
	}
	if err := e.Encode(c); err != nil {
		log.Fatalln(err)
	}
}
