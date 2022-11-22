<?xml version="1.0" encoding="utf-8"?>
<lists version="1">
<Paragraph>
    <list>11</list>
    <Paragraph>
          <list>1111</list>
          <list>1122</list>

    </Paragraph>
</Paragraph>
<Paragraph>
    <list>22</list>
    <Paragraph>
            <list>2222</list>
            <list>3333</list>
            <Paragraph>
                <list>222222</list>
                <list>333333</list>
            </Paragraph>
            <list>4444</list>
            <list>5555</list>
            <Paragraph>
                <list>444444</list>
                <list>555555</list>
            </Paragraph>
    </Paragraph>
</Paragraph>
</lists>

package main

import (
"encoding/xml"
"fmt"
"io/ioutil"
"os"
)

type lists struct {
XMLName xml.Name    `xml:"lists"`
Version string      `xml:"version,attr"`
Svs     []Paragraph `xml:"Paragraph"`
}

type Paragraph struct {
List []string    `xml:"list"`
Svs  []Paragraph `xml:"Paragraph,omitempty"`
}

func main() {
file, err := os.Open("config.xml") // For read access.
if err != nil {
    fmt.Printf("error: %v", err)
    return
}
defer file.Close()
data, err := ioutil.ReadAll(file)
if err != nil {
    fmt.Printf("error: %v", err)
    return
}
v := lists{}
err = xml.Unmarshal(data, &v)
if err != nil {
    fmt.Printf("error: %v", err)
    return
}

//str := v.Svs
for _, vf := range v.Svs {
    fmt.Println(vf.List)
    fmt.Println(vf.Svs[0].List)

}

}







func main() {
    // ...
    PrintParagraphs(v.Svs)
}

func PrintParagraphs(ps []Paragraph) {
    for _, p := range ps {
        fmt.Println(p.List)
        PrintParagraphs(p.Svs)
    }
}

Output

[11]
[1111 1122]
[22]
[2222 3333 4444 5555]
[222222 333333]
[444444 555555]
