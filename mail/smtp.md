
## Go email


Go email tutorial shows how to send emails in Golang with smtp package. In our examples, we use the Mailtrap service.
SMTP

The Simple Mail Transfer Protocol (SMTP) is an internet standard communication protocol for electronic mail transmission. Mail servers and clients use SMTP to send and receive mail messages.
Go smtp

The smtp package implements the Simple Mail Transfer Protocol. It also supports additional extensions.

Note: Gmail is not ideal for testing applications. We should use an online service such as Mailtrap or Mailgun, or use an SMTP server provided by a webhosting company.
The SendMail function

The SendMail function is a high-level function to send emails.

func SendMail(addr string, a Auth, from string, to []string, msg []byte) error

It connects to the server at addr, switches to TLS if possible, authenticates with the optional mechanism a if possible, and then sends an email from address from, to addresses to, with message msg.

The msg parameter should be an RFC 822-style email; such email starts with headers, a blank line, and then the message body. The lines of msg should be terminated with CRLF characters.
Go email simple example

The following is a simple email example.
simple.go

package main

import (
    "fmt"
    "log"
    "net/smtp"
)

func main() {

    from := "john.doe@example.com"

    user := "9c1d45eaf7af5b"
    password := "ad62926fa75d0f"

    to := []string{
        "roger.roe@example.com",
    }

    addr := "smtp.mailtrap.io:2525"
    host := "smtp.mailtrap.io"

    msg := []byte("From: john.doe@example.com\r\n" +
        "To: roger.roe@example.com\r\n" +
        "Subject: Test mail\r\n\r\n" +
        "Email body\r\n")

    auth := smtp.PlainAuth("", user, password, host)

    err := smtp.SendMail(addr, auth, from, to, msg)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Email sent successfully")
}

We send a simple email to the Mailtrap service.

import (
    "fmt"
    "log"
    "net/smtp"
)

We import the net/smtp package.

from := "john.doe@example.com"

This is the email sender.

user := "9c1d45eaf7af5b"
password := "ad62926fa75d0f"

We ge the username and password from the Mailtrap account.

to := []string{
    "roger.roe@example.com",
}

We store the recipients in the to slice.

addr := "smtp.mailtrap.io:2525"
host := "smtp.mailtrap.io"

The address is the host name and the port. Mailtrap listens on port 2525.

msg := []byte("From: john.doe@example.com\r\n" +
    "To: roger.roe@example.com\r\n" +
    "Subject: Test mail\r\n\r\n" +
    "Email body\r\n")

We build the email message. The message lines are separated with CRLF characters.

auth := smtp.PlainAuth("", user, password, host)

The PlainAuth function begins an authentication with a server; it returns an authentication object that implements the plain authentication mechanism. It will only send the credentials if the connection is using TLS or is connected to localhost.

err := smtp.SendMail(addr, auth, from, to, msg)

The email message is sent with the SendMail function. WE pass the function the address, the authentication object, the sender, the recipients and the message.
Go smtp HTML message

The following example sends an email with a body message in HTML.
send_html.go

package main

import (
    "fmt"
    "log"
    "net/smtp"
    "strings"
)

type Mail struct {
    Sender  string
    To      []string
    Subject string
    Body    string
}

func main() {

    sender := "john.doe@example.com"

    to := []string{
        "roger.roe@example.com",
    }

    user := "9c1d45eaf7af5b"
    password := "ad62926fa75d0f"

    subject := "Simple HTML mail"
    body := `<p>An old <b>falcon</b> in the sky.</p>`

    request := Mail{
        Sender:  sender,
        To:      to,
        Subject: subject,
        Body:    body,
    }

    addr := "smtp.mailtrap.io:2525"
    host := "smtp.mailtrap.io"

    msg := BuildMessage(request)
    auth := smtp.PlainAuth("", user, password, host)
    err := smtp.SendMail(addr, auth, sender, to, []byte(msg))

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Email sent successfully")
}

func BuildMessage(mail Mail) string {
    msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
    msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
    msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
    msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
    msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

    return msg
}

In the message body, we use HTML tags. The content type of the message is set to text/html.
Go email carbon copy/blind carbon copy

Carbon copy (CC) recipients are visible to all other recipients while Blind Carbon Copy (BCC) recipients are not visible to anyone. CC recipiens are included in the to parameter and the CC msg field. Sending BCC messages is accomplished by including an email address in the to parameter but not including it in the msg headers.
cc_bc.go

package main

import (
    "fmt"
    "log"
    "net/smtp"
    "strings"
)

type Mail struct {
    Sender  string
    To      []string
    Cc      []string
    Bcc     []string
    Subject string
    Body    string
}

func main() {

    sender := "john.doe@example.com"

    to := []string{
        "roger.roe@example.com",
        "adam.smith@example.com",
        "thomas.wayne@example.com",
        "oliver.holmes@example.com",
    }

    cc := []string{
        "adam.smith@example.com",
        "thomas.wayne@example.com",
    }

    // not used
    bcc := []string{
        "oliver.holmes@example.com",
    }

    user := "9c1d45eaf7af5b"
    password := "ad62926fa75d0f"

    subject := "simple testing mail"
    body := "email body message"

    request := Mail{
        Sender:  sender,
        To:      to,
        Cc:      cc,
        Subject: subject,
        Body:    body,
    }

    addr := "smtp.mailtrap.io:2525"
    host := "smtp.mailtrap.io"

    msg := BuildMessage(request)
    auth := smtp.PlainAuth("", user, password, host)

    err := smtp.SendMail(addr, auth, sender, to, []byte(msg))

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Emails sent successfully")
}

func BuildMessage(mail Mail) string {

    msg := ""
    msg += fmt.Sprintf("From: %s\r\n", mail.Sender)

    if len(mail.To) > 0 {
        msg += fmt.Sprintf("To: %s\r\n", mail.To[0])
    }

    if len(mail.Cc) > 0 {
        msg += fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Cc, ";"))
    }

    msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
    msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

    return msg
}

In the examle, we send an emails to multiple recipiens. Some of them are included CC and BCC recipiens.

to := []string{
    "roger.roe@example.com",
    "adam.smith@example.com",
    "thomas.wayne@example.com",
    "oliver.holmes@example.com",
}

The email is sent to all these emails.

cc := []string{
    "adam.smith@example.com",
    "thomas.wayne@example.com",
}

These two emails will be carbon copied; that is, their email addresses will be visible to anyone.

if len(mail.To) > 0 {
    msg += fmt.Sprintf("To: %s\r\n", mail.To[0])
}

The first email address is displayed in the To field.

if len(mail.Cc) > 0 {
    msg += fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Cc, ";"))
}

Here we build the Cc message header field. Bcc emails are not included in the message headers; therefore, they are not visible to others.
Go email attachment

In the next example, we send an attachment with the email. An email attachment is a computer file sent along with an email message.

Modern email systems use the MIME standard; a message and all its attachments are encapsulated in a single multipart message, with base64 encoding used to convert binary into 7-bit ASCII text.
words.txt

sky
blud
rock
water
poem

We send this text file in the attachment.
attachment.go

package main

import (
    "bytes"
    "encoding/base64"
    "fmt"
    "io/ioutil"
    "log"
    "net/smtp"
    "strings"
)

type Mail struct {
    Sender  string
    To      []string
    Subject string
    Body    string
}

func main() {

    sender := "john.doe@example.com"

    to := []string{
        "roger.roe@example.com",
    }

    user := "9c1d45eaf7af5b"
    password := "ad62926fa75d0f"

    subject := "testing mail with attachment"
    body := "email body message"

    request := Mail{
        Sender:  sender,
        To:      to,
        Subject: subject,
        Body:    body,
    }

    addr := "smtp.mailtrap.io:2525"
    host := "smtp.mailtrap.io"

    data := BuildMail(request)
    auth := smtp.PlainAuth("", user, password, host)
    err := smtp.SendMail(addr, auth, sender, to, data)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Email sent successfully")
}

func BuildMail(mail Mail) []byte {

    var buf bytes.Buffer

    buf.WriteString(fmt.Sprintf("From: %s\r\n", mail.Sender))
    buf.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";")))
    buf.WriteString(fmt.Sprintf("Subject: %s\r\n", mail.Subject))

    boundary := "my-boundary-779"
    buf.WriteString("MIME-Version: 1.0\r\n")
    buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", 
        boundary))

    buf.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
    buf.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
    buf.WriteString(fmt.Sprintf("\r\n%s", mail.Body))

    buf.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
    buf.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
    buf.WriteString("Content-Transfer-Encoding: base64\r\n")
    buf.WriteString("Content-Disposition: attachment; filename=words.txt\r\n")
    buf.WriteString("Content-ID: <words.txt>\r\n\r\n")

    data := readFile("words.txt")

    b := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
    base64.StdEncoding.Encode(b, data)
    buf.Write(b)
    buf.WriteString(fmt.Sprintf("\r\n--%s", boundary))

    buf.WriteString("--")

    return buf.Bytes()
}

func readFile(fileName string) []byte {

    data, err := ioutil.ReadFile(fileName)
    if err != nil {
        log.Fatal(err)
    }

    return data
}

In the code example, we attach a text file to email.

boundary := "my-boundary-779"
buf.WriteString("MIME-Version: 1.0\r\n")
buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", 
    boundary))

A multipart/mixed MIME message is composed of a mix of different data types. Each body part is delineated by a boundary. The boundary parameter is a text string used to delineate one part of the message body from another.

buf.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
buf.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
buf.WriteString(fmt.Sprintf("\r\n%s", mail.Body))

Here we define the body part, which is plain text.

buf.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
buf.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
buf.WriteString("Content-Transfer-Encoding: base64\r\n")
buf.WriteString("Content-Disposition: attachment; filename=words.txt\r\n")
buf.WriteString("Content-ID: <words.txt>\r\n\r\n")

Thi is a part for the text file attachment. The content is encoded in base64.

data := readFile("words.txt")

We read the data from the words.txt file.

b := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
base64.StdEncoding.Encode(b, data)
buf.Write(b)
buf.WriteString(fmt.Sprintf("\r\n--%s", boundary))

buf.WriteString("--")

We write the base64 encoded data into the buffer. The last boundary is ended with two dash characters.

From: john.doe@example.com
To: roger.roe@example.com
Subject: testing mail with attachment
MIME-Version: 1.0
Content-Type: multipart/mixed; boundary=my-boundary-779

--my-boundary-779
Content-Type: text/plain; charset="utf-8"

email body message
--my-boundary-779
Content-Type: text/plain; charset="utf-8"
Content-Transfer-Encoding: base64
Content-Disposition: attachment; filename=words.txt
Content-ID: <words.txt>

c2t5CmJsdWQKcm9jawp3YXRlcgpwb2VtCg==
--my-boundary-779--

This is how the raw email looks like.

$ echo c2t5CmJsdWQKcm9jawp3YXRlcgpwb2VtCg== | base64 -d
sky
blud
rock
water
poem

We can decode the attachment with the base64 command.
Go email templates

In the following example, we use an email template to generate emails for multiple users.

$ mkdir template
$ cd template 
$ go mod init com/zetcode.TemplateEmail
$ go get github.com/shopspring/decimal 

We initiate the project and add the external github.com/shopspring/decimal package.
template.go

package main

import (
    "bytes"
    "fmt"
    "log"
    "net/smtp"
    "text/template"

    "github.com/shopspring/decimal"
)

type Mail struct {
    Sender  string
    To      string
    Subject string
    Body    bytes.Buffer
}

type User struct {
    Name  string
    Email string
    Debt  decimal.Decimal
}

func main() {

    sender := "john.doe@example.com"

    var users = []User{
        {"Roger Roe", "roger.roe@example.com", decimal.NewFromFloat(890.50)},
        {"Peter Smith", "peter.smith@example.com", decimal.NewFromFloat(350)},
        {"Lucia Green", "lucia.green@example.com", decimal.NewFromFloat(120.80)},
    }

    my_user := "9c1d45eaf7af5b"
    my_password := "ad62926fa75d0f"
    addr := "smtp.mailtrap.io:2525"
    host := "smtp.mailtrap.io"

    subject := "Amount due"

    var template_data = `
    Dear {{ .Name }}, your debt amount is ${{ .Debt }}.`

    for _, user := range users {

        t := template.Must(template.New("template_data").Parse(template_data))
        var body bytes.Buffer

        err := t.Execute(&body, user)
        if err != nil {
            log.Fatal(err)
        }

        request := Mail{
            Sender:  sender,
            To:      user.Email,
            Subject: subject,
            Body:    body,
        }

        msg := BuildMessage(request)
        auth := smtp.PlainAuth("", my_user, my_password, host)
        err2 := smtp.SendMail(addr, auth, sender, []string{user.Email}, []byte(msg))

        if err2 != nil {
            log.Fatal(err)
        }
    }

    fmt.Println("Emails sent successfully")
}

func BuildMessage(mail Mail) string {
    msg := ""
    msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
    msg += fmt.Sprintf("To: %s\r\n", mail.To)
    msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
    msg += fmt.Sprintf("\r\n%s\r\n", mail.Body.String())

    return msg
}

The example sends emails to multiple users to remind them about their debt. The text/template package is used to create an email template.

var users = []User{
    {"Roger Roe", "roger.roe@example.com", decimal.NewFromFloat(890.50)},
    {"Peter Smith", "peter.smith@example.com", decimal.NewFromFloat(350)},
    {"Lucia Green", "lucia.green@example.com", decimal.NewFromFloat(120.80)},
}

These are the borrowers.

var template_data = `
    Dear {{ .Name }}, your debt amount is ${{ .Debt }}.`

This is the template; it contains a generic message in which the .Name and .Debt placeholders are replaced with actual values.

for _, user := range users {

    t := template.Must(template.New("template_data").Parse(template_data))
    var body bytes.Buffer

    err := t.Execute(&body, user)
    if err != nil {
        log.Fatal(err)
    }

    request := Mail{
        Sender:  sender,
        To:      user.Email,
        Subject: subject,
        Body:    body,
    }

    msg := BuildMessage(request)
    auth := smtp.PlainAuth("", my_user, my_password, host)
    err2 := smtp.SendMail(addr, auth, sender, []string{user.Email}, []byte(msg))

    if err2 != nil {
        log.Fatal(err)
    }
}

We go over the borrowers and generate a email message for each of them. The Execute function applies a parsed template to the specified data object. After the message is generated, it is sent with SendMail.

In this tutorial, we have worked with emails in Go with the smtp package.
https://zetcode.com/golang/email-smtp/


