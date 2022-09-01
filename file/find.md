# Go find file

last modified January 26, 2022

Go find file tutorial shows how to find files with filepath.Walk.

The `filepath` package implements utility routines for manipulating filename paths.

The `filepath.Walk` walks the file tree, calling the specified function for each file or directory in the tree, including root. The function is recursively walking all subdirectories.

## Go find text files

In the first example, we search for text files.

find\_text\_files.go

package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
)

func main() {

    var files \[\]string

    root := "/home/janbodnar/Documents"

    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

        if err != nil {

            fmt.Println(err)
            return nil
        }

        if !info.IsDir() && filepath.Ext(path) == ".txt" {
            files = append(files, path)
        }

        return nil
    })

    if err != nil {
        log.Fatal(err)
    }

    for \_, file := range files {
        fmt.Println(file)
    }
}

In the code example, we search for files with `.txt` extension.

var files \[\]string

The matching files are stored in the `files` slice.

root := "/home/janbodnar/Documents"

This is the root directory where we start searching.

err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

The first parameter of the `filepath.Walk` is the root directory. The second parameter is the walk function; the function called by `filepath.Walk` to visit each each file or directory.

if err != nil {

    fmt.Println(err)
    return nil
}

Print the error if there is one, but continue searching elsewhere.

if !info.IsDir() && filepath.Ext(path) == ".txt" {
    files = append(files, path)
}

We append the file to the `files` slice if the file is not a directory and it has the `.txt` extension.

for \_, file := range files {
    fmt.Println(file)
}

Finally, we go over the `files` slice and print all matching files to the console.

## Go find file by size

In the next example, we find files by their size.

findbysize.go

package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
)

var files \[\]string

func VisitFile(path string, info os.FileInfo, err error) error {

    if err != nil {

        fmt.Println(err)
        return nil
    }

    file\_size := info.Size()

    if !info.IsDir() && file\_size > 1024\*1024 {

        files = append(files, path)
    }

    return nil
}

func main() {

    root := "/home/janbodnar/Documents"

    err := filepath.Walk(root, VisitFile)

    if err != nil {
        log.Fatal(err)
    }

    for \_, file := range files {
        fmt.Println(file)
    }
}

The example lists files that have size greater than 1 MB.

file\_size := info.Size()

With the `Size` method, we determine the size of the file in bytes.

if !info.IsDir() && file\_size > 1024\*1024 {

    files = append(files, path)
}

If the file is not a directory and its size is greater than 1 MB, we add the file to the `files` slice.

## Go find files by modification time

In the next example, we find files by modification time.

findmodif.go

package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "time"
)

func FindFilesAfter(dir string, t time.Time) (files \[\]string, err error) {

    err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

        if err != nil {

            fmt.Println(err)
            return nil
        }

        if !info.IsDir() && filepath.Ext(path) == ".txt" && info.ModTime().After(t) {
            files = append(files, path)
        }

        return nil
    })

    return
}

func main() {

    root := "/home/janbodnar/Documents"

    t, err := time.Parse("2006-01-02T15:04:05-07:00", "2021-05-01T00:00:00+00:00")

    if err != nil {
        log.Fatal(err)
    }

    files, err := FindFilesAfter(root, t)

    if err != nil {

        log.Fatal(err)
    }

    for \_, file := range files {

        fmt.Println(file)
    }
}

We find all text files that were modified after May 1, 2021.

if !info.IsDir() && filepath.Ext(path) == ".txt" && info.ModTime().After(t) {
    files = append(files, path)
}

The `ModTime` gives the last modification time of the file. With the `After` function, we check whether the time instant is after the given time.

## Go find file by regex pattern

In the following example, we search for file names based on a regular expression.

findregex.go

package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "regexp"
)

var files \[\]string

func VisitFile(path string, info os.FileInfo, err error) error {

    if err != nil {

        fmt.Println(err)
        return nil
    }

    if info.IsDir() || filepath.Ext(path) != ".txt" {

        return nil
    }

    reg, err2 := regexp.Compile("^\[la\]")

    if err2 != nil {

        return err2
    }

    if reg.MatchString(info.Name()) {

        files = append(files, path)
    }

    return nil
}

func main() {

    err := filepath.Walk("/home/janbodnar/Documents", VisitFile)

    if err != nil {

        log.Fatal(err)
    }

    for \_, file := range files {

        fmt.Println(file)
    }
}

In the code example, we search for text files that begin with letter l or letter a.

reg, err2 := regexp.Compile("^\[la\]")

We compile a regular expression with `regexp.Compile`.

if reg.MatchString(info.Name()) {

    files = append(files, path)
}

If the regular expression matches the file name, we add the file to the `files` slice.
