## Crypto/tls.Listen and NewListener functions example

#### package crypto/tls
Listen creates a TLS listener accepting connections on the given network address using net.Listen.   
The configuration config must be non-nil and must have at least one certificate.    
NewListener creates a Listener which accepts connections from an inner Listener and wraps each connection with Server.    
The configuration config must be non-nil and must have at least one certificate.     
For this example to work, first you need to have pem and key files.    
On Linux/Unix machines, you can generate the files with openssl   

```
openssl req -new -nodes -x509 -out server.pem -keyout server.key -days 365
```

Golang Listen() and NewListener() functions usage example


```golang
 package main

 import (
   "fmt"
   "net"
   "os"
   "crypto/tls"
   "crypto/rand"
 )

 func main() {

   certificate, err := tls.LoadX509KeyPair("server.pem", "server.key")

   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }


   // For ClientAuth : tls.RequireAnyClientCert. See http://golang.org/pkg/crypto/tls/#ClientAuthType
   config := tls.Config{Certificates : []tls.Certificate{certificate}, ClientAuth: tls.RequireAnyClientCert}
   config.Rand = rand.Reader


   var newnetlistener, netlistener net.Listener

   // The network net must be a stream-oriented network: "tcp", "tcp4", "tcp6", "unix" or "unixpacket".
   // see http://golang.org/pkg/net/#Listen
   var network = "tcp"

   var laddr = "0.0.0.0:8000" // laddr = address to listen

   netlistener, err = tls.Listen(network, laddr, &config)

   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   fmt.Println("[Server Status] : Listening")
   newnetlistener = tls.NewListener(netlistener, &config) // New Listener
   fmt.Println("[Server Status] : New Listener Listening")

   for {
      conn, err := netlistener.Accept()
      newconn, err := newnetlistener.Accept()

      if err != nil {
         fmt.Println(err)
      }

      fmt.Printf("[Server Status] : Accepted connection from %s", conn.RemoteAddr())
      fmt.Printf("[Server Status] : Accepted connection from %s", newconn.RemoteAddr())
    }

 }
 ```
 
Output :
```
Server Status] : Listening
[Server Status] : New Listener Listening
