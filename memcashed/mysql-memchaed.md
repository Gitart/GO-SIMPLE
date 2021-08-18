## Sampple connect to memcahed and MySQL



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
