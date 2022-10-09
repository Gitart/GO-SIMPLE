package main

import (
	"context"
	"fmt"
	"github.com/nuttech/bell/v2"
	"sort"
	"time"
)

type CustomStruct struct {
	name  string
	param int32
}

type Usr struct{
    Title  string 
    Name   string 
    Status string 
} 

type Sys struct{
    Title  string 
    Name   string 
    Status string 
} 

var(
	EvtWrk = "EwWork"
	EvtUsr = "EwUser"
	EvtSys = "EwSys" 
)


func init(){

  // add listener on event event_name
	bell.Listen(EvtWrk, func(message bell.Message) {
		customStruct := message.(CustomStruct)
		fmt.Println("WRK :", customStruct)
	})

  // add listener on event event_name
	bell.Listen(EvtUsr, func(message bell.Message) {
		usr := message.(Usr)
		fmt.Println("USR :", usr)
	})

  // add listener on event event_name
	bell.Listen(EvtSys, func(message bell.Message) {
		sys := message.(Sys)
		fmt.Println("SYS :", sys)
	})
}



func main(){
   GoRuns()   
   GoRun()
   GoStr()

	 // wait until the event completes its work
	 
}

// Tores
// GeForce
// 

func GoRun(){

	    time.Sleep(time.Second*7) 
     	bell.Ring(EvtWrk, CustomStruct{name: "Pause go test", param: 1332})
      // bell.Wait()
}

func GoRuns(){
  	 Dat:=CustomStruct{name: "testName", param: 12}
  	bell.Ring(EvtWrk, Dat)
     
    time.Sleep(time.Second * 2) 
  	bell.Ring(EvtWrk, CustomStruct{name: "Pause test", param: 1332})

    time.Sleep(time.Second * 1) 
  	bell.Ring(EvtWrk, CustomStruct{name: "Пауза задержки 5 сек", param: 1332})

    time.Sleep(time.Second * 2) 
  	bell.Ring(EvtUsr, Usr{Title: "Пользовательская  5 сек", Name: "Djon"})
}


func GoStr(){
	   f:=[]string{"Mandarin","Apelsin","Apple", "Orange", "Potates", "Anuc","Fructs"}
	   st:=[]string{"Warning","Info","System", "Critical", "Other", "Prim", "Cultural"}

   for _, stt:=range st{
	   
	   for _, t:=range f{
       bell.Ring(EvtSys, Sys{Title: t, Name: "sus", Status: stt})	   	
	   }
	 }

}




func Example() {
	// Use via global state

	event  := "event_name"
	event2 := "event_name_2"

	// add listener on event event_name
	bell.Listen(event, func(message bell.Message) {
		// we extend CustomStruct in message
		customStruct := message.(CustomStruct)
		fmt.Println(customStruct)
	})

	// add listener on event event_name_2
	bell.Listen(event2, func(message bell.Message) {})

	// get event list
	list := bell.List()

	// only for test
	sort.Strings(list)
	fmt.Println(list)

	// remove listeners on event_name_2
	bell.Remove(event2)

	// get event list again
	fmt.Println(bell.List())

	// check if exists event_name_2 event in storage
	fmt.Println(bell.Has(event2))

	// call event event_name
	_ = bell.Ring(event, CustomStruct{name: "testName", param: 12})

	// wait until the event completes its work
	bell.Wait()

	// Output:
	// [event_name event_name_2]
	// [event_name]
	// false
	// {testName 12}
}

func ExampleEvents() {
	// Use events object (without global state)

	eventName := "event_name"

	// make a new events store
	events := bell.New()

	// add listener on event
	events.Listen(eventName, func(msg bell.Message) { fmt.Println(msg) })

	// call event event_name
	_ = events.Ring(eventName, "Hello bell!")

	// wait until the event completes its work
	events.Wait()

	// Output:
	// Hello bell!
}

func Example_usingContext() {
	// Use bell with context

	// create a custom struct for pass a context
	type Custom struct {
		ctx   context.Context
		value interface{}
	}

	// add listener
	bell.Listen("event", func(message bell.Message) {
		for iterationsCount := 1; true; iterationsCount++ {
			select {
			case <-message.(*Custom).ctx.Done():
				return
			default:
				fmt.Printf("Iteration #%d\n", iterationsCount)
				time.Sleep(10 * time.Second)
			}
		}
	})

	// create a global context for all calls
	globalCtx, cancelGlobalCtx := context.WithCancel(context.Background())

	// create a children context for a call with timeout
	ringCtx, ringCancel := context.WithTimeout(globalCtx, time.Minute)
	defer ringCancel()

	_ = bell.Ring("event", &Custom{ringCtx, "value"})

	// wait a second for the handler to perform one iteration
	time.Sleep(time.Second)

	// interrupt all handlers
	cancelGlobalCtx()

	// Output:
	// Iteration #1
}
