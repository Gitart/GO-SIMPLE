# Основные особенности сервера


```golang
//******************************************************************** 
// http://nesv.github.io/golang/2014/02/25/worker-queues-in-go.html
//******************************************************************** 
func main(){
	
   // Create Redis Client
   redisUrl := getEnv("REDIS_URL",      "localhost:6379")
   redisPwd := getEnv("REDIS_PASSWORD", "")

   log.Printf("Connecting to Redis Url '%s'\n", redisUrl)
   log.Printf("Password to '%s'\n",             redisPwd)

   http.HandleFunc("/",                            Startserv)        // Регистрация в сервисе
   http.HandleFunc("/api/1/",                      test)             // Регистрация в сервисе

   // Ловитель жемчуга
   // go Worker()

   srv := &http.Server{Addr:":8080", ReadTimeout:10*time.Second, WriteTimeout:10 * time.Second}

    // Start Server
    go func() {
 	log.Println("Starting Server")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
    }()

    // Graceful Shutdown
    waitForShutdown(srv)
}
```


Static Page
```golang
/********************************************************************************************************************************
 *
 *   Статические странички
 *   c установкой разрешений и доступов на операции
 *
 *   http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {http.ServeFile(w, r, r.URL.Path[1:])})
 *   http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {http.FileServer(http.Dir("/static/"))})
 *
 *   /static/....
 *
 *********************************************************************************************************************************/
func StaticPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allows
	// Allows
	   // if origin := r.Header().Get("Origin"); origin != "" {
	   //  w.Header().Set("Access-Control-Allow-Origin", origin)
	   //  w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	   //  w.Header().Set("Access-Control-Allow-Headers",  "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	   //  w.Header().Set("Access-Control-Max-Age", "86400") // 24 hours
	   // }
	
	//  File static page
	http.ServeFile(w, r, r.URL.Path[1:])
}
```

Остановка сервера

```golang
// ***************************************
// Shootdown server
// ***************************************
func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down server")
	os.Exit(0)
}
```

Чтение переменных среды

```golang
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
```

Чтение тела

```golang
//************************************************************
// Read Body from http
// Htt("http://golang.org/pkg/net/http/")
//************************************************************
func Htt(link string) {
	res, err := http.Get(link)
	checkErr(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	checkErr(err)
	fmt.Println(string(body))
}
```


Обработчик ошибок

```golang
//************************************************************
// Check error
// Простой обработчик ошибок
//************************************************************
func checkErr(e error) {
	if e != nil {
		// panic(e)
		log.Println("Error : ", e.Error())
	}
}
```


Получение парметра из строки запроса
```golang
// *****************************************************************************
// Получение параметра
// *****************************************************************************
func Urp(r *http.Request, Len string) string {
	 return r.URL.Path[len(Len):]
}
```


Копирование файла 
```golang
// ********************************************************
// Copy files from -> to
// https://opensource.com/article/18/6/copying-files-go
// ********************************************************
func CopyFile(src, dst string) (int64, error) {
        sourceFileStat, err := os.Stat(src)
        if err != nil {
           return 0, err
        }

        if !sourceFileStat.Mode().IsRegular() {
           return 0, fmt.Errorf("%s is not a regular file", src)
        }

        source, err := os.Open(src)
        if err != nil {
           return 0, err
        }
        defer source.Close()

        destination, err := os.Create(dst)
        if err != nil {
           return 0,  err
        }
        defer destination.Close()
        nBytes, err := io.Copy(destination, source)
        return nBytes, err
}
```

Удаление элмента из массива
```golang
//*****************************************************************************************************************  
// Удаление элмента из массива
//*****************************************************************************************************************  
func DelInArray(DeletedElement string, Ins []string) []string{
var na []string

for _, v := range Ins {
    if v == DeletedElement {
        continue
    } else {
        na = append(na, v)
    }
}
     return na
}
```


Seeding with the same value results in the same random sequence each run.

```golang
/ ***********************************************
// Можно использовать для совета дня 
// или для ответа на вопросы
//***********************************************

func Answer(){
	// Seeding with the same value results in the same random sequence each run.
	// For different numbers, seed with a different value, such as
	// time.Now().UnixNano(), which yields a constantly-changing number.
	rand.Seed(42)
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])
}
```


Создание структуры сайта по файлу JSON
```golang
// ******************************************************************
// Создание структуры сайта по файлу JSON
// ******************************************************************
func CreateStructureSite() {
	var Dt Setting
    
    
    // Директории в которых будет созданы поддиретктории
	Dr := []string{"in", "out", "bak"}

	
	file, _ := ioutil.ReadFile("./setting.json")
	json.Unmarshal(file, &Dt)
	// fmt.Println(Dt.Menu[1].Head)

	for _, l := range Dr {
		// Созданеи верхнего уровня каталога
		os.Mkdir(l, 0777)

		for _, t := range Dt.Groups {
			d := l + "/" + t.Code
			err := os.Mkdir(d, 0777)

			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("Created directory", d)
		}
	}

	fmt.Println("All directory was creating....")
}
```
