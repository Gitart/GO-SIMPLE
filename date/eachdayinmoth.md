## Loop each day of the current month example

### Problem:
You need to get the current month, find out the number of days in the month programmatically and loop   
through each day to perform certain tasks such as restoring archived data from backups or interpolate    
data base on historical trend. How to do that?

##  Solution:
Find out the current month and the number of days in the month with time.Now() and time.Date()    
functions. Once you have determined the number of days in a the current calendar month, use a for    
loop to loop from day 1 to the end day of the month. 

```golang
 package main

  import (
          "fmt"
          "time"
  )

  func main() {

          // get the current month
          year, month, _ := time.Now().Date()

          fmt.Printf("Current month: [%v]\n", month)

          // get the number of days of the current month
          t := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
          fmt.Printf("Total number of days in [%v], [%v] is [%v]\n", year, month, t.Day())

          // loop each day of the month
          for day := 1; day <= t.Day(); day++ {
                  // do whatever you want here ...
                  fmt.Println(day, month, year)
          }
  }
  
  ```
Sample output:

```
Current month: [February]
Total number of days in [2016], [February] is [29]
1 February 2016
...
29 February 2016
```
