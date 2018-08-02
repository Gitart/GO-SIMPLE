# Read file
This is the most basic way of how to read a file into buffer and display its content chunk by chunk. This example read plain text file, if you are reading a binary file... change fmt.Println(string(buffer[:n])) to fmt.Println(buffer[:n]) (without the string).
For now, this is the most basic example of reading a file in Go

```go
 package main

  import (
      "fmt"
      "io"
      "os"
  )

  func main() {

      file, err := os.Open("sometextfile.txt")

      if err != nil {
          fmt.Println(err)
          return
      }

     defer file.Close()


     // create a buffer to keep chunks that are read

     buffer := make([]byte, 1024)
     for {
         // read a chunk
         n, err := file.Read(buffer)
         if err != nil && err != io.EOF { panic(err) }
         if n == 0 { break }

         // out the file content
         fmt.Println(string(buffer[:n]))

     }
 }
 ```
 
