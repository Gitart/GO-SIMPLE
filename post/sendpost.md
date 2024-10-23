# Для создания PDF-файла в Go и отправки его на электронную почту, мы будем использовать несколько пакетов:

github.com/jung-kurt/gofpdf для создания PDF-файла.
net/smtp для отправки электронной почты.
Шаг 1: Установка зависимостей
Чтобы установить пакет для создания PDF-файла, используйте команду:

bash
Копировать код
go get github.com/jung-kurt/gofpdf
Шаг 2: Создание PDF-файла
Создадим простой PDF-файл с текстом.


```go
package main

import (
    "github.com/jung-kurt/gofpdf"
    "os"
)

func createPDF(filename string) error {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 16)
    pdf.Cell(40, 10, "Hello, this is a test PDF!")
    
    return pdf.OutputFileAndClose(filename)
}

func main() {
    err := createPDF("test.pdf")
    if err != nil {
        panic(err)
    }
    println("PDF created successfully")
}
```

Шаг 3: Отправка PDF-файла на почту
Теперь добавим функцию для отправки этого PDF-файла на почту. Мы будем использовать стандартный пакет net/smtp.

```go
package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "log"
    "mime/multipart"
    "mime/quotedprintable"
    "net/smtp"
    "path/filepath"
    "strings"
)

func sendEmail(smtpHost, smtpPort, senderEmail, senderPassword, recipientEmail, subject, body, attachmentPath string) error {
    // Setup headers
    headers := make(map[string]string)
    headers["From"] = senderEmail
    headers["To"] = recipientEmail
    headers["Subject"] = subject
    headers["MIME-Version"] = "1.0"
    headers["Content-Type"] = `multipart/mixed; boundary="BOUNDARY"`

    // Setup message
    var message bytes.Buffer
    for key, value := range headers {
        message.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
    }

    message.WriteString("\r\n--BOUNDARY\r\n")
    message.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
    message.WriteString("Content-Transfer-Encoding: quoted-printable\r\n\r\n")
    message.WriteString(quotedprintable.EncodeToString([]byte(body)))
    message.WriteString("\r\n--BOUNDARY\r\n")

    // Add attachment
    attachment, err := ioutil.ReadFile(attachmentPath)
    if err != nil {
        return err
    }
    fileName := filepath.Base(attachmentPath)
    message.WriteString(fmt.Sprintf("Content-Type: application/octet-stream; name=\"%s\"\r\n", fileName))
    message.WriteString("Content-Transfer-Encoding: base64\r\n")
    message.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\r\n\r\n", fileName))
    message.Write(attachment)
    message.WriteString("\r\n--BOUNDARY--\r\n")

    // Connect to SMTP server
    auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)
    return smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{recipientEmail}, message.Bytes())
}

func main() {
    // Create the PDF file
    err := createPDF("test.pdf")
    if err != nil {
        log.Fatalf("Error creating PDF: %v", err)
    }

    // Email configuration
    smtpHost := "smtp.gmail.com"
    smtpPort := "587"
    senderEmail := "your-email@gmail.com"
    senderPassword := "your-email-password"
    recipientEmail := "recipient-email@example.com"
    subject := "Test PDF"
    body := "Please find the attached PDF file."

    // Send the email with PDF attachment
    err = sendEmail(smtpHost, smtpPort, senderEmail, senderPassword, recipientEmail, subject, body, "test.pdf")
    if err != nil {
        log.Fatalf("Error sending email: %v", err)
    }

    fmt.Println("Email sent successfully")
}
```


## Создание PDF-файла:
Используем gofpdf для создания простого PDF-файла с текстом.

## Отправка PDF на почту:
В функции sendEmail мы формируем письмо с заголовками и вложением (PDF-файл), а затем отправляем его через SMTP-сервер.

Параметры для отправки:
smtpHost: Укажите хост вашего SMTP-сервера (например, для Gmail — smtp.gmail.com).
smtpPort: Порт SMTP-сервера (обычно 587 для Gmail).
senderEmail: Электронная почта отправителя.
senderPassword: Пароль или ключ доступа для почты отправителя (в случае Gmail может потребоваться создание специального пароля для приложений).
recipientEmail: Электронная почта получателя.
attachmentPath: Путь к файлу, который нужно отправить.
Дополнительные настройки:
Для работы с Gmail, возможно, потребуется включить доступ для ненадёжных приложений или настроить двухфакторную аутентификацию и использовать пароль для приложений.
Результат:
Этот код создаст PDF-файл и отправит его на указанный адрес электронной почты.
