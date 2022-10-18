# Варианты использования канала

Прежде чем читать эту статью, прочтите статью о [каналах в Go](https://go101.org/article/channel.html) , в которой подробно описаны типы и значения каналов. Начинающим сусликам может понадобиться несколько раз прочитать эту и текущую статьи, чтобы освоить программирование канала Go.

В оставшейся части этой статьи будут показаны многие варианты использования канала. Надеюсь, эта статья убедит вас в том, что

*   Асинхронное и параллельное программирование с каналами Go — это просто и приятно.
*   метод синхронизации каналов имеет более широкий спектр применения и имеет больше вариаций, чем решения синхронизации, используемые в некоторых других языках, таких как [модель](https://en.wikipedia.org/wiki/Actor_model) акторов и [шаблон async/await](https://en.wikipedia.org/wiki/Async/await) .

Обратите внимание, что цель этой статьи — показать как можно больше вариантов использования каналов. Мы должны знать, что канал — не единственный метод параллельной синхронизации, поддерживаемый в Go, и в некоторых случаях канал может быть не лучшим решением. Пожалуйста, прочитайте [атомарные операции](https://go101.org/article/concurrent-atomic-operation.html) и [некоторые другие методы синхронизации,](https://go101.org/article/concurrent-synchronization-more.html) чтобы узнать больше о методах параллельной синхронизации в Go.

### Используйте каналы как фьючерсы/обещания

Фьючерсы и обещания используются во многих других популярных языках. Они часто связаны с запросами и ответами.

#### Возвращать каналы только для приема в качестве результатов

В следующем примере `sumSquares` одновременно запрашиваются значения двух аргументов вызова функции. Каждая из двух операций приема канала будет заблокирована до тех пор, пока операция отправки не будет выполнена на соответствующем канале. Для возврата окончательного результата требуется около трех секунд вместо шести секунд.

```go
package main

import (
	"time"
	"math/rand"
	"fmt"
)

func longTimeRequest() <-chan int32 {
	r := make(chan int32)

	go func() {
		// Simulate a workload.
		time.Sleep(time.Second * 3)
		r <- rand.Int31n(100)
	}()

	return r
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano())

	a, b := longTimeRequest(), longTimeRequest()
	fmt.Println(sumSquares(<-a, <-b))
}

```

#### Передавать каналы только для отправки в качестве аргументов

Как и в предыдущем примере, в следующем примере значения двух аргументов `sumSquares` вызова функции запрашиваются одновременно. В отличие от последнего примера, `longTimeRequest` функция принимает в качестве параметра канал только для отправки, а не возвращает результат канала только для приема.

```go
package main

import (
	"time"
	"math/rand"
	"fmt"
)

func longTimeRequest(r chan<- int32)  {
	// Simulate a workload.
	time.Sleep(time.Second * 3)
	r <- rand.Int31n(100)
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ra, rb := make(chan int32), make(chan int32)
	go longTimeRequest(ra)
	go longTimeRequest(rb)

	fmt.Println(sumSquares(<-ra, <-rb))
}

```

На самом деле, для указанного выше примера нам не нужны два канала для передачи результатов. Использование одного канала — это нормально.

```go
...

	// The channel can be buffered or not.
	results := make(chan int32, 2)
	go longTimeRequest(results)
	go longTimeRequest(results)

	fmt.Println(sumSquares(<-results, <-results))
}

```

Это своего рода агрегация данных, которая будет специально представлена ​​ниже.

#### Выигрывает первый ответ

Это усовершенствование варианта с использованием только одного канала в последнем примере.

Иногда часть данных может быть получена из нескольких источников, чтобы избежать больших задержек. Для многих факторов продолжительность отклика этих источников может сильно различаться. Даже для указанного источника длительность его отклика также непостоянна. Чтобы сделать продолжительность ответа как можно короче, мы можем отправить запрос каждому источнику в отдельной горутине. Будет использован только первый ответ, остальные более медленные будут отброшены.

Обратите внимание: если есть *N* источников, пропускная способность канала связи должна быть не менее *N-1* , чтобы избежать блокировки горутин, соответствующих отброшенным ответам, навсегда.

```go
package main

import (
	"fmt"
	"time"
	"math/rand"
)

func source(c chan<- int32) {
	ra, rb := rand.Int31(), rand.Intn(3) + 1
	// Sleep 1s/2s/3s.
	time.Sleep(time.Duration(rb) * time.Second)
	c <- ra
}

func main() {
	rand.Seed(time.Now().UnixNano())

	startTime := time.Now()
	// c must be a buffered channel.
	c := make(chan int32, 5)
	for i := 0; i < cap(c); i++ {
		go source(c)
	}
	// Only the first response will be used.
	rnd := <- c
	fmt.Println(time.Since(startTime))
	fmt.Println(rnd)
}

```

Есть и другие способы реализовать вариант использования «выигрывает первый ответ», используя механизм выбора и буферизованный канал с пропускной способностью, равной единице. Другие способы будут представлены ниже.

#### Больше вариантов запроса-ответа

Каналы параметров и результатов могут быть буферизованы, так что сторонам ответа не нужно будет ждать, пока стороны запроса извлекут переданные значения.

Иногда не гарантируется, что на запрос будет возвращено допустимое значение. Вместо этого по разным причинам может быть возвращена ошибка. В таких случаях мы можем использовать тип структуры `struct{v T; err error}` или пустой тип интерфейса в качестве типа элемента канала.

Иногда по некоторым причинам для ответа может потребоваться гораздо больше времени, чем ожидалось, или он никогда не придет. Мы можем использовать механизм тайм-аута, представленный ниже, для обработки таких обстоятельств.

Иногда со стороны ответа может быть возвращена последовательность значений, это своего рода механизм потока данных, упомянутый ниже.

### Используйте каналы для уведомлений

Уведомления можно рассматривать как специальные запросы/ответы, в которых отвеченные значения не важны. Как правило, мы используем пустой тип структуры `struct{}` в качестве типов элементов каналов уведомлений, поскольку размер типа `struct{}` равен нулю, поэтому значения `struct{}` не потребляют память.

#### Уведомление 1-к-1 путем отправки значения в канал

Если нет значений, которые нужно получить из канала, то следующая операция приема на канале будет заблокирована до тех пор, пока другая горутина не отправит значение в канал. Таким образом, мы можем отправить значение в канал, чтобы уведомить другую горутину, ожидающую получения значения из того же канала.

В следующем примере канал `done` используется в качестве сигнального канала для уведомлений.

```go
package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"sort"
)

func main() {
	values := make([]byte, 32 * 1024 * 1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	done := make(chan struct{}) // can be buffered or not

	// The sorting goroutine
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		// Notify sorting is done.
		done <- struct{}{}
	}()

	// do some other things ...

	<- done // waiting here for notification
	fmt.Println(values[0], values[len(values)-1])
}

```

#### Уведомление 1-к-1 путем получения значения из канала

Если очередь буфера значений канала заполнена (очередь буфера небуферизованного канала всегда заполнена), операция отправки на канале будет заблокирована до тех пор, пока другая горутина не получит значение из канала. Таким образом, мы можем получить значение из канала, чтобы уведомить другую горутину, ожидающую отправки значения в тот же канал. Как правило, канал должен быть небуферизованным.

Этот способ уведомления используется гораздо реже, чем способ, представленный в последнем примере.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
		// The capacity of the signal channel can
		// also be one. If this is true, then a
		// value must be sent to the channel before
		// creating the following goroutine.

	go func() {
		fmt.Print("Hello")
		// Simulate a workload.
		time.Sleep(time.Second * 2)

		// Receive a value from the done
		// channel, to unblock the second
		// send in main goroutine.
		<- done
	}()

	// Blocked here, wait for a notification.
	done <- struct{}{}
	fmt.Println(" world!")
}

```

На самом деле принципиальных различий между получением и отправкой значений для создания уведомлений нет. Оба они могут быть суммированы, поскольку более быстрые уведомляются более медленными.

#### Уведомления N-to-1 и 1-to-N

Немного расширив приведенные выше два варианта использования, можно легко выполнять уведомления N-to-1 и 1-to-N.

```go
package main

import "log"
import "time"

type T = struct{}

func worker(id int, ready <-chan T, done chan<- T) {
	<-ready // block here and wait a notification
	log.Print("Worker#", id, " starts.")
	// Simulate a workload.
	time.Sleep(time.Second * time.Duration(id+1))
	log.Print("Worker#", id, " job done.")
	// Notify the main goroutine (N-to-1),
	done <- T{}
}

func main() {
	log.SetFlags(0)

	ready, done := make(chan T), make(chan T)
	go worker(0, ready, done)
	go worker(1, ready, done)
	go worker(2, ready, done)

	// Simulate an initialization phase.
	time.Sleep(time.Second * 3 / 2)
	// 1-to-N notifications.
	ready <- T{}; ready <- T{}; ready <- T{}
	// Being N-to-1 notified.
	<-done; <-done; <-done
}

```

На самом деле способы выполнения уведомлений 1-к-N и N-к-1, представленные в этом подразделе, на практике обычно не используются. На практике мы часто используем `sync.WaitGroup` уведомления N-to-1, а уведомления 1-to-N делаем по закрытым каналам. Подробности читайте в следующем подразделе.

#### Широковещательные (от 1 до N) уведомления путем закрытия канала

Способ выполнения уведомлений 1-к-N, показанный в последнем подразделе, редко используется на практике, поскольку есть способ получше. Используя возможность получения бесконечных значений из закрытого канала, мы можем закрыть канал для широковещательных уведомлений.

В примере из последнего подраздела мы можем заменить три операции отправки канала `ready <- struct{}{}` в последнем примере одной операцией закрытия канала `close(ready)` для выполнения уведомлений 1-к-N.

```go
...
	close(ready) // broadcast notifications
...

```

Конечно, мы также можем закрыть канал, чтобы сделать уведомление один к одному. На самом деле, это наиболее часто используемый способ уведомления в Go.

Возможность получения бесконечных значений из закрытого канала будет использоваться во многих других случаях использования, представленных ниже. На самом деле эта функция широко используется в стандартных пакетах. Например, `context` пакет использует эту функцию для подтверждения отмены.

#### Таймер: уведомление по расписанию

Каналы легко использовать для реализации одноразовых таймеров.

Пользовательская реализация одноразового таймера:

```go
package main

import (
	"fmt"
	"time"
)

func AfterDuration(d time.Duration) <- chan struct{} {
	c := make(chan struct{}, 1)
	go func() {
		time.Sleep(d)
		c <- struct{}{}
	}()
	return c
}

func main() {
	fmt.Println("Hi!")
	<- AfterDuration(time.Second)
	fmt.Println("Hello!")
	<- AfterDuration(time.Second)
	fmt.Println("Bye!")
}

```

На самом деле `After` функция в `time` стандартном пакете обеспечивает ту же функциональность, но гораздо более эффективную реализацию. Вместо этого мы должны использовать эту функцию, чтобы код выглядел чистым.

Обратите внимание, `<-time.After(aDuration)` что текущая горутина перейдет в состояние блокировки, а `time.Sleep(aDuration)` вызов функции — нет.

Использование `<-time.After(aDuration)` часто используется в механизме тайм-аута, который будет представлен ниже.

### Используйте каналы в качестве блокировки мьютекса

В одном из приведенных выше примеров упоминалось, что буферизованные каналы с одной емкостью можно использовать в качестве одноразового [двоичного семафора](https://en.wikipedia.org/wiki/Semaphore_(programming)) . На самом деле, такие каналы также можно использовать как многократные двоичные семафоры, также известные как мьютексы, хотя такие мьютексы не эффективны, как мьютексы, предоставляемые в `sync` стандартном пакете.

Существует два способа использования буферизованных каналов с одной емкостью в качестве мьютексов.

1.  Блокировка через отправку, разблокировка через получение.
2.  Блокировка через получение, разблокировка через отправку.

Ниже приведен пример блокировки сквозной отправки.

```go
package main

import "fmt"

func main() {
	// The capacity must be one.
	mutex := make(chan struct{}, 1)

	counter := 0
	increase := func() {
		mutex <- struct{}{} // lock
		counter++
		<-mutex // unlock
	}

	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			increase()
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase1000(done)
	go increase1000(done)
	<-done; <-done
	fmt.Println(counter) // 2000
}

```

Ниже приведен пример блокировки через получение. Он просто показывает измененную часть на основе приведенного выше примера блокировки через отправку.

```go
...
func main() {
	mutex := make(chan struct{}, 1)
	mutex <- struct{}{} // this line is needed.

	counter := 0
	increase := func() {
		<-mutex // lock
		counter++
		mutex <- struct{}{} // unlock
	}
...

```

### Используйте каналы в качестве счетных семафоров

Буферизованные каналы могут использоваться как [счетные семафоры](https://en.wikipedia.org/wiki/Semaphore_(programming)) . Счетные семафоры можно рассматривать как блокировки с несколькими владельцами. Если пропускная способность канала равна `N` , то его можно рассматривать как замок, который может иметь большинство `N` владельцев в любое время. Бинарные семафоры (мьютексы) — это специальные счетные семафоры, каждый из бинарных семафоров в любой момент времени может иметь не более одного владельца.

Счетные семафоры часто используются для обеспечения максимального количества одновременных запросов.

Как и при использовании каналов в качестве мьютексов, существует два способа получения права собственности на семафор канала.

1.  Получите право собственности через отправку, освободите через получение.
2.  Получите право собственности через получение, освободите через отправку.

Пример приобретения права собственности через получение значений из канала.

```go
package main

import (
	"log"
	"time"
	"math/rand"
)

type Seat int
type Bar chan Seat

func (bar Bar) ServeCustomer(c int) {
	log.Print("customer#", c, " enters the bar")
	seat := <- bar // need a seat to drink
	log.Print("++ customer#", c, " drinks at seat#", seat)
	time.Sleep(time.Second * time.Duration(2 + rand.Intn(6)))
	log.Print("-- customer#", c, " frees seat#", seat)
	bar <- seat // free seat and leave the bar
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// the bar has 10 seats.
	bar24x7 := make(Bar, 10)
	// Place seats in an bar.
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		// None of the sends will block.
		bar24x7 <- Seat(seatId)
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		go bar24x7.ServeCustomer(customerId)
	}

	// sleeping != blocking
	for {time.Sleep(time.Second)}
}

```

В приведенном выше примере пить могут только клиенты, каждый из которых получил место. Таким образом, в любой момент времени будет больше десяти клиентов, которые пьют.

Последний `for` цикл в `main` функции предотвращает выход из программы. Существует лучший способ, который будет представлен ниже, для выполнения этой работы.

В приведенном выше примере, несмотря на то, что в любой момент времени будет больше десяти клиентов, которые пьют, в баре может одновременно обслуживаться более десяти клиентов. Некоторые клиенты ждут свободных мест. Хотя каждая клиентская горутина потребляет гораздо меньше ресурсов, чем системный поток, общими ресурсами, потребляемыми большим количеством горутин, нельзя пренебречь. Поэтому лучше всего создавать клиентскую горутину, только если есть свободное место.

```go
... // same code as the above example

func (bar Bar) ServeCustomerAtSeat(c int, seat Seat) {
	log.Print("++ customer#", c, " drinks at seat#", seat)
	time.Sleep(time.Second * time.Duration(2 + rand.Intn(6)))
	log.Print("-- customer#", c, " frees seat#", seat)
	bar <- seat // free seat and leave the bar
}

func main() {
	rand.Seed(time.Now().UnixNano())

	bar24x7 := make(Bar, 10)
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		bar24x7 <- Seat(seatId)
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		// Need a seat to serve next customer.
		seat := <- bar24x7
		go bar24x7.ServeCustomerAtSeat(customerId, seat)
	}
	for {time.Sleep(time.Second)}
}

```

В приведенной выше оптимизированной версии будет сосуществовать не более десяти активных клиентских горутин (но все еще будет много клиентских горутин, которые будут созданы за время существования программы).

В более эффективной реализации, показанной ниже, за время существования программы будет создано не более десяти горутин, обслуживающих клиентов.

```go
... // same code as the above example

func (bar Bar) ServeCustomerAtSeat(consumers chan int) {
	for c := range consumers {
		seatId := <- bar
		log.Print("++ customer#", c, " drinks at seat#", seatId)
		time.Sleep(time.Second * time.Duration(2 + rand.Intn(6)))
		log.Print("-- customer#", c, " frees seat#", seatId)
		bar <- seatId // free seat and leave the bar
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	bar24x7 := make(Bar, 10)
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		bar24x7 <- Seat(seatId)
	}

	consumers := make(chan int)
	for i := 0; i < cap(bar24x7); i++ {
		go bar24x7.ServeCustomerAtSeat(consumers)
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		consumers <- customerId
	}
}

```

Не по теме: конечно, если нас не волнуют идентификаторы мест (что распространено на практике), то `bar24x7` семафор вообще не важен:

```go
... // same code as the above example

func ServeCustomer(consumers chan int) {
	for c := range consumers {
		log.Print("++ customer#", c, " drinks at the bar")
		time.Sleep(time.Second * time.Duration(2 + rand.Intn(6)))
		log.Print("-- customer#", c, " leaves the bar")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const BarSeatCount = 10
	consumers := make(chan int)
	for i := 0; i < BarSeatCount; i++ {
		go ServeCustomer(consumers)
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		consumers <- customerId
	}
}

```

Способ получения права собственности на семафор посредством отправки сравнительно проще. Этап расстановки сидений не нужен.

```go
package main

import (
	"log"
	"time"
	"math/rand"
)

type Customer struct{id int}
type Bar chan Customer

func (bar Bar) ServeCustomer(c Customer) {
	log.Print("++ customer#", c.id, " starts drinking")
	time.Sleep(time.Second * time.Duration(3 + rand.Intn(16)))
	log.Print("-- customer#", c.id, " leaves the bar")
	<- bar // leaves the bar and save a space
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// The bar can serve most 10 customers
	// at the same time.
	bar24x7 := make(Bar, 10)
	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second * 2)
		customer := Customer{customerId}
		// Wait to enter the bar.
		bar24x7 <- customer
		go bar24x7.ServeCustomer(customer)
	}
	for {time.Sleep(time.Second)}
}

```

### Диалог (пинг-понг)

Две горутины могут вести диалог через канал. Ниже приведен пример, который напечатает серию чисел Фибоначчи.

```go
package main

import "fmt"
import "time"
import "os"

type Ball uint64

func Play(playerName string, table chan Ball) {
	var lastValue Ball = 1
	for {
		ball := <- table // get the ball
		fmt.Println(playerName, ball)
		ball += lastValue
		if ball < lastValue { // overflow
			os.Exit(0)
		}
		lastValue = ball
		table <- ball // bat back the ball
		time.Sleep(time.Second)
	}
}

func main() {
	table := make(chan Ball)
	go func() {
		table <- 1 // throw ball on table
	}()
	go Play("A:", table)
	Play("B:", table)
}

```

### Канал, инкапсулированный в канал

Иногда мы можем использовать тип канала в качестве типа элемента другого типа канала. В следующем примере `chan chan<- int` это тип канала, тип элемента которого является типом канала только для отправки `chan<- int` .

```go
package main

import "fmt"

var counter = func (n int) chan<- chan<- int {
	requests := make(chan chan<- int)
	go func() {
		for request := range requests {
			if request == nil {
				n++ // increase
			} else {
				request <- n // take out
			}
		}
	}()

	// Implicitly converted to chan<- (chan<- int)
	return requests
}(0)

func main() {
	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			counter <- nil
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase1000(done)
	go increase1000(done)
	<-done; <-done

	request := make(chan int, 1)
	counter <- request
	fmt.Println(<-request) // 2000
}

```

Хотя здесь реализация инкапсуляции может быть не самым эффективным способом для указанного выше примера, вариант использования может быть полезен для некоторых других сценариев.

### Проверка длины и пропускной способности каналов

Мы можем использовать встроенные функции `len` и `cap` для проверки длины и пропускной способности канала. Однако мы редко делаем это на практике. Причина, по которой мы редко используем `len` функцию для проверки длины канала, заключается в том, что длина канала могла измениться после `len` возврата вызова функции. Причина того, что мы редко используем эту `cap` функцию для проверки пропускной способности канала, заключается в том, что пропускная способность канала часто известна или не важна.

Однако есть несколько сценариев, в которых нам нужно использовать две функции. Например, иногда мы хотим получить все значения, буферизованные в незакрытом канале `c` , в который больше никто не будет отправлять значения, тогда мы можем использовать следующий код для получения оставшихся значений.

```go
// Assume the current goroutine is the only
// goroutine tries to receive values from
// the channel c at present.
for len(c) > 0 {
	value := <-c
	// use value ...
}

```

Мы также можем использовать механизм try-receive, представленный ниже, для выполнения той же работы. Эффективность обоих способов почти одинакова. Преимущество механизма try-receive заключается в том, что текущая горутина не обязана быть единственной принимающей горутиной.

Иногда горутина может захотеть записать некоторые значения в буферизованный канал, `c` пока он не заполнится, не входя в состояние блокировки в конце, и горутина является единственным отправителем канала, тогда мы можем использовать следующий код для выполнения этой работы.

```go
for len(c) < cap(c) {
	c <- aValue
}

```

Конечно, мы также можем использовать механизм try-send, представленный ниже, для выполнения той же работы.

### Заблокировать текущую горутину навсегда

Механизм выбора — уникальная функция Go. Он предлагает множество шаблонов и приемов для параллельного программирования. О правилах выполнения кода механизма select читайте в [каналах статей на Go](https://go101.org/article/channel.html#select) .

Мы можем использовать пустой блок выбора, `select{}` чтобы навсегда заблокировать текущую горутину. Это самый простой вариант использования механизма выбора. На самом деле, некоторые варианты использования `for {time.Sleep(time.Second)}` в приведенных выше примерах можно заменить на `select{}` .

Как правило, `select{}` используется для предотвращения выхода основной горутины, потому что, если основная горутина выйдет, вся программа также завершится.

Пример:

```go
package main

import "runtime"

func DoSomething() {
	for {
		// do something ...

		runtime.Gosched() // avoid being greedy
	}
}

func main() {
	go DoSomething()
	go DoSomething()
	select{}
}

```

Кстати, есть и [другие способы](https://go101.org/article/summaries.html#block-forever) заставить горутину навсегда остаться в заблокированном состоянии. Но `select{}` способ самый простой.

### Попробуйте-отправить и попробовать-получить

Блок `select` с одной `default` ветвью и только с одной `case` ветвью называется операцией канала "попытка-отправка" или "попытка-получить" в зависимости от того, является ли операция канала, следующая за `case` ключевым словом, операцией отправки или приема канала.

*   Если операция, следующая за `case` ключевым словом, является операцией отправки, то `select` блок называется операцией попытки отправки. Если операция отправки будет заблокирована, то `default` ветвь будет выполнена (не будет отправлена), в противном случае отправка будет успешной, и `case` будет выполнена единственная ветвь.
*   Если операция, следующая за `case` ключевым словом, является операцией приема, то `select` блок называется операцией попытки получения. Если операция получения будет заблокирована, то `default` ветвь будет выполнена (не будет получена), в противном случае получение будет выполнено успешно, и `case` будет выполнена единственная ветвь.

Операции try-send и try-receive никогда не блокируются.

Стандартный компилятор Go делает специальные оптимизации для блоков выбора try-send и try-receive, эффективность их выполнения намного выше, чем у блоков select с несколькими регистрами.

Ниже приведен пример, показывающий, как работают попытки отправки и получения.

```go
package main

import "fmt"

func main() {
	type Book struct{id int}
	bookshelf := make(chan Book, 3)

	for i := 0; i < cap(bookshelf) * 2; i++ {
		select {
		case bookshelf <- Book{id: i}:
			fmt.Println("succeeded to put book", i)
		default:
			fmt.Println("failed to put book")
		}
	}

	for i := 0; i < cap(bookshelf) * 2; i++ {
		select {
		case book := <-bookshelf:
			fmt.Println("succeeded to get book", book.id)
		default:
			fmt.Println("failed to get book")
		}
	}
}

```

Вывод вышеуказанной программы:

```
succeed to put book 0
succeed to put book 1
succeed to put book 2
failed to put book
failed to put book
failed to put book
succeed to get book 0
succeed to get book 1
succeed to get book 2
failed to get book
failed to get book
failed to get book

```

В следующих подразделах будет показано больше вариантов использования «попробуй-отправь» и «попробуй-получи».

#### Проверить, закрыт ли канал без блокировки текущей горутины

Предположим, что гарантируется, что никакие значения никогда не отправлялись (и не будут) отправлены в канал, мы можем использовать следующий код, чтобы (одновременно и безопасно) проверить, закрыт ли уже канал, не блокируя текущую горутину, где `T` тип элемента соответствующего типа канала.

```go
func IsClosed(c chan T) bool {
	select {
	case <-c:
		return true
	default:
	}
	return false
}

```

Способ проверки того, закрыт ли канал, широко используется в параллельном программировании Go, чтобы проверить, пришло ли уведомление. Уведомление будет отправлено путем закрытия канала в другой горутине.

#### Пиковое/импульсное ограничение

Мы можем реализовать пиковое ограничение, комбинируя [каналы использования в качестве счетных семафоров](https://go101.org/article/channel-use-cases.html#semaphore) и попытки отправки/попробования получения. Пиковый лимит (или пакетный лимит) часто используется для ограничения количества одновременных запросов без блокировки каких-либо запросов.

Ниже приведена модифицированная версия последнего примера в разделе [использования каналов в качестве счетных семафоров](https://go101.org/article/channel-use-cases.html#semaphore) .

```go
...
	// Can serve most 10 customers at the same time
	bar24x7 := make(Bar, 10)
	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		customer := Consumer{customerId}
		select {
		case bar24x7 <- customer: // try to enter the bar
			go bar24x7.ServeConsumer(customer)
		default:
			log.Print("customer#", customerId, " goes elsewhere")
		}
	}
...

```

#### Еще один способ реализации варианта использования «выигрывает первый ответ».

Как упоминалось выше, мы можем использовать механизм выбора (try-send) с буферизованным каналом, пропускная способность которого равна (как минимум) единице, чтобы реализовать вариант использования «выигрывает первый ответ». Например,

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func source(c chan<- int32) {
	ra, rb := rand.Int31(), rand.Intn(3)+1
	// Sleep 1s, 2s or 3s.
	time.Sleep(time.Duration(rb) * time.Second)
	select {
	case c <- ra:
	default:
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// The capacity should be at least 1.
	c := make(chan int32, 1)
	for i := 0; i < 5; i++ {
		go source(c)
	}
	rnd := <-c // only the first response is used
	fmt.Println(rnd)
}

```

Обратите внимание, пропускная способность канала, используемого в приведенном выше примере, должна быть не менее единицы, чтобы первая отправка не была пропущена, если сторона получателя/запроса не подготовилась вовремя.

#### Третий способ реализации варианта использования «Побеждает первый ответ».

Для варианта использования «выигрывает первый ответ», если количество источников невелико, например, два или три, мы можем использовать `select` блок кода для одновременного получения ответов источника. Например,

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func source() <-chan int32 {
	// c must be a buffered channel.
	c := make(chan int32, 1)
	go func() {
		ra, rb := rand.Int31(), rand.Intn(3)+1
		time.Sleep(time.Duration(rb) * time.Second)
		c <- ra
	}()
	return c
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var rnd int32
	// Blocking here until one source responses.
	select{
	case rnd = <-source():
	case rnd = <-source():
	case rnd = <-source():
	}
	fmt.Println(rnd)
}

```

Примечание: если канал, используемый в приведенном выше примере, является небуферизованным каналом, то две горутины будут зависать навсегда после выполнения `select` блока кода. Это [случай утечки памяти](https://go101.org/article/memory-leaking.html#hanging-goroutine) .

Два способа, представленные в текущем и последнем подразделах, также можно использовать для уведомлений N-to-1.

#### Тайм-аут

В некоторых сценариях запрос-ответ по разным причинам может потребоваться много времени для ответа на запрос, а иногда он даже никогда не ответит. В таких случаях мы должны вернуть сообщение об ошибке на сторону клиента, используя решение тайм-аута. Такое решение тайм-аута может быть реализовано с помощью механизма выбора.

В следующем коде показано, как сделать запрос с тайм-аутом.

```go
func requestWithTimeout(timeout time.Duration) (int, error) {
	c := make(chan int)
	// May need a long time to get the response.
	go doRequest(c)

	select {
	case data := <-c:
		return data, nil
	case <-time.After(timeout):
		return 0, errors.New("timeout")
	}
}

```

#### Бегущая строка

Мы можем использовать механизм try-send для реализации тикера.

```go
package main

import "fmt"
import "time"

func Tick(d time.Duration) <-chan struct{} {
	// The capacity of c is best set as one.
	c := make(chan struct{}, 1)
	go func() {
		for {
			time.Sleep(d)
			select {
			case c <- struct{}{}:
			default:
			}
		}
	}()
	return c
}

func main() {
	t := time.Now()
	for range Tick(time.Second) {
		fmt.Println(time.Since(t))
	}
}

```

На самом деле есть `Tick` функция в `time` стандартном пакете, обеспечивающая тот же функционал, но с гораздо более эффективной реализацией. Вместо этого мы должны использовать эту функцию, чтобы код выглядел чистым и работал эффективно.

#### Ограничение скорости

В одном из приведенных выше разделов показано, как использовать команду try-send для [ограничения пиковых значений](https://go101.org/article/channel-use-cases.html#peak-limiting) . Мы также можем использовать try-send для ограничения скорости (с помощью тикера). На практике ограничение скорости часто используется, чтобы избежать превышения квоты и исчерпания ресурсов.

Ниже показан такой пример, позаимствованный из [официальной вики Go](https://github.com/golang/go/wiki/RateLimiting) . В этом примере количество обработанных запросов за одну минуту не будет превышать 200.

```go
package main

import "fmt"
import "time"

type Request interface{}
func handle(r Request) {fmt.Println(r.(int))}

const RateLimitPeriod = time.Minute
const RateLimit = 200 // most 200 requests in one minute

func handleRequests(requests <-chan Request) {
	quotas := make(chan time.Time, RateLimit)

	go func() {
		tick := time.NewTicker(RateLimitPeriod / RateLimit)
		defer tick.Stop()
		for t := range tick.C {
			select {
			case quotas <- t:
			default:
			}
		}
	}()

	for r := range requests {
		<-quotas
		go handle(r)
	}
}

func main() {
	requests := make(chan Request)
	go handleRequests(requests)
	// time.Sleep(time.Minute)
	for i := 0; ; i++ {requests <- i}
}

```

На практике мы часто используем ограничение скорости и пиковое/всплесковое ограничение вместе.

#### Переключатели

Из статей о [каналах в Go](https://go101.org/article/channel.html) мы узнали, что отправка значения или получение значения из нулевого канала являются блокирующими операциями. Используя этот факт, мы можем изменить задействованные каналы в `case` операциях `select` кодового блока, чтобы повлиять на выбор ветви в `select` кодовом блоке.

Ниже приведен еще один пример пинг-понга, реализованный с использованием механизма выбора. В этом примере одной из двух переменных канала, задействованных в блоке выбора, является `nil` . Ветка, соответствующая нулевому каналу `case` , точно не будет выбрана. Мы можем думать, что такие `case` ветки находятся в выключенном состоянии. В конце каждого шага цикла состояния включения/выключения двух `case` ветвей переключаются.

```go
package main

import "fmt"
import "time"
import "os"

type Ball uint8
func Play(playerName string, table chan Ball, serve bool) {
	var receive, send chan Ball
	if serve {
		receive, send = nil, table
	} else {
		receive, send = table, nil
	}
	var lastValue Ball = 1
	for {
		select {
		case send <- lastValue:
		case value := <- receive:
			fmt.Println(playerName, value)
			value += lastValue
			if value < lastValue { // overflow
				os.Exit(0)
			}
			lastValue = value
		}
		// Switch on/off.
		receive, send = send, receive
		time.Sleep(time.Second)
	}
}

func main() {
	table := make(chan Ball)
	go Play("A:", table, false)
	Play("B:", table, true)
}

```

Ниже приведен еще один (непараллельный) пример, который намного проще и также демонстрирует эффект переключения. Этот пример будет напечатан `1212...` во время работы. Практической пользы от него мало. Он показан здесь только для целей обучения.

```go
package main

import "fmt"
import "time"

func main() {
	for c := make(chan struct{}, 1); true; {
		select {
		case c <- struct{}{}:
			fmt.Print("1")
		case <-c:
			fmt.Print("2")
		}
		time.Sleep(time.Second)
	}
}

```

#### Веса возможности выполнения управляющего кода

Мы можем дублировать `case` ветвь в `select` блоке кода, чтобы увеличить вес возможности выполнения соответствующего кода.

Пример:

```go
package main

import "fmt"

func main() {
	foo, bar := make(chan struct{}), make(chan struct{})
	close(foo); close(bar) // for demo purpose
	x, y := 0.0, 0.0
	f := func(){x++}
	g := func(){y++}
	for i := 0; i < 100000; i++ {
		select {
		case <-foo: f()
		case <-foo: f()
		case <-bar: g()
		}
	}
	fmt.Println(x/y) // about 2
}

```

Возможность `f` вызываемой функции примерно равна двойнику `g` вызываемой функции.

#### Выберите из динамического количества случаев

Хотя количество ветвей в `select` блоке фиксировано, мы можем использовать функциональные возможности, предоставляемые `reflect` стандартным пакетом, для создания блока выбора во время выполнения. Динамически создаваемый блок select может иметь произвольное количество ветвей case. Но обратите внимание, что способ отражения менее эффективен, чем фиксированный способ.

Стандартный `reflect` пакет также предоставляет `TrySend` и `TryRecv` выполняет функции для реализации блоков выбора «один случай плюс значение по умолчанию».

### Манипуляции с потоком данных

В этом разделе будут представлены некоторые варианты использования манипулирования потоками данных с использованием каналов.

Как правило, приложение потока данных состоит из множества модулей. Разные модули выполняют разную работу. Каждому модулю может принадлежать один или несколько воркеров (горутин), которые одновременно выполняют одну и ту же работу, указанную для этого модуля. Вот список некоторых примеров работы модулей на практике:

*   генерация/сбор/загрузка данных.
*   обслуживание/сохранение данных.
*   расчет/анализ данных.
*   проверка/фильтрация данных.
*   агрегация/разделение данных
*   компоновка/декомпозиция данных.
*   дублирование/распространение данных.

Обработчик в модуле может получать данные от нескольких других модулей в качестве входных данных и отправлять данные для обслуживания других модулей в качестве выходных данных. Другими словами, модуль может быть как потребителем данных, так и производителем данных. Модуль, который только отправляет данные некоторым другим модулям, но никогда не получает данные от других модулей, называется модулем только для производителя. Модуль, который получает данные только от некоторых других модулей, но никогда не отправляет данные другим модулям, называется модулем только для потребителя.

Вместе множество модулей образуют систему потока данных.

Ниже показаны некоторые реализации рабочих модулей потока данных. Эти реализации предназначены для пояснения, поэтому они очень просты и могут быть неэффективными.

#### Генерация/сбор/загрузка данных

Существуют все виды модулей только для производителей. Рабочий модуль только для производителя может создавать поток данных

*   загрузив файл, прочитав базу данных или просканировав Интернет.
*   путем сбора всех видов метрик из программной системы или всех видов аппаратного обеспечения.
*   путем генерации случайных чисел.
*   и т.п.

Здесь мы используем генератор случайных чисел в качестве примера. Функция-генератор возвращает один результат, но не принимает параметров.

```go
import (
	"crypto/rand"
	"encoding/binary"
)

func RandomGenerator() <-chan uint64 {
	c := make(chan uint64)
	go func() {
		rnds := make([]byte, 8)
		for {
			_, err := rand.Read(rnds)
			if err != nil {
				close(c)
				break
			}
			c <- binary.BigEndian.Uint64(rnds)
		}
	}()
	return c
}

```

По сути, генератор случайных чисел — это мультивозвратное будущее/обещание.

Производитель данных может закрыть канал выходного потока в любое время, чтобы прекратить генерацию данных.

#### Агрегация данных

Рабочий модуль модуля агрегации данных объединяет несколько потоков данных одного типа в один поток. Предположим, что тип данных — `int64` , тогда следующая функция объединит произвольное количество потоков данных в один.

```go
func Aggregator(inputs ...<-chan uint64) <-chan uint64 {
	out := make(chan uint64)
	for _, in := range inputs {
		go func(in <-chan uint64) {
			for {
				out <- <-in // <=> out <- (<-in)
			}
		}(in)
	}
	return out
}

```

Лучшая реализация должна учитывать, был ли закрыт входной поток. (Также действительно для следующих других реализаций рабочего модуля.)

```go
import "sync"

func Aggregator(inputs ...<-chan uint64) <-chan uint64 {
	output := make(chan uint64)
	var wg sync.WaitGroup
	for _, in := range inputs {
		wg.Add(1)
		go func(int <-chan uint64) {
			defer wg.Done()
			// If in is closed, then the
			// loop will ends eventually.
			for x := range in {
				output <- x
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(output)
	}()
	return output
}

```

Если количество агрегируемых потоков данных очень мало (два или три), мы можем использовать `select` блок для агрегирования этих потоков данных.

```go
// Assume the number of input stream is two.
...
	output := make(chan uint64)
	go func() {
		inA, inB := inputs[0], inputs[1]
		for {
			select {
			case v := <- inA: output <- v
			case v := <- inB: output <- v
			}
		}
	}
...

```

#### Раздел данных

Обработчик модуля разделения данных действует противоположно обработчику модуля агрегации данных. Реализовать работника подразделения легко, но на практике работники подразделения не очень полезны и используются редко.

```go
func Divisor(input <-chan uint64, outputs ...chan<- uint64) {
	for _, out := range outputs {
		go func(o chan<- uint64) {
			for {
				o <- <-input // <=> o <- (<-input)
			}
		}(out)
	}
}

```

#### Состав данных

Обработчик компоновки данных объединяет несколько фрагментов данных из разных потоков входных данных в один фрагмент данных.

Ниже приведен пример рабочего процесса композиции, в котором два `uint64` значения из одного потока и одно `uint64` значение из другого потока составляют одно новое `uint64` значение. Конечно, на практике эти типы элементов потокового канала обычно различны.

```go
func Composer(inA, inB <-chan uint64) <-chan uint64 {
	output := make(chan uint64)
	go func() {
		for {
			a1, b, a2 := <-inA, <-inB, <-inA
			output <- a1 ^ b & a2
		}
	}()
	return output
}

```

#### Декомпозиция данных

Декомпозиция данных — это процесс, обратный компоновке данных. Реализация рабочей функции декомпозиции принимает один параметр входного потока данных и возвращает несколько результатов потока данных. Здесь не будут показаны примеры для декомпозиции данных.

#### Дублирование/распространение данных

Дублирование данных (размножение) можно рассматривать как специальную декомпозицию данных. Один фрагмент данных будет продублирован, и каждый из дублированных данных будет отправлен в разные потоки выходных данных.

Пример:

```go
func Duplicator(in <-chan uint64) (<-chan uint64, <-chan uint64) {
	outA, outB := make(chan uint64), make(chan uint64)
	go func() {
		for x := range in {
			outA <- x
			outB <- x
		}
	}()
	return outA, outB
}

```

#### Расчет/анализ данных

Функциональные возможности модулей расчета и анализа данных различаются, и каждый из них очень специфичен. Как правило, рабочая функция таких модулей преобразует каждую часть входных данных в другую часть выходных данных.

Для простой демонстрации здесь показан рабочий пример, который инвертирует каждый бит каждого передаваемого `uint64` значения.

```go
func Calculator(in <-chan uint64, out chan uint64) (<-chan uint64) {
	if out == nil {
		out = make(chan uint64)
	}
	go func() {
		for x := range in {
			out <- ^x
		}
	}()
	return out
}

```

#### Проверка/фильтрация данных

Модуль проверки или фильтрации данных отбрасывает некоторые переданные данные в потоке. Например, следующая рабочая функция отбрасывает все непростые числа.

```go
import "math/big"

func Filter0(input <-chan uint64, output chan uint64) <-chan uint64 {
	if output == nil {
		output = make(chan uint64)
	}
	go func() {
		bigInt := big.NewInt(0)
		for x := range input {
			bigInt.SetUint64(x)
			if bigInt.ProbablyPrime(1) {
				output <- x
			}
		}
	}()
	return output
}

func Filter(input <-chan uint64) <-chan uint64 {
	return Filter0(input, nil)
}

```

Обратите внимание, что каждая из двух реализаций используется в одном из последних двух примеров, показанных ниже.

#### Обслуживание/сохранение данных

Как правило, модуль обслуживания или сохранения данных является последним или последним модулем вывода в системе потока данных. Здесь просто предоставляется простой воркер, который печатает каждый фрагмент данных, полученных из входного потока.

```go
import "fmt"

func Printer(input <-chan uint64) {
	for x := range input {
		fmt.Println(x)
	}
}

```

#### Сборка системы потока данных

Теперь давайте воспользуемся указанными выше рабочими функциями модуля для сборки нескольких систем потоков данных. Сборка системы потока данных — это просто создание нескольких воркеров из разных модулей и указание входных потоков для каждого воркера.

Пример системы потока данных 1 (линейный конвейер):

```go
package main

... // the worker functions declared above.

func main() {
	Printer(
		Filter(
			Calculator(
				RandomGenerator(), nil,
			),
		),
	)
}

```

Вышеупомянутая система потока данных изображена на следующей диаграмме.

![линейный поток данных](https://go101.org/article/res/data-flow-linear.png)

Пример системы потока данных 2 (направленный конвейер ациклического графа):

```go
package main

... // the worker functions declared above.

func main() {
	filterA := Filter(RandomGenerator())
	filterB := Filter(RandomGenerator())
	filterC := Filter(RandomGenerator())
	filter := Aggregator(filterA, filterB, filterC)
	calculatorA := Calculator(filter, nil)
	calculatorB := Calculator(filter, nil)
	calculator := Aggregator(calculatorA, calculatorB)
	Printer(calculator)
}

```

Вышеупомянутая система потока данных изображена на следующей диаграмме.

![Поток данных DAG](https://go101.org/article/res/data-flow-dag.png)

Более сложная топология потока данных может представлять собой произвольные графы. Например, система потока данных может иметь несколько конечных выходов. Но системы потоков данных с топологией циклического графа редко используются в реальности.

Из приведенных выше двух примеров мы можем сделать вывод, что создавать системы потоков данных с каналами очень просто и интуитивно понятно.

Из последнего примера мы можем обнаружить, что с помощью агрегаторов легко реализовать fan-in и fan-out для количества воркеров для указанного модуля.

Фактически, мы можем использовать простой канал, чтобы заменить роль агрегатора. Например, в следующем примере два агрегатора заменяются двумя каналами.

```go
package main

... // the worker functions declared above.

func main() {
	c1 := make(chan uint64, 100)
	Filter0(RandomGenerator(), c1) // filterA
	Filter0(RandomGenerator(), c1) // filterB
	Filter0(RandomGenerator(), c1) // filterC
	c2 := make(chan uint64, 100)
	Calculator(c1, c2) // calculatorA
	Calculator(c1, c2) // calculatorB
	Printer(c2)
}

```

Модифицированная система потока данных изображена на следующей диаграмме.

![Поток данных DAG](https://go101.org/article/res/data-flow-dag-b.png)

Приведенные выше объяснения для систем потоков данных не учитывают, как закрывать потоки данных. Прочтите [эту статью](https://go101.org/article/channel-closing.html) , чтобы узнать, как корректно закрывать каналы.
