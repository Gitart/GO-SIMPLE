# Программирование в Go
## Часть 2

Когда мы не ставим точку с запятой в конце каждой строки, компилятор автоматически добавляет ее сам. Вот почему следующий код не будет компилироваться:
```golang
 for i := 0; i < 5; i++
 {
   fmt.Println(i)
 }
 ```
Правильный вариант:
```golang
 for i := 0; i < 5; i++ {
   fmt.Println(i)
 }
 ```
В go есть три оператора ветвления:
```
  if
  switch
  select
  ```
Два примера оператора if - они идентичны, с той лишь разницей, что во втором случае мы добавляем лишние фигурные скобки, чтобы ограничить локальную переменную:
```golang
 if α := compute(); α < 0 {
   fmt.Printf("(%d)\n", -α)
 } else {
   fmt.Println(α)
 }
 
 
 {
   α := compute()
   if α < 0 {
     fmt.Printf("(%d)\n", -α)
   } else {
     fmt.Println(α)
   }
 }
 
Оператор switch имеет 2 варианта - expression switch и type switch. Первый имеет общепринятый формат, второй специфичен для go. Пример:

 ### Пример 1:
 ```golang
 func BoundedInt(minimum, value, maximum int) int {
   switch {
   case value < minimum:
     return minimum
   case value > maximum:
     return maximum
   }
   return value
 }
 ```
 
 Пример 2:
 ```golang
 switch Suffix(file) { // Canonical ✓
 case ".gz":
   return GzipFileList(file)
 case ".tar", ".tar.gz", ".tgz":
   return TarFileList(file)
 case ".zip":
   return ZipFileList(file)
 }
 ```
 
Пример type switch:

```golang
 func classifier(items ...interface{}) {
   for i, x := range items {
     switch x.(type) {
     case bool:
       fmt.Printf("param #%d is a bool\n", i)
     case float64:
       fmt.Printf("param #%d is a float64\n", i)
     case int, int8, int16, int32, int64:
       fmt.Printf("param #%d is an int\n", i)
     case uint, uint8, uint16, uint32, uint64:
       fmt.Printf("param #%d is an unsigned int\n", i)
     case nil:
       fmt.Printf("param #%d is nil\n", i)
     case string:
       fmt.Printf("param #%d is a string\n", i)
     default:
       fmt.Printf("param #%d's type is unknown\n", i)
     }
    }
   }
   ```
Go использует два варианта цикла for - стандартный и for ... range:
```golang
 for { // Infinite loop
   block
 }
 
 for booleanExpression { // While loop
   block
 }
 
 for optionalPreStatement; booleanExpression; optionalPostStatement { ➊
   block
 }
 
 for index, char := range aString { // String per character iteration ➋
   block
 }
 
 for index := range aString { // String per character iteration ➌
   block // char, size := utf8.DecodeRuneInString(aString[index:])
 }
 
 for index, item := range anArrayOrSlice { // Array or slice iteration ➍
   block
 }
 
 for index := range anArrayOrSlice { // Array or slice iteration ➎
   block // item := anArrayOrSlice[index]
 }
 
 for key, value := range aMap { // Map iteration ➏
   block
 }
 
 for key := range aMap { // Map iteration ➐
   block // value := aMap[key]
 }
 
 for item := range aChannel { // Channel iteration
   block
 }
 ```
 
Следующие два примера идентичны, второй короче и использует метку:
Пример 1:

 ```golang 
  found := false
  for row := range table {
   for column := range table[row] {
     if table[row][column] == x {
       found = true
       break
     }
  }
  if found {
     break
  }
 }
```

 Пример 2:
 ```golang
  found := false
  FOUND:
  for row := range table {
   for column := range table[row] {
     if table[row][column] == x {
       found = true
       break FOUND
     }
   }
  }
  ```
  
goroutine - это функция или метод, которая выполняется параллельно с программой. Это не поток в обычном понимании, это нечто более легкое, что потребляет еще меньше ресурсов. channel - двунаправленный канал, который может быть использован для обмена данными между goroutine. Они не используют глобальных данных и нет надобности в блокировках. Разница между созданием потока в си и goroutine следующая: каждый раз, когда вы создаете поток в си, операционная система выделяет под него в стеке кусок памяти в один метр (грубо говоря), независимо от того, что потоку надо всего несколько байт, и кроме него, этой памятью никто не может распоряжаться, в результате память начинает сегментироваться. Для создания же goroutine требуется всего несколько килобайт. Когда вы создаете большое количество goroutine, это не означает, что при этом операционная система будет создавать такое же количество потоков - последних будет меньше. Переключение контекста между goroutine происходит быстрее, нежели между сишными потоками. Создаются goroutine так:
```golang
 go function(arguments)
 go func(parameters) { block }(arguments)
 ```
 
Как только код доходит до этого места, параллельно с главной программой main тут же одновременно запускается goroutine. Интерфейс для обмена данными:

```golang
 channel <- value     //  посылка данных
 <-channel            //  получение данных
 x := <-channel       //  получение и сохранение
 x, ok := <-channel   //  получение с проверкой 
 ```
 
Канал создается с помощью функции make():
```golang
 make(chan Type)
 make(chan Type, capacity)
 ```
 
В следующем примере создаются два канала. Создаются две goroutine, каждая работает со своим каналом. Из каждого канала мы должны получить по 5 чисел:
```golang
 counterA := createCounter(2)     // counterA is of type chan int
 counterB := createCounter(102)   // counterB is of type chan int
 for i := 0; i < 5; i++ {
   a := <-counterA
   fmt.Printf("(A→%d, B→%d) ", a, <-counterB)
 }
 fmt.Println()
 
 func createCounter(start int) chan int {
   next := make(chan int)
   go func(i int) {
     for {
 	next <- i
 	i++
     }
   }(start)
   return next
 }
 ```
 Вывод:
 ```
 (A→2, B→102) (A→3, B→103) (A→4, B→104) (A→5, B→105) (A→6, B→106)
 ```
 
Оператор select имеет следующий синтаксис:
```golang
 select {
 case sendOrReceive1: block1
 ...
 case sendOrReceiveN: blockN
 default: blockD
 }
 ```
 
Селект с default называется неблокирующим селектом. В следующем примере создаются 6 каналов, которые пересылают булевское значение. Создается одна goroutine, которая рандомно выбирает канал и посылает через него true. Канал не имеет буфера, поэтому goroutine после отсылки сразу блокируется. Селект не имеет дефалтового значения, поэтому он также блокируется:

```golang
 channels := make([]chan bool, 6)
 for i := range channels {
   channels[i] = make(chan bool)
 }
 go func() {
   for {
     channels[rand.Intn(6)] <- true
   }
 }()
 
 for i := 0; i < 36; i++ {
   var x int
   select {
   case <-channels[0]:
     x = 1
   case <-channels[1]:
     x = 2
   case <-channels[2]:
     x = 3
   case <-channels[3]:
     x = 4
   case <-channels[4]:
     x = 5
   case <-channels[5]:
     x = 6
   }
   fmt.Printf("%d ", x)
 }
 fmt.Println()
```

 Вывод:
 6 4 6 5 4 1 2 1 2 1 5 5 4 6 2 3 6 5 1 5 4 4 3 2 3 3 3 5 3 6 5 2 2 3 6 2

defer - отложенная функция, которая выполняется после того, как закончит выполняться внешняя функция. Их может быть более одной, и выполняться они будут по правилу LIFO. Они часто выполняют такие действия, как проверка закрытия файла или канала, или обрабатывают ошибки:
```golang
 var file *os.File
 var err error
 if file, err = os.Open(filename); err != nil {
   log.Println("failed to open the file: ", err)
   return
 }
 defer file.Close()
 ```
 
Обработка ошибок в go включает использование функций panic() и recover() . В программе бывают ошибки и бывают исключения, в первом случае имеется ввиду ситуация, когда файл есть, а открыть его нельзя - в этом случае мы можем вызвать функцию panic() - в других языках ее аналогом является assert. Когда вызывается panic(), функция, в которой это происходит, немедленно останавливается. После чего срабатывает отложенная функция. Если в отложенной функции есть вызов recover(), то дальше программа работает так, как будто ничего не произошло. Вообще функцией panic() лучше не пользоваться :-)
В go variadic функция - это функция, которая имеет произвольное число параметров, для этого используется многоточие:
```golang

 func MinimumInt1(first int, rest ...int) int {
   for _, x := range rest {
     if x < first {
       first = x
     }
   }
   return first
 }
 
 fmt.Println(MinimumInt1(5, 3), MinimumInt1(7, 3, -2, 4, 0, -8, -5))
 ```
 
 Вывод:
 ```
 3 -8
 ```
 
Функции-дженерики реализованы с помощью интерфейсов. Следующий пример показывает, как можно передать в такую функцию-дженерик слайс произвольного типа:
```golang
 func Index(xs interface{}, x interface{}) int {
   switch slice := xs.(type) {
   case []int:
     for i, y := range slice {
       if y == x.(int) {
 	return i
       }
     }
   case []string:
     for i, y := range slice {
       if y == x.(string) {
 	return i
       }
     }
   }
   return -1
 }
 
 xs := []int{2, 4, 6, 8}
 fmt.Println("5 @", Index(xs, 5), " 6 @", Index(xs, 6))
 ys := []string{"C", "B", "K", "A"}
 fmt.Println("Z @", Index(ys, "Z"), " A @", Index(ys, "A"))
 ```
 
 Вывод:
 ```
 5 @ -1    Z @ -1
 6 @ 2     A @ 3
 ```
 
В go higher order function - это функция, параметром которой может быть другая функция. Пример, в котором SliceIndex() является дженерик-функцией - она возвращает индекс элемента в слайсе, если этот элемент отвечает условию, проверяемому в анонимной функции:
```golang
 func SliceIndex(limit int, predicate func(i int) bool) int {
   for i := 0; i < limit; i++ {
     if predicate(i) {
       return i
     }
   }
   return -1
 }
 
 xs := []int{2, 4, 6, 8}
 ys := []string{"C", "B", "K", "A"}
 fmt.Println(SliceIndex(len(xs), func(i int) bool { return xs[i] == 5 }),
 	    SliceIndex(len(xs), func(i int) bool { return xs[i] == 6 }),
 	    SliceIndex(len(ys), func(i int) bool { return ys[i] == "Z" }),
 	    SliceIndex(len(ys), func(i int) bool { return ys[i] == "A" }))
 ```
 Вывод: 
 ```
  -1 2 -1 3
  ```
Следующий аналогичный пример фильтрует слайс:

```golang
 func IntFilter(slice []int, predicate func(int) bool) []int {
   filtered := make([]int, 0, len(slice))
     for i := 0; i < len(slice); i++ {
       if predicate(slice[i]) {
 	filtered = append(filtered, slice[i])
       }
     }
     return filtered
 }
 
 readings := []int{4, -3, 2, -7, 8, 19, -11, 7, 18, -6}
 even := IntFilter(readings, func(i int) bool { return i%2 == 0 })
 fmt.Println(even)
```

 Вывод:
 ```
 [4 2 8 18 -6]
 ```
 
IntFilter фильтрует целые числа. Можно написать дженерик-функцию, которая будет фильтровать произвольный тип:
```golang
 func Filter(limit int, predicate func(int) bool, appender func(int)) {
   for i := 0; i < limit; i++ {
     if predicate(i) {
       appender(i)
     }
   }
 }
 
 readings := []int{4, -3, 2, -7, 8, 19, -11, 7, 18, -6}
 even := make([]int, 0, len(readings))
 Filter(len(readings), func(i int) bool { return readings[i]%2 == 0 },
   func(i int) { even = append(even, readings[i]) })
 fmt.Println(even)
 ```
 
 Вывод:
 ```
 [4 2 8 18 -6]
 ```
 ```golang
 parts := []string{"X15", "T14", "X23", "A41", "L19", "X57", "A63"}
 var Xparts []string
 Filter(len(parts), func(i int) bool { return parts[i][0] == 'X' },
   func(i int) { Xparts = append(Xparts, parts[i]) })
 fmt.Println(Xparts)
 ```
 Вывод:
 ```
 [X15 X23 X57]
 ```
