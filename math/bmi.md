## Calculate BMI and risk category


I'm building a healthcare IoT device that will ask a person about their weight and report back the risk category.    
Eventually, the plan is to stop asking the question and take the weight measurement from a pressure sensor and display   
the risk category on a display. For now, let's get the basic right first.

### To calculate BMI(Body Mass Index), the formula is

```
BMI = mass / height * height
where mass is in kilograms and height is in meters
and the health risk categories are :

Underweight < 18.5
Normal >= 18.5 and < 25
Overweight >= 25 and < 30
Obese >= 30
```


```golang
 package main

 import (
         "fmt"
         "math"
 )

 var (
         weight, height, bmi float64
 )

 func main() {
         fmt.Println("Enter your weight in kilograms : ")
         fmt.Scanf("%f", &weight)

         fmt.Println("Enter your height in meters : ")
         fmt.Scanf("%f", &height)

         fmt.Println("You have entered weight of : ", weight)
         fmt.Println("You have entered height of : ", height)

         bmi = weight / math.Pow(height, 2)

         fmt.Println("Your BMI is : ", bmi)
         fmt.Print("Your risk category is : ")

         if bmi < float64(18.5) {
                 fmt.Println("Underweight")
         } else if bmi < 25 {
                 fmt.Println("Normal weight")
         } else if bmi < 30 {
                 fmt.Println("Overweight")
         } else {
                 fmt.Println("Obese")
         }

         // calculate normal weight based on height and bmi = 25
         normalWeight := 25 * math.Pow(height, 2)
         delta := weight - normalWeight

         fmt.Printf("The normal weight for your height is : %0.2v kilograms.\n", normalWeight)

         if (delta > 0) && (bmi > 30) {
                 fmt.Printf("You need to reduce %0.2v kilograms.\n", math.Abs(delta))
         }

         if (delta < 0) && (bmi < float64(18.5)) {
                 fmt.Printf("You need to increase %0.2v kilograms.\n", math.Abs(delta))
         }

 }
 ```
 
Sample output:
```
Enter your weight in kilograms :
100
Enter your height in meters :
1.7
You have entered weight of : 100
You have entered height of : 1.7
Your BMI is : 34.602076124567475
Your risk category is : Obese
The normal weight for your height is : 72
You ne
