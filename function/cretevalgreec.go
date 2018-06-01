package main

import (
	"fmt"
)

// http://stackoverflow.com/a/27457144/10278

func romanNumeralDict() func(int) string {
    // innerMap is captured in the closure returned below
	innerMap := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	return func(key int) string {
		return innerMap[key]
	}
}

func main() {
	fmt.Println(romanNumeralDict()(10))
	fmt.Println(romanNumeralDict()(100))

	dict := romanNumeralDict()
	fmt.Println(dict(400))
}
