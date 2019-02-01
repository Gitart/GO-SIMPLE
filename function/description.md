## Введение в функции в Голанге

### Функция 
это блок кода, который принимает некоторые входные данные, выполняет некоторую обработку входных данных и производит некоторые выходные данные.

### Иллюстрация функций Голанга
Функции помогают вам разделить вашу программу на небольшие повторяющиеся фрагменты кода. Они улучшают читаемость, удобство обслуживания и тестируемость вашей программы.

### Объявление и вызов функций в Голанге
В Golang мы объявляем функцию, используя funcключевое слово. Функция имеет имя , список входных параметров, разделенных запятыми, а также их типы, тип (ы) результата и тело .

Ниже приведен пример простой вызываемой функции, avgкоторая принимает два входных параметра типа float64и возвращает среднее значение входных данных. Результат также имеет тип float64-

```golang
func avg(x float64, y float64) float64 {
	return (x + y) / 2
}
```

Теперь вызвать функцию очень просто. Вам просто нужно передать требуемое количество параметров в функцию, как это -

avg(6.56, 13.44)

Вот полный пример -

```golang
package main
import "fmt"

func avg(x float64, y float64) float64 {
	return (x + y) / 2
}

func main() {
	x := 5.75
	y := 6.25

	result := avg(x, y)

	fmt.Printf("Average of %.2f and %.2f = %.2f\n", x, y, result)
}
```

Output
Average of 5.75 and 6.25 = 6.00

### Параметры функции и возвращаемый тип (ы) являются необязательными

Входные параметры и возвращаемые типы являются необязательными для функции. Функция может быть объявлена ​​без какого-либо ввода и вывода.


main()Функция является примером такой функции -

```golang
func main() {
}
```

Вот еще один пример -

```golang
func sayHello() {
	fmt.Println("Hello, World")
}
```

Необходимо указать тип только один раз для нескольких последовательных параметров одного типа

Если функция имеет два или более последовательных параметра одного и того же типа, то достаточно указать тип только один раз для последнего параметра этого типа.

Например, мы можем объявить avgфункцию, которую мы видели в предыдущем разделе, также:

```golang
func avg(x, y float64) float64 { }
// Same as - func avg(x float64, y float64) float64 { }
```

Вот еще один пример -

```golang
func printPersonDetails(firstName, lastName string, age int) { }
// Same as - func printPersonDetails(firstName string, lastName string, age int) { }
```

 
### Функции с несколькими возвращаемыми значениями
Функции Go могут возвращать несколько значений. Вот так! Это то, что большинство языков программирования не поддерживают изначально. Но Го отличается.

Допустим, вы хотите создать функцию, которая принимает предыдущую цену и текущую цену акции и возвращает сумму, на которую цена изменилась, и процент изменения.

Вот как вы можете реализовать такую ​​функцию в Go -

```golang
func getStockPriceChange(prevPrice, currentPrice float64) (float64, float64) {
	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	return change, percentChange
}
```

Просто! не так ли? Вам просто нужно указать возвращаемые типы, разделенные запятой внутри скобок, а затем вернуть несколько значений через запятую из функции.

Давайте посмотрим полный пример с main()функцией -

```golang
package main
import (
	"fmt"
	"math"
)

func getStockPriceChange(prevPrice, currentPrice float64) (float64, float64) {
	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	return change, percentChange
}


func main() {
	prevStockPrice := 75000.0
	currentStockPrice := 100000.0

	change, percentChange := getStockPriceChange(prevStockPrice, currentStockPrice)

	if change < 0 {
		fmt.Printf("The Stock Price decreased by $%.2f which is %.2f%% of the prev price\n", math.Abs(change), math.Abs(percentChange))
	} else {
		fmt.Printf("The Stock Price increased by $%.2f which is %.2f%% of the prev price\n", change, percentChange)
	}
}
```

Output
The Stock Price increased by $25000.00 which is 33.33% of the prev price

### Возврат значения ошибки из функции
Многократные возвращаемые значения часто используются в Golang для возврата ошибки из функции вместе с результатом.

Давайте рассмотрим пример - getStockPriceChangeфункция, которую мы видели в предыдущем разделе, вернет ±Inf(бесконечность), если prevPriceесть 0. Если вы хотите вернуть ошибку вместо этого, вы можете сделать это, добавив еще одно возвращаемое значение типа errorи вернуть значение ошибки, например, так:

```golang
func getStockPriceChangeWithError(prevPrice, currentPrice float64) (float64, float64, error) {
	if prevPrice == 0 {
		err := errors.New("Previous price cannot be zero")
		return 0, 0, err
	}
	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	return change, percentChange, nil
}
```

errorТип встроенного типа в Golang. Программы Go используют errorзначения для обозначения ненормальной ситуации. Не волнуйтесь, если вы пока не понимаете errors. Вы узнаете больше об обработке ошибок в следующей статье.

Ниже приведен полный пример, демонстрирующий вышеуказанную концепцию с main()функцией -

```golang
package main
import (
	"errors"
	"fmt"
	"math"
)

func getStockPriceChangeWithError(prevPrice, currentPrice float64) (float64, float64, error) {
	if prevPrice == 0 {
		err := errors.New("Previous price cannot be zero")
		return 0, 0, err
	}
	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	return change, percentChange, nil
}

func main() {
	prevStockPrice := 0.0
	currentStockPrice := 100000.0

	change, percentChange, err := getStockPriceChangeWithError(prevStockPrice, currentStockPrice)

	if err != nil {
		fmt.Println("Sorry! There was an error: ", err)
	} else {
		if change < 0 {
			fmt.Printf("The Stock Price decreased by $%.2f which is %.2f%% of the prev price\n", math.Abs(change), math.Abs(percentChange))
		} else {
			fmt.Printf("The Stock Price increased by $%.2f which is %.2f%% of the prev price\n", change, percentChange)
		}
	}
}
```

Output
Sorry! There was an error:  Previous price cannot be zero

 
### Функции с именованными возвращаемыми значениями
Возвращаемые значения функции в Golang могут быть названы. Именованные возвращаемые значения ведут себя так, как будто вы определили их в верхней части функции.

Давайте перепишем getStockPriceChangeфункцию, которую мы видели в предыдущем разделе, с именованными возвращаемыми значениями -

```golang
// Function with named return values
func getNamedStockPriceChange(prevPrice, currentPrice float64) (change, percentChange float64) {
	change = currentPrice - prevPrice
	percentChange = (change / prevPrice) * 100
	return change, percentChange
}
```

Обратите внимание, как мы изменили :=(короткие объявления) с =(присваиваниями) в теле функции. Это связано с тем, что Go сам определяет все именованные возвращаемые значения и делает их доступными для использования в функции. Поскольку они уже определены, вы не можете определить их снова, используя короткие объявления.

Именованные возвращаемые значения позволяют использовать так называемый возврат Naked ( returnоператор без каких-либо аргументов). Когда вы задаете returnоператор без аргументов, он возвращает именованные возвращаемые значения по умолчанию. Таким образом, вы можете написать вышеупомянутую функцию, как это -

```golang
// Function with named return values and naked return
func getNamedStockPriceChange(prevPrice, currentPrice float64) (change, percentChange float64) {
	change = currentPrice - prevPrice
	percentChange = (change / prevPrice) * 100
	return
}
```

Давайте использовать вышеуказанную функцию в полном примере с main()функцией и проверим вывод:

```golang
package main
import (
	"fmt"
	"math"
)

func getNamedStockPriceChange(prevPrice, currentPrice float64) (change, percentChange float64) {
	change = currentPrice - prevPrice
	percentChange = (change / prevPrice) * 100
	return
}

func main() {
	prevStockPrice := 100000.0
	currentStockPrice := 90000.0

	change, percentChange := getNamedStockPriceChange(prevStockPrice, currentStockPrice)

	if change < 0 {
		fmt.Printf("The Stock Price decreased by $%.2f which is %.2f%% of the prev price\n", math.Abs(change), math.Abs(percentChange))
	} else {
		fmt.Printf("The Stock Price increased by $%.2f which is %.2f%% of the prev price\n", change, percentChange)
	}
}
```

# Output
### The Stock Price decreased by $10000.00 which is 10.00% of the prev price
Именованные возвращаемые значения улучшают читаемость ваших функций. Использование значимых имен позволит потребителям вашей функции знать, что функция возвратит, просто взглянув на ее сигнатуру.

Голые операторы return хороши для коротких функций. Но не используйте их, если ваши функции длинные. Они могут повредить читабельности. Вы должны явно указать возвращаемые значения в более длинных функциях.

### Пустой идентификатор

Иногда вы можете игнорировать некоторые результаты функции, которая возвращает несколько значений.

Например, допустим, что вы используете getStockPriceChangeфункцию, которую мы определили в предыдущем разделе, но вас интересует только количество изменений, а не процентное изменение.

Теперь вы можете просто объявить локальные переменные и сохранить все значения, возвращаемые функцией, следующим образом:

change, percentChange := getStockPriceChange(prevStockPrice, currentStockPrice)
Но в этом случае вы будете вынуждены использовать percentChangeпеременную, потому что Go не позволяет создавать переменные, которые вы никогда не используете.

Так в чем же решение? Ну, вместо этого вы можете использовать пустой идентификатор -


change, _ := getStockPriceChange(prevStockPrice, currentStockPrice)
Пустой идентификатор используется для указания Go, что вам не нужно это значение. Следующий пример демонстрирует эту концепцию -

```golang
package main

import (
	"fmt"
	"math"
)

func getStockPriceChange(prevPrice, currentPrice float64) (float64, float64) {
	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	return change, percentChange
}

func main() {
	prevStockPrice := 80000.0
	currentStockPrice := 120000.0

	change, _ := getStockPriceChange(prevStockPrice, currentStockPrice)

	if change < 0 {
		fmt.Printf("The Stock Price decreased by $%.2f\n", math.Abs(change))
	} else {
		fmt.Printf("The Stock Price increased by $%.2f\n", change)
	}
}
```

Output
The Stock Price increased by $40000.00

### Заключение
В этой статье вы узнали, как объявлять и вызывать функции в Golang, как определять функции с несколькими возвращаемыми и именованными возвращаемыми значениями, как возвращать ошибку из функции и как использовать пустой идентификатор.
