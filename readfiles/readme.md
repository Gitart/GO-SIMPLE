# Go read file


Go read file tutorial shows how to read files in Golang. We read text and binary files. Learn how to write to files in Go in [Go write file](/golang/writefile/).

```
$ go version
go version go1.18.1 linux/amd64
```

To read files in Go, we use the `os`, `ioutil`, `io`, and `bufio` packages.

thermopylae.txt

```
The Battle of Thermopylae was fought between an alliance of Greek city-states,
led by King Leonidas of Sparta, and the Persian Empire of Xerxes I over the
course of three days, during the second Persian invasion of Greece.
```

We use this text file in some examples.

# Advertisements
## Go read file into string

The `ioutil.ReafFile` function reads the whole file into a string. This function is convenient but should not be used with very large files.

read\_file.go

```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {

    content, err := ioutil.ReadFile("thermopylae.txt")

     if err != nil {
          log.Fatal(err)
     }

    fmt.Println(string(content))
}
```

The example reads the whole file and prints it to the console.

```
$ go run read_file.go
The Battle of Thermopylae was fought between an alliance of Greek city-states,
led by King Leonidas of Sparta, and the Persian Empire of Xerxes I over the
course of three days, during the second Persian invasion of Greece.
```

## Go read file line by line

The `Scanner` provides a convenient interface for reading data such as a file of newline-delimited lines of text. It reads data by tokens; the `Split` function defines the token. By default, the function breaks the data into lines with line-termination stripped.

read\_line\_by\_line.go

```go
package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {

    f, err := os.Open("thermopylae.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {

        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
```

The example reads the file line by line. Each line is printed.

```
f, err := os.Open("thermopylae.txt")
```

The `Open` function opens the file for reading.

```
defer f.Close()
```

The file descriptor is closed at the end of the `main` function.

```
scanner := bufio.NewScanner(f)
```

A new scanner is created.

```
for scanner.Scan() {

     fmt.Println(scanner.Text())
}
```

The `Scan` advances the Scanner to the next token, which will then be available through the `Bytes` or `Text` function.

## Go read file by words

The default split function of a scanner is `ScanLines`. With `SplitWords`, we split the content by words.

read\_by\_word.go

```go
package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {

    f, err := os.Open("thermopylae.txt")

    if err != nil {
        fmt.Println(err)
     }

    defer f.Close()

    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanWords)

    for scanner.Scan() {

        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}
```

The example reads the file word by word.

```
$ go run read_by_word.go
The
Battle
of
Thermopylae
was
fought
between
...
```

## Go read file in chunks

We can read files in chunks of data.


```go
package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "io"
)

func main() {

    f, err := os.Open("thermopylae.txt")

    if err != nil {
         log.Fatal(err)
    }

    defer f.Close()

    reader := bufio.NewReader(f)
    buf := make([]byte, 16)

    for {
        n, err := reader.Read(buf)

        if err != nil {

           if err != io.EOF {

               log.Fatal(err)
           }

           break
        } 

        fmt.Print(string(buf[0:n]))
    }

    fmt.Println()
}
```

The example reads the file by small 16 byte portions.

```
buf := make([]byte, 16)
```

We define an array of 16 bytes.

```
for {
    n, err := reader.Read(buf)

    if err != nil {

         if err != io.EOF {

             log.Fatal(err)
         }

         break
    }

    fmt.Print(string(buf[0:n]))
}
```

In the for loop, we read data into the buffer with `Read`, and print the array buffer to the console with `Print`.

## Go read binary file

The `hex` package implements hexadecimal encoding and decoding.

read\_binary\_file.go

```go
package main

import (  
    "bufio"
    "encoding/hex"
    "fmt"
    "log"
    "os"
    "io"
)

func main() {  

    f, err := os.Open("sid.jpg")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    reader := bufio.NewReader(f)
    buf := make([]byte, 256)

    for {
        _, err := reader.Read(buf)

        if err != nil {
            if err != io.EOF {
                fmt.Println(err)
            }
            break
        }
        
        fmt.Printf("%s", hex.Dump(buf))
    }
}
```

In the code example, we read an image and print it in hexadecimal format.

```
fmt.Printf("%s", hex.Dump(buf))
```

The `Dump` returns a string that contains a hex dump of the given data.

```
$ go run read_binary_file.go
00000000  ff d8 ff e0 00 10 4a 46  49 46 00 01 01 00 00 01  |......JFIF......|
00000010  00 01 00 00 ff e1 00 2f  45 78 69 66 00 00 49 49  |......./Exif..II|
00000020  2a 00 08 00 00 00 01 00  0e 01 02 00 0d 00 00 00  |*...............|
00000030  1a 00 00 00 00 00 00 00  6b 69 6e 6f 70 6f 69 73  |........kinopois|
00000040  6b 2e 72 75 00 ff fe 00  3b 43 52 45 41 54 4f 52  |k.ru....;CREATOR|
00000050  3a 20 67 64 2d 6a 70 65  67 20 76 31 2e 30 20 28  |: gd-jpeg v1.0 (|
00000060  75 73 69 6e 67 20 49 4a  47 20 4a 50 45 47 20 76  |using IJG JPEG v|
00000070  38 30 29 2c 20 71 75 61  6c 69 74 79 20 3d 20 39  |80), quality = 9|
00000080  31 0a ff db 00 43 00 03  02 02 03 02 02 03 03 02  |1....C..........|
00000090  03 03 03 03 03 04 07 05  04 04 04 04 09 06 07 05  |................|
000000a0  07 0a 09 0b 0b 0a 09 0a  0a 0c 0d 11 0e 0c 0c 10  |................|
000000b0  0c 0a 0a 0e 14 0f 10 11  12 13 13 13 0b 0e 14 16  |................|
```
