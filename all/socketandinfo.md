## Info Load
http://www.ozon.ru/context/detail/id/139528350/

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



## Chanal

```
go run test.go &
echo Hello| nc localhost:5000
kill %1
exit 2
go run test.go
```

```golang
// https://www.youtube.com/watch?time_continue=881&v=HoEn7lXNQOU
// 
package main
import (
   "net"
   )

func serve(conn net.Conn){
	var (buf=make([]byte,1024); rint; err error)
	defer conn.Close()

	for{
		if r, err = conn.Read(buf); err!=nil{
			break
		}

		if _,err=conn.Write(buf[0:r]); err!=nil{
			break
		}
	}
}


func main() {
	sock,_:=net.Listen("tcp",":5000")
	for{
		con,_:=sock.Accept(); go serve(con)
	}
}
```




## Chanels

```
 go run test.go &
 (while [$(curl -s http://localhost:8080) -lt 1000]; do true; done) &
 curl -s http://localhost:8080
 ```
 
 ```golang
// https://www.youtube.com/watch?time_continue=881&v=HoEn7lXNQOU  (14:52)
// https://www.youtube.com/watch?v=k27Oga3Wmxs
// разработка на ГО сервис
// https://www.youtube.com/watch?v=k27Oga3Wmxs&t=14s
// https://www.youtube.com/watch?v=ZvKuEfsqurc
// https://www.youtube.com/watch?v=53WkeeUVoTY&t=7s
// https://www.youtube.com/watch?v=-2NhrYlMum4
// https://www.youtube.com/watch?v=haN-0O3l8uU
// https://www.youtube.com/watch?v=vmGUPuO6Rls

// Сокет
// https://www.youtube.com/watch?v=KQcWXRlAiyA
// https://www.youtube.com/watch?v=CIh8qN7LO8M
package main
import (
   "net/http"
   "strconv"
   )
var ch chan int 

func count(){
	i:=0
	for{ch<-i;i+=1}
}


func handler(w http.ResponseWriter, r *http.Request) {
  i:<-ch
  w.Write([]byte(strconv.Itoa(i)+"\n")
}

func main(){
	ch=make(chan int)
	go count()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
```
