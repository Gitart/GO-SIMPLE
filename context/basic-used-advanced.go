package main

import (
    "context"
    "time"
)

type MyContext struct{
     Ctx context.Context    
     Title string 
}

func main() {

    // создаём контекст с функцией завершения
    ctx, cancel := context.WithCancel(context.Background())

    d:=MyContext{
        ctx, "Работает нулевая рутина ============================ ",
    }


    go  taskN(d)

    // запускаем нашу первую горутину
    go task(ctx, "Работает первая  рутина")
    
    // запускаем нашу вторую горутину
    go task2(ctx, "Работает вторая рутина")

    // делаем паузу, чтобы дать горутине поработать
    time.Sleep(2 * time.Minute)
    
    // завершаем контекст, чтобы завершить горутину
    cancel()
}

// Task#0
func taskN(ctx MyContext ) {
   
   // запускаем бесконечный цикл
    for {
        select {

        // проверяем не завершён ли ещё контекст и выходим, если завершён
        case <-ctx.Ctx.Done():
            return

        // выполняем нужный нам код
        default:
            println("---------------------------", ctx.Title)
        }

        // делаем паузу перед следующей итерацией
        time.Sleep(time.Second * 1)
    }
}


// Task#1
func task(ctx context.Context, txt string) {

    // запускаем бесконечный цикл
    for {
        select {

        // проверяем не завершён ли ещё контекст и выходим, если завершён
        case <-ctx.Done():
            return

        // выполняем нужный нам код
        default:
            println("-------------", txt)
        }

        // делаем паузу перед следующей итерацией
        time.Sleep(time.Second * 1)
    }
}



// Task#2
func task2(ctx context.Context, txt string) {

    // запускаем бесконечный цикл
    for {
        select {

        // проверяем не завершён ли ещё контекст и выходим, если завершён
        case <-ctx.Done():
            return

        // выполняем нужный нам код
        default:
            println("*******************")
            println(txt)
        }

        // делаем паузу перед следующей итерацией
         time.Sleep(time.Second * 10)

    }
}
