
# Пример реализации простой авторизации по паролю и сохранении сессии на Go.
----
Простейший пример с пакетами github.com/gorilla/mux и github.com/gorilla/sessions. 
Сначала / выдаст вам "you are anonymous", а после посещения /login выдаст "you are user". 


```golang
var cookieStore = sessions.NewCookieStore([]byte("secret"))

const cookieName = "MyCookie"

type sesKey int

const (
    sesKeyLogin sesKey = iota
)

func main() {
    gob.Register(sesKey(0))

    router := mux.NewRouter()
    router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        ses, err := cookieStore.Get(r, cookieName)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        ses.Values[sesKeyLogin] = "user"
        err = cookieStore.Save(r, w, ses)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
    })
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        ses, err := cookieStore.Get(r, cookieName)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        login, ok := ses.Values[sesKeyLogin].(string)
        if !ok {
            login = "anonymous"
        }

        w.Write([]byte("you are " + login))
    })

    http.ListenAndServe(":3000", router)
}
```
