package main

import (
    "bytes"
    "code.google.com/p/go-charset/charset"
    _ "code.google.com/p/go-charset/data" // Import charset configuration files
    "encoding/xml"
    "io/ioutil"
    "log"
    "net/http"
    "testing"
)

type RssFeed struct {
    XMLName xml.Name  `xml:"rss"`
    Items   []RssItem `xml:"channel>item"`
}

type RssItem struct {
    XMLName     xml.Name `xml:"item"`
    Title       string   `xml:"title"`
    Link        string   `xml:"link"`
    Description string   `xml:"description"`
    //NestedTag    string      xml:">nested>tags>"
}

func fetchURL(url string) []byte {
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalf("unable to GET '%s': %s", url, err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("unable to read body '%s': %s", url, err)
    }
    return body
}

func parseXML(xmlDoc []byte, target interface{}) {
    reader := bytes.NewReader(xmlDoc)
    decoder := xml.NewDecoder(reader)
    // Fixes "xml: encoding \"windows-1252\" declared but Decoder.CharsetReader is nil"
    decoder.CharsetReader = charset.NewReader
    if err := decoder.Decode(target); err != nil {
        log.Fatalf("unable to parse XML '%s':\n%s", err, xmlDoc)
    }
}

func TestParseReport(t *testing.T) {
    var rssFeed = &RssFeed{}
    xmlDoc := fetchURL("https://news.ycombinator.com/rss")
    parseXML(xmlDoc, &rssFeed)
    for _, item := range rssFeed.Items {
        log.Printf("%s: %s", item.Title, item.Link)
    }
}
