## Пример использования 

```go
package main

import (
        	"fmt"
)

type Whitelist string
type Horn      string

func (w Whitelist) Makesong() {
	  fmt.Println("Tweet White list...")
}

func (w Horn) Makesong() {
	  fmt.Println("Tweet Horn...")
}


type Noise interface{
	 Makesong()
}

func Play(n Noise){
     n.Makesong()
}

func main(){
     Play(Whitelist("W list")) 
     Play(Horn("Horn list")) 
}


func main_noise(){
	var toy Noise

	toy = Whitelist("Test Whilist")
	toy.Makesong()

	toy = Horn("Test Horn")
	toy.Makesong()
}
```
