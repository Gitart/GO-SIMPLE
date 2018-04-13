# Using Custom Template Functions in Go
## The Go language comes with a powerful built-in template engine. In this article I show how to add custom template functions (functions you can call from within a template).
In an earlier post I showed one way of creating a web application in Go. There I added support for templates. At its most basic, adding a bundle of templates looks like this:

```golang
import (
    "os"
    "html/template"
)

func main() {

    // Create templates
    tpl := template.Must(template.New("main").ParseGlob("*.html"))

}
```

The main line above creates a new bundle of templates by parsing all of the files that match the patterh *.html.
Suppose we begin with a template called index.html that looks like this:

```
<!DOCTYPE html>
<html>
  <head>
    <title>{{.Title}}</title>
  </head>
  <body>
    <h1>{{.Title}}</h1>
    {{.Content}}
  </body>
</html>
```

In this template there are two template directives: {{.Title}} (which appears twice) and {{.Content}}.
We can very quickly turn the code above into a simple template renderer by adding this:

```
tplVars := map[string]string {
    "Title": "Hello world",
    "Content": "Hi there",
}

tpl.ExecuteTemplate(os.StdOut, "index.html", tplVars)
```

If we were to run this (go run main.go), we would get output like this:
```
<!DOCTYPE html>
<html>
  <head>
    <title>Hello world</title>
  </head>
  <body>
    <h1>Hello world</h1>
    Hi there
  </body>
</html>
```

But let's say that we wanted to do a little additional formatting. Generally, the right place to handle presentation logic is in the template. And Go provides some basic tools for this. But say we wanted to make sure that our title was always in "Title Case". There's no built-in template function for this, even though there is a function in strings that does this.
It would be convenient to be able to execute strings.Title inside of a template. While we can't do that by default, adding this feature is pretty easy. We just add a function map to the template renderer.

Here's the new code:

```
package main

import (
    "html/template"
    "os"
    "strings"
)

func main() {

    funcMap := template.FuncMap {
        "title": strings.Title,
    }

    tpl := template.Must(template.New("main").Funcs(funcMap).ParseGlob("*.html"))
    tplVars := map[string]string {
        "Title": "Hello world",
        "Content": "Hi there",
    }
    tpl.ExecuteTemplate(os.Stdout, "index.html", tplVars)
}
```
We've built out a new template function map (template.FuncMap, actually just a map[string]interface{}), and we've declared one new template function called title. When executed, title just calls strings.Title.
In order to tell the template engine about our new functions, we have to pass the function map in, and we have to do it before we parse the template files:

```
tpl := template.Must(template.New("main").Funcs(funcMap).ParseGlob("*.html"))
Now let's adjust our template to use just this new function:
```
```
<!DOCTYPE html>
<html>
  <head>
    <title>{{.Title | title}}</title>
  </head>
  <body>
    <h1>{{.Title}}</h1>
    {{.Content}}
  </body>
</html>
```

Notice that we use the new function on the fourth line, but not later on. Now when we run our program, the output looks like this:
```
<!DOCTYPE html>
<html>
  <head>
    <title>Hello World</title>
  </head>
  <body>
    <h1>Hello world</h1>
    Hi there
  </body>
</html>
```

Only the first .Title was transformed into title case, since we only executed the template function title on that first instance.
You can build your own template functions, too. You're not restricted just to existing functions:

```
    funcMap := template.FuncMap {
        "title": strings.Title,
        "tableflip": func () string { return "(╯°□°）╯︵ ┻━┻" },
    }
```
Now we've added a tableflip function. Each time we embed {{tableflip}} in our template, it will produce the string (╯°□°）╯︵ ┻━┻.
That's how you can add custom template functions to Go templates. There were a number of "basic" functions that I wanted, 
so I created a package that collects them. Sprig has a dozen or two utility functions for handling dates, formatting, 
and basic integer math. (Feel free to contribute your additions via pull request!)
