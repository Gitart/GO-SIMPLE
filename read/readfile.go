package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    file, err := os.Open("/path/to/file.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}



// Readlines From String
func StringToLines(s string) []string {
      var lines []string

      scanner := bufio.NewScanner(strings.NewReader(s))
      for scanner.Scan() {
              lines = append(lines, scanner.Text())
      }

      if err := scanner.Err(); err != nil {
              fmt.Fprintln(os.Stderr, "reading standard input:", err)
      }

      return lines
}
