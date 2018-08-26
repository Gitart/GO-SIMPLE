// client.go
package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

type Args struct {
	X, Y int
}

func main() {

	client, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
	  log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := &Args{7, 8}
	
	var reply int
	c   := jsonrpc.NewClient(client)
	err = c.Call("Calculator.Add", args, &reply)

	if err != nil {
		log.Fatal("arith error:", err)
	}
	
	fmt.Printf("Result: %d+%d=%d\n", args.X, args.Y, reply)


    var otvet string
    err = c.Call("Calculator.Test", "GGGGGGGGG11", &otvet)	

    
    fmt.Println("Dторая процедурка = ", otvet)

}
