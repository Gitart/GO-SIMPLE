## Маленькие помощники

### Работа с картой функций
```golang
func Test(w http.ResponseWriter, rq *http.Request){
     p := rq.URL.Path[len("/api/test/"):]
     
    mf:=map[string] func() int {
   	  "one":  func()int{return 10},
  	  "two":  func()int{return 20},
  	  "tree": func()int{return 30},
  	  "four": func()int{return 40},
     }



   dat:=mf[p]()

   if dat<1{
   	  fmt.Println("Error")
   	  return
   }

   fmt.Println(dat)   
   fmt.Println("---->",mf["one"]())
   fmt.Fprintln(w,dat)
}
```

## Получение IP 
```golang
func TestIP(w http.ResponseWriter, rq *http.Request){
  t   :=rq.RemoteAddr
  ip  :=strings.Split(t, ":")[0]
  port:=strings.Split(t, ":")[1]
  
  c:=`
      <html><body>
            <h2>Показатели порта</h2>
           <h4>IP : %s</h4>
           <h4>Port : %s</h4>
      </html></body>
   `

  fmt.Fprintf(w, c, ip,port)
  fmt.Println(t, ip, port)
}
```


## Пример работы с очередью
```golang
func Test_Que(w http.ResponseWriter, rq *http.Request){
	 io.WriteString(w, "hello, world2!\n")
	 io.WriteString(w, "hello, world3!\n")
	 io.WriteString(w, "hello, world4!\n")


   c1 := make(chan string, 1)
    go func() {
    	c1 <- "Result 0"
        time.Sleep(time.Second * 2)
        c1 <- "result 1"
        c1 <- "result 2"
        c1 <- "result 3"

    }()
fmt.Println(<-c1)

  time.Sleep(time.Second * 2)
  fmt.Println(<-c1)
  fmt.Println(<-c1)
  fmt.Println(<-c1)

    go func() {
    	c1 <- "Result 10"
        time.Sleep(time.Second * 2)
        c1 <- "result 11"
        c1 <- "result 12"
        c1 <- "result 13"

    }()

  time.Sleep(time.Second * 2)
  fmt.Println(<-c1)
  fmt.Println(<-c1)
  fmt.Println(<-c1)
  fmt.Println(<-c1)


	   // fmt.Fprintln(w, "Входите пожалуйста")
	   // w.Write([]byte("Hello word"))

}
```


### Тестовая процедура
```golang
func Test_Auth(w http.ResponseWriter, rq *http.Request){

	 	    L,R,V:=rq.BasicAuth()
		    
		    fmt.Println("Basic Authontefication : ", L, R, V)
		    fmt.Println("Ok Task Get")   

		    Key := rq.Header.Get("Key")  

		    // Описание ошибок
		    // var ErrMissingFile = errors.New("http: ERRROR ")
		    // log.Fatal(ErrMissingFile)

		    agent:=rq.UserAgent()		    
            fmt.Println(agent)
            
            if strings.Contains(agent, "insomnia") {
               fmt.Println("Правильный вход")	
               // fmt.Fprintln(w, "Входите пожалуйста")
            }else{
               fmt.Println("Другой браузер")		
            }

            // Проверка ключиков
		    if V {
		       w.Header().Set("Access-Control-Allow-Origin", "*")
		       w.Header().Set("Password",  L)
		       w.Header().Set("UsersName", R)
		       w.Header().Set("Key",       Key)
		       // w.Write([]byte("Можно получить данные"))
		       w.Header().Set("Description", "Можно получить данные")
               w.WriteHeader(220)
		       
		    }else{
 	    	   w.WriteHeader(420)
		       return 
		    }
}
```
