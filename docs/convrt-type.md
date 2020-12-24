# Converting and Checking Types

Package **strconv** implements conversions to and from string representations of basic data types. **Atoi** is equivalent to ParseInt(s, 10, 0), converted to type int. **ParseInt** interprets a string s in the given base (0, 2 to 36) and bit size (0 to 64) and returns the corresponding value i.

package main

import "strconv"

func main() {
	strVar := "100"
	intVar, \_ := strconv.Atoi(strVar)

	strVar1 := "\-52541"
	intVar1, \_ := strconv.ParseInt(strVar1, 10, 32)

	strVar2 := "101010101010101010"
	intVar2, \_ := strconv.ParseInt(strVar2, 10, 64)
}

---

## How to Convert string to float type in Go?

**ParseFloat** converts the string s to a floating\-point number with the precision specified by bitSize: 32 for float32, or 64 for float64. When bitSize=32, the result still has type float64, but it will be convertible to float32 without changing its value.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "3.1415926535"
	f, \_ := strconv.ParseFloat(s, 8)
	fmt.Printf("%T, %v\\n", f, f)

	s1 := "\-3.141"
	f1, \_ := strconv.ParseFloat(s1, 8)
	fmt.Printf("%T, %v\\n", f1, f1)

	s2 := "\-3.141"
	f2, \_ := strconv.ParseFloat(s2, 32)
	fmt.Printf("%T, %v\\n", f2, f2)
}

float64, 3.1415926535
float64, \-3.141
float64, \-3.1410000324249268

---

## How to convert String to Boolean Data Type Conversion in Go?

**ParseBool** returns the boolean value represented by the string. It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False. Any other value returns an error.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := "true"
	b1, \_ := strconv.ParseBool(s1)
	fmt.Printf("%T, %v\\n", b1, b1)

	s2 := "t"
	b2, \_ := strconv.ParseBool(s2)
	fmt.Printf("%T, %v\\n", b2, b2)

	s3 := "0"
	b3, \_ := strconv.ParseBool(s3)
	fmt.Printf("%T, %v\\n", b3, b3)

	s4 := "F"
	b4, \_ := strconv.ParseBool(s4)
	fmt.Printf("%T, %v\\n", b4, b4)
}

bool, true
bool, true
bool, false
bool, false

---

## How to convert Boolean Type to String in Go?

**FormatBool** function used to convert Boolean variable into String.

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var b bool = true
	fmt.Println(reflect.TypeOf(b))

	var s string = strconv.FormatBool(true)
	fmt.Println(reflect.TypeOf(s))
}

bool
string

---

## How to Convert Float to String type in Go?

FormatFloat converts the floating\-point number f to a string s.

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var f float64 = 3.1415926535
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(f)

	var s string = strconv.FormatFloat(f, 'E', \-1, 32)
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(s)
}

float64
3.1415926535
string
3.1415927E+00

---

## Convert Integer Type to String Type

FormatInt converts the Integer number i to a String s.

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var i int64 = \-654
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(i)

	var s string = strconv.FormatInt(i, 10)
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(s)
}

int64
\-654
string
\-654

---

## Convert Int data type to Int16 Int32 Int64

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int = 10
	fmt.Println(reflect.TypeOf(i))

	i16 := int16(i)
	fmt.Println(reflect.TypeOf(i16))

	i32 := int32(i)
	fmt.Println(reflect.TypeOf(i32))

	i64 := int64(i)
	fmt.Println(reflect.TypeOf(i64))
}

int
int16
int32
int64

---

## Convert Float32 to Float64 and Float64 to Float32

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f32 float32 = 10.6556
	fmt.Println(reflect.TypeOf(f32))

	f64 := float64(f32)
	fmt.Println(reflect.TypeOf(f64))

	f64 = 1097.655698798798
	fmt.Println(f64)

	f32 = float32(f64)
	fmt.Println(f32)
}

float32
float64
1097.655698798798
1097.6556

---

## Converting Int data type to Float in Go

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f32 float32 = 10.6556
	fmt.Println(reflect.TypeOf(f32))

	i32 := int32(f32)
	fmt.Println(reflect.TypeOf(i32))
	fmt.Println(i32)

	f64 := float64(i32)
	fmt.Println(reflect.TypeOf(f64))
}

float32
int32
10
float64
