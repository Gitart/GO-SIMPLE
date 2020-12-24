# Various examples of Carbon date\-time package in Golang

The package called Carbon can help make dealing with date/time in Golang much easier and more semantic so that our code can become more readable and maintainable.
Carbon is a package for Golang is available in Github. https://github.com/uniplaces/carbon

It provides some nice functionality to deal with dates in Golang. Specifically things like:
Easily calculate difference between dates
Dealing with timezones
Getting current time easily Converting a datetime into something readable
Parse an English phrase into datetime (first day of January 2016)
Add and Subtract dates (+ 2 weeks, \-6 months)
Semantic way of dealing with dates

Install Carbon:
go get github.com/uniplaces/carbon

Add to your imports to start using Carbon
import "github.com/uniplaces/carbon"

package main

import (
    "fmt"
    "github.com/uniplaces/carbon"
)

func main() {
    fmt.Printf("Right now is %s\\n", carbon.Now().DateTimeString())
    today, \_ := carbon.NowInLocation("Europe/London")
    fmt.Printf("Right now in London is %s\\n", today)

    fmt.Printf("\\n#######################################\\n")
    fmt.Printf("Tomorrow is %s\\n", carbon.Now().AddDay())
    fmt.Printf("Yesterday is %s\\n", carbon.Now().SubDay())
    fmt.Printf("After 5 Days %s\\n", carbon.Now().AddDays(5))
    fmt.Printf("Before 5 Days %s\\n", carbon.Now().SubDays(5))

    fmt.Printf("\\n#######################################\\n")
    fmt.Printf("Next Month is %s\\n", carbon.Now().AddMonth())
    fmt.Printf("Last Month is %s\\n", carbon.Now().SubMonth())

    fmt.Printf("\\n#######################################\\n")
    fmt.Printf("Next week is %s\\n", carbon.Now().AddWeek())
    fmt.Printf("Last week is %s\\n", carbon.Now().SubWeek())

    fmt.Printf("\\n#######################################\\n")
    fmt.Printf("Next Year %s\\n", carbon.Now().AddYear())
    fmt.Printf("Last Year %s\\n", carbon.Now().SubYear())
    fmt.Printf("After 5 Years %s\\n", carbon.Now().AddYears(5))
    fmt.Printf("Before 5 Years %s\\n", carbon.Now().SubYears(5))

    fmt.Printf("\\n#######################################\\n")
    fmt.Printf("Next Hour %s\\n", carbon.Now().AddHour())
    fmt.Printf("Last Hour %s\\n", carbon.Now().SubHour())
    fmt.Printf("After 5 Mins %s\\n", carbon.Now().AddMinutes(5))
    fmt.Printf("Before 5 Mins %s\\n", carbon.Now().SubMinutes(5))

    fmt.Printf("\\n#######################################\\n")
    fmt.Printf("Weekday? %t\\n", carbon.Now().IsWeekday())
    fmt.Printf("Weekend? %t\\n", carbon.Now().IsWeekend())
    fmt.Printf("LeapYear? %t\\n", carbon.Now().IsLeapYear())
    fmt.Printf("Past? %t\\n", carbon.Now().IsPast())
    fmt.Printf("Future? %t\\n", carbon.Now().IsFuture())

    fmt.Printf("\\n#######################################\\n")
    fmt.Printf("Start of day:   %s\\n", today.StartOfDay())
    fmt.Printf("End of day: %s\\n", today.EndOfDay())
    fmt.Printf("Start of month: %s\\n", today.StartOfMonth())
    fmt.Printf("End of month:   %s\\n", today.EndOfMonth())
    fmt.Printf("Start of year:  %s\\n", today.StartOfYear())
    fmt.Printf("End of year:    %s\\n", today.EndOfYear())
    fmt.Printf("Start of week:  %s\\n", today.StartOfWeek())
    fmt.Printf("End of week:    %s\\n", today.EndOfWeek())

}

C:\\golang\\time>go run t7.go
Right now is 2017\-08\-27 22:38:07
Right now in London is 2017\-08\-27 18:08:07

#######################################
Tomorrow is 2017\-08\-28 22:38:07
Yesterday is 2017\-08\-26 22:38:07
After 5 Days 2017\-09\-01 22:38:07
Before 5 Days 2017\-08\-22 22:38:07

#######################################
Next Month is 2017\-09\-27 22:38:07
Last Month is 2017\-07\-27 22:38:07

#######################################
Next week is 2017\-09\-03 22:38:07
Last week is 2017\-08\-20 22:38:07

#######################################
Next Year 2018\-08\-27 22:38:07
Last Year 2016\-08\-27 22:38:07
After 5 Years 2022\-08\-27 22:38:07
Before 5 Years 2012\-08\-27 22:38:07

#######################################
Next Hour 2017\-08\-27 23:38:07
Last Hour 2017\-08\-27 21:38:07
After 5 Mins 2017\-08\-27 22:43:07
Before 5 Mins 2017\-08\-27 22:33:07

#######################################
Weekday? false
Weekend? true
LeapYear? false
Past? false
Future? false

#######################################
Start of day: 2017\-08\-27 00:00:00
End of day: 2017\-08\-27 23:59:59
Start of month: 2017\-08\-01 00:00:00
End of month: 2017\-08\-31 23:59:59
Start of year: 2017\-01\-01 00:00:00
End of year: 2017\-12\-31 23:59:59
Start of week: 2017\-08\-21 00:00:00
End of week: 2017\-09\-03 23:59:59

C:\\golang\\time>
