## Simple Web Server in Go to log Request Headers

What headers does my proxy add? I’ve been experimenting with Vulcand over the last weeks. A great proxy server that uses etcd directly for its configuration. While I was creating a setup with some advance proxy configuration like HTTPS, I was curious whether Vulcand correctly set the request headers like X-Forwarded-For and X-Forwarded-Proto.

Go Docker Go In order to figure this out I’ve created a simple Go application that creates a HTML page displaying the original request headers. Off course there are other ways to achieve this, but this was a good excuse for creating a Go application and running it in a container!

Go application The application is very simple, I tell Go to start listing on port 8080 and assign a function as the handler for all requests. This handler function takes the request headers from the original request, sorts them and prints them in the reponse. This way the user will see the request headers in the browser! One could access this server directly or through the proxy in front of it and see the difference in request headers.

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

Run the server on Docker We want to run this server on Docker and register it as a backend with our proxy server. I’m using vulcand, but this setup can be used with any proxy. To create a Docker image we need to add a Dockerfile to the same folder that contains the webserver.go file, listed above.


### EXPOSE 8080
This simple Dockerfile tells Docker to use a base image that directly compiles and installs your Go application.
For more information on how the Golang images work, have a look at: https://registry.hub.docker.com/u/library/golang/

To build & run the Docker image, use the following commands:

```
docker build -t my-go-webserver .
docker run -d -P --name proxy_this_webserver my-go-webserver
```
