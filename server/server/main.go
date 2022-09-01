package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// hello() is home handler.
func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, htmlStr)
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Hello, POT method. ParseForm() err: %v", err)
			return
		}

		// Post form from website
		switch r.FormValue("post_from") {
		case "web":
			fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
			s := r.FormValue("key")
			fmt.Fprintf(w, "key = %s, len = %v\n", s, len(s))

		case "client":
			fmt.Fprintf(w, "Post from client! r.PostForm = %v\n", r.PostForm)

		default:
			fmt.Fprintf(w, "Unknown post source:-(\n")
		}
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "/json only support POST method", http.StatusBadRequest)
		return
	}

	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	u := User{}

	// Create a JSON decoder to read and decode JSON stream from request.Body.
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&u); err != nil {
		http.Error(w, "decode JSON error", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "decode JSON successfully. Name: %v, Email: %v\n", u.Name, u.Email)
	return
}

func main() {
	// Index Handler
	http.HandleFunc("/", hello)

	// JSON handler
	http.HandleFunc("/json", jsonHandler)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

var htmlStr = `
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8" />
</head>
<body>
  <div>
      <form method="POST" action="/">
          <input name="post_from" type="hidden" value="web" >
          <input name="key" type="text" value="Hello, 世界">
	  <input type="submit" value="submit" />
      </form>
  </div>
</body>
</html>
`
