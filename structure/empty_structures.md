# Пустые cтруктуры Go

Posted on 2019, Jan 15 3 мин. чтения

# Что это такое

Под пустой структурой имеется ввиду конструкция вида

```go
type Empty struct{}
// или
var a struct{}
```

т.е. структура без полей

# Размер

Пакадж `unsafe`, из стандартной библиотеки go, позволяет получить размер переменных в байтах

```go
var i int32
var s string
var b bool

unsafe.Sizeof(i)  // 4 байта
unsafe.Sizeof(s)  // 8 байт
unsafe.Sizeof(b)  // 1 байт
```

[https://play.golang.org/p/uARBi60z\_E1](https://play.golang.org/p/uARBi60z_E1)

посмотрим размер для пустых структур

```go
type Empty struct{}
var a struct{}
var b Empty

unsafe.Sizeof(a)  // 0 байт
unsafe.Sizeof(b)  // 0 байт
```

[https://play.golang.org/p/4sLaRthBBlB](https://play.golang.org/p/4sLaRthBBlB)

Первый вопрос который приходит в голову \- как такое возможно?

Посмотрим адрес пустой структуры (в смысле переменной).

```go
	type Empty struct{}
	var a struct{}
	var b Empty
	c := struct{}{}

	fmt.Println(unsafe.Pointer(&a)) // 0x1b5498
	fmt.Println(unsafe.Pointer(&b)) // 0x1b5498
	fmt.Println(unsafe.Pointer(&c)) // 0x1b5498
```

[https://play.golang.org/p/MwL7RPOrtDZ](https://play.golang.org/p/MwL7RPOrtDZ)

Видно, что это какая\-то константа

Если заглянуть в файл [https://golang.org/src/runtime/malloc.go](https://golang.org/src/runtime/malloc.go) ,то там есть глобальная переменная

```go
// base address for all 0-byte allocations
var zerobase uintptr
```

и дальше в функции выделяющей память

```go
// Allocate an object of size bytes.
// Small objects are allocated from the per-P cache's free lists.
// Large objects (> 32 kB) are allocated straight from the heap.
func mallocgc(size uintptr, typ *_type, needzero bool) unsafe.Pointer {
	if gcphase == _GCmarktermination {
		throw("mallocgc called with gcphase == _GCmarktermination")
	}

	if size == 0 {
		return unsafe.Pointer(&zerobase)
	}
```

т.е. действительно ноль.

# Использование, как имплементация интерфейса

Из это следуют любопытные вещи. Начнем с “интуитивной” \- если нам нужно просто заимплементить интерфейс, а данные не важны, то есть смысл использовать пустую структуру.

```go
type NeedInterface interface {
	MethodA()
}

type Empty struct {
}

func (Empty) MethodA() {
}
```

если не знать особеннойстей с размерами структур, то есть соблазн “прицепить” метод к простому типу, например `int32` но такое решение займет на 4 байта больше

```go
type myInt int32

func (myInt) MethodA() {
	fmt.Println("method A")
}

type empty struct{}

func (empty) MethodB() {
	fmt.Println("method B")
}

func main() {
	var i myInt
	i.MethodA()

	var e empty
	e.MethodB()

	fmt.Println(unsafe.Sizeof(i))
	fmt.Println(unsafe.Sizeof(e))
}
```

[https://play.golang.org/p/etp2PPNTYOK](https://play.golang.org/p/etp2PPNTYOK)

# Использование, как Set

Иногда нужна структура данных, которая может хранить уникальные значения, при этом порядок, в котором расположены элементы не важен. Можно использовать встроенный `map`, где элементы \- это ключи, а значения \- `struct{}`

```go
type intSet map[int]struct{} // Set для интов

func (p intSet) Add(i int) { // (p intSet) - ссылка, т.к. map
	p[i] = struct{}{}
}

func main() {
	uniqInts := make(intSet, 0)
	uniqInts.Add(1)
	uniqInts.Add(2)
	uniqInts.Add(1) // повтор
	uniqInts.Add(1) // повтор

    // выведет только 2 значения
	for i, _ := range uniqInts {
		fmt.Println(i)
	}
}
```

[https://play.golang.org/p/vedMA78eJQQ](https://play.golang.org/p/vedMA78eJQQ)

# Использование с каналами

Если мы используем каналы для передачи сигналов между горутинами, то есть смысл использовать не `bool`, но `struct{}`
т.к.

```go
	var s struct{}
	var b bool

	fmt.Println(unsafe.Sizeof(s)) // 0 байт
	fmt.Println(unsafe.Sizeof(b)) // 1 байт
```

[https://play.golang.org/p/QMo7jT\-UYY1](https://play.golang.org/p/QMo7jT-UYY1)

Наш сигнальный канал

`var signal chan struct{}`

и доспустим какая\-то длительная операция в горутине, которую нужно дождаться

```go
var signal = make(chan struct{})

func doWork() {
	for i := 0; i < 20; i++ {
		fmt.Println("work:", i)
	}
	//signal <- struct{}{}
	close(signal)
}
func main() {
	go doWork()
	<- signal
	fmt.Println("completed")
}
```

[https://play.golang.org/p/QqlF5de5GdT](https://play.golang.org/p/QqlF5de5GdT)

напомню, что `<- signal`

блокируется, пока мы не передадим значение в канал

`signal <- struct{}{}`

либо не закроем его

`close(signal)`

тогда блокировка тоже снимается

Ссылки:

*   [https://dave.cheney.net/2014/03/25/the\-empty\-struct](https://dave.cheney.net/2014/03/25/the-empty-struct)
*   [https://povilasv.me/go\-memory\-management/](https://povilasv.me/go-memory-management/)
*   [https://blog.golang.org/ismmkeynote](https://blog.golang.org/ismmkeynote)
*   [https://medium.com/@matryer/cool\-golang\-trick\-nil\-structs\-can\-just\-be\-a\-collection\-of\-methods\-741ae57ab262](https://medium.com/@matryer/cool-golang-trick-nil-structs-can-just-be-a-collection-of-methods-741ae57ab262)
*   [https://medium.com/@matryer/golang\-advent\-calendar\-day\-two\-starting\-and\-stopping\-things\-with\-a\-signal\-channel\-f5048161018](https://medium.com/@matryer/golang-advent-calendar-day-two-starting-and-stopping-things-with-a-signal-channel-f5048161018)
*   [множество set](https://ru.wikipedia.org/wiki/%D0%9C%D0%BD%D0%BE%D0%B6%D0%B5%D1%81%D1%82%D0%B2%D0%BE_(%D1%82%D0%B8%D0%BF_%D0%B4%D0%B0%D0%BD%D0%BD%D1%8B%D1%85)#%D0%9E%D0%B1%D1%8A%D0%B5%D0%BA%D1%82_Set)
