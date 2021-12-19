// https://tproger.ru/video/uchimsja-razrabatyvat-na-golang-urok-14-rabota-s-context/

package main

import "fmt"
import "time"
import "context"



func main() {

    // Ожидание процесса
    dur := time.Millisecond*1000
    ctx := context.Background() 
    ctx, cancel :=context.WithTimeout(ctx, dur)
    defer cancel()

    doReguest(ctx, "Ok")
}


// *************************************
// Do thamting
// *************************************
func doReguest(ctx context.Context, key string) {

// Do thamthing ...

select {
        // Ожидание после котрого выход
        case <-time.After(1500*time.Millisecond):
            fmt.Println("timeout отмена")
            return
        case<- ctx.Done():
            fmt.Println("done ")
} 

// time.Sleep(time.Millisecond*1500)


fmt.Println("KEY :",key)

}
