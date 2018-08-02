# Read a text file and replace certain words
golang read replace 

### Problem :
You have a text file with some words that you need to replace with another word.

Solution :
Use the bytes.Replace() function. For example :

```go
 package main

 import (
         "bytes"
         "fmt"
         "io/ioutil"
         "os"
 )

 func main() {

         input, err := ioutil.ReadFile("original.txt")
         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         output := bytes.Replace(input, []byte("replaceme"), []byte("ok"), -1)

         if err = ioutil.WriteFile("modified.txt", output, 0666); err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }
 }
 ```
 
### original.txt

 this is a text file that contains couple of replaceme words that need to be replaced.
 for example, a replaceme word in this line


 and another [replaceme] word in this line too
and after executing the above code

### modified.txt

 this is a text file that contains couple of ok words that need to be replaced.
 for example, a ok word in this line


 and another [ok] word in this line too
