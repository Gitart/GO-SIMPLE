# Шпаргалка по работе с JSON в Golang

[feeeper](https://ashirobokov.wordpress.com/author/feeeper/ "Записи feeeper") [Cheat sheets](https://ashirobokov.wordpress.com/category/cheat-sheets/) 22 сентября, 2016 1 Minute

Парсинг JSON — одна из наиболее частых задач: в JSON приходят данные в REST API, конфигурационные файлы часто оформляются в виде JSON и пр.

Go предоставляют довольно удобные механизмы для этих целей расположенные в пакете `"encoding/json"` включающий в себя необходимые методы.

# Преобразование в JSON
## Преобразование простых типов (`bool`, `string`, `int`)

```golang
boolVar, \_ := json.Marshal(true)
fmt.Println(string(boolVar))
// true
intVar, \_ := json.Marshal(1)
fmt.Println(string(intVar))
// 1

fltVar, \_ := json.Marshal(2.34)
fmt.Println(string(fltVar))
// 2.34

strVar, \_ := json.Marshal("something")
fmt.Println(string(strVar))
// "something"
```


## Преобразование массивов, слайсов и словарей
```golang
sliceVar1 := \[\]string{"John", "Andrew", "Robert"}
sliceVar2, \_ := json.Marshal(sliceVar)
fmt.Println(string(sliceVar2))
// \["John", "Andrew", "Robert"\]

mapVar1 := map\[string\]string{"John": "Accepted", "Andrew": "Waiting", "Robert": "Cancelled"}
mapVar2, \_ := json.Marshal(mapVar1)
fmt.Println(string(mapVar2))
// {"John": "Accepted", "Andrew": "Waiting", "Robert": "Cancelled"}
```

## Преобразование пользовательских типов данных
```golang
type User struct {
    FirstName string
    LastName string
    Books: \[\]string
}
userVar1 := &User{
    FirstName: "John",
    LastName: "Smith",
    Books: \[\]string{ "The Art of Programming", "Golang for Dummies" }}
userVar2, \_ := json.Marshal(userVar1)
fmt.Println(string(userVar2))
// {"FirstName":"John","LastName":"Smith","Books":\["The Art of Programming","Golang for Dummies"\]}

По\-умолчанию ключи в JSON будут соответствовать именам свойств структуры (`FirstName`, `LastName`, `Books`), если требуется изменить это поведение, то необходимо добавить теги:

type User2 struct {
    FirstName string \`json:"name"\` // свойство FirstName будет преобразовано в ключ "name"
    LastName string \`json:"lastname"\` // свойство LastName будет преобразовано в ключ "lastname"
    Books \[\]string \`json:"ordered\_books"\` // свойство Books будет преобразовано в ключ "ordered\_books"
}
userVar3 := &User2{
    FirstName: "John",
    LastName: "Smith",
    Books: \[\]string{ "The Art of Programming", "Golang for Dummies" }}
userVar4, \_ := json.Marshal(userVar3)
fmt.Println(string(userVar4))
// {"name":"John","lastname":"Smith","ordered\_books":\["The Art of Programming","Golang for Dummies"\]}
```

# Преобразование из JSON
## Стандартные типы

```golang
byt := \[\]byte(\`{"num":6.13,"strs":\["a","b"\]}\`)
var dat map\[string\]interface{}

if err := json.Unmarshal(byt, &dat); err != nil {
    panic(err)
}
fmt.Println(dat)
```


В `dat` получим объект типа `map[string]interface{}`. Для того, чтобы работать с «внутренностями» этого объекта придётся немного поcastовать:
```golang
num := dat\["num"\].(float64) // для того, чтобы получить из свойства num число
fmt.Println(num) strs := dat\["strs"\].(\[\]interface{}) // для того, чтобы получить массив интерфейсов...
str1 := strs\[0\].(string) // ... и потом получить из него строку
fmt.Println(str1)
```

## Пользовательские типы
Пользовательские типы так же просто можно получить из JSON:

```golang
user := User1{}
userJson := "{\\"FirstName\\":\\"John\\",\\"LastName\\":\\"Smith\\",\\"Books\\":\[\\"The Art of Programming\\",\\"Golang for Dummies\\"\]}"
bytes := \[\]byte(userJson)
json.Unmarshal(bytes, &user)
fmt.Println(user.FirstName, user.LastName, user.Books)
// John Smith \[The Art of Programming Golang for Dummies\]
```

Но что же делать, если ключи в JSON отличаются от свойств в структуре или необходимо получать не все данные из JSON? Для этого необходимо воспользоваться знакомыми по предыдущим примерам тегами `json`:

```golang
user2 := User2{}
userJson2 := "{\\"name\\":\\"John\\",\\"lastname\\":\\"Smith\\",\\"ordered\_books\\":\[\\"The Art of Programming\\",\\"Golang for Dummies\\"\]}"
bytes2 := \[\]byte(userJson2)
json.Unmarshal(bytes2, &user2)
fmt.Println(user2.FirstName, user2.LastName, user2.Books)
// John Smith \[The Art of Programming Golang for Dummies\]
```

Получили результат аналогичный предыдущему.

На этом основы работы с JSON можно и завершить. Поиграться с преобразование можно

Link :

https://play.golang.org/p/dqn5UdqFfJt
