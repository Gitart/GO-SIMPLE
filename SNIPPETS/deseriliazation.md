# Go: десериализация JSON с неправильной типизацией, или как обходить ошибки разработчиков API

*   [Go](https://habr.com/ru/hub/go/ "Вы подписаны на этот хаб")

*   [Из песочницы](https://habr.com/ru/sandbox/ "Перейти в песочницу")

![image](https://habrastorage.org/webt/pf/hz/xm/pfhzxmp9antvvecnexqujg6pkzk.jpeg)

Недавно мне довелось разрабатывать на Go http-клиент для сервиса, предоставляющего REST API с json-ом в роли формата кодирования. Стандартная задача, но в ходе работы мне пришлось столкнуться с нестандартной проблемой. Рассказываю в чем суть.

Как известно, формат json имеет типы данных. Четыре примитивных: строка, число, логический, null; и два структурных типа: объект и массив. В данном случае нас интересуют примитивные типы. Вот пример json кода с четырьмя полями разных типов:

```
{
	"name":"qwerty",
	"price":258.25,
	"active":true,
	"description":null,
}
```

Как видно в примере, строковое значение заключается в кавычки. Числовое — не имеет кавычек. Логический тип может иметь только одно из двух значений: true или false (без кавычек). И тип null соответственно имеет значение null (также без кавычек).

А теперь собственно сама проблема. В какой-то момент, при детальном рассмотрении получаемого от стороннего сервиса json-кода, я обнаружил, что одно из полей (назовем его price) помимо числового значения периодически имеет строковое значение (число в кавычках). Т. е. один и тот же запрос с разными параметрами может вернуть число в виде числа, а может вернуть это же число в виде строки. Ума не приложу, как на том конце организован код, возвращающий такие результаты, но видимо, это связано с тем, что сервис сам является агрегатором и тянет данные из разных источников, а разработчики не привели json ответа сервера к единому формату. Тем не менее, надо работать с тем что есть.

Но далее меня ждало еще большее удивление. Логическое поле (назовем его active), помимо значений true и false, возвращало строковые значения «true», «false», и даже числовые 1 и 0 (истина и ложь соответственно).

Вся эта путаница с типами данных не была бы критичной, если бы я обрабатывал json скажем на слаботипизированном PHP, но Go имеет сильную типизацию, и требует четкого указания типа десериализуемого поля. В итоге возникла необходимость реализовать механизм, позволяющий в процессе десериализации преобразовывать все значения поля active в логический тип, и любые значения поля price — в числовой.

Начнем с числового поля price.

Предположим что у нас есть json-код следующего вида:

```
[
	{"id":1,"price":2.58},
	{"id":2,"price":7.15}
]
```

Т. е. json содержит массив объектов с двумя полями числового типа. Стандартный код десериализации данного json-а на Go выглядит так:

```
type Target struct {
	Id    int     `json:"id"`
	Price float64 `json:"price"`
}

func main() {
	jsonString := `[{"id":1,"price":2.58},
					{"id":4,"price":7.15}]`

	targets := []Target{}

	err := json.Unmarshal([]byte(jsonString), &targets)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, t := range targets {
		fmt.Println(t.Id, "-", t.Price)
	}
}
```

В данном коде мы десериализуем поле id в тип int, а поле price в тип float64. Теперь предположим, что наш json-код выглядит так:

```
[
	{"id":1,"price":2.58},
	{"id":2,"price":"2.58"},
	{"id":3,"price":7.15},
	{"id":4,"price":"7.15"}
]
```

Т. е. поле price содержит значения как числового типа, так и строкового. В данном случае только числовые значения поля price могут быть декодированы в тип float64, строковые же значения вызовут ошибку о несовместимости типов. Это значит, что ни float64, ни любой другой примитивный тип не подходят для десериализации данного поля, и нам необходим свой пользовательский тип со своей логикой десериализации.

В качестве такого типа объявим структуру CustomFloat64 с единственным полем Float64 типа float64.

```
type CustomFloat64 struct{
	Float64 float64
}
```

И сразу укажем данный тип для поля Price в структуре Target:

```
type Target struct {
	Id    int           `json:"id"`
	Price CustomFloat64 `json:"price"`
}
```

Теперь необходимо описать собственную логику декодирования поля c типом CustomFloat64.

В пакете «encoding/json» предусмотрены два специальных метода: [MarshalJSON](https://godoc.org/encoding/json#RawMessage.MarshalJSON) и [UnmarshalJSON](https://godoc.org/encoding/json#RawMessage.UnmarshalJSON), которые и предназначены для кастомизации логики кодирования и декодирования конкретного пользовательского типа данных. Достаточно переопределить эти методы и описать собственную реализацию.

Переопределим метод UnmarshalJSON для произвольного типа CustomFloat64. При этом необходимо строго следовать сигнатуре метода, иначе он просто не сработает, а главное не выдаст при этом ошибку.

```
func (cf *CustomFloat64) UnmarshalJSON(data []byte) error {
```

На входе данный метод принимает слайс байт (data), в котором содержится значение конкретного поля декодируемого json. Если преобразовать данную последовательность байт в строку, то мы увидим значение поля именно в том виде, в каком оно записано в json. Т. е. если это строковый тип, то мы увидим именно строку с двойными кавычками («258»), если числовой тип, то увидим строку без кавычек (258).

Чтобы отличить числовое значение от строкового, необходимо проверить, является ли первый символ кавычкой. Так как символ двойной кавычки в таблице UNICODE занимает один байт, нам достаточно проверить первый байт слайса data, сравнив его с номером символа в UNICODE. Это номер 34. Обратите внимание, что в общем случае, символ не равнозначен байту, так как может занимать больше одного байта. Символу в Go равнозначен тип rune (руна). В нашем же случае достаточно данного условия:

```
if data[0] == 34 {
```

Если условие выполняется, то значение имеет строковый тип, и нам необходимо получить строку между кавычками, т. е. слайс байт между первым и последним байтом. Именно в этом слайсе содержится числовое значение, которое может быть декодировано в примитивный тип float64. Это значит, что мы можем применить к нему метод json.Unmarshal, при этом результат сохраняя в поле Float64 структуры CustomFloat64.

```
err := json.Unmarshal(data[1:len(data)-1], &cf.Float64)
```

Если же слайс data начинается не с кавычки, то значит в нем уже содержится числовой тип данных, и мы можем применить метод json.Unmarshal непосредственно ко всему слайсу data.

```
err := json.Unmarshal(data, &cf.Float64)
```

Вот полный код метода UnmarshalJSON:

```
func (cf *CustomFloat64) UnmarshalJSON(data []byte) error {
	if data[0] == 34 {
		err := json.Unmarshal(data[1:len(data)-1], &cf.Float64)
		if err != nil {
			return errors.New("CustomFloat64: UnmarshalJSON: " + err.Error())
		}
	} else {
		err := json.Unmarshal(data, &cf.Float64)
		if err != nil {
			return errors.New("CustomFloat64: UnmarshalJSON: " + err.Error())
		}
	}
	return nil
}
```

В итоге, с применением метода json.Unmarshal к нашему json-коду, все значения поля price будут прозрачно для нас преобразованы в примитивный тип float64, и результат запишется в поле Float64 структуры CustomFloat64.

Теперь нам может понадобиться преобразовать структуру Target обратно в json. Но, если мы применяем метод json.Marshal непосредственно к типу CustomFloat64, то сериализуем данную структуру в виде объекта. Нам же необходимо кодировать поле price в числовое значение. Чтобы кастомизировать логику кодирования пользовательского типа CustomFloat64, реализуем для него метод MarshalJSON, при этом строго соблюдая сигнатуру метода:

```
func (cf CustomFloat64) MarshalJSON() ([]byte, error) {
	json, err := json.Marshal(cf.Float64)
	return json, err
}
```

Все, что нужно сделать в этом методе, это опять же использовать метод json.Marshal, но уже применять его не к структуре CustomFloat64, а к ее полю Float64. Из метода возвращаем полученный слайс байт и ошибку.

Вот полный код с выводом результатов сериализации и десериализации (проверка ошибок опущена для краткости, номер байта с символом двойных кавычек вынесен в константу):

```
package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type CustomFloat64 struct {
	Float64 float64
}

const QUOTES_BYTE = 34

func (cf *CustomFloat64) UnmarshalJSON(data []byte) error {
	if data[0] == QUOTES_BYTE {
		err := json.Unmarshal(data[1:len(data)-1], &cf.Float64)
		if err != nil {
			return errors.New("CustomFloat64: UnmarshalJSON: " + err.Error())
		}
	} else {
		err := json.Unmarshal(data, &cf.Float64)
		if err != nil {
			return errors.New("CustomFloat64: UnmarshalJSON: " + err.Error())
		}
	}
	return nil
}

func (cf CustomFloat64) MarshalJSON() ([]byte, error) {
	json, err := json.Marshal(cf.Float64)
	return json, err
}

type Target struct {
	Id    int           `json:"id"`
	Price CustomFloat64 `json:"price"`
}

func main() {
	jsonString := `[{"id":1,"price":2.58},
					{"id":2,"price":"2.58"},
					{"id":3,"price":7.15},
					{"id":4,"price":"7.15"}]`

	targets := []Target{}

	_ := json.Unmarshal([]byte(jsonString), &targets)

	for _, t := range targets {
		fmt.Println(t.Id, "-", t.Price.Float64)
	}

	jsonStringNew, _ := json.Marshal(targets)
	fmt.Println(string(jsonStringNew))
}
```

Результат выполнения кода:

```
1 - 2.58
2 - 2.58
3 - 7.15
4 - 7.15
[{"id":1,"price":2.58},{"id":2,"price":2.58},{"id":3,"price":7.15},{"id":4,"price":7.15}]
```

Перейдем ко второй части и реализуем аналогичный код для десериализации json-а с несогласованными значениями логического поля.

Предположим что у нас есть json-код следующего вида:

```
[
	{"id":1,"active":true},
	{"id":2,"active":"true"},
	{"id":3,"active":"1"},
	{"id":4,"active":1},
	{"id":5,"active":false},
	{"id":6,"active":"false"},
	{"id":7,"active":"0"},
	{"id":8,"active":0},
	{"id":9,"active":""}
]
```

В данном случае поле active подразумевает логический тип и наличие только одного из двух значений: true и false. Значения не логического типа необходимо будет преобразовать в логический в ходе десериализации.

В текущем примере мы допускаем следующие соответствия. Значению true соответствуют: true (логическое), «true» (строковое), «1» (строковое), 1 (числовое). Значению false соответствуют: false (логическое), «false» (строковое), «0» (строковое), 0 (числовое), "" (пустая строка).

Для начала объявим целевую структуру десериализации. В качестве типа поля Active сразу указываем пользовательский тип CustomBool:

```
type Target struct {
	Id     int        `json:"id"`
	Active CustomBool `json:"active"`
}
```

CustomBool является структурой с одним единственным полем Bool типа bool:

```
type CustomBool struct {
	Bool bool
}
```

Реализуем для данной структуры метод UnmarshalJSON. Сразу приведу код:

```
func (cb *CustomBool) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"true"`, `true`, `"1"`, `1`:
		cb.Bool = true
		return nil
	case `"false"`, `false`, `"0"`, `0`, `""`:
		cb.Bool = false
		return nil
	default:
		return errors.New("CustomBool: parsing \"" + string(data) + "\": unknown value")
	}
}
```

Так как поле active в нашем случае имеет ограниченное количество значений, мы можем с помощью конструкции switch-case принять решение, о том, чему должно быть равно значение поля Bool структуры CustomBool. Для проверки понадобится всего два блока case. В первом блоке мы проверяем значение на соответствие true, во втором — false.

При записи возможных значений, следует обратить внимание на роль грависа (это такая кавычка на клавише с буквой Ё в английской раскладке). Данный символ позволяет экранировать двойные кавычки в строке. Для наглядности данным символом я обрамил и значения с кавычками и без кавычек. Таким образом, \`false\` соответствует строке false (без кавычек, тип bool в json), а \`«false»\` соответствует строке «false» (с кавычками, тип string в json). Тоже самое и со значениями \`1\` и \`«1»\` Первое — это число 1 (в json записано без кавычек), второе — строка «1» (в json записано с кавычками). Вот эта запись \`""\` — это пустая строка, Т. е. в формате json она выглядит так: "".

Соответствующее значение (true или false) мы записываем непосредственно в поле Bool структуры CustomBool:

```
cb.Bool = true
```

В блоке defaul возвращаем ошибку о том, что поле имеет неизвестное значение:

```
return errors.New("CustomBool: parsing \"" + string(data) + "\": unknown value")
```

Теперь мы можем применять метод json.Unmarshal к нашему json-коду, и значения поля active будут преобразовываться в примитивный тип bool.

Реализуем метод MarshalJSON для структуры CustomBool:

```
func (cb CustomBool) MarshalJSON() ([]byte, error) {
	json, err := json.Marshal(cb.Bool)
	return json, err
}
```

Здесь ничего нового. Метод выполняет сериализацию поля Bool структуры CustomBool.

Вот полный код с выводом результатов сериализации и десериализации (проверка ошибок опущена для краткости):

```
package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type CustomBool struct {
	Bool bool
}

func (cb *CustomBool) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"true"`, `true`, `"1"`, `1`:
		cb.Bool = true
		return nil
	case `"false"`, `false`, `"0"`, `0`, `""`:
		cb.Bool = false
		return nil
	default:
		return errors.New("CustomBool: parsing \"" + string(data) + "\": unknown value")
	}
}

func (cb CustomBool) MarshalJSON() ([]byte, error) {
	json, err := json.Marshal(cb.Bool)
	return json, err
}

type Target struct {
	Id     int        `json:"id"`
	Active CustomBool `json:"active"`
}

func main() {
	jsonString := `[{"id":1,"active":true},
					{"id":2,"active":"true"},
					{"id":3,"active":"1"},
					{"id":4,"active":1},
					{"id":5,"active":false},
					{"id":6,"active":"false"},
					{"id":7,"active":"0"},
					{"id":8,"active":0},
					{"id":9,"active":""}]`

	targets := []Target{}

	_ = json.Unmarshal([]byte(jsonString), &targets)

	for _, t := range targets {
		fmt.Println(t.Id, "-", t.Active.Bool)
	}

	jsonStringNew, _ := json.Marshal(targets)
	fmt.Println(string(jsonStringNew))
}
```

Результат выполнения кода:

```
1 - true
2 - true
3 - true
4 - true
5 - false
6 - false
7 - false
8 - false
9 - false
[{"id":1,"active":true},{"id":2,"active":true},{"id":3,"active":true},{"id":4,"active":true},{"id":5,"active":false},{"id":6,"active":false},{"id":7,"active":false},{"id":8,"active":false},{"id":9,"active":false}]
```

#### Выводы

Во-первых. Переопределение методов MarshalJSON и UnmarshalJSON для произвольных типов данных позволяет кастомизировать сериализацию и десериализацию конкретного поля json-кода. Помимо указанных вариантов использования, данные функции применяются для работы с полями, допускающими значение null.

Во-вторых. Формат текстового кодирования json — это широко используемый инструмент для обмена информацией, и одним из его преимуществ перед другими форматами является наличие типов данных. За соблюдением этих типов надо строго следить.
