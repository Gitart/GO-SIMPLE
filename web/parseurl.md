# Parse URL

```golang
 package main

 import (
         "fmt"
         "net/url"
         "strings"
 )

 func main() {

         rawURL := "http://username:password@searchengine.com:8080/testpath/?q=socketloop.com#TestFunc"

         fmt.Println("URL : ", rawURL)

         // Parse the URL and ensure there are no errors.
         url, err := url.Parse(rawURL)
         if err != nil {
                 panic(err)
         }

         // see http://golang.org/pkg/net/url/#URL
         // scheme://[userinfo@]host/path[?query][#fragment]

         // get the Scheme
         fmt.Println("Scheme : ", url.Scheme)

         // get the User information
         fmt.Println("Username : ", url.User.Username())

         password, set := url.User.Password()
         fmt.Println("Password : ", password)
         fmt.Println("Password set : ", set)

         // get the Host
         fmt.Println("Raw host : ", url.Host)

         // to get the Port number, split the Host
         hostport := strings.Split(url.Host, ":")
         host := hostport[0]
         port := hostport[1]

         fmt.Println("Host : ", host)
         fmt.Println("Port : ", port)

         // get the Path
         fmt.Println("Path : ", url.Path)

         // get the RawQuery
         fmt.Println("Raw Query ", url.RawQuery)

         // get the fragment
         fmt.Println("Fragment : ", url.Fragment)

 }
 ```
 

## Get URI segments by number and assign as variable example

Coming from PHP and CodeIgniter(framework for PHP) background. One thing that I missed is the URI helper
's $this->uri->segment(n) where it permits me to retrieve the segment by number and assign the value to
string variable. For example :

URL : http://example.com/index.php/news/local/metro/crimeisup

The segment numbers would be this:
news
local
metro
crimeisup


news/local/metro/crime_is_up is known as URL Path.

In this tutorial, we will learn how to get the URI segments by number and assign the segment as variable with Golang.
See codes below :


```golang
 package main

 import (
         "fmt"
         "net/url"
         "strings"
 )

 func main() {
         rawURL := "http://example.com/index.php/news/local/metro/crime_is_up"

         fmt.Println("URL : ", rawURL)

         url, err := url.Parse(rawURL)

         if err != nil {
                 panic(err)
         }

         path := url.Path
         uriSegments := strings.Split(path, "/")
         fmt.Println(uriSegments) // count starts from 1
         var metro = uriSegments[3] // assign to variable
         fmt.Println("[segment 3] : ", metro)
         fmt.Println("[segment 4] : ", uriSegments[4])
 }
 ```
 
Output :

```
URL : http://example.com/index.php/news/local/metro/crimeisup
[ index.php news local metro crimeisup]
[segment 3] : local
[segment 4] : metro
```

Hang on..this is processing a static URL string. How about getting the URL from browser ?

No worries, here it is :

```golang
 package main

 import (
         "net/http"
         "strings"
 )

 func SayHello(w http.ResponseWriter, r *http.Request) {
         w.Write([]byte("Hello, World!"))
 }

 func GetURISegment(w http.ResponseWriter, r *http.Request) {
         uriSegments := strings.Split(r.URL.Path, "/")
         var metro = uriSegments[3]

         w.Write([]byte("[segment 3] : " + metro + "\r\n"))
         w.Write([]byte("[segment 4] : " + uriSegments[4]))
 }

 func main() {
         // http.Handler
         mux := http.NewServeMux()
         mux.HandleFunc("/", SayHello)
         mux.HandleFunc("/news/local/metro/crime_is_up", GetURISegment)

         http.ListenAndServe(":8080", mux)
 }
 ```
 
Run this modified code and point your browser URL to :

```
http://localhost:8080/news/local/metro/crime_is_up
```

and you will see the following output :

[segment 3] : metro
[segment 4] : crimeisup
