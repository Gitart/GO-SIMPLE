# MySQL Tutorial: Creating a Table and Inserting Rows

07 March 2021

Welcome to tutorial no. 2 in our MySQL tutorial series. In the [first tutorial](https://golangbot.com/connect-create-db-mysql/), we discussed how to [connect to MySQL and create a database](https://golangbot.com/connect-create-db-mysql/). In this tutorial, we will learn how to create a table and insert records into that table.

#### MySQL Series Index

[Connecting to MySQL and creating a Database](https://golangbot.com/connect-create-db-mysql/)
[Creating a Table and Inserting Rows](https://golangbot.com/mysql-create-table-insert-row/)
[Selecting single and multiple rows](https://golangbot.com/mysql-select-single-multiple-rows/)
Prepared statements - WIP
Updating rows - WIP
Deleting rows - WIP

### Create Table

We will be creating a table named `product` with the fields `product_id`, `product_name`, `product_price`, `created_at` and `updated_at`.

The MySQL query to create this table is provided below,

```
CREATE TABLE IF NOT EXISTS product(product_id int primary key auto_increment, product_name text, product_price int, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)

```

*product\_id* is an auto incremented `int` and it serves as the primary key. The default values of `created_at` and `updated_at` is set as the current timestamp. Now that we have query, let's convert it into Go code and create our table.

The [ExecContext](https://golang.org/pkg/database/sql/#DB.ExecContext) method of the DB package executes any query that doesn't return any rows. In our case, the create table query doesn't return any rows and hence we will use the `ExecContext()` context method to create our table.

Let's be a responsible developer and create a context with a timeout so that the create table query times out in case of any network partition or runtime errors.

```go
query := `CREATE TABLE IF NOT EXISTS product(product_id int primary key auto_increment, product_name text,
        product_price int, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)`

ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
defer cancelfunc()

```

In the above code, we have created a context with a 5 second timeout. Let's go ahead and use this context in the `ExecContext()` method.

```go
res, err := db.ExecContext(ctx, query)
if err != nil {
    log.Printf("Error %s when creating product table", err)
    return err
}

```

We pass the created context and the MySQL query as parameters to the `ExecContext` method and return errors if any. The `db` is the database connection pool that was created in the previous tutorial [https://golangbot.com/connect-create-db-mysql/](https://golangbot.com/connect-create-db-mysql/). Please go through it to understand how to connect to MySQL and create a connection pool.

Now the table is created successfully. The result set returned from the call to `ExecContext()` contains a method that returns the number of rows affected. The create table statement doesn't affect any rows but still, let's check this out by calling the `res.RowsAffected()` method.

```go
rows, err := res.RowsAffected()
    if err != nil {
        log.Printf("Error %s when getting rows affected", err)
        return err
    }
log.Printf("Rows affected when creating table: %d", rows)

```

The above code will print `Rows affected when creating table: 0` since `create table` doesn't affect any rows.

The entire code is provided below.

```go
package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

const (
    username = "root"
    password = "naveenr123"
    hostname = "127.0.0.1:3306"
    dbname   = "ecommerce"
)

func dsn(dbName string) string {
    return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func dbConnection() (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn(""))
    if err != nil {
        log.Printf("Error %s when opening DB\n", err)
        return nil, err
    }
    //defer db.Close()

    ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
    if err != nil {
        log.Printf("Error %s when creating DB\n", err)
        return nil, err
    }
    no, err := res.RowsAffected()
    if err != nil {
        log.Printf("Error %s when fetching rows", err)
        return nil, err
    }
    log.Printf("rows affected %d\n", no)

    db.Close()
    db, err = sql.Open("mysql", dsn(dbname))
    if err != nil {
        log.Printf("Error %s when opening DB", err)
        return nil, err
    }
    //defer db.Close()

    db.SetMaxOpenConns(20)
    db.SetMaxIdleConns(20)
    db.SetConnMaxLifetime(time.Minute * 5)

    ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    err = db.PingContext(ctx)
    if err != nil {
        log.Printf("Errors %s pinging DB", err)
        return nil, err
    }
    log.Printf("Connected to DB %s successfully\n", dbname)
    return db, nil
}

func createProductTable(db *sql.DB) error {
    query := `CREATE TABLE IF NOT EXISTS product(product_id int primary key auto_increment, product_name text,
        product_price int, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)`
    ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    res, err := db.ExecContext(ctx, query)
    if err != nil {
        log.Printf("Error %s when creating product table", err)
        return err
    }
    rows, err := res.RowsAffected()
    if err != nil {
        log.Printf("Error %s when getting rows affected", err)
        return err
    }
    log.Printf("Rows affected when creating table: %d", rows)
    return nil
}

func main() {
    db, err := dbConnection()
    if err != nil {
        log.Printf("Error %s when getting db connection", err)
        return
    }
    defer db.Close()
    log.Printf("Successfully connected to database")
    err = createProductTable(db)
    if err != nil {
        log.Printf("Create product table failed with error %s", err)
        return
    }
}

```

I have included the code to connect to MySQL and create a database from the [previous tutorial](https://golangbot.com/connect-create-db-mysql/) inside the `dbConnection()` function. The only change from the previous tutorial is that the defer statements in line no. 30 and line no. 52 are commented since we do not want the database to be closed immediately after returning from this function.

The `main()` function creates a new DB connection pool in line no. 89 and passes that to the `createProductTable` function in line no. 95. We defer the database close in line no. 93 so that the connection to the DB is closed when the program terminates. Run this program and you can see the following output,

```
2020/10/25 20:30:51 rows affected 1
2020/10/25 20:30:51 Connected to DB ecommerce successfully
2020/10/25 20:30:51 Successfully connected to database
2020/10/25 20:30:51 Rows affected when creating table: 0

```

To verify whether the table has been created successfully, you can run `desc product;` in [MySQL query browser](https://downloads.mysql.com/archives/query/) and you can see that it returns the table schema.

![](https://golangbot.com/content/images/2020/10/describe-product-table.png)

### Insert Row

The next step is to insert rows into the `product` table we just created. The query to insert a row into the product table is provided below,

```
INSERT INTO product(product_name, product_price) VALUES ("iPhone", 800);

```

Let's discuss how to use the above query in Go and insert rows into the table.

Let's first create a product struct to represent our product.

```go
type product struct {
    name      string
    price     int
}

```

The second step is to create a [prepared statement](https://en.wikipedia.org/wiki/Prepared_statement). Prepared statements are used to parametrize a SQL query so that the same query can be run with different arguments efficiently. It also helps prevent [sql injection](https://en.wikipedia.org/wiki/SQL_injection).

In our case, the parameters to the query are `product_name` and `product_price`. The way to create a prepared statement template is to replace the parameters with question mark `?`. The prepared statement template of the following query

```
INSERT INTO product(product_name, product_price) VALUES ("iPhone", 800);

```

is

```
INSERT INTO product(product_name, product_price) VALUES (?, ?);

```

You can see that `"iPhone"` and `800` are replaced with question marks.

```go
func insert(db *sql.DB, p product) error {
    query := "INSERT INTO product(product_name, product_price) VALUES (?, ?)"
    ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    stmt, err := db.PrepareContext(ctx, query)
    if err != nil {
    log.Printf("Error %s when preparing SQL statement", err)
    return err
    }
    defer stmt.Close()
}

```

Line no. 2 of the above code has the prepared statement template. In line no. 5, we create a prepared statement for our insert query using this template. As usual, we use a context with a timeout to handle network errors. The statement should be closed after use. So in the next line we defer the statement close.

The next step is to pass the necessary parameters to the prepared statement and execute it.

```go
res, err := stmt.ExecContext(ctx, p.name, p.price)
if err != nil {
    log.Printf("Error %s when inserting row into products table", err)
    return err
}
rows, err := res.RowsAffected()
if err != nil {
    log.Printf("Error %s when finding rows affected", err)
    return err
}
log.Printf("%d products created ", rows)
return nil

```

The prepared statement expects two arguments namely the product name and the product price. The `ExecContext` method accepts a variadic list of [interface{}](https://golangbot.com/interfaces-part-1/) arguments. The number of variadic arguments passed to it should match the number of question marks `?` in the prepared statement template, else there will be a runtime error `Column count doesn't match value count at row 1 when preparing SQL statement`.

In our case, there are two question marks in the template and hence in the above code snippet, in line no. 1, we pass the two parameters product name and price to the [ExecContext](https://golang.org/pkg/database/sql/#Stmt.ExecContext) method.

The entire `insert` function is provided below.

```go
func insert(db *sql.DB, p product) error {
    query := "INSERT INTO product(product_name, product_price) VALUES (?, ?)"
    ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    stmt, err := db.PrepareContext(ctx, query)
    if err != nil {
        log.Printf("Error %s when preparing SQL statement", err)
        return err
    }
    defer stmt.Close()
    res, err := stmt.ExecContext(ctx, p.name, p.price)
    if err != nil {
        log.Printf("Error %s when inserting row into products table", err)
        return err
    }
    rows, err := res.RowsAffected()
    if err != nil {
        log.Printf("Error %s when finding rows affected", err)
        return err
    }
    log.Printf("%d products created ", rows)
    return nil
}

```

Please add the following code to the end of `main` function to call the `insert` function.

```go
func main() {
...

p := product{
        name:  "iphone",
        price: 950,
    }
err = insert(db, p)
if err != nil {
    log.Printf("Insert product failed with error %s", err)
    return
    }
}

```

If everything goes well, the program will print `1 products created`

You can check that the product has been inserted successfully by running `select * from product;` and you can see the following output in MySQL query browser.

![](https://golangbot.com/content/images/2020/10/MySQL-Insert-Query.png)

### Last Inserted ID

There might be a need to get the last inserted ID of an insert query with auto increment primary key. In our case, the `product_id` is an auto incremented int primary key. We might need the last inserted product id to reference in other tables. Say, we have a supplier table and would like to map suppliers once a new product is created. In this case, fetching the last inserted ID is essential. The `LastInsertId` method of the result set can be used to fetch this ID. Add the following code to the end of the `insert` function before `return nil`.

```go
func insert(db *sql.DB, p product) error {
...

    prdID, err := res.LastInsertId()
    if err != nil {
    log.Printf("Error %s when getting last inserted product",     err)
    return err
    }
    log.Printf("Product with ID %d created", prdID)
    return nil
}

```

When the program is run with the above code added, the line `Product with ID 2 created` will be printed. We can see that the ID of the last inserted product is `2`.

### Insert Multiple Rows

Let's take our insert statement to the next level and try to insert multiple rows using a single query.

The MySQL syntax for inserting multiple rows is provided below

```
insert into product(product_name, product_price) values ("Galaxy","990"),("iPad","500")

```

The different rows to be inserted are separated by commas. Let's see how to achieve this using Go.

The logic is to generate the `("Galaxy","990"),("iPad","500")` after the `values` part of the query dynamically based on the number of products needed to be inserted. In this case, two products namely `Galaxy` and `iPad` have to be inserted. So there is a need to generate a prepared statement template of the following format.

```
insert into product(product_name, product_price) values (?,?),(?,?)

```

Let's write the function to do this right away.

```go
func multipleInsert(db *sql.DB, products []product) error {
    query := "INSERT INTO product(product_name, product_price) VALUES "
    var inserts []string
    var params []interface{}
    for _, v := range products {
        inserts = append(inserts, "(?, ?)")
        params = append(params, v.name, v.price)
    }
}

```

We iterate over the `products` parameter passed to the function and for each product we append `(?, ?)` to the `inserts` [slice](https://golangbot.com/arrays-and-slices/) in line no. 6. In the same `for` loop we append the actual parameters that should substitute the question marks `?` to the `params` slice.
There is one more step remaining before the prepared statement template is ready. The `inserts` slice is of length 2 and it contains `(?, ?)` and `(?, ?)`. These two have to be concatenated with a comma in the middle. The [Join](https://golang.org/pkg/strings/#Join) can be used to do that. It takes a string slice and a separator as parameters and joins the elements of the slice with the separator.

```go
queryVals := strings.Join(inserts, ",")
query = query + queryVals

```

*queryVals* now contains `(?, ?),(?, ?)`. We then concatenate `query` and `queryVals` to generate the final prepared statement template `INSERT INTO product(product_name, product_price) VALUES (?, ?),(?, ?)`.

The remaining code is similar to the single row insert function. Here is the full function.

```go
func multipleInsert(db *sql.DB, products []product) error {
    query := "INSERT INTO product(product_name, product_price) VALUES "
    var inserts []string
    var params []interface{}
    for _, v := range products {
        inserts = append(inserts, "(?, ?)")
        params = append(params, v.name, v.price)
    }
    queryVals := strings.Join(inserts, ",")
    query = query + queryVals
    log.Println("query is", query)
    ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    stmt, err := db.PrepareContext(ctx, query)
    if err != nil {
        log.Printf("Error %s when preparing SQL statement", err)
        return err
    }
    defer stmt.Close()
    res, err := stmt.ExecContext(ctx, params...)
    if err != nil {
        log.Printf("Error %s when inserting row into products table", err)
        return err
    }
    rows, err := res.RowsAffected()
    if err != nil {
        log.Printf("Error %s when finding rows affected", err)
        return err
    }
    log.Printf("%d products created simulatneously", rows)
    return nil
}

```

The one difference which you could see is in line no. 20. We pass the slice as a [variadic argument](https://golangbot.com/variadic-functions/) since `ExecContext` expects a variadic argument. The remaining code is the same.

Add the following lines to the end of the main function to call the `multipleInsert` function.

```go
func main() {
...

p1 := product{
    name:  "Galaxy",
    price: 990,
}
p2 := product{
    name:  "iPad",
    price: 500,
}
err = multipleInsert(db, []product{p1, p2})
if err != nil {
    log.Printf("Multiple insert failed with error %s", err)
    return
}

```

On running the program you can see

```
query is INSERT INTO product(product_name, product_price) VALUES (?, ?),(?, ?)
2 products created simultaneously

```

printed. On querying the table, it can be confirmed that two products are inserted.

The entire code is available in at [https://github.com/golangbot/mysqltutorial/blob/master/insert/main.go](https://github.com/golangbot/mysqltutorial/blob/master/insert/main.go)

This brings us to an end of this tutorial. Please leave your comments and feedback.

If you would like to advertise on this website, hire me, or if you have any other development requirements please email to *naveen\[at\]golangbot\[dot\]com*.

**Previous tutorial - [Connecting to](https://golangbot.com/connect-create-db-mysql/)**
