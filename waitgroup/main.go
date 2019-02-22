// https://www.youtube.com/watch?v=0BPSR-W4GSY&app=desktop

package main

import(
 "fmt"
 "log"
 "net/http"
)


var urls=[]string{
	"http://google.com",
	"https://tutorialage.com",
	"https://twitter.com",
}

func fetchStatus(w http.ResponseWrite, r *http.Reguest){
	for _, url :=range urls{
	wg.Add(1)
	    go func(url string){
	        resp, err:=http.Get(url)
	        if err!=nil{
               fmt.Println(w, "%v\n",err)     
	        }
	        fmt.Println(w, "%v\n",resp.Status)     // or resp
	        wg.Done()
	    }(url)
	}
	wg.Wait()
}



func main(){
	fmt.Println("Go WaitGroup Tutorial")
	http.HandleFunc("/",fethStaus)
	log.Fatal(http.ListenAndServ(":8080",nil))


}
