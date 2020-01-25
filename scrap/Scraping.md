# Web Scraping with Go

Submitted by [NanoDano](https://www.devdungeon.com/users/nanodano "View user profile.") on Sat, 03/24/2018 \- 20:43

## Overview

*   [Introduction](#intro)
*   [Ethics and guidelines of scraping](#ethics_and_guidelines)
*   [Prerequisites](#prerequisites)
*   [Make an HTTP GET request](#make_http_get_request)
*   [Make an HTTP GET request with timeout](#make_http_get_request_with_timeout)
*   [Set HTTP headers (Change user agent)](#set_http_headers_change_user_agent)
*   [Download a URL](#download_a_url)
*   [Use substring matching to find page title](#substring_matching)
*   [Use regular expressions to find HTML comments](#using_regular_expressions_to_find_html_comments)
*   [Use goquery to find all links on a page](#find_all_links_on_page)
*   [Parse URLs](#parse_urls)
*   [Use goquery to find all images on a page](#find_all_images_on_page)
*   [Make an HTTP POST request with data](#make_http_post_request_with_data)
*   [Make HTTP request with cookie](#make_http_request_with_cookie)
*   [Log in to a website](#log_in_to_website)
*   [Web crawling](#crawling)
*   [DevDungeon Project: Web Genome \- http://www.webgeno.me](#webgenome_project)
*   [Conclusion](#conclusion)

## Introduction

Web scraping ([Wikipedia entry](https://en.wikipedia.org/wiki/Web_scraping)) is a handy tool to have in your arsenal. It can be useful in a variety of situations, like when a website does not provide an API, or you need to parse and extract web content programmatically. This tutorial walks through using the standard library to perform a variety of tasks like making requests, changing headers, setting cookies, using regular expressions, and parsing URLs. It also covers the basics of the **goquery** package (a jQuery like tool) to scrape information from an HTML web page on the internet.

If you need to reverse engineering a web application based on the network traffic, it may also be helpful to learn how to do [packet capture, injection, and analysis with Gopacket](https://www.devdungeon.com/content/packet-capture-injection-and-analysis-gopacket).

If you are downloading and storing content from a site you scrape, you may be interested in [working with files in Go](https://www.devdungeon.com/content/working-files-go).

## Ethics and guidelines of scraping

Before doing any web scraping, it is important to understand what you are doing technically. If you use this information irresponsibly, you could potentially cause a denial\-of\-service, incur bandwidth costs to yourself or the website provider, overload log files, or otherwise stress computing resources. If you are unsure of the repercussions of your actions, do not perform any scraping without consulting a knowledgable person. You are responsible for the actions you take including any cost or repercussion that comes along with it.

When doing any scraping or crawling, you should be considerate of the server owners and use good rate limiting, prevent overloading a single site, and use reasonable settings and limits.

It is important to understand that some sites have terms of service that do not allow scraping. While you might not face legal problems, they could ban your account if you have one, block your IP address, or otherwise revoke your access to the website or service. Before scraping any site, find out if there are any rules or guidelines explicitly stated in the terms of service.

Also keep in mind that some websites do provide APIs. Check to see if an API is avaiable before scraping. If a website or service provides an API, you should use that. APIs are intended to be used programmatically and are also much more efficient.

## Prerequisites

*   [Go](https://golang.org) \- The Go programming language (tested with 1.6)

*   [goquery](https://github.com/PuerkitoBio/goquery) (for some examples) \- Go version of jQuery for DOM parsing

The only dependency, other than Go itself, is the goquery package. Goquery is not needed for every example, as the majority of examples rely only on the standard library. To install the **goquery** dependency, use **go get**:

```
go get github.com/PuerkitoBio/goquery
```

If you have issues with your $GOPATH when using **go get**, be sure to read up about [Workspaces](https://golang.org/doc/code.html#Workspaces) and [the GOPATH environment variable](https://golang.org/doc/code.html#GOPATH) and make sure you have a **GOPATH** set.

## Make an HTTP GET request

The first step to web scraping is being able to make an HTTP request. Let's look a very basic HTTP GET request and how to check the response code and view the content. Note the default timeout of an HTTP request using the default **transport** is forever.

```
// make_http_request.gopackage mainimport (    "io"    "log"    "net/http"    "os")func main() {    // Make HTTP GET request    response, err := http.Get("https://www.devdungeon.com/")    if err != nil {        log.Fatal(err)    }    defer response.Body.Close()    // Copy data from the response to standard output    n, err := io.Copy(os.Stdout, response.Body)    if err != nil {        log.Fatal(err)    }    log.Println("Number of bytes copied to STDOUT:", n)}
```

## Make an HTTP GET request with timeout

When using **http.Get()** to make a request, it uses the default HTTP client with default settings. If you want to override the settings you need to create your own client and use that to make the request. This example demonstrates how to create an **http.Client** and use it to make a request.

```
// make_http_request_with_timeout.gopackage mainimport (    "io"    "log"    "net/http"    "os"    "time")func main() {    // Create HTTP client with timeout    client := &http.Client{        Timeout: 30 * time.Second,    }    // Make request    response, err := client.Get("https://www.devdungeon.com/")    if err != nil {        log.Fatal(err)    }    defer response.Body.Close()    // Copy data from the response to standard output    n, err := io.Copy(os.Stdout, response.Body)    if err != nil {        log.Fatal(err)    }    log.Println("Number of bytes copied to STDOUT:", n)}
```

## Set HTTP headers (Change user agent)

In the first example we saw how to use the default HTTP client to make a request. Then we saw how to create out own client so we could customize the settings, like the timeout. Similarly, the HTTP clients use a default **Request** type which we can also customize. This example will walk through creating a request and modifying the headers before sending.

I highly recommed being a good net citizen and providing a descriptive user agent with a string that is easily parsable with a regular expression and contains a link to a website or GitHub repo so a network admin can learn about what the bot is and rate limit or block your bot if it causes problems.

```
# Example of a decent bot user agentMyScraperBot v1.0 https://www.github.com/username/MyNanoBot - This bot does x, y, z
```

Another reason to change your user agent might be to impersonate a different user agent. The default Go user agent may get blocked and you might have to impersonate a Firefox browser. It can also be useful for testing applications to see how they behave when various mobile and desktop user agents are presented.

This example will demonstrate how to change the HTTP headers before sending your request. To set your user agent, you will need to add/override the User\-Agent header. Note you can change any header this way, including your cookies, if you wanted to manually manage them. We'll talk more about the cookies later. This only requires the standard library.

```
// http_request_change_headers.gopackage mainimport (    "io"    "log"    "net/http"    "os"    "time")func main() {    // Create HTTP client with timeout    client := &http.Client{        Timeout: 30 * time.Second,    }    // Create and modify HTTP request before sending    request, err := http.NewRequest("GET", "https://www.devdungeon.com", nil)    if err != nil {        log.Fatal(err)    }    request.Header.Set("User-Agent", "Not Firefox")    // Make request    response, err := client.Do(request)    if err != nil {        log.Fatal(err)    }    defer response.Body.Close()    // Copy data from the response to standard output    _, err = io.Copy(os.Stdout, response.Body)    if err != nil {        log.Fatal(err)    }}
```

## Download a URL

You may want to simply download the contents of a page and store it for offline review at a later date, or download a binary file after determining what URL contains the file you want. This example demonstrates how to make an HTTP request and stream the contents to a file. This only requires the standard library.

```
// download_url.gopackage mainimport (    "io"    "log"    "net/http"    "os")func main() {    // Make request    response, err := http.Get("https://www.devdungeon.com/archive")    if err != nil {        log.Fatal(err)    }    defer response.Body.Close()    // Create output file    outFile, err := os.Create("output.html")    if err != nil {        log.Fatal(err)    }    defer outFile.Close()    // Copy data from HTTP response to file    _, err = io.Copy(outFile, response.Body)    if err != nil {        log.Fatal(err)    }}
```

## Use substring matching to find page title

Probably the simplest way to search for something in an HTML document is to do a regular substring match. You will need to first convert the response in to a string and then use the **strings** package in the standard library to do substring searches. This is not my preferred way of searching for things, but it can be viable depending on what you are looking for. It is definitely worth knowing and understanding this technique in case you want to use it. Thanks [xiegeo](https://www.reddit.com/r/golang/comments/86xrek/web_scraping_with_go/dw9i8yb/) for reminding me to include this section.

Next we will look at using regular expressions, which are even more powerful than simple substring matches. After that, we'll look at using the **goquery** package to parse the HTML DOM and look for data in a structured way using jQuery like syntax.

```
// substring_matching.gopackage mainimport (    "fmt"    "io/ioutil"    "log"    "net/http"    "os"    "strings")func main() {    // Make HTTP GET request    response, err := http.Get("https://www.devdungeon.com/")    if err != nil {        log.Fatal(err)    }    defer response.Body.Close()    // Get the response body as a string    dataInBytes, err := ioutil.ReadAll(response.Body)    pageContent := string(dataInBytes)    // Find a substr    titleStartIndex := strings.Index(pageContent, "<title>")    if titleStartIndex == -1 {        fmt.Println("No title element found")        os.Exit(0)    }    // The start index of the title is the index of the first    // character, the < symbol. We don't want to include    // <title> as part of the final value, so let's offset    // the index by the number of characers in <title>    titleStartIndex += 7    // Find the index of the closing tag    titleEndIndex := strings.Index(pageContent, "</title>")    if titleEndIndex == -1 {        fmt.Println("No closing tag for title found.")        os.Exit(0)    }    // (Optional)    // Copy the substring in to a separate variable so the    // variables with the full document data can be garbage collected    pageTitle := []byte(pageContent[titleStartIndex:titleEndIndex])    // Print out the result    fmt.Printf("Page title: %s\n", pageTitle)}
```

## Use regular expressions to find HTML comments

Regular expressions are a powerful way of searching for text patterns. I am providing one example of using regular expressions for reference, but I do not recommend using this method unless you have no other choice. In the next examples, I will look at using goquery, an easier way of finding data in a structured HTML document.

```
// find_html_comments_with_regex.gopackage mainimport (    "fmt"    "io/ioutil"    "log"    "net/http"    "regexp")func main() {    // Make HTTP request    response, err := http.Get("https://www.devdungeon.com")    if err != nil {        log.Fatal(err)    }    defer response.Body.Close()    // Read response data in to memory    body, err := ioutil.ReadAll(response.Body)    if err != nil {        log.Fatal("Error reading HTTP body. ", err)    }    // Create a regular expression to find comments    re := regexp.MustCompile("<!--(.|\n)*?-->")    comments := re.FindAllString(string(body), -1)    if comments == nil {        fmt.Println("No matches.")    } else {        for _, comment := range comments {            fmt.Println(comment)        }    }}
```

## Use goquery to find all links on a page

This example will make use of the goquery package to parse the HTML DOM and let us search for specific elements in a convenient, jQuery\-like way. We perform the HTTP request like normal, and then create a goquery document using the response. With the goquery document object, we can call **Find()** and process each element found. In this case, we will search for **a** elements, or links.

I am only scratching the surface of what [goquery](https://github.com/PuerkitoBio/goquery) can do. Here is an example of what it can do:

```
// Example of a more complex goquery to find an element in the DOM// https://github.com/PuerkitoBio/goquerydocument.Find(".sidebar-reviews article .content-block")
```

This is a full working example of how to use goquery to find all the links on a page and print them out.

```
// find_links_in_page.gopackage mainimport (    "fmt"    "log"    "net/http"    "github.com/PuerkitoBio/goquery")// This will get called for each HTML element foundfunc processElement(index int, element *goquery.Selection) {    // See if the href attribute exists on the element    href, exists := element.Attr("href")    if exists {        fmt.Println(href)    }}func main() {    // Make HTTP request    response, err := http.Get("https://www.devdungeon.com")    if err != nil {        log.Fatal(err)    }    defer response.Body.Close()    // Create a goquery document from the HTTP response    document, err := goquery.NewDocumentFromReader(response.Body)    if err != nil {        log.Fatal("Error loading HTTP response body. ", err)    }    // Find all links and process them with the function    // defined earlier    document.Find("a").Each(processElement)}
```

## Parse URLs

In the previous example we looked at finding all the links on a page. A common task after that is to examine the URL and determine if it is a relative URL that leads somewhere on the same site, or a URL that leads off\-site somewhere. You can use the string functions to search and parsae the URL manually, but there is a better way!

The Go standard library provides a convenient **URL** type that can handle all of the URL string parsing for us. Let it handle the heavy lifting with string parsing, and just get the hostname, port, query, requestURI, using the predefined functions. Read more about the **url** package and the **url.URL** type at [https://golang.org/pkg/net/url/](https://golang.org/pkg/net/url/).

```
// parse_urls.gopackage mainimport (    "fmt"    "log"    "net/url")func main() {    // Parse a complex URL    complexUrl := "https://www.example.com/path/to/?query=123&this=that#fragment"    parsedUrl, err := url.Parse(complexUrl)    if err != nil {        log.Fatal(err)    }    // Print out URL pieces    fmt.Println("Scheme: " + parsedUrl.Scheme)    fmt.Println("Host: " + parsedUrl.Host)    fmt.Println("Path: " + parsedUrl.Path)    fmt.Println("Query string: " + parsedUrl.RawQuery)    fmt.Println("Fragment: " + parsedUrl.Fragment)    // Get the query key/values as a map    fmt.Println("\nQuery values:")    queryMap := parsedUrl.Query()    fmt.Println(queryMap)    // Craft a new URL from scratch    var customURL url.URL    customURL.Scheme = "https"    customURL.Host = "google.com"    newQueryValues := customURL.Query()    newQueryValues.Set("key1", "value1")    newQueryValues.Set("key2", "value2")    customURL.Fragment = "bookmarkLink"    customURL.RawQuery = newQueryValues.Encode()    fmt.Println("\nCustom URL:")    fmt.Println(customURL.String())}
```

## Use goquery to find all images on a page

We can also leverage the **goquery** package to search for other elements. This is another simple example similar to finding the links on a page. This example will show how to search for images on a page and list the URLs. This example is written slightly different, to demonstrate how to create an anonymous function to handle the processing instead of a named function.

```
// find_images_in_page.gopackage mainimport (    "fmt"    "log"    "net/http"    "github.com/PuerkitoBio/goquery")func main() {    // Make HTTP request    response, err := http.Get("https://www.devdungeon.com")    if err != nil {        log.Fatal(err)    }    defer response.Body.Close()    // Create a goquery document from the HTTP response    document, err := goquery.NewDocumentFromReader(response.Body)    if err != nil {        log.Fatal("Error loading HTTP response body. ", err)    }    // Find and print image URLs    document.Find("img").Each(func(index int, element *goquery.Selection) {        imgSrc, exists := element.Attr("src")        if exists {            fmt.Println(imgSrc)        }    })}
```

## Make HTTP POST request with data

Making a POST request is similar to a GET request. In fact, it is as simple as changing the word "GET" to "POST" in the request. However, a POST request is often accompanied with a payload. This could be a binary file or a URL encoded form. This example will demonstrate how to make a POST request with URL encoded form data and how to post a file like when uploading a file..

```
// http_post_with_payload.gopackage mainimport (    "log"    "net/http"    "net/url")func main() {    response, err := http.PostForm(        "http://example.com/form",        url.Values{            "username": {"MyUsername"},            "password": {"123"},        },    )    if err != nil {        log.Fatal(err)    }    defer response.Body.Close()    log.Println(response.Header) // Print the response headers    // To upload a file, use Post instead of PostForm, provide    // a content type like application/json or application/octet-stream,    // and then provide the an io.Reader with the data    // http.Post("http://example.com/upload", "image/jpeg", &buff)}
```

## Make HTTP request with cookie

Since cookies are simply HTTP headers, you can manually set and manage cookies yourself by checking and setting the header values as needed.

Go offers a better way of managing cookies with a **Cookie** type that is used by the **Request** and **Response** type. You can see the source code for the **Cookie** at [https://golang.org/src/net/http/cookie.go](https://golang.org/src/net/http/cookie.go)

These are some of the cookie functions available on the **Request** and **Response** types:

```
// Cookie functions for Request // https://golang.org/pkg/net/http/#RequestRequest.AddCookie()  // Add cookie to requestRequest.Cookie()     // Get specific cookieRequest.Cookies()    // Get all cookies// Cookie functions for Response// https://golang.org/pkg/net/http/#ResponseResponse.Cookies()   // Get all cookies
```

Alternatively, you could use a library that is not part of the standard library like the [sessions package provided by Gorilla](https://github.com/gorilla/sessions), but that will not be covered here.

There is also a **cookiejar** type. It is essentially a collection of cookies separated by URL. You can read more about at [https://golang.org/pkg/net/http/cookiejar/](https://golang.org/pkg/net/http/cookiejar/). It is useful if you need to manage cookies for multiple sites.

```
// http_request_with_cookie.gopackage mainimport (    "fmt"    "log"    "net/http")func main() {    request, err := http.NewRequest("GET", "https://www.devdungeon.com", nil)    if err != nil {        log.Fatal(err)    }    // Create a new cookie with the only required fields    myCookie := &http.Cookie{        Name:  "cookieKey1",        Value: "value1",    }    // Add the cookie to your request    request.AddCookie(myCookie)    // Ask the request to tell us about itself,    // just to confirm the cookie attached properly    fmt.Println(request.Cookies())    fmt.Println(request.Header)    // Do something with the request    // client := &http.Client{}    // client.Do(request)}
```

## Log in to a website

Logging in to a site is relatively simple conceptually. You make an HTTP POST to a specific URL containing your username and password, and it returns a cookie, which is simply an HTTP header, that contains a unique key that matches your session on the server. Most websites work the same in this regard, although custom authentication mechanisms, CAPTCHAs, two\-factor authentication, and other security measures complicate this process.

Logging in to a site is going to have to be tailored specifically to your target website. You will have to reverse engineer the authentication process from the site. Many websites use a simple form\-based login system. Inside a browser like Chrome or Firefox, you can right click on one of the form fields and choose "inspect". This will allow you to see how the form is constructed, what the target action url is, and how the form fields are named in order to recreate the request programmatically.

You can inspect the form in the source of the HTML, or you can monitor the network traffic itself. The brwoser extensions will let you see the POST requests going on behind the scene on a website, but you could use other tools as well like jsfiddler, burp suite, Zed Attack Proxy (ZAP), or any other man\-in\-the\-middle proxying tool.

Typically, you will need to get the URL in the **action** attribute of the **form**, and the **name** attribute of the of the username and password **input** fields. Once you have that information, you can make the POST request to the URL, and then store the session cookie the server provides in its response. You will need to pass the session cookie with any subsequent requests you make to the server.

Because every website has it's own mechanism for authentication, I am only covering this at the conceptual level and not providing a code example.

## Web crawling

Crawling is simply an extension of scraping. We already looked at how to [find all links on a page](#find_all_links_on_page), and how to [parse URLs](#parse_urls), which are the important steps. You want to find all the links on a page, parse the url, decide if you want to follow it, and then make a request to the new url, repeating the process.

After parsing a URL, you can determine whether it belongs to the same site you are already on, or leads to another website. You can also look for a file extension at the end of the URL for clues about what it leads to.

You can crawl in a breadth\-first or a depth\-first manner. One depth\-first approach would be to crawl only URLs from the same website before crawling the next website in the list. A breadth\-first approach would be to prioritize links that lead to websites you have never seen before.

For a code example of a web crawler, check out the DevDungeon Web Genome project in the next section.

## DevDungeon Project: Web Genome

Web Genome is a breadth first web crawler that stores HTTP headers in a MongoDB database with a web interface all written in Go. Read more on the [Web Genome project page](https://www.devdungeon.com/content/web-genome) and browse the source code at [https://github.com/DevDungeon/WebGenome](https://github.com/DevDungeon/WebGenome).

Visit the Web Genome website at [http://www.webgeno.me](http://www.webgeno.me).

## Conclusion

With this reference code, you should be able to perform basic web scraping tasks to suit your needs. There are also more features in the goquery library that I have not covered. Refer to the official repository at [https://github.com/PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery) for the latest information.
