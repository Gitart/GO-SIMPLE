# Get Hours, Days, Minutes and Seconds difference between two dates \[Future and Past\]

package main

import (
         "fmt"
         "time"
)

func main() {
loc, \_ := time.LoadLocation("UTC")
now := time.Now().In(loc)
fmt.Println("\\nToday : ", loc, " Time : ", now)

pastDate := time.Date(2015, time.May, 21, 23, 10, 52, 211, time.UTC)
fmt.Println("\\n\\nPast  : ", loc, " Time : ", pastDate) //
fmt.Printf("###############################################################\\n")
diff := now.Sub(pastDate)

hrs := int(diff.Hours())
fmt.Printf("Diffrence in Hours : %d Hours\\n", hrs)

mins := int(diff.Minutes())
fmt.Printf("Diffrence in Minutes : %d Minutes\\n", mins)

second := int(diff.Seconds())
fmt.Printf("Diffrence in Seconds : %d Seconds\\n", second)

days := int(diff.Hours() / 24)
fmt.Printf("Diffrence in days : %d days\\n", days)

fmt.Printf("###############################################################\\n\\n\\n")

futureDate := time.Date(2019, time.May, 21, 23, 10, 52, 211, time.UTC)
fmt.Println("Future  : ", loc, " Time : ", futureDate) //
fmt.Printf("###############################################################\\n")
diff = futureDate.Sub(now)

hrs = int(diff.Hours())
fmt.Printf("Diffrence in Hours : %d Hours\\n", hrs)

mins = int(diff.Minutes())
fmt.Printf("Diffrence in Minutes : %d Minutes\\n", mins)

second = int(diff.Seconds())
fmt.Printf("Diffrence in Seconds : %d Seconds\\n", second)

days = int(diff.Hours() / 24)
fmt.Printf("Diffrence in days : %d days\\n", days)

}

C:\\golang\\time>go run t3.go

Today : UTC Time : 2017\-08\-27 05:15:53.7106215 +0000 UTC

Past : UTC Time : 2015\-05\-21 23:10:52.000000211 +0000 UTC
###############################################################
Diffrence in Hours : 19878 Hours
Diffrence in Minutes : 1192685 Minutes
Diffrence in Seconds : 71561101 Seconds
Diffrence in days : 828 days
###############################################################

Future : UTC Time : 2019\-05\-21 23:10:52.000000211 +0000 UTC
###############################################################
Diffrence in Hours : 15185 Hours
Diffrence in Minutes : 911154 Minutes
Diffrence in Seconds : 54669298 Seconds
Diffrence in days : 632 days

C:\\golang\\time>
