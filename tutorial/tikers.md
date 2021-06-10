В этом руководстве мы рассмотрим тикеры в Go и то, как вы можете эффективно использовать тикеры в ваших собственных приложениях Go.

Тикеры исключительно полезны, когда вам нужно многократно выполнять действие в заданные промежутки времени, и мы можем использовать тикеры в сочетании с горутинами для выполнения этих задач в фоновом режиме наших приложений.

## Тикеры против таймеров

Прежде чем мы углубимся, полезно знать разницу между обоими `tickers` и `timers.`

*   `Tickers` \- Они отлично подходят для повторяющихся задач
*   `Timers` \- Они используются для разовых задач.

## Простой пример

Давайте начнем с действительно простого, в котором мы многократно выполняем простую `fmt.Println` инструкцию каждые 5 секунд.

main.go

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Go Tickers Tutorial")
	// this creates a new ticker which will
    // `tick` every 1 second.
    ticker := time.NewTicker(1 * time.Second)

    // for every `tick` that our `ticker`
    // emits, we print `tock`
	for _ = range ticker.C {
		fmt.Println("tock")
	}
}

```

Теперь, когда мы запускаем это, наше приложение Go будет работать бесконечно, пока мы не выйдем `ctrl-c` из программы, и каждую секунду оно будет выводиться `tock` на терминал.

$ go, запустите main.go

```output
Go Tickers Tutorial
Tock
Tock
^Csignal: interrupt

```

## Работа в фоновом режиме

Итак, мы смогли реализовать действительно простое приложение Go, которое использует a `ticker` для многократного выполнения действия. Однако что произойдет, если мы хотим, чтобы это действие выполнялось в фоновом режиме нашего приложения Go?

Что ж, если бы у нас была задача, которую мы хотели бы запустить в фоновом режиме, мы могли бы переместить наш `for` цикл, который повторяется, `ticker.C` внутрь, `goroutine` что позволит нашему приложению выполнять другие задачи.

Давайте переместим код для создания тикера и цикла в новую вызываемую функцию, `backgroundTask()` а затем в нашей `main()` функции мы назовем это как горутину, используя `go` ключевое слово следующим образом:

main.go

```go
package main

import (
	"fmt"
	"time"
)

func backgroundTask() {
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		fmt.Println("Tock")
	}
}

func main() {
	fmt.Println("Go Tickers Tutorial")

	go backgroundTask()

    // This print statement will be executed before
    // the first `tock` prints in the console
	fmt.Println("The rest of my application can continue")
	// here we use an empty select{} in order to keep
    // our main function alive indefinitely as it would
    // complete before our backgroundTask has a chance
    // to execute if we didn't.
	select{}

}

```

Круто, поэтому, если мы продолжим и запустим это, мы увидим, что наша `main()` функция продолжает выполняться после запуска нашей горутины фоновой задачи.

$ go, запустите main.go

```output
Go Tickers Tutorial
The rest of my application can continue
Tock
Tock
Tock
^Csignal: interrupt

```

## Заключение

Итак, в этом руководстве мы рассмотрели, как вы можете использовать тикеры в ваших собственных приложениях Go для выполнения повторяющихся задач, как в основном потоке, так и в качестве фоновой задачи.

### Дальнейшее чтение

Если вам понравилось, и вы хотите увидеть, как вы можете использовать его `tickers` в более продвинутом контексте, я рекомендую ознакомиться с другой моей статьей, которая представляет собой систему мониторинга статистики YouTube в реальном времени.
