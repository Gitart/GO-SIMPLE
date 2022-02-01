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
Output:

ANSIC:  Tue Nov 10 23:00:00 2009
UnixDate:  Tue Nov 10 23:00:00 UTC 2009
RFC1123:  Tue, 10 Nov 2009 23:00:00 UTC
RFC3339Nano:  2009-11-10T23:00:00Z
RubyDate:  Tue Nov 10 23:00:00 +0000 2009
