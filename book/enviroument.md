## Как установить, получить и перечислить переменные среды?
Используйте os.Environ, чтобы перечислить все ключи / значения среды. os.Environ () возвращает фрагмент строки в форме KEY = value. Затем мы используем strings.Split, чтобы получить ключ и значение. Затем распечатайте все ключи и значения.

```golang
package main
 
import (
  "os"
  "fmt" 
  "strings" 
)
 
func main() {   
    // Set custom env variable
    os.Setenv("CUSTOM", "500")
     
    // fetcha all env variables
    for _, element := range os.Environ() {
        variable := strings.Split(element, "=")
        fmt.Println(variable[0],"=>",variable[1])        
    }
     
    // fetch specific env variables
    fmt.Println("CUSTOM=>", os.Getenv("CUSTOM"))
    fmt.Println("GOROOT=>", os.Getenv("GOROOT"))
}
```

