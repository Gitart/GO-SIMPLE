
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

func Settings(){
    s:=Cor{}
    content, err := ioutil.ReadFile("./set.yaml")
    
    if err != nil {
      log.Fatal(err)
    }

    yaml.Unmarshal([]byte(content), &s)
    fmt.Println ("Cored-adress........  ",  s.Core.Address)
    fmt.Println ("Cored-port..........  ",  s.Core.Port)
    fmt.Println ("Cored-pid...........  ",  s.Core.Pid.Enabled)

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
        }
      }  
}
