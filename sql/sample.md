## Connect to SQL

```golang
package main import ( _ "code.google.com/p/odbc" "database/sql" "log" ) 
func main() { 
// Replace the DSN value with the name of your ODBC data source. 

db, err := sql.Open("odbc", "DSN=SQLSERVER_SAMPLE") 
if err != nil { log.Fatal(err) } 
var ( id int name string ) 
// This is a SQL Server AdventureWorks database query. 
rows, err := db.Query("select departmentid, name from humanresources.department where departmentid = ?", 1) 
if err != nil { log.Fatal(err) } 
defer rows.Close() 

for rows.Next() { 
err := rows.Scan(&id, &name)
if err != nil { log.Fatal(err) } 
log.Println(id, name) } 
err = rows.Err() 
if err != nil { log.Fatal(err) } 
defer db.Close() 
}
```

### Saple 2

```golang
package main import ( "odbc" "log" ) // // CREATE TABLE USERS( ID INTEGER, USERNAME VARCHAR( 50 )); // INSERT INTO USERS VALUES( 1, 'admin' ); // INSERT INTO USERS VALUES( 2, 'sid' ); // INSERT INTO USERS VALUES( 3, 'joe' ); // func main() { conn, err := odbc.Connect("DSN=MYDSN") if err != nil { log.Fatal(err) } stmt, err := conn.Prepare("SELECT * FROM USERS WHERE USERNAME = ? OR USERNAME = ? ") if err != nil { log.Fatal(err) } err = stmt.Execute("admin", "sid" ) if err != nil { log.Fatal(err) } nfields, err := stmt.NumFields(); if err != nil { log.Fatal(err) } println( "Number of fields", nfields ); for i := 0; i < nfields; i ++ { field, err := stmt.FieldMetadata( i + 1 ); if err != nil { log.Fatal(err) } println( "\tField:", i + 1, "Name:", field.Name ); } println( "" ); row, err := stmt.FetchOne() if err != nil { log.Fatal(err) } for row != nil { println( "Row" ) ival := row.GetInt( 0 ) println( "\tField:", 1, "Int:", ival ); sval := row.GetString( 1 ) println( "\tField:", 2, "String:", sval ); row, err = stmt.FetchOne() if err != nil { log.Fatal(err) } } stmt.Close() conn.Close() }
```



