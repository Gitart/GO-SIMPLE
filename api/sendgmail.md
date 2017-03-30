# Send email with attachment(RFC2822) using Gmail API example



This is a simple tutorial on how to send email using Gmail API in Golang since the examples found in   
https://developers.google.com/gmail/api/guides/sending#sending_messages are in Java and Python.  
IMPORTANT! Please do further test by sending email to non-Gmail account.    
So far, I've no problem sending to Gmail account. However, some non-Gmail account such as ZOHO might 
not be able receive the email send by this example program.

Before you start, please turn on the Gmail API and download the credential file - client_secret.json 
by following the Step 1: Turn on the Gmail API section found in https://developers.google.com/gmail/api/quickstart/go#prerequisites

Once you've download the credential file, move it to the same directory as the source code below before running the program.

This code below demonstrate how to send email with attachment in Golang using Gmail API. The portion that 
handles the attachment is:

```golang
 func createMessageWithAttachment(from string, to string, subject string, content string, fileDir string, fileName string) gmail.Message {

         var message gmail.Message

         // read file for attachment purpose
         // ported from https://developers.google.com/gmail/api/sendEmail.py

         fileBytes, err := ioutil.ReadFile(fileDir + fileName)
         if err != nil {
                 log.Fatalf("Unable to read file for attachment: %v", err)
         }

         fileMIMEType := http.DetectContentType(fileBytes)

         // https://www.socketloop.com/tutorials/golang-encode-image-to-base64-example
         fileData := base64.StdEncoding.EncodeToString(fileBytes)

         boundary := randStr(32, "alphanum")

         messageBody := []byte("Content-Type: multipart/mixed; boundary=" + boundary + " \n" +
                 "MIME-Version: 1.0\n" +
                 "to: " + to + "\n" +
                 "from: " + from + "\n" +
                 "subject: " + subject + "\n\n" +

                 "--" + boundary + "\n" +
                 "Content-Type: text/plain; charset=" + string('"') + "UTF-8" + string('"') + "\n" +
                 "MIME-Version: 1.0\n" +
                 "Content-Transfer-Encoding: 7bit\n\n" +
                 content + "\n\n" +
                 "--" + boundary + "\n" +

                 "Content-Type: " + fileMIMEType + "; name=" + string('"') + fileName + string('"') + " \n" +
                 "MIME-Version: 1.0\n" +
                 "Content-Transfer-Encoding: base64\n" +
                 "Content-Disposition: attachment; filename=" + string('"') + fileName + string('"') + " \n\n" +
                 chunkSplit(fileData, 76, "\n") +
                 "--" + boundary + "--")

         // see https://godoc.org/google.golang.org/api/gmail/v1#Message on .Raw
         // use URLEncoding here !! StdEncoding will be rejected by Google API

         message.Raw = base64.URLEncoding.EncodeToString(messageBody)

         return message
 }
 ```
 
and below is the full program source code.

NOTE:
You will need to change the email send and recipients address. Also change the img.pdf file to be attached to 
something else that you have.

When prompted for the authorization code for the first time, cut-n-paste the URL in your browser, then you will get 
a string(token), cut-n-paste that string into your terminal where you execute the program.


```golang
package main

 import (
         "crypto/rand"
         "encoding/base64"
         "encoding/json"
         "fmt"
         "golang.org/x/net/context"
         "golang.org/x/oauth2"
         "golang.org/x/oauth2/google"
         "google.golang.org/api/gmail/v1"
         "io/ioutil"
         "log"
         "net/http"
         "net/url"
         "os"
         "os/user"
         "path/filepath"
 )

 // NOTE : we don't want to visit CSRF URL to get the authorization code
 // and paste into the terminal each time we want to send an email
 // therefore we will retrieve a token for our client, save the token into a file
 // you will be prompted to visit a link in your browser for authorization code only ONCE
 // and subsequent execution of the program will not prompt you for authorization code again
 // until the token expires.

 // getClient uses a Context and Config to retrieve a Token
 // then generate a Client. It returns the generated Client.
 func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
         cacheFile, err := tokenCacheFile()
         if err != nil {
                 log.Fatalf("Unable to get path to cached credential file. %v", err)
         }
         tok, err := tokenFromFile(cacheFile)
         if err != nil {
                 tok = getTokenFromWeb(config)
                 saveToken(cacheFile, tok)
         }
         return config.Client(ctx, tok)
 }

 // getTokenFromWeb uses Config to request a Token.
 // It returns the retrieved Token.
 func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
         authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
         fmt.Printf("Go to the following link in your browser then type the "+
                 "authorization code: \n%v\n", authURL)

         var code string
         if _, err := fmt.Scan(&code); err != nil {
                 log.Fatalf("Unable to read authorization code %v", err)
         }

         tok, err := config.Exchange(oauth2.NoContext, code)
         if err != nil {
                 log.Fatalf("Unable to retrieve token from web %v", err)
         }
         return tok
 }

 // tokenCacheFile generates credential file path/filename.
 // It returns the generated credential path/filename.
 func tokenCacheFile() (string, error) {
         usr, err := user.Current()
         if err != nil {
                 return "", err
         }
         tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
         os.MkdirAll(tokenCacheDir, 0700)
         return filepath.Join(tokenCacheDir,
                 url.QueryEscape("gmail-go-sendemail.json")), err
 }

 // tokenFromFile retrieves a Token from a given file path.
 // It returns the retrieved Token and any read error encountered.
 func tokenFromFile(file string) (*oauth2.Token, error) {
         f, err := os.Open(file)
         if err != nil {
                 return nil, err
         }
         t := &oauth2.Token{}
         err = json.NewDecoder(f).Decode(t)
         defer f.Close()
         return t, err
 }

 // saveToken uses a file path to create a file and store the
 // token in it.
 func saveToken(file string, token *oauth2.Token) {
         fmt.Printf("Saving credential file to: %s\n", file)
         f, err := os.Create(file)
         if err != nil {
                 log.Fatalf("Unable to cache oauth token: %v", err)
         }
         defer f.Close()
         json.NewEncoder(f).Encode(token)
 }

 func sendMessage(service *gmail.Service, userID string, message gmail.Message) {
         _, err := service.Users.Messages.Send(userID, &message).Do()
         if err != nil {
                 log.Fatalf("Unable to send message: %v", err)
         } else {
                 log.Println("Email message sent!")
         }

 }

 func createMessage(from string, to string, subject string, content string) gmail.Message {

         var message gmail.Message

         messageBody := []byte("From: " + from + "\r\n" +
                 "To: " + to + "\r\n" +
                 "Subject: " + subject + "\r\n\r\n" +
                 content)

         // see https://godoc.org/google.golang.org/api/gmail/v1#Message on .Raw
         message.Raw = base64.StdEncoding.EncodeToString(messageBody)

         return message
 }

 func chunkSplit(body string, limit int, end string) string {

         var charSlice []rune

         // push characters to slice
         for _, char := range body {
                 charSlice = append(charSlice, char)
         }

         var result string = ""

         for len(charSlice) >= 1 {
                 // convert slice/array back to string
                 // but insert end at specified limit

                 result = result + string(charSlice[:limit]) + end

                 // discard the elements that were copied over to result
                 charSlice = charSlice[limit:]

                 // change the limit
                 // to cater for the last few words in
                 //
                 if len(charSlice) < limit {
                         limit = len(charSlice)
                 }

         }

         return result

 }

 func randStr(strSize int, randType string) string {

         var dictionary string

         if randType == "alphanum" {
                 dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
         }

         if randType == "alpha" {
                 dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
         }

         if randType == "number" {
                 dictionary = "0123456789"
         }

         var bytes = make([]byte, strSize)
         rand.Read(bytes)
         for k, v := range bytes {
                 bytes[k] = dictionary[v%byte(len(dictionary))]
         }
         return string(bytes)
 }

 func createMessageWithAttachment(from string, to string, subject string, content string, fileDir string, fileName string) gmail.Message {

         var message gmail.Message

         // read file for attachment purpose
         // ported from https://developers.google.com/gmail/api/sendEmail.py

         fileBytes, err := ioutil.ReadFile(fileDir + fileName)
         if err != nil {
                 log.Fatalf("Unable to read file for attachment: %v", err)
         }

         fileMIMEType := http.DetectContentType(fileBytes)

         // https://www.socketloop.com/tutorials/golang-encode-image-to-base64-example
         fileData := base64.StdEncoding.EncodeToString(fileBytes)

         boundary := randStr(32, "alphanum")

         messageBody := []byte("Content-Type: multipart/mixed; boundary=" + boundary + " \n" +
                 "MIME-Version: 1.0\n" +
                 "to: " + to + "\n" +
                 "from: " + from + "\n" +
                 "subject: " + subject + "\n\n" +

                 "--" + boundary + "\n" +
                 "Content-Type: text/plain; charset=" + string('"') + "UTF-8" + string('"') + "\n" +
                 "MIME-Version: 1.0\n" +
                 "Content-Transfer-Encoding: 7bit\n\n" +
                 content + "\n\n" +
                 "--" + boundary + "\n" +

                 "Content-Type: " + fileMIMEType + "; name=" + string('"') + fileName + string('"') + " \n" +
                 "MIME-Version: 1.0\n" +
                 "Content-Transfer-Encoding: base64\n" +
                 "Content-Disposition: attachment; filename=" + string('"') + fileName + string('"') + " \n\n" +
                 chunkSplit(fileData, 76, "\n") +
                 "--" + boundary + "--")

         // see https://godoc.org/google.golang.org/api/gmail/v1#Message on .Raw
         // use URLEncoding here !! StdEncoding will be rejected by Google API

         message.Raw = base64.URLEncoding.EncodeToString(messageBody)

         return message
 }

 func main() {
         ctx := context.Background()

         // process the credential file
         credential, err := ioutil.ReadFile("client_secret.json")
         if err != nil {
                 log.Fatalf("Unable to read client secret file: %v", err)
         }

         // Use GmailSendScope for this example.
         // See the rest at https://godoc.org/google.golang.org/api/gmail/v1#pkg-constants

         config, err := google.ConfigFromJSON(credential, gmail.GmailSendScope)
         if err != nil {
                 log.Fatalf("Unable to parse client secret file to config: %v", err)
         }

         client := getClient(ctx, config)

         // initiate a new gmail client service
         gmailClientService, err := gmail.New(client)
         if err != nil {
                 log.Fatalf("Unable to initiate new gmail client: %v", err)
         }

         // create message without attachment
         msgContent := `Hello!
                        This is a test email send via Gmail API
                        Good Bye!`

         //message := createMessage("from@gmail.com", "to@gmail.com", "Email from GMail API", msgContent)

         // send out our message
         //user := "me"
         //sendMessage(gmailClientService, user, message)

         messageWithAttachment := createMessageWithAttachment("from@gmail.com", "to@gmail.com, to2@gmail.com", "Email WITH ATTACHMENT from GMail API", msgContent, "./", "img.pdf")

         // send out our message
         user := "me"
         sendMessage(gmailClientService, user, messageWithAttachment)

 }
 ```
 
Output:

```
Go to the following link in your browser then type the authorization code:
https://accounts.google.com/o/oauth2/auth?accesstype=offline&clientid=[redacted].apps.googleusercontent.com&redirecturi=urn%3Aietf%3Awg%3Aoauth%3A2.0%3Aoob&responsetype= code&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fgmail.send&state=state-token
4/VruYyZ8uBbf--[redacted]-----M64Eq9DT3rQnX-Zw
Saving credential file to: /Users/[redacted]/.credentials/gmail-go-sendemail.json
2016/07/05 15:01:33 Email message sent!
If everything else goes smoothly, you should receive an email with attachment. You can view the email original form in Gmail by clicking the Show Original option in the drop down button.

A sample email with attachment original data
MIME-Version: 1.0
from: ###@gmail.com
Date: Tue, 5 Jul 2016 03:01:33 -0400
Message-ID: CAPyVj_q_j9Y4KoySoaLuxNni-Emn-6YnuRLMg=1jy3b0MHM2qA@mail.gmail.com
Subject: Email WITH ATTACHMENT from GMail API
To: ####@gmail.com
Content-Type: multipart/mixed; boundary=001a113dea20efb69a0536de0574
--001a113dea20efb69a0536de0574
Content-Type: text/plain; charset=UTF-8
Hello!
This is a test email send via Gmail API
Good Bye!
--001a113dea20efb69a0536de0574
Content-Type: application/pdf; name="img.pdf"
Content-Disposition: attachment; filename="img.pdf"
Content-Transfer-Encoding: base64
X-Attachment-Id: 9b113a0233cf96f7_0.1
JVBERi0xLjMKJcTl8uXrp/Og0MTGCjQgMCBvYmoKPDwgL0xlbmd0aCA1IDAgUiAvRmlsdGVyIC9G
bGF0ZURlY29kZSA+PgpzdHJlYW0KeAErVAhUKFQwAEJTS1MFCxMjhaJUhXCFPAX9gNSi5NSCktLE
HIWiTKAaYwOQKgMwbWhirmeqYGRuqJCcq6DvmWuo4JLPFagQCAADyxMlCmVuZHN0cmVhbQplbmRv
```
