
## Go word frequency

last modified January 26, 2022

Go word frequency tutorial shows how to calculate word frequency in Golang.

Our examples will work only with latin words and are specifically targeted at analyzing the Bible.

$ wget https://raw.githubusercontent.com/janbodnar/data/main/the-king-james-bible.txt

We use the King James Bible.

To cut the text into words, we use Go's strings.FieldsFunc and regular expressions.
Go word frequency example I

The FieldsFunc function splits the string at each run of Unicode code points satisfying the provided function and returns an array of slices.
read_freq.go

package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "sort"
    "strings"
)

func main() {

    fileName := "the-king-james-bible.txt"

    bs, err := ioutil.ReadFile(fileName)

    if err != nil {

        log.Fatal(err)
    }

    text := string(bs)

    fields := strings.FieldsFunc(text, func(r rune) bool {

        return !('a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' || r == '\'')
    })

    wordsCount := make(map[string]int)

    for _, field := range fields {

        wordsCount[field]++
    }

    keys := make([]string, 0, len(wordsCount))

    for key := range wordsCount {

        keys = append(keys, key)
    }

    sort.Slice(keys, func(i, j int) bool {

        return wordsCount[keys[i]] > wordsCount[keys[j]]
    })

    for idx, key := range keys {

        fmt.Printf("%s %d\n", key, wordsCount[key])

        if idx == 10 {
            break
        }
    }
}

We count the frequency of the words from the King James Bible.

fields := strings.FieldsFunc(text, func(r rune) bool {

    return !('a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' || r == '\'')
})

The FieldsFunc cuts the text by characters that are not alphabetic and apostrophe. This will also disregard all the verse numbers.

wordsCount := make(map[string]int)

for _, field := range fields {

    wordsCount[field]++
}

Each word and its frequency is stored in the wordsCount map.

keys := make([]string, 0, len(wordsCount))

for key := range wordsCount {

    keys = append(keys, key)
}

sort.Slice(keys, func(i, j int) bool {

    return wordsCount[keys[i]] > wordsCount[keys[j]]
})

In order to sort the words by frequency, we create a new keys slice. We put all the words there and sort them by their frequency values.

for idx, key := range keys {

    fmt.Printf("%s %d\n", key, wordsCount[key])

    if idx == 10 {
        break
    }
}

We print the top ten frequent words from the Bible.

$ go run word_freq.go
the 62103
and 38848
of 34478
to 13400
And 12846
that 12576
in 12331
shall 9760
he 9665
unto 8942
I 8854

Go word frequency example II

In the second example, we use a regular expression to divide text into words.
word_freq2.go

package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "regexp"
    "sort"
)

type WordFreq struct {
    word string
    freq int
}

func (p WordFreq) String() string {
    return fmt.Sprintf("%s %d", p.word, p.freq)
}

func main() {

    fileName := "the-king-james-bible.txt"

    reg := regexp.MustCompile("[a-zA-Z']+")
    bs, err := ioutil.ReadFile(fileName)

    if err != nil {
        log.Fatal(err)
    }

    text := string(bs)
    matches := reg.FindAllString(text, -1)

    words := make(map[string]int)

    for _, match := range matches {
        words[match]++
    }

    var wordFreqs []WordFreq
    for k, v := range words {
        wordFreqs = append(wordFreqs, WordFreq{k, v})
    }

    sort.Slice(wordFreqs, func(i, j int) bool {

        return wordFreqs[i].freq > wordFreqs[j].freq
    })

    for i := 0; i < 10; i++ {

        fmt.Println(wordFreqs[i])
    }
}

We store the words and their frequencies in the WordFreq structure.

reg := regexp.MustCompile("[a-zA-Z']+")

In our regular expression, one or more alphabetic characters or an apostrophe constitues a word.

matches := reg.FindAllString(text, -1)

The FindAllString function returns a slice of all successive matches of the expression.

words := make(map[string]int)

for _, match := range matches {
    words[match]++
}

We go over the matches and calculate their frequencies in the file. The words and the number of their occurrences are store in the words map.

var wordFreqs []WordFreq
for k, v := range words {
    wordFreqs = append(wordFreqs, WordFreq{k, v})
}

We build a slice of WordFreq structures out of the words map.

sort.Slice(wordFreqs, func(i, j int) bool {

    return wordFreqs[i].freq > wordFreqs[j].freq
})

We sort the wordFreqs slice by the freq field.

for i := 0; i < 10; i++ {

    fmt.Println(wordFreqs[i])
}

We print the first ten most common words.
Go word frequency example III

In the next example, we also use a regular expression.
word_freq3.go

package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "regexp"
    "sort"
)

type WordFreq struct {
    word string
    freq int
}

func (p WordFreq) String() string {
    return fmt.Sprintf("%s %d", p.word, p.freq)
}

type byFreq []WordFreq

func (a byFreq) Len() int           { return len(a) }
func (a byFreq) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byFreq) Less(i, j int) bool { return a[i].freq < a[j].freq }

func main() {

    fileName := "the-king-james-bible.txt"
    bs, err := ioutil.ReadFile(fileName)

    if err != nil {
        log.Fatal(err)
    }

    text := string(bs)

    re := regexp.MustCompile("[a-zA-Z']+")
    matches := re.FindAllString(text, -1)

    words := make(map[string]int)

    for _, match := range matches {
        words[match]++
    }

    var wordFreqs []WordFreq
    for k, v := range words {
        wordFreqs = append(wordFreqs, WordFreq{k, v})
    }

    sort.Sort(sort.Reverse(byFreq(wordFreqs)))

    for i := 0; i < 10; i++ {
        fmt.Printf("%v\n", wordFreqs[i])
    }
}

This example also implements a custom sorting interface.

type byFreq []WordFreq

func (a byFreq) Len() int           { return len(a) }
func (a byFreq) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byFreq) Less(i, j int) bool { return a[i].freq < a[j].freq }

We implement the sort.Interface for []WordFreq based on the freq field.

sort.Sort(sort.Reverse(byFreq(wordFreqs)))

To sort the WordFreq structures in descending order, we use the sort.Reverse function.

In this tutorial, we have counted word frequencies in the King James Bible.

