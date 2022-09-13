# Flat Application Structure in Go

Rather than spending time trying to figure out how to break code into packages, an app with a flat structure would just place all of the `.go` files in a single package.

myapp/
 main.go server.go user.go lesson.go course.go

A flat application structure is what almost everyone begins with when diving into Go. Every program in the [Go tour](https://tour.golang.org/welcome/1), most exercises in [Gophercises](https://gophercises.com/), and many others early Go programs don’t get broken into any packages at all. Instead we just create a few `.go` files and put all of our code in the same (often `main`) package.

At first this sounds awful. Won’t the code become unwieldy extremely quickly? How will I separate my business logic from my UI rendering code? How will I find the right source files? After all, a big part of why we use packages is to separate concerns while making it easier to navigate to the correct source files quickly.

## Using a flat structure effectively

When using a flat structure you should still try to adhere to coding best practices. You will want to separate different parts of your application using different `.go` files:

```bash
myapp/
  main.go # read configs and start your app here
  server.go # overall http handling logic goes here
  user_handler.go # user http handler logic goes here
  user_store.go # user DB logic goes here
  # and so on...

```

Globals can still become problematic, so you should consider using types with methods to keep them out of your code:

```go
type Server struct {
  apiClient *someapi.Client
  router *some.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  s.router.ServeHTTP(w, r)
}

```

And your `main()` function should probably still be stripped of most logic outside of setting up the application:

```go
// Warning: This example is VERY contrived and may not even compile.

type Config struct {
  SomeAPIKey     string
  Port           string
  EnableTheThing bool
}

func main() {
  var config Config
  config.SomeAPIKey = os.Getenv("SOMEAPI_KEY")
  config.Port = os.Getenv("MYAPP_PORT")
  if config.Port == "" {
    config.Port = "3000"
  }
  config.EnableTheThing = true

  if err := run(config); err != nil {
    log.Fatal(err)
  }
}

func run(config Config) error {
  server := myapp.Server{
    APIClient: someapi.NewClient(config.SomeAPIKey),
  }
  return http.ListenAndServe(":" + config.Port, server)
}

```

In fact, you could really use what is basically a flat structure with your code all in a single package and a separate `main` package where you define your command. This would allow you to use the common `cmd` subdirectory pattern:

```bash
myapp/
  cmd/
    web/
      # package main
      main.go
    cli/
      # package main
      main.go
  # package myapp
  server.go
  user_handler.go
  user_store.go
  ...

```

In this example your application is still basically flat, but you pulled out the `main` package because you had a need - like perhaps needing to support two commands using the same core application.

## Why should I use a flat structure?

The key benefit of a flat structure isn’t that we are keeping all of our code in a single directory or anything silly like that. The core benefit of this structure is that you can stop worrying about how to organize things and instead get on with solving the problems you set out to solve with your application.

I absolutely love how reminiscent of my PHP days this application structure feels. When I was first learning to code I started off with random PHP files with logic intermingled with all sorts of HTML and it was a mess. I’m not suggesting that we should build large applications way - it would suck - but I was less worried about where everything should go and more concerned with learning how to write code and solve my particular problems. Using a flat structure just makes it easier to focus on learning and building, whether you are learning about your application’s needs, your domain, or how to code in general.

This is true because we can stop worrying about things like, “Where should this logic go?” because it is easy to fix a mistake if we make one. If it is a function, we can move it to any new source file in our package. If it is a method on the wrong type, we can create two new types and split up the logic from the original. And with all of this we don’t have to worry about running into weird cyclical dependency issues because we only have one package.

Another big reason to consider a flat structure is that it is much easier for your structure to evolve as your application grows in complexity. When it becomes apparent that you could benefit from breaking code into a separate package, all you often need to do is move a few source files into a subdirectory, change their package, and update any reference to use the new package prefix. Eg if we had `SqlUser` and decided we would benefit from having a separate `sql` package to handle all our database related logic, we would update any references to now use `sql.User` after moving the type to the new package. I have found that structures like MVC are a bit more challenging to refactor, albeit not impossible or as hard as it might be in other programming languages.

A flat structure can be especially useful for beginners who are often too quick to create packages. I can’t really say why this phenomenon happens, but newcomers to Go love to create tons of packages and this almost always leads to stuttering (`user.User`), cyclical dependencies, or some other issue.

*In the next article on MVC we will explore how this phenomenon of creating too many packages can make MVC seem impossible in Go, despite that being far from the truth.*

By putting off decisions to create new packages until our application grows a bit and we understand it better, budding Gophers are far less likely to make this mistake.

This is also why many people will encourage developers to avoid breaking their code into microservices too early - you often don’t have enough knowledge to really know what should and shouldn’t be split into a microservice early on and preemptive microservicing (*I kinda hope that becomes a saying*) will just lead to more work in the future.

## A flat structure isn’t all sunshine and rainbows

It would be disingenuous of me to pretend like there aren’t any downsides to using a flat structure, so we should talk about those as well.

For starters, a flat structure can only get you so far. It will work for a while - probably longer than you think - but at some point your app will become complex enough that you need to start breaking it up. The upside to using a flat structure is that you can put this off and you will probably understand your code better when you do break it up. The downside is that you will need to spend some time refactoring at some point, and you might (*maybe - but its a stretch*) find yourself refactoring to a structure you wanted to start with anyway.

Naming collisions can also be awkward at times with a flat structure. For instance, let’s say you want a `Course` type in your application, but how you represent a course in the database isn’t the same as how you render a course in JSON. A quick solution to this is to create two types, but since they are both in the same package you need different names for each and may end up with something like: `SqlCourse` and `JsonCourse`. This isn’t really that big of a deal, but it kinda sucks that we ended up with zero types named simply `Course`.

It also isn’t always super simple to refactor code into a new package. Yes, it is usually pretty easy, but because all of your code is in one package you can occasionally run into code that is cyclical by nature. For example, imagine if our courses had an ID that always started with `crs_` in the JSON response, and we wanted to return the price in various currencies. We might create a `JsonCourse` to handle that:

```go
type JsonCourse struct {
  ID       string `json:"id"`
  Price struct {
    USD string `json:"usd"`
  } `json:"price"`
}

```

Meanwhile the `SqlCourse` only needs to store an integer ID and a single price in cents that we can format in various currencies.

```go
type SqlCourse struct {
  ID    int
  Price int
}

```

Now we need a way to convert from `SqlCourse` to `JsonCourse`, so we might make this a method on the `SqlCourse` type:

```go
func (sc SqlCourse) ToJson() (JsonCourse, error) {
  jsonCourse := JsonCourse{
    ID: fmt.Sprintf("crs_%v", sc.ID),
  }
  jsonCourse.Price.USD = Price: fmt.Sprintf("%d.%2d", sc.Price/100, sc.Price%100)
  return jsonCourse, nil
}

```

And then later we might need a way to parse incoming JSON and convert it into our SQL equivalent, so we add that to the `JsonCourse` type as another method:

```go
func (jc JsonCourse) ToSql() (SqlCourse, error) {
  var sqlCourse SqlCourse
  // JSON ID is "crs_123" and we convert to "123"
  // for SQL IDs
  id, err := strconv.Atoi(strings.TrimPrefix(jc.ID, "crs_"))
  if err != nil {
    // Note: %w is a Go 1.13 thing that I haven't really
    // tested out, so let me know if I'm using it wrong 😂
    return SqlCourse{}, fmt.Errorf("converting json course to sql: %w", err)
  }
  sqlCourse.ID = id
  // JSON price is 1.00 and we may convert to 100 cents
  sqlCourse.Price = ...
  return sqlCourse, nil
}

```

Every step we have taken here made sense and felt logical, but we are now left with two types that MUST be in the same package otherwise they will present a cyclical dependency.

I find that issues like this are far less likely to occur when MVC, Domain Driven Design, and other app structures are used ahead of time, but if we are being honest this isn’t really *that* hard to fix. All we really need to do is extract the conversion logic and place it wherever we use both types.

```go
func JsonCourseToSql(jsonCourse json.Course) (sql.Course, error) {
  // move the `ToSql()` functionality here
}

func SqlCourseToJson(sqlCourse sql.Course) (json.Course, error) {
  // Move the `ToJson()` functionality here
}

```

Lastly, flat structures aren’t hip; if you want to rock a sweet mustache and show your buddies at the coffee shop how awesome you are, this might not earn you bonus points. On the other hand, if you just want your code to work, this could be a good fit. 🤷‍♂️

## Is a flat structure right for me?

First, let me make a general recommendation: **Don’t try to skip to the end in an attempt to avoid ever needing to refactor code on your way.** It doesn’t work, ever, and you will probably just end up doing more work that way. It is nearly impossible to predict future requirements for your software, and this is just one more way we try to do it as developers.

Not only is it unlikely to save you any time, you could also be doing yourself a disservice. Large enterprise organizations use more complicated code structures because they need to. Whether it is because they need to test with a variety of configurations, need rock-solid unit testing, or whatever else, there is pretty much always a reason they use the complicated structure they do. If you are a solo developer learning to code, or a small team trying to move quickly, your needs aren’t the same. Trying to pretend like you are a large org without understanding why they chose the structure they did is more likely to slow you down than actually help you.

*The caveat here is that if you know what you are doing this isn’t always true, but I still find it to be true in more cases than not.*

What this all means is that you should pick the structure that best suites your situation. If you are unsure of how complicated your application is going to be or are just learning, a flat structure can be a great starting point. Then you can refactor and/or extract packages once you have a better understanding of what your app needs. This is a point that many developers love to ignore - without building out an application it can often be hard to understand how it should be split up. This problem can also pop up when people jump to microservices too quickly.

On the other hand, if you already know your application is going to be massive - perhaps you are porting a large application from one stack to another - then this might be a bad starting point because you already have a lot of context to work from.

## Additional considerations

A few additional things to keep in mind if you do opt to try out a flat structure:

*   Just because you are only working within one package doesn’t mean you should avoid best practices; globals are generally bad, configuration should likely happen in `main()` (or perhaps a `run()` if you use that pattern), and `init()` is almost always a mistake.
*   Starting with a flat structure doesn’t lock you into a single package. Break code into separate packages as soon as it becomes clear this will be beneficial.
*   You can still benefit from breaking code up into separate source files and using custom types.

**Disclaimer** - *You are welcome to try things like `init()` and globals when learning. In fact, as a junior dev I think getting your code working and understanding it is more important than perfect structure because you will learn more by coding than by fretting over doing something wrong. Writing the initial, working version is usually far harder than refactoring it using Go best practices, and you are more likely to understand WHY seasoned developers make the recommendations they do after writing it the “bad” way. A similar example is using Redux in React; without ever experiencing the problems redux solves, you can’t really appreciate what it does.*
