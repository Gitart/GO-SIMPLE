package main

 import (
         "fmt"
         "io/ioutil"
         "net/http"
         "os"
 )

 func main() {

         // http.Get() can handle gzipped data response
         // automagically

         resp, err := http.Get("https://golang.org")

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         defer resp.Body.Close()

         htmlData, err := ioutil.ReadAll(resp.Body)

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         fmt.Println(os.Stdout, string(htmlData))

 }
