# Read gzipped http response



There are times when we want to grab a website content for parsing(crawling) and found out that the content is gzipped.   
Normally, to deal with gzipped HTML reply, you can use Exec package to execute curl from command line   
and pipe the gzipped content to gunzip, such as this :  

```
 curl -H "Accept-Encoding: gzip" http://www.thestar.com.my | gunzip
 ```
 
Another way to process gzipped http response can be done in Golang as well. The following codes  
will demonstrate how to get same result as the curl command via Golang.   


```golang
 package main

 import (
         "compress/gzip"
         "fmt"
         "io"
         "net/http"
         "os"
 )

 func main() {
         client := new(http.Client)

         request, err := http.NewRequest("Get", " http://www.thestar.com.my", nil)

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         request.Header.Add("Accept-Encoding", "gzip")

         response, err := client.Do(request)
         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }
         defer response.Body.Close()

         // Check that the server actual sent compressed data
         var reader io.ReadCloser
         switch response.Header.Get("Content-Encoding") {
         case "gzip":
                 reader, err = gzip.NewReader(response.Body)
                 if err != nil {
                         fmt.Println(err)
                         os.Exit(1)
                 }
                 defer reader.Close()
         default:
                 reader = response.Body
         }

         // to standard output
         _, err = io.Copy(os.Stdout, reader)

         // see https://www.socketloop.com/tutorials/golang-saving-and-reading-file-with-gob
         // on how to save to file

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

 }
 ```
