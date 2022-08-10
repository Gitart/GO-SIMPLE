package main

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

func main() {
	tournaments := []struct {
		Place string
		Date  time.Time
	}{
		// for clarity - date is sorted, we don't need sort it again
		{"Town1", time.Date(2015, time.November, 10, 23, 0, 0, 0, time.Local)},
		{"Town2", time.Date(2015, time.October, 10, 23, 0, 0, 0, time.Local)},
		{"Town3", time.Date(2014, time.November, 10, 23, 0, 0, 0, time.Local)},
	}
	t, err := template.New("").Parse(`
{{$prev_year:=0}}
{{range .}}
	{{with .Date}}
		{{$year:=.Year}}
            		{{if ne $year $prev_year}}
            			Actions in year {{$year}}:
           		{{$prev_year:=$year}}
        	{{end}}
	{{end}}
        
        {{.Place}}, {{.Date}}
    {{end}}

	`)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, tournaments)
	if err != nil {
		fmt.Println("executing template:", err)
	}
}
