#Usage Context 
```golang
package main

import (
	"context"
	"log"
	"net/http"
	"time"
	"fmt"
)


func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/",       StatusPage)
	mux.HandleFunc("/login",  LoginPage)
	mux.HandleFunc("/logout", LogoutPage)
	mux.HandleFunc("/test",   TestPage)

	log.Println("Start server on port :8085")
	
	contextedMux := AddContext(mux)
	
	log.Fatal(http.ListenAndServe(":8085", contextedMux))
}




// *************************************************************************
// Эта функция будет выполняться для всех хендлеров
// *************************************************************************
func AddContext(next http.Handler) http.Handler {
	 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("METHOD ",r.Method, " URL:", r.RequestURI)
		
	})
}



func AddContext1(next http.Handler) http.Handler {
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		log.Println(r.Method, "-", r.RequestURI)
		
		cookie, _ := r.Cookie("username")
	
		if cookie != nil {
			//Add data to context
			ctx := context.WithValue(r.Context(), "Username", cookie.Value)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}





func TestPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("FFFF TEST")
}

func StatusPage(w http.ResponseWriter, r *http.Request) {
	//Get data from context
	if username := r.Context().Value("Username"); username != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello " + username.(string) + "\n"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Logged in"))
	}
}



func LoginPage(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(365 * 24 * time.Hour)  ////Set to expire in 1 year
	cookie     := http.Cookie{Name: "username", Value: "alice_cooper@gmail.com", Expires: expiration}
	http.SetCookie(w, &cookie)
}


func LogoutPage(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().AddDate(0, 0, -1)	//Set to expire in the past
	cookie     := http.Cookie{Name: "username", Value: "alice_cooper@gmail.com", Expires: expiration}
	http.SetCookie(w, &cookie)
}
```
