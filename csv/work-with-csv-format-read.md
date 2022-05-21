## We will process the tabdata.csv file
- https://www.socketloop.com/tutorials/golang-read-tab-delimited-file-with-encoding-csv-package?spot_im_comment_id=sp_L6HOxGDv_tutorial-790_c_Cait1x

## Format CSV
 Adam    36      CEO  
 Eve     34      CFO  
 Mike    38      COO  

### with this code :

```Go
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

         csvFile, err := os.Open("./tabdata.csv")
         if err != nil {
                 fmt.Println(err)
         }

         defer csvFile.Close()
         reader := csv.NewReader(csvFile)
         reader.Comma = '\t' // Use tab-delimited instead of comma <---- here!
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
 
