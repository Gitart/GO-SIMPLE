/* FTPClient
 */
package main

import (
    "fmt"
    "net"
    // "os"
    // "bufio"
    "strings"
    // "bytes"
    "time"
    "math/rand"
)

func main() {
    fmt.Println("\nIM CLIENT\n\n")
    tcpAddr, _  := net.ResolveTCPAddr("tcp", "127.0.0.1:1203")
    listener, _ := net.ListenTCP("tcp", tcpAddr)
    
    go Seder("IM STARTED AND LISTING YOU")

    go Seder(Answer())


    for {
        conn, err := listener.Accept()
        if err != nil {
           continue
        }

        go handleClient(conn)
    }
  
}

// Слушает ответ от клиента
func handleClient(conn net.Conn) {
    defer conn.Close()

    var buf [512]byte
    n, _ := conn.Read(buf[0:])
    s := string(buf[0:n])
    fmt.Println(">>", s)
      

 
     // reader := bufio.NewReader(os.Stdin) 


    for {
        //  line, err := reader.ReadString('\n')
        // // lose trailing whitespace
        // line = strings.TrimRight(line, " \t\r\n")
        // if err != nil {
        //     break
        // } 

        // // split into command + arg
        // strs := strings.SplitN(line, " ", 2)

        // decode user request
        
        if strings.Contains(s,"comm"){
           fmt.Println("Я услышал  твою комманду Comm")
           Seder("Coom - услышана")
        }

        if strings.Contains(s,"test"){
           fmt.Println("Отсылаю тебе комманду для теста")
           Seder("Test - услышана")
        }
    

        // fmt.Println("----", strs[0])
        n, err := conn.Read(buf[0:])
        if err != nil {
           conn.Close()
           return
        }




        s := string(buf[0:n])
        fmt.Println("EEE->",s) 
       
    }

}

func Sms() {
    for i := 0; i < 10; i++ {
      Seder("Hi all")        
      Seder("Hi friend?")        
      Seder("Что нового ?")        
      time.Sleep(time.Second*2)
   }
}


// Послание серверу от клиента на адрес сервера
func Seder(Txt string){
      conn, err := net.Dial("tcp", "127.0.0.1:1202")
      
      if err != nil {
         fmt.Println("INFO: Waiting Server ")
         return
       }else{
         conn.Write([]byte("CLIENT : "+Txt))
       }
     defer conn.Close()
}

func Sendr(conn net.Conn, Txt string){
    conn.Write([]byte(Txt))
    fmt.Println("Sender.....")
    
}

// ***********************************************
// Можно использовать для совета дня 
// или для ответа на вопросы
//***********************************************

func Answer() string{
    // Seeding with the same value results in the same random sequence each run.
    // For different numbers, seed with a different value, such as
    // time.Now().UnixNano(), which yields a constantly-changing number.
    
    
    rand.Seed(time.Now().UnixNano())

    answers := []string{
        "It is certain", 
        "It is decidedly so",
        "Without a doubt",
        "Yes definitely",
        "You may rely on it",
        "As I see it yes",
        "Most likely",
        "Outlook good",
        "Yes",
        "Signs point to yes",
        "Reply hazy try again",
        "Ask again later",
        "Better not tell you now",
        "Cannot predict now",
        "Concentrate and ask again",
        "Don't count on it",
        "My reply is no",
        "My sources say no",
        "Outlook not so good",
        "Very doubtful",
    }

    r:=answers[rand.Intn(len(answers))]

    // fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])
    return r
}
