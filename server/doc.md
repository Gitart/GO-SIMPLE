# Go’s http package by example
Go’s http package has turned into one of my favorite things about the Go programming language. Initially it appears to be somewhat complex, but in reality it can be broken down into a couple of simple components that are extremely flexible in how they can be used. This guide will cover the basic ideas behind the http package, as well as examples in using, testing, and composing apps built with it.

This guide assumes you have some basic knowledge of what an interface in Go is, and some idea of how HTTP works and what it can do.

### Handler
The building block of the entire http package is the http.Handler interface, which is defined as follows:

```golang
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

Once implemented the http.Handler can be passed to http.ListenAndServe, which will call the ServeHTTP method on every incoming request.

http.Request contains all relevant information about an incoming http request which is being served by your http.Handler.

The http.ResponseWriter is the interface through which you can respond to the request. It implements the io.Writer interface, so you can use methods like fmt.Fprintf to write a formatted string as the response body, or ones like io.Copy to write out the contents of a file (or any other io.Reader). The response code can be set before you begin writing data using the WriteHeader method.

Here’s an example of an extremely simple http server:

```golang
package main

import (
	"fmt"
	"log"
	"net/http"
)

type helloHandler struct{}

func (h helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, you've hit %s\n", r.URL.Path)
}

func main() {
	err := http.ListenAndServe(":9999", helloHandler{})
	log.Fatal(err)
}
```

http.ListenAndServe serves requests using the handler, listening on the given address:port. It will block unless it encounters an error listening, in which case we log.Fatal.

Here’s an example of using this handler with curl:

```
 ~ $ curl localhost:9999/foo/bar
 hello, you've hit /foo/bar
```
 
HandlerFunc
Often defining a full type to implement the http.Handler interface is a bit overkill, especially for extremely simple ServeHTTP functions like the one above. The http package provides a helper function, http.HandlerFunc, which wraps a function which has the signature func(w http.ResponseWriter, r *http.Request), returning an http.Handler which will call it in all cases.

The following behaves exactly like the previous example, but uses http.HandlerFunc instead of defining a new type.

```golang
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, you've hit %s\n", r.URL.Path)
	})

	err := http.ListenAndServe(":9999", h)
	log.Fatal(err)
}
```

## ServeMux
On their own, the previous examples don’t seem all that useful. If we wanted to have different behavior for different endpoints we would end up with having to parse path strings as well as numerous if or switch statements. Luckily we’re provided with http.ServeMux, which does all of that for us. Here’s an example of it being used:

```golang
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	h := http.NewServeMux()

	h.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, you hit foo!")
	})

	h.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, you hit bar!")
	})

	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "You're lost, go home")
	})

	err := http.ListenAndServe(":9999", h)
	log.Fatal(err)
}
```

The http.ServeMux is itself an http.Handler, so it can be passed into http.ListenAndServe. When it receives a request it will check if the request’s path is prefixed by any of its known paths, choosing the longest prefix match it can find. We use the / endpoint as a catch-all to catch any requests to unknown endpoints. Here’s some examples of it being used:

```
 ~ $ curl localhost:9999/foo
Hello, you hit foo!

 ~ $ curl localhost:9999/bar
Hello, you hit bar!

 ~ $ curl localhost:9999/baz
You're lost, go home
```

http.ServeMux has both Handle and HandleFunc methods. These do the same thing, except that Handle takes in an http.Handler while HandleFunc merely takes in a function, implicitly wrapping it just as http.HandlerFunc does.

## Other muxes
There are numerous replacements for http.ServeMux like gorilla/mux which give you things like automatically pulling variables out of paths, easily asserting what http methods are allowed on an endpoint, and more. Most of these replacements will implement http.Handler like http.ServeMux does, and accept http.Handlers as arguments, and so are easy to use in conjunction with the rest of the things I’m going to talk about in this post.

### Composability
When I say that the http package is composable I mean that it is very easy to create re-usable pieces of code and glue them together into a new working application. The http.Handler interface is the way all pieces communicate with each other. Here’s an example of where we use the same http.Handler to handle multiple endpoints, each slightly differently:

```golang
package main

import (
	"fmt"
	"log"
	"net/http"
)

type numberDumper int

func (n numberDumper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here's your number: %d\n", n)
}

func main() {
	h := http.NewServeMux()

	h.Handle("/one", numberDumper(1))
	h.Handle("/two", numberDumper(2))
	h.Handle("/three", numberDumper(3))
	h.Handle("/four", numberDumper(4))
	h.Handle("/five", numberDumper(5))

	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "That's not a supported number!")
	})

	err := http.ListenAndServe(":9999", h)
	log.Fatal(err)
}
```

numberDumper implements http.Handler, and can be passed into the http.ServeMux multiple times to serve multiple endpoints. Here’s it in action:

```
 ~ $ curl localhost:9999/one
Here's your number: 1
 ~ $ curl localhost:9999/five
Here's your number: 5
 ~ $ curl localhost:9999/bazillion
That's not a supported number!
Testing
```

Testing http endpoints is extremely easy in Go, and doesn’t even require you to actually listen on any ports! The httptest package provides a few handy utilities, including NewRecorder which implements http.ResponseWriter and allows you to effectively make an http request by calling ServeHTTP directly. Here’s an example of a test for our previously implemented numberDumper, commented with what exactly is happening:

```golang
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	. "testing"
)

func TestNumberDumper(t *T) {
	// We first create the http.Handler we wish to test
	n := numberDumper(1)

	// We create an http.Request object to test with. The http.Request is
	// totally customizable in every way that a real-life http request is, so
	// even the most intricate behavior can be tested
	r, _ := http.NewRequest("GET", "/one", nil)

	// httptest.Recorder implements the http.ResponseWriter interface, and as
	// such can be passed into ServeHTTP to receive the response. It will act as
	// if all data being given to it is being sent to a real client, when in
	// reality it's being buffered for later observation
	w := httptest.NewRecorder()

	// Pass in our httptest.Recorder and http.Request to our numberDumper. At
	// this point the numberDumper will act just as if it was responding to a
	// real request
	n.ServeHTTP(w, r)

	// httptest.Recorder gives a number of fields and methods which can be used
	// to observe the response made to our request. Here we check the response
	// code
	if w.Code != 200 {
		t.Fatalf("wrong code returned: %d", w.Code)
	}

	// We can also get the full body out of the httptest.Recorder, and check
	// that its contents are what we expect
	body := w.Body.String()
	if body != fmt.Sprintf("Here's your number: 1\n") {
		t.Fatalf("wrong body returned: %s", body)
	}
}
```

In this way it’s easy to create tests for your individual components that you are using to build your application, keeping the tests near to the functionality they’re testing.

Note: if you ever do need to spin up a test server in your tests, httptest also provides a way to create a server listening on a random open port for use in tests as well.

## Middleware
Serving endpoints is nice, but often there’s functionality you need to run for every request before the actual endpoint’s handler is run. For example, access logging. A middleware component is one which implements http.Handler, but will actually pass the request off to another http.Handler after doing some set of actions. The http.ServeMux we looked at earlier is actually an example of middleware, since it passes the request off to another http.Handler for actual processing. Here’s an example of our previous example with some logging middleware:

```golang
package main

import (
	"fmt"
	"log"
	"net/http"
)

type numberDumper int

func (n numberDumper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here's your number: %d\n", n)
}

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s requested %s", r.RemoteAddr, r.URL)
		h.ServeHTTP(w, r)
	})
}

func main() {
	h := http.NewServeMux()

	h.Handle("/one", numberDumper(1))
	h.Handle("/two", numberDumper(2))
	h.Handle("/three", numberDumper(3))
	h.Handle("/four", numberDumper(4))
	h.Handle("/five", numberDumper(5))

	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "That's not a supported number!")
	})

	hl := logger(h)

	err := http.ListenAndServe(":9999", hl)
	log.Fatal(err)
}
```

logger is a function which takes in an http.Handler called h, and returns a new http.Handler which, when called, will log the request it was called with and then pass off its arguments to h. To use it we pass in our http.ServeMux, so all incoming requests will first be handled by the logging middleware before being passed to the http.ServeMux.

Here’s an example log entry which is output when the /five endpoint is hit:
```
2015/06/30 20:15:41 [::1]:34688 requested /five
Middleware chaining
```

Being able to chain middleware together is an incredibly useful ability which we get almost for free, as long as we use the signature func(http.Handler) http.Handler. A middleware component returns the same type which is passed into it, so simply passing the output of one middleware component into the other is sufficient.

However, more complex behavior with middleware can be tricky. For instance, what if you want a piece of middleware which takes in a parameter upon creation? Here’s an example of just that, with a piece of middleware which will set a header and its value for all requests:

```golang
package main

import (
	"fmt"
	"log"
	"net/http"
)

type numberDumper int

func (n numberDumper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here's your number: %d\n", n)
}

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s requested %s", r.RemoteAddr, r.URL)
		h.ServeHTTP(w, r)
	})
}

type headerSetter struct {
	key, val string
	handler  http.Handler
}

func (hs headerSetter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(hs.key, hs.val)
	hs.handler.ServeHTTP(w, r)
}

func newHeaderSetter(key, val string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return headerSetter{key, val, h}
	}
}

func main() {
	h := http.NewServeMux()

	h.Handle("/one", numberDumper(1))
	h.Handle("/two", numberDumper(2))
	h.Handle("/three", numberDumper(3))
	h.Handle("/four", numberDumper(4))
	h.Handle("/five", numberDumper(5))

	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "That's not a supported number!")
	})

	hl := logger(h)
	hhs := newHeaderSetter("X-FOO", "BAR")(hl)

	err := http.ListenAndServe(":9999", hhs)
	log.Fatal(err)
}
```

And here’s the curl output:

```
 ~ $ curl -i localhost:9999/three
 HTTP/1.1 200 OK
 X-Foo: BAR
 Date: Wed, 01 Jul 2015 00:39:48 GMT
 Content-Length: 22
 Content-Type: text/plain; charset=utf-8
 Here's your number: 3
 ```
 

newHeaderSetter returns a function which accepts and returns an http.Handler. Calling that returned function with an http.Handler then gets you an http.Handler which will set the header given to newHeaderSetter before continuing on to the given http.Handler.

This may seem like a strange way of organizing this; for this example the signature for newHeaderSetter could very well have looked like this:

func newHeaderSetter(key, val string, h http.Handler) http.Handler
And that implementation would have worked fine. But it would have been more difficult to compose going forward. In the next section I’ll show what I mean.

## Composing middleware with alice
Alice is a very simple and convenient helper for working with middleware using the function signature we’ve been using thusfar. Alice is used to create and use chains of middleware. Chains can even be appended to each other, giving even further flexibility. Here’s our previous example with a couple more headers being set, but also using alice to manage the added complexity.

```golang
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/justinas/alice"
)

type numberDumper int

func (n numberDumper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here's your number: %d\n", n)
}

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s requested %s", r.RemoteAddr, r.URL)
		h.ServeHTTP(w, r)
	})
}

type headerSetter struct {
	key, val string
	handler  http.Handler
}

func (hs headerSetter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(hs.key, hs.val)
	hs.handler.ServeHTTP(w, r)
}

func newHeaderSetter(key, val string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return headerSetter{key, val, h}
	}
}

func main() {
	h := http.NewServeMux()

	h.Handle("/one", numberDumper(1))
	h.Handle("/two", numberDumper(2))
	h.Handle("/three", numberDumper(3))
	h.Handle("/four", numberDumper(4))

	fiveHS := newHeaderSetter("X-FIVE", "the best number")
	h.Handle("/five", fiveHS(numberDumper(5)))

	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "That's not a supported number!")
	})

	chain := alice.New(
		newHeaderSetter("X-FOO", "BAR"),
		newHeaderSetter("X-BAZ", "BUZ"),
		logger,
	).Then(h)

	err := http.ListenAndServe(":9999", chain)
	log.Fatal(err)
}
```

In this example all requests will have the headers X-FOO and X-BAZ set, 
but the /five endpoint will also have the X-FIVE header set.

## Fin
Starting with a simple idea of an interface, the http package allows us to create for ourselves 
an incredibly useful and flexible (yet still rather simple) ecosystem for building web apps with re-usable components,
all without breaking our static checks.
