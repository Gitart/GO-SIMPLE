// MongoDB utils
package for_mgo

import (
	"errors"
	"fmt"
	"mgo"
	"strings"
)

const (
	sc_Error      = ", Error: "
	sc_ErrURL     = "URL is empty"
	sc_ErrConnect = "Connect to URL: '%s' is failed"
	sc_ErrDB      = "Database name is empty"
	sc_ErrOpenDB  = "Open Database: '%s' to URL: '%s' is failed"
	sc_ErrNotDB   = "Database is not exists"
	sc_ReadTables = "Read list collections names"
)

//строка соединения (URL) =//[mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
//примеры:
// = mongodb://myuser:mypass@localhost:40001,otherhost:40001/mydb
// = localhost
// = localhost/test

//возвращает обьекты сессии и БД
//>> для обьекта сессии в конце обязательно вызвать defer s.Close()
//наприер:
func OpenDB(url, db_name string) (s *mgo.Session, db *mgo.Database, err error) {
	url = strings.TrimSpace(url)
	if url == "" {
		return s, db, errors.New(sc_ErrURL)
	}
	_s, err := mgo.Dial(url)
	if err != nil {
		return s, db, errors.New(fmt.Sprintf(sc_ErrConnect, url) +
			sc_Error + err.Error())
	}
	db_name = strings.TrimSpace(db_name)
	if db_name == "" {
		_s.Close()
		return s, db, errors.New(sc_ErrDB)
	}
	L, err := _s.DatabaseNames()
	if err != nil {
		_s.Close()
		return s, db, errors.New(fmt.Sprintf(sc_ErrOpenDB, db_name, url) +
			sc_Error + err.Error())
	}
	y := false
	_name := strings.ToLower(db_name)
	for _, n := range L {
		if strings.ToLower(n) == _name {
			y = true
			break
		}
	}
	if y {
		_db := _s.DB(db_name)
		return _s, _db, nil

	} else {
		return s, db, errors.New(fmt.Sprintf(sc_ErrOpenDB, db_name, url) +
			sc_Error + sc_ErrNotDB)

	}
}

//true - коллекция "c_name" имеется в БД "db"
func CollectionExists(db *mgo.Database, c_name string) (r bool, err error) {
	n, err := db.CollectionNames()
	if err != nil {
		return false, errors.New(fmt.Sprint(sc_ReadTables + sc_Error + ". " + err.Error()))
	}
	r = false
	if len(n) > 0 {
		for _, colName := range n {
			if colName == c_name {
				r = true
				break
			}
		}
	}
	return r, nil
}
