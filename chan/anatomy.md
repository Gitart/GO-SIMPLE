# Анатомия каналов в Go

[Go \*](https://habr.com/ru/hub/go/)

Из песочницы

Привет, Хабр! Представляю вашему вниманию перевод статьи ["Anatomy of Channels in Go"](https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb) автора Uday Hiwarale.

## Что такое каналы?

Канал — это объект связи, с помощью которого горутины обмениваются данными. Технически это конвейер (или труба), откуда можно считывать или помещать данные. То есть одна горутина может отправить данные в канал, а другая — считать помещенные в этот канал данные.

## Создание канала

Go для создания канала предоставляет ключевое слово chan. Канал может передавать данные только одного типа, данные других типов через это канал передавать невозможно.

```
package main

import "fmt"

func main() {
    var c chan int
    fmt.Println(c)
}
```

[Пример в play.golang.org](https://play.golang.org/p/iWOFLfcgfF-)

Программа выше создает канал `c`, который будет передавать `int`. Данная программа выведет `<nil>`, потому что нулевое значение канала — это `nil`. Такой канал абсолютно бесполезен. Вы не можете передать или получить данные из канала, так как он не был создан (инициализирован). Для его создания необходимо использовать `make`.

```
package main

import "fmt"

func main() {
    c := make(chan int)

    fmt.Printf("type of `c` is %T\n", c)
    fmt.Printf("value of `c` is %v\n", c)
}
```

[Пример в play.golang.org](https://play.golang.org/p/N4dU7Ql9bK7)

В данном примере используется короткий синтаксис `:=` для создания канала с использованием функции `make`. Программа выше выводит следующий результат:

```
type of `c` is chan int
value of `c` is 0xc0420160c0
```

Обратите внимание на значение переменной `c`, это адрес в памяти. В go каналы являются указателями. В большинстве своем, когда вам необходимо взаимодействовать с горутиной, вы помещаете канал как аргумент в функцию или метод. Горутина получает этот канал как аргумент, и вам не нужно разыменовывать его для того, чтобы извлечь или передать данные через этот канал.

## Запись и чтение данных

Go предоставляет простой синтаксис для чтения `<-` и записи в канал

```
c <- data
```

В этом примере мы передаем данные в канал `c`. Направление стрелки указывает на то, что мы извлекаем данные из `data` и помещаем в канал `c`.

```
<- c
```

А здесь мы считываем данные с канала `c`. Эта операция не сохраняет данные в переменную и она является корректной. Если вам необходимо сохранить данные с канала в переменную, вы можете использовать следующий синтаксис:

```
var data int
data = <- c
```

Теперь данные из канала `c`, который имеет тип `int`, могут быть записаны в переменную `data`. Так же можно упростить запись, используя короткий синтаксис:

```
data := <- c
```

Go определит тип данных, передаваемый каналу `c`, и предоставит `data` корректный тип данных.

Все вышеобозначенные операции с каналом являются блокируемыми. Когда вы помещаете данные в канал, горутина блокируется до тех пор, пока данные не будут считаны другой горутиной из этого канала. В то же время операции канала говорят планировщику о планировании другой горутины, поэтому программа не будет заблокирована полностью. Эти функции весьма полезны, так как отпадает необходимость писать блокировки для взаимодействия горутин.

## Каналы на практике

```
package main

import "fmt"

func greet(c chan string) {
    fmt.Println("Hello " + <-c + "!")
}

func main() {
    fmt.Println("main() started")
    c := make(chan string)

    go greet(c)

    c <- "John"
    fmt.Println("main() stopped")
}
```

[Пример в play.golang.org](https://play.golang.org/p/OeYLKEz7qKi)

Разберем программу по шагам:

1.  Мы объявили функцию `greet`, которая принимает канал `c` как аргумент. В этой функции мы считываем данные из канала `c` и выводим в консоль.
2.  В функции `main` программа сначала выводит `"main() started"`.
3.  Затем мы, используя `make`, создаем канал `c` с типом даных `string`.
4.  Помещаем канал `с` в функцию `greet` и запускаем функцию как горутину, используя ключевое слово `go`.
5.  Теперь у нас имеется две горутины `main` и `greet`, `main` по-прежнему остается активной.
6.  Помещаем данные в канал `с` и в этот момент `main` блокируется до тех пор, пока другая горутина (`greet`) не считает данные из канала `c`. Планировщик Go планирует запуск `greet` и выполняет описанное в первом пункте.
7.  После чего `main` снова становится активной и выводит в консоль `"main() stopped"`.

## Deadlock (Взаимная блокировка)

Как уже ранее говорилось, чтение или запись данных в канал блокирует горутину и контроль передается свободной горутине. Представим, что такие горутины отсутствуют, либо они все "спят". В этот момент возникает deadlock, который приведет к аварийному завершению программы.

> Если вы попытаетесь считать данные из канала, но в канале будут отсутствовать данные, планировщик заблокирует текущую горутину и разблокирует другую в надежде, что какая-либо горутина передаст данные в канал. То же самое произойдет в случае отправки данных: планировщик заблокирует передающую горутину, пока другая не считает данные из канала.

Примером deadlock может быть `main` горутина, которая эксклюзивно производит операции с каналом.

```
package main

import "fmt"

func main() {
    fmt.Println("main() started")

    c := make(chan string)
    c <- "John"

    fmt.Println("main() stopped")
}
```

[Пример в play.golang.org](https://play.golang.org/p/2KTEoljdci_f)

Программа выше выведет следующее при попытке ее исполнить:

```
main() started
fatal error: all goroutines are asleep - deadlock!
goroutine 1 [chan send]:
main.main()
        program.go:10 +0xfd
exit status 2
```

### Закрытие канала

В Go так же можно закрыть канал, через закрытый канал невозможно будет передать или принять данные. Горутина может проверить закрыт канал или нет, используя следующую конструкцию: `val, ok := <- channel`, где ok будет истиной в случае, если канал открыт или операция чтения может быть выполнена, иначе `ok` будет `false`, если канал закрыт и отсутствуют данных для чтения из него. Закрыть канал можно, используя встроенную функцию `close`, используя следующий синтаксис `close(channel)`. Давайте рассмотрим следующий пример:

```
package main

import "fmt"

func greet(c chan string) {
    <-c // for John
    <-c // for Mike
}

func main() {
    fmt.Println("main() started")

    c := make(chan string, 1)

    go greet(c)
    c <- "John"

    close(c) // closing channel

    c <- "Mike"
    fmt.Println("main() stopped")
}
```

[Пример в play.golang.org](https://play.golang.org/p/LMmAq4sgm02)

> Для понимания концепта блокировки первая операция отправки `c <- "John"` будет блокирующей, и другая горутина должна будет считать данные из канала, следовательно `greet` горутина будет запланирована планировщиком. Затем первая операция чтения будет неблокируемой, поскольку присутствуют данные для чтения в канале `c`. Вторая операция чтения будет блокируемой, потому что в канале `c` отсутствуют данные, поэтому планировщик переключится на `main` горутину и программа выполнит закрытие канала `close(c)`.

Вывод программы:

```
main() started
panic: send on closed channel

goroutine 1 [running]:
main.main()
    program.go:20 +0x120
exit status 2
```

Как вы можете заметить, программа завершилась с ошибкой, которая говорит, что запись в закрытый канал невозможна. Для дальнейшего понимания закрытия каналов давайте рассмотрим пример с циклом `for`.

### Пример с циклом for

```
package main

import "fmt"

func squares(c chan int) {
    for i := 0; i <= 9; i++ {
        c <- i * i
    }

    close(c) // close channel
}

func main() {
    fmt.Println("main() started")
    c := make(chan int)

    go squares(c) // start goroutine

    // periodic block/unblock of main goroutine until chanel closes
    for {
        val, ok := <-c
        if ok == false {
            fmt.Println(val, ok, "<-- loop broke!")
            break // exit break loop
        } else {
            fmt.Println(val, ok)
        }
    }

    fmt.Println("main() stopped")
}
```

[Пример в play.golang.org](https://play.golang.org/p/X58FTgSHhXi)

Бесконечный цикл может быть полезен для чтения данных из канала, когда мы не знаем сколько данных мы ожидаем. В этом примере мы создаем горутину `squares`, которая последовательно возвращает квадраты чисел от 0 до 9. В `main` мы считываем эти числа внутри цикла `for`.

В цикле мы считываем данные из канала, используя ранее рассмотренный синтаксис `val, ok := <-c`, где `ok` предоставляет нам информацию о том, что канал закрыт. В горутине `squares` после того, как записали все данные, мы закрываем канал, используя функцию `close`. Когда `ok` будет `true`, программа выведет значение `val` и статус канала (переменная `ok`). Когда `ok` станет `false`, мы завершим цикл, используя ключевое слово `break`. Таким образом мы получим следующий результат:

```
main() started
0 true
1 true
4 true
9 true
16 true
25 true
36 true
49 true
64 true
81 true
0 false <-- loop broke!
main() stopped
```

> Когда канал закрыт, значение `val`, считанное горутиной, является нулевым значением, в зависимости от типа данных канала. Так как в нашем случае тип данных канала `int`, то нулевое значение будет 0, как раз это мы и видим в этой строке: `0 false <-- loop broke!`

Для того, чтобы избежать столь громоздкой проверки закрытия канала в случае цикла `for`, Go предоставляет ключевое слово `range`, которое автоматически останавливает цикл, когда канал будет закрыт. Давайте перепишем нашу программу с использованием `range`:

```
package main

import "fmt"

func squares(c chan int) {
    for i := 0; i <= 9; i++ {
        c <- i * i
    }

    close(c) // close channel
}

func main() {
    fmt.Println("main() started")
    c := make(chan int)

    go squares(c) // start goroutine

    // periodic block/unblock of main goroutine until chanel closes
    for val := range c {
        fmt.Println(val)
    }

    fmt.Println("main() stopped")
}
```

[Пример в play.golang.org](https://play.golang.org/p/ICCYbWO7ZvD)

В этом примере мы использовали `val := range c` вместо бесконечного цикла, где `range` будет считывать данные из канала до тех пор, пока канал не будет закрыт. В результате программа выведет следующее:

```
main() started
0
1
4
9
16
25
36
49
64
81
main() stopped
```

> Если вы не закроете канал для цикла `for` с использованием `range`, то программа будет завершена аварийно из-за `dealock` во время выполнения.

### Размер буфера канала

Как вы уже заметили, каждая операция отправки данных в канал блокирует текущую горутину. Но мы еще не использовали функцию `make` с 2-мя аргументами. Второй аргумент — это размер буфера канала. По-умолчанию размер буфера канала равен 0, такой канал называется небуферизированным каналом. То есть все, что мы пишем в канал, сразу доступно для чтения.

Когда размер буфера больше 0, горутина не блокируется до тех пор, пока буфер не будет заполнен. Когда буфер заполнен, любые значения отправляемые через канал, добавляются к буферу, отбрасывая предыдущее значение, которое доступно для чтения (где горутина будет заблокирована). Но есть один подвох, операция чтения на буферизированном канале является жадной, таким образом, как только операция чтения началась, она не будет завершена до полного опустошения буфера. Это означает, что горутина будет считывать буфер канала без блокировки до тех пор, пока буфер не станет пустым.

Для объявления буферизированного канала мы можем использовать следующий синтаксис:

```
c := make(chan Type, n)
```

Это выражение создаст канал с типом данных `Type` и размером буфера `n`. Текущая горутина не будет заблокирована, пока в канал не будет передано n+1 данных.

Давайте докажем, что горутина не блокируется, пока буфер не заполнится и не переполнится:

```
package main

import "fmt"

func squares(c chan int) {
    for i := 0; i <= 3; i++ {
        num := <-c
        fmt.Println(num * num)
    }
}

func main() {
    fmt.Println("main() started")
    c := make(chan int, 3)

    go squares(c)

    c <- 1
    c <- 2
    c <- 3

    fmt.Println("main() stopped")
}
```

[Пример в play.golang.org](https://play.golang.org/p/k0usdYZfp3D)

В этом примере канал `c` имеет размер буфера равным 3. Это означает, что он может содержать 3 значения(`c <- 3`), но поскольку буфер не переполняется (так как мы не поместили новое значение в буфер), `main` не будет блокироваться, и программа будет успешно завершена без вывода чисел. Вывод программы:

```
main() started
main() stopped
```

Теперь давайте поместим еще одно значение в канал:

```
package main

import "fmt"

func squares(c chan int) {
    for i := 0; i <= 3; i++ {
        num := <-c
        fmt.Println(num * num)
    }
}

func main() {
    fmt.Println("main() started")
    c := make(chan int, 3)

    go squares(c)

    c <- 1
    c <- 2
    c <- 3
    c <- 4 // blocks here

    fmt.Println("main() stopped")
}
```

[Пример в play.golang.org](https://play.golang.org/p/KGyiskRj1Wi)

Как упоминалось ранее, теперь мы помещаем дополнительное значение в буфер и `main` блокируется, затем стартует горутина `squares`, которая вычитывает все значения из буфера, пока он не станет пустым.

## Длина и емкость канала

Подобно срезам, буферизированный канал имеет длину и емкость. Длина канала — это количество значений в очереди (не считанных) в буфере канала, емкость — это размер самого буфера канала. Для того, чтобы вычислить длину, мы используем функцию `len`, а, используя функцию `cap`, получаем размер буфера.

```
package main

import "fmt"

func main() {
    c := make(chan int, 3)
    c <- 1
    c <- 2

    fmt.Printf("Length of channel c is %v and capacity of channel c is %v", len(c), cap(c))
    fmt.Println()
}

```

[Пример в play.golang.org](https://play.golang.org/p/qsDZu6pXLT7)

Вывод программы:

```
Length of channel c is 2 and capacity of channel c is 3
```

Вышеприведенная программа работает нормально и `deadlock` не возникает, потому что размер буфера канала равен 3, а мы записали только 2 значения в буфер, поэтому планировщик не попытался запланировать другую горутину и не заблокировал `main`. Вы даже можете считать эти данные в `main`, если вам это необходимо, потому что **буфер не заполнен**.

Другой пример:

```
package main

import "fmt"

func sender(c chan int) {
    c <- 1 // len 1, cap 3
    c <- 2 // len 2, cap 3
    c <- 3 // len 3, cap 3
    c <- 4 // <- goroutine blocks here
    close(c)
}

func main() {
    c := make(chan int, 3)

    go sender(c)

    fmt.Printf("Length of channel c is %v and capacity of channel c is %v\n", len(c), cap(c))

    // read values from c (blocked here)
    for val := range c {
        fmt.Printf("Length of channel c after value '%v' read is %v\n", val, len(c))
    }
}
```

[Пример в play.golang.org](https://play.golang.org/p/-gGpm08-wzz)

Вывод программы:

```
Length of channel c is 0 and capacity of channel c is 3
Length of channel c after value '1' read is 3
Length of channel c after value '2' read is 2
Length of channel c after value '3' read is 1
Length of channel c after value '4' read is 0
```

Дополнительный пример с буферизированным каналом:

```
package main

import (
    "fmt"
    "runtime"
)

func squares(c chan int) {
    for i := 0; i < 4; i++ {
        num := <-c
        fmt.Println(num * num)
    }
}

func main() {
    fmt.Println("main() started")
    c := make(chan int, 3)
    go squares(c)

    fmt.Println("active goroutines", runtime.NumGoroutine())
    c <- 1
    c <- 2
    c <- 3
    c <- 4 // blocks here

    fmt.Println("active goroutines", runtime.NumGoroutine())

    go squares(c)

    fmt.Println("active goroutines", runtime.NumGoroutine())

    c <- 5
    c <- 6
    c <- 7
    c <- 8 // blocks here

    fmt.Println("active goroutines", runtime.NumGoroutine())
    fmt.Println("main() stopped")
}

```

[Пример в play.golang.org](https://play.golang.org/p/sdHPDx64aor)

Вывод программы:

```
main() started
active goroutines 2
1
4
9
16
active goroutines 1
active goroutines 2
25
36
49
64
active goroutines 1
main() stopped

```

Используя буферизованный канал и цикл `for range`, мы можем читать с закрытых каналов. Поскольку у закрытых каналов данные все еще живут в буфере, их можно считать:

```
package main

import "fmt"

func main() {
    c := make(chan int, 3)
    c <- 1
    c <- 2
    c <- 3
    close(c)

    // iteration terminates after receiving 3 values
    for elem := range c {
        fmt.Println(elem)
    }
}
```

[Пример в play.golang.org](https://play.golang.org/p/vULFyWnpUoj)

## Работа с несколькими горутинами

Давайте напишем 2 горутины, одна для вычисления квадрата целого числа, а другая для вычисления куба:

```
package main

import "fmt"

func square(c chan int) {
    fmt.Println("[square] reading")
    num := <-c
    c <- num * num
}

func cube(c chan int) {
    fmt.Println("[cube] reading")
    num := <-c
    c <- num * num * num
}

func main() {
    fmt.Println("[main] main() started")

    squareChan := make(chan int)
    cubeChan := make(chan int)

    go square(squareChan)
    go cube(cubeChan)

    testNum := 3
    fmt.Println("[main] sent testNum to squareChan")

    squareChan <- testNum

    fmt.Println("[main] resuming")
    fmt.Println("[main] sent testNum to cubeChan")

    cubeChan <- testNum

    fmt.Println("[main] resuming")
    fmt.Println("[main] reading from channels")

    squareVal, cubeVal := <-squareChan, <-cubeChan
    sum := squareVal + cubeVal

    fmt.Println("[main] sum of square and cube of", testNum, " is", sum)
    fmt.Println("[main] main() stopped")
}
```

[Пример в play.golang.org](https://play.golang.org/p/6wdhWYpRfrX)

Разберем программу по шагам:

1.  Мы создали 2 функции `square` и `cube`, которые мы запускаем как горутины. Обе получают канал `c` c типом данных `int`, и считывают данные из него в переменную `num`. Затем мы пишем данные в канал `c`.
2.  В `main` горутине мы создаем два канала `squareChan` и `cubeChan` c типом данных `int`.
3.  Запускаем `square` и `cube` горутины.
4.  Так как контроль по-прежнему внутри `main` `testNum` получает значение 3.
5.  Затем мы отправляем данные в канал `squareChan` и `cubeChan`. Горутина `main` будет заблокирована, пока данные из каналов не будут считаны. Как только значение будет считано, горутина снова станет активной.
6.  Когда в `main` мы попытаемся прочитать данные из заданных каналов(`squareChan` и `cubeChan`), управление будет заблокировано, пока другие горутины (`square` и `cube`) не запишут данные в эти каналы. Мы также использовали сокращенный синтаксис `:=` для получения данных из каналов.
7.  Когда операция записи канала завершена, начинает выполняться `main`, после чего мы рассчитываем сумму и выводим ее.

Результат выполнения программы:

```
[main] main() started
[main] sent testNum to squareChan
[cube] reading
[square] reading
[main] resuming
[main] sent testNum to cubeChan
[main] resuming
[main] reading from channels
[main] sum of square and cube of 3  is 36
[main] main() stopped
```

### Однонаправленные каналы

До сих пор мы видели каналы, которые могут передавать и принимать данные. Но мы также можем создать канал, который будет однонаправленным. Например, канал, который сможет только считывать данные, и канал который сможет только записывать их.

Однонаправленный канал также создается с использованием `make`, но с дополнительным стрелочным синтаксисом.

```
roc := make(<-chan int)
soc := make(chan<- int)
```

Где `roc` канал для чтения, а `soc` канал для записи. Следует заметить, что каналы также имеют разный тип.

```
package main

import "fmt"

func main() {
    roc := make(<-chan int)
    soc := make(chan<- int)

    fmt.Printf("Data type of roc is `%T`\n", roc)
    fmt.Printf("Data type of soc is `%T\n", soc)
}

```

[Пример в play.golang.org](https://play.golang.org/p/JZO51IoaMg8)

Вывод программы:

```
Data type of roc is `<-chan int`
Data type of soc is `chan<- int
```

Но в чем смысл использования однонаправленного канала? Использование однонаправленного канала улучшает безопасность типов в программe, что, как следствие, порождает меньше ошибок.

Но допустим, что у вас есть программа, в которой вам нужно только читать данные из канала, а основная программа должна иметь возможность читать и записывать данные из/в тот же канал. Как это будет работать?

К счастью Go предоставляет простой синтаксис для преобразования двунаправленного канала в однонаправленный канал.

```
import "fmt"

func greet(roc <-chan string) {
    fmt.Println("Hello " + <-roc + "!")
}

func main() {
    fmt.Println("main() started")
    c := make(chan string)

    go greet(c)

    c <- "John"
    fmt.Println("main() stopped")
}
```

[Пример в play.golang.org](https://play.golang.org/p/k3B3gCelrGv)

Мы только что изменили параметры `greet` для того, чтобы преобразовать двунаправленный канал на канал для чтения данных. Теперь мы можем только считывать данные из этого канала, а любые операции чтения приведут к аварийному завершению программы со следующей ошибкой:

`"invalid operation: roc <- "some text" (send to receive-only type <-chan string)"`

### Анонимные горутины

Каналы также могут работать и с анонимными горутинами. Давайте изменим предыдущий пример, используя анонимные горутины.
Вот что у нас получилось:

```
package main

import "fmt"

func main() {
    fmt.Println("main() started")
    c := make(chan string)

    // launch anonymous goroutine
    go func(c chan string) {
        fmt.Println("Hello " + <-c + "!")
    }(c)

    c <- "John"
    fmt.Println("main() stopped")
}
```

[Пример в play.golang.org](https://play.golang.org/p/cM5nFgRha7c)

Как вы можете заметить вывод программы остался тот же самый.

### Канал с типом данных канала

Каналы являются [объектами первого класса](https://ru.wikipedia.org/wiki/%D0%9E%D0%B1%D1%8A%D0%B5%D0%BA%D1%82_%D0%BF%D0%B5%D1%80%D0%B2%D0%BE%D0%B3%D0%BE_%D0%BA%D0%BB%D0%B0%D1%81%D1%81%D0%B0), то есть они могут быть использованы как значение элемента структуры, или аргументы функции, как возврат значения из функции/метода и даже как тип для другого канала. В примере ниже мы используем канал в качестве типа данных для другого канала:

```
package main

import "fmt"

// gets a channel and prints the greeting by reading from channel
func greet(c chan string) {
    fmt.Println("Hello " + <-c + "!")
}

// gets a channels and writes a channel to it
func greeter(cc chan chan string) {
    c := make(chan string)
    cc <- c
}

func main() {
    fmt.Println("main() started")

    // make a channel `cc` of data type channel of string data type
    cc := make(chan chan string)

    go greeter(cc) // start `greeter` goroutine using `cc` channel

    // receive a channel `c` from `greeter` goroutine
    c := <-cc

    go greet(c) // start `greet` goroutine using `c` channel

    // send data to `c` channel
    c <- "John"

    fmt.Println("main() stopped")
}
```

[Пример в play.golang.org](https://play.golang.org/p/xVQvvb8O4De)

### select

`select` похож на `switch` без аргументов, но он может использоваться только для операций с каналами. Оператор `select` используется для выполнения операции только с одним из множества каналов, условно выбранного блоком case.

Давай взглянем на пример ниже, и обсудим как он работает:

```
package main

import (
    "fmt"
    "time"
)

var start time.Time
func init() {
    start = time.Now()
}

func service1(c chan string) {
    time.Sleep(3 * time.Second)
    c <- "Hello from service 1"
}

func service2(c chan string) {
    time.Sleep(5 * time.Second)
    c <- "Hello from service 2"
}

func main() {
    fmt.Println("main() started", time.Since(start))

    chan1 := make(chan string)
    chan2 := make(chan string)

    go service1(chan1)
    go service2(chan2)

    select {
    case res := <-chan1:
        fmt.Println("Response from service 1", res, time.Since(start))
    case res := <-chan2:
        fmt.Println("Response from service 2", res, time.Since(start))
    }

    fmt.Println("main() stopped", time.Since(start))
}
```

[Пример в play.golang.org](https://play.golang.org/p/ar5dZUQ2ArH)

В этом примере мы используем оператор `select` как `switch`, но вместо булевых операций, мы используем операции для чтения данных из канала. Оператор `select` также является блокируемым, за исключением использования `default`(позже вы увидите пример с его использованием). После выполнения одного из блоков `case`, горутина `main` будет разблокирована. Задались вопросом когда `case` условие выполнится?

Если все блоки `case` являются блокируемыми, тогда `select` будет ждать до момента, пока один из блоков `case` разблокируется и будет выполнен. Если несколько или все канальные операции не блокируемы, тогда один из неблокируемых `case` будет выбран случайным образом (Примечание переводчика: имеется ввиду случай, когда пришли одновременно данные из двух и более каналов).

Давайте наконец разберем программу, которую написали ранее. Мы запустили 2 горутины с независимыми каналами. Затем мы использовали оператор `select` c двумя `case` операторами. Один `case` считывает данные из `chan1` а другой из `chan2`. Так как каналы не используют буфер, операция чтения будет блокируемой. Таким образом оба `case` будут блокируемыми и `select` будет ждать до тех пор, пока один из `case` не разблокируется.

Когда программа находится в блоке `select` горутина `main` будет заблокирована и будут запланированы все горутины (по одной за раз), которые используются в блоке `select`, в нашем случае это `service1` и `service2`. `service1` ждет 3 секунды, после чего будет разблокирован и сможет записать данные в `chan1`. Таким же образом как и `service1` действует `service2`, только он ожидает 5 секунд и осуществляет запись в `chan2`. Так как `service1` разблокируется раньше, чем `service2`, первый `case` разблокируется раньше и произведет чтение из `chan1`, а второй `case` будет проигнорирован. После чего управление вернется в `main`, и программа завершится после вывода в консоль.

Вывод программы:

```
main() started 0s
Response from service 1 Hello from service 1 3s
main() stopped 3s
```

> Вышеприведенная программа имитирует реальный веб-сервис, в котором балансировщик нагрузки получает миллионы запросов и должен возвращать ответ от одной из доступных служб. Используя стандартные горутины, каналы и select, мы можем запросить ответ у нескольких сервисов, и тот, который ответит раньше всех, может быть использован.

Для того, чтобы симулировать случай, когда все блоки `case` разблокируются в одно и тоже время, мы может просто удалить вызов Sleep из горутин.

```
package main

import (
    "fmt"
    "time"
)

var start time.Time
func init() {
    start = time.Now()
}

func service1(c chan string) {
    c <- "Hello from service 1"
}

func service2(c chan string) {
    c <- "Hello from service 2"
}

func main() {
    fmt.Println("main() started", time.Since(start))

    chan1 := make(chan string)
    chan2 := make(chan string)

    go service1(chan1)
    go service2(chan2)

    select {
    case res := <-chan1:
        fmt.Println("Response from service 1", res, time.Since(start))
    case res := <-chan2:
        fmt.Println("Response from service 2", res, time.Since(start))
    }

    fmt.Println("main() stopped", time.Since(start))
}
```

[Пример в play.golang.org](https://play.golang.org/p/giSkkqt8XHb)

Данная программа выводит следующий результат:

```
main() started 0s
service2() started 481µs
Response from service 2 Hello from service 2 981.1µs
main() stopped 981.1µs
```

Но иногда вы можете получить следующий результат:

```
main() started 0s
service1() started 484.8µs
Response from service 1 Hello from service 1 984µs
main() stopped 984µs
```

Это происходит потому, что операции `chan1` и `chan2` выполняются практически одновременно, но все же существует некоторая разница во времени при исполнении и планировании горутин.

Для того, чтобы сделать все блоки `case` неблокируемыми, мы можем использовать каналы с буфером.

```
package main

import (
    "fmt"
    "time"
)

var start time.Time

func init() {
    start = time.Now()
}

func main() {
    fmt.Println("main() started", time.Since(start))
    chan1 := make(chan string, 2)
    chan2 := make(chan string, 2)

    chan1 <- "Value 1"
    chan1 <- "Value 2"
    chan2 <- "Value 1"
    chan2 <- "Value 2"

    select {
    case res := <-chan1:
        fmt.Println("Response from chan1", res, time.Since(start))
    case res := <-chan2:
        fmt.Println("Response from chan2", res, time.Since(start))
    }

    fmt.Println("main() stopped", time.Since(start))
}
```

[Пример в play.golang.org](https://play.golang.org/p/RLRGEmFQP3f)

Вывод может быть следующим:

```
main() started 0s
Response from chan2 Value 1 0s
main() stopped 1.0012ms
```

Или таким:

```
main() started 0s
Response from chan1 Value 1 0s
main() stopped 1.0012ms
```

В приведенной программе оба канала имеют буфер размером 2. Так как мы отправляем 2 значения в буфер, горутина не будет заблокирована и программа перейдет в блок `select`. Чтение из буферизированного канала не является блокируемой операцией, если буфер не пустой, поэтому все блоки `case` будут неблокируемыми, и во время выполнения Go выберет `case` случайным образом.

### default case

Так же как и `switch`, оператор `select` поддерживает оператор `default`. Оператор `default` является неблокируемым, но это еще не все, оператор `default` делает блок `select` всегда неблокируемым. Это означает, что операции отправки и чтение на любом канале (не имеет значения будет ли канал с буфером или без) всегда будут неблокируемыми.

Если значение будет доступно на каком-либо канале, то `select` выполнит этот `case`. Если нет, то он немедленно выполнит `default`.

```
package main

import (
    "fmt"
    "time"
)

var start time.Time

func init() {
    start = time.Now()
}

func service1(c chan string) {
    fmt.Println("service1() started", time.Since(start))
    c <- "Hello from service 1"
}

func service2(c chan string) {
    fmt.Println("service2() started", time.Since(start))
    c <- "Hello from service 2"
}

func main() {
    fmt.Println("main() started", time.Since(start))

    chan1 := make(chan string)
    chan2 := make(chan string)

    go service1(chan1)
    go service2(chan2)

    select {
    case res := <-chan1:
        fmt.Println("Response from service 1", res, time.Since(start))
    case res := <-chan2:
        fmt.Println("Response from service 2", res, time.Since(start))
    default:
        fmt.Println("No response received", time.Since(start))
    }

    fmt.Println("main() stopped", time.Since(start))
}

```

[Пример в play.golang.org](https://play.golang.org/p/rFMpc80EuT3)

Вывод программы:

```
main() started 0s
No response received 0s
main() stopped 0s
```

Так как в приведенной программе каналы используются без буфера, и значение еще отсутствует, в обоих каналах будет исполнен `default`. Если бы в блоке `select` отсутствовал `default`, то произошла бы блокировка и результат был бы другим.

Так как с `default` `select` не блокируется, планировщик не запускает доступные горутины. Но `main` можно заблокировать, вызвав `time.Sleep`. Таким образом все горутины будут исполнены, и когда управление перейдет в `main`, каналы будут иметь данные для чтения.

```
package main

import (
    "fmt"
    "time"
)

var start time.Time

func init() {
    start = time.Now()
}

func service1(c chan string) {
    fmt.Println("service1() started", time.Since(start))
    c <- "Hello from service 1"
}

func service2(c chan string) {
    fmt.Println("service2() started", time.Since(start))
    c <- "Hello from service 2"
}

func main() {
    fmt.Println("main() started", time.Since(start))

    chan1 := make(chan string)
    chan2 := make(chan string)

    go service1(chan1)
    go service2(chan2)

    time.Sleep(3 * time.Second)

    select {
    case res := <-chan1:
        fmt.Println("Response from service 1", res, time.Since(start))
    case res := <-chan2:
        fmt.Println("Response from service 2", res, time.Since(start))
    default:
        fmt.Println("No response received", time.Since(start))
    }

    fmt.Println("main() stopped", time.Since(start))
}
```

[Пример в play.golang.org](https://play.golang.org/p/eD0NHxHm9hN)

По итогу мы получим следующий результат:

```
main() started 0s
service1() started 0s
service2() started 0s
Response from service 1 Hello from service 1 3.0001805s
main() stopped 3.0001805s
```

Или такой, в некоторых случаях:

```
main() started 0s
service1() started 0s
service2() started 0s
Response from service 2 Hello from service 2 3.0000957s
main() stopped 3.0000957s
```

## Deadlock

Для того, чтобы избежать `deadlock`, можно использовать `default`, чтобы операции с каналами стали неблокируемыми, планировщик Go не будет планировать горутины для отправки данных в канал, даже если данные не доступны на данный момент.

```
package main

import (
    "fmt"
    "time"
)

var start time.Time

func init() {
    start = time.Now()
}

func main() {
    fmt.Println("main() started", time.Since(start))

    chan1 := make(chan string)
    chan2 := make(chan string)

    select {
    case res := <-chan1:
        fmt.Println("Response from chan1", res, time.Since(start))
    case res := <-chan2:
        fmt.Println("Response from chan2", res, time.Since(start))
    default:
        fmt.Println("No goroutines available to send data", time.Since(start))
    }

    fmt.Println("main() stopped", time.Since(start))
}
```

[Пример в play.golang.org](https://play.golang.org/p/S3Wxuqb8lMF)

Вывод программы:

```
main() started 0s
No goroutines available to send data 0s
main() stopped 0s
```

Аналогично получению данных, операция отправки данных будет работать также в случае использования оператора `default`, если присутствуют другие горутины, готовые принять отправленные данные (в режиме ожидания).

### nil каналы

Как мы уже знаем, нулевое значение в случае канала — это `nil`, из-за этого мы не может выполнять операции отправки или приема данных. При попытке отправить или принять данные через этот канал в блоке `select`, мы получим ошибку.

```
package main

import "fmt"

func service(c chan string) {
    c <- "response"
}

func main() {
    fmt.Println("main() started")

    var chan1 chan string

    go service(chan1)

    select {
    case res := <-chan1:
        fmt.Println("Response from chan1", res)
    }

    fmt.Println("main() stopped")
}
```

[Пример в play.golang.org](https://play.golang.org/p/uhraFubcF4S)

Вывод программы:

```
main() started
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [select (no cases)]:
main.main()
    program.go:17 +0xc0

goroutine 6 [chan send (nil chan)]:
main.service(0x0, 0x1)
    program.go:6 +0x40
created by main.main
    program.go:14 +0xa0
```

Из полученного результата мы можем заметить, что `select (no cases)` означает, что `select` оператор пустой, потому что конструкции `case` с нулевым каналом игнорируются. Но так как пустой `select{}` блокирует `main` горутину, активируется горутина `service`, которая попытается записать данные в `nil` канал, что впоследствии приведет к аварийному завершению программы со следующей ошибкой: `chan send (nil chan)`. Для того, чтобы этого избежать, можно использовать оператор `default`.

```
package main

import "fmt"

func service(c chan string) {
    c <- "response"
}

func main() {
    fmt.Println("main() started")

    var chan1 chan string

    go service(chan1)

    select {
    case res := <-chan1:
        fmt.Println("Response from chan1", res)
    default:
        fmt.Println("No response")
    }

    fmt.Println("main() stopped")
}
```

[Пример в play.golang.org](https://play.golang.org/p/upLsz52_CrE)

Вывод программы:

```
main() started
No response
main() stopped
```

В приведенной программе блоки `case` игнорируются, так как блок `default` исполняется первым. Поэтому планировщик не запускает горутину `service`. Такие программы, естественно, писать не стоит, необходимо всегда проверять, что канал не `nil`.

### Добавляем timeout

Ранее написанная программа не особенно полезна из-за того, что блок `default` выполнится раньше. Но иногда необходимо, чтобы определенный сервис ответил за определенное время, если он не отвечает, тогда должен выполниться блок `default`. Этого можно добиться, используя `case` с канальными операциями, которые будут разблокированы после заданного времени. Такая канальная операция предоставляется функцией `After` из пакета (package) `time`. Давайте рассмотрим следующий пример:

```
package main

import (
    "fmt"
    "time"
)

var start time.Time

func init() {
    start = time.Now()
}

func service1(c chan string) {
    time.Sleep(3 * time.Second)
    c <- "Hello from service 1"
}

func service2(c chan string) {
    time.Sleep(5 * time.Second)
    c <- "Hello from service 2"
}

func main() {
    fmt.Println("main() started", time.Since(start))

    chan1 := make(chan string)
    chan2 := make(chan string)

    go service1(chan1)
    go service2(chan2)

    select {
    case res := <-chan1:
        fmt.Println("Response from service 1", res, time.Since(start))
    case res := <-chan2:
        fmt.Println("Response from service 2", res, time.Since(start))
    case <-time.After(2 * time.Second):
        fmt.Println("No response received", time.Since(start))
    }

    fmt.Println("main() stopped", time.Since(start))
}
```

[Пример в play.golang.org](https://play.golang.org/p/mda2t2IQK__X)

Данная программа выдаст следующий результат через 2 секунды:

```
main() started 0s
No response received 2s
main() stopped 2s
```

В этой программе, благодаря конструкции `<-time.After(2 * time.Second)` горутина `main` будет разблокирована через 2 секунды. `time.After` создаёт канал, по которому посылаются метки времени с заданным интервалом. Так как данные из каналов `chan1` и `chan2` не были получены, выполняется 3-й блок, после чего программа успешно завершается.

Это может быть полезно в случае, когда вы не хотите ждать ответа от сервера продолжительное время. Если изменить `time.After(2 * time.Second)` на `time.After(10 * time.Second)` мы получим результат из `service1`.

### Пустой select

Подобно пустому `for{}`, пустой `select{}` так же является валидным, но есть подвох. Как мы уже знаем `select` блокируется до тех пор, пока один из блоков `case` не будет выполнен, но так как в пустом `select` отсутствуют блоки `case`, горутина не будет разблокирована, и как результат, мы получим `deadlock`.

```
package main

import "fmt"

func service() {
    fmt.Println("Hello from service!")
}

func main() {
    fmt.Println("main() started")

    go service()

    select {}

    fmt.Println("main() stopped")
}
```

[Пример в play.golang.org](https://play.golang.org/p/-pBd-BLMFOu)

В результате мы получим следующий вывод:

```
main() started
Hello from service!
fatal error: all goroutines are asleep - deadlock!
goroutine 1 [select (no cases)]:
main.main()
        program.go:16 +0xba
exit status 2
```

### WaitGroup

Теперь давайте представим состояние, когда вам нужно узнать, что все горутины были выполнены (Примечание переводчика: например, операция сложения запущенная в нескольких горутинах). Такая задача является прямо противоположной тому, что мы делали с `select`. Здесь мы дожидаемся полного завершения всех горутин.

На помощь нам приходит WaitGroup. Это структура со счетчиком, которая отслеживает сколько горутин вами было создано, и сколько из них было завершено (Примечание переводчика: сама она это делать не умеет, но есть методы, которые позволят вам добиться этого, так же подобного можно добиться с использованием каналов, но это считается устаревшим подходом и, как вы уже могли заметить, имеет ряд недостатков). Достижение счетчиком нуля будет означать, что все горутины были выполнены.

Давайте разберем следующий пример:

```
package main

import (
    "fmt"
    "sync"
    "time"
)

func service(wg *sync.WaitGroup, instance int) {
    time.Sleep(2 * time.Second)
    fmt.Println("Service called on instance", instance)
    wg.Done() // decrement counter
}

func main() {
    fmt.Println("main() started")
    var wg sync.WaitGroup // create waitgroup (empty struct)

    for i := 1; i <= 3; i++ {
        wg.Add(1) // increment counter
        go service(&wg, i)
    }

    wg.Wait() // blocks here
    fmt.Println("main() stopped")
}
```

[Пример в play.golang.org](https://play.golang.org/p/8qrAD9ceOfJ)

В этой программе мы создали пустой WaitGroup, внутри себя эта структура содержит приватные поля `noCopy` и `state1` ([https://golang.org/src/sync/waitgroup.go?s=574:929#L10](https://golang.org/src/sync/waitgroup.go?s=574:929#L10)). Структура имеет три метода: `Add`, `Wait` и `Done`. Давайте их рассмотрим.

Метод `Add` принимает `int` аргумент, который является `delta` (дельтой) для счетчика `WaitGroup`. Где счетчика — это число со значением, по умолчанию равным 0. Он хранит число запущенных горутин. Когда `WaitGroup` создана, значение счетчика будет равно 0, и мы можем увеличивать его, передавая `delta` как параметр метода `Add`. Счетчика не понимает автоматически, когда была запущена программа, поэтому нам нужно вручную увеличивать его, используя функцию `Add`.

Метод `Wait` используется для блокировки текущей горутины, когда мы его вызываем. Как только счетчик достигнет 0, горутина будет разблокирована. Поэтому нам необходимо как-то уменьшать значение счетчика.

Метод `Done` уменьшает значение счетчика. Он не принимает никаких параметров. (Примечание переводчика: если посмотреть исходники пакета `sync`, то можно увидеть, что внутри себя он просто вызывает [Add(-1)](https://golang.org/src/sync/waitgroup.go?s=574:929#L98)).

И так, после создания `wg`, мы запускаем итерацию в цикле `for` от 1 до 3х включительно. На каждой итерации мы запускаем горутину и инкрементируем счетчик на 1. Таким образом у нас будет 3 запущенных горутины, которые необходимо выполнить и `WaitGroup` со значением счетчика равным 3. Заметьте, что мы передали указатель на `wg` в горутину. Это необходимо, чтобы вызвать `Done` в горутине после завершения работы, что в свою очередь уменьшит значение счетчика.

После выполнения цикла `for`, мы запускаем `wg.Wait()`, чтобы передать управление другим горутинам, и, как следствие, это заблокирует наш `main` до тех пор, пока все горутины не будут завершены, и значение счетчика не будет равно 0. После чего `main` будет разблокирована, и программа будет успешно завершена.

Таким образом мы получим следующий вывод:

```
main() started
Service called on instance 1
Service called on instance 3
Service called on instance 2
main() stopped
```

Результат выше может отличаться, из-за порядка выполнения горутин.

### Пул воркеров

Как следует из названия, пул воркеров — это набор горутин, работающих одновременно для определенной задачи. В примере `WaitGroup` мы увидели набор горутин, работающих одновременно, но у них не было определенной задачи. Как только вы добавляете каналы в горутины, у них появляется какая-то работа, и они становятся пулом воркеров.

```
package main

import (
    "fmt"
    "time"
)

// worker than make squares
func sqrWorker(tasks <-chan int, results chan<- int, id int) {
    for num := range tasks {
        time.Sleep(time.Millisecond) // simulating blocking task
        fmt.Printf("[worker %v] Sending result by worker %v\n", id, id)
        results <- num * num
    }
}

func main() {
    fmt.Println("[main] main() started")

    tasks := make(chan int, 10)
    results := make(chan int, 10)

    // launching 3 worker goroutines
    for i := 0; i < 3; i++ {
        go sqrWorker(tasks, results, i)
    }

    // passing 5 tasks
    for i := 0; i < 5; i++ {
        tasks <- i * 2 // non-blocking as buffer capacity is 10
    }

    fmt.Println("[main] Wrote 5 tasks")

    // closing tasks
    close(tasks)

    // receving results from all workers
    for i := 0; i < 5; i++ {
        result := <-results // blocking because buffer is empty
        fmt.Println("[main] Result", i, ":", result)
    }

    fmt.Println("[main] main() stopped")
}
```

[Пример в play.golang.org](https://play.golang.org/p/IYiMV1I4lCj)

Вывод программы:

```
[main] main() started
[main] Wrote 5 tasks
[worker 0] Sending result by worker 0
[worker 2] Sending result by worker 2
[worker 1] Sending result by worker 1
[main] Result 0 : 4
[main] Result 1 : 0
[main] Result 2 : 16
[worker 2] Sending result by worker 2
[main] Result 3 : 64
[worker 0] Sending result by worker 0
[main] Result 4 : 36
[main] main() stopped
```

Итак, давайте разберемся с тем, что тут происходит:

1.  Функция `sqrWorker` принимает канал `tasks`, канал `results`, а так же `id`. Задача этой горутины — отправлять квадрат числа, полученного из канала `tasks`, в канал `results`.
2.  В функции `main`, мы создали каналы `tasks` и `result` с размером буфера, равной 10. Следовательно, любая операция отправки будет неблокируемой, пока буфер не заполнится. Поэтому канал с буфером большого размера — это неплохая идея.
3.  Затем мы порождаем несколько экземпляров `sqrWorker` в виде горутин с двумя вышеописанными каналами и параметром `id`, чтобы позже получить информацию о том, какой воркер выполняет задачу.
4.  Далее мы передали 5 значений каналу `tasks`, операция будет неблокируемой, так как размер буфера не превышен.
5.  Так как мы закончили с каналом `tasks`, закрываем его. В этом нет необходимости, но это сэкономит много времени в будущем, если появятся ошибки.
6.  Используя цикл `for` с 5ю итерациями, мы извлекаем результат из канала `results`. Так как операция чтения на пустом буфере является блокируемой, планировщик запустит горутину из пула воркеров. До тех пор, пока горутина не вернет результат, `main` будет заблокирован.
7.  Поскольку мы симулируем операцию блокировки в горутине, это приведет к вызову планировщиком другой доступной горутины для запуска. Когда горутина запустится, она запишет результат в канал `results`, а так как операция записи в канал с буфером является неблокируемым до тех пор, пока буфер не заполнен, блокировки при записи не произойдет. Таким образом как только одна из горутин завершится, запустятся другие горутины и считают данные из канала `tasks`. После того, как все горутины считают данные из `tasks`, цикл `for` завершится, а канал `tasks` будет пустым. Так же не произойдет ошибка `deadlock`, так как канал `tasks` был закрыт.
8.  Иногда все воркеры могут находиться в режиме ожидания, поэтому `main` программа будет работать до тех пор, пока канал `results` не будет пуст.
9.  После того, как все воркеры отработают, `main` восстановит контроль, выведет оставшиеся результаты из канала `results`, и продолжит выполнение.

Приведенный пример достаточно большой, но прекрасно объясняет, как несколько горутин могут извлекать данные из канала и выполнять свою работу. Горутины весьма эффективны, когда они могут блокироваться. Если убрать вызов `time.Sleep()`, то только одна горутина будет выполняться, так как другие горутины не будут запланированы, до тех пор пока цикл не закончится и горутина не завершится.

> Вы можете получить другой результат в приведенном примере, в зависимости от скорости работы вашей системы.

Давайте воспользуемся концепцией `WaitGroup` для синхронизации горутин. Используя предыдущий пример с `WaitGroup`, мы можем получить те же результаты, но более элегантно.

```
package main

import (
    "fmt"
    "sync"
    "time"
)

// worker than make squares
func sqrWorker(wg *sync.WaitGroup, tasks <-chan int, results chan<- int, instance int) {
    for num := range tasks {
        time.Sleep(time.Millisecond)
        fmt.Printf("[worker %v] Sending result by worker %v\n", instance, instance)
        results <- num * num
    }

    // done with worker
    wg.Done()
}

func main() {
    fmt.Println("[main] main() started")

    var wg sync.WaitGroup

    tasks := make(chan int, 10)
    results := make(chan int, 10)

    // launching 3 worker goroutines
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go sqrWorker(&wg, tasks, results, i)
    }

    // passing 5 tasks
    for i := 0; i < 5; i++ {
        tasks <- i * 2 // non-blocking as buffer capacity is 10
    }

    fmt.Println("[main] Wrote 5 tasks")

    // closing tasks
    close(tasks)

    // wait until all workers done their job
    wg.Wait()

    // receving results from all workers
    for i := 0; i < 5; i++ {
        result := <-results // non-blocking because buffer is non-empty
        fmt.Println("[main] Result", i, ":", result)
    }

    fmt.Println("[main] main() stopped")
}
```

[Пример в play.golang.org](https://play.golang.org/p/0rRfchn7sL1)

Результат работы программы:

```
[main] main() started
[main] Wrote 5 tasks
[worker 0] Sending result by worker 0
[worker 2] Sending result by worker 2
[worker 1] Sending result by worker 1
[worker 2] Sending result by worker 2
[worker 0] Sending result by worker 0
[main] Result 0 : 4
[main] Result 1 : 0
[main] Result 2 : 16
[main] Result 3 : 64
[main] Result 4 : 36
[main] main() stopped
```

В приведенном результате мы получили немного другой, более аккуратный вывод, потому что операция чтения из канала `results` в `main` не блокируется, так как канал `results` уже содержит данные из-за вызванного ранее `wg.Wait()`. Используя `WaitGroup`, мы можем предотвратить много (ненужных) переключений контекста (планирование горутин и их запуск), в данном случае 7 против 9 в предыдущем примере. Но при этом вам приходится ожидать завершения всех горутин.

### Мьютекс

Мьютекс — это один из самых простых концепций в Go. Но прежде чем разобраться в нем, давайте для начала разберемся в понятии `race condition`([состоянии гонки](https://ru.wikipedia.org/wiki/%D0%A1%D0%BE%D1%81%D1%82%D0%BE%D1%8F%D0%BD%D0%B8%D0%B5_%D0%B3%D0%BE%D0%BD%D0%BA%D0%B8)). Горутины имеют независимый стек, следовательно нет необходимости в обмене данными между ними. Но, иногда, необходимо использовать общие данные между несколькими горутинами. В этом случае несколько горутин пытаются взаимодействовать с данными в общей области памяти, что иногда приводит к непредсказуемому результату. Рассмотрим простой пример:

```
package main

import (
    "fmt"
    "sync"
)

var i int // i == 0

// goroutine increment global variable i
func worker(wg *sync.WaitGroup) {
    i = i + 1
    wg.Done()
}

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go worker(&wg)
    }

    // wait until all 1000 goroutines are done
    wg.Wait()

    // value of i should be 1000
    fmt.Println("value of i after 1000 operations is", i)
}
```

[Пример в play.golang.org](https://play.golang.org/p/MQNepChxiEa)

В приведенной программе мы порождаем 1000 горутин, которые увеличивают значение глобальной переменной `i`, равной изначально 0. Мы написали программу с использованием `WaitGroup`, поскольку мы хотим, чтобы все 1000 горутин увеличивали значение `i` последовательно, и, в результате, итоговое значение было равно 1000. Когда `main` восстанавливается после вызова `wg.Wait()`, мы выводим значение `i`. Давайте посмотрим на конечный результат:

```
value of i after 1000 operations is 937
```

Что? Почему мы получили значение меньше 1000? Возможно часть горутин не отработала. Но, в действительности, произошло `race condition`. Посмотрим, как это могло случиться.

Вычисление `i = i + 1` состоит из трех шагов:

1.  Получить значение `i`
2.  Увеличить на 1
3.  Обновить значение `i` с новым значением

Давайте представим следующий сценарий, в котором между этими шагами были запланированы разные горутины. К примеру, рассмотрим 2 горутины из пула 1000 горутин, а именно. G1 и G2.

G1 запускается, когда `i` равна 0, после второго шага `i` стала равной 1. Но перед тем, как G1 изменит значение `i` на 1 в шаге 3, новая горутина G2 уже была запланирована, и эта горутина выполнит те же шаги. В случае G2, значение `i` все еще 0, поэтому на третьем шаге значение `i` будет равно 1, в это время G1 собирается закончить третий шаг и изменить значение `i` на 1. В идеальном мире, где горутины планируются после выполнения всех 3-х шагов, успешное выполнение 2х горутин привело бы к значению `i` равному 2, но это не так. Поэтому, мы можем предположить, почему наша программа не выдает значение `i` равным 1000.

Как мы знаем, горутины планируются совместно и до тех пор, пока горутина не заблокируется по одному из условий, другая горутина не будет запланирована. Но операция `i = i + 1` не является блокируемой, тогда почему планировщик Go планирует другие горутины?

Вы можете посмотреть ответ на [stackoverflow](https://stackoverflow.com/questions/37469995/goroutines-are-cooperatively-scheduled-does-that-mean-that-goroutines-that-don). В любом случае не следует полагаться на алгоритм планирования Go и реализовывать собственную логику для синхронизации различных программ.

Один из способов удостовериться, что горутина выполнит все 3 вышеуказанных шага за раз, это использовать мьютекс. [Мьютекс](https://ru.wikipedia.org/wiki/%D0%9C%D1%8C%D1%8E%D1%82%D0%B5%D0%BA%D1%81) — это концепция в программировании, где только один поток может выполнять несколько операций одновременно. Это делается с помощью подпрограммы, получающей блокировку для выполнения любых манипуляции со значением, которое она должна изменить, а затем снимает блокировку после. Когда значение заблокировано, никакая другая подпрограмма не может читать или записывать его.

В Go мьютексы — это структура данных, которую предоставляет пакет `sync`. В Go перед выполнением любой операции со значением, которое может вызвать `race condition`, мы получаем эксклюзивную блокировку, используя метод `mutex.Lock()`. Как только мы выполнили операцию `i = i + 1` в ранее написанной программе, мы снимаем блокировки, используя метод `mutext.Unlock()`. Когда любая другая горутина попытается прочитать или записать значение `i` при наличии блокировки, эта программа будет блокироваться до тех пор, пока мьютекс не будет разблокирован. И горутина сможет безопасно читать и писать данные в переменную `i`. Запомните, что любые переменные, находящиеся между `Lock` и `Unlock`, будут недоступны для других горутин до тех пор, пока не выполнится операция снятия блокировки.

Давайте изменим предыдущий пример, используя мьютекс.

```
package main

import (
    "fmt"
    "sync"
)

var i int // i == 0

// goroutine increment global variable i
func worker(wg *sync.WaitGroup, m *sync.Mutex) {
    m.Lock() // acquire lock
    i = i + 1
    m.Unlock() // release lock
    wg.Done()
}

func main() {
    var wg sync.WaitGroup
    var m sync.Mutex

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go worker(&wg, &m)
    }

    // wait until all 1000 goroutines are done
    wg.Wait()

    // value of i should be 1000
    fmt.Println("value of i after 1000 operations is", i)
}
```

В данной программе мы создали мьютекс и передали его указатель во все горутины, прежде чем выполнить операцию с переменной `i`, мы получили эксклюзивную блокировку, используя `m.Lock()`, а после операций с переменной `i` мы сняли блокировку, используя `m.Unlock()`. Таким образом мы получим следующий результат:

```
value of i after 1000 operations is 1000
```

Из приведенного результата видно, что мьютекс помог нам разрешить `race condition`. Но старайтесь избегать использования общих ресурсов между горутинами.

> Вы можете проверить программу на `race condition` в Go, используя флаг `race`, при запуске программы. `go run -race program.go`. Более подробно об этом можно прочитать [здесь](https://blog.golang.org/race-detector).

## Паттерны конкурентного программирования

Существует множество способов, с помощью которых параллелизм делает нашу повседневную жизнь проще. Ниже приведены несколько концепций и методологий, с помощью которых мы можем сделать программы быстрее и надежнее.

### Генератор

Используя каналы, мы можем достаточно просто реализовать генератор. Так как вычисления в генераторе могут являться вычислительно дорогими, то мы могли бы сделать генерацию данных конкурентно. Таким образом, программе не нужно ждать, пока все данные будут сгенерированы. Например, генерация ряда Фибоначчи.

```
package main

import "fmt"

// fib returns a channel which transports fibonacci numbers
func fib(length int) <-chan int {
    // make buffered channel
    c := make(chan int, length)

    // run generation concurrently
    go func() {
        for i, j := 0, 1; i < length; i, j = i+j, i {
            c <- i
        }
        close(c)
    }()

    // return channel
    return c
}

func main() {
    // read 10 fibonacci numbers from channel returned by `fib` function
    for fn := range fib(10) {
        fmt.Println("Current fibonacci number is", fn)
    }
}
```

```
Current fibonacci number is 0
Current fibonacci number is 1
Current fibonacci number is 1
Current fibonacci number is 2
Current fibonacci number is 3
Current fibonacci number is 5
Current fibonacci number is 8
```

Используя функцию fib, мы получаем канал, который мы можем использовать в цикле. Находясь внутри функции fib, мы создаем и возвращаем канал только для приема. Возвращаемый канал преобразуется из двунаправленного канала в однонаправленный канал для приема. Используя анонимную горутину, мы помещаем числа Фибоначчи в этот канал. Как только мы закончили с циклом `for`, мы закрываем канал внутри анонимной горутины. В `main`, используя `range`, мы итерируем данные канала, полученные после вызова функции `fib`.

### Fan-in и Fan-out

Fan-in — это стратегия мультиплексирования, при которой входы нескольких каналов объединяются в один выходной канал. Fan-out — это обратная операция, при которой один канал разделяется на несколько каналов.

```

package main

import (
    "fmt"
    "sync"
)

// return channel for input numbers
func getInputChan() <-chan int {
    // make return channel
    input := make(chan int, 100)

    // sample numbers
    numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

    // run goroutine
    go func() {
        for num := range numbers {
            input <- num
        }
        // close channel once all numbers are sent to channel
        close(input)
    }()

    return input
}

// returns a channel which returns square of numbers
func getSquareChan(input <-chan int) <-chan int {
    // make return channel
    output := make(chan int, 100)

    // run goroutine
    go func() {
        // push squares until input channel closes
        for num := range input {
            output <- num * num
        }

        // close output channel once for loop finishes
        close(output)
    }()

    return output
}

// returns a merged channel of `outputsChan` channels
// this produce fan-in channel
// this is variadic function
func merge(outputsChan ...<-chan int) <-chan int {
    // create a WaitGroup
    var wg sync.WaitGroup

    // make return channel
    merged := make(chan int, 100)

    // increase counter to number of channels `len(outputsChan)`
    // as we will spawn number of goroutines equal to number of channels received to merge
    wg.Add(len(outputsChan))

    // function that accept a channel (which sends square numbers)
    // to push numbers to merged channel
    output := func(sc <-chan int) {
        // run until channel (square numbers sender) closes
        for sqr := range sc {
            merged <- sqr
        }
        // once channel (square numbers sender) closes,
        // call `Done` on `WaitGroup` to decrement counter
        wg.Done()
    }

    // run above `output` function as groutines, `n` number of times
    // where n is equal to number of channels received as argument the function
    // here we are using `for range` loop on `outputsChan` hence no need to manually tell `n`
    for _, optChan := range outputsChan {
        go output(optChan)
    }

    // run goroutine to close merged channel once done
    go func() {
        // wait until WaitGroup finishes
        wg.Wait()
        close(merged)
    }()

    return merged
}

func main() {
    // step 1: get input numbers channel
    // by calling `getInputChan` function, it runs a goroutine which sends number to returned channel
    chanInputNums := getInputChan()

    // step 2: `fan-out` square operations to multiple goroutines
    // this can be done by calling `getSquareChan` function multiple times where individual function call returns a channel which sends square of numbers provided by `chanInputNums` channel
    // `getSquareChan` function runs goroutines internally where squaring operation is ran concurrently
    chanOptSqr1 := getSquareChan(chanInputNums)
    chanOptSqr2 := getSquareChan(chanInputNums)

    // step 3: fan-in (combine) `chanOptSqr1` and `chanOptSqr2` output to merged channel
    // this is achieved by calling `merge` function which takes multiple channels as arguments
    // and using `WaitGroup` and multiple goroutines to receive square number, we can send square numbers
    // to `merged` channel and close it
    chanMergedSqr := merge(chanOptSqr1, chanOptSqr2)

    // step 4: let's sum all the squares from 0 to 9 which should be about `285`
    // this is done by using `for range` loop on `chanMergedSqr`
    sqrSum := 0

    // run until `chanMergedSqr` or merged channel closes
    // that happens in `merge` function when all goroutines pushing to merged channel finishes
    // check line no. 86 and 87
    for num := range chanMergedSqr {
        sqrSum += num
    }

    // step 5: print sum when above `for loop` is done executing which is after `chanMergedSqr` channel closes
    fmt.Println("Sum of squares between 0-9 is", sqrSum)
}
```

Пройдем по шагам.

1.  Получаем канал `chanInputNums`, посредством вызова функции `getInputChan`. Функция `getInputChan` создает канал и возвращает его как канал, доступный только для чтения, а также запускает анонимную горутину, которая последовательно помещает в канал числа из массива `numbers` и закрывает канал.
2.  Разделяем наш канал (fan-out) на два канала(`chanOptSqr1` и `chanOptSqr2`), передавая его два раза функции `getSquareChan`. Функция `getSquareChan` создает канал и возвращает его как канал, доступный только для чтения, а также запускает анонимную горутину для вычисления квадрата чисел на основе данных канала, полученного в качестве аргумента функции.
3.  Собираем данные из каналов в один (fan-in), используя функцию `merge`. В функции `merge` мы создаем `WaitGroup`, а также новый канал(`merged`), где мы объединим все данные из списка каналов `outputsChan`, после, мы увеличиваем счетчик на основании числа полученных каналов, подготавливаем анонимную функцию для чтения данных из канала и группировки данных в наш новый канал `merged`, а также уменьшим значение счетчика, когда все данные из переданного канала будут считаны. Вызываем нашу анонимную функцию для каждого канала в качестве горутины. А так же создаем и стартуем еще одну анонимную горутину для того, чтобы дождаться выполнения операции объединения всех данных в один канал и после этого закрываем канал в рамках анонимной функции. После чего возвращаем наш новый канал `merged`.
4.  Считываем данные из канала `chanMergedSqr` используя `for` и `range`, и суммируем полученные данные.
5.  В конце выводим наш результат.

Вывод программы:

```
Sum of squares between 0-9 is 285
```
