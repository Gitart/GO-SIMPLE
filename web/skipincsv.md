## Skip blank/empty lines in CSV file and trim whitespaces example

### Problem:
Your Golang program is reading from a CSV file and you want to skip blank lines in the CSV file. 
Also, you want to trim whitespaces in the final output. How to do that?  

### Solution:
After initiating the csv.Read() or csv.NewReader() functions. Set the reader's FieldsPerRecord and TrimLeadingSpace to:

 csvReader.FieldsPerRecord = -1 // optional
 csvReader.TrimLeadingSpace = true
 
### The CSV reader will automagically ignore those empty lines.

```golang
 package main

 import (
 	"encoding/csv"
 	"fmt"
 	"io"
 	"log"
 	"strings"
 )

 func main() {
 	csvDataWithEmptyLines := `first_name,last_name,username
 
 "Adam","Sandler",     adam
 Steve,     McQueen,steve
 "Robert","Spacey","robspacy"
 `
 	csvReader := csv.NewReader(strings.NewReader(csvDataWithEmptyLines))

 	// add these
 	csvReader.FieldsPerRecord = -1
 	csvReader.TrimLeadingSpace = true

 	for {
 		record, err := csvReader.Read()
 		if err == io.EOF {
 			break
 		}
 		if err != nil {
 			log.Fatal(err)
 		}

 		fmt.Println(record)
 	}
 }
 ```
 
Output:
// properly formatted and look nice
```
[firstname lastname username]
[Adam Sandler adam]
[Steve McQueen steve]
[Robert Spacey robspacy]
```
