# Golang Goroutines

---

Concurrency in Golang is the ability for functions to run independent of each other. Goroutines are functions that are run concurrently. Golang provides Goroutines as a way to handle operations concurrently.

New goroutines are created by the go statement.

To run a function as a goroutine, call that function prefixed with the go statement. Here is the example code block:

```jsx
sum()     // A normal function call that executes sum synchronously and waits for completing it
go sum()  // A goroutine that executes sum asynchronously and doesn't wait for completing it

```

The go keyword makes the function call to return immediately, while the function starts running in the background as a goroutine and the rest of the program continues its execution. The **main** function of every Golang program is started using a goroutine, so every Golang program runs at least one goroutine.

---

## Creating Goroutines

Added the go keyword before each call of function **responseSize**. The three **responseSize** goroutines starts up concurrently and three calls to *http.Get* are made concurrently as well. The program doesn't wait until one response comes back before sending out the next request. As a result the three response sizes are printed much sooner using goroutines.

```jsx
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
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
	go responseSize("https://www.golangprograms.com")
	go responseSize("https://coderwall.com")
	go responseSize("https://stackoverflow.com")
	time.Sleep(10 * time.Second)
}
```

We have added a call to *time.Sleep* in the **main** function which prevents the main goroutine from exiting before the **responseSize** goroutines can finish. Calling time.Sleep(10 \* time.Second) will make the **main** goroutine to sleep for 10 seconds.

You may see the following output when you run the above program −

```markup
C:\Golang\goroutines\create-simple-goroutine>go run main.go
Step1:  https://www.golangprograms.com
Step1:  https://stackoverflow.com
Step1:  https://coderwall.com
Step2:  https://stackoverflow.com
Step3:  https://stackoverflow.com
Step4:  116749
Step2:  https://www.golangprograms.com
Step3:  https://www.golangprograms.com
Step4:  79551
Step2:  https://coderwall.com
Step3:  https://coderwall.com
Step4:  203842

```

---

## Waiting for Goroutines to Finish Execution

The WaitGroup type of sync package, is used to wait for the program to finish all goroutines launched from the main function. It uses a counter that specifies the number of goroutines, and Wait blocks the execution of the program until the WaitGroup counter is zero.

The Add method is used to add a counter to the WaitGroup.

The Done method of WaitGroup is scheduled using a defer statement to decrement the WaitGroup counter.

The Wait method of the WaitGroup type waits for the program to finish all goroutines.

The Wait method is called inside the main function, which blocks execution until the WaitGroup counter reaches the value of zero and ensures that all goroutines are executed.

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

You may see the following output when you run the above program −

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

---

## Fetch Values from Goroutines

The most natural way to fetch a value from a goroutine is channels. Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine or in a synchronous function.

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

func responseSize(url string, nums chan int) {
	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Send value to the unbuffered channel
	nums <- len(body)
}

func main() {
	nums := make(chan int) // Declare a unbuffered channel
	wg.Add(1)
	go responseSize("https://www.golangprograms.com", nums)
	fmt.Println(<-nums) // Read the value from unbuffered channel
	wg.Wait()
	close(nums) // Closes the channel
}
```

You may see the following output when you run the above program −

```markup

C:\Golang\goroutines\create-simple-goroutine>go run main.go
79655
```

---

## Play and Pause Execution of Goroutine

Using channels we can play and pause execution of goroutine. A **channel** handles this communication by acting as a conduit between goroutines.

```jsx
package main

import (
	"fmt"
	"sync"
	"time"
)

var i int

func work() {
	time.Sleep(250 * time.Millisecond)
	i++
	fmt.Println(i)
}

func routine(command <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var status = "Play"
	for {
		select {
		case cmd := <-command:
			fmt.Println(cmd)
			switch cmd {
			case "Stop":
				return
			case "Pause":
				status = "Pause"
			default:
				status = "Play"
			}
		default:
			if status == "Play" {
				work()
			}
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	command := make(chan string)
	go routine(command, &wg)

	time.Sleep(1 * time.Second)
	command <- "Pause"

	time.Sleep(1 * time.Second)
	command <- "Play"

	time.Sleep(1 * time.Second)
	command <- "Stop"

	wg.Wait()
}
```

You can see the following output when you run the above program −

```markup
1
2
3
4
Pause
Play
5
6
7
8
9
Stop
```

---

## Fix Race Condition using Atomic Functions

Race conditions occur due to unsynchronized access to shared resource and attempt to read and write to that resource at the same time.

Atomic functions provide low\-level locking mechanisms for synchronizing access to integers and pointers. Atomic functions generally used to fix the race condition.

The functions in the **atomic** under **sync** packages provides support to synchronize goroutines by locking access to shared resources.

```jsx
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int32          // counter is a variable incremented by all goroutines.
	wg      sync.WaitGroup // wg is used to wait for the program to finish.
)

func main() {
	wg.Add(3) // Add a count of two, one for each goroutine.

	go increment("Python")
	go increment("Java")
	go increment("Golang")

	wg.Wait() // Wait for the goroutines to finish.
	fmt.Println("Counter:", counter)

}

func increment(name string) {
	defer wg.Done() // Schedule the call to Done to tell main we are done.

	for range name {
		atomic.AddInt32(&counter, 1)
		runtime.Gosched() // Yield the thread and be placed back in queue.
	}
}
```

The AddInt32 function from the atomic package synchronizes the adding of integer values by enforcing that only one goroutine can perform and complete this add operation at a time. When goroutines attempt to call any atomic function, they're automatically synchronized against the variable that's referenced.

You can see the following output when you run the above program −

```markup
C:\Golang\goroutines>go run -race main.go
Counter: 15

C:\Golang\goroutines>
```

Note if you replace the code line **atomic.AddInt32(&counter, 1)** with **counter++**, then you will see the below output\-

```markup
C:\Golang\goroutines>go run -race main.go
==================
WARNING: DATA RACE
Read at 0x0000006072b0 by goroutine 7:
  main.increment()
      C:/Golang/goroutines/main.go:31 +0x76

Previous write at 0x0000006072b0 by goroutine 8:
  main.increment()
      C:/Golang/goroutines/main.go:31 +0x90

Goroutine 7 (running) created at:
  main.main()
      C:/Golang/goroutines/main.go:18 +0x7e

Goroutine 8 (running) created at:
  main.main()
      C:/Golang/goroutines/main.go:19 +0x96
==================
Counter: 15
Found 1 data race(s)
exit status 66

C:\Golang\goroutines>
```

---

## Define Critical Sections using Mutex

A mutex is used to create a critical section around code that ensures only one goroutine at a time can execute that code section.

```jsx
package main

import (
	"fmt"
	"sync"
)

var (
	counter int32          // counter is a variable incremented by all goroutines.
	wg      sync.WaitGroup // wg is used to wait for the program to finish.
	mutex   sync.Mutex     // mutex is used to define a critical section of code.
)

func main() {
	wg.Add(3) // Add a count of two, one for each goroutine.

	go increment("Python")
	go increment("Go Programming Language")
	go increment("Java")

	wg.Wait() // Wait for the goroutines to finish.
	fmt.Println("Counter:", counter)

}

func increment(lang string) {
	defer wg.Done() // Schedule the call to Done to tell main we are done.

	for i := 0; i < 3; i++ {
		mutex.Lock()
		{
			fmt.Println(lang)
			counter++
		}
		mutex.Unlock()
	}
}
```

A critical section defined by the calls to **Lock()** and **Unlock()** protects the actions against the counter variable and reading the text of name variable. You can see the following output when you run the above program −

```markup

C:\Golang\goroutines>go run -race main.go
PHP stands for Hypertext Preprocessor.
PHP stands for Hypertext Preprocessor.
The Go Programming Language, also commonly referred to as Golang
The Go Programming Language, also commonly referred to as Golang
Counter: 4

C:\Golang\goroutines>
```
