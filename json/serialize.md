# JSON сериализация структур Go

Posted on 2020, Apr 07 3 мин. чтения

Сериализация/десериализация JSON достаточно часто встречающаяся задача. В библиотеке go есть пакадж [encoding/json](https://golang.org/pkg/encoding/json/), который ее решает.

С помощью json.Marshal() мы сериализуем в JSON и с помощью json.Unmarshal() конвертируем обратно в go объект.

Например сериализуем/десериализуем массив строк:

```go
	arr := []string{"test", "test"}

	resJSON, err := json.Marshal(arr)
	if err != nil {
		log.Fatal(err)
	}

	resArr := []string{}
	if err := json.Unmarshal(resJSON, &resArr); err != nil {
		log.Fatal(err)
	}
```

[https://play.golang.org/p/VJas8MWm1IZ](https://play.golang.org/p/VJas8MWm1IZ)

Рассмотрим сериализацию структур:

```go
	testStruct := Test{
		FieldA: "test",
		FieldB: 10,
	}

	resJSON, err := json.Marshal(testStruct)
	if err != nil {
		log.Fatal(err)
	}
```

[https://play.golang.org/p/b5yRvST2l6y](https://play.golang.org/p/b5yRvST2l6y)

Из вывода программы видно, что названия полей сериализуются как “FieldA”, “FieldB”. Но в случае json, хотелось бы такой формат имен “fieldA”, “fieldB”. Для управления сериализацией `json.Marshal` использует систему тегов.

Например укажем имена полей:

```go
type Test struct {
	FieldA string `json:"fieldA"`
	FieldB int64  `json:"fieldB"`
}
```

[https://play.golang.org/p/4UxCLqV0aJc](https://play.golang.org/p/4UxCLqV0aJc)

Типы json полей (string, number, boolean) выводятся из типов полей структуры, но с помощью тегов можно указать, что поле должно сериализоваться, как строка:

```go
type Test struct {
	FieldA string `json:"fieldA"`
	FieldB int64  `json:"fieldB,string"`
	FieldC time.Time
}
```

[https://play.golang.org/p/HI4M97qTpIg](https://play.golang.org/p/HI4M97qTpIg)

Для number, boolean таких подсказок не предусмотрено.

C помощью метки “\-” можно выключать поле из сериализации:

```go
type Test struct {
	FieldA string    `json:"fieldA"`
	FieldB int64     `json:"fieldB,string"`
	FieldC time.Time `json:"-"`
}
```

[https://play.golang.org/p/9VOY5Der\_hf](https://play.golang.org/p/9VOY5Der_hf)

А с помощью метки “omitempty” можно обрезать пустые поля при сериализации. Например инт поле с значение 0, не будет сериализоваться:

```go
    type Test struct {
        FieldA string    `json:"fieldA"`
        FieldB int64     `json:"fieldB,string,omitempty"`
        FieldC time.Time `json:"-"`
    }

	withEmptyB := Test{
		FieldA: "test",
		FieldB: 0,
		FieldC: time.Now(),
	}
```

[https://play.golang.org/p/FSA8tpFC8S5](https://play.golang.org/p/FSA8tpFC8S5)

Для дальнейшего управления json сериализацией, пакадж предлагает заимплементить интерфейсы Marshaler, Unmarshaler:

```go
type Marshaler interface {
    MarshalJSON() ([]byte, error)
}

type Unmarshaler interface {
    UnmarshalJSON([]byte) error
}
```

Т.е. для структуры, которую мы будем сериализовать нужно добавить метод `MarshalJSON() ([]byte, error)`, либо, если десериализуем \- `UnmarshalJSON([]byte) error`.

Например, реальный случай, когда из внешнего апи приходит json, где одно и тоже поле может быть как строка, так и массив строк:

```js
{
  ... // другие поля
  "languages": "russian"
}

```

и

```js
{
  ... // другие поля
  "languages": ["russian", "english"]
}

```

Но при десериализации, хотелось бы получать всегда массив строк, даже, когда приходит одна строка:

```go
type APIResponse struct {
	Languages []string `json:"languages"`
}
```

Как вариант, можно “обычные” поля выделить в некую “base part”, “languages” \- в отдельную структуру и связать все через [embendding](https://golang.org/doc/effective_go.html#embedding):

```go
type APIResponseBase struct {
	FieldA int64
	FieldB string
}

type APIResponse struct {
	*APIResponseBase
	Languages []string
}

func (p *APIResponse) UnmarshalJSON(data []byte) error {
	basePart := &APIResponseBase{}
	if err := json.Unmarshal(data, basePart); err != nil {
		return err
	}

	withArray := &struct {
		Languages []string
	}{}
	if err := json.Unmarshal(data, withArray); err == nil {
		p.APIResponseBase = basePart
		p.Languages = withArray.Languages
		return nil
	}

	withStr := &struct {
		Languages string
	}{}
	if err := json.Unmarshal(data, withStr); err == nil {
		p.APIResponseBase = basePart
		p.Languages = []string{withStr.Languages}
		return nil
	} else {
		return err
	}
}
```

[https://play.golang.org/p/H\-BkWZzEFq\-](https://play.golang.org/p/H-BkWZzEFq-)

Или другой вариант \- определить новый тип `type ArrayStr []string`, написать для него `UnmarshalJSON` и дальше использовать в `APIResponse` для поля `Languages`:

```go
type ArrayStr []string

func (p *ArrayStr) UnmarshalJSON(data []byte) error {
	var (
		resArr []string
		resStr string
	)

	if err := json.Unmarshal(data, &resArr); err == nil {
		*p = ArrayStr(resArr)
		return nil
	}

	if err := json.Unmarshal(data, &resStr); err != nil {
		return err
	}

	*p = ArrayStr([]string{resStr})
	return nil

}

type APIResponse struct {
	Languages ArrayStr
}
```

[https://play.golang.org/p/AorXMKVo1Lm](https://play.golang.org/p/AorXMKVo1Lm)

Ссылки:

*   [https://golang.org/pkg/encoding/json/](https://golang.org/pkg/encoding/json/)
*   [https://blog.golang.org/json\-and\-go](https://blog.golang.org/json-and-go)
*   [https://stackoverflow.com/a/38757780](https://stackoverflow.com/a/38757780)
*   [https://engineering.bitnami.com/articles/dealing\-with\-json\-with\-non\-homogeneous\-types\-in\-go.html](https://engineering.bitnami.com/articles/dealing-with-json-with-non-homogeneous-types-in-go.html)
