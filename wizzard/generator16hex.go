// https://play.golang.org/p/2KNFA7SLyqb

package main

import "fmt"

func main() {
	val := "Y032039-47848-DSKJDK"
	
	hex := fmt.Sprintf("%x",val)
	fmt.Printf(hex)
}
