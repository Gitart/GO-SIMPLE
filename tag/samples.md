# Work with tag

```golang
// https://blog.golang.org/slices

package main
import "fmt"

// Structure
type TY struct{
	Name string
	Tags []string
}


//  Main
func main() {
	var Y TY
	Z := []string{"Первый","Второй","Третий","Четвертый"}
	Y.Tags=Z
    DeleteEl(Y,2)	
    fmt.Println(Y)
}


// Delete elements in tag
func DeleteEl(L TY, Num int) TY {
     RemoveElement(L.Tags,Num)
	 return L
}

// Adding
func Chng1(){
	var Y TY
	Z := []string{1:"t",2:"g","df",10:"ddd"}
	Z=append(Z,"Adding element")

	Y.Name="Name"
	Y.Tags=Z
	Y.Tags[0]="Change"

	fmt.Println(Y.Tags[4])
	fmt.Println(Y.Tags[0])
	// fmt.Println(Y)
}


// Adding
func Chng(){
	var Y TY
	Z := make([]string,5)

	Y.Name="Name"
	Y.Tags=Z
	Y.Tags[0]="Change"
	Y.Tags[4]="Changes"

	fmt.Println(Y.Tags[0])
	// fmt.Println(Y)
}



// Remove (v.1)
func RemoveElement(s []string, i int) []string {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

// Remove (v.2)
func RemoveIndex(s []int, index int) []int {
    return append(s[:index], s[index+1:]...)
}
```


### Function delete
```golang
package main

import (
    "fmt"
)

func RemoveIndex(s []int, index int) []int {
    return append(s[:index], s[index+1:]...)
}

func main() {
    all := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    fmt.Println(all) //[0 1 2 3 4 5 6 7 8 9]
    n := RemoveIndex(all, 5)
    fmt.Println(n) //[0 1 2 3 4 6 7 8 9]
}
```

## Sample

```golang
package main

import "fmt"

func Exclude(xs *[]string, excluded map[string]bool) {
    w := 0
    for _, x := range *xs {
        if !excluded[x] {
            (*xs)[w] = x
            w++
        }
    }
    *xs = (*xs)[:w]
}

func mapFromSlice(ex []string) map[string]bool {
    r := map[string]bool{}
    for _, e := range ex {
        r[e] = true
    }
    return r
}

func main() {
    urls := []string{"test", "abc", "def", "ghi"}
    remove := mapFromSlice([]string{"abc", "test"})
    Exclude(&urls, remove)
    fmt.Println(urls)
}
```


## Sample 3

```golang
package main

import "fmt"

func main() {
    urlList := []string{"test", "abc", "def", "ghi"}
    remove := []string{"abc", "test"}

    new_list := make([]string, 0)

    my_map := make(map[string]bool, 0)
    for _, ele := range remove {
        my_map[ele] = true
    }

    for _, ele := range urlList {
        _, is_in_map := my_map[ele]
        if is_in_map {
            fmt.Printf("Have to ignore : %s\n", ele)
        } else {
            new_list = append(new_list, ele)    
        }
    }

    fmt.Println(new_list)

}
```


