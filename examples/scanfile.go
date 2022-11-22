//File System Scanning in Golang
package main

import (
    "path/filepath"
    "os"
    "flag"
)

type visitor int

func (v visitor) VisitDir(path string, f *os.FileInfo) bool {
    println(path)
    return true
} 

func (v visitor) VisitFile(path string, f *os.FileInfo) {
    println(path)
}

func main() {
    root := flag.Arg(0)
    filepath.Walk(root, visitor(0), nil)
}
	
package main

import (
    "path/filepath"
    "fmt"
    "os"
)

func main() {
    dirname := "." + string(filepath.Separator)
    d, err := os.Open(dirname)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fi, err := d.Readdir(-1)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    for _, fi := range fi {
        if fi.IsRegular() {
            fmt.Println(fi.Name, fi.Size, "bytes")
        }
    }
}

//Output:
//Makefile 92 bytes
//charset.go 2087 bytes
//reverse.go 234 bytes
//wordcnt.go 2387 bytes
//readln.txt 45 bytes
//stat.go 166 bytes
//readln.go 1289 bytes
