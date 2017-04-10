/******************************************************************************************
 *      
 *		SERVICE 
 *		Copyright (2014-2017)
 *
 *      MODULE           : Vars    
 *
 *		Description      : Service REST API
 *		Version          : Version 4.1.2
 *		Date Started     : 03.11.2014
 *		Date Changed     : 05.02.2016
 *		Author           : Savchenko Arthur
 *		Last Upadte Date : 05.02.2016 11:10
 *
 *		Условные обозначения, сокращения и соглашения о коде :
 *      Термины, соглашения и обозначения, принятые в документации
 *
 ****************************************************************************************/
package main

import (
	"time"                                                      
	r "github.com/dancannon/gorethink"                          
)

//  Gloabal Variables for connection 
var (
	sessionArray []*r.Session
	CurTime              = time.Now().Format("2006-01-02 15:04:05")   // Формат
	CurTimeShort         = time.Now().Format("2006-01-02")            // Формат
	CurTimeUnix          = time.Now().Format(time.RFC3339)            // Дата UNIX  
	CurTimeNano          = time.Now().Format(time.RFC3339Nano)        // Дата UNIX nano
	ActiveIp      string = "localhost"                                // Активный адрес 
	TempStr       string = ":5555"                                    // Активный порт
	Remarks              = "Version testing - 5.027 17.06.2016 10:00"
	Term                 = "TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION"
	Servicestatus        = "Ok"
	Notify               = "Programm alredy started or port 5555 is busy other service."
)

// Declaration inetrfaces & structure & other type
type Mii interface{}                                          // Interface
type Mif []interface{}                                        // Cрез Interface
type Msr []string                                             // Срез String
type Mst map[string]interface{}                               // Map - string - interface
type Mss map[string]string                                    // Map - string - string
type Msi map[string]int64                                     // Map - string - int64
type Mis map[int64]string                                     // Map - int64 - string
type Msl map[int]string                                       // Map - int   - string
type Mil []int64



// System Constants
const (
	AdressService          = "127.0.0.1"                      // Local адресс
	PortService            = ":5555"                          // Local Port
	DatabaseName           = "System"                         // System - Test Databse
	AdressProductionServer = "10.0.50.16:28015"               // IP Server Morion
	AdressMorionServer     = "10.0.50.16"                     // IP Local Server
	AdressExternal         = "195.128.18.66:28015"            // IP World External for Connect
	AddressMainserver      = "localhost:28015"                // Local Host Mashine
	ServerVersion          = "2.0.5"                          // Current Version
	CodeMirrorVer          = "4.8.001"                        // Editor version
	ServerDescript         = "Draft"                          // Draft Version
	TBN                    = "Docmove"                        // Table Name for Documeents
	DBN                    = "test"                           // Database for Tset
	DHO                    = "HO"                             // Database Head Office
	DSYS                   = "System"                         // System Database
	Organization           = "Ionica"                         // Organization Name
	rupor                  = "Finish Work"                    // Check for Test
    ServiceClientKey       = "S0864AA791CE7E7B00R1T$$"        // Secret Client Key By Default  
    ServiceSecretKey       = "A0AEC09A647A688A64A28"          // Secret Service Key By Default 
    DatabaseKey            = "000Morion000"                   // Secret Key for Database
    SecKey                 = "KeySecret$"                     // AccessKey 
)    

// DocHeader - базовая структура шапки документа
type DocHeader struct {
     ID_OWNER_CORP      int64 `json:",string"`                // ИД корпорации, где создан документ (SEQ)
     ID_OWNER_CONTR     int64 `json:",string"`                // ИД организации (сети), где создан документ
     ID_OWNER_STRUCT    int64 `json:",string"`                // ИД точки (а, склад, аптечный пункт), где создан документ
     ID_PARTNER_CORP    int64 `json:",string"`                // ИД корпорации партнера  
     ID_PARTNER_CONTR   int64 `json:",string"`                // ИД организации (сети) партнера
     ID_PARTNER_STRUCT  int64 `json:",string"`                // ИД точки партнера (а, склад, аптечный пункт)
     ID_DOC             int64 `json:",string"`                // ИД документа
     ID_TYPE            int64                                 // ИД типа документа
     NAME_TYPE          string                                // Название типа документа
     DOC_NUM            string                                // Номер документа  
     DOC_DATE_UNX       int64  `json:",string"`               // Дата документа (дата + время) формат UnixNano
     DOC_DATE_STR       string                                // Дата документа (дата + время) формат 2006-01-02T15:04:05.000
     USER_NOTE          string                                // Примечание от пользователя
     HDF_PARENT         int64 `json:",string"`                // ИД документа основания
     HDF_SEQ            int64                                 // Последовательность приращения движения документов в HDF
     HDF_TIME_UNX       int64 `json:",string"`                // Время создания - изменения в реестре формат UnixNano
     HDF_TIME_STR       string                                // Время создания - изменения в реестре формат 2006-01-02T15:04:05.000
     HDF_STATUS         int16                                 // Статус регистрации документа   
}

// Matrix
type Point struct {
	 STATUS            int64                                  // Статус обработки
	 DATETIME          string                                 // Дата обработки
	 STRUCTURE         int64                                  // Ид структуры и
	 ABCC              string                                 // АВС
	 ABCF              string                                 // АВС
	 XYZ               string                                 // XYZ
	 BCQ               string                                 // BCQ
	 Description       string                                 // Описание
	 LIM               int64                                  // Лимит минимального остатка на складе
	 DAYS              int64                                  // Количсетво дней - рекомендуемое
	 Recomendation     string                                 // Рекомендации общего характера
}

//  Документ перемещение - регистрация
type Document struct {
	 ID              string                                   // Ключ
	 ID_COORP        int                                      // Коорпорация
	 ID_BUSINESS     int                                      // Сеть - организация
	 ID_STRUCTURE    int                                      // Точка а
	 ID_COORP_TO     int                                      // Коорпорация        (куда)
	 ID_BUSINESS_TO  int                                      // Сеть - организация (куда)
	 ID_STRUCTURE_TO int                                      // Точка а       (куда)
	 ID_DOC          int                                      // Ид документа
	 DOC_NUM         int                                      // Номер документа
	 DOC_DATE        int                                      // Дата документа (дата + время)
	 ID_TYPE         int                                      // Тип документа  (тип документа)
	 HDF_SEQ         int                                      // Последовательность приращения движения документов в Head Office
	 HDF_TIME        int                                      // Время создания - изменения в реестре
	 HDF_STATUS      int                                      // Статус регистрации документа
}      

// Log
type Logofile        struct {
	 Date            string                                   // Дата операции
	 OperationName   string                                   // Имя Операции - краткое описание операции
	 NumberDoc       string                                   // Номер документа с которым производилась операция
	 TypeOperation   string                                   // Тип операции удаление - добавление - изменение
	 Sataus          string                                   // Саттус выполнения Ok, Err
}

// Log structure
type Logstruct       struct {
	 Datetime        string                                   // Дата операции и время операции
	 Operation       string                                   // Описание операции
	 Document        string                                   // Номер документа 
	 Type            string                                   // Тип операции удаление
	 Status          string                                   // Саттус выполнения
	 Group           string                                   // Расчет, оповещение, временное сообщение
	 Code            string                                   // Код операции Info, Warning, Attetion, Notify,  
	 Createdtime     string                                   // Дата время операции   
	 Createddate     string                                   // Дата время операции   
	 Key             string                                   // Ключ уникальной опреации       
}

// Dictonary
type Directory       struct {
	 Title           string                                   // Описание операции
	 TitleEng        string                                   // Описание операции по английски
}

// BookMark
type Bookmark        struct {
	 Title           string                                   // Описание закладки 
	 Url             string                                   // Адрес 
}

//  Version
type Verstruc       struct {
	 Id             int        `gorethink:"id"`               // Id Key
	 Ver            string     `gorethink:"Ver"`              // Version programs
	 Descript       string     `gorethink:"Description"`      // Description operation 
	 Datevers       string     `gorethink:"Datever"`          // Date last version
}

// Max Sequence
type MaxIntt        struct {
	 Mx             int64                                     // Max number for SEQ
  // Id             string
}

// Коєфициенты для расчета потребности по-умолчанию
type KoefAnalyz    struct {
     Acmin         float64                                    // Колич мин А
     Bcmin         float64                                    // Колич мин B
     Bcmax         float64                                    // Колич мax B
     Ccmax         float64                                    // Колич мин C
     Asmin         float64                                    // Cумм  мин А
     Bsmin         float64                                    // Cумм  мин В
     Bsmax         float64                                    // Cумм  мах B
     Csmin         float64                                    // Cумм  мин C 
     Xmin          float64                                    // Mин       Х 
     Ymin          float64                                    // Mин       Y
     Ymax          float64                                    // Max       Y
     Zmax          float64                                    // Maх       Z
   }  

// Cтруктура корпорации
type Corporation struct {
	 ID                string                                 // Ид в сервисе
	 NAME              string                                 // Сокращенное наименование
	 CODE              string                                 // Код для внешних систем
	 DESCRIPTION       string                                 // Описание
	 FULLNAME          string                                 // Полное имя
	 GROUPID           string                                 // Группа
	 ADRESS_CITY       string                                 // Город
	 ADRESS_AREA       string                                 // Район
	 ADRESS_STREET     string                                 // Улица
	 ADRESS_OFFICE     string                                 // Комната 
	 ADRESS_TEL        string                                 // Контактный телефон
	 ADRESS_MOB        string                                 // Адрес тел
	 CODE_EDRPOU       string                                 // Код ЕДРПОУ
	 DIRECTOR          string                                 // Директор
	 ACCAUNTAN         string                                 // Главный бухгалтер
	 REMARK            string                                 // Пояснения пользователя
	 CREATED           string                                 // Дата создания
	 CHANGED           string                                 // Дата корректировки
	 STATUS            string                                 // [P]-Подготовлен, [A]-Активный, [B]-Блокирован
	 EMAIL             string                                 // Электронный адрес для писем
	 BLOCKDATE         string                                 // Дата блокировки
	 BLOCKREASON       string                                 // Причина блокировки
	 SKYPE             string                                 // Адрес скайпа
	 ADMIN             string                                 // Название логина для всей сети
	 PASS              string                                 // Пароль (зашифрован)
	 FLAG              string                                 // Флаг - системное поле
}


type Corp struct{
	 ID                string                                 // "C105" ,
     ACTIVE            string                                 // Активная "Y" Бокирована "N"          
     Address           string                                 // Адрес
     Street            string                                 // Улица
     Area              string                                 // Область
     City              string                                 // Город
     Code              int64                                  // Код корпорации 578245
     Codekey           string                                 // Код корпорации "C12233"
     Corporation       string                                 // Наименование
     Country           string                                 // Страна 
     Frm               string                                 // Форма властности             
     Status            string                                 // "Old" New ,
     TS                string                                 // Код ТС
     TSNAME            string                                 // Имя Фамилия ТС
     TSVIZIT           string                                 // Время визита ТС
     WTS               string                                 // Код ВТС  
     WTSNAME           string                                 // Имя ВТС    
     WTSVIZIT          string                                 // Дата визита ВТС
     DESCRIPTION       string                                 // Примечание
     Phonework         string                                 // Рабочий телефон
     Phonemob          string                                 // Рабочий мобильный телефон
     Email             string                                 // Email
}


//   Aliance - Coorporation
type Aliance struct {
	 ID                int64                                  // Уникальный код 
	 Id                string                                 // Код в базе
	 CODE              string                                 // Уникальный код
	 GROUPID           string                                 // Группа
     TITLE             string                                 // Расширенное наименование
	 NAME              string                                 // Статус обработки
	 FULLNAME          string                                 // Дата обработки
     ADDRESS           string                                 // Адрес
     CITY              string                                 // Город
     CREATED           string                                 // Дата создания записи в базе 
     DESCRIPTION       string                                 // Описание
     EMAIL             string                                 // Электронная почта 
     SKYPE             string                                 // Скайп
     SEQ               int64                                  // Cчетчик
     STATUS            string                                 // Статус (New, Old, Block, Active) 
     TELEPHONE         string                                 // Телефон
     URL               string                                 // URL  
	 ADMIN             string                                 // Администратор фио
	 PASSWORD          string                                 // Пароль для администратора
	 BLOCKDATE         string                                 // Дата работы после котрой блокируется 
	 CHENGED           string                                 // Дата изменения
	 REMARK            string                                 // Примечание для внутреннего использования
	 NETS              *[]Netss                               // Коды сетей
	 STRUCTURES        *[]Structures                          // Коды 
}

// Сеть 
type Netss struct {
	 ID               string                                  // Primary Key
	 CODE             string                                  // Уникальный код
	 GROUPID          string                                  // Группа
     TITLE            string                                  // Расширенное наименование
	 NAME             string                                  // Статус обработки
	 FULLNAME         string                                  // Дата обработки
	 // STRUCTURES       []Structures                         // Коды 
}

// а
type Structures       struct {
	 ID                string                                 // Уникальный код int64   
	 Id                string                                 // Код в базе
	 ID_STRUCTURE      int64                                  // Ид структуры   
	 IDSYS             int64                                  // Unix time     
	 NET               string                                 // Код Точек
	 CORP              string                                 // Код корпорации
	 CORPID            string                                 // Код корпорации 
	 CORPNAME          string                                 // Наименование корпорации
	 CODE              string                                 // Уникальный код
	 CATEGORY          string                                 // Категория 
	 MAIN              string                                 // Центральня - основная    
	 CREATED           string                                 // Дата создания
	 CHANGED           string                                 // Дата корректировки
	 GROUPID           string                                 // Группа 
     TITLE             string                                 // Расширенное наименование
	 NAME              string                                 // Статус обработки
	 FULLNAME          string                                 // Дата обработки
     ADDRESS	       string                                 // Адрес
     ADRESS_AREA       string                                 // Область
	 ADRESS_CITY       string                                 // Город
	 ADRESS_MOB        string                                 // Мобильный
	 ADRESS_OFFICE     string                                 // Дом
	 ADRESS_STREET     string                                 // Улица
	 ADRESS_TEL        string                                 // Городской телефон
	 ADRESS_DISTRICT   string                                 // Район
     EDRPOU            string                                 // Едрпоу
     DESCRIPTION       string                                 // Описание
     EMAIL             string                                 // Электронная почта 
     SKYPE             string                                 // Скайп
     SEQ               int64                                  // Cчетчик
     STATUS            string                                 // Статус (New, Old, Block, Active) 
     TELEPHONE         string                                 // Телефон
     URL               string                                 // URL  
     ACCOUNTAN         string                                 // Главный бухгалтер Фио
     DIRECTOR          string                                 // Генеральный директор фио
	 ADMIN             string                                 // Администратор фио
	 PASSWORD          string                                 // Пароль для администратора
	 BLOCKDATE         string                                 // Дата работы после котрой блокируется 
	 CHENGED           string                                 // Дата изменения
	 REMARK            string                                 // Примечание для внутреннего использования
	 FLAG              string                                 // Системное поле
	 NETS              []Netss                                // Коды сетей
	 NETNAME           string                                 // Наименование сети 
	 STRUCTURES        string                                 // Коды  Old-  []Structures
     CALCREC           int64                                  // Количество обработанных записей после расчета потребности
     ABCMIN            float64                                // Настройка АВС минимума для анализа
     ABCMAX            float64                                // Настройка АВС маx для анализа
     BCGMIN            float64                                // Настройка ВСG минимума для анализа
     BCGMAX            float64                                // Настройка ВСG маx для анализа
     BCGDISCOUNT       float64                                // Настройка ВСG коєфициент скидки
     EXTREM            int64                                  // Коєфициент всплеска 
     TIMECNT           string                                 // Затраченное время на расчет в секундах (int64 - старый враинат)
     CALCVERS          int64                                  // Версия расчета
     REGION            string                                 // Код региона
     WTS               string                                 // Код Ведущего техничсеского специалиста
     TS                string                                 // Код замещающего техничсеского специалиста 
     TS_START          string                                 // Дата замены с которой начинает работать заместитель
     TS_FINISH         string                                 // Дата по которую будет заместитель работать вместо основного технического специалиста
     TS_NAME           string                                 // Имя ТС
     ACTIVE            string                                 // Кативирована  Y - N
     OPERATION         string                                 // Имя последней операции в системе
     DEEPDAY           int64                                  // Глубина обработки при расчете
     DEEPDAYA          int64                                  // Глубина обработки при расчете для АВС анализа
     COUNTCALC         int64                                  // Количество обработанных записей  
     STARTDATE         string                                 // Дата начала расчета 
     STATUSCALC        string                                 // Статус обработки для вывода в отчет
     STATUSCODE        int64                                  // Статус обработки для отчета         
     LASTDOCDATE       string                                 // Дата последнего обновления при корректировки 
     LASTDOCDATA       string                                 // Дата последнего вход ТС
} 

// Function
func (v *Structures) Codekey() string {return v.DIRECTOR + "  " + v.NAME }

func (v *Structures) Log(){ 
	        r.DB("System").Table("Log").Insert(Mst{"STRUCTURES":v.ID_STRUCTURE, "REMARK":"Добавлена", "CODE":"ST"}).Run(sessionArray[0])
            // var c chan string 
            // RT1000(c, v.DESCRIPTION)
            // RT1000(c, v.DESCRIPTION)
 }


// Добавление в очередь
func Chainadd(C chan string, Text string  ){
        C<-Text
}

// Права пользователя
type Userrigth struct {
     Access           bool                                    // Доступ к системе 
     View             bool                                    // Просмотр   
     Insert           bool                                    // Вставка 
     Delete           bool                                    // Удаление
     Update           bool                                    // Обновление     0-запрет 1-разрешено   
     Admin            bool                                    // Права администратора  
} 

// Пользователи короткий вид
type User struct {
	 Name              string  `gorethink:"name"`             // Полное имя    
	 Password          string  `gorethink:"password"`         // Пароль
}

// Пользователи системы
type USER struct {
     Id                string                 // Уникальный код `gorethink:"Id"`
     id                string
	 ID                int64   `gorethink:"ID"`               // Id Key 
	 Ip                string                                 // Имя комп    
	 Os                string                                 // Операционная система  
	 Name              string  `gorethink:"Name"`             // Полное имя  (Имя Фамилия)  
	 Login             string  `gorethink:"Login"`            // Login
	 Password          string  `gorethink:"Password"`         // Пароль
	 Lname             string                                 // Фамилия 
	 Mname             string                                 // Имя
	 Fname             string                                 // Отчество
	 Telephone         string                                 // Tелефон рабочий
	 Telephonemob      string                                 // Tелефон мобильный
	 Position          string                                 // Должность
	 Role              string                                 // Роль в системе (Админ, Ттех. специалист, ВТС - ведущий тех.сс пециалист, директор, бухгалтер, менеджер)
	 Description       string                                 // Описание
	 Status            string                                 // Cтатус
	 Insert_at         string                                 // Дата вставки
	 Update_at         string                                 // Дата обновления
	 Visit_at          string                                 // Дата последнего посещения
	 Structure         int64                                  // Код структуры
	 Coorporation      int64                                  // Код кооропорации
	 Structcode        string                                 // Код структуры
	 Coorpcode         string                                 // Код кооропорации
	 Corpcode          string                                 // Код кооропорации
	 Flag              int                                    // Признак для системных нужд (обработка временная блокировка)
	 Deadline          string                                 // Дата до которой действует разрешения
	 Email             string                                 // Электронная почта 
	 Right             Userrigth                              // Права и полномочия в системе 
	 Key               string                                 // Kлюч передаваемый из формы  
	 Region_code       string                                 // Код региона
	 Region_name       string                                 // Наименование региона    
	 Wts_code          string                                 // Код регионального техничсекского представителя
	 Structures        []int64                                // Коды закрепленные за техничсеким специалистом
	 Code              string                                 // Личный код Технического специалиста 
	 Skype             string                                 // Skype   


}



// Возврат для сервиса пермещения
 var mr struct {
   	 ID               int64 `json:",string"`                  // Уникальный идентификатор
     HDF_SEQ          int64                                   // Счетчик модификации данных
     HDF_TIME_UNX     int64 `json:",string"`                  // Время модификации (UnixNano)
     HDF_TIME_STR     string                                  // Время модификации (2006-01-02T15:04:05.000)
}

type Mreturn struct {
     ID              int64 `json:",string"`                   // Уникальный идентификатор
     HDF_SEQ         int64                                    // Счетчик модификации данных
     HDF_TIME_UNX    int64 `json:",string"`                   // Время модификации (UnixNano)
     HDF_TIME_STR    string                                   // Время модификации (2006-01-02T15:04:05.000)
     // HDF_DEL       int64                                   // Признак удаления 0-нормальная запись 1 - запись удалена
     HDF_EDITOR      int64 `json:",string"`                   // Код и которая редактировала данные
}	

var AdresPort         string                                  // Port for service   

type Systemfunction struct {
     Title    string                                          // Время модификации (2006-01-02T15:04:05.000)
     Other    string                                          // прочие возможности 
     Ltt      *Systemfunction                                 // Cтруктура систем  
}	

var (
	 Sst_f  = func (intt string) string{return "Title" + intt}
     Sst_n  = func (nm   string) string{return "Name"  + nm + " "}
)

// Свойства системных функций
func (v *Systemfunction) Add(myname    string) string {return myname + " Calculation"}
func (v *Systemfunction) Delete(myname string) string {return myname + " Deleted"}
func (v *Systemfunction) Count(myname  int64)  int64  {return myname + 24*60*60}

// Свойства лог файла
// func (v *Logstruct) DateTime() string {return  time.Now().Format(Ymdtimer)} 
// func (v *USER) Codeskode(Mname,Lname string) string {return Mname+ " " + Lname}

// Структура  импорта данных получаемых от точек
// которая инициализирует как точки так и корпорацию
// Пример JSON

/*

{
   "CODE":        "",
   "NAME":        "Приватне підприємство "Пана"",
   "DESCRIPTION": "",
   "BUSINESS": [ {
         "ID_BUSINESS": 5012249,
         "NAME_BUSINESS": "Приватне підприємство \"Пана\"" ,
         "ADDRESS":  "Просп. Леніна, 122, м.Харків, 11111",
      "STRUCTURES":[
            {"ID_STRUCTURE":5012250,"NAME":"Точка","NAME_EXT":"","ADDRESS":"м.Харків, просп. Леніна, 122" },
            {"ID_STRUCTURE":7231103,"NAME":"Основная","NAME_EXT":"AP1","ADDRESS":"м.Харків, просп. Леніна, 122" },
            {"ID_STRUCTURE":7231104,"NAME":"Точка 3","NAME_EXT":"AP3","ADDRESS":"м.Харків, вул. М. Жукова, 11" },
            {"ID_STRUCTURE":49274747,"NAME":"Point 1 Ленина 22","NAME_EXT":"AL1","ADDRESS":"м.Харків, просп. Леніна, 212" }
        ],
       "REMARKS": "Примечание"
     }]
}

*/
 
// Описание структуры инициализация организации
// Дополнение информации обновление информации
type IBuss struct {
	 ID_BUSINESS        int64
	 NAME_BUSINESS      string
	 ADDRESS            string
	 REMARKS            string  
	 STRUCTURES         []Istruct
}

type Istruct struct {
	ID_STRUCTURE        int64 
	NAME                string 
	NAME_EXT            string 
    ADDRESS             string	
}

type ICorp struct {
    CODE                string
    NAME                string
    DESCRIPTION         string  
    BUSINESS            []IBuss
}
