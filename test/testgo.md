# Тесты в Go

Posted on 2018, Jul 30 4 мин. чтения

В го есть встроенный фреймворк для тестирования [(https://golang.org/pkg/testing/).](https://golang.org/pkg/testing/) Для создания нового набора тестов создаем go файл, название которого заканчивается как *\_test.go* Т.е., к примеру, основной код у нас в файле good.go, а тесты в good\_test.go. В названии файла с тестами не обязательно привязываться к названию основного файла. Допустим у нас в файле good.go какой\-то сложный функционал и мы создаем к нему 2 набора тестов, которые сгрупированы в 2 файлах suite1\_test.go, suite2\_test.go.

Чтобы тест выполнился фреймворком, функция должна иметь определенную сигнатуру:

```golang
func TestXxx(t *testing.T)
```

Имя начинается с Test и один параметр \*testing.T

Например, в основом коде у нас есть функция

```golang
func Power(a, b int) int {
}
```

и тест

```golang
func TestNegativeNumbers(*testing.T) {
    res:= Mult(-2, -2)
    if res != 4 {
		t.Errorf("Expected %d, but was %d", 4, res)
    }
}
```

В go часто практикуется табличный подход для тестов (table\-driven), когда кейсы для тестируемой функции собираются в таблицу и дальше по ней итерируется.

```golang
package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
```

Так же есть возможность организовать иерархию тестов

```golang
func TestFoo(t *testing.T) {
    // <setup code>
    t.Run("A=1", func(t *testing.T) { ... })
    t.Run("A=2", func(t *testing.T) { ... })
    t.Run("B=1", func(t *testing.T) { ... })
    // <tear-down code>
}
```

Основной тест `TestFoo` последовательно запускает набор тестов. В начале `TestFoo` можно добавить инициализирующий код, который выполнится перед началом тестов, и в конце код очистки. К примеру, перед этим набором тестов пишем какие\-то тестовые данные в базу, на них прогоняем тесты и после последнего теста `B1` очищаем базу.

Первый параметр в Run \- это название дочернего теста. Если запустить тесты с опцией \-v

```
go test -v ./...

```

То увидим что\-то в таком духе

```
=== RUN   TestFoo
=== RUN   TestFoo/A=1
=== RUN   TestFoo/A=2
=== RUN   TestFoo/B=1

```

Можно управлять тем, какие тесты мы отправляем на запуск: \- запускаем все тесты

```
go test -run ''

```

*   запускаем родительские тесты в названии которых есть “Foo”, т.е. “TestFoo”, “TestFooBar”. `go test -run Foo`
*   находим родительские тесты в названии которых есть “Foo”, и которые содержат дочерние тесты в названии которых есть “A=”. Запускаем эти дочерние тесты и связанные с ними родительские `go test -run Foo/A=`
*   запускаем родительские тесты, у которых есть дочерние тесты с названием “A=1” `go test -run /A=1`

Тесты можно выполнять параллельно

```golang
func TestFooParallel(t *testing.T) {
	// <setup code>
	setUp()

	t.Run("A=1", func(t *testing.T) {
		t.Parallel()
		t.Log("A1")
	})
	t.Run("A=2", func(t *testing.T) {
		t.Parallel()
		t.Log("A2")
	})
	t.Run("B=1", func(t *testing.T) {
		t.Parallel()
		t.Log("B1")
	})

	// <tear-down code>
	tearDown()
}
```

по выводу видно, что тесты выполняются не последовательно

```bash
--- PASS: TestFooParallel (0.00s)
    --- PASS: TestFooParallel/A=1 (0.00s)
        log_test.go:50: A1
    --- PASS: TestFooParallel/B=1 (0.00s)
        log_test.go:58: B1
    --- PASS: TestFooParallel/A=2 (0.00s)
        log_test.go:54: A2
```

Кроме того можно стартовать группу параллельных тестов

```golang
func TestTeardownParallel(t *testing.T) {
    // This Run will not return until the parallel tests finish.
    t.Run("group", func(t *testing.T) {
        t.Run("Test1", parallelTest1)
        t.Run("Test2", parallelTest2)
        t.Run("Test3", parallelTest3)
    })
    // <tear-down code>
}
```

При этом в TestTeardownParallel код за `Run("group")` не будет выполняться, пока все параллельные тесты не выполнятся (Test1, Test2, Test3). Таким образом перед `Run("group")` можно добавить код инициализации, а после `Run("group")` код очистки для этой группы тестов.

Если нужен глобальная инициализация и очистка для всего тестового пакаджа (для каждого пакаджа отдельно), то можно использовать функцию `TestMain`:

```golang
func TestMain(m *testing.M) {

	fmt.Println("setup")
	res := m.Run()
	fmt.Println("tear-down")

	os.Exit(res)
}
```

код `fmt.Println("setup")` выполнится до всех тестов в пакадже и `fmt.Println("tear-down")` после.

Еще хочется отметить, что команда go build ./… собирает, только основной код, тесты при этом не собираются, и если там закралась ошибка компиляции, то вы будете не в курсе. Если вы пишете тесты и нужно просто собрать их без запуска, то можно воспользоваться забавным хаком \- в команде `-run` указать не существующий тест.

```bash
go test -run none
```

Ссылки:

*   [https://golang.org/pkg/testing/](https://golang.org/pkg/testing/)
*   [https://golang.org/cmd/go/#hdr\-Testing\_functions](https://golang.org/cmd/go/#hdr-Testing_functions)
*   [https://golang.org/cmd/go/#hdr\-Testing\_flags](https://golang.org/cmd/go/#hdr-Testing_flags)
*   [https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721](https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721)
*   [https://github.com/golang/go/wiki/TableDrivenTests](https://github.com/golang/go/wiki/TableDrivenTests)
*   [https://blog.golang.org/subtests](https://blog.golang.org/subtests)
