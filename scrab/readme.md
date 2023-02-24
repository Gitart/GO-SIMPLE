# Web Scraping With Go

----
[Colly](http://go-colly.org/docs/examples/random_delay)     
[Go Query](https://github.com/PuerkitoBio/goquery?ref=morioh.com&utm_source=morioh.com#examples)    



Web scraping is an essential tool that every developer uses at some point in their career. Hence, it is essential for developers to understand what a web scraper is, as well as how to build one. [Wikipedia](https://en.wikipedia.org) defines *[web scraping](https://en.wikipedia.org/wiki/Web_scraping)* as follows:

> Web scraping, web harvesting, or web data extraction is data scraping used for extracting data from websites. The web scraping software may directly access the World Wide Web using the Hypertext Transfer Protocol or a web browser. While web scraping can be done manually by a software user, the term typically refers to automated processes implemented using a bot or web crawler. It is a form of copying in which specific data is gathered and copied from the web, typically into a central local database or spreadsheet, for later retrieval or analysis.

In other words, *web scraping* is a process for extracting data from websites and is used in many cases, ranging from data analysis to lead generation. The task can be completed manually or can be automated through a script or software.

There are a variety of use cases for web scraping. Take a look at just a few:

*   **Gathering data:** The most useful application or use of web scraping is data gathering. Data is compelling, and analyzing data in the right way can put one company ahead of another. Web scraping is an essential tool for data gathering—writing a simple script can make data gathering much more accessible and faster than doing manual jobs. Furthermore, the data can also be inputted into a spreadsheet to be better visualized and analyzed.

*   **Performing market research and lead generation:** Doing market research and generating leads are crucial web scraping tasks. Emails, phone numbers, and other important information from various websites can be scraped and later used for these important tasks.

*   **Building price comparison tools:** You might have noticed browser extensions that alert you of a price change for products on e-commerce platforms. Such tools are also built using web scrapers.

In this article, you will learn how to create a simple web scraper using [Go](https://golang.org/).

Robert Griesemer, Rob Pike, and Ken Thompson created the Go programming language at Google, and it has been in the market since 2009. Go, also known as Golang, has many brilliant features. Getting started with Go is fast and straightforward. As a result, this comparatively newer language is gaining a lot of attraction in the developer world.

## Implementing Web Scraping with Go

The support for concurrency has made Go a fast, powerful language, and because the language is easy to get started with, you can build your web scraper with only a few lines of code. For creating web scrapers with Go, two libraries are very popular:

1.  [goquery](https://github.com/PuerkitoBio/goquery)
2.  [Colly](http://go-colly.org/)

In this article, you’ll be using Colly to implement the scraper. At first, you’ll be learning the very basics of building a scraper, and you’ll implement a URL scraper from a Wikipedia page. Once you know the basic building blocks of web scraping with Colly, you’ll level up the skill and implement a more advanced scraper.

## Prerequisites

Before moving forward in this article, be sure that the following tools and libraries are installed on your computer. You’ll need the following:

*   Basic understanding of Go
*   [Go](https://go.dev/) (preferably the latest version—1.17.2, as of writing this article)
*   IDE or text editor of your choice ([Visual Studio Code](https://code.visualstudio.com/) preferred)
*   [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go) for the IDE (if available)

## Understanding Colly and the `Collector` Component

The Colly package is used for building web crawlers and scrapers. It is based on Go’s Net/HTTP and goquery package. The goquery package gives a jQuery-like syntax in Go to target HTML elements. This package alone is also used to build scrapers.

The main component of Colly is the `Collector`. According to the [docs](http://go-colly.org/docs/introduction/start/), the `Collector` component manages the network communications, and it is also responsible for the callbacks attached to it when a `Collector` job is running. This component is configurable, and you can modify the `UserAgent` string or add `Authentication` headers, restricting or allowing URLs with the help of this component.

Tired of getting blocked while scraping the web?

Join 20,000 users using our API to get the data they need!

[Try ScrapingBee for Free](https://app.scrapingbee.com/account/register)

## Understanding Colly Callbacks

Callbacks can also be added to the `Collector` component. The Colly library has callbacks, such as `OnHTML` and `OnRequest`. You can refer to the [docs](http://go-colly.org/docs/introduction/start/) to learn about all the callbacks. These callbacks run at different points in the life cycle of the `Collector`. For example, the `OnRequest` callback is run just before the `Collector` makes an HTTP request.

The `OnHTML` method is the most common callback used in building web scrapers. It allows registering a callback for the `Collector` when it reaches a specific HTML tag on the web page.

## Initializing Project Directory and Installing Colly

Before starting to write code, you have to initialize the project directory. Open the IDE of your choice and open a folder where you will save all your project files. Now, open a terminal window, and locate your directory. After, type the following command in the terminal:

```bash
go mod init github.com/Username/Project-Name

```

In the above command, change `[github.com](http://github.com)` to the domain where you store your files, such as [Bitbucket](https://bitbucket.org/) or [Gitlab](https://about.gitlab.com). Also, change `Username` to your username and `Project-Name` with whatever project name you would like to give it.

Once you type in the command and press enter, you’ll find that a new file is created with the name of `go.mod`. This file holds the information about the direct and indirect dependencies that the project needs. The next step is to install the Colly dependency. To install the dependency, type the following command in the terminal:

```bash
go get -u github.com/go-colly/colly/...

```

This will download the Colly library and will generate a new file called `go.sum`. You can now find the dependency in the `go.mod` file. The `go.sum` file lists the checksum of the direct and indirect dependencies, along with the version. You can read more about the `go.sum` and `go.mod` files [here](https://golangbyexample.com/go-mod-sum-module/).

## Building a Basic Scraper

Now that you have set up the project directory with the necessary dependency, you can move forward with writing some codes. The basic scraper aims to scrape all the links from a specific Wikipedia page and print them on the terminal. This scraper is built to make you comfortable with the building blocks of the Colly library.

Create a new file in the folder with an extension of `.go`—for example, `main.go`. All the logic will go into this file. Start by writing `package main`. This line tells the compiler that the package should compile as an executable program instead of a shared library.

```go
package main

```

The next step is to start writing the `main` function. If you are using Visual Studio Code, it will do the importing of the necessary packages automatically. Otherwise, in the case of other IDEs, you may have to do it manually. The `Collector` of Colly is initialized with the following line of code:

```go
func main() {
    c := colly.NewCollector(
        colly.AllowedDomains("en.wikipedia.org"),
    )
}

```

Here, the `NewCollector` is initialized, and as an option, [en.wikipedia.org](http://en.wikipedia.org) is passed as an allowed domain. The same `Collector` can also be initialized without passing any option to it. Now, if you save the file, Colly will be automatically imported to your `main.go` file; if not, add the following lines after the `package main` line:

```go
import (
    "fmt"

    "github.com/gocolly/colly"
)

```

The above lines import two packages in the `main.go` file. The first package is the `fmt` package and the second one is the Colly library.

Now, open [this URL](https://en.wikipedia.org/wiki/Web_scraping) in your browser. This is the Wikipedia page on web scraping. The web scraper is going to scrape all the links from this page. Understanding the browser developer tools well is an invaluable skill in web scraping. Open the browser inspect tools by right-clicking on the page and selecting **Inspect**. This will open the page inspector. You’ll be able to see the whole HTML, CSS, network calls, and other important information from here. For this example specifically, find the `mw-parser-output` div:

![Wikipedia in Dev Tools](https://www.scrapingbee.com/blog/web-scraping-go/30ihmtB.png)

This div element contains the body of the page. Targeting the links inside this div will provide all the links used inside the article.

Next, you will use the `OnHTML` method. Here is the remaining code for the scraper:

```go
// Find and print all links
    c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
        links := e.ChildAttrs("a", "href")
        fmt.Println(links)
    })
    c.Visit("https://en.wikipedia.org/wiki/Web_scraping")

```

The `OnHTML` method takes in two parameters. The first parameter is the HTML element. Reaching it is going to execute the callback function, which is passed as the second parameter. Inside the callback function, the `links` variable is assigned to a method that returns all the child attributes matching the element’s attributes. The `e.ChildAttrs("a", "href")` function [returns](https://pkg.go.dev/github.com/gocolly/colly?utm_source=godoc#HTMLElement.ChildAttrs) a slice of strings of all the links inside the `mw-parser-output` div. The `fmt.Println(links)` function prints the links in the terminal.

Finally, visit the URL using the `c.Visit("https://en.wikipedia.org/wiki/Web_scraping")` command. The complete scraper code will look like this:

```go
package main

import (
    "fmt"

    "github.com/gocolly/colly"
)

func main() {
    c := colly.NewCollector(
        colly.AllowedDomains("en.wikipedia.org"),
    )

    // Find and print all links
    c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
        links := e.ChildAttrs("a", "href")
        fmt.Println(links)
    })
    c.Visit("https://en.wikipedia.org/wiki/Web_scraping")
}

```

Running this code with the command `go run main.go` will get all the links on the page.

## Scraping Table Data

![W3Schools HTML Table](https://www.scrapingbee.com/blog/web-scraping-go/c3Owwna.png)

To scrape the table data, you can either remove the codes you have written inside `c.OnHTML` or create a new project by following the same steps mentioned above. To make and write a CSV file, you’ll be using the `encoding/csv` library available in Go. Here is the starter code:

```go
package main

import (
    "encoding/csv"
    "log"
    "os"
)

func main() {
    fName := "data.csv"
    file, err := os.Create(fName)
    if err != nil {
        log.Fatalf("Could not create file, err: %q", err)
        return
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()
}

```

Inside the `main` function, the first action is to define the file name. Here, it is defined as `data.csv`. Then using the `os.Create(fName)` method, the file is created with the name `data.csv`. If any error occurs during the creation of the file, it’ll also log the error and exit the program. The `defer file.Close()` command will close the file when the surrounding function returns.

The `writer := csv.NewWriter(file)` command initializes the CSV writer to write to the file, and the `writer.Flush()` will throw everything from the buffer to the writer.

Once the file creation process is done, the scraping process can be started. This is similar to the above example.

Next, add the below lines of code after the `defer writer.Flush()` line ends:

```go
c := colly.NewCollector()
    c.OnHTML("table#customers", func(e *colly.HTMLElement) {
        e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
            writer.Write([]string{
                el.ChildText("td:nth-child(1)"),
                el.ChildText("td:nth-child(2)"),
                el.ChildText("td:nth-child(3)"),
            })
        })
        fmt.Println("Scrapping Complete")
    })
    c.Visit("https://www.w3schools.com/html/html_tables.asp")

```

In this code, Colly is being initialized. Colly uses the `ForEach` method to iterate through the content. Because the table has three columns or `td` elements, using the `nth-child` pseudo selector, three columns are selected. `el.ChildText` returns the text inside the element. Putting it inside the `writer.Write` method will write the elements into the CSV file. Finally, the print statement prints a message when the scraping is complete. Because this code is not targeting the table headers, it will not print the heading. The complete code for this scraper will be like this:

```go
package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"

    "github.com/gocolly/colly"
)

func main() {
    fName := "data.csv"
    file, err := os.Create(fName)
    if err != nil {
        log.Fatalf("Could not create file, err: %q", err)
        return
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    c := colly.NewCollector()
    c.OnHTML("table#customers", func(e *colly.HTMLElement) {
        e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
            writer.Write([]string{
                el.ChildText("td:nth-child(1)"),
                el.ChildText("td:nth-child(2)"),
                el.ChildText("td:nth-child(3)"),
            })
        })
        fmt.Println("Scrapping Complete")
    })
    c.Visit("https://www.w3schools.com/html/html_tables.asp")
}

```

Once successful, the output will appear like this:

![Output CSV File in Excel](https://www.scrapingbee.com/blog/web-scraping-go/iet6dxt.png)

## Conclusion

In this article, you learned what web scrapers are as well as some use cases and how they can be implemented with Go, with the help of the Colly library.

However, the methods described in this tutorial are not the only possible way of implementing a scraper. Consider experimenting with this yourself and finding new ways to do it. Colly can also work together with the `goquery` library to make a more powerful scraper.

Depending on your use case, you can modify Colly to satisfy your needs. Web scraping is very handy for keyword research, brand protection, promotion, website testing, and many other things. So knowing how to build your own web scraper can help you become a better developer.
