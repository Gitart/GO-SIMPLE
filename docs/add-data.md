# Add N number of Year, Month, Day, Hour, Minute, Second, Millisecond, Microsecond and Nanosecond to current date\-time

In below example AddDate and Add function used from Golang Time package.

**Syntax:**

func (t Time) AddDate(years int, months int, days int) Time
func (t Time) Add(d Duration) Time

package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Println("\\nToday:", now)

    after := now.AddDate(1, 0, 0)
    fmt.Println("\\nAdd 1 Year:", after)

    after = now.AddDate(0, 1, 0)
    fmt.Println("\\nAdd 1 Month:", after)

    after = now.AddDate(0, 0, 1)
    fmt.Println("\\nAdd 1 Day:", after)

    after = now.AddDate(2, 2, 5)
    fmt.Println("\\nAdd multiple values:", after)

    after = now.Add(10\*time.Minute)
    fmt.Println("\\nAdd 10 Minutes:", after)

    after = now.Add(10\*time.Second)
    fmt.Println("\\nAdd 10 Second:", after)

    after = now.Add(10\*time.Hour)
    fmt.Println("\\nAdd 10 Hour:", after)

    after = now.Add(10\*time.Millisecond)
    fmt.Println("\\nAdd 10 Millisecond:", after)

    after = now.Add(10\*time.Microsecond)
    fmt.Println("\\nAdd 10 Microsecond:", after)

    after = now.Add(10\*time.Nanosecond)
    fmt.Println("\\nAdd 10 Nanosecond:", after)
}

C:\\golang\\time>go run t4.go

Today: 2017\-08\-27 11:17:54.1224628 +0530 IST

Add 1 Year: 2018\-08\-27 11:17:54.1224628 +0530 IST

Add 1 Month: 2017\-09\-27 11:17:54.1224628 +0530 IST

Add 1 Day: 2017\-08\-28 11:17:54.1224628 +0530 IST

Add multiple values: 2019\-11\-01 11:17:54.1224628 +0530 IST

Add 10 Minutes: 2017\-08\-27 11:27:54.1224628 +0530 IST

Add 10 Second: 2017\-08\-27 11:18:04.1224628 +0530 IST

Add 10 Hour: 2017\-08\-27 21:17:54.1224628 +0530 IST

Add 10 Millisecond: 2017\-08\-27 11:17:54.1324628 +0530 IST

Add 10 Microsecond: 2017\-08\-27 11:17:54.1224728 +0530 IST

Add 10 Nanosecond: 2017\-08\-27 11:17:54.12246281 +0530 IST

C:\\golang\\time>
