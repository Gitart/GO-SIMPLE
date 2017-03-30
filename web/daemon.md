## godaemon

```
go get github.com/icattlecoder/godaemon
```


```golang
package main

import (
	_ "github.com/icattlecoder/godaemon"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/index", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("hello, golang!\n"))
	})
	log.Fatalln(http.ListenAndServe(":7070", mux))
}
```

### Run
```
./example -d=true
~$ curl http://127.0.0.1:7070/index
```

hello, golang!
