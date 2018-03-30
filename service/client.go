package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "http://localhost:9200/_search"

	payload := strings.NewReader("{\r\n  \"query\": {\r\n    \"match\": {\r\n      \"name\": \"New-York\"\r\n    }\r\n  }\r\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "b9905cb8-87f9-4c2d-8496-06bf81e42c8b")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
