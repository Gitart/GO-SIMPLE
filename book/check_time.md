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

