package main

import (
    "fmt"
    "net"
    "os"
    // "strings"
    "encoding/json"
)


// Main procedure
func main() {
     ServerStart()
}


func ServerStart(){
    fmt.Println("Server is started on the 10001 ports")
    fmt.Println("Lets prepare a address at any address at port 10001")


    /* Lets prepare a address at any address at port 10001*/
    ServerAddr, err := net.ResolveUDPAddr("udp",":10001")
    CheckError(err)

    /* Now listen at selected port */
    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    CheckError(err)
    defer ServerConn.Close()

    buf := make([]byte, 1024)


     for {
        n, addr, err := ServerConn.ReadFromUDP(buf)
        fmt.Println("Count:",n, string(buf))
         
         ss:=string(buf[0:n])
        // sf:=strings.Split(ss, " ")
        // ss:=buf[0:n]
      
        if err != nil {
           fmt.Println("Error: ",err)
        }

        fmt.Println(addr,ss)   
     }
}


func ServerStart_old(){
    fmt.Println("Server is started on the 10001 ports")
    fmt.Println("Lets prepare a address at any address at port 10001")


    /* Lets prepare a address at any address at port 10001*/
    ServerAddr, err := net.ResolveUDPAddr("udp",":10001")
    CheckError(err)

    /* Now listen at selected port */
    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    CheckError(err)
    defer ServerConn.Close()

    // buf := make([]byte, 1024)

    // buf := make([]byte, 5000)
    var buf []byte
    buf=[]byte(`aaaaaa`)


     for {
        n, addr, err := ServerConn.ReadFromUDP(buf)
        fmt.Println("Count:",n, string(buf))
         
        // ss:=string(buf[0:n])
        // sf:=strings.Split(ss, " ")
        ss:=buf[0:n]
      
        if err != nil {
           fmt.Println("Error: ",err)
        }
    

         m := make(map[string]string)
         json.Unmarshal(ss, &m)

        fmt.Println(addr, ">", m["Msg"])   

        // for i := 0; i < n; i++ {
        // // fmt.Println("Received ",string(buf[0:n]), " Получение от адреса ", addr)
        // cc=cc+string(buf[0:n])
        // fmt.Println(">", cc)   

        // }
      
     }
}



/* 
  A Simple function to verify error 

*/
func CheckError(err error) {
    if err  != nil {
       fmt.Println("Error: " , err)
       os.Exit(0)
    }
}
