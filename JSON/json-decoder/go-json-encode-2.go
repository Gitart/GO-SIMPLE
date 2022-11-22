package main

import (
	"encoding/json"
	"net/http"
)

type Profile struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Blogs    []Blog `json:"blogs,omitempty"`
}

type Blog struct {
	BlogName string `json:"name"`
	URL      string `json:"url"`
}

func main() {

	profile := Profile{Email: "abhirockzz@gmail.com", Username: "abhirockzz", Blogs: []Blog{
		Blog{BlogName: "devto", URL: "https://dev.to/abhirockzz/"},
		Blog{BlogName: "medium", URL: "https://medium.com/@abhishek1987/"},
	}}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		err := encoder.Encode(&profile)
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}
