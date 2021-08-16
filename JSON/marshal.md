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
https://play.golang.org/p/7nk5ZEbVLw

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


```go
package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type CustomBool struct {
	Bool bool
}

func (cb *CustomBool) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"true"`, `true`, `"1"`, `1`:
		cb.Bool = true
		return nil
	case `"false"`, `false`, `"0"`, `0`, `""`:
		cb.Bool = false
		return nil
	default:
		return errors.New("CustomBool: parsing \"" + string(data) + "\": unknown value")
	}
}

func (cb CustomBool) MarshalJSON() ([]byte, error) {
	json, err := json.Marshal(cb.Bool)
	return json, err
}

type Target struct {
	Id     int        `json:"id"`
	Active CustomBool `json:"active"`
}

func main() {
	jsonString := `[{"id":1,"active":true},
					{"id":2,"active":"true"},
					{"id":3,"active":"1"},
					{"id":4,"active":1},
					{"id":5,"active":false},
					{"id":6,"active":"false"},
					{"id":7,"active":"0"},
					{"id":8,"active":0},
					{"id":9,"active":""}]`

	targets := []Target{}

	_ = json.Unmarshal([]byte(jsonString), &targets)

	for _, t := range targets {
		fmt.Println(t.Id, "-", t.Active.Bool)
	}

	jsonStringNew, _ := json.Marshal(targets)
	fmt.Println(string(jsonStringNew))
}
```
