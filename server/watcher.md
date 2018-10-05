## Watcher

```golang
func main() {
    f, err := os.OpenFile("testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if err != nil {
        fmt.Println("error opening file: %v", err)
    }
    defer f.Close()

    log.SetOutput(f)

    http.HandleFunc("/", defaultHandler)
    http.HandleFunc("/check", checkHandler)

    serverErr := http.ListenAndServe("127.0.0.1:8080", nil) // set listen port

    if serverErr != nil {
        log.Println("Error starting server")

    } else {
        fmt.Println("Started server on - 127.0.0.1:8080" )
    }
}
```

The above code will start a local server on 8080 and I am able to hit the routes via browser. It's all good!

However, now I want to run a separate go routine, that watches a file -

```golang
func initWatch() string{
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        fmt.Println(err)
    }
    defer watcher.Close()

    done := make(chan bool)
    go func() {
        for {
            select {
                case event := <-watcher.Events:
                    if ( event.Op&fsnotify.Remove == fsnotify.Remove || event.Op&fsnotify.Rename == fsnotify.Rename ) {
                        fmt.Println("file removed - ", event.Name)
                    }

                case err := <-watcher.Errors:
                    fmt.Println("error:", err)
                }
        }
    }()

    err = watcher.Add("sampledata.txt")
    if err != nil {
        fmt.Println(err)
    }


    <-done
}
```

And now, if I call the function initWatch() 
BEFORE http.ListenAndServe("127.0.0.1:8080", nil) 

## initWatch()

```golang
serverErr := http.ListenAndServe("127.0.0.1:8080", nil) // set listen port
```

And if I call the function  initWatch() AFTER http.ListenAndServe("127.0.0.1:8080", nil) 
then the file watcher function is not working. Ex -

serverErr := http.ListenAndServe("127.0.0.1:8080", nil) // set listen port
initWatch()

