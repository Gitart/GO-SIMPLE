// https://play.golang.org/p/aEOEEQrIL1

package main

import (
	"crypto/sha1"
	"fmt"
	"encoding/hex"
	"strconv"
	"io"     
	"crypto/rand"
)

func main() {

//t:=fmt.Sprintf("% x", sha1.Sum([]byte("erewrewrew-ewrtretretret-ertretr-ertgrety")))


//fmt.Println(t,"---", t[:8], hex.EncodeToString(t[:8]))
//y,_:=strconv.ParseInt(hex.EncodeToString(t[:8]), 16, 64)
//fmt.Println("Answer = ", y,"=", strconv.FormatInt(y,10))
fmt.Println(GenGUIDSHA())


}


// Генерация GUID
func GENGUID() string {
	uuid   := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	
	if n != len(uuid) || err != nil {
	   return ""
	}

	uuid[8] = uuid[8]&^0xc0 | 0x80               
	uuid[6] = uuid[6]&^0xf0 | 0x40               
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])//, nil
	
}


// Генерация HSA1
// fmt.Println(GenGUIDSHA())
func GenGUIDSHA() string {
    g:=GENGUID()
    fmt.Println(g)
    t:=sha1.Sum([]byte(g))
    y,_:=strconv.ParseInt(hex.EncodeToString(t[:8]), 16, 64)
    r:=strconv.FormatInt(y,10)
    return r
}

