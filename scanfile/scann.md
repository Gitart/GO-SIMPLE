# Scan file sample

## This is the cleanest way to read from a Reader line by line.

```golang
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
```

## Other sample for usage


```golang
package main

import (
    "bufio"
    "bytes"
    "fmt"
    "io"
    "os"
)

func readFileWithReadString(fn string) (err error) {
    fmt.Println("readFileWithReadString")

    file, err := os.Open(fn)
    defer file.Close()

    if err != nil {
        return err
    }

    // Start reading from the file with a reader.
    reader := bufio.NewReader(file)

    var line string
    for {
        line, err = reader.ReadString('\n')

        fmt.Printf(" > Read %d characters\n", len(line))

        // Process the line here.
        fmt.Println(" > > " + limitLength(line, 50))

        if err != nil {
            break
        }
    }

    if err != io.EOF {
        fmt.Printf(" > Failed!: %v\n", err)
    }

    return
}

func readFileWithScanner(fn string) (err error) {
    fmt.Println("readFileWithScanner - this will fail!")

    // Don't use this, it doesn't work with long lines...

    file, err := os.Open(fn)
    defer file.Close()

    if err != nil {
        return err
    }

    // Start reading from the file using a scanner.
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        fmt.Printf(" > Read %d characters\n", len(line))

        // Process the line here.
        fmt.Println(" > > " + limitLength(line, 50))
    }

    if scanner.Err() != nil {
        fmt.Printf(" > Failed!: %v\n", scanner.Err())
    }

    return
}

func readFileWithReadLine(fn string) (err error) {
    fmt.Println("readFileWithReadLine")

    file, err := os.Open(fn)
    defer file.Close()

    if err != nil {
        return err
    }

    // Start reading from the file with a reader.
    reader := bufio.NewReader(file)

    for {
        var buffer bytes.Buffer

        var l []byte
        var isPrefix bool
        for {
            l, isPrefix, err = reader.ReadLine()
            buffer.Write(l)

            // If we've reached the end of the line, stop reading.
            if !isPrefix {
                break
            }

            // If we're just at the EOF, break
            if err != nil {
                break
            }
        }

        if err == io.EOF {
            break
        }

        line := buffer.String()

        fmt.Printf(" > Read %d characters\n", len(line))

        // Process the line here.
        fmt.Println(" > > " + limitLength(line, 50))
    }

    if err != io.EOF {
        fmt.Printf(" > Failed!: %v\n", err)
    }

    return
}

func main() {
    testLongLines()
    testLinesThatDoNotFinishWithALinebreak()
}

func testLongLines() {
    fmt.Println("Long lines")
    fmt.Println()

    createFileWithLongLine("longline.txt")
    readFileWithReadString("longline.txt")
    fmt.Println()
    readFileWithScanner("longline.txt")
    fmt.Println()
    readFileWithReadLine("longline.txt")
    fmt.Println()
}

func testLinesThatDoNotFinishWithALinebreak() {
    fmt.Println("No linebreak")
    fmt.Println()

    createFileThatDoesNotEndWithALineBreak("nolinebreak.txt")
    readFileWithReadString("nolinebreak.txt")
    fmt.Println()
    readFileWithScanner("nolinebreak.txt")
    fmt.Println()
    readFileWithReadLine("nolinebreak.txt")
    fmt.Println()
}

func createFileThatDoesNotEndWithALineBreak(fn string) (err error) {
    file, err := os.Create(fn)
    defer file.Close()

    if err != nil {
        return err
    }

    w := bufio.NewWriter(file)
    w.WriteString("Does not end with linebreak.")
    w.Flush()

    return
}

func createFileWithLongLine(fn string) (err error) {
    file, err := os.Create(fn)
    defer file.Close()

    if err != nil {
        return err
    }

    w := bufio.NewWriter(file)

    fs := 1024 * 1024 * 4 // 4MB

    // Create a 4MB long line consisting of the letter a.
    for i := 0; i < fs; i++ {
        w.WriteRune('a')
    }

    // Terminate the line with a break.
    w.WriteRune('\n')

    // Put in a second line, which doesn't have a linebreak.
    w.WriteString("Second line.")

    w.Flush()

    return
}

func limitLength(s string, length int) string {
    if len(s) < length {
        return s
    }

    return s[:length]
}
```
