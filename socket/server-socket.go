package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"log"
	"net/http"
	"time"
)



func main() {
	
	http.Handle("/echo",         websocket.Handler(echoHandler))
    http.Handle("/eh",           websocket.Handler(echoHand))
    


    // Имитация отправки в канал 
    // http.HandleFunc("/chan/",  echoHandChans)

    http.HandleFunc("/in/",      echoHandChans_inp)
    http.HandleFunc("/out/",     echoHandChans_read)




    // Здесь ловим канал в скоетах
    http.Handle("/ehs",        websocket.Handler(echoHandChan))



    fmt.Println("Server start 4444 port...")
	err := http.ListenAndServe(":4444", nil)

	if err != nil {
	   log.Println("ListenAndServe: " + err.Error())
	}
}


type CCC struct{
	Ch chan string 
}



func work(messages chan<- string) {
    messages <- "golangcode.com"
}

var Mss = make(chan string, 10)

// var Mc chan string


var Ii int=0


// Запись в канал ообщения
func echoHandChans_inp(w http.ResponseWriter, r *http.Request){
	pt        := r.URL.Path[len("/in/"):]
   
   Ii++
   I:=fmt.Sprintf("%v",Ii)

   go   func (){
         Mss<-I + " " + pt
    }()

    log.Println(I+" "+pt)
    
}


// Чтение из канала
func echoHandChans_read(w http.ResponseWriter, r *http.Request){
   for{
   select {
     case res:=<-Mss : 
     	fmt.Println(res)
     // default : 
     // 	fmt.Println("No")
   }
}
}

// Отправка в канал тествового сообщения
func echoHandChans2(w http.ResponseWriter, r *http.Request){
	// Mc:=make(chan string, 2)
    // go Tochan(Mc, "Test lkz")
    // go Tochan(Mc, "Test lkz2")
    // log.Println("Ok отослано ")
    // time.Sleep(1 * time.Second)	
    // // c:=<-Mc
    // log.Println("Получено ",c)
    // // c=<-Mc
    // log.Println("Получено 3",c)
}


func Tochan(c chan string, S string)   {
     c<-S
}


// func Mcc(Chs string chan){
    
//        Mc<-"Test chan for ez"  
    
	
// }












// Пример создания нововго соединения
// Новости получают из этого сообщения
func echoHandChan(ws *websocket.Conn) {
      fmt.Println("Channel Open ")   
    	
// select {
//      case res:=<-Mc : 
//      	ws.Write([]byte(res))
//      default : 
//      	ws.Write([]byte("break......"))
// }


 for{
   select {
     case res:=<-Mss : 
     	
     	 ws.Write([]byte(res)	)
     // default : 
     // 	fmt.Println("No")
   }
}

   
      	
       

}



// Пример создания нововго соединения
// Новости получают из этого сообщения
func echoHandChan1(ws *websocket.Conn) {
       
      log.Println("Cлушет очоередь ")

      // for {
      // 	   res, ok := <-Mc 

      //    if ok == false { 
      //       fmt.Println("Channel Close ", ok) 
      //       break
      //   } 
           
      // 	   ws.Write([]byte(res))
      // 	   fmt.Println("Channel Open ", res, ok) 
      // }

      ws.Write([]byte("создания"))

}




// Пример создания нововго соединения
// Новости получают из этого сообщения
func echoHand(ws *websocket.Conn) {

for i:=1;i<10; i++{
    fmt.Println("Ok ",i)
    t:=time.Now().Format("15:04:05")
	
    ss:=fmt.Sprintf("%v*Запись № %v ",i,t) 
    m:=[]byte(ss)
	_, errs := ws.Write(m)

	if errs != nil {
	   log.Println("Ошибка:",errs)
	   break
	}
	
	log.Println(string(m))
	time.Sleep(1 * time.Second)	
}



ws.Write([]byte("Остановка на перекур...."))
time.Sleep(11 * time.Second)	




for i:=1;i<15; i++{
    fmt.Println("Ok ",i)
    t:=time.Now().Format("15:04:05")
	
    ss:=fmt.Sprintf("%v* Дополнительная запись № %v ",i,t) 
    m:=[]byte(ss)
	_, errs := ws.Write(m)

	if errs != nil {
	   log.Println("Ошибка:",errs)
	   break
	}
	
	log.Println(string(m))
	time.Sleep(100 * time.Millisecond)	
}




}



// Новости можно еще больше переключать
// Пример создания нововго соединения
// Новости получают из этого сообщения
func echoHandler(ws *websocket.Conn) {

    fmt.Println("Ok")

	msg    := make([]byte, 512)
	n, err := ws.Read(msg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Receive: %s\n", msg[:n])

	m, err := ws.Write(msg[:n])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Send: %s\n", msg[:m])
}
