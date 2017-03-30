## Compress and decompress file with compress/flate example    
A quick and simple tutorial on how to compress a file with Golang's compress/flate package.   
Package flate implements the DEFLATE compressed data format, as described in RFC 1951.   
To compress a file, pipe and flush the data out with NewWriter() function :   

```golang
package main

 import (
         "compress/flate"
         "fmt"
         "io"
         "os"
 )

 func main() {
         inputFile, err := os.Open("file.txt")

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         defer inputFile.Close()

         outputFile, err := os.Create("file.txt.compressed")

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         defer outputFile.Close()

         flateWriter, err := flate.NewWriter(outputFile, flate.BestCompression)

         if err != nil {
                 fmt.Println("NewWriter error ", err)
                 os.Exit(1)
         }

         defer flateWriter.Close()
         io.Copy(flateWriter, inputFile)

         flateWriter.Flush()
 }
 ```
 
Sample test data :

```
>cat file.txt

This is a test file for Golang Compress/Flate
This is a test file for Golang Compress/Flate
This is a test file for Golang Compress/Flate
This is a test file for Golang Compress/Flate
This is a test file for Golang Compress/Flate
This is a test file for Golang Compress/Flate
>cat file.txt.compressed
```

��,V�D������T���"��ļt��܂���b}��ĒT�Q�#U5����
To decompress the file, read in the compressed data with NewReader() function before piping out the decompressed data to file


```golang
 package main

 import (
         "compress/flate"
         "fmt"
         "io"
         "os"
 )

 func main() {
         inputFile, err := os.Open("file.txt.compressed")

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         defer inputFile.Close()

         outputFile, err := os.Create("file.txt.decompressed")

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         defer outputFile.Close()

         flateReader := flate.NewReader(inputFile)

         defer flateReader.Close()
         io.Copy(outputFile, flateReader)

 }
 ```
