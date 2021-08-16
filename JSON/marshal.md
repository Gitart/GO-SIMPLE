## Переопределение стандартных методов 
https://play.golang.org/p/7g5i-53SVz7

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



## Преобразование
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Animal uint

const (
	CommonAnimal Animal = iota
	Zebra
	Gopher
)

func (a *Animal) String() string {
	switch *a {
	case Gopher:
		return "gopher"
	case Zebra:
		return "zebra"
	default:
		return "common-animal"
	}
}

func sToAnimal(s string) Animal {
	switch strings.ToLower(s) {
	case "gopher":
		return Gopher
	case "zebra":
		return Zebra
	default:
		return CommonAnimal
	}
}

func (a *Animal) UnmarshalJSON(b []byte) error {
	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return nil
	}
	aAnimal := sToAnimal(unquoted)
	*a = aAnimal
	return nil
}

func (a *Animal) MarshalJSON() ([]byte, error) {
	quoted := strconv.Quote(a.String())
	return []byte(quoted), nil
}

func Example_marshalJSON() {
	zoo := []Animal{Gopher, Zebra, CommonAnimal, Gopher, Gopher, Zebra}
	marshaled, err := json.Marshal(&zoo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", marshaled)

	// Output:
	// ["gopher","zebra","common-animal","gopher","gopher","zebra"]
}

func Example_unmarshalJSON() {
	rawZooManifest := `["zebra", "zebra", "common-animal", "gopher", "gopher"]`
	var zooManifest []*Animal
	if err := json.Unmarshal([]byte(rawZooManifest), &zooManifest); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", zooManifest)
	// Output:
	// [zebra zebra common-animal gopher gopher]
}

func main() {
	Example_unmarshalJSON()
	Example_marshalJSON()
}
```

# Serialize
```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {
	a := make(Stuff)
	a[1] = "asdf"
	a[-1] = "qwer"
	fmt.Println("Initial:     ",a)

	stuff, err := json.Marshal(a)
	fmt.Println("Serialized:  ", string(stuff), err)

	b := make(Stuff)
	err = json.Unmarshal(stuff, &b)
	fmt.Println("Deserialized:", b, err)
}

type Stuff map[int]string

func (this Stuff) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")
	length := len(this)
	count := 0
	for key, value := range this {
		jsonValue, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		buffer.WriteString(fmt.Sprintf("\"%d\":%s", key, string(jsonValue)))
		count++
		if count < length {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}

func (this Stuff) UnmarshalJSON(b []byte) error {
	var stuff map[string]string
	err := json.Unmarshal(b, &stuff)
	if err != nil {
		return err
	}
	for key, value := range stuff {
		numericKey, err := strconv.Atoi(key)
		if err != nil {
			return err
		}
		this[numericKey] = value
	}
	return nil
}
```
