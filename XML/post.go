package main

import (
    "bytes"
    "fmt"
    "net/http"
)

func main() {
    // or you can use []byte(`...`) and convert to Buffer later on
    body := "<request> <parameters> <email>test@test.com</email> <password>test</password> </parameters> </request>"

    client := &http.Client{}
    // build a new request, but not doing the POST yet
    req, err := http.NewRequest("POST", "http://localhost:8080/", bytes.NewBuffer([]byte(body)))
    if err != nil {
        fmt.Println(err)
    }
    // you can then set the Header here
    // I think the content-type should be "application/xml" like json...
    req.Header.Add("Content-Type", "application/xml; charset=utf-8")
    // now POST it
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(resp)
}
