// https://www.socketloop.com/tutorials/zip-compress-file-in-go

package main

 import (
         "archive/zip"
         "flag"
         "fmt"
         "io"
         "os"
 )

 // use RegisterCompressor, if you plan to use your own compression algorithm
 // other that the standard algorithm 8(deflate)
 // the init block is commentted out as it is not used in this example
 //func init() {
 // see https://pkware.cachefly.net/webdocs/casestudies/APPNOTE.TXT
 // section 4.4.5 for the available compression methods...which are not implemented by Golang....
 // only algorithm 0 and number 8(deflate) available

 //      var comp zip.Compressor
 //      zip.RegisterCompressor(999, comp)
 //}

 func main() {
         flag.Parse()

         filename := flag.Arg(0)

         if filename == "" {
                 fmt.Println("Usage go-zip sourcefile")
                 os.Exit(1)
         }

         fmt.Printf("Zipping file. \n")

         err := zipfile(filename)

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         fmt.Println("Zipped to ", filename+".zip")
 }

 func zipfile(filename string) error {

         newfile, err := os.Create(filename + ".zip")
         if err != nil {
                 return err
         }
         defer newfile.Close()

         zipit := zip.NewWriter(newfile)
         defer zipit.Close()

         zipfile, err := os.Open(filename)
         if err != nil {
                 return err
         }
         defer zipfile.Close()

         // get the file information
         info, err := zipfile.Stat()
         if err != nil {
                 return err
         }

         header, err := zip.FileInfoHeader(info)
         if err != nil {
                 return err
         }

         // default is Store 0(no compression!) 
         // forking important !!
         // change to deflate
         // see http://golang.org/pkg/archive/zip/#pkg-constants

         header.Method = zip.Deflate

         writer, err := zipit.CreateHeader(header)
         if err != nil {
                 return err
         }
         _, err = io.Copy(writer, zipfile)
         return err

 }
