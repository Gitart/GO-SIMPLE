# Атомарные операции в `sync/atomic` стандартном пакете

Атомарные операции более примитивны, чем другие методы синхронизации. Они не блокируются и обычно реализуются непосредственно на аппаратном уровне. На самом деле они часто используются при реализации других методов синхронизации.

Обратите внимание, что многие примеры ниже не являются параллельными программами. Они предназначены только для демонстрации и объяснения, чтобы показать, как использовать атомарные функции, предоставляемые в `sync/atomic` стандартном пакете.

### Обзор атомарных операций, представленных до версии Go 1.19-

Стандартный `sync/atomic` пакет предоставляет следующие пять атомарных функций для целочисленного типа , `T` где `T` должны быть любые из `int32` , `int64` , `uint32` и `uint64` . `uintptr`

```go
func AddT(addr *T, delta T)(new T)
func LoadT(addr *T) (val T)
func StoreT(addr *T, val T)
func SwapT(addr *T, new T) (old T)
func CompareAndSwapT(addr *T, old, new T) (swapped bool)

```

Например, для type предусмотрены следующие пять функций `int32` .

```go
func AddInt32(addr *int32, delta int32)(new int32)
func LoadInt32(addr *int32) (val int32)
func StoreInt32(addr *int32, val int32)
func SwapInt32(addr *int32, new int32) (old int32)
func CompareAndSwapInt32(addr *int32,
				old, new int32) (swapped bool)

```

Следующие четыре элементарные функции предоставляются для (безопасных) типов указателей. Когда эти функции были введены в стандартную библиотеку, Go не поддерживал пользовательские дженерики, поэтому эти функции реализованы через [небезопасный тип указателя](https://go101.org/article/unsafe.html) `unsafe.Pointer` (аналог C в Go `void*` ).

```go
func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)
func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)
func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer,
				) (old unsafe.Pointer)
func CompareAndSwapPointer(addr *unsafe.Pointer,
				old, new unsafe.Pointer) (swapped bool)

```

Для указателей нет `AddPointer` функции, поскольку указатели Go (безопасные) не поддерживают арифметические операции.

Стандартный `sync/atomic` пакет также предоставляет тип `Value` , соответствующий типу указателя которого `*Value` имеет четыре метода (перечислены ниже, последние два были введены в Go 1.17). Мы можем использовать эти методы для выполнения атомарных операций со значениями любого типа.

```go
func (*Value) Load() (x interface{})
func (*Value) Store(x interface{})
func (*Value) Swap(new interface{}) (old interface{})
func (*Value) CompareAndSwap(old, new interface{}) (swapped bool)

```

### Обзор новых атомарных операций, появившихся начиная с Go 1.19

Go 1.19 представил несколько типов, каждый из которых владеет набором методов атомарных операций, для достижения тех же эффектов, что и функции уровня пакета, перечисленные в последнем разделе.

Среди этих типов , , `Int32` и предназначены для целочисленных атомарных операций. Методы типа перечислены ниже. Методы остальных четырех типов представлены аналогичным образом. `Int64` `Uint32` `Uint64` `Uintptr` `atomic.Int32`

```go
func (*Int32) Add(delta int32) (new int32)
func (*Int32) Load() int32
func (*Int32) Store(val int32)
func (*Int32) Swap(new int32) (old int32)
func (*Int32) CompareAndSwap(old, new int32) (swapped bool)

```

Начиная с Go 1.18, Go уже поддерживает пользовательские дженерики. И некоторые стандартные пакеты начали принимать пользовательские дженерики начиная с Go 1.19. Пакет `sync/atomic` является одним из этих пакетов. Тип `Pointer[T any]` , представленный в этом пакете Go 1.19, является универсальным типом. Его методы перечислены ниже.

```go
(*Pointer[T]) Load() *T
(*Pointer[T]) Store(val *T)
(*Pointer[T]) Swap(new *T) (old *T)
(*Pointer[T]) CompareAndSwap(old, new *T) (swapped bool)

```

Go 1.19 также представил `Bool` тип для выполнения логических атомарных операций.

### Атомарные операции для целых чисел

Оставшаяся часть этой статьи показывает несколько примеров использования атомарных операций, предоставляемых в Go.

В следующем примере показано, как выполнить `Add` атомарную операцию над `int32` значением с помощью `AddInt32` функции. В этом примере основная горутина создает 1000 новых параллельных горутин. Каждая вновь созданная горутина увеличивает целое число `n` на единицу. Атомарные операции гарантируют отсутствие гонок данных среди этих горутин. В конце концов, `1000` гарантировано будет напечатано.

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var n int32
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&n, 1)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(atomic.LoadInt32(&n)) // 1000
}

```

Если выражение `atomic.AddInt32(&n, 1)` заменить на `n++` , то вывод может быть не `1000` .

Следующий код повторно реализует вышеуказанную программу, используя `atomic.Int32` тип и его методы (начиная с Go 1.19). Этот код выглядит немного аккуратнее.

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var n atomic.Int32
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			n.Add(1)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(n.Load()) // 1000
}

```

Атомарные `StoreT` функции `LoadT` /методы часто используются для реализации методов установки и получения (соответствующего типа указателя) типа, если значения типа необходимо использовать одновременно. Например, версия функции:

```go
type Page struct {
	views uint32
}

func (page *Page) SetViews(n uint32) {
	atomic.StoreUint32(&page.views, n)
}

func (page *Page) Views() uint32 {
	return atomic.LoadUint32(&page.views)
}

```

И версия type+methods (начиная с Go 1.19):

```go
type Page struct {
	views atomic.Uint32
}

func (page *Page) SetViews(n uint32) {
	page.views.Store(n)
}

func (page *Page) Views() uint32 {
	return page.views.Load()
}

```

Для целочисленного типа со знаком `T` ( `int32` или `int64` ) второй аргумент для вызова `AddT` функции может быть отрицательным значением, чтобы выполнить операцию атомарного уменьшения. Но как выполнять операции атомарного уменьшения для значений беззнакового типа , `T` таких как `uint32` и `uint64` ? `uintptr` Есть два обстоятельства для второго беззнакового аргумента.

1.  Для беззнаковой переменной `v` типа `T` допустимо `-v` в Go. Таким образом, мы можем просто передать `-v` его как второй аргумент `AddT` вызова.
2.  Для положительного постоянного целого числа `c` запрещено `-c` использовать в качестве второго аргумента `AddT` вызова (где `T` обозначает беззнаковый целочисленный тип). Вместо этого мы можем использовать его `^T(c-1)` в качестве второго аргумента.

Этот `^T(v-1)` трюк также работает для беззнаковой переменной `v` , но `^T(v-1)` менее эффективен, чем `T(-v)` .

В трюке `^T(c-1)` , если `c` это типизированное значение и его тип точно `T` , то форма может быть сокращена как `^(c-1)` .

Пример:

```go
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var (
		n uint64 = 97
		m uint64 = 1
		k int    = 2
	)
	const (
		a        = 3
		b uint64 = 4
		c uint32 = 5
		d int    = 6
	)

	show := fmt.Println
	atomic.AddUint64(&n, -m)
	show(n) // 96 (97 - 1)
	atomic.AddUint64(&n, -uint64(k))
	show(n) // 94 (96 - 2)
	atomic.AddUint64(&n, ^uint64(a - 1))
	show(n) // 91 (94 - 3)
	atomic.AddUint64(&n, ^(b - 1))
	show(n) // 87 (91 - 4)
	atomic.AddUint64(&n, ^uint64(c - 1))
	show(n) // 82 (87 - 5)
	atomic.AddUint64(&n, ^uint64(d - 1))
	show(n) // 76 (82 - 6)
	x := b; atomic.AddUint64(&n, -x)
	show(n) // 72 (76 - 4)
	atomic.AddUint64(&n, ^(m - 1))
	show(n) // 71 (72 - 1)
	atomic.AddUint64(&n, ^uint64(k - 1))
	show(n) // 69 (71 - 2)
}

```

Вызов `SwapT` функции похож на `StoreT` вызов функции, но возвращает старое значение.

Вызов `CompareAndSwapT` функции применяет операцию сохранения только тогда, когда текущее значение соответствует переданному старому значению. `bool` Возвращаемый результат `CompareAndSwapT` вызова функции указывает, применяется ли операция сохранения .

Пример:

```go
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var n int64 = 123
	var old = atomic.SwapInt64(&n, 789)
	fmt.Println(n, old) // 789 123
	swapped := atomic.CompareAndSwapInt64(&n, 123, 456)
	fmt.Println(swapped) // false
	fmt.Println(n)       // 789
	swapped = atomic.CompareAndSwapInt64(&n, 789, 456)
	fmt.Println(swapped) // true
	fmt.Println(n)       // 456
}

```

Ниже приведена соответствующая версия type+methods (начиная с Go 1.19):

```go
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var n atomic.Int64
	n.Store(123)
	var old = n.Swap(789)
	fmt.Println(n.Load(), old) // 789 123
	swapped := n.CompareAndSwap(123, 456)
	fmt.Println(swapped)  // false
	fmt.Println(n.Load()) // 789
	swapped = n.CompareAndSwap(789, 456)
	fmt.Println(swapped)  // true
	fmt.Println(n.Load()) // 456
}

```

Обратите внимание, что до сих пор (Go 1.19) атомарные операции для 64-битных слов, также называемых значениями int64 и uint64, требовали, чтобы 64-битные слова были выровнены по 8 байтам в памяти. В Go 1.19 введены операции атомарного метода, это требование всегда выполняется как в 32-битной, так и в 64-битной архитектуре, но это не так для операций атомарной функции в 32-битной архитектуре. Пожалуйста, ознакомьтесь с информацией о [схеме памяти](https://go101.org/article/memory-layout.html) .

### Атомарные операции для указателей

Выше упоминалось, что в стандартном пакете есть четыре функции `sync/atomic` для выполнения операций с атомарными указателями с помощью небезопасных указателей.

Из статьи [type-unsafe pointers](https://go101.org/article/unsafe.html) мы узнали, что в Go значения любого типа указателя могут быть явно преобразованы в `unsafe.Pointer` , и наоборот. Таким образом, значения `*unsafe.Pointer` типа также могут быть явно преобразованы в `unsafe.Pointer` , и наоборот.

Следующий пример не является параллельной программой. Он просто показывает, как выполнять операции с атомарными указателями. В этом примере `T` может быть произвольным типом.

```go
package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type T struct {x int}

func main() {
	var pT *T
	var unsafePPT = (*unsafe.Pointer)(unsafe.Pointer(&pT))
	var ta, tb = T{1}, T{2}
	// store
	atomic.StorePointer(
		unsafePPT, unsafe.Pointer(&ta))
	fmt.Println(pT) // &{1}
	// load
	pa1 := (*T)(atomic.LoadPointer(unsafePPT))
	fmt.Println(pa1 == &ta) // true
	// swap
	pa2 := atomic.SwapPointer(
		unsafePPT, unsafe.Pointer(&tb))
	fmt.Println((*T)(pa2) == &ta) // true
	fmt.Println(pT) // &{2}
	// compare and swap
	b := atomic.CompareAndSwapPointer(
		unsafePPT, pa2, unsafe.Pointer(&tb))
	fmt.Println(b) // false
	b = atomic.CompareAndSwapPointer(
		unsafePPT, unsafe.Pointer(&tb), pa2)
	fmt.Println(b) // true
}

```

Да, использование атомарных функций указателя довольно многословно. На самом деле, использование не только многословно, но и не защищено [рекомендациями по совместимости с Go 1](https://golang.org/doc/go1compat) , поскольку для этих применений требуется импорт `unsafe` стандартного пакета.

Напротив, код будет намного проще и чище, если мы будем использовать общий `Pointer` тип Go 1.19 и его методы для выполнения операций с атомарными указателями, как показано в следующем коде.

```go
package main

import (
	"fmt"
	"sync/atomic"
)

type T struct {x int}

func main() {
	var pT atomic.Pointer[T]
	var ta, tb = T{1}, T{2}
	// store
	pT.Store(&ta)
	fmt.Println(pT.Load()) // &{1}
	// load
	pa1 := pT.Load()
	fmt.Println(pa1 == &ta) // true
	// swap
	pa2 := pT.Swap(&tb)
	fmt.Println(pa2 == &ta) // true
	fmt.Println(pT.Load())  // &{2}
	// compare and swap
	b := pT.CompareAndSwap(&ta, &tb)
	fmt.Println(b) // false
	b = pT.CompareAndSwap(&tb, &ta)
	fmt.Println(b) // true
}

```

Что еще более важно, реализация с использованием универсального `Pointer` типа защищена рекомендациями по совместимости с Go 1.

### Атомарные операции для значений произвольных типов

Тип `Value` , предоставленный в `sync/atomic` стандартном пакете, может использоваться для атомарной загрузки и хранения значений любого типа.

Type `*Value` имеет несколько методов: `Load` , `Store` , `Swap` и `CompareAndSwap` (последние два представлены в Go 1.17). Типы входных параметров этих методов — все `interface{}` . Таким образом, в вызовы этих методов может быть передано любое значение. Но для адресуемого `Value` значения `v` , как только вызов `v.Store()` (сокращение от `(&v).Store()` ) когда-либо был вызван, последующие вызовы метода для значения `v` также должны принимать значения аргументов с тем же [конкретным типом](https://go101.org/article/type-system-overview.html#concrete-type) , что и аргумент первого `v.Store()` вызова, иначе произойдет паника. Аргумент `nil` интерфейса также `v.Store()` вызывает панику вызова.

Пример:

```go
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	type T struct {a, b, c int}
	var ta = T{1, 2, 3}
	var v atomic.Value
	v.Store(ta)
	var tb = v.Load().(T)
	fmt.Println(tb)       // {1 2 3}
	fmt.Println(ta == tb) // true

	v.Store("hello") // will panic
}

```

Другой пример (для Go 1.17+):

```go
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	type T struct {a, b, c int}
	var x = T{1, 2, 3}
	var y = T{4, 5, 6}
	var z = T{7, 8, 9}
	var v atomic.Value
	v.Store(x)
	fmt.Println(v) // {{1 2 3}}
	old := v.Swap(y)
	fmt.Println(v)       // {{4 5 6}}
	fmt.Println(old.(T)) // {1 2 3}
	swapped := v.CompareAndSwap(x, z)
	fmt.Println(swapped, v) // false {{4 5 6}}
	swapped = v.CompareAndSwap(y, z)
	fmt.Println(swapped, v) // true {{7 8 9}}
}

```

На самом деле, мы также можем использовать атомарные функции указателя, описанные в предыдущем разделе, для выполнения атомарных операций со значениями любого типа с еще одним косвенным уровнем. Оба способа имеют свои преимущества и недостатки. Какой способ следует использовать, зависит от требований на практике.

### Гарантия порядка памяти, сделанная Atomic Operations в Go

Для простоты использования атомарные операции Go , входящие в `sync/atomic` стандартный пакет, спроектированы без какой-либо связи с упорядочением памяти. По крайней мере, в официальной документации не указаны какие-либо гарантии порядка памяти, обеспечиваемые `sync/atomic` стандартным пакетом. Подробную информацию см. в разделе Модель памяти [Go](https://go101.org/article/memory-model.html#atomic) .
