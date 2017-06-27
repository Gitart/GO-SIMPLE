
// https://golang.org/pkg/encoding/csv/#example_Reader
// https://www.dotnetperls.com/csv-go

package main


import (
	"encoding/csv"
	"crypto/rand"
	"encoding/hex"     
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
  r "github.com/dancannon/gorethink"
)



type Mst map[string]interface{}
// // Текущее время
// var CurTime = time.Now().Format("2006-01-02 15:04:05")

// Структура данных
type DTp struct {
	ID                     int64
	QUANT                  int64
	QUANT_STOCK            int64
	NO_USE_STOCK           int64
	COEF_VAT               int
	PRICE_FACTORY          int64
	PRICE_BUY              int64
	PRICE_BUY_VAT          int64
	PRICE_BUY_SUM          int64
	PRICE_SELL             int64
	PRICE_SELL_VAT         int64
	PRICE_SELL_SUM         int64
	PRICE_SRC              int64
	PRICE_SRC_VAT          int64
	PRICE_SRC_SUM          int64
	ID_DOC                 int64
	ID_BUSINESS            int64
	ID_STRUCTURE           int64
	IS_CLOSED              int64
	ID_DOC_TYPE            int64
	ID_VAT                 int64
	NO_VAT                 int64
	ID_CONTRACTOR          int64
	DOC_NAME               string
	DOC_DATE_TIME          string
	TAX_NAME               int64
	TAX_DATE               string
	ID_PARSEL              int64
	ID_PARSEL_PARENT       int64
	ID_DRUG                int64
	ID_DOC_BUY             int64
	ID_STRUCTURE_BUY       int64
	DOC_NAME_BUY           string
	DOC_DATE_TIME_BUY      string
	ID_SUPPLIER_BUY        int64
	DOC_NAME_SUPP_BUY      int64
	DOC_DATE_SUPP_BUY      int64
	DATE_PAY_BUY           int
	CODE                   int
	SERIES                 int
	DATE_EXPIRE            string
	CERT_NAME              int
	CERT_DATE              int
	NON_FISC               int
	CARD_NUMBER            int
	AMOUNT_SELL_DIS        int
	ID_SOURCE_RECEIPT_ITEM int
}



type TestRecord struct {
	Email string
	Date  string
}


// Настройка экспорта и чтение файла
type Reader struct {
	Comma            rune // field delimiter (set to ',' by NewReader)
	Comment          rune // comment character for start of line
	FieldsPerRecord  int  // number of expected fields per record
	LazyQuotes       bool // allow lazy quotes
	TrailingComma    bool // ignored; here for backwards compatibility
	TrimLeadingSpace bool // trim leading space
	                      // contains filtered or unexported fields
}




/*
	uid	
	data	
	code	
	company	
	companycode	
	customer	
	customercode	
	client	
	contractdate	
	contractcode	
	periodbegin	
	periodend	
	currency	
	ndsrate	
	totalnet	
	total	
	brand	
	cf_code	
	cf_name	
	cf_uid
*/


func main() {

	// Параметры
	NameBase    := os.Args[1]   // Database
	NameTables  := os.Args[2]   // Tables
	NameFile    := os.Args[3]   // Name file CSV
    DelData     := os.Args[4]   // A = Added D = Deleted
    Namedata    := os.Args[5]   // Name
	Ro          := r.InsertOpts{ReturnChanges: false, Durability : "soft"}
	k           := 0

    

	csvfile, err := os.Open(NameFile)

	if err != nil {
	   log.Fatal("Erorr сonvert file CSV.")
	}

	defer csvfile.Close()

    // Установки для считывания файла
	reader := csv.NewReader(csvfile)
	       reader.FieldsPerRecord  = 0                 // Строка считывания 0-игнорировать строку наименования полей -1 если нет наименования полей
	       reader.Comma            = '\t'              // Разделитель \t - табулятор   ;-точка с запятой
	       reader.Comment          = '#'               // Комментарии
           reader.LazyQuotes       = true              // Игнорирует пропуски в поле после ТАБ   
           reader.TrailingComma    = true              // retain rather than remove empty slots
           reader.TrimLeadingSpace = false             // retain rather than remove empty slots 

	// Чтение
	rawCSVdata, err := reader.ReadAll()

	// Обработка ошибок
	if err != nil {
	   log.Println("ERR LOAD FILE : ", err)
	   // os.Exit(1)
	}

	// Newd, err := GenUUID()
	// Mss := time.Millisecond.String()
	// Mss := time.Now().String()
	// var Newd = "St_" + time.Now().String() + Mss
	// fmt.Printf("Ñòàðò  !!!  ", CurTime + " " + Newd + "\n" )
    // var CurTime = time.Now().Format("2006-01-02 15:04:05")

	// Контроль данных
	log.Printf("Start load data.")
    log.Printf("Before load %v ", CountRec("Test"))


	// Подключение
  AddressPort  :="111.111.111.111:2222"
	session, err := r.Connect(r.ConnectOpts{Address: AddressPort, Database: NameBase})

	// Обработка ошибка
	if err != nil {
	   log.Fatalln(err)
	}

	// Создание таблицы для заливки с мягкой вставкой
	// r.Db("test").CreateTable("Docum",{durability:soft})


	// База  данных таблица
	DT := r.DB(NameBase).Table(NameTables)
 

    // Deleted data before load data in table
    if DelData == "d" || DelData == "D" || DelData == "Del" || DelData == "del" || DelData == "Delete" || DelData == "delete" {
       DT.Delete().Exec(session)
       log.Println("Data was deleted.")
     } else {
       log.Println("Data was append to existing data.")
     }


	// Чтение файла СSV
	for _, each := range rawCSVdata {
		    k++
		
		err:=DT.Insert(Mst{
			"Ids":                     k,
			"Date_load":               time.Now().Format("2006-01-02 15:04:05"),
			"Id":                      each[0],
			"Date":                    each[1],
			"Code":                    each[2],
			"Company":                 each[3],
			"Company_code":            each[4],
			"Customer":                each[5],
			"Customer_code":           each[6],
			"Client":                  each[7],
			"Contractdate":            each[8],
			"Contractcode":            each[9],
			"Periodbegin":             each[10],
			"Periodend":               each[11],
			"Currency":                each[12],
			"Ndsrate":                 each[13],
			"Totalnet":                each[14],
			"Total":                   Cvt(each[15]),
			"Brand":                   Cvt(each[16]),
			"Cf_code":                 each[17],
			"Cf_name":                 each[18],
			"Cf_uid":                  each[19],
			"Flag":                    Namedata,
			},Ro).Exec(session)
	    

	        if err!=nil {
		       fmt.Println(err)
			   return
		    }
	
		// Просмотр 
		// fmt.Printf("Record %v  ID %s \n", k, each[0])
		// csvfile.Close()
	}

	// second sanity check, dump out allRecords and see if
	// individual record can be accessible
	// fmt.Println(allRecords)
	// fmt.Println(allRecords[1].Email)
	// fmt.Println(allRecords[1].Date)
	// Mss = time.Now().String()
	// var Newd = "St_" + time.Now().String() + Mss

	// Окончание загрузки
	// CurTime = time.Now().Format("2006-01-02 15:04:05")
	log.Printf("Load records  : %v", k)
    log.Printf("All records in table %v", CountRec("Test"))
}



///*********************************************************************************************************
// Generation GUID code
//*********************************************************************************************************
func GenUUID() (string, error) {
	uuid   := make([]byte, 16)
	n, err := rand.Read(uuid)
	
	if n != len(uuid) || err != nil {
	   return "", err
	}

	// TODO: verify the two lines implement RFC 4122 correctly
	uuid[8] = 0x80 // variant bits see page 5
	uuid[4] = 0x40 // version 4 Pseudo Random, see page 7
	return hex.EncodeToString(uuid), nil
}

//*********************************************************************************************************
// Convert string to float  with replace comma
// 
//*********************************************************************************************************
func Cvt(Strr string) float64 {
	
    // Замена запятой точкой
    St:=strings.Replace(Strr, ",", ".", 2)
 
	if len(Strr) == 0 {
	   Strr = "0.00"
	}

	ttt, err := strconv.ParseFloat(St, 64)
    // ttt, err := strconv.ParseInt(Strr, 10, 64)
    // fmt.Printf("SRTCONVERY....   %s \n)", Strr )

    // Возврат
	if err != nil {
		return 0.00
	  } else {
	    return ttt
	}
}



// Подсчет количества записей
func CountRec(NameTable string) int64 {

	var response int64
    // Подключение
    AddressPort  :="111.222.444.111:2222"
	session, err := r.Connect(r.ConnectOpts{Address: AddressPort, Database: "Work"})

	// Обработка ошибка
	if err != nil {
	   log.Fatalln(err)
	}

	// База  данных таблица
	res,_ := r.DB("Work").Table(NameTable).Count().Run(session)
    defer res.Close()

    // Error
	if err != nil {
		log.Println(err)
		panic("No document ...")
	}

	err = res.One(&response)

	// Error
	if err != nil {
	   fmt.Println("Document # Absent ")
	   log.Println(err)
	   panic("No document")
	}
	return response
}
