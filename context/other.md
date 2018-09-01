# Sample

```
Secret message
func secretMessageHandler(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "42")
}
Here's a way to write middleware with just the standard library:

func requireKey(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.FormValue("key") != "12345" {
            http.Error(w, "bad key", http.StatusForbidden)
            return
        }
        h(w, r)
    }
}

- In main.go:

http.HandleFunc("/secret/message", requireKey(secretMessageHandler))
```
### There's been a change of plans...
We were hard-coding the key, but your boss says now we need to check Redis.

Let's just make our Redis connection a global variable for now...

```
var redisDB *redis.Client

func main() {
    redisDB = ... // set up redis
}
func requireKeyRedis(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        key := r.FormValue("key")
        userID, err := redisDB.Get("auth:" + key).Result()
        if key == "" || err != nil {
            http.Error(w, "bad key", http.StatusForbidden)
            return
        }
        log.Println("user", userID, "viewed message")
        h(w, r)
    }
}
```
ust one quick addition...
We need to issue temporary session tokens for some use cases, so we need to check if either a key or a session is provided.

func requireKeyOrSession(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        key := r.FormValue("key")
        // set key from db if we have a session
        if session := r.FormValue("session"); session != "" {
            var err error
            if key, err = redisDB.Get("session:" + session).Result(); err != nil {
                http.Error(w, "bad session", http.StatusForbidden)
                return
            }
        }
        userID, err := redisDB.Get("auth:" + key).Result()
        if key == "" || err != nil {
            http.Error(w, "bad key", http.StatusForbidden)
            return
        }
        log.Println("user", userID, "viewed message")
        h(w, r)
    }
}

By the way...
Your boss also asks:

Can we also check the X-API-Key header?
Can we restrict certain keys to certain IP addresses?
Can we ...?
There's too much to shove into one middleware: so we make an auth package.

package auth

type Auth struct {
    Key     string
    Session string
    UserID  string
}

func Check(r *http.Request) (auth Auth, ok bool) {
    // lots of complicated checks
}

What about Redis?
We need to reference the DB from our new auth package as well.

Should we pass the connection to Check?

func Check(redisDB *redis.Client, r *http.Request) (Auth, bool) { ... }
What happens we need to check MySQL as well?

func Check(redisDB *redis.Client, archiveDB *sql.DB, r *http.Request) (Auth, bool) { ... }
Your boss says MongoDB is web scale, so that gets added too.

func Check(redisDB *redis.Client, archiveDB *sql.DB, mongo *mgo.Session, r *http.Request) (Auth, bool) { ... }
This isn't going to work...


How about an init method?
Making a global here too?

var redisDB *redis.Client
func Init(r *redis.Client, ...) {
    redisDB = r
}
That doesn't solve our arguments problem. Let's shove them in a struct.

package config
type Context struct {
    RedisDB   *redis.Client
    ArchiveDB *sql.DB
    ...
}
Init with this guy?

auth.Init(appContext)
Who inits who? 
What about tests?

Session table
Let's try making a global map of connections to auths.

var authMap map[*http.Request]*auth.Auth
Then populate it during our check.

func requireKeyOrSession(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ...
        a, ok := auth.Check(dbContext, r)
        authMapMutex.Lock()
        authMap[r] = &a
        ...
    }
}
Should work, but will our *http.Requests leak? We need to make sure to clean them up. 
What happens when we need to keep track of more than just Auth? 
How do we coordinate this data across packages? What about concurrency? 
(This is kind of how gorilla/sessions works)

Attempt #2: Goji
Goji is a popular web micro-framework. Goji handlers take an extra parameter called web.C (probably short for Context).

c.Env is a map[interface{}]interface{} for storing arbitrary data — perfect for our auth token! This used to be a map[string]interface{}, more on this later.

Let's rewrite our auth middleware for Goji:

func requiresKey(c *web.C, h http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        a := c.Env["auth"]
        if a == nil {
            http.Error(w, "bad key", http.StatusForbidden)
            return
        }
        h.ServeHTTP(w, r)
    }
    return http.HandlerFunc(fn)
}
Goji groups
We can set up groups of routes:

package main

import (
    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
)

func main() {
    ...
    secretGroup := web.New()
    secretGroup.Use(requiresKey)
    secretGroup.Get("/secret/message", secretMessageHandler)
    goji.Handle("/secret/*", secretGroup)
    goji.Serve()
}
This will run our checkAuth for all routes under /secret/.

Downside: Goji-flavored context
Let's say we want to re-use our auth package elsewhere, like a batch process.

Do we want to put our database connections in web.C, even if we're not running a web server? Should all of our internal packages be importing Goji?

package auth
func Check(c web.C, session, key string) bool {
    // How do we call this if we're not using goji?
    redisDB, _ := c.Env["redis"].(*redis.Client) // kind of ugly...
}
Having to do a type assertion every time we use this DB is annoying. Also, what happens when some other library wants to use this "redis" key?

Downside: Groups need to be set up once, in main.go
Defining middleware for a group is tricky. What happens if you have code like...

package addon

func init() {
    goji.Get("/secret/addon", addonHandler) // will secretGroup handle this?
}
Everything works will if your entire app is set up in main.go, but in my experience it's very finicky and hard to reason about handlers that are set up in other ways.

Attempt #3: kami & x/net/context
What is x/net/context?

It's an almost-standard package for sharing context across your entire app.
Includes facilities for setting deadlines and cancelling requests.
Includes a way to store data similar to Goji's web.C.
Immutable, must be replaced to update
Check out this official blog post, which focuses mostly on x/net/context for cancellation:

blog.golang.org/context

Quick example:

ctx := context.Background() // blank context
ctx = context.WithValue(ctx, "my_key", "my_value")
fmt.Println(ctx.Value("my_key").(string)) // "my_value"
kami
kami is a mix of HttpRouter, x/net/context, and Goji, with a very simple middleware system included.

package main

import (
    "fmt"
    "net/http"

    "github.com/guregu/kami"
    "golang.org/x/net/context"
)

func hello(ctx context.Context, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", kami.Param(ctx, "name"))
}

func main() {
    kami.Get("/hello/:name", hello)
    kami.Serve()
    
    Example: sharing DB connections
import "github.com/guregu/db"
I made a simple package for storing DB connections in your context. At Gunosy, we use something similar. db.OpenSQL() returns a new context containing a named SQL connection.

func main() {
    ctx := context.Background()
    mysqlURL := "root:hunter2@unix(/tmp/mysql.sock)/myCoolDB"
    ctx = db.OpenSQL(ctx, "main", "mysql", mysqlURL)
    defer db.Close(ctx)                              // closes all DB connections
    kami.Context = ctx
    kami.Get("/hello/:name", hello)
    kami.Serve()
}
kami.Context is our "god context" from which all request contexts are derived.

Example: sharing DB connections (2)
Within a request, we use db.SQL(ctx, name) to retrieve the connection.

func hello(ctx context.Context, w http.ResponseWriter, r *http.Request) {
    mainDB := db.SQL(ctx, "main") // *sql.DB
    var greeting string
    mainDB.QueryRow("SELECT content FROM greetings WHERE name = ?", kami.Param(ctx, "name")).
        Scan(&greeting)
    fmt.Fprintf(w, "Hello, %s!", greeting)
}

Tests
For tests, you can put a mock DB connection in your context.

main_test.go:

import _ "github.com/mycompany/testhelper"
testhelper/testhelper.go:

import (
    "github.com/guregu/db"
    "github.com/guregu/kami"
    _ "github.com/guregu/mogi"
)

func init() {
    ctx := context.Background()
    // use mogi for tests
    ctx = db.OpenSQL("main", "mogi", "")
    kami.Context = ctx
}
How does it work?
Because context.Value() takes an interface{}, we can use unexported type as the key to "protect" it. This way, other packages can't screw with your data. In order to interact with a database, you have to use the exported functions like OpenSQL, and Close.

package db

import (
    "database/sql"
    "golang.org/x/net/context"
)

type sqlkey string // lowercase!

// SQL retrieves the *sql.DB with the given name or nil.
func SQL(ctx context.Context, name string) *sql.DB {
    db, _ := ctx.Value(sqlkey(name)).(*sql.DB)
    return db
}
BTW: This is why Goji switched its web.C from a map[string]interface{} to map[interface{}]interface{}.

Graceful
This is the "Goji" part of kami. Literally copy and pasted from Goji.

kami.Serve() // works *exactly* like goji.Serve()
Supports Einhorn for graceful restarts.

Thank you, Goji.

Downsides
kami isn't perfect. But everything was was on this slide is now taken care of!

Pull requests are always welcome.

Production ready!
We use kami to power the Gunosy API and it works just fine!

Switching to x/net/context eliminates nearly all global variables.

No more somepkg.Init() madness.

Easy to test: just put mocks inside your context.

Check it out!
