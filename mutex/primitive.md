Примитивы синхронизации в Go
German Gorelkin
German Gorelkin
Feb 9, 2020 · 5 min read

Продолжаем серию статей о проблемах многопоточности, параллелизме, concurrency и других интересных штуках.

    Race condition и Data Race
    Deadlocks, Livelocks и Starvation
    Примитивы синхронизации в Go
    Безопасная работа с каналами в Go
    Goroutine Leaks

    Пакет sync содержит примитивы, которые наиболее полезны для низкоуровневой синхронизации доступа к памяти.

WaitGroup

    WaitGroup — это отличный способ дождаться завершения набора одновременных операций.

Запустим несколько goroutine и дождемся завершения их работы:

```go
var wg sync.WaitGroup

wg.Add(1)
go func() {
   defer wg.Done()
   
   fmt.Println("1st goroutine sleeping...")
   time.Sleep(100 * time.Millisecond)
}()

wg.Add(1)
go func() {
   defer wg.Done()

   fmt.Println("2nd goroutine sleeping...")
   time.Sleep(200 * time.Millisecond)
}()

wg.Wait()
fmt.Println("All goroutines complete.")
```

У нас нет гарантий когда будут запущены наши goroutine. Возможна ситуация когда при вызове 
Wait еще не будет ни одной запущенной goroutine. По этому важно вызвать Add за пределами 
процедур, которые они помогают отслеживать.

Пример неопределенного поведения:

```go
var wg sync.WaitGroupgo func() {
  wg.Add(1)
  defer wg.Done()
  fmt.Println("1st goroutine sleeping...")
  time.Sleep(1)
 }()wg.Wait()
fmt.Println("All goroutines complete.")
```

## О WaitGroup можно думать как о concurrent-safe счетчике.

Вызовы Add увеличивает счетчик на переданное число, а вызовы Done уменьшают счетчик на единицу. Wait блокируется пока счетчик не станет равным нулю.

Обычно Add вызывают как можно ближе к goroutine. Но иногда удобно используют Add для отслеживания группы goroutine одновременно. Например в таких циклах:

```go
hello := func(wg *sync.WaitGroup, id int) {
   defer wg.Done()
   fmt.Printf("Hello from %v!\n", id)
}

const numGreeters = 5
var wg sync.WaitGroup

wg.Add(numGreeters)
for i := 0; i < numGreeters; i++ {
   go hello(&wg, i+1)
}

wg.Wait()
```

## GermanGorelkin/go-patterns
Design patterns implemented in Golang. Contribute to GermanGorelkin/go-patterns development by creating an account on…

```go
github.com
Mutex and RWMutex
```

**Mutex** означает mutual exclusion(взаимное исключение) и 
является способом защиты critical section(критическая секция) вашей программы.

**Критическая секция** — это область вашей программы, которая требует эксклюзивного доступа к общему ресурсу. 
При нахождении в критической секции двух (или более) потоков возникает состояние race(гонки). 
Так же возможны проблемы взаимной блокировки(deadlock).

Mutex обеспечивает безопасный доступ к общим ресурсам.

Простой пример счетчика:

```go
type counter struct{
   count int
}
func (c *counter) Increment() {
   c.count++
}
func (c *counter) Decrement() {
   c.count--
}
```

Напишем тест, который будет в разных goroutine увеличивать или уменьшать общее значение:

```go
c := new(counter)

var wg sync.WaitGroup
numLoop := 1000

wg.Add(numLoop)
for i := 0; i < numLoop; i++ {
   go func() {
      defer wg.Done()
      c.Increment()
   }()
}

wg.Add(numLoop)
for i := 0; i < numLoop; i++ {
   go func() {
      defer wg.Done()
      c.Decrement()
   }()
}

wg.Wait()

expected := 0
assert.Equal(t, expected, c.count)
```

Результат:

```
expected: 0
actual:   52
```


Используем Mutex для синхронизации доступа:

```go
type counter struct{
   sync.Mutex
   count int
}
func (c *counter) Increment() {
   c.Lock()
   defer c.Unlock()
   c.count++
}
func (c *counter) Decrement() {
   c.Lock()
   defer c.Unlock()
   c.count--
}
```

Мы вызываем Unlock в defer. Это очень распространенная идиома при использовании Mutex, 
чтобы гарантировать, что вызов всегда происходит, даже при панике. Несоблюдение этого 
требования может привести к deadlock вашей программы. Хотя defer и несет небольшие затраты.

Критическая секция названа так, потому что она отражает узкое место в вашей программе. 
Вход в критическую секцию и выход из нее обходится довольно дорого, поэтому обычно 
люди пытаются минимизировать время, проведенное в критических секциях.

Возможно не все процессы будут читать и записывать в общую память. 
В этом случае вы можете воспользоваться мьютексом другого типа.
RWMutex

    RWMutex концептуально то же самое, что и Mutex: он защищает доступ к памяти. 
    Тем не менее, RWMutex дает вам немного больше контроля над памятью. Вы можете 
    запросить блокировку для чтения, и в этом случае вам будет предоставлен доступ, 
    если блокировка не удерживается для записи.

Это означает, что произвольное число читателей может удерживать 
блокировку читателя, пока ничто другое не удерживает блокировку писателя.

Посмотрим как это работает:

```go
func (c *counter) CountV1() int {
   c.Lock()
   defer c.Unlock()
   return c.count
}
func (c *counter) CountV2() int {
   c.RLock()
   defer c.RUnlock()
   return c.count
}
```

CountV2 не блокирует count если не было блокировок на запись.

Немного бенчмарков:

```go
func BenchmarkCountV1(b *testing.B) {
   c := new(counter)
   var wg sync.WaitGroup
   for i := 0; i < b.N; i++ {
      for j := 0; j < 1000; j++ {
         wg.Add(1)
         go func() {
            defer wg.Done()
            c.CountV1()
         }()
      }
      wg.Wait()
   }
}

func BenchmarkCountV2(b *testing.B) {
   c := new(counter)
   var wg sync.WaitGroup
   for i := 0; i < b.N; i++ {
      for j := 0; j < 1000; j++ {
         wg.Add(1)
         go func() {
            defer wg.Done()
            c.CountV2()
         }()
      }
      wg.Wait()
   }
}
```

Результаты:
```
enchmarkCountV1-8           2132            501896 ns/op
BenchmarkCountV2-8          3358            306254 ns/op
```


