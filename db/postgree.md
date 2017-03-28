# Practical Persistence in Go: Organising Database Access
## 16th July 2015

A few weeks ago someone created a thread on Reddit asking:
In the context of a web application what would you consider a Go best practice for accessing the database 
in (HTTP or other) handlers?  

The replies it got were a genuinely interesting mix. Some people advised using dependency injection, 
a few espoused the simplicity of using global variables, others suggested putting the connection pool pointer into x/net/context.

Me? I think the right answer depends on the project.

What's the overall structure and size of the project? What's your approach to testing?  
How is it likely to grow in the future? All these things and more should play a part when you pick an approach to take. 

So in this post I'll take a look at four different methods for organising your code and structuring access to your  
database connection pool.

### Global variables
The first approach we'll look at is a common and straightforward one – putting the pointer to your database connection
pool in a global variable.

To keep code nice and DRY, you'll sometimes see this combined with an initialisation function that allows the connection 
pool global to be set from other packages and tests.

I like concrete examples, so let's carry on working with the online bookstore database and code from my previous post.   
We'll create a simple application with an MVC-like structure – with the HTTP handlers in main and a models package containing   
a global DB variable,  InitDB() function, and our database logic.  


```
bookstore
├── main.go
└── models
    ├── books.go
    └── db.go
File: main.go


### package main

```golang

import (
    "bookstore/models"
    "fmt"
    "net/http"
)

func main() {
    models.InitDB("postgres://user:pass@localhost/bookstore")

    http.HandleFunc("/books", booksIndex)
    http.ListenAndServe(":3000", nil)
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, http.StatusText(405), 405)
        return
    }
    bks, err := models.AllBooks()
    if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }
    for _, bk := range bks {
        fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
    }
}
File: models/db.go
package models

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

var db *sql.DB

func InitDB(dataSourceName string) {
    var err error
    db, err = sql.Open("postgres", dataSourceName)
    if err != nil {
        log.Panic(err)
    }

    if err = db.Ping(); err != nil {
        log.Panic(err)
    }
}
File: models/books.go
package models

type Book struct {
    Isbn   string
    Title  string
    Author string
    Price  float32
}

func AllBooks() ([]*Book, error) {
    rows, err := db.Query("SELECT * FROM books")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    bks := make([]*Book, 0)
    for rows.Next() {
        bk := new(Book)
        err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
        if err != nil {
            return nil, err
        }
        bks = append(bks, bk)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return bks, nil
}
```

If you run the application and make a request to /books you should get a response similar to:

```
$ curl -i localhost:3000/books
HTTP/1.1 200 OK
Content-Length: 205
Content-Type: text/plain; charset=utf-8

978-1503261969, Emma, Jayne Austen, £9.44
978-1505255607, The Time Machine, H. G. Wells, £5.99
978-1503379640, The Prince, Niccolò Machiavelli, £6.99
```

Using a global variable like this is potentially a good fit if:

All your database logic is contained in the same package.
Your application is small enough that keeping track of globals in your head isn't a problem.
Your approach to testing means that you don't need to mock the database or run tests in parallel.
For the example above using a global works just fine. But what happens in more complicated applications 
where database logic is spread over multiple packages?

One option is to have multiple InitDB calls, but that can quickly become cluttersome and I've personally 
found it a bit flaky (it's easy to forget to initialise a connection pool and get nil-pointer panics at runtime). 
A second option is to create a separate config package with an exported DB variable and import "yourproject/config"
into every file that needs it. Just in case it isn't immediately understandable what I mean, I've included a simple 
example in this gist.

### Dependency injection

The second approach we'll look at is dependency injection. In our example, we want to explicitly pass a connection 
pool pointer to our HTTP handlers and then onward to our database logic.

In a real-world application there are probably extra application-level (and concurrency-safe) items that you want
your handlers to have access to. Things like pointers to your logger or template cache, as well as the database connection pool.

So for projects where all your handlers are in the same package, a neat approach is to put these items into a custom Env type:

```golang
type Env struct {
    db *sql.DB
    logger *log.Logger
    templates *template.Template
}
```

… and then define your handlers as methods against Env. This provides a clean and idiomatic way of making the 
connection pool (and potentially other items) available to your handlers. Here's a full example:  


### File: main.go

```golang
package main

import (
    "bookstore/models"
    "database/sql"
    "fmt"
    "log"
    "net/http"
)

type Env struct {
    db *sql.DB
}

func main() {
    db, err := models.NewDB("postgres://user:pass@localhost/bookstore")
    if err != nil {
        log.Panic(err)
    }
    env := &Env{db: db}

    http.HandleFunc("/books", env.booksIndex)
    http.ListenAndServe(":3000", nil)
}

func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, http.StatusText(405), 405)
        return
    }
    bks, err := models.AllBooks(env.db)
    if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }
    for _, bk := range bks {
        fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
    }
}
File: models/db.go
package models

import (
    "database/sql"
    _ "github.com/lib/pq"
)

func NewDB(dataSourceName string) (*sql.DB, error) {
    db, err := sql.Open("postgres", dataSourceName)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}
File: models/books.go
package models

import "database/sql"

type Book struct {
    Isbn   string
    Title  string
    Author string
    Price  float32
}

func AllBooks(db *sql.DB) ([]*Book, error) {
    rows, err := db.Query("SELECT * FROM books")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    bks := make([]*Book, 0)
    for rows.Next() {
        bk := new(Book)
        err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
        if err != nil {
            return nil, err
        }
        bks = append(bks, bk)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return bks, nil
}
```

Or using a closure…

If you don't want to define your handlers as methods on Env an alternative approach is to put your handler    
logic into a closure and close over the Env variable like so:


### File: main.go

```golang
package main

import (
    "bookstore/models"
    "database/sql"
    "fmt"
    "log"
    "net/http"
)

type Env struct {
    db *sql.DB
}

func main() {
    db, err := models.NewDB("postgres://user:pass@localhost/bookstore")
    if err != nil {
        log.Panic(err)
    }
    env := &Env{db: db}

    http.Handle("/books", booksIndex(env))
    http.ListenAndServe(":3000", nil)
}

func booksIndex(env *Env) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "GET" {
            http.Error(w, http.StatusText(405), 405)
            return
        }
        bks, err := models.AllBooks(env.db)
        if err != nil {
            http.Error(w, http.StatusText(500), 500)
            return
        }
        for _, bk := range bks {
            fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
        }
    })
}
```

### Dependency injection in this way is quite a nice approach when:

All your handlers are contained in the same package.
There is a common set of dependencies that each of your handlers need.
Your approach to testing means that you don't need to mock the database or run tests in parallel.
Again, you could still use this general approach if your handlers and database logic are spread across multiple packages.   
One way to achieve this would be to setup a separate  config package exporting the Env type and close over config.Env    
in the same way as the example above. Here's a basic gist.   


### Using an interface

We can take this dependency injection example a little further. Let's change the models package so that it exports   
a custom DB type (which embeds *sql.DB) and implement our database logic as methods against the DB type.

The advantages of this are twofold: first it gives our code a really clean structure, but – more importantly 
– it also opens up the potential to mock our database for unit testing.  

Let's amend the example to include a new Datastore interface, which implements exactly the same methods as our new DB type.

```golang
type Datastore interface {
    AllBooks() ([]*Book, error)
}
```

We can then use this interface instead of the direct DB type throughout our application. Here's the updated example:

### File: main.go

```golang
package main

import (
    "fmt"
    "log"
    "net/http"
    "bookstore/models"
)

type Env struct {
    db models.Datastore
}

func main() {
    db, err := models.NewDB("postgres://user:pass@localhost/bookstore")
    if err != nil {
        log.Panic(err)
    }

    env := &Env{db}

    http.HandleFunc("/books", env.booksIndex)
    http.ListenAndServe(":3000", nil)
}

func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, http.StatusText(405), 405)
        return
    }
    bks, err := env.db.AllBooks()
    if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }
    for _, bk := range bks {
        fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
    }
}
```

### File: models/db.go

```golang
package models

import (
    _ "github.com/lib/pq"
    "database/sql"
)

type Datastore interface {
    AllBooks() ([]*Book, error)
}

type DB struct {
    *sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
    db, err := sql.Open("postgres", dataSourceName)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return &DB{db}, nil
}
```

### File: models/books.go

```golang
package models

type Book struct {
    Isbn   string
    Title  string
    Author string
    Price  float32
}

func (db *DB) AllBooks() ([]*Book, error) {
    rows, err := db.Query("SELECT * FROM books")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    bks := make([]*Book, 0)
    for rows.Next() {
        bk := new(Book)
        err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
        if err != nil {
            return nil, err
        }
        bks = append(bks, bk)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return bks, nil
}
```

Because our handlers are now using the Datastore interface we can easily create mock database responses for any unit tests:

```golang
package main

import (
    "bookstore/models"
    "net/http"
    "net/http/httptest"
    "testing"
)

type mockDB struct{}

func (mdb *mockDB) AllBooks() ([]*models.Book, error) {
    bks := make([]*models.Book, 0)
    bks = append(bks, &models.Book{"978-1503261969", "Emma", "Jayne Austen", 9.44})
    bks = append(bks, &models.Book{"978-1505255607", "The Time Machine", "H. G. Wells", 5.99})
    return bks, nil
}

func TestBooksIndex(t *testing.T) {
    rec := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/books", nil)

    env := Env{db: &mockDB{}}
    http.HandlerFunc(env.booksIndex).ServeHTTP(rec, req)

    expected := "978-1503261969, Emma, Jayne Austen, £9.44\n978-1505255607, The Time Machine, H. G. Wells, £5.99\n"
    if expected != rec.Body.String() {
        t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
    }
}
```

### Request-scoped context

Finally let's look at using request-scoped context to store and pass around the database connection pool.      
Specifically, we'll make use of the x/net/context package.

Personally I'm not a fan of storing application-level variables in request-scoped context – it feels clunky
and burdensome to me. The x/net/context documentation kinda advises against it too:   

Use context Values only for request-scoped data that transits processes and APIs, not for passing optional parameters to functions.
That said, people do use this approach. And if your project consists of a sprawling set of packages – and using a global config    
is out of the question – it's quite an attractive proposition.

Let's adapt the bookstore example one last time, passing context.Context to our handlers using the pattern suggested 
in this excellent article by Joe Shaw.

### File: main.go

```golang
package main

import (
  "bookstore/models"
  "fmt"
  "golang.org/x/net/context"
  "log"
  "net/http"
)

type ContextHandler interface {
  ServeHTTPContext(context.Context, http.ResponseWriter, *http.Request)
}

type ContextHandlerFunc func(context.Context, http.ResponseWriter, *http.Request)

func (h ContextHandlerFunc) ServeHTTPContext(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
  h(ctx, rw, req)
}

type ContextAdapter struct {
  ctx     context.Context
  handler ContextHandler
}

func (ca *ContextAdapter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
  ca.handler.ServeHTTPContext(ca.ctx, rw, req)
}

func main() {
  db, err := models.NewDB("postgres://user:pass@localhost/bookstore")
  if err != nil {
    log.Panic(err)
  }
  ctx := context.WithValue(context.Background(), "db", db)

  http.Handle("/books", &ContextAdapter{ctx, ContextHandlerFunc(booksIndex)})
  http.ListenAndServe(":3000", nil)
}

func booksIndex(ctx context.Context, w http.ResponseWriter, r *http.Request) {
  if r.Method != "GET" {
    http.Error(w, http.StatusText(405), 405)
    return
  }
  bks, err := models.AllBooks(ctx)
  if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }
  for _, bk := range bks {
    fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
  }
}
```

### File: models/db.go

```golang
package models

import (
    "database/sql"
    _ "github.com/lib/pq"
)

func NewDB(dataSourceName string) (*sql.DB, error) {
    db, err := sql.Open("postgres", dataSourceName)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}
File: models/books.go
package models

import (
    "database/sql"
    "errors"
    "golang.org/x/net/context"
)

type Book struct {
    Isbn   string
    Title  string
    Author string
    Price  float32
}

func AllBooks(ctx context.Context) ([]*Book, error) {
    db, ok := ctx.Value("db").(*sql.DB)
    if !ok {
        return nil, errors.New("models: could not get database connection pool from context")
    }

    rows, err := db.Query("SELECT * FROM books")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    bks := make([]*Book, 0)
    for rows.Next() {
        bk := new(Book)
        err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
        if err != nil {
            return nil, err
        }
        bks = append(bks, bk)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return bks, nil
}

```
