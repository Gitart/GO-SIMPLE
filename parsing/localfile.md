## Go Colly scrape local files

We can scrape files on a local disk.

words.html

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Document title</title>
</head>
<body>
<p>List of words</p>
<ul>
    <li>dark</li>
    <li>smart</li>
    <li>war</li>
    <li>cloud</li>
    <li>park</li>
    <li>cup</li>
    <li>worm</li>
    <li>water</li>
    <li>rock</li>
    <li>warm</li>
</ul>
<footer>footer for words</footer>
</body>
</html>
```
We have this HTML file.

local.go

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/gocolly/colly/v2"
)

func main() {

    t := &http.Transport{}
    t.RegisterProtocol("file", http.NewFileTransport(http.Dir(".")))

    c := colly.NewCollector()
    c.WithTransport(t)

    words := \[\]string{}

    c.OnHTML("li", func(e \*colly.HTMLElement) {
        words = append(words, e.Text)
    })

    c.Visit("file://./words.html")

    for \_, p := range words {
        fmt.Printf("%s\\n", p)
    }
}
```
To scrape local files, we must register a file protocol. We scrape all the words from the list.

$ go run local.go

```
dark
smart
war
cloud
park
cup
worm
water
rock
warm
```
