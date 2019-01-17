/* FTP Server
 */
package main

import (
    "fmt"
    "net"
    "os"
    "time"
    "strings"
)

const (
    DIR  = "DIR"
    CD   = "CD"
    PWD  = "PWD"
    SND  = "SND"
)

func main() {
    fmt.Println("\nIM SERVER\n\n")
    
    tcpAddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:1202")
    checkError(err)

    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)
    

     go Coversage()
     go Seder("I`M STERTED")
    

    for {
        conn, err := listener.Accept()
        if err != nil {
           continue
        }
        go handleClient(conn)
    }
}


// Послание клиенту
func Coversage(){
      Seder("Привет  давай поговорим")
      time.Sleep(time.Second*10)
      Seder("Расскажи о себе немного")
     time.Sleep(time.Second*10)
      Seder("Я послал тебе команду - comm")
       
}    

// Послание
func Seder(Txt string){
     conn, err := net.Dial("tcp", "127.0.0.1:1203")
     

    if err != nil {
       return
    }else{
           conn.Write([]byte("SERVER :" + Txt))
    }
    
    defer conn.Close()  

  
    // time.Sleep(time.Second*2)
}


func handleClient(conn net.Conn) {
    defer conn.Close()

     Seder("Привет клиент я слушаю твои комманды !")
     go Resp()

    var buf [512]byte
    n, _ := conn.Read(buf[0:])
    s := string(buf[0:n])
    fmt.Println(s)
    return


    for {
        n, err := conn.Read(buf[0:])
        if err != nil {
           conn.Close()
           return
        }

        s := string(buf[0:n])
        
        if strings.Contains(s,"Test"){
          fmt.Println("Спасибо за команду ТЕСТ")
          Seder("Test - услышана")
        }


      
        // decode request
        if s[0:2] == CD {
            chdir(conn, s[3:])
        } else if s[0:3] == DIR {
            dirList(conn)
        } else if s[0:3] == SND {
            snd(conn,s)
            
        } else if s[0:3] == PWD {
            pwd(conn)
        }
    }
}


func Resp(){

       Seder("SERVER: Ответ моему другу клиенту!")    
}



// sender
func snd(conn net.Conn, s string) {
     conn.Write([]byte(s))
    
}



func chdir(conn net.Conn, s string) {
    if os.Chdir(s) == nil {
       conn.Write([]byte("OK"))
    } else {
       conn.Write([]byte("ERROR"))
    }
}

func pwd(conn net.Conn) {
    s, err := os.Getwd()
    if err != nil {
       conn.Write([]byte(""))
       return
    }
    conn.Write([]byte(s))
}


func dirList(conn net.Conn) {
    defer conn.Write([]byte("\r\n"))

    dir, err := os.Open(".")
    if err != nil {
       return
    }

    names, err := dir.Readdirnames(-1)
    if err != nil {
       return
    }
    
    for _, nm := range names {
        conn.Write([]byte(nm + "\r\n"))
    }
}

func checkError(err error) {
    if err != nil {
       fmt.Println("Fatal error ", err.Error())
       os.Exit(1)
    }
}
