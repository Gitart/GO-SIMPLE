## Request

```golang
package main

import (
	"bytes"
	"fmt"
	"net/http"
	//"net/http/httputil"
)

func main() {
	var b bytes.Buffer

	r, err := http.NewRequest("POST", "http://example.com", &b)
	if err != nil {
		panic(err)
	}
	r.Header.Add("X-Custom", "Copy me!")

	rc, err := http.NewRequest("POST", r.URL.String(), &b)
	if err != nil {
		panic(err)
	}

	rc.Header = r.Header // note shallow copy

	fmt.Println("Headers", r.Header, rc.Header)
	
	
	// Adjust copy adjusts original 
	rc.Header.Add("X-Hello","World")

	fmt.Println("Headers", r.Header, rc.Header)
	
}
```


## Round Trip

```golang
package main

import (
    "fmt"
    "net/http"
    "sort"
)

func (c *Client) RoundTrip(action string, in, out Message) error {
    fmt.Println("****************************************************************")
    headerFunc := func(r *http.Request) {
        r.Header.Add("Content-Type", fmt.Sprintf("text/xml; charset=utf-8"))
        r.Header.Add("SOAPAction", fmt.Sprintf(action))
        r.Cookies()
    }
    return doRoundTrip(c, headerFunc, in, out)
}

func doRoundTrip(c *Client, setHeaders func(*http.Request), in, out Message) error {
    req := &Envelope{
        EnvelopeAttr: c.Envelope,
        NSAttr:       c.Namespace,
        Header:       c.Header,
        Body:         Body{Message: in},
    }

    if req.EnvelopeAttr == "" {
        req.EnvelopeAttr = "http://schemas.xmlsoap.org/soap/envelope/"
    }
    if req.NSAttr == "" {
        req.NSAttr = c.URL
    }
    var b bytes.Buffer
    err := xml.NewEncoder(&b).Encode(req)
    if err != nil {
        return err
    }
    cli := c.Config
    if cli == nil {
        cli = http.DefaultClient
    }
    r, err := http.NewRequest("POST", c.URL, &b)
    if err != nil {
        return err
    }
    setHeaders(r)    
    ```  
    
 ## Simple Web Server in Go to log Request Headers  
    
#What headers does my proxy add? I’ve been experimenting with Vulcand over the last weeks. A great proxy server that uses etcd directly for its configuration. While I was creating a setup with some advance proxy configuration like HTTPS, I was curious whether Vulcand correctly set the request headers like X-Forwarded-For and X-Forwarded-Proto.

#Go Docker Go In order to figure this out I’ve created a simple Go application that creates a HTML page displaying the original request headers. Off course there are other ways to achieve this, but this was a good excuse for creating a Go application and running it in a container!

#Go application The application is very simple, I tell Go to start listing on port 8080 and assign a function as the handler for all requests. This handler function takes the request headers from the original request, sorts them and prints them in the reponse. This way the user will see the request headers in the browser! One could access this server directly or through the proxy in front of it and see the difference in request headers.

```golang
package main

import (
    "fmt"
    "net/http"
    "sort"
)

func handler(w http.ResponseWriter, r *http.Request) {

    var keys []string
    for k := range r.Header {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    
    fmt.Fprintln(w, "<b>Request Headers:</b></br>", r.URL.Path[1:])
    for _, k := range keys {
        fmt.Fprintln(w, k, ":", r.Header[k], "</br>", r.URL.Path[1:])
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```
    
