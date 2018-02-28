## Value


```golang
package main

import (
	"fmt"
	"regexp"
)

func main() {
	content := []byte(`
	# comment line
	option1: value1
	option2: value2

	# another comment line
	option3: value3
`)

	
	pattern  := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)
	template := []byte("$key=$value\n")
	result   := []byte{}
	
	for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {
		result = pattern.Expand(result, template, content, submatches)
	}
	fmt.Println(string(result))
}
```

