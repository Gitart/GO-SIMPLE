# Как ждать, пока Goroutines завершит выполнение?

Тип пакета синхронизации WaitGroup используется для ожидания завершения программой всех процедур, запускаемых из основной функции. Он использует счетчик, который определяет количество процедур, и Wait блокирует выполнение программы до тех пор, пока счетчик WaitGroup не станет равным нулю.

Метод Add используется для добавления счетчика в WaitGroup.

Метод Done для группы WaitGroup планируется с помощью оператора defer для уменьшения счетчика WaitGroup.

Метод Wait типа WaitGroup ожидает, пока программа завершит все процедуры.

Метод Wait вызывается внутри главной функции, которая блокирует выполнение до тех пор, пока счетчик WaitGroup не достигнет нулевого значения, и обеспечит выполнение всех процедур.

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

Вы можете увидеть следующий результат при запуске вышеуказанной программы \-

```markup

C:\Golang\goroutines\create-simple-goroutine>go run main.go
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
