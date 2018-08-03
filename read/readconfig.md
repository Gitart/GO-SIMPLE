## Read Config

```go
package main

/*
cat config.json
{"object": 
    {
       "buffer_size": 10,
       "Databases":
       [
               {
                       "host": "localhost",
                       "user": "root",
                       "pass": "",
                       "type": "mysql",
                       "name": "go",
                       "Tables":
                       [
                               {
                                       "name": "testing",
                                       "statment": "teststring",
                                       "regex": "teststring ([0-9]+) ([A-z]+)",
                                       "Types": 
                                        [
                                           {
                                               "id": "int",
                                               "value": "string"
                                           }
                                        ]
                               }
                       ]
               }
       ]
    }
}
*/

import (
    "fmt"
    "os"
    "json"
    "io/ioutil"
)

type jsonobject struct {
    Object ObjectType
}

type ObjectType struct {
    Buffer_size int
    Databases   []DatabasesType
}

type DatabasesType struct {
    Host   string
    User   string
    Pass   string
    Type   string
    Name   string
    Tables []TablesType
}

type TablesType struct {
    Name     string
    Statment string
    Regex    string
    Types    []TypesType
}

type TypesType struct {
    Id    string
    Value string
}

// Main function
// I realize this function is much too simple I am simply at a loss to

func main() {
    file, e := ioutil.ReadFile("./config.json")
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))

    //m := new(Dispatch)
    //var m interface{}
    var jsontype jsonobject
    json.Unmarshal(file, &jsontype)
    fmt.Printf("Results: %v\n", jsontype)
}

```
