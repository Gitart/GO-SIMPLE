# Sample work with YAML
// https://play.golang.org/p/QN69Y-w4GZJ

```golang
package main

import (
    "fmt"
    "log"
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

//*****************************************
// Main
//*****************************************
func main() {
     Settings()
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
  pid:
    enabled:    true
    path:       "c:/out"
    override:   "c:/in"
    other: 
      enabled:  "lddldlldldlldldl:ssssssssss:sssssd:sssssss:ssssss:kskkskks"
system:
  path:      "inputpast"
  password:  "stttstsdpassWord001" 
  name:      "Timeran"
  googlekey: "KN-1002233" 


    
```  
