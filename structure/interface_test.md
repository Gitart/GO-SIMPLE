# Samples



### Interfaces
```golang
package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    // ********************* Marshal *********************
    u := map[string]interface{}{}
    u["name"] = "kish"
    u["age"] = 28
    u["work"] = "engine"
    //u["hobbies"] = []string{"art", "football"}
    u["hobbies"] = "art"

    b, err := json.Marshal(u)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(b))

    // ********************* Unmarshal *********************
    var a interface{}
    err = json.Unmarshal(b, &a)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Println(a)
}
```

### Templates
```golang
{{ $TotalPrice := 0.0 }}
{{ range $i, $tx := .Transactions }}
{{ $TotalPrice := FloatInc $TotalPrice (StrToFloat .TotalPrice) }}
  <tr>
    <td>{{ inc $i 1 }}</td> 
    <td>{{ .Description.String }}</td>
    <td>{{ .Type }}</td>
    <td>{{ .TotalPrice }}</td>
    <td>{{ .Note }}</td>  
  </tr>  
{{ end }}
<tr>
  <td></td> 
  <td></td>
  <td></td>
  <td>{{ $TotalPrice }}</td>
  <td></td>
  <td></td>
</tr> 
Transactions are Money Transaction with TotalPrice DB Fields and I have 4 functions according Iris framework spec.

tmpl.AddFunc("dec", func(num int, step int) int {
    return num - step
})

tmpl.AddFunc("inc", func(num int, step int) int {
    return num + step
})

tmpl.AddFunc("FloatDec", func(num float64, step float64) float64 {
    return num - step
})

tmpl.AddFunc("FloatInc", func(num float64, step float64) float64 {
    return num + step
})

tmpl.AddFunc("StrToFloat", func(s string) (float64, error) {
    return strconv.ParseFloat(s, 64)
}) 
```






