JWT или веб-токены JSON, как они более формально известны, представляют собой компактное, безопасное для URL-адресов средство представления требований, передаваемых между двумя сторонами. По сути, это сбивающий с толку способ сказать, что JWT позволяют передавать информацию от клиента на сервер без сохранения состояния, но безопасным способом.

## Предпосылки

Прежде чем вы сможете следовать этой статье, вам понадобится следующее:

*   Вам понадобится Go версии 1.11+, установленной на вашем компьютере для разработки.

## Вступление

Стандарт JWT использует либо секрет, используя алгоритм HMAC, либо пару открытого / закрытого ключей, используя RSA или ECDSA.

> **Примечание.** Если вас интересует формальное определение того, что такое JWT, я рекомендую ознакомиться с RFC: [RFC-7519.](https://tools.ietf.org/html/rfc7519)

Они широко используются в одностраничных приложениях (SPA) в качестве средств безопасной связи, поскольку позволяют нам делать две ключевые вещи:

*   **Аутентификация** \- наиболее часто используемая практика. Когда пользователь входит в ваше приложение или каким-либо образом аутентифицируется, каждый запрос, который затем отправляется клиентом от имени пользователя, будет содержать JWT.
*   **Обмен информацией** . Второе применение JWT - безопасная передача информации между различными системами. Эти JWT могут быть подписаны с использованием пар открытого / закрытого ключей, поэтому вы можете безопасно проверить каждую систему в этой транзакции, а JWT содержат механизм защиты от несанкционированного доступа, поскольку они подписываются на основе заголовка и полезной нагрузки.

Итак, если вы еще не догадались, в этом руководстве мы рассмотрим, что именно нужно для создания безопасного REST API на основе Go, использующего для связи веб-токены JSON!

## Простой REST API

Итак, мы собираемся использовать код из одной из моих других статей « [Создание простого REST API в Go»](https://tutorialedge.net/golang/creating-restful-api-with-golang/) , чтобы начать работу. Это будет иметь действительно простую `Hello World` конечную точку, и она будет работать на порту 8081.

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
    handleRequests()
}

```

Когда мы запускаем эту попытку перейти на нашу домашнюю страницу `http://localhost:8081/` , мы должны увидеть сообщение `Hello World` в нашем браузере.

## JWT аутентификация

Итак, теперь, когда у нас есть простой API, который мы можем защитить с помощью подписанных токенов JWT, давайте создадим клиентский API, который будет пытаться запрашивать данные из этого исходного API.

Для этого мы можем использовать JWT, подписанный защищенным ключом, о котором и наш клиент, и сервер не будут знать. Давайте рассмотрим, как это будет работать:

1.  Наш клиент сгенерирует подписанный JWT на основе нашей общей парольной фразы.
2.  Когда наш клиент обращается к нашему серверному API, он включает этот JWT как часть запроса.
3.  Наш сервер сможет прочитать этот JWT и проверить токен, используя ту же парольную фразу.
4.  Если JWT действителен, он вернет строго конфиденциальное `hello world` сообщение обратно клиенту, в противном случае оно вернется `not authorized` .

Наша архитектурная диаграмма будет выглядеть примерно так:

![диаграмма архитектуры](https://images.tutorialedge.net/images/golang/go-jwt-tutorial/diagram-01.png)

### Наш Сервер

Итак, давайте посмотрим на это в действии, давайте создадим действительно простой сервер:

```go
package main

import (
    "fmt"
    "log"
    "net/http"

    jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
    fmt.Println("Endpoint Hit: homePage")

}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        if r.Header["Token"] != nil {

            token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
                if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                    return nil, fmt.Errorf("There was an error")
                }
                return mySigningKey, nil
            })

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            if token.Valid {
                endpoint(w, r)
            }
        } else {

            fmt.Fprintf(w, "Not Authorized")
        }
    })
}

func handleRequests() {
    http.Handle("/", isAuthorized(homePage))
    log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
    handleRequests()
}

```

Давайте разберемся с этим. Мы создали действительно простой API с единственной конечной точкой, которая защищена нашим `isAuthorized` декоратором промежуточного программного обеспечения. В этой `isAuthorized` функции мы проверяем, содержит ли входящий запрос `Token` заголовок в запросе, а затем проверяем, действителен ли токен на основе нашего частного `mySigningKey` .

Если это действительный токен, мы обслуживаем защищенную конечную точку.

> **Примечание. В** этом примере используются декораторы. Если вам не нравится концепция декораторов в Go, я рекомендую вам ознакомиться с другой моей статьей здесь: [Начало работы с декораторами в Go.](https://tutorialedge.net/golang/go-decorator-function-pattern-tutorial/)

### Наш клиент

Теперь, когда у нас есть сервер с защищенной конечной точкой JWT, давайте создадим что-то, что может с ним взаимодействовать.

Мы создадим простое клиентское приложение, которое попытается вызвать нашу `/` конечную точку нашего сервера.

```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "time"

    jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func homePage(w http.ResponseWriter, r *http.Request) {
    validToken, err := GenerateJWT()
    if err != nil {
        fmt.Println("Failed to generate token")
    }

    client := &http.Client{}
    req, _ := http.NewRequest("GET", "http://localhost:9000/", nil)
    req.Header.Set("Token", validToken)
    res, err := client.Do(req)
    if err != nil {
        fmt.Fprintf(w, "Error: %s", err.Error())
    }

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Fprintf(w, string(body))
}

func GenerateJWT() (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)

    claims := token.Claims.(jwt.MapClaims)

    claims["authorized"] = true
    claims["client"] = "Elliot Forbes"
    claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

    tokenString, err := token.SignedString(mySigningKey)

    if err != nil {
        fmt.Errorf("Something Went Wrong: %s", err.Error())
        return "", err
    }

    return tokenString, nil
}

func handleRequests() {
    http.HandleFunc("/", homePage)

    log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
    handleRequests()
}

```

Давайте разберемся, что происходит в приведенном выше коде. Опять же, мы определили действительно простой API с единственной конечной точкой. Эта конечная точка при срабатывании генерирует новый JWT, используя нашу безопасность `mySigningKey` , затем создает новый http-клиент и устанавливает `Token` заголовок, равный только что сгенерированной нами строке JWT.

Затем он пытается поразить наше `server` приложение, которое работает `http://localhost:9000` с этим подписанным токеном JWT. Затем наш сервер проверяет токен, который мы создали на клиенте, и отправляет нам наше сверхсекретное `Hello World` сообщение.

## Заключение

Надеюсь, этот учебник помог демистифицировать искусство защиты ваших приложений Go и REST API с помощью веб-токенов JSON. Было очень весело писать эту статью, и я надеюсь, что она помогла вам в ваших путешествиях по разработке Go.

Если вам понравился этот урок, дайте мне знать в разделе комментариев ниже или поделитесь этой статьей в социальных сетях, это действительно помогает мне и моему сайту!

> **Примечание.** Если вы хотите отслеживать, когда на сайте публикуются новые статьи о Go, подписывайтесь на меня в твиттере и я буду следить за последними новостями: [@Elliot\_F](https://twitter.com/elliot_f) .

### Дальнейшее чтение

Если вы хотите больше узнать о веб-токенах JSON и о том, как они используются, я могу полностью порекомендовать следующие статьи:
