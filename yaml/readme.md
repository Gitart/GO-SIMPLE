# Sample work with YAML


```golang
package main

import (
	"fmt"
	"flag"
	"os"
	"runtime"
  "github.com/fatih/color"
  "gopkg.in/yaml.v2"
  "log"
  "io/ioutil"
  "strings"
  "encoding/json"
)


func Settings(){
    s:=Cor{}
    content, err := ioutil.ReadFile("./set.yaml")
    
    if err != nil {
      log.Fatal(err)
    }

    yaml.Unmarshal([]byte(content), &s)
    fmt.Println ("Cored-adress........  ",  s.Core.Address)
    fmt.Println ("Cored-port..........  ",  s.Core.Port)
    
}
```

## Model structure
```golang
type Cor struct{
      Core struct{
        Port          string `yaml:"port"`
        Address       string `yaml:"address"`
        Enabled       string `yaml:"enabled"`
        Worker:       string `yaml:"worker_num"`
        Queue_num:    string `yaml:"queue_num"` 
        Sync:         string `yaml:"sync"` 
        Mode:         string `yaml:"mode"` 
        Ssl:          string `yaml:"ssl"` 
        Cert_path:    string `yaml:"cert_path"` 
        Key_path:     string `yaml:"key_path"` 
      }  
}
```


## setting.yaml
```yaml
core:
  enabled:    true                    # enabale httpd server
  address:    "192.168.10.11"         # ip address to bind (default: any)
  port:       "8088"                  # ignore this port number if auto_tls is enabled (listen 443).
  worker_num: 0                       # default worker number is runtime.NumCPU()
  queue_num:  0                       # default queue number is 8192
  max_notification: 100
  sync:       false                   # set true if you need get error message from fail push notification in API response.
  mode:       "release"
  ssl:        false
  cert_path:  "cert.pem"
  key_path:   "key.pem"
```  
