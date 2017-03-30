## Identifying Golang HTTP client request

### Problem:
You've read about the Mirai DDoS malware that infects IoT devices and turning them into a massive botnet. 
You're developing a new server program for collecting data from your IoT devices(clients) 
with Golang on HTTP protocol rather than MQTT. (but why!? - oh WiFI instead of Bluetooth)

To secure the devices, you change the factory default usernames and passwords upon unpacking/setup stage. Great!

For added security, you want to eliminate unwanted connections from potential malware or some random HTTP 
requests such as from browsers, you also want your program to accept valid connection from Golang client program only. How to identify a connection make by another Golang client?

### NOTES:

1) This tutorial is written "for" a friend.
2)This is not a foolproof solution, but good enough to be used in early stage of requests filtering. A malicious coder can still spoof the user agent. For better security, consider implementing authentication mechanism with HMAC. See https://www.socketloop.com/tutorials/golang-simple-client-server-hmac-authentication-without-ssl-example.
3)The potential problem is that .... you might have to hard code the HMAC's secret key in your program.

### SOLUTION:

Check the user agent string in the request header. If the user agent starts with Go-http, you can be certain that the request 
is made by a Golang program.

```
 GET / HTTP/1.1
 User-Agent [Go-http-client/1.1] <------- here
 Accept-Encoding [gzip]
 Detected IP address is :  46.166.190.130:3619
 Real IP address could be :
 Go HTTP client detected <---------- and here
 ```
 
Here is the code example that you can use to identify Golang HTTP client request. Run the code and visit the server
with another Golang HTTP client. You might want to do an anonymous visit via Tor network as well.
See https://www.socketloop.com/tutorials/golang-accessing-content-anonymously-with-tor

```golang
 package main

 import (
 	"fmt"
 	"net/http"
 	"strings"
 )

 func checkRequest(w http.ResponseWriter, r *http.Request) {

 	// similar to PHP's $_SERVER["REQUEST_METHOD"], $_SERVER["REQUEST_URI"] & $_SERVER["REQUEST_PROTOCOL"]
 	basicInfo := r.Method + " " + r.RequestURI + " " + r.Proto

 	fmt.Println(basicInfo)

 	// some extra header information
 	headerInfo := r.Header

 	for key, value := range headerInfo {
 		fmt.Println(key, value)
 	}

 	// similar to $_SERVER["REMOTE_ADDR"]

 	fmt.Println("Detected IP address is : ", r.RemoteAddr)

 	// in case request is via proxy server
 	fmt.Println("Real IP address could be : ", r.Header.Get("X-Forwarded-For"))

 	if strings.HasPrefix(r.Header.Get("User-Agent"), "Go-http-client") {
 		fmt.Println("Go HTTP client detected")
             // then continue to process request or perform HMAC authentication
             // else ignore or block further request from the originating IP address.
 	}

 	fmt.Println("-------------------------------------------------------")
 }

 func main() {
 	http.HandleFunc("/", checkRequest)
 	http.ListenAndServe(":8080", nil)
 }
 ```
