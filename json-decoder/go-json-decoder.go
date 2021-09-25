package main

import (
	"encoding/json"
	"fmt"
	"strings"
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
	jsonData := `{"email":"abhirockzz@gmail.com","username":"abhirockzz","blogs":[{"name":"devto","url":"https://dev.to/abhirockzz/"},{"name":"medium","url":"https://medium.com/@abhishek1987/"}]}`
	jsonDataReader := strings.NewReader(jsonData)
	decoder := json.NewDecoder(jsonDataReader)

	var profile Profile
	err := decoder.Decode(&profile)
	if err != nil {
		panic(err)
	}
	fmt.Println(profile.Username)
	fmt.Println(profile.Email)
	fmt.Println(profile.Blogs[0].URL)
	fmt.Println(profile.Blogs[1].URL)
}
