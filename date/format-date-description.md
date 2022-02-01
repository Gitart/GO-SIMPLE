## Format date

|Format|	Example|
|------|---------|
|ANSIC	|“Mon Jan _2 15:04:05 2006”|
|UnixDate|	“Mon Jan _2 15:04:05 MST 2006”|
|RubyDate|	“Mon Jan 02 15:04:05 -0700 2006”|
|RFC822	|“02 Jan 06 15:04 MST”|
|RFC822Z|	“02 Jan 06 15:04 -0700”|
|RFC850	|“Monday, 02-Jan-06 15:04:05 MST”|
|RFC1123	|“Mon, 02 Jan 2006 15:04:05 MST”|
|RFC1123Z	|“Mon, 02 Jan 2006 15:04:05 -0700”|
|RFC3339	|“2006-01-02T15:04:05Z07:00”|
|RFC3339Nano	|“2006-01-02T15:04:05.999999999Z07:00”|

Layouts must use the reference time Mon Jan 2 15:04:05 MST 2006 to show the pattern with which to format/parse a given time/string.

## Example 1:

```go
// Golang program to illustrate the time
// formatting using custom layouts
  
package main
  
import (
    "fmt"
    "time"
)
  
func main() {
  
    // this function returns the present time
    current_time := time.Now()
  
    // individual elements of time can
    // also be called to print accordingly
    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
    current_time.Year(), current_time.Month(), current_time.Day(),
    current_time.Hour(), current_time.Minute(), current_time.Second())
  
    // formatting time using
    // custom formats
    fmt.Println(current_time.Format("2006-01-02 15:04:05"))
    fmt.Println(current_time.Format("2006-January-02"))
    fmt.Println(current_time.Format("2006-01-02 3:4:5 pm"))
}
```

Output:
```
2009-11-10T23:00:00-00:00
2009-11-10 23:00:00
2009-November-10
2009-11-10 11:0:0 pm
```

## Example 2:

```go
// Golang program to illustrate the time
// formatting using format constants
package main
  
import (
    "fmt"
    "time"
)
  
func main() {
  
    // this function returns the present time
    current_time := time.Now()
  
    // using inbuilt format constants
    // shown in the table above
    fmt.Println("ANSIC: ", current_time.Format(time.ANSIC))
    fmt.Println("UnixDate: ", current_time.Format(time.UnixDate))
    fmt.Println("RFC1123: ", current_time.Format(time.RFC1123))
    fmt.Println("RFC3339Nano: ", current_time.Format(time.RFC3339Nano))
    fmt.Println("RubyDate: ", current_time.Format(time.RubyDate))
}
```

Output:
```
ANSIC:  Tue Nov 10 23:00:00 2009
UnixDate:  Tue Nov 10 23:00:00 UTC 2009
RFC1123:  Tue, 10 Nov 2009 23:00:00 UTC
RFC3339Nano:  2009-11-10T23:00:00Z
RubyDate:  Tue Nov 10 23:00:00 +0000 2009
```
