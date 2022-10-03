package main

import (
  "encoding/xml"
  "github.com/russross/blackfriday"
  "net/http"
  "time"
)

func main() {
  http.HandleFunc("/", rssHandler)
  http.ListenAndServe(":3000", nil)
}

func rssHandler(w http.ResponseWriter, r *http.Request) {
  type Item struct {
    Title       string `xml:"title"`
    Link        string `xml:"link"`
    Description string `xml:"description"`
    PubDate     string `xml:"pubDate"`
  }

  type rss struct {
    Version     string `xml:"version,attr"`
    Description string `xml:"channel>description"`
    Link        string `xml:"channel>link"`
    Title       string `xml:"channel>title"`

    Item []Item `xml:"channel>item"`
  }

  articles := []Item{
    {"Bold style", "http://mywebsite.com/foo", "**lorem ipsum with a bold style**", time.Now().Format(time.RFC1123Z)},
    {"Italic style", "http://mywebsite.com/foo2", "*lorem ipsum with an italic style*", time.Now().Format(time.RFC1123Z)}}

  data := []Item{}
  for _, article := range articles {
    Title := article.Title
    Link := article.Link
    Description := string(blackfriday.MarkdownCommon([]byte(article.Description)))
    PubDate := article.PubDate
    data = append(data, Item{Title, Link, Description, PubDate})
  }

  feed := &rss{
    Version:     "2.0",
    Description: "My super website",
    Link:        "http://mywebsite.com",
    Title:       "Mywebsite",
    Item:        data,
  }

  x, err := xml.MarshalIndent(feed, "", "  ")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/xml")
  w.Write(x)
}
