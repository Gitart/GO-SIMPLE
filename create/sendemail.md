# Send email with attachment
golang smtp email-attachment 

In my previous tutorial on Go's send email and configure SMTP example, I left out a part where any decent email package should have - that is the - email attachment part.
So, how to attach files onto email in Golang ?
After trying to write a decent tutorial to demonstrate this...only then I realized that the codes can be super messy and unreadable. Therefore, I would rather settle for this simple to use third party package https://github.com/scorredoira/email to get this tutorial published. (I'm lazy)

Here it is :

```go
 package main

 import (
         "github.com/scorredoira/email"
         "fmt"
         "net/smtp"
         "strconv"
 )

 type EmailConfig struct {
         Username string
         Password string
         Host     string
         Port     int
 }

 func main() {
         // authentication configuration
         smtpHost := "smtp.***.com" // change to your SMTP provider address
         smtpPort := *** // change to your SMPT provider port number
         smtpPass := "****"      // change here
         smtpUser := "******" // change here

         emailConf := &EmailConfig{smtpUser, smtpPass, smtpHost, smtpPort}

         emailauth := smtp.PlainAuth("", emailConf.Username, emailConf.Password, emailConf.Host)

         sender := "******@mail.com" // change here

         receivers := []string{
                 "*****@mail.com",
                 "****@mail.com",
         } // change here

         message := "Please see the email attachment for the images"
         subject := "Attached my photos!"

         emailContent := email.NewMessage(subject, message)

         emailContent.From = sender
         emailContent.To = receivers

         files := []string{
                 "big.jpg",
                 "small.jpg",
         } // change here to your own files

         for _, filename := range files {
                 err := emailContent.Attach(filename)

                 if err != nil {
                         fmt.Println(err)
                 }
         }

         // send out the email
         err := email.Send(smtpHost+":"+strconv.Itoa(emailConf.Port), //convert port number from int to string
                 emailauth,
                 emailContent)

         if err != nil {
                 fmt.Println(err)
         }
 }
 ```
 
