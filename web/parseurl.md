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
 

