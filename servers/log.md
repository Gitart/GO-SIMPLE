
## Log operations

```golang
package main

import (
  "log"
  "os"
)

func main() {
logfile, _ := os.Create("./log.txt")
defer logfile.Close()
logger := log.New(logfile, "example ", log.LstdFlags|log.Lshortfile)
logger.Println("This is a regular message.")
logger.Fatalln("This is a fatal error.")
logger.Println("This is the end of the function.")
}
```

## UDP loging

```golang
package main

import (
"log"
"net"
"time"
)

func main() {

timeout := 30 * time.Second
conn, err := net.DialTimeout("udp", "localhost:1902", timeout)

if err != nil {
   panic("Failed to connect to localhost:1902")
}

defer conn.Close()

f := log.Ldate | log.Lshortfile
logger := log.New(conn, "example ", f)
logger.Println("This is a regular message.")
logger.Panicln("This is a panic.")
}
```


## Log to 

```golang
package main
import (
"fmt"
"log"
"log/syslog"
)

func main() {
priority := syslog.LOG_LOCAL3 | syslog.LOG_NOTICE
flags := log.Ldate | log.Lshortfile
logger, err := syslog.NewLogger(priority, flags)

if err != nil {
  fmt.Printf("Can't attach to syslog: %s", err)
  return
}

logger.Println("This is a test log message.")
}
```
