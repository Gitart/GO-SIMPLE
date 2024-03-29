# Use uint8

**Применение** : Может использоваться для проверки принадлдежности к опредленным группам знаков

Проверка символов на приндлежность к опредленным категориям
Проверка произвождится через использование uint8.
Есть два способа преобразовать символьную переменную в uint8.


Для получения значения в числовом формате используются синтаксис
```go
     fmt.Println('s') 
```
Output: 115



Обратите внимание что используются одинарные кавычки.
Например при сравненни и анализе 
```go
  s[2] == '}'
```


## Samples
```go
// getShellName returns the name that begins the string and the number of bytes
// consumed to extract it. If the name is enclosed in {}, it's part of a ${}
// expansion and two more bytes are needed than the length of the name.
func getShellName(s string) (string, int) {
	switch {
	case s[0] == '{':
		if len(s) > 2 && isShellSpecialVar(s[1]) && s[2] == '}' {
			return s[1:2], 3
		}
		// Scan to closing brace
		for i := 1; i < len(s); i++ {
			if s[i] == '}' {
				if i == 1 {
					return "", 2 // Bad syntax; eat "${}"
				}
				return s[1:i], i + 1
			}
		}
		return "", 1 // Bad syntax; eat "${"
	case isShellSpecialVar(s[0]):
		return s[0:1], 1
	}
	// Scan alphanumerics.
	var i int
	for i = 0; i < len(s) && isAlphaNum(s[i]); i++ {
	}
	return s[:i], i
}
```


1. Спопоб 
```go
newStr := []uint8("c")
```

2. Способ
```go
vr:="c"
newStr := vr[0]
```

В следующих шагах можно проверять их на принадлежность определенной группе
```go
// isAlphaNum reports whether the byte is an ASCII letter, number, or underscore
func isAlphaNum(c uint8) bool {
	return c == '_' || '0' <= c && c <= '9' || 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
}
```


## Samples
```go
func IsWordBreak(i rune) bool {
	switch {
	case i >= 'a' && i <= 'z':
	case i >= 'A' && i <= 'Z':
	case i >= '0' && i <= '9':
	default:
		return true
	}
	return false
}
```

## Полный пример
```go
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	vr := "g"
	newStr := []uint8("*")
	fmt.Println(newStr)
	c := isAlphaNum(newStr[0])
	fmt.Println(c, 's')
	f := isAlphaNum(vr[0])
	fmt.Println(f)

}

// isAlphaNum reports whether the byte is an ASCII letter, number, or underscore
func isAlphaNum(c uint8) bool {
	return c == '_' || '0' <= c && c <= '9' || 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
}

// isShellSpecialVar reports whether the character identifies a special
// shell variable such as $*.
func isShellSpecialVar(c uint8) bool {
	switch c {
	case '*', '#', '$', '@', '!', '?', '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	}
	return false
}
```

// Sample
```go
package main

import "fmt"

func main() {
	var htmlNospaceReplacementTable = []string{
		0:    "&#xfffd;",
		'\t': "&#9;",
		'\n': "&#10;",
		'\v': "&#11;",
		'\f': "&#12;",
		'\r': "&#13;",
		' ':  "&#32;",
		'"':  "&#34;",
		'&':  "&amp;",
		'\'': "&#39;",
		'+':  "&#43;",
		'<':  "&lt;",
		'=':  "&#61;",
		'>':  "&gt;",
		// A parse error in the attribute value (unquoted) and
		// before attribute value states.
		// Treated as a quoting character by IE.
		'`': "&#96;",
	}
	fmt.Println(htmlNospaceReplacementTable['\t'])
}
```
## Примеры
[Пример использования](https://go.dev/play/p/YgXj4xzaGZL)   
[Пример](https://www.socketloop.com/tutorials/golang-convert-cast-string-to-uint8-type-and-back-to-string)   
[Использование в стандартной библиотеке](https://go.dev/src/os/env.go)    
