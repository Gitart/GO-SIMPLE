## Get path name to current directory or folder
golang current-directory  

### Problem :
You need to get the path name of the current directory where the executable is residing. How to do that in Golang?
Solution :
Use the os.Getwd() function.

For example :

```go
 package main

 import (
         "fmt"
         "os"
         "strings"
 )

 func main() {
         dir, _ := os.Getwd()
         fmt.Println(strings.Replace(dir, " ", "\\ ", -1))
 }
 ```
