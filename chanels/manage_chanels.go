# How to Manage Go Channels With Range and Close

## Read data from a channel and then close the channel

[![Abhishek Gupta](https://miro.medium.com/fit/c/56/56/1*ZdSjzBPYoAeJU5cB5f2_xA.jpeg)](https://abhishek1987.medium.com/?source=post_page-----98f93f6e8c0c-----------------------------------)

[

Abhishek Gupta

](https://abhishek1987.medium.com/?source=post_page-----98f93f6e8c0c-----------------------------------)

[

Apr 20, 2020·4 min read

](https://betterprogramming.pub/manging-go-channels-with-range-and-close-98f93f6e8c0c?source=post_page-----98f93f6e8c0c-----------------------------------)

![](https://miro.medium.com/max/1400/1*QlFeCL_pOFLHgvUitMT1ZA.jpeg)

Photo by [Drew Beamer](https://unsplash.com/@drew_beamer?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) on [Unsplash](https://unsplash.com/s/photos/numbers?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText)

Welcome another part of the “[Just Enough Go” series](https://medium.com/@abhishek1987/just-enough-go-blog-series-c1cd62b04beb) in which we’ll go over how to use `range` to read data from a channel and `close` to shut it down.

[

## “Just Enough Go” — blog series

### Welcome! 👋👋

medium.com

](https://medium.com/@abhishek1987/just-enough-go-blog-series-c1cd62b04beb)

To run the different cases/examples, please use the [code on GitHub](https://github.com/abhirockzz/just-enough-go/blob/master/channels-range-close/channels-range-close.go).

You can use `<-` (e.g., `<-myChannel`) to accept values from a channel.

# Simple Scenario

func f1() {
	c := make(chan int) //producer
	go func() {
	 for i := 1; i <= 5; i++ {
	  c <- i
	  time.Sleep(1 \* time.Second)
	 }
	}() //consumer
	go func() {
	 for x := 1; x <= 5; x++ {
	  i := <-c
          fmt.Println("i =", i)
         }
         fmt.Println("press ctrl+c to exit")
        }() e := make(chan os.Signal)
	signal.Notify(e, syscall.SIGINT, syscall.SIGTERM)
	<-e
}

The producer goroutine sends five integers, and the consumer goroutine accepts them. The fact that the number of records exchanged (five in this example) is fixed/known means this is an ideal scenario. The consumer knows exactly when to finish/exit

If you run this: the receiver will print 1-5, asking you to exit.

To run, simply uncomment `f1()` in `main`, and run the program using `go run channels-range-close.go`.

i = 1
i = 2
i = 3
i = 4
i = 5
consumer finished. press ctrl+c to exit
producer finished
^C

# Goroutine Leak

This was an oversimplified case. Let’s make a small change by removing the `for` loop counter and converting it into an `infinite` loop. This is to simulate a scenario where the receiver wants to get all the values sent by the producer but doesn’t know the specifics — i.e., how many values will be sent (in real applications, this is often the case).

//consumer
go func() {
  for {
   i := <-c
   fmt.Println("i =", i)
  }
  fmt.Println("consumer finished. press ctrl+c to exit")
 }()

The output from the modified program is:

To run, simply uncomment `f2()` in `main`, and run the program using `go run channels-range-close.go`.

i = 1
i = 2
i = 3
i = 4
i = 5
producer finished

Notice the producer goroutine exited, but you didn’t see the `consumer finished. press ctrl+c to exit` message. Once the producer is done sending five integers and the consumer receives all of them, it’s just stuck there waiting for the next value, `i := <-c`, and won’t be able to return/exit.

In long-running programs, this results in a `goroutine leak`.

# `'range’` and '`close’` to the Rescue

This is where `range` and `close` can help:

*   `range` provides a way to iterate over values of a channel (just like you would for a slice)
*   `close` makes it possible to signal to consumers of a channel that nothing else will be sent on this channel

Let’s refactor the program. First, change the consumer to use `range`. Remove the `i := <-c` bit, and replace it with `for i := range c`.

go func() {
  for i := range c {
   fmt.Println("i =", i)
  }
  fmt.Println("consumer finished. press ctrl+c to exit")
}()

Update the producer goroutine to add `close(c)` outside the `for` loop. This will ensure the consumer goroutine gets the signal that there’s nothing more to come from the channel, and the `range` loop will terminate.

go func() {
   for i := 1; i <= 5; i++ {
    c <- i
    time.Sleep(1 \* time.Second)
 }
 close(c)
 fmt.Println("producer finished")
}()

If you run the program now, you should see this output:

To run, simply uncomment `f3()` in `main`, and run the program using `go run channels-range-close.go`.

i = 1
i = 2
i = 3
i = 4
i = 5
producer finished
consumer finished. press ctrl+c to exit

# Bonus

The consumer goroutine doesn’t have to coexist with the producer goroutine to receive the values — i.e. even if the producer goroutine finishes (and closes the channel), the consumer goroutine `range` loop will receive all the values. This is helpful when the consumer is processing the records sequentially.

We can simulate this scenario by using a combination of:

*   A buffered channel in the producer
*   Delaying the consumer goroutine by adding a `time.Sleep()`

In the producer, we create a buffered channel of capacity five `c := make(chan int, 5)`. This is to ensure the producer goroutine won’t block in the absence of a consumer:

c := make(chan int, 5)

//producer
go func() {
  for i := 1; i <= 5; i++ {
  c <- i
 }
 close(c)
 fmt.Println("producer finished")
}()

The consumer remains the same, except for `time.Sleep(5 * time.Second)`, which allows the producer goroutine to exit before the consumer can start off.

go func() {
	time.Sleep(5 \* time.Second)
        fmt.Println("consumer started")
	for i := range c {
	  fmt.Println("i =", i)
	} fmt.Println("consumer finished. press ctrl+c to exit")
}()

Here’s the output you should see:

To run, simply uncomment `f4()` in `main`, and run the program using `go run channels-range-close.go`.

producer finished
consumer started
i = 1
i = 2
i = 3
i = 4
i = 5
consumer finished. press ctrl+c to exit

The producer goroutine finished sending five records. The consumer woke up after a while, receiving and printing out all e five messages sent by the producer.
