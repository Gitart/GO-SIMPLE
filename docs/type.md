# Data Type

Go is a statically typed programming language. This means that variables always have a specific type and that type cannot change. The keyword var is used for declaring variables of a particular data type. Here is the syntax for declaring variables:

var name type = expression

On the left we use the var keyword to declare a variable and then assign a value to it. We can declare mutiple variables of the same type in a single statement as shown here:

var fname,lname string

Multiple variables of the same type can also be declared on a single line: var x, y int makes x and y both int variables. You can also make use of parallel assignment: a, b := 20, 16 If you are using an initializer expression for declaring variables, you can omit the type using short variable declaration as shown here:

country, state := "Germany", "Berlin"

We use the operator : = for declaring and initializing variables with short variable declaration. When you declare variables with this method, you can't specify the type because the type is determined by the initializer expression.

package main

import "fmt"

//Global variable declaration
var (m int
    n int)

func main(){
    var x int = 1 // Integer Data Type
    var y int    //  Integer Data Type
    fmt.Println(x)
    fmt.Println(y)

    var a,b,c = 5.25,25.25,14.15 // Multiple float32 variable declaration
    fmt.Println(a,b,c)

    city:="Berlin" // String variable declaration
    Country:="Germany" // Variable names are case sensitive
    fmt.Println(city)
    fmt.Println(Country) // Variable names are case sensitive

    food,drink,price:="Pizza","Pepsi",125  // Multiple type of variable declaration in same line
    fmt.Println(food,drink,price)
    m,n=1,2
    fmt.Println(m,n)
}
