## Versions in modules




```golang
package main

 import (
         "fmt"
         "time"
 )

 var (
         version   string
         timeStamp string
 )

 func main() {
         timeStamp := time.Now()

         fmt.Println("Running version : ", version)
         fmt.Println("Build time: ", timeStamp)
 }
 
 ```
 
 
For example:
```
go build -ldflags "-X main.version=v1.1" buildinfo.go
```
when executing the binary, you will see this during run time:

$ ./buildinfo

```
Running version : v1.1
Build time: 2017-02-15 14:35:35.598182767 +0800 SGT
```

### NOTES:
If you encounter this error message during compilation time:
```
-X flag requires argument of the form importpath.name=value
```

simply change "-X main.version v0.1" to "-X main.version=v0.1"


## Other

```golang
 package main

 import "fmt"

 var str string

 func main() {
     fmt.Println(str)
 }
 ```
 
and execute the program with -ldflags to initialize the str variable value
```
$ go run -ldflags "-X main.str abc" main.go
```

or if you prefer, build it first into executable binary :
```
$ go build -ldflags "-X main.str 'abc'" main.go && ./main
```


Output :

abc
