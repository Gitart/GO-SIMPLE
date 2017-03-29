# Введение

### В этой статье рассмотрено:   
— Создание структуры данных с методами загрузки и сохранения    
— Использование пакета http для создания веб-приложения    
— Использование пакета template для обработки HTML-шаблонов    
— Использование пакета regexp для проверки вводимых пользователем данных   
— Использование замыканий.    

###  Предпологаемые знания:   
— Опыт программирования   
— Понимание основных веб-технологий (HTTP, HTML)    
— Знание некоторых UNIX-команд     

###  Начало
Для начала, вам необходим компьютер с ОС Linux, OS X или FreeBSD для запуска Go. Если у вас такового не имеется,  
вы можете установить Linux на виртуальную машину (с помощью VirtualBox или подобного ПО) или VPS.  


Создайте новую директорию для нашего примера и перейдите (cd) в неё:   
```
$ mkdir ~/gowiki   
$ cd ~/gowiki
```

Создайте файл с названием wiki.go, откройте в своем любимом текстовом редакторе и добавьте следующие строки:
package main    

```golang
import (
	"fmt"
	"io/ioutil"
	"os"
)
```

Мы произвели импорт пакетов fmt, ioutil и os из стандартной библиотеки Go. Позже, для добавления дополнительной      
функциональности, мы добавим больше пакетов в блок import.

### Структуры данных
Начнем с определения структур данных. Вики состоит из набора связанных друг с другом страниц, каждая из которых     
имеет заголовок и тело (содержимое страницы). На этом этапе мы определим Page как структуру с двумя элементами, 
соответствующих заголовку и телу:

```golang
type Page struct {
	Title	string
	Body	[]byte
}
```

Тип []byte означает «срез byte» (см. «Эффективный Go» для подробной информации о срезах).
Элемент Body имеет тип данных []byte, а не string для того, чтобы позже использовать его библиотеками io.

Структура Page описывает, как информация о странице будет храниться в памяти. А как насчет нашего постоянного хранилища?   
Мы можем сделать его, создав метод save для Page:

```golang
func (p *Page) save() os.Error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}
```

Заголовок метода читается как: «Этот метод называется save, он воспринимает в качестве отправителя p,     
который является указателем на Page. Он не принимает параметров и возвращает значения типа os.Error».
Этот метод будет сохранять Body из Page в текстовый файл. Для простоты мы будем использовать Title в качестве имени файла.     
Метод save возвращает значение типа os.Error, т.к. такого типа возвращаемое значение функции WriteFile (функция из стандартной      
библиотеки, записывающая срез байтов в файл). Метод save возвращает код ошибки чтобы дать возможность приложению обработать     
её если что-то пойдет не так при записи файла. Если все прошло должным образом, Page.save() вернет значение nil     
(нулевое значение для указателей, интерфейсов и некоторых других типов).

Восьмеричная целая постоянная 0600, переданная в качестве третьего параметра WriteFile обозначает,      
что файл необходимо создать с правами на запись и чтение только для текущего пользователя (см. man для open(2)).

Также мы хотим загружать страницы:   

```golang
func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}
}
```

Функция loadPage формирует имя файла из Title, считывает содержимое файла в новую Page и возвращает указатель на новую страницу.  
Функции могут возвращать несколько значений. Функция стандартной библиотеки io.ReadFile возвращает типы []byte и os.Error.    
В loadPage ошибка не обрабатывается. «Пустой индетификатор», представленный в виде символа подчеркивания (_)   
используется для того, чтобы отбросить возвращаемое значение (на деле присваивая значение в никуда).

Но что произойдет, если ReadFile вернет ошибку? Например, файл может не существовать. Мы не должны игнорировать
такие ошибки. Давайте отредактируем функцию и вернем *Page и os.Error.  

```golang
func loadPage(title string) (*Page, os.Error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
```

Вызывающие функцию части программы теперь могут проверить второй параметр. Если он будет равен nil,      
страница загружена успешна. Если нет, os.Error может быть обработана (за деталями см. документацию по пакету os.

На данный момент у нас есть простая структура данных и возможность сохранять её и загружать из файла.     
Давайте напишем функцию main, чтобы протестировать написанный код:

```golang
func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
````

После компиляции и выполнения этого кода, должен появиться файл TestPage.txt, содержащий в себе содержимое p1.     
Файл должен быть считан в структуру p2, а его параметр Body выведен на экран.

Вы можете скомпилировать и запустить программу следующим образом:
```
$ 8g wiki.go
$ 8l wiki.8
$ ./8.out
This is a sample page.
```

(Команды 8g и 8l предназначены для GOARCH=386. Если у вас система amd64, поставьте шестерки вместо восьмерок)

### Представляем пакет http
Ниже приведен пример работающего простого веб-сервера

```golang
package main

import (
	"fmt"
	"http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
```

Функция main начинается с вызова http.HandleFunc, которая сообщает пакету http обрабатывать все запросы в корневой ("/")
директории с помощью handler.

Затем она вызывает http.ListenAndServe, указывая, что она должна слушать порт 8080 любого интерфейса (":8080").   
(Пока что не обращайте на второй (nil) параметр). Эта функция будет блокироваться, пока программа не завершиться.

Функция handler имеет тип http.HandlerFunc. В качестве аргументов она принимает http.ResponseWriter и http.Request.   

Значение http.ResponseWriter компонует ответ HTTP-сервера. Записывая его, мы отправляем данные HTTP-клиенту.   

http.Request является структурой данных, которая представляет собой HTTP-запрос клиента. Строка r.URL.Path
часть запрашиваемого URL, представляющая собой путь к запрашиваемой директории. Замыкание [1:]   
обозначает «создать подслой Path от первого символа до конца». Это отсечет первую "/" от пути к директории.  

Если вы запустите программу и запросите URL http://localhost:8080/monkeys, программа выдаст страницу,      
содержащую «Hi there, I love monkeys!».

### Использование http для обработки wiki-страниц


Для использования пакета http, его необходимо импортировать:
```golang
import (
	"fmt"
	"http"
	"io/ioutil"
	"os"
)
```

Давайте создадим обработчик для просмотра вики-страницы:

```golang
const lenPath = len("/view/")

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
```

Сначала эта функция извлекает заголовок страницы из r.URL.Path, часть из запрашевоемого URL, обозначающую путь.    
Глобальная постоянная lenPath содержит в себе длину "/view/", с которой начинается имя запрашиваемого пути.     
От Path отрезается с помощью [lenPath:] первые 6 символов строки. Это сделано потому, что путь всегда начинается 
со строки "/view/", которая не является частью заголовка страницы.
Затем функция загружает данные страницы, форматирует страницу с помощью строки простого HTML и записывает всё в w,
тип которой http.ResponseWriter.

Снова обратите внимание на использование _ для игнорирования возвращаемого значения типа os.Error из loadPage.
Здесь это сделано для упращения, но являет собой пример плохого программирования. Мы вернемся к этому позже.

Для использования этого обработчика, мы создадим функцию main, которая проинициализирует http, 
используя viewHandler для обработки любых запросов к директории /view/.   

```golang
func main() {
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
```

Полный исходный код примера

Давайте создадим файл с данными страницы (test.txt), скомпилируем наш код и попробуем обработать вики-страницу:

```
$ echo "Hello world" > test.txt
$ 8g wiki.go
$ 8l wiki.8
$ ./8.out
```

Запустив этот веб-сервер, запрос http://localhost:8080/view/test должен отобразить страницу с заголовком «test»    
и содержащую слова «Hello world»

### Редактирование страниц
Вики не была бы вики без возможности редактировать страницы. Давайте создадим два новых обработчика:     
один назовем editHandler для отображения формы редактирования страницы, а второй, с названием saveHandler,     
для сохранения данных, введенных в форму.

Для начала мы добавим их в main():

```golang
func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}
```

Функция editHandler загружает страницу (или, если страница не существует, создает пустую структуру Page) и отображает HTML-форму.

```golang
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>",
		p.Title, p.Title, p.Body)
}
```

Эта функция будет работать хорошо, но «вшитый» HTML-код в тело функции выглядит уродливо. Конечно, есть лучший способ.

### Пакет template
Пакет template является частью стандартной библиотеки Go. Мы можем использовать template для хранения HTML в отдельном файле,     
что позволит нам изменят отображение страницы редактирования без правки кода Go.

Для начала, мы должны добавить template в список импорта

```golang
import (
	"http"
	"io/ioutil"
	"os"
	"template"
)
```

Давайте создадим файл шаблона, содержащий HTML-форму. Откройте новый файл с именем edit.html и добавьте следующие строки:
```html
<h1>Editing {Title}</h1>

<form action="/save/{Title}" method="POST">
<div><textarea name="body" rows="20" cols="80">{Body|html}</textarea></div>
<div><input type="submit" value="Save"></div>
</form>
```

Измените editHandler для использования шаблона вместо вставленного в неё HTML-кода:

```golang
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFile("edit.html", nil)
	t.Execute(p, w)
}
```

Метод t.Execute заменяет все вхождения {Title} и {Body} на значения p.Title и p.Body, записывая результирующий    
HTML в http.ResponseWriter.
Обратите внимание, что мы использовали {Body|html} выше в шаблоне. Часть |html запрашивает обработчик   
шаблона пропустить значение Body через обработчик html, прежде чем вывести его, что позволит экранировать 
HTML-символы (например, замена > на &gt;). Это позволит избежать нарушение HTML-формы пользователем.

Теперь, когда мы удалили выражение fmt.Sprintf, мы можем удалить "fmt" из списка import.   

Давайте также создадим шаблон для нашего viewHandler под названием 
view.html:   

```html
<h1>{Title}</h1>
<p>[<a href="/edit/{Title}">edit</a>]</p>
<div>{Body}</div>
```

Измените viewHandler следующим образом:

```golang
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	p, _ := loadPage(title)
	t, _ := template.ParseFile("view.html", nil)
	t.Execute(p, w)
}
```

Обратите внимание, что мы используем практически одинаковый код в обоих обработчиках.     
Давайте уберем это дублирование переносом кода обработки шаблона в отдельную функцию:

```golang
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	p, _ := loadPage(title)
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFile(tmpl+".html", nil)
	t.Execute(p, w)
}
```


Теперь обработчики короче и яснее.
Обработка несуществующих страниц
Что случится, если вы запросите /view/APageThatDoesntExist? Программа завершится с ошибкой. Это произойдет потому,    
что она игнорирует возвращаемое значение ошибки от loadPage. Вместо этого, если запрашиваемая страница не существует,    
она должна перенаправить клиента на редактирование новой страницы:

```golang
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}
```

Функция http.Redirect добавляет код статуса HTTP http.StatuseFound (302) и заголовок Location в ответ сервера.

### Сохранение страниц
Функция saveHandler будет обрабатывать данные формы.

```golang
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
```

Заголовок страницы (взятый из URL) и единственное поле в форме, Body, хранятся в новой Page. Метод save()     
затем вызывается для записи данный в файл, а клиент будет переадресован на страницу /view/.   
Значение, возвращаемое из FormValue имеет тип string. Мы должны преобразовать это значение в тип []byte    
до присваивания в структуру Page. Для этого преобразования мы используем []byte(body).

### Обработка ошибок
В программе есть несколько мест, где ошибки проигнорированы. Это неправильный подход, т.к. при возникновении    
ошибки программа завершится аварийно. Лучшим решением будет обработать ошибки и вернуть сообщение об ошибке пользователю.    
В таком случае, если что-то пойдет неправильно, сервер продолжит функционировать, а пользователь будет проинформирован.  

Прежде всего, давайте обработаем ошибки в renderTemplate:

```golang
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFile(tmpl+".html", nil)
	if err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(p, w)
	if err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
	}
}
```

Функция http.Error отправляет специальный код HTTP (в данном случае, «Internal Server Error») и сообщение об ошибке.     
Теперь давайте исправим saveHandler:

```golang
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err = p.save()
	if err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
```

Любые ошибки, которые произойдут во время p.save() отобразятся пользователю.

### Кэширование шаблона
В коде имеется неэффективный код: renderTemplate вызывает ParseFile каждый раз, когда генерирует страницу.     
Лучшим подходом было бы вызывать ParseFile единожды для каждого шаблона при инициализации программы,     
а результирующие значения *Template хранить в структуре данных для дальнейшего использования.

Для начала, мы создадим глобальное отображение с именем templates, в котором будем хранить наши значения *Template,     
с ключевыми значениями string (имена шаблонов).

```golang
var templates = make(map[string]*template.Template)
```

Затем мы создадим функцию init, которая будет вызываться перед main при инициализации программы.     
Функция template.MustParseFile удобная оболочка вокруг ParseFile, которая не возвращает код ошибки.    
Вместо этого, она экстренно завершает программу. Это подходит в данном случае — если шаблоны     
не смогут загрузиться, единственное, что останется сделать, это выйти из программы.

```golang
func init() {
	for _, tmpl := range []string{"edit", "view"} {
		templates[tmpl] = template.MustParseFile(tmpl+".html", nil)
	}
}
```

Цикл for используется с выражением range для прохода по массиву постоянных, составленного из имен шаблонов,     
которые мы хотим пропарсить. Если мы добавляем новые шаблоны в нашу программу, мы должны добавить их имена в этот массив.

Затем мы редактируем нашу функцию renderTemplate для вызова метода Execute для соответствующего Template из templates:  

```golang
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates[tmpl].Execute(w, p)
	if err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
	}
}
```

### Валидация
Как вы могли заметить, у нашей программы есть серьезная дыра в безопасности: пользователь может перейти в любую    
директорию на сервере и выполнить там чтение или запись. Чтобы избежать этого, мы можем написать функцию валидации   
с регулярным выражением.   

Прежде всего, добавьте "regexp" в список import. Затем мы можем создать глобальную переменную для хранения нашего    
валидационного регулярного выражения:
```golang
var titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")   
```

Функция regexp.MustCompile будет парсить и выполнять регулярное выражение и возвращать regexp.Regexp.MustCompile;    
так же, как template.ParseFile, в отличие от Compile, если возникнет ошибка при выполнении регулярного выражения,    
завершится экстренно, в то время, как Compile возвратит os.Error в качестве второго параметра.   

Теперь давайте напишем функцию, которая вырезает строку заголовка из запрашиваемого URL и проверяет     
его нашим выражением titleValidator:

```golang
func getTitle(w http.ResponseWriter, r *http.Request) (title string, err os.Error) {
	title = r.URL.Path[lenPath:]
	if !titleValidator.MatchString(title) {
		http.NotFound(w, r)
		err = os.NewError("Invalid Page Title")
	}
	return
}
```

Если заголовок валидный, он будет возвращен вместе со значением nil в качестве ошибки.    
Если заголовок невалиден, функция напишет «404 Not Found» и возвратит ошибку обработчику.

Давайте добавим вызов getTitle в каждый из обработчиков:

```golang
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err = p.save()
	if err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
		return
	}
```


### Представляем литералы функций и замыкания
Отлов ошибок в каждом обработчике привел к созданию большого количества повторяющегося кода.     
Что если бы мы смогли обернуть каждый из обработчиков в функцию, которая будет делать валидацию     
и проверку ошибок? Литералы функций в Go дают мощную абстрактную функциональность, которая может нам помочь с этой задачей.   

Для начала, мы перепишим обределение функции каждого обработчика, чтобы получить доступ к строке заголовка:  

```golang
func viewHandler(w http.ResponseWriter, r *http.Request, title string)
func editHandler(w http.ResponseWriter, r *http.Request, title string)
func saveHandler(w http.ResponseWriter, r *http.Request, title string)
```

Теперь давайте определим обертку, которая принимает функцию указанного типа, а возвращает функцию     
типа http.HandlerFunc (подходящую для передачи в функцию http.HandleFunc):    

```golang
func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Здесь мы будем извлекать заголовок страницы из Request
		// и вызывать переданный обработчик 'fn'
	}
}
```

Возвращаемая функция называется замыканием, т.к. она связывает значения, определенные вне неё. В этом случае,    
переменная fn (единственный аргумент makeHandler) заключена замыканием. Переменная fn будет единственная для    
наших обработчиков сохранения, редактирования и загрузки.   

```golang
Теперь мы можем взять код из getTitle и использовать его тут (с небольшими изменениями):
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Path[lenPath:]
		if !titleValidator.MatchString(title) {
			http.NotFound(w, r)
			return
		}
		fn(w, r, title)
	}
}
```

Замыкание, возвращаемое makeHandler, является функцией, которая принимает http.ResponseWriter и 
http.Request (другими словами, http.HandlerFunc). Замыкание извлекает title из запрашиваемого пути   
к директории, проверяет его с помощью регулярного выражения TitleValidator. Если title невалиден,    
ошибка будет записана в ResponseWriter с помощью функции http.NotFound. Если title валиден, функция   
fn будет вызвана с аргументами ResponseWriter, Request и title.

Теперь мы можем обернуть обработчик функций с помощью makeHandler, прежде чем она будет зарегистрирована пакетом http.

```golang
func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.ListenAndServe(":8080", nil)
}
```

Наконец, мы уберем вызовы getTitle из функций обработки, сделав их гораздо более простыми:

```golang
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
```


### Тестирование
Нажмите здесь для просмотра полученной версии кода.

Перекомпилируйте код и запустите приложение:
```
$ 8g wiki.go
$ 8l wiki.8
$ ./8.out
```

При посещении http://localhost:8080/view/ANewPage должна отобразиться страница с формой редактирования.   
У вас должна быть возможность ввести текст, нажать 'Save' и быть перенаправленным на вновь созданную страницу.  

