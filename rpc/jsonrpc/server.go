// server.go
package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	X, Y int
}




type Calculator struct{}

// 
func (t *Calculator) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

// 
func (t *Calculator) Prn(Txt string, rep *string) error {
	 *rep = "OMNMMMMMMMMMMM----> " + Txt
	 return nil
}




func main() {
	cal    := new(Calculator)
	server := rpc.NewServer()
	server.Register(cal)
	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	listener, e := net.Listen("tcp", ":1234")
	
	if e != nil {
		log.Fatal("listen error:", e)
	}
	
	for {
		if conn, err := listener.Accept(); err != nil {
	  	   log.Fatal("accept error: " + err.Error())
		} else {
	  	  log.Printf("new connection established\n")
		  go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}

}




