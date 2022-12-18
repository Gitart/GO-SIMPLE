https://play.golang.org/p/esopq0SMhG_T

```go
package main

import (
	"encoding/xml"
	"log"
	"strings"
)

type Character struct {
	Name        string `xml:"name"`
	Surname     string `xml:"surname"`
	Job         string `xml:"job,omitempty"`
	YearOfBirth int    `xml:"year_of_birth,omitempty"`
}

func main() {
	r := strings.NewReader(`<?xml version="1.0" encoding="UTF-8"?>
<character>
    <name>Herbert</name>
    <surname>West</surname>
    <job>Scientist</job>
</character>
}`)
	d := xml.NewDecoder(r)
	var c Character
	if err := d.Decode(&c); err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v", c)
}
```
