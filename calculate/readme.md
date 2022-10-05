# Calculate with subtotal sum

### Figure
For the example application, this means that the map of products is processed sequentially so that each
product category is processed in turn, and, within each category, each product is processed, as illustrated by

![image](https://user-images.githubusercontent.com/3950155/194041223-35e6cf22-e171-403b-a133-b5016863dc0d.png)

## Output
![image](https://user-images.githubusercontent.com/3950155/194041589-e34f288e-6ceb-43f6-92cf-a26c4aadb2d0.png)


```go
package main

import (
 "fmt"
 "time"
 "strconv"
)

// Structure
type Product struct {
    Name, Category string
    Price float64
}

var ProductList = []*Product {
 { "Kayak",            "Watersports", 279.60 },
 { "Lifejacket",       "Watersports", 49.95 },
 { "Soccer Ball",      "Soccer",      19.50 },
 { "Corner Flags",     "Soccer",      34.95 },
 { "Stadium",          "Soccer",      79500 },
 { "Thinking Cap",     "Chess",       16.50 },
 { "Unsteady Chair",   "Chess",       75.88 },
 { "Bling-Bling King", "Chess",       1200  },
}

type ProductGroup []*Product
type ProductData  = map[string]ProductGroup
var  Products     = make(ProductData)

//-------------------------------------------------------
// ■ Main
//-------------------------------------------------------
func main() {

 for _, p := range ProductList {
    
    if _, ok := Products[p.Category]; ok {
        Products[p.Category] = append(Products[p.Category], p)
    } else {
        Products[p.Category] = ProductGroup{p}
    }
 }

 // fmt.Println(Products)
 CalcStoreTotal(Products)
}

//-------------------------------------------------------
// ■ Currency
//------------------------------------------------------- 
func ToCurrency(val float64) string {
     return "$" + strconv.FormatFloat(val, 'f', 2, 64)
}

//-------------------------------------------------------
// ■ Total sum
//-------------------------------------------------------
func CalcStoreTotal(data ProductData) {
 var storeTotal float64
 
 for category, group := range data {
     storeTotal += group.TotalPrice(category)
 }

 fmt.Println("Total:", ToCurrency(storeTotal))
}

//-------------------------------------------------------
// ■ Total price
//-------------------------------------------------------
func (group ProductGroup) TotalPrice(category string, ) (total float64) {
     for _, p := range group {
         total += p.Price
         time.Sleep(time.Millisecond * 100)
     }

     fmt.Println(category, "subtotal:", ToCurrency(total))
     return
}
```

# Calculate from chanels 
## Receiving a Result Using a Channel

```go
package main

import (
 "fmt"
 "time"
 "strconv"
)

// Structure
type Product struct {
    Name, Category string
    Price float64
}

var ProductList = []*Product {
 { "Kayak",            "Watersports", 279.60 },
 { "Lifejacket",       "Watersports", 49.95 },
 { "Soccer Ball",      "Soccer",      19.50 },
 { "Corner Flags",     "Soccer",      34.95 },
 { "Stadium",          "Soccer",      79500 },
 { "Thinking Cap",     "Chess",       16.50 },
 { "Unsteady Chair",   "Chess",       75.88 },
 { "Bling-Bling King", "Chess",       1200  },
}

type ProductGroup []*Product
type ProductData  = map[string]ProductGroup
var  Products     = make(ProductData)

//-------------------------------------------------------
// ■ Main
//-------------------------------------------------------
func main() {

 for _, p := range ProductList {
    
    if _, ok := Products[p.Category]; ok {
        Products[p.Category] = append(Products[p.Category], p)
    } else {
        Products[p.Category] = ProductGroup{p}
    }
 }

 // fmt.Println(Products)
 CalcStoreTotal(Products)

 // Remove before goroutine 
 // time.Sleep(time.Second * 1)
}

//-------------------------------------------------------
// ■ Currency
//------------------------------------------------------- 
func ToCurrency(val float64) string {
     return "$" + strconv.FormatFloat(val, 'f', 2, 64)
}

//-------------------------------------------------------
// ■ Total sum
//-------------------------------------------------------
func CalcStoreTotal(data ProductData) {
 var storeTotal float64
  var channel chan float64 = make(chan float64)

 for category, group := range data {
     // storeTotal += group.TotalPrice(category)
     // go group.TotalPrice(category)
     go group.TotalPrice(category, channel)
 }

 for i := 0; i < len(data); i++ {
     storeTotal += <- channel
 }

 fmt.Println("Total:", ToCurrency(storeTotal))
}

//-------------------------------------------------------
// ■ Using Goroutines and Channels
//-------------------------------------------------------
// func (group ProductGroup) TotalPrice(category string, ) (total float64) {
func (group ProductGroup) TotalPrice(category string, resultChannel chan float64) {
      var total float64
     for _, p := range group {
         total += p.Price
         time.Sleep(time.Millisecond * 100)
     }

     fmt.Println(category, "subtotal:", ToCurrency(total))
     resultChannel <- total
     // return
}



```



