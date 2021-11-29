package main

import "time"
import "fmt"

// ***********************************************
// Basic main
// ***********************************************
func main(){

  ch:= make(chan string)
  ct:= make(chan string)
  ss:= make(chan string)

  go Worker(ch, ct, ss)

  // Не выполнится если применить defer
  func(){
     ch <- "Start - > Finish data ..."
  }()

  go func(){
     ch <- "Пример нового старта ..."
  }()
 
  go func(){
     ch <- "Пример нового старта ..."
  }()

  // Cahnel #1
  go func(){
    for i := 1; i<=10; i++ {
      ct <- fmt.Sprintf("S: %v",i)
      time.Sleep(time.Millisecond * 40)
     }
  }()
    
  // Cahnel #2
  go func(){
    for i := 1; i<=10; i++ {
      ss <- fmt.Sprintf("S: %v",i)
      time.Sleep(time.Millisecond * 10)
    }
  }()
   
  // Cahnel #3
  for i := 1; i<=5; i++ {
      ch <- fmt.Sprintf("D: %v",i)
      time.Sleep(time.Millisecond * 110)
  }

}


// ***********************************************
// Go worker
// ***********************************************
func Worker(c chan string, t chan string, ss chan string ){

    for {
      select {
        case res := <-c:  fmt.Println("ONE : ", res)
        case res := <-t:  fmt.Println("TWO : ", res)
        case res := <-ss: fmt.Println("TRI : ", res) 
    }
  }
}
