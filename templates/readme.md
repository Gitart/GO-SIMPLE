# Summary
In this chapter, I described the standard library for creating HTML and text templates. The templates can
contain a wide range of actions, which are used to include content in the output. The syntax for templates
can be awkward—and care must be taken to express the content exactly as the template engine requires—
but the template engine is flexible and extensible and, as I demonstrate in Part 3, can easily be modified to
alter its behavior

```go
package main
import (
 "text/template"
 "os"
 "strings"
)
func GetCategories(products []Product) (categories []string) {
 catMap := map[string]string {}
 for _, p := range products {
 if (catMap[p.Category] == "") {
 catMap[p.Category] = p.Category
 categories = append(categories, p.Category)
 }
 }
 return
}
func Exec(t *template.Template) error {
 productMap := map[string]Product {}
 for _, p := range Products {
 productMap[p.Name] = p
 }
 return t.Execute(os.Stdout, &productMap)
}
func main() {
 allTemplates := template.New("allTemplates")
 allTemplates.Funcs(map[string]interface{} {
 "getCats": GetCategories,
 "lower": strings.ToLower,
 })
 allTemplates, err := allTemplates.ParseGlob("templates/*.txt")
 if (err == nil) {
 selectedTemplated := allTemplates.Lookup("mainTemplate")
 err = Exec(selectedTemplated)
 }
 if (err != nil) {
 Printfln("Error: %v %v", err.Error())
 }
}
```
