package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
)

// global var
var (
    port     int
    visits   int = 1
    httpAddr string
)

func visitServer(visitChan chan bool) {
    // run continuously but do something only when the channel receives 'true'
    for {
        if <-visitChan {
            visits++
            log.Printf("%v visits", visits)
        }
    }
}

type withChanHandler struct {
    visitChan chan bool
}

func (h *withChanHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    h.visitChan <- true
    fmt.Fprintf(w, "Hello, World!\n number of visits: %d\n", visits)
}

// done once
func init() {
    flag.IntVar(&port, "port", 8080, "HTTP Server Port")
    flag.Parse()
    httpAddr = fmt.Sprintf(":%v", port)
}
func main() {
    var visitChan = make(chan bool)
    go visitServer(visitChan)
    http.Handle("/", &withChanHandler{visitChan})
    log.Printf("Listening to %v", httpAddr)
    log.Fatal(http.ListenAndServe(httpAddr, nil))
}
