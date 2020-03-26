## Конвертация стринговой даты   
Конвертация стринговую дату одного формата в стринговую дату другого формата



### Пример
https://play.golang.org/p/3NTKaMpqdPA

### Описание
https://yourbasic.org/golang/format-parse-string-time-date-example/

Решение
```golang
package main

import (
	"fmt"
	"time"
	"log"
)

func main() {
     
        date:="2020-03-26T15:20:06+02:00"
         
       // Правильный формат
       layout:=time.RFC3339
     
       // Формат даты в кторый необходимо преобразовать
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


