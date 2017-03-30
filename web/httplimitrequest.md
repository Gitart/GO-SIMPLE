## Delay or limit HTTP requests example


### Problem:
For some technical limitation reason, you want to delay or slow down HTTP requests to one of your machines  
in order not to overwhelm the CPU. How to do that?  

### Solution:
Introduce a delay to the HTTP handler or prioritise the request with job queues.  
In this code example that follows, we will use the simplest delay mechanism for delaying HTTP requests.  


```golang
 package main

 import (
 	"log"
 	"net/http"
 	"time"
 )

 // some global variables
 var (
 	delayer <-chan time.Time
 	counter int
 	seconds time.Duration
 )

 func delayedHelloWorld(w http.ResponseWriter, r *http.Request) {
 	<-delayer
 	counter++
 	log.Printf("Processing request #%d - delayed HTTP request by %d seconds", counter, seconds)
 	w.Write([]byte("Hello, processed your request."))

 }

 func HelloWorld(w http.ResponseWriter, r *http.Request) {
 	<-delayer
 	counter++
 	log.Printf("Processing request #%d - delayed HTTP request by %d seconds", counter, seconds)
 	w.Write([]byte("Hello, processed your request."))
 }

 func main() {

 	counter = 0
 	seconds = 3
 	delayer = time.Tick(seconds * time.Second)

 	mux := http.NewServeMux()
 	mux.Handle("/", http.HandlerFunc(delayedHelloWorld)) // for Handle() method
 	mux.HandleFunc("/helloworld", HelloWorld)            // for HandleFunc() method

 	http.ListenAndServe(":8080", mux)

 }
 ```
 
Run the code and see the log output on the terminal.

Next, point your browsers to http://localhost:8080 and make couple of requests by clicking on the refresh
button(or press F5 button repeatedly). You should see that all the requests are delayed by 3 seconds.

```
2016/10/17 13:06:57 Processing request #16 - delayed HTTP request by 3 seconds
2016/10/17 13:07:00 Processing request #17 - delayed HTTP request by 3 seconds
2016/10/17 13:07:03 Processing request #18 - delayed HTTP request by 3 seconds
2016/10/17 13:07:06 Processing request #19 - delayed HTTP request by 3 seconds
2016/10/17 13:07:09 Processing request #20 - delayed HTTP request by 3 seconds
```

### NOTES:

If you're looking for job queuing example, see https://gist.github.com/harlow/dbcd639cf8d396a2ab73
Apart from that, you can also choose to filter requests made by legit Go client or HTTP POST or GET.
