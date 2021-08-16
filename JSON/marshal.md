## Переопределение стандартных методов 
* UnmarshalJSON()  вызывается автоматически
* MarshalJSON() в ручном варианте


## Sample
```go
package main

import (
	"encoding/json"
	"log"
	"strings"
)

type Character struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Job         string `json:"job,omitempty"`
	YearOfBirth int    `json:"year_of_birth,omitempty"`
}

func (c *Character) UnmarshalJSON(b []byte) error {
	log.Println("Calling UnmarshalJSON")
	type C Character
	var v C
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	*c = Character(v)
	if c.Job == "" {
		c.Job = "unknown"
	} 
	return nil
}


func (c *Character)  MarshalJSON() (r []byte, er error) {
	log.Println("Calling ")
	return []byte(`dd11`),nil
}


func main() {
	r := strings.NewReader(`{"name":"Lavinia","surname":"Whateley","year_of_birth":1878}`)
	d := json.NewDecoder(r)
	var c Character
	if err := d.Decode(&c); err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v", c)
}
```
