## Setup API server or gateway with Caddy and http.ListenAndServe() function example

Writing this down for my own future references and hopefully it can be useful to you too. 
This is a record of my own attempt in setting up Caddy server with a connection to(proxy) another Golang application.
The Golang application acts as an API server that will pump out HTML and JavaScript codes to Caddy via http.ListenAndServe() function.
If you haven't configure your Caddy server yet. Please read https://caddyserver.com/docs/getting-started first.
Below is my own Caddyfile configuration. You will need to replace example.com to your domain name or if you're 
testing on localhost. Replace https://api.example.com to localhost:


```
 https://api.example.com {
    root /opt/api
    gzip
    browse
    ext .html
    log access.log
    errors error.log

    # connect to external Golang program(API server) on port 9999 via http.ListenAndServe()  
    proxy / :9999 {  

        # except some things for caddy to serve directly
      except /assets /files /robots.txt /favicon.ico
    }


    header / {
        # Allow CORS, while removing the Server field
        Access-Control-Allow-Origin *
        Access-Control-Allow-Methods "GET, POST, OPTIONS"
        -Server

        # Enable HTTP Strict Transport Security (HSTS) to force clients to always
        # connect via HTTPS (do not use if only testing)
        Strict-Transport-Security "max-age=31536000;"
        # Enable cross-site filter (XSS) and tell browser to block detected attacks
        X-XSS-Protection "1; mode=block"
        # Prevent some browsers from MIME-sniffing a response away from the declared Content-Type
        X-Content-Type-Options "nosniff"
        # Disallow the site to be rendered within a frame (clickjacking protection)
        X-Frame-Options "DENY"
     }
 }
```

```
 http://api.example.com {
    # always HTTP(S)
    redir https://api.example.com
 }
 ```
 
In the Caddyfile configuration above, the most important block is the proxy configuration that tells Caddy to pay attention to port 9999.
Next, configure the Golang program(API server) that will pump out the HTML and JavaScript codes, serve other API stuff and etc to port 9999 to be picked up by Caddy.

api-server.go:

```golang
 package main

 import (
 	"net/http"
 )

 func APIPumpCodes(w http.ResponseWriter, r *http.Request) {
 	html := `<html>
 <head>
 <title>Serving HTML and JavaScript codes via port 9999</title>
 </head>

 <body>
 <p>Serving HTML and JavaScript codes via port 9999</p>
 </body>

 </html>`

 	w.Write([]byte(html))
 }
```

```golang
 func main() {
 	mux := http.NewServeMux()
 	mux.HandleFunc("/", APIPumpCodes)

 	http.ListenAndServe(":9999", mux)
 } 
```

Build and execute the api-server.go program. Use & command to let it run in the background.

```
>./api-server &
```

and fire up the Caddy server

```
>./caddy &
```

If everything goes smoothly, you should see the following output on your terminal:

```
Activating privacy features... done.
http://api.example.com
https://api.example.com
Pointing your browser to https://api.example.com should give you this message:

Serving HTML codes via port 9999
Alright! You have a Caddy server communicating successfully with your Golang program via port 9999.
```


NOTE: If you get the 502 Bad Gateway error message, it is because the port number mismatch or the
Golang program(api-server.go) is not running.
