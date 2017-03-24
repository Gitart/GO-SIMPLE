## Запись в файл

```golang
package main

import (
    "io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // To start, here's how to dump a string (or just bytes) into a file.
    d1 := []byte("does<b>this</b>work")
    err := ioutil.WriteFile("/root/go/src/webapp/1.html", d1, 0644)
    check(err)

}
```
