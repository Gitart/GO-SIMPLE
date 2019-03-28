package main
import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "net/http"
    "html/template"
    "log"
)
type Product struct{
    Id int
    Model string
    Company string
    Price int
}
var database *sql.DB
 
// функция добавления данных
func CreateHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
 
        err := r.ParseForm()
        if err != nil {
            log.Println(err)
        }
        model := r.FormValue("model")
        company := r.FormValue("company")
        price := r.FormValue("price")
 
        _, err = database.Exec("insert into productdb.Products (model, company, price) values (?, ?, ?)", 
          model, company, price)
 
        if err != nil {
            log.Println(err)
          }
        http.Redirect(w, r, "/", 301)
    }else{
        http.ServeFile(w,r, "templates/create.html")
    }
}
 
func IndexHandler(w http.ResponseWriter, r *http.Request) {
 
    rows, err := database.Query("select * from productdb.Products")
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()
    products := []Product{}
     
    for rows.Next(){
        p := Product{}
        err := rows.Scan(&p.Id, &p.Model, &p.Company, &p.Price)
        if err != nil{
            fmt.Println(err)
            continue
        }
        products = append(products, p)
    }
 
    tmpl, _ := template.ParseFiles("templates/index.html")
    tmpl.Execute(w, products)
}
 
func main() {
      
    db, err := sql.Open("mysql", "root:password@/productdb")
     
    if err != nil {
        log.Println(err)
    }
    database = db
    defer db.Close()
    http.HandleFunc("/", IndexHandler)
    http.HandleFunc("/create", CreateHandler)
 
    fmt.Println("Server is listening...")
    http.ListenAndServe(":8181", nil)
}
