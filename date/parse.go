package main

import (
	"fmt"
	"time"
)

func main() {
	//formatting works in a one-two-three... pattern
	//in the date portion, 01 is the month (American format) and 02 is the day and 06 is the year
	//in the time portion, 03 or 15 is the hour and 04 is the minutes while 05 is the seconds
	//at the end the UTC offset will always begin with - (negative) and 0700 [-0700]

	layout := "Mon 01/02/06 03:04:05PM -07:00"
	str := "Mon 03/27/17 02:21:00.215PM +02:00"

	printTime := func() {
		t, err := time.Parse(layout, str)
		if err != nil {
			fmt.Println("ERROR:", err.Error())
		}

		out := "02/01/2006 03:04"
		fmt.Println("DateTime:", t.Format(out))
	}
	printTime()

	layout = "Jan 02 2006"
	str = "Mar 27 2017"
	printTime()

	utcTime := time.Now().UTC()
	fmt.Println("UTC:", utcTime)

	localTime := utcTime.Local()
	fmt.Println("Local:", localTime)
}
