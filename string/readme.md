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

## Как найти длину строки?
В строке Golang вы можете найти длину строки, используя две функции: одна — len(), а другая — RuneCountInString(). 
```go
// Go program to illustrate how to
// find the length of the string
  
package main
  
import (
    "fmt"
    "unicode/utf8"
)
  
// Main function
func main() {
  
    // Creating and initializing a string
    // using shorthand declaration
    mystr := "Welcome to GeeksforGeeks ??????"
  
    // Finding the length of the string
    // Using len() function
    length1 := len(mystr)
  
    // Using RuneCountInString() function
    length2 := utf8.RuneCountInString(mystr)
  
    // Displaying the length of the string
    fmt.Println("string:", mystr)
    fmt.Println("Length 1:", length1)
    fmt.Println("Length 2:", length2)
  
}
```
