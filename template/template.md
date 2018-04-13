# Using Functions Inside Go Templates
In this tutorial we are going to cover how to use template functions like and, eq, and index to add some basic logic to our templates. Once we have a pretty good understanding of how to use these functions we will explore how to go about adding some custom functions to our templates and using them.

## This article is part of a series
This is part three of a four part series introducing the html/template (and text/template) packages in Go. If you haven’t already, I suggest you check out the rest of the series here: An Introduction to Templates in Go. They aren’t required reading, but I think you’ll like them.
If you are enjoying this series, consider signing up for my mailing list to get notified when I release new articles like it. I promise I don’t spam.

## The and function
By default the if action in templates will evaluate whether or not an argument is empty, but what happens when you want to evaluate multiple arguments? You could write nested if/else blocks, but that would get ugly quickly.
Instead, the html/template package provides the and function. Using it is similar to how you would use the and function in Lisp (another programming language). This is easier shown than explained, so lets just jump into some code. Open up main.go and add the following:

```golang
package main

import (
  "html/template"
  "net/http"
)

var testTemplate *template.Template

type User struct {
  Admin bool
}

type ViewData struct {
  *User
}

func main() {
  var err error
  testTemplate, err = template.ParseFiles("hello.gohtml")
  if err != nil {
    panic(err)
  }

  http.HandleFunc("/", handler)
  http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")

  vd := ViewData{&User{true}}
  err := testTemplate.Execute(w, vd)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
```


## Then open up hello.gohtml and add the following to your template.

```
{{if and .User .User.Admin}}
  You are an admin user!
{{else}}
  Access denied!
{{end}}
```

If you run this code you should see the output You are an admin user!. If you update main.go to either not include a *User object, or set Admin to false, or even if you provide nil to the testTemplate.Execute() method you will instead see Access denied!.
The and function takes in two arguments, lets call them a and b, and then runs logic roughly equivalent to if a then b else a. The weirdest part is that and is indeed a function instead of something you place between two variables. Just remember that this is a function and not a logic operation and you should be fine.
Likewise, the template package also provides an or function that operates much like and except it will short circuit when true. IE the logic for or a b is roughly equivalent to if a then a else b so b will never be evaluated if a is not empty.

## Comparison functions (equals, less than, etc)
So far we have been dealing with relatively simple logic revolving around whether or not something is empty or not, but what happens when we need to do some comparisons? For example, what if we want to adjust the class on an object based on whether the user was getting close to going over his usage limit?

The html/template package provides us with a few classes to help do comparison. These are


```
eq - Returns the boolean truth of arg1 == arg2
ne - Returns the boolean truth of arg1 != arg2
lt - Returns the boolean truth of arg1 < arg2
le - Returns the boolean truth of arg1 <= arg2
gt - Returns the boolean truth of arg1 > arg2
ge - Returns the boolean truth of arg1 >= arg2
```

## These are used similarly to how and and or are used, where you first type the function and then type the arguments. For example, you might use the following code in your template to determine which text to render with regards to their API usage.

```
{{if (ge .Usage .Limit)}}
  <p class="danger">
    You have reached your API usage limit. Please upgrade or contact support for more help.
  </p>
{{else if (gt .Usage .Warning)}}
  <p class="warning">
    You have used {{.Usage}} of {{.Limit}} API calls and are nearing your limit. Have you considered upgrading?
  </p>
{{else if (eq .Usage 0)}}
  <p>
    You haven't used the API yet! What are you waiting for?
  </p>
{{else}}
  <p>
    You have used {{.Usage}} of {{.Limit}} API calls.
  </p>
{{end}}
```

## if...else if...else

If you have been following along with the series it is also worth noting that this code also demonstrates how to create an if...elseif...else block, which we haven’t covered yet. These work pretty much like an if...else block so, but they allow you to have a few different conditional clauses.

Using function variables
Up until now we have mostly dealt with data structures inside of our templates, but what happens if we want to call our own functions from within a template? For example, lets imagine we have a User type and we need to find out if the current user has permission to access our enterprise-only feature when creating the UI. We could create a customer struct for the view and add in a field for the permission.

```
type ViewData struct {
  Permissions map[string]bool
}

// or

type ViewData struct {
  Permissions struct {
    FeatureA bool
    FeatureB bool
  }
}
```

The problem with this approach is that we always need to know every feature that is used in the current view, or if we instead used a map[string]bool we would need to fill it with a value for every possible feature. It would be much easier if we could just call a function when we wanted to know whether or not a user had access to a feature. There are a few ways to go about doing this in Go, so I am going to cover a few possible ways to do this.

## 1.Create a method on the User type

The first is the simplest - lets say we have a User type that we already provide to the view, we can just add a HasPermission() method to the object and then use that. To see this in action, add the following to hello.gohtml.

```
{{if .User.HasPermission "feature-a"}}
  <div class="feature">
    <h3>Feature A</h3>
    <p>Some other stuff here...</p>
  </div>
{{else}}
  <div class="feature disabled">
    <h3>Feature A</h3>
    <p>To enable Feature A please upgrade your plan</p>
  </div>
{{end}}

{{if .User.HasPermission "feature-b"}}
  <div class="feature">
    <h3>Feature B</h3>
    <p>Some other stuff here...</p>
  </div>
{{else}}
  <div class="feature disabled">
    <h3>Feature B</h3>
    <p>To enable Feature B please upgrade your plan</p>
  </div>
{{end}}

<style>
  .feature {
    border: 1px solid #eee;
    padding: 10px;
    margin: 5px;
    width: 45%;
    display: inline-block;
  }
  .disabled {
    color: #ccc;
  }
</style>
```

## And then add the following to main.go in the same directory.

```
package main

import (
  "html/template"
  "net/http"
)

var testTemplate *template.Template

type ViewData struct {
  User User
}

type User struct {
  ID    int
  Email string
}

func (u User) HasPermission(feature string) bool {
  if feature == "feature-a" {
    return true
  } else {
    return false
  }
}

func main() {
  var err error
  testTemplate, err = template.ParseFiles("hello.gohtml")
  if err != nil {
    panic(err)
  }

  http.HandleFunc("/", handler)
  http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")

  vd := ViewData{
    User: User{1, "jon@calhoun.io"},
  }
  err := testTemplate.Execute(w, vd)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
```

## After you run your code you should see something like this in your browser:
We are successfully enabling and disabling features on the front end depending on whether the user has access to them! When we declare functions on types we are able to call these in the same manner that we would access data inside of the struct, so this should all feel pretty familiar to you.
Now that we have seen how to call methods let’s check out a more dynamic way to call functions inside of a template using the call function.

## 2.Calling function variables and fields
Lets imagine that for whatever reason that you can’t use the approach above because your method for determining logic needs to change at times. In this case it makes sense to create a HasPermission func(string) bool attribute on the User type and then assign it with a function. Open up main.go and change your code to reflect the following.

```
package main

import (
  "html/template"
  "net/http"
)

var testTemplate *template.Template

type ViewData struct {
  User User
}

type User struct {
  ID            int
  Email         string
  HasPermission func(string) bool
}

func main() {
  var err error
  testTemplate, err = template.ParseFiles("hello.gohtml")
  if err != nil {
    panic(err)
  }

  http.HandleFunc("/", handler)
  http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")

  vd := ViewData{
    User: User{
      ID:    1,
      Email: "jon@calhoun.io",
      HasPermission: func(feature string) bool {
        if feature == "feature-b" {
          return true
        }
        return false
      },
    },
  }
  err := testTemplate.Execute(w, vd)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
```

Everything looks good, but if you visit localhost:3000 in your browser after starting the server you will notice that that we get an error like

template: hello.gohtml:1:10: executing "hello.gohtml" at <.User.HasPermission>: HasPermission has arguments but cannot be invoked as function
When we assign functions to variables, we need to tell the html/template package that we want to call the function. Open up your hello.gohtml file and add the word call right after your if statements, like so.

```
{{if (call .User.HasPermission "feature-a")}}
...
{{if (call .User.HasPermission "feature-b")}}
...
```

## Parenthesis can be used in templates

While parethesis aren’t generally required in Go templates, they can be incredibly useful for making it clear which arguments need to be passed into which functions and specifying a clear order of operations. Keep them in mind as you use templates!
Go ahead and restart your server and check out localhost again. You should see the same page as before, but this time Feature B is enabled instead of Feature A.
call is a function already provided by the html/template package that calls the first argument given to it (the .User.HasPermission function in our case) using the rest of the arguments as arguments to the function call.

## 3.Creating custom functions with a template.FuncMap
The final way of calling our own functions that I am going to cover is creating custom functions with a template.FuncMap. This is, in my opinion, the most useful and powerful way to define functions because it allows us to create global helper methods that can be used throughout our app.
To get started, first head over to the docs for template.FuncMap. The first thing to note is that this type appears to just be a map[string]interface{}, but there is a note below that every interface must be a function with a single return value, or a function with two return values where the first is the data you need to access in the template, and the second is an error that will terminate template execution if it isn’t nil.
This might be confusing at first, so let’s just jump righ into an example. Open main.go again and update it to match the code below.

```
package main

import (
  "html/template"
  "net/http"
)

var testTemplate *template.Template

type ViewData struct {
  User User
}

type User struct {
  ID    int
  Email string
}

func main() {
  var err error
  testTemplate, err = template.New("hello.gohtml").Funcs(template.FuncMap{
    "hasPermission": func(user User, feature string) bool {
      if user.ID == 1 && feature == "feature-a" {
        return true
      }
      return false
    },
  }).ParseFiles("hello.gohtml")
  if err != nil {
    panic(err)
  }

  http.HandleFunc("/", handler)
  http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")

  user := User{
    ID:    1,
    Email: "jon@calhoun.io",
  }
  vd := ViewData{user}
  err := testTemplate.Execute(w, vd)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
```

## And once again open up hello.gohtml and update each if statement to use the new function like so.

```
{{if hasPermission .User "feature-a"}}
...
{{if hasPermission .User "feature-b"}}
...
```

The hasPermission function should now be powering your logic that determines if a feature is enabled or not. In main.go we defined a template.FuncMap that mapped the method name ("hasPermission") to a function that takes in two arguments (a User and a feature string) and then returns true or false. We then called the template.New() function to create a new template, called the Funcs() method on this new template to define our custom functions, and then finally we parsed our hello.gohtml file as the source for our template.

## Define functions before parsing templates
In previous examples we were creating our template by calling the template.ParseFiles function provided by the html/template package. This is a package level function and returns a template after parsing the files. Now we are calling the ParseFiles method on the template.Template type, which has the same return values but applies the changes to the existing template (rather than a brand new one) and then returns the result.
In this situation we need to use the method because we need to first define any custom functions we plan to use in our templates, and once we do this with the template package it will return a *template.Template. After defining those custom functions we can then proceed to parse templates that make use of the functions. If we were to first parse the templates you would see an error related to an undefined function being called in your template.
Next up we will look into how to make this function work without having to pass in a User object every time we call it.

## Making our functions globally useful
The hasPermission function we defined in the last section is great, but one problem with it is that we can only use it when we have access to the User object as well. Passing this around might not be to bad at first, but as an app grows it will end up having many templates and it is pretty easy to forget to pass the User object to a template, or to miss it on a nested template.
Our function would be much simpler if we could we can simplify it and only needed to pass in a feature name, so lets go ahead and update our code to make this happen.
The first thing we need to do is create a function for when no User is present. We will set this in the template.FuncMap before parsing our template so that we don’t get parsing errors, and to make sure we have some logic in place in case the user is not available.

Open up main.go and update the main() function to match the code below.

```
func main() {
  var err error
  testTemplate, err = template.New("hello.gohtml").Funcs(template.FuncMap{
    "hasPermission": func(feature string) bool {
      return false
    },
  }).ParseFiles("hello.gohtml")
  if err != nil {
    panic(err)
  }

  http.HandleFunc("/", handler)
  http.ListenAndServe(":3000", nil)
}
Next we need to define our function that uses a closure. This is basically a fancy way of saying we are going to define a dynamic function that has access to variables that are not necessarily passed into it, but are available when we define the function. In our case that variable will be the User object. Update the handler() function inside of main.go with the following code.

func handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")

  user := User{
    ID:    1,
    Email: "jon@calhoun.io",
  }
  vd := ViewData{user}
  err := testTemplate.Funcs(template.FuncMap{
    "hasPermission": func(feature string) bool {
      if user.ID == 1 && feature == "feature-a" {
        return true
      }
      return false
    },
  }).Execute(w, vd)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
```

## Want to see more closure examples?
If you are interesting in learning a bit more about closures, including a few examples of them in action, I suggest checking out the related article (click the button below). In the article I explain what anonymous functions and closures are, provide examples, and there is even a followup article with common uses for closures in Go.

## RELATED ARTICLE
What is a Closure?
Even though we defined the hasPermission function in our main() function, we are overwriting it inside of our handler when we have access to the User object, but before we execute the template. This is really powerful because we can now use the hasPermission function in any template without worrying about whether the User object was passed to the template or not.

## HTML safe strings and HTML comments
In An Intro to Templates in Go - Contextual Encoding I mentioned that if you need to prevent certain HTML comments from being stripped out of templates that it is possible, but at the time we didn’t cover how. In this section we are going to not only cover how to make this happen, but also how to make any string skip the default encoding process that happens when executing an html/template.
To refresh your memory, imagine you have some HTML in your layout that needs a comment for IE compatibility like so.

```
<!--[if IE]>
<meta http-equiv="Content-Type" content="text/html; charset=Unicode">
<![endif]-->
```

Unfortunately the html/template package will strip out these comments by default, so we need to come up with a way to make comments that are HTML safe. Specifically, we need to create a function that provides us with a template.HTML object with the contents <!--[if IE]> and another for the contents <![endif]-->.

Open main.go and replace its contents with the following.

```
package main

import (
  "html/template"
  "net/http"
)

var testTemplate *template.Template

func main() {
  var err error
  testTemplate, err = template.New("hello.gohtml").Funcs(template.FuncMap{
    "ifIE": func() template.HTML {
      return template.HTML("<!--[if IE]>")
    },
    "endif": func() template.HTML {
      return template.HTML("<![endif]-->")
    },
  }).ParseFiles("hello.gohtml")
  if err != nil {
    panic(err)
  }

  http.HandleFunc("/", handler)
  http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")

  err := testTemplate.Execute(w, nil)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
```

In the main function we implement the functions I described before and name then ifIE and endif. This allows us to update our template (hello.gohtml) like so.

```
{{ifIE}}
<meta http-equiv="Content-Type" content="text/html; charset=Unicode">
{{endif}}
```

And then if you restart the server, reload the page, and then view the page source you should see the following in it:

```
<!--[if IE]>
<meta http-equiv="Content-Type" content="text/html; charset=Unicode">
<![endif]-->
```

This works great, but creating a function for every single comment we might ever want to use in our app would get tedious very quickly. For really common comments (like the endif above) creating its own function makes sense, but we need a way to pass in any HTML comment and ensure that it doesn’t get encoded. To do this we need to define a function that takes in a string and converts it into a template.HTML. Open up main.go again and update your template.FuncMap to match the one below.

```
func main() {
  // ...
  testTemplate, err = template.New("hello.gohtml").Funcs(template.FuncMap{
    "ifIE": func() template.HTML {
      return template.HTML("<!--[if IE]>")
    },
    "endif": func() template.HTML {
      return template.HTML("<![endif]-->")
    },
    "htmlSafe": func(html string) template.HTML {
      return template.HTML(html)
    },
  }).ParseFiles("hello.gohtml")
  //...
}
```

With our new htmlSafe function we can add custom comments as we need to, like an if statement for IE6 specifically.

```
{{htmlSafe "<!--[if IE 6]>"}}
<meta http-equiv="Content-Type" content="text/html; charset=Unicode">
{{htmlSafe "<![endif]-->"}}
```

The last line in this example could also be {{endif}} since we still have that function defined, but I opted to use htmlSafe for consistency.
Our htmlSafe function could even be used in conjunction with other methods (eg {{htmlSafe .User.Widget}}) if we wanted, but, generally speaking, if you want those methods to return HTML safe strings you should probably update their return type to be template.HTML so that your intentions are clarified for future developers.

Summing Up
After followed along with all of the examples you should have a solid grasp on how to use functions in templates as well as how to define your own functions and make them accessible inside of your templates.
In the final article in this series - Creating the V in MVC - I cover how to combine everything that we have learned so far in this series in order to create a reusable view layer for a web application. We will even start to make our pages look prettier with Bootstrap, a popular HTML, CSS, and JS framework, in order to illustrate how this doesn’t affect the rest of our code complexity at all; instead, the view logic is all isolated to our newly created view type.

Link https://www.calhoun.io/intro-to-templates-p3-functions/
