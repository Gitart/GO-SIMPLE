

## TCP

```golang

// Check 
// telnet 127.0.0.1 5000
// или Putty 
// https://www.youtube.com/watch?v=lgh6zic15EA


package main 


import (
	"fmt"
	"bufio"
	"net"
)


func main() {

	listener,_:=net.Listen("tcp",":5000")

	for {
		conn, err:=listener.Accept()

		if err!=nil{
		   fmt.Println("Error")
		   conn.Close()
		   continue
		}

        fmt.Println("Connected...")
        bufReader:=bufio.NewReader(conn)
        fmt.Println("Start reading")

      // Для запуска с нескольких теминалов 
      go func (conn net.Conn){
            for{

         	rbyte,err:=bufReader.ReadByte()
         	
         	if err!=nil{
         	   fmt.Println("Error read")
         	   break
         	}

              fmt.Println(string(rbyte))
         	
         }

       }(conn)

	}
}
```

