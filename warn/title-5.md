# Глава 5 - Лакомые кусочки

В этой главе мы поговорим о возможностях Go, которые не вписываются в остальные разделы.

## Обработка ошибок

Предпочтительным способом обработки ошибок в Go является возвращение значений вместо исключений. Взглянем на функцию strconv.Atoi, которая принимает строку и пытается конвертировать её в целое число:

```golang
package main

import (
  "fmt"
  "os"
  "strconv"
)

func main() {
  if len(os.Args) != 2 {
    os.Exit(1)
  }

  n, err := strconv.Atoi(os.Args[1])
  if err != nil {
    fmt.Println("не является числом")
  } else {
    fmt.Println(n)
  }
}
```

Вы можете создать свой тип ошибок. Единственное требование, которое необходимо выполнить, это реализовать встроенный интерфейс error:

```golang
type error interface {
  Error() string
}
```

Также мы можем создать свою ошибку с помощью импорта пакета errors и вызова функции New:

```golang
import (
  "errors"
)


func process(count int) error {
  if count < 1 {
    return errors.New("Invalid count")
  }
  ...
  return nil
}
```

Это общепринятый способ использования переменных ошибок в стандартной библиотеке Go. Например, в пакете io есть переменная EOL, которая определяется так:

```golang
var EOF = errors.New("EOF")
```

Это переменная пакета (определённая вне функции), которая имеет публичный доступ (имя начинается с большой буквы). Различные функции могут возвращать эту ошибку во время чтения из файла или STDIN. Если в вашем контексте она имеет смысл, вам тоже нужно её использовать. Так можно обработать чтение только одного файла:

```golang
package main

import (
  "fmt"
  "io"
)

func main() {
  var input int
  _, err := fmt.Scan(&input)
  if err == io.EOF {
    fmt.Println("no more input!")
  }
}
```

И последнее замечание, в Go есть функции panic и recover. panic похожа на выброс исключения, а recover на catch. Они редко используются.

## Defer

Хотя в Go и есть сборщик мусора, некоторые ресурсы требуют, чтобы мы явно освобождали их. Например, нам нужно вызывать Close() для закрытия файла, после того, как работа с ним окончена. Такой код всегда опасен. С одной стороны, когда мы пишем функцию, легко забыть вызвать Close, для того, что мы открыли на 10 строк выше. С другой, функция может иметь несколько мест с возвращением результата, и нужно вызывать закрытие ресурса в каждом из них. Решением в Go является ключевое слово defer:

```golang
package main

import (
  "fmt"
  "os"
)

func main() {
  file, err := os.Open("a_file_to_read")
  if err != nil {
    fmt.Println(err)
    return
  }
  defer file.Close()
  // read the file
}
```

Если вы попытаетесь выполнить этот код, вы, вероятно, получите ошибку (файл не существует). Смысл в том, чтобы показать как работает defer. Неважно, где находится ваш defer, он всё равно будет выполнен после того, как метод вернет результат. Это позволяет вам освобождать ресурсы прямо там же, где вы их инициализировали и спасает от дублирования кода в случае нескольких return.


**go fmt**

Большинство программ, написанных на Go, используют одинаковые правила форматирования. Символ табуляции используется для отступа, а скобка ставится на той же строке, что и инструкция.

Я знаю, у вас есть свой собственный стиль и вы придерживаетесь его. Я следовал ему долгое время, но я рад, что в конечном итоге сдался. Главная причина этому была в команде go fmt. Она проста в использовании и не вызывает споров (по поводу личных предпочтений).

Когда вы находитесь внутри проекта, вы можете применить правила форматирования для него и всех под-проектов с помощью:

**go fmt ./...**

Попробуйте. Она делает больше, чем просто расставляет отступы. Она выравнивает объявления полей, а так же сортирует ваши импорты в алфавитном порядке.

## Инициализация в условии

Go поддерживает немного модифицированные условные блоки. Инициализация значения в них имеет приоритет при вычислении условия:

```golang
if x := 10; count > x {
  ...
}
```

Это довольно простой пример. В реальном коде будет что-то такое:

```golang
if err := process(); err != nil {
  return err
}
```

Интересно, что в то время, как эти значения недоступны вне инструкции if, они доступны внутри else if или else.

## Пустой интерфейс и преобразования

В многих объектно-ориентированных языках существует базовый класс, часто называемый object, который является супер классом для всех остальных классов. В Go нет наследования и нет такого супер класса. Что у него есть, так это пустой интерфейс без методов: interface{}. Так как каждый тип реализует все 0 методов этого интерфейса, то можно сказать, что в неявном виде каждый тип реализует пустой интерфейс.

Если бы нам было нужно, мы бы могли написать функцию add со следующей сигнатурой:

```go
func add(a interface{}, b interface{}) interface{} {
  ...
}
```

Для преобразования переменной в определенный тип, используйте .(ТИП):

```golang
return a.(int) + b.(int)
Также вы можете использовать такой switch:

switch a.(type) {
  case int:
    fmt.Printf("a теперь int и равно %d\n", a)
  case bool, string:
    // ...
  default:
    // ...
}
```

Вы встретите и, возможно, будете использовать пустой интерфейс чаще, чем может показаться с первого взгляда. Хотя стоит признать, что он не способствует чистоте кода. Преобразование значений туда и обратно – это и некрасиво и опасно. Но иногда в языках со статической типизацией это единственный выбор.

## Строки и массивы байтов

Строки и массивы байтов тесто связаны. Мы можем легко конвертировать одно в другое:

```golang
stra := "the spice must flow"
byts := []byte(stra)
strb := string(byts)
```

На самом деле, такой способ преобразования, является также общим для всех типов. Некоторые функции ожидают явно int32, или int64, или их беззнаковые эквиваленты. Используется это так:

**int64(count)**

Тем не менее, иметь дело с байтами и строками вы, вероятно, будете часто. Заметьте, что когда вы используете []byte(X) или string(X), вы создаёте копию данных. Это необходимо, так как строки в Go неизменяемы.

Строки состоят из рун (в Go, тип rune является псевдонимом для int32), которые являются частями Юникода. Если взять длину строки, мы можете получить не то, чего ожидали. Следующий код выведет 3:

```golang
fmt.Println(len("椒"))
```

Когда вы производите итерацию по строке с использованием range, вы получаете руны, а не байты. Конечно, когда вы переводите строку в []byte, вы получаете корректные данные.

## Тип функция

**Функция** – это тип первого порядка:

```golang
type Add func(a int, b int) int
```

который можно использовать как угодно – как поле структуры, как параметр, как возвращаемое значение.

```golang
package main

import (
  "fmt"
)

type Add func(a int, b int) int

func main() {
  fmt.Println(process(func(a int, b int) int{
      return a + b
  }))
}

func process(adder Add) int {
  return adder(1, 2)
}
```

Использование таких функций может помочь отделить код от реализации, как и в случае с интерфейсами.

## Перед тем как продолжить

Мы рассмотрели различные аспекты программирования с Go. В частности, мы увидели, как происходит обработка ошибок и как освобождаются ресурсы на примере открытых файлов. Многие люди не любят подход Go к обработке ошибок. Она кажется шагом назад. Иногда, я согласен. Тем не менее, я считаю, что такой код легче отслеживать. defer является необычным, но практичным подходом к управлению ресурсами. На самом деле он не привязан конкретно к ресурсам. Вы можете использовать defer для любых целей, таких как логирование после завершения функции.

Конечно мы рассмотрели не все вкусности, которые предлагает Go. Но вы должны чувствовать себя достаточно комфортно при понимании того, с чем вы столкнётесь.