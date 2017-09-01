## Sample Service for Gorilla mux
[Youtube](https://www.youtube.com/watch?v=t96hBT53S4U&t=174s)    
[Sample](https://github.com/HakaseLabs/source-blog/blob/master/rest-api/main.go)


```golang



package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
)


type Person struct {
	ID        string   `json:"id,omitempty`
	Firstname string   `json:"firstname,omitempty`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person



// При добавлении пользователя все данные храняться в памяти - пока работает сервис
// После выключения сервиса все ломается
// 
func main() {

	fmt.Println("Start...")
	router := mux.NewRouter()

    // Изначальное заполнение массива данными для примера
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})

    // Roter
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")                  // Получение всех пользователей
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")             // Получение сведений по ид пользователя
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")         // Добавление пользователя указав его ИД и в боди передав JSON
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")       // Удаление всех из   
	
	// Log & run server
	log.Fatal(http.ListenAndServe(":8000", router))
}


// Получение всех пользователей
func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

 // Получение сведений по ид пользователя
func GetPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}


   // Добавление пользователя указав его ИД и в боди передав JSON
func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people    = append(people, person)
	json.NewEncoder(w).Encode(people)
}


// Удаление всех из массива
func DeletePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}
```


