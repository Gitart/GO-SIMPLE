## Ввод с консоли


```golang
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)


func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\r\n", "", -1)

		// if strings.Compare("hi", text) == 0 {
		// 	fmt.Println("hello, Yourself")
		// }


		switch text{
		       case "Ok": fmt.Println("Да ок")
		       case "Or": fmt.Println("Да OR")
		   }


}




func Inputs() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\r\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}

	}

	

}



func Input_d(){
	fmt.Print("\nContinue? [Y/N] ")

	reader := bufio.NewReader(os.Stdin)
	c, num, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("You entered: %q\n", c)
	fmt.Println("The size entered: ", num)
    
      

	if c == 'L' || c == 'Y' {
		fmt.Println("Thank you for pressing Y to continue!")
	} else {
		fmt.Println("No? Ok, we'll exit.")
	}



}



func Input_v() {
    var text string

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    text, _ = reader.ReadString('\n')
    
   tt:= strings.Replace(text, "\n", "", -1) 
   fmt.Println(tt)

   if tt=="Ok"{
   fmt.Println("hhhhhhhhhhhhhhh")   	
   }


      switch tt {
      case "Ok": 
        fmt.Println("Ok2222")
        break
      case "No": 
        fmt.Println("No33333")
        break
}

}
```
