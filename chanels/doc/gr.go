package main

import (
	"os"
   "fmt"
   "runtime"
	"strconv"
	"sync"
)

type bankOp struct { // bank operation: deposit or withdraw
   howMuch int       // amount
   confirm chan int  // confirmation channel
}

var accountBalance = 0          // shared account
var bankRequests chan *bankOp   // channel to banker

func updateBalance(amt int) int {
   update := &bankOp{howMuch: amt, confirm: make(chan int)} 
   bankRequests <- update                                           
   newBalance := <-update.confirm
	return newBalance
}

// For now a no-op, but could save balance with a timestamp.
func logBalance(current int) { } 

func reportAndExit(msg string) {
   fmt.Println(msg)
	os.Exit(-1) // all 1s in binary
}

func main() {
	if len(os.Args) < 2 {
		reportAndExit("\nUsage: go ms1.go <number of updates per thread>")
	}
	iterations, err := strconv.Atoi(os.Args[1])
	if err != nil {
		reportAndExit("Bad command-line argument: " + os.Args[1]);
	}
   	
	bankRequests = make(chan *bankOp, 8) // 8 is channel buffer size

	var wg sync.WaitGroup     
   // The banker: handles all requests for deposits and withdrawals through a channel.
   go func() {
      for {   
         /* The select construct is non-blocking:
            -- if there's something to read from a channel, do so
            -- otherwise, fall through to the next case, if any */
         select {
         case request := <-bankRequests:
            accountBalance += request.howMuch   // update account
            request.confirm <- accountBalance   // confirm with current balance
         }
      }
   }()

   // miser increments the balance
   wg.Add(1)           // increment WaitGroup counter
   go func() {
      defer wg.Done()  // invoke Done on the WaitGroup when finished
      for i := 0; i < iterations ; i++ {
			newBalance := updateBalance(1)
			logBalance(newBalance)
         runtime.Gosched()  // yield to another goroutine
      }
   }()
   
	// spendthrift decrements the balance
   wg.Add(1)           // increment WaitGroup counter
   go func() {
      defer wg.Done() 
      for i := 0; i < iterations; i++ {
			newBalance := updateBalance(-1)
			logBalance(newBalance)
         runtime.Gosched()  // be nice--yield 
      }
   }()
	
   wg.Wait()  // await completion of miser and spendthrift
	fmt.Println("Final balance: ", accountBalance)	// confirm the balance is zero
}
