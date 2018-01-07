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
