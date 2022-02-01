package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func repeat(w io.Writer, word string, reps int) {
	for i := 0; i < reps; i++ {
		io.WriteString(w, word)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	repeat(w, "Sammy", 5)
}
