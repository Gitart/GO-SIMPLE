# Strings in Golang

View Discussion

Improve Article

Save Article

Like Article

*   Last Updated : 09 Aug, 2019

*   Read
*   Discuss

View Discussion

Improve Article

Save Article

Like Article

In Go language, strings are different from other languages like [Java](https://www.geeksforgeeks.org/java/#Strings%20in%20Java), [C++](https://www.geeksforgeeks.org/c-plus-plus/#Arrays%20and%20Strings), [Python](https://www.geeksforgeeks.org/python-strings/), etc. it is a sequence of variable-width characters where each and every character is represented by one or more bytes using [UTF-8 Encoding](https://en.wikipedia.org/wiki/UTF-8?). Or in other words, strings are the immutable chain of arbitrary bytes(including bytes with zero value) or ***string is a read-only slice of bytes*** and the bytes of the strings can be represented in the Unicode text using UTF-8 encoding.
Due to UTF-8 encoding Golang string can contain a text which is the mixture of any language present in the world, without any confusion and limitation of the page. Generally, strings are enclosed in *double-quotes””*, as shown in the below example:

**Example:**

|

`// Go program to illustrate`

`// how to create strings`

`package main`

`import` `"fmt"`

`func main() {`

`// Creating and initializing a`

`// variable with a string`

`// Using shorthand declaration`

`My_value_1 :=` `"Welcome to GeeksforGeeks"`

`// Using var keyword`

`var My_value_2 string`

`My_value_2 =` `"GeeksforGeeks"`

`// Displaying strings`

`fmt.Println(``"String 1: "``, My_value_1)`

`fmt.Println(``"String 2: "``, My_value_2)`

`}`

 |

**Output:**

String 1:  Welcome to GeeksforGeeks
String 2:  GeeksforGeeks

**Note:** String can be empty, but they are not nil.

#### String Literals

In Go language, string literals are created in two different ways:

*   **Using double quotes(“”):** Here, the string literals are created using double-quotes(“”). This type of string support escape character as shown in the below table, but does not span multiple lines. This type of string literals is widely used in Golang programs.

    | Escape character | Description |
    | --- | --- |
    | **\\\\** | Backslash(\\) |
    | **\\000** | Unicode character with the given 3-digit 8-bit octal code point |
    | **\\’** | Single quote (‘). It is only allowed inside character literals |
    | **\\”** | Double quote (“). It is only allowed inside interpreted string literals |
    | **\\a** | ASCII bell (BEL) |
    | **\\b** | ASCII backspace (BS) |
    | **\\f** | ASCII formfeed (FF) |
    | **\\n** | ASCII linefeed (LF |
    | **\\r** | ASCII carriage return (CR) |
    | **\\t** | ASCII tab (TAB) |
    | **\\uhhhh** | Unicode character with the given 4-digit 16-bit hex code point. |
    |  | Unicode character with the given 8-digit 32-bit hex code point. |
    | **\\v** | ASCII vertical tab (VT) |
    | **\\xhh** | Unicode character with the given 2-digit 8-bit hex code point. |

*   **Using backticks(“):** Here, the string literals are created using backticks(“) and also known as **`raw literals`**. Raw literals do not support escape characters, can span multiple lines, and may contain any character except backtick. It is, generally, used for writing multiple line message, in the regular expressions, and in HTML.

    **Example:**

    |

    `// Go program to illustrate string literals`

    `package main`

    `import` `"fmt"`

    `func main() {`

    `// Creating and initializing a`

    `// variable with a string literal`

    `// Using double-quote`

    `My_value_1 :=` `"Welcome to GeeksforGeeks"`

    `// Adding escape character`

    `My_value_2 :=` `"Welcome!\nGeeksforGeeks"`

    `// Using backticks`

    ``My_value_3 := `Hello!GeeksforGeeks` ``

    `// Adding escape character`

    `// in raw literals`

    ``My_value_4 := `Hello!\nGeeksforGeeks` ``

    `// Displaying strings`

    `fmt.Println(``"String 1: "``, My_value_1)`

    `fmt.Println(``"String 2: "``, My_value_2)`

    `fmt.Println(``"String 3: "``, My_value_3)`

    `fmt.Println(``"String 4: "``, My_value_4)`

    `}`

     |

    **Output:**

    String 1:  Welcome to GeeksforGeeks
    String 2:  Welcome!
    GeeksforGeeks
    String 3:  Hello!GeeksforGeeks
    String 4:  Hello!\\nGeeksforGeeks

#### Important Points About String

*   **Strings are immutable:** In Go language, strings are immutable once a string is created the value of the string cannot be changed. Or in other words, strings are read-only. If you try to change, then the compiler will throw an error.

    **Example:**

    |

    `// Go program to illustrate`

    `// string are immutable`

    `package main`

    `import` `"fmt"`

    `// Main function`

    `func main() {`

    `// Creating and initializing a string`

    `// using shorthand declaration`

    `mystr :=` `"Welcome to GeeksforGeeks"`

    `fmt.Println(``"String:"``, mystr)`

    `/* if you trying to change`

    `the value of the string`

    `then the compiler will`

    `throw an error, i.e,`

    `cannot assign to mystr[1]`

    `mystr[1]= 'G'`

    `fmt.Println("String:", mystr)`

    `*/`

    `}`

     |

    **Output:**

    String: Welcome to GeeksforGeeks

*   **How to iterate over a string?:** You can iterate over string using for rang loop. This loop can iterate over the Unicode code point for a string.

    **Syntax:**

    for index, chr:= range str{
         // Statement..
    }

    Here, the index is the variable which store the first byte of UTF-8 encoded code point and *chr* store the characters of the given string and str is a string.

    **Example:**

    |

    `// Go program to illustrate how`

    `// to iterate over the string`

    `// using for range loop`

    `package main`

    `import` `"fmt"`

    `// Main function`

    `func main() {`

    `// String as a range in the for loop`

    `for` `index, s := range` `"GeeksForGeeKs"` `{`

    `fmt.Printf(``"The index number of %c is %d\n"``, s, index)`

    `}`

    `}`

     |

    **Output:**

    The index number of G is 0
    The index number of e is 1
    The index number of e is 2
    The index number of k is 3
    The index number of s is 4
    The index number of F is 5
    The index number of o is 6
    The index number of r is 7
    The index number of G is 8
    The index number of e is 9
    The index number of e is 10
    The index number of K is 11
    The index number of s is 12

*   **How to access the individual byte of the string?:** The string is of a byte so, we can access each byte of the given string.

    **Example:**

    |

    `// Go program to illustrate how to`

    `// access the bytes of the string`

    `package main`

    `import` `"fmt"`

    `// Main function`

    `func main() {`

    `// Creating and initializing a string`

    `str :=` `"Welcome to GeeksforGeeks"`

    `// Accessing the bytes of the given string`

    `for` `c := 0; c < len(str); c++ {`

    `fmt.Printf(``"\nCharacter = %c Bytes = %v"``, str, str)`

    `}`

    `}`

     |

    **Output:**

    Character = W Bytes = 87
    Character = e Bytes = 101
    Character = l Bytes = 108
    Character = c Bytes = 99
    Character = o Bytes = 111
    Character = m Bytes = 109
    Character = e Bytes = 101
    Character =   Bytes = 32
    Character = t Bytes = 116
    Character = o Bytes = 111
    Character =   Bytes = 32
    Character = G Bytes = 71
    Character = e Bytes = 101
    Character = e Bytes = 101
    Character = k Bytes = 107
    Character = s Bytes = 115
    Character = f Bytes = 102
    Character = o Bytes = 111
    Character = r Bytes = 114
    Character = G Bytes = 71
    Character = e Bytes = 101
    Character = e Bytes = 101
    Character = k Bytes = 107
    Character = s Bytes = 115

*   **How to create a string form the slice?:** In Go language, you are allowed to create a string from the slice of bytes.

    **Example:**

    |

    `// Go program to illustrate how to`

    `// create a string from the slice`

    `package main`

    `import` `"fmt"`

    `// Main function`

    `func main() {`

    `// Creating and initializing a slice of byte`

    `myslice1 := []byte{0x47, 0x65, 0x65, 0x6b, 0x73}`

    `// Creating a string from the slice`

    `mystring1 := string(myslice1)`

    `// Displaying the string`

    `fmt.Println(``"String 1: "``, mystring1)`

    `// Creating and initializing a slice of rune`

    `myslice2 := []rune{0x0047, 0x0065, 0x0065, `

    `0x006b, 0x0073}`

    `// Creating a string from the slice`

    `mystring2 := string(myslice2)`

    `// Displaying the string`

    `fmt.Println(``"String 2: "``, mystring2)`

    `}`

     |

    **Output:**

    String 1:  Geeks
    String 2:  Geeks

*   **How to find the length of the string?:** In Golang string, you can find the length of the string using two functions one is **len()** and another one is **RuneCountInString()**. The RuneCountInString() function is provided by UTF-8 package, this function returns the total number of rune presents in the string. And the *len()* function returns the number of bytes of the string.

    **Example:**

    |

    `// Go program to illustrate how to`

    `// find the length of the string`

    `package main`

    `import (`

    `"fmt"`

    `"unicode/utf8"`

    `)`

    `// Main function`

    `func main() {`

    `// Creating and initializing a string`

    `// using shorthand declaration`

    `mystr :=` `"Welcome to GeeksforGeeks ??????"`

    `// Finding the length of the string`

    `// Using len() function`

    `length1 := len(mystr)`

    `// Using RuneCountInString() function`

    `length2 := utf8.RuneCountInString(mystr)`

    `// Displaying the length of the string`

    `fmt.Println(``"string:"``, mystr)`

    `fmt.Println(``"Length 1:"``, length1)`

    `fmt.Println(``"Length 2:"``, length2)`

    `}`

     |

    **Output:**

    string: Welcome to GeeksforGeeks ??????
    Length 1: 31
    Length 2: 31
