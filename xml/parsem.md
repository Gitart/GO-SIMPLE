

```xml
<ROWDATA>
<ROW>
  <ПІБ>    ПОПКО    РУСЛАН ВАСИЛЬОВИЧ</ПІБ>
  <Місце_проживання>61112, Харківська обл., місто Харків, Московський район, ПРОСПЕКТ П'ЯТДЕСЯТИРІЧЧЯ ВЛКСМ, будинок 86, квартира 65</Місце_проживання>
  <Основний_вид_діяльності>45.32 Роздрібна торгівля деталями та приладдям для автотранспортних засобів</Основний_вид_діяльності>
  <Стан>зареєстровано</Стан>
</ROW>
</ROWDATA>
```


## Program
```golang
 package main
    import (
    "encoding/xml"
    "fmt"
    "golang.org/x/text/encoding/charmap"
    "golang.org/x/text/transform"
    "io/ioutil"
    "os"
    "strings"
)

type Rowdata struct {
    XMLName xml.Name `xml:"ROWDATA"`
    Rowdata []Row    `xml:"ROW"`
}

type Row struct {
    XMLName  xml.Name `xml:"ROW"`
    Location string   `xml:"Місце_проживання"`
    Director string   `xml:"ПІБ"`
    Activity string   `xml:"Основний_вид_діяльності"`
    City     string   `xml:"Стан"`
}

func main() {
    xmlFile, err := os.Open("FOP_1.xml")
    if err != nil {
        fmt.Println(err)
    }

    defer xmlFile.Close()
    byteValue, _ := ioutil.ReadAll(xmlFile)
    koi8rString := transform.NewReader(strings.NewReader(string(byteValue)), charmap.Windows1251.NewDecoder())
    decBytes, _ := ioutil.ReadAll(koi8rString)
    var entries Rowdata
    xml.Unmarshal(decBytes, &entries)

    for i := 0; i < len(entries.Rowdata); i++ {
        fmt.Println("Name: " + entries.Rowdata[i].Director)
    }
```


### Test
```golang
import "golang.org/x/net/html/charset"

func main() {
    xmlFile, err := os.Open("FOP_1.xml")
    if err != nil {
        fmt.Println(err)
    }
    defer xmlFile.Close()

    var entries Rowdata
    parser := xml.NewDecoder(xmlFile)
    parser.CharsetReader = charset.NewReaderLabel
    err = parser.DecodeElement(&entries, nil)
    if err != nil {
        fmt.Println(err)
        return
    }

    for i := 0; i < len(entries.Rowdata); i++ {
        fmt.Println("Name: " + entries.Rowdata[i].Director)
    }
}
```
