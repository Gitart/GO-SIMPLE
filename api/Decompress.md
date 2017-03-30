# Decompress zlib file example

Continuation from previous tutorial on how to zlib compress a file. In this part, we will learn how  
to decompress a file that was compressed with zlib compression algorithm.
  

Here are the codes :

```golang
 package main

 import (
         "compress/zlib"
         "flag"
         "fmt"
         "io"
         "os"
         "strings"
 )

 func main() {

         flag.Parse() // get the arguments from command line

         filename := flag.Arg(0)

         if filename == "" {
                 fmt.Println("Usage : unzlib sourcefile.zlib")
                 os.Exit(1)
         }

         zlibfile, err := os.Open(filename)

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         reader, err := zlib.NewReader(zlibfile)
         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }
         defer reader.Close()

         newfilename := strings.TrimSuffix(filename, ".zlib")

         writer, err := os.Create(newfilename)

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         defer writer.Close()

         if _, err = io.Copy(writer, reader); err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         fmt.Println("Decompressed to ", newfilename)

 }
 ```
 
Sample output :


```
./unzlib uncompressed.txt.zlib
Decompressed to uncompressed.txt
./unzlib
Usage : unzlib sourcefile.zlib
```
