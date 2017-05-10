## Serch element in Array

```golang
package main

import (
	"fmt"
)

func main() {

parts := []string{"X15", "T14", "X23", "A41", "L19", "X57", "A63"}
var Xparts []string

Filter(len(parts), func(i int) bool { return parts[i][0] == 'X' },

func(i int) { Xparts = append(Xparts, parts[i]) })

fmt.Println(Xparts)
}

func Filter(limit int, predicate func(int) bool, appender func(int)) {
for i := 0; i < limit; i++ {
if predicate(i) {
 appender(i)
 }
 }
}
```
