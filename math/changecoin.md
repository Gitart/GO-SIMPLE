# Cahge coin

```golang
package main

 import (
         "fmt"
         "os"
 )

 var (
         amount, originalAmount, quarters, dimes, nickels, pennies int
 )

 func sanityCheck(input, min, max int) {

         if !((input >= min) && (input <= max)) {
                 fmt.Printf("Input parameter must in between %d and %d.\n", min, max)
                 os.Exit(-1)
         }

 }
```

```golang
 func main() {
         fmt.Println("Enter an integer value from 1 to 100 : ")
         fmt.Println("and I will show you a combination of coins")
         fmt.Println("that equals to the number.")

         _, err := fmt.Scanf("%d", &amount)

         if err != nil {
                 fmt.Println(err)
         }

         sanityCheck(amount, 1, 100)

         fmt.Println("You have entered : ", amount)

         originalAmount = amount
         quarters = amount / 25
         amount = amount % 25
         dimes = amount / 10
         amount = amount % 10
         nickels = amount / 5
         amount = amount % 5
         pennies = amount

         fmt.Println("Original amount entered : ", originalAmount)
         fmt.Println("Has a combination in coins of : ")
         fmt.Println("Quarters : ", quarters)
         fmt.Println("Dimes : ", dimes)
         fmt.Println("Nickels : ", nickels)
         fmt.Println("Pennies : ", pennies)

 }
 ```
