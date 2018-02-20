# Sample import csv
http://rosettacode.org/wiki/CSV_data_manipulation#Go


```golang
package main
 
import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)
 
func main() {
	rows := readSample()
	appendSum(rows)
	writeChanges(rows)
}
 
func readSample() [][]string {
	f, err := os.Open("sample.csv")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := csv.NewReader(f).ReadAll()
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	return rows
}
 
func appendSum(rows [][]string) {
	rows[0] = append(rows[0], "SUM")
	for i := 1; i < len(rows); i++ {
		rows[i] = append(rows[i], sum(rows[i]))
	}
}
 
func sum(row []string) string {
	sum := 0
	for _, s := range row {
		x, err := strconv.Atoi(s)
		if err != nil {
			return "NA"
		}
		sum += x
	}
	return strconv.Itoa(sum)
}
 
func writeChanges(rows [][]string) {
	f, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}
	err = csv.NewWriter(f).WriteAll(rows)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
```


## sample.csv:

C1,C2,C3,C4,C5  
1,5,9,13,17   
2,six,10,14,18   
3,7,11,15,19   
4,8,12,16,20  

## output.csv:

C1,C2,C3,C4,C5,SUM  
1,5,9,13,17,45  
2,six,10,14,18,NA  
3,7,11,15,19,55   
4,8,12,16,20,60  


