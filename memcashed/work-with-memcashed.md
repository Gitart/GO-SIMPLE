https://medium.com/@litanin/basic-sql-memcached-golang-9e8fd5b7efe1
# базовый | SQL, Memcached и Golang

Создание небольшого простого приложения Golang с тяжелыми запросами MySQL и использование Memcached для повышения производительности API.

# Когда вам нужен кеш?
Допустим, у вас есть приложение Golang, которое подключено к базе данных, использует таблицу «продукты» с примерно 1 млн записей, и вы хотите, чтобы ваши пользователи могли выполнять поиск продуктов по ее имени. Очень простая настройка:

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//Product representing our product model
type Product struct {
	ID   uint
	Name string
}

var db *gorm.DB

func main() {

	var err error

	// open DB connection
	{

		db, err = gorm.Open("mysql",
			"root:root@/tmp?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			panic(err)
		}
		defer db.Close()
	}

	// make a simple handler for our API
	http.HandleFunc("/", myAPI)

	// start HTTP server on port 8080
	http.ListenAndServe(":8080", nil)
}

func myAPI(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	// good practice to limit search by having more precise search criteria - at least more than 3 symbols
	if len(name) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("too short 'name' search parameter"))
		return
	}

	var list []*Product

	// get list of products
	err := db.
		Model(&Product{}).
		Where("name like ?", fmt.Sprintf("%%%v%%", name)). // %my search%
		Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {

		log.Printf("db error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return

	} else if err == gorm.ErrRecordNotFound {

		w.WriteHeader(http.StatusNotFound) // 404
		return

	}

	// marshal product list array into JSON bytes
	payloadBytes, err := json.MarshalIndent(list, "	", " ")
	if err != nil {
		log.Printf("json marshal error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(payloadBytes)
}
```

На моем ноутбуке (MacBook Pro, 2019 г.) обработка одного запроса занимает около 350–400 мс:

![](https://miro.medium.com/max/30/1*3vAJ5_hTtOhHnsSa5t6rHQ.jpeg?q=20)

![](https://miro.medium.com/max/700/1*3vAJ5_hTtOhHnsSa5t6rHQ.jpeg)

Когда я пытаюсь протестировать его с любой нагрузкой, используя эй ( [https://github.com/rakyll/hey](https://github.com/rakyll/hey) ), все становится довольно медленно:

![](https://miro.medium.com/max/24/1*3SKm13NjY6d_CCiNcPuX7g.jpeg?q=20)

![](https://miro.medium.com/max/700/1*3SKm13NjY6d_CCiNcPuX7g.jpeg)

Средняя соответственно о времени NSE составляет около 8 секунд , 50 запросов 20 одновременных рабочих.

Это пример, в котором мы можем использовать кеш для повышения производительности нашего приложения. Кэш полезен, когда вы знаете, что у вас много одинаковых вызовов api на конечной точке api, поэтому вы можете кешировать частичный или весь ответ.

# Memcached

Memcached - одна из самых простых систем кеширования, и для ее установки на Mac:

$ brew установить memcached

Убедитесь, что он запущен и работает:

![](https://miro.medium.com/max/30/1*Mv2DmKVZ6FPHDuhtOqD94g.jpeg?q=20)

![](https://miro.medium.com/max/700/1*Mv2DmKVZ6FPHDuhtOqD94g.jpeg)

# Подключиться к Memcached из Go

Я использую пакет [github.com/rainycape/memcache](http://github.com/rainycape/memcache) для подключения к серверу memcached:

var mc \* memcache.Client
mc, err = memcache.New ("127.0.0.1:11211")
if err! = nil {
 panic (err)
}

Нам нужно использовать в основном 2 функции пакета Memcached: Get и Set, которые будут извлекать данные из кеша и устанавливать данные в кеш. Давайте расширим наш обработчик конечных точек API следующими методами:

```go
func myAPI(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	if len(name) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var list []*Product

	// try get cached data
	{
		listCached, err := mc.Get(fmt.Sprintf("products_%v", name))
		if err == nil {

			w.Header().Set("cached", "1")
			w.Write(listCached.Value)
			return

		} else if err != memcache.ErrCacheMiss {

			log.Printf("memcached error: %v", err)

		}
	}

	// get list of products from DB
	err := db.
		Model(&Product{}).
		Where("name like ?", fmt.Sprintf("%%%v%%", name)).
		Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {

		log.Printf("db error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return

	} else if err == gorm.ErrRecordNotFound {

		w.WriteHeader(http.StatusNotFound)
		return

	}

	// marshal product list array into JSON bytes
	payloadBytes, err := json.MarshalIndent(list, "	", " ")
	if err != nil {
		log.Printf("json marshal error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = mc.Set(&memcache.Item{
		Key:        fmt.Sprintf("products_%v", name),
		Value:      payloadBytes,
		Expiration: 10,
	})
	if err != nil {
		log.Println(err)
	}

	w.Write(payloadBytes)
}
```
В нашем обработчике, прежде чем мы начнем извлекать данные из БД, мы проверяем, есть ли они у нас в кеше:
```go
// try get cached data
	{
		listCached, err := mc.Get(fmt.Sprintf("products_%v", name))
		if err == nil {

			w.Header().Set("cached", "1")
			w.Write(listCached.Value)
			return

		} else if err != memcache.ErrCacheMiss {

			log.Printf("memcached error: %v", err)

		}
	}
  ```
  
  и если данные есть, мы просто записываем их обратно клиенту, а также устанавливаем флаг заголовка «cached» в «1», чтобы уведомить клиента API о том, что обслуживаемые данные были фактически получены из кеша. Иногда вам это нужно, в зависимости от вашей бизнес-логики, в нашем случае мы будем использовать его, чтобы проверить и увидеть, что данные на самом деле поступают из кеша, а теперь и из БД.
Если данные не найдены в кеше, мы следуем нашему обычному потоку и получаем данные из БД, но перед тем, как передать их клиенту, мы сохраним их в кеше с истечением срока действия 10 секунд:
```go
err = mc.Set(&memcache.Item{
		Key:        fmt.Sprintf("products_%v", name),
		Value:      payloadBytes,
		Expiration: 10,
	})
	if err != nil {
		log.Println(err)
	}
  ```
  Таким образом, следующий вызов данного API с тем же параметром поиска в течение 10 секунд получит данные из кеша, а не из БД.
Полный main.go:

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rainycape/memcache"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var mc *memcache.Client

//Product representing our product model
type Product struct {
	ID   uint
	Name string
}

func main() {

	var err error

	// open DB connection
	{
		db, err = gorm.Open("mysql", "root:root@/tmp?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			panic(err)
		}
		defer db.Close()
	}

	// open memcached connection
	{
		mc, err = memcache.New("127.0.0.1:11211")
		if err != nil {
			panic(err)
		}
	}

	// make a simple handler for our API
	http.HandleFunc("/", myAPI)

	// start HTTP server on port 8080
	http.ListenAndServe(":8080", nil)

}

func myAPI(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	if len(name) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var list []*Product

	// try get cached data
	{
		listCached, err := mc.Get(fmt.Sprintf("products_%v", name))
		if err == nil {

			w.Header().Set("cached", "1")
			w.Write(listCached.Value)
			return

		} else if err != memcache.ErrCacheMiss {

			log.Printf("memcached error: %v", err)

		}
	}

	// get list of products from DB
	err := db.
		Model(&Product{}).
		Where("name like ?", fmt.Sprintf("%%%v%%", name)).
		Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {

		log.Printf("db error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return

	} else if err == gorm.ErrRecordNotFound {

		w.WriteHeader(http.StatusNotFound)
		return

	}

	// marshal product list array into JSON bytes
	payloadBytes, err := json.MarshalIndent(list, "	", " ")
	if err != nil {
		log.Printf("json marshal error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = mc.Set(&memcache.Item{
		Key:        fmt.Sprintf("products_%v", name),
		Value:      payloadBytes,
		Expiration: 10,
	})
	if err != nil {
		log.Println(err)
	}

	w.Write(payloadBytes)
}
```

