### Пакет xml в Golang - метод Marshal

Пакет **xml** реализует простой анализатор XML 1.0, который понимает пространства имен XML.

```
func Marshal(v interface{}) ([]byte, error)

```

Marshal возвращает XML кодированный v.

Marshal обрабатывает массив или срез путем маршалинга каждого из элементов. Marshal обрабатывает указатель путем маршалинга значения, на которое он указывает, или, если указатель равен nil, ничего не записывая. Marshal обрабатывает значение интерфейса путем маршалирования содержащегося в нем значения или, если значение интерфейса равно nil, ничего не записывая. Marshal обрабатывает все остальные данные, записывая один или несколько элементов XML, содержащих данные.

Имя для элементов XML взято в порядке предпочтения:

*   тег в поле XMLName, если данные представляют собой структуру
*   значение поля XMLName типа Name
*   тег поля структуры, используемого для получения данных
*   имя поля структуры, используемого для получения данных
*   название маршалированного типа

Элемент XML для структуры содержит маршалированные элементы для каждого из экспортируемых полей структуры, за исключением:

*   поле XMLName, описанное выше, опущено.
*   поле с тегом "-" опущено.
*   поле с тегом "name,attr" становится атрибутом с данным name в элементе XML.
*   поле с тегом "attr" становится атрибутом с именем поля в элементе XML.
*   поле с тэгом ",chardata" записывается как символьные данные, не как элемент XML.
*   поле с тегом ",cdata" записывается как символьные данные, обернутые в один или несколько тегов <!\[CDATA\[ ... \]\]>, а не как элемент XML.
*   поле с тэгом ",innerxml" записывается дословно, не подчинено к обычной процедуре маршалинга.
*   поле с тегом ",comment" записывается как комментарий XML, не объект для обычной процедуры маршалинга. Он не должен содержать строк "--" внутри него.
*   поле с тегом, включающим опцию "omitempty", опущено если значение поля пусто. Пустые значения false, 0, любой nil указатель или значение интерфейса и любой массив, срез, карта или строка нулевой длины.
*   анонимное структурное поле обрабатывается так, как если бы поля его значение было частью внешней структуры.
*   поле, реализующее Marshaler, пишется путем вызова его MarshalXML метода.
*   поле, реализующее encoding.TextMarshaler записывается путем кодирования результат его метода MarshalText как текст.

Если в поле используется тег"a>b>c", то элемент c будет вложен в родительские элементы a и b. Поля, которые появляются рядом друг с другом и называют одного и того же родителя, будут заключены в один элемент XML.

Если имя XML для поля структуры определяется как тегом поля, так и полем XMLName структуры, имена должны совпадать.

Marshal вернет ошибку, если будет предложено маршалировать канал, функцию или карту.

Метод MarshalIndent

```
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)

```

MarshalIndent работает как Marshal, но каждый элемент XML начинается с новой строки с отступом, которая начинается с prefix и сопровождается одной или несколькими копиями indent в зависимости от глубины вложения.

#### Пример использования Marshal (MarshalIndent)

```go
package main

import (
    "encoding/xml"
    "fmt"
    "os"
)

func main() {
    type Address struct {
        City, State string
    }
    type Person struct {
        XMLName   xml.Name `xml:"person"`
        Id        int      `xml:"id,attr"`
        FirstName string   `xml:"name>first"`
        LastName  string   `xml:"name>last"`
        Age       int      `xml:"age"`
        Height    float32  `xml:"height,omitempty"`
        Married   bool
        Address
        Comment string `xml:",comment"`
    }

    v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
    v.Comment = " Need more details. "
    v.Address = Address{"Hanga Roa", "Easter Island"}

    output, err := xml.MarshalIndent(v, "  ", "    ")
    if err != nil {
        fmt.Printf("error: %v\n", err)
    }

    os.Stdout.Write(output)
}

```

Вывод:

```
  <person id="13">
      <name>
          <first>John</first>
          <last>Doe</last>
      </name>
      <age>42</age>
      <Married>false</Married>
      <City>Hanga Roa</City>
      <State>Easter Island</State>
      <!-- Need more details. -->
  </person>
```
