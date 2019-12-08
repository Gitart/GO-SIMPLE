## HTTP(S) прокси на Go в 100 строчек кода

![](https://4gophers.ru/img/proxy/main.png)

Перевод “[HTTP(S) Proxy in Golang in less than 100 lines of code](https://medium.com/@mlowicki/http-s-proxy-in-golang-in-less-than-100-lines-of-code-6a51c2f2c38c)“

В этой статье я опишу реализацию HTTP и HTTPS прокси сервера. С HTTP все просто: сначала парсим запрос от клиента, передаем этот запрос дальше на сервер, получаем ответ от сервера и передаем его обратно клиенту. Нам достаточно использовать HTTP сервер и клиент из пакета `net/http`. С HTTPS все несколько сложнее. Технически это будет туннелирование HTTP с использованием метода CONNECT. Клиент отправляет запрос, указав метод CONNECT, с помощью которого устанавливается соединение между клиентом и удаленным сервером. Как только наш туннель из 2х TCP соединений готов, клиент обменивается TLS рукопожатием с сервером, посылает запрос и ждет ответ.

### Сертификаты

Наш прокси будет работать как HTTPS сервер(если используется параметр `—-proto https`), а это значит нам нужны сертификаты и приватные ключи. В качестве примера будем использовать самоподписанные сертификаты, которые можно сгенерировать вот таким скриптом:

```bash
#!/usr/bin/env bash
case `uname -s` in
    Linux*)     sslConfig=/etc/ssl/openssl.cnf;;
    Darwin*)    sslConfig=/System/Library/OpenSSL/openssl.cnf;;
esac
openssl req \
    -newkey rsa:2048 \
    -x509 \
    -nodes \
    -keyout server.key \
    -new \
    -out server.pem \
    -subj /CN=localhost \
    -reqexts SAN \
    -extensions SAN \
    -config <(cat $sslConfig \
        <(printf '[SAN]\nsubjectAltName=DNS:localhost')) \
    -sha256 \
    -days 3650

```

Необходимо убедить вашу операционную систему доверять получившимся сертификатам. Для этого в OS X можно использовать [Keychain Access](https://tosbourn.com/getting-os-x-to-trust-self-signed-ssl-certificates/).

### HTTP

Для работы с HTTP будем использовать встроенный [клиент и сервер](https://golang.org/pkg/net/http/). Прокся будет обрабатывать полученный запрос, передавать его нужному серверу и возвращать ответ клиенту.

```
   +------+        +-----+        +-----------+
   |client|        |proxy|        |destination|
   +------+        +-----+        +-----------+
1          --Req-->
2                         --Req-->
3                         <--Res--
4          <--Res--

```

### HTTP туннелирование с использованием CONNECT

Если мы хотим использовать HTTPS или WebSockets, то придется поменять тактику. Нам нужен метод HTTP [CONNECT](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/CONNECT). Этот метод работает как приказ серверу установить TCP соединение с необходимым сервером и рулить TCP стримом между сервером и клиентом. В таком случае SSL не будет разрываться и все данные будут передаваться по этому своеобразному туннелю.

```
    +------+            +-----+                   +-----------+
    |client|            |proxy|                   |destination|
    +------+            +-----+                   +-----------+
1           --CONNECT-->
2                              <--TCP handshake-->
3           <--------------Tunnel---------------->

```

### Реализация

```go
package main
import (
    "crypto/tls"
    "flag"
    "io"
    "log"
    "net"
    "net/http"
    "time"
)
func handleTunneling(w http.ResponseWriter, r *http.Request) {
    dest_conn, err := net.DialTimeout("tcp", r.Host, 10*time.Second)
    if err != nil {
        http.Error(w, err.Error(), http.StatusServiceUnavailable)
        return
    }
    w.WriteHeader(http.StatusOK)
    hijacker, ok := w.(http.Hijacker)
    if !ok {
        http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
        return
    }
    client_conn, _, err := hijacker.Hijack()
    if err != nil {
        http.Error(w, err.Error(), http.StatusServiceUnavailable)
    }
    go transfer(dest_conn, client_conn)
    go transfer(client_conn, dest_conn)
}
func transfer(destination io.WriteCloser, source io.ReadCloser) {
    defer destination.Close()
    defer source.Close()
    io.Copy(destination, source)
}
func handleHTTP(w http.ResponseWriter, req *http.Request) {
    resp, err := http.DefaultTransport.RoundTrip(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusServiceUnavailable)
        return
    }
    defer resp.Body.Close()
    copyHeader(w.Header(), resp.Header)
    w.WriteHeader(resp.StatusCode)
    io.Copy(w, resp.Body)
}
func copyHeader(dst, src http.Header) {
    for k, vv := range src {
        for _, v := range vv {
            dst.Add(k, v)
        }
    }
}
func main() {
    var pemPath string
    flag.StringVar(&pemPath, "pem", "server.pem", "path to pem file")
    var keyPath string
    flag.StringVar(&keyPath, "key", "server.key", "path to key file")
    var proto string
    flag.StringVar(&proto, "proto", "https", "Proxy protocol (http or https)")
    flag.Parse()
    if proto != "http" && proto != "https" {
        log.Fatal("Protocol must be either http or https")
    }
    server := &http.Server{
        Addr: ":8888",
        Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            if r.Method == http.MethodConnect {
                handleTunneling(w, r)
            } else {
                handleHTTP(w, r)
            }
        }),
        // Disable HTTP/2.
        TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
    }
    if proto == "http" {
        log.Fatal(server.ListenAndServe())
    } else {
        log.Fatal(server.ListenAndServeTLS(pemPath, keyPath))
    }
}

```

Предупреждаю, что это не готовый к продакшену код. Это только пример. В этом коде не хватает передачи необходимых [hop\-by\-hop заголовков](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#hbh) и правильной настройки таймаутов(об этом можно почитать в прекрасной статье “[Руководство по net/http таймаутам в Go](https://4gophers.ru/articles/rukovodstvo-po-nethttp-taimautam-v-go/)“

Наша прокся будет поддерживать оба способа. По умолчанию будем работать по простой схеме, но создадим туннель если указан метод `CONNECT`

```go
http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodConnect {
        handleTunneling(w, r)
    } else {
        handleHTTP(w, r)
    }
})

```

Функция `handleHTTP` очень простая, поэтому сконцентрируемся на `handleTunneling`. Все начинается с установки соединения:

```go
dest_conn, err := net.DialTimeout("tcp", r.Host, 10*time.Second)
if err != nil {
    http.Error(w, err.Error(), http.StatusServiceUnavailable)
    return
 }
 w.WriteHeader(http.StatusOK)

```

Затем используем интерфейс `Hijacker` чтобы получить соединение с которым работает наш http сервер.

```go
hijacker, ok := w.(http.Hijacker)
if !ok {
    http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
    return
}
client_conn, _, err := hijacker.Hijack()
if err != nil {
    http.Error(w, err.Error(), http.StatusServiceUnavailable)
}

```

Если мы перехватываем соединение, то и обслуживать его дальше должны сами.

Теперь мы можем передавать данные напрямую между двумя TCP соединениями. Собственно, это и будет тем самым туннелем.

```go
go transfer(dest_conn, client_conn)
go transfer(client_conn, dest_conn)

```

В этих рутинах данные передаются от клиента к серверу и обратно.

### Проверяем

Чтобы проверить как все это работает можно использовать хром:

```bash
chrome --proxy-server=https://localhost:8888

```

Или сurl:

```bash
curl -Lv --proxy https://localhost:8888 --proxy-cacert server.pem https://google.com

```

Curl должен быть собран с поддержкой [HTTPS\-прокси](https://daniel.haxx.se/blog/2016/11/26/https-proxy-with-curl/)

### HTTP/2

К сожалению, у нас не получится так просто реализовать прокси для HTTP/2. Все дело в интерфейсе `Hijacker`. Подробности можно узнать тут [#14797](https://github.com/golang/go/issues/14797#issuecomment-196103814).
