# Как я могу преобразовать строковую переменную в тип Boolean, Integer или Float в Golang?

В следующем фрагменте исходного кода показана сокращенная программа с несколькими значениями переменных String, проанализированными в Float, Boolean и Integer.

Синтаксический анализ функции `ParseBool` , `ParseFloat` и `ParseInt` преобразования строк значений. Следовательно, значения переменных String сохраняются как Boolean, Float и Integer в соответствующих типах данных.

[?](#)

|

1

2

3

4

5

6

7

8

9

10

11

12

13

14

15

16

17

18

19

20

21

22

23

24

25

26

27

28

 |

`package` `main`

`import` `(   `

`"fmt"`

`"reflect"`

`"strconv"`

`)`

`func` `main() {  `

`fmt.Println(``"\nConvert String into Boolean Data type\n"``)`

`str1 :=` `"japan"`

`fmt.Println(``"Before :"``, reflect.TypeOf(str1))`

`bolStr,_ := strconv.ParseBool(str1)`

`fmt.Println(``"After :"``, reflect.TypeOf(bolStr)) `

`fmt.Println(``"\nConvert String into Float64 Data type\n"``)`

`str2 :=` `"japan"`

`fmt.Println(``"Before :"``, reflect.TypeOf(str2))      `

`fltStr,_ := strconv.ParseFloat(str2,``64``)`

`fmt.Println(``"After :"``, reflect.TypeOf(fltStr)) `

`fmt.Println(``"\nConvert String into Integer Data type\n"``)`

`str3 :=` `"japan"`

`fmt.Println(``"Before :"``, reflect.TypeOf(str3))`

`intStr,_ := strconv.ParseInt(str3,``10``,``64``)`

`fmt.Println(``"After :"``, reflect.TypeOf(intStr))`

`}`

 |

C: \\ golang \\ codes> go run example2.go

Преобразовать строку в логический тип данных

Before: string
After: bool

Преобразовать строку в Float64 Тип данных

Before: string
After: float64

Преобразовать строку в целочисленный тип данных

Before: string
After: int64

C: \\ golang \\ коды>
