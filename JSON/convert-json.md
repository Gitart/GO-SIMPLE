## Convert To Company

```go
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type SearchResult struct {
	Date        string      `json:"date"`
	IdCompany   int         `json:"idCompany"`
	Company     string      `json:"company"`
	IdIndustry  interface{} `json:"idIndustry"`
	Industry    string      `json:"industry"`
	IdContinent interface{} `json:"idContinent"`
	Continent   string      `json:"continent"`
	IdCountry   interface{} `json:"idCountry"`
	Country     string      `json:"country"`
	IdState     interface{} `json:"idState"`
	State       string      `json:"state"`
	IdCity      interface{} `json:"idCity"`
	City        string      `json:"city"`
}

func fieldSet(fields ...string) map[string]bool {
	set := make(map[string]bool, len(fields))
	for _, s := range fields {
		set[s] = true
	}
	return set
}

func (s *SearchResult) SelectFields(fields ...string) map[string]interface{} {
	fs := fieldSet(fields...)
	rt, rv := reflect.TypeOf(*s), reflect.ValueOf(*s)
	out := make(map[string]interface{}, rt.NumField())
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		jsonKey := field.Tag.Get("json")
		if fs[jsonKey] {
			out[jsonKey] = rv.Field(i).Interface()
		}
	}
	return out
}

func main() {
	result := &SearchResult{
		Date:     "to be honest you should probably use a time.Time field here, just sayin",
		Industry: "rocketships",
		IdCity:   "interface{} is kinda inspecific, but this is the idcity field",
		City:     "New York Fuckin' City",
	}
	b, err := json.MarshalIndent(result.SelectFields("idCity", "city", "company"), "", "  ")
	if err != nil {
		panic(err.Error())
	}
	fmt.Print(string(b))
}
```
