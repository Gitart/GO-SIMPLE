# Система пакаджей Go

Posted on 2019, Jul 26 5 мин. чтения

Когда пишешь на го, что такое пакадж уже интуитивно понятно, но попробуем все же сформулировать - что такое пакадж?

По [спецификации](https://golang.org/ref/spec#Packages) пакадж у нас - это набор `.go` файлов в одной директории, где каждый файл помечен выражением `package <название_пакаджа>` и `<название_пакаджа>` одинаково для всех файлов.

Между собой пакаджи мы подключаем директивой `import` и здесь уже используются не названия пакаджей, а пути к ним. Пути в импорте указываются относительные и точка отсчета может меняться.

Для пакаджей, которые входят в стандартную библиотеку точка отсчета `GOROOT/src`, для всех сторонних пакаджей - `GOPATH/src`.

В связке с пакаджами работает тулза `go get`, благодаря которой можно устанавливать пакадж из удаленного [репозитория](https://github.com/golang/go/wiki/GoGetTools).

Например команда

```
go get github.com/gorilla/mux

```

установит пакадж `mux` из репозитория `https://github.com/gorilla/mux` в директорию `GOPATH/src/github.com/gorilla/mux` и дальше в коде подключаем как

```
import github.com/gorilla/mux

```

Проблема с `go get` что она умеет “проматывать” версии зависимостей только вперед, обновляя до последней версии.

При работе со стандартной библиотекой, локальными пакаджами и командой из одного человека это может сработать. Во всех остальных случаях возникает вопрос с версионостью пакаджей и воспроизводимостью сборки проекта.

Например у нас 1 разработчик на проекте `bar`, в проекте используется `github.com/gorilla/mux`. Кодовая база ширится и в помощь приходит 2-ой разработчик, он успешно забирает проект и зависимости с помощью `go get`. Но мы легко можем представить, что в это время вендор пакаджа `github.com/gorilla/mux` выпустил новую версию и второй разработчик уже будет использовать эту свежую версию. И кто его знает, как изменится поведение разрабатываемого проекта `bar` на второй машине. Еще веселее если мы представим, что собираем релиз для продакшена и забираем зависимости с помощью `go get`.

Но с версии go1.5 у нас появился быстрофикс в виде директории `vendor`, который разруливает эти ситуации.

```bash
project-bar
├── ...
├── ...
└── vendor
```

Теперь при сборке `go build`, сначала ищет зависимости в директории проекта `vendor`, а потом уже в `GOPATH`. Т.е. мы можем отдельно в вендоре сохранить все свои зависимости с необходимыми версиями и тем самым добиться воспроизводимого билда на другой машине.

Но ручное управление зависимостями чревато ошибками, хотелось бы пакетный менеджер, который все это контролирует.

Мне для этих целей нравится [Dep](https://golang.github.io/dep/). Как и ряд других инструментов, он позволяет выполнять все необходимые действия с зависимостями - добавить, обновить, откатить, удалить. Позволяет указывать версии зависимостей по комиту, по тегу, добавлять правила зависимостей (к примеру - версия не ниже такой-то).

(До `go1.5` [пакетные менеджеры](https://github.com/golang/go/wiki/PackageManagementTools) использовали переопределение `GOPATH` для работы с зависимостями.)

## Go Modules

С версии 1.11 `go` получил встроенную подержку зависимостей на основе модулей ([Go modules](https://github.com/golang/go/wiki/Modules)). В режиме модуля, исходники проекта теперь не обязаны располагаться в директории `GOPATH/src`, а для зависимостей можно указывать версии.

Модулем считается коллекция пакаджей собранная в одной директории, с файлом `go.mod` в корне. Файл `go.mod` определяет путь к модулю, который используется с `import` директивой в проектах. Так же `go.mod` определяет зависимости модуля с указанием версий.

```bash
<диретория_модуля>
├── <внутренний_пакадж_модуля_1>
├── <внутренний_пакадж_модуля_2>
├── <...>
├── <исходный_файл_1>
├── <исходный_файл_1>
├── <...>
└── go.mod
```

В go 1.11 поддержка модулей реализована в форме эксперимента и включается, если проект расположен вне $GOPATH/src и в корне размещен файл go.mod, либо через переменную среды `GO111MODULE` (см. [доки](https://golang.org/cmd/go/#hdr-Preliminary_module_support)).

Например выполнить команду в корне проекта для сборки в режиме модулей:

```bash
env GO111MODULE=on go build ./...
```

Рассмотрим на примере работу с модулями.

Создадим тестовый проект `project-bar` и расположим его не в `GOPATH`, а например, в домашнем каталоге. Добавим исходик `hello.go`

```go
package projectBar

func Hello() string {
	return "Hello, world."
}
```

и инициализируем проект для работы с модулями

```bash
$go mod init example.com/project-bar

go: creating new go.mod: module example.com/project-bar
```

добавим тест `hello_test.go`

```go
package projectBar

import "testing"

func TestHello(t *testing.T) {
    want := "Hello, world."
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
```

и выполним его

```bash
$go test
PASS
ok      example.com/project-bar 0.001s
```

Наш пакадж собрался не находясь в GOPATH, отлично!

Давайте добавим теперь метод, который будет поднимать сервер и в ответ на роут `/hello` возвращать `Hello, world.`

```go
package projectBar

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Hello() string {
	return "Hello, world."
}

func HelloServer() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler)
	http.Handle("/", r)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, Hello())
}
```

в коде мы использовали стороннюю библиотеку `github.com/gorilla/mux` и соответсвенно добавили ее в импорт.

Теперь попробуем собрать наш пакадж

```bash
go build
go: finding github.com/gorilla/mux v1.7.3
go: downloading github.com/gorilla/mux v1.7.3
```

тулчейн нашел и забрал необходимый пакадж, а go.mod обновился требованиями к зависимостям

```bash
module example.com/project-bar

require github.com/gorilla/mux v1.7.3
```

Давайте попробуем даунгрейднуть зависимость

```bash
$ go get github.com/gorilla/mux@'<v1.6'
go: finding github.com/gorilla/mux v1.5.0
go: downloading github.com/gorilla/mux v1.5.0
```

и наш go.mod теперь

```bash
module example.com/project-bar

require github.com/gorilla/mux v1.5.0
```

На мой взгляд, модули с успехом заменяют менеджеры зависимостей, которые работают через `vendor`.

Но в тоже время время можно представить ситуацию, когда каталог вендоров полезен. Например мы используем определенную библиотеку, но в какой-то момент автор убрал ее из открытого доступа. Соотвественно, если мы забираем зависимости только с удаленного источника, наш проект сломается.

Каталог с вендорами гарантирует даже в этой ситуации повторяемость билда. И тулчейн поддерживает такую работу с модулями.

Продолжим наш пример с `project_bar` и скинем зависимоси в вендоры:

```bash
$go mod vendor
go: downloading github.com/gorilla/context v1.1.1
```

в корень проекта добавилась директория `vendor`

```bash
.
├── go.mod
├── go.sum
├── hello.go
├── hello_test.go
└── vendor
    ├── github.com
    │   └── gorilla
    │       ├── context
    │       └── mux
    └── modules.txt
```

собраться с вендорами

```bash
$go build -mod=vendor
```

Полезные ссылки:

*   [Спецификация на пакаджи go](https://golang.org/ref/spec#Packages)
*   [Организация кода для go](https://golang.org/doc/code.html#Workspaces)
*   [GOPATH](https://github.com/golang/go/wiki/GOPATH)
*   [Vendors](https://codeengineered.com/blog/2015/go-1.5-vendor-handling/)
*   [Dep](https://golang.github.io/dep/)
*   [https://github.com/golang/go/wiki/PackageManagementTools](https://github.com/golang/go/wiki/PackageManagementTools)
*   [https://github.com/golang/go/wiki/Modules](https://github.com/golang/go/wiki/Modules)
*   [https://blog.golang.org/modules2019](https://blog.golang.org/modules2019)
*   [https://blog.golang.org/using-go-modules](https://blog.golang.org/using-go-modules)
*   [Vendors + modules](https://github.com/golang/go/wiki/Modules#how-do-i-use-vendoring-with-modules-is-vendoring-going-away)
