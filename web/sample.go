package main
 
import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
 
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
)
 
var conn *sql.DB
 
type Users struct {
    Users []User `json:"users"`
}
 
type User struct {
    ID    int    "json:id"
    Name  string "json:username"
    Email string "json:email"
    First string "json:first"
    Last  string "json:last"
}
 
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
 
    NewUser := User{}
    NewUser.Name = r.FormValue("user")
    NewUser.Email = r.FormValue("email")
    NewUser.First = r.FormValue("first")
    NewUser.Last = r.FormValue("last")
    
    output, err := json.Marshal(NewUser)
    fmt.Println(string(output))
    
    if err != nil {
        fmt.Println("Error while marhalling User")
    }
 
    sql := "INSERT INTO users set user_nickname='" + NewUser.Name + "', user_first='" + NewUser.First + "', user_last='" + NewUser.Last + "', user_email='" + NewUser.Email + "'"
    query, err := database.Exec(sql)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(q)
}
 
func RetrieveUserHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("starting retrieval")
    start := 0
    limit := 10
 
    next := start + limit
 
    w.Header().Set("Pragma", "no-cache")
    w.Header().Set("Link", "<http://localhost:8282/api/users?start="+string(next)+"; rel=\"next\"")
 
    rows, _ := database.Query("select * from users LIMIT 10")
    Response := Users{}
 
    for rows.Next() {
 
        user := User{}
        rows.Scan(&user.ID, &user.Name, &user.First, &user.Last, &user.Email)
 
        Response.Users = append(Response.Users, user)
    }
 
    output, err := json.Marshal(Response)
    if err != nil {
        fmt.Print("Error while marshalling output")
    }
    fmt.Fprintln(w, string(output))
}
func main() {
 
    conn, err := sql.Open("mysql", "root@/social_network")
    if err != nil {
        panic("Cant establish connection to DB")
    }
    database = conn
    routes := mux.NewRouter()
    routes.HandleFunc("/api/users", UserCreate).Methods("POST")
    routes.HandleFunc("/api/users", UsersRetrieve).Methods("GET")
    http.Handle("/", routes)
    http.ListenAndServe(":8080", nil)
}
