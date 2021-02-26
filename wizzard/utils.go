package services

import (
  "net"
  "net/http"
	"io"
	"fmt"
	"time"
	"strings"
	"strconv"
	"crypto/rand"
	"encoding/json"
	"app/models"
	"reflect"
	"math"
	 mrn "math/rand"
	"regexp"
	"errors"
	"app/vars"
	"github.com/shopspring/decimal"      // Библиотека для работы с числами большой точностью округления
)

// **********************************************************
// Ping 
// **********************************************************
func PingIp (url string) (int, error) {
	tr := &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    1 * time.Second,
			Dial:               (&net.Dialer{ 
				                      Timeout:   2 * time.Second, 
				                      KeepAlive: 1 * time.Second}).Dial,
			DisableCompression: true,
    }


    client := &http.Client{Transport: tr}
    
    req, err := http.NewRequest("HEAD", url, nil)
    if err != nil {
       return 400, err
    }
    resp, err := client.Do(req)
    if err != nil {
       return 500, err
    }
    resp.Body.Close()
    return resp.StatusCode, nil
}

// *********************************************************
// Tcp
// *********************************************************
func PingIpTcp (url string) bool {
    Ret :=false

    timeout := time.Duration(1 * time.Second)
    _, err := net.DialTimeout("tcp", url, timeout)
  
    if err != nil {
       fmt.Printf("%s %s %s\n", url, "not responding", err.Error())
       Ret=true
    } else {
       fmt.Printf("%s %s %s\n", url, "responding on port")
       Ret=false
    }
      
    return Ret

}

// *****************************************************************
// Округление до 2 цифр и возвращение float64
// https://play.golang.org/p/KNhgeuU5sT
// https://www.socketloop.com/tutorials/golang-display-float-in-2-decimal-points-and-rounding-up-or-down
// *****************************************************************
func FloatRound (num float64) float64 {
   j1      := fmt.Sprintf("%.2f", num)        
   j2, err := strconv.ParseFloat(j1, 64)
   
   if err != nil {
      return 0.0
    }
   return j2
}

// *****************************************************************
// Округление до 2 цифр и возвращение string
// *****************************************************************
func FloatRoundStr (num float64) string {
     return fmt.Sprintf("%.2f", num)        
}

// *****************************************************************
// Преобразование string/Int64
// *****************************************************************
func Str_int64 (param string) int64 {
    number, err:= strconv.ParseInt(param, 10, 64)
    if err!=nil{
      return 0.0
    }
    return number
}

// *****************************************************************
// Преобразование string/Int
// *****************************************************************
func Str_int (param string) int {
   number, err := strconv.Atoi(param)
   if err!=nil{
      return 0.0
    }
    return number
}

// *****************************************************************
// Преобразование string/float64
// *****************************************************************
func Str_float64 (param string) float64 {
  
  // Для решения проблемы с gorm - который не позволяет сохранять 
  // нулевые и пустые значения
  if param == "0" {
  	 return 0.00001
  }

  if  number, err := strconv.ParseFloat(param, 64); err == nil{
      return number
  }else{
      return 0.0
  }
}

// *****************************************************************
// Преобразование int64/string
// *****************************************************************
func Int_str(num int) string {
    s := strconv.Itoa(num)
	return s
} 

// *****************************************************************
// Преобразование int64/string
// *****************************************************************
func Int64_str(num int64) string {
	s:=strconv.FormatInt(num, 10)
	return s
} 

// *****************************************************************
// Преобразование float64/string
// *****************************************************************
func Float64_str(num float64) string {
	s := fmt.Sprintf("%f", num)
	return s
} 

// *****************************************************************
// Код сформирован случайным образом
// на основе GUID и приведен к верхнему регистру
// Используется для разных целей
// *****************************************************************
func Sys_keyid() string {
	s := strings.ToUpper(strings.Replace(GENID(), "-", "", 1))
	return s
}

// *****************************************************************
 // Формирование UID (6 знаков)
 //
 // newUUID generates a random UUID according to RFC 4122
 // uuid, err := newUUID()
 //  	if err != nil {
 //  		fmt.Printf("error: %v\n", err)
 //  	}
 //  	fmt.Printf("%s\n", uuid)
 // *****************************************************************
func genUuid() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)

	if n != len(uuid) || err != nil {
		return " "
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80 
	uuid[6] = uuid[6]&^0xf0 | 0x40 
	return fmt.Sprintf("-%x", uuid[0:10]) 
}

// *****************************************************************
 // Формирование UID (6 знаков)
 //
 // newUUID generates a random UUID according to RFC 4122
 // uuid, err := newUUID()
 // 	if err != nil {
 // 		fmt.Printf("error: %v\n", err)
 // 	}
 // 	fmt.Printf("%s\n", uuid)
 // *****************************************************************
func GENID() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)

	if n != len(uuid) || err != nil {
		return " "
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80 
	uuid[6] = uuid[6]&^0xf0 | 0x40 
    return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

// *****************************************************************
// Округление до N знаков
// Библиотека с высокой точностьюdecimal.Decimal
// *****************************************************************
func ToFixedNum(num float64, precision int32)  decimal.Decimal {
     sumd:= decimal.NewFromFloat(num).Round(precision)
     return sumd
 }

// *****************************************************************
// Округление до N знаков
// *****************************************************************
func ToFixed(num float64, precision int) float64 {
     output := math.Pow(10, float64(precision))
     return float64(round(num * output)) / output
}

func round(num float64) int {
     return int(num + math.Copysign(0.5, num))
}

// *****************************************************************
// Marshal order to string
// *****************************************************************
func Marshl(order models.Order)  {
     slice, _ := json.Marshal(order)
     fmt.Println(string(slice))
}

// *****************************************************************
// TITLE           : Translate str date format to str date froamt2
// DATE & TIME     : 10.05.2016 14:39 
// DESCRIPTION     : Возврат двух дат с разницей в часах между текущим временем и вычисляемым
//                   Дата возвращается в строковом виде
//                   На вход подается параметр в часах (строковой тип) 
//                   на выходе два параметра текущая дата и дата на несколько часов назад (входной параметр)
// *****************************************************************
func FormHrs(Hrs string) (string) {
    F  := "02-01-2006 15:04"
    S,_:= time.Parse(F, Hrs)
    // return     r.Term(S.String())
    return S.String()
}

// *****************************************************************
// TITLE         : Проверка ввода корректного пароля
// DESCRIPTION   : Проверка пароля по маске
// USAGE         : checkPassword("pass")
// *****************************************************************
func checkPassword(password string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", password); !ok {
		return false
	}
	return true
}

// *****************************************************************
// TITLE         : Генерация символьных строк - кодов
// ORGANIZATION  : 
// DATE          : 01-04-2015 10:59
// DESCRIPTION   : Functions for generating random strings
//                 Три функции генерирующие последовательность кодов
//                 Применяется для использования создания уникальных имен Cookies
//                 Хотя область применения может быть различной
// USAGE         : RandomString(10), RandomHumanFriendlyString(10), RandomCookieFriendlyString(5)
// AUTHOR        : Savchenko Arthur
// CODE          : Worked with Cookies
// SOURCE        : https://github.com/xyproto/permissions2/blob/master/cookies.go
// IMPORT        : import (	mrn "math/rand")

// Generate a random string of the given length.
// Выводит вместе с нечетабильными символами !!!
// *****************************************************************
func RandomString(length int) string {
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = byte(mrn.Int63() & 0xff)
	}
	return string(b)
}

// *****************************************************************
//  Generate a random, but human-friendly, string of the given length.
//  Should be possible to read out loud and send in an email without problems.
//  Samples : ubipamus
// *****************************************************************
func RandomHumanFriendlyString(length int) string {
	const (
		vowels     = "aeiouy"                   // email+browsers didn't like "æøå" too much
		consonants = "bcdfghjklmnpqrstvwxz"
	)

	b := make([]byte, length)
	for i := 0; i < length; i++ {
		if i%2 == 0 {
			b[i] = vowels[mrn.Intn(len(vowels))]
		} else {
			b[i] = consonants[mrn.Intn(len(consonants))]
		}
	}
	return string(b)
}

// *****************************************************************
// Generate a random, but cookie-friendly, string of the given length.
// *****************************************************************
func RandomCookieFriendlyString(length int) string {
	 const allowed = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	 b     := make([]byte, length)
	 for i := 0; i < length; i++ {
	 	   b[i] = allowed[mrn.Intn(len(allowed))]
	 }
	 return string(b)
}

// *****************************************************************
// Проверка правильного формата почты
// *****************************************************************
func checkEmail(str string) error {
	if m, err := regexp.MatchString(`^[^@]+@[^@]+$`, str); !m {
	   fmt.Println("err = ", err)
	   return errors.New("Please enter a valid email address.")
	}
	   return nil
}

// *****************************************************************
// *  Convert date 
// *****************************************************************
func checkDate(str string) error {
	_, err := time.Parse("01022006", str)
	
	if err != nil {
		_, err = time.Parse("20060102", str)
	}

	if str == "" || err != nil {
	   return errors.New("Please enter a valid Date.")
	}
	return nil
}


/****************************************************************************************
 * Check null is interface
/****************************************************************************************/
func isNil(a interface{}) bool {
     defer func() { 
           recover() 
      }()
     return a == nil || reflect.ValueOf(a).IsNil()
}


/****************************************************************************************
 * Check null is interface
/****************************************************************************************/
func ISNIL(Vars string) string{
     defer func() { recover() }()
     
     if len(Vars) == 0 || reflect.ValueOf(Vars).IsNil()  {
        return "33"
     }
     return Vars
}

// ***************************************************
// Форматирование чисел с разбивкой на триады
// 1,234,567.45
// https://play.golang.org/p/vnsAV23nUXv
// https://play.golang.org/p/633kC8cKTcF
// ***************************************************
func FormatCommas(num float64) string {
    str := fmt.Sprintf("%.2f", num)
    re := regexp.MustCompile("(\\d+)(\\d{3})")
    for n := ""; n != str; {
        n = str
        str = re.ReplaceAllString(str, "$1 $2")
    }
    return str
}

// Test // 1,234,567.45
func FormatCommasTest() {
	   fmt.Println(FormatCommas(12345678.23))   
}


// **************************************************************
// Format date to string  
// **************************************************************
func FormatDateString(datetime time.Time) string {
     return datetime.Format("02.01.2006")
}

func FormatXls(number float64) string {
	   num := fmt.Sprintf("%.2f", number)
     num =  strings.Replace(num,".",",",-1)
     return  num
}

// **************************************************************
// Increment 
// **************************************************************
func Inc(n int) int {
     return n + 1
}

// **************************************************************
// Generator Number doc
// **************************************************************
func NumGenerator(typeorder int64) string {
         Pr:=""
         // Если не был вставлен номер 
         if typeorder == -1 {
            Pr = "R"
         }else{
            Pr = "P" 
         }
         return  Pr + NewNumber()
}

// *****************************************************************
// Формирование номера 
// *****************************************************************
func NewNumber() string {
	return strconv.FormatInt(time.Now().Unix(),10)
     // return time.Now().Format("02012006154050") // + "-" +  RandomCookieFriendlyString(5)
}

func NewNumberStr() string {
     return time.Now().Format("02012006154050") // + "-" +  RandomCookieFriendlyString(5)
}

// ******************************************************
// StructToMap Converts a struct to a map while maintaining the json alias as keys
// ******************************************************
func StructToMap(obj interface{}) (newMap map[string]interface{}, err error) {
    data, err := json.Marshal(obj)

    if err != nil {
        return
    }

    err = json.Unmarshal(data, &newMap) // Convert to a map
    return
}

// ******************************************************
// Get Hash <=Login Password
// ******************************************************
func GetHash (Login, Password string) string{
    var userCredsString = Login + Password
    var hash = CryptoSha256(userCredsString, string(vars.SIGN_KEY))
    return hash
}

// **************************************************
// Unix Time to Date
// **************************************************
func DateFormat(layout string, d float64) string{
    intTime := int64(d)
    t := time.Unix(intTime, 0)
    if layout == "" {
       layout = "02.01.2006 15:04:05"
    }
    return t.Format(layout)
}

// **************************************************
// Generator key 16
// **************************************************
func KeyLink(key string) string {
  hex := fmt.Sprintf("%x", key)
  return hex
}

// **************************************************
// Chek key
// **************************************************
func ChekKeyLink(key, id string) bool {
  if KeyLink(key) == id {
     return true
  }
  return false
}

