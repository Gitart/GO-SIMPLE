// https://play.golang.org/p/I17ZSM6LQb


package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func report(r *http.Request) {
	r.Host = "stackoverflow.com"
	r.URL.Host = r.Host
	r.URL.Scheme = "http"
}

func main() {
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "stackoverflow.com",
	})
	proxy.Director = report
	http.Handle("/", proxy)
	http.ListenAndServe(":8080", nil)
}
