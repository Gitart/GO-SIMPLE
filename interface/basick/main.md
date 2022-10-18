# Интерфейсы в Go

Типы интерфейсов — это особый вид типов в Go. Интерфейсы играют несколько важных ролей в Go. По сути, типы интерфейсов заставляют Go поддерживать упаковку значений. Следовательно, благодаря упаковке значений поддерживаются отражение и полиморфизм.

Начиная с версии 1.18 Go уже поддерживает пользовательские дженерики. В пользовательских дженериках тип интерфейса может (всегда) также использоваться в качестве ограничений типа. На самом деле все ограничения типов на самом деле являются интерфейсными типами. До версии Go 1.18 в качестве типов значений можно было использовать все типы интерфейсов. Но начиная с Go 1.18 некоторые типы интерфейсов можно использовать только как ограничения типов. Типы интерфейсов, которые могут использоваться в качестве типов значений, называются базовыми типами интерфейсов.

Эта статья в основном была написана до того, как Go стал поддерживать пользовательские дженерики, поэтому в ней в основном рассказывается об основных интерфейсах. Подробнее о типах интерфейсов только с ограничениями читайте в книге [Go generics 101 .](https://go101.org/generics/101.html)

### Типы интерфейсов и наборы типов

Тип интерфейса определяет некоторые (тип) требования. Все неинтерфейсные типы, удовлетворяющие этим требованиям, образуют набор типов, который называется набором типов интерфейсного типа.

Требования, определенные для типа интерфейса, выражаются путем встраивания некоторых элементов интерфейса в тип интерфейса. В настоящее время (Go 1.19) существует два типа элементов интерфейса: элементы метода и элементы типа.

*   Элемент метода представляет собой [спецификацию метода](https://go101.org/article/method.html#method-set) . Спецификация метода, встроенная в тип интерфейса, не может использовать пустой идентификатор `_` в качестве своего имени.
*   Элемент типа может быть именем типа, литералом типа, приближенным типом или объединением типов. В текущей статье мало говорится о последних двух, а речь идет только об именах типов и литералах, которые обозначают типы интерфейсов.

Например, предварительно объявленный [`error` тип интерфейса](https://golang.org/pkg/builtin/#error) , определение которого показано ниже, включает в себя спецификацию метода `Error() string` . В определении `interface{...}` называется литералом типа интерфейса, а слово `interface` является ключевым словом в Go.

```go
type error interface {
        Error() string
}

```

Мы также можем сказать, что `error` тип интерфейса (непосредственно) определяет метод `Error() string` . Его набор типов состоит из всех неинтерфейсных типов, у которых есть [метод](https://go101.org/article/method.html) со спецификацией `Error() string` . Теоретически набор типов бесконечен. Конечно, для указанного проекта Go он конечен.

Ниже приведены некоторые другие определения типов интерфейсов и объявления псевдонимов.

```go
// This interface directly specifies two methods and
// embeds two other interface types, one of which
// is a type name and the other is a type literal.
type ReadWriteCloser = interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	error                      // a type name
	interface{ Close() error } // a type literal
}

// This interface embeds an approximation type. Its type
// set inlcudes all types whose underlying type is []byte.
type AnyByteSlice = interface {
	~[]byte
}

// This interface embeds a type union. Its type set inlcudes
// 6 types: uint, uint8, uint16, uint32, uint64 and uintptr.
type Unsigned interface {
	uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

```

Встраивание типа интерфейса (обозначаемого либо именем типа, либо литералом типа) в другой эквивалентно (рекурсивному) расширению элементов первого во второй. Например, тип интерфейса, обозначенный псевдонимом типа `ReadWriteCloser` , эквивалентен типу интерфейса, обозначенному следующим литералом, который непосредственно определяет четыре метода.

```go
interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Error() string
	Close() error
}

```

Набор типов вышеуказанного типа интерфейса состоит из всех неинтерфейсных типов, которые имеют как минимум четыре метода, указанные типом интерфейса. Набор типов также бесконечен. Это определенно подмножество набора типов `error` интерфейсного типа.

Обратите внимание, что до версии Go 1.18 в типы интерфейсов можно встраивать только имена типов интерфейсов.

Все типы интерфейсов, показанные в следующем коде, называются пустыми типами интерфейсов, которые ничего не встраивают.

```go
// The unnamed blank interface type.
interface{}

// Nothing is a defined blank interface type.
type Nothing interface{}

```

На самом деле, в Go 1.18 появился предварительно объявленный псевдоним, `any` обозначающий пустой тип интерфейса `interface{}` .

Набор типов пустого типа интерфейса состоит из всех неинтерфейсных типов.

### Наборы методов типов

С каждым типом связан [набор методов .](https://go101.org/article/method.html#method-set)

*   Для неинтерфейсного типа его набор методов состоит из спецификаций всех [методов (явных или неявных), объявленных](https://go101.org/article/method.html) для него.
*   Для типа интерфейса его набор методов состоит из всех спецификаций методов, которые он указывает, прямо или косвенно, посредством встраивания других типов.

В примерах, показанных в последнем разделе,

*   набор методов типа интерфейса, обозначенного значком, `ReadWriteCloser` содержит четыре метода.
*   набор методов предварительно объявленного типа интерфейса `error` содержит только один метод.
*   набор методов пустого типа интерфейса пуст.

Для удобства набор методов типа часто также называют набором методов любого значения типа.

### Основные типы интерфейсов

Базовые типы интерфейсов — это типы интерфейсов, которые можно использовать в качестве типов значений. Небазовый тип интерфейса также называется типом интерфейса только с ограничениями.

В настоящее время (Go 1.19) каждый базовый тип интерфейса может быть полностью определен набором методов (может быть пустым). Другими словами, базовый тип интерфейса не требует определения элементов типа.

В примерах, показанных в предыдущем разделе, тип интерфейса, обозначенный псевдонимом `ReadWriteCloser` , является базовым типом, а `Unsigned` тип интерфейса и тип, обозначенный псевдонимом `AnyByteSlice` , — нет. Последние два являются типами интерфейса только с ограничениями.

Пустые типы интерфейса и предварительно объявленный `error` тип интерфейса также являются базовыми типами интерфейса.

Два безымянных базовых типа интерфейса идентичны, если их наборы методов идентичны. Обратите внимание, что неэкспортированные имена методов (начинающиеся со строчных букв) из разных пакетов всегда будут рассматриваться как два разных имени метода, даже если сами имена двух методов совпадают.

### Реализации

Если неинтерфейсный тип содержится в наборе типов интерфейсного типа, то мы говорим, что неинтерфейсный тип реализует интерфейсный тип. Если набор типов интерфейса является подмножеством другого типа интерфейса, то говорят, что первый реализует второй.

Тип интерфейса всегда реализует сам себя, поскольку набор типов всегда является подмножеством (или надмножеством) самого себя. Точно так же два типа интерфейса с одним и тем же набором методов реализуют друг друга. На самом деле два безымянных типа интерфейса идентичны, если их наборы типов идентичны.

Если тип `T` реализует тип интерфейса `X` , то набор методов `T` должен быть надмножеством `X` , независимо от того, `T` является ли он типом интерфейса или типом, не являющимся интерфейсом. Как правило, не наоборот. А если `X` это базовый интерфейс, то наоборот. Например, в примерах, представленных в предыдущем разделе, тип интерфейса, обозначенный как `ReadWriteCloser` реализует `error` тип интерфейса.

Все реализации неявны в Go. Компилятор не требует явного указания отношений реализации в коде. `implements` В Go нет ключевого слова. Компиляторы Go будут автоматически проверять отношения реализации по мере необходимости.

Например, в следующем примере наборы методов типа указателя структуры `*Book` , целочисленного типа `MyInt` и типа указателя `*MyInt` содержат спецификацию метода `About() string` , поэтому все они реализуют вышеупомянутый тип интерфейса `Aboutable` .

```go
type Aboutable interface {
	About() string
}

type Book struct {
	name string
	// more other fields ...
}

func (book *Book) About() string {
	return "Book: " + book.name
}

type MyInt int

func (MyInt) About() string {
	return "I'm a custom integer value"
}

```

Дизайн неявной реализации позволяет типам, определенным в других пакетах библиотек, таких как стандартные пакеты, пассивно реализовывать некоторые типы интерфейсов, объявленные в пользовательских пакетах. Например, если мы объявим тип интерфейса следующим образом, то тип `DB` и тип, `Tx` объявленные в [стандартном `database/sql` пакете](https://golang.org/pkg/database/sql/) , будут автоматически реализовывать тип интерфейса, поскольку они оба имеют три соответствующих метода, указанных в интерфейсе.

```go
import "database/sql"

...

type DatabaseStorer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

```

Обратите внимание, что набор типов пустого типа интерфейса состоит из всех неинтерфейсных типов, поэтому все типы реализуют любой пустой тип интерфейса. Это важный факт в Go.

### Ценностный бокс

Опять же, в настоящее время (Go 1.19) типы значений интерфейса должны быть базовыми типами интерфейса. В оставшемся содержании текущей статьи, когда упоминается тип значения, тип значения может быть не интерфейсным типом или базовым интерфейсным типом. Это никогда не тип интерфейса только с ограничениями.

Мы можем рассматривать каждое значение интерфейса как поле для инкапсуляции значения, не связанного с интерфейсом. Чтобы упаковать/инкапсулировать неинтерфейсное значение в интерфейсное значение, тип неинтерфейсного значения должен реализовывать тип интерфейсного значения.

Если тип `T` реализует (базовый) тип интерфейса `I` , то любое значение типа `T` может быть неявно преобразовано в тип `I` . Другими словами, любое значение типа `T` может быть [присвоено](https://go101.org/article/constants-and-variables.html#assignment) (изменяемым) значениям типа `I` . Когда `T` значение преобразуется (присваивается) в `I` значение,

*   если type `T` не является интерфейсным типом, то копия `T` значения упаковывается (или инкапсулируется) в результирующее (или целевое) `I` значение. Временная сложность копии `*O*(n)` , где `n` \- размер копируемого `T` значения.
*   если тип `T` также является интерфейсным типом, то копия значения, заключенного в `T` значение, помещается (или инкапсулируется) в значение результата (или назначения) `I` . Стандартный компилятор Go делает здесь оптимизацию, поэтому временная сложность копии составляет `*O*(1)` , а не `*O*(n)` .

Информация о типе упакованного значения также хранится в значении интерфейса результата (или назначения). (Это будет объяснено ниже.)

Когда значение заключено в значение интерфейса, это значение называется ***динамическим значением значения*** интерфейса. Тип динамического значения называется ***динамическим типом*** значения интерфейса.

Прямая часть динамического значения интерфейсного значения является неизменной, хотя мы можем заменить динамическое значение интерфейсного значения другим динамическим значением.

В Go нулевые значения любого типа интерфейса представлены предварительно объявленным `nil` идентификатором. Ничто не заключено в нулевое значение интерфейса. Присвоение нетипизированного `nil` значения интерфейса очистит динамическое значение, заключенное в поле значения интерфейса.

*(Обратите внимание, что нулевые значения многих неинтерфейсных типов в Go также представлены `nil` в Go. Неинтерфейсные нулевые значения также могут быть заключены в интерфейсные значения. не является нулевым значением интерфейса.)*

Поскольку любой тип реализует все типы пустых интерфейсов, любое значение, не относящееся к интерфейсу, может быть заключено (или назначено) в пустое значение интерфейса. По этой причине пустые типы интерфейсов можно рассматривать как `any` типы во многих других языках.

Когда нетипизированное значение (кроме нетипизированных `nil` значений) присваивается пустому значению интерфейса, нетипизированное значение будет сначала преобразовано в его тип по умолчанию. (Другими словами, мы можем думать, что нетипизированное значение выводится как значение его типа по умолчанию).

Давайте рассмотрим пример, который демонстрирует некоторые назначения со значениями интерфейса в качестве мест назначения.

```go
package main

import "fmt"

type Aboutable interface {
	About() string
}

// Type *Book implements Aboutable.
type Book struct {
	name string
}
func (book *Book) About() string {
	return "Book: " + book.name
}

func main() {
	// A *Book value is boxed into an
	// interface value of type Aboutable.
	var a Aboutable = &Book{"Go 101"}
	fmt.Println(a) // &{Go 101}

	// i is a blank interface value.
	var i interface{} = &Book{"Rust 101"}
	fmt.Println(i) // &{Rust 101}

	// Aboutable implements interface{}.
	i = a
	fmt.Println(i) // &{Go 101}
}

```

Обратите внимание, что прототип `fmt.Println` функции, много раз использованный в предыдущих статьях,

```go
func Println(a ...interface{}) (n int, err error)

```

Вот почему `fmt.Println` вызовы функций могут принимать аргументы любых типов.

Ниже приведен еще один пример, показывающий, как пустое значение интерфейса используется для упаковывания значений любого неинтерфейсного типа.

```go
package main

import "fmt"

func main() {
	var i interface{}
	i = []int{1, 2, 3}
	fmt.Println(i) // [1 2 3]
	i = map[string]int{"Go": 2012}
	fmt.Println(i) // map[Go:2012]
	i = true
	fmt.Println(i) // true
	i = 1
	fmt.Println(i) // 1
	i = "abc"
	fmt.Println(i) // abc

	// Clear the boxed value in interface value i.
	i = nil
	fmt.Println(i) // <nil>
}

```

Компиляторы Go создадут глобальную таблицу, содержащую информацию о каждом типе во время компиляции. Информация включает в себя [тип](https://go101.org/article/type-system-overview.html#type-kinds) типа, какими методами и полями владеет тип, тип элемента типа контейнера, размеры типа и т. д. Глобальная таблица будет загружена в память при запуске программы.

Во время выполнения, когда значение, не являющееся интерфейсом, заключено в значение интерфейса, среда выполнения Go (по крайней мере, для стандартной среды выполнения Go) будет анализировать и создавать информацию о реализации для пары типов двух значений и сохранять информацию о реализации. в значении интерфейса. Информация о реализации для каждой пары неинтерфейсного типа и типа интерфейса будет построена только один раз и кэширована в глобальной карте для рассмотрения эффективности выполнения. Количество записей глобальной карты никогда не уменьшается. На самом деле ненулевое значение интерфейса просто использует [поле внутреннего указателя, которое ссылается на кэшированную запись информации о реализации](https://go101.org/article/value-part.html#interface-structure) .

Информация о реализации для каждой пары (тип интерфейса, динамический тип) включает две части информации:

1.  информация динамического типа (не интерфейсного типа)
2.  и таблицу методов (срез), в которой хранятся все соответствующие методы, указанные типом интерфейса и объявленные для неинтерфейсного типа (динамический тип).

Эти две части информации необходимы для реализации двух важных функций в Go:

1.  Информация о динамическом типе является ключом к реализации [отражения](https://go101.org/article/interface.html#reflection) в Go.
2.  Информация таблицы методов является ключом к реализации полиморфизма (полиморфизм будет объяснен в следующем разделе).

### Полиморфизм

Полиморфизм — это одна из ключевых функций, предоставляемых интерфейсами, и это важная особенность Go.

Когда неинтерфейсное значение `t` типа `T` заключено в интерфейсное значение `i` type `I` , вызов метода, указанного интерфейсным типом `I` для интерфейсного значения , фактически `i` вызовет соответствующий метод, объявленный для неинтерфейсного типа , для неинтерфейсного значения. Другими словами, **вызов метода значения интерфейса фактически вызовет соответствующий метод динамического значения значения интерфейса** . Например, вызов метода вызовет метод на самом деле. Когда в значение интерфейса заключены разные динамические значения разных динамических типов, значение интерфейса ведет себя по-разному. Это называется полиморфизмом. `T` `t` `i.m` `t.m`

При `i.m` вызове метода таблица методов в информации о реализации, хранящейся в `i` , будет просматриваться, чтобы найти и вызвать соответствующий метод `t.m` . Таблица методов представляет собой срез, а поиск — это просто индексация элементов среза, так что это быстро.

*(Обратите внимание, что вызов методов для нулевого значения интерфейса вызовет панику во время выполнения, поскольку нет доступных объявленных методов для вызова.)*

Пример:

```go
package main

import "fmt"

type Filter interface {
	About() string
	Process([]int) []int
}

// UniqueFilter is used to remove duplicate numbers.
type UniqueFilter struct{}
func (UniqueFilter) About() string {
	return "remove duplicate numbers"
}
func (UniqueFilter) Process(inputs []int) []int {
	outs := make([]int, 0, len(inputs))
	pusheds := make(map[int]bool)
	for _, n := range inputs {
		if !pusheds[n] {
			pusheds[n] = true
			outs = append(outs, n)
		}
	}
	return outs
}

// MultipleFilter is used to keep only
// the numbers which are multiples of
// the MultipleFilter as an int value.
type MultipleFilter int
func (mf MultipleFilter) About() string {
	return fmt.Sprintf("keep multiples of %v", mf)
}
func (mf MultipleFilter) Process(inputs []int) []int {
	var outs = make([]int, 0, len(inputs))
	for _, n := range inputs {
		if n % int(mf) == 0 {
			outs = append(outs, n)
		}
	}
	return outs
}

// With the help of polymorphism, only one
// "filterAndPrint" function is needed.
func filterAndPrint(fltr Filter, unfiltered []int) []int {
	// Calling the methods of "fltr" will call the
	// methods of the value boxed in "fltr" actually.
	filtered := fltr.Process(unfiltered)
	fmt.Println(fltr.About() + ":\n\t", filtered)
	return filtered
}

func main() {
	numbers := []int{12, 7, 21, 12, 12, 26, 25, 21, 30}
	fmt.Println("before filtering:\n\t", numbers)

	// Three non-interface values are boxed into
	// three Filter interface slice element values.
	filters := []Filter{
		UniqueFilter{},
		MultipleFilter(2),
		MultipleFilter(3),
	}

	// Each slice element will be assigned to the
	// local variable "fltr" (of interface type
	// Filter) one by one. The value boxed in each
	// element will also be copied into "fltr".
	for _, fltr := range filters {
		numbers = filterAndPrint(fltr, numbers)
	}
}

```

Выход:

```
before filtering:
	 [12 7 21 12 12 26 25 21 30]
remove duplicate numbers:
	 [12 7 21 26 25 30]
keep multiples of 2:
	 [12 26 30]
keep multiples of 3:
	 [12 30]

```

В приведенном выше примере полиморфизм делает ненужным написание одной `filterAndPrint` функции для каждого типа фильтра.

Помимо вышеупомянутого преимущества, полиморфизм также позволяет разработчикам пакета библиотечного кода объявить экспортируемый тип интерфейса и объявить функцию (или метод), которая имеет параметр типа интерфейса, так что пользователь пакета может объявить тип, который реализует тип интерфейса в пользовательском коде и передает аргументы пользовательского типа в вызовы функции (или метода). Разработчикам пакета кода не нужно заботиться о том, как объявлен пользовательский тип, если пользовательский тип удовлетворяет поведению, заданному типом интерфейса, объявленным в пакете кода библиотеки.

На самом деле полиморфизм не является существенным свойством языка. Существуют альтернативные способы выполнения той же работы, такие как функции обратного вызова. Но способ полиморфизма чище и элегантнее.

### Отражение

Информация о динамическом типе, хранящаяся в значении интерфейса, может использоваться для проверки динамического значения значения интерфейса и управления значениями, на которые ссылается динамическое значение. В программировании это называется отражением.

В этой статье не будут объясняться функции, предоставляемые [стандартным `reflect` пакетом](https://golang.org/pkg/reflect/) . Пожалуйста, прочитайте [размышления в Go](https://go101.org/article/reflection.html) , чтобы узнать, как использовать этот пакет. Ниже будут представлены только встроенные функции отражения в Go. В Go встроенные отражения достигаются с помощью утверждений типов и `type-switch` блоков кода потока управления.

#### Утверждение типа

В Go есть четыре типа случаев преобразования значений, связанных с интерфейсом:

1.  преобразовать неинтерфейсное значение в интерфейсное значение, где тип неинтерфейсного значения должен реализовывать тип интерфейсного значения.
2.  преобразовать значение интерфейса в значение интерфейса, где тип исходного значения интерфейса должен реализовывать тип значения целевого интерфейса.
3.  преобразовать значение интерфейса в значение, не являющееся интерфейсом, где тип значения, не являющегося интерфейсом, должен реализовывать тип значения интерфейса.
4.  преобразовать значение интерфейса в значение интерфейса, где тип значения исходного интерфейса не реализует тип интерфейса назначения, но динамический тип значения исходного интерфейса может реализовать тип интерфейса назначения.

Мы уже объяснили первые два вида случаев. Оба они требуют, чтобы тип исходного значения реализовывал тип интерфейса назначения. Конвертируемость первых двух проверяется во время компиляции.

Здесь будут объяснены последние два вида случаев. Конвертируемость для последних двух проверяется во время выполнения с помощью синтаксиса, называемого ***утверждением типа*** . На самом деле, синтаксис также применим ко второму типу преобразования в нашем списке выше.

Форма выражения утверждения типа: `i.(T)` , где `i` — значение интерфейса, а `T` — имя типа или литерал типа. Тип `T` должен быть

*   либо произвольный неинтерфейсный тип,
*   или произвольный тип интерфейса.

В утверждении типа `i.(T)` называется `i` утвержденным значением и `T` называется утвержденным типом. Утверждение типа может быть успешным или неудачным.

*   В случае `T` неинтерфейсного типа, если динамический тип `i` существует и идентичен `T` , то утверждение будет выполнено успешно, в противном случае утверждение завершится ошибкой. Когда утверждение выполняется успешно, результатом оценки утверждения является копия динамического значения `i` . Мы можем рассматривать утверждения такого рода как попытки распаковки значения.
*   В случае типа `T` интерфейса, если динамический тип `i` существует и реализует `T` , то утверждение будет выполнено успешно, в противном случае утверждение завершится ошибкой. Когда утверждение выполняется успешно, копия динамического значения `i` будет заключена в `T` значение, и это `T` значение будет использоваться в качестве результата оценки утверждения.

Когда утверждение типа не выполняется, результатом его оценки является нулевое значение утвержденного типа.

По правилам, описанным выше, если утвержденное значение в утверждении типа является нулевым значением интерфейса, то утверждение всегда будет давать сбой.

В большинстве сценариев утверждение типа используется как выражение с одним значением. Однако, когда утверждение типа используется в качестве единственного выражения исходного значения в присваивании, оно может привести ко второму необязательному нетипизированному логическому значению и рассматриваться как выражение с несколькими значениями. Второе необязательное нетипизированное логическое значение указывает, успешно ли выполняется утверждение типа.

Обратите внимание: если утверждение типа не выполняется и утверждение типа используется как выражение с одним значением (второй необязательный логический результат отсутствует), то произойдет паника.

Пример, который показывает, как использовать утверждения типа (утверждаемые типы не являются интерфейсными типами):

```go
package main

import "fmt"

func main() {
	// Compiler will deduce the type of 123 as int.
	var x interface{} = 123

	// Case 1:
	n, ok := x.(int)
	fmt.Println(n, ok) // 123 true
	n = x.(int)
	fmt.Println(n) // 123

	// Case 2:
	a, ok := x.(float64)
	fmt.Println(a, ok) // 0 false

	// Case 3:
	a = x.(float64) // will panic
}

```

Другой пример, показывающий, как использовать утверждения типа (утверждаемые типы являются типами интерфейса):

```go
package main

import "fmt"

type Writer interface {
	Write(buf []byte) (int, error)
}

type DummyWriter struct{}
func (DummyWriter) Write(buf []byte) (int, error) {
	return len(buf), nil
}

func main() {
	var x interface{} = DummyWriter{}
	var y interface{} = "abc"
	// Now the dynamic type of y is "string".
	var w Writer
	var ok bool

	// Type DummyWriter implements both
	// Writer and interface{}.
	w, ok = x.(Writer)
	fmt.Println(w, ok) // {} true
	x, ok = w.(interface{})
	fmt.Println(x, ok) // {} true

	// The dynamic type of y is "string",
	// which doesn't implement Writer.
	w, ok = y.(Writer)
	fmt.Println(w, ok) //  false
	w = y.(Writer)     // will panic
}

```

Фактически для значения интерфейса `i` с динамическим типом `T` вызов метода `i.m(...)` эквивалентен вызову метода `i.(T).m(...)` .

#### `type-switch` блок управления потоком

Синтаксис `type-switch` блока кода может быть самым странным синтаксисом в Go. Его можно рассматривать как расширенную версию утверждения типа. Кодовый `type-switch` блок в чем-то похож на `switch-case` кодовый блок потока управления. Это выглядит как:

```go
switch aSimpleStatement; v := x.(type) {
case TypeA:
	...
case TypeB, TypeC:
	...
case nil:
	...
default:
	...
}

```

Эта `aSimpleStatement;` часть не является обязательной в `type-switch` кодовом блоке. `aSimpleStatement` должно быть [простое утверждение](https://go101.org/article/expressions-and-statements.html#simple-statements) . `x` должно быть значением интерфейса, и оно называется утвержденным значением. `v` называется результатом утверждения, он должен быть представлен в форме короткого объявления переменной.

За каждым `case` ключевым словом в `type-switch` блоке может следовать предварительно объявленный `nil` идентификатор или список, разделенный запятыми, состоящий как минимум из одного имени типа и литерала типа. Ни один из таких элементов ( `nil` , имена типов и литералы типов) не может дублироваться в одном и том же `type-switch` блоке кода.

Если тип, обозначенный именем типа или литералом типа, следующим за `case` ключевым словом в `type-switch` блоке кода, не является типом интерфейса, то он должен реализовывать тип интерфейса утвержденного значения.

Вот пример, в котором используется `type-switch` блок кода потока управления.

```go
package main

import "fmt"

func main() {
	values := []interface{}{
		456, "abc", true, 0.33, int32(789),
		[]int{1, 2, 3}, map[int]bool{}, nil,
	}
	for _, x := range values {
		// Here, v is declared once, but it denotes
		// different variables in different branches.
		switch v := x.(type) {
		case []int: // a type literal
			// The type of v is "[]int" in this branch.
			fmt.Println("int slice:", v)
		case string: // one type name
			// The type of v is "string" in this branch.
			fmt.Println("string:", v)
		case int, float64, int32: // multiple type names
			// The type of v is "interface{}",
			// the same as x in this branch.
			fmt.Println("number:", v)
		case nil:
			// The type of v is "interface{}",
			// the same as x in this branch.
			fmt.Println(v)
		default:
			// The type of v is "interface{}",
			// the same as x in this branch.
			fmt.Println("others:", v)
		}
		// Note, each variable denoted by v in the
		// last three branches is a copy of x.
	}
}

```

Выход:

```
number: 456
string: abc
others: true
number: 0.33
number: 789
int slice: [1 2 3]
others: map[]
<nil>

```

Приведенный выше пример эквивалентен следующему в логике:

```go
package main

import "fmt"

func main() {
	values := []interface{}{
		456, "abc", true, 0.33, int32(789),
		[]int{1, 2, 3}, map[int]bool{}, nil,
	}
	for _, x := range values {
		if v, ok := x.([]int); ok {
			fmt.Println("int slice:", v)
		} else if v, ok := x.(string); ok {
			fmt.Println("string:", v)
		} else if x == nil {
			v := x
			fmt.Println(v)
		} else {
			_, isInt := x.(int)
			_, isFloat64 := x.(float64)
			_, isInt32 := x.(int32)
			if isInt || isFloat64 || isInt32 {
				v := x
				fmt.Println("number:", v)
			} else {
				v := x
				fmt.Println("others:", v)
			}
		}
	}
}

```

`type-switch` кодовые блоки подобны `switch-case` кодовым блокам в некоторых аспектах.

*   Как и `switch-case` блоки, в `type-switch` блоке кода может быть не более одной `default` ветви.
*   Подобно `switch-case` блокам, в `type-switch` кодовом блоке, если присутствует `default` ветвь, это может быть последняя ветвь, первая ветвь или средняя ветвь.
*   Как и `switch-case` блоки, `type-switch` блок кода может не содержать никаких ветвей, он будет рассматриваться как незадействованный.

Но, в отличие от `switch-case` блоков кода, `fallthrough` операторы нельзя использовать в блоках ветвей блока `type-switch` кода.

### Подробнее об интерфейсах в Go

#### Сравнения с использованием значений интерфейса

Есть два случая сравнений с использованием интерфейсных значений:

1.  сравнения между значением, не являющимся интерфейсом, и значением интерфейса.
2.  сравнения между двумя значениями интерфейса.

В первом случае тип неинтерфейсного значения должен реализовывать тип (предположим, что это `I` ) интерфейсного значения, поэтому неинтерфейсное значение может быть преобразовано (упаковано) в интерфейсное значение `I` . Это означает, что сравнение между значением, не относящимся к интерфейсу, и значением интерфейса может быть преобразовано в сравнение между двумя значениями интерфейса. Поэтому ниже будут объяснены только сравнения между двумя значениями интерфейса.

Сравнение двух значений интерфейса фактически сравнивает их соответствующие динамические типы и динамические значения.

Этапы сравнения двух значений интерфейса (с `==` оператором):

1.  если одно из двух значений интерфейса является нулевым значением интерфейса, то результатом сравнения является то, является ли другое значение интерфейса также нулевым значением интерфейса.
2.  если динамические типы двух значений интерфейса являются двумя разными типами, то результатом сравнения будет `false` .
3.  в случае, когда динамические типы двух значений интерфейса имеют один и тот же тип,
    *   если тот же динамический тип является [несравнимым типом](https://go101.org/article/value-conversions-assignments-and-comparisons.html#comparison-rules) , произойдет паника.
    *   в противном случае результат сравнения является результатом сравнения динамических значений двух значений интерфейса.

Короче говоря, два значения интерфейса равны, только если выполняется одно из следующих условий.

1.  Они оба являются нулевыми значениями интерфейса.
2.  Их динамические типы идентичны и сопоставимы, а их динамические значения равны друг другу.

По правилам, два значения интерфейса, оба из которых являются динамическими значениями, `nil` могут не совпадать. Пример:

```go
package main

import "fmt"

func main() {
	var a, b, c interface{} = "abc", 123, "a"+"b"+"c"
	// A case of step 2.
	fmt.Println(a == b) // false
	// A case of step 3.
	fmt.Println(a == c) // true

	var x *int = nil
	var y *bool = nil
	var ix, iy interface{} = x, y
	var i interface{} = nil
	// A case of step 2.
	fmt.Println(ix == iy) // false
	// A case of step 1.
	fmt.Println(ix == i) // false
	// A case of step 1.
	fmt.Println(iy == i) // false

	// []int is an incomparable type
	var s []int = nil
	i = s
	// A case of step 1.
	fmt.Println(i == nil) // false
	// A case of step 3.
	fmt.Println(i == i) // will panic
}

```

#### Внутренняя структура интерфейсных значений

Для официального компилятора/среды выполнения Go пустые значения интерфейса и непустые значения интерфейса представлены двумя разными внутренними структурами. Пожалуйста, ознакомьтесь с [ценными частями](https://go101.org/article/value-part.html#interface-structure) для получения подробной информации.

#### Динамическое значение указателя и динамическое значение без указателя

Официальный компилятор/среда выполнения Go делает оптимизацию, которая делает упаковку значений указателя в значения интерфейса более эффективной, чем упаковка значений без указателя. Для [малых значений размера](https://go101.org/article/value-copy-cost.html) различия в эффективности невелики, но для больших значений размеров различия могут быть немалыми. Для той же оптимизации утверждения типа с типом указателя также более эффективны, чем утверждения типа с базовым типом типа указателя, если базовый тип является типом большого размера.

Поэтому, пожалуйста, старайтесь не упаковывать значения большого размера, вместо этого упаковывайте их указатели.

#### Значения `[]T` нельзя напрямую преобразовать в `[]I` , даже если тип `T` реализует тип интерфейса `I` .

Например, иногда нам может понадобиться преобразовать `[]string` значение в `[]interface{}` тип. В отличие от некоторых других языков, здесь нет прямого способа преобразования. Мы должны сделать преобразование вручную в цикле:

```go
package main

import "fmt"

func main() {
	words := []string{
		"Go", "is", "a", "high",
		"efficient", "language.",
	}

	// The prototype of fmt.Println function is
	// func Println(a ...interface{}) (n int, err error).
	// So words... can't be passed to it as the argument.

	// fmt.Println(words...) // not compile

	// Convert the []string value to []interface{}.
	iw := make([]interface{}, 0, len(words))
	for _, w := range words {
		iw = append(iw, w)
	}
	fmt.Println(iw...) // compiles okay
}

```

#### Каждый метод, указанный в типе интерфейса, соответствует неявной функции.

Для каждого метода с именем `m` в наборе методов, определяемом типом интерфейса `I` , компиляторы будут неявно объявлять функцию с именем `I.m` , которая имеет на один входной параметр типа `I` больше, чем метод `m` . Дополнительный параметр является первым входным параметром функции `I.m` . Предположим `i` , что значение интерфейса равно `I` , тогда вызов метода `i.m(...)` эквивалентен вызову функции `I.m(i, ...)` .

Пример:

```go
package main

import "fmt"

type I interface {
	m(int)bool
}

type T string
func (t T) m(n int) bool {
	return len(t) > n
}

func main() {
	var i I = T("gopher")
	fmt.Println(i.m(5))                        // true
	fmt.Println(I.m(i, 5))                     // true
	fmt.Println(interface{m(int)bool}.m(i, 5)) // true

	// The following lines compile okay,
	// but will panic at run time.
	I(nil).m(5)
	I.m(nil, 5)
	interface {m(int) bool}.m(nil, 5)
}
```