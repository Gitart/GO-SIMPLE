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
