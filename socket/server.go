package main
import (
"fmt"
"net"
"os"
)
const (
SVR_HOST = "localhost"
SVR_PORT = "9982"
SVR_TYPE = "tcp"
)

func main() {
     fmt.Println("server is running")
     
     svr, err := net.Listen(SVR_TYPE, SVR_HOST+":"+SVR_PORT)
     if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
     }
     
     defer svr.Close()
     
     fmt.Println("Listening on " + SVR_HOST + ":" + SVR_PORT)
     fmt.Println("Waiting client…")

for {
     conn, err := svr.Accept()
	 if err != nil {
		 fmt.Println("Error accepting: ", err.Error())
		 os.Exit(1)
	 }

    fmt.Println("client connected")
    go handleClient(conn)
    }

}


func handleClient(conn net.Conn) {
	buff := make([]byte, 1024)
	msgLen, err := conn.Read(buff)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ",string(buff[:msgLen]))
	conn.Close()
}
