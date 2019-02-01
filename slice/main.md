## Введение в SLICE 

### Срез 
- это сегмент массива. Срезы основаны на массивах и обеспечивают большую мощность, гибкость и удобство по сравнению с массивами.

Как и массивы, ломтики индексируются и имеют длину. Но в отличие от массивов, они могут быть изменены.

Внутренне, Slice - это просто ссылка на базовый массив. В этой статье мы узнаем, как создавать и использовать фрагменты, а также узнаем, как они работают под капотом.

### Объявление SLICE
Срез типа Tобъявлен с использованием []T. Например, вот как вы можете объявить фрагмент типа int-

```golang
// Slice of type `int`
var s []int
```

Срез объявляется так же, как массив, за исключением того, что мы не указываем размер в скобках [].


 
### Создание и инициализация среза
1. Создание среза с использованием литерала среза
Вы можете создать срез с помощью литерала среза, как этот -

```golang
// Creating a slice using a slice literal
var s = []int{3, 5, 7, 9, 11, 13, 17}
```

Выражение в правой части приведенного выше оператора является литералом фрагмента. Литерал фрагмента объявляется так же, как литерал массива , за исключением того, что вы не указываете размер в квадратных скобках [].

Когда вы создаете срез с использованием литерала среза, он сначала создает массив, а затем возвращает ссылку на него.

Давайте посмотрим полный пример -

```golang
package main
import "fmt"

func main() {
	// Creating a slice using a slice literal
	var s = []int{3, 5, 7, 9, 11, 13, 17}

	// Short hand declaration
	t := []int{2, 4, 8, 16, 32, 64}

	fmt.Println("s = ", s)
	fmt.Println("t = ", t)
}
```

 Output
s =  [3 5 7 9 11 13 17]
t =  [2 4 8 16 32 64]

### 2. Создание среза из массива
Поскольку срез является сегментом массива, мы можем создать срез из массива.

Чтобы создать срез из массива a, мы указываем два индекса low(нижняя граница) и high(верхняя граница), разделенные двоеточием -

```golang
// Obtaining a slice from an array `a`
a[low:high]
```

Вышеупомянутое выражение выбирает срез из массива a. Результирующий фрагмент включает в себя все элементы, начиная с индекса lowдо high, но исключая элемент с индексомhigh .

Давайте рассмотрим пример, чтобы прояснить ситуацию -

```golang
package main
import "fmt"

func main() {
	var a = [5]string{"Alpha", "Beta", "Gamma", "Delta", "Epsilon"}

	// Creating a slice from the array
	var s []string = a[1:4]

	fmt.Println("Array a = ", a)
	fmt.Println("Slice s = ", s)
}
Array a =  [Alpha Beta Gamma Delta Epsilon]
Slice s =  [Beta Gamma Delta]
```

В low и highиндексы в выражении среза не являются обязательными. Значение по умолчанию для lowэто 0, и highдлина среза.

```golang
package main
import "fmt"

func main() {
	a := [5]string{"C", "C++", "Java", "Python", "Go"}

	slice1 := a[1:4]
	slice2 := a[:3]
	slice3 := a[2:]
	slice4 := a[:]

	fmt.Println("Array a = ", a)
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
	fmt.Println("slice3 = ", slice3)
	fmt.Println("slice4 = ", slice4)
}
```

Output  
Array a =  [C C++ Java Python Go]
slice1 =  [C++ Java Python]
slice2 =  [C C++ Java]
slice3 =  [Java Python Go]
slice4 =  [C C++ Java Python Go]

### 3. Создание среза из другого среза
Срез также может быть создан путем нарезки существующего среза.

```golang
package main
import "fmt"

func main() {
	cities := []string{"New York", "London", "Chicago", "Beijing", "Delhi", "Mumbai", "Bangalore", "Hyderabad", "Hong Kong"}

	asianCities := cities[3:]
	indianCities := asianCities[1:5]

	fmt.Println("cities = ", cities)
	fmt.Println("asianCities = ", asianCities)
	fmt.Println("indianCities = ", indianCities)
}
```

Output
cities =  [New York London Chicago Beijing Delhi Mumbai Bangalore Hyderabad Hong Kong]
asianCities =  [Beijing Delhi Mumbai Bangalore Hyderabad Hong Kong]
indianCities =  [Delhi Mumbai Bangalore Hyderabad]

### Модификация среза
Ломтики являются ссылочными типами. Они ссылаются на базовый массив. Изменение элементов среза приведет к изменению соответствующих элементов в указанном массиве. Другие срезы, которые ссылаются на тот же массив, также увидят эти изменения.

```golang
package main
import "fmt"

func main() {
	a := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	slice1 := a[1:]
	slice2 := a[3:]

	fmt.Println("------- Before Modifications -------")
	fmt.Println("a  = ", a)
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)

	slice1[0] = "TUE"
	slice1[1] = "WED"
	slice1[2] = "THU"

	slice2[1] = "FRIDAY"

	fmt.Println("\n-------- After Modifications --------")
	fmt.Println("a  = ", a)
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
}
```

Output   
------- Before Modifications -------
a  =  [Mon Tue Wed Thu Fri Sat Sun]
slice1 =  [Tue Wed Thu Fri Sat Sun]
slice2 =  [Thu Fri Sat Sun]

-------- After Modifications --------
a  =  [Mon TUE WED THU FRIDAY Sat Sun]
slice1 =  [TUE WED THU FRIDAY Sat Sun]
slice2 =  [THU FRIDAY Sat Sun]

 
### Длина и емкость среза

Срез состоит из трех вещей -

Указатель (ссылка) к нижележащему массива.
Длина отрезка массива , что срез содержит.
Емкость (максимальный размер , до которого может вырасти сегмент).

### Иллюстрация ломтиков Голанга
Рассмотрим в качестве примера следующий массив и полученный из него фрагмент:

```golang
var a = [6]int{10, 20, 30, 40, 50, 60}
var s = [1:4]
```

Вот как представлен фрагмент sв приведенном выше примере:

### Длина и вместимость кусочков Голанга

***Длина среза*** - это количество элементов в срезе, как 3в приведенном выше примере.
***Емкость*** - это количество элементов в базовом массиве, начиная с первого элемента в срезе. это5 в приведенном выше примере.

Вы можете найти длину и емкость среза, используя встроенные функции len()и cap()-

```golang
package main
import "fmt"

func main() {
	a := [6]int{10, 20, 30, 40, 50, 60}
	s := a[1:4]

	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
}
```

Output
s = [20 30 40], len = 3, cap = 5

Длина среза может быть увеличена до его емкости путем повторной нарезки. Любая попытка расширить его длину за пределы доступной емкости приведет к ошибке времени выполнения.

Посмотрите на следующий пример, чтобы понять, как повторная нарезка данного фрагмента изменяет его длину и емкость -

```golang
package main
import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("Original Slice")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	s = s[1:5]
	fmt.Println("\nAfter slicing from index 1 to 5")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	s = s[:8]
	fmt.Println("\nAfter extending the length")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	s = s[2:]
	fmt.Println("\nAfter dropping the first two elements")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
}
```

 Output
Original Slice
s = [10 20 30 40 50 60 70 80 90 100], len = 10, cap = 10

After slicing from index 1 to 5
s = [20 30 40 50], len = 4, cap = 9

After extending the length
s = [20 30 40 50 60 70 80 90], len = 8, cap = 9

After dropping the first two elements
s = [40 50 60 70 80 90], len = 6, cap = 7

### Создание среза с помощью встроенной make()функции

Теперь, когда мы знаем о длине и емкости среза. Давайте посмотрим на другой способ создания среза.

Golang предоставляет библиотечную функцию make()для вызова слайсов. Ниже подпись make()функции -

```golang
func make([]T, len, cap) []T
```

Функция make принимает тип, длину и дополнительную емкость. Он выделяет базовый массив с размером, равным заданной емкости, и возвращает фрагмент, который ссылается на этот массив.

```golang
package main
import "fmt"

func main() {
	// Creates an array of size 10, slices it till index 5, and returns the slice reference
	s := make([]int, 5, 10)
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
}
```

Output
s = [0 0 0 0 0], len = 5, cap = 10

Параметр производительности в make()функции является необязательным. Если опущено, по умолчанию используется указанная длина -

```golang
package main
import "fmt"

func main() {
	// Creates an array of size 5, and returns a slice reference to it
	s := make([]int, 5)
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
}
```

Output
s = [0 0 0 0 0], len = 5, cap = 5

### Нулевое значение ломтиков
Нулевое значение из среза nil. Нулевой срез не имеет базового массива, а имеет длину и емкость 0-

```golang
package main
import "fmt"

func main() {
	var s []int
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	if s == nil {
		fmt.Println("s is nil")
	}
}
```

Output
s = [], len = 0, cap = 0
s is nil

### Функции среза

1. Функция copy (): копирование фрагмента   
Эти copy()функции копируют элементы из одного среза в другой. Его подпись выглядит так -

```golang
func copy(dst, src []T) int
```

Требуется два среза - целевой срез и исходный срез. Затем он копирует элементы из источника в место назначения и возвращает количество скопированных элементов.

Обратите внимание, что элементы копируются, только если целевой раздел имеет достаточную емкость.

```golang
package main
import "fmt"

func main() {
	src := []string{"Sublime", "VSCode", "IntelliJ", "Eclipse"}
	dest := make([]string, 2)

	numElementsCopied := copy(dest, src)

	fmt.Println("src = ", src)
	fmt.Println("dest = ", dest)
	fmt.Println("Number of elements copied from src to dest = ", numElementsCopied)
}
```

Output   
src  =  [Sublime VSCode IntelliJ Eclipse]
dest =  [Sublime VSCode]
Number of elements copied from src to dest =  2


### 2. Функция append (): добавление к фрагменту
append() Функция добавляет новые элементы в конце данной секции. Ниже приведена подпись appendфункции.

```golang
func append(s []T, x ...T) []T
```

Она принимает кусочек и переменное число аргументов х ... T . Затем он возвращает новый фрагмент, содержащий все элементы из данного фрагмента, а также новые элементы.

Если данный срез не обладает достаточной емкостью для размещения новых элементов, тогда выделяется новый базовый массив с большей емкостью. Все элементы из базового массива существующего среза копируются в этот новый массив, а затем добавляются новые элементы.

Однако, если срез имеет достаточную емкость для размещения новых элементов, то append() функция повторно использует свой базовый массив и добавляет новые элементы в тот же массив.

Давайте посмотрим на пример, чтобы лучше понять вещи -

```golang
package main
import "fmt"

func main() {
	slice1 := []string{"C", "C++", "Java"}
	slice2 := append(slice1, "Python", "Ruby", "Go")

	fmt.Printf("slice1 = %v, len = %d, cap = %d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 = %v, len = %d, cap = %d\n", slice2, len(slice2), cap(slice2))

	slice1[0] = "C#"
	fmt.Println("\nslice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
}
```

Output   
slice1 = [C C++ Java], len = 3, cap = 3
slice2 = [C C++ Java Python Ruby Go], len = 6, cap = 6
slice1 =  [C# C++ Java]
slice2 =  [C C++ Java Python Ruby Go]

В приведенном выше примере, поскольку slice1имеет емкость 3, он не может вместить больше элементов. Таким образом, новый основной массив выделяется с большей емкостью, когда мы добавляем к нему больше элементов.

Так что, если вы измените slice1, slice2не увидите эти изменения, потому что это относится к другому массиву.

Но что, если бы slice1было достаточно возможностей для размещения новых элементов? Ну, в этом случае новый массив не будет выделен, и элементы будут добавлены в тот же базовый массив.

Кроме того, в этом случае изменения в slice1также повлияют, slice2поскольку оба будут ссылаться на один и тот же базовый массив. Это продемонстрировано в следующем примере -

```golang
package main
import "fmt"

func main() {
	slice1 := make([]string, 3, 10)
	copy(slice1, []string{"C", "C++", "Java"})

	slice2 := append(slice1, "Python", "Ruby", "Go")

	fmt.Printf("slice1 = %v, len = %d, cap = %d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 = %v, len = %d, cap = %d\n", slice2, len(slice2), cap(slice2))

	slice1[0] = "C#"
	fmt.Println("\nslice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
}
```

Output  
slice1 = [C C++ Java], len = 3, cap = 10
slice2 = [C C++ Java Python Ruby Go], len = 6, cap = 10

slice1 =  [C# C++ Java]
slice2 =  [C# C++ Java Python Ruby Go]

### Добавление к нулевому срезу

Когда вы добавляете значения к nilсрезу, он выделяет новый срез и возвращает ссылку на новый срез.

```golang
package main
import "fmt"

func main() {
	var s []string

	// Appending to a nil slice
	s = append(s, "Cat", "Dog", "Lion", "Tiger")

	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
}
```

Output  
s = [Cat Dog Lion Tiger], len = 4, cap = 4


### Добавление одного среза к другому
Вы можете напрямую добавить один фрагмент к другому с помощью ...оператора. Этот оператор расширяет срез до списка аргументов. Следующий пример демонстрирует его использование -

```golang
package main
import "fmt"

func main() {
	slice1 := []string{"Jack", "John", "Peter"}
	slice2 := []string{"Bill", "Mark", "Steve"}

	slice3 := append(slice1, slice2...)

	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
	fmt.Println("After appending slice1 & slice2 = ", slice3)
}
```

Output
slice1 =  [Jack John Peter]
slice2 =  [Bill Mark Steve]
After appending slice1 & slice2 =  [Jack John Peter Bill Mark Steve]

### Ломтик ломтика
Ломтики могут быть любого типа. Они также могут содержать другие ломтики. В приведенном ниже примере создается фрагмент среза -

```golang
package main

import "fmt"

func main() {
	s := [][]string{
		{"India", "China"},
		{"USA", "Canada"},
		{"Switzerland", "Germany"},
	}

	fmt.Println("Slice s = ", s)
	fmt.Println("length = ", len(s))
	fmt.Println("capacity = ", cap(s))
}
```

Output
Slice s =  [[India China] [USA Canada] [Switzerland Germany]]
length =  3
capacity =  3

### Итерация по срезу
Вы можете выполнять итерацию по срезу так же, как и по массиву. Ниже приведены два способа перебора фрагмента:

1. Перебор фрагмента с использованием ***for*** цикла

```golang
package main
import "fmt"

func main() {
	countries := []string{"India", "America", "Russia", "England"}

	for i := 0; i < len(countries); i++ {
		fmt.Println(countries[i])
	}
}
```

Output  
India
America
Russia
England

### Перебор фрагмента с использованием rangeформы ***for*** цикла

```golang
package main
import "fmt"

func main() {
	primeNumbers := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	for index, number := range primeNumbers {
		fmt.Printf("PrimeNumber(%d) = %d\n", index+1, number)
	}
}
```

Output   
PrimeNumber(1) = 2
PrimeNumber(2) = 3
PrimeNumber(3) = 5
PrimeNumber(4) = 7
PrimeNumber(5) = 11
PrimeNumber(6) = 13
PrimeNumber(7) = 17
PrimeNumber(8) = 19
PrimeNumber(9) = 23
PrimeNumber(10) = 29

Игнорирование ***index*** из rangeформы forцикла с использованием пустого идентификатора
range Форма for петли дает вам indexи valueпо этому показателю в каждой итерации. 
Если вы не хотите использовать index, то вы можете отказаться от него, используя подчеркивание _.

Подчеркивание ( _) называется пустым идентификатором. Он используется, чтобы сказать компилятору Go, что нам не нужно это значение.

```golang
package main
import "fmt"

func main() {
	numbers := []float64{3.5, 7.4, 9.2, 5.4}

	sum := 0.0
	for _, number := range numbers {
		sum += number
	}

	fmt.Printf("Total Sum = %.2f\n", sum)
}
```

Output  
Total Sum = 25.50

### Заключение

В этой статье вы узнали, как создавать срезы, как срезы работают внутренне, 
как использовать встроенные функции copy()иappend() выращивать срезы.
