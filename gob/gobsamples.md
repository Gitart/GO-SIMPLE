### Load

```golang
package main

 import (
 	"encoding/gob"
 	"fmt"
 	"os"
 )

 func main() {
 	data := []int{101, 102, 103}

 	// create a file
 	dataFile, err := os.Create("integerdata.gob")

 	if err != nil {
 		fmt.Println(err)
 		os.Exit(1)
 	}

      // serialize the data
 	dataEncoder := gob.NewEncoder(dataFile)
 	dataEncoder.Encode(data)

 	dataFile.Close()
 }
```

### and retrieve the serialized objects or values:
readfilegob.go

```golang
 package main

 import (
 	"encoding/gob"
 	"fmt"
 	"os"
 )

 func main() {
 	var data []int

 	// open data file
 	dataFile, err := os.Open("integerdata.gob")

 	if err != nil {
 		fmt.Println(err)
 		os.Exit(1)
 	}

 	dataDecoder := gob.NewDecoder(dataFile)
 	err = dataDecoder.Decode(&data)

 	if err != nil {
 		fmt.Println(err)
 		os.Exit(1)
 	}

 	dataFile.Close()

 	fmt.Println(data)
 }
 ```
 
