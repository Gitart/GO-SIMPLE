Переменные пакета времени выполнения
Ниже приведена короткая программа для отображения компилятора, номера процессора, языковой версии, GOOS, GOARCH и GOROOT во время выполнения.

```
package main
 
import (
    "fmt"
    "runtime"
)
 
func main() {
    fmt.Printf("\nGOOS:%s", runtime.GOOS)
    fmt.Printf("\nGOARCH:%s", runtime.GOARCH)
    fmt.Printf("\nGOROOT:%s", runtime.GOROOT())
    fmt.Printf("\nCompiler:%s", runtime.Compiler)
    fmt.Printf("\nNo. of CPU:%d", runtime.NumCPU()) 
}
```
