```
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

```

1.  Использовать omitempty, если значение пустое.

```
type User struct {
    Name string `json:"name,omtiempty"`
    Age  int    `json:"age,omitempty"`
}

```

```
User{"Alex", 0} --> {"name":"Alex"}
{"name":"Alex", 21} --> User{"Alex", 21}

```

2.  Использовать другую структуру скрыв нужные поля тем или иным способом

```
type UserResponse struct {
    Name string `json:"name"`
    age  int                  // ну или Age int `json:"-"`
}

```

3.  Использовать свой метод MarshalJSON

```
func (u *User) MarshalJSON() ([]byte, error) {
    type userResponse struct {
        Name string
    }
    var reply userResponse
    reply.Name = u.Name
    return json.Marshal(&reply)
}
```
