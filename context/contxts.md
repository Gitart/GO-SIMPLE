# Go's net/context and http.Handler

The approaches in this post are now obsolete thanks to Go 1.7, which adds the context package to the standard library and uses it in the net/http *http.Request type. The background info here may still be helpful, but I wrote a follow-up post that revisits things for Go 1.7 and beyond.

A summary of this post is available in Japanese thanks to @craftgear. 
The golang.org/x/net/context package (hereafter referred as net/context although it’s not yet in the standard library) is a wonderful tool for the Go programmer’s toolkit. The blog post that introduced it shows how useful it is when dealing with external services and the need to cancel requests, set deadlines, and send along request-scoped key/value data.
The request-scoped key/value data also makes it very appealing as a means of passing data around through middleware and handlers in Go web servers. Most Go web frameworks have their own concept of context, although none yet use net/context directly.

Questions about using net/context for this kind of server-side context keep popping up on the /r/golang subreddit and the Gopher’s Slack community. Having recently ported a fairly large API surface from Martini to http.ServeMux and net/context, I hope this post can answer those questions.

## About http.Handler
The basic unit in Go’s HTTP server is its http.Handler interface, which is defined as:

```
type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
}
```

http.ResponseWriter is another simple interface and http.Request is a struct that contains data corresponding to the HTTP request, things like URL, headers, body if any, etc.

Notably, there’s no way to pass anything like a context.Context here.

## About context.Context
Much more detail about contexts can be found in the introductory blog post, but the main aspect I want to call attention to in this post is that contexts are derived from other contexts. Context values become arranged as a tree, and you only have access to values set on your context or one of its ancestor nodes.  


For example, let’s take context.Background() as the root of the tree, and derive a new context by attaching the content of the X-Request-ID HTTP header.
```
type key int
const requestIDKey key = 0

func newContextWithRequestID(ctx context.Context, req *http.Request) context.Context {
    return context.WithValue(ctx, requestIDKey, req.Header.Get("X-Request-ID"))
}

func requestIDFromContext(ctx context.Context) string {
    return ctx.Value(requestIDKey).(string)
}

ctx := context.Background()
ctx = newContextWithRequestID(ctx, req)
```

This derived context is the one we would then pass to the next layer of the system. Perhaps that would create its own contexts with values, deadlines, or timeouts, or it could extract values we previously stored.

### Approaches
These approaches are now obsolete as of Go 1.7. Read my follow-up post that revisits this topic for Go 1.7 and beyond.
So, without direct support for net/context in the standard library, we have to find another way to get a context.Context into our handlers.
There are three basic approaches:

Use a global request-to-context mapping
Create a http.ResponseWriter wrapper struct
Create your own handler types
Let’s examine each.

Global request-to-context mapping
In this approach we create a global map of requests to contexts, and wrap our handlers in a middleware that handles the lifetime of the context associated with a request. This is the approach taken by Gorilla’s context package, although with its own context type rather than net/context.   

Because every HTTP request is processed in its own goroutine and Go’s maps are not safe for concurrent access for performance reasons, it is crucial that we protect all map accesses with a sync.Mutex. This also introduces lock contention among concurrently processed requests. Depending on your application and workload, this could become a bottleneck.   

In general, though, this approach works well for Gorilla’s context, because its context value is simply a map of key/value pairs. Our context is arranged like a tree, and it’s important that the map always hold a reference to the leaf. This places a burden on the programmer to manually update the pointer’s value as new contexts are derived.   

An example usage might look like this:

```
var cmap = map[*http.Request]*context.Context{}
var cmapLock sync.Mutex

// Note that we are returning a pointer to the context, not the
// context itself.
func contextFromRequest(req *http.Request) *context.Context {
    cmapLock.Lock()
    defer cmapLock.Unlock()
    return cmap[req]
}

// Necessary wrapper around all handlers.  Must be the first middleware.
func contextHandler(ctx context.Context, h http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
        ctx2 := ctx // make a copy of the root context reference
        cmapLock.Lock()
        cmap[req] = &ctx2
        cmapLock.Unlock()

        h.ServeHTTP(rw, req)

        cmapLock.Lock()
        delete(cmap, req)
        cmapLock.Unlock()
    })
}

func middleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
        ctxp := contextFromRequest(req)
        *ctxp = newContextWithRequestID(*ctxp, req)

        h.ServeHTTP(rw, req)
    })
}

func handler(rw http.ResponseWriter, req *http.Request) {
    ctxp := contextFromRequest(req)

    reqID := requestIDFromContext(*ctxp)
    fmt.Fprintf(rw, "Hello request ID %s\n", reqID)
}

func main() {
    h := contextHandler(context.Background(), middleware(http.HandlerFunc(handler)))
    http.ListenAndServe(":8080", h)
}
```


Dereferencing the context pointer and updating it by hand is ugly, tedious and error-prone, which is why I don’t recommend this approach.
Update: Good question on Reddit asking why use a pointer to a context.Context here. It’s not necessary, but if you don’t use a pointer you must modify the underlying map any time you derive a new context. Doing so greatly increases the lock contention problem, because you must now lock around the map any time you update the context for a request.

```
http.ResponseWriter wrapper struct
In this approach we create a new struct type that embeds an existing http.ResponseWriter and attaches additional functionality to it. This approach is often used by Go web frameworks to do things like capturing the status code for the purpose of logging it later. Like the first approach, you’ll need to wrap handlers in a middleware that wraps the http.ResponseWriter and passes it into subsequent middleware and your handler.

type contextResponseWriter struct {
    http.ResponseWriter
    ctx context.Context
}

func contextHandler(ctx context.Context, h http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
        crw := &contextResponseWriter{rw, ctx}
        h.ServeHTTP(crw, req)
    })
}

func middleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
        crw := rw.(*contextResponseWriter)
        crw.ctx = newContextWithRequestID(crw.ctx, req)

        h.ServeHTTP(rw, req)
    })
}

func handler(rw http.ResponseWriter, req *http.Request) {
    crw := rw.(*contextResponseWriter)

    reqID := requestIDFromContext(crw.ctx)
    fmt.Fprintf(rw, "Hello request ID %s\n", reqID)
}

func main() {
    h := contextHandler(context.Background(), middleware(http.HandlerFunc(handler)))
    http.ListenAndServe(":8080", h)
}
```


This approach just feels dirty to me. The context is associated with the request, not the response, so sticking it on http.ResponseWriter feels out of place. The ResponseWriter’s purpose is simply to give handlers a way to write data to the output socket.
Piggybacking on http.ResponseWriter requires a type assertion to your wrapper struct type before you can access the context. The details of this can be hidden away in a safe helper function, but it doesn’t hide the fact that the runtime assertion is necessary.
There is also another hidden downside. There is a concrete value (with a type internal to package net/http) underlying the http.ResponseWriter that is passed into your handler. That value also implements additional interfaces from the net/http package. If you simply wrap http.ResponseWriter, your wrapper will not be implementing these additional interfaces.
You must implement these interfaces with wrapper functions if you hope to match the base http.ResponseWriter’s functionality. In some cases, like http.Flusher, this is easy with a simple conditional type assertion:

```
func (crw *contextResponseWriter) Flush() {
    if f, ok := crw.ResponseWriter.(http.Flusher); ok {
        f.Flush()
    }
}
```

However, http.CloseNotifier is quite a bit harder. Its definition contains a method that returns a <-chan bool. That channel has certain semantics that existing code likely depends upon1. We have a couple different options here, none of them good:
Ignore the interface and don’t implement it, making the functionality unavailable even if the underlying http.ResponseWriter supports it.
Implement the interface and wrap to the underlying implementation. But what if the underlying http.ResponseWriter does not support this interface? We can’t guarantee the proper semantics of the API.
These are just two interfaces that the standard library implements today. This approach is not future-proof, because additional interfaces may be added to the standard library and implemented internally within net/http.
I don’t recommend this approach because of the interface issue, but if you’re ok with ignoring them, this is probably the simplest to implement.
Custom context handler types
In this approach, we eschew http.Handler for a new type of our own creation. This has obvious downsides: you cannot use existing de facto middleware or handlers without wrappers. Ultimately, though, I think this is the cleanest way to pass a context.Context around.
Let’s create a new ContextHandler type, following in the model of http.Handler. We’ll also create an analog to http.HandlerFunc.

```
type ContextHandler interface {
    ServeHTTPContext(context.Context, http.ResponseWriter, *http.Request)
}

type ContextHandlerFunc func(context.Context, http.ResponseWriter, *http.Request)

func (h ContextHandlerFunc) ServeHTTPContext(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
    h(ctx, rw, req)
}
Middleware can now derive new contexts from the one passed to the handler, and pass them onto the next middleware or handler in the chain.

func middleware(h ContextHandler) ContextHandler {
    return ContextHandlerFunc(func(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
        ctx = newContextWithRequestID(ctx, req)
        h.ServeHTTPContext(ctx, rw, req)
    })
}
```

The final context handler has access to all of the request data set by middleware above it.

```
func handler(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
    reqID := requestIDFromContext(ctx)
    fmt.Fprintf(rw, "Hello request ID %s\n", reqID)
}
```
The last trick is converting our ContextHandler into something that is http.Handler compatible, so we can use it anywhere standard handlers are used.

```
type ContextAdapter struct{
    ctx context.Context
    handler ContextHandler
}

func (ca *ContextAdapter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
    ca.handler.ServeHTTPContext(ca.ctx, rw, req)
}

func main() {
    h := &ContextAdapter{
        ctx: context.Background(),
        handler: middleware(ContextHandlerFunc(handler)),
    }
    http.ListenAndServe(":8080", h)
}
```
The ContextAdapter type also allows us to use existing http.Handler middleware, as long as they run before it does. 
Existing logging, panic recovery, and form validation middleware should all continue to work great with our 
context handlers plus our adapter.


This is my preferred method for integrating net/context with my server. 
I recently converted an approximately 30-route server from Martini to this method, 
and things are working great. The code is much cleaner, easier to follow, and performs better. 
This API service does both HTTP basic and OAuth authentication, passing along client and user 
information via contexts. It extracts request IDs that are passed across to other services via contexts. 
Context-aware middleware handles setting CORS headers, handling OPTIONS requests, recovering 
from panics by returning JSON-encoded errors, logging request and response info, and recording statsd metrics.

