# Profiling between to time

```golang
package main

 import (
         "fmt"
         "log"
         "time"
 )

 func main() {
         t1 := time.Now()
         fmt.Println("Do something...")
         t2 := time.Now()
         log.Printf("[Do something] takes about %v\n", t2.Sub(t1))
 }
 
 ```
Sample output :

Do something...
2015/03/30 10:24:52 [Do something] takes about 36.466µs
Do something...
2015/03/30 10:24:54 [Do something] takes about 36.193µs
another example for web GET operation. This must take into account of your Internet connection speed as well.
```

```
 func indexHandler(w http.ResponseWriter, r *http.Request) {
   t1 := time.Now()
   fmt.Fprintf(w, "Hello World!")
   t2 := time.Now()
   log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
 }

 func main() {
   http.HandleFunc("/", indexHandler)
   http.ListenAndServe(":8080", nil)
 }
 ```
