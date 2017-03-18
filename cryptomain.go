package controllers

import (
        "github.com/revel/revel"
        "bars/app/models"
        "time"
        "net/http"
        "fmt"
        "crypto/rand"
        "strings"
        "encoding/base64"                                                                                                                                                                                       
        "crypto/aes"                                                                                                                                                                                            
        "crypto/cipher"  
        "io" 
)

const (
          sepr   = "*"
       // secKey = "SecKsCsertxfY0d2sdx003-d2s93VA2Y"                                // Ключ шифрования - расшифрования
          secKey = "XYUlF_SecretCode"                                                // Ключ шифрования - расшифрования
          salt   = "SweorpXc00937CvvsdfdhfhjkkGhd74jfkdmbvkfHjdkJHG-384hdhs"         // Ключ шифрования - расшифрования 
)

// Cookies structure
type SecCook struct {
     Uid      string 
     Name     string
     Surname  string
     Email    string      
} 

type AuthController struct {
     BaseController
}

func (c AuthController) LoginPage() revel.Result {
     return c.RenderTemplate("App/auth/login.html")
}

func (c AuthController) RegPage() revel.Result {
     return c.RenderTemplate("App/auth/reg.html")
}


/**
* @api {post} /signin/   Login()
* @apiName               Login()
* @apiDescription        Авторизация пользователя
* @apiGroup              EventsController
* @apiVersion            0.0.1
*/
// func (c AuthController) Login() revel.Result {
    // email    := c.Request.FormValue("email")
    // password := c.Request.FormValue("password")

    // Проверка данных

    // Проверка данных
    // if c.AuthValidator(email, password) {
    //    return c.Redirect("/register/")
    // }
//     return c.RenderText("Its Login action")
// }


/**
* @api {post} /signup/   Register()
* @apiName               Register()
* @apiDescription        Регистрация нового пользователя
* @apiGroup              EventsController
* @apiVersion            0.0.1
*/
func (c AuthController) Register(baseModel models.BaseModel) revel.Result {

    // сбор данных
    name        := c.Request.FormValue("firstname")
    surname     := c.Request.FormValue("lastname")
    email       := c.Request.FormValue("email")
    password    := c.Request.FormValue("password")
    re_password := c.Request.FormValue("re_password")

    if password != re_password{
       fmt.Println("Dont match")
       return c.Redirect("/register/")
    }

    // Проверка данных
    if c.AuthValidator(email, password, name, surname) {
       fmt.Println("--------------------------------------------------------------------------------------")
       fmt.Println(email, password, name, surname)
       fmt.Println("Not valid")
       return c.Redirect("/register/")
    }

    // хеширование пароля
    hashpass := c.Sys_sha(email, password)

    // занос в БД
    data := models.Mst{"name":name, "surname":surname, "email":email, "password":hashpass }
    baseModel.CREATE(c.GENGUID(), "Users", data)

    // установка кук  (BR_DATA = {name:name, surname:surname, Uid:Uid})
    expiration := time.Now().Add(365 * 24 * time.Hour)
    new_cookie := &http.Cookie{Name: "Test", Value: "rutest",  Expires: expiration, Path: "/"}
    c.SetCookie(new_cookie)

    // c.Session["name"] =
    // c.Session["surname"] =  // Error - value needs to be a string
    return c.Redirect("/")
}


/*
 *  Валидация данных пользователя
 */
func (c AuthController) AuthValidator(email, password, name, surname string) bool {

    // сообщения
    emailValidMsg      := fmt.Sprintf(c.Message("emailValid") ,  "email")
    requiredEmailMsg   := fmt.Sprintf(c.Message("required") , "email")
    maxSizeEmailMsg    := fmt.Sprintf(c.Message("maxSize") ,  "email", "100")

    requiredPassword   := fmt.Sprintf(c.Message("required") , "password")
    maxSizePasswordMsg := fmt.Sprintf(c.Message("maxSize") ,  "password", "100")

    requiredNameMsg    := fmt.Sprintf(c.Message("required") , "name")
    maxSizeNameMsg     := fmt.Sprintf(c.Message("maxSize") ,  "name", "100")

    requiredSurnameMsg := fmt.Sprintf(c.Message("required") , "surname")
    maxSizeSurnameMsg  := fmt.Sprintf(c.Message("maxSize") ,  "surname", "100")

    // проверка
    c.Validation.Email(email).Message(emailValidMsg)
    c.Validation.Required(email).Message(requiredEmailMsg)
    c.Validation.MaxSize(email, 40).Message(maxSizeEmailMsg)

    c.Validation.Required(password).Message(requiredPassword)
    c.Validation.MaxSize(password, 40).Message(maxSizePasswordMsg)

    c.Validation.Required(name).Message(requiredNameMsg)
    c.Validation.MaxSize(name, 40).Message(maxSizeNameMsg)

    c.Validation.Required(surname).Message(requiredSurnameMsg)
    c.Validation.MaxSize(surname, 40).Message(maxSizeSurnameMsg)

    // если нашлись ошибки после проверки
	if c.Validation.HasErrors() {
       c.Validation.Keep()           // в переменную errors записуем все ошибки
       c.FlashParams()               // все данные которые пришли формы доступны в Flash
       return true
	} else{
	   return false
	}
}


// func CodeCooke(Uid,Name,Surname,Email string) SecCook {
//      var ck = SecCook{Uid,Name,Surname,Email}
//      return ck
// }



// **************************************************************************************
// Test Cookers
// **************************************************************************************
func (c AuthController) CookeTest() revel.Result {

    // Запись кук 
    c.CookAdd("Allcook", "UID-A00120302", "Gerda", "Lvaovna", "email@email.com")
    c.CookAdd("Nmm",     "UID-A0012000w", "Gerda", "Lvaovna", "email@email.com")


    // Получение и расшифровка кук
    T:=c.CookeRead("Allcook")
    // Вывод данных 
    fmt.Println(T.Surname, T.Email, T.Name, T.Uid)

    // Получение и расшифровка кук
    Z:=c.CookeRead("Nmm")
    
    // Вывод данных 
    fmt.Println(Z.Surname)
    fmt.Println(Z.Email)

    // Удаление кук по имени
    // c.CookDelete("Nmm")  
    

    // Работа с сессиями
     c.SessionAdd("GG","ghh")
     fmt.Println(c.SessionRead("GG"))
     
     return c.RenderText("Ok")
}

// ****************************************************************
// Запись сессии
// c.SessionAdd("GG","ghh")
// ****************************************************************
func (c AuthController) SessionAdd(Name,Value string){
      // key   := []byte(secKey)    
      c.Session["Name"]=Value //encrypt(key, Value)
}


// ****************************************************************
//  Чтение сессии
//  fmt.Println(c.SessionRead("GG"))
// ****************************************************************
func (c AuthController) SessionRead(Name string) string {
      // key   := []byte(secKey)    
      v:=c.Session["Name"]
      // s:=decrypt(key, v)
      return v //s
}

// ******************************************************************************************
// Сохранение кук 
// NameCookies - Имя кук
// Uid         - Код пользователя
// Login       - Логин
// Password    - Пароль
// Email       - Почта
// ******************************************************************************************
func (c AuthController) CookAdd(NameCook, Uid, Login, Pass, Email string) bool {

    if NameCook=="" || Uid=="" || Login =="" || Pass=="" || Email=="" {
       return  false
    }

    l          := c.CookeCode(Uid, Login, Pass, Email)
    expiration := time.Now().Add(365 * 24 * time.Hour)
    new_cookie := &http.Cookie{Name: NameCook, Value: l,  Expires: expiration, Path: "/"}

    // Запись кук
    c.SetCookie(new_cookie)
    return  true
}


// ******************************************************************************************
// Удаление кук по его имени
// NameCookies - Имя кук
// Пример : c.CookDelete("Nmm")  
// ******************************************************************************************
func (c AuthController) CookDelete(NameCook string) bool {

    if NameCook=="" {
       return  false
    }
    
    expiration := time.Now().Add(-1 * time.Hour)
    new_cookie := &http.Cookie{Name: NameCook,  Expires: expiration, Path: "/"}
    // Запись кук
    c.SetCookie(new_cookie)
    return  true
}


// *************************************************************************
//  Шифрование кук
//  И сборка их в одну строку 
//  вида : Uid*Name*Saurname*Email   
// *************************************************************************
func (c AuthController) CookeCode(Uid, Name, Surname, Email string) string {

    key   := []byte(secKey)             // Ключ шифрования - расшифрования
    uid   := encrypt(key, Uid)
    name  := encrypt(key, Name)
    sname := encrypt(key, Surname)
    email := encrypt(key, Email)
    
    // Формирование страки возврата
    pr := uid + sepr + name + sepr + sname + sepr + email
    return pr
}


// *************************************************************************
//  Чтение кук с рашифровкой
// func (c AuthController) CookeRead(Ncook, Text string) SecCook {
// *************************************************************************
func (c AuthController) CookeRead(Ncook string ) SecCook {
    var Ck SecCook
    key  := []byte(secKey)   
    ck, _ := c.Request.Cookie(Ncook)
    L     := string(ck.Value)
     
    uids, name, sname, email := c.CookeDecode(L)

    Ck.Uid     = decrypt(key, uids)
    Ck.Name    = decrypt(key, name)
    Ck.Surname = decrypt(key, sname)
    Ck.Email   = decrypt(key, email)
     
    // fmt.Println(Ck) 
     return Ck
}

     
// **************************************************************************************************
// Возврат четыре параметра в перемнные
// **************************************************************************************************
func (c AuthController) CookeDecode(Cooks string) (string, string, string, string) {
     
     if Cooks== "" {
        return "","","",""
     }

     lg := c.Splt (Cooks, sepr, 0)
     nm := c.Splt (Cooks, sepr, 1)
     sm := c.Splt (Cooks, sepr, 2)
     tm := c.Splt (Cooks, sepr, 3)
     return lg, nm, sm, tm
}


/******************************************************************************************************
 * TITLE  :  Разделение строки
 * REMARK :  Парс по признаку
 ******************************************************************************************************/
func (c AuthController) Splt(Text, Sep string, Nb int64) string {
    if Text=="" || Sep=="" { return "" }
    lg := strings.Split(Text, Sep)[Nb]
    return lg
}


/**************************************************************************************************
    Пример использования шифрования - дешифрования

    key          := []byte()
    key          := []byte("SecKsCsertxfY0d2sdx003-d2s93VA2Y")
    key          := []byte("Gerda_SecretCode")
    originalText := "mes@ms.ua"

    cryptoText := encrypt(key, originalText)        // encrypt value to base64
    text       := decrypt(key, cryptoText)          // encrypt base64 crypto to original value 

    fmt.Println(cryptoText)
    fmt.Printf(text)
**************************************************************************************************/

//**************************************************************************************************
// Encrypt string to base64 crypto using AES
//**************************************************************************************************
func encrypt(key []byte, text string) string {
    // key := []byte(keyText)
    plaintext  := []byte(text)
    block, err := aes.NewCipher(key)
    
    if err != nil {
       panic(err)
     }
    
    // The IV needs to be unique, but not secure. 
    // Therefore it's common to include it at the 
    // beginning of the ciphertext.
    ciphertext := make([]byte, aes.BlockSize+len(plaintext))
    iv         := ciphertext[:aes.BlockSize]
    if _, err  := io.ReadFull(rand.Reader, iv); err != nil {  panic(err)}
    stream     := cipher.NewCFBEncrypter(block, iv)
    
    stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

    // convert to base64
    return base64.URLEncoding.EncodeToString(ciphertext)         
}


//**************************************************************************************************
// decrypt from base64 to decrypted string
//**************************************************************************************************
func decrypt(key []byte, cryptoText string) string {

    ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)
    block, err    := aes.NewCipher(key)
    
    if err!= nil {
       panic(err)
    }
    
    // The IV needs to be unique, but not secure. 
    // Therefore it's common to include it at the 
    // beginning of the ciphertext.
    if len(ciphertext) < aes.BlockSize {
       panic("ciphertext too short")
    }

    iv        := ciphertext[:aes.BlockSize]
    ciphertext = ciphertext[aes.BlockSize:]
    stream    := cipher.NewCFBDecrypter(block, iv)
    
    // XORKeyStream can work in-place if the two 
    // arguments are the same.
    stream.XORKeyStream(ciphertext, ciphertext)
    return fmt.Sprintf("%s", ciphertext)
}
