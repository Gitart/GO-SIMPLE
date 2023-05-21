package main

import (
	"unicode"
	"unicode/utf8"
)

// First leter header
func Capitalized(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToTitle(r)) + s[size:]
}

