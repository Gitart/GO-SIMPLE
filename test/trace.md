## go tool trace


```golang
package main

import (
    "os"
    "runtime/trace"
)

func main() {
    f, err := os.Create("trace.out")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    err = trace.Start(f)
    if err != nil {
        panic(err)
    }
    defer trace.Stop()

  // всея остальная логика вашей программы
}
```

Это позволить сохранят все события, произошедшие в программе, в бинарный файл trace.out. 
После этого можно запускать go tool trace trace.out
