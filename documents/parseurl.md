## Parse URL
### Пример получения строки запроса параметров


```golang
package main

import (
	"fmt"
	"log"
	"net/url"
	"net"
	)

func main() {
	u, err := url.Parse("http://User12:passworduser@bing.com:8900/search?q=dotnet&uu=333&zzz=ddddd&k=newtype#Fragment1")
	if err != nil {
		log.Fatal(err)
	}

	R  := u.RawQuery
	S  := u.Scheme 
	H  := u.Host   
	Q  := u.Query()
	P1 := Q["uu"]
	P2 := Q["zzz"]
	M  := u.Opaque
	F  := u.Fragment
	U  := u.User
	I  := U.Username()
	W,_:= U.Password() 
	P  := u.Path
	m,_:= url.ParseQuery(R)
	
	//v:= u.Values{}
	//v  := url.Values{"name=Ava&friend=Jess&friend=Sarah&friend=Zoe"}
	//vv:=v.Get("name")
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println("-----------")
    fmt.Println("HOST          : "+host)
    fmt.Println("PORT          : "+port)
    fmt.Println("PATH          : "+P)
    fmt.Println("PARAM1        : "+m["k"][0])
    fmt.Println("PARAM2        : "+m["zzz"][0])
    fmt.Println("PARAM3        : "+m["q"][0])
    fmt.Println("PARAM4        : "+m["uu"][0])
    fmt.Println("RAWQUERY      : "+R)
    fmt.Println("SHEME         : "+S)
    fmt.Println("HOST:PORT     : "+H)
    fmt.Println("UU            : "+P1[0])
    fmt.Println("ZZZ           : "+P2[0])
    fmt.Println("OPAGUE        : "+M)
    fmt.Println("FRAGMENT      : "+F)
    fmt.Println("USER          : "+I)
    fmt.Println("PASSWORD      : "+W)
    fmt.Println("-----------")
	


	fmt.Println(R,S,H,P1,P2,Q, F,M,U )
	fmt.Println( "Фрагмент="+F )
}
```

#### Применение :

[Пример](http://play.golang.org/p/_wwlMD6eiR)

```html
<form class="form-horizontal" action="/api/clear/?id=ddfff344" method="POST">
```

```golang
func taem(){
  
       u,_:=url.Parse(req.URL.String())
       m,_:= url.ParseQuery(u.RawQuery)
       L:=m["id"][0]
       println(L)
}
```
