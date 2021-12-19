package main
import (
  "fmt"
  "sync"
)
type Fetcher interface {
  // Fetch returns the body of URL and
  // a slice of URLs found on that page.
  Fetch(url string) (body string, urls []string, err error)
}
var cache = make(Cache)
var wg sync.WaitGroup
var mux sync.Mutex
// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
  defer wg.Done()
  if cache.get(url) {
    fmt.Printf("xx Skipping: %s\n", url)
    return
  }
  fmt.Printf("** Crawling: %s\n", url)
  cache.set(url, true)

  if depth <= 0 {
    return
  }
  body, urls, err := fetcher.Fetch(url)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Printf("found: %s %q\n", url, body)
  for _, u := range urls {
    wg.Add(1)
    go Crawl(u, depth-1, fetcher)
  }
  return
}
func main() {
  wg.Add(1)
  Crawl("https://golang.org/", 4, fetcher)
  wg.Wait()
}
type Cache map[string]bool
func (ch Cache) get(key string) bool {
  mux.Lock()
  defer mux.Unlock()
  return cache[key]
}
func (ch Cache) set(key string, value bool) {
  mux.Lock()
  defer mux.Unlock()
  cache[key] = value
}
// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult
type fakeResult struct {
  body string
  urls []string
}
func (f fakeFetcher) Fetch(url string) (string, []string, error) {
  if res, ok := f[url]; ok {
    return res.body, res.urls, nil
  }
  return "", nil, fmt.Errorf("not found: %s", url)
}
// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
  "https://golang.org/": &fakeResult{
   "The Go Programming Language",
  []string{
   "https://golang.org/pkg/",
   "https://golang.org/cmd/",
  },
 },
 "https://golang.org/pkg/": &fakeResult{
  "Packages",
  []string{
   "https://golang.org/",
   "https://golang.org/cmd/",
   "https://golang.org/pkg/fmt/",
   "https://golang.org/pkg/os/",
  },
 },
 "https://golang.org/pkg/fmt/": &fakeResult{
  "Package fmt",
  []string{
   "https://golang.org/",
   "https://golang.org/pkg/",
  },
 },
 "https://golang.org/pkg/os/": &fakeResult{
  "Package os",
  []string{
   "https://golang.org/",
   "https://golang.org/pkg/",
  },
 },
}
