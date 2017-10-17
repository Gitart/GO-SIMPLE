
// Simple Golang HTTP Request Context Example
// https://gocodecloud.com/blog/2016/11/15/simple-golang-http-request-context-example/

package main

import (
	"context"
	"log"
	"net/http"
	"time"
)


// **********************************************************************
// Start
// **********************************************************************
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/",         StatusPage)
	mux.HandleFunc("/login",    LoginPage)
	mux.HandleFunc("/logout",   LogoutPage)

	log.Println("Start server on port :8085")

	contextedMux := AddContext(mux)
	
	log.Fatal(http.ListenAndServe(":8085", contextedMux))
}



// **********************************************************************
// Добавление в контекст
// **********************************************************************
func AddContext(next http.Handler) http.Handler {
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


// **********************************************************************
// Status page
// **********************************************************************
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


// **********************************************************************
// Логирование операции добавление в куки на год
// **********************************************************************
func LoginPage(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(365 * 24 * time.Hour)  
	cookie     := http.Cookie{Name: "username", Value: "alice_cooper@gmail.com", Expires: expiration}
	http.SetCookie(w, &cookie)
}


// **********************************************************************
// Выход из страницы
// Set to expire in the past
// **********************************************************************
func LogoutPage(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().AddDate(0, 0, -1)	
	cookie     := http.Cookie{Name: "username", Value: "alice_cooper@gmail.com", Expires: expiration}
	http.SetCookie(w, &cookie)
}

