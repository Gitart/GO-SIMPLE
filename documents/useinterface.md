## Использование интерфейса

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
