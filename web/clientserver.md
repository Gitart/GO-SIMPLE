## Simple client server example
This a quick tutorial on creating a bare bone client-server program with Golang. We will start with the server 
part first and then followed by client part. You can use this as a foundation to create your own client-server program.  



server.go

```golang
 package main

 import (
 	"fmt"
 	"log"
 	"net"
 )

 func handleConnection(c net.Conn) {

 	log.Printf("Client %v connected.", c.RemoteAddr())

 	// stuff to do... like read data from client, process it, write back to client
 	// see what you can do with (c net.Conn) at
 	// http://golang.org/pkg/net/#Conn

 	// buffer := make([]byte, 4096)

 	//for {
 	//		n, err := c.Read(buffer)
 	//		if err != nil || n == 0 {
 	//			c.Close()
 	//			break
 	//		}
 	//		n, err = c.Write(buffer[0:n])
 	//		if err != nil {
 	//			c.Close()
 	//			break
 	//		}
 	//	}
 	log.Printf("Connection from %v closed.", c.RemoteAddr())
 }

 func main() {
 	ln, err := net.Listen("tcp", ":6000")
 	if err != nil {
 		log.Fatal(err)
 	}

 	fmt.Println("Server up and listening on port 6000")

 	for {
 		conn, err := ln.Accept()
 		if err != nil {
 			log.Println(err)
 			continue
 		}
 		go handleConnection(conn)
 	}
 }
 ```
 
compile the server.go and run it as a background process :

```
./server &
```

then run the client on a separate machine(if possible) and connect to the server

### client.go

```golang
 package main

 import (
         "fmt"
         "net"
 )

 func main() {
         hostName := "example.com" // change this
         portNum := "6000"

         conn, err := net.Dial("tcp", hostName+":"+portNum)

         if err != nil {
                 fmt.Println(err)
                 return
         }

         fmt.Printf("Connection established between %s and localhost.\n", hostName)
         fmt.Printf("Remote Address : %s \n", conn.RemoteAddr().String())
         fmt.Printf("Local Address : %s \n", conn.LocalAddr().String())

 }
``` 

Sample output :

For client

```
./client-dial
Connection established between socketloop.com and localhost.
Remote Address : 162.243.5.230:6000
Local Address : 192.168.1.65:49774
```

For server
```
./server
Server up and listening on port 6000
2015/05/09 08:55:55 Client 14.192.213.197:8017 connected.
2015/05/09 08:55:55 Connection from 14.192.213.197:8017 closed.
This should be the basis of creating client-server programs with Golang. Hope you find this tutorial useful.
```

Check out this example as well on establishing secure(SSL/TLS) connection between client-server :
