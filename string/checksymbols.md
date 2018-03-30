## Chrck Symbols in string


```golang
package main

var allowed = []rune{'a','b','c','d','e','f','g'}

func haveSpecial(input string) bool {
	for _, char := range input {
		found := false
		for _, c := range allowed {
			if c == char {
				found = true
				break
			}
		}
		if !found {
			return true
		}
	}
	return false
}

func main() {
	cases := []string{"abcdef","abc$€f"}
	
	for _, input := range cases {
		if haveSpecial(input) {
			println(input + ": NOK")
		} else {
			println(input + ": OK")
		}
	}
}
```

