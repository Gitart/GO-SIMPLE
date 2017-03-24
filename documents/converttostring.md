## Convert int64 to string

```golang
package main

import (
    "strconv"
    "fmt"
)

func main() {
    var num int64
    num = 32534535567545675
    t := strconv.FormatInt(num, 10)
    t += "string"
    fmt.Println(t)
}
```
