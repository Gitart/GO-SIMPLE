# Go write file

Go write file tutorial shows how to write to files in Golang. Learn how to read files in Go in [Go read file](/golang/readfile/).

To write to files in Go, we use the `os`, `ioutil`, and `fmt` packages.

```
func (f *File) WriteString(s string) (n int, err error)
```

The functions that we use typically return the number of bytes written and an error, if any.

```
$ go version
go version go1.18.1 linux/amd64
```

We use Go version 1.18.

## Go write to file with File.WriteString

The `File.WriteString` function writes the contents of a string to a file.

write\_string.go

```go
package main

import (
    "fmt"
    "log"
    "os"
)

func main() {

    f, err := os.Create("data.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    _, err2 := f.WriteString("old falcon\n")

    if err2 != nil {
        log.Fatal(err2)
    }

    fmt.Println("done")
}
```

The example writes a string to a text file.

```
f, err := os.Create("data.txt")
```

The `os.Create` creates or truncates the named file. If the file already exists, it is truncated.

```
_, err2 := f.WriteString("old falcon\n")
```

The string is written to the created file with `WriteString`.

```
$ go run write_string.go
done
$ cat data.txt
old falcon
```

We run the program and show the contents of the created file.

## Go write to file with File.Write/File.WriteAt

The `File.Write` writes n bytes to a file. The `File.WriteAt` writes n bytes to a file starting at the specified byte offset.

write\_bytes.go

```go
package main

import (
    "fmt"
    "log"
    "os"
)

func main() {

    f, err := os.Create("data.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    val := "old falcon"
    data := []byte(val)

    _, err2 := f.Write(data)

    if err2 != nil {
        log.Fatal(err2)
    }

    val2 := " and red fox\n"
    data2 := []byte(val2)

    var idx int64 = int64(len(data))

    _, err3 := f.WriteAt(data2, idx)

    if err3 != nil {
        log.Fatal(err3)
    }

    fmt.Println("done")
}
```

First, we write "old falcon" to the file with `File.Write`, then we write the " and red fox" next to it with `File.WriteAt`.

```
val := "old falcon"
data := []byte(val)
```

We transform the string to bytes.

```
_, err2 := f.Write(data)
```

We write the bytes with `Write`.

```
var idx int64 = int64(len(data))
```

We calculate the length of the previously written string.

```
_, err3 := f.WriteAt(data2, idx)
```

We write the string and the specified index with `WriteAt`.

```
$ go run write_bytes.go
done
$ cat data.txt
old falcon and red fox
```

We run the program and check the file.

## Go write to file with ioutil.WriteFile

The `outil.WriteFile` writes data to the specified file. This is a higher-level convenience function. The opening and closing of the file is handled for us.

write\_file.go

```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {

    val := "old falcon\n"
    data := []byte(val)

    err := ioutil.WriteFile("data.txt", data, 0)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("done")
}
```

The example writes a string to a file with `ioutil.WriteFile`.

## Go write a slice to file

The following example writes a slice of words to a file.

write\_slice.go

```go
package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    f, err := os.Create("data.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    words := []string{"sky", "falcon", "rock", "hawk"}

    for _, word := range words {

        _, err := f.WriteString(word + "\n")

        if err != nil {
            log.Fatal(err)
        }
    }

    fmt.Println("done")
}
```

We define a slice of words. In a for loop, we go through the slice and write all words to the file.

```
$ go run write_slice.go 
done
$ cat data.txt 
sky
falcon
rock
hawk
```

We run the program and show the contents of the created file.

## Go write formatted string

The `fmt.Fprintln` formats the string using the default formats for its operands and writes it to the writer. Spaces are always added between operands and a newline is appended.

write\_formatted.go

```go
package main

import (
    "fmt"
    "log"
    "os"
)

func main() {

    f, err := os.Create("data.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    const name, age = "Johne Doe", 34

    n, err := fmt.Fprintln(f, name, "is", age, "years old.")

    if err != nil {

        log.Fatal(err)
    }

    fmt.Println(n, "bytes written")
    fmt.Println("done")
}
```

In the code example, we have two variables: `name` and `age`. We build a string from the variables and write it to the created file. We also print the number of bytes that were written.

```
$ go run write_formatted.go 
27 bytes written
done
$ cat data.txt 
Johne Doe is 34 years old.
```

We run the program and show the contents of the created file.

In this article, we have covered writing to files in Go.
