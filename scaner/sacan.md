# Scan

```golang
package main

import (
	"fmt"
	"github/ho/datafile"
	"log"
)


func main(){
	lines,err := datafile.GetStrings("votes.txt")
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(lines)

    var names []string
    var counts []int

    for _, line := range lines{
           matched:=false

           for i, name :=range names{
           	   if name ==line{
           	   	  counts[i]++
           	   	  matched = true
           	   }
           }

           if matched==false{
           	  names=append(names,line)
           	  consts =append(counts,1)
           }



          for i,name :=range names{
          	  fmt.Printf("%s: %d\n", name, counts[i])    
          } 

    }

}
```

# Lib

```golang
package datafile

import (
	"bufio"
	"os"
)

func GetStrings(filename string) ([]string,error){
var lines []string 
file,err:=os.Open(filename)
if err!=nil{
	return nil,err
}

scaner:=bufio.NewScaner(file)
for scaner.Scan(){
	line:=scaner.Text()
	lines=append(lines,line)
}

err = file.Close()
if err!=nil{
	return nil,err
}

if scaner.Err()!=nil{
	return nil, scaner.Err()
}
   return lines, nil

}
```




