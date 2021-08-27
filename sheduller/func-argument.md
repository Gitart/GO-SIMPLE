## Sample
```go
package main

import "fmt"

type fn func(int) 

func myfn1(i int) {
	fmt.Printf("\ni is %v", i)
}
func myfn2(i int) {
	fmt.Printf("\ni is %v", i)
}
func test(f fn, val int) {
	f(val)
}
func main() {
	test(myfn1, 123)
	test(myfn2, 321)
}
```

```go
package main

    import "fmt"

    func plusTwo() (func(v int) (int)) {
        return func(v int) (int) {
            return v+2
        }
    }

    func plusX(x int) (func(v int) (int)) {
       return func(v int) (int) {
           return v+x
       }
    }

    func main() {
        p := plusTwo()
        fmt.Printf("3+2: %d\n", p(3))

        px := plusX(3)
        fmt.Printf("3+3: %d\n", px(3))
    }
    ```
    
    # Sample
    ```go
    package main

import "fmt"

func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {

    sum(1, 2)
    sum(1, 2, 3)

    nums := []int{1, 2, 3, 4}
    sum(nums...)
}
```

