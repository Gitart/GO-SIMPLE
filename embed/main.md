# Go embed

last modified February 15, 2022

Go embed tutorial shows how to access embedded files from within a running Go program.

The `embed` package allows to access static files such as images and HTML files from within a running Go binary. It was introduced in Go 1.16.

## Go embed a text file

In the first example, we embed a text file into a slice of bytes.

data/words.txt

sky
blue
rock
water
array
karma
falcon

We have a few words in the text file.

main.go

package main

import (
    "bytes"
    \_ "embed"
    "fmt"
)

var (
    //go:embed data/words.txt
    data \[\]byte
)

func main() {

    fmt.Println(string(data))

    fmt.Println("----------------------")

    words := bytes.Split(data, \[\]byte{'\\n'})

    for \_, w := range words {

        fmt.Println(string(w))
    }
}

We embed a text file into the program and print the data.

var (
    //go:embed data/words.txt
    data \[\]byte
)

Embedding is done by using the `//go:embed` directive above the variable declaration.

$ go build
$ ./txtfile.exe
sky
blue
rock
water
array
karma
falcon
----------------------
sky
blue
rock
water
array
karma
falcon

## Go embed multiple files

In the following example, we embed two text files.

data/langs.txt

Perl
Raku
F#
Clojure
Go
C#

This is the `langs.txt` file.

data/words.txt

sky
blue
rock
falcon
war
tree
storm
cup

This is the `words.txt` file.

main.go

package main

import (
    "embed"
    "fmt"
)

//go:embed data/\*
var f embed.FS

func main() {

    langs, \_ := f.ReadFile("data/langs.txt")
    fmt.Println(string(langs))

    words, \_ := f.ReadFile("data/words.txt")
    fmt.Println(string(words))
}

We embed two text files and print their contents.

//go:embed data/\*
var f embed.FS

With the \* wildcard character, we embed all files within the data directory.

$ go build
$ ./files.exe
Perl
Raku
F#
Clojure
Go
C#

sky
blue
rock
falcon
war
tree
storm
cup

## Go embed static files

In the following example, we embed static files into a binary of a web application.

public/index.html

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Home</title>
</head>
<body>
    <p>
        Home page
    </p>
</body>
</html>

This is the `index.html` file.

public/about.html

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>About</title>
</head>
<body>
    <p>
        About page
    </p>
</body>
</html>

This is the `about.html` file.

main.go

package main

import (
    "embed"
    "io/fs"
    "net/http"
)

//go:embed public
var content embed.FS

func handler() http.Handler {

    fsys := fs.FS(content)
    html, \_ := fs.Sub(fsys, "public")

    return http.FileServer(http.FS(html))
}

func main() {

    mux := http.NewServeMux()
    mux.Handle("/", handler())

    http.ListenAndServe(":8080", mux)
}

The code example runs a server which serves two static files. We embed a whole directory.

//go:embed public
var content embed.FS

The `embed.FS` allows to embed a tree of files.

func handler() http.Handler {

    fsys := fs.FS(content)
    html, \_ := fs.Sub(fsys, "public")

    return http.FileServer(http.FS(html))
}

The handler serves static files from the `public` directory. In Go, the `http.FileServer` is used to serve static content.

In this tutorial, we have showed how to access static files from within a running Go program using the `embed` package.
