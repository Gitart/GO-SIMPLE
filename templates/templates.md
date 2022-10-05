
# 📒 ZetCode

All Go Python C# Java JavaScript Subscribe

Ebooks

    PyQt5 ebook
    Tkinter ebook
    SQLite Python
    wxPython ebook
    Windows API ebook
    Java Swing ebook
    Java games ebook
    MySQL Java ebook

Go template

last modified April 23, 2022

Go template tutorial shows how to create templates in Golang with standard library.

A template engine or template processor is a library designed to combine templates with a data model to produce documents. Template engines are often used to generate large amounts of emails, in source code preprocessing, or producing dynamic HTML pages.

We create a template engine, where we define static parts and dynamic parts. The dynamic parts are later replaced with data. The rendering function later combines the templates with data. A template engine is used to combine templates with a data model to produce documents.

Go contains two template packages: text/template and html/template. Both share the same interface. The html/template automatically secures HTML output against certain attacks.

Within the template API, the Parse function parses template strings present in the program, the ParseFiles loads and parses template files, and Execute renders a template to output using specific data fields. The New function allocates a new, undefined template with the given name.

Similar to many other template engines, data evaluations and control structures are delimited by {{ and }}.
Go template Parse

The Parse function parses template strings within a Go program.
main.go

package main

import (
    "log"
    "os"
    "text/template"
)

type User struct {
    Name       string
    Occupation string
}

func main() {

    user := User{"John Doe", "gardener"}

    tmp := template.New("simple")
    tmp, err := tmp.Parse("{{.Name}} is a {{.Occupation}}")

    if err != nil {
        log.Fatal(err)
    }

    err2 := tmp.Execute(os.Stdout, user)

    if err2 != nil {
        log.Fatal(err2)
    }
}

The example creates a simple text message.

type User struct {
    Name       string
    Occupation string
}

This is the data type used in the template; the fields must be exported, that is capitalized.

tmp := template.New("simple")

A new template is created.

tmp, err := tmp.Parse("{{.Name}} is a {{.Occupation}}")

We parse the template string with Parse. Using the dot operator, we access the fields that are passed to the template engine.

$ go run main.go
John Doe is a gardener

Go template Must

The Must function is a helper function which takes care of error checking.
main.go

package main

import (
    "log"
    "os"
    "text/template"
)

type User struct {
    Name       string
    Occupation string
}

func main() {

    user := User{"John Doe", "gardener"}

    tmp := template.Must(template.New("simple").Parse("{{.Name}} is a {{.Occupation}}"))

    f, err := os.Create("output.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    err2 := tmp.Execute(f, user)

    if err2 != nil {
        log.Fatal(err2)
    }
}

In the example, we use a template to write a message to a file. We also use the Must function.

$ go run main.go
$ cat output.txt
John Doe is a gardener

Go template ParseFiles

The ParseFiles function creates a new template and parses the template definitions from the give file names.
message.txt

{{.Name}} is a {{.Occupation}}

This is the template file.
main.go

package main

import (
    "log"
    "os"
    "text/template"
)

type User struct {
    Name       string
    Occupation string
}

func main() {

    user := User{"John Doe", "gardener"}

    tmp, err := template.ParseFiles("message.txt")

    if err != nil {
        log.Fatal(err)
    }

    err2 := tmp.Execute(os.Stdout, user)

    if err2 != nil {
        log.Fatal(err2)
    }
}

In the example, we create a simple message from a template file.
Go template range

The range directive goes through items of an array, slice, map, or channel insice a template.
words.txt

{{range .Words -}}
    {{ .}}
{{end}}

In the template, we use the range directive to go through the elements of the Words data structure. The - character strips whitespace characters.
main.go

package main

import (
    "log"
    "os"
    "text/template"
)

type Data struct {
    Words []string
}

func main() {

    data := Data{Words: []string{"sky", "blue", "forest", "tavern", "cup", "cloud"}}

    tmp, err := template.ParseFiles("words.txt")

    if err != nil {
        log.Fatal(err)
    }

    err2 := tmp.Execute(os.Stdout, data)

    if err2 != nil {
        log.Fatal(err2)
    }
}

In the program, we pass a slice of words to the tempate engine. We get a list of words as the output.

$ go run main.go
sky
blue
forest
tavern
cup
cloud

Go template conditions

Conditions can be created with if/else if/else directives.
data.txt

{{- range .Todos -}}
    {{if .Done}}
        {{- .Title -}}
    {{end}}
{{end}}

In the template file, we use the if directive to output only tasks that are finished.
main.go

package main

import (
    "log"
    "os"
    "text/template"
)

type Todo struct {
    Title string
    Done  bool
}

type Data struct {
    Todos []Todo
}

func main() {

    data := Data{Todos: []Todo{
        {Title: "Task 1", Done: false},
        {Title: "Task 2", Done: true},
        {Title: "Task 3", Done: true},
        {Title: "Task 4", Done: false},
        {Title: "Task 5", Done: true},
    }}

    tmp := template.Must(template.ParseFiles("data.txt"))

    err2 := tmp.Execute(os.Stdout, data)

    if err2 != nil {
        log.Fatal(err2)
    }
}

We generate on output from a slice of todos. Only tasks that are done are included in the output.
Go template server example

The following example uses templates in a server application.
layout.html

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Users</title>
</head>

<body>
    <table>
        <thead>
            <tr>
                <th>Name</th>
                <th>Occupation</th>
            </tr>
        </thead>
        <tbody>
            {{ range .Users}}
            <tr>
                <td>{{.Name}}</td>
                <td>{{.Occupation}}</td>
            </tr>
            {{ end }}
        </tbody>
    </table>
</body>
</html>

The output is an HTML file. The data is inserted into HTML table.
main.go

package main

import (
    "html/template"
    "log"
    "net/http"
)

type User struct {
    Name       string
    Occupation string
}

type Data struct {
    Users []User
}

func main() {
    tmp := template.Must(template.ParseFiles("layout.html"))
    http.HandleFunc("/users", func(w http.ResponseWriter, _ *http.Request) {

        data := Data{
            Users: []User{
                {Name: "John Doe", Occupation: "gardener"},
                {Name: "Roger Roe", Occupation: "driver"},
                {Name: "Peter Smith", Occupation: "teacher"},
            },
        }
        tmp.Execute(w, data)
    })

    log.Println("Listening...")
    http.ListenAndServe(":8080", nil)
}

The web server returns an HTML page with a table of users for the /users URL path.
Go email templates

In the following example, we use an email template to generate emails for multiple users. We use the Mailtrap email testing service.

$ mkdir template
$ cd template 
$ go mod init com/zetcode/TemplateEmail
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

    my_user := "username"
    my_password := "password"
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

In the example, we send emails to multiple users to remind them about their debt.

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

We iterate over the slice of borrowers and generate a email message for each of them. TheExecute function 
applies a parsed template to the specified data object. After the message is generated, it is sent with SendMail.
In this tutorial, we have created dynamic documents using the built-in template package.
https://zetcode.com/golang/template/


