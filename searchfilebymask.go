package main

import (
  "fmt"
  "path/filepath"
)

func main() {
  files, _ := filepath.Glob("*.txt")
  fmt.Printf("%q\n", files)
}

//$ go run example.go
//["1.txt" "2.txt"]
