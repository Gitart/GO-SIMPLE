# Get current date and time in various format in golang

package main

import (
    "fmt"
    "time"
)
func main() {

    currentTime := time.Now()

    fmt.Println("Current Time in String: ", currentTime.String())

    fmt.Println("MM\-DD\-YYYY : ", currentTime.Format("01\-02\-2006"))

    fmt.Println("YYYY\-MM\-DD : ", currentTime.Format("2006\-01\-02"))

    fmt.Println("YYYY.MM.DD : ", currentTime.Format("2006.01.02 15:04:05"))

    fmt.Println("YYYY#MM#DD {Special Character} : ", currentTime.Format("2006#01#02"))

    fmt.Println("YYYY\-MM\-DD hh:mm:ss : ", currentTime.Format("2006\-01\-02 15:04:05"))

    fmt.Println("Time with MicroSeconds: ", currentTime.Format("2006\-01\-02 15:04:05.000000"))

    fmt.Println("Time with NanoSeconds: ", currentTime.Format("2006\-01\-02 15:04:05.000000000"))

    fmt.Println("ShortNum Month : ", currentTime.Format("2006\-1\-02"))

    fmt.Println("LongMonth : ", currentTime.Format("2006\-January\-02"))

    fmt.Println("ShortMonth : ", currentTime.Format("2006\-Jan\-02"))

    fmt.Println("ShortYear : ", currentTime.Format("06\-Jan\-02"))

    fmt.Println("LongWeekDay : ", currentTime.Format("2006\-01\-02 15:04:05 Monday"))

    fmt.Println("ShortWeek Day : ", currentTime.Format("2006\-01\-02 Mon"))

    fmt.Println("ShortDay : ", currentTime.Format("Mon 2006\-01\-2"))

    fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006\-01\-02 3:4:5"))

    fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006\-01\-02 3:4:5 PM"))

    fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006\-01\-02 3:4:5 pm"))
}

C:\\golang\\time>go run t6.go
Current Time in String: 2017\-07\-04 00:47:20.1424751 +0530 IST
MM\-DD\-YYYY : 07\-04\-2017
YYYY\-MM\-DD : 2017\-07\-04
YYYY.MM.DD : 2017.07.04 00:47:20
YYYY#MM#DD {Special Character} : 2017#07#04
YYYY\-MM\-DD hh:mm:ss : 2017\-07\-04 00:47:20
Time with MicroSeconds: 2017\-07\-04 00:47:20.142475
Time with NanoSeconds: 2017\-07\-04 00:47:20.142475100
ShortNum Month : 2017\-7\-04
LongMonth : 2017\-July\-04
ShortMonth : 2017\-Jul\-04
ShortYear : 17\-Jul\-04
LongWeekDay : 2017\-07\-04 00:47:20 Tuesday
ShortWeek Day : 2017\-07\-04 Tue
ShortDay : Tue 2017\-07\-4
Short Hour Minute Second: 2017\-07\-04 12:47:20
Short Hour Minute Second: 2017\-07\-04 12:47:20 AM
Short Hour Minute Second: 2017\-07\-04 12:47:20 am

C:\\golang\\time>
