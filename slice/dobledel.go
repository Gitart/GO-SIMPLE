package main
  
import (
"fmt"
)
  
func unique(intSlice []int) []int {
    keys := make(map[int]bool)
    list := []int{} 
    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }    
    return list
}
  
func main() {
    intSlice := []int{1,5,3,6,9,9,4,2,3,1,5}
    fmt.Println(intSlice) 
    uniqueSlice := unique(intSlice)
    fmt.Println(uniqueSlice)
}
