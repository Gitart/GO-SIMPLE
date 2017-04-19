## Go routines wit chanels


Программа запускает три goroutines:    
Один для генерации значений Фибоначчи 10 случайным образом
Сгенерированные числа;   
Другой - для генерирования квадратичных значений 10 случайных чисел; а также  
Последний - для печати результирующих значений, генерируемых первой и второй горутинами.   
Из основных Функции, два канала создаются для связи значений Фибоначчи и генерируемых квадратов   
Соответствующими горатинами. Функция generateFibonacci запускается как goroutine, которая выполняет   
Отправьте операцию в каналы fibs для получения значений Фибоначчи. Запускается функция generateSquare   
Как goroutine, который выполняет операцию отправки в каналы sqrs для получения значений квадрата. 
Функция  PrintValues запускается как goroutine, который опрашивает как fibs, так и sqrs каналы для печати результирующего
Значения, когда значения могут принимать оба канала.

Внутри функции printValues выражение select используется с двумя блоками case. 
Блок выбора. В 20 раз, используя выражение цикла for. Мы используем 20 раз для печати 10 значений Фибоначчи и 10 квадратных
значения. В сценарии реального мира вы можете запускать это в бесконечном цикле, в котором вы могли бы быть.

Постоянно взаимодействуя с каналами.

```golang
package main
import (
    "fmt"
    "math"
    "math/rand"
    "sync"
)

type (
fibvalue struct {input, value int}
squarevalue struct {input, value int}
)

func generateSquare(sqrs chan<- squarevalue) {
defer wg.Done()
    for i := 1; i <= 10; i++ {
        num := rand.Intn(50)
        sqrs <- squarevalue{input: num,value: num * num}
    }

}

func generateFibonacci(fibs chan<- fibvalue) {
defer wg.Done()
for i := 1; i <= 10; i++ {
    num := float64(rand.Intn(50))
    // Fibonacci using Binet's formula
    Phi := (1 + math.Sqrt(5)) / 2
    phi := (1 - math.Sqrt(5)) / 2
    result := (math.Pow(Phi, num) - math.Pow(phi, num)) / math.Sqrt(5)
    fibs <- fibvalue{input: int(num),value: int(result)}
}
}

func printValues(fibs <-chan fibvalue, sqrs <-chan squarevalue) {
defer wg.Done()
    for i := 1; i <= 20; i++ {
        select {
          case fib := <-fibs: fmt.Printf("Fibonacci value of %d is %d\n", fib.input, fib.value)
          case sqr := <-sqrs: fmt.Printf("Square value of %d is %d\n", sqr.input, sqr.value)
        }
    }
}

// wg is used to wait for the program to finish.
var wg sync.WaitGroup


func main() {
    wg.Add(3)
    // Create Channels
    fibs := make(chan fibvalue)
    sqrs := make(chan squarevalue)
    // Launching 3 goroutines
    go generateFibonacci(fibs)
    go generateSquare(sqrs)
    go printValues(fibs, sqrs)
    wg.Wait()// Wait for completing all goroutines
}
```
