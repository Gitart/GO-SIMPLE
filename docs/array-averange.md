# Program in Go language to Calculate Average Using Arrays

This Go language program takes n number of element from user, stores data in an array and calculates the average of those numbers.

// Golang Program to Calculate Average Using Arrays.

package main
import "fmt"

func main(){
	var num\[100\] int
	var temp,sum,avg int
	fmt.Print("Enter number of elements: ")
	fmt.Scanln(&temp)
	for i := 0; i < temp; i++ {
		fmt.Print("Enter the number : ")
		fmt.Scanln(&num\[i\])
		sum += num\[i\]
	}

	avg = sum/temp
	fmt.Printf("The Average of entered %d number(s) is %d",temp,avg)
}

/\*

Output:

C:\\golang>go run example21.go
Enter number of elements: 5
Enter the number : 2
Enter the number : 65
Enter the number : 18
Enter the number : 59
Enter the number : 54
The Average of entered 5 number(s) is 39
\*/
