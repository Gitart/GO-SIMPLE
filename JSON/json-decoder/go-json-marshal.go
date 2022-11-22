package main

import (
	"encoding/json"
	"fmt"
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
	//jsonData, err := json.Marshal(&myprofile)
	jsonData, err := json.MarshalIndent(&profile, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
}
