## Работа с БОТОМ
### Бот отвечает на разные запросы
### Необходимо поменять apiToken на свой



```golang
// Copyright 2015-2016 mrd0ll4r and contributors. All rights reserved.
// Use of this source code is governed by the MIT license, which can be found in
// the LICENSE file.

package main

import (
	"fmt"
	"github.com/mrd0ll4r/tbotapi"
	"github.com/mrd0ll4r/tbotapi/examples/boilerplate"
)

const apiToken = "11111111111:XXXXXXXXXXX"

// *************************************************************************
// Main procedure
// *************************************************************************
func main() {
    test2()	
}

// *************************************************************************
// Пример использования бота
// *************************************************************************
func test2(){

	updateFunc := func(update tbotapi.Update, api *tbotapi.TelegramBotAPI) {

		 	 msg := update.Message
			 typ := msg.Type()
             fmt.Println("TYPE : ", typ)

			 fmt.Printf("<-%d, From:\t%s, Text: %s \n", msg.ID, msg.Chat, *msg.Text)
            
             s:=fmt.Sprintf("%s", *msg.Text)
             v:=api.NewOutgoingMessage

             switch s {
                     case "Привет", "Hi", "hi", "привет", "Здраствуйте":
                          v(tbotapi.NewRecipientFromChat(msg.Chat),  "Здраствуй дорогой друг рад сообщить тебе что я тебя узнал").SetMarkdown(true).Send()     	
                     case "Артур", "Лена", "Никитка", "Никита", "Сынок":     
                          v(tbotapi.NewRecipientFromChat(msg.Chat),  "Никитка приветик я рада тебя видеть").SetMarkdown(true).Send()   
                     case "Что нового":
                           // v(tbotapi.NewRecipientFromChat(msg.Chat),"<h3>Спасибо все ок</h3><hr><p>Вот думаю, что рассказать тебе о Сахалине.</p>").SetHTML(true).Send()       	
                     	   // v(tbotapi.NewRecipientFromChat(msg.Chat),  "Да ниче так вот болтаю с тобой.").SetMarkdown(true).Send()  
                     	   api.NewOutgoingMessage(tbotapi.NewRecipientFromChat(msg.Chat),"<b>Ответ</b> <a href=\"https://google.com\">links</a>").SetHTML(true).Send() 
                     default:
			              v(tbotapi.NewRecipientFromChat(msg.Chat),  "Я тебя незнаю").SetMarkdown(true).Send()     	
            }

	 }

     boilerplate.RunBot(apiToken, updateFunc, "Бот готов.", "Слушаю и отвечаю.")
}



// *************************************************************************
// Тест второй 
// *************************************************************************
func Test(){
	updateFunc := func(update tbotapi.Update, api *tbotapi.TelegramBotAPI) {

		switch update.Type() {

		case tbotapi.MessageUpdate:
		 	 msg := update.Message
			 typ := msg.Type()

			if typ != tbotapi.TextMessage {
				// Ignore non-text messages for now.
				fmt.Println("Ignoring non-text message")
				return
			}
			// Note: Bots cannot receive from channels, at least no text messages. So we don't have to distinguish anything here.

			// Display the incoming message.
			// msg.Chat implements fmt.Stringer, so it'll display nicely.
			// We know it's a text message, so we can safely use the Message.Text pointer.
			fmt.Printf("<-%d, From:\t%s, Text: %s \n", msg.ID, msg.Chat, *msg.Text)

			// Now let's send a markdown-formatted message.
			outMsg, err := api.NewOutgoingMessage(tbotapi.NewRecipientFromChat(msg.Chat),"Сообщение _formatted_ *text* with [links](https://google.com)").SetMarkdown(true).Send()
			if err != nil {
				fmt.Printf("Error sending: %s\n", err)
				return
			}

			fmt.Printf("->%d, To:\t%s, Text: %s\n", outMsg.Message.ID, outMsg.Message.Chat, *outMsg.Message.Text)

			// And now with HTML.
			outMsg, err = api.NewOutgoingMessage(tbotapi.NewRecipientFromChat(msg.Chat),"This is <i>formatted</i> <b>text</b> with <a href=\"https://google.com\">links</a>").SetHTML(true).Send()
			if err != nil {
				fmt.Printf("Error sending: %s\n", err)
				return
			}

			fmt.Printf("->%d, To:\t%s, Text: %s\n", outMsg.Message.ID, outMsg.Message.Chat, *outMsg.Message.Text)

		case tbotapi.InlineQueryUpdate:
			 fmt.Println("Ignoring received inline query: ", update.InlineQuery.Query)

		case tbotapi.ChosenInlineResultUpdate:
			 fmt.Println("Ignoring chosen inline query result (ID): ", update.ChosenInlineResult.ID)

		default:
			fmt.Println("Ignoring unknown Update type.")
		}

	}

	// Run the bot, this will block.
	boilerplate.RunBot(apiToken, updateFunc, "Markup", "Demonstrates markdown and HTML markup")
}
```
