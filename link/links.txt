GO
http://go-lang.cat-v.org/go-code                                                 Коды на GO
http://habrahabr.ru/post/229799/                                                 Go + Heroku: развертывание web-приложения
http://go-lang.cat-v.org/pure-go-libs                                            - Описание GO
http://habrahabr.ru/post/160833/                                                 Google Drive теперь поддерживает публикацию веб-сайтов
https://developers.google.com/drive/web/publish-site 
https://developers.google.com/drive/web/folder
https://developers.google.com/drive/v2/reference/files
http://habrahabr.ru/post/235417/                                                -- Переезд на ГО
http://habrahabr.ru/company/yandex/blog/237985/                                 -- в Яндексе
http://habrahabr.ru/post/229799/                                                -- Go HEROKU
http://gopractice.ru/                                                           -- GO
http://habrahabr.ru/post/234693                                                 -- Сервис загрузки
https://groups.google.com/forum/#!forum/golang-nuts                             -- Сообщество
                    
 
Разные сведения
http://habrahabr.ru/post/145416/ Google Drive. Отчет с данными из таблицы. Создание простенькой БД. Часть 
http://habrahabr.ru/post/145832/ Кластеризация на клиенте или как показать 10000 точек на карте
http://habrahabr.ru/post/145447/ - Оповещение о новых письмах в Gmail по SMS средствами Google Calendar + Google Apps Script
http://habrahabr.ru/post/147179/ - Сохранение контактной информации в Google Contacts
http://habrahabr.ru/post/190428/ - Модификация стоковых прошивок для Android. Часть 4
http://habrahabr.ru/post/149365/ - Программирование в Android — зачем такие сложности?
http://habrahabr.ru/post/122095/ - Веб-разработка на Go
http://myrusakov.ru/html.html 
https://golang.org/pkg/encoding/json/    - Jason
http://play.golang.org/p/R1D0NoClbz - my program
http://golang.org/ref/spec#TypeName - type
http://golang.org/doc/effective_go.html - efective
http://habrahabr.ru/post/205268/ - Блокнот с графическим интерфейсом на языке Go
http://habrahabr.ru/post/225481/ - Разбор строки адреса (улица [дом]) средствами Golang и Postgis
http://habrahabr.ru/post/197598/  -  пишем граббер веб страниц с многопоточностью и блудницами из песочницы
http://habrahabr.ru/post/219459/ - Язык Go для начинающих
http://habrahabr.ru/post/186568/  Автоматическое оповещение об изменениях статуса почтовых посылок через SMS
http://habrahabr.ru/post/195950/ - Быстрые треки на google maps
 
Подборка библиотек для бекенда
Go*
Мы пишем свои бекенды на Go. Собираем метрики кода и балансируем запросы на шарды.
Шифруем RPC. Общаемся с Монгой. За год разработки сформировался стек проверенных библиотек.
Например, goagain сэкономил кучу времени и дебага после обрыва клиентов внутреннего RPC.
 
Делимся подборкой библиотек, проверенных и работающих в бою.
 
github.com/rcrowley/goagain                              Перезагрузка HTTP или RPC сервера без отключения клиентов.
github.com/cheggaaa/pb                                   Консольный прогресс-бар. Поддерживает интерфейсы Reader и Writer.
github.com/rcrowley/go-metrics                           Метрики в коде. Счетчики, Перцентили, Гистограммы. Красиво дампит в консоль и умеет складывать в Graphite и InfluxDB
github.com/golang/glog                                   Логгер. Порт Google Log. Можно задавать уровень детализации (verbosity). Еще детализация фильтруется по модулям.
labix.org/mgo                                           Лучший драйвер для MongoDB. 
github.com/camlistore/lock                              Лок-файл. Пишет PID в файл. Если лок есть, проверит не умер-ли процесс. Поддерживает FreeBSD, MacOS, ARM и Plan9.
github.com/codegangsta/cli                              Разделяет логику по флагам и действиям. Генерирует справку и автодополнение для консоли.
godoc.org/code.google.com/p/go-uuid/uuid                Генерация стандартного UUID.
godoc.org/code.google.com/p/go.crypto/ssh               Полноценный SSH транспорт. Умеет парсить приватный и публичный ключ.
github.com/kr/pretty                                    Выводит глубокие структуры в читаемом виде. Достойная замена fmt.Printf("%+v", ...)
github.com/google/btree                                 Реализация B-Tree от Гугла.
github.com/bitly/dablooms                               Фильтры Блума со счетчиками и удалением.
github.com/bitly/go-hostpool                            Балансер любых ресурсов с обратной связью. Мы используем для медленнего пессимизации лагающих шардов. В основе алгоритм Многорукого Бандита.
github.com/influxdb/influxdb                            База для метрик. Быстрее, чем Graphite. Поддерживает шардирование и выборки аля SQL.
godoc.org/code.google.com/p/go.crypto/nacl              Прикладное быстрое шифрование на элиптических кривых.
