package main

import (
    "context"
    "io"
    "log"
    "net/http"
    "os"
    "time"
)

func getContent(ctx context.Context) {
    req, err := http.NewRequest("GET", "http://example.com", nil)
    if err != nil {
        log.Fatal(err)
    }
    ctx, cancel := context.WithDeadline(ctx, time.Now().Add(3 * time.Second))
    defer cancel()

    req.WithContext(ctx)

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    io.Copy(os.Stdout, resp.Body)
}

func main() {
    ctx := context.Background()
    getContent(ctx)
}
If you want to make cancel trigger on main:

func main() {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)

    sc := make(chan os.Signal, 1)
    signal.Notify(sc, os.Interrupt)
    go func(){
        <-sc
        cancel()
    }()

    getContent(ctx)
}
