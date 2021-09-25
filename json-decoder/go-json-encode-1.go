package main

import (
	"encoding/json"
	"os"
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

	encoder := json.NewEncoder(os.Stdout)
	err := encoder.Encode(&profile)
	if err != nil {
		panic(err)
	}
}
