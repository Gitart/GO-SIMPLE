Package main
import (
“fmt”
“regexp”
“strconv”
)
func main() {
// string to search
searchIn := “John: 2578.34 William: 4567.23 Steve: 5632.18”
pat := “[0-9]+.[0-9]+” // pattern to search for in searchIn
f := func (s string) string {
v, _ := strconv.ParseFloat(s, 32)
200
Ivo Balbaert
return strconv.FormatFloat(v * 2, ‘f’, 2, 32)
}
if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
fmt.Println(“Match found!”)
}
re, _ := regexp.Compile(pat)
// replace pat with “##.#”
str := re.ReplaceAllString(searchIn, “##.#”)
fmt.Println(str)
// using a function :
str2 := re.ReplaceAllStringFunc(searchIn, f)
fmt.Println(str2)
}
