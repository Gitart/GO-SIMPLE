## Go pongo2

[Source](https://zetcode.com/golang/pongo2/)

Go pongo2 tutorial shows how to work with templates in Golang with pongo2 template engine.

A template engine is a library designed to combine templates with a data to produce documents. Template engines are used to generate large amounts of emails, in source code preprocessing, or to produce dynamic HTML pages.

A template consists of static data and dynamic regions. The dynamic regions are later replaced with data. The rendering function later combines the templates with data. A template engine is used to combine templates with a data model to produce documents.

The pongo2 library is a Go template engine inspired by Django's template engine.

The pongo2 uses various delimiters in template string:

    {% %} - statements
    {{ }} - expressions to print to the template output
    {# #} - comments which are not included in the template output
    # ## - line statements

Templates can be read from strings with pongo2.FromString, files with pongo2.FromFile, or bytes with pongo2.FromBytes.

The documents are rendered with Execute, ExecuteWriter, or ExecuteBytes functions. These functions accept a Context, which provides constants, variables, instances or functions to a template.
Go pongo2.FromString

The pongo2.FromString reads a template from a string.
main.go

package main

import (
    "fmt"
    "log"

    "github.com/flosch/pongo2/v5"
)

func main() {

    tpl, err := pongo2.FromString("Hello {{ name }}!")

    if err != nil {
        log.Fatal(err)
    }

    res, err := tpl.Execute(pongo2.Context{"name": "John Doe"})

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(res)
}

The example produces a simple text message.

tpl, err := pongo2.FromString("Hello {{ name }}!")

The variable to print is placed within the {{ }} brackets.

res, err := tpl.Execute(pongo2.Context{"name": "John Doe"})

We render the final string with Execute. In the context, we pass a value for the name variable.

$ go run main.go
Hello John Doe!

main.go

package main

import (
    "fmt"
    "log"

    "github.com/flosch/pongo2/v5"
)

func main() {

    tpl, err := pongo2.FromString("{{ name }} is a {{ occupation }}")

    if err != nil {
        log.Fatal(err)
    }

    name, occupation := "John Doe", "gardener"
    ctx := pongo2.Context{"name": name, "occupation": occupation}

    res, err := tpl.Execute(ctx)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(res)
}

In this example, we pass two variables in the context.

$ go run main.go
John Doe is a gardener

Go pongo2.FromFile

With the pongo2.FromFile function, we read the template from a file.
message.tpl

{{ name }} is a {{ occupation }}

This is the template file.
main.go

package main

import (
    "fmt"
    "log"

    "github.com/flosch/pongo2/v5"
)

func main() {

    tpl, err := pongo2.FromFile("message.tpl")

    if err != nil {
        log.Fatal(err)
    }

    name, occupation := "John Doe", "gardener"
    ctx := pongo2.Context{"name": name, "occupation": occupation}

    res, err := tpl.Execute(ctx)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(res)
}

The example produces a simple message, while reading the template from a file.
Go pongo2 for directive

The for directive is used to iterate over a data collection in a template.
words.tpl

{% for word in words -%}
    {{ word }}
{% endfor %}

In the template, we use the for directive to go through the elements of the words data structure. The - character strips whitespace characters.
main.go

package main

import (
    "fmt"
    "log"

    "github.com/flosch/pongo2/v5"
)

func main() {

    tpl, err := pongo2.FromFile("words.tpl")

    if err != nil {
        log.Fatal(err)
    }

    words := []string{"sky", "blue", "storm", "nice", "barrack", "stone"}

    ctx := pongo2.Context{"words": words}

    res, err := tpl.Execute(ctx)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(res)
}

In the program, we pass a slice of words to the tempate engine. We get a list of words as the output.

$ go run main.go
sky
blue
storm
nice
barrack
stone

Go pongo2 filter

A filter can be applied to data to modify them. Filters are applied after the | character.
words.tpl

{% for word in words -%}
    {{ word }} has {{ word | length }} characters
{% endfor %}

The length filter returns the size of the string.
main.go

package main

import (
    "fmt"
    "log"

    "github.com/flosch/pongo2/v5"
)

func main() {

    tpl, err := pongo2.FromFile("words.tpl")

    if err != nil {
        log.Fatal(err)
    }

    words := []string{"sky", "blue", "storm", "nice", "barrack", "stone"}

    ctx := pongo2.Context{"words": words}

    res, err := tpl.Execute(ctx)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(res)
}

In the program, we pass a slice of words to the template. We print each word and its size.

$ go run main.go
sky has 3 characters
blue has 4 characters
storm has 5 characters
nice has 4 characters
barrack has 7 characters
stone has 5 characters

Go pongo2 if condition

Conditions can be created with if/endif directives.
todos.tpl

{% for todo in todos -%}
    {% if todo.Done %}
        {{- todo.Title -}}
    {% endif %}
{% endfor %}

In the template file, we use the if directive to output only tasks that are finished.
main.go

package main

import (
    "fmt"
    "log"

    "github.com/flosch/pongo2/v5"
)

type Todo struct {
    Title string
    Done  bool
}

type Data struct {
    Todos []Todo
}

func main() {

    tpl, err := pongo2.FromFile("todos.tpl")

    if err != nil {
        log.Fatal(err)
    }

    todos := []Todo{
        {Title: "Task 1", Done: false},
        {Title: "Task 2", Done: true},
        {Title: "Task 3", Done: true},
        {Title: "Task 4", Done: false},
        {Title: "Task 5", Done: true},
    }

    ctx := pongo2.Context{"todos": todos}

    res, err := tpl.Execute(ctx)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(res)
}

We generate on output from a slice of todos. In the output we include only finished tasks.
Server example

In the next example, we use templates in a server application.
users.html

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Users</title>
</head>

<body>
    <table>
        <thead>
            <tr>
                <th>Name</th>
                <th>Occupation</th>
            </tr>
        </thead>
        <tbody>
            {% for user in users %}
            <tr>
                <td>{{ user.Name }} </td>
                <td>{{ user.Occupation }}</td>
            </tr>
            {% endfor %}
        </tbody>
    </table>
</body>
</html>

The output is an HTML file. The users are displayed in an HTML table.
main.go

package main

import (
    "net/http"

    "github.com/flosch/pongo2/v5"
)

type User struct {
    Name       string
    Occupation string
}

var tpl = pongo2.Must(pongo2.FromFile("users.html"))

func usersHandler(w http.ResponseWriter, r *http.Request) {

    users := []User{
        {Name: "John Doe", Occupation: "gardener"},
        {Name: "Roger Roe", Occupation: "driver"},
        {Name: "Peter Smith", Occupation: "teacher"},
    }

    err := tpl.ExecuteWriter(pongo2.Context{"users": users}, w)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {

    http.HandleFunc("/users", usersHandler)
    http.ListenAndServe(":8080", nil)
}

The web server returns an HTML page with a table of users for the /users URL path.

var tpl = pongo2.Must(pongo2.FromFile("index.html"))

The pongo2.Must is a helper function which pre-compiles the templates at application startup.

err := tpl.ExecuteWriter(pongo2.Context{"users": users}, w)

The ExecuteWriter renders the template with the given context and writes the output to the response writer on success. Nothing is written on error; instead the error is being returned.

In this tutorial, we have created dynamic documents using third-party pongo2 templating engine. 
