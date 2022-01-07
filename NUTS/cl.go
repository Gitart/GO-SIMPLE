package main

import (
    "fmt"
    "log"
    "time"
    "flag"
    "github.com/nats-io/nats.go"
)

func main(){
    log.SetPrefix("INFO: ")

    useCh := flag.String("ch", "foo", "display colorized output")
    flag.Parse()
    fmt.Println("Listen chanel :", *useCh)
    
    go Clss(*useCh,500)
    go Clss("testing",500)
    go Clss("t.com",500)

     Clss("fooo",500)


}


func Clss(chname string, pause time.Duration){

// fmt.Println(nats.DefaultURL)    
// Connect to a server
nc, err := nats.Connect(nats.DefaultURL)

if err != nil {
   log.Println(err)
}

 defer nc.Close()

   nc.Subscribe(chname, func(m *nats.Msg) {
       log.Printf("Ch: %s  Mes: %s\n", chname, string(m.Data))
   })

for {
    time.Sleep(time.Millisecond * pause)
 }


}

