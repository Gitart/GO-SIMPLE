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

### Output
```
[[1 3 2 2]]
[[1 5 2 4]]
[[1 3 2 2] [4 7 5 6]]
[[1 5 2 4] [6 8 7 7]]
[]
```

### Oписание
Такой же вариант как и первый только выводим индексы



### Вариант с стринг
```golang
func main() {
	re := regexp.MustCompile("foo.?")
	fmt.Printf("%q\n", re.FindString("seafood fool"))
	fmt.Printf("%q\n", re.FindString("meat"))
}
```

### Output
```
"food"
```

### Использование варианта с или

```
func main() {
	re := regexp.MustCompile("a(x*)b(y|z)c")
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))
}
```


### Output
```
["axxxbyc" "xxx" "y"]
["abzc" "" "z"]
```

### Поиск повторений

```golang
func main() {
	re := regexp.MustCompile("(gopher){2}")
	fmt.Println(re.MatchString("gopher"))
	fmt.Println(re.MatchString("gophergopher"))
	fmt.Println(re.MatchString("gophergophergopher"))
}
```
### Output
```
false
true
true
```


## Cоздание темплейта на основании поиска

```golang
func main() {
	re := regexp.MustCompile("a(x*)b")
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "T"))
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "$1"))
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "${1}"))
}
```
### Output
```
-T-T-
-$1-$1-
-${1}-${1}-
```






