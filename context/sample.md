## Context sample

https://play.golang.org/p/yTY_iJUIQBT


```go
package main

import (
    "context"
    "fmt"
)

type keyOne string
type keyTwo string

var ctx = context.Background()

func main() {
    
    ctx = context.WithValue(ctx, keyOne("one"), "valueOne")
    ctx = context.WithValue(ctx, keyTwo("one"), "valueTwo")
    ctx = context.WithValue(ctx, "ssss", "sssssssvalueTwo")  
    mm(ctx)
    mms(ctx)
    
}

func mms(ctx context.Context){
    fmt.Println(ctx.Value("ssss"))     
}


func mm(ctx context.Context){
      fmt.Println(ctx.Value(keyOne("one")).(string))
      fmt.Println(ctx.Value(keyTwo("one")).(string))
      fmt.Println(ctx.Value("ssss"))
}
```

