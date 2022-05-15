 ### Подробное объяснение использования контекста 


На сервере go запрос для каждого запроса выполняется в отдельной горутине. Обработка запроса может также 
спроектировать взаимодействие между несколькими горутинами. Использование контекста может сделать удобной для 
разработчиков передачу запросов в этих горутинах. Данные, отмените сигнал или крайний срок горутины.
Структура контекста
```go
    // A Context carries a deadline, cancelation signal, and request-scoped values
    // across API boundaries. Its methods are safe for simultaneous use by multiple
    // goroutines.
    type Context interface {
        // Done returns a channel that is closed when this Context is canceled
        // or times out.
        Done() <-chan struct{}
     
        // Err indicates why this context was canceled, after the Done channel
        // is closed.
        Err() error
     
        // Deadline returns the time when this Context will be canceled, if any.
        Deadline() (deadline time.Time, ok bool)
     
        // Value returns the value associated with key or nil if none.
        Value(key interface{}) interface{}
    }
```

Done Метод возвращает закрытый канал, когда контекст отменен или истекло время ожидания. Закрытый канал можно использовать в качестве широковещательного уведомления, чтобы сообщить функциям, связанным с контекстом, чтобы они остановили текущую работу и затем вернулись.

Когда родительская операция запускает горутину для дочерних операций, эти дочерние операции не могут отменить родительскую операцию. Функция WithCancel, описанная ниже, позволяет отменить вновь созданный контекст.

Контекст может безопасно использоваться несколькими горутинами. Разработчики могут передать контекст любому количеству горутин, а затем все горутины могут быть уведомлены об отмене контекста.

ErrМетод возвращает причину отмены контекста.

DeadlineВернитесь, когда контекст истечет.

ValueВозвращает данные, связанные с контекстом.
Унаследованный контекст
## BackGround
```go
    // Background returns an empty Context. It is never canceled, has no deadline,
    // and has no values. Background is typically used in main, init, and tests,
    // and as the top-level Context for incoming requests.
    func Background() Context
```
BackGound является корнем всех контекстов и не может быть отменен.
## WithCancel
```go
    // WithCancel returns a copy of parent whose Done channel is closed as soon as
    // parent.Done is closed or cancel is called.
    func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
```
WithCancel возвращает унаследованный контекст. Этот контекст закрывает свой канал Done, когда Done родительского контекста закрывается, или закрывает свой Done, когда он отменяется.
WithCancel также возвращает функцию отмены cancel, которая используется для отмены текущего контекста.

видеоAdvanced Go Concurrency PatternsЯ изменил первый пример кода WithCancel в
```go
    package main
     
    import (
        "context"
        "log"
        "os"
        "time"
    )
     
    var logg *log.Logger
     
    func someHandler() {
        ctx, cancel := context.WithCancel(context.Background())
        go doStuff(ctx)
     
    // Отмена doStuff через 10 секунд
        time.Sleep(10 * time.Second)
        cancel()
     
    }
     
     // Работаем каждую секунду, и в то же время он будет определять, был ли отменен ctx, и если да, выходим
    func doStuff(ctx context.Context) {
        for {
            time.Sleep(1 * time.Second)
            select {
            case <-ctx.Done():
                logg.Printf("done")
                return
            default:
                logg.Printf("work")
            }
        }
    }
     
    func main() {
        logg = log.New(os.Stdout, "", log.Ltime)
        someHandler()
        logg.Printf("down")
    }
   ```  
     

результат
```
    E:\wdy\goproject>go run context_learn.go
    15:06:44 work
    15:06:45 work
    15:06:46 work
    15:06:47 work
    15:06:48 work
    15:06:49 work
    15:06:50 work
    15:06:51 work
    15:06:52 work
    15:06:53 down
```
## withDeadline withTimeout
```go
    WithTimeout func(parent Context, timeout time.Duration) (Context, CancelFunc)
    WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).
```
WithTimeout эквивалентен WithDeadline (parent, time.Now (). Add (timeout)).

Измените приведенный выше пример кода
```go
     
    func timeoutHandler() {
        // ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
        // go doTimeOutStuff(ctx)
        go doStuff(ctx)
     
        time.Sleep(10 * time.Second)
     
        cancel()
     
    }
     
    func main() {
        logg = log.New(os.Stdout, "", log.Ltime)
        timeoutHandler()
        logg.Printf("end")
    }
 ```    

Выход
```
    15:59:22 work
    15:59:24 work
    15:59:25 work
    15:59:26 work
    15:59:27 done
    15:59:31 end
```
Вы можете видеть, что doStuff был отменен, когда истекло время ожидания контекста, а ctx.Done () был закрыт.
замените context.WithDeadline на context.WithTimeout
```go
    func timeoutHandler() {
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        // ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
        // go doTimeOutStuff(ctx)
        go doStuff(ctx)
     
        time.Sleep(10 * time.Second)
     
        cancel()
    }
```
Выход
```
    16:02:47 work
    16:02:49 work
    16:02:50 work
    16:02:51 work
    16:02:52 done
    16:02:56 end
```
По видео Advanced Go Concurrency PatternsНапишите код через 5 минут 48 секунд doTimeOutStuff замените doStuff
     
```go     
    func doTimeOutStuff(ctx context.Context) {
        for {
            time.Sleep(1 * time.Second)
     
                     если крайний срок, ok: = ctx.Deadline (); ok {// Deadl установлен
                logg.Printf("deadline set")
                if time.Now().After(deadline) {
                    logg.Printf(ctx.Err().Error())
                    return
                }
     
            }
     
            select {
            case <-ctx.Done():
                logg.Printf("done")
                return
            default:
                logg.Printf("work")
            }
        }
    }
     
    func timeoutHandler() {
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        // ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
        go doTimeOutStuff(ctx)
        // go doStuff(ctx)
     
        time.Sleep(10 * time.Second)
     
        cancel()
     
    }
```
Выход:
```
    16:03:55 deadline set
    16:03:55 work
    16:03:56 deadline set
    16:03:56 work
    16:03:57 deadline set
    16:03:57 work
    16:03:58 deadline set
    16:03:58 work
    16:03:59 deadline set
    16:03:59 context deadline exceeded
    16:04:04 end
```
## context deadline exceeded
Это сообщение об ошибке ctx.Err при истечении времени ожидания ctx.
Поисковая тестовая программа

См. Полный код в официальной документацииGo Concurrency Patterns: Context, Ключом является функция httpDo
```go
    func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
        // Run the HTTP request in a goroutine and pass the response to f.
        tr := &http.Transport{}
        client := &http.Client{Transport: tr}
        c := make(chan error, 1)
        go func() { c <- f(client.Do(req)) }()
        select {
        case <-ctx.Done():
            tr.CancelRequest(req)
            <-c // Wait for f to return.
            return ctx.Err()
        case err := <-c:
            return err
        }
    }
```
Ключевым моментом httpDo является
```go
        select {
        case <-ctx.Done():
            tr.CancelRequest(req)
            <-c // Wait for f to return.
            return ctx.Err()
        case err := <-c:
            return err
        }
```
Либо ctx был отменен, либо произошла ошибка запроса.
## WithValue
```go
func WithValue(parent Context, key interface{}, val interface{}) Context
```

Посмотреть программу поискаuseripКод в
Код клавиши следующий:
```go
    // NewContext returns a new Context carrying userIP.
    func NewContext(ctx context.Context, userIP net.IP) context.Context {
        return context.WithValue(ctx, userIPKey, userIP)
    }
     
    // FromContext extracts the user IP address from ctx, if present.
    func FromContext(ctx context.Context) (net.IP, bool) {
        // ctx.Value returns nil if ctx has no value for the key;
        // the net.IP type assertion returns ok=false for nil.
        userIP, ok := ctx.Value(userIPKey).(net.IP)
        return userIP, ok
    }
```
## Информация в go doc

    The WithCancel, WithDeadline, and WithTimeout functions take a Context (the
    parent) and return a derived Context (the child) and a CancelFunc. Calling
    the CancelFunc cancels the child and its children, removes the parent's
    reference to the child, and stops any associated timers.

Следует отметить, что вызов CancelFunc отменяет дочерний элемент и контекст, сгенерированный дочерним элементом,
удаляет ссылку родительского контекста на этот дочерний элемент и останавливает связанный счетчик.
