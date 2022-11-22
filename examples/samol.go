package main

import "os"
import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/json"

type Tracks struct {
    Toptracks []Toptracks_info
}

type Toptracks_info struct {
    Track []Track_info
    Attr  []Attr_info
}

type Track_info struct {
    Name       string
    Duration   string
    Listeners  string
    Mbid       string
    Url        string
    Streamable []Streamable_info
    Artist     []Artist_info
    Attr       []Track_attr_info
}

type Attr_info struct {
    Country    string
    Page       string
    PerPage    string
    TotalPages string
    Total      string
}

type Streamable_info struct {
    Text      string
    Fulltrack string
}

type Artist_info struct {
    Name string
    Mbid string
    Url  string
}

type Track_attr_info struct {
    Rank string
}

func get_content() {
    // json data
    url := "http://ws.audioscrobbler.com/2.0/?method=geo.gettoptracks&api_key=c1572082105bd40d247836b5c1819623&format=json&country=Netherlands"

    res, err := http.Get(url)

    if err != nil {
        panic(err.Error())
    }

    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
        panic(err.Error())
    }

    var data Tracks
    json.Unmarshal(body, &data)
    fmt.Printf("Results: %v\n", data)
    os.Exit(0)
}

func main() {
    get_content()
}





var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
    r, err := myClient.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}
Example use:

type Foo struct {
    Bar string
}

func main() {
    foo1 := new(Foo) // or &Foo{}
    getJson("http://example.com", foo1)
    println(foo1.Bar)

    // alternately:

    foo2 := Foo{}
    getJson("http://example.com", &foo2)
    println(foo2.Bar)
}
