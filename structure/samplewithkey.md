## Key

```golang

package main

import "fmt"

type myStruct struct {
	atrib1 string
	atrib2 string
}

type mapKey struct {
	Key    string
	Option string
}

func main() {
	apiKeyTypeRequest := make(map[mapKey]myStruct)

	apiKeyTypeRequest[mapKey{"Key", "MyFirstOption"}] = myStruct{"first Value first op", "second Value first op"}
	apiKeyTypeRequest[mapKey{"Key", "MysecondtOption"}] = myStruct{atrib1: "first Value second op"}

	fmt.Printf("%+v\n", apiKeyTypeRequest)
}
```
