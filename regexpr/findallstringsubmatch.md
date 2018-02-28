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

## Вариант второй

```golang
func main() {
	re := regexp.MustCompile("a(x*)b")
	// Indices:
	//    01234567   012345678
	//    -ab-axb-   -axxb-ab-
	fmt.Println(re.FindAllStringSubmatchIndex("-ab-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-axxb-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-ab-axb-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-axxb-ab-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-foo-", -1))
}
```
### Oписание
Такой же вариант как и первый только выводим индексы


