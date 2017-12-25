## Send to mail
### based on http://code.google.com/p/go-wiki/wiki/SendingMail
//run: "godoc net/smtp" see more detail

```golang
package main

import (
        "log"
        "net/smtp"
)



 // Отправка почты в режиме 
 func main() {


    addr     := "smtp.gmail.com"
    sender   := "post@gmail.com"
    password := "mypasswordsecret"
    to       := "post@meta.ua"
    msg      := "subject:Сообщение автоматическое от автоматического рассылочного робота \n\n Send.\n Новости от портала \n"
    auth     := smtp.PlainAuth("go agent",  sender, password, addr)
    err      := smtp.SendMail(addr+":25", auth, sender, []string{to}, []byte(msg))

    // Error
    if err != nil {
            log.Fatal(err)
    }
 }
 ```
 
