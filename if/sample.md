## Утверждения потока управления Golang: если, переключить и для

### Если заявление
If операторы используются, чтобы указать, должен ли блок кода выполняться или нет в зависимости от заданного условия.

Ниже приводится синтаксис ifвысказываний на Голанге:

```golang
if(condition) {
	// Code to be executed if the condition is true.
}
```

Вот простой пример -

```golang
package main
import "fmt"

func main() {
	var x = 25
	if(x % 5 == 0) {
		fmt.Printf("%d is a multiple of 5\n", x)
	}
}
```

Output  
25 is a multiple of 5

Обратите внимание, что Вы можете опустить скобки ()в ifвыражении в Golang, но фигурные скобки {}обязательны -

```golang
var y = -1
if y < 0 {
	fmt.Printf("%d is negative\n", y)
}
```

Вы можете комбинировать несколько условий, используя операторы короткого замыкания &&и ||так -

```golang
var age = 21
if age >= 17 && age <= 30 {
	fmt.Println("My Age is between 17 and 30")
}	
```
 
### If-Else Заявление
Выписка ifможет быть объединена с elseблоком. elseБлок выполняется , если выполняется условие , указанное в ifзаявлении ложно -

```golang
if condition {
	// code to be executed if the condition is true
} else {
	// code to be executed if the condition is false
}
```

Вот простой пример

```golang
package main
import "fmt"

func main() {
	var age = 18
	if age >= 18 {
		fmt.Println("You're eligible to vote!")
	} else {
		fmt.Println("You're not eligible to vote!")
	}
}
```

Output   
You're eligible to vote!

### If-Else-If Chain
ifУ операторов также может быть несколько else ifчастей, составляющих цепочку таких условий:

```golang
package main
import "fmt"

func main() {
	var BMI = 21.0
	if BMI < 18.5 {
		fmt.Println("You are underweight");
	} else if BMI >= 18.5 && BMI < 25.0 {
		fmt.Println("Your weight is normal");
	} else if BMI >= 25.0 && BMI < 30.0 {
		fmt.Println("You're overweight")
	} else {
		fmt.Println("You're obese")
	}
}
```

Output
Your weight is normal

### Если с коротким заявлением
if Заявление в Golang может также содержать краткое заявление декларации , предшествующее условное выражение -

```golang
if n := 10; n%2 == 0 {
	fmt.Printf("%d is even\n", n)
} 
```

Переменная, объявленная в кратком выражении, доступна только внутри ifблока и его elseили else-ifветвей -

```golang
if n := 15; n%2 == 0 {
	fmt.Printf("%d is even\n", n)
} else {
	fmt.Printf("%d is odd\n", n)
}
```

Обратите внимание, что если вы используете короткое утверждение, то вы не можете использовать скобки. Таким образом, следующий код сгенерирует синтаксическую ошибку -

```golang
// You can't use parentheses when `if` contains a short statement
if (n := 15; n%2 == 0) { // Syntax Error
}
```

 
### Переключатель Заявление
Оператор Switch принимает выражение и сопоставляет его со списком возможных случаев. Как только совпадение найдено, он выполняет блок кода, указанный в сопоставленном регистре.

Вот простой пример оператора switch -

```golang
package main
import "fmt"

func main() {
	var dayOfWeek = 6
	switch dayOfWeek {
		case 1: fmt.Println("Monday")
		case 2: fmt.Println("Tuesday")
		case 3: fmt.Println("Wednesday")
		case 4: fmt.Println("Thursday")
		case 5: fmt.Println("Friday")
		case 6: {
			fmt.Println("Saturday")
			fmt.Println("Weekend. Yaay!")
		}
		case 7: {
			fmt.Println("Sunday")
			fmt.Println("Weekend. Yaay!")
		}
		default: fmt.Println("Invalid day")
	}
}
```

Output  
Saturday
Weekend. Yaay!

Go оценивает все варианты переключения один за другим сверху вниз, пока дело не будет успешно выполнено. Как только дело успешно выполняется, он запускает блок кода, указанный в этом деле, а затем останавливается (он не оценивает дальнейшие дела).

Это противоречит другим языкам, таким как C, C ++ и Java, где вам необходимо явно вставлять breakоператор после тела каждого случая, чтобы остановить оценку последующих случаев.

Если ни один из случаев не завершился успешно, выполняется случай по умолчанию.

### Переключить с коротким заявлением  
Точно так же if, switchможет также содержать краткое объявление объявления, предшествующее условному выражению. Таким образом, вы могли бы также написать предыдущий пример переключения, как это -

```golang
switch dayOfWeek := 6; dayOfWeek {
	case 1: fmt.Println("Monday")
	case 2: fmt.Println("Tuesday")
	case 3: fmt.Println("Wednesday")
	case 4: fmt.Println("Thursday")
	case 5: fmt.Println("Friday")
	case 6: {
		fmt.Println("Saturday")
		fmt.Println("Weekend. Yaay!")
	}
	case 7: {
		fmt.Println("Sunday")
		fmt.Println("Weekend. Yaay!")
	}
	default: fmt.Println("Invalid day")
}
```

Единственное отличие состоит в том, что переменная, объявленная оператором short ( dayOfWeek), доступна только внутри блока switch.

### Объединение нескольких случаев коммутатора
Вы можете объединить несколько switchдел в один, например, так -

```golang
package main
import "fmt"

func main() {
	switch dayOfWeek := 5; dayOfWeek {
		case 1, 2, 3, 4, 5:
			fmt.Println("Weekday")
		case 6, 7:
			fmt.Println("Weekend")
		default:
			fmt.Println("Invalid Day")		
	}
}
```

 Output
Weekday
Это удобно, когда вам нужно запустить общую логику для нескольких случаев.

### Переключить без выражения
В Golang выражение, которое мы указываем в switchвыражении, является необязательным. switchЗаявление без выражения такого же , как switch true. Он оценивает все случаи один за другим и запускает первый случай, который оценивается как true -

```golang
package main
import "fmt"

func main() {
	var BMI = 21.0 
	switch {
		case BMI < 18.5:
			fmt.Println("You're underweight")
		case BMI >= 18.5 && BMI < 25.0:
			fmt.Println("Your weight is normal")
		case BMI >= 25.0 && BMI < 30.0:
			fmt.Println("You're overweight")
		default:
			fmt.Println("You're obese")	
	}
}
```

Переключение без выражения - это просто лаконичный способ написания if-else-ifцепочек.

### Для петли
Цикл используется для многократного выполнения блока кода. У Голанга есть только одно утверждение цикла - forцикл.

Ниже приводится общий синтаксис forцикла в Go -
```golang
for initialization; condition; increment {
	// loop body
}
```

Оператор инициализации выполняется ровно один раз до первой итерации цикла. На каждой итерации условие проверяется. Если условие trueвыполняется, тогда выполняется тело цикла, в противном случае цикл завершается. Оператор приращения выполняется в конце каждой итерации.

```golang
Вот простой пример цикла for -

package main
import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}
```

Output
0 1 2 3 4 5 6 7 8 9 
В отличие от других языков, таких как C, C ++ и Java, цикл for в Go не содержит скобок, а фигурные скобки обязательны.

Обратите внимание, что операторы инициализации и приращения в forцикле являются необязательными и могут быть опущены

### Пропуск инструкции инициализации

```golang
package main
import "fmt"

func main() {
    i := 2
    for ;i <= 10; i += 2 {
        fmt.Printf("%d ", i)
    }
}
```

Output
2 4 6 8 10 

### Пропуск инструкции приращения

```golang
package main
import "fmt"

func main() {
    i := 2
    for ;i <= 20; {
        fmt.Printf("%d ", i)
        i *= 2
    }
}
```

 Output
2 4 8 16

Обратите внимание, что вы также можете опустить точки с запятой в forцикле в вышеприведенном примере и написать это так:
```golang
package main
import "fmt"

func main() {
    i := 2
    for i <= 20 {
        fmt.Printf("%d ", i)
        i *= 2
    }
}   
```

Вышеуказанный forцикл похож на whileцикл в других языках. Go не имеет whileцикла, потому что мы можем легко представить whileцикл с помощью for.

Наконец, вы также можете опустить условие из forцикла в Golang. 
### Это даст вам бесконечный цикл -

```golang
package main

func main() {
	// Infinite Loop
	for {
	}
}
```

### заявление о нарушении
Вы можете использовать breakоператор, чтобы выйти из цикла до его нормального завершения. Вот пример -
```golang
package main
import "fmt"

func main() {
	for num := 1; num <= 100; num++ {
		if num%3 == 0 && num%5 == 0 {
			fmt.Printf("First positive number divisible by both 3 and 5 is %d\n", num)
			break
		}
	}
}
```

 Output
First positive number divisible by both 3 and 5 is 15

### продолжить заявление
Оператор continueиспользуется для остановки выполнения тела цикла на полпути и перехода к следующей итерации цикла.

```golang
package main
import "fmt"

func main() {
	for num := 1; num <= 10; num++ {
		if num%2 == 0 {
			continue;
		}
		fmt.Printf("%d ", num)
	}
}
```

Output
1 3 5 7 9 

### Заключение
В этой статье вы узнали, как работать с операторами потока управления, например if, switchи forв Golang.
