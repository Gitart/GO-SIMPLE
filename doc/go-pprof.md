### [Пакет pprof в Golang](https://golang-blog.blogspot.com/2020/08/package-pprof-in-golang.html)

Пакет **pprof** обслуживает данные профилирования среды выполнения HTTP\-сервера в формате, ожидаемом инструментом визуализации pprof.

Пакет обычно импортируется только из\-за побочного эффекта регистрации его обработчиков HTTP. Все обрабатываемые пути начинаются с /debug/pprof/.

Чтобы использовать pprof, свяжите этот пакет со своей программой:

```
import _ "net/http/pprof"

```

Если ваше приложение еще не запустило http\-сервер, вам необходимо запустить его. Добавьте "net/http" и "log" к вашему импорту и следующий код к вашей main функции:

```
go func() {
    log.Println(http.ListenAndServe("localhost:6060", nil))
}()

```

Если вы не используете DefaultServeMux, вам нужно будет зарегистрировать обработчики в используемом мультиплексоре (mux). Например:

```
package main

import (
    "log"
    "net/http"
    "net/http/pprof"
)

func main() {
    mux := http.NewServeMux()

    // Здесь можно указать свой путь
    mux.HandleFunc("/debug/pprof/profile", pprof.Profile)

    log.Fatal(http.ListenAndServe(":6060", mux))
}

```

Затем используйте инструмент pprof, чтобы посмотреть профиль кучи (heap profile):

```
go tool pprof http://localhost:6060/debug/pprof/heap

```

Или посмотреть 30\-секундный профиль процессора:

```
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

```

Или посмотреть профиль блокировки goroutine после вызова runtime.SetBlockProfileRate в вашей программе:

```
go tool pprof http://localhost:6060/debug/pprof/block

```

Или посмотреть на держателей конкурирующих мьютексов после вызова runtime.SetMutexProfileFraction в вашей программе:

```
go tool pprof http://localhost:6060/debug/pprof/mutex

```

Пакет также экспортирует обработчик, который обслуживает данные трассировки выполнения для команды "go tool trace". Чтобы собрать 5\-секундную трассировку выполнения:

```
wget -O trace.out http://localhost:6060/debug/pprof/trace?seconds=5
go tool trace trace.out

```

Чтобы просмотреть все доступные профили, откройте в браузере http://localhost:6060/debug/pprof/.

#### Функция Cmdline

```
func Cmdline(w http.ResponseWriter, r *http.Request)

```

Cmdline отвечает командной строкой запущенной программы с аргументами, разделенными NUL байтами. Инициализация пакета регистрирует его как /debug/pprof/cmdline.

#### Функция Handler

```
func Handler(name string) http.Handler

```

Handler возвращает обработчик HTTP, который обслуживает именованный профиль.

#### Функция Index

```
func Index(w http.ResponseWriter, r *http.Request)

```

Index отвечает профилем в формате pprof, указанным в запросе. Например, "/debug/pprof/heap" обслуживает профиль "heap". Index отвечает на запрос "/debug/pprof/" HTML\-страницей со списком доступных профилей.

#### Функция Profile

```
func Profile(w http.ResponseWriter, r *http.Request)

```

Profile отвечает профилем процессора в формате pprof. Профилирование длится в течение времени, указанного в параметре GET в секундах, или 30 секунд, если не указано иное. Инициализация пакета регистрирует его как /debug/pprof/profile.

#### Функция Symbol

```
func Symbol(w http.ResponseWriter, r *http.Request)

```

Symbol ищет программные счетчики, перечисленные в запросе, отвечая таблицей, отображающей программные счетчики на имена функций. Инициализация пакета регистрирует его как /debug/pprof/symbol.

#### Функция Trace

```
func Trace(w http.ResponseWriter, r *http.Request)

```

Trace отвечает трассировкой выполнения в двоичной форме. Трассировка длится в течение времени, указанного в параметре GET в секундах, или в течение 1 секунды, если не указано иное. Инициализация пакета регистрирует его как /debug/pprof/trace.
