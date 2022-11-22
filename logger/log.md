

```golang
f, err := os.OpenFile("testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
if err != nil {
    log.Fatalf("error opening file: %v", err)
}
defer f.Close()

log.SetOutput(f)
log.Println("This is a test log entry")
```

### Log

```golang
package logger

import (
  "flag"
  "os"
  "log"
  "go/build"
)

var (
  Log      *log.Logger
)


func init() {
    // set location of log file
    var logpath = build.Default.GOPATH + "/src/chat/logger/info.log"

   flag.Parse()
   var file, err1 = os.Create(logpath)

   if err1 != nil {
      panic(err1)
   }
      Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
      Log.Println("LogFile : " + logpath)
}
```

### import the package wherever you want to log e.g main.go

```golang
package main

import (
   "logger"
)

const (
   VERSION = "0.13"
 )

func main() {

    // time to use our logger, print version, processID and number of running process
    logger.Log.Printf("Server v%s pid=%d started with processes: %d", VERSION, os.Getpid(),runtime.GOMAXPROCS(runtime.NumCPU()))

}
```


### Logs

```golang
package main

import (
    "log"
    "os"
)
var (
    outfile, _ = os.Create("path/to/my.log") // update path for your needs
    l      = log.New(outfile, "", 0)
)

func main() {
    l.Println("hello, log!!!")
}
```

```golang
import (
    "os/exec"
)

func main() {
    // check error here...
    exec.Command("/bin/sh", "-c", "echo "+err.Error()+" >> log.log").Run()
}
```
