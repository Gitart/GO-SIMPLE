# Веб\-приложение для генерации QR\-кода в Голанге

Код быстрого отклика \- это двумерный пиктографический код, используемый для его быстрой читаемости и сравнительно большой емкости памяти. Код состоит из черных модулей, расположенных в виде квадратного шаблона на белом фоне. Закодированная информация может состоять из данных любого типа (например, двоичных, буквенно\-цифровых символов или символов кандзи).

Основной пример веб\-приложения для генерации штрих\-кода. Алгоритм или внутренняя логика для генерации штрих\-кода доступны в стороннем пакете штрих\-кодов. Здесь цель \- показать пример использования пакета и создания веб\-приложения.

---

## 1\. Установите необходимый пакет

Пакет **штрих\-кода** может быть использован для создания различных типов штрих\-кодов. Вы можете установить этот пакет, выполнив следующую команду в вашем терминале git bash:

```sh
go get github.com/boombuler/barcode
```
## 2\. Разработка

### 2.1 Исходный код main.go

**Основная** функция начинается с вызова http.HandleFunc, который говорит пакет HTTP , чтобы обрабатывать все запросы к веб \- корня ( «/») с **homeHandler** . Функция **homeHandler** имеет тип http.HandlerFunc. В качестве аргументов он принимает http.ResponseWriter и **http.Request** .

Функция **viewCodeHandler** , которая позволит пользователям просматривать сгенерированный QR\-код на новой странице. Он будет обрабатывать URL с префиксом **"/ generator /"** . Функция template.ParseFiles будет читать содержимое файла **generator.html** и возвращать \* template.Template.

```go
package main

import (
	"image/png"
	"net/http"
	"text/template"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type Page struct {
	Title string
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/generator/", viewCodeHandler)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "QR Code Generator"}

	t, _ := template.ParseFiles("generator.html")
	t.Execute(w, p)
}

func viewCodeHandler(w http.ResponseWriter, r *http.Request) {
	dataString := r.FormValue("dataString")

	qrCode, _ := qr.Encode(dataString, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 512, 512)

	png.Encode(w, qrCode)
}
```
Функция **FormValue** выдаст значение поля ввода dataString, которое будет использоваться для генерации QR\-кода с использованием функции **Encode** .

### 2.2 Исходный код генератора .html.

Файл шаблона, содержащий форму HTML.

```go
<h1>{{.Title}}</h1>
<div>Please enter the string you want to QRCode.</div>
<form action="generator/" method=post>
    <input type="text" name="dataString">
    <input type="submit" value="Submit">
</form>
```

## 3\. Исполнение

#### С помощью командной строки или команды putty run "go run main.go"

Вы увидите вывод, подобный изображенному ниже.

![](https://www.golangprograms.com/media/wysiwyg/golangwebapps/qr-code/first-screen.JPG) ![](https://www.golangprograms.com/media/wysiwyg/golangwebapps/qr-code/second-screen.JPG)

### Самые полезные на этой неделе

*   [Преобразовать целочисленный тип в строковый тип](https://www.golangprograms.com/convert-integer-type-to-string-in-go.html)
*   [Как ждать, пока Goroutines завершит выполнение?](https://www.golangprograms.com/program-demonstrates-how-to-wait-for-goroutines-to-finish-execution.html)
*   [Регулярное выражение для совпадения формата времени ЧЧ: ММ в Голанге](https://www.golangprograms.com/regular-expression-for-matching-hh-mm-time-format-in-golang.html)
*   [Как создать миниатюру изображения?](https://www.golangprograms.com/how-to-create-thumbnail-of-an-image-in-golang.html)
*   [Пример использования функции Weekday и YearDay](https://www.golangprograms.com/example-to-use-weekday-and-yearday-function.html)
*   [Golang Читать Написать Создать и удалить текстовый файл](https://www.golangprograms.com/golang-read-write-create-and-delete-text-file.html)
*   [Как объединить два или более фрагментов в Голанге?](https://www.golangprograms.com/how-to-concatenate-two-or-more-slices-in-golang.html)
*   [Как прочитать имена всех файлов и папок в текущем каталоге?](https://www.golangprograms.com/how-to-read-names-of-all-files-and-folders-in-current-directory.html)
*   [Как конвертировать строку в целочисленный тип в Go?](https://www.golangprograms.com/how-to-convert-string-to-integer-type-in-go.html)
*   [Golang Читать, писать и](https://www.golangprograms.com/golang-read-write-and-process-data-in-csv.html)
