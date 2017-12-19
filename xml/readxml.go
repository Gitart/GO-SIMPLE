 package main

 import (
         "fmt"
         "io/ioutil"
         "os"
         "encoding/xml"
 )


 type Staff struct {
        XMLName xml.Name `xml:"staff"`
        ID int `xml:"id"`
        FirstName string `xml:"firstname"`
        LastName string `xml:"lastname"`
        UserName string `xml:"username"`
 }

 type Company struct {
        XMLName xml.Name `xml:"company"`
        Staffs []Staff `xml:"staff"`
 }

 func (s Staff) String() string {
         return fmt.Sprintf("\t ID : %d - FirstName : %s - LastName : %s - UserName : %s \n", s.ID,  s.FirstName , s.LastName, s.UserName)
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

         fmt.Println(c.Staffs)
 }
