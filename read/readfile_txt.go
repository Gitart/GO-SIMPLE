package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Data struct {
    AString string
    AFloat  float64
    AnInt   int64
}

func ParseLine(line string) (*Data, error) {
    data := new(Data)
    var err error
    text := strings.TrimRight(line, " ")
    i := strings.LastIndex(text, " ")
    i++
    text = text[i:]
    data.AnInt, err = strconv.ParseInt(text, 10, 64)
    if err != nil {
        return nil, err
    }
    line = line[:i]
    text = strings.TrimRight(line, " ")
    i = strings.LastIndex(text, " ")
    i++
    text = text[i:]
    data.AFloat, err = strconv.ParseFloat(text, 64)
    if err != nil {
        return nil, err
    }
    line = line[:i]
    data.AString = line
    return data, nil
}

func main() {
    f, err := os.Open("data.txt")
    if err != nil {
        fmt.Fprintln(os.Stderr, "opening input:", err)
        return
    }
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        data, err := ParseLine(line)
        if err != nil {
            fmt.Fprintln(os.Stderr, "reading input:", err, ":", line)
            continue
        }
        fmt.Println(*data)
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading input:", err)
    }
}
