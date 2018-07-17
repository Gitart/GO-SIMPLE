## Поддержка сервера на плаву
Добавим в этот пример еще один обработчик верхнего уровня, который будет обрабатывать ошибки хэндлеров и поддерживать сервер на плаву:


```golang
 func recoverHandler(next http.Handler) http.Handler {
     fn := func(w http.ResponseWriter, r *http.Request) {
     defer func() {
       if err := recover(); err != nil {
         log.Printf("panic: %+v", err)
         http.Error(w, http.StatusText(500), 500)
       }
     }()
 
     next.ServeHTTP(w, r)
   }
 
   return http.HandlerFunc(fn)
 }
 ```
 
Главная функция будет выглядеть так:

```golang
 func main() {
   commonHandlers := New(loggingHandler, recoverHandler)
   http.Handle("/about/", commonHandlers.ThenFunc(aboutHandler))
   http.Handle("/", commonHandlers.ThenFunc(indexHandler))
   http.ListenAndServe(":8000", nil)
 }
 ```
