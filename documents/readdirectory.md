## Read directory and file

```golang
package main
 
import (
    "fmt"
    "os"
    "path/filepath"
)
 
func VisitFile(fp string, fi os.FileInfo, err error) error {
    if err != nil {
        fmt.Println(err) // can't walk here,
        return nil       // but continue walking elsewhere
    }
    if !!fi.IsDir() {
        fmt.Println("directory:"+fp)
        return nil // not a file.
    }
    fmt.Println("file:"+fp)
    return nil
}
 
func main() {
    //specify directory below or walk through /
    filepath.Walk("/", VisitFile)
}
```

## Output

```text
directory:/
directory:/dev
file:/dev/null
file:/dev/random
file:/dev/urandom
file:/dev/zero
directory:/etc
file:/etc/group
file:/etc/hosts
file:/etc/passwd
file:/etc/resolv.conf
directory:/tmp
directory:/usr
directory:/usr/local
directory:/usr/local/go
directory:/usr/local/go/lib
directory:/usr/local/go/lib/time
file:/usr/local/go/lib/time/zoneinfo.zip
```


## Read in array

```golang
package main
 
import (
    "fmt"
    "os"
    "path/filepath"
)

func dir(thepath string) []string {

  var files []string

  filepath.Walk(thepath, func(path string, _ os.FileInfo, _ error) error {
    //fmt.Println(path)
    files = append(files, path)
    return nil
  })

  return files

}

func main() {

  path := "/"
  fmt.Println(dir(path))

}

```


## Current directory

```golang
package main

import (
    "fmt"
    "os"
)

func main() {
    pwd, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println(pwd)
}
```


## Write wiyh buf

```golang
package main

import (
    "bufio"
    "fmt"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    //specify folder
    path := "/"
    folder := path+"test"

    //create it with permissions
    os.Mkdir(folder, 0700)
	
    //define filename
    file := folder+"/1.html"

    //open a file for writing.
    f, err := os.Create(file)
    check(err)

    // `bufio` provides buffered writers and readers
    w := bufio.NewWriter(f)
    n4, err := w.WriteString("does<b>this</b>work")
    fmt.Printf("wrote %d bytes\n", n4)

    // Use `Flush` to ensure all buffered operations have been applied to the underlying writer.
    w.Flush()

}
```


