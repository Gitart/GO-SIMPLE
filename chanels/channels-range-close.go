package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//v_1()
	//v_2()
	//v_3()
	//v_4()
}

func f4() {
	c := make(chan int, 5)

	//producer
	go func() {
		for i := 1; i <= 5; i++ {
			c <- i
		}
		close(c)
		fmt.Println("producer finished")
	}()

	//consumer
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("consumer started")
		for i := range c {
			fmt.Println("i =", i)
		}
		fmt.Println("consumer finished. press ctrl+c to exit")
	}()

	e := make(chan os.Signal)
	signal.Notify(e, syscall.SIGINT, syscall.SIGTERM)
	<-e
}

func f3() {
	c := make(chan int)

	//producer
	go func() {
		for i := 1; i <= 5; i++ {
			c <- i
			time.Sleep(1 * time.Second)
		}
		close(c)
		fmt.Println("producer finished")

	}()

	//consumer
	go func() {
		for i := range c {
			fmt.Println("i =", i)
		}
		fmt.Println("consumer finished. press ctrl+c to exit")
	}()

	e := make(chan os.Signal)
	signal.Notify(e, syscall.SIGINT, syscall.SIGTERM)
	<-e
}

func f2() {
	c := make(chan int)

	//producer
	go func() {
		for i := 1; i <= 5; i++ {
			c <- i
			time.Sleep(1 * time.Second)
		}
		fmt.Println("producer finished")
	}()

	//consumer
	go func() {
		for {
			i := <-c
			fmt.Println("i =", i)
		}
		fmt.Println("consumer finished. press ctrl+c to exit")
	}()

	e := make(chan os.Signal)
	signal.Notify(e, syscall.SIGINT, syscall.SIGTERM)
	<-e
}

func f1() {
	c := make(chan int)

	//prodcuer
	go func() {
		for i := 1; i <= 5; i++ {
			c <- i
			time.Sleep(1 * time.Second)
		}
		fmt.Println("producer finished")
	}()

	//consumer
	go func() {
		for x := 1; x <= 5; x++ {
			i := <-c
			fmt.Println("i =", i)
		}
		fmt.Println("consumer finished. press ctrl+c to exit")
	}()

	e := make(chan os.Signal)
	signal.Notify(e, syscall.SIGINT, syscall.SIGTERM)
	<-e
}
