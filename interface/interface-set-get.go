package main
import "fmt"

// *******************************************************
// Cache (BASIC INTERACE)
// *******************************************************
type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
}

// *******************************************************
// (STRUCT)
// *******************************************************
type Repository struct {
     container  map[string]interface{}
}

func NewRepository() *Repository {
	return &Repository{make(map[string]interface{})}
}

func (this *Repository) Set(key string, value interface{}) {
	 this.container[key] = value
}

func (this *Repository) Get(key string) interface{} {
	 return this.container[key]
}

// *******************************************************
// (STRUCT)
// *******************************************************
func NewSt() *St {
	return &St{make(map[string]interface{})}
}

// Kv 
type St struct {
     Kv map[string]interface{}
}

// Set 
func (m *St) Set(Key string, Value interface{}){
	 m.Kv[Key] = Value
}

func (m *St) Get(key string) interface{} {
	 return m.Kv[key]
}

// Get All
func (m *St) GetAll() *St {
	 return m
}

// *****************************************
// Preview work interface
// *****************************************
func All(c Cache) {
	fmt.Println("    >:", c)

    fmt.Println("tel  :----------->" , c.Get("te"))
	fmt.Println("pr   :----------->" , c.Get("pr"))
	fmt.Println("prom :----------->" , c.Get("prom"))
	fmt.Println("\n\n")
}

// *****************************************
// Main process 
// *****************************************
func main() {

    ss := NewSt()
    ss.Set("te",     "tel  |   telephone")

    srep := NewRepository()
    srep.Set("pr",   "pr   | printer")
    srep.Set("prom", "prom | Telegraph")
	
	All(srep)
	All(ss)
}


// *****************************************
// Additional Information
// *****************************************
func Info(){
    
    // Repository
    rep := NewRepository()
    rep.Set("News",     "Новостная лента")
    rep.Set("Old",      "Old news ")
    rep.Set("Forecast", "Prognoz")
    frep := rep.Get("Old")

    fmt.Println(frep)
    fmt.Println("-----------------------")

    // New Stat
	sr := NewSt()
	sr.Set("London", "England")
	sr.Set("Paris",  "France")
	fmt.Println("Country ", sr.Get("London"))

	sr.Set("tts34", "Элемент который знается")
	fmt.Println("Получение 34 елемента    : ", sr.Get("tts34"))
	fmt.Println("Получение всех елементов : ", sr.GetAll())
	fmt.Println("Получение всех елементов : ", *sr.GetAll() )
	fmt.Println("Количество елементов     : ", len(sr.GetAll().Kv) )
    
    dd:=sr.GetAll()
    fmt.Println(dd.Kv)
    
    // Range
    for k, r := range dd.Kv {
    	fmt.Println(k, "=", r)
    }
}
