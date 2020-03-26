## Конвертация стринговой даты   
Конвертация стринговую дату одного формата в стринговую дату другого формата

### Решение
```golang
package main

import (
	"fmt"
	"time"
	"log"
)

func main() {
        
	// Дата которая поступает на вход
        date:="2020-03-26T15:20:06+02:00"
         
       // Правильный формат который ожидается на вход
       // и которому должна соответствовать входящая дата
       layout:=time.RFC3339
     
       // Формат даты в который необходимо преобразовать
       format := "2006-01-02 15:04:05"         
       t, err := time.Parse(layout , date)

       // Если ошибка парсинга возвращаем сегодняшнюю дату
       if err!=nil{
          log.Println("JOURNAL: Ошибка парсинга даты.")
       }

       Finish := t.Format(format)
       fmt.Println(Finish)
}


```


## Link

### Пример
https://play.golang.org/p/3NTKaMpqdPA

### Описание
https://yourbasic.org/golang/format-parse-string-time-date-example/

### Дискуссия
https://stackoverflow.com/questions/37937794/string-date-to-date
