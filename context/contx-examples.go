package main
import (
    "fmt"
    "net/http"
    "time"
)

// Handler for /example request
func example(w http.ResponseWriter, req *http.Request) {

    fmt.Println("example handler started")

    // Accessing the context of the request
    context := req.Context()

    select {

    // Simulating some work by the server
    // Waits 10 seconds and then responds with "example\n"
    case <-time.After(10 * time.Second):
        fmt.Fprintf(w, "example\n")

    // Handling request cancellation
    case <-context.Done():
        err := context.Err()
        fmt.Println("server:", err)
    }

    fmt.Println("example handler ended")
}

func main() {

    http.HandleFunc("/example", example)
    http.ListenAndServe(":5000", nil)
}
