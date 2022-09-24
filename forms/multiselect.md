## Мультивыбор в форме и обработка

## From
```
<select name="cars" id="cars" multiple>
  <option value="1">Volvo</option>
  <option value="2">Saab</option>
  <option value="3">Opel</option>
  <option value="4">Audi</option>
  <option value="5">Alfa Romeo</option>
  <option value="6">BMW</option>
  <option value="7">KIA</option>
</select>
```                

## Обработка поля с мультивыбором


Этот метод не работает с мультивыбором
```
c.FormValue("title")
```

### Этот метод работает
```go
// Read fields in cicle
c.Request().ParseForm()
for key, value := range c.Request().PostForm {
    fmt.Println(key,value)
}

// Использование для мультивыбора
ccc := c.Request().PostForm["remark"]
        
fmt.Println("Remark post    : ", ccc[0])
fmt.Println("Remark post nom: ", ccc)

cars := c.Request().PostForm["cars"]
        
fmt.Println("Cars : ", c.FormValue("cars"))
        
for _, rncar:=range(cars){
  fmt.Println("!!!!!!!!!! -----------", rncar)
}
```
