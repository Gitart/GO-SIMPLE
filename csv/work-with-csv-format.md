## Чтение CSV файла из сети по URL
- https://gist.github.com/stupidbodo/71f2b164744a18a18e74  
- https://www.jernejsila.com/2015/11/12/exporting-data-as-csv-file-from-web-apps-with-golang/ 


## Загрузка CSV файл по URL с записью в базу данных

```Go
// Загрузка в формате CSV
// Формат должен быть опредлен заранее
func Api_bi_load_csv(w http.ResponseWriter, req *http.Request){
     type Needs struct{
      	 Iddrug  string        
       	 Count   string            
     }

     var Nd []Needs
     var Nrec Needs

     reader := csv.NewReader(req.Body)
	 
     // Settings Reader
     // https://golang.org/pkg/encoding/csv/#example_Reader
     reader.Comma =';'
     reader.FieldsPerRecord=-1
     rawCSVdata, err:= reader.ReadAll()
         
     if err != nil {
        fmt.Println(err)
        os.Exit(1)
     }
         
     for i, each := range rawCSVdata {
          if i>1{
             fmt.Printf("#%v  DRUG %s  = COUNT %s\n",i, each[0], each[1])
             Nrec.Count  = each[1]
             Nrec.Iddrug = each[0]
	
	     // Через индекс
             // Nd[i].Iddrug =   each[0]
             // Nd[i].Count  =   each[1]
             Nd = append(Nd, Nrec)
	    }	
       }
       
       // Records in database
       err:=r.DB("Bi").Table("Analysis").Insert(Nd).Exec(sessionArray[0])
       if err!=nil{
          return
       }
}

```

## Вариант с адресом

```Go
package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
)

func readCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ';'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}


func main() {
    url := "https://raw.githubusercontent.com/mledoze/countries/master/dist/countries.csv"
	data, err := readCSVFromUrl(url)
	if err != nil {
		panic(err)
	}

	for idx, row := range data {
		// skip header
		if idx == 0 {
			continue
		}

		if idx == 6 {
			break
		}

		fmt.Println(row[2])
	}
}

// Will Print:
// AF
// AX
// AL
// DZ
// AS
```

## Чтение из сети JSON файл

```Go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Country struct {
	CountryName        CountryName       `json:"name"`
	TLD                []string          `json:"tld"`
	CCA2               string            `json:"cca2"`
	CCN3               string            `json:"ccn3"`
	CCA3               string            `json:"cca3"`
	Currency           []string          `json:"currency"`
	CallingCode        []string          `json:"callingCode"`
	Capital            string            `json:"capital"`
	AlternateSpellings []string          `json:"altSpellings"`
	Relevance          string            `json:"relevance"`
	Region             string            `json:"region"`
	Subregion          string            `json:"subregion"`
	NativeLanguage     string            `json:"nativeLanguage"`
	Languages          map[string]string `json:"languages"`
	Translations       map[string]string `json:"translations"`
	LatLng             [2]float64        `json:"latlng"`
	Demonym            string            `json:"demonym"`
	Borders            []string          `json:"borders"`
	Area               float64           `json:"area"`
}

type CountryName struct {
	Common   string            `json:"common"`
	Official string            `json:"official"`
	Native   CountryNameNative `json:"native"`
}

type CountryNameNative struct {
	Common   string `json:"common"`
	Official string `json:"official"`
}

func readJSONFromUrl(url string) ([]Country, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var countryList []Country
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, &countryList); err != nil {
		return nil, err
	}

	return countryList, nil
}

func main() {
    url := "https://raw.githubusercontent.com/mledoze/countries/master/dist/countries.json"
	countryList, err := readJSONFromUrl(url)
	if err != nil {
		panic(err)
	}

	for idx, row := range countryList {
		// skip header
		if idx == 0 {
			continue
		}

		if idx == 6 {
			break
		}

		fmt.Println(row.CountryName.Common)
	}
}

// Will Print:
// Åland Islands
// Albania
// Algeria
// American Samoa
// Andorra
```
