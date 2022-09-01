# Нарезаем массивы правильно в Go

[Программирование \*](https://habr.com/ru/hub/programming/)[Go \*](https://habr.com/ru/hub/go/)

![](https://habrastorage.org/r/w1560/getpro/habr/upload_files/3c6/ec0/919/3c6ec0919c6d6174feb07704e4dca5b1.png)

Второй очерк из цикла приключений в мире сусликов.

Это вторая статья серии небольших рассказов о необычных подводных камнях, которые можно встретить в начале разработки на Go. Напоминаю, что в статьях есть примеры кода, будьте с ними аккуратнее - не все из них будут компилироваться и работать, читайте внимательно комментарии, везде указано, на какой строке происходит ошибка. Также в блоках кода везде табуляция заменена на пробелы - это сделано намеренно, чтобы статьи выглядели у всех одинаково.

Начинать рассказывать снова буду издалека, с самых основ, но это необходимо для полного понимания. В конце же рассказа вас, как и прежде, будет ждать самое интересное, но всё равно стоит читать его с самого начала.

Как писал ранее, я всю жизнь занимаюсь разработкой программного обеспечения, в основном в сфере WEB, успел познакомиться с многими языками программирования и поработать в разных крупных компаниях. Сейчас руковожу разработкой в компании NUT.Tech, мы там делаем классные и интересные вещи. В данный момент в основном разработка в отделе построена вокруг Go, поэтому о нём я и решил рассказывать.

Статьи серии:

1.  [Интерфейсы в Go - как красиво выстрелить себе в ногу](https://habr.com/ru/post/597461/)

2.  Нарезаем массивы правильно в Go

3.  ...

Расскажу я сегодня об одной из базовых структур языка, некоторые особенности которой при первом знакомстве вгоняют в ступор. Речь пойдёт о срезах и о том, какие интересные “фичи” нам приносит их внутреннее устройство в языке. Но начнем мы издалека - с массивов.

### Массивы

Для начала посмотрим, что язык Go даёт нам для работы со структурами данных, известные в других языках как списки, массивы, векторы и тому подобное.

Под массивом в Go обычно понимается структура данных фиксированного размера, хранящая элементы одного типа. Фиксированный размер означает то, что после создания в массив нельзя будет добавить новые элементы и количество элементов уже не может стать меньше. Размер, он же длина, как и тип элементов, должны быть известны на этапе компиляции, поэтому они задаются сразу в коде.

```
a := [3]int{1, 2, 3}
```

Что же у нас тут происходит. А происходит примерно следующее: инициализируется переменная `a` и в неё помещаются элементы типа `int` с длиной три элемента (записана в квадратных скобках). Также есть более удобная форма подобной записи.

```
a := [...]int{1, 2, 3}
```

Здесь происходит всё ровно то же самое, но три точки в скобках позволяют нам не записывать длину самим, вместо этого она будет автоматически выведена компилятором из количества элементов.

У массивов есть одна важная особенность: переменные, имеющие в качестве типа массивы разной длины не взаимозаменяемы, то есть с точки зрения Go, это переменные совершенно разных типов. Например:

```
// объявляем переменную с типом массива из двух элементов
var a [2]int

b := [1]int{1}
a = b
// ура! ошибка компиляции
// cannot use b (type [1]int) as type [2]int in assignment

c := [2]int{1, 2}
a = c
// а так ошибки нет, всё компилируется и работает
```

Посмотрим на несколько простых примеров работы с массивами:

```
// создание
a := [3]int{1, 2, 3}

// получение элемента
fmt.Println(a[0]) // 1

fmt.Println(a[4])
// ошибка времени компиляции
// invalid array index 4 (out of bounds for 3-element array)

// но есть нюанс, если задавать элемент через переменную -
// всё упадёт во время выполнения
i := 4
fmt.Println(a[i])
// panic: runtime error: index out of range [4] with length 3

// изменение элемента
a[0] = 42
fmt.Printf("%#v\n", a) // [3]int{42, 2, 3}

// одной интересной особенностью является то,
// что у массива без заданных значений
// всё равно можно получить элемент по индексу
var b [3]int
fmt.Println(b[1]) // 0

// всё дело в том, что Go сам инициализирует
// все элементы значениями по умолчанию
fmt.Printf("%#v\n", b) // [3]int{0, 0, 0}
```

В оперативной памяти же массивы выглядят просто как последовательность значений одного размера. Они занимают фиксированный объем и имеют постоянное расположение в памяти на всё время жизни, пока за ними не придет сборщик мусора.

И чтобы быть совсем уверенными в том, что разобрались в массивах, заглянем в исходный код Go, который к нашему счастью написан на Go.

```
// src/go/types/array.go

type Array struct {
    len  int64
    elem Type
}
```

Достаточно просто выглядит структура, всего два поля: длина - поле `len` (length) и тип элементов - поле `elem` (element). Тут всё кажется понятным, так что время перейти к срезам.

### Срезы

Эта структура данных обычно чуть менее привычна, чем массивы, тем не менее в Go преимущественно встретить можно именно её. Срезы очень похожи на массивы с точки зрения использования. Главным отличием от массивов является то, что срез не имеет строго фиксированного размера и количество элементов в нём может меняться в течение жизни программы. То есть, простыми словами, в него можно добавить и удалить из него элементы.

На этот раз сразу начнем с исходного кода суслика.

```
// src/runtime/slice.go

type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}
```

Тут всё немного сложнее: появилось поле `cap` (capacity), оно отражает размер (емкость) коллекции и может отличаться от `len`, которое тут фактически просто количество элементов. Ну и вместо поля `elem`, у нас поле `array`, которое хранит ссылку на определённый элемент исходного массива в памяти, с этого элемента и начинается срез.

Небольшое отступление: в русскоязычных источниках чаще можно прочитать перевод поля `cap`, как "объем" или "емкость", встречал ещё вариант "вместимость", всё это ближе к английскому, чем размер, поэтому далее в тексте я буду употреблять термин "ёмкость".

Длине среза и его ёмкости нужно уделять особое внимание. С одной стороны, Go самостоятельно полностью управляет значениями в этих полях и делает все процедуры по перераспределению занимаемой памяти, с другой стороны, иногда это ведет к интересным и неожиданным последствиям.

Для начала посмотрим на простые примеры работы со срезами:

```
// в отличии от массива срез нужно инициализировать без указания размера
a := []int{1, 2, 3}

// во время инициализации среза сначала инициализируется массив,
// там сохраняются переданные значения и сразу создаётся срез,
// указывающий на этот массив и хранящий длину и емкость

// аналогично массивам можно получать и изменять элементы
fmt.Println(a[0]) // 1
a[0] = 42
fmt.Printf("%#v\n", a) // [3]int{42, 2, 3}

// при получении элемента за границами среза
// произойдет ошибка времени выполнения
fmt.Println(a[4])
// panic: runtime error: index out of range [4] with length 3

// также срез можно создать на основе существующего массива
m := [3]string{"Шито", "Крыто", "Корыто"}
b := m[0:2]
fmt.Printf("%#v\n", b) // []string{"Шито", "Крыто"}
```

Тут есть одна особенность, которая не бросается сразу в глаза и происходит из-за того, что срез - это ссылка на исходный массив, значит, чтобы существовать этому срезу, и массив должен существовать. Поэтому, если мы определили массив (или срез) из большого количества элементов, взяли небольшой новый срез от него и дальше используем только значения нового среза, значит большое количество элементов из исходного массива (среза) будут жить в памяти, несмотря на то, что они нам больше не нужны. Таким образом, мы получаем некоторую утечку памяти, за этим нужно аккуратно следить.

С основами вроде бы разобрались, теперь посмотрим внимательнее на то, что же такое ёмкость и длина среза.

Нужно не забывать, что практически срез - это просто ссылка на массив, и каждый элемент среза - это ссылки на элементы связанного массива. По этой причине нужно быть аккуратным, когда мы в функцию передаем срез. В Go всё можно передать по значению, и срез не исключение, но значением среза является ссылка, поэтому при такой передаче функция будет менять значения исходного массива, вот так вот запутанно всё, но привыкнуть можно. Посмотрим на простой пример:

```
func main() {
    m := [...]int{1, 2, 3} // создаём массив
    s := m[:] // таким образом можно создать срез, содержащий весь массив

    // теперь нам понадобится функция, которая будет менять нулевой
    // элемент переданного среза
    func(l []int) {
            l[0] = 42
    }(s)

    // выводим исходный массив
    fmt.Printf("%#v\n", m) // [3]int{42, 2, 3}
    // первый элемент поменялся в исходном массиве,
    // несмотря на то, что срез в функцию был передан "по значению"
}
```

Благодаря этой хитрости работает пакет сортировки, который принимает на вход срез и ничего не возвращает, сортируя значения сразу в переданном срезе, без лишних выделений памяти.

Теперь пора разобраться до конца с длиной и размером. Для получения значения этих параметров есть одноименные функции `len()` и `cap()`.

```
func main() {
    a := []int{1, 2, 3}

    fmt.Println(cap(a)) // 3
    fmt.Println(len(a)) // 3

    // ёмкость и длина у нас тут равны, но что если удалить один элемент
    a = append(a[:1], a[2:]...)

    fmt.Printf("%#v\n", a) // []int{1, 3}

    fmt.Println(cap(a)) // 3
    fmt.Println(len(a)) // 2
    // ёмкость массива осталась неизменной, но длина его стала меньше

    // теперь добавим несколько элементов
    a = append(a, 10, 11, 12, 13, 14, 15)

    fmt.Println(cap(a)) // 8
    fmt.Println(len(a)) // 8
    // ёмкость и длина изменились автоматически, всё хорошо
}
```

И тут мы наконец подошли к самому интересному вопросу: “Что происходит, когда ёмкость среза увеличивается?”

А происходит следующее: создается в памяти новый массив, туда копируются все значения из среза, к ним добавляются новые значения, которые и привели к увеличению ёмкости среза, в срезе в поле `array` устанавливается ссылка на новый массив. Это можно увидеть воспользовавшись пакетом reflect:

```
func main() {
    a := []string{"Шито"}
    fmt.Println(len(a), cap(a)) // 1 1
    fmt.Printf("%#v\n", reflect.ValueOf(a).Pointer()) // 0xc00010c220

    a = append(a, "Крыто")
    fmt.Println(len(a), cap(a)) // 2 2
    fmt.Printf("%#v\n", reflect.ValueOf(a).Pointer()) // 0xc000130000

    a = append(a, "Корыто")
    fmt.Println(len(a), cap(a)) // 3 4 (внезапно)
    fmt.Printf("%#v\n", reflect.ValueOf(a).Pointer()) // 0xc00012e040

    // на самом деле ссылки будут разными при каждом запуске,
    // но тут главное можно увидеть, что после каждого добавления элемента
    // и, как следствие, увеличения емкости массива,
    // срез начинает ссылаться на новую область памяти

    // также в последнем примере мы добавили один элемент,
    // а ёмкость увеличился на два, это связано с оптимизацией -
    // постоянно перемещать элементы в новую область памяти дорого,
    // поэтому Go пытается нам помочь и заранее увеличить размер массива

    // если мы теперь добавим четвертый элемент, ссылка уже не поменяется,
    // потому что емкость останется прежней

    a = append(a, "_")
    fmt.Println(len(a), cap(a)) // 4 4
    fmt.Printf("%#v\n", reflect.ValueOf(a).Pointer()) // 0xc00012e040
}
```

Операция эта не самая дешёвая: нужно и память новую выделить и значения туда скопировать. Поэтому в мире Go стараются такого избегать, несмотря на оптимизации, которые пытается делать язык. А избежать этого часто достаточно просто: нужно создать срез сразу с нужной ёмкостью, для этого можно использовать встроенную функцию `make()`. Для создания среза первым аргументом передается тип среза, вторым длина и третьим ёмкость. Посмотрим на примере:

```
a := make([]int, 2, 3)
// тут мы создаём срез длинной два и ёмкостью три
// так как мы задаем длину, в срезе уже будут элементы
// со значениями по умолчанию, это просто для примера,
// вторым аргументом можно было просто передать 0

fmt.Printf("%#v\n", a)      // []int{0, 0}
fmt.Println(len(a), cap(a)) // 2 3
```

Таким образом, в случаях с заранее известным числом элементов можно уменьшить количество выполняемых операций. Для этого, кстати, есть неплохой статический анализатор, который умеет подсказывать места, в которых можно использовать технику предварительного аллоцирования массива с нужной емкостью: [https://github.com/alexkohler/prealloc](https://github.com/alexkohler/prealloc). Там же есть и бенчмарки, показывающие насколько замедляют код лишние увеличения ёмкости.

Но на этом мой рассказ не заканчивается, впереди есть ещё кое-что интересное. Следующий пример является следствием того, что мы рассмотрели выше. Но, тем не менее, лучше один раз увидеть, чем сойти с ума в поисках ошибки.

```
// создаем некоторый массив и срез от него
a := [...]int{1, 2}
s := a[:]

// как уже делали ранее, меняем первый элемент массива через срез
s[0] = 42
fmt.Printf("%#v\n", a) // []int{42, 2}
// отлично, поменялся первый элемент в исходном массиве

// теперь добавим к срезу одно значение
s = append(s, 3)

// и снова поменяем первый элемент
s[0] = 1
fmt.Printf("%#v\n", a) // []int{42, 2}
// если вы читали внимательно,
// то уже понимаете, почему ничего не изменилось

```

Итак, рассказываю, при втором присвоении ничего не происходит, потому что функция `append` при добавлении элемента создала новый массив и скопировала все элементы среза туда. После этого срез уже никак не связан с первоначальным массивом и через этот срез больше нельзя менять элементы исходного массива.

Для полноты картины приведу еще один пример, на этот раз с сортировкой.

```
a := [...]int{2, 1, 3}
s := a[:]

// сортируем срез от меньшего к большему
sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

fmt.Println(a) // [1 2 3]
// отлично, исходный массив отсортирован
// срез следовательно тоже, так как он
// просто ссылается на отсортированный массив
fmt.Println(s) // [1 2 3]

// теперь изменим размер среза и попробуем
// ещё раз отсортировать, но уже от большего
s = append(s, 42)

sort.Slice(s, func(i, j int) bool { return s[i] > s[j] })
fmt.Println(a) // [1 2 3]
// как и ожидалось, исходный массив остался не отсортирован,
// но зато отсортирован срез s
fmt.Println(s) // [42 3 2 1]

```

Но и функция `append` не так проста. В комментариях мне подсказали [очень хороший пример](https://habr.com/ru/post/597521/#comment_23879937), который лёг в основу следующего.

```
func main() {
    // на этот раз создадим срез, а не массив,
    // на основе которого дальше создадим новый срез,
    // так как в жизни, чаще в основе наших срезов
    // будут лежать другие срезы, но в этом примере с массивом
    // всё работало бы аналогично
    a := []int{1, 2, 3, 4, 5, 6}

    // создаём новый срез, от нашего базового среза
    s := a[:2]
    fmt.Printf("%#v\n", s)      // []int{1, 2}
    fmt.Println(len(s), cap(s)) // 2 6
    // тут нужно обратить внимание, что cap равняется шести,
    // всё потому что новый срез создаётся с таким объёмом,
    // чтобы уместиться в объём базового,
    // при этом быть максимально большим
    // например, у среза a[2:3] объём был бы равен четырём,
    // так как мы "отрезали" кусок от исходного среза,
    // начиная со второго элемента (цифры 3), и до конца изначально
    // выделенной памяти осталось ещё четыре ячейки памяти

    // дальше делаем несколько добавлений в конец нашего нового среза
    s = append(s, 800)
    s = append(s, 900)

    // и смотрим что же получилось
    fmt.Printf("%#v\n", s) // []int{1, 2, 800, 900}
    // выглядит всё хорошо: два элемента добавились в срез
    fmt.Println(len(s), cap(s)) // 4 6
    // тут тоже всё понятно: к отрезанным двум элементам
    // добавили ещё два элемента и длина стала четыре

    // но посмотрим что стало с исходным нашим срезом
    fmt.Printf("%#v\n", a) // []int{1, 2, 800, 900, 5, 6}
    // append не просто добавил элементы к новому срезу,
    // он их записал поверх элементов исходного среза
}
```

### Заключение

В процессе написания статьи мне задали вопрос: “А что же в итоге использовать? Срезы или массивы? И зачем нужны массивы, если можно просто всегда использовать срезы?”. Ответ следующий: всегда используй массивы, где это возможно и не ведет к усложнению кода. К сожалению, в реальном коде мест, где получается использовать массивы, не так много. Тем не менее, массив остаётся более строгим типом данных, который заставляет программиста правильнее работать с оперативной памятью, что иногда может защитить от ненужного замедления некоторых участков кода.

Как и многое в Go, срезы на первый взгляд кажутся несложными, но имеют интересные подводные камни. Для прочтения очень рекомендую эту статью [Go Slices: usage and internals](https://go.dev/blog/slices-intro) из официального блога Go.

Будьте аккуратны с Go, особенно с вещами, которые в нём происходят автоматически. Старайтесь использовать статические анализаторы, они иногда здорово могут помочь с обнаружением ошибок, в том числе при не оптимальном создании слайсов. О других интересных ошибках я расскажу в следующих сериях.