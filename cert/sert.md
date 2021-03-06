## Бесплатные и автоматические SSL сертификаты

![](https://4gophers.ru/img/ssl/main.png)

Перевод статьи “[Free and Automated SSL Certificates with Go](https://goenning.net/2017/11/08/free-and-automated-ssl-certificates-with-go/)“

Сейчас без HTTPS никуда. И это связанно не только с безопасностью. Поисковоки лучше ранжируют сайты работающие по HTTPS в отличии от обычного HTTP.

Уже не то время чтобы искаь оправдания и не использовать HTTPS на вашем сайте.

В этой статье я я расскажу как написать приложение, которе автоматически генерирует SSL сертификаты и использует их для HTTPS. И самое главное: это совершенно бесплатно!

### Требования

Чтобы попробовать все что тут описывается вам нужно:

*   Компилятор Go
*   Серевер, который доступен из интернета. Если у вас такого нет, то можете быстро запустить его на [https://vscale.io/](https://vscale.io/)
*   Нужен домен и настроенный DNS. Большинство облачных сервисов предоставляют их бесплатно, например: `yourvn0001.yourcloud.net`.

### Let’s Encrypt и ACME протокол

[Let’s Encrypt](https://letsencrypt.org/) это известный и доверенный SSL эмиттер который предоставляет бесплатные сертификаты. А самое главное \- с его помощью можно автоматизировать получение сертификатов. С его помощью можно получить сертификат за секунды без регистрации и СМС.

**Autocert** это пакет, который реализует ACME протокол для генерации сертификатов чз Let’s Encrypt. Это единственный пакет, который вам нужен.

Его можно установить с помощью `go get`

```
go get golang.org/x/crypto/acme/autocert

```

Из доклада [2016 \- Matthew Holt \- Go with ACME](https://www.youtube.com/watch?v=KdX51QJWQTA) можно получить больше информации о ACME или альтернативных пакетах.

### ~Магия~ код с пошаговыми объяснениями

```go
package main

import (
    "crypto/tls"
    "fmt"
    "net/http"

    "golang.org/x/crypto/acme/autocert"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello Secure World")
    })

    certManager := autocert.Manager{
        Prompt: autocert.AcceptTOS,
        Cache:  autocert.DirCache("certs"),
    }

    server := &http.Server{
        Addr:    ":443",
        Handler: mux,
        TLSConfig: &tls.Config{
            GetCertificate: certManager.GetCertificate,
        },
    }

    go http.ListenAndServe(":80", certManager.HTTPHandler(nil))
    server.ListenAndServeTLS("", "")
}

```

Все начинается с создания `mux` в функции `main`. Потом делаем хендлер, который выводит “Hello World” сообщение при заходе на корень `/` сайта. В примере мы используем дефолтный mux, но можно использовать любой сторонний роутер.

На следующем шаге мы создаем экземпляр `autocert.Manager`. Эта структура будет использоваться для работы с Let’s Encrypt и получения SSl сертификатов. Поле `Cache` определяет где сохранять сертификаты и откуда их загружать. В этом примере мы используем `autocert.DirCache` и сохраняем сертификаты локальной папке. Это самый простой способ, но не всегда самый лучший: если ваш сайт хостится на нескольких серверах то каждый сервер будут иметь свой локальный кеш.

В конце создаем `http.Server` на `443` и используем наш `mux`. Нам нужен `tls.Config` который мы используем при создании сервера. Тут происходит вся магия. Метод `GetCertificate` используется для указания серверу откуда грузить сертификаты при HTTPS запросе. Это удобно, потому что мы можем выбирать какой сертификат использовать для определенного запроса. Метод `certManager.GetCertificate` сначала пытается загрузить сертификат из кеша, а если там ничего нет, то новый сертификат выкачивается с Let’s Encrypt по ACME протоколу.

В начале 2018 [Let’s Encrypt отключили TLS\-SNI challenge](https://community.letsencrypt.org/t/2018-01-11-update-regarding-acme-tls-sni-and-shared-hosting-infrastructure/50188) из\-за проблем с безопасностью. Рекомендуется использовать [HTTP challenge](https://tools.ietf.org/html/draft-ietf-acme-acme-07#section-8.3). Поэтому мы подключаем еще один листнер на 80 порт и используем `certManager.HTTPHandler(nil)`

Остается только запустить сервер `server.ListenAndServeTLS("", "")`. Если вы работали с HTTPS сертификатами до этой статьи, то должны помнить что необходимо указывать два параметра: `Certificate` и `Privatey Key`. В случае с `autocert` нам ничего не нужно указывать и мы просто оставляем поля пустыми.

Важно упомянуть, что когда мы используем `certManager.HTTPHandler(nil)`, весь трафик автоматически уходит на HTTPS. Это поведение можно исправить, если вместо `nil` указать кастомный хендлер.

### Запускаем

Наше приложение можно запустить как и любое другое приложение на Go. Но локально оно не будет работать. Для Let’s Encrypt необходимо чтобы вебсайт был доступен из вне по известному DNS имени. Когда вы запускаете приложение локально \- Let’s Encrypt не может достучаться до вашего домена для верификации.

1.  Создайте DNS A запись для вашего домена.
2.  Соберите приложение: `CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o autossl`. Тут мы компилируем для linux.
3.  Залейте файл на сервер.
4.  Запустите приложение на сервере.
5.  Откройте ваше приложение в браузере и укажите адрес вашего сайта.

![](https://4gophers.ru/img/ssl/auto-ssl-golang.png)

И все работает!

Вы должны увидеть сообщение `Hello Secure World` и зеленую SSL иконку.

Первый запрос может занять несколько секунд. Но как только сертификат будет сгенерирован и закеширован, то запросы сразу станут молниеносными.

### Важные замечания и советы

1.  Существует ограничение на количество генерируемых сертификатов для одного домена. На момент написания статьи можно было сгенерировать не больше 20 сертификатов в неделю. Это более чем достаточно, но если вы не будете следить за кешом, то очень быстро исчерпаете лимит. Больше информации по лимитам можно найти в документации: [https://letsencrypt.org/docs/rate\-limits/](https://letsencrypt.org/docs/rate-limits/).
2.  То как хранить сертификаты остается на ваше усмотрение. Но хранить их локально на сервере становится неудобно как только у вас появляется необходимость использовать кластер. Много серверов, много запросов на генерацию сертификатов и вот вы уже не вкладываетесь в лимиты.
3.  Сертификаты полученные через Let’s Encrypt действительны только 90 дней. Правда, **autocert** может сам их перегенерировать и нам ничего не нужно будет делать. Тем не менее, стоит мониторить срок истечения сертификата.
4.  Сертификаты ограничены и вы не можете использовать для поддоменов(для этого нужно использовать wildcard сертификаты)
5.  Приложение с сертификатом не получится запускать локально, а значит вам нужна настройка которая будет переключать режимы HTTP/HTTPS

Если у вас есть предложения и вопросы \- пишите их в комментариях.

Спасибо!
