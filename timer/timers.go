// Module timers
package main

import (
	"time"
  	"log"
)


func Timers(){
     go Timer_seconds()
     go Timer_minutes()
     go Timer_hours()
}

// Bacgraund Process
func Timer_seconds(){
	for {
		 time.Sleep(12 * time.Second) 
    	      log.Println("Too second")
    }
}

// Bacgraund Process
func Timer_minutes(){

	for {
		 time.Sleep(10 * time.Minute) 
    	 log.Println("One minute")
    }
}

// Bacgraund Process
func Timer_hours(){
	for {
		 time.Sleep(1 * time.Hour) 
    	 log.Println("One minute")
    }
}
