## Simple client
```golang
package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

func main() {
  res, _ :=http.Get("http://goinpracticebook.com")
  b, _ := ioutil.ReadAll(res.Body)
  res.Body.Close()
  fmt.Printf("%s", b)
}
```

## delete request
```golang

package main
  import (
  "fmt"
  "net/http"
)

func main() {
  req, _ := http.NewRequest("DELETE","http://example.com/foo/bar", nil)
  res, _ := http.DefaultClient.Do(req)
  fmt.Printf("%s", res.Status)
}
```

## custom client

```golang

func main() {
  cc := &http.Client{Timeout: time.Second}
  res, err :=cc.Get("http://goinpracticebook.com")

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  b, _ := ioutil.ReadAll(res.Body)
  res.Body.Close()
  fmt.Printf("%s", b)
}
```



## Has time out

```golang

func hasTimedOut(err error) bool {
  switch err := err.(type) {
  case *url.Error:
  if err, ok := err.Err.(net.Error); ok && err.Timeout() {
  return true
}


case net.Error:
if err.Timeout() {
  return true
}

case *net.OpError:
if err.Timeout() {
  return true
}
}

errTxt := "use of closed network connection"
if err != nil && strings.Contains(err.Error(), errTxt) {
  return true
}
return false
}


//This function provides the capability to detect a variety of timeout situations. The fol-
//lowing snippet is an example of using that function to check whether an error was
//caused by a timeout:

res, err := http.Get("http://example.com/test.zip")
if err != nil && hasTimedOut(err) {
fmt.Println("A timeout error occured")
return
}
```
