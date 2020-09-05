package hacknews

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Initializer -- represents the kind of requests that we want to fetch
type Initializer struct {
	Story   string
	NbPosts int
}

// Post -- represents the json object returned by the API
type Post struct {
	Id          int    `json:"id"`
	Deleted     bool   `json:"deleted"`
	Type        string `json:"type"`
	By          string `json:"by"`
	Time        int    `json:"time"`
	Text        string `json:"text"`
	Dead        bool   `json:"dead"`
	Parent      int    `json:"parent"`
	Poll        int    `json:"poll"`
	Kids        []int  `json:"kids"`
	Url         string `json:"url"`
	Score       int    `json:"score"`
	Title       string `json:"title"`
	Parts       []int  `json:"parts"`
	Descendants int    `json:"descendants"`
}

// GetCodesStory -- Return the ids of a story
func (init Initializer) GetCodesStory() ([]int, error) {
	resp, errFetch := http.Get("https://hacker-news.firebaseio.com/v0/" + init.Story + ".json?print=pretty")
	if errFetch != nil {
		return nil, fmt.Errorf("Error while fetching codes story : %v", errFetch)
	}
	defer resp.Body.Close()
	body, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		return nil, fmt.Errorf("Error while reading Body of codes story : %v", errRead)
	}
	jsonDecoded := []int{}
	errDecoded := json.Unmarshal(body, &jsonDecoded)
	if errDecoded != nil {
		return nil, fmt.Errorf("Error while decoding json from codes story : %v", errDecoded)
	}
	return jsonDecoded, nil
}

// GetPostStory -- Return the posts of story thanks their ids
func (init Initializer) GetPostStory(codes []int) ([]Post, error) {
	post := Post{}
	myData := []Post{}
	for i := 0; i < init.NbPosts; i++ {
		resp, errFetch := http.Get("https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(codes[i]) + ".json?print=pretty")
		if errFetch != nil {
			return nil, fmt.Errorf("Error while fetching posts story : %v", errFetch)
		}
		body, errRead := ioutil.ReadAll(resp.Body)
		if errRead != nil {
			return nil, fmt.Errorf("Error while reading Body of posts story : %v", errRead)
		}
		errDecoded := json.Unmarshal(body, &post)
		if errDecoded != nil {
			return nil, fmt.Errorf("Error while decoding json from posts story : %v", errDecoded)
		}
		myData = append(myData, post)
	}
	return myData, nil
}
