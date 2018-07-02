package main

import (
	"fmt"
)

func main() {
	fmt.Println(nameStart('Z'))

}

// *************************************************************************
// nameStart returns whether c can be the first character of an identifier
// (not counting an initial hyphen, or an escape sequence).
// *************************************************************************
func nameStart(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_' || c > 127
}

func hexDigit(c byte) bool {
	return '0' <= c && c <= '9' || 'a' <= c && c <= 'f' || 'A' <= c && c <= 'F'
}

// nameStart returns whether c can be the first character of an identifier
// (not counting an initial hyphen, or an escape sequence).
func nameStart(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_' || c > 127
}

// nameChar returns whether c can be a character within an identifier
// (not counting an escape sequence).
func nameChar(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_' || c > 127 ||
		c == '-' || '0' <= c && c <= '9'
}
