package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegex := regexp.MustCompile("abc|xyz|123|yyy")

	match := sampleRegex.Match([]byte("abc"))
	fmt.Println(match)

	match = sampleRegex.Match([]byte("xyz"))
	fmt.Println(match)

	match = sampleRegex.Match([]byte("123"))
	fmt.Println(match)

	match = sampleRegex.Match([]byte("a2bc2xy2z1223yy2yy"))
	fmt.Println(match)

	match = sampleRegex.Match([]byte("abd"))
	fmt.Println(match)
}


/*
true
true
true
false
false
*/
