# Горутины Голанга

[❮ Предыдущий](https://www.golangprograms.com/go-language/interface.html) [Следующий ❯](https://www.golangprograms.com/go-language/channels.html)

---

## Горутины

Параллелизм в Golang - это способность функций работать независимо друг от друга. Горутины - это функции, которые выполняются одновременно. Golang предоставляет горутины как способ одновременной обработки операций.

Новые горутины создаются оператором go .

Чтобы запустить функцию как горутину, вызовите эту функцию с префиксом go. Вот пример блока кода:

### Пример

```jsx
sum()     // A normal function call that executes sum synchronously and waits for completing it
go sum()  // A goroutine that executes sum asynchronously and doesn't wait for completing it

```

Идти ключевое слово делает вызов функции немедленно вернуться, в то время как функция начинает работать в фоновом режиме в качестве goroutine и остальной частью программы продолжает выполнение. **Основная** функция каждой программы Golang запускается с помощью goroutine, поэтому каждая программа Golang работает по крайней мере один goroutine.

---

## Создание горутин

Добавлено ключевое слово go перед каждым вызовом функции **responseSize** . Три горутины **responseSize** запускаются одновременно, и одновременно *выполняются* три вызова *http.Get* . Программа не ждет, пока вернется один ответ, перед отправкой следующего запроса. В результате три размера ответа печатаются гораздо раньше с помощью горутин.

### Пример

```jsx
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
	fmt.Println("Step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step2: ", url)
	defer response.Body.Close()

	fmt.Println("Step3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step4: ", len(body))
}

func main() {
	go responseSize("https://www.golangprograms.com")
	go responseSize("https://coderwall.com")
	go responseSize("https://stackoverflow.com")
	time.Sleep(10 * time.Second)
}
```

### Выход

```jsx
Step1:  https://www.golangprograms.com
Step1:  https://stackoverflow.com
Step1:  https://coderwall.com
Step2:  https://stackoverflow.com
Step3:  https://stackoverflow.com
Step4:  116749
Step2:  https://www.golangprograms.com
Step3:  https://www.golangprograms.com
Step4:  79551
Step2:  https://coderwall.com
Step3:  https://coderwall.com
Step4:  203842
```

Мы добавили вызов функции *time.Sleep* в **основную** функцию, которая предотвращает выход из основной *горутины* до того, как *горутины* **responseSize** могут завершиться. Вызов time.Sleep (10 \* time.Second) заставит **основную** горутину спать на 10 секунд.

---

## Ожидание завершения выполнения горутинов

Тип пакета синхронизации WaitGroup используется для ожидания завершения программой всех горутин, запущенных из основной функции. Он использует счетчик, который указывает количество горутин, а Wait блокирует выполнение программы до тех пор, пока счетчик WaitGroup не станет нулевым.

Метод Add используется для добавления счетчика в WaitGroup.

Метод Done класса WaitGroup планируется с использованием оператора defer для уменьшения счетчика WaitGroup.

Метод Wait типа WaitGroup ожидает, пока программа завершит все горутины.

Метод Wait вызывается внутри основной функции, которая блокирует выполнение до тех пор, пока счетчик WaitGroup не достигнет нулевого значения, и гарантирует выполнение всех горутин.

### Пример

```jsx
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

// WaitGroup is used to wait for the program to finish goroutines.
var wg sync.WaitGroup

func responseSize(url string) {
	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done()

	fmt.Println("Step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step2: ", url)
	defer response.Body.Close()

	fmt.Println("Step3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step4: ", len(body))
}

func main() {
	// Add a count of three, one for each goroutine.
	wg.Add(3)
	fmt.Println("Start Goroutines")

	go responseSize("https://www.golangprograms.com")
	go responseSize("https://stackoverflow.com")
	go responseSize("https://coderwall.com")

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Terminating Program")
}
```

### Выход

```jsx
Start Goroutines
Step1:  https://coderwall.com
Step1:  https://www.golangprograms.com
Step1:  https://stackoverflow.com
Step2:  https://stackoverflow.com
Step3:  https://stackoverflow.com
Step4:  116749
Step2:  https://www.golangprograms.com
Step3:  https://www.golangprograms.com
Step4:  79801
Step2:  https://coderwall.com
Step3:  https://coderwall.com
Step4:  203842
Terminating Program
```

---

## Получение значений из горутин

Самый естественный способ получить значение из горутины - это каналы. Каналы - это каналы, которые соединяют параллельные горутины. Вы можете отправлять значения в каналы из одной горутины и получать эти значения в другую горутину или в синхронной функции.

### Пример

```jsx
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

// WaitGroup is used to wait for the program to finish goroutines.
var wg sync.WaitGroup

func responseSize(url string, nums chan int) {
	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Send value to the unbuffered channel
	nums <- len(body)
}

func main() {
	nums := make(chan int) // Declare a unbuffered channel
	wg.Add(1)
	go responseSize("https://www.golangprograms.com", nums)
	fmt.Println(<-nums) // Read the value from unbuffered channel
	wg.Wait()
	close(nums) // Closes the channel
}
```

### Выход

```jsx
79655
```

---

## Воспроизвести и приостановить выполнение горутина

Используя каналы, мы можем воспроизводить и приостанавливать выполнение горутины. **Канала** обрабатывает это сообщение, действуя в качестве канала между goroutines.

### Пример

```jsx
package main

import (
	"fmt"
	"sync"
	"time"
)

var i int

func work() {
	time.Sleep(250 * time.Millisecond)
	i++
	fmt.Println(i)
}

func routine(command <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var status = "Play"
	for {
		select {
		case cmd := <-command:
			fmt.Println(cmd)
			switch cmd {
			case "Stop":
				return
			case "Pause":
				status = "Pause"
			default:
				status = "Play"
			}
		default:
			if status == "Play" {
				work()
			}
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	command := make(chan string)
	go routine(command, &wg)

	time.Sleep(1 * time.Second)
	command <- "Pause"

	time.Sleep(1 * time.Second)
	command <- "Play"

	time.Sleep(1 * time.Second)
	command <- "Stop"

	wg.Wait()
}
```

### Выход

```jsx
1
2
3
4
Pause
Play
5
6
7
8
9
Stop
```

---

## Исправить состояние гонки с помощью атомарных функций

Состояния состязания возникают из-за несинхронизированного доступа к общему ресурсу и попытки одновременного чтения и записи в этот ресурс.

Атомарные функции обеспечивают низкоуровневые механизмы блокировки для синхронизации доступа к целым числам и указателям. Атомарные функции обычно используются для исправления состояния гонки.

Функции в пакетах **atomic** under **sync** обеспечивают поддержку синхронизации горутин путем блокировки доступа к общим ресурсам.

### Пример

```jsx
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int32          // counter is a variable incremented by all goroutines.
	wg      sync.WaitGroup // wg is used to wait for the program to finish.
)

func main() {
	wg.Add(3) // Add a count of two, one for each goroutine.

	go increment("Python")
	go increment("Java")
	go increment("Golang")

	wg.Wait() // Wait for the goroutines to finish.
	fmt.Println("Counter:", counter)

}

func increment(name string) {
	defer wg.Done() // Schedule the call to Done to tell main we are done.

	for range name {
		atomic.AddInt32(&counter, 1)
		runtime.Gosched() // Yield the thread and be placed back in queue.
	}
}
```

### Выход

Функция AddInt32 из атомарного пакета синхронизирует добавление целочисленных значений, обеспечивая, чтобы только одна горутина могла выполнять и завершать эту операцию добавления за раз. Когда горутины пытаются вызвать любую атомарную функцию, они автоматически синхронизируются с указанной переменной.

Обратите внимание, если вы замените строку кода **atomic.AddInt32 (& counter, 1)** на **counter ++** , вы увидите следующий результат:

```markup
C:\Golang\goroutines>go run -race main.go
==================
WARNING: DATA RACE
Read at 0x0000006072b0 by goroutine 7:
  main.increment()
      C:/Golang/goroutines/main.go:31 +0x76

Previous write at 0x0000006072b0 by goroutine 8:
  main.increment()
      C:/Golang/goroutines/main.go:31 +0x90

Goroutine 7 (running) created at:
  main.main()
      C:/Golang/goroutines/main.go:18 +0x7e

Goroutine 8 (running) created at:
  main.main()
      C:/Golang/goroutines/main.go:19 +0x96
==================
Counter: 15
Found 1 data race(s)
exit status 66

C:\Golang\goroutines>
```

### Выход

```jsx
C:\Golang\goroutines>go run -race main.go
Counter: 15
```

---

## Определите критические разделы с помощью Mutex

Мьютекс используется для создания критического раздела вокруг кода, который гарантирует, что только одна горутина может одновременно выполнить этот раздел кода.

### Пример

```jsx
package main

import (
	"fmt"
	"sync"
)

var (
	counter int32          // counter is a variable incremented by all goroutines.
	wg      sync.WaitGroup // wg is used to wait for the program to finish.
	mutex   sync.Mutex     // mutex is used to define a critical section of code.
)

func main() {
	wg.Add(3) // Add a count of two, one for each goroutine.

	go increment("Python")
	go increment("Go Programming Language")
	go increment("Java")

	wg.Wait() // Wait for the goroutines to finish.
	fmt.Println("Counter:", counter)

}

func increment(lang string) {
	defer wg.Done() // Schedule the call to Done to tell main we are done.

	for i := 0; i < 3; i++ {
		mutex.Lock()
		{
			fmt.Println(lang)
			counter++
		}
		mutex.Unlock()
	}
}
```

### Выход

```jsx
C:\Golang\goroutines>go run -race main.go
PHP stands for Hypertext Preprocessor.
PHP stands for Hypertext Preprocessor.
The Go Programming Language, also commonly referred to as Golang
The Go Programming Language, also commonly referred to as Golang
Counter: 4

C:\Golang\goroutines>
```

Критическая секция, определяемая вызовами **Lock ()** и **Unlock (),** защищает действия от переменной счетчика и чтения текста переменной имени.
