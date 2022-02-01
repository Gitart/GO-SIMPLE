## Наполнение Bytes 
[String byte Go Palygraond](https://go.dev/play/p/Rhx9-WuSqdy)

```go
package main

import (
	"bytes"
	"fmt"
)

func main() {

	// Creating and initializing strings
	// Using bytes.Buffer with
	// WriteString() function
	var b bytes.Buffer

	b.WriteString("<h1>Hello</h1>")
	b.WriteString("<p>Сеттинг</p>")
	b.WriteString("<h>ssssssss</h1>")
	b.WriteString("kw---------w")
	b.WriteString("wwws")

	fmt.Println("String: ", b.String())

	b.WriteString("f")
	b.WriteString("o")
	b.WriteString("r")
	b.WriteString("G")
	b.WriteString("e")
	b.WriteString("e")
	b.WriteString("k")
	b.WriteString("s")

	fmt.Println("String: ", b.String())

}
```

```
String:  <h1>Hello</h1><p>Сеттинг</p><h>ssssssss</h1>kw---------wwwws
String:  <h1>Hello</h1><p>Сеттинг</p><h>ssssssss</h1>kw---------wwwwsforGeeks
```

