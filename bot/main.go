package main

import (
	     "flag"
	     "github.com/Syfaro/telegram-bot-api"
	     "log"
	     "os"
	     // "time"
	     "io/ioutil"
	     "net/http"
	     // "fmt"
	     "encoding/json"
	     // "bytes"
)


type Mst map[string]interface{}

var (
	// глобальная переменная в которой храним токен
	telegramBotToken string
)


// *******************************************************************
// Инициализация подключения к боту
// *******************************************************************
func init() {
	// принимаем на входе флаг -telegrambottoken
	flag.StringVar(&telegramBotToken, "5yyyy:xxx", "5yyy:xxx04", "Telegram Bot Token")
	flag.Parse()

	// без него не запускаемся
	if telegramBotToken == "" {
		log.Print("-telegrambottoken is required")
		os.Exit(1)
	}
}


// *******************************************************************
// 
// Основная процедура
// 
// *******************************************************************
func main() {
	// используя токен создаем новый инстанс бота
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)

	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// u - структура с конфигом для получения апдейтов
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	zapros := ""
	// используя конфиг u создаем канал в который будут прилетать новые сообщения

	updates, err := bot.GetUpdatesChan(u)

	// в канал updates прилетают структуры типа Update
	// вычитываем их и обрабатываем
	for update := range updates {

		zapros = update.Message.Text
		reply := "Введите номер заявки :"

		reply = api_jira_get_task(zapros)

		// создаем ответное сообщение
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		// отправляем
		bot.Send(msg)
		reply = "Какую заявку еще хотите просмотреть ?"

		// // универсальный ответ на любое сообщение
		// reply := "Не знаю что сказать"
		// if update.Message == nil {
		// 	continue
		// }

		// логируем от кого какое сообщение пришло
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// свитч на обработку комманд комманда - сообщение, начинающееся с "/"
		switch update.Message.Command() {
		case "start":		reply = "Добрый день! Я Unity-Bars бот. Я знаю все о Ваших заявках. Введите номер заявки в Jire. Например: COBUMMFO-5112 или COBUMMFO-5113."
		case "hello":		reply = "world"
		case "аccount":		reply = "26220"
		case "rls":  		reply = "Текущий релиз № 68 В него вошли 45678-34568 изменения"
		case "dev":			reply = "Текущий разработчик : Сидоренко Олег Иванович"
		}

		// создаем ответное сообщение
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		// отправляем
		bot.Send(msg)
	}
}


// **************************************************************
//   https://skarlso.github.io/2015/11/20/go-jira-api-client
// **************************************************************/
func api_jira_get_task(Num string) string {

	//    var r  *http.Request
	Login   := "loginname"
	Pass    := "123"
	Return  := ""
    Descr   := ""
    Descrru := ""
    

	// r.SetBasicAuth(Suser, Spass)
	// r.Header.Set("Content-Type", "application/json")
    // url := `http://jira.com:12000/rest/api/2/search?jql=key%20%3D"`+ Num + `"&maxResults=1`
	url := "http://jira.com:12000/rest/api/2/issue/" + Num

	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(Login, Pass)

	res, err := http.DefaultClient.Do(req)
	Err(err, "Error detect default client.")
	defer res.Body.Close()
	
	body, err := ioutil.ReadAll(res.Body)
	Err(err, "Error read body.")

	var data Mst
	errj := json.Unmarshal([]byte(body), &data)
	Err(errj, "Error read body.")

	tt, errtt := data["fields"].(interface{})

	if !errtt {
		Return = "Извините уважаемый коллега но такой заявки я не помню... Пожалуйста введите номер другой заявки."
		return Return
	}
    
    // Формирование заявки 
	Dtt       := tt.(map[string]interface{})
	// Tids   := Dtt["fields"].(map[string]interface{})
	Tid       := Dtt["summary"].(string)                       // Описание краткое

	if Dtt["customfield_15213"] !=nil{
	   Descr = Dtt["customfield_15213"].(string)               // Описание полное
    }else{
       Descr="Релиз не опредлен."
    }

	if Dtt["customfield_15214"] !=nil{
	   Descrru = Dtt["customfield_15214"].(string)             // Описание полное
    }else{
       Descrru ="Релиз РУ не опредлен"
    }

	Status    := Dtt["status"].(map[string]interface{})        // Статус 
	Statusnm  := Status["name"].(string)                      
   

    // Добавление
    // lnk:="http://jira.com:12000/browse/"+Num
    // fmt.Println(Dtt)
    // Формирование ответа на заявку
	Return     = "   Заявка              : " + Num + 
	             "\n Статус              : " + Statusnm + 
	             "\n Релиз ММФО    : " + Descr +
              	 "\n Релиз РУ          : " + Descrru +
	             "\n\nОписание :\n" + Tid +
	             "\n Info #UNITY-BARS"

	return Return

	// fmt.Println("response Status:",  res.Status)
	// fmt.Println("response Headers:", res.Header)
	// fmt.Println("response Body:",    string(body))

	// var data Mst
	// errj:=json.Unmarshal(body, &data)
	//  if errj != nil {
	//     fmt.Println(errj.Error())
	// }

	// fmt.Printf("Results: %v\n", data)
	// fmt.Fprint(w, data)
	// os.Exit(0)
}



/***************************************************************
  Author      Savchenko Arthur
  Company     Name company
  Description Error
  Time        11-12-2017
  Title
 ****************************************************************/
func Err(Er error, Txt string) {
	if Er != nil {
 	   log.Println("ERROR : "+Txt, "Description : "+Er.Error())
	   return
	}
}



//**************************************************************************
// Старая копия
// для тестирования
//**************************************************************************
func mains() {
	// используя токен создаем новый инстанс бота
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)

	if err != nil {
	   log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// u - структура с конфигом для получения апдейтов
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// используя конфиг u создаем канал в который будут прилетать новые сообщения
	updates, err := bot.GetUpdatesChan(u)

	// в канал updates прилетают структуры типа Update
	// вычитываем их и обрабатываем
	for update := range updates {

		reply := "Привет как зовут"
		// log.Println(update.Message.Text)

		switch update.Message.Text {
		case "Артур":			reply = "Привет  друг."
		case "Валерий":			reply = "Добрый день Валерий! Рад тебя видеть."
		case "Анна":			reply = "Я заню тебя ты возглавляешь отдел разработки."
		case "Роман":			reply = "Я заню тебя ты администратор на Барсе."
		case "Рома":	   	    reply = "Я уже занкомился с тобой !! Ты администратор на Барсе. Рад тебя видеть снова."
		case "Релиз":			reply = "На сегодняшнюю дату "
		case "Счет":			reply = "Ваш счет в банке "
		case "Остаток":			reply = "На вашем счету : 1000 грн. и $200"
		case "Дата":			reply = "17.01.2018"
		case "11111":			reply = "Заявка в работе и ждет ответа от разработчика"
		case "00000":			reply = "Заявка не обработана и ждет утверждение от банка"
		case "22222":			reply = "Ожидание утверждения от отдела администрации"
		case "Ярослав":			reply = "Привет Ярик ! Рад видеть службу поддержки у себя в гостях."
		case "Ярик":			reply = "Привет Ярик ! Повторно рад тебя видеть!! Привет службе поддержки!"
		case "Евгений":			reply = "Привет Женя ! Рад видеть руководителя поддержки у себя в гостях!"
		case "Андрей":			reply = "Привет Андрей ! Фамилия твоя не Савченко случайно ?? Знаю я одного такого "
		case "Test":			reply = `Документы ."`
		case "Оля": 			reply = `Привет Оля ! Я тебя знаю! Ты на Барсе директор !!!! `
		case "cc12":	reply = api_jira_get_task("ccc5112")
		}

		// создаем ответное сообщение
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		// отправляем
		bot.Send(msg)
		reply = "Что еще пожелаете узнать ?"

		// // универсальный ответ на любое сообщение
		// reply := "Не знаю что сказать"
		// if update.Message == nil {
		// 	continue
		// }

		// логируем от кого какое сообщение пришло
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// свитч на обработку комманд комманда - сообщение, начинающееся с "/"
		switch update.Message.Command() {
		case "start":		reply = "Привет. Я  бот. Кака тебя зовут?"
		case "hello":		reply = "world"
		case "аccount":		reply = "26220"
		case "rls":			reply = "Текущий релиз № 68 В него вошли 45678-34568 изменения"
		case "dev":			reply = "Текущий разработчик : Сидоренко Олег Иванович"
		}

		// создаем ответное сообщение
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		// отправляем
		bot.Send(msg)
	}
}
