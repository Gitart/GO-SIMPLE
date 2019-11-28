# 50 Shades of Go: ловушки, ошибки и типичные ошибки для разработчиков New Golang

#### 50 оттенков го на других языках

*   Перевод на китайский язык: [запись в блоге](https://wuyin.io/2018/03/07/50-shades-of-golang-traps-gotchas-mistakes/) , [gmentfault](https://segmentfault.com/a/1190000013739000) (by [wuYin](https://twitter.com/@wuYinBest) ) \- нужны обновления
*   Другой китайский перевод: [сообщение в блоге](http://colobu.com/2015/09/07/gotchas-and-common-mistakes-in-go-golang/) (от Shadowwind LEY) \- нужны обновления
*   Русский перевод: [сообщение в блоге](https://habr.com/company/mailru/blog/314804/) ( [Илья Ожерелиев](https://habr.com/users/3vilhamst3r/) , [Mail.Ru Group Blog](https://habr.com/company/mailru/) ) \- нужны обновления

#### обзор

Go \- это простой и забавный язык, но, как и любой другой язык, в нем есть несколько ошибок ... Многие из этих ошибок не являются полностью ошибкой Go. Некоторые из этих ошибок являются естественными ловушками, если вы пришли с другого языка. Другие из\-за ошибочных предположений и недостающих деталей.

Многие из этих ошибок могут показаться очевидными, если вы потратите время на изучение языка, читая официальные спецификации, вики, обсуждения в списках рассылки, множество замечательных постов и презентаций Роба Пайка и исходный код. Не все начинаются одинаково, и это нормально. Если вы новичок в Go, информация здесь сэкономит вам часы отладки вашего кода.

[Всего начинающих](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#total_beginner) :

*   [Открывающая скобка не может быть размещена на отдельной линии](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#opening_braces)
*   [Неиспользуемые переменные](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#unused_vars)
*   [Неиспользованный импорт](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#unused_imports)
*   [Короткие объявления переменных могут использоваться только внутри функций](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#short_vars)
*   [Переопределение переменных с помощью кратких объявлений переменных](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#vars_redeclare)
*   [Невозможно использовать короткие объявления переменных для установки значений полей](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#short_fields)
*   [Случайное изменение переменных](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#vars_shadow)
*   [Нельзя использовать «ноль» для инициализации переменной без явного типа](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#nil_init)
*   [Использование "ноль" ломтиков и карт](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#nil_slices_maps)
*   [Емкость карты](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#map_cap)
*   [Струны не могут быть "ноль"](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#nil_strings)
*   [Аргументы функций массива](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#array_func_args)
*   [Неожиданные значения в разделах Slice и Array](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#unexpected_slice_arr_vals)
*   [Ломтики и массивы являются одномерными](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#one_dim_slice_arr)
*   [Доступ к несуществующим ключам карты](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#map_key_ne)
*   [Струны неизменны](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#imm_strings)
*   [Преобразования между строками и байтовыми срезами](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#string_byte_slice_conv)
*   [Строки и оператор индекса](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#string_idx)
*   [Строки не всегда в тексте UTF8](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#strings_na_utf)
*   [Длина строки](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#string_length)
*   [Отсутствует запятая в многострочных литералах Slice / Array / Map](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#mline_lit_comma)
*   [log.Fatal и log.Panic Do Больше, чем Log](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#log_fatal_exit)
*   [Операции со встроенной структурой данных не синхронизированы](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#coll_no_sync)
*   [Значения итерации для строк в предложениях «range»](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#string_range_vals)
*   [Итерация по карте с помощью предложения "for range"](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#map_range)
*   [Поведенческое поведение в «переключателях»](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#switch_fall)
*   [Увеличение и уменьшение](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#inc_dec)
*   [Побитовый оператор НЕ](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#bit_not)
*   [Различия в приоритетах операторов](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#op_precedence)
*   [Неэкспортированные поля структуры не кодируются](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#unexp_struct_field_enc)
*   [Выход из приложения с активными программами](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#gor_app_exit)
*   [Отправка на небуферизованный канал возвращается, как только целевой приемник готов](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#unbuf_ch_send_done)
*   [Отправка на закрытый канал вызывает панику](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#closed_ch_send)
*   [Использование "ноль" каналов](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#using_nil_ch)
*   [Методы с получателями значений не могут изменить исходное значение](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#method_val_receiver)

[Средний Новичок](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#intermediate_beginner) :

*   [Закрытие HTTP ответа тела](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#close_http_resp_body)
*   [Закрытие HTTP\-соединений](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#close_http_conn)
*   [JSON Encoder добавляет символ новой строки](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#json_encode_newline)
*   [Пакет JSON исключает специальные символы HTML в ключах и строковых значениях](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#json_escape_html)
*   [Распаковка чисел JSON в значения интерфейса](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#json_num)
*   [Строковые значения JSON не будут в порядке с шестнадцатеричным или другими не\-UTF8 Escape\-последовательностями](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#json_utf8_strings)
*   [Сравнение структур, массивов, фрагментов и карт](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#compare_struct_arr_slice_map)
*   [Восстановление от паники](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#panic_recover)
*   [Обновление и ссылка на значения элементов в срезах, массивах и картах «для диапазона»](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#range_val_update)
*   [«Скрытые» данные в срезах](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#slice_hidden_data)
*   [Повреждение данных среза](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#slice_data_corruption)
*   ["Несвежие" ломтики](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#stale_slices)
*   [Объявления типов и методы](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#type_decl_methods)
*   [Выход из «для переключателя» и «для выбора» кодовых блоков](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#deep_for_breakout)
*   [Переменные итераций и замыкания в выражениях "for"](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#closure_for_it_vars)
*   [Оценка аргумента вызова отложенной функции](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#deferred_calls)
*   [Выполнение отложенного вызова функции](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#deferred_call_exe)
*   [Неудачные утверждения типа](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#failed_type_assert)
*   [Заблокированные рутины и утечки ресурсов](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#blocked_goroutines)
*   [Один и тот же адрес для разных переменных нулевого размера](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#zero_size_var)
*   [Первое использование йоты не всегда начинается с нуля](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#iota_zero)

[Продвинутый начинающий](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#advanced_beginner) :

*   [Использование методов получения указателя на экземплярах значений](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#ptr_receiver_val_inst)
*   [Обновление полей значений карты](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#map_value_field_update)
*   [Интерфейсы "nil" и значения интерфейсов "nil"](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#nil_in_nil_in_vals)
*   [Переменные стека и кучи](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#stack_heap_vars)
*   [GOMAXPROCS, параллелизм и параллелизм](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#gomaxprocs)
*   [Перезапись операции чтения и записи](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#rw_reorder)
*   [Упреждающее планирование](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#psched)

[Cgo (он же Храбрый Новичок)](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#cgo) :

*   [Импортировать блоки C и Multiline Import](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#cgo_multiline_import_c)
*   [Нет пустых строк между Import C и Cgo Комментарии](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#cgo_noblanks)
*   [Невозможно вызвать функции C с переменными аргументами](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#cgo_no_var_args)

#### Ловушки, ошибки и распространенные ошибки

###### Открывающая скобка не может быть размещена на отдельной линии

*   уровень: начинающий

На большинстве других языков, которые используют фигурные скобки, вы можете выбрать, где их разместить. Иди по\-другому. Вы можете поблагодарить автоматическую инъекцию точки с запятой (без предупреждения) за это поведение. Да, у Go есть точки с запятой :\-)

Сбой:

```
package main

import "fmt"

func main()
{ //error, can't have the opening brace on a separate line
    fmt.Println("hello there!")
}

```

Ошибка компиляции:

> /tmp/sandbox826898458/main.go:6: синтаксическая ошибка: неожиданная точка с запятой или символ новой строки перед {

Работает:

```
package main

import "fmt"

func main() {
    fmt.Println("works!")
}

```

###### Неиспользуемые переменные

*   уровень: начинающий

Если у вас есть неиспользуемая переменная, ваш код не скомпилируется. Хотя есть исключение. Вы должны использовать переменные, которые вы объявляете внутри функций, но это нормально, если у вас есть неиспользуемые глобальные переменные. Также нормально иметь неиспользуемые аргументы функции.

Если вы присвоите новое значение неиспользуемой переменной, ваш код все равно не будет скомпилирован. Вам нужно как\-то использовать значение переменной, чтобы компилятор был доволен.

Сбой:

```
package main

var gvar int //not an error

func main() {
    var one int   //error, unused variable
    two := 2      //error, unused variable
    var three int //error, even though it's assigned 3 on the next line
    three = 3

    func(unused string) {
        fmt.Println("Unused arg. No compile error")
    }("what?")
}

```

Ошибки компиляции:

> /tmp/sandbox473116179/main.go:6: один объявлен и не используется /tmp/sandbox473116179/main.go:7: два объявлен и не используется /tmp/sandbox473116179/main.go:8: три объявлены и не используются

Работает:

```
package main

import "fmt"

func main() {
    var one int
    _ = one

    two := 2
    fmt.Println(two)

    var three int
    three = 3
    one = three

    var four int
    four = four
}

```

Другой вариант \- закомментировать или удалить неиспользуемые переменные :\-)

###### Неиспользованный импорт

*   уровень: начинающий

Ваш код не сможет скомпилироваться, если вы импортируете пакет без использования каких\-либо его экспортируемых функций, интерфейсов, структур или переменных.

Если вам действительно нужен импортированный пакет, вы можете использовать пустой идентификатор в `_` качестве имени пакета, чтобы избежать этой ошибки компиляции. Пустой идентификатор используется для импорта пакетов с учетом их побочных эффектов.

Сбой:

```
package main

import (
    "fmt"
    "log"
    "time"
)

func main() {
}

```

Ошибки компиляции:

> /tmp/sandbox627475386/main.go:4: импортировано и не используется: "fmt" /tmp/sandbox627475386/main.go:5: импортировано и не используется: "log" /tmp/sandbox627475386/main.go:6: импортировано и не используется: «время»

Работает:

```
package main

import (
    _ "fmt"
    "log"
    "time"
)

var _ = log.Println

func main() {
    _ = time.Now
}

```

Другой вариант \- удалить или закомментировать неиспользованный импорт :\-) [`goimports`](http://godoc.org/golang.org/x/tools/cmd/goimports) Инструмент может помочь вам в этом.

###### Короткие объявления переменных могут использоваться только внутри функций

*   уровень: начинающий

Сбой:

```
package main

myvar := 1 //error

func main() {
}

```

Ошибка компиляции:

> /tmp/sandbox265716165/main.go:3: оператор без объявления вне тела функции

Работает:

```
package main

var myvar = 1

func main() {
}

```

###### Переопределение переменных с помощью кратких объявлений переменных

*   уровень: начинающий

Вы не можете повторно объявить переменную в отдельном операторе, но это разрешено в объявлениях с несколькими переменными, где также объявлена ​​хотя бы одна новая переменная.

Переименованная переменная должна находиться в том же блоке, иначе вы получите теневую переменную.

Сбой:

```
package main

func main() {
    one := 0
    one := 1 //error
}

```

Ошибка компиляции:

> /tmp/sandbox706333626/main.go:5: нет новых переменных в левой части: =

Работает:

```
package main

func main() {
    one := 0
    one, two := 1,2

    one,two = two,one
}

```

###### Невозможно использовать короткие объявления переменных для установки значений полей

*   уровень: начинающий

Сбой:

```
package main

import (
  "fmt"
)

type info struct {
  result int
}

func work() (int,error) {
    return 13,nil
  }

func main() {
  var data info

  data.result, err := work() //error
  fmt.Printf("info: %+v\n",data)
}

```

Ошибка компиляции:

> prog.go: 18: безымянный data.result в левой части: =

Несмотря на то, что есть билет для решения этой проблемы, он вряд ли изменится, потому что Робу Пайку нравится «как есть» :\-)

Используйте временные переменные или предварительно объявите все свои переменные и используйте стандартный оператор присваивания.

Работает:

```
package main

import (
  "fmt"
)

type info struct {
  result int
}

func work() (int,error) {
    return 13,nil
  }

func main() {
  var data info

  var err error
  data.result, err = work() //ok
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Printf("info: %+v\n",data) //prints: info: {result:13}
}

```

###### Случайное изменение переменных

*   уровень: начинающий

Синтаксис объявления коротких переменных настолько удобен (особенно для динамического языка), что его легко рассматривать как обычную операцию присваивания. Если вы сделаете эту ошибку в новом блоке кода, ошибки компилятора не будет, но ваше приложение не будет делать то, что вы ожидаете.

```
package main

import "fmt"

func main() {
    x := 1
    fmt.Println(x)     //prints 1
    {
        fmt.Println(x) //prints 1
        x := 2
        fmt.Println(x) //prints 2
    }
    fmt.Println(x)     //prints 1 (bad if you need 2)
}

```

Это очень распространенная ловушка даже для опытных разработчиков Go. Это легко сделать, и это может быть трудно обнаружить.

Вы можете использовать [`vet`](http://godoc.org/golang.org/x/tools/cmd/vet) команду, чтобы найти некоторые из этих проблем. По умолчанию `vet` не выполняется проверка теневых переменных. Обязательно используйте `-shadow` флаг: `go tool vet -shadow your_file.go`

Обратите внимание, что `vet` команда не будет сообщать обо всех скрытых переменных. Используйте [`go-nyet`](https://github.com/barakmich/go-nyet) для более агрессивного обнаружения теневых переменных.

###### Нельзя использовать «ноль» для инициализации переменной без явного типа

*   уровень: начинающий

Идентификатор «ноль» может использоваться в качестве «нулевого значения» для интерфейсов, функций, указателей, карт, срезов и каналов. Если вы не укажете тип переменной, компилятор не сможет скомпилировать ваш код, потому что он не может угадать тип.

Сбой:

```
package main

func main() {
    var x = nil //error

    _ = x
}

```

Ошибка компиляции:

> /tmp/sandbox188239583/main.go:4: использование нетипизированного nil

Работает:

```
package main

func main() {
    var x interface{} = nil

    _ = x
}

```

###### Использование "ноль" ломтиков и карт

*   уровень: начинающий

Это нормально, чтобы добавить элементы к «нулевому» срезу, но то же самое с картой вызовет панику во время выполнения.

Работает:

```
package main

func main() {
    var s []int
    s = append(s,1)
}

```

Сбой:

```
package main

func main() {
    var m map[string]int
    m["one"] = 1 //error

}

```

###### Емкость карты

*   уровень: начинающий

Вы можете указать емкость карты при ее создании, но вы не можете использовать `cap()` функцию на картах.

Сбой:

```
package main

func main() {
    m := make(map[string]int,99)
    cap(m) //error
}

```

Ошибка компиляции:

> /tmp/sandbox326543983/main.go:5: неверный аргумент m (тип map \[string\] int) для cap

###### Струны не могут быть "ноль"

*   уровень: начинающий

Это недоработка для разработчиков, которые привыкли присваивать «нулевые» идентификаторы строковым переменным.

Сбой:

```
package main

func main() {
    var x string = nil //error

    if x == nil { //error
        x = "default"
    }
}

```

Ошибки компиляции:

> /tmp/sandbox630560459/main.go:4: нельзя использовать nil в качестве строки типа в присваивании /tmp/sandbox630560459/main.go:6: недопустимая операция: x == nil (несоответствие типов string и nil)

Работает:

```
package main

func main() {
    var x string //defaults to "" (zero value)

    if x == "" {
        x = "default"
    }
}

```

###### Аргументы функций массива

*   уровень: начинающий

Если вы разработчик C или C ++, то для вас это указатели. При передаче массивов в функции функции ссылаются на одну и ту же ячейку памяти, чтобы они могли обновлять исходные данные. Массивы в Go являются значениями, поэтому при передаче массивов функциям функции получают копию исходных данных массива. Это может быть проблемой, если вы пытаетесь обновить данные массива.

```
package main

import "fmt"

func main() {
    x := [3]int{1,2,3}

    func(arr [3]int) {
        arr[0] = 7
        fmt.Println(arr) //prints [7 2 3]
    }(x)

    fmt.Println(x) //prints [1 2 3] (not ok if you need [7 2 3])
}

```

Если вам нужно обновить исходные данные массива, используйте типы указателей массива.

```
package main

import "fmt"

func main() {
    x := [3]int{1,2,3}

    func(arr *[3]int) {
        (*arr)[0] = 7
        fmt.Println(arr) //prints &[7 2 3]
    }(&x)

    fmt.Println(x) //prints [7 2 3]
}

```

Другим вариантом является использование ломтиков. Даже если ваша функция получает копию переменной slice, она все равно ссылается на исходные данные.

```
package main

import "fmt"

func main() {
    x := []int{1,2,3}

    func(arr []int) {
        arr[0] = 7
        fmt.Println(arr) //prints [7 2 3]
    }(x)

    fmt.Println(x) //prints [7 2 3]
}

```

###### Неожиданные значения в разделах Slice и Array

*   уровень: начинающий

Это может произойти, если вы привыкли использовать операторы for\-in или foreach на других языках. Предложение «range» в Go отличается. Он генерирует два значения: первое значение \- это индекс элемента, а второе значение \- данные элемента.

Плохо:

```
package main

import "fmt"

func main() {
    x := []string{"a","b","c"}

    for v := range x {
        fmt.Println(v) //prints 0, 1, 2
    }
}

```

Хорошо:

```
package main

import "fmt"

func main() {
    x := []string{"a","b","c"}

    for _, v := range x {
        fmt.Println(v) //prints a, b, c
    }
}

```

###### Ломтики и массивы являются одномерными

*   уровень: начинающий

Может показаться, что Go поддерживает многомерные массивы и фрагменты, но это не так. Тем не менее, возможно создание массивов массивов или срезов слайсов. Для приложений числовых вычислений, которые используют динамические многомерные массивы, это далеко не идеально с точки зрения производительности и сложности.

Вы можете создавать динамические многомерные массивы, используя необработанные одномерные массивы, фрагменты «независимых» фрагментов и фрагменты «общих данных».

Если вы используете необработанные одномерные массивы, вы отвечаете за индексацию, проверку границ и перераспределение памяти, когда массивы должны расти.

Создание динамического многомерного массива с использованием срезов «независимых» срезов является двухэтапным процессом. Сначала вы должны создать внешний срез. Затем вы должны выделить каждый внутренний срез. Внутренние срезы не зависят друг от друга. Вы можете вырастить и сжать их, не затрагивая другие внутренние кусочки.

```
package main

func main() {
    x := 2
    y := 4

    table := make([][]int,x)
    for i:= range table {
        table[i] = make([]int,y)
    }
}

```

Создание динамического многомерного массива с использованием срезов «общих данных» представляет собой трехэтапный процесс. Во\-первых, вы должны создать «контейнерный» фрагмент данных, который будет содержать необработанные данные. Затем вы создаете внешний срез. Наконец, вы инициализируете каждый внутренний срез путем повторного среза необработанных данных.

```
package main

import "fmt"

func main() {
    h, w := 2, 4

    raw := make([]int,h*w)
    for i := range raw {
        raw[i] = i
    }
    fmt.Println(raw,&raw[4])
    //prints: [0 1 2 3 4 5 6 7] <ptr_addr_x>

    table := make([][]int,h)
    for i:= range table {
        table[i] = raw[i*w:i*w + w]
    }

    fmt.Println(table,&table[1][0])
    //prints: [[0 1 2 3] [4 5 6 7]] <ptr_addr_x>
}

```

There's a spec/proposal for multi\-dimensional arrays and slices, but it looks like it's a low priority feature at this point in time.

###### Accessing Non\-Existing Map Keys

*   level: beginner

This is a gotcha for developers who expect to get "nil" identifiers (like it's done in other languages). The returned value will be "nil" if the "zero value" for the corresponding data type is "nil", but it'll be different for other data types. Checking for the appropriate "zero value" can be used to determine if the map record exists, but it's not always reliable (e.g., what do you do if you have a map of booleans where the "zero value" is false). The most reliable way to know if a given map record exists is to check the second value returned by the map access operation.

Bad:

```
package main

import "fmt"

func main() {
    x := map[string]string{"one":"a","two":"","three":"c"}

    if v := x["two"]; v == "" { //incorrect
        fmt.Println("no entry")
    }
}

```

Good:

```
package main

import "fmt"

func main() {
    x := map[string]string{"one":"a","two":"","three":"c"}

    if _,ok := x["two"]; !ok {
        fmt.Println("no entry")
    }
}

```

###### Струны неизменны

*   уровень: начинающий

Попытка обновить отдельный символ в строковой переменной с помощью оператора индекса приведет к сбою. Строки являются байтовыми срезами только для чтения (с несколькими дополнительными свойствами). Если вам нужно обновить строку, тогда используйте байтовый фрагмент вместо преобразования его в тип строки, когда это необходимо.

Сбой:

```
package main

import "fmt"

func main() {
    x := "text"
    x[0] = 'T'

    fmt.Println(x)
}

```

Ошибка компиляции:

> /tmp/sandbox305565531/main.go:7: невозможно назначить x \[0\]

Работает:

```
package main

import "fmt"

func main() {
    x := "text"
    xbytes := []byte(x)
    xbytes[0] = 'T'

    fmt.Println(string(xbytes)) //prints Text
}

```

Обратите внимание, что это не совсем правильный способ обновления символов в текстовой строке, поскольку данный символ может храниться в нескольких байтах. Если вам нужно обновить текстовую строку, сначала преобразуйте ее в фрагмент руны. Даже с кусочками рун один символ может охватывать несколько рун, что может произойти, например, если у вас есть символы с серьезным акцентом. Эта сложная и неоднозначная природа «символов» является причиной, по которой строки Go представляются в виде байтовых последовательностей.

###### Преобразования между строками и байтовыми срезами

*   уровень: начинающий

Когда вы конвертируете строку в байтовый фрагмент (и наоборот), вы получаете полную копию оригинальных данных. Это не похоже на операцию приведения в других языках, и это не похоже на пересчет, где новая переменная слайса указывает на тот же базовый массив, который использовался в оригинальном байтовом слайсе.

Go имеет несколько оптимизаций для `[]byte` к `string` и `string` для `[]byte` преобразования , чтобы избежать дополнительных ассигнований (с большим количеством оптимизаций в списке TODO).

Первая оптимизация позволяет избежать дополнительных размещений , когда `[]byte` ключи используются для поиска записей в `map[string]` коллекции: `m[string(key)]` .

Вторая оптимизация позволяет избежать дополнительных ассигнований в `for range` пунктах , где строки преобразуются в `[]byte` : `for i,v := range []byte(str) {...}` .

###### Строки и оператор индекса

*   уровень: начинающий

Оператор индекса в строке возвращает значение байта, а не символа (как это делается в других языках).

```
package main

import "fmt"

func main() {
    x := "text"
    fmt.Println(x[0]) //print 116
    fmt.Printf("%T",x[0]) //prints uint8
}

```

Если вам нужен доступ к определенной строке «символов» (кодовые точки / руны Юникода), используйте `for range` предложение. Официальный пакет "unicode / utf8" и экспериментальный пакет utf8string (golang.org/x/exp/utf8string) также полезны. Пакет utf8string включает в себя удобный `At()` метод. Преобразование строки в кусок руны также является опцией.

###### Строки не всегда в тексте UTF8

*   уровень: начинающий

Строковые значения не обязательно должны быть текстом UTF8. Они могут содержать произвольные байты. Единственными временными строками являются UTF8, когда используются строковые литералы. Даже тогда они могут включать другие данные, используя escape\-последовательности.

Чтобы узнать, есть ли у вас текстовая строка UTF8, используйте `ValidString()` функцию из пакета "unicode / utf8".

```
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    data1 := "ABC"
    fmt.Println(utf8.ValidString(data1)) //prints: true

    data2 := "A\xfeC"
    fmt.Println(utf8.ValidString(data2)) //prints: false
}

```

###### Длина строки

*   уровень: начинающий

Допустим, вы разработчик Python и у вас есть следующий фрагмент кода:

```
data = u'♥'
print(len(data)) #prints: 1

```

Когда вы конвертируете его в похожий фрагмент кода Go, вы можете быть удивлены.

```
package main

import "fmt"

func main() {
    data := "♥"
    fmt.Println(len(data)) //prints: 3
}

```

Встроенная `len()` функция возвращает количество байтов вместо количества символов, как это делается для строк Unicode в Python.

Чтобы получить те же результаты в Go, используйте `RuneCountInString()` функцию из пакета "unicode / utf8".

```
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    data := "♥"
    fmt.Println(utf8.RuneCountInString(data)) //prints: 1

```

Технически `RuneCountInString()` функция не возвращает количество символов, поскольку один символ может занимать несколько рун.

```
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    data := "é"
    fmt.Println(len(data))                    //prints: 3
    fmt.Println(utf8.RuneCountInString(data)) //prints: 2
}

```

###### Отсутствует запятая в многолинейных фрагментах Slice, Array и Map

*   уровень: начинающий

Сбой:

```
package main

func main() {
    x := []int{
    1,
    2 //error
    }
    _ = x
}

```

Ошибки компиляции:

> /tmp/sandbox367520156/main.go:6: синтаксическая ошибка: нужна запятая перед новой строкой в ​​составном литерале /tmp/sandbox367520156/main.go:8: оператор без объявления вне тела функции /tmp/sandbox367520156/main.go:9 : синтаксическая ошибка, неожиданно }

Работает:

```
package main

func main() {
    x := []int{
    1,
    2,
    }
    x = x

    y := []int{3,4,} //no error
    y = y
}

```

Вы не получите ошибку компилятора, если оставите завершающую запятую, когда свернете объявление в одну строку.

###### log.Fatal и log.Panic Do Больше, чем Log

*   уровень: начинающий

Библиотеки журналов часто предоставляют разные уровни журналов. В отличие от этих библиотек логов, пакет логов в Go делает больше, чем логи, если вы вызываете его `Fatal*()` и `Panic*()` функции. Когда ваше приложение вызывает эти функции, Go также завершает работу вашего приложения :\-)

```
package main

import "log"

func main() {
    log.Fatalln("Fatal Level: log entry") //app exits here
    log.Println("Normal Level: log entry")
}

```

###### Операции со встроенной структурой данных не синхронизированы

*   уровень: начинающий

Несмотря на то, что Go имеет ряд функций для поддержки параллелизма изначально, безопасные для параллелизма сборы данных \- это не одно из них :\-) Вы несете ответственность за то, чтобы обновления сбора данных были атомарными. Рекомендованные способы реализации этих элементарных операций \- распорядки и каналы, но вы также можете использовать пакет «sync», если это имеет смысл для вашего приложения.

###### Значения итерации для строк в предложениях «range»

*   уровень: начинающий

Значение индекса (первое значение, возвращаемое операцией «диапазон») является индексом первого байта для текущего «символа» (кодовая точка / руна Юникода), возвращенного во втором значении. Это не индекс для текущего «символа», как это делается в других языках. Обратите внимание, что действительный символ может быть представлен несколькими рунами. Обязательно ознакомьтесь с пакетом «norm» (golang.org/x/text/unicode/norm), если вам нужно работать с символами.

Предложения `for range` со строковыми переменными будут пытаться интерпретировать данные как текст UTF8. Для любых байтовых последовательностей, которые он не понимает, он вернет 0xfffd руны (или символы замены юникода) вместо реальных данных. Если в строковых переменных хранятся произвольные (не текстовые UTF8) данные, убедитесь, что они преобразованы в байтовые фрагменты, чтобы получить все сохраненные данные как есть.

```
package main

import "fmt"

func main() {
    data := "A\xfe\x02\xff\x04"
    for _,v := range data {
        fmt.Printf("%#x ",v)
    }
    //prints: 0x41 0xfffd 0x2 0xfffd 0x4 (not ok)

    fmt.Println()
    for _,v := range []byte(data) {
        fmt.Printf("%#x ",v)
    }
    //prints: 0x41 0xfe 0x2 0xff 0x4 (good)
}

```

###### Итерация по карте с помощью предложения "for range"

*   уровень: начинающий

Это ошибка, если вы ожидаете, что элементы будут в определенном порядке (например, упорядочены по значению ключа). Каждая итерация карты будет давать разные результаты. Среда выполнения Go пытается пройти лишнюю милю, рандомизируя порядок итераций, но это не всегда удается, поэтому вы можете получить несколько идентичных итераций карты. Не удивляйтесь, увидев 5 одинаковых итераций подряд.

```
package main

import "fmt"

func main() {
    m := map[string]int{"one":1,"two":2,"three":3,"four":4}
    for k,v := range m {
        fmt.Println(k,v)
    }
}

```

И если вы используете Go Playground ( [https://play.golang.org/](https://play.golang.org/) ), вы всегда получите те же результаты, потому что он не перекомпилирует код, если вы не внесете изменения.

###### Поведенческое поведение в «переключателях»

*   уровень: начинающий

Блоки case в выражениях switch выключаются по умолчанию. Это отличается от других языков, где поведение по умолчанию заключается в следующем блоке case.

```
package main

import "fmt"

func main() {
    isSpace := func(ch byte) bool {
        switch(ch) {
        case ' ': //error
        case '\t':
            return true
        }
        return false
    }

    fmt.Println(isSpace('\t')) //prints true (ok)
    fmt.Println(isSpace(' '))  //prints false (not ok)
}

```

Вы можете заставить блоки «case» проваливаться, используя инструкцию «fallthrough» в конце каждого блока «case». Вы также можете переписать оператор switch, чтобы использовать списки выражений в блоках case.

```
package main

import "fmt"

func main() {
    isSpace := func(ch byte) bool {
        switch(ch) {
        case ' ', '\t':
            return true
        }
        return false
    }

    fmt.Println(isSpace('\t')) //prints true (ok)
    fmt.Println(isSpace(' '))  //prints true (ok)
}

```

###### Увеличение и уменьшение

*   уровень: начинающий

Многие языки имеют операторы увеличения и уменьшения. В отличие от других языков, Go не поддерживает префиксную версию операций. Вы также не можете использовать эти два оператора в выражениях.

Сбой:

```
package main

import "fmt"

func main() {
    data := []int{1,2,3}
    i := 0
    ++i //error
    fmt.Println(data[i++]) //error
}

```

Ошибки компиляции:

> /tmp/sandbox101231828/main.go:8: синтаксическая ошибка: неожиданно ++ /tmp/sandbox101231828/main.go:9: синтаксическая ошибка: неожиданно ++, ожидается:

Работает:

```
package main

import "fmt"

func main() {
    data := []int{1,2,3}
    i := 0
    i++
    fmt.Println(data[i])
}

```

###### Побитовый оператор НЕ

*   уровень: начинающий

Многие языки используют `~` в качестве унарного оператора NOT (иначе говоря, побитового дополнения), но Go использует для этого оператор XOR ( `^` ).

Сбой:

```
package main

import "fmt"

func main() {
    fmt.Println(~2) //error
}

```

Ошибка компиляции:

> /tmp/sandbox965529189/main.go:6: побитовый оператор дополнения равен ^

Работает:

```
package main

import "fmt"

func main() {
    var d uint8 = 2
    fmt.Printf("%08b\n",^d)
}

```

Go по\-прежнему использует `^` в качестве оператора XOR, что может вводить в заблуждение некоторых людей.

If you want you can represent a unary NOT operation (e.g, `NOT 0x02`) with a binary XOR operation (e.g., `0x02 XOR 0xff`). This could explain why `^` is reused to represent unary NOT operations.

Go also has a special 'AND NOT' bitwise operator (`&^`), which adds to the NOT operator confusion. It looks like a special feature/hack to support `A AND (NOT B)` without requiring parentheses.

```
package main

import "fmt"

func main() {
    var a uint8 = 0x82
    var b uint8 = 0x02
    fmt.Printf("%08b [A]\n",a)
    fmt.Printf("%08b [B]\n",b)

    fmt.Printf("%08b (NOT B)\n",^b)
    fmt.Printf("%08b ^ %08b = %08b [B XOR 0xff]\n",b,0xff,b ^ 0xff)

    fmt.Printf("%08b ^ %08b = %08b [A XOR B]\n",a,b,a ^ b)
    fmt.Printf("%08b & %08b = %08b [A AND B]\n",a,b,a & b)
    fmt.Printf("%08b &^%08b = %08b [A 'AND NOT' B]\n",a,b,a &^ b)
    fmt.Printf("%08b&(^%08b)= %08b [A AND (NOT B)]\n",a,b,a & (^b))
}

```

###### Operator Precedence Differences

*   level: beginner

Aside from the "bit clear" operators (`&^`) Go has a set of standard operators shared by many other languages. The operator precedence is not always the same though.

```
package main

import "fmt"

func main() {
    fmt.Printf("0x2 & 0x2 + 0x4 -> %#x\n",0x2 & 0x2 + 0x4)
    //prints: 0x2 & 0x2 + 0x4 -> 0x6
    //Go:    (0x2 & 0x2) + 0x4
    //C++:    0x2 & (0x2 + 0x4) -> 0x2

    fmt.Printf("0x2 + 0x2 << 0x1 -> %#x\n",0x2 + 0x2 << 0x1)
    //prints: 0x2 + 0x2 << 0x1 -> 0x6
    //Go:     0x2 + (0x2 << 0x1)
    //C++:   (0x2 + 0x2) << 0x1 -> 0x8

    fmt.Printf("0xf | 0x2 ^ 0x2 -> %#x\n",0xf | 0x2 ^ 0x2)
    //prints: 0xf | 0x2 ^ 0x2 -> 0xd
    //Go:    (0xf | 0x2) ^ 0x2
    //C++:    0xf | (0x2 ^ 0x2) -> 0xf
}

```

###### Unexported Structure Fields Are Not Encoded

*   level: beginner

The struct fields starting with lowercase letters will not be (json, xml, gob, etc.) encoded, so when you decode the structure you'll end up with zero values in those unexported fields.

```
package main

import (
    "fmt"
    "encoding/json"
)

type MyData struct {
    One int
    two string
}

func main() {
    in := MyData{1,"two"}
    fmt.Printf("%#v\n",in) //prints main.MyData{One:1, two:"two"}

    encoded,_ := json.Marshal(in)
    fmt.Println(string(encoded)) //prints {"One":1}

    var out MyData
    json.Unmarshal(encoded,&out)

    fmt.Printf("%#v\n",out) //prints main.MyData{One:1, two:""}
}

```

###### App Exits With Active Goroutines

*   level: beginner

The app will not wait for all your goroutines to complete. This is a common mistake for beginners in general. Everybody starts somewhere, so there's no shame in making rookie mistakes :\-)

```
package main

import (
    "fmt"
    "time"
)

func main() {
    workerCount := 2

    for i := 0; i < workerCount; i++ {
        go doit(i)
    }
    time.Sleep(1 * time.Second)
    fmt.Println("all done!")
}

func doit(workerId int) {
    fmt.Printf("[%v] is running\n",workerId)
    time.Sleep(3 * time.Second)
    fmt.Printf("[%v] is done\n",workerId)
}

```

You'll see:

> \[0\] is running
> \[1\] is running
> all done!

One of the most common solutions is to use a "WaitGroup" variable. It will allow the main goroutine to wait until all worker goroutines are done. If your app has long running workers with message processing loops you'll also need a way to signal those goroutines that it's time to exit. You can send a "kill" message to each worker. Another option is to close a channel all workers are receiving from. It's a simple way to signal all goroutines at once.

```
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    done := make(chan struct{})
    workerCount := 2

    for i := 0; i < workerCount; i++ {
        wg.Add(1)
        go doit(i,done,wg)
    }

    close(done)
    wg.Wait()
    fmt.Println("all done!")
}

func doit(workerId int,done <-chan struct{},wg sync.WaitGroup) {
    fmt.Printf("[%v] is running\n",workerId)
    defer wg.Done()
    <- done
    fmt.Printf("[%v] is done\n",workerId)
}

```

If you run this app you'll see:

> \[0\] is running
> \[0\] is done
> \[1\] is running
> \[1\] is done

Looks like the workers are done before the main goroutine exits. Great! However, you'll also see this:

> fatal error: all goroutines are asleep \- deadlock!

That's not so great :\-) What's going on? Why is there a deadlock? The workers exited and they executed `wg.Done()`. The app should work.

The deadlock happens because each worker gets a copy of the original "WaitGroup" variable. When workers execute `wg.Done()` it has no effect on the "WaitGroup" variable in the main goroutine.

```
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    done := make(chan struct{})
    wq := make(chan interface{})
    workerCount := 2

    for i := 0; i < workerCount; i++ {
        wg.Add(1)
        go doit(i,wq,done,&wg)
    }

    for i := 0; i < workerCount; i++ {
        wq <- i
    }

    close(done)
    wg.Wait()
    fmt.Println("all done!")
}

func doit(workerId int, wq <-chan interface{},done <-chan struct{},wg *sync.WaitGroup) {
    fmt.Printf("[%v] is running\n",workerId)
    defer wg.Done()
    for {
        select {
        case m := <- wq:
            fmt.Printf("[%v] m => %v\n",workerId,m)
        case <- done:
            fmt.Printf("[%v] is done\n",workerId)
            return
        }
    }
}

```

Now it works as expected :\-)

###### Sending to an Unbuffered Channel Returns As Soon As the Target Receiver Is Ready

*   level: beginner

The sender will not be blocked until your message is processed by the receiver. Depending on the machine where you are running the code, the receiver goroutine may or may not have enough time to process the message before the sender continues its execution.

```
package main

import "fmt"

func main() {
    ch := make(chan string)

    go func() {
        for m := range ch {
            fmt.Println("processed:",m)
        }
    }()

    ch <- "cmd.1"
    ch <- "cmd.2" //won't be processed
}

```

###### Sending to an Closed Channel Causes a Panic

*   level: beginner

Receiving from a closed channel is safe. The `ok` return value in a receive statement will be set to `false` indicating that no data was received. If you are receiving from a buffered channel you'll get the buffered data first and once it's empty the `ok` return value will be `false`.

Sending data to a closed channel causes a panic. It is a documented behavior, but it's not very intuitive for new Go developers who might expect the send behavior to be similar to the receive behavior.

```
package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan int)
    for i := 0; i < 3; i++ {
        go func(idx int) {
            ch <- (idx + 1) * 2
        }(i)
    }

    //get the first result
    fmt.Println(<-ch)
    close(ch) //not ok (you still have other senders)
    //do other work
    time.Sleep(2 * time.Second)
}

```

Depending on your application the fix will be different. It might be a minor code change or it might require a change in your application design. Either way, you'll need to make sure your application doesn't try to send data to a closed channel.

The buggy example can be fixed by using a special cancellation channel to signal the remaining workers that their results are no longer neeeded.

```
package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan int)
    done := make(chan struct{})
    for i := 0; i < 3; i++ {
        go func(idx int) {
            select {
            case ch <- (idx + 1) * 2: fmt.Println(idx,"sent result")
            case <- done: fmt.Println(idx,"exiting")
            }
        }(i)
    }

    //get first result
    fmt.Println("result:",<-ch)
    close(done)
    //do other work
    time.Sleep(3 * time.Second)
}

```

###### Using "nil" Channels

*   level: beginner

Send and receive operations on a `nil` channel block forver. It's a well documented behavior, but it can be a surprise for new Go developers.

```
package main

import (
    "fmt"
    "time"
)

func main() {
    var ch chan int
    for i := 0; i < 3; i++ {
        go func(idx int) {
            ch <- (idx + 1) * 2
        }(i)
    }

    //get first result
    fmt.Println("result:",<-ch)
    //do other work
    time.Sleep(2 * time.Second)
}

```

If you run the code you'll see a runtime error like this: `fatal error: all goroutines are asleep - deadlock!`

This behavior can be used as a way to dynamically enable and disable `case` blocks in a `select` statement.

```
package main

import "fmt"
import "time"

func main() {
    inch := make(chan int)
    outch := make(chan int)

    go func() {
        var in <- chan int = inch
        var out chan <- int
        var val int
        for {
            select {
            case out <- val:
                out = nil
                in = inch
            case val = <- in:
                out = outch
                in = nil
            }
        }
    }()

    go func() {
        for r := range outch {
            fmt.Println("result:",r)
        }
    }()

    time.Sleep(0)
    inch <- 1
    inch <- 2
    time.Sleep(3 * time.Second)
}

```

###### Methods with Value Receivers Can't Change the Original Value

*   level: beginner

Method receivers are like regular function arguments. If it's declared to be a value then your function/method gets a copy of your receiver argument. This means making changes to the receiver will not affect the original value unless your receiver is a map or slice variable and you are updating the items in the collection or the fields you are updating in the receiver are pointers.

```
package main

import "fmt"

type data struct {
    num int
    key *string
    items map[string]bool
}

func (this *data) pmethod() {
    this.num = 7
}

func (this data) vmethod() {
    this.num = 8
    *this.key = "v.key"
    this.items["vmethod"] = true
}

func main() {
    key := "key.1"
    d := data{1,&key,make(map[string]bool)}

    fmt.Printf("num=%v key=%v items=%v\n",d.num,*d.key,d.items)
    //prints num=1 key=key.1 items=map[]

    d.pmethod()
    fmt.Printf("num=%v key=%v items=%v\n",d.num,*d.key,d.items)
    //prints num=7 key=key.1 items=map[]

    d.vmethod()
    fmt.Printf("num=%v key=%v items=%v\n",d.num,*d.key,d.items)
    //prints num=7 key=v.key items=map[vmethod:true]
}

```

###### Closing HTTP Response Body

*   level: intermediate

When you make requests using the standard http library you get a http response variable. If you don't read the response body you still need to close it. Note that you must do it for empty responses too. It's very easy to forget especially for new Go developers.

Some new Go developers do try to close the response body, but they do it in the wrong place.

```
package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
    resp, err := http.Get("https://api.ipify.org?format=json")
    defer resp.Body.Close()//not ok
    if err != nil {
        fmt.Println(err)
        return
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(body))
}

```

Этот код работает для успешных запросов, но в случае сбоя http\-запроса `resp` переменная может быть `nil` , что вызовет панику во время выполнения.

Наиболее распространенная причина закрытия тела ответа \- использование `defer` вызова после проверки ошибок HTTP\-ответа.

```
package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
    resp, err := http.Get("https://api.ipify.org?format=json")
    if err != nil {
        fmt.Println(err)
        return
    }

    defer resp.Body.Close()//ok, most of the time :-)
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(body))
}

```

В большинстве случаев, когда ваш http\-запрос завершается неудачей, `resp` переменная будет `nil` и `err` переменная будет `non-nil` . Однако, когда вы получаете ошибку перенаправления, обе переменные будут `non-nil` . Это означает, что вы все равно можете получить утечку.

Эту утечку можно устранить, добавив вызов для закрытия `non-nil` тел ответов в блоке обработки ошибок http. Другой вариант \- использовать один `defer` вызов для закрытия тел ответа для всех неудачных и успешных запросов.

```
package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
    resp, err := http.Get("https://api.ipify.org?format=json")
    if resp != nil {
        defer resp.Body.Close()
    }

    if err != nil {
        fmt.Println(err)
        return
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(body))
}

```

Оригинальная реализация для `resp.Body.Close()` также считывает и отбрасывает оставшиеся данные тела ответа. Это гарантировало возможность повторного использования http\-соединения для другого запроса, если включено поведение keepalive http\-соединения. Последнее поведение http\-клиента отличается. Теперь вы несете ответственность за чтение и удаление оставшихся данных ответов. Если вы этого не сделаете, http\-соединение может быть закрыто вместо повторного использования. Эта маленькая ошибка должна быть документирована в Go 1.5.

Если повторное использование http\-соединения важно для вашего приложения, вам может потребоваться добавить что\-то вроде этого в конце логики обработки ответа:

```
_, err = io.Copy(ioutil.Discard, resp.Body)

```

Это будет необходимо, если вы сразу не прочитаете все тело ответа, что может произойти, если вы обрабатываете ответы API json с помощью следующего кода:

```
json.NewDecoder(resp.Body).Decode(&data)

```

###### Закрытие HTTP\-соединений

*   level: intermediate

Some HTTP servers keep network connections open for a while (based on the HTTP 1.1 spec and the server "keep\-alive" configurations). By default, the standard http library will close the network connections only when the target HTTP server asks for it. This means your app may run out of sockets/file descriptors under certain conditions.

You can ask the http library to close the connection after your request is done by setting the `Close` field in the request variable to `true`.

Another option is to add a `Connection` request header and set it to `close`. The target HTTP server should respond with a `Connection: close` header too. When the http library sees this response header it will also close the connection.

```
package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
    req, err := http.NewRequest("GET","http://golang.org",nil)
    if err != nil {
        fmt.Println(err)
        return
    }

    req.Close = true
    //or do this:
    //req.Header.Add("Connection", "close")

    resp, err := http.DefaultClient.Do(req)
    if resp != nil {
        defer resp.Body.Close()
    }

    if err != nil {
        fmt.Println(err)
        return
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(len(string(body)))
}

```

You can also disable http connection reuse globally. You'll need to create a custom http transport configuration for it.

```
package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
    tr := &http.Transport{DisableKeepAlives: true}
    client := &http.Client{Transport: tr}

    resp, err := client.Get("http://golang.org")
    if resp != nil {
        defer resp.Body.Close()
    }

    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(resp.StatusCode)

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(len(string(body)))
}

```

If you send a lot of requests to the same HTTP server it's ok to keep the network connection open. However, if your app sends one or two requests to many different HTTP servers in a short period of time it's a good idea to close the network connections right after your app receives the responses. Increasing the open file limit might be a good idea too. The correct solution depends on your application though.

###### JSON Encoder Adds a Newline Character

*   level: intermediate

You are writing a test for your JSON encoding function when you discover that your tests fail because you are not getting the expected value. What happened? If you are using the JSON Encoder object then you'll get an extra newline character at the end of your encoded JSON object.

```
package main

import (
  "fmt"
  "encoding/json"
  "bytes"
)

func main() {
  data := map[string]int{"key": 1}

  var b bytes.Buffer
  json.NewEncoder(&b).Encode(data)

  raw,_ := json.Marshal(data)

  if b.String() == string(raw) {
    fmt.Println("same encoded data")
  } else {
    fmt.Printf("'%s' != '%s'\n",raw,b.String())
    //prints:
    //'{"key":1}' != '{"key":1}\n'
  }
}

```

The JSON Encoder object is designed for streaming. Streaming with JSON usually means newline delimited JSON objects and this is why the Encode method adds a newline character. This is a documented behavior, but it's commonly overlooked or forgotten.

###### JSON Package Escapes Special HTML Characters in Keys and String Values

*   level: intermediate

This is a documented behavior, but you have to be careful reading all of the JSON package documentation to learn about it. The `SetEscapeHTML` method description talks about the default encoding behavior for the and, less than and greater than characters.

This is a very unfortunate design decision by the Go team for a number of reasons. First, you can't disable this behavior for the `json.Marshal` calls. Second, this is a badly implemented security feature because it assumes that doing HTML encoding is sufficient to protect against XSS vulnerabilities in all web applications. There are a lot of different contexts where the data can be used and each context requires its own encoding method. And finally, it's bad because it assumes that the primary use case for JSON is a web page, which breaks the configuration libraries and the REST/HTTP APIs by default.

```
package main

import (
  "fmt"
  "encoding/json"
  "bytes"
)

func main() {
  data := "x < y"

  raw,_ := json.Marshal(data)
  fmt.Println(string(raw))
  //prints: "x \u003c y" <- probably not what you expected

  var b1 bytes.Buffer
  json.NewEncoder(&b1).Encode(data)
  fmt.Println(b1.String())
  //prints: "x \u003c y" <- probably not what you expected

  var b2 bytes.Buffer
  enc := json.NewEncoder(&b2)
  enc.SetEscapeHTML(false)
  enc.Encode(data)
  fmt.Println(b2.String())
  //prints: "x < y" <- looks better
}

```

A suggestion to the Go team... Make it an opt\-in.

###### Unmarshalling JSON Numbers into Interface Values

*   level: intermediate

By default, Go treats numeric values in JSON as `float64` numbers when you decode/unmarshal JSON data into an interface. This means the following code will fail with a panic:

```
package main

import (
  "encoding/json"
  "fmt"
)

func main() {
  var data = []byte(`{"status": 200}`)

  var result map[string]interface{}
  if err := json.Unmarshal(data, &result); err != nil {
    fmt.Println("error:", err)
    return
  }

  var status = result["status"].(int) //error
  fmt.Println("status value:",status)
}

```

Runtime Panic:

> panic: interface conversion: interface is float64, not int

If the JSON value you are trying to decode is an integer you have serveral options.

Option one: use the float value as\-is :\-)

Option two: convert the float value to the integer type you need.

```
package main

import (
  "encoding/json "
  "fmt "
)

func main() {
  var data = []byte(`{"status ": 200}`)

  var result map[string]interface{}
  if err := json.Unmarshal(data, &result); err != nil {
    fmt.Println("error: ", err)
    return
  }

  var status = uint64(result["status "].(float64)) //ok
  fmt.Println("status value: ",status)
}

```

Option three: use a `Decoder` type to unmarshal JSON and tell it to represent JSON numbers using the `Number` interface type.

```
package main

import (
  "encoding/json "
  "bytes "
  "fmt "
)

func main() {
  var data = []byte(`{"status ": 200}`)

  var result map[string]interface{}
  var decoder = json.NewDecoder(bytes.NewReader(data))
  decoder.UseNumber()

  if err := decoder.Decode(&result); err != nil {
    fmt.Println("error: ", err)
    return
  }

  var status,_ = result["status "].(json.Number).Int64() //ok
  fmt.Println("status value: ",status)
}

```

You can use the string representation of your `Number` value to unmarshal it to a different numeric type:

```
package main

import (
  "encoding/json "
  "bytes "
  "fmt "
)

func main() {
  var data = []byte(`{"status ": 200}`)

  var result map[string]interface{}
  var decoder = json.NewDecoder(bytes.NewReader(data))
  decoder.UseNumber()

  if err := decoder.Decode(&result); err != nil {
    fmt.Println("error: ", err)
    return
  }

  var status uint64
  if err := json.Unmarshal([]byte(result["status "].(json.Number).String()), &status); err != nil {
    fmt.Println("error: ", err)
    return
  }

  fmt.Println("status value: ",status)
}

```

Option four: use a `struct` type that maps your numeric value to the numeric type you need.

```
package main

import (
  "encoding/json "
  "bytes "
  "fmt "
)

func main() {
  var data = []byte(`{"status ": 200}`)

  var result struct {
    Status uint64 `json:"status "`
  }

  if err := json.NewDecoder(bytes.NewReader(data)).Decode(&result); err != nil {
    fmt.Println("error: ", err)
    return
  }

  fmt.Printf("result=> %+v ",result)
  //prints: result => {Status:200}
}

```

Option five: use a `struct` that maps your numeric value to the `json.RawMessage` type if you need to defer the value decoding.

This option is useful if you have to perform conditional JSON field decoding where the field type or structure might change.

```
package main

import (
  "encoding/json "
  "bytes "
  "fmt "
)

func main() {
  records := [][]byte{
    []byte(`{"status ": 200, "tag ":"one "}`),
    []byte(`{"status ":"ok ", "tag ":"two "}`),
  }

  for idx, record := range records {
    var result struct {
      StatusCode uint64
      StatusName string
      Status json.RawMessage `json:"status "`
      Tag string             `json:"tag "`
    }

    if err := json.NewDecoder(bytes.NewReader(record)).Decode(&result); err != nil {
      fmt.Println("error: ", err)
      return
    }

    var sstatus string
    if err := json.Unmarshal(result.Status, &sstatus); err == nil {
      result.StatusName = sstatus
    }

    var nstatus uint64
    if err := json.Unmarshal(result.Status, &nstatus); err == nil {
      result.StatusCode = nstatus
    }

    fmt.Printf("[%v] result=> %+v\n ",idx,result)
  }
}

```

###### JSON String Values Will Not Be Ok with Hex or Other non\-UTF8 Escape Sequences

*   level: intermediate

Go expects string values to be UTF8 encoded. This means you can't have arbitrary hex escaped binary data in your JSON strings (and you also have to escape the backslash character). This is really a JSON gotcha Go inherited, but it happens often enough in Go apps that it makes sense to mention it anyways.

```
package main

import (
  "fmt "
  "encoding/json "
)

type config struct {
  Data string `json:"data "`
}

func main() {
  raw := []byte(`{"data ":"\xc2 "}`)
  var decoded config

  if err := json.Unmarshal(raw, &decoded); err != nil {
        fmt.Println(err)
    //prints: invalid character 'x' in string escape code
    }

}

```

Вызовы Unmarshal / Decode потерпят неудачу, если Go увидит шестнадцатеричную escape\-последовательность. Если вам нужно иметь обратную косую черту в вашей строке, убедитесь, что вы избежали ее с другой обратной реакцией. Если вы хотите использовать двоичные данные в шестнадцатеричном формате, вы можете избежать обратной косой черты, а затем выполнить собственное шестнадцатеричное экранирование с декодированными данными в вашей строке JSON.

```
package main

import (
  "fmt "
  "encoding/json "
)

type config struct {
  Data string `json:"data "`
}

func main() {
  raw := []byte(`{"data ":"\\xc2 "}`)

  var decoded config

  json.Unmarshal(raw, &decoded)

  fmt.Printf("%#v ",decoded) //prints: main.config{Data:"\\xc2 "}
  //todo: do your own hex escape decoding for decoded.Data
}

```

Другой вариант \- использовать тип данных байтового массива / слайса в вашем объекте JSON, но двоичные данные должны быть в кодировке base64.

```
package main

import (
  "fmt "
  "encoding/json "
)

type config struct {
  Data []byte `json:"data "`
}

func main() {
  raw := []byte(`{"data ":"wg=="}`)
  var decoded config

  if err := json.Unmarshal(raw, &decoded); err != nil {
          fmt.Println(err)
      }

  fmt.Printf("%#v ",decoded) //prints: main.config{Data:[]uint8{0xc2}}
}

```

Остается остерегаться символа замены Unicode (U + FFFD). Go будет использовать символ замены вместо недопустимого UTF8, поэтому вызов Unmarshal / Decode не завершится неудачно, но полученное строковое значение может не соответствовать ожидаемому.

###### Сравнение структур, массивов, фрагментов и карт

*   уровень: средний

Вы можете использовать оператор равенства `==`, для сравнения структурных переменных, если каждое структурное поле можно сравнить с оператором равенства.

```
package main

import "fmt "

type data struct {
    num int
    fp float32
    complex complex64
    str string
    char rune
    yes bool
    events <-chan string
    handler interface{}
    ref *byte
    raw [10]byte
}

func main() {
    v1 := data{}
    v2 := data{}
    fmt.Println("v1==v 2: ",v1 == v2) //prints: v1 == v2: true
}

```

Если какое\-либо из полей структуры несопоставимо, то использование оператора равенства приведет к ошибкам времени компиляции. Обратите внимание, что массивы сравнимы, только если их элементы данных сопоставимы.

```
package main

import "fmt "

type data struct {
    num int                //ok
    checks [10]func() bool //not comparable
    doit func() bool       //not comparable
    m map[string] string   //not comparable
    bytes []byte           //not comparable
}

func main() {
    v1 := data{}
    v2 := data{}
    fmt.Println("v1==v 2: ",v1 == v2)
}

```

Go предоставляет ряд вспомогательных функций для сравнения переменных, которые нельзя сравнивать с помощью операторов сравнения.

Наиболее распространенным решением является использование `DeepEqual()`функции в отражающем пакете.

```
package main

import (
    "fmt "
    "reflect "
)

type data struct {
    num int                //ok
    checks [10]func() bool //not comparable
    doit func() bool       //not comparable
    m map[string] string   //not comparable
    bytes []byte           //not comparable
}

func main() {
    v1 := data{}
    v2 := data{}
    fmt.Println("v1==v 2: ",reflect.DeepEqual(v1,v2)) //prints: v1 == v2: true

    m1 := map[string]string{"one ": "a ","two
                ": "b "}
    m2 := map[string]string{"two ": "b ", "one
                ": "a "}
    fmt.Println("m1==m 2: ",reflect.DeepEqual(m1, m2)) //prints: m1 == m2: true

    s1 := []int{1, 2, 3}
    s2 := []int{1, 2, 3}
    fmt.Println("s1==s 2: ",reflect.DeepEqual(s1, s2)) //prints: s1 == s2: true
}

```

Помимо того, что он медленный (который может или не может быть прерывателем сделки для вашего приложения), он `DeepEqual()`также имеет свои собственные ошибки.

```
package main

import (
    "fmt "
    "reflect "
)

func main() {
    var b1 []byte = nil
    b2 := []byte{}
    fmt.Println("b1==b 2: ",reflect.DeepEqual(b1, b2)) //prints: b1 == b2: false
}

```

`DeepEqual()`не считает пустой срез равным «нулевому» срезу. Это поведение отличается от поведения, которое вы получаете, используя `bytes.Equal()`функцию. `bytes.Equal()`считает "ноль " и пустые ломтики равными.

```
package main

import (
    "fmt "
    "bytes "
)

func main() {
    var b1 []byte = nil
    b2 := []byte{}
    fmt.Println("b1==b 2: ",bytes.Equal(b1, b2)) //prints: b1 == b2: true
}

```

`DeepEqual()` не всегда идеально сравнивать ломтики

```
package main

import (
    "fmt "
    "reflect "
    "encoding/json "
)

func main() {
    var str string = "one "
    var in interface{} = "one "
    fmt.Println("str==i n: ",str == in,reflect.DeepEqual(str, in))
    //prints: str == in: true true

    v1 := []string{"one ","two "}
    v2 := []interface{}{"one ","two "}
    fmt.Println("v1==v 2: ",reflect.DeepEqual(v1, v2))
    //prints: v1 == v2: false (not ok)

    data := map[string]interface{}{
        "code ": 200,
        "value ": []string{"one ","two "},
    }
    encoded, _ := json.Marshal(data)
    var decoded map[string]interface{}
    json.Unmarshal(encoded, &decoded)
    fmt.Println("data==d ecoded: ",reflect.DeepEqual(data, decoded))
    //prints: data == decoded: false (not ok)
}

```

Если байты ломтик (или строка) содержат текстовые данные , которые вы могли бы возникнуть соблазн использовать `ToUpper()`или `ToLower()`от «байт» и «строки» пакетов , когда нужно сравнить значения в случае нечувствительным образом (перед использованием `==`, `bytes.Equal()`или `bytes.Compare()`). Он будет работать для английского текста, но не будет работать для текста на многих других языках. `strings.EqualFold()`и `bytes.EqualFold()`должен использоваться вместо

Если байты ломтики содержат секреты (например, криптографические хэш, маркеры и т.д.) , которые должны быть проверены в отношении пользовательских данных, не следует использовать `reflect.DeepEqual()`, `bytes.Equal()`или `bytes.Compare()`потому , что эти функции сделают ваше приложение уязвимой для [**атак синхронизации**](http://en.wikipedia.org/wiki/Timing_attack ) . Чтобы избежать утечки информации о времени, используйте функции из пакета «crypto / subtle» (например, `subtle.ConstantTimeCompare()`).

###### Восстановление от паники

*   уровень: средний

Эта `recover()`функция может быть использована, чтобы поймать / перехватить панику. Вызов `recover()`будет работать только тогда, когда он выполняется в отложенной функции.

Неправильно:

```
package main

import "fmt "

func main() {
    recover() //doesn't do anything
    panic("not good ")
    recover() //won't be executed :)
    fmt.Println("ok ")
}

```

Работает:

```
package main

import "fmt "

func main() {
    defer func() {
        fmt.Println("recovered: ",recover())
    }()

    panic("not good ")
}

```

Вызов to `recover()`работает только в том случае, если он вызывается непосредственно в вашей отложенной функции.

Сбой:

```
package main

import "fmt "

func doRecover() {
    fmt.Println("recovered=> ",recover()) //prints: recovered => <nil>
}

func main() {
    defer func() {
        doRecover() //panic is not recovered
    }()

    panic("not good ")
}

```

###### Обновление и ссылка на значения элементов в разделах «Срез», «Массив» и «Карта»

*   уровень: средний

The data values generated in the "range " clause are copies of the actual collection elements. They are not references to the original items. This means that updating the values will not change the original data. It also means that taking the address of the values will not give you pointers to the original data.

```
package main

import "fmt "

func main() {
    data := []int{1,2,3}
    for _,v := range data {
        v *= 10 //original item is not changed
    }

    fmt.Println("data: ",data) //prints data: [1 2 3]
}

```

If you need to update the original collection record value use the index operator to access the data.

```
package main

import "fmt "

func main() {
    data := []int{1,2,3}
    for i,_ := range data {
        data[i] *= 10
    }

    fmt.Println("data: ",data) //prints data: [10 20 30]
}

```

If your collection holds pointer values then the rules are slightly different. You still need to use the index operator if you want the original record to point to another value, but you can update the data stored at the target location using the second value in the "for range " clause.

```
package main

import "fmt "

func main() {
    data := []*struct{num int} {{1},{2},{3}}

    for _,v := range data {
        v.num *= 10
    }

    fmt.Println(data[0],data[1],data[2]) //prints &{10} &{20} &{30}
}

```

###### "Hidden " Data in Slices

*   level: intermediate

When you reslice a slice, the new slice will reference the array of the original slice. If you forget about this behavior it can lead to unexpected memory usage if your application allocates large temporary slices creating new slices from them to refer to small sections of the original data.

```
package main

import "fmt "

func get() []byte {
    raw := make([]byte,10000)
    fmt.Println(len(raw),cap(raw),&raw[0]) //prints: 10000 10000 <byte_addr_x>
    return raw[:3]
}

func main() {
    data := get()
    fmt.Println(len(data),cap(data),&data[0]) //prints: 3 10000 <byte_addr_x>
}

```

To avoid this trap make sure to copy the data you need from the temporary slice (instead of reslicing it).

```
package main

import "fmt "

func get() []byte {
    raw := make([]byte,10000)
    fmt.Println(len(raw),cap(raw),&raw[0]) //prints: 10000 10000 <byte_addr_x>
    res := make([]byte,3)
    copy(res,raw[:3])
    return res
}

func main() {
    data := get()
    fmt.Println(len(data),cap(data),&data[0]) //prints: 3 3 <byte_addr_y>
}

```

###### Slice Data "Corruption "

*   level: intermediate

Let's say you need to rewrite a path (stored in a slice). You reslice the path to reference each directory modifying the first folder name and then you combine the names to create a new path.

```
package main

import (
    "fmt "
    "bytes "
)

func main() {
    path := []byte("AAAA/BBBBBBBBB ")
    sepIndex := bytes.IndexByte(path,'/')
    dir1 := path[:sepIndex]
    dir2 := path[sepIndex+1:]
    fmt.Println("dir1=> ",string(dir1)) //prints: dir1 => AAAA
    fmt.Println("dir2=> ",string(dir2)) //prints: dir2 => BBBBBBBBB

    dir1 = append(dir1,"suffix "...)
    path = bytes.Join([][]byte{dir1,dir2},[]byte{'/'})

    fmt.Println("dir1=> ",string(dir1)) //prints: dir1 => AAAAsuffix
    fmt.Println("dir2=> ",string(dir2)) //prints: dir2 => uffixBBBB (not ok)

    fmt.Println("new path=> ",string(path))
}

```

It didn't work as you expected. Instead of "AAAAsuffix/BBBBBBBBB " you ended up with "AAAAsuffix/uffixBBBB ". It happened because both directory slices referenced the same underlying array data from the original path slice. This means that the original path is also modified. Depending on your application this might be a problem too.

This problem can fixed by allocating new slices and copying the data you need. Another option is to use the full slice expression.

```
package main

import (
    "fmt "
    "bytes "
)

func main() {
    path := []byte("AAAA/BBBBBBBBB ")
    sepIndex := bytes.IndexByte(path,'/')
    dir1 := path[:sepIndex:sepIndex] //full slice expression
    dir2 := path[sepIndex+1:]
    fmt.Println("dir1=> ",string(dir1)) //prints: dir1 => AAAA
    fmt.Println("dir2=> ",string(dir2)) //prints: dir2 => BBBBBBBBB

    dir1 = append(dir1,"suffix "...)
    path = bytes.Join([][]byte{dir1,dir2},[]byte{'/'})

    fmt.Println("dir1=> ",string(dir1)) //prints: dir1 => AAAAsuffix
    fmt.Println("dir2=> ",string(dir2)) //prints: dir2 => BBBBBBBBB (ok now)

    fmt.Println("new path=> ",string(path))
}

```

The extra parameter in the full slice expression controls the capacity for the new slice. Now appending to that slice will trigger a new buffer allocation instead of overwriting the data in the second slice.

###### "Stale " Slices

*   level: intermediate

Multiple slices can reference the same data. This can happen when you create a new slice from an existing slice, for example. If your application relies on this behavior to function properly then you'll need to worry about "stale " slices.

At some point adding data to one of the slices will result in a new array allocation when the original array can't hold any more new data. Now other slices will point to the old array (with old data).

```
import "fmt "

func main() {
    s1 := []int{1,2,3}
    fmt.Println(len(s1),cap(s1),s1) //prints 3 3 [1 2 3]

    s2 := s1[1:]
    fmt.Println(len(s2),cap(s2),s2) //prints 2 2 [2 3]

    for i := range s2 { s2[i] += 20 }

    //still referencing the same array
    fmt.Println(s1) //prints [1 22 23]
    fmt.Println(s2) //prints [22 23]

    s2 = append(s2,4)

    for i := range s2 { s2[i] += 10 }

    //s1 is now "stale "
    fmt.Println(s1) //prints [1 22 23]
    fmt.Println(s2) //prints [32 33 14]
}

```

###### Type Declarations and Methods

*   level: intermediate

When you create a type declaration by defining a new type from an existing (non\-interface) type, you don't inherit the methods defined for that existing type.

Fails:

```
package main

import "sync "

type myMutex sync.Mutex

func main() {
    var mtx myMutex
    mtx.Lock() //error
    mtx.Unlock() //error
}

```

Compile Errors:

> /tmp/sandbox106401185/main.go:9: mtx.Lock undefined (type myMutex has no field or method Lock) /tmp/sandbox106401185/main.go:10: mtx.Unlock undefined (type myMutex has no field or method Unlock)

If you do need the methods from the original type you can define a new struct type embedding the original type as an anonymous field.

Works:

```
package main

import "sync "

type myLocker struct {
    sync.Mutex
}

func main() {
    var lock myLocker
    lock.Lock() //ok
    lock.Unlock() //ok
}

```

Interface type declarations also retain their method sets.

Works:

```
package main

import "sync "

type myLocker sync.Locker

func main() {
    var lock myLocker = new(sync.Mutex)
    lock.Lock() //ok
    lock.Unlock() //ok
}

```

###### Breaking Out of "for switch " and "for select " Code Blocks

*   level: intermediate

A "break " statement without a label only gets you out of the inner switch/select block. If using a "return " statement is not an option then defining a label for the outer loop is the next best thing.

```
package main

import "fmt "

func main() {
    loop:
        for {
            switch {
            case true:
                fmt.Println("breaking out... ")
                break loop
            }
        }

    fmt.Println("out! ")
}

```

A "goto " statement will do the trick too...

###### Iteration Variables and Closures in "for " Statements

*   level: intermediate

Это самая распространенная ошибка в Go. Переменные итерации в `for`операторах повторно используются в каждой итерации. Это означает, что каждое замыкание (также называемое литералом функции), созданное в вашем `for`цикле, будет ссылаться на одну и ту же переменную (и они получат значение этой переменной в тот момент, когда эти программы будут выполняться).

Неправильно:

```
package main

import (
    "fmt "
    "time "
)

func main() {
    data := []string{"one ","two ","three "}

    for _,v := range data {
        go func() {
            fmt.Println(v)
        }()
    }

    time.Sleep(3 * time.Second)
    //goroutines print: three, three, three
}

```

Самое простое решение (которое не требует каких\-либо изменений в процедуре) \- сохранить текущее значение переменной итерации в локальной переменной внутри `for`блока цикла.

Работает:

```
package main

import (
    "fmt "
    "time "
)

func main() {
    data := []string{"one ","two ","three "}

    for _,v := range data {
        vcopy := v //
        go func() {
            fmt.Println(vcopy)
        }()
    }

    time.Sleep(3 * time.Second)
    //goroutines print: one, two, three
}

```

Другое решение состоит в том, чтобы передать текущую переменную итерации в качестве параметра анонимной программе.

Работает:

```
package main

import (
    "fmt "
    "time "
)

func main() {
    data := []string{"one ","two ","three "}

    for _,v := range data {
        go func(in string) {
            fmt.Println(in)
        }(v)
    }

    time.Sleep(3 * time.Second)
    //goroutines print: one, two, three
}

```

Вот немного более сложная версия ловушки.

Неправильно:

```
package main

import (
    "fmt "
    "time "
)

type field struct {
    name string
}

func (p *field) print() {
    fmt.Println(p.name)
}

func main() {
    data := []field{{"one "},{"two "},{"three "}}

    for _,v := range data {
        go v.print()
    }

    time.Sleep(3 * time.Second)
    //goroutines print: three, three, three
}

```

Работает:

```
package main

import (
    "fmt "
    "time "
)

type field struct {
    name string
}

func (p *field) print() {
    fmt.Println(p.name)
}

func main() {
    data := []field{{"one "},{"two "},{"three "}}

    for _,v := range data {
        v := v
        go v.print()
    }

    time.Sleep(3 * time.Second)
    //goroutines print: one, two, three
}

```

Как вы думаете, что вы увидите, когда запустите этот код (и почему)?

```
package main

import (
    "fmt "
    "time "
)

type field struct {
    name string
}

func (p *field) print() {
    fmt.Println(p.name)
}

func main() {
    data := []*field{{"one "},{"two "},{"three "}}

    for _,v := range data {
        go v.print()
    }

    time.Sleep(3 * time.Second)
}

```

###### Оценка аргумента вызова отложенной функции

*   уровень: средний

Arguments for a deferred function call are evaluated when the `defer` statement is evaluated (not when the function is actually executing). The same rules apply when you defer a method call. The structure value is also saved along with the explicit method parameters and the closed variables.

```
package main

import "fmt "

func main() {
    var i int = 1

    defer fmt.Println("result=> ",func() int { return i * 2 }())
    i++
    //prints: result => 2 (not ok if you expected 4)
}

```

If you have pointer parameters it is possible to change the values they point to because only the pointer is saved when the `defer` statement is evaluated.

```
package main

import (
  "fmt "
)

func main() {
  i := 1
  defer func (in *int) { fmt.Println("result=>
                ", *in) }(&i)

  i = 2
  //prints: result => 2
}

```

###### Deferred Function Call Execution

*   level: intermediate

Отложенные вызовы выполняются в конце содержащей функции (и в обратном порядке), а не в конце содержащего код блока. Для начинающих разработчиков Go легко ошибиться, путая правила выполнения отложенного кода с правилами области видимости переменных. Это может стать проблемой, если у вас есть долго выполняющаяся функция с `for`циклом, который пытается `defer`очищать вызовы ресурсов на каждой итерации.

```
package main

import (
    "fmt "
    "os "
    "path/filepath "
)

func main() {
    if len(os.Args) != 2 {
        os.Exit(-1)
    }

    start, err := os.Stat(os.Args[1])
    if err != nil || !start.IsDir(){
        os.Exit(-1)
    }

    var targets []string
    filepath.Walk(os.Args[1], func(fpath string, fi os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !fi.Mode().IsRegular() {
            return nil
        }

        targets = append(targets,fpath)
        return nil
    })

    for _,target := range targets {
        f, err := os.Open(target)
        if err != nil {
            fmt.Println("bad target: ",target,"error: ",err) //prints error: too many open files
            break
        }
        defer f.Close() //will not be closed at the end of this code block
        //do something with the file...
    }
}

```

Один из способов решить эту проблему \- заключить блок кода в функцию.

```
package main

import (
    "fmt "
    "os "
    "path/filepath "
)

func main() {
    if len(os.Args) != 2 {
        os.Exit(-1)
    }

    start, err := os.Stat(os.Args[1])
    if err != nil || !start.IsDir(){
        os.Exit(-1)
    }

    var targets []string
    filepath.Walk(os.Args[1], func(fpath string, fi os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !fi.Mode().IsRegular() {
            return nil
        }

        targets = append(targets,fpath)
        return nil
    })

    for _,target := range targets {
        func() {
            f, err := os.Open(target)
            if err != nil {
                fmt.Println("bad target: ",target,"error: ",err)
                return
            }
            defer f.Close() //ok
            //do something with the file...
        }()
    }
}

```

Другой вариант \- избавиться от `defer`высказывания :\-)

###### Неудачные утверждения типа

*   уровень: средний

Утверждения с ошибочными типами возвращают «нулевое значение» для целевого типа, используемого в операторе подтверждения. Это может привести к неожиданному поведению, когда оно смешано с изменением теней.

Неправильно:

```
package main

import "fmt "

func main() {
    var data interface{} = "great "

    if data, ok := data.(int); ok {
        fmt.Println("[is an int] value=> ",data)
    } else {
        fmt.Println("[not an int] value=> ",data)
        //prints: [not an int] value => 0 (not "great ")
    }
}

```

Работает:

```
package main

import "fmt "

func main() {
    var data interface{} = "great "

    if res, ok := data.(int); ok {
        fmt.Println("[is an int] value=> ",res)
    } else {
        fmt.Println("[not an int] value=> ",data)
        //prints: [not an int] value => great (as expected)
    }
}

```

###### Заблокированные рутины и утечки ресурсов

*   уровень: средний

Роб Пайк рассказал о ряде фундаментальных шаблонов параллелизма в своей презентации [«Шаблоны Go Concurrency»](https://talks.golang.org/2012/concurrency.slide#1
                ) в Google I / O в 2012 году. Одним из них является получение первого результата из ряда целей.

```
func First(query string, replicas ...Search) Result {
    c := make(chan Result)
    searchReplica := func(i int) { c <- replicas[i](query) }
    for i := range replicas {
        go searchReplica(i)
    }
    return <-c
}

```

Функция запускает процедуры для каждой реплики поиска. Каждая программа отправляет свой результат поиска в канал результатов. Первое значение из канала результатов возвращается.

А как насчет результатов других goroutines? Как насчет самих горутин?

Канал результата в `First()`функции не буферизован. Это означает, что возвращается только первая программа. Все остальные goroutines застряли, пытаясь отправить свои результаты. Это означает, что если у вас есть более одной реплики, каждый вызов приведет к утечке ресурсов.

Чтобы избежать утечек, вы должны убедиться, что все goroutines выход. Одним из возможных решений является использование буферизованного канала результатов, достаточно большого, чтобы вместить все результаты.

```
func First(query string, replicas ...Search) Result {
    c := make(chan Result,len(replicas))
    searchReplica := func(i int) { c <- replicas[i](query) }
    for i := range replicas {
        go searchReplica(i)
    }
    return <-c
}

```

Другим потенциальным решением является использование `select`оператора с `default`регистром и буферизованным каналом результата, который может содержать одно значение. Этот `default`случай гарантирует, что программы не застрянут, даже если канал результатов не может получать сообщения.

```
func First(query string, replicas ...Search) Result {
    c := make(chan Result,1)
    searchReplica := func(i int) {
        select {
        case c <- replicas[i](query):
        default:
        }
    }
    for i := range replicas {
        go searchReplica(i)
    }
    return <-c
}

```

Вы также можете использовать специальный канал отмены, чтобы прерывать рабочих.

```
func First(query string, replicas ...Search) Result {
    c := make(chan Result)
    done := make(chan struct{})
    defer close(done)
    searchReplica := func(i int) {
        select {
        case c <- replicas[i](query):
        case <- done:
        }
    }
    for i := range replicas {
        go searchReplica(i)
    }

    return <-c
}

```

Почему в презентации содержались эти ошибки? Роб Пайк просто не хотел усложнять слайды. Это имеет смысл, но это может быть проблемой для новых разработчиков Go, которые будут использовать код как есть, не думая, что у него могут быть проблемы.

###### Один и тот же адрес для разных переменных нулевого размера

*   уровень: средний

Если у вас есть две разные переменные, не должны ли они иметь разные адреса? Ну, это не относится к Go :\-) Если у вас переменные нулевого размера, они могут использовать один и тот же адрес в памяти.

```
package main

import (
  "fmt "
)

type data struct {
}

func main() {
  a := &data{}
  b := &data{}

  if a == b {
    fmt.Printf("same address - a=%p b=%p\n ",a,b)
    //prints: same address - a=0x1953e4 b=0x1953e4
  }
}

```

###### Первое использование йоты не всегда начинается с нуля

*   уровень: средний

Может показаться, что `iota`идентификатор похож на оператор приращения. Вы начинаете новую декларацию констант, и при первом использовании `iota`вы получаете ноль, при втором ее использовании вы получаете один и так далее. Это не всегда так.

```
package main

import (
  "fmt "
)

const (
  azero = iota
  aone  = iota
)

const (
  info  = "processing "
  bzero = iota
  bone  = iota
)

func main() {
  fmt.Println(azero,aone) //prints: 0 1
  fmt.Println(bzero,bone) //prints: 1 2
}

```

Это `iota`действительно оператор индекса для текущей строки в блоке объявления констант, поэтому, если первое использование `iota`не является первой строкой в ​​блоке объявления констант, начальное значение не будет равно нулю.

###### Использование методов получения указателя на экземплярах значений

*   уровень: продвинутый

Можно вызывать метод\-приемник указателя для значения, пока оно является адресуемым. Другими словами, в некоторых случаях вам не нужно иметь версию метода для получателя значения.

Хотя не каждая переменная является адресуемой. Элементы карты не адресуемы. Переменные, на которые ссылаются через интерфейсы, также не являются адресуемыми.

```
package main

import "fmt "

type data struct {
    name string
}

func (p *data) print() {
    fmt.Println("name: ",p.name)
}

type printer interface {
    print()
}

func main() {
    d1 := data{"one "}
    d1.print() //ok

    var in printer = data{"two "} //error
    in.print()

    m := map[string]data {"x ":data{"three "}}
    m["x "].print() //error
}

```

Ошибки компиляции:

> /tmp/sandbox017696142/main.go:21: невозможно использовать литерал данных (тип данных) в качестве принтера типа в присваивании: данные не реализуют принтер (метод печати имеет приемник указателя)
> /tmp/sandbox017696142/main.go:25: невозможно вызвать метод указателя на m \["x "\] /tmp/sandbox017696142/main.go:25: не может получить адрес m \["x "\]

###### Обновление полей значений карты

*   уровень: продвинутый

Если у вас есть карта структурных значений, вы не можете обновить отдельные структурные поля.

Сбой:

```
package main

type data struct {
    name string
}

func main() {
    m := map[string]data {"x ":{"one "}}
    m["x "].name = "two " //error
}

```

Ошибка компиляции:

> /tmp/sandbox380452744/main.go:9: невозможно назначить m \["x "\]. name

Это не работает, потому что элементы карты не адресуемы.

Что может быть еще более запутанным для новых разработчиков Go, так это тот факт, что элементы слайса являются адресуемыми.

```
package main

import "fmt "

type data struct {
    name string
}

func main() {
    s := []data {{"one "}}
    s[0].name = "two " //ok
    fmt.Println(s)    //prints: [{two}]
}

```

Обратите внимание, что некоторое время назад было возможно обновить поля элемента карты в одном из компиляторов Go (gccgo), но это поведение было быстро исправлено :\-) Это также рассматривалось как потенциальная возможность для Go 1.3. Это не было достаточно важно, чтобы поддерживать в тот момент, поэтому он все еще в списке задач.

Первый способ \- использовать временную переменную.

```
package main

import "fmt "

type data struct {
    name string
}

func main() {
    m := map[string]data {"x ":{"one "}}
    r := m["x "]
    r.name = "two "
    m["x "] = r
    fmt.Printf("%v ",m) //prints: map[x:{two}]
}

```

Другой обходной путь \- использовать карту указателей.

```
package main

import "fmt "

type data struct {
    name string
}

func main() {
    m := map[string]*data {"x ":{"one "}}
    m["x "].name = "two " //ok
    fmt.Println(m["x "]) //prints: &{two}
}

```

Кстати, что происходит, когда вы запускаете этот код?

```
package main

type data struct {
    name string
}

func main() {
    m := map[string]*data {"x ":{"one "}}
    m["z "].name = "what? " //???
}

```

###### Интерфейсы "nil " и значения интерфейсов "nil "

*   уровень: продвинутый

Это второй самый распространенный недостаток в Go, потому что интерфейсы не являются указателями, даже если они могут выглядеть как указатели. Переменные интерфейса будут иметь значение «ноль» только в том случае, если их поля типа и значения имеют значение «ноль».

Поля типа и значения интерфейса заполняются на основе типа и значения переменной, используемой для создания соответствующей переменной интерфейса. Это может привести к неожиданному поведению, когда вы пытаетесь проверить, равна ли переменная интерфейса нулю.

```
package main

import "fmt "

func main() {
    var data *byte
    var in interface{}

    fmt.Println(data,data == nil) //prints: <nil> true
    fmt.Println(in,in == nil)     //prints: <nil> true

    in = data
    fmt.Println(in,in == nil)     //prints: <nil> false
    //'data' is 'nil', but 'in' is not 'nil'
}

```

Следите за этой ловушкой, когда у вас есть функция, которая возвращает интерфейсы.

Неправильно:

```
package main

import "fmt "

func main() {
    doit := func(arg int) interface{} {
        var result *struct{} = nil

        if(arg > 0) {
            result = &struct{}{}
        }

        return result
    }

    if res := doit(-1); res != nil {
        fmt.Println("good result: ",res) //prints: good result: <nil>
        //'res' is not 'nil', but its value is 'nil'
    }
}

```

Работает:

```
package main

import "fmt "

func main() {
    doit := func(arg int) interface{} {
        var result *struct{} = nil

        if(arg > 0) {
            result = &struct{}{}
        } else {
            return nil //return an explicit 'nil'
        }

        return result
    }

    if res := doit(-1); res != nil {
        fmt.Println("good result: ",res)
    } else {
        fmt.Println("bad result (res is nil) ") //here as expected
    }
}

```

###### Переменные стека и кучи

*   уровень: продвинутый

Вы не всегда знаете, расположена ли ваша переменная в стеке или куче. В C ++ создание переменных с использованием `new`оператора всегда означает, что у вас есть куча переменных. В Go компилятор решает, где будет размещена переменная, даже если используются функции `new()`или `make()`. Компилятор выбирает место для хранения переменной на основе ее размера и результата «анализа выхода». Это также означает, что можно возвращать ссылки на локальные переменные, что не так в других языках, таких как C или C ++.

Если вам нужно знать, где размещены ваши переменные, передайте флаг \-m "gm " в "go build " или "go run " (например, `go run -gcflags -m app.go`).

###### GOMAXPROCS, параллелизм и параллелизм

*   уровень: продвинутый

В версии 1.4 и ниже используется только один контекст выполнения / поток ОС. Это означает, что в любой момент времени может выполняться только одна программа. Начиная с 1.5 Go устанавливает количество контекстов выполнения равным количеству логических ядер ЦП, возвращаемых `runtime.NumCPU()`. Это число может совпадать или не совпадать с общим количеством логических ядер ЦП в вашей системе в зависимости от настроек соответствия ЦП вашего процесса. Вы можете изменить это число, изменив `GOMAXPROCS`переменную окружения или вызвав `runtime.GOMAXPROCS()`функцию.

Существует распространенное заблуждение, `GOMAXPROCS`представляющее количество процессоров, которые Go будет использовать для запуска программ. Документация по `runtime.GOMAXPROCS()`функциям добавляет больше путаницы. Описание `GOMAXPROCS`переменной ( [https://golang.org/pkg/runtime/](https://golang.org/pkg/runtime/ ) ) лучше говорит о потоках ОС.

Вы можете установить `GOMAXPROCS`больше, чем количество ваших процессоров. Начиная с 1.10, для GOMAXPROCS больше нет ограничения. Максимальное значение, которое `GOMAXPROCS`раньше составляло 256, было позже увеличено до 1024 в 1,9.

```
package main

import (
    "fmt "
    "runtime "
)

func main() {
    fmt.Println(runtime.GOMAXPROCS(-1)) //prints: X (1 on play.golang.org)
    fmt.Println(runtime.NumCPU())       //prints: X (1 on play.golang.org)
    runtime.GOMAXPROCS(20)
    fmt.Println(runtime.GOMAXPROCS(-1)) //prints: 20
    runtime.GOMAXPROCS(300)
    fmt.Println(runtime.GOMAXPROCS(-1)) //prints: 256
}

```

###### Перезапись операции чтения и записи

*   уровень: продвинутый

Go может изменить порядок некоторых операций, но это гарантирует, что общее поведение в программе, где это происходит, не изменится. Тем не менее, это не гарантирует порядок выполнения нескольких групп.

```
package main

import (
    "runtime "
    "time "
)

var _ = runtime.GOMAXPROCS(3)

var a, b int

func u1() {
    a = 1
    b = 2
}

func u2() {
    a = 3
    b = 4
}

func p() {
    println(a)
    println(b)
}

func main() {
    go u1()
    go u2()
    go p()
    time.Sleep(1 * time.Second)
}

```

Если вы запустите этот код несколько раз, вы можете увидеть эти `a`и `b`переменные комбинации:

> 1
> 2
>
> 3
> 4
>
> 0
> 2
>
> 0
> 0
>
> 1
> 4

Наиболее интересным сочетанием для `a`и `b`является «02». Это показывает, что `b`было обновлено ранее `a`.

Если вам нужно сохранить порядок операций чтения и записи в нескольких программах, вам нужно использовать каналы или соответствующие конструкции из пакета «sync».

###### Упреждающее планирование

*   уровень: продвинутый

Возможно иметь мошенническую программу, которая мешает запуску других программ. Это может произойти, если у вас есть `for`цикл, который не позволяет планировщику работать.

```
package main

import "fmt "

func main() {
    done := false

    go func(){
        done = true
    }()

    for !done {
    }
    fmt.Println("done! ")
}

```

`for`Цикл не должен быть пустым. Это будет проблемой, если он содержит код, который не запускает выполнение планировщика.

Планировщик будет запускаться после GC, операторов "go ", блокировки операций канала, блокировки системных вызовов и операций блокировки. Он также может запускаться при вызове не встроенной функции.

```
package main

import "fmt "

func main() {
    done := false

    go func(){
        done = true
    }()

    for !done {
        fmt.Println("not done! ") //not inlined
    }
    fmt.Println("done! ")
}

```

Чтобы выяснить, является ли функция, которую вы вызываете в `for`цикле, встроенной, передайте флаг \-m "gm " в "go build " или "go run " (например, `go build -gcflags -m`).

Другим вариантом является явный вызов планировщика. Вы можете сделать это с помощью `Gosched()`функции из пакета "runtime ".

```
package main

import (
    "fmt "
    "runtime "
)

func main() {
    done := false

    go func(){
        done = true
    }()

    for !done {
        runtime.Gosched()
    }
    fmt.Println("done! ")
}

```

Обратите внимание, что код выше содержит условие гонки. Это было сделано намеренно, чтобы показать гочин планирования.

###### Импортировать блоки C и Multiline Import

*   уровень: Cgo

Вам необходимо импортировать пакет "C ", чтобы использовать Cgo. Вы можете сделать это одной строкой `import`или `import`блоком.

```
package main

/*
#include <stdlib.h>
*/
import (
  "C "
)

import (
  "unsafe "
)

func main() {
  cs := C.CString("my go string ")
  C.free(unsafe.Pointer(cs))
}

```

Если вы используете `import`формат блока, вы не можете импортировать другие пакеты в том же блоке.

```
package main

/*
#include <stdlib.h>
*/
import (
  "C "
  "unsafe "
)

func main() {
  cs := C.CString("my go string ")
  C.free(unsafe.Pointer(cs))
}

```

Ошибка компиляции:

> ./main.go:13:2: не удалось определить тип имени для C.free

###### Нет пустых строк между Import C и Cgo Комментарии

*   уровень: Cgo

Одна из первых ошибок, связанных с Cgo, \- это расположение комментариев cgo над `import "C "`утверждением.

```
package main

/*
#include <stdlib.h>
*/

import "C "

import (
  "unsafe "
)

func main() {
  cs := C.CString("my go string ")
  C.free(unsafe.Pointer(cs))
}

```

Ошибка компиляции:

> ./main.go:15:2: не удалось определить тип имени для C.free

Убедитесь, что у вас нет пустых строк над `import "C "`заявлением.

###### Невозможно вызвать функции C с переменными аргументами

*   уровень: Cgo

Вы не можете вызывать функции C с переменными аргументами напрямую.

```
package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C "

import (
  "unsafe "
)

func main() {
  cstr := C.CString("go ")
  C.printf("%s\n ",cstr) //not ok
  C.free(unsafe.Pointer(cstr))
}

```

Ошибка компиляции:

> ./main.go:15:2: неожиданный тип: ...

Вы должны обернуть свои переменные C\-функции в функции с известным числом параметров.

```
package main

/*
#include <stdio.h>
#include <stdlib.h>

void out(char* in) {
  printf("%s\n ", in);
}
*/
import "C "

import (
  "unsafe "
)

func main() {
  cstr := C.CString("go ")
  C.out(cstr) //ok
  C.free(unsafe.Pointer(cstr))
}

```
