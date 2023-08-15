# Out JSON file to relation tables

![image](https://github.com/Gitart/GO-SIMPLE/assets/3950155/1d742da1-d760-482c-ad98-a8b46c9af45a)

> [!WARNING]
> Critical content demanding immediate user attention due to potential risks.


## Table relations
```sql
CREATE TABLE `m_acts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `num` varchar(45) DEFAULT NULL,
  `company` varchar(45) DEFAULT NULL,
  `created_at` varchar(45) DEFAULT NULL,
  `total` float DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;

CREATE TABLE `m_acts_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_acts` int DEFAULT NULL,
  `title` varchar(45) DEFAULT NULL,
  `summ` float DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fg_acts_idx` (`id_acts`),
  CONSTRAINT `fg_acts` FOREIGN KEY (`id_acts`) REFERENCES `m_acts` (`id`)
) ENGINE=InnoDB;
```

## GO 
```go
package project

import (
	"boiler/controllers/db"
	"github.com/labstack/echo/v4"
	"log"
)

type MActs struct {
	Id        int           `json:"id"`
	Num       string        `json:"num"`
	Company   string        `json:"company"`
	CreatedAt string        `json:"created_at"`
	Total     float64       `json:"total"`
	Items     []*MActsItems `json:"items" gorm:"foreignKey:IdActs"`
}

type MActsItems struct {
	Id     int     `json:"id"`
	IdActs int     `json:"id_acts"`
	Title  string  `json:"title"`
	Sum    float64 `json:"sum"`
}

// Get JSON Acts
func GetAct(e echo.Context) error {

	mActsList := []MActs{}

	result := db.DB.Preload("Items").Find(&mActsList)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	jsonData, err := json.Marshal(mActsList)
	if err != nil {
		log.Fatal("Error marshaling JSON:", err)
	}
	
	fmt.Println(string(jsonData))

	return e.JSON(200, mActsList)
}
```



classDiagram
direction BT
class areas {
   varchar(45) title
   int id
}
class node3 {
   varchar(45) title
   int location_id
   varchar(100) location
   int id
}
class brigades_items {
   int user_id
   int brigad_id
   varchar(100) brigad
   varchar(100) user
   int id
}
class districts {
   varchar(100) title
   int id
}
class locations {
   varchar(45) code
   varchar(100) Uuid  /* Guid - Для формирования ссылки на котельню */
   datetime created_at  /* Cоздан запись */
   datetime updated_at
   varchar(100) title  /* Нименование котельни */
   int area_id
   varchar(45) area  /* Область */
   int district_id
   varchar(100) district
   varchar(45) city_type
   varchar(45) city  /* Город */
   varchar(45) street_type
   varchar(100) street  /* улица */
   varchar(45) build  /* Дом */
   varchar(45) kv
   int contragent_id
   varchar(150) contragent
   varchar(45) object
   int object_id
   varchar(100) otg
   int otg_id
   int manager_id  /* ВЕД - Ид менеджера */
   int operator_id  /* Ид пользователя */
   int block
   varchar(150) remark
   int id
}
class node6 {
   int contragent_id
   varchar(150) contragent
   varchar(45) num  /* Автоматически код для системы */
   datetime created_at
   varchar(150) title  /* Название для системы */
   varchar(200) long  /* Длинное название для документов */
   int location_id
   varchar(150) location
   int boiler_id
   json contacts
   varchar(45) note
   varchar(45) status
   int id
}
class otgs {
   varchar(100) title
   varchar(100) type_id
   varchar(45) email
   varchar(100) center
   varchar(45) district
   int id
}
class roles {
   varchar(45) code
   varchar(50) title
   varchar(50) description
   varchar(45) objects
   bigint id
}
class users {
   varchar(50) uuid  /* UUID */
   varchar(45) name  /* Имя Фамилия */
   varchar(45) first_name  /* Имя */
   varchar(45) middle_name  /* Фамилия */
   varchar(45) last_name  /* Отчество */
   timestamp created_at  /* Дата создания */
   timestamp updated_at  /* Дата обновления */
   timestamp deleted_at  /* Дата удаления */
   int boiler_id
   int role_id  /* Роль ид */
   varchar(45) role
   varchar(50) email  /* Почта */
   varchar(50) login
   varchar(150) password  /* Пароль в зашифрованном виде SHA256 */
   int is_admin  /* Признак администратора системы */
   varchar(20) mobile  /* Мобильный телефон */
   varchar(45) telephone_ext
   varchar(150) remark  /* Примечание */
   int present  /* Статус присутствия =1 Отсутствие =-1 */
   varchar(100) status  /* Статус входа Ок - если удачная регистрация Err - при плохой */
   varchar(45) active  /* Активный пользователь = A Заблокирован = B */
   int department_id
   bigint id  /* Id */
}

node3  -->  locations : location_id:id
brigades_items  -->  users : user_id:id
locations  -->  areas : area_id:id
locations  -->  districts : district_id:id
locations  -->  node6 : object_id:id
locations  -->  otgs : otg_id:id
node6  -->  locations : location_id:id
users  -->  roles : role_id:id

