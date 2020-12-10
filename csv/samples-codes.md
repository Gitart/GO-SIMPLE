## Test

```go
func parseLocation(file string) (map[string]Point, error) {
    f, err := os.Open(file)
    defer f.Close()
    if err != nil {
        return nil, err
    }
    lines, err := csv.NewReader(f).ReadAll()
    if err != nil {
        return nil, err
    }
    locations := make(map[string]Point)
    for _, line := range lines {
        name := line[0]
        lat, laterr := strconv.ParseFloat(line[1], 64)
        if laterr != nil {
            return nil, laterr
        }
        lon, lonerr := strconv.ParseFloat(line[2], 64)
        if lonerr != nil {
            return nil, lonerr
        }
        locations[name] = Point{lat, lon}
    }
    return locations, nil
}

```

## This answer is not useful
[](https://stackoverflow.com/posts/24999351/timeline)


## Go is a very verbose language, however you could use something like this:

```go
// predeclare err
func parseLocation(file string) (locations map[string]*Point, err error) {
    f, err := os.Open(file)
    if err != nil {
        return nil, err
    }
    defer f.Close() // this needs to be after the err check

    lines, err := csv.NewReader(f).ReadAll()
    if err != nil {
        return nil, err
    }

    //already defined in declaration, no need for :=
    locations = make(map[string]*Point, len(lines))
    var lat, lon float64 //predeclare lat, lon
    for _, line := range lines {
        // shorter, cleaner and since we already have lat and err declared, we can do this.
        if lat, err = strconv.ParseFloat(line[1], 64); err != nil {
            return nil, err
        }
        if lon, err = strconv.ParseFloat(line[2], 64); err != nil {
            return nil, err
        }
        locations[line[0]] = &Point{lat, lon}
    }
    return locations, nil
}

```

## A more efficient and proper version

```go
func parseLocation(file string) (map[string]*Point, error) {
    f, err := os.Open(file)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    csvr := csv.NewReader(f)

    locations := map[string]*Point{}
    for {
        row, err := csvr.Read()
        if err != nil {
            if err == io.EOF {
                err = nil
            }
            return locations, err
        }

        p := &Point{}
        if p.lat, err = strconv.ParseFloat(row[1], 64); err != nil {
            return nil, err
        }
        if p.lon, err = strconv.ParseFloat(row[2], 64); err != nil {
            return nil, err
        }
        locations[row[0]] = p
    }
}

```

## This answer is not useful
Go now has a csv package for this. Its is `encoding/csv`. You can find the docs here: [https://golang.org/pkg/encoding/csv/](https://golang.org/pkg/encoding/csv/)
There are a couple of good examples in the docs. Here is a helper method I created to read a csv file and returns its records.

```go
package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
)

func readCsvFile(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}

func main() {
    records := readCsvFile("../tasks.csv")
    fmt.Println(records)
}

```



## Show activity on this post.

```go
import (
    "bufio"
    "encoding/csv"
    "os"
    "fmt"
    "io"
)

func ReadCsvFile(filePath string)  {
    // Load a csv file.
    f, _ := os.Open(filePath)

    // Create a new reader.
    r := csv.NewReader(f)
    for {
        record, err := r.Read()
        // Stop at EOF.
        if err == io.EOF {
            break
        }

        if err != nil {
            panic(err)
        }
        // Display record.
        // ... Display record length.
        // ... Display all individual elements of the slice.
        fmt.Println(record)
        fmt.Println(len(record))
        for value := range record {
            fmt.Printf("  %v\n", record[value])
        }
    }
}

```

## Show activity on this post.

```go
package main
import "encoding/csv"
import "io"

type Scanner struct {
   Reader *csv.Reader
   Head map[string]int
   Row []string
}

func NewScanner(o io.Reader) Scanner {
   csv_o := csv.NewReader(o)
   a, e := csv_o.Read()
   if e != nil {
      return Scanner{}
   }
   m := map[string]int{}
   for n, s := range a {
      m[s] = n
   }
   return Scanner{Reader: csv_o, Head: m}
}

func (o *Scanner) Scan() bool {
   a, e := o.Reader.Read()
   o.Row = a
   return e == nil
}

func (o Scanner) Text(s string) string {
   return o.Row[o.Head[s]]
}

```

### Example:

```go
package main
import "strings"

func main() {
   s := `Month,Day
January,Sunday
February,Monday`

   o := NewScanner(strings.NewReader(s))
   for o.Scan() {
      println(o.Text("Month"), o.Text("Day"))
   }
}

```
You can also read contents of a directory to load all the CSV files. And then read all those CSV files 1 by 1 with `goroutines`

### `csv` file:

```
101,300.00,11000901,1155686400
102,250.99,11000902,1432339200

```

### `main.go` file:

```go
const sourcePath string = "./source"

func main() {
    dir, _ := os.Open(sourcePath)
    files, _ := dir.Readdir(-1)

    for _, file := range files {
        fmt.Println("SINGLE FILE: ")
        fmt.Println(file.Name())
        filePath := sourcePath + "/" + file.Name()
        f, _ := os.Open(filePath)
        defer f.Close()
        // os.Remove(filePath)

        //func
        go func(file io.Reader) {
            records, _ := csv.NewReader(file).ReadAll()
            for _, row := range records {
                fmt.Println(row)
            }
        }(f)

        time.Sleep(10 * time.Millisecond)// give some time to GO routines for execute
    }
}

```

### And the OUTPUT will be:

> $ go run main.go

```
SINGLE FILE:
batch01.csv
[101 300.00 11000901 1155686400]
[102 250.99 11000902 1432339200]

```

###  Below example with the `Invoice struct`

```go
func main() {
    dir, _ := os.Open(sourcePath)
    files, _ := dir.Readdir(-1)

    for _, file := range files {
        fmt.Println("SINGLE FILE: ")
        fmt.Println(file.Name())
        filePath := sourcePath + "/" + file.Name()
        f, _ := os.Open(filePath)
        defer f.Close()

        go func(file io.Reader) {
            records, _ := csv.NewReader(file).ReadAll()
            for _, row := range records {
                invoice := new(Invoice)
                invoice.InvoiceNumber = row[0]
                invoice.Amount, _ = strconv.ParseFloat(row[1], 64)
                invoice.OrderID, _ = strconv.Atoi(row[2])
                unixTime, _ := strconv.ParseInt(row[3], 10, 64)
                invoice.Date = time.Unix(unixTime, 0)

                fmt.Printf("Received invoice `%v` for $ %.2f \n", invoice.InvoiceNumber, invoice.Amount)
            }
        }(f)

        time.Sleep(10 * time.Millisecond)
    }
}

type Invoice struct {
    InvoiceNumber string
    Amount        float64
    OrderID       int
    Date          time.Time
```
