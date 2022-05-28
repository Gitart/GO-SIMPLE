// https://www.geeksforgeeks.org/function-as-a-field-in-golang-structure/

package main
import "fmt"

type Tet func(int, int) int

type Terik struct {
	Title      string
	Tarticles  int
	Fn         Tet
}


func (t *Terik) Add() {
	fn := func(Ma int, pay int) int {
          return Ma * pay*18289
    }

      t.Title     = "Total number of published articles"
      t.Tarticles = 123
      t.Fn        = fn
}


// Function with structure
func main() {

	fn := func(Ma int, pay int) int {
          return Ma * pay
    }

    fn = func(Ma int, pay int) int {
         return Ma * pay * 122
    }

   V:= Terik {
   	 Title : "Описание наименования систем...",
   	 Fn    : fn ,
   }

   fmt.Println("Function  : ", V.Title)
   V.Title = "Описание наименования"
   V.Fn(111,33)

   fmt.Println("Function  : ", V.Title)
   fmt.Println("Function  : ", V.Fn(111,33))
   fmt.Println("Результат : ", V.Fn(11123,3453))
   V.Add()

   fmt.Println("Результат 3 : ", V.Fn(11123,3453))
   fmt.Println("Function  3 : ", V.Title)

   Tyr(1)
   Tyr(2)

   Userss() 
}

// Used point
func Tyr(ty int ){
	fmt.Println("Pending article")
	if ty==1{
		goto tets
	}
	 
	return
	tets:
	fmt.Println("Before test ....")
}


type User struct {
    name       string
    occupation string
    country    string
}

// Users array
func Userss() {

    users := []User{

        {"John Doe", "gardener", "USA"},
        {"Roger Roe", "driver", "UK"},
        {"Paul Smith", "programmer", "Canada"},
        {"Lucia Mala", "teacher", "Slovakia"},
        {"Patrick Connor", "shopkeeper", "USA"},
        {"Tim Welson", "programmer", "Canada"},
        {"Tomas Smutny", "programmer", "Slovakia"},
    }

     fmt.Println(users)

    for _, user := range users {
        fmt.Println(user.name)
    }
 }
