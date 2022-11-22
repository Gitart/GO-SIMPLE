Go Fragments: a collection of annotated Go programs examples
gofragments.net logo	
FUNDAMENTALS
CONCURRENCY
NET AND WEB
SITE INFO
Fragment Description:



The server program issues Google search requests and demonstrates the use of the go.net Context API. 
It serves on port 8080. 
The /search endpoint accepts these query params: 
q=the Google search query timeout=a timeout for the request, in time.Duration format For example, http://localhost:8080/search?q=golang&timeout=1s serves the first few Google search results for 'golang' or a 'deadline exceeded' error if the timeout expires. 
see article here: 
http://blog.golang.org/context 
serverUsingCONTEXT

Last update, on 2015, Fri 9 Oct, 16:15:40

/* ... <== see fragment description ... */

package main

import (
    "go.net/context"
    "go.net/google"
    "go.net/userip"
    "html/template"
    "log"
    "net/http"
    "time"
)

func main() {
    http.HandleFunc("/search", handleSearch)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// handleSearch handles URLs like /search?q=golang&timeout=1s by forwarding
// the
// query to google.Search. If the query param includes timeout, the search is
// canceled after that duration elapses.
func handleSearch(w http.ResponseWriter, req *http.Request) {
    // ctx is the Context for this handler. Calling cancel closes the
    // ctx.Done channel, which is the cancellation signal for requests
    // started by this handler.
    var (
        ctx    context.Context
        cancel context.CancelFunc
    )
    timeout, err := time.ParseDuration(req.FormValue("timeout"))
    if err == nil {
        // The request has a timeout, so create a context that is
        // canceled automatically when the timeout expires.
        ctx, cancel = context.WithTimeout(context.Background(), timeout)
    } else {
        ctx, cancel = context.WithCancel(context.Background())
    }
    defer cancel() // Cancel ctx as soon as handleSearch returns.
    // Check the search query.
    query := req.FormValue("q")
    if query == "" {
        http.Error(w, "no query", http.StatusBadRequest)
        return
    }
    // Store the user IP in ctx for use by code in other packages.
    userIP, err := userip.FromRequest(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    ctx = userip.NewContext(ctx, userIP)
    // Run the Google search and print the results.
    start := time.Now()
    results, err := google.Search(ctx, query)
    elapsed := time.Since(start)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := resultsTemplate.Execute(w, struct {
        Results          google.Results
        Timeout, Elapsed time.Duration
    }{
        Results: results,
        Timeout: timeout,
        Elapsed: elapsed,
    }); err != nil {
        log.Print(err)
        return
    }
}

var resultsTemplate = template.Must(template.New("results").Parse(`
<html>
<head/>
<body>
  <ol>
  {{range .Results}}
    <li>{{.Title}} - <a href="{{.URL}}">{{.URL}}</a></li>
  {{end}}
  </ol>
  <p>{{len .Results}}  results in {{.Elapsed}}; timeout {{.Timeout}}</p>
</body>
</html>
`))

/* More about Context:
 import "golang.org/x/net/context"
 Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.
Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context. The chain of function calls between must propagate the Context, optionally replacing it with a modified copy created using WithDeadline, WithTimeout, WithCancel, or WithValue.
see details here: http://godoc.org/golang.org/x/net/context
*/

Comments



ABOUT GO FRAGMENTS

Go fragments, gather snippets of Go code that I found interesting and helpful.

I really enjoyed 'Go' these last 2 years (both the programming language, the Open Source project and its community).

I hope this Go fragments collection will be useful for you too, and also pleasant to visit and browse.

More about this site
TODOS

gofragments.net logo
gofragments.net logo
CONTACT ME

 contact me
RSS FEEDS

tockell consulting logoHome About Licensing and Policy
All material here is licensed under a - Creative Commons License - "gofragments.net" by pmjtoca is licensed under a Creative Commons Attribution 4.0 International License. Based on a work at http://gofragments.net.
