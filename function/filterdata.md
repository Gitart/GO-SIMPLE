## Func filter

```go
package main

import "fmt"

type User struct {
    name       string
    occupation string
    married    bool
}

func main() {

    u1 := User{"John Doe", "gardener", false}
    u2 := User{"Richard Roe", "driver", true}
    u3 := User{"Bob Martin", "teacher", true}
    u4 := User{"Lucy Smith", "accountant", false}
    u5 := User{"James Brown", "teacher", true}

    users := []User{u1, u2, u3, u4, u5}

    married := filter(users, func(u User) bool {
        if u.married == true {
            return true
        }
        return false
    })

    teachers := filter(users, func(u User) bool {

        if u.occupation == "teacher" {
            return true
        }
        return false
    })

    fmt.Println("Married:")
    fmt.Printf("%v\n", married)

    fmt.Println("Teachers:")
    fmt.Printf("%v\n", teachers)

}

func filter(s []User, f func(User) bool) []User {
    var res []User

    for _, v := range s {

        if f(v) == true {
            res = append(res, v)
        }
    }
    return res
}
```
We have a slice of User structures. We filter the slice to form new slices of married users and users that are teachers.

```
married := filter(users, func(u User) bool {
    if u.married == true {
        return true
    }
    return false
})
```

We call the filter function. It accepts an anonymous function as a parameter. The function returns true for married users. A function that returns a boolean value is known also as a predicate.
```
func filter(s []User, f func(User) bool) []User {
    var res []User

    for _, v := range s {

        if f(v) == true {
            res = append(res, v)
        }
    }
    return res
}
```

### The filter function forms a new slice for all users that satisfy the given condition.

$ go run filtering.go 
Married:
[{Richard Roe driver true} {Bob Martin teacher true} {James Brown teacher true}]
Teachers:
[{Bob Martin teacher true} {James Brown teacher true}]
