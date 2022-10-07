# Golang: Gracefully stop application

To shutdown go application gracefully, you can use open source libraries or write your own code.

Following are popular libraries to stop go application gracefully

1.  [https://github.com/tylerb/graceful](http://github.com/tylerb/graceful)
2.  [https://github.com/braintree/manners](http://github.com/braintree/manners)

In this article, I will explain how to write your own code to stop go app gracefully

**Step 1:** make channel which can listen for signals from OS. Refer [os.Signal](http://golang.org/pkg/os/signal/) package for more detail. os.Signal package is used to access incoming signals from OS.

var gracefulStop = make(chan os.Signal)

**Step 2:** Use notify method of os.Signal to register system calls. For gracefully stop. we should listen to SIGTERM and SIGINT. signal.Notify method takes two arguments 1. channel 2. constant from syscall.

signal.Notify(gracefulStop, syscall.SIGTERM)
signal.Notify(gracefulStop, syscall.SIGINT)

**Step 3:** Now, We needs to create Go routine to listen channel “gracefulStop” for incoming signals. the following Go routine will block until it receives signals from OS. Now, you can perform clean up your stuff it can be closing DB connections, clearing buffered channels, write something to file, etc.. In the following code, I just put wait for 2 seconds. After completing your work you need to send a signal to OS by using os.Exit function. os.Exit function takes integer argument normally, it can be 0 or 1. 0 means clean exit without any error or problem. 1 means exit with an error or some issue. The exit status will help caller to identify the last status when process end.

go func() {
       sig := <-gracefulStop
       fmt.Printf("caught sig: %+v", sig)
       fmt.Println("Wait for 2 second to finish processing")
       time.Sleep(2\*time.Second)
       os.Exit(0)
}()

**Full Source**

For the demo, I use simple HTTP server which will display “Server is running” message on the browser.

package mainimport (
       "os"
       "os/signal"
       "syscall"
       "fmt"
       "time"
       "net/http"
)func main() {
 http.HandleFunc("/", func(w http.ResponseWriter, r \*http.Request) {
              fmt.Fprint(w,"Server is running")
       }) var gracefulStop = make(chan os.Signal)
       signal.Notify(gracefulStop, syscall.SIGTERM)
       signal.Notify(gracefulStop, syscall.SIGINT)
       go func() {
              sig := <-gracefulStop
              fmt.Printf("caught sig: %+v", sig)
              fmt.Println("Wait for 2 second to finish processing")
              time.Sleep(2\*time.Second)
              os.Exit(0)
       }() http.ListenAndServe(":8080",nil)
}

**Console log when you stop go lang app**

 ![](https://miro.medium.com/max/700/0*I83wny7SxDNiPRiF.png)

**References**
