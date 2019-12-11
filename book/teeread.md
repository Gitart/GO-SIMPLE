# Как использовать TeeReader из IO Package в Golang?
TeeReader возвращает Reader, который пишет в w то, что он читает из r. Все чтения из r, выполненные через него, сопоставляются с соответствующими записями в w.

```golang
package main
 
import (
    "bytes"
    "io"
    "fmt"
    "strings"
)
 
func main() {
    testString := strings.NewReader("Jobs, Code, Videos and News for Go hackers.")
    var bufferRead bytes.Buffer
    example := io.TeeReader(testString, &bufferRead)
    readerMap := make([]byte, testString.Len())
    length, err := example.Read(readerMap)
    fmt.Printf("\nBufferRead: %s", &bufferRead)
    fmt.Printf("\nRead: %s", readerMap)
    fmt.Printf("\nLength: %d, Error:%v", length, err)
}
```
