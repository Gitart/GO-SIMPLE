# How do I limit the rate of requests in Colly to avoid being blocked?

In web scraping, it's important to respect the server's resources and adhere to the website's terms of service. One way to do this is by limiting the rate of your requests to avoid overwhelming the server and reduce the chances of being blocked. Colly, a popular web scraping framework for Go (Golang), provides options to adjust the rate of your requests.

Here's how you can limit the rate of requests in Colly:

1. **Set the `Limit` rule**: You can limit the number of simultaneous requests to a domain and also introduce a delay between requests.

Copy

```go
package main

import (
    "fmt"
    "github.com/gocolly/colly/v2"
    "github.com/gocolly/colly/v2/queue"
    "log"
    "time"
)

func main() {
    // Instantiate the collector
    c := colly.NewCollector(
        // You may set other options as needed
    )

    // Create a request queue with a concurrency of 2
    // This means Colly will only make 2 simultaneous requests to the domain
    q, _ := queue.New(
        2, // Number of consumer threads
        &queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
    )

    // Set the delay between requests to 1 second
    // This means after each request Colly will wait for 1 second before making another request
    c.Limit(&colly.LimitRule{
        DomainGlob:  "*",
        Delay:       1 * time.Second,
        RandomDelay: 1 * time.Second, // Add some randomness to the delay
    })

    // On every a element which has href attribute call callback
    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        link := e.Attr("href")
        // Add link found on the page to the queue
        err := q.AddURL(link)
        if err != nil {
            log.Printf("Error adding URL to the queue: %s", err)
        }
    })

    // Start scraping on the website
    err := q.AddURL("http://example.com")
    if err != nil {
        log.Printf("Error adding URL to the queue: %s", err)
    }

    // Consume URLs
    q.Run(c)

    fmt.Println("Scraping finished.")
}
```

In the example above, the `LimitRule` is set with a `Delay` of 1 second between requests and a `RandomDelay` which adds up to 1 second of random delay to the `Delay` time. This helps to make the scraping process appear more human-like.

Remember, it's important to read the website's `robots.txt` file and comply with it, as well as with the website's terms of service.

Lastly, it's also a good practice to identify yourself by setting a custom `User-Agent` and providing contact information through the `From` field in case the website owner needs to reach you:

Copy

```go
c.UserAgent = "YourCustomUserAgent/1.0"
c.From = "your_email@example.com"
```

By setting a `User-Agent`, you let the server know which client is making the request, and providing a contact email is a courteous way to enable communication in case your scraping activities cause any issues.

# How do I troubleshoot common issues when using Colly?

Colly is a popular scraping framework for Golang, designed to simplify the process of extracting data from websites. When using Colly, you might encounter various issues, such as unexpected errors, blocked requests, or incomplete data extraction. Here are some common troubleshooting steps to help you address these issues:

### 1. Debugging Output

Colly has built-in debugging capabilities. You can enable the debugger to get detailed information about the requests and responses:

Copy

```go
// Enable the debugger by attaching the default logger
c := colly.NewCollector(
    colly.Debugger(&debug.LogDebugger{}),
)
```

This will output detailed logs to the console, which can help you understand what's happening under the hood.

### 2. Error Handling

Always ensure you have error handlers in place to catch and diagnose any issues that arise:

Copy

```go
c.OnError(func(r *colly.Response, err error) {
    fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
})
```

### 3. Check for Blocked Requests

Websites might block your requests if they detect unusual behavior (e.g., too many requests from the same IP in a short period). To troubleshoot this:

* Slow down your requests using `Limit` rules.
* Rotate user agents.
* Use proxy servers.

Copy

```go
// Set up rate limit
err := c.Limit(&colly.LimitRule{
    DomainGlob:  "*",
    Parallelism: 1,
    Delay:       5 * time.Second,
})

// Rotate User Agents
c.UserAgent = "AnotherUserAgentString"

// Use a Proxy
c.SetProxy("http://myproxy.com:3128")
```

### 4. Inspect the Response

If you're not getting the data you expect, inspect the raw response:

Copy

```go
c.OnResponse(func(r *colly.Response) {
    fmt.Println("Visited", r.Request.URL)
    fmt.Println("Response", string(r.Body))
})
```

### 5. Check for JavaScript

If the content you're scraping is loaded via JavaScript, Colly won't be able to see it since it doesn't process JavaScript. You can:

* Look for API calls in the browser's network tab and scrape the API directly.
* Use a tool like Chromedp or go-rod to render JavaScript before scraping.

### 6. Element Selectors

Ensure your selectors are correct. If you're not getting the expected elements:

* Double-check your selectors.
* Make sure the elements exist in the raw HTML response.
* Use the browser's developer tools to validate your selectors.

### 7. Handling Redirects

Colly follows redirects by default. If you want to control this behavior, you can disable automatic redirects:

Copy

```go
c.WithTransport(&http.Transport{
    DisableKeepAlives: true,
    Proxy: http.ProxyFromEnvironment,
    DialContext: (&net.Dialer{
        Timeout:   30 * time.Second,
        KeepAlive: 30 * time.Second,
        DualStack: true,
    }).DialContext,
    ForceAttemptHTTP2:     true,
    MaxIdleConns:          100,
    IdleConnTimeout:       90 * time.Second,
    TLSHandshakeTimeout:   10 * time.Second,
    ExpectContinueTimeout: 1 * time.Second,
})
```

### 8. Cookies and Sessions

If the site uses cookies or sessions to manage state, make sure you're preserving them between requests:

Copy

```go
// This will get called for each request
c.OnRequest(func(r *colly.Request) {
    // Set cookies
    r.Headers.Set("Cookie", "name=value")
})
```

### 9. Captchas and JavaScript Challenges

Some sites have captchas or JavaScript challenges to block scrapers. Handling these requires more advanced techniques like captcha solving services or headless browsers.

### 10. Check for Legal Compliance

Ensure that your scraping activities comply with the website's `robots.txt` file, terms of service, and legal regulations.



### Conclusion

When troubleshooting Colly, the key is to gather as much information as possible about the issue you're facing. Use the debugging techniques mentioned above, inspect the responses thoroughly, and make sure you're adhering to the site's scraping policies. If you're still stuck, consider reaching out to the community or checking the official Colly documentation for more insights.





# Rate limit

```go
package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

func main() {
	url := "https://httpbin.org/delay/2"

	// Instantiate default collector
	c := colly.NewCollector(
		// Turn on asynchronous requests
		colly.Async(true),
		// Attach a debugger to the collector
		colly.Debugger(&debug.LogDebugger{}),
	)

	// Limit the number of threads started by colly to two
	// when visiting links which domains' matches "*httpbin.*" glob
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*httpbin.*",
		Parallelism: 2,
		//Delay:      5 * time.Second,
	})

	// Start scraping in five threads on https://httpbin.org/delay/2
	for i := 0; i < 5; i++ {
		c.Visit(fmt.Sprintf("%s?n=%d", url, i))
	}
	// Wait until threads are finished
	c.Wait()
}
```
