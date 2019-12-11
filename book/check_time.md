## Регулярное выражение для совпадения формата времени ЧЧ: ММ в Голанге

```golang
package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "8:2"
	str2 := "9:9"
	str3 := "12:29"
	str4 := "02:5"
	str5 := "23:59"
	str6 := "55:59"
	str7 := "0:01"

	re := regexp.MustCompile(`^([0-9]|0[0-9]|1[0-9]|2[0-3]):([0-9]|[0-5][0-9])$`)

	fmt.Printf("Pattern: %v\n", re.String()) // print pattern
	fmt.Printf("Time: %v\t:%v\n", str1, re.MatchString(str1))
	fmt.Printf("Time: %v\t:%v\n", str2, re.MatchString(str2))
	fmt.Printf("Time: %v\t:%v\n", str3, re.MatchString(str3))
	fmt.Printf("Time: %v\t:%v\n", str4, re.MatchString(str4))
	fmt.Printf("Time: %v\t:%v\n", str5, re.MatchString(str5))
	fmt.Printf("Time: %v\t:%v\n", str6, re.MatchString(str6))
	fmt.Printf("Time: %v\t:%v\n", str7, re.MatchString(str7))
}
```

Regular expression to validate the date format in "dd/mm/yyyy"
// Regular expression validate the date format in "dd/mm/yyyy"

```golang
package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "31/07/2010"
	str2 := "1/13/2010"
	str3 := "29/2/2007"
	str4 := "31/08/2010"
	str5 := "29/02/200a"
	str6 := "29/02/200a"
	str7 := "55/02/200a"
	str8 := "2_/02/2009"

	re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")

	fmt.Printf("Pattern: %v\n", re.String()) // print pattern
	fmt.Printf("\nDate: %v :%v\n", str1, re.MatchString(str1))
	fmt.Printf("Date: %v :%v\n", str2, re.MatchString(str2))
	fmt.Printf("Date: %v :%v\n", str3, re.MatchString(str3))
	fmt.Printf("Date: %v :%v\n", str4, re.MatchString(str4))
	fmt.Printf("Date: %v :%v\n", str5, re.MatchString(str5))
	fmt.Printf("Date: %v :%v\n", str6, re.MatchString(str6))
	fmt.Printf("Date: %v :%v\n", str7, re.MatchString(str7))
	fmt.Printf("Date: %v :%v\n", str8, re.MatchString(str8))
}
```
