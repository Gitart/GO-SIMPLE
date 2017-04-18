
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

##delete request
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
