package main

 import (
         "fmt"
         "io"
         "net/http"
         "os"
 )

 func main() {
         response, err := http.Get("https://www.socketloop.com")

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         defer response.Body.Close()

         htmlfile, err := os.Create("file.html")

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         defer htmlfile.Close()

         // save response body into a file
         io.Copy(htmlfile, response.Body)

         fmt.Println("HTML data saved into file.html")
 }
