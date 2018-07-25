
// hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)


package main

import (
    "golang.org/x/crypto/bcrypt"
    "fmt"
)

func main() {
    password := []byte("MyDarkSecret")

    // Hashing the password with the default cost of 10
    hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(hashedPassword))

    // Comparing the password with the hash
    err = bcrypt.CompareHashAndPassword(hashedPassword, password)
    fmt.Println(err) // nil means it is a match
}
