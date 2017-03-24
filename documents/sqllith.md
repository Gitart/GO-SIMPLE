    	fmt.Println(affect)
    
    	// запрос
    	rows, err := db.Query("SELECT * FROM userinfo")
    	checkErr(err)
    
    	for rows.Next() {
    		var uid int
    		var username string
    		var department string
    		var created string
    		err = rows.Scan(&uid, &username, &department, &created)
    		checkErr(err)
    		fmt.Println(uid)
    		fmt.Println(username)
    		fmt.Println(department)
    		fmt.Println(created)
    	}
    
    	// удаление
    	stmt, err = db.Prepare("delete from userinfo where uid=?")
    	checkErr(err)
    
    	res, err = stmt.Exec(id)
    	checkErr(err)
    
    	affect, err = res.RowsAffected()
    	checkErr(err)
    
    	fmt.Println(affect)
    
    	db.Close()
    
    }
    
    func checkErr(err error) {
    	if err != nil {
    		panic(err)
    	}
    }


Вы наверняка заметили, что код очень похож на пример из предыдущего раздела, и мы изменили только имя зарегистрированного драйвера и вызвали `sql.Open` для соединения с SQLite по-другому.

Ниже дана ссылка на инструмент управления SQLite: [http://sqliteadmin.orbmu2k.de/](http://sqliteadmin.orbmu2k.de/)

## Ссылки

- [Содержание](preface.md)
- Предыдущий раздел: [MySQL](05.2.md)
- Следующий раздел: [PostgreSQL](05.4.md)
