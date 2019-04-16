package main

import (
	"bytes"
	"fmt"
)

func main() {

	var (
		fp, sp [16]byte
		key    = []byte("0123456789abcdef:fedcba9876543210")
	)

	i := bytes.Index(key, []byte(":"))
	if i < 0 {
		fmt.Println("error ':' not found in key")
		return
	}
	copy(fp[:], key[:i])
	copy(sp[:], key[i+1:])

	fmt.Println(string(fp[:]))
	fmt.Println(string(sp[:]))
}
