### Golang program that uses csv, NewReader on file

```go


package main

import (
    "bufio"
    "encoding/csv"
    "os"
    "fmt"
    "io"
)

func main() {
    // Load a TXT file.
    f, _ := os.Open("C:\\programs\\file.txt")

    // Create a new reader.
    r := csv.NewReader(bufio.NewReader(f))
    for {
        record, err := r.Read()
        // Stop at EOF.
        if err == io.EOF {
            break
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



### Contents: file.txt
```
cat,dog,bird
10,20,30,40
fish,dog,snake
```

### Output

```
[cat dog bird]
3
  cat
  dog
  bird
[10 20 30 40]
4
  10
  20
  30
  40
[fish dog snake]
3
  fish
  dog
  snake
  ```
  
## Golang program that uses ReadAll, strings.NewReader

```go
package main

import (
    "encoding/csv"
    "fmt"
    "strings"
)

func main() {
    // Create a 3-line string.
    data := `fish,blue,water
fox,red,farm
sheep,white,mountain
frog,green,pond`

    // Use strings.NewReader.
    // ... This creates a new Reader for passing to csv.NewReader.
    r := csv.NewReader(strings.NewReader(data))
    // Read all records.
    result, _ := r.ReadAll()

    fmt.Printf("Lines: %v", len(result))
    fmt.Println()

    for i := range result {
        // Element count.
        fmt.Printf("Elements: %v", len(result[i]))
        fmt.Println()
        // Elements.
        fmt.Println(result[i])
    }
}
```

## Output

```txt


Lines: 4
Elements: 3
[fish blue water]
Elements: 3
[fox red farm]
Elements: 3
[sheep white mountain]
Elements: 3
[frog green pond]
```

