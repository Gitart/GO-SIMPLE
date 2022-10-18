## Как изящно закрыть каналы
Несколько дней назад я написал статью, в которой объясняются правила канала в Go . Эта статья получила много голосов на Reddit и HN , но есть и некоторые критические замечания по поводу дизайна канала Go.

Я собрал некоторые критические замечания по поводу следующего дизайна и правил каналов Go:
нет простых и универсальных способов проверить, закрыт ли канал, без изменения статуса канала.
закрытие закрытого канала вызовет панику, поэтому опасно закрывать канал, если доводчики не знают, закрыт канал или нет.
отправка значений в закрытый канал вызовет панику, поэтому опасно отправлять значения в канал, если отправители не знают, закрыт ли канал.
Критика выглядит разумной (на самом деле нет). Да, действительно нет встроенной функции проверки, закрылся канал или нет.

На самом деле существует простой способ проверить, закрыт ли канал, если вы можете убедиться, что никакие значения не были (и не будут) когда-либо отправлены в канал. Метод был показан в прошлой статье . Здесь для большей согласованности метод снова указан в следующем примере.

```go
package main

import "fmt"

type T int

func IsClosed(ch <-chan T) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}

func main() {
	c := make(chan T)
	fmt.Println(IsClosed(c)) // false
	close(c)
	fmt.Println(IsClosed(c)) // true
}
```

Как было сказано выше, это не универсальный способ проверить, закрыт ли канал.

На самом деле, даже если есть простая встроенная `closed` функция для проверки того, был ли закрыт канал, ее полезность будет очень ограниченной, как и встроенная `len` функция для проверки текущего количества значений, хранящихся в буфере значений. канала. Причина в том, что статус проверяемого канала мог измениться сразу после возврата из вызова таких функций, так что возвращаемое значение уже не могло отражать последний статус только что проверенного канала. Хотя прекратить отправку значений в канал `ch` при `closed(ch)` возврате вызова `true` можно, закрывать канал или продолжать отправлять значения в канал при `closed(ch)` возврате вызова небезопасно `false` .

### Принцип закрытия канала

Один общий принцип использования каналов Go — **не закрывать канал со стороны получателя и не закрывать канал, если у канала есть несколько одновременных отправителей** . Другими словами, мы должны закрывать канал в горутине отправителя только в том случае, если отправитель является единственным отправителем канала.

*(Ниже мы будем называть вышеуказанный принцип принципом **закрытия канала** .)*

Конечно, это не универсальный принцип закрытия каналов. Универсальный принцип — **не закрывать (и не отправлять значения) закрытые каналы** . Если мы можем гарантировать, что никакие горутины больше не будут закрываться и отправлять значения в незакрытый канал, отличный от nil, тогда горутина может безопасно закрыть канал. Однако предоставление таких гарантий получателем или одним из многих отправителей канала обычно требует больших усилий и часто усложняет код. Наоборот, гораздо проще придерживаться упомянутого выше **принципа закрытия канала .**

### Решения, которые грубо закрывают каналы

Если вы в любом случае закроете канал со стороны получателя или одного из нескольких отправителей канала, то вы можете использовать [механизм восстановления,](https://go101.org/article/control-flows-more.html#panic-recover) чтобы предотвратить возможную панику из-за сбоя вашей программы. Вот пример (предположим, что тип элемента канала — `T` ).


```go
func SafeClose(ch chan T) (justClosed bool) {
	defer func() {
		if recover() != nil {
			// The return result can be altered
			// in a defer function call.
			justClosed = false
		}
	}()

	// assume ch != nil here.
	close(ch)   // panic if ch is closed
	return true // <=> justClosed = true; return
}

```

Это решение явно нарушает **принцип закрытия канала** .

Та же идея может быть использована для отправки значений в потенциально закрытый канал.

```go
func SafeSend(ch chan T, value T) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()

	ch <- value  // panic if ch is closed
	return false // <=> closed = false; return
}

```

Мало того, что грубое решение нарушает **принцип закрытия канала** , так еще и в процессе может произойти гонка данных.

### Решения, которые вежливо закрывают каналы

Многие люди предпочитают использовать `sync.Once` для закрытия каналов:

```go
type MyChannel struct {
	C    chan T
	once sync.Once
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan T)}
}

func (mc *MyChannel) SafeClose() {
	mc.once.Do(func() {
		close(mc.C)
	})
}

```

Конечно, мы также можем использовать `sync.Mutex` , чтобы избежать многократного закрытия канала:

```go
type MyChannel struct {
	C      chan T
	closed bool
	mutex  sync.Mutex
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan T)}
}

func (mc *MyChannel) SafeClose() {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	if !mc.closed {
		close(mc.C)
		mc.closed = true
	}
}

func (mc *MyChannel) IsClosed() bool {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	return mc.closed
}

```

Эти способы могут быть вежливыми, но они не могут избежать гонки данных. В настоящее время спецификация Go не гарантирует отсутствие гонок данных, когда канал закрывается, а операции отправки канала выполняются одновременно. Если `SafeClose` функция вызывается одновременно с операцией отправки канала в тот же канал, может произойти гонка данных (хотя такая гонка данных обычно не приносит никакого вреда).

### Решения, которые изящно закрывают каналы

Одним из недостатков вышеупомянутой `SafeSend` функции является то, что ее вызовы нельзя использовать в качестве операций отправки, которые следуют за `case` ключевым словом в `select` блоках. Другим недостатком вышеизложенного `SafeSend` и `SafeClose` функций является то, что многие люди, включая меня, сочли бы приведенные выше решения с использованием `panic` / `recover` и `sync` пакета некрасивыми. Далее будут представлены некоторые чисто канальные решения без использования `panic` / `recover` и `sync` пакета для всех видов ситуаций.

*(В следующих примерах `sync.WaitGroup` используется для полноты примеров. Его использование на практике может быть не всегда необходимым.)*

#### 1\. М получателей, один отправитель, отправитель говорит "больше не отправляет", закрывая канал данных

Это самая простая ситуация, просто позвольте отправителю закрыть канал данных, когда он не хочет больше отправлять.

```go
package main

import (
	"time"
	"math/rand"
	"sync"
	"log"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumReceivers = 100

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)

	// the sender
	go func() {
		for {
			if value := rand.Intn(Max); value == 0 {
				// The only sender can close the
				// channel at any time safely.
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()

			// Receive values until dataCh is
			// closed and the value buffer queue
			// of dataCh becomes empty.
			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()
}

```

#### 2\. Один получатель, N отправителей, единственный получатель говорит «пожалуйста, прекратите отправлять больше», закрывая дополнительный сигнальный канал.

Это ситуация немного сложнее, чем описанная выше. Мы не можем позволить приемнику закрыть канал данных, чтобы остановить передачу данных, так как это нарушит **принцип закрытия канала** . Но мы можем позволить получателю закрыть дополнительный сигнальный канал, чтобы уведомить отправителей о прекращении отправки значений.

```go
package main

import (
	"time"
	"math/rand"
	"sync"
	"log"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)

	// ...
	dataCh := make(chan int)
	stopCh := make(chan struct{})
		// stopCh is an additional signal channel.
		// Its sender is the receiver of channel
		// dataCh, and its receivers are the
		// senders of channel dataCh.

	// senders
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				// The try-receive operation is to try
				// to exit the goroutine as early as
				// possible. For this specified example,
				// it is not essential.
				select {
				case <- stopCh:
					return
				default:
				}

				// Even if stopCh is closed, the first
				// branch in the second select may be
				// still not selected for some loops if
				// the send to dataCh is also unblocked.
				// But this is acceptable for this
				// example, so the first select block
				// above can be omitted.
				select {
				case <- stopCh:
					return
				case dataCh <- rand.Intn(Max):
				}
			}
		}()
	}

	// the receiver
	go func() {
		defer wgReceivers.Done()

		for value := range dataCh {
			if value == Max-1 {
				// The receiver of channel dataCh is
				// also the sender of stopCh. It is
				// safe to close the stop channel here.
				close(stopCh)
				return
			}

			log.Println(value)
		}
	}()

	// ...
	wgReceivers.Wait()
}

```

Как упоминалось в комментариях, для дополнительного канала сигнала его отправитель является получателем канала данных. Канал дополнительного сигнала закрывается его единственным отправителем, который придерживается **принципа закрытия канала** .

В этом примере канал `dataCh` никогда не закрывается. Да, каналы не должны быть закрыты. Канал в конечном итоге будет удален сборщиком мусора, если никакие горутины больше не ссылаются на него, независимо от того, закрыт он или нет. Таким образом, грациозность закрытия канала здесь не в том, чтобы закрыть канал.

#### 3\. M получателей, N отправителей, любой из них говорит «давайте закончим игру», уведомляя модератора о закрытии дополнительного канала сигнала.

Это сложнейшая ситуация. Мы не можем позволить ни получателям, ни отправителям закрыть канал данных. И мы не можем позволить любому из получателей закрыть дополнительный сигнальный канал, чтобы уведомить всех отправителей и получателей о выходе из игры. Выполнение любого из них нарушит **принцип закрытия канала** . Однако мы можем ввести роль модератора, чтобы закрыть дополнительный сигнальный канал. Один трюк в следующем примере заключается в том, как использовать операцию try-send, чтобы уведомить модератора о закрытии дополнительного сигнального канала.

```go
package main

import (
	"time"
	"math/rand"
	"sync"
	"log"
	"strconv"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumReceivers = 10
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)
	stopCh := make(chan struct{})
		// stopCh is an additional signal channel.
		// Its sender is the moderator goroutine shown
		// below, and its receivers are all senders
		// and receivers of dataCh.
	toStop := make(chan string, 1)
		// The channel toStop is used to notify the
		// moderator to close the additional signal
		// channel (stopCh). Its senders are any senders
		// and receivers of dataCh, and its receiver is
		// the moderator goroutine shown below.
		// It must be a buffered channel.

	var stoppedBy string

	// moderator
	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					// Here, the try-send operation is
					// to notify the moderator to close
					// the additional signal channel.
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}

				// The try-receive operation here is to
				// try to exit the sender goroutine as
				// early as possible. Try-receive and
				// try-send select blocks are specially
				// optimized by the standard Go
				// compiler, so they are very efficient.
				select {
				case <- stopCh:
					return
				default:
				}

				// Even if stopCh is closed, the first
				// branch in this select block might be
				// still not selected for some loops
				// (and for ever in theory) if the send
				// to dataCh is also non-blocking. If
				// this is unacceptable, then the above
				// try-receive operation is essential.
				select {
				case <- stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wgReceivers.Done()

			for {
				// Same as the sender goroutine, the
				// try-receive operation here is to
				// try to exit the receiver goroutine
				// as early as possible.
				select {
				case <- stopCh:
					return
				default:
				}

				// Even if stopCh is closed, the first
				// branch in this select block might be
				// still not selected for some loops
				// (and forever in theory) if the receive
				// from dataCh is also non-blocking. If
				// this is not acceptable, then the above
				// try-receive operation is essential.
				select {
				case <- stopCh:
					return
				case value := <-dataCh:
					if value == Max-1 {
						// Here, the same trick is
						// used to notify the moderator
						// to close the additional
						// signal channel.
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}

					log.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	// ...
	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)
}

```

В этом примере **сохраняется принцип закрытия канала** .

Обратите внимание, что размер буфера (емкость) канала `toStop` равен единице. Это делается для того, чтобы избежать пропуска первого уведомления при его отправке до того, как горутина модератора будет готова получить уведомление от `toStop` .

Мы также можем установить пропускную способность `toStop` канала как сумму отправителей и получателей, тогда нам не нужен блок попытки отправки `select` для уведомления модератора.

```go
...
toStop := make(chan string, NumReceivers + NumSenders)
...
			value := rand.Intn(Max)
			if value == 0 {
				toStop <- "sender#" + id
				return
			}
...
				if value == Max-1 {
					toStop <- "receiver#" + id
					return
				}
...

```

#### 4\. Вариант ситуации «М получателей, один отправитель»: запрос на закрытие делает сторонняя горутина

Иногда необходимо, чтобы сигнал закрытия был сделан сторонней горутиной. В таких случаях мы можем использовать дополнительный сигнальный канал, чтобы уведомить отправителя о закрытии канала данных. Например,

```go
package main

import (
	"time"
	"math/rand"
	"sync"
	"log"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumReceivers = 100
	const NumThirdParties = 15

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)
	closing := make(chan struct{}) // signal channel
	closed := make(chan struct{})

	// The stop function can be called
	// multiple times safely.
	stop := func() {
		select {
		case closing<-struct{}{}:
			<-closed
		case <-closed:
		}
	}

	// some third-party goroutines
	for i := 0; i < NumThirdParties; i++ {
		go func() {
			r := 1 + rand.Intn(3)
			time.Sleep(time.Duration(r) * time.Second)
			stop()
		}()
	}

	// the sender
	go func() {
		defer func() {
			close(closed)
			close(dataCh)
		}()

		for {
			select{
			case <-closing: return
			default:
			}

			select{
			case <-closing: return
			case dataCh <- rand.Intn(Max):
			}
		}
	}()

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()

			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()
}

```

Идея, использованная в `stop` функции, взята из [комментария](https://groups.google.com/forum/#!msg/golang-nuts/lEKehHH7kZY/SRmCtXDZAAAJ) Роджера Пеппе.

#### 5\. Вариант ситуации «N отправителя»: канал данных должен быть закрыт, чтобы сообщить получателям, что отправка данных завершена.

В решениях для описанных выше ситуаций с N-отправителями, чтобы придерживаться **принципа закрытия канала** , мы избегаем закрытия каналов данных. Однако иногда требуется, чтобы каналы данных были закрыты в конце, чтобы получатели знали, что отправка данных завершена. В таких случаях мы можем преобразовать ситуацию с N отправителями в ситуацию с одним отправителем, используя промежуточный канал. Средний канал имеет только одного отправителя, поэтому мы можем закрыть его вместо закрытия исходного канала данных.

```go
package main

import (
	"time"
	"math/rand"
	"sync"
	"log"
	"strconv"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const Max = 1000000
	const NumReceivers = 10
	const NumSenders = 1000
	const NumThirdParties = 15

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)     // will be closed
	middleCh := make(chan int)   // will never be closed
	closing := make(chan string) // signal channel
	closed := make(chan struct{})

	var stoppedBy string

	// The stop function can be called
	// multiple times safely.
	stop := func(by string) {
		select {
		case closing <- by:
			<-closed
		case <-closed:
		}
	}

	// the middle layer
	go func() {
		exit := func(v int, needSend bool) {
			close(closed)
			if needSend {
				dataCh <- v
			}
			close(dataCh)
		}

		for {
			select {
			case stoppedBy = <-closing:
				exit(0, false)
				return
			case v := <- middleCh:
				select {
				case stoppedBy = <-closing:
					exit(v, true)
					return
				case dataCh <- v:
				}
			}
		}
	}()

	// some third-party goroutines
	for i := 0; i < NumThirdParties; i++ {
		go func(id string) {
			r := 1 + rand.Intn(3)
			time.Sleep(time.Duration(r) * time.Second)
			stop("3rd-party#" + id)
		}(strconv.Itoa(i))
	}

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					stop("sender#" + id)
					return
				}

				select {
				case <- closed:
					return
				default:
				}

				select {
				case <- closed:
					return
				case middleCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for range [NumReceivers]struct{}{} {
		go func() {
			defer wgReceivers.Done()

			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	// ...
	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)
}

```

#### Больше ситуаций?

Вариантов ситуаций должно быть больше, но показанные выше являются наиболее распространенными и основными. Я считаю, что при грамотном использовании каналов (и других методов параллельного программирования) всегда можно найти решение, поддерживающее **принцип закрытия канала для каждого варианта ситуации.**

### Вывод

Нет ситуаций, которые заставят вас нарушить **принцип закрытия канала** . Если вы столкнулись с такой ситуацией, пожалуйста, переосмыслите свой дизайн и перепишите код.

Программирование с помощью каналов Go похоже на искусство.














