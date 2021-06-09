Декораторы, безусловно, более заметны в других языках программирования, таких как Python и TypeScript, но это не значит, что вы не можете использовать их в Go. Фактически, для некоторых проблем использование декораторов - идеальное решение, как мы, надеюсь, узнаем в этом руководстве.

## Понимание шаблона декоратора

> **По** сути, **декораторы** позволяют вам обернуть существующую функциональность и добавить или добавить свои собственные пользовательские функции поверх.

В Go функции считаются объектами первого класса, что по сути означает, что вы можете передавать их так же, как и переменную. Давайте посмотрим на это в действии на очень простом примере:

```go
package main

import (
  "fmt"
  "time"
)

func myFunc() {
  fmt.Println("Hello World")
  time.Sleep(1 * time.Second)
}

func main() {
  fmt.Printf("Type: %T\n", myFunc)
}

```

Итак, в этом примере мы определили функцию с именем `myFunc` , которая просто распечатывает `Hello World` . Однако в теле нашей `main()` функции мы вызвали `fmt.Printf` и использовали `%T` для вывода типа значения, которое мы передаем в качестве второго аргумента. В этом случае мы передаем, `myFunc` который впоследствии распечатает следующее:

```s
$ go run test.go
Type: func()

```

Итак, что это значит для нас, разработчиков Go? Что ж, это подчеркивает тот факт, что **функции могут передаваться и использоваться в качестве аргументов** в других частях нашей кодовой базы.

Давайте посмотрим на это в действии, немного расширив нашу кодовую базу и добавив `coolFunc()` функцию, которая принимает функцию как единственный параметр:

```go
package main

import (
  "fmt"
  "time"
)

func myFunc() {
  fmt.Println("Hello World")
  time.Sleep(1 * time.Second)
}

// coolFunc takes in a function
// as a parameter
func coolFunc(a func()) {
    // it then immediately calls that functino
  a()
}

func main() {
  fmt.Printf("Type: %T\n", myFunc)
  // here we call our coolFunc function
  // passing in myFunc
    coolFunc(myFunc)
}

```

Когда мы попытаемся запустить это, мы должны увидеть, что наш новый вывод содержит нашу `Hello World` строку, как мы и ожидали:

```s
$ go run test.go
Type: func()
Hello World

```

Поначалу это может показаться вам немного странным. Зачем вам нужно делать что-то подобное? По сути, это добавляет уровень абстракции к вашему вызову `myFunc` и усложняет код, не добавляя особой ценности.

## Простой декоратор

Давайте посмотрим, как мы можем использовать этот шаблон, чтобы добавить некоторую ценность нашей кодовой базе. При желании мы могли бы добавить дополнительное ведение журнала выполнения определенной функции, чтобы выделить время ее начала и окончания.

```go
package main

import (
    "fmt"
    "time"
)

func myFunc() {
  fmt.Println("Hello World")
    time.Sleep(1 * time.Second)
}

func coolFunc(a func()) {
    fmt.Printf("Starting function execution: %s\n", time.Now())
    a()
    fmt.Printf("End of function execution: %s\n", time.Now())
}

func main() {
    fmt.Printf("Type: %T\n", myFunc)
    coolFunc(myFunc)
}

```

После этого вы должны увидеть журналы, которые выглядят примерно так:

```s
$ go run test.go
Type: func()
Starting function execution: 2018-10-21 11:11:25.011873 +0100 BST m=+0.000443306
Hello World
End of function execution: 2018-10-21 11:11:26.015176 +0100 BST m=+1.003743698

```

Как видите, мы смогли эффективно обернуть мою исходную функцию, не внося изменений в ее реализацию. Теперь мы можем четко видеть, когда эта функция была запущена и когда она завершила выполнение, и это подчеркивает, что функция завершает выполнение примерно за секунду.

## Примеры из реального мира

Давайте рассмотрим еще несколько примеров того, как мы можем использовать декораторы для дальнейшей славы и богатства. Мы возьмем действительно простой веб-сервер http и украсим наши конечные точки, чтобы мы могли проверить, имеет ли входящий запрос определенный заголовок.

> **Если вы хотите** узнать больше о написании простого REST API в Go, я рекомендую проверить мою другую статью здесь: [Создание REST API в Go](https://tutorialedge.net/golang/creating-restful-api-with-golang/)

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: homePage")
    fmt.Fprintf(w, "Welcome to the HomePage!")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
    handleRequests()
}

```

Как видите, ничего особо сложного в нашем коде нет. Мы настраиваем `net/http` маршрутизатор, который обслуживает одну `/` конечную точку.

Давайте добавим действительно простую функцию декоратора аутентификации, которая будет проверять, установлен ли `Authorized` заголовок `true` во входящем запросе.

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        fmt.Println("Checking to see if Authorized header set...")

        if val, ok := r.Header["Authorized"]; ok {
            fmt.Println(val)
            if val[0] == "true" {
                fmt.Println("Header is set! We can serve content!")
                endpoint(w, r)
            }
        } else {
            fmt.Println("Not Authorized!!")
            fmt.Fprintf(w, "Not Authorized!!")
        }
    })
}

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: homePage")
    fmt.Fprintf(w, "Welcome to the HomePage!")
}

func handleRequests() {

    http.Handle("/", isAuthorized(homePage))
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
    handleRequests()
}

```

> **Примечание.** Это абсолютно неправильный способ защиты вашего REST API. Я бы рекомендовал использовать JWT или OAuth2 для достижения этой цели!

Итак, давайте разберемся с этим и попытаемся понять, что происходит!

Мы создали новую функцию-декоратор, `isAuthorized()` которая принимает функцию, которая соответствует той же сигнатуре, что и наша исходная `homePage` функция. Затем это возвращает файл `http.Handler` .

В теле нашей `isAuthorized()` функции мы возвращаем new, `http.HandlerFunc` внутри которого выполняется проверка того, что наш `Authorized` заголовок установлен и равен `true` . Это значительно упрощенная версия `OAuth2` аутентификации / авторизации. Есть несколько небольших несоответствий, но она дает вам общее представление о том, как это будет работать.

Однако **важно отметить** тот факт, что нам удалось украсить существующую конечную точку и добавить некоторую форму аутентификации вокруг указанной конечной точки без необходимости изменять существующую реализацию этой функции.

Теперь, если бы мы добавили новую конечную точку, которую мы хотели защитить, мы могли бы легко это сделать:

```go
// define our newEndpoint function. Notice how, yet again,
// we don't do any authentication based stuff in the body
// of this function
func newEndpoint(w http.ResponseWriter, r *http.Request) {
    fmt.Println("My New Endpoint")
    fmt.Fprintf(w, "My second endpoint")
}

func handleRequests() {

    http.Handle("/", isAuthorized(homePage))
  // register our /new endpoint and decorate our
  // function with our isAuthorized Decorator
  http.Handle("/new", isAuthorized(newEndpoint))
    log.Fatal(http.ListenAndServe(":8081", nil))
}

```

Это подчеркивает ключевые преимущества шаблона декоратора, при котором обернуть код в нашу кодовую базу невероятно просто. Мы можем легко добавить новые аутентифицированные конечные точки, используя тот же метод.

## Заключение

Надеюсь, это руководство помогло демистифицировать чудеса декоратора и то, как вы можете использовать шаблон декоратора в ваших собственных программах на основе Go. Мы узнали о преимуществах шаблона декоратора и о том, как мы можем использовать его, чтобы дополнить существующие функциональные возможности новыми.

Во второй части руководства мы рассмотрели более реалистичный пример того, как вы потенциально можете использовать это в своих собственных Go-системах производственного уровня.

Если вам понравился этот урок, то, пожалуйста, не стесняйтесь делиться статьей, это действительно помогает сайту, и я буду очень признателен! Если у вас есть какие-либо вопросы и / или комментарии, дайте мне знать в разделе комментариев ниже!
