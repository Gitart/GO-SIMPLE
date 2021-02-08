package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type User2 struct {
    Name string    `json:"name"`
    Age  int       `json:"age"`
    Test string    `json:"-"`   
    Ot   string    `json:"ot,omtiempty"`  
}

type User struct {
    Name string    `json:"name"`
    Age  int       `json:"age"`
    Ot   string    `json:"ot"`  
}


// При получении данных с BODY (JSON) - используется метод 


// https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go-ru
func (u *User) MarshalJSON() ([]byte, error) {
    type userResponse struct {
        Name string
    }
    var reply userResponse
    reply.Name = u.Name
    return json.Marshal(&reply)
}

func main(){
      host := "0.0.0.0:3000"
	  http.HandleFunc("/",   AdminHandler) 
      http.HandleFunc("/1/", AdminHandlers) 
	  log.Fatal(http.ListenAndServe(host, nil))
}

// *********************************************************
// Исключение полей
// *********************************************************
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	         var Dt User 
	        // json.NewDecoder(r.Body).Decode(&Dt)
             // out,  := json.Marshal(&Dt, "", "  ")
            reads, _ := ioutil.ReadAll(r.Body) 
            json.Unmarshal([]byte(reads), &Dt)
            fmt.Printf("%+v \n", Dt)


            //  out, err := json.MarshalIndent(Dt, "", "  ")
            //  if err != nil {
            //      log.Println(err)
               
            //    }

            // fmt.Println(string(out)) 



            // fmt.Fprintln(w, string(out))
            // fmt.Fprintln(w, Dt)



            data, _ := json.Marshal(Dt)
		    w.Header().Set("Content-Type", "application/json; charset=utf-8")
		    w.Write(data)
}

func AdminHandle1(w http.ResponseWriter, r *http.Request) {
	         var Dt User 
 	        json.NewDecoder(r.Body).Decode(&Dt)
            fmt.Fprintln(w, Dt)
}

type User1 struct {
    Name      string    `json:"name"`
    Password  string    `json:"-"`
    CreatedAt time.Time `json:"createdAt"`
    Test      string    `json:"-"`   
    Ot        string    `json:"ot,omtiempty"`  
}

func AdminHandlers(w http.ResponseWriter, r *http.Request) {
    u := &User1{
        Name:      "Sammy the Shark",
        Password:  "fisharegreat",
        CreatedAt: time.Now(),
    }

    out, err := json.MarshalIndent(u, "", "  ")
    if err != nil {
        log.Println(err)
        os.Exit(1)
    }

    fmt.Println(string(out))
}
