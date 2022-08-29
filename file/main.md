# Reading and Writing Different File Types

[❮ Previous](https://www.golangprograms.com/files-directories-examples.html) [Next ❯](https://www.golangprograms.com/regular-expressions.html)

---

## Reading and Writing Different File Types

Learn how to read and write data in common file types(Text, CSV, JSON, and XML) using bufio, encoding and io packages.

---

## Reading XML file

The xml package includes `Unmarshal()` function that supports decoding data from a byte slice into values. The `xml.Unmarshal()` function is used to decode the values from the XML formatted file into a `Notes` struct.
Sample XML file:

The notes.xml file is read with the ioutil.ReadFile() function and a byte slice is returned, which is then decoded into a struct instance with the xml.Unmarshal() function. The struct instance member values are used to print the decoded data.

### Example

<note\>
<to\>Tove</to\>
<from\>Jani</from\>
<heading\>Reminder</heading\>
<body\>Don't forget me this weekend!</body\>
</note\>

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Notes struct {
	To      string \`xml:"to"\`
	From    string \`xml:"from"\`
	Heading string \`xml:"heading"\`
	Body    string \`xml:"body"\`
}

func main() {
	data, \_ := ioutil.ReadFile("notes.xml")

	note := &Notes{}

	\_ = xml.Unmarshal(\[\]byte(data), &note)

	fmt.Println(note.To)
	fmt.Println(note.From)
	fmt.Println(note.Heading)
	fmt.Println(note.Body)
}

### Output

```jsx
Tove
Jani
Reminder
Don't forget me this weekend!
```

---

## Writing XML file

The xml package has an `Marshal()` function which is used to serialized values from a struct and write them to a file in XML format.

The notes struct is defined with an uppercase first letter and ″xml″ field tags are used to identify the keys. The struct values are initialized and then serialize with the xml.Marshal() function. The serialized XML formatted byte slice is received which then written to a file using the ioutil.WriteFile() function.

### Example

```jsx
package main

import (
	"encoding/xml"
	"io/ioutil"
)

type notes struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Heading string `xml:"heading"`
	Body    string `xml:"body"`
}

func main() {
	note := ¬es{To: "Nicky",
		From:    "Rock",
		Heading: "Meeting",
		Body:    "Meeting at 5pm!",
	}

	file, _ := xml.MarshalIndent(note, "", " ")

	_ = ioutil.WriteFile("notes1.xml", file, 0644)

}
```

---

## Reading JSON file

The json package includes `Unmarshal()` function which supports decoding data from a byte slice into values. The decoded values are generally assigned to struct fields, the field names must be exported and should be in capitalize format.

The JSON file test.json is read with the ioutil.ReadFile() function, which returns a byte slice that is decoded into the struct instance using the json.Unmarshal() function. At last, the struct instance member values are printed using for loop to demonstrate that the JSON file was decoded.

### Example

```jsx
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type CatlogNodes struct {
	CatlogNodes []Catlog `json:"catlog_nodes"`
}

type Catlog struct {
	Product_id string `json: "product_id"`
	Quantity   int    `json: "quantity"`
}

func main() {
	file, _ := ioutil.ReadFile("test.json")

	data := CatlogNodes{}

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.CatlogNodes); i++ {
		fmt.Println("Product Id: ", data.CatlogNodes[i].Product_id)
		fmt.Println("Quantity: ", data.CatlogNodes[i].Quantity)
	}

}
```

---

## Writing JSON file

The json package has a `MarshalIndent()` function which is used to serialized values from a struct and write them to a file in JSON format.

The Salary struct is defined with json fields. The struct values are initialized and then serialize with the json.MarshalIndent() function. The serialized JSON formatted byte slice is received which then written to a file using the ioutil.WriteFile() function.

### Example

```jsx
package main

import (
	"encoding/json"
	"io/ioutil"
)

type Salary struct {
        Basic, HRA, TA float64
    }

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func main() {
	data := Employee{
        FirstName: "Mark",
        LastName:  "Jones",
        Email:     "mark@gmail.com",
        Age:       25,
        MonthlySalary: []Salary{
            Salary{
                Basic: 15000.00,
                HRA:   5000.00,
                TA:    2000.00,
            },
            Salary{
                Basic: 16000.00,
                HRA:   5000.00,
                TA:    2100.00,
            },
            Salary{
                Basic: 17000.00,
                HRA:   5000.00,
                TA:    2200.00,
            },
        },
    }

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("test.json", file, 0644)
}
```

### Output

```jsx
{
 "FirstName": "Mark",
 "LastName": "Jones",
 "Email": "mark@gmail.com",
 "Age": 25,
 "MonthlySalary": [
  {
   "Basic": 15000,
   "HRA": 5000,
   "TA": 2000
  },
  {
   "Basic": 16000,
   "HRA": 5000,
   "TA": 2100
  },
  {
   "Basic": 17000,
   "HRA": 5000,
   "TA": 2200
  }
 ]
}
```

---

## Reading Text File

The bufio package `Scanner` generally used for reading the text by lines or words from a file. The following source code snippet shows reading text line-by-line from the plain text file as below.

The os.Open() function is used to open a specific text file in read-only mode and this returns a pointer of type os.File. The method os.File.Close() is called on the os.File object to close the file and there is a loop to iterates through and prints each of the slice values. The program after execution shows the below output line-by-line as they read it from the file.

### Example

```jsx
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for _, eachline := range txtlines {
		fmt.Println(eachline)
	}
}
```

### Output

```jsx
Lorem ipsum dolor sit amet, consectetur adipiscing elit.
Nunc a mi dapibus, faucibus mauris eu, fermentum ligula.
Donec in mauris ut justo eleifend dapibus.
Donec eu erat sit amet velit auctor tempus id eget mauris.
```

---

## Writing Text File

The bufio package provides an efficient buffered `Writer` which queues up bytes until a threshold is reached and then finishes the write operation to a file with minimum resources. The following source code snippet shows writing a string slice to a plain text file line-by-line.

The sampledata is represented as a string slice which holds few lines of data which will be written to a new line within the file. The function os.OpenFile() is used with a flag combination to create a write-only file if none exists and appends to the file when writing.

### Example

```jsx
package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	sampledata := []string{"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		"Nunc a mi dapibus, faucibus mauris eu, fermentum ligula.",
		"Donec in mauris ut justo eleifend dapibus.",
		"Donec eu erat sit amet velit auctor tempus id eget mauris.",
	}

	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, data := range sampledata {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()
	file.Close()
}
```

---

## Reading CSV File

The csv package have a `NewReader()` function which returns a `Reader` object to process CSV data. A `csv.Reader` converts \\r\\n sequences in its input to just \\n, which includes multi line field values also.

The file test.csv have few records is opened in read-only mode using the `os.Open()` function, which returns an pointer type instance of `os.File`. The `csv.Reader.Read()` method is used to decode each file record into pre-defined struct `CSVData` and then store them in a slice until `io.EOF` is returned.

### Example

```jsx
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for _, eachline := range txtlines {
		fmt.Println(eachline)
	}
}
```

### Output

```jsx
Name -- City -- Job
John -- London -- CA
Micky -- Paris -- IT
```

---

## Writing CSV File

The csv package have a `NewWriter()` function which returns a `Writer` object which is used for writing CSV data. A `csv.Writer` writes csv records which are terminated by a newline and uses a comma as the field delimiter. The following source code snippet shows how to write data to a CSV file.

A two-dimensional slice rows contains sample csv records. The os.Create() function creates a csv file test.csv; truncate all it's records if already exists and returning an instance of os.File object. The csvwriter.Write(row) method is called to write each slice of strings to the file as CSV records.

### Example

```jsx
package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	rows := [][]string{
		{"Name", "City", "Language"},
		{"Pinky", "London", "Python"},
		{"Nicky", "Paris", "Golang"},
		{"Micky", "Tokyo", "Php"},
	}

	csvfile, err := os.Create("test.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvfile)

	for _, row := range rows {
		_ = csvwriter.Write(row)
	}

	csvwriter.Flush()

	csvfile.Close()
}
```
