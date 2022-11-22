package main

import (
    "database/sql"
    "fmt"
    "io/ioutil"
    "log"
    "strings"
    "golang.org/x/text/encoding/charmap"
    "golang.org/x/text/transform"
    _ "code.google.com/p/odbc"
)
var (
    name_otdel string
    name_utf   string
    query      string
)
func main() {
    db, err := sql.Open("odbc", "DSN=DBS0")
    if err != nil {
        fmt.Println("Error in connect DB")
        log.Fatal(err)
    }
    query = "select t.DepartmentNAME from dbo.oms_Department t where t.rf_LPUID = 1000"
    rows, err := db.Query(query)
    if err != nil {
        log.Fatal(err)
    }
    for rows.Next() {
        if err := rows.Scan(&name_otdel); err != nil {
            log.Fatal(err)
        }
        sr := strings.NewReader(name_otdel)
        tr := transform.NewReader(sr, charmap.Windows1251.NewDecoder())
        buf, err := ioutil.ReadAll(tr)
        if err != nil {
            log.Fatal(err)
        }
        name_utf = string(buf)
        fmt.Println(name_utf)
    }
    defer rows.Close()
}
