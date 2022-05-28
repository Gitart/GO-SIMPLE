
ZetCode
All Go Python C# Java JavaScript Subscribe
Ebooks

    PyQt5 ebook
    Tkinter ebook
    SQLite Python
    wxPython ebook
    Windows API ebook
    Java Swing ebook
    Java games ebook
    MySQL Java ebook

Go regular expressions

last modified January 26, 2022

Go regular expressions tutorial shows how to parse text in Go using regular expressions.
Regular expressions

Regular expressions are used for text searching and more advanced text manipulation. Regular expressions are built into tools including grep and sed, text editors including vi and emacs, programming languages including Go, Java, and Python.

Go has built-in API for working with regular expressions; it is located in regexp package.

A regular expression defines a search pattern for strings. It is used to match text, replace text, or split text. A regular expression may be compiled for better performance. The Go syntax of the regular expressions accepted is the same general syntax used by Perl, Python, and other languages.
Regex examples

The following table shows a couple of regular expression strings.
Regex 	Meaning
. 	Matches any single character.
? 	Matches the preceding element once or not at all.
+ 	Matches the preceding element once or more times.
* 	Matches the preceding element zero or more times.
^ 	Matches the starting position within the string.
$ 	Matches the ending position within the string.
| 	Alternation operator.
[abc] 	Matches a or b, or c.
[a-c] 	Range; matches a or b, or c.
[^abc] 	Negation, matches everything except a, or b, or c.
\s 	Matches white space character.
\w 	Matches a word character; equivalent to [a-zA-Z_0-9]
Go regex MatchString

The MatchString function reports whether a string contains any match of the regular expression pattern.
matchstring.go

package main

import (
    "fmt"
    "log"
    "regexp"
)

func main() {

    words := [...]string{"Seven", "even", "Maven", "Amen", "eleven"}

    for _, word := range words {

        found, err := regexp.MatchString(".even", word)

        if err != nil {
            log.Fatal(err)
        }

        if found {

            fmt.Printf("%s matches\n", word)
        } else {

            fmt.Printf("%s does not match\n", word)
        }
    }
}

In the code example, we have five words in an array. We check which words match the .even regular expression.

words := [...]string{"Seven", "even", "Maven", "Amen", "eleven"}

We have an array of words.

for _, word := range words {

We go through the array of words.

found, err := regexp.MatchString(".even", word)

We check if the current word matches the regular expression with MatchString. We have the .even regular expression. The dot (.) metacharacter stands for any single character in the text.

if found {

    fmt.Printf("%s matches\n", word)
} else {

    fmt.Printf("%s does not match\n", word)
}

We print if of the word matches the regular expression or not.

$ go run matchstring.go 
Seven matches
even does not match
Maven does not match
Amen does not match
eleven matches

Two words in the array match the our regular expression.
Go compiled regular expression

The Compile function parses a regular expression and returns, if successful, a Regexp object that can be used to match against text. Compiled regular expressions yield faster code.

The MustCompile function is a convenience function which compiles a regular expression and panics if the expression cannot be parsed.
compiled.go

package main

import (
    "fmt"
    "log"
    "regexp"
)

func main() {

    words := [...]string{"Seven", "even", "Maven", "Amen", "eleven"}

    re, err := regexp.Compile(".even")

    if err != nil {
        log.Fatal(err)
    }

    for _, word := range words {

        found := re.MatchString(word)

        if found {

            fmt.Printf("%s matches\n", word)
        } else {

            fmt.Printf("%s does not match\n", word)
        }
    }
}

In the code example, we used a compiled regular expression.

re, err := regexp.Compile(".even")

We compile the regular expression with Compile.

found := re.MatchString(word)

The MatchString function is called on the returned regex object.
compiled2.go

package main

import (
    "fmt"
    "regexp"
)

func main() {

    words := [...]string{"Seven", "even", "Maven", "Amen", "eleven"}

    re := regexp.MustCompile(".even")

    for _, word := range words {

        found := re.MatchString(word)

        if found {

            fmt.Printf("%s matches\n", word)
        } else {

            fmt.Printf("%s does not match\n", word)
        }
    }
}

The example is simplified with MustCompile.
Go regex FindAllString

The FindAllString function returns a slice of all successive matches of the regular expression.
findall.go

package main

import (
    "fmt"
    "os"
    "regexp"
)

func main() {

    var content = `Foxes are omnivorous mammals belonging to several genera 
of the family Canidae. Foxes have a flattened skull, upright triangular ears, 
a pointed, slightly upturned snout, and a long bushy tail. Foxes live on every 
continent except Antarctica. By far the most common and widespread species of 
fox is the red fox.`

    re := regexp.MustCompile("(?i)fox(es)?")

    found := re.FindAllString(content, -1)

    fmt.Printf("%q\n", found)

    if found == nil {
        fmt.Printf("no match found\n")
        os.Exit(1)
    }

    for _, word := range found {
        fmt.Printf("%s\n", word)
    }

}

In the code example, we find all occurrences of the word fox, including its plural form.

re := regexp.MustCompile("(?i)fox(es)?")

With the (?i) syntax, the regular expression is case insensitive. The (es)? indicates that "es" characters might be included zero times or once.

found := re.FindAllString(content, -1)

We look for all occurrences of the defined regular expression with FindAllString. The second parameter is the maximum matches to look for; -1 means search for all possible matches.

$ go run findall.go 
["Foxes" "Foxes" "Foxes" "fox" "fox"]
Foxes
Foxes
Foxes
fox
fox

We have found five matches.
Go regex FindAllStringIndex

The FindAllStringIndex returns a slice of all successive indexes of matches of the expression.
allindex.go

package main

import (
    "fmt"
    "regexp"
)

func main() {

    var content = `Foxes are omnivorous mammals belonging to several genera 
of the family Canidae. Foxes have a flattened skull, upright triangular ears, 
a pointed, slightly upturned snout, and a long bushy tail. Foxes live on every 
continent except Antarctica. By far the most common and widespread species of 
fox is the red fox.`

    re := regexp.MustCompile("(?i)fox(es)?")

    idx := re.FindAllStringIndex(content, -1)

    for _, j := range idx {
        match := content[j[0]:j[1]]
        fmt.Printf("%s at %d:%d\n", match, j[0], j[1])
    }
}

In the code example, we find all occurrences of the fox word and their indexes in the text.

$ go run allindex.go 
Foxes at 0:5
Foxes at 81:86
Foxes at 196:201
fox at 296:299
fox at 311:314

Go regex Split

The Split function cuts a string into substrings separated by the defined regular expression. It returns a slice of the substrings between those expression matches.
splittext.go

package main

import (
    "fmt"
    "log"
    "regexp"
    "strconv"
)

func main() {

    var data = `22, 1, 3, 4, 5, 17, 4, 3, 21, 4, 5, 1, 48, 9, 42`

    sum := 0

    re := regexp.MustCompile(",\\s*")

    vals := re.Split(data, -1)

    for _, val := range vals {

        n, err := strconv.Atoi(val)

        sum += n

        if err != nil {
            log.Fatal(err)
        }
    }

    fmt.Println(sum)
}

In the code example, we have a comma-separated list of values. We cut the values from the string and calculate their sum.

re := regexp.MustCompile(",\\s*")

The regular expresion includes a comma character and any number of adjacend spaces.

vals := re.Split(data, -1)

We get the slice of values.

for _, val := range vals {

    n, err := strconv.Atoi(val)

    sum += n

    if err != nil {
        log.Fatal(err)
    }
}

We go through the slice and calculate the sum. The slice contains strings; therefore, we convert each string into an integer with strconv.Atoi function.

$ go run splittext.go 
189

The sum of values is 189.
Go regex capturing groups

Round brackets () are used to create capturing groups. This allows us to apply a quantifier to the entire group or to restrict alternation to a part of the regular expression.

To find capturing groups (Go uses the term subexpressions), we use the FindStringSubmatch function.
capturegroups.go

package main

import (
    "fmt"
    "regexp"
)

func main() {

    websites := [...]string{"webcode.me", "zetcode.com", "freebsd.org", "netbsd.org"}

    re := regexp.MustCompile("(\\w+)\\.(\\w+)")

    for _, website := range websites {

        parts := re.FindStringSubmatch(website)

        for i, _ := range parts {
            fmt.Println(parts[i])
        }

        fmt.Println("---------------------")
    }
}

In the code example, we divide the domain names into two parts by using groups.

re := regexp.MustCompile("(\\w+)\\.(\\w+)")

We define two groups with parentheses.

parts := re.FindStringSubmatch(website)

The FindStringSubmatch returns a slice of strings holding the matches, including those from the capturing groups.

$ go run capturegroups.go 
webcode.me
webcode
me
---------------------
zetcode.com
zetcode
com
---------------------
freebsd.org
freebsd
org
---------------------
netbsd.org
netbsd
org
---------------------

Go regex replacing strings

It is possible to replace strings with ReplaceAllString. The method returns the modified string.
replacing.go

package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "regexp"
    "strings"
)

func main() {

    resp, err := http.Get("http://webcode.me")

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {

        log.Fatal(err)
    }

    content := string(body)

    re := regexp.MustCompile("<[^>]*>")
    replaced := re.ReplaceAllString(content, "")

    fmt.Println(strings.TrimSpace(replaced))
}

The example reads HTML data of a web page and strips its HTML tags using a regular expression.

resp, err := http.Get("http://webcode.me")

We create a GET request with Get function from the http package.

body, err := ioutil.ReadAll(resp.Body)

We read the body of the response object.

re := regexp.MustCompile("<[^>]*>")

This pattern defines a regular expression that matches HTML tags.

replaced := re.ReplaceAllString(content, "")

We remove all the tags with ReplaceAllString method.
Go regex ReplaceAllStringFunc

The ReplaceAllStringFunc returns a copy of a string in which all matches of the regular expression have been replaced by the return value of the specified function.
replacing2.go

package main

import (
    "fmt"
    "regexp"
    "strings"
)

func main() {

    content := "an old eagle"

    re := regexp.MustCompile(`[^aeiou]`)

    fmt.Println(re.ReplaceAllStringFunc(content, strings.ToUpper))
}

In the code example, we apply the strings.ToUpper function on all wovels of a string.

$ go run replaceallfunc.go 
aN oLD eaGLe

In this tutorial, we have worked with regular expression in Go.

https://zetcode.com/golang/regex/
