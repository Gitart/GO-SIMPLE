# Design philosophy of Golang interface type

Implementing a Go interface is done implicitly. In Go, there is no need to explicitly implement an interface into a concrete type by specifying any keyword. To implement an interface into a concrete type, just provide the methods with the same signature that is defined in the interface type.

An interface type is defined with the keyword interface. An interface defines a set of methods (the method set), but these methods do not contain code: they are not implemented (they are abstract). A method set is a list of methods that a type must have in order to implement the interface. Also an interface cannot contain variables.

#### An interface is declared in the format

Interface Type catalog

type catalog interface {
	shipping() float64
	tax() float64
}

The interface type catalog is a contract for creating various product types in a catalog. The catalog interface provides two behaviors in its contract: shipping and tax.

## Implementing an interface

The following source code shows the configurable type as an implementation of the interface type catalog. The configurable type is defined as a struct with receiver methods shipping and tax. This fact automatically qualifies configurable as an implementation of catalog:

package main
import "fmt"

type catalog interface {
   shipping() float64
   tax() float64
}

type configurable struct {
   name string
   price, qty float64
}

func (c \*configurable) tax() float64{
  return c.price \* c.qty \* 0.05
}

func (c \*configurable) shipping() float64{
  return c.qty \* 5
}

func  main() {
  tshirt := configurable{}
  tshirt.price = 250
  tshirt.qty = 2
  fmt.Println("Shipping Charge: ", tshirt.shipping())
  fmt.Println("Tax: ", tshirt.tax())
}

###### When you run the program, you get the following output:

C:\\golang>go run example.go
Shipping Charge: 10
Tax: 25

## Subtyping with Go interfaces

Go support composition (has\-a) relationships when building objects using subtyping via interfaces. This can be explained that the configurable, download, simple type (and any other type that implements the methods shipping and tax) can be treated as a subtype of catalog, as shown in the following figure:

![Implementing an interface](https://www.golangprograms.com/media/wysiwyg/subtyping.jpg)

package main
import "fmt"

type catalog interface {
   shipping() float64
   tax() float64
}

type configurable struct {
   name string
   price, qty float64
}

func (c \*configurable) tax() float64{
  return c.price \* c.qty \* 0.05
}

func (c \*configurable) shipping() float64{
  return c.qty \* 5
}

type download struct{
    name string
    price, qty float64
}

func (d \*download) tax() float64{
  return d.price \* d.qty \* 0.07
}

type simple struct {
  name string
  price, qty float64
}

func (s \*simple) tax() float64{
  return s.price \* s.qty \* 0.03
}

func (s \*simple) shipping() float64{
  return s.qty \* 3
}

func  main() {
  tshirt := configurable{}
  tshirt.price = 250
  tshirt.qty = 2
  fmt.Println("Configurable Product")
  fmt.Println("Shipping Charge: ", tshirt.shipping())
  fmt.Println("Tax: ", tshirt.tax())

  mobile := simple{"Samsung S\-7",10,25}
  fmt.Println("\\nSimple Product")
  fmt.Println("Shipping Charge: ", mobile.shipping())
  fmt.Println("Tax: ", mobile.tax())

  book := download{"Python in 24 Hours",19,1}
  fmt.Println("\\nDownloadable Product")
  fmt.Println("Tax: ", book.tax())
}

###### When you run the program, you get the following output:

C:\\golang>go run example.go
Configurable Product
Shipping Charge: 10
Tax: 25

Simple Product
Shipping Charge: 75
Tax: 7.5

Downloadable Product
Tax: 1.33

C:\\golang>

## Multiple interfaces

In GO, the implicit mechanism of interfaces allows to satisfy multiple interface type at once. This can be implemented by intergrating the method set of a given type intersect with the methods of each interface type. Let us re\-implement the previous code. New interface discount has been created. This is illustrated by the following figure:

![Implementing Multiple interfaces](https://www.golangprograms.com/media/wysiwyg/subtyping1.jpg)

package main
import "fmt"

type catalog interface {
   shipping() float64
   tax() float64
}

type discount interface{
    offer() float64
}

type configurable struct {
   name string
   price, qty float64
}

func (c \*configurable) tax() float64{
  return c.price \* c.qty \* 0.05
}

func (c \*configurable) shipping() float64{
  return c.qty \* 5
}

func (c \*configurable) offer() float64{
  return c.price \* 0.15
}

type download struct{
    name string
    price, qty float64
}

func (d \*download) tax() float64{
  return d.price \* d.qty \* 0.10
}

type simple struct {
  name string
  price, qty float64
}

func (s \*simple) tax() float64{
  return s.price \* s.qty \* 0.03
}

func (s \*simple) shipping() float64{
  return s.qty \* 3
}

func (s \*simple) offer() float64{
  return s.price \* 0.10
}

func  main() {
  tshirt := configurable{}
  tshirt.price = 250
  tshirt.qty = 2
  fmt.Println("Configurable Product")
  fmt.Println("Shipping Charge: ", tshirt.shipping())
  fmt.Println("Tax: ", tshirt.tax())
  fmt.Println("Discount: ", tshirt.offer())

  mobile := simple{"Samsung S\-7",3000,2}
  fmt.Println("\\nSimple Product")
  fmt.Println(mobile.name)
  fmt.Println("Shipping Charge: ", mobile.shipping())
  fmt.Println("Tax: ", mobile.tax())
  fmt.Println("Discount: ", mobile.offer())

  book := download{"Python in 24 Hours",50,1}
  fmt.Println("\\nDownloadable Product")
  fmt.Println(book.name)
  fmt.Println("Tax: ", book.tax())
}

###### When you run the program, you get the following output:

C:\\golang>go run example.go
Configurable Product
Shipping Charge: 10
Tax: 25
Discount: 37.5

Simple Product
Samsung S\-7
Shipping Charge: 6
Tax: 180
Discount: 300

Downloadable Product
Python in 24 Hours
Tax: 5

C:\\golang>

## Interface embedding

In GO, the interface type also support for type embedding (similar to the struct type). This gives you the flexibility to structure your types in ways that maximize type reuse.
Continuing with the catalog example, a struct configurable is declared in which the type discount and giftpack is embedded. Here you create more concrete types of the catalog interface. Because type giftpack and discount is an implementation of the catalog interface, the type configurable is also an implementation of the catalog interface. All fields and methods defined in the Type discount and giftpack types are also available in the configurable type.
The following illustration shows how the interface types may be combined so the is\-a relationship still satisfies the relationships between code components:

![Implementing Interface embedding](https://www.golangprograms.com/media/wysiwyg/subtyping2.jpg)

package main
import "fmt"

type discount interface{
    offer() float64
}

type giftpack interface{
    available() string
}

type catalog interface {
   discount
   giftpack
   shipping() float64
   tax() float64
}

type configurable struct {
   name string
   price, qty float64
}

func (c \*configurable) tax() float64{
  return c.price \* c.qty \* 0.05
}

func (c \*configurable) shipping() float64{
  return c.qty \* 5
}

func (c \*configurable) offer() float64{
  return c.price \* 0.15
}

func (c \*configurable) available() string{
    if c.price > 1000{
      return "Gift Pack Available"
    }
    return "Gift Pack not Available"
}

type download struct{
    name string
    price, qty float64
}

func (d \*download) tax() float64{
  return d.price \* d.qty \* 0.10
}

func (d \*download) available() string{
    if d.price > 500{
      return "Gift Pack Available"
    }
    return "Gift Pack not Available"
}

type simple struct {
  name string
  price, qty float64
}

func (s \*simple) tax() float64{
  return s.price \* s.qty \* 0.03
}

func (s \*simple) shipping() float64{
  return s.qty \* 3
}

func (s \*simple) offer() float64{
  return s.price \* 0.10
}

func  main() {
  tshirt := configurable{}
  tshirt.price = 1550
  tshirt.qty = 2
  fmt.Println("Configurable Product")
  fmt.Println("Shipping Charge: ", tshirt.shipping())
  fmt.Println("Tax: ", tshirt.tax())
  fmt.Println("Discount: ", tshirt.offer())
  fmt.Println(tshirt.available())

  mobile := simple{"Samsung S\-7",3000,2}
  fmt.Println("\\nSimple Product")
  fmt.Println(mobile.name)
  fmt.Println("Shipping Charge: ", mobile.shipping())
  fmt.Println("Tax: ", mobile.tax())
  fmt.Println("Discount: ", mobile.offer())

  book := download{"Python in 24 Hours",50,1}
  fmt.Println("\\nDownloadable Product")
  fmt.Println(book.name)
  fmt.Println("Tax: ", book.tax())
  fmt.Println(book.available())
}

###### When you run the program, you get the following output:

C:\\golang>go run example.go
Configurable Product
Shipping Charge: 10
Tax: 155
Discount: 232.5
Gift Pack Available

Simple Product
Samsung S\-7
Shipping Charge: 6
Tax: 180
Discount: 300

Downloadable Product
Python in 24 Hours
Tax: 5
Gift Pack not Available

C:\\golang>
