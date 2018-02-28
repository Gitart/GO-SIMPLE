## Поиск

```golang
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("a(x*)b")
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-abtt-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-axb-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("axxx-axxxb-ab-ffxx", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("axxxb-ab-axxbggg-aaxxxxb", -1))
}
```

### Output
```
[["ab" ""]]
[["axxb" "xx"]]
[["ab" ""] ["axb" "x"]]
[["axxxb" "xxx"] ["ab" ""]]
[["axxxb" "xxx"] ["ab" ""] ["axxb" "xx"] ["axxxxb" "xxxx"]]
```

### Oписание
Находим все сочетания **ab** и если есть все сочетания **axb** и разбиваем на два объекта равное количеству сочетаний

