package main

import (
    "fmt"

    "github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector()

    i    := 0
    scan := true

    c.OnHTML("#stores-list--section-16266 td.data-cell-0, td.data-cell-1, td.data-cell-2, td.data-cell-3",
            func(e *colly.HTMLElement) {

            if scan {

                fmt.Printf("%s ", e.Text)
            }

            i++

            if i%4 == 0 && i < 40 {
                fmt.Println()
            }

            if i == 40 {
                scan = false
                fmt.Println()
            }
        })

    c.Visit("https://nrf.com/resources/top-retailers/top-100-retailers/top-100-retailers-2019")
}
