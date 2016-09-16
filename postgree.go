package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    _ "github.com/lib/pq"
)

type Book struct {
    isbn   string
    title  string
    author string
    price  float32
}

var b []Book

func main() {

    db, err := sql.Open("postgres", "postgres://****:****@localhost/postgres?sslmode=disable")

    if err != nil {
        log.Fatal(err)
    }
    rows, err := db.Query("SELECT * FROM books")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var bks []Book
    for rows.Next() {
        bk := new(Book)
        err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
        if err != nil {
            log.Fatal(err)
        }
        bks = append(bks, *bk)
    }
    if err = rows.Err(); err != nil {
        log.Fatal(err)
    }

    b = bks

    http.HandleFunc("/db", getBooksFromDB)
    http.ListenAndServe("localhost:1337", nil)

}

func getBooksFromDB(w http.ResponseWriter, r *http.Request) {

    fmt.Println(b)
    response, err := json.Marshal(b)
    if err != nil {
        panic(err)

    }

    fmt.Fprintf(w, string(response))
}
