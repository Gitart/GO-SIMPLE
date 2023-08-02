// Цель - обновлять несколько таблиц JSON
// с подчиненными вложенными записями

package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm/logger"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DbPassword string = "MainDBuser123$"
	DbName            = "work"
	DbUser            = "root"
	DB         *gorm.DB

	Data = `{
               "title"    : "New compa222ny",
               "emploies" : [
                               {"id2":11, "name":"Оля1"},
                               {"id2":12, "name":"Конста1нт"},
                               {"i2d":13, "name":"Рлялч1я"}
                 ]
               } `
)

func init() {
	DB = DBC()
}

// DBC Db connect
// If get error : gorm.Open(mysql.Open(dsn), gorm.Config{}) password yes
// GRANT ALL PRIVILEGES ON mydb.* TO 'root'@'%' IDENTIFIED BY 'Gerda3000';
// ALTER USER 'root'@'localhost' IDENTIFIED BY 'Gerda3000';
func DBC() *gorm.DB {

	// Change password to
	// database depended on from OS
	cfg := &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent), // Off - отключение сообщения о медленном выполнении и других ошибках,
		SkipDefaultTransaction:                   true,                                  // Skip Transaction
		PrepareStmt:                              true,                                  // Prepare
		AllowGlobalUpdate:                        true,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: true,
		DisableNestedTransaction:                 true,
	}

	// Connect string
	dsn := DbUser + ":" + DbPassword + "@tcp(127.0.0.1:3306)/" + DbName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), cfg)

	if err != nil {
		fmt.Println("  DATABASE: DESCRIPTION : ", err.Error())
	} else {

		// Db Config
		// https://stackoverflow.com/questions/63517123/how-to-set-sql-connection-config-on-gorm-v2
		// Setting connections
		// Необходимо понаблюдать некоторое время (Need watching some times)
		dbConfig, _ := db.DB()
		dbConfig.SetMaxIdleConns(150)
		dbConfig.SetMaxOpenConns(200)
		dbConfig.SetConnMaxLifetime(time.Hour)
	}

	return db
}

type Companies struct {
	Id       int64      `json:"id"`
	Title    string     `json:"title"`
	Emploies []Emploies `gorm:"polymorphic:Company; polymorphicValue:company"`
}

// "polymorphic:Company - говорит о том что в подчиненной таблице должно быть поле - CompanyID !
// CompanyType - тоже ожидается
// polymorphicValue:master - указывает что писать в поле тип в связной таблице
// если его не будет - то вставляется имя родительской таблицы

type Emploies struct {
	Id          int64 `json:"id"`
	CompanyID   int   `json:"company_id"`
	CompanyType string
	Name        string `json:"name"`
}

// Main
func main() {
	comp := Companies{}
	json.Unmarshal([]byte(Data), &comp)

	//comp.Id = time.Now().Unix()
	//comp.Id = 1674936765

	//comp := Companies{
	//	Id:    time.Now().Unix(),
	//	Title: "sss",
	//	Emploies: []Emploies{
	//		{Name: "Stepanov"},
	//		{Name: "Пастушенко"},
	//		{Name: "Степаненко"},
	//	},
	//}

	sess := DB.Session(&gorm.Session{FullSaveAssociations: true})

	//er := DB.Debug().Updates(&comp)
	// er := DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(&comp)

	// Добавляет
	//  Но - если уже есть такая запись в родительской таблице то ее меняет а подчиненные нет!

	// sess.Create(&comp)

	// Обновляет и связанняе
	// Здесь меняет и в родительской таблице и в подчиненной
	// Кнонечно если в подчиненной таблице будут указаны ИД записей которые надо поменять !!!!!
	sess.Save(&comp)

	//Model(&Companies{}).
	//Association("IdCommpany").
	//Append(comp)

	//if er.Error != nil {
	//	fmt.Println("ERROR:", er.Error)
	//}

	fmt.Println("ok")
}
