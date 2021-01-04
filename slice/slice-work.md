### Спецификация Go: добавление в срезы и копирование срезов

Встроенные функции **append** и **copy** помогают в общих операциях срезов. Для обеих функций результат не зависит от того, перекрывается ли память, на которую ссылаются аргументы.

[variadic](https://golang-blog.blogspot.com/2019/06/go-specification-variadic-parameters.html) функция **append** добавляет ноль или более значений **x** к **s** типа **S**, который должен быть типом среза, и возвращает полученный срез, также типа S. Значения x передаются параметру типа ... T, где T тип элемента S и применяются соответствующие правила передачи параметров. В особом случае append также принимает первый аргумент, присваиваемый (assignable) типу `[]byte`, со вторым аргументом строкового типа, за которым следует .... В этой форме добавляются байты строки.

```
append(s S, x ...T) S  // T это элемент типа S

```

Если емкость s недостаточно велика, чтобы соответствовать дополнительным значениям, append выделяет новый, достаточно большой базовый массив, который подходит как существующим элементам среза, так и дополнительным значениям. В противном случае append повторно использует базовый массив.

```
s0 := []int{0, 0}
s1 := append(s0, 2)                // добавляем отдельный элемент     s1 == []int{0, 0, 2}
s2 := append(s1, 3, 5, 7)          // добавляем несколько элементов   s2 == []int{0, 0, 2, 3, 5, 7}
s3 := append(s2, s0...)            // добавляем срез                  s3 == []int{0, 0, 2, 3, 5, 7, 0, 0}
s4 := append(s3[3:6], s3[2:]...)   // добавляем перекрывающийся срез  s4 == []int{3, 5, 7, 2, 3, 5, 7, 0, 0}

var t []interface{}
t = append(t, 42, 3.1415, "foo")   // t == []interface{}{42, 3.1415, "foo"}

var b []byte
b = append(b, "bar"...)            // добавляем содержимое строки      b == []byte{'b', 'a', 'r' }

```

Функция **copy** копирует элементы среза из исходного src в dst назначения и возвращает количество скопированных элементов. Оба аргумента должны иметь одинаковый тип элемента T и должны быть присваиваемыми срезу типа `[]T`. Количество копируемых элементов является минимумом len(src) и len(dst). В особом случае copy также принимает целевой аргумент, назначаемый типу `[]byte` с исходным аргументом строкового типа. Эта форма копирует байты из строки в срез байтов.

```
copy(dst, src []T) int
copy(dst []byte, src string) int

```

Примеры:

```
var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
var s = make([]int, 6)
var b = make([]byte, 5)
n1 := copy(s, a[0:])            // n1 == 6, s == []int{0, 1, 2, 3, 4, 5}
n2 := copy(s, s[2:])            // n2 == 4, s == []int{2, 3, 4, 5, 4, 5}
n3 := copy(b, "Hello, World!")  // n3 == 5, b == []byte("Hell
```