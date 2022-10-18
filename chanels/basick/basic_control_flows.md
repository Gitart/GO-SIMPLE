# Основные потоки управления

Блоки кода потока управления в Go очень похожи на другие популярные языки программирования, но есть и много отличий. Эта статья покажет эти сходства и различия.

### Введение в потоки управления в Go

В Go есть три вида основных блоков кода потока управления:

*   `if-else` двусторонний условный блок выполнения.
*   `for` блок петель.
*   `switch-case` многоходовой условный блок выполнения.

Есть также несколько блоков кода потока управления, которые связаны с определенными типами в Go.

*   `for-range` блок цикла для типов [контейнеров .](https://go101.org/article/container.html#iteration)
*   `type-switch` многосторонний условный блок выполнения для [интерфейсных](https://go101.org/article/interface.html#type-switch) типов.
*   `select-case` блок для типов [каналов .](https://go101.org/article/channel.html#select)

Как и многие другие популярные языки, Go также поддерживает `break` операторы перехода `continue` и `goto` выполнения кода. Помимо этого, в Go есть специальный оператор перехода по коду, `fallthrough` .

Среди шести видов блоков потока управления, за исключением `if-else` потока управления, остальные пять называются **разрушаемыми блоками потока управления** . Мы можем использовать `break` операторы, чтобы заставить выполнение выпрыгивать из разрушаемых блоков потока управления.

`for` а `for-range` блоки цикла называются **блоками потока управления циклом** . Мы можем использовать `continue` операторы, чтобы закончить шаг цикла заранее в блоке потока управления циклом, т.е. перейти к следующей итерации цикла.

Обратите внимание, что каждый из упомянутых выше блоков потока управления является оператором и может содержать множество других подоператоров.

Вышеупомянутые операторы потока управления являются все в узком смысле. Механизмы, представленные в следующей статье, [горутины, отложенные вызовы функций и panic/recover](https://go101.org/article/control-flows-more.html) , а также методы синхронизации параллелизма, представленные в [обзоре синхронизации параллелизма](https://go101.org/article/concurrent-synchronization-overview.html) в более поздней статье , можно рассматривать как операторы потока управления в широком смысле.

В этой статье будут объяснены только основные блоки кода потока управления и операторы перехода кода, другие будут объяснены во многих других статьях Go 101 позже.

### `if-else` Блоки потока управления

Полная форма `if-else` блока кода похожа на

```go
if InitSimpleStatement; Condition {
	// do something
} else {
	// do something
}

```

`if` и `else` являются ключевыми словами. Как и во многих других языках программирования, `else` ветвь не является обязательной.

Часть `InitSimpleStatement` также необязательна. Это должно быть [простое утверждение](https://go101.org/article/expressions-and-statements.html#simple-statements) , если оно присутствует. Если он отсутствует, мы можем рассматривать его как пустой оператор (один из видов простых операторов). На практике `InitSimpleStatement` часто используется короткое объявление переменной или чистое присваивание. A `Condition` должно быть [выражением](https://go101.org/article/expressions-and-statements.html#expressions) , результатом которого является логическое значение. Порция `Condition` может быть заключена в пару `()` или нет, но не может быть заключена вместе с `InitSimpleStatement` порцией.

Если `InitSimpleStatement` в `if-else` блоке присутствует, он будет выполнен перед выполнением других операторов в `if-else` блоке. Если `InitSimpleStatement` отсутствует, то точка с запятой после нее не обязательна.

Каждый `if-else` поток управления формирует один неявный блок кода, один `if` явный блок кода ветвления и один необязательный `else` блок кода ветвления. Оба блока кода ответвления вложены в блок неявного кода. При выполнении, если `Condition` выражение приводит к результату `true` , тогда `if` будет выполнен блок ответвления, в противном случае `else` будет выполнен блок ответвления.

Пример:

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if n := rand.Int(); n%2 == 0 {
		fmt.Println(n, "is an even number.")
	} else {
		fmt.Println(n, "is an odd number.")
	}

	n := rand.Int() % 2 // this n is not the above n.
	if n % 2 == 0 {
		fmt.Println("An even number.")
	}

	if ; n % 2 != 0 {
		fmt.Println("An odd number.")
	}
}

```

Если `InitSimpleStatement` в `if-else` блоке кода содержится короткое объявление переменной, то объявленные переменные будут рассматриваться как объявленные в верхнем вложенном блоке неявного кода блока `if-else` кода.

Блок `else` кода ответвления может быть неявным, если за соответствующим блоком `else` следует другой `if-else` блок кода, в противном случае он должен быть явным.

Пример:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	if h := time.Now().Hour(); h < 12 {
		fmt.Println("Now is AM time.")
	} else if h > 19 {
		fmt.Println("Now is evening time.")
	} else {
		fmt.Println("Now is afternoon time.")
		h := h // the right one is declared above
		// The just new declared "h" variable
		// shadows the above same-name one.
		_ = h
	}

	// h is not visible here.
}

```

### `for` Блоки управления циклом

Полная форма `for` блока цикла:

```go
for InitSimpleStatement; Condition; PostSimpleStatement {
	// do something
}

```

`for` является ключевым словом. Части `InitSimpleStatement` и `PostSimpleStatement` должны быть простыми операторами, и `PostSimpleStatement` часть не должна быть коротким объявлением переменной. `Condition` должно быть выражением, результатом которого является логическое значение. Все три части являются необязательными.

В отличие от многих других языков программирования, только что упомянутые три части, следующие за `for` ключевым словом, не могут быть заключены в пару `()` .

Каждый `for` поток управления формирует как минимум два блока кода, один неявный, а другой явный. Явный вложен в неявный.

Блок `InitSimpleStatement` in a `for` loop будет выполнен (только один раз) перед выполнением других операторов в `for` блоке цикла.

Выражение `Condition` будет оцениваться на каждом шаге цикла. Если результат оценки равен `false` , то цикл завершится. В противном случае будет выполнено тело (также известное как явный блок кода) цикла.

Будет `PostSimpleStatement` выполняться в конце каждого шага цикла.

Пример `for` цикла. В примере будут напечатаны целые числа от `0` до `9` .

```go
for i := 0; i < 10; i++ {
	fmt.Println(i)
}

```

Если обе части `InitSimpleStatement` и `PostSimpleStatement` отсутствуют (просто рассматривайте их как пустые операторы), две точки с запятой рядом с ними можно опустить. Форма называется `for` формой цикла только с условием. Это то же самое, что `while` цикл в других языках.

```go
var i = 0
for ; i < 10; {
	fmt.Println(i)
	i++
}
for i < 20 {
	fmt.Println(i)
	i++
}

```

Если `Condition` часть отсутствует, компиляторы будут рассматривать ее как `true` .

```go
for i := 0; ; i++ { // <=> for i := 0; true; i++ {
	if i >= 10 {
		// "break" statement will be explained below.
		break
	}
	fmt.Println(i)
}

// The following 4 endless loops are
// equivalent to each other.
for ; true; {
}
for true {
}
for ; ; {
}
for {
}

```

Если `InitSimpleStatement` in a `for` block является коротким оператором объявления переменных, то объявленные переменные будут рассматриваться как объявленные в верхнем вложенном блоке неявного кода `for` блока. Например, следующий фрагмент кода печатается `012` вместо `0` .

```go
for i := 0; i < 3; i++ {
	fmt.Print(i)
	// The left i is a new declared variable,
	// and the right i is the loop variable.
	i := i
	// The new declared variable is modified, but
	// the old one (the loop variable) is not yet.
	i = 10
	_ = i
}

```

Оператор `break` может использоваться для предварительного перехода выполнения из `for` блока потока управления циклом, если `for` блок потока управления циклом является самым внутренним разрушаемым блоком потока управления, содержащим `break` оператор.

```go
i := 0
for {
	if i >= 10 {
		break
	}
	fmt.Println(i)
	i++
}

```

Оператор `continue` может быть использован для завершения текущего шага цикла заранее ( `PostSimpleStatement` все равно будет выполнен), если `for` блок потока управления циклом является самым внутренним блоком потока управления циклом, содержащим `continue` оператор. Например, следующий фрагмент кода напечатает `13579` .

```go
for i := 0; i < 10; i++ {
	if i % 2 == 0 {
		continue
	}
	fmt.Print(i)
}

```

### `switch-case` Блоки потока управления

`switch-case` блок потока управления является одним из видов блоков потока управления условным выполнением.

Полная форма `switch-case` блока

```go
switch InitSimpleStatement; CompareOperand0 {
case CompareOperandList1:
	// do something
case CompareOperandList2:
	// do something
...
case CompareOperandListN:
	// do something
default:
	// do something
}

```

В полной форме,

*   `switch` , `case` и `default` три ключевых слова.
*   Часть `InitSimpleStatement` должна быть простым заявлением. Часть `CompareOperand0` представляет собой выражение, которое рассматривается как типизированное значение (если это нетипизированное значение, то оно рассматривается как значение типа своего типа по умолчанию), поэтому оно не может быть нетипизированным `nil` . `CompareOperand0` называется выражением переключения в спецификации Go.
*   Каждая из частей `CompareOperandListX` ( `X` может представлять from `1` to `N` ) должна быть списком выражений, разделенных запятыми. Каждое из этих выражений должно быть сравнимо с `CompareOperand0` . Каждое из этих выражений называется выражением case в спецификации Go. Если выражение case является нетипизированным значением, оно должно быть неявно преобразовано в тип выражения switch в том же `switch-case` потоке управления. Если преобразование невозможно, компиляция завершается ошибкой.

Каждый `case CompareOperandListX:` или `default:` открывает (и за ним следует) неявный блок кода. Неявный блок кода и тот `case CompareOperandListX:` или `default:` образуют ветвь. Каждая такая ветвь не является обязательной. Позже мы назовем блок неявного кода в такой ветви блоком кода ветви.

В блоке потока управления может быть не более одной `default` ветви . `switch-case`

Помимо блоков кода ответвления, каждый `switch-case` поток управления формирует два блока кода, один неявный и один явный. Явный вложен в неявный. Все блоки кода ветвления вложены в явный (и косвенно вложены в неявный).

`switch-case` Блоки потока управления являются разрушаемыми, поэтому `break` операторы также могут использоваться в любом блоке кода ответвления в блоке потока управления, чтобы заранее `switch-case` выполнить переход из блока потока управления. `switch-case`

Блок `InitSimpleStatement` in a `for` loop будет выполнен (только один раз) перед выполнением других операторов в `for` блоке цикла.

Сначала `InitSimpleStatement` будет выполнено, когда будет выполнен `switch-case` поток управления, затем выражение переключения `CompareOperand0` будет оцениваться и оцениваться только один раз. Результат оценки всегда является типизированным значением. Результат оценки будет сравниваться (с помощью `==` оператора) с результатом оценки каждого выражения case в `CompareOperandListX` списках выражений сверху вниз и слева направо. Если обнаруживается, что выражение case равно `CompareOperand0` , процесс сравнения останавливается и выполняется соответствующий блок кода ответвления выражения. Если ни одно из выражений case не окажется равным `CompareOperand0` , будет выполнен блок кода ответвления по умолчанию (если он присутствует).

Пример `switch-case` потока управления:

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	switch n := rand.Intn(100); n%9 {
	case 0:
		fmt.Println(n, "is a multiple of 9.")

		// Different from many other languages,
		// in Go, the execution will automatically
		// jumps out of the switch-case block at
		// the end of each branch block.
		// No "break" statement is needed here.
	case 1, 2, 3:
		fmt.Println(n, "mod 9 is 1, 2 or 3.")
		// Here, this "break" statement is nonsense.
		break
	case 4, 5, 6:
		fmt.Println(n, "mod 9 is 4, 5 or 6.")
	// case 6, 7, 8:
		// The above case line might fail to compile,
		// for 6 is duplicate with the 6 in the last
		// case. The behavior is compiler dependent.
	default:
		fmt.Println(n, "mod 9 is 7 or 8.")
	}
}

```

Функция `rand.Intn` возвращает неотрицательное `int` случайное значение, которое меньше указанного аргумента.

Обратите внимание: если любые два выражения case в `switch-case` потоке управления могут быть обнаружены как равные во время компиляции, то компилятор может отклонить последнее. Например, стандартный компилятор Go считает `case 6, 7, 8` строку в приведенном выше примере недопустимой, если она не закомментирована. Но другие компиляторы могут подумать, что эта строка подходит. На самом деле, текущий стандартный компилятор Go (версия 1.19) [допускает дублирование выражений регистра логических](https://github.com/golang/go/issues/28357) значений, а gccgo (v8.2) допускает дублирование выражений регистра логических и строковых значений.

Как видно из комментариев в приведенном выше примере, в отличие от многих других языков, в Go в конце каждого блока кода ветвления выполнение автоматически выходит за пределы соответствующего `switch-case` блока управления. Тогда как позволить выполнению перейти к следующему блоку кода ветвления? Go предоставляет `fallthrough` ключевое слово для выполнения этой задачи. Например, в следующем примере каждый блок кода ответвления будет выполняться в соответствии с их порядком сверху вниз.

```go
rand.Seed(time.Now().UnixNano())
switch n := rand.Intn(100) % 5; n {
case 0, 1, 2, 3, 4:
	fmt.Println("n =", n)
	// The "fallthrough" statement makes the
	// execution slip into the next branch.
	fallthrough
case 5, 6, 7, 8:
	// A new declared variable also called "n",
	// it is only visible in the current
	// branch code block.
	n := 99
	fmt.Println("n =", n) // 99
	fallthrough
default:
	// This "n" is the switch expression "n".
	fmt.Println("n =", n)
}

```

Пожалуйста, обрати внимание,

*   оператор `fallthrough` должен быть последним оператором в ветке.
*   оператор `fallthrough` не может отображаться в последней ветви в `switch-case` блоке потока управления.

Например, все следующие `fallthrough` виды использования являются незаконными.

```go
switch n := rand.Intn(100) % 5; n {
case 0, 1, 2, 3, 4:
	fmt.Println("n =", n)
	// The if-block, not the fallthrough statement,
	// is the final statement in this branch.
	if true {
		fallthrough // error: not the final statement
	}
case 5, 6, 7, 8:
	n := 99
	fallthrough // error: not the final statement
	_ = n
default:
	fmt.Println(n)
	fallthrough // error: show up in the final branch
}

```

Части `InitSimpleStatement` и `CompareOperand0` в `switch-case` потоке управления являются необязательными. Если `CompareOperand0` часть отсутствует, она будет рассматриваться как `true` типизированное значение встроенного типа `bool` . Если `InitSimpleStatement` часть отсутствует, точка с запятой после нее может быть опущена.

И, как упоминалось выше, все ветки необязательны. Таким образом, все следующие блоки кода являются допустимыми, все они могут рассматриваться как недействующие.

```go
switch n := 5; n {
}

switch 5 {
}

switch _ = 5; {
}

switch {
}

```

Для последних двух `switch-case` блоков потока управления в последнем примере, как упоминалось выше, каждая из отсутствующих `CompareOperand0` частей рассматривается как типизированное значение `true` встроенного типа `bool` . Таким образом, следующий фрагмент кода напечатает `hello` .

```go
switch { // <=> switch true {
case true: fmt.Println("hello")
default: fmt.Println("bye")
}

```

Другим очевидным отличием от многих других языков является то, что порядок `default` ветвления в `switch-case` блоке потока управления может быть произвольным. Например, следующие три `switch-case` блока потока управления эквивалентны друг другу.

```go
switch n := rand.Intn(3); n {
case 0: fmt.Println("n == 0")
case 1: fmt.Println("n == 1")
default: fmt.Println("n == 2")
}

switch n := rand.Intn(3); n {
default: fmt.Println("n == 2")
case 0: fmt.Println("n == 0")
case 1: fmt.Println("n == 1")
}

switch n := rand.Intn(3); n {
case 0: fmt.Println("n == 0")
default: fmt.Println("n == 2")
case 1: fmt.Println("n == 1")
}

```

### `goto` Заявление и декларация этикетки

Как и многие другие языки, Go также поддерживает `goto` операторы. За `goto` ключевым словом должна следовать метка, чтобы сформировать утверждение. Метка объявляется с формой `LabelName:` , где `LabelName` должен быть идентификатор. Метка, имя которой не является пустым идентификатором, должна использоваться хотя бы один раз.

Оператор `goto` заставит выполнение перейти к следующему оператору, следующему за объявлением метки, используемой в `goto` операторе. Таким образом, за объявлением метки должен следовать один оператор.

Метка должна быть объявлена ​​в теле функции. Использование метки может появиться до или после объявления метки. Но метка не видна (и не может появиться) за пределами самого внутреннего блока кода, в котором она объявлена.

В следующем примере `goto` оператор и метка используются для реализации потока управления циклом.

```go
package main

import "fmt"

func main() {
	i := 0

Next: // here, a label is declared.
	fmt.Println(i)
	i++
	if i < 5 {
		goto Next // execution jumps
	}
}

```

Как упоминалось выше, метка не видна (и не может появиться) за пределами самого внутреннего блока кода, в котором она объявлена. Поэтому следующий пример не компилируется.

```go
package main

func main() {
goto Label1 // error
	{
		Label1:
		goto Label2 // error
	}
	{
		Label2:
	}
}

```

Обратите внимание, что если метка объявлена ​​в области действия переменной, то использование метки не может появиться перед объявлением переменной. Области действия идентификатора будут объяснены в статьях о [блоках и областях действия в Go](https://go101.org/article/blocks-and-scopes.html) позже.

Следующий пример также не компилируется.

```go
package main

import "fmt"

func main() {
	i := 0
Next:
	if i >= 5 {
		// error: jumps over declaration of k
		goto Exit
	}

	k := i + i
	fmt.Println(k)
	i++
	goto Next

// This label is declared in the scope of k,
// but its use is outside of the scope of k.
Exit:
}

```

Только что упомянутое правило [может измениться позже](https://github.com/golang/go/issues/26058) . В настоящее время, чтобы приведенный выше код нормально компилировался, мы должны настроить область действия переменной `k` . Есть два способа решить проблему в последнем примере.

Один из способов — уменьшить область действия переменной `k` .

```go
func main() {
	i := 0
Next:
	if i >= 5 {
		goto Exit
	}
	// Create an explicit code block to
	// shrink the scope of k.
	{
		k := i + i
		fmt.Println(k)
	}
	i++
	goto Next
Exit:
}

```

Другой способ — расширить область видимости переменной `k` .

```go
func main() {
	var k int // move the declaration of k here.
	i := 0
Next:
	if i >= 5 {
		goto Exit
	}

	k = i + i
	fmt.Println(k)
	i++
	goto Next
Exit:
}

```

### `break` и `continue` заявления с метками

Оператор `goto` должен содержать метку. Оператор `break` or `continue` также может содержать метку, но метка необязательна. Как правило, `break` содержащие метки используются во вложенных прерываемых блоках потока управления, а `continue` операторы, содержащие метки, используются во вложенных блоках потока управления циклом.

If a `break` statement contains a label, the label must be declared just before a breakable control flow block which contains the `break` statement. We can view the label name as the name of the breakable control flow block. The `break` statement will make execution jump out of the breakable control flow block, even if the breakable control flow block is not the innermost breakable control flow block containing `break` statement.

Если `continue` оператор содержит метку, метка должна быть объявлена ​​непосредственно перед блоком потока управления циклом, который содержит `continue` оператор. Мы можем рассматривать имя метки как имя блока потока управления циклом. Оператор `continue` завершит текущий шаг блока потока управления циклом заранее, даже если блок потока управления циклом не является самым внутренним блоком потока управления циклом, содержащим `continue` оператор.

Ниже приведен пример использования операторов `break` and `continue` с метками.

```go
package main

import "fmt"

func FindSmallestPrimeLargerThan(n int) int {
Outer:
	for n++; ; n++{
		for i := 2; ; i++ {
			switch {
			case i * i > n:
				break Outer
			case n % i == 0:
				continue Outer
			}
		}
	}
	return n
}

func main() {
	for i := 90; i < 100; i++ {
		n := FindSmallestPrimeLargerThan(i)
		fmt.Print("The smallest prime number larger than ")
		fmt.Println(i, "is", n)
	}
}
```
