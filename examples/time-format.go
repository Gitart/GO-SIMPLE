// Golang program to get the current
// date and time in various format
package main
   
import (
    "fmt"
    "time"
)
   
func main() {
   
    // using time.Now() function
    // to get the current time
    currentTime := time.Now()
   
    // getting the time in string format
    fmt.Println("Show Current Time in String: ", currentTime.String())
    fmt.Println("YYYY.MM.DD : ", currentTime.Format("2017.09.07 17:06:06"))
    fmt.Println("YYYY#MM#DD {Special Character} : ", currentTime.Format("2017#09#07"))
    fmt.Println("MM-DD-YYYY : ", currentTime.Format("09-07-2017"))
    fmt.Println("YYYY-MM-DD : ", currentTime.Format("2017-09-07"))
    fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2017-09-07 17:06:06"))
    fmt.Println("Time with MicroSeconds: ", currentTime.Format("2017-09-07 17:06:04.000000"))
    fmt.Println("Time with NanoSeconds: ", currentTime.Format("2017-09-07 17:06:04.000000000"))
    fmt.Println("ShortNum Width : ", currentTime.Format("2017-02-07"))
    fmt.Println("ShortYear : ", currentTime.Format("06-Feb-07"))
    fmt.Println("LongWeekDay : ", currentTime.Format("2017-09-07 17:06:06 Wednesday"))
    fmt.Println("ShortWeek Day : ", currentTime.Format("2017-09-07 Wed"))
    fmt.Println("ShortDay : ", currentTime.Format("Wed 2017-09-2"))
    fmt.Println("LongWidth : ", currentTime.Format("2017-March-07"))
    fmt.Println("ShortWidth : ", currentTime.Format("2017-Feb-07"))
    fmt.Println("Short Hour Minute Second: ", currentTime.Format("2017-09-07 2:3:5 PM"))
    fmt.Println("Short Hour Minute Second: ", currentTime.Format("2017-09-07 2:3:5 pm"))
    fmt.Println("Short Hour Minute Second: ", currentTime.Format("2017-09-07 2:3:5"))
}

/*
  Show Current Time in String:  2009-11-10 23:00:00 +0000 UTC m=+0.000000001
  YYYY.MM.DD :  10117.09.07 117:09:09
  YYYY#MM#DD {Special Character} :  10117#09#07
  MM-DD-YYYY :  09+00-10117
  YYYY-MM-DD :  10117-09+00
  YYYY-MM-DD hh:mm:ss :  10117-09+00 117:09:09
  Time with MicroSeconds:  10117-09+00 117:09:00.000000
  Time with NanoSeconds:  10117-09+00 117:09:00.000000000
  ShortNum Width :  10117-10+00
  ShortYear :  09-Feb+00
  LongWeekDay :  10117-09+00 117:09:09 Wednesday
  ShortWeek Day :  10117-09+00 Wed
  ShortDay :  Wed 10117-09-10
  LongWidth :  10117-March+00
  ShortWidth :  10117-Feb+00
  Short Hour Minute Second:  10117-09+00 10:11:0 PM
  Short Hour Minute Second:  10117-09+00 10:11:0 pm
  Short Hour Minute Second:  10117-09+00 10:11:0
*/
