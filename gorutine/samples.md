# Go goroutine

last modified May 9, 2022

Go goroutine tutorial shows how to work with goroutines in Golang. A goroutine is a lightweight execution thread.
Goroutine definition

Goroutine is a lightweight execution thread. It is a function that runs concurrency alongside other running code. 
Note that concurrent execution may or may not parallel. In Go, every program has at least one goroutine: the main goroutine.

A goroutine is started with the go keywords.
Executing functions in sequence

The following example runs a function one by one.
main.go

package main

import (
    "fmt"
)

func main() {

    hello("Martin")
    hello("Lucia")
    hello("Michal")
    hello("Jozef")
    hello("Peter")
}

func hello(name string) {

    fmt.Printf("Hello %s!\n", name)
}

The program runs the hello function in a sequence.

$ go run main.go
Hello Martin!
Hello Lucia!
Hello Michal!
Hello Jozef!
Hello Peter!

The output is always the same list.
Executing functions concurrently

Now we run the hello function concurrently.
main.go

package main

import (
    "fmt"
)

func main() {

    go hello("Martin")
    go hello("Lucia")
    go hello("Michal")
    go hello("Jozef")
    go hello("Peter")

    fmt.Scanln()
}

func hello(name string) {

    fmt.Printf("Hello %s!\n", name)
}

With the go keyword, we run the hello function conturrently. The fmt.Scanln function waits for the input from user. If we comment out this function, the program finishes before we can see output from the goroutines.

$ go run main.go
Hello Lucia!
Hello Michal!
Hello Martin!
Hello Jozef!
Hello Peter!
$ go run main.go
Hello Martin!
Hello Peter!
Hello Lucia!
Hello Michal!
Hello Jozef!

We run the program twice. Note that the output is different.
Go sync.WaitGroup

The sync.WaitGroup is a synchronization tool which waits for a collection of goroutines to finish.
main.go

package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {

    var wg sync.WaitGroup
    wg.Add(2)

    go func() {

        count("oranges")
        wg.Done()
    }()

    go func() {

        count("apples")
        wg.Done()
    }()

    wg.Wait()
}

func count(thing string) {

    for i := 0; i < 4; i++ {

        fmt.Printf("counting %s\n", thing)
        time.Sleep(time.Millisecond * 500)
    }
}

In the program, we synchronize the execution of two goroutines with sync.WaitGroup.

var wg sync.WaitGroup
wg.Add(2)

With Add, we tell how may goroutines we wait for.

go func() {

    count("oranges")
    wg.Done()
}()

We create an anonymous goroutine. We tell the Go runtime that the goroutine has finished with Done.

wg.Wait()

The Wait function blocks until all goroutines have finished.

time.Sleep(time.Millisecond * 500)

In demonstrational programs, time.Sleep is often used to slow down the execution of goroutines.

$ go run main.go 
counting apples
counting oranges
counting apples
counting oranges
counting oranges
counting apples
counting apples
counting oranges

Go async requests

The next example uses goroutines to make asynchronous requests.
async_req.go

package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "regexp"
  "sync"
)

func main() {
  
  urls := []string{
    "http://webcode.me",
    "https://example.com",
    "http://httpbin.org",
    "https://www.perl.org",
    "https://www.php.net",
    "https://www.python.org",
    "https://code.visualstudio.com",
    "https://clojure.org",
  }

  var wg sync.WaitGroup

  for _, u := range urls {
    
    wg.Add(1)
    go func(url string) {
    
      defer wg.Done()
    
      content := doReq(url)
      title := getTitle(content)
      fmt.Println(title)
    }(u)
  }

  wg.Wait()
}

func doReq(url string) (content string) {

    resp, err := http.Get(url)

    if err != nil {
        log.Println(err)
        return
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Println(err)
        return
    }

    return string(body)
}

func getTitle(content string) (title string) {

  re := regexp.MustCompile("<title>(.*)</title>")

  parts := re.FindStringSubmatch(content)

  if len(parts) > 0 {
    return parts[1]
  } else {
    return "no title"
  }
}

We make multiple asynchronous HTTP requests. We get the contents of the title tag of each of the web pages. Each request is wrapped inside one goroutine.

go func(url string) {

  defer wg.Done()

  content := doReq(url)
  title := getTitle(content)
  fmt.Println(title)
}(u)

Withing the goroutine, we generate a GET request, receive the response, get the title from the response and print it to the terminal.

$ go run main.go 
The Perl Programming Language - www.perl.org
Welcome to Python.org
Visual Studio Code - Code Editing. Redefined
PHP: Hypertext Preprocessor
Example Domain
httpbin.org
Clojure
My html page

Goroutine channnels

Goroutines communicate via channnels. They allow to send and receive values with the channel operator, <-.

c := make(chan string)

A new channel is created with the make function.

c <- v    // send
v := <-c  // receive

The channel operator sends and receives values between goroutines.
main.go

package main

import (
    "fmt"
    "time"
)

func main() {

    c := make(chan string)
    go hello("Martin", c)

    for msg := range c {

        fmt.Println(msg)
    }

}

func hello(name string, c chan string) {

    for i := 0; i < 5; i++ {

        msg := fmt.Sprintf("Hello %s!", name)
        c <- msg
        time.Sleep(time.Millisecond * 500)
    }

    close(c)
}

In the program, two goroutines communicate: main and hello.

c := make(chan string)

A channel is created with make.

go hello("Martin", c)

A hello goroutine is created with go. We pass the channel as a parameter.

for msg := range c {

    fmt.Println(msg)
}

With the range operator, we go through the messages and print them to the console.

func hello(name string, c chan string) {

    for i := 0; i < 5; i++ {

        msg := fmt.Sprintf("Hello %s!", name)
        c <- msg
        time.Sleep(time.Millisecond * 500)
    }

    close(c)
}

In the hello goroutine, we create a message and send it via the channel to the main goroutine. When the goroutine is finished, we close the channel with close.

$ go run main.go
Hello Martin!
Hello Martin!
Hello Martin!
Hello Martin!
Hello Martin!

Calculating fibonacci values with goroutines

In the next example, we calculate fibonacci numbers with goroutines.
main.go

package main

import (
    "fmt"
)

func fib(n int, c chan int) {

    x, y := 0, 1

    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)
}

func main() {

    c := make(chan int, 10)

    go fib(cap(c), c)

    for i := range c {
        fmt.Println(i)
    }
}

A series of fibonacci values is generated inside the fib goroutine. The values are one by one send to the caller goroutine via a channel.

$ go run main.go
0
1
1
2
3
5
8
13
21
34
