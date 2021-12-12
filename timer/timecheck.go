package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	time.Sleep(time.Millisecond * 1223)

	secs := time.Since(start).Milliseconds()
	out := fmt.Sprintf("%d ", secs)
	fmt.Println(out)

}
