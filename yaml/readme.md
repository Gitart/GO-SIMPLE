# Sample work with YAML
// https://play.golang.org/p/QN69Y-w4GZJ

```go
package main

import (
    "fmt"
    "log"
    "io/ioutil"
    "gopkg.in/yaml.v3"
)

//*****************************************
// Main
//*****************************************
func main() {
     Settings()
      Marsh()
}

type Cor struct{
      Core struct{
        Port          string `yaml:"port"`
        Address       string `yaml:"address"`
        Enabled       string `yaml:"enabled"`
        Worker        string `yaml:"worker_num"`
        Queue_num     string `yaml:"queue_num"` 
        Sync          string `yaml:"sync"` 
        Mode          string `yaml:"mode"` 
        Ssl           string `yaml:"ssl"` 
        Cert_path     string `yaml:"cert_path"` 
        Key_path      string `yaml:"key_path"` 
        Pid struct {
                Enabled       string `yaml:"enabled"`
                Path          string `yaml:"path"`
                Override      string `yaml:"override"`
                Other struct {
                               Enabled  string `yaml:"enabled"`
                              }
        }
      }  
   System struct{
    Path        string `yaml:"path"`
    Password    string `yaml:"password"`
    Name        string `yaml:"name"`
    Title       string `yaml:"title"`
    Key         string `yaml:"key"`  
    Googlekey   string `yaml:"googlekey"`
   }
}


func Settings(){
    s:=Cor{}
    content, err := ioutil.ReadFile("./setting.yaml")
    
    if err != nil {
      log.Fatal(err)
    }

    yaml.Unmarshal([]byte(content), &s)
    fmt.Println ("Cored-adress........  ",  s.Core.Address)
    fmt.Println ("Cored-port..........  ",  s.Core.Port)
    fmt.Println ("Cored-pid...........  ",  s.Core.Pid.Enabled)
    fmt.Println ("Cored-pid...........  ",  s.Core.Pid.Override)
    fmt.Println ("Cored-pid...........  ",  s.Core.Cert_path)
    fmt.Println ("Cored-pid...........  ",  s.Core.Pid.Other.Enabled)
    fmt.Println ("Sysyte:Password.....  ",  s.System.Password)
    fmt.Println ("Sysyte:Googlekey....  ",  s.System.Googlekey)
}


// Marshal to string
func Marsh(){

data:=`
core:
  enabled:    true                    # enabale httpd server
  address:    "192.168.10.11"         # ip address to bind (default: any)
  port:       "8088"                  # ignore this port number if auto_tls is enabled (listen 443).
`

     // m := make(map[interface{}]interface{})
     t:= Cor{}

      err := yaml.Unmarshal([]byte(data), &t)
        if err != nil {
                log.Fatalf("error: %v", err)
        }
    
    d, err := yaml.Marshal(&t)
        if err != nil {
           log.Fatalf("error: %v", err)
        }
     
    fmt.Printf("--- m dump:\n%s\n\n", string(d))
}
```
