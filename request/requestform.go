Making a POST Request
Similarly, we can also make a POST request with body params to create a post using the API.

package main
import (
 "encoding/json"
 "io/ioutil"
 "log"
 "net/http"
 "net/url"
)
type Post struct {
 Userid string `json:"userId"`
 ID     int    `json:"id"`
 Title  string `json:"title"`
 Body   string `json:"body"`
}
func main() {
  params := url.Values{}
  params.Add("title", "foo")
  params.Add("body", "bar")
  params.Add("userId", "1")
resp, err := http.PostForm("https://jsonplaceholder.typicode.com/posts",
  params)
 if err != nil {   
   log.Printf("Request Failed: %s", err)
   return
 }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  // Log the request body 
  bodyString := string(body)
  log.Print(bodyString)
  // Unmarshal result
  post := Post{}
  err = json.Unmarshal(body, &post)
  if err != nil {
     log.Printf("Reading body failed: %s", err)
     return
  }
  
  log.Printf("Post added with ID %d", post.ID)
}
