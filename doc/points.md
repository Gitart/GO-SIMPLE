# Как избежать Го Готча

Опубликовано 4 ноября 2016 г.

[#golang](https://divan.dev//tags/golang)

##### TL; DR путем обучения внутренним

> Гоча является действительной конструкцией в системе, программе или языке программирования, которая работает как документированная, но противоречит интуитивному принципу и почти допускает ошибки, потому что ее легко вызывать, и она неожиданна или неразумна по своему результату (источник: [wikipedia](https://en.wikipedia.org/wiki/Gotcha_(programming)) )

Язык программирования Go имеет несколько подводных камней , и существует [ряд](https://go-traps.appspot.com) из [хороших статей ,](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html) [объясняя](https://medium.com/@Jarema./golang-slice-append-gotcha-e9020ff37374#.xvfl7r4ti) их. Я нахожу эти статьи очень важными, особенно для новичков в Go, так как я вижу, что люди время от времени сталкиваются с этими проблемами.

Но один вопрос долго беспокоил меня \- почему я никогда не сталкивался с большинством этих ошибок? Серьезно, самые известные из них, такие как путаница с «нулевым интерфейсом» или «добавление фрагментов», никогда не были проблемой для меня. Я как\-то избежал этих проблем с первых дней работы с Go. Почему это так?

И ответ был на самом деле прост. Оказалось, что мне посчастливилось прочитать несколько статей о внутренних представлениях структур данных Go и изучить некоторые основы того, как все работает внутри Go. И этого знания было достаточно, чтобы построить интуицию о Го и избежать этих ошибок.

Помните, *«ошибки есть ... действительные конструкции ... но нелогичные»* ? Вот и все. У вас есть только два варианта:

*   «Исправить» язык
*   исправить интуицию

Второе на самом деле было бы лучше рассматривать как *построить интуицию* . Если у вас есть четкое представление о том, как срезы или интерфейсы работают под капотом, почти невозможно столкнуться с этими ошибками.

Таким образом, это сработало для меня и, вероятно, будет работать для других. Вот почему я решил собрать эти базовые знания о некоторых внутренностях Go в этом посте и помочь людям построить интуицию о представлении в памяти различных вещей.

Давайте начнем с базового понимания того, как вещи представлены в памяти. Краткий обзор того, что мы собираемся изучить:

*   [указатели](#pointers)
*   [Массивы и ломтики](#arrays-and-slices)
*   [Append](#append)
*   [Интерфейсы](#interfaces)
*   [Пустой интерфейс](#empty-interface)

# указатели

На самом деле, Go довольно близок к оборудованию. Когда вы создаете 64\-битную `int64` переменную integer ( ), вы точно знаете, сколько памяти она занимает, и вы можете использовать [unsafe.Sizeof ()](https://golang.org/pkg/unsafe/#Sizeof) для вычисления размера любого другого типа.

Я часто использую визуализацию блоков памяти, чтобы «увидеть» размеры переменных, массивов и структур данных. Визуальное представление дает вам простой способ получить представление о типах и часто помогает рассуждать о поведении и производительности.

Для разминки давайте визуализируем большинство основных типов в Go: ![Основные типы](https://divan.dev/images/basic_types.png) Предполагая, что вы работаете на 32\-битной машине *(что, вероятно, в настоящее время неверно)* , вы можете видеть, что *int64* занимает вдвое больше памяти, чем *int* .

Немного сложнее внутреннее представление указателей \- фактически это один блок в памяти, который содержит адрес памяти для некоторой другой области в памяти, где хранятся фактические данные. Когда вы слышите причудливое слово *«разыменование указателя»,* это на самом деле означает *«добраться до реальных блоков памяти по адресу, сохраненному в переменной указателя»* . Вы можете представить это так: ![указатели](https://divan.dev/images/pointers.png) адрес в памяти обычно представлен шестнадцатеричным значением, поэтому  на рисунке *«0x…»* .  Но знание того, что «значение указателя» может находиться в одном месте, а «фактические данные, на которые ссылается указатель» \- в другом, поможет нам в будущем.

Теперь одна из «хитростей» для новичков в Go, особенно у тех, кто ранее не знал языков с указателями, \- это путаница между «передачей по значению» параметров функции. Как вы знаете, в Go все передается «по значению», то есть путем копирования. Это должно быть намного проще, если вы попытаетесь визуализировать это копирование:

![Func params](https://divan.dev/images/func_params.png) В первом случае вы копируете все эти блоки памяти \- а на самом деле их часто намного больше, чем 2 \- это может быть легко 2 миллиона блоков, и вам нужно скопировать их все, что является одной из самых дорогих операций.  Но во втором случае вы копируете только один блок памяти, который содержит адрес фактических данных, и это быстро и дешево.

Теперь вы, естественно, можете видеть, что изменение `p` в функции `Foo()` не будет изменять исходные данные в первом случае, но определенно изменится во втором случае, так как сохраненный адрес `p` ссылается на исходные блоки данных.

Хорошо, если вы поняли, как знание внутренних представлений может помочь вам избежать распространенных ошибок, давайте углубимся немного глубже.

# Массивы и ломтики

Ломтики в начале часто путают с массивами. Итак, давайте посмотрим на массивы.

### Массивы

```go
var arr [5]int
var arr [5]int{1,2,3,4,5}
var arr [...]int{1,2,3,4,5}
```

Массивы \- это просто непрерывные блоки памяти, и если вы проверите исходный код среды выполнения Go ( [src / runtime / malloc.go](https://golang.org/src/runtime/malloc.go#L793) ), вы можете увидеть, что создание массива \- это, по сути, выделение фрагмента памяти заданного размера. Старый добрый маллок, просто умнее :)

```go
// newarray allocates an array of n elements of type typ.
func newarray(typ *_type, n int) unsafe.Pointer {
    if n < 0 || uintptr(n) > maxSliceCap(typ.size) {
        panic(plainError("runtime: allocation size out of range"))
    }
    return mallocgc(typ.size*uintptr(n), typ, true)
}
```

Что это значит для нас? Это означает, что мы можем просто представить массив как набор блоков в памяти, расположенных рядом друг с другом: ![массив](https://divan.dev/images/array.png) элементы массива всегда инициализируются *нулевыми значениями* его типа, в нашем случае 0 `[5]int` .  Мы можем проиндексировать их и получить длину, используя `len()` встроенную команду.  Больше ничего, в принципе.  Когда вы ссылаетесь на один элемент в массиве по индексу и делаете что\-то вроде этого:

```go
var arr [5]int
arr[4] = 42
```

вы берете пятый (4 + 1) элемент и меняете его значение: ![Массив 2](https://divan.dev/images/array2.png) теперь мы готовы исследовать фрагменты.

### Ломтики

На первый взгляд кусочки похожи на массивы, а объявление действительно похоже:

```go
var foo []int
```

Но если мы перейдем к исходному коду Go ( [src / runtime / slice.go](https://golang.org/src/runtime/slice.go#L11) ), мы увидим, что на самом деле срезы Go \- это структуры с тремя полями \- указатель на массив, длину и емкость:

```go
type slice struct {
        array unsafe.Pointer
        len   int
        cap   int
}
```

Когда вы создаете новый фрагмент, среда выполнения Go создает в памяти этот объект из трех блоков с указателем, установленным в `nil` и `len` и `cap` равным 0. Давайте представим его визуально: ![Срез 1](https://divan.dev/images/slice1.png) это не очень интересно, поэтому давайте использовать `make` для инициализации фрагмента заданного размера. :

```go
foo = make([]int, 5)
```

создаст срез с базовым массивом из 5 элементов, инициализированный с 0, и установит оба значения `len` и `cap` на 5. Cap означает емкость и поможет зарезервировать больше места для будущего роста. Вы можете использовать `make([]int, len, cap)` синтаксис для указания емкости. На самом деле, вам почти никогда не придется иметь дело с этим, но важно понимать концепцию потенциала.

```go
foo = make([]int, 3, 5)
```

Давайте посмотрим на оба случая: ![Срез 2](https://divan.dev/images/slice2.png)

Теперь, когда вы обновляете некоторые элементы среза, вы фактически изменяете значения в базовом массиве.

```go
foo = make([]int, 5)
foo[3] = 42
foo[4] = 100
```

![Срез 3](https://divan.dev/images/slice3.png)

Это было просто. Но что произойдет, если вы создадите еще один подлис и измените некоторые элементы? Давай попробуем:

```go
foo = make([]int, 5)
foo[3] = 42
foo[4] = 100

bar  := foo[1:4]
bar[1] = 99
```

![Срез 4](https://divan.dev/images/slice4.png)

Теперь вы видите это! Изменяя `bar` , вы фактически изменили базовый массив, на который также ссылается slice `foo` . И это на самом деле реально, вы можете написать что\-то вроде этого:

```go
var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    return digitRegexp.Find(b)
}
```

Считая, скажем, 10 МБ данных в срез и ища только 3 цифры, вы можете предположить, что вы возвращаете 3 байта, но на самом деле базовый массив будет храниться в памяти, независимо от его размера. ![Ломтик 5](https://divan.dev/images/slice5.png)

И это одна из самых распространенных ошибок Го, о которой вы можете прочитать. Но как только у вас в голове будет эта картина внутреннего среза, держу пари, что с ней почти невозможно столкнуться!

# Append

Рядом с гочами срезов, есть некоторые ошибки, связанные со встроенной универсальной функцией `append()` . По сути, он выполняет одну операцию \- добавляет значение к срезу, но внутренне выполняет много сложных задач, выделяя память разумным и эффективным способом, если это необходимо.

Давайте возьмем следующий код:

```go
a := make([]int, 32)
a = append(a, 1)
```

Он создает новый фрагмент размером 32 дюйма и добавляет новый (33\-й) элемент.

Помните `cap` \- емкость в ломтиках? Емкость означает *способность расти* . `append` проверяет, имеет ли срез больше возможностей для роста и, если нет, выделяет больше памяти. Выделение памяти \- довольно дорогая операция, поэтому она `append` пытается предвидеть эту операцию и запрашивает не 1 байт, а больше 32 байта \- в два раза больше первоначальной емкости. Опять же, выделение большего количества памяти за один раз, как правило, дешевле и быстрее, чем выделение меньшего количества памяти много раз.

Запутанная часть здесь заключается в том, что по многим причинам выделение большего объема памяти обычно означает выделение ее по другому адресу и перемещение данных из старого места в новое. Это означает, что адрес базового массива в срезе также изменится. Давайте визуализируем это: ![Append](https://divan.dev/images/append.png) легко увидеть два базовых массива \- старый и новый.  Не похоже на возможную ошибку, верно?  Старый массив будет освобожден GC позже, если другой фрагмент не ссылается на него.  Этот случай, на самом деле, является одним из ошибок с добавлением.  Что если вы создадите подлису `b` , а затем  добавите  значение `a` , предполагая, что они совместно используют общий базовый массив?

```go
a := make([]int, 32)
b := a[1:16]
a = append(a, 1)
a[2] = 42
```

Вы получите это: ![Добавить 2](https://divan.dev/images/append2.png) Да, у вас будет два разных базовых массива, и это может быть довольно нелогичным для начинающих.  Так что, как правило, будьте осторожны при использовании субликсов, особенно субликсов с добавлением.

Кстати, `append` ломтик растет, удваивая его емкость только до 1024, после чего он будет использовать так называемые [классы размера памяти,](https://golang.org/src/runtime/msize.go) чтобы гарантировать, что рост будет не более ~ 12,5%. Запросить 64 байта для массива 32 байта \- это нормально, но если ваш слайс равен 4 ГБ, выделение еще 4 ГБ для добавления 1 элемента довольно дорого, так что это имеет смысл.

# Интерфейсы

Хорошо, это самая запутанная вещь для многих людей. Требуется некоторое время, чтобы обдумать правильное использование интерфейсов в Go, особенно после того, как вы испытали травматический опыт работы с языками классов. И одним из источников путаницы является другое значение `nil` ключевого слова в контексте интерфейсов.

Чтобы помочь понять эту тему, давайте снова взглянем на исходный код Go. Что такое интерфейс под капотом? Вот код из [src / runtime / runtime2.go](https://golang.org/src/runtime/runtime2.go#L143) :

```go
type iface struct {
    tab  *itab
    data unsafe.Pointer
}
```

`itab` обозначает *интерфейсную таблицу* и также является структурой, которая содержит необходимую метаинформацию об интерфейсе и базовом типе:

```go
type itab struct {
    inter  *interfacetype
    _type  *_type
    link   *itab
    bad    int32
    unused int32
    fun    [1]uintptr // variable sized
}
```

Мы не собираемся изучать логику того, как работает утверждение типа интерфейса, но важно понимать, что *интерфейс* представляет собой соединение информации интерфейса и статического типа плюс указатель на фактическую переменную (поле `data` в `iface` ). Давайте создадим переменную `err` типа интерфейса `error` и представим ее визуально:

```go
var err error
```

![Interface1](https://divan.dev/images/iface1.png)

Фактически, то, что вы видите на этой картинке, называется *ниль интерфейсом* . Когда вы возвращаете nil в функцию с типом возврата `error` , вы возвращаете этот объект. Он содержит информацию об интерфейсе ( `itab.inter` ), но имеет поля `nil` in `data` и `itab.type` . Этот объект будет оцениваться как истинный в `if err == nil {}` условии.

```go
func foo() error {
    var err error // nil
    return err
}

err := foo()
if err == nil {...} // true
```

Теперь, известная ошибка должна вернуть `*os.PathError` переменную, которая есть `nil` .

```go
func foo() error {
    var err *os.PathError // nil
    return err
}

err := foo()
if err == nil {...} // false
```

Эти две части кода похожи, если только вы не знаете, как выглядит интерфейс внутри. Давайте представим эту `nil` переменную типа `*os.PathError` , обернутую в `error` интерфейсе: ![Interface2](https://divan.dev/images/iface2.png) вы можете ясно видеть `*os.PathError` переменную \- это просто блок памяти, содержащий `nil` значение, потому что это нулевое значение для указателей.  Но фактическое `error` возвращение `foo()` \- это очень сложная структура с информацией об интерфейсе, об основном типе и адресе памяти этого блока памяти, содержащего `nil` указатель.  Почувствуйте разницу?

В обоих случаях мы имеем `nil` , но есть огромная разница между *«иметь интерфейс с переменной, значение которой равно nil»* и *«интерфейс без переменной»* . Имея это знание внутренней структуры интерфейсов, попробуйте теперь перепутать эти два примера: сейчас ![Interface3](https://divan.dev/images/iface3.png) должно быть намного сложнее столкнуться с этой ошибкой.

### Пустой интерфейс

Несколько слов о *пустом интерфейсе* \- `interface{}` . В исходном коде Go ( [src / runtime / malloc.go](https://golang.org/src/runtime/runtime2.go#L148) это реализовано с использованием собственной структуры \- `eface` :

```go
type eface struct {
    _type *_type
    data  unsafe.Pointer
}
```

Как вы можете видеть, он похож `iface` , но не имеет интерфейсной таблицы. Он просто не нужен, потому что по определению пустой интерфейс реализован любым статическим типом. Поэтому, когда вы заключаете что\-то \- явно или неявно (например, передавая в качестве аргумента функции) \- в `interface{}` , вы фактически работаете с этой структурой:

```go
func foo() interface{} {
    foo := int64(42)
    return foo
}
```

![Пустой интерфейс](https://divan.dev/images/eface.png)

Одним из `interface{}` связанных с этим недостатков является разочарование, что вы не можете легко назначить фрагмент интерфейсов фрагменту конкретных типов и наоборот. Что\-то типа

```go
func foo() []interface{} {
    return []int{1,2,3}
}
```

Вы получите ошибку времени компиляции:

```bash
$ go build
cannot use []int literal (type []int) as type []interface {} in return argument
```

Это сбивает с толку и начало. Мол, почему я могу сделать это преобразование с одной переменной, но не могу сделать со слайсом? Но как только вы узнаете, что такое пустой интерфейс (еще раз посмотрите на рисунок выше), становится совершенно ясно, что это «преобразование» на самом деле является довольно дорогой операцией, которая включает в себя выделение группы памяти и занимает около O (n). времени и пространства. И один из распространенных подходов в дизайне Go \- **«если хочешь сделать что\-то дорогое \- делай это явно»** . ![Slice Interface](https://divan.dev/images/eface_slice.png) Надеюсь, теперь это имеет смысл и для вас.

# Вывод

Не каждый готч может быть атакован обучением внутренних органов. Некоторые из них \- это просто разница между вашим прошлым и новым опытом, и у всех нас есть какой\-то другой опыт и опыт. Тем не менее, есть много ошибок, которых можно успешно избежать, просто поняв немного глубже, как работает Go. Я надеюсь, что объяснения в этом посте помогут вам понять, что происходит внутри ваших программ, и сделают вас лучшим разработчиком. Иди \- твой друг, и, зная это немного лучше, не повредит в любом случае.

Если вам интересно больше узнать о внутренностях Go, вот список ссылок, которые мне помогли:

*   [Структуры данных Go](http://research.swtch.com/godata)
*   [Структуры данных Go: Интерфейсы](http://research.swtch.com/interfaces)
*   [Go Slices: использование и внутренности](https://blog.golang.org/go-slices-usage-and-internals)
*   [Gopher Puzzlers](http://talks.godoc.org/github.com/davecheney/presentations/gopher-puzzlers.slide)

И, конечно, вечный источник полезных вещей :)

*   [Перейти исходный код](https://golang.org/src/)
*   [Эффективный Go](https://golang.org/doc/effective_go.html)
*   [Перейти спец](https://golang.org/ref/spec)

Счастливого взлома!

PS. Я также выступил с аналогичным докладом на [Golang BCN](http://www.meetup.com/Golang-Barcelona/) Meetup в ноябре 16 года.

Вот слайды: [Как избежать Go Gotchas.pdf](http://divan.dev/talks/2016/bcn/HowToAvoidGoGotchas.pdf)
