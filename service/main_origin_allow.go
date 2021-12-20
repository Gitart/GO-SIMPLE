package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	BookName string `json:"book_name"`
	Slug     string `json:"slug"`
	Author   string `json:"author"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getSingleBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, getItem := range books {
		if getItem.Slug == params["slug"] {
			json.NewEncoder(w).Encode(getItem)
			return
		}
	}
}

func main() {
	books = append(books, Book{
		BookName: "Half of the Yellow Sun",
		Slug:     "half-of-the-yellow-sun",
		Author:   "Chimamanda Adichie",
	}, Book{
		BookName: "Americanah",
		Slug:     "americanah",
		Author:   "Chimamanda Adichie",
	})

	router := mux.NewRouter()
	router.HandleFunc("/", getBooks).Methods("GET")
	router.HandleFunc("/{slug}", getSingleBook).Methods("GET")

	log.Println("Starting your server at :8080")
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}
