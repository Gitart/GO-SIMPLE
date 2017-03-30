##  Save map/struct to JSON or XML file

Previous tutorial on converting map/slice/array to JSON or XML format is for output to web via net/http package. 
This tutorial is a slight modification and save the output to JSON or XML file instead.


```golang
 package main

 import (
         "encoding/json"
         "encoding/xml"
         "fmt"
         "io"
         "os"
         "strconv"
 )

 type Person struct {
         Name string `json:"name"`
         Age  int    `json:"age"`
 }

 func main() {

         // create and populate a map from dummy JSON data

         dataStr := `{"Name":"Dummy","Age":0}`

         personMap := make(map[string]interface{})

         err := json.Unmarshal([]byte(dataStr), &personMap)

         if err != nil {
                 panic(err)
         }

         var onePerson Person

         // convert map to Person struct
         onePerson.Name = fmt.Sprintf("%s", personMap["Name"])
         onePerson.Age, _ = strconv.Atoi(fmt.Sprintf("%v", personMap["Age"]))

         jsonData, err := json.Marshal(onePerson)

         if err != nil {
                 panic(err)
         }

         // sanity check
         fmt.Println(string(jsonData))

         // write to JSON file

         jsonFile, err := os.Create("./Person.json")

         if err != nil {
                 panic(err)
         }
         defer jsonFile.Close()

         jsonFile.Write(jsonData)
         jsonFile.Close()
         fmt.Println("JSON data written to ", jsonFile.Name())

         // write to XML file

         xmlFile, err := os.Create("./Person.xml")
         if err != nil {
                 panic(err)
         }
         defer xmlFile.Close()

         xmlWriter := io.Writer(xmlFile)

         enc := xml.NewEncoder(xmlWriter)
         enc.Indent("  ", "    ")
         if err := enc.Encode(onePerson); err != nil {
                 fmt.Printf("error: %v\n", err)
         }

         xmlFile.Close()
         fmt.Println("XML data written to ", xmlFile.Name())

 }
 ```
 
Output:
```
{"name":"Dummy","age":0}
JSON data written to ./Person.json
XML data written to ./Person.xml
```



## Covert map/slice/array to JSON or XML format

### Problem :
You have a data struct in map, slice or array format and you have to convert the data to JSON or XML format. How to do that?

### Solution :
Convert the map, slice or array data to JSON with json.Marshal() function. In the struct, remember to add struct tags for JSON or XML.
NOTE : Set Content-Type to application/json or application/xml when writing out as HTTP response.

```golang
For example :

 package main

 import (
 	"encoding/json"
 	"fmt"
 	"net/http"
 )

 type KeyPair struct {
 	Id   int    `json:"id"` // <--- json struct tags
 	Name string `json:"name"` // <--- json struct tags
 }

 func Home(w http.ResponseWriter, r *http.Request) {

 	KP := KeyPair{Id: 1, Name: "Adam"}
 	fmt.Println(KP)

 	byte, err := json.Marshal(KP) // <---- here !

 	if err != nil {
 		return
 	}

 	w.Header().Set("Content-Type", "application/json") // <---- here !
 	fmt.Fprint(w, string(byte))

 	fmt.Println(string(byte))
 }

 func main() {

 	http.HandleFunc("/", Home)
 	http.ListenAndServe(":8080", nil)
 }
 ```
 
 ## XML to JSON example

For this tutorial, we will learn how to read data from XML file, process the data and save the output to JSON format.  
Converting XML to JSON data format can be done easily with the Golang's encoding/xml and encoding/json packages.  

Create the Employees.xml file with this content :


```xml
 <?xml version="1.0"?>
  <company>
          <staff>
                  <id>101</id>
                  <firstname>Derek</firstname>
                  <lastname>Young</lastname>
                  <username>derekyoung</username>
          </staff>
          <staff>
                  <id>102</id>
                  <firstname>John</firstname>
                  <lastname>Smith</lastname>
                  <username>johnsmith</username>
          </staff>
  </company>
  ```
  
and the code to eat this XML data and poop out JSON file :


```golang
 package main

 import (
         "encoding/json"
         "encoding/xml"
         "fmt"
         "io/ioutil"
         "os"
 )

 type jsonStaff struct {
         ID        int
         FirstName string
         LastName  string
         UserName  string
 }

 type Staff struct {
         XMLName   xml.Name `xml:"staff"`
         ID        int      `xml:"id"`
         FirstName string   `xml:"firstname"`
         LastName  string   `xml:"lastname"`
         UserName  string   `xml:"username"`
 }

 type Company struct {
         XMLName xml.Name `xml:"company"`
         Staffs  []Staff  `xml:"staff"`
 }

 func (s Staff) String() string {
         return fmt.Sprintf("\t ID : %d - FirstName : %s - LastName : %s - UserName : %s \n", s.ID, s.FirstName, s.LastName, s.UserName)
 }

 func main() {
         xmlFile, err := os.Open("Employees.xml")
         if err != nil {
                 fmt.Println("Error opening file:", err)
                 return
         }
         defer xmlFile.Close()

         XMLdata, _ := ioutil.ReadAll(xmlFile)

         var c Company
         xml.Unmarshal(XMLdata, &c)

         // sanity check - XML level
         fmt.Println(c.Staffs)

         // convert to JSON
         var oneStaff jsonStaff
         var allStaffs []jsonStaff

         for _, value := range c.Staffs {
                 oneStaff.ID = value.ID
                 oneStaff.FirstName = value.FirstName
                 oneStaff.LastName = value.LastName
                 oneStaff.UserName = value.UserName

                 allStaffs = append(allStaffs, oneStaff)
         }

         jsonData, err := json.Marshal(allStaffs)

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         // sanity check - JSON level

         fmt.Println(string(jsonData))

         // now write to JSON file

         jsonFile, err := os.Create("./Employees.json")

         if err != nil {
                 fmt.Println(err)
         }
         defer jsonFile.Close()

         jsonFile.Write(jsonData)
         jsonFile.Close()

 }
 ```
 
run the code above and you should be able to see a new Employees.json file appear in the same directory.

Employees.json
```json
  [
     {
         "ID": 101,
         "FirstName": "Derek",
         "LastName": "Young",
         "UserName": "derekyoung"
     },
     {
         "ID": 102,
         "FirstName": "John",
         "LastName": "Smith",
         "UserName": "johnsmith"
     }
 ]
 ```
 
 
## Convert CSV data to JSON format and save to file
Need to load a CSV data file and save it to JSON encoded file or stream it out ... like to a http service ? 
This tutorial will cover just that :

The Golang code below will first read this data.csv data file :

```csv
  Adam,36,CEO
  Eve,34,CFO
  Mike,38,COO
```

and output to data.json file
```json
  [
  {"Name":"Adam","Age":36,"Job":"CEO"},
  {"Name":"Eve","Age":34,"Job":"CFO"},
  {"Name":"Mike","Age":38,"Job":"COO"}
  ]
```

### csv2json.go


```golang
 package main

 import (
         "encoding/csv"
         "encoding/json"
         "fmt"
         "os"
         "strconv"
 )

 type Employee struct {
         Name string
         Age  int
         Job  string
 }

 func main() {
         // read data from CSV file

         csvFile, err := os.Open("./data.csv")

         if err != nil {
                 fmt.Println(err)
         }

         defer csvFile.Close()

         reader := csv.NewReader(csvFile)

         reader.FieldsPerRecord = -1

         csvData, err := reader.ReadAll()

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         var oneRecord Employee
         var allRecords []Employee

         for _, each := range csvData {
                 oneRecord.Name = each[0]
                 oneRecord.Age, _ = strconv.Atoi(each[1]) // need to cast integer to string
                 oneRecord.Job = each[2]
                 allRecords = append(allRecords, oneRecord)
         }

         jsondata, err := json.Marshal(allRecords) // convert to JSON

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         // sanity check
         // NOTE : You can stream the JSON data to http service as well instead of saving to file
         fmt.Println(string(jsondata))

         // now write to JSON file

         jsonFile, err := os.Create("./data.json")

         if err != nil {
                 fmt.Println(err)
         }
         defer jsonFile.Close()

         jsonFile.Write(jsondata)
         jsonFile.Close()
 }
 ```
 
Output :
```
[{"Name":"Adam","Age":36,"Job":"CEO"},{"Name":"Eve","Age":34,"Job":"CFO"},{"Name":"Mike","Age":38,"Job":"COO"}]
```


## Encoding 


```golang
 package main

 import (
         "encoding/json"
         "fmt"
         "os"
 )

 type Employee struct {
         Name string
         Age  int
         Job  string
 }

 func main() {

         worker := Employee{
                 Name: "Adam",
                 Age:  36,
                 Job:  "CEO",
         }

         output, err := json.Marshal(worker) // <--- here

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         fmt.Println(string(output))
       // os.Stdout.Write(b) -- also ok

 }
 ```
 
Output :
```
{"Name":"Adam","Age":36,"Job":"CEO"}
```




## How to read CSV file
One way or another, CSV file is going to be part and parcel of a developer life. 
A programmer will bound to meet up with CSV file one day. In this tutorial, we will show you how to read CSV  
file wtih Go. Below is a simple code in Go demonstrating the capability.   


```golang
 package main

 import (
         "encoding/csv"
         "fmt"
         "os"
 )

 func main() {

         csvfile, err := os.Open("somecsvfile.csv")

         if err != nil {
                 fmt.Println(err)
                 return
         }

         defer csvfile.Close()

         reader := csv.NewReader(csvfile)

         reader.FieldsPerRecord = -1 // see the Reader struct information below

         rawCSVdata, err := reader.ReadAll()

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         // sanity check, display to standard output
         for _, each := range rawCSVdata {
                 fmt.Printf("email : %s and timestamp : %s\n", each[0], each[1])
         }
 }
 ```
 
See how to load CSV values into Struct at https://www.socketloop.com/tutorials/how-to-unmarshal-or-load-csv-record-into-struct-go

The CSV file contains the following data
more somecsvfile.csv
```csv
"jenniferlcl@*****.com","2012-07-03 18:38:06"
"norazlinjumali@*****.com","2010-06-26 19:46:08"
"wilfred5571@*****.com","2010-07-02 21:49:55"
"nas_kas81@*****.com","2010-07-06 12:49:31"
"tammyu3622@*****.com","2010-07-06 13:55:21"
"wakrie@*****.com","2012-03-02 11:00:59"
"yst.shirin@*****.com","2010-07-07 10:19:11"
"annl_107@*****.com","2010-07-07 20:55:59"
"jen_5831@*****.com","2010-07-07 21:12:27"
"hsheyli@*****.com","2011-09-07 00:39:11"
The Reader has the following data structure :
```


```golang
 type Reader struct {
         Comma            rune // field delimiter (set to ',' by NewReader)
         Comment          rune // comment character for start of line
         FieldsPerRecord  int  // number of expected fields per record
         LazyQuotes       bool // allow lazy quotes
         TrailingComma    bool // ignored; here for backwards compatibility
         TrimLeadingSpace bool // trim leading space
         // contains filtered or unexported fields
 }
 ```
