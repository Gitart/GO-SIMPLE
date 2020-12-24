# Subtract N number of Year, Month, Day, Hour, Minute, Second, Millisecond, Microsecond and Nanosecond to current date\-time.

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
    fmt.Println("Today:", now)

    after := now.AddDate(\-1, 0, 0)
    fmt.Println("Subtract 1 Year:", after)

    after = now.AddDate(0, \-1, 0)
    fmt.Println("Subtract 1 Month:", after)

    after = now.AddDate(0, 0, \-1)
    fmt.Println("Subtract 1 Day:", after)

    after = now.AddDate(\-2, \-2, \-5)
    fmt.Println("Subtract multiple values:", after)

    after = now.Add(\-10\*time.Minute)
    fmt.Println("Subtract 10 Minutes:", after)

    after = now.Add(\-10\*time.Second)
    fmt.Println("Subtract 10 Second:", after)

    after = now.Add(\-10\*time.Hour)
    fmt.Println("Subtract 10 Hour:", after)

    after = now.Add(\-10\*time.Millisecond)
    fmt.Println("Subtract 10 Millisecond:", after)

    after = now.Add(\-10\*time.Microsecond)
    fmt.Println("Subtract 10 Microsecond:", after)

    after = now.Add(\-10\*time.Nanosecond)
    fmt.Println("Subtract 10 Nanosecond:", after)
}

C:\\golang\\time>go run t5.go
Today: 2017\-08\-27 12:21:17.8379942 +0530 IST
Subtract 1 Year: 2016\-08\-27 12:21:17.8379942 +0530 IST
Subtract 1 Month: 2017\-07\-27 12:21:17.8379942 +0530 IST
Subtract 1 Day: 2017\-08\-26 12:21:17.8379942 +0530 IST
Subtract multiple values: 2015\-06\-22 12:21:17.8379942 +0530 IST
Subtract 10 Minutes: 2017\-08\-27 12:11:17.8379942 +0530 IST
Subtract 10 Second: 2017\-08\-27 12:21:07.8379942 +0530 IST
Subtract 10 Hour: 2017\-08\-27 02:21:17.8379942 +0530 IST
Subtract 10 Millisecond: 2017\-08\-27 12:21:17.8279942 +0530 IST
Subtract 10 Microsecond: 2017\-08\-27 12:21:17.8379842 +0530 IST
Subtract 10 Nanosecond: 2017\-08\-27 12:21:17.83799419 +0530 IST

C:\\golang\\time>
