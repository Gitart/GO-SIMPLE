# Methods and Objects

## Go methods

A method is defined just like any other Go function. When a Go function is defined with a limited scope or attached with a specific type it is known as a method. Methods provide a way to add behavior to user\-defined types. Methods are really functions that contain an extra parameter that's declared between the keyword func and the function name.

In Go a method is a special kind of function that acts on variable of a certain type, called the receiver, which is an extra parameter placed before the method's name, used to specify the moderator type to which the method is attached. The receiver type can be anything, not only a struct type: any type can have methods, even a function type or alias types for int, bool, string or array.

The general format of a method is:
func (recv receiver\_type) methodName(parameter\_list) (return\_value\_list) { … }
The receiver is specified in ( ) before the method name after the func keyword.

#### Here is an example of methods on a non struct type:

package main
import "fmt"

type multiply int

func (m multiply) tentimes() int {
	return int(m \* 10)
}

func main() {
	var num int
	fmt.Print("Enter any positive integer: ")
    fmt.Scanln(&num)
	mul:= multiply(num)
	fmt.Println("Ten times of a given number is: ",mul.tentimes())
}

#### When you run the program, you get the following output:

C:\\golang>go run example42.go
Enter any positive integer: 5
Ten times of a given number is: 50

C:\\golang>

The parameter between the keyword func and the function name is called a receiver and binds the function to the specified type. When a function has a receiver, that function is called a method.

There are two types of receivers in Go: value receivers and pointer receivers. In above program the tentimes method is declared with a value receiver. The receiver for tentimes is declared as a value of type int. When you declare a method using a value receiver, the method will always be operating against a copy of the value used to make the method call.

The mul variable is initialized as the multiply type. Therefore, the tentimes method can be accessed using mul.tentimes() when we call the tentimes method, the value of mul is the receiver value for the call and the tentimes method is operating on a copy of this value.

#### Here's a program that shows implementation of number of methods attached to a type, via the receiver parameter, which is known as the type's method set.

package main
import "fmt"

type salary float64
func (s salary) total() total {
   return total(s)
}

type total float64
func (t total) hra() hra {
   t += t \* 0.3   // 30% HRA Addition
   return hra(t)
}
func (t total) salary() salary {
   t \-=t \* 0.10    // 10% Tax Deduction
   return salary(t)
}

type hra float64
func (h hra) basic() basic {
   h += h \* 0.3   // 30% HRA Addition
  return basic(h)
}
func (h hra) total() total {
  return total(h)
}

type basic float64
func (b basic) total() total {
   return total(b)
}

func main() {
    fmt.Println("Salary calculation for First Employee:")
    sal1 := basic(9000.00)
    fmt.Println(sal1.total())
    fmt.Println(sal1.total().hra().total())
    fmt.Println(sal1.total().hra().total().salary())

    fmt.Println("\\nSalary calculation for Second Employee:")
    sal2 := basic(5000.00)
    fmt.Println(sal2.total())
    fmt.Println(sal2.total().salary())
}

###### When you run the program, you get the following output:

C:\\golang>go run example43.go
Salary calculation for First Employee:
9000
11700
10530

Salary calculation for Second Employee:
5000
4500

C:\\golang>

## Method Overloading

Method Overloading is possible based on the receiver type, a method with the same name can exist on 2 of more different receiver types,e.g. this is allowed in the same package:
func (s \*inclTax) Salary(e Employee) Employee
func (s \*exclTax) Salary(e Employee) Employee

Receiver parameters can be passed as either values of or pointers of the base type. Pointer receiver parameters are widely used in Go.

###### You can also declare methods with pointer receivers.

package main
import "fmt"

type multiply int
type addition int

func (m \*multiply) twice() {
  \*m = multiply(\*m \* 2)
}

func (a \*addition) twice() {
  \*a = addition(\*a + \*a)
}

func main() {
  var mul multiply = 15
  mul.twice()
  fmt.Println(mul)

  var add addition = 15
  add.twice()
  fmt.Println(add)
}

###### When you run the program, you get the following output:

C:\\golang>go run example.go
30
30
C:\\golang>

Let's take another example in which we call to a method which receives a pointer to Struct. We call to that method by value and by pointer and we got the same result.

###### Example to declare methods with pointer receivers.

package main
import "fmt"

type multiply struct {
	num int
}

func (m \*multiply) twice(n int) {
  	m.num = n\*2
}

func (m multiply) display() int{
  	return m.num
}

func main() {
  fmt.Println("Call by value")
  var mul1 multiply	// mul1 is a value
  mul1.twice(10)
  fmt.Println(mul1.display())

  fmt.Println("Call by pointer")
  mul2 := new(multiply) // mul2 is a pointer
  mul2.twice(10)
  fmt.Println(mul2.display())
}

###### When you run the program, you get the following output:

C:\\golang>go run example.go
Call by value
20
Call by pointer
20

In the main() we ourselves do not have to figure out whether to call the methods on a pointer or not, Go does that for us. mul1 is a value and mul2 is a pointer, but the methods calls work just fine.

## Objects in Go

There is no concept of a class type that serves as the basis for objects in Go. Any data type in Go can be used as an object.struct type in Go can receive a set of method to define its behavior. There is no special type called a class or object exist in GO. strct type in Go comes the closet to what is commonly refer as object in other programming language. Go supports the majority of concepts that are usually attributed to object\-oriented programming.

With the help of concepts like packages and an extensible type system Go supports physical and logical modularity at its core; hence we able to achieve Modularity and encapsulation in Go.

A newly declared name type does not inherit all attributes of its underlying type and are treated variously by the type system. Hence GO doesn't support polymorphism through inheritance.But it is possible to create objects and express their polymorphic relationships through composition using a type such as a struct or an interface.

Let us start with the following simple example to demonstrate how the struct type may be used as an object that can achieve polymorphic composition.

package main
import "fmt"

type gadgets uint8
const (
    Camera gadgets = iota
    Bluetooth
    Media
    Storage
    VideoCalling
    Multitasking
    Messaging
)
type mobile struct {
    make string
    model string
}

type smartphone struct {
   gadgets gadgets
}

func (s \*smartphone) launch() {
   fmt.Println ("New Smartphone Launched:")
}

type android struct {
   mobile
   smartphone
   waterproof string
}
func (a \*android) samsung() {
   fmt.Printf("%s %s\\n",
          a.make, a.model)
}

type iphone struct {
   mobile
   smartphone
   sensor int
}
func (i \*iphone) apple() {
   fmt.Printf("%s %s\\n",
          i.make, i.model)
}

func main() {
   t := &android {}
   t.make ="Samsung"
   t.model ="Galaxy J7 Prime"
   t.gadgets = Camera+Bluetooth+Media+Storage+VideoCalling+Multitasking+Messaging
   t.launch()
   t.samsung()
}

C:\\golang>go run example.go
New Smartphone Launched:
Samsung Galaxy J7 Prime

In above program the composition over inheritance principle is used to achieve polymorphism using the type embedding mechanism supported by the struct type. Here each type is independent and is considered to be different from all others. The above program show that the types iphone and android is a mobile via a subtype relationship.

The methods t.launch() being invoked however, neither type, iphone nor android, are receivers of a method named launch(). The launch()method is defined for the smartphone type. Since the smartphone type is embedded in the types iphone nor android, the launch() method is promoted in scope to these enclosing types and is therefore accessible.
