# Compare


```go
// Go program to illustrate the concept
// of == and != operator with strings


package main
 
import "fmt"
 
// Main function
func main() {
 
    // Creating and initializing strings
    // using shorthand declaration
    str1 := "Geeks"
    str2 := "Geek"
    str3 := "GeeksforGeeks"
    str4 := "Geeks"
 
    // Checking the string are equal
    // or not using == operator
    result1 := str1 == str2
    result2 := str2 == str3
    result3 := str3 == str4
    result4 := str1 == str4
     
    fmt.Println("Result 1: ", result1)
    fmt.Println("Result 2: ", result2)
    fmt.Println("Result 3: ", result3)
    fmt.Println("Result 4: ", result4)
 
    // Checking the string are not equal
    // using != operator
    result5 := str1 != str2
    result6 := str2 != str3
    result7 := str3 != str4
    result8 := str1 != str4
     
    fmt.Println("\nResult 5: ", result5)
    fmt.Println("Result 6: ", result6)
    fmt.Println("Result 7: ", result7)
    fmt.Println("Result 8: ", result8)
 
}
```

**Output: **

```
Result 1:  false
Result 2:  false
Result 3:  false
Result 4:  true

Result 5:  true
Result 6:  true
Result 7:  true
Result 8:  false
```

**Example 2: **

```go
// Go program to illustrate the concept
// of comparison operator with strings
package main
 
import "fmt"
 
// Main function
func main() {
 
    // Creating and initializing
    // slice of string using the
    // shorthand declaration
    myslice := []string{"Geeks", "Geeks",
                    "gfg", "GFG", "for"}
     
    fmt.Println("Slice: ", myslice)
 
    // Using comparison operator
    result1 := "GFG" > "Geeks"
    fmt.Println("Result 1: ", result1)
 
    result2 := "GFG" < "Geeks"
    fmt.Println("Result 2: ", result2)
 
    result3 := "Geeks" >= "for"
    fmt.Println("Result 3: ", result3)
 
    result4 := "Geeks" <= "for"
    fmt.Println("Result 4: ", result4)
 
    result5 := "Geeks" == "Geeks"
    fmt.Println("Result 5: ", result5)
 
    result6 := "Geeks" != "for"
    fmt.Println("Result 6: ", result6)
 
}
```

**Output: **

```
Slice:  [Geeks Geeks gfg GFG for]
Result 1:  false
Result 2:  true
Result 3:  false
Result 4:  true
Result 5:  true
Result 6:  true
2. Using Compare() method: You can also compare two strings using the built-in function Compare() provided by the strings package. This function returns an integer value after comparing two strings lexicographically. The return values are: 

Return 0, if str1 == str2.
Return 1, if str1 > str2.
Return -1, if str1 < str2.
**Syntax: **

func Compare(str1, str2 string) int

**Example:**

```go
// Go program to illustrate how to compare
// string using compare() function
package main
 
import (
    "fmt"
    "strings"
)
 
func main() {
 
    // Comparing string using Compare function
    fmt.Println(strings.Compare("gfg", "Geeks"))
     
    fmt.Println(strings.Compare("GeeksforGeeks",
                               "GeeksforGeeks"))
     
    fmt.Println(strings.Compare("Geeks", " GFG"))
     
    fmt.Println(strings.Compare("GeeKS", "GeeKs"))
 
}
Output: 

1
0
1
-1
```
