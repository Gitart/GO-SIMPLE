package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func download(url string) error {
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	filename := path.Base(url)
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func worker(q chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		url, ok := <-q
		if !ok {
			return
		}
		if err := download(url); err == nil {
			log.Printf("downloaded %s", url)
		} else {
			log.Println(err)
		}
	}
}

func main() {
	doc, err := goquery.NewDocument("https://www.asahi.com/articles/ASL4J669JL4JUEHF016.html")
	if err != nil {
		log.Fatal(err)
	}

	q := make(chan string)

	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go worker(q, &wg)
	}

	doc.Find(".PlainMod .title a").Each(func(n int, qs *goquery.Selection) {
		if a, ok := qs.Attr("href"); ok && strings.HasSuffix(a, ".pdf") {
			q <- a
		}
	})

	close(q)
	wg.Wait()
}
