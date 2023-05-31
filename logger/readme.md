## Log 

**Saving log messages to a custom log file in Golang**

### logging
Loggin in Unix log server in Golang
The log package in Golang, allows you to send log messages to log files. This log file can be syslog, mail.log, or your custom log file.
In this article, you'll learn how to do this.

### Logging to syslog
The example below shows you the way of sending log messages to /var/log/syslog file in Unix os. We are going to use log/syslog package to set the log file.

```go
package main

import (
    "log"
    "log/syslog"
    "os"
)


func main() {
    // Log to syslog
    logWriter, err := syslog.New(syslog.LOG_SYSLOG, "My Awesome App")
    if err != nil {
        log.Fatalln("Unable to set logfile:", err.Error())
    }
    // set the log output
    log.SetOutput(logWriter)

    log.Println("This is a log from GOLANG")
}
```

Now, if you run the application, This is a log from GOLANG will be logged in the syslog file.
You can check it using this command:

```go
package main

import (
    "log"
    "os"
    "sync"
    "time"
)

func main() {
    logFile, err := os.OpenFile("clog", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
    if err != nil {
        log.Fatalln(err)
    }
    log.SetOutput(logFile)

    var wg sync.WaitGroup

    wg.Add(3)

    go func() {
        defer wg.Done()
        for i := 0; i < 10; i++ {
            log.Println("F1; loop:", i)
            time.Sleep(time.Millisecond)
        }
    }()

    go func() {
        defer wg.Done()
        for i := 0; i < 10; i++ {
            log.Println("F2; loop:", i)
            time.Sleep(time.Millisecond * 2)
        }
    }()

    go func() {
        defer wg.Done()
        for i := 0; i < 10; i++ {
            log.Println("F3; loop:", i)
            time.Sleep(time.Millisecond * 3)
        }
    }()

    wg.Wait()

}
```
