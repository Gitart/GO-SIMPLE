package main
import (
	    "net/http"
	    "time"
	    "fmt"
)

// **********************************************************
// Write cookies
// Cookies_write(w, "Key",  "Values",  12) 12-Days
// **********************************************************
func Cookies_write(w http.ResponseWriter, Name, Value string, Days int){
	   Expire := time.Now().AddDate(0, 0, Days)
	   http.SetCookie(w, &http.Cookie{Name:Name, Value:Value, Expires:Expire, HttpOnly: true, Path:"/"})
}

// **********************************************************
// Read cookies
// Cookies_read(w, "Key")
// https://stackoverflow.com/questions/54275704/how-to-read-cookies-from-golang/54276297
// vl:=Cookies_read
// vl.Name 
// vl.Value
// **********************************************************
func Cookies_read(r *http.Request, Name string) *http.Cookie {
	   ck, _ := r.Cookie(Name)
	   fmt.Println(ck)
	   return ck
}


func Cookies_read_str(r *http.Request, Name string) string {
	   ck, _ := r.Cookie(Name)
	   
	   return ck.Value
}
