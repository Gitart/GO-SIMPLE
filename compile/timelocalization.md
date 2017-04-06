## Get local time and equivalent time in different time zone

```golang
package main

 import (
         "fmt"
         "time"
 )

 func main() {

         t := time.Now()

         fmt.Println("Location : ", t.Location(), " Time : ", t)

         // get the list of available time zones
         location, err := time.LoadLocation("America/New_York")

         if err != nil {
                 fmt.Println(err)
         }

         fmt.Println("Location : ", location, " Time : ", t.In(location))
 }
 ```
 
Sample output :

```
Location : Local Time : 2015-07-30 16:39:28.158021072 +0800 MYT
Location : America/New_York Time : 2015-07-30 04:39:28.158021072 -0400 EDT
```
