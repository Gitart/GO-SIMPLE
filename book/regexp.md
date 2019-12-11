# Регулярное выражение (Regex) \- примеры и решения

Регулярное выражение (или короткое регулярное выражение) является очень полезным инструментом, который используется для описания шаблона поиска для сопоставления текста. Regex \- это не что иное, как последовательность некоторых символов, которая определяет шаблон поиска. Regex используется для анализа, фильтрации, проверки и извлечения значимой информации из большого текста, такого как журналы и выходные данные, генерируемые другими программами.

---

## Регулярное выражение для извлечения текста в квадратных скобках

//Regular expression to extract text between square brackets in Golang
```golang
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str1 := "this is a \[sample\] \[\[string\]\] with \[SOME\] special words"

	re := regexp.MustCompile(\`\\\[(\[^\\\[\\\]\]\*)\\\]\`)
	fmt.Printf("Pattern: %v\\n", re.String())      // print pattern
	fmt.Println("Matched:", re.MatchString(str1)) // true

	fmt.Println("\\nText between square brackets:")
	submatchall := re.FindAllString(str1, \-1)
	for \_, element := range submatchall {
		element = strings.Trim(element, "\[")
		element = strings.Trim(element, "\]")
		fmt.Println(element)
	}
}
```

Pattern: \\\[(\[^\\\[\\\]\]\*)\\\]
Matched: true
Text between square brackets:
sample
string
SOME

---

## Regular expression to extract all Non\-Alphanumeric Characters from a String
```golang
//Regular expression to extract all Non\-Alphanumeric Characters from a String

package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "We @@@Love@@@@ #Go!$! \*\*\*\*Programming\*\*\*\*Language^^^"

	re := regexp.MustCompile(\`\[^a\-zA\-Z0\-9\]+\`)

	fmt.Printf("Pattern: %v\\n", re.String()) // print pattern
	fmt.Println(re.MatchString(str1))        // true

	submatchall := re.FindAllString(str1, \-1)
	for \_, element := range submatchall {
		fmt.Println(element)
	}
}
```

Pattern: \[^a\-zA\-Z0\-9\]+
true
 @@@
@@@@ #
!$! \*\*\*\*
\*\*\*\*
^^^

---

## Regular expression to extract date(YYYY\-MM\-DD) from string

//Regular expression to extract date(YYYY\-MM\-DD) from string
```golang
package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "If I am 20 years 10 months and 14 days old as of August 17,2016 then my DOB would be 1995\-10\-03"

	re := regexp.MustCompile(\`\\d{4}\-\\d{2}\-\\d{2}\`)

	fmt.Printf("Pattern: %v\\n", re.String()) // print pattern

	fmt.Println(re.MatchString(str1)) // true

	submatchall := re.FindAllString(str1, \-1)
	for \_, element := range submatchall {
		fmt.Println(element)
	}
}
```
Pattern: \\d{4}\-\\d{2}\-\\d{2}
true
1995\-10\-03

---

## Regular expression to extract DNS host\-name or IP Address from string

//Regular expression to extract DNS hostname or IP Address from string
```golang
package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := \`Proxy Port Last Check Proxy Speed Proxy Country Anonymity 118.99.81.204
	118.99.81.204 8080 34 sec Indonesia \- Tangerang Transparent 2.184.31.2 8080 58 sec
	Iran Transparent 93.126.11.189 8080 1 min Iran \- Esfahan Transparent 202.118.236.130
	7777 1 min China \- Harbin Transparent 62.201.207.9 8080 1 min Iraq Transparent\`

	re := regexp.MustCompile(\`(25\[0\-5\]|2\[0\-4\]\[0\-9\]|\[01\]?\[0\-9\]\[0\-9\]?)(\\.(25\[0\-5\]|2\[0\-4\]\[0\-9\]|\[01\]?\[0\-9\]\[0\-9\]?)){3}\`)

	fmt.Printf("Pattern: %v\\n", re.String()) // print pattern
	fmt.Println(re.MatchString(str1)) // true

	submatchall := re.FindAllString(str1, \-1)
	for \_, element := range submatchall {
		fmt.Println(element)
	}
}
```

Pattern: (25\[0\-5\]|2\[0\-4\]\[0\-9\]|\[01\]?\[0\-9\]\[0\-9\]?)(\\.(25\[0\-5\]|2\[0\-4\]\[0\-9\]|\[01\]?\[0\-9\]\[0\-9\]?)){3}
true
118.99.81.204
118.99.81.204
2.184.31.2
93.126.11.189
202.118.236.130
62.201.207.9

---

## Regular expression to extract domain from URL

//Regular expression to extract domain from URL
```golang
package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := \`http://www.suon.co.uk/product/1/7/3/\`

	re := regexp.MustCompile(\`^(?:https?:\\/\\/)?(?:\[^@\\/\\n\]+@)?(?:www\\.)?(\[^:\\/\\n\]+)\`)
	fmt.Printf("Pattern: %v\\n", re.String()) // print pattern
	fmt.Println(re.MatchString(str1)) // true

	submatchall := re.FindAllString(str1,\-1)
	for \_, element := range submatchall {
		fmt.Println(element)
	}
}
```

Pattern: ^(?:https?:\\/\\/)?(?:\[^@\\/\\n\]+@)?(?:www\\.)?(\[^:\\/\\n\]+)
true
http://www.suon.co.uk

---

## Regular expression to validate email address

//Regular expression to validate email address
```golang
package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "ç$€§/az@gmail.com"
	str2 := "abcd@gmail\_yahoo.com"
	str3 := "abcd@gmail\-yahoo.com"
	str4 := "abcd@gmailyahoo"
	str5 := "abcd@gmail.yahoo"

	re := regexp.MustCompile("^\[a\-zA\-Z0\-9.!#$%&'\*+/=?^\_\`{|}~\-\]+@\[a\-zA\-Z0\-9\](?:\[a\-zA\-Z0\-9\-\]{0,61}\[a\-zA\-Z0\-9\])?(?:\\\\.\[a\-zA\-Z0\-9\](?:\[a\-zA\-Z0\-9\-\]{0,61}\[a\-zA\-Z0\-9\])?)\*$")

	fmt.Printf("Pattern: %v\\n", re.String()) // print pattern
	fmt.Printf("\\nEmail: %v :%v\\n", str1, re.MatchString(str1))
	fmt.Printf("Email: %v :%v\\n", str2, re.MatchString(str2))
	fmt.Printf("Email: %v :%v\\n", str3, re.MatchString(str3))
	fmt.Printf("Email: %v :%v\\n", str4, re.MatchString(str4))
	fmt.Printf("Email: %v :%v\\n", str5, re.MatchString(str5))
}
```

Pattern: ^\[a\-zA\-Z0\-9.!#$%&'\*+/=?^\_\`{|}~\-\]+@\[a\-zA\-Z0\-9\](?:\[a\-zA\-Z0\-9\-\]{0,61}\[a\-zA\-Z0\-9\])?(?:\\.\[a\-zA\-Z0\-9\](?:\[a\-zA\-Z0\-9\-\]{0,61}\[a\-zA\-Z0\-9\])?)\*$

Email: ç$?§/az@gmail.com :false
Email: abcd@gmail\_yahoo.com :false
Email: abcd@gmail\-yahoo.com :true
Email: abcd@gmailyahoo :true
Email: abcd@gmail.yahoo :true

---

## Regular expression to validate phone number

// Regular expression to validate phone number
```golang
package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "1(234)5678901x1234"
	str2 := "(+351) 282 43 50 50"
	str3 := "90191919908"
	str4 := "555\-8909"
	str5 := "001 6867684"
	str6 := "001 6867684x1"
	str7 := "1 (234) 567\-8901"
	str8 := "1\-234\-567\-8901 ext1234"

	re := regexp.MustCompile(\`^(?:(?:\\(?(?:00|\\+)(\[1\-4\]\\d\\d|\[1\-9\]\\d?)\\)?)?\[\\\-\\.\\ \\\\\\/\]?)?((?:\\(?\\d{1,}\\)?\[\\\-\\.\\ \\\\\\/\]?){0,})(?:\[\\\-\\.\\ \\\\\\/\]?(?:#|ext\\.?|extension|x)\[\\\-\\.\\ \\\\\\/\]?(\\d+))?$\`)

	fmt.Printf("Pattern: %v\\n", re.String()) // print pattern
	fmt.Printf("\\nPhone: %v\\t:%v\\n", str1, re.MatchString(str1))
	fmt.Printf("Phone: %v\\t:%v\\n", str2, re.MatchString(str2))
	fmt.Printf("Phone: %v\\t\\t:%v\\n", str3, re.MatchString(str3))
	fmt.Printf("Phone: %v\\t\\t\\t:%v\\n", str4, re.MatchString(str4))
	fmt.Printf("Phone: %v\\t\\t:%v\\n", str5, re.MatchString(str5))
	fmt.Printf("Phone: %v\\t\\t:%v\\n", str6, re.MatchString(str6))
	fmt.Printf("Phone: %v\\t\\t:%v\\n", str7, re.MatchString(str7))
	fmt.Printf("Phone: %v\\t:%v\\n", str8, re.MatchString(str8))
}
```

Pattern: ^(?:(?:\\(?(?:00|\\+)(\[1\-4\]\\d\\d|\[1\-9\]\\d?)\\)?)?\[\\\-\\.\\ \\\\\\/\]?)?((?:\\(?\\d{1,}\\)?\[\\\-\\.\\ \\\\\\/\]?){0,})(?:\[\\\-\\.\\ \\\\\\/\]?(?:#|ext\\.?|extension|x
)\[\\\-\\.\\ \\\\\\/\]?(\\d+))?$

Phone: 1(234)5678901x1234       :true
Phone: (+351) 282 43 50 50      :true
Phone: 90191919908              :true
Phone: 555\-8909                 :true
Phone: 001 6867684              :true
Phone: 001 6867684x1            :true
Phone: 1 (234) 567\-8901         :true
Phone: 1\-234\-567\-8901 ext1234   :true

---

## Regular expression to validate the date format in "dd/mm/yyyy"

// Regular expression validate the date format in "dd/mm/yyyy"
```golang
package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "31/07/2010"
	str2 := "1/13/2010"
	str3 := "29/2/2007"
	str4 := "31/08/2010"
	str5 := "29/02/200a"
	str6 := "29/02/200a"
	str7 := "55/02/200a"
	str8 := "2\_/02/2009"

	re := regexp.MustCompile("(0?\[1\-9\]|\[12\]\[0\-9\]|3\[01\])/(0?\[1\-9\]|1\[012\])/((19|20)\\\\d\\\\d)")

	fmt.Printf("Pattern: %v\\n", re.String()) // print pattern
	fmt.Printf("\\nDate: %v :%v\\n", str1, re.MatchString(str1))
	fmt.Printf("Date: %v :%v\\n", str2, re.MatchString(str2))
	fmt.Printf("Date: %v :%v\\n", str3, re.MatchString(str3))
	fmt.Printf("Date: %v :%v\\n", str4, re.MatchString(str4))
	fmt.Printf("Date: %v :%v\\n", str5, re.MatchString(str5))
	fmt.Printf("Date: %v :%v\\n", str6, re.MatchString(str6))
	fmt.Printf("Date: %v :%v\\n", str7, re.MatchString(str7))
	fmt.Printf("Date: %v :%v\\n", str8, re.MatchString(str8))
}
```

Pattern: (0?\[1\-9\]|\[12\]\[0\-9\]|3\[01\])/(0?\[1\-9\]|1\[012\])/((19|20)\\d\\d)
Date: 31/07/2010 :true
Date: 1/13/2010 :false
Date: 29/2/2007 :true
Date: 31/08/2010 :true
Date: 29/02/200a :false
Date: 29/02/200a :false
Date: 55/02/200a :false
Date: 2\_/02/2009 :false

---

## Regular expression to validate common Credit Card Numbers

// Regular expression validate Visa, MasterCard, American Express, Diners Club, Discover, and JCB cards
```golang
package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "4111111111111111"
	str2 := "346823285239073"
	str3 := "370750517718351"
	str4 := "4556229836495866"
	str5 := "5019717010103742"
	str6 := "76009244561"
	str7 := "4111\-1111\-1111\-1111"
	str8 := "5610591081018250"
	str9 := "30569309025904"
	str10 := "6011111111111117"

	re := regexp.MustCompile(\`^(?:4\[0\-9\]{12}(?:\[0\-9\]{3})?|\[25\]\[1\-7\]\[0\-9\]{14}|6(?:011|5\[0\-9\]\[0\-9\])\[0\-9\]{12}|3\[47\]\[0\-9\]{13}|3(?:0\[0\-5\]|\[68\]\[0\-9\])\[0\-9\]{11}|(?:2131|1800|35\\d{3})\\d{11})$\`)

	fmt.Printf("Pattern: %v\\n", re.String()) // print pattern
	fmt.Printf("\\nCC : %v :%v\\n", str1, re.MatchString(str1))
	fmt.Printf("CC : %v :%v\\n", str2, re.MatchString(str2))
	fmt.Printf("CC : %v :%v\\n", str3, re.MatchString(str3))
	fmt.Printf("CC : %v :%v\\n", str4, re.MatchString(str4))
	fmt.Printf("CC : %v :%v\\n", str5, re.MatchString(str5))
	fmt.Printf("CC : %v :%v\\n", str6, re.MatchString(str6))
	fmt.Printf("CC : %v :%v\\n", str7, re.MatchString(str7))
	fmt.Printf("CC : %v :%v\\n", str8, re.MatchString(str8))
	fmt.Printf("CC : %v :%v\\n", str9, re.MatchString(str9))
	fmt.Printf("CC : %v :%v\\n", str10, re.MatchString(str10))
}
```

Pattern: ^(?:4\[0\-9\]{12}(?:\[0\-9\]{3})?|\[25\]\[1\-7\]\[0\-9\]{14}|6(?:011|5\[0\-9\]\[0\-9\])\[0\-9\]{12}|3\[47\]\[0\-9\]{13}|3(?:0\[0\-5\]|\[68\]\[0\-9\])\[0\-9\]{11}|(?:2131|18
00|35\\d{3})\\d{11})$
CC : 4111111111111111 :true
CC : 346823285239073 :true
CC : 370750517718351 :true
CC : 4556229836495866 :true
CC : 5019717010103742 :false
CC : 76009244561 :false
CC : 4111\-1111\-1111\-1111 :false
CC : 5610591081018250 :true
CC : 30569309025904 :true
CC : 6011111111111117 :true

---

## Replace any non\-alphanumeric character sequences with a dash using Regex

```golang
package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	reg, err := regexp.Compile("\[^A\-Za\-z0\-9\]+")
	if err != nil {
		log.Fatal(err)
	}
	newStr := reg.ReplaceAllString("#Golang#Python$Php&Kotlin@@", "\-")
	fmt.Println(newStr)
}
```

\-Golang\-Python\-Php\-Kotlin\-

---

## Replace first occurrence of string using Regexp
```golang
package main

import (
	"fmt"
	"regexp"
)

func main() {
	strEx := "Php\-Golang\-Php\-Python\-Php\-Kotlin"
	reStr := regexp.MustCompile("^(.\*?)Php(.\*)$")
	repStr := "${1}Java$2"
	output := reStr.ReplaceAllString(strEx, repStr)
	fmt.Println(output)
}
```

Java\-Golang\-Php\-Python\-Php\-Kotlin

---

## Regular expression to split string on white spaces
```golang
package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "Split   String on \\nwhite    \\tspaces."

	re := regexp.MustCompile(\`\\S+\`)

	fmt.Printf("Pattern: %v\\n", re.String()) // Print Pattern

	fmt.Printf("String contains any match: %v\\n", re.MatchString(str1)) // True

	submatchall := re.FindAllString(str1, \-1)
	for \_, element := range submatchall {
		fmt.Println(element)
	}
}
```

Pattern: \\S+
String contains any match: true
Split
String
on
white
spaces.

---

## Regular expression to extract numbers from a string in Golang
```golang
package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "Hello X42 I'm a Y\-32.35 string Z30"

	re := regexp.MustCompile(\`\[\-\]?\\d\[\\d,\]\*\[\\.\]?\[\\d{2}\]\*\`)

	fmt.Printf("Pattern: %v\\n", re.String()) // Print Pattern

	fmt.Printf("String contains any match: %v\\n", re.MatchString(str1)) // True

	submatchall := re.FindAllString(str1, \-1)
	for \_, element := range submatchall {
		fmt.Println(element)
	}
}
```

Pattern: \[\-\]?\\d\[\\d,\]\*\[\\.\]?\[\\d{2}\]\*
String contains any match: true
42
\-32.35
30

---

## Regular expression to extract filename from given path in Golang

```golang
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(\`^(.\*/)?(?:$|(.+?)(?:(\\.\[^.\]\*$)|$))\`)

	str1 := \`http://www.golangprograms.com/regular\-expressions.html\`
	match1 := re.FindStringSubmatch(str1)
	fmt.Println(match1\[2\])

	str2 := \`/home/me/dir3/dir3a/dir3ac/filepat.png\`
	match2 := re.FindStringSubmatch(str2)
	fmt.Println(match2\[2\])
}
```

regular\-expressions
filepat

---

## Split a string at uppercase letters using regular expression in Golang

```golang
package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "Hello X42 I'm a Y\-32.35 string Z30"

	re := regexp.MustCompile(\`\[A\-Z\]\[^A\-Z\]\*\`)

	fmt.Printf("Pattern: %v\\n", re.String()) // Print Pattern

	submatchall := re.FindAllString(str1, \-1)
	for \_, element := range submatchall {
		fmt.Println(element)
	}
}
```

Pattern: \[A\-Z\]\[^A\-Z\]\*
Hello
X42
I'm a
Y\-32.35 string
Z30

---

## Regular Expression to get a string between parentheses in Golang

```golang
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str1 := "This is a (sample) ((string)) with (SOME) special words"

	re := regexp.MustCompile(\`\\((.\*?)\\)\`)
	fmt.Printf("Pattern: %v\\n", re.String()) // print pattern

	fmt.Println("\\nText between parentheses:")
	submatchall := re.FindAllString(str1, \-1)
	for \_, element := range submatchall {
		element = strings.Trim(element, "(")
		element = strings.Trim(element, ")")
		fmt.Println(element)
	}
}
```

Pattern: \\((.\*?)\\)
Text between parentheses:
sample
string
SOME

---

## How to remove symbols from a string in Golang?

```golang
package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	str1 := "how much for the maple syrup? $20.99? That's ridiculous!!!"

	re, err := regexp.Compile(\`\[^\\w\]\`)
	if err != nil {
		log.Fatal(err)
	}
	str1 = re.ReplaceAllString(str1, " ")
	fmt.Println(str1)
}
```

How much for the apple cider   120 99  It is too much

---

## Regex to extract image name from HTML in Golang

```golang
package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := \`<img src="1.png"><x><z?>
			 <img czx zcxz src='2.png'><x><z?>\`

	re := regexp.MustCompile(\`<img\[^>\]+\\bsrc=\["'\](\[^"'\]+)\["'\]\`)

	submatchall := re.FindAllStringSubmatch(str1, \-1)
	for \_, element := range submatchall {
		fmt.Println(element\[1\])
	}
}
```

1.png
2.png

---

## How to replace emoji characters in string using regex in Golang?

```golang
package main

import (
	"fmt"
	"regexp"
)

func main() {
	var emojiRx = regexp.MustCompile(\`\[\\x{1F600}\-\\x{1F6FF}|\[\\x{2600}\-\\x{26FF}\]\`)
	var str = emojiRx.ReplaceAllString("Thats a nice joke ðŸ˜†ðŸ˜†ðŸ˜† ðŸ˜›", \`\[e\]\`)
	fmt.Println(str)
}
```

Thats a nice joke \[e\]\[e\]\[e\] \[e\]

---
```golang
## How to extract text from between html tag using Regular Expressions in Golang?

package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := \`<html><body>
			<form name="query" action="http://www.example.net/action.php" method="post">
				<textarea type="text" name="nameiknow">The text I want</textarea>
				<div id="button">
					<input type="submit" value="Submit" />
				</div>
			</form>
			</body></html>\`

	re := regexp.MustCompile(\`<textarea.\*?>(.\*)</textarea>\`)

	submatchall := re.FindAllStringSubmatch(str1, \-1)
	for \_, element := range submatchall {
		fmt.Println(element\[1\])
	}
}
```

The text I want

---

## Regular expression for matching HH:MM time format in Golang
```golang
package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "8:2"
	str2 := "9:9"
	str3 := "12:29"
	str4 := "02:5"
	str5 := "23:59"
	str6 := "55:59"
	str7 := "0:01"

	re := regexp.MustCompile(\`^(\[0\-9\]|0\[0\-9\]|1\[0\-9\]|2\[0\-3\]):(\[0\-9\]|\[0\-5\]\[0\-9\])$\`)

	fmt.Printf("Pattern: %v\\n", re.String()) // print pattern
	fmt.Printf("Time: %v\\t:%v\\n", str1, re.MatchString(str1))
	fmt.Printf("Time: %v\\t:%v\\n", str2, re.MatchString(str2))
	fmt.Printf("Time: %v\\t:%v\\n", str3, re.MatchString(str3))
	fmt.Printf("Time: %v\\t:%v\\n", str4, re.MatchString(str4))
	fmt.Printf("Time: %v\\t:%v\\n", str5, re.MatchString(str5))
	fmt.Printf("Time: %v\\t:%v\\n", str6, re.MatchString(str6))
	fmt.Printf("Time: %v\\t:%v\\n", str7, re.MatchString(str7))
}
```

Pattern: ^(\[0\-9\]|0\[0\-9\]|1\[0\-9\]|2\[0\-3\]):(\[0\-9\]|\[0\-5\]\[0\-9\])$
Time: 8:2       :true
Time: 9:9       :true
Time: 12:29     :true
Time: 02:5      :true
Time: 23:59     :true
Time: 55:59     :false
Time: 0:01      :true
