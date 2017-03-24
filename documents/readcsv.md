## Чтение СSV файла
[Rossettacode CODE](http://rosettacode.org/wiki/CSV_data_manipulation#Go)


### Программа загружает СSV файл и суммирует эементы в строке и в конце дописывает сумму последним элементом


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

sample.csv:
```
C1,C2,C3,C4,C5
1,5,9,13,17
2,six,10,14,18
3,7,11,15,19
4,8,12,16,20
```

output.csv:
```
C1,C2,C3,C4,C5,SUM
1,5,9,13,17,45
2,six,10,14,18,NA
3,7,11,15,19,55
4,8,12,16,20,60
```



# Втрой пример

```golang
package main

import (
"os"
"fmt"
"encoding/csv"
)

func main() {

// Open CSV file, handle any errors
r := csv.NewReader(csvdata/csvtest)
r.Comma = '\t' // Use tab-separated values
row, e := r.Read()
if e != nil {
  panic(e)
}

// Get Longitude and Latitude Coordinates from CSV file
fmt.Println("Getting EPA Geographical Location Data...")
//Lat, e := Reader(r, "LatitudeMeasure")
//Long, e := Reader(r, "LongitudeMeasure")
if e != nil {
  panic(e)
}

}

func (r *Reader) ReadAll() (records [][]string, err os.Error) {
        for {
                record, err := r.Read()
                if err == os.EOF {
                        return records, nil
                }
                if err != nil {
                        return nil, err
                }
                records = append(records, record)
        }
        panic("unreachable")
}

type Reader struct {
    Comma            rune // Field delimiter (set to '\t' "tab" by NewReader)
    Comment          rune // Comment character for start of line
    FieldsPerRecord  int  // Number of expected fields per record
    LazyQuotes       bool // Allow lazy quotes
    TrailingComma    bool // Allow trailing comma
    TrimLeadingSpace bool // Trim leading space
    // contains filtered or unexported fields
}
```
