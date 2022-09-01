## Go Colly crawl links

In more complex tasks, we need to crawl links that we find in our documents.

links.go

package main

import (
    "fmt"

    "github.com/gocolly/colly/v2"
)

func main() {

    c := colly.NewCollector()

    c.OnHTML("title", func(e \*colly.HTMLElement) {
        fmt.Println(e.Text)
    })

    c.OnHTML("a\[href\]", func(e \*colly.HTMLElement) {

        link := e.Attr("href")
        c.Visit(e.Request.AbsoluteURL(link))
    })

    c.Visit("http://webcode.me/small.html")
}
