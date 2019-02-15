
## Get data for trand

```golang

func Test(w http.ResponseWriter, r *http.Request){
     // d:=Doget("https://api.bitfinex.com/v1/pubticker/btcusd","GET")
     d:=Doget("https://kuna.io/api/v2/tickers/btcuah","GET")

     // fmt.Println(d["ticker"].(map[string]interface {})["high"])
     f:=Mstr(d["ticker"],"high")
     l:=Mstr(d["ticker"],"low")
     s:=Mstw(d["ticker"])


      fmt.Println(s["low"])

     fmt.Println(f)
     fmt.Println(l)

     Mstw1(d)


      // convert map to json
    jsonString, _ := json.Marshal(d)
    fmt.Println(string(jsonString))

    // convert json to struct
    ss := Vl{}
    json.Unmarshal(jsonString, &ss)
    fmt.Println("At",ss.At)
    fmt.Println("Ticker",ss.Ticker)
    fmt.Println("Buy",ss.Ticker.Buy)
     // fmt.Printf("%X",d)       
    
}
```



## Structure

```golang
type Vl struct{
   
  At int              `json:"at"          db:"at"`
  Ticker struct {
    Buy string         `json:"buy"        db:"buy"`
    Sell string        `json:"sell"       db:"sell"`
    Low string         `json:"low"        db:"low"`
    High string        `json:"high"       db:"high"`
    Last string        `json:"last"       db:"last"`
    Vol string         `json:"vol"        db:"vol"`
    Price string       `json:"price"      db:"price"`
  }                    `json:"ticker"     db:"ticker"`
}
```

## Универсальная процедура конвертации json

```golang
func Mstw(Tr interface {}) map[string]interface {} {
        return Tr.(map[string]interface {})
}
```

## Универсальная процедура конвертации json

```golang
func Mstr(Tr interface {}, Field string) string {
        return Tr.(map[string]interface {})[Field].(string)  
}
```



## Универсальная процедура конвертации json

```golang
// *********************************************
// Get 
// Base get procedure
// *********************************************
func Doget(Url, Method string) map[string]interface{} {
  
    var Dat map[string]interface{}  
    client:= &http.Client{}

    reqest, err := http.NewRequest(Method, Url, nil)
    if err != nil {
       fmt.Println("GetAccountBalance http.NewRequest err",err)
       return nil
    }

    reqest.Header.Set("Content-Type", "application/json; charset=utf-8")
    reqest.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Trident/7.0; rv:11.0) like Gecko")
    reqest.Header.Set("Connection", "keep-alive")
    
    response, errdo := client.Do(reqest)
    if errdo != nil {
       fmt.Println("client.Do err",errdo)
       return nil
    }
    
    defer response.Body.Close()
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
       fmt.Println("ioutil.ReadAll err", err)
       return nil
    }
  
    jsonerr := json.Unmarshal(body, &Dat)
    if jsonerr != nil {
       fmt.Println("json unmarshal  err:", jsonerr.Error())   
       return nil
    } 

    // else {
       // if balanceresult.Status == "ok" {
       //    fmt.Println("GetAccountBalance result:",balanceresult)
       // }else{
       // fmt.Println("GetAccountBalance result err:")
    // }
    // }
   return Dat
}
```

# ID generetaor
```golang
func Ids() string{
     return strconv.FormatInt(time.Now().UnixNano()/1000000, 10) 
}
```


## Конвертация Map to JSON
```golang
package main

import (
    "encoding/json"
    "fmt"
)

type MyStruct struct {
    Id           string `json:"id"`
    Name         string `json:"name"`
    UserId       string `json:"user_id"`
    CreatedAt    int64  `json:"created_at"`
}

func main() {
    m := make(map[string]interface{})
    m["id"] = "2"
    m["name"] = "jack"
    m["user_id"] = "123"
    m["created_at"] = 5
    fmt.Println(m)

    // convert map to json
    jsonString, _ := json.Marshal(m)
    fmt.Println(string(jsonString))

    // convert json to struct
    s := MyStruct{}
    json.Unmarshal(jsonString, &s)
    fmt.Println(s)

}
```


## Второй вариант

```golang
package main

import (
    "fmt"
    "github.com/mitchellh/mapstructure"
)

type MyStruct struct {
    Id        string `json:"id"`
    Name      string `json:"name"`
    UserId    string `json:"user_id"`
    CreatedAt int64  `json:"created_at"`
}

func main() {
    input := map[string]interface{} {
        "id": "1",
        "name": "Hello",
        "user_id": "123",
        "created_at": 123,
    }
    var output MyStruct
    cfg := &mapstructure.DecoderConfig{
        Metadata: nil,
        Result:   &output,
        TagName:  "json",
    }
    decoder, _ := mapstructure.NewDecoder(cfg)
    decoder.Decode(input)

    fmt.Printf("%#v\n", output)
    // main.MyStruct{Id:"1", Name:"Hello", UserId:"123", CreatedAt:123}
}
```

