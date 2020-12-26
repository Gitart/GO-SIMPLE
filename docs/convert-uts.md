# Convert specific UTC date time to PST, HST, MST and SGT

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    t, err := time.Parse("2006 01 02 15 04", "2015 11 11 16 50")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(t)

    loc, err := time.LoadLocation("America/Los\_Angeles")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(loc)

    t = t.In(loc)
    fmt.Println(t.Format(time.RFC822))

    loc, err = time.LoadLocation("Singapore")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(loc)

    t = t.In(loc)
    fmt.Println(t.Format(time.RFC822))

    loc, err = time.LoadLocation("US/Hawaii")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(loc)

    t = t.In(loc)
    fmt.Println(t.Format(time.RFC822))

    loc, err = time.LoadLocation("EST")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(loc)

    t = t.In(loc)
    fmt.Println(t.Format(time.RFC822))

    loc, err = time.LoadLocation("MST")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(loc)

    t = t.In(loc)
    fmt.Println(t.Format(time.RFC822))

}
```
```
C:\\golang\\time>go run t7.go
2015\-11\-11 16:50:00 +0000 UTC
America/Los\_Angeles
11 Nov 15 08:50 PST
Singapore
12 Nov 15 00:50 SGT
US/Hawaii
11 Nov 15 06:50 HST
EST
11 Nov 15 11:50 EST
MST
11 Nov 15 09:50 MST
```

C:\\golang\\time>
