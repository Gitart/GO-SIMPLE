 package main

 import (
         "fmt"
         "golang.org/x/text/transform"
         "golang.org/x/text/unicode/norm"
         "regexp"
         "strings"
         "unicode"
 )

 func isMn(r rune) bool {
         return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
 }

 func SEOURL(s string) string {
         seoStr := strings.ToLower(s)
         //seoStr = strings.Replace(seoStr, "/", "-", -1)

         //regE := regexp.MustCompile("/s+/")
         //seoStrByte := regE.ReplaceAll([]byte(seoStr), []byte("-"))
         //seoStr = string(seoStrByte) // convert []byte to string

         // convert all spaces to dash
         regE := regexp.MustCompile("[[:space:]]")
         seoStrByte := regE.ReplaceAll([]byte(seoStr), []byte("-"))
         seoStr = string(seoStrByte) // convert []byte to string

         // remove all blanks such as tab
         regE = regexp.MustCompile("[[:blank:]]")
         seoStrByte = regE.ReplaceAll([]byte(seoStr), []byte(""))
         seoStr = string(seoStrByte) // convert []byte to string

         // remove all punctuations with the exception of dash
         //regE = regexp.MustCompile("[[:punct:]]")

         regE = regexp.MustCompile("[!/:-@[-`{-~]")
         seoStrByte = regE.ReplaceAll([]byte(seoStr), []byte(""))
         seoStr = string(seoStrByte) // convert []byte to string

         // \x9\xA\xD will cause non-hex character in escape sequence error
         // regE = regexp.MustCompile("/[^\x9\xA\xD\x20-\x7F]/")

         //regE = regexp.MustCompile("[[:xdigit:]]") -- will remove some alphabet. Bug?

         regE = regexp.MustCompile("/[^\x20-\x7F]/")
         seoStrByte = regE.ReplaceAll([]byte(seoStr), []byte(""))
         seoStr = string(seoStrByte) // convert []byte to string

         regE = regexp.MustCompile("`&(amp;)?#?[a-z0-9]+;`i")
         seoStrByte = regE.ReplaceAll([]byte(seoStr), []byte("-"))
         seoStr = string(seoStrByte) // convert []byte to string

         regE = regexp.MustCompile("`&([a-z])(acute|uml|circ|grave|ring|cedil|slash|tilde|caron|lig|quot|rsquo);`i")
         seoStrByte = regE.ReplaceAll([]byte(seoStr), []byte("\\1"))
         seoStr = string(seoStrByte) // convert []byte to string

         regE = regexp.MustCompile("`[^a-z0-9]`i")
         seoStrByte = regE.ReplaceAll([]byte(seoStr), []byte("-"))
         seoStr = string(seoStrByte) // convert []byte to string

         regE = regexp.MustCompile("`[-]+`")
         seoStrByte = regE.ReplaceAll([]byte(seoStr), []byte("-"))
         seoStr = string(seoStrByte) // convert []byte to string

         // normalize unicode strings and remove all diacritical/accents marks
         // see https://www.socketloop.com/tutorials/golang-normalize-unicode-strings-for-comparison-purpose

         t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
         seoStr, _, _ = transform.String(t, seoStr)

         return strings.TrimSpace(seoStr)
 }

 func main() {

         NonSEOString := "@<ElNi\u00f1o coming?  > #% sooner this year!"

         fmt.Println("BEFORE : ", NonSEOString)

         SEOedString := SEOURL(NonSEOString)

         fmt.Println("AFTER : ", SEOedString)
 }
