## Connect to SQL

```golang
package main import ( _ "code.google.com/p/odbc" "database/sql" "log" ) func main() { // Replace the DSN value with the name of your ODBC data source. db, err := sql.Open("odbc", "DSN=SQLSERVER_SAMPLE") if err != nil { log.Fatal(err) } var ( id int name string ) // This is a SQL Server AdventureWorks database query. rows, err := db.Query("select departmentid, name from humanresources.department where departmentid = ?", 1) if err != nil { log.Fatal(err) } defer rows.Close() for rows.Next() { err := rows.Scan(&id, &name) if err != nil { log.Fatal(err) } log.Println(id, name) } err = rows.Err() if err != nil { log.Fatal(err) } defer db.Close() }
```
