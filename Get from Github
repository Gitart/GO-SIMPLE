// Get from GITHUB
package main

import (
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

func main() {
    //resp, err := http.Get("https://api.github.com/repos/dotcloud/docker")
    resp, err := http.Get("https://raw.githubusercontent.com/Gitart/Projects/master/HeadOffice")
	
    
    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    _, err = os.Stdout.Write(body)

    if err != nil {
        log.Fatal(err)
    }
}
