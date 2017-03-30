# Generated Code  or password


```golang
package main

 import (
         "fmt"
         "math/rand"
         "time"
 )

 //var vowels = []rune{'a', 'e', 'i', 'o', 'u'}
 //var consonants = []rune{'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n','p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z'}

 func humanReadablePassword(alphabetSize, numberSize int) string {

         vowels := "aeiou"
         consonants := "bcdfghjklmnpqrstvwxyz"
         digits := "0123456789"

         prefixSize := alphabetSize / 2
         if alphabetSize%2 != 0 {
                 prefixSize = int(alphabetSize/2) + 1
         }
         suffixSize := alphabetSize - prefixSize

         var prefixPart = make([]byte, prefixSize)

         for i := 0; i <= prefixSize-1; i++ {
                 if i%2 == 0 {
                         // use consonants
                         prefixPart[i] = consonants[rand.Intn(len(consonants)-1)]
                 } else {
                         // use vowels
                         prefixPart[i] = vowels[rand.Intn(len(vowels)-1)]
                 }
         }

         var midPart = make([]byte, numberSize)

         // use digits
         for k, _ := range midPart {
                 midPart[k] = digits[rand.Intn(len(digits))]
         }

         var suffixPart = make([]byte, suffixSize)

         for i := 0; i <= suffixSize-1; i++ {
                 if i%2 == 0 {
                         // use consonants
                         suffixPart[i] = consonants[rand.Intn(len(consonants)-1)]
                 } else {
                         // use vowels
                         suffixPart[i] = vowels[rand.Intn(len(vowels)-1)]
                 }
         }

         return string(prefixPart) + string(midPart) + string(suffixPart)
 }

 func main() {
         rand.Seed(time.Now().UnixNano())
         fmt.Println("6 alphabets with 2 digits : ", humanReadablePassword(6, 2)) // best option
         fmt.Println("3 alphabets with 8 digits : ", humanReadablePassword(3, 8))
         fmt.Println("9 alphabets with 9 digits : ", humanReadablePassword(9, 9))

 }
 ```
 
 ## Rendom string
 
 ```golang
 package main

 import (
    "encoding/base64"
    "crypto/rand"
    "fmt"
 )

 func main() {
   size := 32 // change the length of the generated random string here

   rb := make([]byte,size)
   _, err := rand.Read(rb)


   if err != nil {
      fmt.Println(err)
   }

   rs := base64.URLEncoding.EncodeToString(rb)

   fmt.Println(rs)
 }
 ```
 
 Output :

```
>go run randomstr.go
k38b1e3c4YVQJ4FMrgcRBLU2DyEuDlyxVNjy3UA7sIw=
>go run randomstr.go
b6NW5LJY1gI5Hui7oWZbTJa_LCsL__JaJ33zZ5NIXHE=
```


# Variant 2


```golang
package main

 import (
         "crypto/rand"
         "fmt"
 )

 func randStr(strSize int, randType string) string {

         var dictionary string

         if randType == "alphanum" {
                 dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
         }

         if randType == "alpha" {
                 dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
         }

         if randType == "number" {
                 dictionary = "0123456789"
         }

         var bytes = make([]byte, strSize)
         rand.Read(bytes)
         for k, v := range bytes {
                 bytes[k] = dictionary[v%byte(len(dictionary))]
         }
         return string(bytes)
 }

 func main() {

         fmt.Println("Alphanum : ", randStr(16, "alphanum"))

         fmt.Println("Alpha : ", randStr(16, "alpha"))

         fmt.Println("Numbers : ", randStr(16, "number"))

 }
 ```
 
Output (sample) :

```
Alphanum : jem7veyr7VVi2gNT
Alpha : UYqbhKPACfHsvjLh
Numbers : 8892932200439488
``` 
