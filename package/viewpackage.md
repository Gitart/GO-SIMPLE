# Видимость имен в go на уровне файлов пакетов

Posted on 2015, Oct 05 2 мин. чтения

В go есть возможность ограничить видимость типа/метода/объекта в зависимости с большой или маленькой буквы начинается имя. Если с маленькой, то видимость только внутри пакета и разные пакеты могут одинаковые имена внутри использовать. Интересно это же правило действует для разных файлов внутри 1 пакета? Т.е. допустим у меня есть пакет user и в нем 2 файла generate.go и auth.go и в каждом хотелось бы свой тип

```go
type request struct {
   ...
}
```

А не называть их GenerateRequest, AuthRequest и т.д.

Так напишем тест

## helloworld.go

```go
package main

import "fmt"

func main() {
	fmt.Printf("Hello, world.\n")
}
```

## separate.go

```go
package main

func TestUpper(){
	fmt.Printf("I'm visible!\n")
}

func testLower(){
	fmt.Printf("I'm visible!\n")
}
```

`go build`

```
# helloworld
./separate.go:4: undefined: fmt in fmt.Printf
./separate.go:8: undefined: fmt in fmt.Printf

```

ага, т.е. в каждом файле нужно все импорты прописывать, ничего не мержится. Оч, хорошо.

## helloworld.go

```go
package main

import "fmt"

func main() {
	fmt.Printf("Hello, world.\n")
	TestUpper()
}
```

## separate.go

```go
package main

import "fmt"

func TestUpper(){
	fmt.Printf("I'm visible!\n")
}

func testLower(){
	fmt.Printf("I'm visible!\n")
}
```

`go build`

Собирается, как ожидалось. Так, добавим метод начинающийся с маленькой буквы.

## helloworld.go

```go
package main

import "fmt"

func main() {
	fmt.Printf("Hello, world.\n")
	TestUpper()
	testLower()
}
```

`go build`

Собирается, жаль. Т.е. видимость с помощью имен можно только на уровне пакетов (package) ограничивать.

А если такой же метод добавить?

## helloworld.go

```go
package main

import "fmt"

func main() {
	fmt.Printf("Hello, world.\n")
	TestUpper()
	testLower()
}

func testLower(){
	fmt.Printf("I'm visible!\n")
}
```

`go build`

```
# helloworld
./separate.go:9: testLower redeclared in this block
	previous declaration at ./helloworld.go:11
```
