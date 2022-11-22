package main

import "fmt"
import "io"
import "crypto/rand"                     // Crypto 

func main() {
uuid   := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	
	if n != len(uuid) || err != nil {
	   return 
	}

	uuid[8] = uuid[8]&^0xc0 | 0x80               // variant bits; see section 4.1.1
	uuid[6] = uuid[6]&^0xf0 | 0x40               // version 4 (pseudo-random); see section 4.1.3 	// return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
	 fmt.Printf("%X", uuid[0:8])         // return fmt.Sprintf("%x", uuid[0:6]), nil
}
