
### Получить поле в опредленном элементе интерфейса

```golang
package main
import "fmt"
type Mst map[string]interface{}     // map - string - interface
type Mif []interface{}                       // interface


func main() {

	 OSS:= Mif{
		Mst{"id": 1, "Title": "Дата последненго обновления", "Value": "Y"},
		Mst{"id": 2, "Title": "Количество обработанных строк", "Value": 0},
		Mst{"id": 3, "Title": "Дата рабочей версии", "Value": "sss"},
		Mst{"id": 4, "Title": "Номер версии", "Value": "0.122.22"},
		Mst{"id": 5, "Title": "Путь к файлу", "Value": "C:/log.txt"},
		Mst{"id": 6, "Title": "Наименование организации", "Value": "Morion"},
		Mst{"id": 7, "Title": "Код", "Value": "M-00120-1273783-20141110"},
		Mst{"id": 8, "Title": "Имя администратора", "Value": "Admin"},
		Mst{"id": 9, "Title": "Пароль для администратора", "Value": "AAAAA00012S"},
		Mst{"id": 10, "Title": "Адрес", "Value": "Киев "},
		Mst{"id": 11, "Title": "Контактная информация", "Value": ""},
		Mst{"id": 13, "Title": "Уникальный код", "Value": "12345-12345-00988"},
		Mst{"id": 14, "Title": "Обработанный ID", "Value": 1},
		Mst{"id": 15, "Title": "Дата последнего визита", "Value": ""},
	}
	
	
	
	// Как получить Title из Id:5 ? 
	// fmt.Printf("%s %d\n", OSS["a"], m["b"])

	// Элементарно (типа) :)
	fmt.Printf("%v\n", OSS[4].(Mst)["Title"])
}
```
