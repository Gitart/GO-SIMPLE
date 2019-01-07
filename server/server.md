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
