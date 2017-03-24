## Replace in string

```golang
package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "/one/two/thee/1/3/red.jpg"
    str = strings.Replace(str, "/one/two/thee/", "", -1)
    fmt.Println(str)
}
```
