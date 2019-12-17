## Error response

[Error go play](https://play.golang.org/p/_LxcVyD4-KG)

```golang
package main

import (
	"encoding/base64"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func urlGetter(w http.ResponseWriter, r *http.Request) {
	// get the url we're going to fetch
	vars := mux.Vars(r)
	url, err := base64.URLEncoding.DecodeString(vars["url"])
	if err != nil {
		http.Error(w, "bad url", 400)
		return
	}

	log.Println("GETTING:", string(url))
	resp, err := http.Get(string(url))
	if err != nil {
		log.Printf("ERROR: %s (%s)", url, err)
		http.Error(w, "UPSTREAM ERROR", http.StatusBadGateway)
		return
	}
	resp.Close = true
	defer resp.Body.Close()

	// send back the upstream response on error
	if resp.StatusCode != http.StatusOK {
		log.Printf("Upstream ERROR: %d (%s)", resp.StatusCode, url)
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
		return
	}

	// just dump the bytes for now
	io.Copy(ioutil.Discard, resp.Body)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/get/{url}", urlGetter)
	http.Handle("/", r)

	log.Println("URL Getter Listening on :8888")
	http.ListenAndServe(":8888", nil)
}
```
