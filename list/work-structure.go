package main

 import (
 	"fmt"
 	"strings"
 )

 type citySize struct {
 	Name       string
 	Population float64
 }

 // cities with the largest population in the world
 var citiesList = []citySize{
 	{"Tokyo", 38001000},
 	{"Delhi", 25703168},
 	{"Shanghai", 23740778},
 	{"Sao Paulo", 21066245},
 }

 func findRecordsByCityName(name string) bool {
 	for _, v := range citiesList {
 		// convert to lower case for exact matching
 		if strings.ToLower(v.Name) == strings.ToLower(name) {
 			return true
 		}
 	}
 	return false
 }

 func returnRecordsByCityName(name string) citySize {
 	for _, v := range citiesList {
 		// convert to lower case for exact matching
 		if strings.ToLower(v.Name) == strings.ToLower(name) {
 			return v
 		}
 	}
 	return citySize{"", 0}
 }

 func main() {

 	if !findRecordsByCityName("New York") {
 		fmt.Println("New York is not in the list!")
 	}

 	// Check if Tokyo is in the list
 	cityData := returnRecordsByCityName("Tokyo")
 	if cityData.Name != "" {
 		fmt.Println(cityData)
 	}

 	// Check if Teluk Intan is in the list
 	cityData = returnRecordsByCityName("Teluk Intan")
 	if cityData.Name != "" {
 		fmt.Println(cityData)
 	} else {
 		fmt.Println("Teluk Intan is not in the list")
 	}

 }

// Output:

//  New York is not in the list!
//  {Tokyo 3.8001e+07}
//  Teluk Intan is not in the list
