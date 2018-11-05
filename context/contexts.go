package main

import (
        "fmt"
        "net/http"
        "time"

        "golang.org/x/net/context"

        "github.com/husobee/backdrop"
        "github.com/husobee/vestigo"
        "github.com/satori/go.uuid"
        "github.com/tylerb/graceful"
)

// middleware is a definition of  what a middleware is, 
// take in one handlerfunc and wrap it within another handlerfunc
type middleware func(http.HandlerFunc) http.HandlerFunc

// buildChain builds the middlware chain recursively, functions are first class
func buildChain(f http.HandlerFunc, m ...middleware) http.HandlerFunc {
        // if our chain is done, use the original handlerfunc
        if len(m) == 0 {
                return f
        }
        // otherwise nest the handlerfuncs
        return m[0](buildChain(f, m[1:cap(m)]...))
}

func main() {
        // backdrop is my context solution
        backdrop.Start(nil)

        // create an endpoint grouping called publicChain
        // which has the public middlewares
        var publicChain = []middleware{
                PublicMiddleware,
        }
        // create an endpoint grouping called privateChain for
        // urls we want to protect with middlewares
        var privateChain = []middleware{
                PublicMiddleware,
                AuthMiddleware,
                PrivateMiddleware,
        }

        // set up awesome router ;)
        router := vestigo.NewRouter()
        // public has the public middleware chain
        router.Get("/v:version/public", buildChain(f, publicChain...))
        // private has the private middleware chain
        router.Get("/v:version/private", buildChain(f, privateChain...))

        // graceful start/stop server
        srv := &graceful.Server{
                Timeout: 5 * time.Second,
                Server: &http.Server{
                        Addr: ":1234",
                        // top level handler needs to clear the context
                        // per each request, use this wrapper handler
                        Handler: backdrop.NewClearContextHandler(router),
                },
        }
        srv.ListenAndServe()
}

// AuthMiddleware - takes in a http.HandlerFunc, and returns a http.HandlerFunc
var AuthMiddleware = func(f http.HandlerFunc) http.HandlerFunc {
        // one time scope setup area for middleware
        return func(w http.ResponseWriter, r *http.Request) {
                // ... pre handler functionality
                fmt.Println("start auth")
                f(w, r)
                fmt.Println("end auth")
                // ... post handler functionality
        }
}

// PrivateMiddleware - takes in a http.HandlerFunc, and returns a http.HandlerFunc
var PrivateMiddleware = func(f http.HandlerFunc) http.HandlerFunc {
        // one time scope setup area for middleware
        return func(w http.ResponseWriter, r *http.Request) {
                // ... pre handler functionality
                fmt.Println("start private")
                f(w, r)
                fmt.Println("end private")
                // ... post handler functionality
        }
}

// PublicMiddleware - takes in a http.HandlerFunc, and returns a http.HandlerFunc
var PublicMiddleware = func(f http.HandlerFunc) http.HandlerFunc {
        // one time scope setup area for middleware
        return func(w http.ResponseWriter, r *http.Request) {
                // add a request id..
                backdrop.Set(r, "id", uuid.NewV4())
                // ... pre handler functionality
                fmt.Println("start public")
                f(w, r)
                fmt.Println("end public")
                // ... post handler functionality
        }
}

// this is the handler func we are wrapping with middlewares
func f(w http.ResponseWriter, r *http.Request) {
        // get the id from the context
        id, err := backdrop.Get(r, "id")
        if err != nil {
                fmt.Println("err: ", err.Error())
        }
        fmt.Printf("request id is: %v\n", id)
        // you can also get the entire context if you are more comfortable with that
        ctx := backdrop.GetContext(r)
        ctx = context.WithValue(ctx, "key", "value")
        // and setting the newly created context in backdrop
        backdrop.SetContext(r, ctx)
}
