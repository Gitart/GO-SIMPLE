# Month

```golang
package main

 import (
 	"fmt"
 	"time"
 )

 func main() {

 	now := time.Now()

 	fmt.Println("Now is : ", now)

 	// get month name from now
 	_, month, _ := now.Date()

 	fmt.Println("and the month is : ", month)

 }
 ```
 
