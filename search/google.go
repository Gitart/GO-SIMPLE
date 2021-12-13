package main

import (
	"fmt"
	"context"
	"github.com/rocketlaunchr/google-search"
)

//  Main process
func main(){

      Search("Стартап, news", 10)
      Search("IT, news", 10)
      Search("Новости, В мире", 10)
      Search("Автомобили, последние новости",5)
      Search("Технологии, последние новости",3)
}

// Search news in Google
func Search(txt string, cnt int) {

    var ctx = context.Background()
    opts := googlesearch.SearchOptions{Limit: cnt}
    dat,_:= googlesearch.Search(ctx, txt, opts)

    fmt.Println("***************",txt, "************************")    
    for _, rd:=range dat{
         fmt.Println(rd.Rank, rd.URL)
         fmt.Println(rd.Title)
         fmt.Println(rd.Description)
         fmt.Println("")
   }

   fmt.Println("..........................................................................")
   fmt.Println("")
}
