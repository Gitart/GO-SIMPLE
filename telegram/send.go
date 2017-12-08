
/*******************************************************************************************************************
  
  Посылка сообщения в телеграм
  Bot_send("Начало процесса")

  https://github.com/LibreLabUCM/teleg-api-bot/wiki/Getting-started-with-the-Telegram-Bot-API
  https://www.shellhacks.com/ru/telegram-api-send-message-personal-notification-bot/
********************************************************************************************************************/
func Bot_send(Text string){
  	action      := "https://api.telegram.org/bot358911111:AAAA....ZZZZ/sendMessage?chat_id=235188412&text="+Text
    contentType := "Text"
    http.Post(action, contentType, nil) 
}
