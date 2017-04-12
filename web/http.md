# HTTPS и Go
Это перевод статьи "HTTPS and Go". Статья больше для новичков, чем для матерых гоферов, но есть полезная информация для всех программистов. 

Работа с HTTP сервером - это одна из первых задач, с которой сталкивается начинающий Go программист. 

Реализовать простенький HTTP сервер на Go легко. Необходимо написать всего пару строк кода и у вас готов и работает сервер на 8080 порту: 

```golang
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Привет!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

Откройте страничку https://127.0.0.1:8080 в вашем браузере и вы увидите сообщение "Привет!". 

Но что если вам нужно работаться с защищенным HTTPS соединением? В первом приближении, это достаточно просто. Для этого можно использовать метод 
```golang
ListenAndServeTLS, вместо http.ListenAndServe(":8080", nil).
```
http.ListenAndServeTLS(":8081", "cert.pem", "key.pem", nil)

И все готово. Ну, почти. Эта функция получает на два аргумента больше: "cert.pem" - ваш серверный сертификат в PEM формате,
"key.pem" - приватный ключ в PEM формате. 

Получение сертификата для сервера и приватного ключа
### Использование OpenSSL
Вы можете легко сгенерировать оба файла с помощью OpenSSL. OpenSSL поставляется в Mac OS X и Linux. Если вы используете Windows, то вам нужно установить бинарники отдельно.

К счастью, для генерирование сертификата и приватного ключа с помощью OpenSSL достаточно одной команды:
```
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem
```

Вам нужно будет ответить на пару вопросов в момент генерации. Самая важная часть, это поле "Common Name (e.g. server FQDN or YOUR name)". Тут вы должны указать имя вашего сервера (например myblog.com, или 127.0.0.1:8081 если вам нужен доступ к вашей локальной машине на 8081 порту).
После этого, вы обнаружите два файла "cert.pem" и "key.pem" в той папке, где вы запускали OpenSSL команду. 
Учтите, что эти файлы называются самоподписанным сертификатом. Это значит, что вы можете использовать эти файлы, но браузер будет
определять соединение как небезопасное.

Вы можете сами проверить это, как только запустите сервер 
```golang
http.ListenAndServeTLS(":8081", "cert.pem", "key.pem", nil)
```

с указанием сгенерированного файла. Перейдя на страничку https://127.0.0.1:8081 в браузере вы увидите предупреждение безопасности. 

Это означает, что сертификат на сервере не подписан доверенным центром сертификации. В Firefox нужно кликнуть "Дополнительно" и затем "Добавить исключение..." после этого браузер перейдет на сайт. В Google Chrome нужно кликнуть "Дополнительные" и затем "Перейти на сайт (небезопасно)", тогда последует переход на страницу.

## Используем Go

Есть другой способ генерации файлов сертификата и ключа - вы можете сделать это непосредственно с помощью Go кода. 
В стандартной поставке Go есть пример программы, которая демонстрирует как это делается. Она называется ***generate_cert.go.***

Для удобства, я собрал все это в отдельную библиотеку, названную httpscerts. Мы можем модифицировать нашу программу для использования 
httpscerts и автоматической генерации необходимых сертификатов:

```golang
package main

import (
    "fmt"
    "github.com/kabukky/httpscerts"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Привет")
}

func main() {
    // Проверяем, доступен ли cert файл.
    err := httpscerts.Check("cert.pem", "key.pem")
    // Если он недоступен, то генерируем новый.
    if err != nil {
        err = httpscerts.Generate("cert.pem", "key.pem", "127.0.0.1:8081")
        if err != nil {
            log.Fatal("Ошибка: Не можем сгенерировать https сертификат.")
        }
    }
    http.HandleFunc("/", handler)
    http.ListenAndServeTLS(":8081", "cert.pem", "key.pem", nil)
}
```

Конечно, ваш браузер все также отобразит предупреждение о самоподписанном сертификате.

## Использование StartSSL
Самодписанные сертификаты это удобно для тестирования. Но как только вы запустите сервер в продакшене, вам станет нужен сертификат 
подписанный в доверенном центре, который будет нормально принимать браузер и операционные системы.

К сожалению, это платная услуга. Для примера, Comodo сдерет с вас $100 за сертификат на 1 год. Internet Security Research Group 
работает над этой проблемой, но пока нет возможности получить бесплатный сертификат.

Единственная альтернатива, это использовать сервис StartSSL. StartSSL выдает сертификат для одного домена, за который не нужно будет 
платить в течении первого года. Конечно, вам придется им заплатить за отзыв сертификата, в случае Heartbleed например, 
но сейчас это единственный вариант получить бесплатный сертификат, хоть и на ограниченное время.

Зарегистрируйтесь на ***StartSSL*** и сгенерируйте сертификат и приватный ключ для своего домена. Внимательно прочитайте
инструкцию или найдите пару туториалов по использованию StartSSL.

В дальнейшем, будем считать, что вы сохранили сертификат как "cert.pem" и приватный ключ как "key.pem".

Ваш сертификат может быть защищен паролем. Для этого откройте "key.pem" в обычном текстовом редакторе. Если он действительно зашифрован 
паролем, то вы увидите что-то вроде:

```
Proc-Type: 4,ENCRYPTED
```

Чтобы удалить пароли из приватного ключа, используйте OpenSSL команду: 
```
openssl rsa -in key.pem -out key_unencrypted.pem
```

В конце концов, вам нужно добавить StartSSL Intermediate CA и StartSSL Root CA в "cert.pem" 
Скачайте "Class 1 Intermediate Server CA" 
и "StartCom Root CA (PEM encoded)" из StartSSL Tool Box (Log In > Tool Box > StartCom CA Certificates) и 
положите файлы рядом с вашим "cert.pem". Используя Linux или Mac OS X, запустите:

```
cat cert.pem sub.class1.server.ca.pem ca.pem > cert_combined.pem
```

Используя Windows, запустите: 
```
type cert.pem sub.class1.server.ca.pem ca.pem > cert_combined.pem
```

Теперь вы можете использовать "cert_combined.pem" и "key_unencrypted.pem" в вашей Go программе.
Если хотите, можете переименовать их в "cert.pem" и "key.pem". 

## Обработка HTTP соединения

Сейчас ваш HTTPS сервер работает отлично. Используя StartSSL сертификат вы не будете видеть предупреждения, 
заходя на ваш сайт https://yourdomain.com.
Но как быть http://yourdomain.com? Так как HTTP сервер теперь не запущен, то страничка не загрузится. 
Есть два способа решить эту проблему.

## Раздача одинакового контента по HTTP и HTTPS

Это очень просто реализовать:

```golang
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there!")
}

func main() {
    http.HandleFunc("/", handler)
    // Запуск HTTPS сервера в отдельной go-рутине
    go http.ListenAndServeTLS(":8081", "cert.pem", "key.pem", nil)
    // Запуск HTTP сервера
    http.ListenAndServe(":8080", nil)
}
```

## Редирект с HTTP на HTTPS

Это наиболее верный подход, если вы хотите заинкриптить весь ваш трафик. Для достижения этого, вам нужна функция,
которая будет выполнять редирект с HTTP на HTTPS:

```golang
package main

import (
    "fmt"
    "net/http"
)

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
    // Перенаправляем входящий HTTP запрос. Учтите, 
    // что "127.0.0.1:8081" работает только для вашей локальной машина
    http.Redirect(w, r, "https://127.0.0.1:8081"+r.RequestURI, 
                            http.StatusMovedPermanently)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there!")
}

func main() {
    http.HandleFunc("/", handler)
    // Запуск HTTPS сервера в отдельной go-рутине
    go http.ListenAndServeTLS(":8081", "cert.pem", "key.pem", nil)
    // Запуск HTTP сервера и редирект всех входящих запросов на HTTPS
    http.ListenAndServe(":8080", http.HandlerFunc(redirectToHttps))
}
```

Или, используя два разных ServeMux для HTTP и HTTPS серверов, вы можете редиректить на HTTPS только по специфическим путям
(например /admin/): 

```golang
package main

import (
    "fmt"
    "net/http"
)

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
    // Перенаправляем входящий HTTP запрос. Учтите, 
    // что "127.0.0.1:8081" работает только для вашей локальной машина
    http.Redirect(w, r, "https://127.0.0.1:8081"+r.RequestURI, 
                                http.StatusMovedPermanently)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there!")
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi admin!")
}

func main() {
    // Создаем новый ServeMux для HTTP соединений
    httpMux := http.NewServeMux()
    // Создаем новый ServeMux для HTTPS соединений
    httpsMux := http.NewServeMux()
    // Перенаправляем /admin/ на HTTPS
    httpMux.Handle("/admin/", http.HandlerFunc(redirectToHttps))
    // Обрабатываем все остальное
    httpMux.Handle("/", http.HandlerFunc(homeHandler))
    // Так же, обрабатываем все по HTTPS
    httpsMux.Handle("/", http.HandlerFunc(homeHandler))
    httpsMux.Handle("/admin/", http.HandlerFunc(adminHandler))
    // Запуск HTTPS сервера в отдельной go-рутине
    go http.ListenAndServeTLS(":8081", "cert.pem", "key.pem", httpsMux)
    // Запуск HTTPS сервера
    http.ListenAndServe(":8080", httpMux)
}
```

И на этом все. Экспериментируйте с HTTPS и Go в удовольствие!
