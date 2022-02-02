package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

type html struct {
	Body body `xml:"body>div"`
}
type body struct {
	Content string `xml:",innerxml"`
}

func main() {
	b := []byte(`<!DOCTYPE html>
<html>
    <head>
        <title>
            Title of the document
        </title>
    </head>
    <body>
        body content 
        <p>more content</p>
        <div>tyes</div>
        <div>tye22s</div>
    </body>
</html>`)

	h := html{}
	err := xml.NewDecoder(bytes.NewBuffer(b)).Decode(&h)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Println(h.Body.Content)
}
