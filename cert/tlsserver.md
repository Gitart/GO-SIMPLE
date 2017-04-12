# TLS

[Материал](https://vk.com/videos-27669892?z=video-27669892_171547629%2Fclub27669892%2Fpl_-27669892_-2)

```golang
import (
   "net/http"
   "golang.org/x/net/http2"
   )

func main(){
  server:=new(http.Server)
  n2conf:=new(http2.Server)
  
  n2conf.ConfigureServer(server,n2conf)
  http.Handle(http.FileServer(http.Dir("public")))
  serevr.Addr=":3001"
  server.ListenAndServeTLS("sert.pm","key.pm")
}
```
