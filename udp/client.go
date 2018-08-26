

// https://play.golang.org/p/KbLVd0VySe


package main


import (
    "fmt"
    "net"
    // "time"
    "strconv"
    "log"
    "bufio"
    "os"
)


func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}


func main() {
     // В цикле передача
     // forwrite()   
     
     formanual()
    // mains()

    // !!!!!!!!!!!!!!!
    // В юниксу команда ss  показ порты

}



// ************************************************************************************
//  В одной строке черз пробел вводятся слова которые и попадают в разные переменные
//  пробел и является разделителем на разные перменнвые!!!!!!!!!11
// ************************************************************************************
func formanual(){
    
    
      //c := make(chan string)
      // var l:=""

     // i:=0 

     // fmt.Scan(&c)

     
// a:=func(Tx string) string{
//      l:=l+Tx
//      fmt.Println("dddd ", l)
//      //cc<-l
//      // fmt.Println(cc)
//      return l
// } 




 var c string 
 var o string

 i:=0

// Непывный ввод с подсчетм количества слов 
// когда 4 слова выход 
// Интересная функция scan - работает интересно
// В бесконесном цикле крутиться но берет по одному слову
// или несколько если указать через запятую 
// нет возможности получения количества слов введенных 
// нужно опредлять каким угодно способом 
// fmt.Scan(&o, &v, &b) - напримре для трех вводимых слов
 

for {
     i++    
     fmt.Scan(&o)
     c=c+" "+o
     fmt.Println("...", c)     
     
    // После 4 слов выход
    if i==4 {
       break
    }
}
     
///fmt.Println("____",c)     

     
Sendtoserver(c)                     




     // for{
     //      d,_:= fmt.Scan(&c)
          

     //     i++ 
           
     //     fmt.Println(i,d)  
     //      // o <- c  

     //    }


      // fmt.Println(o)

          // Sendtoserver(c)                     
}










func mains() {
    //reading an integer
    var age int
    fmt.Println("What is your age?")
    _, err:=fmt.Scan(&age)

    if err!=nil{
        return
    }

    //reading a string
    reader := bufio.NewReader(os.Stdin)
    var nam string
    fmt.Println("What is your name?")
    
     nam,_=reader.ReadString('\n')

    fmt.Println("Your name is ", nam, " and you are age ", age)
}

// 
func formanuals(){
     c:=""
     fmt.Println("Ввод для передачи на сервер.")
     for  {
                
           fmt.Scanln(&c)    
           // fmt.Println(c)
           Sendtoserver(c)                     
     }

}






// ********************************************************
// Send to server for manual type
// ********************************************************
func Sendtoserver(Snd string){

    ServerAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:10001"); CheckError(err)
    LocalAddr, err := net.ResolveUDPAddr("udp","127.0.0.1:0");     CheckError(err)
    Conn, err      := net.DialUDP("udp", LocalAddr, ServerAddr);   CheckError(err)
    defer Conn.Close()

     // msg:=[]byte(`{"Msg": "`+Snd+`"}`)
    
    // buf := []byte(Snd)
    msg:=[]byte(Snd)

   // msg=[]byte("ssssssssssssssssssssssss ssssssssssss")

    _,errs := Conn.Write(msg)

    if errs != nil {
       fmt.Println(Snd, err)
    }
}



// ********************************************************
// Send to server for manual type
// ********************************************************
func SendtoserverJson(Snd string){

     ServerAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:10001"); CheckError(err)
     LocalAddr, err := net.ResolveUDPAddr("udp","127.0.0.1:0");     CheckError(err)
     Conn, err      := net.DialUDP("udp", LocalAddr, ServerAddr);   CheckError(err)
     defer Conn.Close()
    
    // input := []byte(`{"key1": "value1", "key2": "value2"}`)
    // var m  map[string]string
    // json.Unmarshal(input, &m)
     

     buf := []byte(Snd)
    
    _,errs := Conn.Write(buf)

    if errs != nil {
       fmt.Println(Snd, err)
    }
}








func forwrite(){
     ServerAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:10001")
     CheckError(err)

     LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
     CheckError(err)

     Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
     CheckError(err)

     defer Conn.Close()
     i := 0

    for {
        msg := strconv.Itoa(i)
        log.Println("Send->", msg)
    
        i++
        buf := []byte("Testing -> " + msg)
    
        _,err := Conn.Write(buf)
    
        if err != nil {
           fmt.Println(msg, err)
        }
    

        // time.Second * 1 самая маленькая
        // time.Millisecond * 1 средняя скорость
        // Microsecund * 1 смая большая скорость
        // time.Sleep(time.Microsecond * 1)
        // time.Sleep(time.Microsecond * 1)
    }
}

