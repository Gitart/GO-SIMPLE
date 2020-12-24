# Get Year, Month, Day, Hour, Min and Second from a specified date

package main
import (
         "fmt"
         "time"
)
func main() {
    t := time.Date(2015, 02, 21, 23, 10, 52, 211, time.UTC)
    fmt.Println(t)
    fmt.Println("\\n######################################\\n")

    y := t.Year()
    mon := t.Month()
    d := t.Day()
    h := t.Hour()
    m := t.Minute()
    s := t.Second()
    n := t.Nanosecond()

    fmt.Println("Year   :",y)
    fmt.Println("Month  :",mon)
    fmt.Println("Day    :",d)
    fmt.Println("Hour   :",h)
    fmt.Println("Minute :",m)
    fmt.Println("Second :",s)
    fmt.Println("Nanosec:",n)
}

C:\\golang\\time>go run t7.go
2015\-02\-21 23:10:52.000000211 +0000 UTC

######################################

Year : 2015
Month : February
Day : 21
Hour : 23
Minute : 10
Second : 52
Nanosec: 211

C:\\golang\\time>
