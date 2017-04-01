# Работа с операционной системой

```golang

package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func execute() {

  // here we perform the pwd command.
  // we can store the output of this in our out variable 
  // and catch any errors in err
	out, err := exec.Command("pwd").Output()

  // if there is an error with our execution
  // handle it here
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Println("Command Successfully Executed")
  // as the out variable defined above is of type []byte we need to convert
  // this to a string or else we will see garbage printed out in our console
  // this is how we convert it to a string
	output := string(out[:])

  // once we have converted it to a string we can then output it.
	fmt.Println(output)
}

func main() {

	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		execute()
	}
}

```
