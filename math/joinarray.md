## Join arrays or slices example

A quick note on how to join arrays or slices in Golang. So used to Python way of joining arrays with + symbol. 
However, it is not available in Golang :(

```golang
 package main

 import (
         "fmt"
 )

 func main() {
         list1 := []int{1, 2, 3}
         list2 := []int{4, 5, 6}

         // python way - will not work in Golang
         //list3 := list1 + list2
         //fmt.Println(list3)

         list3 := list1

         // example
         // to combine two slices or join arrays, use for loop and builtin append function
         for index, _ := range list2 {
                 list3 = append(list3, list2[index])
         }

         fmt.Println(list3)

         // another example
         // super quick way to join arrays
         fmt.Println(append(list1, list2...))

 }
 ```
Output:

```
[1 2 3 4 5 6]
[1 2 3 4 5 6]
```
