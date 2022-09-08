## Как получить доступ к отдельному байту строки?
Строка состоит из байта, поэтому мы можем получить доступ к каждому байту данной строки.

```go
// Go program to illustrate how to
// access the bytes of the string
package main
  
import "fmt"
  
// Main function
func main() {
  
    // Creating and initializing a string
    str := "Welcome to GeeksforGeeks"
  
    // Accessing the bytes of the given string
    for c := 0; c < len(str); c++ {
  
        fmt.Printf("\nCharacter = %c Bytes = %v", str, str)
    }
}
```


