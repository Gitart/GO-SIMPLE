
ZetCode
All Go Python C# Java JavaScript Subscribe
Ebooks

    PyQt5 ebook
    Tkinter ebook
    SQLite Python
    wxPython ebook
    Windows API ebook
    Java Swing ebook
    Java games ebook
    MySQL Java ebook

Go string format

last modified January 26, 2022

Go string format tutorial shows how to format strings in Golang. In fmt package we find functions that implement formatted I/O.

To format strings in Go, we use functions including fmt.Printf, fmt.Sprintf, or fmt.Fscanf.

The functions take the format string and the list of arguments as parameters.

%[flags][width][.precision]verb

The format string has this syntax. The options specified within [] characters are optional.

The verb at the end defines the type and the interpretation of its corresponding argument.

    d - decimal integer
    o - octal integer
    O - octal integer with 0o prefix
    b - binary integer
    x - hexadecimal integer lowercase
    X - hexadecimal integer uppercase
    f - decimal floating point, lowercase
    F - decimal floating point, uppercase
    e - scientific notation (mantissa/exponent), lowercase
    E - scientific notation (mantissa/exponent), uppercase
    g - the shortest representation of %e or %f
    G - the shortest representation of %E or %F
    c - a character represented by the corresponding Unicode code point
    q - a quoted character
    U - Unicode escape sequence
    t - the word true or false
    s - a string
    v - default format
    #v - Go-syntax representation of the value
    T - a Go-syntax representation of the type of the value
    p - pointer address
    % - a double %% prints a single %

The flags is a set of characters that modify the output format. The set of valid flags depends on the conversion character. The width is a non-negative decimal integer indicating the minimum number of runes to be written to the output. If the value to be printed is shorter than the width, the result is padded with blank spaces. The value is not truncated even if the result is larger.

For integer conversion characterss the precision specifies the minimum number of digits to be written. If the value to be written is shorter than this number, the result is padded with leading zeros. For strings it is the maximum number of runes to be printed. For e, E, f and F verbs, it is the number of digits to be printed after the decimal point. For g and G verbs it is the maximum number of significant digits to be printed.
Go string format functions

The formatting functions format the string according to the specified format specifiers.
fmt_funs.go

package main

import (
    "fmt"
)

func main() {

    name := "Jane"
    age := 17

    fmt.Printf("%s is %d years old\n", name, age)

    res := fmt.Sprintf("%s is %d years old", name, age)
    fmt.Println(res)
}

In the code example, we use two string formatting functions: fmt.Printf and fmt.Sprintf.

fmt.Printf("%s is %d years old\n", name, age)

The fmt.Printf function prints a formatted string to the console. The %s expects a string value and the %d an integer value.

res := fmt.Sprintf("%s is %d years old", name, age)

The fmt.Sprintf function formats a string into a variable.

$ go run fmt_funs.go
Jane is 17 years old
Jane is 17 years old

Go string format general verbs

The following example uses some general verbs.
general.go

package main

import (
    "fmt"
)

type User struct {
    name       string
    occupation string
}

func main() {

    msg := "and old falcon"
    n := 16
    w := 12.45
    r := true
    u := User{"John Doe", "gardener"}
    vals := []int{1, 2, 3, 4, 5}
    ctrs := map[string]string{
        "sk": "Slovakia",
        "ru": "Russia",
        "de": "Germany",
        "no": "Norway",
    }

    fmt.Printf("%v %v %v %v %v\n  %v %v\n", msg, n, w, u, r, vals, ctrs)
    fmt.Printf("%v %+v\n", u, u)

    fmt.Println("--------------------")

    fmt.Printf("%#v %#v %#v %#v %#v\n  %#v %#v\n", msg, n, w, u, r, vals, ctrs)
    fmt.Printf("%T %T %T %T %T %T %T\n", msg, n, w, u, r, vals, ctrs)

    fmt.Println("--------------------")

    fmt.Printf("The prices dropped by 12%%\n")
}

The example presents Go's general verbs. The %v and %#v are useful for determining the values of Go data types. The %T is useful for determining the data type of a variable. The %% simply outputs the percent sign.

$ go run general.go 
and old falcon 16 12.45 {John Doe gardener} true
  [1 2 3 4 5] map[de:Germany no:Norway ru:Russia sk:Slovakia]
{John Doe gardener} {name:John Doe occupation:gardener}
--------------------
"and old falcon" 16 12.45 main.User{name:"John Doe", occupation:"gardener"} true
  []int{1, 2, 3, 4, 5} map[string]string{"de":"Germany", "no":"Norway", "ru":"Russia", "sk":"Slovakia"}
string int float64 main.User bool []int map[string]string
--------------------
The prices dropped by 12%

Go string format indexing

The formatting functions apply the format specifiers by the order of the given arguments. The next example shows how to change their order.
indexing.go

package main

import (
    "fmt"
)

func main() {

    n1 := 2
    n2 := 3
    n3 := 4

    res := fmt.Sprintf("There are %d oranges %d apples %d plums", n1, n2, n3)
    fmt.Println(res)

    res2 := fmt.Sprintf("There are %[2]d oranges %d apples %[1]d plums", n1, n2, n3)
    fmt.Println(res2)
}

We format two strings. In the first case, the variables are applied as they are specified. In the second case, we change their order with [2] and [1] characters, which take the third and the second arguments, respectively.

$ go run indexing.go
There are 2 oranges 3 apples 4 plums
There are 3 oranges 4 apples 2 plums

Go string format conversion characters

The format conversion characters define the type and the interpretation of their corresponding arguments.
con_chars.go

package main

import (
    "fmt"
)

func main() {

    fmt.Printf("%d\n", 1671)
    fmt.Printf("%o\n", 1671)
    fmt.Printf("%x\n", 1671)
    fmt.Printf("%X\n", 1671)
    fmt.Printf("%#b\n", 1671)
    fmt.Printf("%f\n", 1671.678)
    fmt.Printf("%F\n", 1671.678)
    fmt.Printf("%e\n", 1671.678)
    fmt.Printf("%E\n", 1671.678)
    fmt.Printf("%g\n", 1671.678)
    fmt.Printf("%G\n", 1671.678)
    fmt.Printf("%s\n", "Zetcode")
    fmt.Printf("%c %c %c %c %c %c %c\n", 'Z', 'e', 't',
        'C', 'o', 'd', 'e')
    fmt.Printf("%p\n", []int{1, 2, 3})
    fmt.Printf("%d%%\n", 1671)
    fmt.Printf("%t\n", 3 > 5)
    fmt.Printf("%t\n", 5 > 3)
}

The example shows the Go's string formt conversion characters.

$ go run con_chars.go
1671
3207
687
687
0b11010000111
1671.678000
1671.678000
1.671678e+03
1.671678E+03
1671.678
1671.678
Zetcode
Z e t C o d e
0xc0000c0000
1671%
false
true

Go string format integers

The following example formats integers.
integers.go

package main

import (
    "fmt"
)

func main() {

    val := 122

    fmt.Printf("%d\n", val)
    fmt.Printf("%c\n", val)
    fmt.Printf("%q\n", val)
    fmt.Printf("%x\n", val)
    fmt.Printf("%X\n", val)
    fmt.Printf("%o\n", val)
    fmt.Printf("%O\n", val)
    fmt.Printf("%b\n", val)
    fmt.Printf("%U\n", val)
}

The integer is formatted in decimal, hexadecimal, octal, and binary notations. It is also formatted as a character literal, quoted charracter literal, and in a Unicode escape sequence.

$ go run integers.go
122
z
'z'
7a
7A
172
0o172
1111010
U+007A

Go string format precision

The following example sets the precision of a floating point value.
precision.go

package main

import (
    "fmt"
)

func main() {

    fmt.Printf("%0.f\n", 16.540)
    fmt.Printf("%0.2f\n", 16.540)
    fmt.Printf("%0.3f\n", 16.540)
    fmt.Printf("%0.5f\n", 16.540)
}

For floating point values, the precision is the number of digits to be printed after the decimal point.

$ go run precision.go
17
16.54
16.540
16.54000

Go string format scientific notation

The e, E, g, and G verbs are used to format numbers in scientific notation.

    f - decimal floating point, lowercase
    F - decimal floating point, uppercase
    e - scientific notation (mantissa/exponent), lowercase
    E - scientific notation (mantissa/exponent), uppercase
    g - uses the shortest representation of %e or %f
    G - Use the shortest representation of %E or %F

scientific.go

package main

import (
    "fmt"
)

func main() {

    val := 1273.78888769000

    fmt.Printf("%f\n", val)
    fmt.Printf("%e\n", val)
    fmt.Printf("%g\n", val)
    fmt.Printf("%E\n", val)
    fmt.Printf("%G\n", val)

    fmt.Println("-------------------------")

    fmt.Printf("%.10f\n", val)
    fmt.Printf("%.10e\n", val)
    fmt.Printf("%.10g\n", val)
    fmt.Printf("%.10E\n", val)
    fmt.Printf("%.10G\n", val)

    fmt.Println("-------------------------")

    val2 := 66_000_000_000.1200

    fmt.Printf("%f\n", val2)
    fmt.Printf("%e\n", val2)
    fmt.Printf("%g\n", val2)
    fmt.Printf("%E\n", val2)
    fmt.Printf("%G\n", val2)
}

The example formats floating point values in normal decimal and scientific notations.

$ go run scientific.go
1273.788888
1.273789e+03
1273.78888769
1.273789E+03
1273.78888769
-------------------------
1273.7888876900
1.2737888877e+03
1273.788888
1.2737888877E+03
1273.788888
-------------------------
66000000000.120003
6.600000e+10
6.600000000012e+10
6.600000E+10
6.600000000012E+10

Go string format flags

The flags is a set of characters that modify the output format. The set of valid flags depends on the specifier character. Perl recognizes the following flags:

    space    prefix non-negative number with a space
    +    prefix non-negative number with a plus sign
    -    left-justify within the field
    0    use zeros, not spaces, to right-justify
    #    puts the leading 0 for any octal, prefix non-zero hexadecimal with 0x or 0X, prefix non-zero binary with 0b

flags.go

package main

import (
    "fmt"
)

func main() {

    fmt.Printf("%+d\n", 1691)

    fmt.Println("---------------------")

    fmt.Printf("%#x\n", 1691)
    fmt.Printf("%#X\n", 1691)
    fmt.Printf("%#b\n", 1691)

    fmt.Println("---------------------")

    fmt.Printf("%10d\n", 1691)
    fmt.Printf("%-10d\n", 1691)
    fmt.Printf("%010d\n", 1691)
}

The example uses flags in the string format specifier.

$ go run flags.go
+1691
---------------------
0x69b
0X69B
0b11010011011
---------------------
      1691
1691
0000001691

Go string format width

The width is the minimum number of runes to be output. It is specified by an optional decimal number immediately preceding the verb. If absent, the width is whatever is necessary to represent the value.

If the width is greater than the value, it is padded with spaces.
width.go

package main

import (
    "fmt"
)

func main() {

    w := "falcon"
    n := 122
    h := 455.67

    fmt.Printf("%s\n", w)
    fmt.Printf("%10s\n", w)

    fmt.Println("---------------------")

    fmt.Printf("%d\n", n)
    fmt.Printf("%7d\n", n)
    fmt.Printf("%07d\n", n)

    fmt.Println("---------------------")

    fmt.Printf("%10f\n", h)
    fmt.Printf("%11f\n", h)
    fmt.Printf("%12f\n", h)
}

The examples uses width with a string, integer, and a float.

fmt.Printf("%07d\n", n)

With a preceding 0 character, the number is not padded with space but with 0 character.

$ go run width.go 
falcon
    falcon
---------------------
122
    122
0000122
---------------------
455.670000
 455.670000
  455.670000

In this tutorial, we have covered string formatting in Golang.

List all Go tutorials.
Home Facebook Twitter Github Subscribe Privacy
© 2007 - 2022 Jan Bodnar admin(at)zetcode.com
