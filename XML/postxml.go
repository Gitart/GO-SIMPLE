package main

import (
    "log"
    "net/http"
    "strings"
)

func main() {
    const myurl = "http://127.0.0.1:8080"
    const xmlbody = `
<request>
    <parameters>
        <email>test@test.com</email>
        <password>test</password>
    </parameters>
</request>`

    resp, err := http.Post(myurl, "text/xml", strings.NewReader(xmlbody))
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    // Do something with resp.Body
}
