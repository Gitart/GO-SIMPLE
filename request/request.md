# How To Make HTTP Requests in Go



```go
package main

import (
  // "net/http"
  "fmt"
  "log"
  "time"
  "crypto/subtle"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
)

var Secretkey = "secret_fmLy5ekiaT1s6Yo"

func main(){
    e:= echo.New()

    // Check password
    // e.Use(middleware.BasicAuth(Basick))

    // Check API KEY
    // e.Use(middleware.KeyAuth(KeyCheck))

    // Middleware 
    // Для всех запросов
    e.Use(BasicMiddelware)

    // Гереация уникальго кода клиенту
    // в Headere clients - X-Request-Id 
    e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{Generator:CustomGenerator}))

    e.POST("/info/:id", Info)
    e.GET("/nfc/:nfc", NfcRead)
    // e.GET("/favicon.ico", faviconHandler)

    e.Logger.Fatal(e.Start(":1323"))
}

// Favico
func faviconHandler (c echo.Context) error {
     return c.String(200, "fvico-get")
}

// NFC использование
func Info (c echo.Context) error {
  par:= c.Param("id")
  // user, password, hasAuth := c.BasicAuth()

  fmt.Println(par)
  return c.String(200, par)
}

// NFC использование
func NfcRead (c echo.Context) error {
  par:= c.Param("nfc")
  // user, password, hasAuth := c.BasicAuth()

  fmt.Println(par)
  return c.String(200, par)
}

// Проверка пароля (логина)
func CheckNom(username string){
     if subtle.ConstantTimeCompare([]byte(username), []byte("joe"))==1{
        fmt.Println("ok")    
     }else{
        fmt.Println("Bad")    
     }
}

// Проверка по логину и паролю
func Basick(username, password string, c echo.Context) (bool, error) {
     CheckNom(username) 

  // Be careful to use constant time comparison to prevent timing attacks
  if  subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
      subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
      return true, nil
  }
  return false, nil
}

// Посылает в head 
// X-Request-Id 
func CustomGenerator() string {
    tg := time.Now().Format("150405")
    return "k:"+tg+":"+RandString(10)
}

// Проверка по ключу
func KeyCheck(key string, c echo.Context) (bool, error) {
     return key == Secretkey, nil
}

// Проверка everything
func AnyCheck(c echo.Context) (bool, error) {
     fmt.Println("okey check")
     return true, nil
}

// *****************************************************
// For all application
// e.Use(BasicMiddelware)    
// logging
// *****************************************************
func BasicMiddelware(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error { 
        

     fmt.Println("*************************** M I D D E L W A R E   G L O B A L  ****************************************") 
     fmt.Println("")

     // Выполнится перед отдчей
     c.Response().Before(func() {
        log.Println("before response")
     })

     // Выполнится после отдчи
     c.Response().After(func() {
       log.Println("after response")
     })

      s  := c.Request().Method
      st := c.Response().Status
      mm := c.Request().URL

      // ssec := map[string]interface{}{
      //                "ggr": "dddf",
      // }

      ssec:=`{"nom":"A234"}`

      // Response to header client
      c.Response().Header().Set("WWW-Authenticate", "Basic realm=Restricted")
      c.Response().Header().Set("Appcode", CustomGenerator())
      c.Response().Header().Set("Appset", ssec)
      
      // Эта опция будет требовать сохранить результат на диск
      // c.Response().Header().Set("Content-Type", "application/json")
      

      // w.Header().Set("Content-Type", "application/json")

      // st := c.Response().Size
      // c.Response().URL   //echo.GetPath(c.Path)

      fmt.Printf("%+v\n",c.Request())

      // get from client
      rq:=c.Request()
      fmt.Println("BasicAuth:---------------",  rq.BasicAuth)                    //  
      fmt.Println("MyKey:-------------------",  rq.Header["Mykey"])              // K-230003333-2030033-skdkkdk
      fmt.Println("Content-Type:------------",  rq.Header["Content-Type"])       // multipart/form-data; boundary=--------------------------007032604719999045714180 
      fmt.Println("Accept:------------------",  rq.Header["Accept"])             // [*/*
      fmt.Println("Accept-Encoding:---------",  rq.Header["Accept-Encoding"])    //  [gzip, deflate, br]
      fmt.Println("Accept-Encoding:---------",  rq.Header["Accept-Encoding"])    //  gzip, deflate, br
      fmt.Println("Authorization:-----------",  rq.Header["Authorization"])      //  [Basic am9lOnNlY3JldA==]
      fmt.Println("Authorization[0]:--------",  rq.Header["Authorization"])      //  Basic am9lOnNlY3JldA==
      fmt.Println("Postman-Token:-----------",  rq.Header["Postman-Token"])      //  Basic am9lOnNlY3JldA==
      fmt.Println("Agent User---------------",  rq.Header["User-Agent"])         //  PostmanRuntime/7.29.2 [PostmanRuntime/7.29.2]
      fmt.Println("Host --------------------",  rq.Host)                         // 
      fmt.Println("ContentLength------------",  rq.ContentLength)                // 
      fmt.Println("RemoteAddr---------------",  rq.RemoteAddr)                   // 192.168.0.101:3097 
      fmt.Println("Close     ---------------",  rq.Close)                        // false
      fmt.Println("RequestURI  -------------",  rq.RequestURI)                   // /info/122
      fmt.Println("TransferEncoding---------",  rq.TransferEncoding)             // []
      fmt.Println("Значение одного поля ----",  rq.FormValue("Tname"))           // Test Name
      fmt.Println("Значение всех полей------",  rq.Form)                         // map[Names:[KKK] Pasport:[KL] Tname:[Yhshs]]
     
      fmt.Println("Basick for Any Midleware ....")
      fmt.Println("Метод  : ", s)
      fmt.Println("Status : ", st)
      fmt.Println("Путь   : ", mm)

      return next(c)
  }
}

// Pause with parameter
func Pause(secondsleep time.Duration){
   time.Sleep(time.Second*secondsleep)
}


      
/*
...
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      fmt.Printf("server: %s /\n", r.Method)
      fmt.Printf("server: query id: %s\n", r.URL.Query().Get("id"))
      fmt.Printf("server: content-type: %s\n", r.Header.Get("content-type"))
      fmt.Printf("server: headers:\n")
      for headerName, headerValue := range r.Header {
          fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
      }

      reqBody, err := ioutil.ReadAll(r.Body)
      if err != nil {
             fmt.Printf("server: could not read request body: %s\n", err)
      }
      fmt.Printf("server: request body: %s\n", reqBody)

      fmt.Fprintf(w, `{"message": "hello!"}`)
  })
...
*/
```

