
// Get client IP Address in Go
// Знание того, кто получает доступ к веб-серверу может быть полезно для отслеживания активности пользователя. 
// В этом уроке мы узнаем, как получить IP-адрес клиента в Go.
// Чтобы получить IP-адрес клиента, просто использовать
// ip,_,_ := net.SplitHostPort(r.RemoteAddr)
// Тем не менее, иногда пользователь получить доступ к веб-серверу через прокси или балансировки нагрузки 
// (например, в облаке хостинг и т.д.) и приведенный выше код даст вам IP-адрес прокси или балансировки нагрузки.
// Чтобы получить реальный IP - адрес клиента, вы можете получить информацию заголовка HTTP Request "X-Forwarded-For (XFF)"
// fmt.Fprintf(w,"X-Forwarded-For :" + r.Header.Get("X-FORWARDED-FOR"));
// Это полный пример кода, который можно выполнить в вашем веб-сервере и проверить с помощью веб-браузера.

 package main

 import (
  "fmt"
  "net/http"
  "net"
 )

 func IndexPage(w http.ResponseWriter, r *http.Request) {
 
  // get client ip address
  ip,_,_ := net.SplitHostPort(r.RemoteAddr)

  // print out the ip address
  fmt.Fprintf(w,ip + "\n\n")

  // sometimes, the user acccess the web server via a proxy or load balancer.
  // The above IP address will be the IP address of the proxy or load balancer and not the user's machine.

  // let's get the request HTTP header "X-Forwarded-For (XFF)"
  // if the value returned is not null, then this is the real IP address of the user.
  fmt.Fprintf(w,"X-Forwarded-For :" + r.Header.Get("X-FORWARDED-FOR"));
 }


 func main() {
  http.HandleFunc("/", IndexPage)
  http.ListenAndServe(":8888", nil)
 }
