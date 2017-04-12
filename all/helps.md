## Подсказки по Golang
### Подсказки по основам языка программирования Go.

Комментарии
// Package path implements utility
// routines for manipulating
// slash-separated filename paths.

/*
Package path implements utility
routines for manipulating
slash-separated filename paths.
*/
Именования
// Первый пакет.
package myname

// Экспортируется
type Buffer []byte

// Экспортируется
func New() Buffer { ... }

// Второй пакет.
package main

import "myname"

// Не экспортируется
func sample() {
  var b myname.Buffer = myname.New()
  ...
}
Константы
const small = 1
const huge = 100

// Или
const (
  red = iota // red == 0
  blue       // blue == 1
  green      // green == 2
)

## Переменные
```golang
// Полная декларация с инициализацией.
var v int = 0
// Полная декларация без инициализации.
var v int
// Вывод типа при инициализации.
var v = 0
// Тоже самое, но еще короче.
v := 0
Точки с запятой
// Как правило, не нужны
f, err := os.Open(name)
if err != nil {
    return err
}
codeUsing(f)

// Но иногда нужны.
if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
}
if
if x > 0 {
    return y
}

if err := file.Chmod(0664); err != nil {
    log.Stderr(err)
    return err
}
for
// Как for в C
for init; condition; post { }

// Как while в C
for condition { }

// Как бесконечный цикл (for(;;)
// или while(1)) в C
for { }

sum := 0
for i := 0; i < 10; i++ {
    sum += i
}

// range
for key, value := range oldMap {
    newMap[key] = value
}

var m map[string]int
sum := 0
// Ключ из map-а игнорируется.
for _, value := range m {
    sum += value
}
Switch
switch {
case '0' <= c && c <= '9':
    return c - '0'
case 'a' <= c && c <= 'f':
    return c - 'a' + 10
case 'A' <= c && c <= 'F':
    return c - 'A' + 10
}

switch c {
case ' ', '?', '&', '=', '#', '+', '%':
    return true
}
```

## Массивы и слайсы
```golang
// Массив
var array [3]float32 =
    [3]float32{7.0, 8.5, 9.1}
array = [...]float32{7.0, 8.5, 9.1}
array = [3]float32{7.0, 8.5, 9.1}

// Слайс
var slice []int = []int{1,2,3}
slice = make([]int, 3) // [0, 0, 0]
slice = append(slice, 1) // [0, 0, 0, 1]
Хэш-таблицы (map-ы)
var timeZone = map[string] int {
  "UTC":  0*60*60,
  "EST": -5*60*60,
  "CST": -6*60*60,
  "MST": -7*60*60,
  "PST": -8*60*60,
}

// Поиск
if seconds, ok := timeZone[tz]; ok {
  return seconds
}
```

## Функции
```golang
// В Go функция может возвращать
// несколько значений:
func idiv(a, b int) (int, int) {
  d := a / b
  r := a % b
  return d, r
}

// Возвращаемым значениям
// можно присваивать имена:
func idiv(a, b int) (d, r int) {
  d = a / b
  r = a % b
  return
}

// Возврат кодов ошибок
func (file *File) Write(b []byte) (n int,
    err Error)
```

## Методы
```golang
type myType struct { i int }
func (p *myType) get() int { return p.i }

var m myType
i := m.get()
new() и make()
Cлайсы, хэш-таблицы и каналы нужно создавать через make. Все остальное - через new
// new(T) выделяет память для объекта
// типа T и инициализирует ее нулями.
// Возвращается значение типа *T
type SyncedBuffer struct {
  lock        sync.Mutex
  buffer      bytes.Buffer
}

p := new(SyncedBuffer)

// make() используется для создания в
// динамической памяти специальных
// объектов. К таким объектам
// относятся слайсы, хэш-таблицы (map-ы)
// и каналы
var v  []int = make([]int, 100)
v := make([]int, 100)
Указатели vs Значения
type ByteSlice []byte

func (slice ByteSlice) Append(
  data []byte) []slice {
  // Точно такое же тело,
  // как показано ранее.
}

func (p *ByteSlice) Append(data []byte) {
  slice := *p
  // Тело такое же, как и раньше,
  // но без return.
  *p = slice
}
Интерфейсы
// Объявляем интерфейс
type myInterface interface {
  get() int
  set(i int)
}

// Определяем интерфейс
type myType struct { i int }
func (p *myType) get() int { 
  return p.i
}

func (p *myType) set(i int) { 
  p.i = i
}

// Используем интерфейс
func getAndSet(x myInterface) {}
func f1() {
  var p myType
  getAndSet(&p)
}
```

## Преобразования
```golang
func (s Sequence) String() string {
  sort.Sort(s)
  return fmt.Sprint([]int(s))
}
Композиция и Агрегация
// Есть 2 интерфейса
type Reader interface {
  Read(p []byte) (n int, err error)
}

type Writer interface {
  Write(p []byte) (n int, err error)
}

// Интерфейс ReadWrite группирует
// методы из базовых
// интерфейсов Read и Write.
type ReadWriter interface {
  Reader
  Writer
}

// Композиция
type ReadWriter struct {
  *Reader
  *Writer
}
var rw ReadWriter
rw.Read( ... )

// Агрегация
type ReadWriter struct {
  reader *Reader
  writer *Writer
}
var rw ReadWriter
rw.reader.Read( ... )
```

## Goroutines

```golang
// Запустить list.Sort в другом потоке;
// не ждать его завершения.
go list.Sort()

func Announce(message string,
                delay int64) {
  go func() {
    time.Sleep(delay)
    fmt.Println(message)
  }()  // Круглые скобки здесь важны
       // -- делается вызов функции.
}
```

## Каналы

```golang
// небуферизированный канал целых чисел
ci := make(chan int)
// небуферизированный канал целых чисел
cj := make(chan int, 0)
// буферизированный канал
// указателей на файлы.
cs := make(chan *os.File, 100)

c := make(chan int)  // Создание канала.
// Сортировка запускается в
// отдельном потоке.
// Когда она завершается в
// канал записывается сигнал.
go func() {
    list.Sort();
    c <- 1   // Отсылка сигнала.
             // Не важно, какое это
             // будет значение.
}()
doSomethingForAWhile()
<-c   // Ожидание завершения сортировки.
```

## Распараллеливание
```golang
type Vector []float64

// Выполнение операции над n
// элементами вектора v,
// начиная с позиции i.
func (v Vector) DoSome(i, n int,
                u Vector, c chan int) {
  for ; i < n; i++ {
      v[i] += u.Op(v[i])
  }
  // сигнализация завершения
  // обработки элементов.
  c <- 1
}

// Количество доступных CPU.
const NCPU = 4

func (v Vector) DoAll(u Vector) {
  // Буферизация не обязательна,
  // но может иметь смысл.
  c := make(chan int, NCPU)

  for i := 0; i < NCPU; i++ {
    go v.DoSome(i*len(v)/NCPU,
            (i+1)*len(v)/NCPU, u, c)
  }
  // Вычитываем все из канала.
  for i := 0; i < NCPU; i++ {
    <-c  // Ждем, пока завершиться
         // очередная операция.
  }
  // Все действия завершены.
}
```


## Ошибки
```golang
// Стандартная ошибка
type Error interface {
    Error() string;
}

// Использование
import "errors"
errors.New("Some problem")

// Пример
type ErrCustomError struct {
    errStr string
}

func (e *ErrCustomError) Error() string {
    return e.errStr
}

&ErrCustomError{"error ='("}

// По-быстрому инстанцировать
// новую ошибку
fmt.Errorf("%s", "Значение ошибки")
```
