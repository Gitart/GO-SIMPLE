# Working With YAML

---

Here's the YAML file.

```
[root@centos helloworld]# cat switches.yaml
datacenters:
  dc01:
    - switch-01
    - switch-02
    - switch-03

```

Parse a YAML file and access its data.

```
package main

import (
    "fmt"
    "github.com/beego/goyaml2"
    "os"
    "reflect"
)

func main() {
    // Open the file to read
    file, err := os.Open("/gosrc/src/helloworld/switches.yaml")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // Read yaml from file
    fileData, err := goyaml2.Read(file)
    if err != nil {
        panic(err)
    }

    // To access data, assert yaml to map[string]interface{}
    yaml, ok := fileData.(map[string]interface{})
    if !ok {
        panic("Yaml not a map")
    }

    // If nested maps, assert it again to access nested data
    datacenters, ok := yaml["datacenters"].(map[string]interface{})
    if !ok {
        panic("Datacenters is not a map")
    }

    // This time, it's a slice that's nested in the map
    dc01, ok := datacenters["dc01"].([]interface{})
    if !ok {
        panic("dc01 is not a slice")
    }

    fmt.Println("Nested Yaml is of type:", reflect.TypeOf(dc01))

    // Iterate through slice
    for _, c := range dc01 {
        fmt.Println(c)
    }
}

```

Output

```
[root@centos helloworld]# ./hello
Nested Yaml is of type: []interface {}
switch-01
switch-02
switch-03
```
