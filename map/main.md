## Map на примере

***Карта*** - это неупорядоченная коллекция пар ключ-значение. Он сопоставляет ключи со значениями. Ключи являются уникальными в пределах карты, в то время как значения могут не быть.

Структура данных карты используется для быстрого поиска, поиска и удаления данных на основе ключей. Это одна из наиболее часто используемых структур данных в информатике.

Go предоставляет встроенный тип карты. В этой статье мы узнаем, как использовать встроенный тип карты Голанга.

### Объявление карты
Карта объявляется с использованием следующего синтаксиса -

```golang
var m map[KeyType]ValueType
```

Например, вот как вы можете объявить карту stringключей к intзначениям -

```golang
var m map[string]int
```

Нулевое значение из карты является nil. nilКарта не имеет ключей. Более того, любая попытка добавить ключи на nilкарту приведет к ошибке во время выполнения.

Давайте посмотрим на пример

```golang
package main
import "fmt"

func main() {
	var m map[string]int
	fmt.Println(m)
	if m == nil {
		fmt.Println("m is nil")
	}

	// Attempting to add keys to a nil map will result in a runtime error
	// m["one hundred"] = 100
}
```

Output
map[]
m is nil

Если вы раскомментируете оператор m["one hundred"] = 100, программа выдаст следующую ошибку:

panic: assignment to entry in nil map
Поэтому необходимо инициализировать карту перед добавлением к ней элементов.


 
### Инициализация карты

1. Инициализация карты с использованием встроенной ***make()*** функции
Вы можете инициализировать карту, используя встроенную make() функцию. 
Вам просто нужно передать тип карты в make() функцию, как в примере ниже. 
Функция вернет инициализированную и готовую к использованию карту 

```golang
// Initializing a map using the built-in make() function

var m = make(map[string]int)
Давайте посмотрим полный пример -

package main
import "fmt"

func main() {
	var m = make(map[string]int)

	fmt.Println(m)

	if m == nil {
		fmt.Println("m is nil")
	} else {
		fmt.Println("m is not nil")
	}

	// make() function returns an initialized and ready to use map.
	// Since it is initialized, you can add new keys to it.
	m["one hundred"] = 100
	fmt.Println(m)
}
```

Output
map[]
m is not nil
map[one hundred:100]

### Инициализация карты с использованием литерала карты
***Литерал карты*** - это очень удобный способ инициализации карты с некоторыми данными. 
Вам просто нужно передать пары ключ-значение, разделенные двоеточием, внутри фигурных скобок, например:

```golang
var m = map[string]int{
	"one": 1,
	"two": 2,
	"three": 3,
}
```

Обратите внимание, что последняя запятая необходима, в противном случае вы получите ошибку компилятора.

Давайте посмотрим полный пример -

```golang
package main
import "fmt"

func main() {
	var m = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5, // Comma is necessary
	}

	fmt.Println(m)
}
```

Output
map[one:1 two:2 three:3 four:4 five:5]

Вы также можете создать пустую карту, используя литерал карты, оставив фигурные скобки пустыми -

```golang
// Initialize an empty map
var m = map[string]int{}
```

Вышеупомянутое утверждение функционально идентично использованию make() функции.

### Добавление элементов (пар ключ-значение) на карту

Вы можете добавить новые элементы в инициализированную карту, используя следующий синтаксис -

```golang
m[key] = value
```

В следующем примере инициализируется карта с помощью make()функции и добавляются некоторые новые элементы.

```golang
package main
import "fmt"

func main() {
	// Initializing a map
	var tinderMatch = make(map[string]string)

	// Adding keys to a map
	tinderMatch["Rajeev"] = "Angelina" // Assigns the value "Angelina" to the key "Rajeev"
	tinderMatch["James"] = "Sophia"
	tinderMatch["David"] = "Emma"

	fmt.Println(tinderMatch)

	/*
	  Adding a key that already exists will simply override
	  the existing key with the new value
	*/
	tinderMatch["Rajeev"] = "Jennifer"
	fmt.Println(tinderMatch)
}
```

Output  
map[Rajeev:Angelina James:Sophia David:Emma]
map[Rajeev:Jennifer James:Sophia David:Emma]


Если вы попытаетесь добавить ключ, который уже существует на карте, он будет просто переопределен новым значением.
 
### Получение значения, связанного с данным ключом на карте


Вы можете получить значение, назначенное ключу на карте, используя синтаксис m[key]. Если ключ существует на карте, вы получите назначенное значение. В противном случае вы получите нулевое значение типа значения карты.

Давайте посмотрим на пример, чтобы понять это -

```golang
package main
import "fmt"

func main() {
	var personMobileNo = map[string]string{
		"John":  "+33-8273658526",
		"Steve": "+1-8579822345",
		"David": "+44-9462834443",
	}

	var mobileNo = personMobileNo["Steve"]
	fmt.Println("Steve's Mobile No : ", mobileNo)

	// If a key doesn't exist in the map, we get the zero value of the value type
	mobileNo = personMobileNo["Jack"]
	fmt.Println("Jack's Mobile No : ", mobileNo)
}
```

# Output
Steve's Mobile No :  +1-8579822345
Jack's Mobile No : 

В приведенном выше примере, поскольку ключ "Jack"не существует на карте, мы получаем нулевое значение типа значения карты.
Так как тип значения карты - stringмы получаем " ".

В отличие от других языков, мы не получаем ошибку времени выполнения в Golang, если ключ не существует на карте.

Но что, если вы хотите проверить наличие ключа? В приведенном выше примере карта вернется, " "даже если ключ "Jack"существует со значением " ". Итак, как мы можем различать случаи, когда существует ключ со значением, равным нулевому значению типа значения, и отсутствием ключа?

Что ж, давайте узнаем.

### Проверка наличия ключа на карте
Когда вы извлекаете значение, назначенное данному ключу, используя синтаксис map[key], 
он также возвращает дополнительное логическое значение, которое существует, true если ключ существует на карте, 
и false если он не существует.

Таким образом, вы можете проверить наличие ключа на карте с помощью следующего двухзначного присваивания:

value, ok := m[key]

Булева переменная ok будет, true если ключ существует, и в false противном случае.

Для примера рассмотрим следующую карту. Он отображает идентификаторы сотрудников для имен -

```golang
var employees = map[int]string{
	1001: "Rajeev",
	1002: "Sachin",
	1003: "James",
}
````

Доступ к ключу 1001вернется "Rajeev"и true, поскольку ключ 1001существует на карте -

```golang
name, ok := employees[1001]  // "Rajeev", true
````

Однако, если вы попытаетесь получить доступ к ключу, который не существует, то карта вернет пустую строку ""(нулевое значение строк), и false-

```golang
name, ok := employees[1010]  // "", false
```

Если вы просто хотите проверить наличие ключа без извлечения значения, связанного с этим ключом, тогда вы можете использовать _(подчеркивание) вместо первого значения -

```golang
_, ok := employees[1005]
```

Теперь давайте проверим полный пример -

```golang
package main
import "fmt"

func main() {
	var employees = map[int]string{
		1001: "John",
		1002: "Steve",
		1003: "Maria",
	}

	printEmployee(employees, 1001)
	printEmployee(employees, 1010)

	if isEmployeeExists(employees, 1002) {
		fmt.Println("EmployeeId 1002 found")
	}
}

func printEmployee(employees map[int]string, employeeId int) {
	if name, ok := employees[employeeId]; ok {
		fmt.Printf("name = %s, ok = %v\n", name, ok)
	} else {
		fmt.Printf("EmployeeId %d not found\n", employeeId)
	}
}

func isEmployeeExists(employees map[int]string, employeeId int) bool {
	_, ok := employees[employeeId]
	return ok
}
```

Output   
name = Rajeev, ok = true
EmployeeId 1010 not found
EmployeeId 1002 found

В приведенном выше примере, я использовал краткую декларацию в ifзаявлении для инициализации nameи okзначения, а затем проверить логическое значение ok. Это делает код более лаконичным.

### Удаление ключа с карты

Вы можете удалить ключ с карты, используя встроенную delete()функцию. Синтаксис выглядит так -

```golang
// Delete the `key` from the `map`
delete(map, key)
```

delete() Функция не возвращает никакого значения. Кроме того, он ничего не делает, если ключ не существует на карте.

Вот полный пример

```golang
package main

import "fmt"

func main() {
	var fileExtensions = map[string]string{
		"Python": ".py",
		"C++":    ".cpp",
		"Java":   ".java",
		"Golang": ".go",
		"Kotlin": ".kt",
	}

	fmt.Println(fileExtensions)

	delete(fileExtensions, "Kotlin")

	// delete function doesn't do anything if the key doesn't exist
	delete(fileExtensions, "Javascript")

	fmt.Println(fileExtensions)
}
```

Output  
map[Python:.py C++:.cpp Java:.java Golang:.go Kotlin:.kt]
map[Python:.py C++:.cpp Java:.java Golang:.go]

### Карты являются ссылочными типами

Карты являются ссылочными типами. Когда вы назначаете карту новой переменной, они оба ссылаются на одну и ту же базовую структуру данных. Поэтому изменения, сделанные одной переменной, будут видны другой -

```golang
package main
import "fmt"

func main() {
	var m1 = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}

	var m2 = m1
	fmt.Println("m1 = ", m1)
	fmt.Println("m2 = ", m2)

	m2["ten"] = 10
	fmt.Println("\nm1 = ", m1)
	fmt.Println("m2 = ", m2)
}
```

Output
m1 =  map[one:1 two:2 three:3 four:4 five:5]
m2 =  map[one:1 two:2 three:3 four:4 five:5]

m1 =  map[one:1 two:2 three:3 four:4 five:5 ten:10]
m2 =  map[one:1 two:2 three:3 four:4 five:5 ten:10]
Та же концепция применяется, когда вы передаете карту в функцию. Любые изменения, внесенные в карту внутри функции, также видны вызывающей стороне.

### Итерация по карте

Вы можете перебрать карту, используя rangeформу цикла for. Это дает вам key, valueпару в каждой итерации 

```golang
package main
import "fmt"

func main() {
	var personAge = map[string]int{
		"Rajeev": 25,
		"James":  32,
		"Sarah":  29,
	}

	for name, age := range personAge {
		fmt.Println(name, age)
	}

}
```

Output   
James 32
Sarah 29
Rajeev 25

Обратите внимание, что карта - это неупорядоченная коллекция, и поэтому порядок итераций карты не всегда будет одинаковым при каждой итерации по ней.

Поэтому, если вы запустите вышеуказанную программу несколько раз, вы получите результаты в разных порядках.

### Заключение
В этой статье вы узнали, как объявлять и инициализировать карты, как добавлять ключи к карте, 
как извлечь значение, связанное с данным ключом на карте, как проверить наличие ключа на карте, 
как удалить ключ с карты, и как перебрать карту.
