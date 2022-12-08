# Building a Basic REST API in Go using Fiber

**Fiber** is a new Go-based web framework which has exploded onto the scene and generated **a lot** of interest from the programming community. The repository for the framework has consistently been on the [GitHub Trending](https://github.com/trending "GitHub Trending") page for the Go programming language and as such, I thought I would open up the old VS Code and try my hand at building a simple REST API.

So, in this tutorial, we’ll be covering how you can get started building your own REST API systems in Go using this new Fiber framework!

**By the end of this tutorial**, we will have covered:

*   Project Setup
*   Building a Simle CRUD REST API for a Book management system
*   Breaking out the project into a more extensible format with additional packages.

Let’s dive in!

## Video Tutorial

This tutorial is also available in video format:

## Why Fiber?

If you are coming from another language and trying your hand at developing Go applications then Fiber is an incredibly easy framework to start working with. It presents a familiar feel to Node.js developers who have previously built systems using Express.js. It’s also built on top of `Fasthttp` which is an incredibly performant and minimal HTTP engine built for Go.

If we have a look at the quick start code from the project’s `README.md` we can see just how quickly and simply we can get a simple `HTTP GET` based endpoint returning a `Hello, World!`:

main.go

```go
package main

import "github.com/gofiber/fiber"

func main() {
  app := fiber.New()

  app.Get("/", func(c *fiber.Ctx) {
    c.Send("Hello, World!")
  })

  app.Listen(3000)
}

```

We can then run this and start up our server on `http://localhost:3000` by first initializing our project using `go mod init` and then running `go run main.go` which will download all of `Fiber`’s dependencies before starting up the server:

```s
$ go mod init github.com/tutorialedge/go-fiber-tutorial
$ go run main.go
Fiber v1.9.1 listening on :3000

```

Awesome, we now have the base upon which we can start building more complex systems on top of! 😎

## Introduction

Let’s start off by modifying the quick-start code and making it more extensible:

main.go

```go
package main

import (
	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(3000)
}

```

Let’s breakdown what we have done here.

*   We’ve created a new function called `setupRoutes` which we pass the pointer to our `app`. Within this `setupRoutes` function we map the endpoints to the named functions. This change allows us to move our route logic out from the app initialization logic which is important if we are going to be writing more complex apps.
*   We’ve created the named function `helloWorld` which we have now mapped the `/` endpoint to. This change allows us to write more complex endpoint functions.

## Building our REST API Endpoints

So, with these new changes in place, let’s now look at extending the functionality of our app and creating some additional endpoints from which we can serve requests. We’ll be building a book management system which will feature an in-memory store of books that we have been reading during this pandemic lock down!

We’ll want to create the following endpoints:

*   `/api/v1/book` - a `HTTP GET` endpoint which will return all of the books that you have read during lock down.
*   `/api/v1/book/:id` - a `HTTP GET` endpoint which takes in a path parameter for the book ID and returns just a solitary book
*   `/api/v1/book` - a `HTTP POST` endpoint which will allow us to add new books to the list
*   `/api/v1/book/:id` - a `HTTP DELETE` endpoint which will allow us to delete a book from the list in case we add any books by mistake?

> **Challenge** - Add the `HTTP PUT` endpoint for updating a book on the list.

Let’s see how we can start building this up now.

### The Book Package

Not enough introductory tutorials break out of the `main.go` file and I’ve been guilty of this in the past. So let’s break this cycle and build some solid foundations which can be easily extended should you wish to build more complex apps off the code in this tutorial.

We’ll start by creating a new package within our Go project. This will house all of the logic for our book endpoints:

```s
$ mkdir -p book
$ cd book
$ touch book.go

```

Within this newly created `book.go` file, let’s start defining the stubs for the functions we’ll be mapping to the endpoints outlined above:

book/book.go

```go
package book

import (
	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx) {
	c.Send("All Books")
}

func GetBook(c *fiber.Ctx) {
	c.Send("Single Book")
}

func NewBook(c *fiber.Ctx) {
	c.Send("New Book")
}

func DeleteBook(c *fiber.Ctx) {
	c.Send("Delete Book")
}

```

With this in place, we can then return to the `main.go` file and within our `setupRoutes` function we can map our endpoints to these new functions like so:

main.go

```go
package main

import (
	"github.com/elliotforbes/go-fiber-tutorial/book"
	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(3000)
}

```

Very cool! We have now imported our new `book` package and mapped the endpoints we wanted to these 4 new functions.

Let’s try hitting these endpoints with a few `curl` commands to see if they respond the way we expect:

```s
$ curl http://localhost:3000/api/v1/book
All Books

$ curl http://localhost:3000/api/v1/book/1
Single Book

$ curl -X POST http://localhost:3000/api/v1/book
New Book

$ curl -X DELETE http://localhost:3000/api/v1/book/1
Delete Book

```

**Brilliant, all 4 endpoints have returned the proper response** for their respective HTTP requests!

## Adding a Database

Now that we have our respective endpoints all defined and working as expected, let’s have a look at setting up a simple database that we’ll interact with using `gorm` which simplifies our life talking to databases!

In the root of your project directory, run the following commands to create a new folder called `database/` and a new file called `database.go`:

```s
$ mkdir -p database
$ cd database
$ touch database.go

```

Within this new database.go file, we will want to define a global `DBConn` variable which will be a pointer to a database connection that our endpoints will be using to interact with a local sqlite database:

database/database.go

```go
package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DBConn *gorm.DB
)

```

With this in place, we’ll want to update our `main.go` file to open up the connection to this sqlite database by creating a new `initDatabase()` function.

main.go

```go

package main

import (
	"fmt"
	"github.com/elliotforbes/go-fiber-tutorial/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}

func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)
	app.Listen(3000)

	defer database.DBConn.Close()
}

```

Next, we’ll have to update our `book/book.go` code so that we define a Book `struct` which we’ll use to create database tables.

```go
package book

import (
	"fmt"

	"github.com/elliotforbes/go-fiber-tutorial/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

```

### Updating our Endpoints

Next, we’ll need to update the functions mapped to each of our endpoints. Let’s start off by updating `GetBooks` to return all books:

```go
func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

```

Using the `c.JSON` method provided to us by fiber, we can quickly and easily serialize the books array into a JSON string and return it in our response!

Next let’s update our single book endpoint:

```go
func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

```

Here we’ve used the `c.Params("id")` function in order to retrieve the path parameter which represents the `ID` of the book we want to retrieve. Once again we can use the `c.JSON` function to return this single book.

> **Note** - I’ve not bothered adding error handling to this particular endpoint, it will always assume that the book exists. I’ll leave this as a challenge to the reader to handle this case.

### Adding and Deleting Books

So far we have just dealt with retrieving books from our database, let’s look at how we can start adding and deleting books by updating the `NewBook` and `DeleteBook` functions.

In the `NewBook` function let’s hard code the book we are going to populate for now so that we can incrementally test our API. This will call `db.Create` in order to push the new book into the database for us and then we’ll return the `JSON` for that book:

```go
func NewBook(c *fiber.Ctx) {
	db := database.DBConn
	var book Book
	book.Title = "1984"
	book.Author = "George Orwell"
	book.Rating = 5
	db.Create(&book)
	c.JSON(book)
}

```

Perfect, now finally let’s update the `DeleteBook` function. Here we will actually perform some error handling and check to see if the book first exists within the database before attempting to delete the book and returning a simple message confirming the deletion:

```go
func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
        c.Status(500).Send("No Book Found with ID")
        return
	}
	db.Delete(&book)
	c.Send("Book Successfully deleted")
}

```

### Migrating our Database

Gorm thankfully handles the creation and any updates of our tables for us, so the complexity of setting all this up is minimal. We need to add the call to `AutoMigrate` passing in the struct that we want to generate our tables based off of:

main.go

```go
func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

```

When we next start our API, it will automatically generate the tables for us within our sqlite database.

### Testing our Endpoints:

Now that we have our endpoints defined and talking to the database, the next step is to test these manually to verify if they work as intended:

```s
$ curl http://localhost:3000/api/v1/book
[{"ID":3,"CreatedAt":"2020-04-24T09:20:37.622829+01:00","UpdatedAt":"2020-04-24T09:20:37.622829+01:00","DeletedAt":null,"name":"1984","author":"George Orwell","rating":5},{"ID":4,"CreatedAt":"2020-04-24T09:29:47.573672+01:00","UpdatedAt":"2020-04-24T09:29:47.573672+01:00","DeletedAt":null,"name":"1984","author":"George Orwell","rating":5}]

$ curl http://localhost:3000/api/v1/book/1
{"ID":3,"CreatedAt":"2020-04-24T09:20:37.622829+01:00","UpdatedAt":"2020-04-24T09:20:37.622829+01:00","DeletedAt":null,"name":"1984","author":"George Orwell","rating":5}

$ curl -X POST http://localhost:3000/api/v1/book
{"ID":5,"CreatedAt":"2020-04-24T09:49:16.405426+01:00","UpdatedAt":"2020-04-24T09:49:16.405426+01:00","DeletedAt":null,"name":"1984","author":"George Orwell","rating":5}

$ curl -X DELETE http://localhost:3000/api/v1/book/1
Book Successfully Deleted

```

All of these have worked as we had intended them too! We now have a mostly functioning REST API which we can interact with and throw a frontend on top of!

### Reading JSON Request Data

The final thing I want to cover in this tutorial is reading the body of an incoming request and parsing that into a `book struct` so that we can populate custom data into our database.

Thankfully, the `fiber` framework features a very handy `BodyParser` method which can read in a request body and then populate a struct for us like so:

```go
func NewBook(c *fiber.Ctx) {
	db := database.DBConn
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&book)
	c.JSON(book)
}

```

With this new change in place, let’s re-run our API and send a [HTTP POST request using the curl command](https://tutorialedge.net/snippet/sending-post-request-with-curl/ "HTTP POST request using the curl command") which we’ll pass in a new book:

```s
$ curl -X POST -H "Content-Type: application/json" --data "{\"title\": \"Angels and Demons\", \"author\": \"Dan Brown\", \"rating\": 4}" http://localhost:3000/api/v1/book
{"ID":6,"CreatedAt":"2020-04-24T10:50:52.658811+01:00","UpdatedAt":"2020-04-24T10:50:52.658811+01:00","DeletedAt":null,"title":"Angels and Demons","author":"Dan Brown","rating":4}

```

Everything is working as expected! We can see the new book being added to the database for us with the information we have provided!

## Conclusion

🔥 Awesome, so in this tutorial, we managed to build a really simple REST API for a book management system in Go using the Fiber framework! 🔥

I hope this helped you out and you enjoy the tutorial! 😄 If you liked it or have any additional questions or comments then please let me know in the comments section below!
