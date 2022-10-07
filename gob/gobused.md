## samples

```go
package main
 
import (
	"bytes"
	"encoding/gob"
	"fmt"
)
 
type User struct {
	FirstName string
	LastName  string
	Age       int
	Active    bool
}
 
func (u User) String() string {
	return fmt.Sprintf(`{"FirstName":%s,"LastName":%s,"Age":%d,"Active":%v }`,
		u.FirstName, u.LastName, u.Age, u.Active)
}
 
type SimpleUser struct {
	FirstName string
	LastName  string
}
 
func (u SimpleUser) String() string {
	return fmt.Sprintf(`{"FirstName":%s,"LastName":%s}`,
		u.FirstName, u.LastName)
}
 
func main() {
 
	var buff bytes.Buffer
 
	// Кодирование значения
	enc := gob.NewEncoder(&buff)
	user := User{
		"Radomir",
		"Sohlich",
		30,
		true,
	}
	enc.Encode(user)
	fmt.Printf("%X\n", buff.Bytes())
 
	// Декодирование значения
	out := User{}
	dec := gob.NewDecoder(&buff)
	dec.Decode(&out)
	fmt.Println(out.String())
 
	enc.Encode(user)
	out2 := SimpleUser{}
	dec.Decode(&out2)
	fmt.Println(out2.String())
 
}
```

### output
```go
40FF81030101045573657201FF82000104010946697273744E616D65010C0001084C6173744E616D65010C0001034167650104000106416374697665010200000019FF8201075261646F6D69720107536F686C696368013C010100
{"FirstName":Radomir,"LastName":Sohlich,"Age":30,"Active":true }
{"FirstName":Radomir,"LastName":Sohlich}
```
