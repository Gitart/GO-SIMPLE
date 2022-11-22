
```go
package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	a := strings.Fields(s)
	for _, v := range a {
		m[v]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
```


## Sample 2

```go
package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    set := strings.Fields(s)
    var st_map map[string]int = make(map[string]int)
    for i := range set {
        count := 0
        for j := range set {
            if set[i] == set[j] {
                count = count + 1
            }
        }
        st_map[set[i]] = count
    }
    return st_map
}

func main() {
    wc.Test(WordCount)
}
```

## Sample 3

```go
package main

import (
"golang.org/x/tour/wc"
"strings"
)

func WordCount(s string) map[string]int {
f := make(map[string]int)
ss:=strings.Fields(s)
for i :=range ss{
f[ss[i]]+=1
}
return f
}

func main() {
wc.Test(WordCount)
}
```


## @gajama
This works as well, but should've realised that the 'if' loop just increments!

```go
package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    m := make(map[string]int)
    words := strings.Fields(s)
    for _, w := range words {
        if _, in := m[w]; in {
            m[w] = m[w] + 1
        } else {
            m[w] = 1
        }
    }
    return m
}

func main() {
    wc.Test(WordCount)
}
```


## @othbert

```go
package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) (x map[string]int) {
    x = make(map[string]int)
    for _, word := range strings.Fields(s) {
        x[word]++
    }
    return 
}

func main() {
    wc.Test(WordCount)
}
```

 
## @ghost
It's works, but it's not an optimized code :

```go
package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    m_return := make(map[string]int)
    slc_s := strings.Fields(s)
    for _, v := range slc_s {
        _, ok := m_return[v]
        if ok {
            m_return[v] += 1
        } else {
            m_return[v] = 1
        }
    }
    return m_return
}

func main() {
    wc.Test(WordCount)
}
```

#With less code :

```go
package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) (x map[string]int) {
  x = make(map[string]int)
  for _, j := range strings.Fields(s) {
  x[j]++
  }
return
}

func main() {
    wc.Test(WordCount)
}
```

## @abjoker
A less elegant version

```go
package main
import (
	`"golang.org/x/tour/wc"`
	"strings"
)
        
`func` WordCount(s string) map[string]int {
	var str []string
	str=strings.Fields(s)
	b:=make(map[string]int)
	
 for i:=range str{
	n:=1	
	st:=str[i]
            
               //inner loop for matching previous words in the string
	        for j:=0 ; j<i ; j++ {  
		if (st==str[j]){     //if their is a previous occurence of the word
			n++
			b[str[i]]=n
		}
             }// end of inner loop
	     
                     if n==1{     //if word occurs for the first time
			b[str[i]]=n 
                          } 

	}//end of outer loop
	

    return b
 }
   
func main() {
	wc.Test(WordCount)
      }
 ```
 
 ##@espaciomore
```go
// gotta love the symplified version

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		//_, wordExists := wordMap[word]
		//if !wordExists {
		//	wordMap[word] = 1
		//} else {
			wordMap[word]++
		//}
	}
	return wordMap
}
```


## My Version with switch-case

```go
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	list := strings.Fields(s)
	for _, val := range list{
		switch _ , ok := ret[val]; ok{
		case ok:
			ret[val]++
		default:
			ret[val] = 1
		}
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
```

 ## @t1maccapp
Using "strings" in this exercise looks like cheating to me :bowtie:

```go
package main

import (
  "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
  wordCount := make(map[string]int)
  currentWord := ""
	
  for i, value := range(s) {
    currentChar := string(value)
	
    if currentChar == " " {
      if len(currentWord) > 0 {
        wordCount[currentWord] += 1
      }
	  
      currentWord = ""
    } else if i == len(s) - 1 {
      currentWord = currentWord + currentChar
	  
      if len(currentWord) > 0 {
        wordCount[currentWord] += 1
      }
    } else {
      currentWord = currentWord + currentChar
    }
  }
  
  return wordCount
}

func main() {
  wc.Test(WordCount)
}
```

## @viter

```go
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	a := strings.Fields(s)
	for i:=0;i<len(a);i++ {
		m[a[i]]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
```

 ## @cahitbeyaz
```go
 package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	
	var counterMap map[string]int
	splittedStrings:=strings.Fields(s)
	counterMap=make(map[string]int)
	for _,v:=range splittedStrings{
		counterMap[v]+=1
	}
	
	return counterMap
}

func main() {
	wc.Test(WordCount)
}
```

 ## csabakollar 
 ```go
 package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Split(s," ")
	m := make(map[string]int)
	for i:=0; i<len(words);i++ {
		m[words[i]] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
```

## Combining strings.Fields and for loop into one line:

```go
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCound(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range strings.Fields(s) {
		m[v]++
	}
	return m
}

func main() {
	wc.Test(WordCound)
}
```

## @nickkaczmarek


```go
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	listOfWords := strings.Fields(s)
	wordMap := make(map[string]int)
	
	for _, word := range(listOfWords) {
		if(wordMap[word] == 1) {
			wordMap[word] = 2
		} else {
			wordMap[word] = 1
		}
	}
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
```

## After reading through this, this is what I came to:
```go
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) (m map[string]int) {
	m = make(map[string]int)
	for _, w:= range(strings.Fields(s)) {
		m[w]++
	}
	return
}

func main() {
	wc.Test(WordCount)
}
```


## @BarryMcAuley

```go
 package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	for _, word := range strings.Fields(s) {
		counts[word]++
	}
	
	return counts
}

func main() {
	wc.Test(WordCount)
}
```

# @stevefoxuser

```go
package main

import (
"golang.org/x/tour/wc"
"strings"
)

func WordCount(s string) map[string]int {
m := make(map[string]int)
for _,v:= range strings.Fields(s) {
m[v] += 1
}
return m
}

func main() {
wc.Test(WordCount)
}
```

 ## @takirohit
 ```go
 package main


import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCount := make(map[string]int)
	
	for word := range words {		
		wordCount[words[word]] = wordCount[words[word]] + 1
	}
	return wordCount
}

func main() {
	wc.Test(WordCount)
}
```
