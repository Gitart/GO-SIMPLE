# Метод цепочки
Одним из очень полезных свойств методов является возможность связывать их вместе, сохраняя при этом ваш код в чистоте. 
Давайте рассмотрим пример установки некоторых атрибутов Personиспользования цепочки:

```go
type Person struct {
	Name string
	Age  int
}

func (p *Person) withName(name string) *Person {
	p.Name = name
	return p
}

func (p *Person) withAge(age int) *Person {
	p.Age = age
	return p
}

func main() {
	p := &Person{}
	p = p.withName("John").withAge(21)

  fmt.Println(*p)
  // {John 21}
}
```

### Если бы мы использовали функции для одной и той же вещи, это выглядело бы довольно ужасно:

```go
p = withName(withAge(p, 18), "John")
```
