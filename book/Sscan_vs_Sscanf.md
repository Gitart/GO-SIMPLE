## Sscan против Sscanf против Sscanln из пакета FMT
Sscan сканирует строку аргумента, сохраняя последовательные значения, разделенные пробелом, в последовательные аргументы. Новые строки считаются пробелом.
Sscanf сканирует строку аргумента, сохраняя последовательные значения, разделенные пробелом, в последовательные аргументы в соответствии с форматом.
Sscanln похож на Sscan, но останавливает сканирование на новой строке, и после последнего элемента должен быть символ новой строки или EOF.

```golang
package main
 
import (
    "fmt"      
)
 
func main(){    
    var X int
    var Y int
 
    fmt.Printf("\nIntital X: %d, Y: %d", X, Y)
 
    fmt.Sscan("100\n200", &X, &Y)
    fmt.Printf("\nSscan X: %d, Y: %d", X, Y)
 
    fmt.Sscanf("(10, 20)", "(%d, %d)", &X, &Y)
    fmt.Printf("\nSscanf X: %d, Y: %d", X, Y)
     
    fmt.Sscanln("50\n50", &X, &Y)
    fmt.Printf("\nSscanln X: %d, Y: %d", X, Y)
}
```
