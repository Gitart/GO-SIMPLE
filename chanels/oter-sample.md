## Switch statements
Switch statements may have multiple case values, and break is implicit:

switch day {
case 1, 2, 3, 4, 5:
    tag = "workday"
case 0, 6:
    tag = "weekend"
default:
    tag = "invalid"
}


## The case values don't have to be constants:

switch {
case day < 0 || day > 6:
    tag = "invalid"
case day == 0 || day == 6:
    tag = "weekend"
default:
    tag = "workday"
}


## For loops
    for i := 0; i < len(primes); i++ {
        fmt.Println(i, primes[i])
    }
A range clause permits easy iteration over arrays and slices:

    for i, x := range primes {
        fmt.Println(i, x)
    }

Unused values are discarded by assigning to the blank (_) identifier:

    var sum int
    for _, x := range primes {
        sum += x
    }



## Methods

package main
import "fmt"

type Point struct{ x, y int }

func PointToString(p Point) string {
    return fmt.Sprintf("Point{%d, %d}", p.x, p.y)
}

func (p Point) String() string {
    return fmt.Sprintf("Point{%d, %d}", p.x, p.y)
}

func main() {
    p := Point{3, 5}
    fmt.Println(PointToString(p)) // static dispatch
    fmt.Println(p.String())       // static dispatch
    fmt.Println(p)
}


cat
package main

import (
    "flag"
    "io"
    "os"
)

func main() {
    flag.Parse()
    for _, arg := range flag.Args() {
        f, err := os.Open(arg)
        if err != nil {
            panic(err)
        }
        defer f.Close()
        _, err = io.Copy(os.Stdout, f)
        if err != nil {
            panic(err)
        }
    }
}



A boring function
We need an example to show the interesting properties of the concurrency primitives.
To avoid distraction, we make it a boring example.

func main() {
    f("Hello, World", 500*time.Millisecond)
}
func f(msg string, delay time.Duration) {
    for i := 0; ; i++ {
        fmt.Println(msg, i)
        time.Sleep(delay)
    }
}



Ignoring it
The go statement runs the function as usual, but doesn't make the caller wait.
It launches a goroutine.
The functionality is analogous to the & on the end of a shell command.

func main() {
    go f("three", 300*time.Millisecond)
    go f("six", 600*time.Millisecond)
    go f("nine", 900*time.Millisecond)
}

Ignoring it a little less
When main returns, the program exits and takes the function f down with it.
We can hang around a little, and on the way show that both main and the launched goroutine are running.

func main() {
    go f("three", 300*time.Millisecond)
    go f("six", 600*time.Millisecond)
    go f("nine", 900*time.Millisecond)
    time.Sleep(3 * time.Second)
    fmt.Println("Done.")
}



Using channels
A channel connects the main and f goroutines so they can communicate.

func main() {
    c := make(chan string)
    go f("three", 300*time.Millisecond, c)
    for i := 0; i < 10; i++ {
        fmt.Println("Received", <-c) // Receive expression is just a value.
    }
    fmt.Println("Done.")
}


func f(msg string, delay time.Duration, c chan string) {
    for i := 0; ; i++ {
        c <- fmt.Sprintf("%s %d", msg, i) // Any suitable value can be sent.
        time.Sleep(delay)
    }
}


## Using channels between many goroutines
func main() {
    c := make(chan string)
    go f("three", 300*time.Millisecond, c)
    go f("six", 600*time.Millisecond, c)
    go f("nine", 900*time.Millisecond, c)
    for i := 0; i < 10; i++ {
        fmt.Println("Received", <-c)
    }
    fmt.Println("Done.")
}

A single channel may be used to communicate between many (not just two) goroutines; 
many goroutines may communicate via one or multiple channels.
This enables a rich variety of concurrency patterns.



## Elements of a work-stealing scheduler
func worker(in chan int, out chan []int) {
    for {
        order := <-in           // Receive a work order.
        result := factor(order) // Do some work.
        out <- result           // Send the result back.
    }
}


The worker uses two channels to communicate: 
- The in channel waits for some work order. 
- The out channel communicates the result. 
- As work load, a worker (very slowly) computes the list of prime factors for a given order.


## A matching producer and consumer
func producer(out chan int) {
    for order := 0; ; order++ {
        out <- order // Produce a work order.
    }
}

func consumer(in chan []int, n int) {
    for i := 0; i < n; i++ {
        result := <-in // Consume a result.
        fmt.Println("Consumed", result)
    }
}

The producer produces and endless supply of work orders and sends them out.
The consumer receives n results from the in channel and then terminates.


## Putting it all together

func main() {
    start := time.Now()

    in := make(chan int)    // Channel on which work orders are received.
    out := make(chan []int) // Channel on which results are returned.
    go producer(in)
    go worker(in, out) // Launch one worker.
    consumer(out, 100)

    fmt.Println(time.Since(start))
}


We use one worker to handle the entire work load.
Because there is only one worker, we see the result coming back in order.
This is running rather slow...


## Using 10 workers

    in  := make(chan int)
    out := make(chan []int)
    go producer(in)
    // Launch 10 workers.
    for i := 0; i < 10; i++ {
        go worker(in, out)
    }
    consumer(out, 100)


A ready worker will read the next order from the in channel and start working on it. Another ready worker will proceed with the next order, and so forth.
Because we have many workers and since different orders take different amounts of time to work on, we see the results coming back out-of-order.
On a multi-core system, many workers may truly run in parallel.
This is running much faster...
