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
