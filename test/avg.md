
```golang
package main
import (
   "net/http"
   "encoding/json"
   stats "github.com/c9s/goprocinfo/linux"
)

func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/json")
	stats,_:=stats.ReadLoadAvg("proc/loadavg")
	resp,_:=json.MarshalIndent(stats, "","  ")
	w.Write(resp)
	
}
```
