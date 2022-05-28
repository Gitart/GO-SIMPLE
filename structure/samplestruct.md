## Go struct tags

last modified February 16, 2022

Go struct tags tutorial shows how to work with struct tags in Golang.

A struct is a user-defined type that contains a collection of fields. It is used to group related data to 
form a single unit. A Go struct can be compared to a lightweight class without the inheritance feature.

A struct tag is additional meta data information inserted into struct fields. The meta data can be acquired 
through reflection. Struct tags usually provide instructions on how a struct field is encoded to or decoded from a format.

Struct tags are used in popular packages including:

    encoding/json
    encoding/xml
    gopkg.in/mgo.v2/bson
    gorm.io/gorm
    github.com/gocarina/gocsv
    gopkg.in/yaml.v2

Go struct tag json

In the next example, we use the json struct tag with the encoding/json package.
main.go

package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    Id         int    `json:"id"`
    Name       string `json:"name"`
    Occupation string `json:"occupation,omitempty"`
}

func (p User) String() string {

    return fmt.Sprintf("User id=%v, name=%v, occupation=%v",
        p.Id, p.Name, p.Occupation)
}

func main() {

    user := User{Id: 1, Name: "John Doe", Occupation: "gardener"}
    res, _ := json.MarshalIndent(user, " ", "  ")

    fmt.Println(string(res))

    user2 := User{Id: 1, Name: "John Doe"}
    res2, _ := json.MarshalIndent(user2, " ", "  ")

    fmt.Println(string(res2))
}

The example uses struct tags to configure how JSON data is encoded.

type User struct {
    Id         int    `json:"id"`
    Name       string `json:"name"`
    Occupation string `json:"occupation,omitempty"`
}

With `json:"id"` struct tag, we encode the Id field in lowercase. In addition, the omitempty omits the Occupation field if it is empty.

$ go run main.go
{
    "id": 1,
    "name": "John Doe",
    "occupation": "gardener"
    }
{
    "id": 1,
    "name": "John Doe"
}

Go struct tag xml

In the following example, we use the xml struct tag with the encoding/xml package.
main.go

package main

import (
    "encoding/xml"
    "fmt"
)

type User struct {
    Id         int    `xml:"id"`
    Name       string `xml:"name"`
    Occupation string `xml:"occupation"`
}

func (p User) String() string {

    return fmt.Sprintf("User id=%v, name=%v, occupation=%v",
        p.Id, p.Name, p.Occupation)
}

func main() {
    user := User{Id: 1, Name: "John Doe", Occupation: "gardener"}

    res, _ := xml.MarshalIndent(user, " ", "  ")

    fmt.Println(xml.Header + string(res))
}

The example turns a Go structure into XML format. Using struct tags we can configure the output.

$ go run main.go
<?xml version="1.0" encoding="UTF-8"?>
 <User>
   <id>1</id>
   <name>John Doe</name>
   <occupation>gardener</occupation>
 </User>

Go struct tag csv

The next example uses the csv struct tag with the github.com/gocarina/gocsv library. The package provides easy serialization and deserialization functions to use CSV in Golang.
main.go

package main

import (
    "fmt"

    "github.com/gocarina/gocsv"
)

type User struct {
    Id         string `csv:"user_id"`
    Name       string `csv:"user_name"`
    Occupation string `csv:"user_occupation"`
}

func (p User) String() string {

    return fmt.Sprintf("User id=%v, name=%v, occupation=%v",
        p.Id, p.Name, p.Occupation)
}

func main() {

    users := []User{}

    users = append(users, User{Id: "1", Name: "John Doe", Occupation: "gardener"})
    users = append(users, User{Id: "2", Name: "Roger Doe", Occupation: "driver"})

    res, _ := gocsv.MarshalString(users)

    fmt.Println(res)
}

We transform a slice of user structures into CSV; the struct tags configure the header names.

$ go run main.go
user_id,user_name,user_occupation
1,John Doe,gardener
2,Roger Doe,driver
