# Некоторые варианты использования при панике/восстановлении

Паника и восстановление были [введены ранее](https://go101.org/article/control-flows-more.html#panic-recover) . Далее в текущей статье будут представлены некоторые (хорошие и плохие) варианты использования паники/восстановления.

### Вариант использования 1: избегайте паники, вызывающей сбой программ

Это должен быть самый популярный вариант использования panic/recover. Вариант использования обычно используется в параллельных программах, особенно в программах клиент-сервер.

Пример:

```go
package main

import "errors"
import "log"
import "net"

func main() {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		// Handle each client connection
		// in a new goroutine.
		go ClientHandler(conn)
	}
}

func ClientHandler(c net.Conn) {
	defer func() {
		if v := recover(); v != nil {
			log.Println("capture a panic:", v)
			log.Println("avoid crashing the program")
		}
		c.Close()
	}()
	panic(errors.New("just a demo.")) // a demo-purpose panic
}

```

Запустив сервер и запустив его `telnet localhost 12345` в другом терминале, мы можем заметить, что сервер не рухнет из-за паники, созданной в горутине каждого обработчика клиента.

Если мы не восстановим потенциальную панику в каждой горутине обработчика клиента, потенциальная паника приведет к сбою программы.

### Вариант использования 2: автоматический перезапуск сбойной горутины

Когда в горутине обнаруживается паника, мы можем создать для нее новую горутину. Пример:

```go
package main

import "log"
import "time"

func shouldNotExit() {
	for {
		// Simulate a workload.
		time.Sleep(time.Second)

		// Simulate an unexpected panic.
		if time.Now().UnixNano() & 0x3 == 0 {
			panic("unexpected situation")
		}
	}
}

func NeverExit(name string, f func()) {
	defer func() {
		if v := recover(); v != nil {
			// A panic is detected.
			log.Println(name, "is crashed. Restart it now.")
			go NeverExit(name, f) // restart
		}
	}()
	f()
}

func main() {
	log.SetFlags(0)
	go NeverExit("job#A", shouldNotExit)
	go NeverExit("job#B", shouldNotExit)
	select{} // block here for ever
}

```

### Вариант использования 3: Использование `panic` / `recover` вызовы для имитации операторов длинного перехода

Иногда мы можем использовать panic/recover как способ симулировать операторы длинного перехода кросс-функции и возвраты кросс-функции, хотя обычно этот способ использовать не рекомендуется. Такой способ вредит как читабельности кода, так и эффективности выполнения. Единственным преимуществом является то, что иногда код может выглядеть менее подробным.

В следующем примере, когда во внутренней функции создается паника, выполнение переходит к отложенному вызову.

```go
package main

import "fmt"

func main() {
	n := func () (result int)  {
		defer func() {
			if v := recover(); v != nil {
				if n, ok := v.(int); ok {
					result = n
				}
			}
		}()

		func () {
			func () {
				func () {
					// ...
					panic(123) // panic on succeeded
				}()
				// ...
			}()
		}()
		// ...
		return 0
	}()
	fmt.Println(n) // 123
}

```

### Вариант использования 4: использование вызовов `panic` / `recover` для уменьшения количества проверок на наличие ошибок

Пример:

```go
func doSomething() (err error) {
	defer func() {
		err = recover()
	}()

	doStep1()
	doStep2()
	doStep3()
	doStep4()
	doStep5()

	return
}

// In reality, the prototypes of the doStepN functions
// might be different. For each of them,
// * panic with nil for success and no needs to continue.
// * panic with error for failure and no needs to continue.
// * not panic for continuing.
func doStepN() {
	...
	if err != nil {
		panic(err)
	}
	...
	if done {
		panic(nil)
	}
}

```

Приведенный выше код менее подробен, чем следующий.

```go
func doSomething() (err error) {
	shouldContinue, err := doStep1()
	if !shouldContinue {
		return err
	}
	shouldContinue, err = doStep2()
	if !shouldContinue {
		return err
	}
	shouldContinue, err = doStep3()
	if !shouldContinue {
		return err
	}
	shouldContinue, err = doStep4()
	if !shouldContinue {
		return err
	}
	shouldContinue, err = doStep5()
	if !shouldContinue {
		return err
	}

	return
}

// If err is not nil, then shouldContinue must be true.
// If shouldContinue is true, err might be nil or non-nil.
func doStepN() (shouldContinue bool, err error) {
	...
	if err != nil {
		return false, err
	}
	...
	if done {
		return false, nil
	}
	return true, nil
}

```

Однако, как правило, этот шаблон использования паники/восстановления не рекомендуется использовать. Это менее идиоматично и менее эффективно.
