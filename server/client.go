package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)


func main() {
     url    := "http://localhost:8888/api/data"
     // method := "GET"

     http.PostForm(url, "name") 

}



func main1() {

  url    := "http://localhost:8888/api/data"
  method := "GET"

  payload := strings.NewReader(`
   {
    "test":"testing",
    "news":"News",
    "system":"Sysytem",
    "other":{"otherd":"mn"}
    }
  	`)

  client   := &http.Client {}
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  
  req.Header.Add("Authorization", "Basic QWRtaW46")
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(string(body))
}
