Stackoverflow questions
In the next example, we scrape Stackoverflow questions.

stack.go
package main

import (
    "fmt"

    "github.com/gocolly/colly/v2"
)

type question struct {
    title   string
    excerpt string
}

func main() {

    c := colly.NewCollector()
    qs := []question{}

    c.OnHTML("div.summary", func(e *colly.HTMLElement) {

        q := question{}
        q.title = e.ChildText("a.question-hyperlink")
        q.excerpt = e.ChildText(".excerpt")

        qs = append(qs, q)

    })

    c.OnScraped(func(r *colly.Response) {
        for idx, q := range qs {

            fmt.Println("---------------------------------")
            fmt.Println(idx + 1)
            fmt.Printf("Q: %s \n\n", q.title)
            fmt.Println(q.excerpt)
        }
    })

    c.Visit("https://stackoverflow.com/questions/tagged/perl")
}
