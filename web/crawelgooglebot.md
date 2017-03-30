 ## How to determine if request or crawl is from Google robots

For this tutorial, we will learn how to detect a visit by Google robots or web crawlers and learn 
how to distinguish fake/spoof user agents from bad actors pretending to be Google.
Taking example from the official guide from Google on how to see which robots Google uses to crawl website. 
We can see that all of the user agents contain the string "google".
The simplest way to detect Google robots or crawlers should look like this :


```golang
 package main

 import (
         "strings"
 	"fmt"
 	"net/http"
 )

 func getUserAgent(w http.ResponseWriter, r *http.Request) {
 	ua := r.Header.Get("User-Agent")

 	fmt.Printf("user agent is: %s \n", ua)
 	w.Write([]byte("user agent is " + ua))

 	ualow := strings.ToLower(ua)

 	if strings.Contains(ualow, "google") {
 		fmt.Println("Visited by Google bot")
 	} else {
 		fmt.Println("Visited by some thing else")
 	}
 }

 func main() {
 	http.HandleFunc("/", getUserAgent)
 	http.ListenAndServe(":8080", nil)
 }
 ```
 
Sample output :

```
These are the results simulated by Google PageSpeed Insights
```

user agent is: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko; Google Page Speed Insights) Chrome/27.0.1453 Safari/537.36
Visited by Google bot
user agent is: Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/537.36 (KHTML, like Gecko; Google Page Speed Insights) Version/8.0 Mobile/12F70 Safari/600.1.4
Visited by Google bot
However, the above code is kinda primitive and can be easily spoofed. We need to verify if the Google robot is indeed genuine by following this guideline from Google on how to verify Google robots.


```golang
 package main

 import (
 	"fmt"
 	"net"
 	"net/http"
 	"os/exec"
 	"strings"
         "os"
 )

 func getUserAgent(w http.ResponseWriter, r *http.Request) {
 	ua := r.Header.Get("User-Agent")

 	fmt.Printf("user agent is: %s \n", ua)
 	w.Write([]byte("user agent is " + ua))

 	ualow := strings.ToLower(ua)

 	if strings.Contains(ualow, "google") {
 		fmt.Println("Visited by Google bot")
 	} else {
 		fmt.Println("Visited by some thing else")
 	}

 	// get IP address
 	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

 	// capture the output of host command to verify Google robots
 	// based on https://support.google.com/webmasters/answer/80553

 	cmd := exec.Command("host", ip)
 	result, err := cmd.Output() // capture the exec output to variable result

 	if err != nil {
 		//fmt.Println(err)
 		fmt.Printf("Host %s command execution failed \n", ip)
 		os.Exit(1)
 	}

 	fmt.Println("Host reply : ", string(result))

 	// if result contain the word google, then it is genuine user agent
 	// else fake

 	if strings.Contains(strings.ToLower(string(result)), "google") {
 		fmt.Println(" and the user agent is real. ")
 	} else {
 		fmt.Println(" and the user agent is determine to be FAKED after verifying with host command. ")
 	}

 }

 func main() {
 	http.HandleFunc("/", getUserAgent)
 	http.ListenAndServe(":8080", nil)
 }
 ```
 
