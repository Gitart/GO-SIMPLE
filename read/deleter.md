## Delete files
Being able to select certain type of files to delete can be useful. In this follow up tutorial from Golang : Delete file, we are going to refine it further by only deleting files with certain extension with filepath.Ext() function.
In our case, it will files with .png extension.
Adapting the codes from earlier tutorial, we will add one additional if statement into the code to only delete files with .png extension.
The extra line is if filepath.Ext(file.Name()) == ".png"


```go
 package main

  import (
      "fmt"
      "os"
      "path/filepath"
  )

  func main() {

      dirname := "." + string(filepath.Separator)

      d, err := os.Open(dirname)
      if err != nil {
          fmt.Println(err)
          os.Exit(1)
      }
      defer d.Close()

      files, err := d.Readdir(-1)
      if err != nil {
          fmt.Println(err)
          os.Exit(1)
      }

      fmt.Println("Reading "+ dirname)

      for _, file := range files {
          if file.Mode().IsRegular() {
              if filepath.Ext(file.Name()) == ".png" {
                os.Remove("file.Name()")
                fmt.Println("Deleted ", file.Name())
              }
          }
      }
  }
  ```
  
