package main

import (
    "fmt"
    "flag"
    "github.com/bradfitz/gomemcache/memcache"
)

type Item struct {
    // Key is the Item's key (250 bytes maximum).
    Key string
    // Value is the Item's value.
    Value []byte
    // Flags are server-opaque flags whose semantics are entirely
    // up to the app.
    Flags uint32

    // Expiration is the cache expiration time, in seconds: either a relative
    // time from now (up to 1 month), or an absolute Unix epoch time.
    // Zero means the Item has no expiration time.
    Expiration int32

    // Compare and swap ID.
    casid uint64
}

const Ip = "127.0.0.1:11211"



var mc *memcache.Client
// https://medium.com/@litanin/basic-sql-memcached-golang-9e8fd5b7efe1

func main() {

    usGet := flag.Bool("get", false, "display Get")
    flag.Parse()

    mc = memcache.New(Ip)

   if !*usGet {
    Set("kuku:23","суп5ер сувет 10",10)
    Set("kuku","суп5ер суветwww 20",20)
    Set("kuku1","суп5ер суветww 30",30)
    Set("kuku2","суп5ер суветws 40",40)
    Set("kuku3","суп5ер сувеss  50",50)
    Set("kuku3","Замена сувеss  60",50)
    fmt.Println("Запись перменных ....")
   }

   
   if *usGet {
      fmt.Println("Чтение  ....")
     fmt.Println(Get("kuku:23"))
     fmt.Println(Get("kuku"))
     fmt.Println(Get("kuku1"))
     fmt.Println(Get("kuku2"))
     fmt.Println(Get("kuku3"))
    }

}


// Get
func Get(key string) string {

    ff,err:=mc.Get(key)
    if err!=nil{
        return ""
    }
    return string(ff.Value)
}

// Set
func Set(key, txt string, exp int32) {
    mc.Set(&memcache.Item{ Key   : key, 
                           Value : []byte(txt), 
                           Expiration: exp })
}
