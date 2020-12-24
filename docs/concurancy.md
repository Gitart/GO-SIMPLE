# Golang Concurrency

Concurrency is an ability of a program to do multiple things at the same time. This means a program that have two or more tasks that run individually of each other, at about the same time, but remain part of the same program. Concurrency is very important in modern software, due to the need to execute independent pieces of code as fast as possible without disturbing the overall flow of the program.

Concurrency in Golang is the ability for functions to run independent of each other. A goroutine is a function that is capable of running concurrently with other functions. When you create a function as a goroutine, it has been treated as an independent unit of work that gets scheduled and then executed on an available logical processor. The Golang runtime scheduler has feature to manages all the goroutines that are created and need processor time. The scheduler binds operating system's threads to logical processors in order to execute the goroutines. By sitting on top of the operating system, scheduler controls everything related to which goroutines are running on which logical processors at any given time.

Popular programming languages such as Java and Python implement concurrency by using threads. Golang has built\-in concurrency constructs: goroutines and channels. Concurrency in Golang is cheap and easy. Goroutines are cheap, lightweight threads. Channels, are the conduits that allow for communication between goroutines.

Communicating Sequential Processes, or CSP for short, is used to describe how systems that feature multiple concurrent models should interact with one another. It typically relies heavily on using channels as a medium for passing messages between two or more concurrent processes, and is the underlying mantra of Golang.

**Goroutines** — A goroutine is a function that runs independently of the function that started it.
**Channels** — A channel is a pipeline for sending and receiving data. Channels provide a way for one goroutine to send structured data to another.

*Concurrency and parallelism comes into the picture when you are examining for multitasking and they are often used interchangeably, concurrent and parallel refer to related but different things.*

**Concurrency** \- Concurrency is about to handle numerous tasks at once. This means that you are working to manage numerous tasks done at once in a given period of time. However, you will only be doing a single task at a time. This tends to happen in programs where one task is waiting and the program determines to drive another task in the idle time. It is an aspect of the problem domain — where your program needs to handle numerous simultaneous events.

**Parallelism** \- Parallelism is about doing lots of tasks at once. This means that even if we have two tasks, they are continuously working without any breaks in between them. It is an aspect of the solution domain — where you want to make your program faster by processing different portions of the problem in parallel.

A concurrent program has multiple logical threads of control. These threads may or may not run in parallel. A parallel program potentially runs more quickly than a sequential program by executing different parts of the computation simultaneously (in parallel). It may or may not have more than one logical thread of control.

---

## Illustration of the dining philosophers problem in Golang

Five silent philosophers sit at a round table with bowls of spaghetti. Forks are placed between each pair of adjacent philosophers. Each philosopher must alternately think and eat. However, a philosopher can only eat spaghetti when they have both left and right forks. Each fork can be held by only one philosopher and so a philosopher can use the fork only if it is not being used by another philosopher. After an individual philosopher finishes eating, they need to put down both forks so that the forks become available to others. A philosopher can take the fork on their right or the one on their left as they become available, but cannot start eating before getting both forks. Eating is not limited by the remaining amounts of spaghetti or stomach space; an infinite supply and an infinite demand are assumed. The problem is how to design a discipline of behavior (a concurrent algorithm) such that no philosopher will starve; i.e., each can forever continue to alternate between eating and thinking, assuming that no philosopher can know when others may want to eat or think.

package main

import (
	"hash/fnv"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

// Number of philosophers is simply the length of this list.
var ph = \[\]string{"Mark", "Russell", "Rocky", "Haris", "Root"}

const hunger = 3                // Number of times each philosopher eats
const think = time.Second / 100 // Mean think time
const eat = time.Second / 100   // Mean eat time

var fmt = log.New(os.Stdout, "", 0)

var dining sync.WaitGroup

func diningProblem(phName string, dominantHand, otherHand \*sync.Mutex) {
	fmt.Println(phName, "Seated")
	h := fnv.New64a()
	h.Write(\[\]byte(phName))
	rg := rand.New(rand.NewSource(int64(h.Sum64())))
	rSleep := func(t time.Duration) {
		time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
	}
	for h := hunger; h > 0; h\-\- {
		fmt.Println(phName, "Hungry")
		dominantHand.Lock() // pick up forks
		otherHand.Lock()
		fmt.Println(phName, "Eating")
		rSleep(eat)
		dominantHand.Unlock() // put down forks
		otherHand.Unlock()
		fmt.Println(phName, "Thinking")
		rSleep(think)
	}
	fmt.Println(phName, "Satisfied")
	dining.Done()
	fmt.Println(phName, "Left the table")
}

func main() {
	fmt.Println("Table empty")
	dining.Add(5)
	fork0 := &sync.Mutex{}
	forkLeft := fork0
	for i := 1; i < len(ph); i++ {
		forkRight := &sync.Mutex{}
		go diningProblem(ph\[i\], forkLeft, forkRight)
		forkLeft = forkRight
	}
	go diningProblem(ph\[0\], fork0, forkLeft)
	dining.Wait() // wait for philosphers to finish
	fmt.Println("Table empty")
}

Table empty
Mark seated
Mark Hungry
Mark Eating
..................
..................
Haris Thinking
Haris Satisfied
Haris Left the table
Table empty

---

## Illustration of Checkpoint Synchronization in Golang

The checkpoint synchronization is a problem of synchronizing multiple tasks. Consider a workshop where several workers assembling details of some mechanism. When each of them completes his work, they put the details together. There is no store, so a worker who finished its part first must wait for others before starting another one. Putting details together is the checkpoint at which tasks synchronize themselves before going their paths apart.

package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func worker(part string) {
	log.Println(part, "worker begins part")
	time.Sleep(time.Duration(rand.Int63n(1e6)))
	log.Println(part, "worker completes part")
	wg.Done()
}

var (
	partList    = \[\]string{"A", "B", "C", "D"}
	nAssemblies = 3
	wg          sync.WaitGroup
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for c := 1; c <= nAssemblies; c++ {
		log.Println("begin assembly cycle", c)
		wg.Add(len(partList))
		for \_, part := range partList {
			go worker(part)
		}
		wg.Wait()
		log.Println("assemble.  cycle", c, "complete")
	}
}

2019/07/15 16:10:32 begin assembly cycle 1
2019/07/15 16:10:32 D worker begins part
2019/07/15 16:10:32 A worker begins part
2019/07/15 16:10:32 B worker begins part
........
2019/07/15 16:10:32 D worker completes part
2019/07/15 16:10:32 C worker completes part
2019/07/15 16:10:32 assemble.  cycle 3 complete

---

## Illustration of Producer Consumer Problem in Golang

The problem describes two processes, the producer and the consumer, who share a common, fixed\-size buffer used as a queue. The producer's job is to generate data, put it into the buffer, and start again. At the same time, the consumer is consuming the data (i.e., removing it from the buffer), one piece at a time. The problem is to make sure that the producer won't try to add data into the buffer if it's full and that the consumer won't try to remove data from an empty buffer. The solution for the producer is to either go to sleep or discard data if the buffer is full. The next time the consumer removes an item from the buffer, it notifies the producer, who starts to fill the buffer again. In the same way, the consumer can go to sleep if it finds the buffer empty. The next time the producer puts data into the buffer, it wakes up the sleeping consumer.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

type Consumer struct {
	msgs \*chan int
}

// NewConsumer creates a Consumer
func NewConsumer(msgs \*chan int) \*Consumer {
	return &Consumer{msgs: msgs}
}

// consume reads the msgs channel
func (c \*Consumer) consume() {
	fmt.Println("consume: Started")
	for {
		msg := <\-\*c.msgs
		fmt.Println("consume: Received:", msg)
	}
}

// Producer definition
type Producer struct {
	msgs \*chan int
	done \*chan bool
}

// NewProducer creates a Producer
func NewProducer(msgs \*chan int, done \*chan bool) \*Producer {
	return &Producer{msgs: msgs, done: done}
}

// produce creates and sends the message through msgs channel
func (p \*Producer) produce(max int) {
	fmt.Println("produce: Started")
	for i := 0; i < max; i++ {
		fmt.Println("produce: Sending ", i)
		\*p.msgs <\- i
	}
	\*p.done <\- true // signal when done
	fmt.Println("produce: Done")
}

func main() {
	// profile flags
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to \`file\`")
	memprofile := flag.String("memprofile", "", "write memory profile to \`file\`")

	// get the maximum number of messages from flags
	max := flag.Int("n", 5, "defines the number of messages")

	flag.Parse()

	// utilize the max num of cores available
	runtime.GOMAXPROCS(runtime.NumCPU())

	// CPU Profile
	if \*cpuprofile != "" {
		f, err := os.Create(\*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	var msgs = make(chan int)  // channel to send messages
	var done = make(chan bool) // channel to control when production is done

	// Start a goroutine for Produce.produce
	go NewProducer(&msgs, &done).produce(\*max)

	// Start a goroutine for Consumer.consume
	go NewConsumer(&msgs).consume()

	// Finish the program when the production is done
	<\-done

	// Memory Profile
	if \*memprofile != "" {
		f, err := os.Create(\*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up\-to\-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}

consume: Started
produce: Started
produce: Sending  0
produce: Sending  1
consume: Received: 0
consume: Received: 1
produce: Sending  2
produce: Sending  3
consume: Received: 2
consume: Received: 3
produce: Sending  4
produce: Done

---

## Illustration of Sleeping Barber Problem in Golang

The barber has one barber's chair in a cutting room and a waiting room containing a number of chairs in it. When the barber finishes cutting a customer's hair, he dismisses the customer and goes to the waiting room to see if there are others waiting. If there are, he brings one of them back to the chair and cuts their hair. If there are none, he returns to the chair and sleeps in it. Each customer, when they arrive, looks to see what the barber is doing. If the barber is sleeping, the customer wakes him up and sits in the cutting room chair. If the barber is cutting hair, the customer stays in the waiting room. If there is a free chair in the waiting room, the customer sits in it and waits their turn. If there is no free chair, the customer leaves.

package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	sleeping = iota
	checking
	cutting
)

var stateLog = map\[int\]string{
	0: "Sleeping",
	1: "Checking",
	2: "Cutting",
}
var wg \*sync.WaitGroup // Amount of potentional customers

type Barber struct {
	name string
	sync.Mutex
	state    int // Sleeping/Checking/Cutting
	customer \*Customer
}

type Customer struct {
	name string
}

func (c \*Customer) String() string {
	return fmt.Sprintf("%p", c)\[7:\]
}

func NewBarber() (b \*Barber) {
	return &Barber{
		name:  "Sam",
		state: sleeping,
	}
}

// Barber goroutine
// Checks for customers
// Sleeps \- wait for wakers to wake him up
func barber(b \*Barber, wr chan \*Customer, wakers chan \*Customer) {
	for {
		b.Lock()
		defer b.Unlock()
		b.state = checking
		b.customer = nil

		// checking the waiting room
		fmt.Printf("Checking waiting room: %d\\n", len(wr))
		time.Sleep(time.Millisecond \* 100)
		select {
		case c := <\-wr:
			HairCut(c, b)
			b.Unlock()
		default: // Waiting room is empty
			fmt.Printf("Sleeping Barber \- %s\\n", b.customer)
			b.state = sleeping
			b.customer = nil
			b.Unlock()
			c := <\-wakers
			b.Lock()
			fmt.Printf("Woken by %s\\n", c)
			HairCut(c, b)
			b.Unlock()
		}
	}
}

func HairCut(c \*Customer, b \*Barber) {
	b.state = cutting
	b.customer = c
	b.Unlock()
	fmt.Printf("Cutting  %s hair\\n", c)
	time.Sleep(time.Millisecond \* 100)
	b.Lock()
	wg.Done()
	b.customer = nil
}

// customer goroutine
// just fizzles out if it's full, otherwise the customer
// is passed along to the channel handling it's haircut etc
func customer(c \*Customer, b \*Barber, wr chan<\- \*Customer, wakers chan<\- \*Customer) {
	// arrive
	time.Sleep(time.Millisecond \* 50)
	// Check on barber
	b.Lock()
	fmt.Printf("Customer %s checks %s barber | room: %d, w %d \- customer: %s\\n",
		c, stateLog\[b.state\], len(wr), len(wakers), b.customer)
	switch b.state {
	case sleeping:
		select {
		case wakers <\- c:
		default:
			select {
			case wr <\- c:
			default:
				wg.Done()
			}
		}
	case cutting:
		select {
		case wr <\- c:
		default: // Full waiting room, leave shop
			wg.Done()
		}
	case checking:
		panic("Customer shouldn't check for the Barber when Barber is Checking the waiting room")
	}
	b.Unlock()
}

func main() {
	b := NewBarber()
	b.name = "Rocky"
	WaitingRoom := make(chan \*Customer, 5) // 5 chairs
	Wakers := make(chan \*Customer, 1)      // Only one waker at a time
	go barber(b, WaitingRoom, Wakers)

	time.Sleep(time.Millisecond \* 100)
	wg = new(sync.WaitGroup)
	n := 10
	wg.Add(10)
	// Spawn customers
	for i := 0; i < n; i++ {
		time.Sleep(time.Millisecond \* 50)
		c := new(Customer)
		go customer(c, b, WaitingRoom, Wakers)
	}

	wg.Wait()
	fmt.Println("No more customers for the day")
}

Checking waiting room: 0
Sleeping Barber \- <nil\>
Customer 120 checks Sleeping barber | room: 0, w 0 \- customer: <nil\>
Woken by 120
..............
..............
Checking waiting room: 0
No more customers for the day

---

## Illustration of Cigarette Smokers Problem in Golang

Assume a cigarette requires three ingredients to make and smoke: tobacco, paper, and matches. There are three smokers around a table, each of whom has an infinite supply of one of the three ingredients — one smoker has an infinite supply of tobacco, another has paper, and the third has matches. A fourth party, with an unlimited supply of everything, chooses at random a smoker, and put on the table the supplies needed for a cigarrette. The chosen smoker smokes, and the process should repeat indefinitely.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	paper = iota
	grass
	match
)

var smokeMap = map\[int\]string{
	paper: "paper",
	grass: "grass",
	match: "match",
}

var names = map\[int\]string{
	paper: "Sandy",
	grass: "Apple",
	match: "Daisy",
}

type Table struct {
	paper chan int
	grass chan int
	match chan int
}

func arbitrate(t \*Table, smokers \[3\]chan int) {
	for {
		time.Sleep(time.Millisecond \* 500)
		next := rand.Intn(3)
		fmt.Printf("Table chooses %s: %s\\n", smokeMap\[next\], names\[next\])
		switch next {
		case paper:
			t.grass <\- 1
			t.match <\- 1
		case grass:
			t.paper <\- 1
			t.match <\- 1
		case match:
			t.grass <\- 1
			t.paper <\- 1
		}
		for \_, smoker := range smokers {
			smoker <\- next
		}
		wg.Add(1)
		wg.Wait()
	}
}

func smoker(t \*Table, name string, smokes int, signal chan int) {
	var chosen = \-1
	for {
		chosen = <\-signal // blocks

		if smokes != chosen {
			continue
		}

		fmt.Printf("Table: %d grass: %d match: %d\\n", len(t.paper), len(t.grass), len(t.match))
		select {
		case <\-t.paper:
		case <\-t.grass:
		case <\-t.match:
		}
		fmt.Printf("Table: %d grass: %d match: %d\\n", len(t.paper), len(t.grass), len(t.match))
		time.Sleep(10 \* time.Millisecond)
		select {
		case <\-t.paper:
		case <\-t.grass:
		case <\-t.match:
		}
		fmt.Printf("Table: %d grass: %d match: %d\\n", len(t.paper), len(t.grass), len(t.match))
		fmt.Printf("%s smokes a cigarette\\n", name)
		time.Sleep(time.Millisecond \* 500)
		wg.Done()
		time.Sleep(time.Millisecond \* 100)
	}
}

const LIMIT = 1

var wg \*sync.WaitGroup

func main() {
	wg = new(sync.WaitGroup)
	table := new(Table)
	table.match = make(chan int, LIMIT)
	table.paper = make(chan int, LIMIT)
	table.grass = make(chan int, LIMIT)
	var signals \[3\]chan int
	// three smokers
	for i := 0; i < 3; i++ {
		signal := make(chan int, 1)
		signals\[i\] = signal
		go smoker(table, names\[i\], i, signal)
	}
	fmt.Printf("%s, %s, %s, sit with \\n%s, %s, %s\\n\\n", names\[0\], names\[1\], names\[2\], smokeMap\[0\], smokeMap\[1\], smokeMap\[2\])
	arbitrate(table, signals)
}

Sandy, Apple, Daisy, sit with
paper, grass, match

Table chooses match: Daisy
Table: 1 grass: 1 match: 0
Table: 1 grass: 0 match: 0
Table: 0 grass: 0 match: 0
Daisy smokes a cigarette
Table chooses paper: Sandy
Table: 0 grass: 1 match: 1
Table: 0 grass: 1 match: 0
Table: 0 grass: 0 match: 0
Sandy smokes a cigarette
Table chooses match: Daisy
Table: 1 grass: 1 match: 0
Table: 1 grass: 0 match: 0
Table: 0 grass: 0 match: 0
Daisy smokes a cigarette
