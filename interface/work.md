## Sample work with interface

```golang
package main

import "fmt"

type Ui struct{
Name, Location string

}

type User struct {
Id int
Name, Location string
ss []int
kk []Ui
}

type Player struct {
User
GameId int
fff [] string
ff []int

}

func main() {
p := Player{}
p.Id = 42
p.Name = "Matt"
p.Location = "LA"
p.GameId = 90404
p.fff=[]string{"fff","ggg"}
p.ff=[]int{12,444,555,3332}
p.User.ss=[]int{12,444,225,333211}
// p.User.kk[0]=Ui{"ddd","fff"}
// p.User.kk[1]=Ui{"ddd","fff"}
p.User.kk=[]Ui{{"ddd","fff"},{"ddd","fff"},{"ddd","fff"},{"ddd","fff"},{"ddd","fff"},{"ddd","fff"}}

fmt.Printf("%+v", p)
}
```


## Samples second


```golang
package main

import (
  "fmt"
	// "encoding/json"
)



type El struct {
   Name string
   Age  int
}


type Et struct {
   Name string
   Id   string
   Last string 
   Date string
   Age  int
   Els  El
   Tags []string
}



func main() {
   var T Et   = Et{"One","A1002","aerrt","2007-303", 12, El{"dddd",23}, []string{"a","b","df"}}
   var Z []Et = []Et{
   	                  Et{"Two", "A1003", "aerrt","2007-303", 14,  El{"Age",33}, []string{"a0","b4","df"}}, 
   	                  Et{"Tree","A1004", "aerrt","2007-303", 142, El{"Age",45}, []string{"a1","b3,23,34","df"}},
   	              } 

   fmt.Println(T)
   fmt.Println(Z[1].Age)
   fmt.Println(Z[1].Els.Age)
   fmt.Println(Z[1].Els.Name)

   Z[1].Els.Name="Replaced"
   Cp:=len(Z)

   fmt.Println("Count:===",Cp)   
   fmt.Println(Z[1].Els.Name)
   

   fmt.Println("Все теги :", Z[1].Tags)
   fmt.Println("Один тег :", Z[1].Tags[2])

   for _, t:=range(Z){
   	   fmt.Println(t)
   	   fmt.Println("    Name:", t.Name)
   }

   fmt.Println(Z)
}
```
