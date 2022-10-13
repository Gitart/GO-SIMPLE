// ***************************************
// Работа с каналами и очередями
// ***************************************
package main

import "fmt"
import "time"

var idx int

// ***************************************
// Запись в канала
// ***************************************
func mch(ch chan int, inn int){
    ch<-inn
}

// ***************************************
// Rider chanel
// ***************************************
func reads(ch chan int){
     for num := range ch {
           fmt.Printf("F(%d): \t%d\n", idx, num)
          //idx=idx+num
          idx++
     }
}

// ***************************************
// Routine in routine
// ***************************************
func posul(ch chan int){
    // Добавление в очередь
    for i:=1; i<=10; i++ {   
        go mch(ch,i)
    }     	   
}

// ***************************************
// Main 
// Можно использовать Bell
// Context
// ***************************************
func main() {
	idx=0
     c := make(chan int,1000)

     go posul(c) 
     // go reads(c)  

    // Добавление в очередь
    for i:=1; i<=10; i++ {   
        go mch(c,i)
    }     
   
   // time.Sleep(time.Millisecond*100)
   for i:=1; i<=10; i++ {   
   	fmt.Println("----------",i)
       go mch(c,i)
   }
   
   time.Sleep(time.Millisecond*100)
   fmt.Println("Count: ", len(c))
   

   time.Sleep(time.Second*1)
    fmt.Println("Before read : ",len(c))
   go reads(c)  

   time.Sleep(time.Second*1)
   fmt.Println("Count: ",len(c))
   
    // time.Sleep(time.Millisecond*10000)
    // Сколько элементов будет показано вычитан за это время из очереди
    time.Sleep(time.Second*1)
    fmt.Println("All elements-------------------- :", idx)      

    // time.Sleep(time.Second*5)
    // fmt.Println("All elements part two :", idx)      
       
}


// Плохо работает!!!!!!!!!
func main2() {
	

	c := make(chan int)
	
	go func(chan int){
	    c<-1
	}(c)

	go func(chan int){
	    c<-2
	}(c)
	
	go func(chan int){
	    c<-256
	}(c)
	
	// close(c)
	
	 time.Sleep(time.Second)
	o:=len(c)
	cp:=cap(c)

    // for _, f:=range c{
    // 	fmt.Println(f)
    // } 

	fmt.Println(o, cp)
	// x:=<-c
	// y:=<-c
	// l:=<-c
	
	// fmt.Println(x,y,l,o)
}
