## Regex to extract image name from HTML in Golang

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := `<img src="1.png"><x><z?>
			 <img czx zcxz src='2.png'><x><z?>`

	re := regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)

	submatchall := re.FindAllStringSubmatch(str1, -1)
	for _, element := range submatchall {
		fmt.Println(element[1])
	}
}
```
