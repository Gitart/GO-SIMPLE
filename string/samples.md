## Using Sprintf: In Go language, you can also concatenate string using Sprintf() method.
**Example:**

```go
// Go program to illustrate how to concatenate strings
// Using Sprintf function
package main
  
import "fmt"
  
func main() {
  
    // Creating and initializing strings
    str1 := "Tutorial"
    str2 := "of"
    str3 := "Go"
    str4 := "Language"
  
    // Concatenating strings using 
    // Sprintf() function
    result := fmt.Sprintf("%s%s%s%s", str1, 
                          str2, str3, str4)
      
    fmt.Println(result)
}
```

Output:
```
TutorialofGoLanguage
Using += operator or String append: In Go strings, you are allowed to append a string using += operator. This operator adds a new or given string to the end of the specified string.
```

**Example:**

```go
// Go program to illustrate how
// to concatenate strings
// Using += operator
package main
  
import "fmt"
  
func main() {
  
    // Creating and initializing strings
    str1 := "Welcome"
    str2 := "GeeksforGeeks"
  
    // Using += operator
    str1 += str2
    fmt.Println("String: ", str1)
  
    str1 += "This is the tutorial of Go language"
    fmt.Println("String: ", str1)
  
    str2 += "Portal"
    fmt.Println("String: ", str2)
  
}
```

Output:

```
String:  WelcomeGeeksforGeeks
String:  WelcomeGeeksforGeeksThis is the tutorial of Go language
String:  GeeksforGeeksPortal
Using Join() function: This function concatenates all the elements present in the slice of string into a single string. This function is available in string package.
```

**Syntax:**

func Join(str []string, sep string) string
Here, str is the string from which we can concatenate elements and sep is the separator which is placed between the elements in the final string.

Example:

```go
// Go program to illustrate how to
// concatenate all the elements
// present in the slice of the string
package main
  
import (
    "fmt"
    "strings"
)
  
func main() {
  
    // Creating and initializing slice of string
    myslice := []string{"Welcome", "To",
              "GeeksforGeeks", "Portal"}
  
    // Concatenating the elements 
    // present in the slice
    // Using join() function
    result := strings.Join(myslice, "-")
    fmt.Println(result)
}
```
