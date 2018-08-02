# Example 1:

```go
 package main

 import (
         "fmt"
         "net"
 )

 func whois(domainName, server string) string {
         conn, err := net.Dial("tcp", server+":43")

         if err != nil {
                 fmt.Println("Error")
         }

         defer conn.Close()

         conn.Write([]byte(domainName + "\r\n"))

         buf := make([]byte, 1024)

         result := []byte{}

         for {
                 numBytes, err := conn.Read(buf)
                 sbuf := buf[0:numBytes]
                 result = append(result, sbuf...)
                 if err != nil {
                         break
                 }
         }

         return string(result)
 }

 func main() {
         result := whois("socketloop.com", "com.whois-servers.net")
         fmt.Println(result)
 }
 ```
 
## Example 2:

```go
 package main

 import (
         "fmt"
         "github.com/likexian/whois-go"
 )

 func main() {
         result, err := whois.Whois("socketloop.com")

         if err != nil {
                 fmt.Println(err)
         }

         fmt.Println(result)
 }
 ```
 
