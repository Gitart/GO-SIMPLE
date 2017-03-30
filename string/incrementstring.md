## Increment string example


### Problem:
You need to create "copies" of a file, prevent user from overwriting files accidentally or n  
eed duplicate database content which has unique titles or slugs. Having increment string function   
is helpful in this kind of situation. How to do that?  

### Solution:
Increments a string by appending a number to it or increasing the number.
Below is a function to increment the string value that you can use or adapt to your own Golang code.

```golang
 package main

 import (
 	"fmt"
 	"os"
 	"strconv"
 	"strings"
 )

 func incrementString(str string, separator string, first int) string {

 	// set default values
 	// see https://www.socketloop.com/tutorials/golang-proper-way-to-set-function-argument-default-value

 	if separator == "" {
 		separator = "_"
 	}

 	if first == 0 || first < 0 {
 		first = 1
 	}

 	// test to see if str already has integer suffix(ends with _#)
 	test := strings.SplitN(str, separator, 2)

 	if len(test) >= 2 {
 		// increase file counter by 1
 		i, err := strconv.Atoi(test[1])

 		if err != nil {
 			fmt.Println(err)
 			os.Exit(1)
 		}

 		increased := i + first
 		return test[0] + separator + strconv.Itoa(increased)
 	} else {
 		return str + separator + strconv.Itoa(first)
 	}
 }

 func main() {

 	result := incrementString("file", "_", 0)
 	fmt.Println(result)

 	result = incrementString("file", "-", 2)
 	fmt.Println(result)

 	// increase by 1
 	result = incrementString("file_2", "", 1)
 	fmt.Println(result)

 	// increase by 2
 	result = incrementString("file_2", "", 2)
 	fmt.Println(result)

 	// increase by 100
 	result = incrementString("file_3", "", 100)
 	fmt.Println(result)

 	// change separator to # sign
 	result = incrementString("imagefiles", "#", 10)
 	fmt.Println(result)

 	// will NOT accept negative number. will change to default value of 1
 	result = incrementString("imagefiles", "#", -99)
 	fmt.Println(result)

 }
 ```
