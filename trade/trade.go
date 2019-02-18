// https://firebase.google.com/pricing/

package main

import (
  "net/http"	
  "log"
  "time"
  "fmt"	
  "io/ioutil"
  "strconv"
  "encoding/json"
  // "database/sql"
  // "database/sql/driver"
  // "errors"
  // "context"
  // "os/signal"
  // "os"
  "strings"
  // "syscall"
  // "github.com/jinzhu/gorm"
  //  shell "github.com/ipfs/go-ipfs-api"
  "bufio"
  "os"
  "flag"

)


var(
   period *int =flag.Int("period", 3, "Period getting...")

)

type Mst map[string]interface{}

/*
    Main procedure
*/
func main() {
    
flag.Parse()
fmt.Println("Period:", *period)

    go GetDataN()
//go GhP()

   // Start page for service 
   http.HandleFunc("/",                                     Startserv)               // Registration
   http.HandleFunc("/static/",                              StaticPage)              // Static  page
   http.HandleFunc("/test/",                                Test)                    // Test

   // Settig 
   srv := &http.Server{Addr:":8888", ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second}

   // Start Service
   fmt.Println("Server started...")

  if err := srv.ListenAndServe(); err != nil {
     log.Println("Error listening servers", err.Error())
	}
}


/*
  Static pages
  /static/....
*/
func StaticPage(w http.ResponseWriter, r *http.Request) {
     w.Header().Set("Access-Control-Allow-Origin", "*") // Allows
     http.ServeFile(w, r,r.URL.Path[1:])
 }


func Startserv(w http.ResponseWriter, r *http.Request) {

}



// *******************************************************************************************************
func Test(w http.ResponseWriter, r *http.Request){
    
     for {
          time.Sleep(time.Second * 10)
          GetDataN()
   }
}


// Structure
type Vl struct{
  At int              `json:"at"          db:"at"`
  Ticker  struct {
    Buy   string       `json:"buy"        db:"buy"`        // Покупка 
    Sell  string       `json:"sell"       db:"sell"`       // Продажа 
    Low   string       `json:"low"        db:"low"`        // Низкая самая
    High  string       `json:"high"       db:"high"`       // Высокая
    Last  string       `json:"last"       db:"last"`       // Последняя на бирже - показ на графике
    Vol   string       `json:"vol"        db:"vol"`        // Объем
    Price string       `json:"price"      db:"price"`      // Цена
  }                    `json:"ticker"     db:"ticker"`
}


// Цикл прокрутки через промежуток времени
func GetDataN(){

fmt.Println("Date:Time      Last   Sell   Buy    Отнош Зн  My Perc")	

   for {
          GetDataS()
          time.Sleep(time.Second * 5)
   }
}



var fbuy int64
var fsel int64
var sts  string  // status


// ****************************************************
// Get data 
// ****************************************************
func GetDataS(){
    // d:=Doget("https://api.bitfinex.com/v1/pubticker/btcusd","GET")
    d:=Doget("https://kuna.io/api/v2/tickers/btcuah","GET")

    // convert map to json
    jsonString, _ := json.Marshal(d)

    // convert json to struct
    ss := Vl{}
    json.Unmarshal(jsonString, &ss)


    myp := 94251.0            // Я купил
    myy := myp*1.02           // Должен продать 
    // tek := myp - tk.Sell      // текущая разница


    tk:=ss.Ticker

    // fmt.Println("Купил по цене                ", myp)
    // fmt.Println("Надо продать с наценкой 2.5% ", myy)
    // fmt.Println("Разница                      ", Str2float64(tk.Sell)-myy)
    // fmt.Println("Разница                      ", Str2float64(tk.Sell)-myp)
    // fmt.Println("Сейчас продают               ", Str2int64(tk.Sell))

    // fmt.Println("-------------------------")
    // fmt.Println("At    ", ss.At)
    
    // fmt.Println("Buy   ", Str2int64(tk.Buy))
    // fmt.Println("Sell  ", Str2int64(tk.Sell))
    
    // fmt.Println("Low   ", Str2int64(tk.Low))
    // fmt.Println("High  ", Str2int64(tk.High))
    // fmt.Println("Last  ", Str2int64(tk.Last))
    // fmt.Println("Vol   ", Str2int64(tk.Vol))
    // fmt.Println("Price ", Str2int64(tk.Price))
    // fmt.Println(" ")
    // Datetime
     
 

// Если поменялось значение продажи или покупки
// Обновляем показания в окне     
if fbuy!=Str2int64(tk.Buy) || fsel!=Str2int64(tk.Sell){
 
   // Статус 
   if fbuy!=Str2int64(tk.Buy) {sts="Покупка"}       // Изменилась продажа  
   if fsel!=Str2int64(tk.Sell){sts="Продажа"}       // Изменилась покупка
   otn    := Str2float64(tk.Sell)/ Str2float64(tk.Buy)    //  Отношение покупки к продаже

   fmt.Printf("%s  %6v %6v %6v %6v %6v %6v  %s \n", dtt(), Str2int64(tk.Last), Str2int64(tk.Buy), Str2int64(tk.Sell), otn, myp, myy, sts)
}


  fbuy=Str2int64(tk.Buy)
  fsel=Str2int64(tk.Sell)
}


func dtt()string{
 tm:=time.Now().Format("15:04:05.000")
 return tm
}


// Конвертация Без дисятичных
func Str2int64(s string) int64 {
  ss:=strings.Split(s,".")[0]
  i, err := strconv.ParseInt(ss, 10, 64)
  // i, err := strconv.ParseFloat(s, 64) 
  if err != nil {
    fmt.Println(err.Error())
  }
  return i
}


func Str2float64(s string) float64 {
   ss:=strings.Split(s,".")[0]
  i, err := strconv.ParseFloat(ss, 64) 
  if err != nil {
    fmt.Println(err.Error())
  }
  return i
}


// Get data
func GetData(){
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

func Tre(){
}


func Mstw1(Tr  Mst) {
  fmt.Println(Tr)
}

func Mstw(Tr interface {}) map[string]interface {} {
        return Tr.(map[string]interface {})
}

func Mstr(Tr interface {}, Field string) string {
     return Tr.(map[string]interface {})[Field].(string)  
}

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


func Ids() string{
     return strconv.FormatInt(time.Now().UnixNano()/1000000, 10) 
}

// https://stackoverflow.com/questions/30109061/golang-parse-html-extract-all-content-with-body-body-tags
// https://mozgcorp.ru/post/parsing-web-stranic-na-go-s-pomoshchyu-goquery
// http://qaru.site/questions/236120/extract-links-from-a-web-page-using-go-lang
// https://benjamincongdon.me/blog/2018/03/01/Scraping-the-Web-in-Golang-with-Colly-and-Goquery/
// http://golang-examples.tumblr.com/post/47426518779/parse-html
// https://mozgcorp.ru/post/parsing-web-stranic-na-go-s-pomoshchyu-goquery

// Описан метод с настройкой байт
// https://medium.com/golangspec/in-depth-introduction-to-bufio-scanner-in-golang-55483bb689b4


// Поиск элемента в документе
func GhP(){
	// Где ищем
	url := "https://ru.tradingview.com/symbols/BTCUSDT/technicals/"
	// url  = "https://ru.tradingview.com/symbols/BTCUSDT/markets/"
	url="https://www.bbc.com/news/world-us-canada-47203706"
	url="https://www.bbc.com/news/world-asia-china-47265411"

	req, _ := http.NewRequest("GET", url, nil)
	// req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("cache-control", "no-cache")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
    Bods    := string(body)
    
    // Ищем все span
    // Или любые другие элеименты на странице
    serarch:="<p "
    // serarch=""
    
     // Scaner 
     scanner := bufio.NewScanner(strings.NewReader(Bods))
     buf := make([]byte, 0, 1024*1024)
     scanner.Buffer(buf, bufio.MaxScanTokenSize)

    // Прокрутка боди документа
    for scanner.Scan() {
	
  		 // Обрезание пустого пространства
		 sk:=strings.TrimSpace(scanner.Text())

		 if strings.Contains(sk,serarch){
		    fmt.Println(">",sk) 
		 }
	   }

    // Check Error
	if err := scanner.Err(); err != nil {
	   fmt.Fprintln(os.Stderr, "reading standard input:", err.Error())
	}
    // sd:=strings.IndexAt(Bods, "speedometerSignal-pyzN--tL- brandColor-1WP1oBmS-")
	// fmt.Println(dd)
	// fmt.Println(string(body))
	// speedometerSignal-pyzN--tL- brandColor-1WP1oBmS-
}

