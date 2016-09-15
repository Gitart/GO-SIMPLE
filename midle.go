package main

import (
    "fmt"
    "log"
    "net/http"
    "net/http/httptest"
)

type ModifierMiddleware struct {
    handler http.Handler
}

func (m *ModifierMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    rec := httptest.NewRecorder()
    // passing a ResponseRecorder instead of the original RW
    m.handler.ServeHTTP(rec, r)
    // after this finishes, we have the response recorded
    // and can modify it before copying it to the original RW
    // we copy the original headers first
    for k, v := range rec.Header() {
        w.Header()[k] = v
    }
    // and set an additional one
    w.Header().Set("X-We-Modified-This", "Yup")
    // only then the status code, as this call writes the headers as well
    w.WriteHeader(480)
    // the body hasn't been written yet, so we can prepend some data.
    w.Write([]byte("Middleware says hello again. From me !\n --> "))
    // then write out the original body
    w.Write(rec.Body.Bytes())
}
func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world! with the middleware help")
}
func main() {
    // uncomment these 3 lines and comment those that follow, then try again.
    // http.HandleFunc("/", rootHandler)
    // println("Listening on port 8080")
    // log.Fatal(http.ListenAndServe(":8080", nil))
    mid := &ModifierMiddleware{
        http.HandlerFunc(rootHandler),
    }
    println("Listening on port 8080")
    // sign: func ListenAndServe(addr string, handler Handler) error
    log.Fatal(http.ListenAndServe(":8080", mid))
}
