# MySQL Tutorial: Connecting to MySQL and Creating a DB using Go

05 October 2020

Welcome to tutorial no. 1 in our MySQL tutorial series. In this tutorial, we will connect to MySQL and create a database. We will also ping the DB to ensure the connection is established properly.

#### MySQL Series Index

[Connecting to MySQL and creating a Database](https://golangbot.com/connect-create-db-mysql/)
[Creating a Table and Inserting Rows](https://golangbot.com/mysql-create-table-insert-row/)
[Selecting single and multiple rows](https://golangbot.com/mysql-select-single-multiple-rows/)
Prepared statements - WIP
Updating rows - WIP
Deleting rows - WIP

### Importing the MySQL driver

The first step in creating the MySQL database is to download the MySQL driver [package](https://golangbot.com/go-packages/) and import it into our application.

Let's create a folder for our app and then download the MySQL package.

I have created a folder in the `Documents` directory. Please feel free to create it wherever you like.

```
mkdir ~/Documents/mysqltutorial
cd ~/Documents/mysqltutorial

```

After creating the directory, let's initialize a go module for the project.

```
go mod init github.com/golangbot/mysqltutorial

```

The above command initializes a module named `github.com/golangbot/mysqltutorial`

The next step is to download the MySql driver. Run the following command to download the MySQL driver package.

```
go get github.com/go-sql-driver/mysql

```

Let's write a program to import the MySQL driver we just downloaded.

Create a file named `main.go` with the following contents.

```go
package main

import (
    "database/sql"

    _ "github.com/go-sql-driver/mysql"
)

```

### Use of blank identifier \_ when importing the driver

The `"database/sql"` package provides generic [interfaces](https://golangbot.com/interfaces-part-1/) for accessing the MySQL database. It contains the types needed to manage the MySQL DB.

In the next line, we import `_ "github.com/go-sql-driver/mysql"` prefixed with an underscore (called as a blank identifier). What does this mean? This means we are importing the MySQL driver package for its side effect and we will not use it explicitly anywhere in our code. **When a package is imported prefixed with a blank identifier, the init function of the package will be called. Also, the Go compiler will not complain if the package is not used anywhere in the code.**

That's all fine, but why is this needed?

The reason is any SQL driver must be registered by calling the [Register](https://golang.org/pkg/database/sql/#Register) function before it can be used. If we take a look at the source code of the MySQL driver, in line [https://github.com/go-sql-driver/mysql/blob/b66d043e6c8986ca01241b990326db395f9c0afd/driver.go#L83](https://github.com/go-sql-driver/mysql/blob/b66d043e6c8986ca01241b990326db395f9c0afd/driver.go#L83) we can see the following `init` function

```go
func init() {
    sql.Register("mysql", &MySQLDriver{})
}

```

The above function registers the SQL driver named `mysql`. When we import the package prefixed with the blank identifier `_ "github.com/go-sql-driver/mysql"`, this `init` function is called and the driver is available for use. Perfect 😃. Just what we wanted.

### Connecting and Creating the Database

Now that we have registered the driver successfully, the next step is to connect to MySQL and create the database.

Let's define [constants](https://golangbot.com/constants/) for our DB credentials.

```go
const (
    username = "root"
    password = "password"
    hostname = "127.0.0.1:3306"
    dbname   = "ecommerce"
)

```

Please replace the above values with your credentials.

The DB can be opened by using [Open](https://golang.org/pkg/database/sql/#Open) function of the sql package. This function takes two parameters, the driver name, and the data source name(DSN). As we have already discussed, the driver name is `mysql`. The DSN is of the following format

```
username:password@protocol(address)/dbname?param=value

```

Let's write a small function that will return us this DSN when the database name is passed as a parameter.

```go
func dsn(dbName string) string {
    return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

```

The above [function](https://golangbot.com/functions/) returns a DSN for the `dbName` passed. The `dbName` is optional and it can be empty. For example, if `ecommerce` is passed, it will return `root:password@tcp(127.0.0.1:3306)/ecommerce`

Since we are actually creating the DB here and do not want to connect an existing DB, an empty `dbName` will be passed to the `dsn` function.

```go
func main() {
    db, err := sql.Open("mysql", dsn(""))
    if err != nil {
        log.Printf("Error %s when opening DB\n", err)
        return
    }
    defer db.Close()
}

```

Please ensure that the user has access rights to create the DB. The above lines of code open and return a connection to the database. The database connection is closed when the function returns using [defer](https://golangbot.com/defer/).

After establishing a connection to the DB, the next step is to create the DB. The following code does that.

```go
ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
defer cancelfunc()
res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
if err != nil {
    log.Printf("Error %s when creating DB\n", err)
    return
}
no, err := res.RowsAffected()
if err != nil {
    log.Printf("Error %s when fetching rows", err)
    return
}
log.Printf("rows affected %d\n", no)

```

After opening the database, we use the [ExecContext](https://golang.org/pkg/database/sql/#DB.ExecContext) method to create the database. This [method](https://golangbot.com/methods/) is used to execute a query without returning any rows. Since we are creating a DB, it returns no rows, and `ExecContext` can be used to create the database. Being a responsible developer, we pass a context with a timeout of 5 seconds to ensure that the control doesn't get stuck when creating the DB in case there is any network error or any other error in the DB. `cancelfunc` is only needed when we want to cancel the context before it times out. There is no use of it here, hence we just defer the `cancelfunc` call.

The `ExecContext` call returns a [result](https://golang.org/pkg/database/sql/#Result) type and an error. We can check the number of rows affected by the query by calling the `RowsAffected()` method. The above code creates a database named `ecommerce`.

### Understanding Connection Pool

The next step after creating the DB is to connect to it and start executing queries. In other programming languages, you might do this by running the `use ecommerce` command to select the database and start executing queries. This can be done in Go by using the code `db.ExecContext("USE ecommerce")`.

While this might seem to be a logical way to proceed, this leads to unexpected runtime errors in Go. Let's understand the reason behind this.

When we first executed `sql.Open("mysql", dsn(""))`, the [DB](https://golang.org/pkg/database/sql/#DB) returned is actually a pool of underlying DB connections. The sql package takes care of maintaining the pool, creating and freeing connections automatically. This DB is also safe to be concurrently accessed by multiple [Goroutines](https://golangbot.com/goroutines/).

Since `DB` is a connection pool, if we execute `use ecommerce` on `DB`, it will be run on only one of the DB connections in the pool. When we execute another query on `DB`, we might end up running the query on some other connection in the pool on which `use ecommerce` was not executed. This will lead to the error `Error Code: 1046. No database selected`.

The solution is simple. We close the existing connection to the DB which we created without specifying a DB name and open a new connection with the DB name `ecommerce` which was just created.

```go
db.Close()
db, err = sql.Open("mysql", dsn(dbname))
if err != nil {
    log.Printf("Error %s when opening DB", err)
    return
}
defer db.Close()

```

In the above lines, we close the existing connection and open a new connection to the DB. This time we specify the DB name `ecommerce` in line no. 2 when opening a connection to the database. Now we have a connection pool connected to the `ecommerce` DB 😃.

### Connection Pool Options

There are few important connection pool options to be set to ensure that network partitions and other runtime errors that may occur with our DB connections are handled properly.

#### SetMaxOpenConns

This option is used to set the maximum number of open connections that are allowed from our application. It's better to set this to ensure that our application doesn't utilize all available connections to MySQL and starve other applications.

The maximum connections for a MySQL Server can be determined by running the folllowing query

```
show variables like 'max_connections';

```

It returns the following output in my case. `151` is the default maximum connections allowed. You can change it to a different value according to your requirement.
![](https://golangbot.com/content/images/2020/10/mysql-max-connections.png)

*151* is the maximum connections allowed for this entire MySQL server which may include other applications accessing the same DB and also access to other databases if any exist.

Ensure that you set a value lower than `max_connections` so that other applications and databases are not starved. I am using `20`. Please feel free to change it according to your requirement.

```go
db.SetMaxOpenConns(20)

```

#### SetMaxIdleConns

This option limits the maximum idle connections. The number of idle connections in the connection pool is controlled by this setting.

```go
db.SetMaxIdleConns(20)

```

#### SetConnMaxLifetime

It's quite common for connections to become unusable because of a number of reasons. For instance, there might be a firewall or middleware that terminates idle connections. This option ensures that the driver closes the idle connection properly before it is terminated by a firewall or middleware.

```go
db.SetConnMaxLifetime(time.Minute * 5)

```

Please feel free to change the above options based on your requirement.

### Pinging the DB

The `Open` function call doesn't make an actual connection to the DB. It just validates whether the DSN is correct. The [PingContext()](https://golang.org/pkg/database/sql/#DB.PingContext) method must be called to verify the actual connection to the database. It pings the DB and verifies the connection.

```go
ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
defer cancelfunc()
err = db.PingContext(ctx)
if err != nil {
    log.Printf("Errors %s pinging DB", err)
    return
}
log.Printf("Connected to DB %s successfully\n", dbname)

```

We create a context with a 5 second timeout to ensure that the control doesn't get stuck when pinging the DB in case there is a network error or any other error.

The full code is provided below.

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
    password = "password"
    hostname = "127.0.0.1:3306"
    dbname   = "ecommerce"
)

func dsn(dbName string) string {
    return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func main() {
    db, err := sql.Open("mysql", dsn(""))
    if err != nil {
        log.Printf("Error %s when opening DB\n", err)
        return
    }
    defer db.Close()

    ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
    if err != nil {
        log.Printf("Error %s when creating DB\n", err)
        return
    }
    no, err := res.RowsAffected()
    if err != nil {
        log.Printf("Error %s when fetching rows", err)
        return
    }
    log.Printf("rows affected %d\n", no)

    db.Close()
    db, err = sql.Open("mysql", dsn(dbname))
    if err != nil {
        log.Printf("Error %s when opening DB", err)
        return
    }
    defer db.Close()

    db.SetMaxOpenConns(20)
    db.SetMaxIdleConns(20)
    db.SetConnMaxLifetime(time.Minute * 5)

    ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    err = db.PingContext(ctx)
    if err != nil {
        log.Printf("Errors %s pinging DB", err)
        return
    }
    log.Printf("Connected to DB %s successfully\n", dbname)
}

```

Running the above code will print

```
2020/08/11 19:17:44 rows affected 1
2020/08/11 19:17:44 Connected to DB ecommerce successfully

```

That's about it for connecting to MySQL and creating a DB.

Please leave your comments and feedback.
