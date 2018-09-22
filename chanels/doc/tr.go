package main

import (
	"os"
   "fmt"
   "runtime"
	"strconv"
   "sync"
)

var accountBalance = 0    // balance for shared bank account
var mutex = &sync.Mutex{} // mutual-exclusion lock

// critical-section code with explicit locking/unlocking
func updateBalance(amt int) { 
	mutex.Lock()          
	accountBalance += amt  // critical section
	mutex.Unlock()
}

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

   var wg sync.WaitGroup      

   // miser increments the balance
   wg.Add(1)           // increment WaitGroup counter
   go func() {
      defer wg.Done()  // invoke Done on the WaitGroup when finished
      for i := 0; i < iterations ; i++ {
			updateBalance(1)
         runtime.Gosched() // yield to another goroutine
      }
   }()
   
	// spendthrift decrements the balance
   wg.Add(1)           // increment WaitGroup counter
   go func() {
      defer wg.Done() 
      for i := 0; i < iterations; i++ {
			updateBalance(-1)
         runtime.Gosched()  // be nice--yield 
      }
   }()
	
   wg.Wait()  // await completion of miser and spendthrift goroutines
	fmt.Println("Final balance: ", accountBalance)	// confirm final balance is zero
}
