## Fscan, Fscanf и Fscanln из пакета FMT
---
Fscan сканирует текст, считанный из r, сохраняя последовательные значения, разделенные пробелом, в последовательные аргументы.
Fscanf сканирует текст, считанный из r, сохраняя последовательные разделенные пробелами значения в последовательных аргументах, как определено форматом.
Fscanln похож на Fscan, но останавливает сканирование на новой строке и после последнего элемента должен быть символ новой строки или EOF.

Для выполнения этой программы необходимо создать текстовый файл с именем testfile.txt и содержимым, приведенным в разделе комментариев к программе ниже.

/* testfile.txt content:
255 5.5 true Australia
200 9.3 false Germany
Paris 5 true 5.5
*/

```golang
package main
 
import (
    "fmt"
    "os"
)
 
var (
    L int32
    M float32
    N bool
    O string
)
 
func main() {
    file, err := os.Open("testfile.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()
     
    fmt.Fscan(file, &L, &M, &N, &O)     
    fmt.Printf("\nL M N O: %v %v %v %v", L,M,N,O)
     
    fmt.Fscanln(file, &L, &M, &N, &O)
    fmt.Printf("\nL M N O: %v %v %v %v", L,M,N,O)
     
    fmt.Fscanf(file, "%s %d %t %f",&O, &L, &N, &M)
    fmt.Printf("\nO L N M: %v %v %v %v", O,L,N,M)
}
```
