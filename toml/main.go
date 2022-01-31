package main

import (
    "github.com/BurntSushi/toml"
    "time"
    "fmt"
)

type Config struct {
    Title   string
    Owner   ownerInfo
    DB      database `toml:"database"`
    Servers map[string]server
    Clients clients
}

type ownerInfo struct {
    Name string
    Org  string `toml:"organization"`
    Bio  string
    DOB  time.Time
}

type database struct {
    Server  string
    Ports   []int
    ConnMax int `toml:"connection_max"`
    Enabled bool
}

type server struct {
    IP string
    DC string
}

type clients struct {
    Data  [][]interface{}
    Hosts []string
}



func main() {
    var conf Config
    if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
        fmt.Println(err)
    }

    fmt.Println(conf.Title)
    fmt.Println(conf.Servers["alpha"])
    fmt.Println(conf.Servers["alpha"].IP)
    fmt.Println(conf.Servers["alpha"].DC)
    fmt.Println(conf.Servers["betta"].DC)
    fmt.Println(conf.Servers["omega"])
    
    fmt.Println(conf.Owner.Name)
    fmt.Println(conf.Owner.DOB)
    fmt.Println(conf.DB.Server)
    fmt.Println(conf.DB.ConnMax)
    fmt.Println(conf.DB.Ports)
    fmt.Println(conf.Clients.Data[1][1])
    fmt.Println(conf.Clients.Hosts[1])

    for _, prt:=range conf.DB.Ports{
        fmt.Println("Port Dtb: ", prt)
    }

}
