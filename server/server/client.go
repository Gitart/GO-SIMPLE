package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	homeURL = "http://localhost:8080"
	jsonURL = homeURL + "/json"
)

func main() {
	// Add form data
	v := url.Values{}
	v.Set("post_from", "client")
	v.Set("o", "omega")
	v.Set("b", "beta")
	v.Set("z", "zeta")
	v.Set("a", "alpha")
	v.Set("g", "gamma")
	v.Set("ch", "chip china cheap")
	v.Set("中", "中文 中国 中心")

	// Values.Encode() encodes the values into "URL encoded" form sorted by key.
	s := v.Encode()
	fmt.Printf("v.Encode(): %v\n", s)

	req, err := http.NewRequest("POST", homeURL, strings.NewReader(s))
	if err != nil {
		fmt.Printf("http.NewRequest() error: %v\n", err)
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("http.Do() error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll() error: %v\n", err)
		return
	}

	fmt.Printf("read resp.Body successfully:\n%v\n", string(data))

	// POST JSON data and read reponse.
	jsonData := `{"name":"Frank", "email":"xx@xx.com"}`
	reqJSON, err := http.NewRequest("POST", jsonURL, strings.NewReader(jsonData))
	if err != nil {
		fmt.Printf("http.NewRequest() error: %v\n", err)
		return
	}

	req.Header.Add("Content-Type", "application/json")

	respJSON, err := c.Do(reqJSON)
	if err != nil {
		fmt.Printf("http.Do() error: %v\n", err)
		return
	}
	defer respJSON.Body.Close()

	data, err = ioutil.ReadAll(respJSON.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll() error: %v\n", err)
		return
	}

	fmt.Printf("read resp.Body successfully:\n%v\n", string(data))

}
