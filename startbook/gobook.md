# Учебник Go

Февраль 7, 2020

## Что такое го?

Go (также известный как Golang) — это язык программирования с открытым исходным кодом, разработанный Google. Это статически скомпилированный язык. Go поддерживает параллельное программирование, то есть позволяет запускать несколько процессов одновременно. Это достигается с помощью каналов, подпрограмм и т. Д. Go имеет сборку мусора, которая сама выполняет управление памятью и обеспечивает отложенное выполнение функций.

**Что вы узнаете:** \[ скрыть \]

*   [Что такое го?](#1)
*   [Как скачать и установить GO](#2)
*   [Ваша программа First Go](#3)
*   [Типы данных](#4)
*   [переменные](#5)
*   [Константы](#6)
*   [Loops](#7)
*   [Если еще](#8)
*   [переключатель](#9)
*   [Массивы](#10)
*   [Кусочек](#11)
*   [функции](#12)
*   [пакеты](#13)
*   [Откладывать и укладывать](#14)
*   [указатели](#15)
*   [сооружения](#16)
*   [Методы (не функции)](#17)
*   [совпадение](#18)
*   [Goroutines](#19)
*   [каналы](#20)
*   [Выбрать](#21)
*   [Mutex](#22)
*   [Обработка ошибок](#23)
*   [Пользовательские ошибки](#24)
*   [Чтение файлов](#25)
*   [Запись файлов](#26)
*   [Шпаргалка](#27)

## Как скачать и установить GO

**Шаг 1)** Перейдите на [https://golang.org/dl/](https://translate.googleusercontent.com/translate_c?depth=1&pto=aue&rurl=translate.google.ru&sl=en&sp=nmt4&tl=ru&u=https://golang.org/dl/&usg=ALkJrhiZOWh0MrfdyJSfmqUUBSj-HBwf8A) . Загрузите бинарный файл для вашей ОС.

[![](https://coderlessons.com/wp-content/uploads/images/gur/b31ad292cb480ac35584f23daf25211d.png)

![](https://coderlessons.com/wp-content/uploads/images/gur/b31ad292cb480ac35584f23daf25211d.png)

](https://coderlessons.com/wp-content/uploads/images/gur/b31ad292cb480ac35584f23daf25211d.png)

**Шаг 2)** Дважды щелкните установщик и нажмите «Выполнить».

[![](https://coderlessons.com/wp-content/uploads/images/gur/7166439b084a83cfaf0080a6b8d27778.png)

![](https://coderlessons.com/wp-content/uploads/images/gur/7166439b084a83cfaf0080a6b8d27778.png)

](https://coderlessons.com/wp-content/uploads/images/gur/7166439b084a83cfaf0080a6b8d27778.png)

**Шаг 3)** Нажмите Далее

[![](https://coderlessons.com/wp-content/uploads/images/gur/be60185eaa6ab86d9c2db89ef80a1b55.png)

![](https://coderlessons.com/wp-content/uploads/images/gur/be60185eaa6ab86d9c2db89ef80a1b55.png)

](https://coderlessons.com/wp-content/uploads/images/gur/be60185eaa6ab86d9c2db89ef80a1b55.png)

**Шаг 4)** Выберите папку для установки и нажмите Далее.

[![](https://coderlessons.com/wp-content/uploads/images/gur/8bfa93177b988ba62bde6a8ddeeb9cc5.png)

![](https://coderlessons.com/wp-content/uploads/images/gur/8bfa93177b988ba62bde6a8ddeeb9cc5.png)

](https://coderlessons.com/wp-content/uploads/images/gur/8bfa93177b988ba62bde6a8ddeeb9cc5.png)

**Шаг 5)** Нажмите Finish после завершения установки.

[![](https://coderlessons.com/wp-content/uploads/images/gur/30b9b01e339d7b55f39eca3461446031.png)

![](https://coderlessons.com/wp-content/uploads/images/gur/30b9b01e339d7b55f39eca3461446031.png)

](https://coderlessons.com/wp-content/uploads/images/gur/30b9b01e339d7b55f39eca3461446031.png)

**Шаг 6)** После завершения установки вы можете проверить это, открыв терминал и набрав

go version

Это покажет версию Go установлен

[![](https://coderlessons.com/wp-content/uploads/images/gur/eb99629126d1f42ec01b600c97786a83.png)

![](https://coderlessons.com/wp-content/uploads/images/gur/eb99629126d1f42ec01b600c97786a83.png)

](https://coderlessons.com/wp-content/uploads/images/gur/eb99629126d1f42ec01b600c97786a83.png)

## Ваша программа First Go

Создайте папку с именем studyGo. Вы создадите наши программы go внутри этой папки. Go файлы создаются с расширением **.go** . Вы можете запускать программы Go, используя синтаксис

go run <filename>

Создайте файл с именем first.go, добавьте в него приведенный ниже код и сохраните

package main
import ("fmt")

func main() {
	fmt.Println("Hello World! This is my first Go program\\n")
}

[![](https://coderlessons.com/wp-content/uploads/images/gur/976647f2de38c2ce313254c74a538b9a.png)

![](https://coderlessons.com/wp-content/uploads/images/gur/976647f2de38c2ce313254c74a538b9a.png)

](https://coderlessons.com/wp-content/uploads/images/gur/976647f2de38c2ce313254c74a538b9a.png)

Перейдите к этой папке в вашем терминале. Запустите программу с помощью команды

иди беги первым

Вы можете увидеть вывод печати

Hello World! This is my first Go program

[![](https://coderlessons.com/wp-content/uploads/images/gur/a67cc480a5d1d149db268c23a741084d.png)

![](https://coderlessons.com/wp-content/uploads/images/gur/a67cc480a5d1d149db268c23a741084d.png)

](https://coderlessons.com/wp-content/uploads/images/gur/a67cc480a5d1d149db268c23a741084d.png)

Теперь давайте обсудим вышеуказанную программу.

основной пакет — каждая программа go должна начинаться с имени пакета. Go позволяет нам использовать пакеты в других программах go и, следовательно, поддерживает повторное использование кода. Выполнение программы go начинается с кода внутри пакета с именем main.

import fmt — импортирует пакет fmt. Этот пакет реализует функции ввода / вывода.

func main () — это функция, с которой начинается выполнение программы. Основная функция всегда должна быть помещена в основной пакет. Под main (), вы можете написать код внутри {}.

fmt.Println — это напечатает текст на экране с помощью функции Println fmt.

Примечание. В следующих разделах, когда вы упоминаете выполнить / запустить код, это означает сохранить код в файл с расширением .go и запустить его с использованием синтаксиса.

    go run <filename>

## Типы данных

Типы (типы данных) представляют тип значения, хранящегося в переменной, тип значения, которое возвращает функция, и т. Д.

В Go есть три основных типа.

**Числовые типы** — представляют числовые значения, которые включают целые числа, числа с плавающей запятой и комплексные значения. Различные числовые типы:

int8 — 8\-битные целые числа со знаком.

int16 — 16\-битные целые числа со знаком.

int32 — 32\-битные целые числа со знаком.

int64 — 64\-разрядные целые числа со знаком.

uint8 — 8\-битные целые числа без знака.

uint16 — 16\-битные целые числа без знака.

uint32 — 32\-битные целые числа без знака.

uint64 — 64\-битные целые числа без знака.

float32 — 32\-битные числа с плавающей точкой.

float64 — 64\-битные числа с плавающей точкой.

complex64 — имеет float32 реальных и мнимых частей.

complex128 — имеет float32 реальных и мнимых частей.

**String types** — Представляет последовательность байтов (символов). Вы можете выполнять различные операции со строками, такие как конкатенация строк, извлечение подстроки и т. Д.

**Булевы типы.** Представляет 2 значения: true или false.

## переменные

Переменные указывают на область памяти, в которой хранится какое\-то значение. Параметр type (в приведенном ниже синтаксисе) представляет тип значения, которое можно сохранить в ячейке памяти.

Переменная может быть объявлена ​​с использованием синтаксиса

    var <variable\_name> <type>

После того, как вы объявите переменную типа, вы можете присвоить переменную любому значению этого типа.

Вы также можете дать начальное значение переменной во время самого объявления, используя

    var <variable\_name> <type> = <value>

Если вы объявляете переменную с начальным значением, перейдите к типу переменной по типу присвоенного значения. Таким образом, вы можете опустить тип во время объявления, используя синтаксис

    var <variable\_name> = <value>

Кроме того, вы можете объявить несколько переменных с синтаксисом

    var <variable\_name1>, <variable\_name2>  = <value1>, <value2>

Программа ниже имеет несколько примеров объявлений переменных

package main
import "fmt"

func main() {
    //declaring a integer variable x
    var x int
    x=3 //assigning x the value 3
    fmt.Println("x:", x) //prints 3

    //declaring a integer variable y with value 20 in a single statement and prints it
    var y int=20
    fmt.Println("y:", y)

    //declaring a variable z with value 50 and prints it
    //Here type int is not explicitly mentioned
    var z=50
    fmt.Println("z:", z)

    //Multiple variables are assigned in single line\- i with an integer and j with a string
    var i, j = 100,"hello"
    fmt.Println("i and j:", i,j)
}

Выход будет

x: 3
y: 20
z: 50
i and j: 100 hello

Go также предоставляет простой способ объявления переменных со значением, опуская ключевое слово var, используя

 <variable\_name> := <value>

Обратите внимание, что вы использовали **: =** вместо **\=** . Вы не можете использовать: = просто для присвоения значения переменной, которая уже объявлена. : = используется для объявления и присвоения значения.

Создайте файл с именем assign.go со следующим кодом

package main
import ("fmt")

func main() {
	a := 20
	fmt.Println(a)

	//gives error since a is already declared
	a := 30
	fmt.Println(a)
}

Выполните go run assign.go, чтобы увидеть результат как

./assign.go:7:4: no new variables on left side of :=

Переменные, объявленные без начального значения, будут иметь 0 для числовых типов, false для Boolean и пустую строку для строк

## Константы

Постоянные переменные — это те переменные, значение которых нельзя изменить после присвоения. Константа в Go объявляется с помощью ключевого слова «const»

Создайте файл с именем constant.go и со следующим кодом

package main
import ("fmt")

func main() {
	const b =10
	fmt.Println(b)
	b = 30
	fmt.Println(b)
}

Выполните go run constant.go, чтобы увидеть результат как

.constant.go:7:4: cannot assign to b

## Loops

Циклы используются для многократного выполнения блока операторов в зависимости от условия. Большинство языков программирования предоставляют 3 типа циклов — для while, а для while. **Но Go поддерживает только цикл.**

Синтаксис цикла for

for initialisation\_expression; evaluation\_expression; iteration\_expression{
   // one or more statement
}

Выражение initialisation\_expression выполняется первым (и только один раз).

Затем оценивается выражение\_производства, и если оно истинно, выполняется код внутри блока.

Идентификатор iteration\_expression выполняется, и выражение\_класса вычисляется снова. Если это правда, блок операторов выполняется снова. Это будет продолжаться до тех пор, пока expression\_expression не станет ложным.

Скопируйте приведенную ниже программу в файл и выполните ее, чтобы увидеть числа циклической печати от 1 до 5

package main
import "fmt"

func main() {
var i int
for i = 1; i <= 5; i++ {
fmt.Println(i)
    }
}

Выход

1
2
3
4
5

## Если еще

Если еще это условное утверждение. Синакс

if condition{
// statements\_1
}else{
// statements\_2
}

Здесь условие оценивается, и если оно истинно, будут выполнены операторы\_1, иначе будут выполнены операторы\_2.

Вы можете использовать оператор if без других также. Вы также можете приковать если еще заявления. Приведенные ниже программы объяснят больше, если еще.

Выполните следующую программу. Он проверяет, меньше ли число х, чем 10. Если это так, он напечатает «х меньше, чем 10»

package main
import "fmt"

func main() {
    var x = 50
    if x < 10 {
        //Executes if x < 10
        fmt.Println("x is less than 10")
    }
}

Здесь, поскольку значение x больше 10, инструкция внутри условия блока не будет выполнена.

Теперь смотрите программу ниже. У нас есть блок else, который будет выполнен при сбое оценки if.

package main
import "fmt"

func main() {
    var x = 50
    if x < 10 {
        //Executes if x is less than 10
        fmt.Println("x is less than 10")
    } else {
        //Executes if x >= 10
        fmt.Println("x is greater than or equals 10")
    }
}

Эта программа даст вам вывод

x is greater than or equals 10

Теперь мы увидим программу с несколькими блоками if else (с цепочкой if else). Выполните приведенный ниже пример. Он проверяет, является ли число меньше 10 или от 10 до 90 или больше 90.

package main
import "fmt"

func main() {
    var x = 100
    if x < 10 {
        //Executes if x is less than 10
        fmt.Println("x is less than 10")
    } else if x >= 10 && x <= 90 {
        //Executes if x >= 10 and x<=90
        fmt.Println("x is between 10 and 90")
    } else {
        //Executes if both above cases fail i.e x>90
        fmt.Println("x is greater than 90")
    }
}

Здесь сначала условие if проверяет, меньше ли x 10, и нет. Таким образом, он проверяет следующее условие (иначе, если), находится ли оно между 10 и 90, что также ложно. Затем он выполняет блок в разделе else, который дает вывод

x is greater than 90

## переключатель

Переключатель — это еще одно условное утверждение. Операторы Switch оценивают выражение, и результат сравнивается с набором доступных значений (случаев). Когда совпадение найдено, выполняются операторы, связанные с этим совпадением (регистром). Если совпадений не найдено, ничего не будет выполнено. Вы также можете добавить регистр по умолчанию для переключения, который будет выполняться, если не найдено других совпадений. Синтаксис переключателя

switch expression {
    case value\_1:
        statements\_1
    case value\_2:
        statements\_2
    case value\_n:
        statements\_n
    default:
        statements\_default
    }

Здесь значение выражения сравнивается со значениями в каждом случае. Когда совпадение найдено, выполняются операторы, связанные с этим делом. Если совпадений не найдено, выполняются операторы в разделе по умолчанию.

Выполните следующую программу

package main
import "fmt"

func main() {
    a,b := 2,1
    switch a+b {
    case 1:
        fmt.Println("Sum is 1")
    case 2:
        fmt.Println("Sum is 2")
    case 3:
        fmt.Println("Sum is 3")
    default:
        fmt.Println("Printing default")
    }
}

Вы получите вывод как

Sum is 3

Измените значение a и b на 3, и результат будет

Printing default

Вы также можете иметь несколько значений в кейсе, разделяя их запятой.

## Массивы

Массив представляет собой фиксированный размер именованной последовательности элементов одного типа. Вы не можете иметь массив, который содержит как целое число, так и символы. Вы не можете изменить размер массива, как только Вы определите размер.

Синтаксис для объявления массива

var arrayname \[size\] type

Каждому элементу массива может быть присвоено значение с использованием синтаксиса

arrayname \[index\] = value

Индекс массива начинается с **0 до размера\-1** .

Вы можете присвоить значения элементам массива во время объявления, используя синтаксис

arrayname := \[size\] type {value\_0,value\_1,…,value\_size\-1}

Вы также можете игнорировать параметр размера при объявлении массива со значениями, заменив размер на **…,** и компилятор найдет длину из числа значений. Синтаксис

arrayname :=  \[…\] type {value\_0,value\_1,…,value\_size\-1}

Вы можете найти длину массива, используя синтаксис

len(arrayname)

Выполните приведенный ниже пример, чтобы понять массив

package main
import "fmt"

func main() {
    var numbers \[3\] string //Declaring a string array of size 3 and adding elements
    numbers\[0\] = "One"
    numbers\[1\] = "Two"
    numbers\[2\] = "Three"
    fmt.Println(numbers\[1\]) //prints Two
    fmt.Println(len(numbers)) //prints 3
    fmt.Println(numbers) // prints \[One Two Three\]

    directions := \[...\] int {1,2,3,4,5} // creating an integer array and the size of the array is defined by the number of elements
    fmt.Println(directions) //prints \[1 2 3 4 5\]
    fmt.Println(len(directions)) //prints 5

    //Executing the below commented statement prints invalid array index 5 (out of bounds for 5\-element array)
    //fmt.Println(directions\[5\])
}

**Output**

Two
3
\[One Two Three\]
\[1 2 3 4 5\]
5

## Slice

A slice is a portion or segment of an array. Or it is a view or partial view of an underlying array to which it points. You can access the elements of a slice using the slice name and index number just as you do in an array. You cannot change the length of an array, but you can change the size of a slice.

Contents of a slice are actually the pointers to the elements of an array. It means **if you change any element in a slice, the underlying array contents also will be affected.**

The syntax for creating a slice is

var slice\_name \[\] type = array\_name\[start:end\]

This will create a slice named slice\_name from an array named array\_name with the elements at the index start to end\-1.

Execute the below program. The program will create a slice from the array and print it. Also, you can see that modifying the contents in the slice will modify the actual array.

package main
import "fmt"

func main() {
    // declaring array
    a := \[5\] string {"one", "two", "three", "four", "five"}
    fmt.Println("Array after creation:",a)

    var b \[\] string = a\[1:4\] //created a slice named b
    fmt.Println("Slice after creation:",b)

    b\[0\]="changed" // changed the slice data
    fmt.Println("Slice after modifying:",b)
    fmt.Println("Array after slice modification:",a)
}

This will print result as

Array after creation: \[one two three four five\]
Slice after creation: \[two three four\]
Slice after modifying: \[changed three four\]
Array after slice modification: \[one changed three four five\]

There are certain functions which you can apply on slices

**len(slice\_name)** — returns the length of the slice

**append(slice\_name, value\_1, value\_2)** — It is used to append value\_1 and value\_2 to an existing slice.

**append(slice\_nale1,slice\_name2…)** – appends slice\_name2 to slice\_name1

Выполните следующую программу.

package main
import "fmt"

func main() {
	a := \[5\] string {"1","2","3","4","5"}
	slice\_a := a\[1:3\]
	b := \[5\] string {"one","two","three","four","five"}
	slice\_b := b\[1:3\]

    fmt.Println("Slice\_a:", slice\_a)
    fmt.Println("Slice\_b:", slice\_b)
    fmt.Println("Length of slice\_a:", len(slice\_a))
    fmt.Println("Length of slice\_b:", len(slice\_b))

    slice\_a = append(slice\_a,slice\_b...) // appending slice
    fmt.Println("New Slice\_a after appending slice\_b :", slice\_a)

    slice\_a = append(slice\_a,"text1") // appending value
    fmt.Println("New Slice\_a after appending text1 :", slice\_a)
}

Выход будет

Slice\_a: \[2 3\]
Slice\_b: \[two three\]
Length of slice\_a: 2
Length of slice\_b: 2
New Slice\_a after appending slice\_b : \[2 3 two three\]
New Slice\_a after appending text1 : \[2 3 two three text1\]

Программа сначала создает 2 среза и печатает их длину. Затем он добавил один фрагмент к другому, а затем добавил строку к полученному фрагменту.

## функции

Функция представляет собой блок операторов, который выполняет определенную задачу. Объявление функции сообщает нам имя функции, тип возвращаемого значения и входные параметры. Определение функции представляет собой код, содержащийся в функции. Синтаксис объявления функции

func function\_name(parameter\_1 type, parameter\_n type) return\_type {
//statements
}

Параметры и типы возврата являются необязательными. Кроме того, вы можете вернуть несколько значений из функции.

Давайте запустим следующий пример. Здесь функция с именем calc будет принимать 2 числа и выполняет сложение и вычитание и возвращает оба значения.

package main
import "fmt"

//calc is the function name which accepts two integers num1 and num2
//(int, int) says that the function returns two values, both of integer type.
func calc(num1 int, num2 int)(int, int) {
    sum := num1 + num2
    diff := num1 \- num2
    return sum, diff
}

func main() {
    x,y := 15,10

    //calls the function calc with x and y an d gets sum, diff as output
    sum, diff := calc(x,y)
    fmt.Println("Sum",sum)
    fmt.Println("Diff",diff)
}

Выход будет

Sum 25
Diff 5

## пакеты

Пакеты используются для организации кода. В большом проекте невозможно написать код в одном файле. Перейти позволяют нам организовать код под разные пакеты. Это повышает читаемость кода и возможность его повторного использования. Исполняемая программа Go должна содержать пакет с именем main, и выполнение программы начинается с функции с именем main. Вы можете импортировать другие пакеты в нашей программе, используя синтаксис

import package\_name

Мы увидим и обсудим, как создавать и использовать пакеты в следующем примере.

**Шаг 1)** Создайте файл с именем package\_example.go и добавьте следующий код

package main
import "fmt"
//the package to be created
import "calculation"

func main() {
	x,y := 15,10
	//the package will have function Do\_add()
sum := calculation.Do\_add(x,y)
fmt.Println("Sum",sum)
}

В приведенной выше программе fmt представляет собой пакет, который Go предоставляет нам в основном для целей ввода / вывода. Кроме того, вы можете увидеть пакет с именем калькуляции. Внутри main () вы можете увидеть сумму шагов: = calculation.Do\_add (x, y). Это означает, что вы вызываете функцию Do\_add из расчета пакета.

**Шаг 2)** Сначала вы должны создать расчет пакета внутри папки с тем же именем в папке src the go. Установленный путь go можно найти в переменной PATH.

Для Mac найдите путь, выполнив echo $ PATH

[![](https://coderlessons.com/wp-content/uploads/images/gur/a5a2c4eaa1b0f06799d50be47d7df3b5.png)

![](https://coderlessons.com/wp-content/uploads/images/gur/a5a2c4eaa1b0f06799d50be47d7df3b5.png)

](https://coderlessons.com/wp-content/uploads/images/gur/a5a2c4eaa1b0f06799d50be47d7df3b5.png)

Таким образом, путь / usr / local / go

Для окон найдите путь, выполнив echo% GOROOT%

[![](https://coderlessons.com/wp-content/uploads/images/gur/58c62a6b1360852b1346a84efa00e462.png)

![](https://coderlessons.com/wp-content/uploads/images/gur/58c62a6b1360852b1346a84efa00e462.png)

](https://coderlessons.com/wp-content/uploads/images/gur/58c62a6b1360852b1346a84efa00e462.png)

Здесь путь C: \\ Go \\

**Шаг 3)** Перейдите в папку src (/ usr / local / go / src для Mac и C: \\ Go \\ src для Windows). Теперь из кода, имя пакета является расчетным. Go требует, чтобы пакет был помещен в каталог с тем же именем в каталоге src. Создайте каталог с именем вычислений в папке src.

**Шаг 4)** Создайте файл с именем calc.go (вы можете дать любое имя, но имя пакета в коде имеет значение. Здесь это должен быть расчет) внутри каталога вычислений и добавить следующий код

package calculation

func Do\_add(num1 int, num2 int)(int) {
    sum := num1 + num2
    return sum
}

**Шаг 5)** Запустите команду go install из каталога вычислений, который скомпилирует calc.go.

[![](https://coderlessons.com/wp-content/uploads/images/gur/931c1612e00e112052860e4d1536b508.png)

![](https://coderlessons.com/wp-content/uploads/images/gur/931c1612e00e112052860e4d1536b508.png)

](https://coderlessons.com/wp-content/uploads/images/gur/931c1612e00e112052860e4d1536b508.png)

**Шаг 6)** Теперь вернитесь к package\_example.go и запустите go, запустите package\_example.go. Выход будет Сумма 25.

Обратите внимание, что имя функции Do\_add начинается с заглавной буквы. Это происходит потому, что в Go, если имя функции начинается с заглавной буквы, это означает, что другие программы могут видеть (получать к нему доступ), иначе другие программы не могут получить к нему доступ. Если бы имя функции было do\_add, то вы бы получили ошибку

не может ссылаться на неэкспортированное имя

## Откладывать и укладывать

Операторы отсрочки используются для отсрочки выполнения вызова функции до тех пор, пока функция, содержащая инструкцию отсрочки, не завершит выполнение.

Давайте узнаем это на примере:

package main
import "fmt"

func sample() {
    fmt.Println("Inside the sample()")
}
func main() {
    //sample() will be invoked only after executing the statements of main()
    defer sample()
    fmt.Println("Inside the main()")
}

Выход будет

Inside the main()
Inside the sample()

Здесь выполнение sample () откладывается до завершения выполнения включающей функции (main ()).

Отсрочка стека использует несколько операторов отсрочки. Предположим, у вас есть несколько операторов отсрочки внутри функции. Go помещает все отложенные вызовы функций в стек, и как только вмещающая функция возвращается, составленные функции выполняются в порядке « **последний пришел — первый вышел» (LIFO).** Вы можете увидеть это в следующем примере.

Выполните код ниже

package main
import "fmt"

func display(a int) {
    fmt.Println(a)
}
func main() {
    defer display(1)
    defer display(2)
    defer display(3)
    fmt.Println(4)
}

Выход будет

4
3
2
1

Здесь сначала выполняется код внутри main (), а затем вызовы отложенных функций выполняются в обратном порядке, то есть 4, 3,2,1.

## указатели

Перед объяснением указателей давайте сначала обсудим оператор ‘&’. Оператор ‘&’ используется для получения адреса переменной. Это означает, что «& a» напечатает адрес памяти переменной a.

Выполните приведенную ниже программу, чтобы отобразить значение переменной и адрес этой переменной

package main
import "fmt"

func main() {
	a := 20
	fmt.Println("Address:",&a)
	fmt.Println("Value:",a)
}

Результат будет

Address: 0xc000078008
Value: 20

Переменная\-указатель хранит адрес памяти другой переменной. Вы можете определить указатель, используя синтаксис

	var variable\_name \*type

Звездочка (\*) представляет переменную\-указатель. Вы поймете больше, выполнив следующую программу

package main
import "fmt"

func main() {
	//Create an integer variable a with value 20
	a := 20

	//Create a pointer variable b and assigned the address of a
	var b \*int = &a

	//print address of a(&a) and value of a
	fmt.Println("Address of a:",&a)
	fmt.Println("Value of a:",a)

	//print b which contains the memory address of a i.e. &a
	fmt.Println("Address of pointer b:",b)

	//\*b prints the value in memory address which b contains i.e. the value of a
	fmt.Println("Value of pointer b",\*b)

	//increment the value of variable a using the variable b
	\*b = \*b+1

	//prints the new value using a and \*b
	fmt.Println("Value of pointer b",\*b)
	fmt.Println("Value of a:",a)}

Выход будет

Address of a: 0x416020
Value of a: 20
Address of pointer b: 0x416020
Value of pointer b 20
Value of pointer b 21
Value of a: 21

## сооружения

Структура — это определенный пользователем тип данных, который сам содержит еще один элемент того же или другого типа.

Использование структуры состоит из двух этапов.

Сначала создайте (объявите) тип структуры

Во\-вторых, создайте переменные этого типа для хранения значений.

Структуры в основном используются, когда вы хотите хранить связанные данные вместе.

Рассмотрим часть информации о сотруднике, которая имеет имя, возраст и адрес. Вы можете справиться с этим двумя способами

Создайте 3 массива: в одном массиве хранятся имена сотрудников, в одном — возраст, а в третьем — возраст.

Объявите тип структуры с 3 полями: имя, адрес и возраст. Создайте массив этого типа структуры, где каждый элемент является структурным объектом, имеющим имя, адрес и возраст.

Первый подход не эффективен. В таких сценариях структуры более удобны.

Синтаксис для объявления структуры

type structname struct {
   variable\_1 variable\_1\_type
   variable\_2 variable\_2\_type
   variable\_n variable\_n\_type
}

Пример объявления структуры

type emp struct {
    name string
    address string
    age int
}

Здесь создается новый пользовательский тип с именем emp. Теперь вы можете создавать переменные типа emp, используя синтаксис

	var variable\_name struct\_name

Примером является

var empdata1 emp

Вы можете установить значения для empdata1 как

empdata1.name = "John"
	empdata1.address = "Street\-1, Bangalore"
	empdata1.age = 30

Вы также можете создать структурную переменную и присвоить значения

empdata2 := emp{"Raj", "Building\-1, Delhi", 25}

Здесь вам нужно поддерживать порядок элементов. Радж будет сопоставлен с именем, следующим элементом по адресу и последним по возрасту.

Выполните код ниже

package main
import "fmt"

//declared the structure named emp
type emp struct {
        name string
        address string
        age int
}

//function which accepts variable of emp type and prints name property
func display(e emp) {
          fmt.Println(e.name)
}

func main() {
// declares a variable, empdata1, of the type emp
var empdata1 emp
//assign values to members of empdata1
empdata1.name = "John"
empdata1.address = "Street\-1, London"
empdata1.age = 30

//declares and assign values to variable empdata2 of type emp
empdata2 := emp{"Raj", "Building\-1, Paris", 25}

//prints the member name of empdata1 and empdata2 using display function
display(empdata1)
display(empdata2)
}

**Вывод**

John
Raj

## Методы (не функции)

Метод — это функция с аргументом получателя. Архитектурно, это между ключевым словом func и именем метода. Синтаксис метода

func (variable variabletype) methodName(parameter1 paramether1type) {
}

Давайте преобразуем приведенный выше пример программы, чтобы использовать методы вместо функции.

package main
import "fmt"

//declared the structure named emp
type emp struct {
    name string
    address string
    age int
}

//Declaring a function with receiver of the type emp
func(e emp) display() {
    fmt.Println(e.name)
}

func main() {
    //declaring a variable of type emp
    var empdata1 emp

    //Assign values to members
    empdata1.name = "John"
    empdata1.address = "Street\-1, Lodon"
    empdata1.age = 30

    //declaring a variable of type emp and assign values to members
    empdata2 := emp {
        "Raj", "Building\-1, Paris", 25}

    //Invoking the method using the receiver of the type emp
   // syntax is variable.methodname()
    empdata1.display()
    empdata2.display()
}

Go не является объектно\-ориентированным языком и не имеет понятия класса. Методы дают представление о том, что вы делаете в объектно\-ориентированных программах, где функции класса вызываются с использованием синтаксиса objectname.functionname ()

## совпадение

Go поддерживает одновременное выполнение задач. Это означает, что Go может выполнять несколько задач одновременно. Это отличается от концепции параллелизма. Параллельно задача разбивается на маленькие подзадачи и выполняется параллельно. Но одновременно, несколько задач выполняются одновременно. Параллелизм достигается в Go с использованием Goroutines и Channels.

## Goroutines

Goroutine — это функция, которая может работать одновременно с другими функциями. Обычно, когда функция вызывается, элемент управления передается в вызываемую функцию, и после завершения ее выполнения управление возвращается к вызывающей функции. Вызывающая функция затем продолжает свое выполнение. Вызывающая функция ожидает, пока вызванная функция завершит выполнение, прежде чем продолжить работу с остальными операторами.

Но в случае goroutine, вызывающая функция не будет ждать завершения вызванной функции. Он будет продолжать выполняться со следующими инструкциями. Вы можете иметь несколько программ в программе.

Кроме того, основная программа завершит работу, как только завершит выполнение своих операторов, и не будет ждать завершения вызванных процедур.

Goroutine вызывается с помощью ключевого слова go, за которым следует вызов функции.

пример

go add(x,y)

Вы разберетесь с рутинами из приведенных ниже примеров. Выполните следующую программу

package main
import "fmt"

func display() {
	for i:=0; i<5; i++ {
		fmt.Println("In display")
	}
}

func main() {
	//invoking the goroutine display()
	go display()
	//The main() continues without waiting for display()
	for i:=0; i<5; i++ {
		fmt.Println("In main")
	}
}

Выход будет

In main
In main
In main
In main
In main

Здесь основная программа завершила выполнение еще до начала программы. Display () — это программа, которая вызывается с использованием синтаксиса

go function\_name(parameter list)

В приведенном выше коде main () не ожидает завершения display (), а main () завершает свое выполнение до того, как display () выполнит свой код. Таким образом, оператор print внутри display () не был напечатан.

Теперь мы модифицируем программу, чтобы также печатать операторы из display (). Мы добавляем задержку в 2 секунды в цикл for функции main () и задержку в 1 секунду в цикле for дисплея ().

package main
import "fmt"
import "time"

func display() {
	for i:=0; i<5; i++ {
		time.Sleep(1 \* time.Second)
		fmt.Println("In display")
	}
}

func main() {
	//invoking the goroutine display()
	go display()
	for i:=0; i<5; i++ {
		time.Sleep(2 \* time.Second)
		fmt.Println("In main")
	}
}

Вывод будет несколько похож на

In display
In main
In display
In display
In main
In display
In display
In main
In main
In main

Здесь вы можете видеть, что оба цикла выполняются перекрывающимся образом из\-за одновременного выполнения.

## каналы

Каналы — это способ взаимодействия функций друг с другом. Его можно рассматривать как средство, в котором одна процедура размещает данные и к которой обращается другая процедура.

Канал может быть объявлен с синтаксисом

channel\_variable := make(chan datatype)

Пример:

	ch := make(chan int)

Вы можете отправить данные на канал, используя синтаксис

channel\_variable <\- variable\_name

пример

    ch <\- x

Вы можете получать данные из канала, используя синтаксис

    variable\_name := <\- channel\_variable

пример

   y := <\- ch

В приведенных выше примерах goroutine вы видели, что основная программа не ждет goroutine. Но это не тот случай, когда каналы вовлечены. Предположим, что если программа отправляет данные в канал, функция main () будет ожидать оператора, получающего данные канала, пока не получит данные.

Вы увидите это в примере ниже. Сначала напишите нормальную программу и посмотрите на поведение. Затем измените программу, чтобы использовать каналы и посмотреть поведение.

Выполните следующую программу

package main
import "fmt"
import "time"

func display() {
	time.Sleep(5 \* time.Second)
	fmt.Println("Inside display()")
}

func main() {
	go display()
	fmt.Println("Inside main()")
}

Выход будет

Inside main()

Функция main () завершила выполнение и завершила работу до выполнения программы. Таким образом, печать внутри дисплея () не была выполнена.

Теперь измените вышеприведенную программу, чтобы использовать каналы и посмотреть поведение.

package main
import "fmt"
import "time"

func display(ch chan int) {
	time.Sleep(5 \* time.Second)
	fmt.Println("Inside display()")
	ch <\- 1234
}

func main() {
	ch := make(chan int)
	go display(ch)
	x := <\-ch
	fmt.Println("Inside main()")
	fmt.Println("Printing x in main() after taking from channel:",x)
}

Выход будет

Inside display()
Inside main()
Printing x in main() after taking from channel: 1234

Вот что происходит: main () при достижении x: = <\-ch будет ожидать данных на канале ch. Дисплей () ожидает 5 секунд, а затем передает данные в канал ch. Функция main () при получении данных из канала разблокируется и продолжает выполнение.

Отправитель, который отправляет данные в канал, может проинформировать получателей о том, что при закрытии канала больше данных не будет добавлено в канал. Это в основном используется, когда вы используете цикл для передачи данных в канал. Канал может быть закрыт с помощью

close(channel\_name)

И на стороне приемника можно проверить, закрыт ли канал, используя дополнительную переменную, при извлечении данных из канала, используя

variable\_name, status := <\- channel\_variable

Если статус Истинный, это означает, что вы получили данные с канала. Если false, это означает, что вы пытаетесь читать с закрытого канала

Вы также можете использовать каналы для общения между программами. Необходимо использовать 2 процедуры — одна отправляет данные в канал, а другая получает данные из канала. Смотрите ниже программу

package main
import "fmt"
import "time"

//This subroutine pushes numbers 0 to 9 to the channel and closes the channel
func add\_to\_channel(ch chan int) {
	fmt.Println("Send data")
	for i:=0; i<10; i++ {
		ch <\- i //pushing data to channel
	}
	close(ch) //closing the channel

}

//This subroutine fetches data from the channel and prints it.
func fetch\_from\_channel(ch chan int) {
	fmt.Println("Read data")
	for {
		//fetch data from channel
x, flag := <\- ch

		//flag is true if data is received from the channel
//flag is false when the channel is closed
if flag == true {
			fmt.Println(x)
		}else{
			fmt.Println("Empty channel")
			break
		}
	}
}

func main() {
	//creating a channel variable to transport integer values
	ch := make(chan int)

	//invoking the subroutines to add and fetch from the channel
	//These routines execute simultaneously
	go add\_to\_channel(ch)
	go fetch\_from\_channel(ch)

	//delay is to prevent the exiting of main() before goroutines finish
	time.Sleep(5 \* time.Second)
	fmt.Println("Inside main()")
}

Здесь есть 2 подпрограммы: одна отправляет данные в канал, а другая печатает данные в канал. Функция add\_to\_channel добавляет числа от 0 до 9 и закрывает канал. Одновременно функция fetch\_from\_channel ожидает в

x, flag: = <\- ch, и как только данные становятся доступными, они печатают данные. Он выходит, как только флаг становится ложным, что означает, что канал закрыт.

Ожидание в main () дается, чтобы предотвратить выход из main () до тех пор, пока горутины не закончат выполнение.

Выполните код и увидите вывод как

Read data
Send data
0
1
2
3
4
5
6
7
8
9
Empty channel
Inside main()

## Выбрать

Выбор можно рассматривать как оператор переключения, который работает на каналах. Здесь операторы case будут операцией канала. Обычно, каждый случай заявления будет считан попыткой чтения из канала. Когда любое из дел готово (канал читается), выполняется оператор, связанный с этим делом. Если несколько случаев готовы, он выберет случайный. Вы можете иметь дело по умолчанию, которое выполняется, если ни одно из дел не готово.

Давайте посмотрим на код ниже

package main
import "fmt"
import "time"

//push data to channel with a 4 second delay
func data1(ch chan string) {
    time.Sleep(4 \* time.Second)
    ch <\- "from data1()"
}

//push data to channel with a 2 second delay
func data2(ch chan string) {
    time.Sleep(2 \* time.Second)
    ch <\- "from data2()"
}

func main() {
    //creating channel variables for transporting string values
    chan1 := make(chan string)
    chan2 := make(chan string)

    //invoking the subroutines with channel variables
    go data1(chan1)
    go data2(chan2)

    //Both case statements wait for data in the chan1 or chan2.
    //chan2 gets data first since the delay is only 2 sec in data2().
    //So the second case will execute and exits the select block
    select {
    case x := <\-chan1:
        fmt.Println(x)
    case y := <\-chan2:
        fmt.Println(y)
    }
}

Выполнение вышеуказанной программы даст вывод:

from data2()

Здесь оператор select ожидает данных, которые будут доступны в любом из каналов. Data2 () добавляет данные в канал после 2 секундного ожидания, что приведет к выполнению второго случая.

Добавьте случай по умолчанию для выбора в той же программе и посмотрите вывод. Здесь при достижении блока выбора, если ни в одном случае нет данных, готовых на канале, он выполнит блок по умолчанию, не дожидаясь доступности данных на каком\-либо канале.

package main
import "fmt"
import "time"

//push data to channel with a 4 second delay
func data1(ch chan string) {
    time.Sleep(4 \* time.Second)
    ch <\- "from data1()"
}

//push data to channel with a 2 second delay
func data2(ch chan string) {
    time.Sleep(2 \* time.Second)
    ch <\- "from data2()"
}

func main() {
    //creating channel variables for transporting string values
    chan1 := make(chan string)
    chan2 := make(chan string)

    //invoking the subroutines with channel variables
    go data1(chan1)
    go data2(chan2)

    //Both case statements check for data in chan1 or chan2.
    //But data is not available (both routines have a delay of 2 and 4 sec)
    //So the default block will be executed without waiting for data in channels.
    select {
    case x := <\-chan1:
        fmt.Println(x)
    case y := <\-chan2:
        fmt.Println(y)
    default:
    	fmt.Println("Default case executed")
    }
}

Эта программа выдаст вывод:

Default case executed

Это связано с тем, что при достижении блока выбора ни у одного канала не было данных для чтения. Итак, случай по умолчанию выполняется.

## Mutex

Мьютекс — это краткая форма взаимного исключения. Mutex используется, когда вы не хотите, чтобы к ресурсу обращались одновременно несколько подпрограмм. У Mutex есть 2 метода — блокировка и разблокировка. Мутекс содержится в пакете синхронизации. Итак, вы должны импортировать пакет синхронизации. Операторы, которые должны выполняться взаимно исключительно, могут быть помещены в mutex.Lock () и mutex.Unlock ().

Давайте изучим мьютекс на примере, который подсчитывает количество выполнений цикла. В этой программе мы ожидаем, что подпрограмма будет запускать цикл 10 раз и счет будет сохранен в сумме. Вы вызываете эту процедуру 3 раза, поэтому общее количество должно быть 30. Счет хранится в глобальной переменной count.

Во\-первых, вы запускаете программу без мьютекса

package main
import "fmt"
import "time"
import "strconv"
import "math/rand"
//declare count variable, which is accessed by all the routine instances
var count = 0

//copies count to temp, do some processing(increment) and store back to count
//random delay is added between reading and writing of count variable
func process(n int) {
	//loop incrementing the count by 10
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Int31n(2)) \* time.Second)
		temp := count
		temp++
		time.Sleep(time.Duration(rand.Int31n(2)) \* time.Second)
		count = temp
	}
	fmt.Println("Count after i="+strconv.Itoa(n)+" Count:", strconv.Itoa(count))
}

func main() {
	//loop calling the process() 3 times
	for i := 1; i < 4; i++ {
		go process(i)
	}

	//delay to wait for the routines to complete
	time.Sleep(25 \* time.Second)
	fmt.Println("Final Count:", count)
}

Увидеть результат

 Count after i=1 Count: 11
Count after i=3 Count: 12
Count after i=2 Count: 13
Final Count: 13

Результат может отличаться при его выполнении, но окончательный результат не будет 30.

Вот что происходит: 3 программы пытаются увеличить количество циклов, хранящихся в переменной count. Предположим, что в данный момент счет равен 5, а goroutine1 собирается увеличить счет до 6. Основные шаги включают

Копировать счетчик в темп

Увеличение температуры

Хранить темп обратно для подсчета

Предположим, вскоре после выполнения шага 3 по goroutine1; другая процедура может иметь старое значение, скажем, 3 выполняет вышеуказанные шаги и хранить 4 обратно, что неправильно. Этого можно избежать, используя мьютекс, который заставляет другие подпрограммы ждать, когда одна подпрограмма уже использует переменную.

Теперь вы запустите программу с мьютексом. Здесь вышеупомянутые 3 шага выполняются в мьютексе.

package main
import "fmt"
import "time"
import "sync"
import "strconv"
import "math/rand"

//declare a mutex instance
var mu sync.Mutex

//declare count variable, which is accessed by all the routine instances
var count = 0

//copies count to temp, do some processing(increment) and store back to count
//random delay is added between reading and writing of count variable
func process(n int) {
	//loop incrementing the count by 10
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Int31n(2)) \* time.Second)
		//lock starts here
		mu.Lock()
		temp := count
		temp++
		time.Sleep(time.Duration(rand.Int31n(2)) \* time.Second)
		count = temp
		//lock ends here
		mu.Unlock()
	}
	fmt.Println("Count after i="+strconv.Itoa(n)+" Count:", strconv.Itoa(count))
}

func main() {
	//loop calling the process() 3 times
	for i := 1; i < 4; i++ {
		go process(i)
	}

	//delay to wait for the routines to complete
	time.Sleep(25 \* time.Second)
	fmt.Println("Final Count:", count)
}

Теперь вывод будет

 Count after i=3 Count: 21
Count after i=2 Count: 28
Count after i=1 Count: 30
Final Count: 30

Здесь мы получаем ожидаемый результат в качестве конечного результата. Поскольку операторы чтения, приращения и обратной записи счетчика выполняются в мьютексе.

## Обработка ошибок

Ошибки — это ненормальные условия, такие как закрытие файла, который не открыт, открытие файла, который не существует, и т. Д. Функции обычно возвращают ошибки как последнее возвращаемое значение.

Пример ниже объясняет больше об ошибке.

package main
import "fmt"
import "os"

//function accepts a filename and tries to open it.
func fileopen(name string) {
    f, er := os.Open(name)

    //er will be nil if the file exists else it returns an error object
    if er != nil {
        fmt.Println(er)
        return
    }else{
    	fmt.Println("file opened", f.Name())
    }
}

func main() {
    fileopen("invalid.txt")
}

Выход будет:

open /invalid.txt: no such file or directory

Здесь мы попытались открыть несуществующий файл, и он вернул ошибку в переменную er. Если файл действителен, то ошибка будет нулевой

## Пользовательские ошибки

Используя эту функцию, вы можете создавать собственные ошибки. Это делается с помощью New () пакета ошибок. Мы перепишем вышеупомянутую программу, чтобы использовать пользовательские ошибки.

Запустите программу ниже

package main
import "fmt"
import "os"
import "errors"

//function accepts a filename and tries to open it.
func fileopen(name string) (string, error) {
    f, er := os.Open(name)

    //er will be nil if the file exists else it returns an error object
    if er != nil {
        //created a new error object and returns it
        return "", errors.New("Custom error message: File name is wrong")
    }else{
    	return f.Name(),nil
    }
}

func main() {
    //receives custom error or nil after trying to open the file
    filename, error := fileopen("invalid.txt")
    if error != nil {
        fmt.Println(error)
    }else{
    	fmt.Println("file opened", filename)
    }
}

The output will be:

Custom error message:File name is wrong

Here the area() returns the area of a square. If the input is less than 1 then area() returns an error message.

## Reading files

Files are used to store data. Go allows us to read data from the files

First create a file, data.txt, in your present directory with the below content.

Line one
Line two
Line three

Now run the below program to see it prints the contents of the entire file as output

package main
import "fmt"
import "io/ioutil"

func main() {
    data, err := ioutil.ReadFile("data.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    fmt.Println("Contents of file:", string(data))
}

Here the data, err := ioutil.ReadFile(«data.txt») reads the data and returns a byte sequence. While printing it is converted to string format.

## Writing files

You will see this with a program

package main
import "fmt"
import "os"

func main() {
    f, err := os.Create("file1.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    l, err := f.WriteString("Write Line one")
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    fmt.Println(l, "bytes written")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
}

Здесь создается файл test.txt. Если файл уже существует, его содержимое усекается. Writeline () используется для записи содержимого в файл. После этого Вы закрыли файл с помощью Close ().

## Шпаргалка

В этом уроке Go мы рассмотрели

| **Тема** | **Описание** | **Синтаксис** |
| Основные типы | Числовой, строковый, bool |  |
| переменные | Объявите и присвойте значения переменным | var имя\_переменной тип var имя\_переменной тип = значение var имя\_переменной1, имя\_переменной2 = значение1, значение2 имя\_переменной: = значение |
| Константы | Переменные, значение которых нельзя изменить после присвоения | константная переменная = значение |
| Для петли | Выполнить операторы в цикле. | для инициализации\_экспрессия; evaluation\_expression; iteration\_expression {// один или несколько операторов} |
| Если еще | Это условное утверждение | условие if {// Statement\_1} else {// Statement\_2} |
| переключатель | Условное утверждение с несколькими случаями | Выражение переключателя {case value\_1: Statement\_1 Case Value\_2: Statement\_2 Case Value\_n: Statement\_n По умолчанию: Statement\_default} |
| массив | Фиксированный размер именованной последовательности элементов одного типа | имя\_прибора: = \[размер\] тип {значение\_0, значение\_1,…, значение\_размер\-1} |
| Кусочек | Часть или сегмент массива | var slice\_name \[\] type = array\_name \[start: end\] |
| функции | Блок утверждений, который выполняет определенную задачу | func имя\_функции (тип параметра\_1, тип параметра\_n) return\_type {// операторы} |
| пакеты | Используются для организации кода. Увеличивает читаемость и возможность повторного использования кода | import package\_nam |
| Перенести | Откладывает выполнение функции, пока содержащая функция не закончит выполнение | defer имя\_функции (список\_параметров) |
| указатели | Хранит адрес памяти другой переменной. | var variable\_name \* type |
| Структура | Определяемый пользователем тип данных, который сам содержит еще один элемент того же или другого типа | тип structname struct {variable\_1 variable\_1\_type variable\_2 variable\_2\_type variable\_n variable\_n\_type} |
| методы | Метод — это функция с аргументом получателя. | func (variable variabletype) methodName (parameter\_list) {} |
| Goroutine | Функция, которая может работать одновременно с другими функциями. | перейти имя\_функции (список\_параметров) |
| канал | Способ для функций общаться друг с другом. Среда, на которую одна подпрограмма помещает данные и доступ к которой осуществляется другой подпрограммой. | Объявление: ch: = make (chan int) Отправка данных на канал: channel\_variable <\- имя переменной. Получение от канала: имя переменной: = <\- переменная канала. |
| Выбрать | Переключатель оператора, который работает на каналах. Операторы case будут операцией канала. Когда какой\-либо канал готов с данными, выполняется оператор, связанный с этим случаем. | выберите {case x: = <\-chan1: fmt.Println (x) case y: = <\-chan2: fmt.Println (y)} |
| Mutex | Mutex используется, когда вы не хотите, чтобы к ресурсу обращались одновременно несколько подпрограмм. Mutex имеет 2 метода — блокировка и разблокировка | mutex.Lock () // операторы mutex.Unlock (). |
| Читать файлы | Читает данные и возвращает последовательность байтов. | Данные, err: = ioutil.ReadFile (имя файла) |
| Написать файл | Записывает данные в файл | l, err: = f.WriteString (text\_to\_write) |
