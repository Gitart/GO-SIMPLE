# Подробнее об отложенных вызовах функций

Отложенные вызовы функций были [введены ранее](https://go101.org/article/control-flows-more.html#defer) . Из-за ограниченного знания Go в то время некоторые дополнительные детали и варианты использования вызовов отложенных функций не затрагиваются в этой статье. Эти детали и варианты использования будут затронуты в оставшейся части этой статьи.

### Вызовы многих встроенных функций с возвращаемыми результатами не могут быть отложены

В Go все результирующие значения вызова пользовательских функций могут отсутствовать (отбрасываться). Однако для встроенных функций с непустыми списками возвращаемых результатов значения результатов их вызовов [не должны отсутствовать](https://go101.org/article/exceptions.html#discard-return-results) (по крайней мере, для Go 1.19), за исключением вызовов встроенных `copy` и `recover` функций. С другой стороны, мы узнали, что значения результата отложенного вызова функции должны быть отброшены, поэтому вызовы многих встроенных функций не могут быть отложены.

К счастью, необходимость откладывать вызовы встроенных функций (с непустыми списками возвращаемых результатов) на практике встречается редко. Насколько я знаю, `append` иногда требуется отложить только вызовы встроенной функции. В этом случае мы можем отложить вызов анонимной функции, которая упаковывает `append` вызов.

```go
package main

import "fmt"

func main() {
	s := []string{"a", "b", "c", "d"}
	defer fmt.Println(s) // [a x y d]
	// defer append(s[:1], "x", "y") // error
	defer func() {
		_ = append(s[:1], "x", "y")
	}()
}

```

### Момент оценки значений отложенной функции

Вызываемая функция (значение) в вызове отложенной функции оценивается, когда вызов помещается в очередь отложенных вызовов текущей горутины. Например, следующая программа напечатает `false` .

```go
package main

import "fmt"

func main() {
	var f = func () {
		fmt.Println(false)
	}
	defer f()
	f = func () {
		fmt.Println(true)
	}
}

```

Вызываемая функция в отложенном вызове функции может быть нулевым значением функции. В таком случае паника будет происходить при вызове функции nil, а не при помещении вызова в очередь отложенных вызовов текущей горутины. Пример:

```go
package main

import "fmt"

func main() {
	defer fmt.Println("reachable 1")
	var f func() // f is nil by default
	defer f()    // panic here
	// The following lines are also reachable.
	fmt.Println("reachable 2")
	f = func() {} // useless to avoid panicking
}

```

### Момент оценки аргументов получателя отложенных вызовов методов

Как объяснялось ранее, аргументы отложенного вызова функции [также оцениваются, когда](https://go101.org/article/control-flows-more.html#argument-evaluation-moment) отложенный вызов помещается в очередь отложенных вызовов текущей горутины.

Аргументы получателя метода также не являются исключением. Например, следующая программа печатает `1342` .

```go
package main

type T int

func (t T) M(n int) T {
  print(n)
  return t
}

func main() {
	var t T
	// "t.M(1)" is the receiver argument of the method
	// call ".M(2)", so it is evaluated when the
	// ".M(2)" call is pushed into deferred call queue.
	defer t.M(1).M(2)
	t.M(3).M(4)
}

```

### Отложенные вызовы делают код более чистым и менее подверженным ошибкам

Пример:

```go
import "os"

func withoutDefers(filepath string, head, body []byte) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}

	_, err = f.Seek(16, 0)
	if err != nil {
		f.Close()
		return err
	}

	_, err = f.Write(head)
	if err != nil {
		f.Close()
		return err
	}

	_, err = f.Write(body)
	if err != nil {
		f.Close()
		return err
	}

	err = f.Sync()
	f.Close()
	return err
}

func withDefers(filepath string, head, body []byte) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Seek(16, 0)
	if err != nil {
		return err
	}

	_, err = f.Write(head)
	if err != nil {
		return err
	}

	_, err = f.Write(body)
	if err != nil {
		return err
	}

	return f.Sync()
}

```

Какой из них выглядит чище? Судя по всему, тот, что с отложенными звонками, хоть и немного. И это менее подвержено ошибкам, поскольку `f.Close()` в функции так много вызовов без отложенных вызовов, что у нее больше шансов пропустить один из них.

Ниже приведен еще один пример, показывающий, что отложенные вызовы могут сделать код менее подверженным ошибкам. Если `doSomething` вызовы паникуют в следующем примере, функция `f2` завершится без разблокировки `Mutex` значения. Таким образом, функция `f1` менее подвержена ошибкам.

```go
var m sync.Mutex

func f1() {
	m.Lock()
	defer m.Unlock()
	doSomething()
}

func f2() {
	m.Lock()
	doSomething()
	m.Unlock()
}

```

### Потери производительности из-за отложенных вызовов функций

Не всегда хорошо использовать отложенные вызовы функций. Для официального компилятора Go до версии 1.13 отложенные вызовы функций приведут к некоторым потерям производительности во время выполнения. Начиная с Go Toolchain 1.13, некоторые распространенные варианты использования отложенных вызовов были значительно оптимизированы, так что, как правило, нам не нужно заботиться о проблеме потери производительности, вызванной отложенными вызовами. Благодарим Дэна Скейлза за отличную оптимизацию.

### Вид утечки ресурсов из-за отложенных вызовов функций

Очень большая очередь отложенных вызовов может также потреблять много памяти, и некоторые ресурсы могут не освобождаться вовремя, если некоторые вызовы задерживаются слишком сильно. Например, если при вызове следующей функции необходимо обработать много файлов, то большое количество обработчиков файлов не будет освобождено до выхода из функции.

```go
func writeManyFiles(files []File) error {
	for _, file := range files {
		f, err := os.Open(file.path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.WriteString(file.content)
		if err != nil {
			return err
		}

		err = f.Sync()
		if err != nil {
			return err
		}
	}

	return nil
}

```

В таких случаях мы можем использовать анонимную функцию для включения отложенных вызовов, чтобы отложенные вызовы функций выполнялись раньше. Например, вышеприведенная функция может быть переписана и улучшена как

```go
func writeManyFiles(files []File) error {
	for _, file := range files {
		if err := func() error {
			f, err := os.Open(file.path)
			if err != nil {
				return err
			}
			// The close method will be called at
			// the end of the current loop step.
			defer f.Close()

			_, err = f.WriteString(file.content)
			if err != nil {
				return err
			}

			return f.Sync()
		}(); err != nil {
			return err
		}
	}

	return nil
}
```
