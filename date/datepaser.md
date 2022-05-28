## Go datetime parse


In this article, we show how to parse datetime values in Golang with time.Parse and time.ParseInLocation.
func Parse(layout, value string) (Time, error)

The parse function parses a formatted string and returns the time value it represents.

func ParseInLocation(layout, value string, loc *Location) (Time, error)

The ParseInLocation function parses a formatted string and returns the time value it represents, while taking a location (timezone) into account.

Unlike most other languages, Go does not use the usual approach of using format specifiers such as yyyy-mm-dd to parse datetime values. Instead, it uses a uniq datetime value of Mon Jan 2 15:04:05 MST 2006. So in order to parse a specifec datetime value, we choose a specific layout of this very instant of time.
```go
const (

    Layout      = "01/02 03:04:05PM '06 -0700"
    ANSIC       = "Mon Jan _2 15:04:05 2006"
    UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
    RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
    RFC822      = "02 Jan 06 15:04 MST"
    RFC822Z     = "02 Jan 06 15:04 -0700"
    RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
    RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
    RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700"
    RFC3339     = "2006-01-02T15:04:05Z07:00"
    RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
    Kitchen     = "3:04PM"

    // Handy time stamps.
    Stamp      = "Jan _2 15:04:05"
    StampMilli = "Jan _2 15:04:05.000"
    StampMicro = "Jan _2 15:04:05.000000"
    StampNano  = "Jan _2 15:04:05.000000000"
)
```
There are several predefined layouts in the time module available.
Go time.Parse example

In the first example, we use the time.Parse function to parse a few datetime values.
```go
main.go

package main

import (
    "fmt"
    "time"
)

// Mon Jan 2 15:04:05 MST 2006

func main() {

    v1 := "2022/05/12"
    v2 := "14:55:23"
    v3 := "2014-11-12T11:45:26.37"

    const (
        layout1 = "2006/01/02"
        layout2 = "15:04:05"
        layout3 = "2006-01-02T15:04:05"
    )

    t, err := time.Parse(layout1, v1)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(t.Format(time.UnixDate))

    t, err = time.Parse(layout2, v2)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(t.Format(time.Kitchen))

    t, err = time.Parse(layout3, v3)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(t.Format(time.UnixDate))
}
```
We parse three datetime values.
```go
v1 := "2022/05/12"
v2 := "14:55:23"
v3 := "2014-11-12T11:45:26.37"
```
We have three datetime strings written in various formats.
```go
const (
    layout1 = "2006/01/02"
    layout2 = "15:04:05"
    layout3 = "2006-01-02T15:04:05"
)
```

In order to read the datime strings, we need to prepare the appropriate layouts.
```go
t, err := time.Parse(layout1, v1)

if err != nil {
    fmt.Println(err)
}

fmt.Println(t.Format(time.UnixDate))
```

We read the first value with time.Parse and check for error value. Then we print the value and format it to time.UnixDate with Format function.

$ go run main.go
Thu May 12 00:00:00 UTC 2022
2:55PM
Wed Nov 12 11:45:26 UTC 2014

Predefined datetime layouts

There are several predefined datetime layouts in the time module.
main.go
```go
package main

import (
    "fmt"
    "time"
)

func main() {

    dates := []string{
        "Sat May 28 11:54:40 CEST 2022",
        "Sat May 28 11:54:40 2022",
        "Sat, 28 May 2022 11:54:40 CEST",
        "28 May 22 11:54 CEST",
        "2022-05-28T11:54:40.809289619+02:00",
        "Sat May 28 11:54:40 +0200 2022",
    }

    layouts := []string{
        time.UnixDate,
        time.ANSIC,
        time.RFC1123,
        time.RFC822,
        time.RFC3339Nano,
        time.RubyDate,
    }

    for i := 0; i < len(dates); i++ {

        parsed, err := time.Parse(layouts[i], dates[i])

        if err != nil {
            fmt.Println(err)
        }

        fmt.Println(parsed)
    }
}
```
In the example, we have a slice of six datetime strings. We use appropriated predefined datetime constants to parse them.

$ go run main.go
```
2022-05-28 11:54:40 +0200 CEST
2022-05-28 11:54:40 +0000 UTC
2022-05-28 11:54:40 +0200 CEST
2022-05-28 11:54:00 +0200 CEST
2022-05-28 11:54:40.809289619 +0200 CEST
2022-05-28 11:54:40 +0200 CEST
```
Go time.ParseInLocation example

The time.ParseInLocation function also takes the timezone into account when parsing datetime values.
main.go
```go
package main

import (
    "fmt"
    "log"
    "time"
)

func main() {

    loc, err := time.LoadLocation("Local")

    if err != nil {
        log.Println(err)
    }

    date := "Sat May 28 11:54:40 2022"

    parsed, err := time.ParseInLocation(time.ANSIC, date, loc)

    if err != nil {
        log.Println(err)
    }

    fmt.Println(parsed)

    loc, err = time.LoadLocation("Europe/Moscow")

    if err != nil {
        log.Println(err)
    }

    parsed, err = time.ParseInLocation(time.ANSIC, date, loc)

    if err != nil {
        log.Println(err)
    }

    fmt.Println(parsed)
}
```
The example parses the given date in local and Europe/Moscow locations.

$ go run main.go
```
2022-05-28 11:54:40 +0200 CEST
2022-05-28 11:54:40 +0300 MSK
```
