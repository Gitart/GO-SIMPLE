## UDP client server read write example
In this tutorial, we will learn how to build a bare bone client-server UDP program with Golang. 
This UDP example is adapted from previous tutorial on how to create TCP client-server example in Golang.

Here is an example of how a UDP server program looks like in Golang. The server will first setup a 
listener and then starts listening for incoming UDP client on localhost. You can change the host name and 
port number to suit your configuration.

udpserver.go

```golang
 package main

 import (
         "fmt"
         "log"
         "net"
 )

 func handleUDPConnection(conn *net.UDPConn) {

         // here is where you want to do stuff like read or write to client

         buffer := make([]byte, 1024)

         n, addr, err := conn.ReadFromUDP(buffer)

         fmt.Println("UDP client : ", addr)
         fmt.Println("Received from UDP client :  ", string(buffer[:n]))

         if err != nil {
                 log.Fatal(err)
         }

         // NOTE : Need to specify client address in WriteToUDP() function
         //        otherwise, you will get this error message
         //        write udp : write: destination address required if you use Write() function instead of WriteToUDP()

         // write message back to client
         message := []byte("Hello UDP client!")
         _, err = conn.WriteToUDP(message, addr)

         if err != nil {
                 log.Println(err)
         }

 }

 func main() {
         hostName := "localhost"
         portNum := "6000"
         service := hostName + ":" + portNum

         udpAddr, err := net.ResolveUDPAddr("udp4", service)

         if err != nil {
                 log.Fatal(err)
         }

         // setup listener for incoming UDP connection
         ln, err := net.ListenUDP("udp", udpAddr)

         if err != nil {
                 log.Fatal(err)
         }

         fmt.Println("UDP server up and listening on port 6000")

         defer ln.Close()

         for {
                 // wait for UDP client to connect
                 handleUDPConnection(ln)
         }

 }
 ```
 
### Build this udpserver.go and run it on the background :

```
./udpserver &
```

then run the client on a separate machine(if possible) and connect to the server

### udpclient.go

```golang
 package main

 import (
         "log"
         "net"
         "fmt"
 )

 func main() {
         hostName := "localhost"
         portNum := "6000"

         service := hostName + ":" + portNum

         RemoteAddr, err := net.ResolveUDPAddr("udp", service)

         //LocalAddr := nil
         // see https://golang.org/pkg/net/#DialUDP

         conn, err := net.DialUDP("udp", nil, RemoteAddr)

         // note : you can use net.ResolveUDPAddr for LocalAddr as well
         //        for this tutorial simplicity sake, we will just use nil

         if err != nil {
                 log.Fatal(err)
         }

         log.Printf("Established connection to %s \n", service)
         log.Printf("Remote UDP address : %s \n", conn.RemoteAddr().String())
         log.Printf("Local UDP client address : %s \n", conn.LocalAddr().String())

         defer conn.Close()

         // write a message to server
         message := []byte("Hello UDP server!")

         _, err = conn.Write(message)

         if err != nil {
                 log.Println(err)
         }

         // receive message from server
         buffer := make([]byte, 1024)
         n, addr, err := conn.ReadFromUDP(buffer)

         fmt.Println("UDP Server : ", addr)
         fmt.Println("Received from UDP server : ", string(buffer[:n]))

 }
 ```
 
Sample output for udpserver.go :

```
UDP server up and listening on port 6000
UDP client : 127.0.0.1:63937
Received from UDP client : Hello UDP server!
```

Sample output for udpclient.go :

```
2015/11/24 11:14:56 Established connection to localhost:6000
2015/11/24 11:14:56 Remote UDP address : 127.0.0.1:6000
2015/11/24 11:14:56 Local UDP client address : 127.0.0.1:63937
UDP Server : 127.0.0.1:6000
Received from UDP server : Hello UDP client!
```
