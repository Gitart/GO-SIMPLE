# Working with Redis in Go
## 26th February 2016

In this post I'm going to be looking at using Redis as a data persistence layer for a Go application. We'll start by 
explaining a few of the essential concepts, and then build a working web application which highlights some techniques 
for using Redis in a concurrency-safe way.  
This post assumes a basic knowledge of Redis itself (and a working installation, if you want to follow along).   
If you haven't used Redis before, I highly recommend reading the Little Book of Redis by Karl Seguin or running  
through the Try Redis interactive tutorial.  

## Installing a driver

First up we need to install a Go driver (or client) for Redis.   
A list of available drivers is located at http://redis.io/clients#go.   

Throughout this post we'll be using the Radix.v2 driver. It's well maintained, and I've found it's API
clean and straightforward to use. If you're following along you'll need to go get it:


$ go get github.com/mediocregopher/radix.v2

Notice that the Radix.v2 package is broken up into 6 sub-packages (cluster, pool, pubsub,  redis, sentinel and util). 
To begin with we'll only need the functionality in the redis package, so our import statements should look like:

```golang
import (
    "github.com/mediocregopher/radix.v2/redis"
)
```

Getting started with Radix.v2 and Go

As an example, let's say that we have an online record shop and want to store information about the albums for sale in Redis.
There's many different ways we could model this data in Redis, but we'll keep things simple and store each album as a hash – with fields for title, artist, price and the number of 'likes' that it has. As the key for each album hash we'll use the pattern album:{id}, where id is a unique integer value.
So if we wanted to store a new album using the Redis CLI, we could execute a HMSET command along the lines of:

```golang
127.0.0.1:6379> HMSET album:1 title "Electric Ladyland" artist "Jimi Hendrix" price 4.95 likes 8
OK
```

To do the same thing from a Go application, we need to combine a couple of functions from the Radix.v2 redis package.
The first is the Dial() function, which returns a new connection (or in Radix.v2 terms, client) to our Redis server.
The second is the client.Cmd() method, which sends a command to our Redis server across the connection. 
This always returns a pointer to a Resp object, which holds the reply from our command (or any error message if it didn't work).

It's quite straightforward in practice:

File: main.go

```golang
package main

import (
    "fmt"
    // Import the Radix.v2 redis package.
    "github.com/mediocregopher/radix.v2/redis"
    "log"
)

func main() {
    // Establish a connection to the Redis server listening on port 6379 of the
    // local machine. 6379 is the default port, so unless you've already
    // changed the Redis configuration file this should work.
    conn, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        log.Fatal(err)
    }
    // Importantly, use defer to ensure the connection is always properly
    // closed before exiting the main() function.
    defer conn.Close()

    // Send our command across the connection. The first parameter to Cmd()
    // is always the name of the Redis command (in this example HMSET),
    // optionally followed by any necessary arguments (in this example the
    // key, followed by the various hash fields and values).
    resp := conn.Cmd("HMSET", "album:1", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
    // Check the Err field of the *Resp object for any errors.
    if resp.Err != nil {
        log.Fatal(resp.Err)
    }

    fmt.Println("Electric Ladyland added!")
}
```

In this example we're not really interested in the reply from Redis (all successful HMSET commands just reply with the string "OK")
so we don't do anything with the *Resp object apart from checking it for any errors.
In such cases, it's common to chain the check against Err like so:

```golang
err = conn.Cmd("HMSET", "album:1", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8).Err
if err != nil {
    log.Fatal(err)
}
```

### Working with replies

When we are interested in the reply from Redis, the Resp object comes with some useful helper functions for    
converting the reply into a Go type we can easily work with.     

These are:

```
Resp.Bytes()     – converts a single reply to a byte slice ([]byte)
Resp.Float64()   – converts a single reply to a Float64
Resp.Int()       – converts a single reply to a int
Resp.Int64()     – converts a single reply to a int64
Resp.Str()       – converts a single reply to a string
Resp.Array()     – converts an array reply to an slice of individual Resp objects ([]*Resp)
Resp.List()      – converts an array reply to an slice of strings ([]string)
Resp.ListBytes() – converts an array reply to an slice of byte slices ([][]byte)
Resp.Map()       – converts an array reply to a map of strings, using each item in the array    
                   reply alternately as the keys and values for the map (map[string]string)  
```

Let's use some of these in conjunction with the HGET command to retrieve information from one of our album hashes:

File: main.go

```golang
package main

import (
    "fmt"
    "github.com/mediocregopher/radix.v2/redis"
    "log"
)

func main() {
    conn, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    // Issue a HGET command to retrieve the title for a specific album, and use
    // the Str() helper method to convert the reply to a string.
    title, err := conn.Cmd("HGET", "album:1", "title").Str()
    if err != nil {
        log.Fatal(err)
    }

    // Similarly, get the artist and convert it to a string.
    artist, err := conn.Cmd("HGET", "album:1", "artist").Str()
    if err != nil {
        log.Fatal(err)
    }

    // And the price as a float64...
    price, err := conn.Cmd("HGET", "album:1", "price").Float64()
    if err != nil {
        log.Fatal(err)
    }

    // And the number of likes as an integer.
    likes, err := conn.Cmd("HGET", "album:1", "likes").Int()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%s by %s: £%.2f [%d likes]\n", title, artist, price, likes)
}
```

It's worth pointing out that, when we use these helper methods, the error they return could relate to one of two things:    
either the failed execution of the command (as stored in the Resp object's Err field), or the conversion of the reply data   
to the desired type (for example, we'd get an error if we tried to convert the reply "Jimi Hendrix" to a Float64).   
There's no way of knowing which kind of error it is unless we examine the error message.   

If you run the code above you should get output which looks like:   

```
$ go run main.go
```

Electric Ladyland by Jimi Hendrix: £4.95 [8 likes]
Let's now look at a more complete example, where we use the HGETALL command to retrieve all fields    
from an album hash in one go and store the information in a custom Album struct.    

File: main.go

```golang
package main

import (
    "fmt"
    "github.com/mediocregopher/radix.v2/redis"
    "log"
    "strconv"
)

// Define a custom struct to hold Album data.
type Album struct {
    Title  string
    Artist string
    Price  float64
    Likes  int
}

func main() {
    conn, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    // Fetch all album fields with the HGETALL command. Because HGETALL
    // returns an array reply, and because the underlying data structure in
    // Redis is a hash, it makes sense to use the Map() helper function to
    // convert the reply to a map[string]string.
    reply, err := conn.Cmd("HGETALL", "album:1").Map()
    if err != nil {
        log.Fatal(err)
    }

    // Use the populateAlbum helper function to create a new Album object from
    // the map[string]string.
    ab, err := populateAlbum(reply)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(ab)
}

// Create, populate and return a pointer to a new Album struct, based on data
// from a map[string]string.
func populateAlbum(reply map[string]string) (*Album, error) {
    var err error
    ab := new(Album)
    ab.Title = reply["title"]
    ab.Artist = reply["artist"]
    // We need to use the strconv package to convert the 'price' value from a
    // string to a float64 before assigning it.
    ab.Price, err = strconv.ParseFloat(reply["price"], 64)
    if err != nil {
        return nil, err
    }
    // Similarly, we need to convert the 'likes' value from a string to an
    // integer.
    ab.Likes, err = strconv.Atoi(reply["likes"])
    if err != nil {
        return nil, err
    }
    return ab, nil
}
```

Running this code should give an output like:

```
$ go run main.go
&{Electric Ladyland Jimi Hendrix 4.95 8}
Using in a web application
```

One important thing to know about Radix.v2 is that the redis package (which we've used so far) is not safe for concurrent use.   
If we want to access a single Redis server from multiple goroutines, as we would in a web application, we must use 
the pool package instead. This allows us to establish a pool of Redis connections and each time we want to use a connection
we fetch it from the pool, execute our command on it, and return it too the pool.    

We'll illustrate this in a simple web application, building on the online record store example we've already used.    
Our finished app will support 3 functions:

```
Method	Path	Function

GET	/album?id=1	Show details of a specific album (using the id provided in the query string)
POST	/like	Add a new like for a specific album (using the id provided in the request body)
GET	/popular	List the top 3 most liked albums in order
If you'd like to follow along, head on over into your Go workspace and create a basic application scaffold…
```

$ cd $GOPATH/src
$ mkdir -p recordstore/models
$ cd recordstore
$ touch main.go models/albums.go
$ tree
.
├── main.go
└── models
    └── albums.go

…And use the Redis CLI to add a few additional albums, along with a new likes sorted set.   
This sorted set will be used within the GET /popular route to help us quickly and efficiently retrieve    
the ids of albums with the most likes. Here's the commands to run:  

```
HMSET album:1 title "Electric Ladyland" artist "Jimi Hendrix" price 4.95 likes 8
HMSET album:2 title "Back in Black" artist "AC/DC" price 5.95 likes 3
HMSET album:3 title "Rumours" artist "Fleetwood Mac" price 7.95 likes 12
HMSET album:4 title "Nevermind" artist "Nirvana" price 5.95 likes 8
ZADD likes 8 1 3 2 12 3 8 4
```

We'll follow an MVC-ish pattern for our application and use the models/albums.go file for all our Redis-related logic.
In the models/albums.go file we'll use the init() function to establish a Redis connection pool on startup, and we'll   
repurpose the code we wrote earlier into a FindAlbum() function that we can use from our HTTP handlers.   


File: models/albums.go

```golang
package models

import (
    "errors"
    // Import the Radix.v2 pool package, NOT the redis package.
    "github.com/mediocregopher/radix.v2/pool"
    "log"
    "strconv"
)

// Declare a global db variable to store the Redis connection pool.
var db *pool.Pool

func init() {
    var err error
    // Establish a pool of 10 connections to the Redis server listening on
    // port 6379 of the local machine.
    db, err = pool.New("tcp", "localhost:6379", 10)
    if err != nil {
        log.Panic(err)
    }
}

// Create a new error message and store it as a constant. We'll use this
// error later if the FindAlbum() function fails to find an album with a
// specific id.
var ErrNoAlbum = errors.New("models: no album found")

type Album struct {
    Title  string
    Artist string
    Price  float64
    Likes  int
}

func populateAlbum(reply map[string]string) (*Album, error) {
    var err error
    ab := new(Album)
    ab.Title = reply["title"]
    ab.Artist = reply["artist"]
    ab.Price, err = strconv.ParseFloat(reply["price"], 64)
    if err != nil {
        return nil, err
    }
    ab.Likes, err = strconv.Atoi(reply["likes"])
    if err != nil {
        return nil, err
    }
    return ab, nil
}

func FindAlbum(id string) (*Album, error) {
    // Use the connection pool's Get() method to fetch a single Redis
    // connection from the pool.
    conn, err := db.Get()
    if err != nil {
        return nil, err
    }
    // Importantly, use defer and the connection pool's Put() method to ensure
    // that the connection is always put back in the pool before FindAlbum()
    // exits.
    defer db.Put(conn)

    // Fetch the details of a specific album. If no album is found with the
    // given id, the map[string]string returned by the Map() helper method
    // will be empty. So we can simply check whether it's length is zero and
    // return an ErrNoAlbum message if necessary.
    reply, err := conn.Cmd("HGETALL", "album:"+id).Map()
    if err != nil {
        return nil, err
    } else if len(reply) == 0 {
        return nil, ErrNoAlbum
    }

    return populateAlbum(reply)
}
```

Something worth elaborating on is the pool.New() function. In the above code we specify a pool size of 10,     
which simply limits the number of idle connections waiting in the pool to 10 at any one time.    
If all 10 connections are in use when an additional pool.Get() call is made a new connection will be created on the fly.    

When you're only issuing one command on a connection, like in the FindAlbum() function above, it's possible to     
use the pool.Cmd() shortcut. This will automatically get a new connection from the pool, execute a given command,      
and then put the connection back in the pool.

Here's the FindAlbum() function re-written to use this shortcut:

```golang
func FindAlbum(id string) (*Album, error) {
    reply, err := db.Cmd("HGETALL", "album:"+id).Map()
    if err != nil {
        return nil, err
    } else if len(reply) == 0 {
        return nil, ErrNoAlbum
    }

    return populateAlbum(reply)
}
```

Alright, let's head over to the main.go file and set up a simple web server and HTTP handler for the GET /album route.    

File: main.go

```golang
package main

import (
    "fmt"
    "net/http"
    "recordstore/models"
    "strconv"
)

func main() {
    // Use the showAlbum handler for all requests with a URL path beginning
    // '/album'.
    http.HandleFunc("/album", showAlbum)
    http.ListenAndServe(":3000", nil)
}

func showAlbum(w http.ResponseWriter, r *http.Request) {
    // Unless the request is using the GET method, return a 405 'Method Not
    // Allowed' response.
    if r.Method != "GET" {
        w.Header().Set("Allow", "GET")
        http.Error(w, http.StatusText(405), 405)
        return
    }

    // Retrieve the id from the request URL query string. If there is no id
    // key in the query string then Get() will return an empty string. We
    // check for this, returning a 400 Bad Request response if it's missing.
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, http.StatusText(400), 400)
        return
    }
    // Validate that the id is a valid integer by trying to convert it,
    // returning a 400 Bad Request response if the conversion fails.
    if _, err := strconv.Atoi(id); err != nil {
        http.Error(w, http.StatusText(400), 400)
        return
    }

    // Call the FindAlbum() function passing in the user-provided id. If
    // there's no matching album found, return a 404 Not Found response. In
    // the event of any other errors, return a 500 Internal Server Error
    // response.
    bk, err := models.FindAlbum(id)
    if err == models.ErrNoAlbum {
        http.NotFound(w, r)
        return
    } else if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }

    // Write the album details as plain text to the client.
    fmt.Fprintf(w, "%s by %s: £%.2f [%d likes] \n", bk.Title, bk.Artist, bk.Price, bk.Likes)
}
```


If you run the application:

$ go run main.go
And make a request for one of the albums using cURL. You should get a response like:

```
$ curl -i localhost:3000/album?id=2
HTTP/1.1 200 OK
Content-Length: 42
Content-Type: text/plain; charset=utf-8

Back in Black by AC/DC: £5.95 [3 likes]
Using transactions
```

### The second route, POST /likes, is quite interesting.

When a user likes an album we need to issue two distinct commands: a HINCRBY to increment the likes field in the album hash,    
and a ZINCRBY to increment the relevant score in our  likes sorted set.

This creates a problem. Ideally we would want both keys to be incremented at exactly the same time as a single atomic action.    
Having one key updated after the other opens up the potential for data races to occur.    

The solution to this is to use Redis transactions, which let us run multiple commands together as an atomic group.    
To do this we use the MULTI command to start a transaction, followed by the commands (in our case a HINCRBY and ZINCRBY),    
and finally the EXEC command (which then executes our both our commands together as an atomic group).   

Let's create a new IncrementLikes() function in the albums model which uses this technique.    

File: models/albums.go
...

```golang
func IncrementLikes(id string) error {
    conn, err := db.Get()
    if err != nil {
        return err
    }
    defer db.Put(conn)

    // Before we do anything else, check that an album with the given id
    // exists. The EXISTS command returns 1 if a specific key exists
    // in the database, and 0 if it doesn't.
    exists, err := conn.Cmd("EXISTS", "album:"+id).Int()
    if err != nil {
        return err
    } else if exists == 0 {
        return ErrNoAlbum
    }

    // Use the MULTI command to inform Redis that we are starting a new
    // transaction.
    err = conn.Cmd("MULTI").Err
    if err != nil {
        return err
    }

    // Increment the number of likes in the album hash by 1. Because it
    // follows a MULTI command, this HINCRBY command is NOT executed but
    // it is QUEUED as part of the transaction. We still need to check
    // the reply's Err field at this point in case there was a problem
    // queueing the command.
    err = conn.Cmd("HINCRBY", "album:"+id, "likes", 1).Err
    if err != nil {
        return err
    }
    // And we do the same with the increment on our sorted set.
    err = conn.Cmd("ZINCRBY", "likes", 1, id).Err
    if err != nil {
        return err
    }

    // Execute both commands in our transaction together as an atomic group.
    // EXEC returns the replies from both commands as an array reply but,
    // because we're not interested in either reply in this example, it
    // suffices to simply check the reply's Err field for any errors.
    err = conn.Cmd("EXEC").Err
    if err != nil {
        return err
    }
    return nil
}
```

We'll also update the main.go file to add an addLike() handler for the route:


```golang
File: main.go
func main() {
    http.HandleFunc("/album", showAlbum)
    http.HandleFunc("/like", addLike)
    http.ListenAndServe(":3000", nil)
}
...
func addLike(w http.ResponseWriter, r *http.Request) {
    // Unless the request is using the POST method, return a 405 'Method Not
    // Allowed' response.
    if r.Method != "POST" {
        w.Header().Set("Allow", "POST")
        http.Error(w, http.StatusText(405), 405)
        return
    }

    // Retreive the id from the POST request body. If there is no parameter
    // named "id" in the request body then PostFormValue() will return an
    // empty string. We check for this, returning a 400 Bad Request response
    // if it's missing.
    id := r.PostFormValue("id")
    if id == "" {
        http.Error(w, http.StatusText(400), 400)
        return
    }
    // Validate that the id is a valid integer by trying to convert it,
    // returning a 400 Bad Request response if the conversion fails.
    if _, err := strconv.Atoi(id); err != nil {
        http.Error(w, http.StatusText(400), 400)
        return
    }

    // Call the IncrementLikes() function passing in the user-provided id. If
    // there's no album found with that id, return a 404 Not Found response.
    // In the event of any other errors, return a 500 Internal Server Error
    // response.
    err := models.IncrementLikes(id)
    if err == models.ErrNoAlbum {
        http.NotFound(w, r)
        return
    } else if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }

    // Redirect the client to the GET /ablum route, so they can see the
    // impact their like has had.
    http.Redirect(w, r, "/album?id="+id, 303)
}
```

If you make a POST request to like one of the albums you should now get a response like:

```
$ curl -i -L -d "id=2" localhost:3000/like
HTTP/1.1 303 See Other
Location: /album?id=2
Date: Thu, 25 Feb 2016 17:08:19 GMT
Content-Length: 0
Content-Type: text/plain; charset=utf-8

HTTP/1.1 200 OK
Content-Length: 42
Content-Type: text/plain; charset=utf-8

Back in Black by AC/DC: £5.95 [4 likes]
```

### Using the Watch command

OK, on to our final route: GET /popular. This route will display the details of the top 3 albums with the most likes,     
so to facilitate this we'll create a FindTopThree() function in the  models/albums.go file. In this function we need to:   

Use the ZREVRANGE command to fetch the 3 album ids with the highest score (i.e. most likes) from our likes sorted set.  
Loop through the returned ids, using the HGETALL command to retrieve the details of each album and add them to a []*Album slice.   
Again, it's possible to imagine a race condition occurring here. If a second client happens to like an album at the exact moment      
between our ZREVRANGE command and the HGETALLs for all 3 albums being completed, our user could end up being sent wrong     
or mis-ordered data.  

The solution here is to use the Redis WATCH command in conjunction with a transaction.    
WATCH instructs Redis to monitor a specific key for any changes. If another client modifies our watched key between     
our WATCH instruction and our subsequent transaction's EXEC, the transaction will fail and return a nil reply.    
If no client changes the value before our EXEC, the transaction will complete as normal. We can execute our code in a     
loop until the transaction is successful.   


File: models/albums.go

```golang
package models

import (
    "errors"
    "github.com/mediocregopher/radix.v2/pool"
    // Import the Radix.v2 redis package (we need access to its Nil type).
    "github.com/mediocregopher/radix.v2/redis"
    "log"
    "strconv"
)
...
func FindTopThree() ([]*Album, error) {
    conn, err := db.Get()
    if err != nil {
        return nil, err
    }
    defer db.Put(conn)

    // Begin an infinite loop.
    for {
        // Instruct Redis to watch the likes sorted set for any changes.
        err = conn.Cmd("WATCH", "likes").Err
        if err != nil {
            return nil, err
        }

        // Use the ZREVRANGE command to fetch the album ids with the highest
        // score (i.e. most likes) from our 'likes' sorted set. The ZREVRANGE
        // start and stop values are zero-based indexes, so we use 0 and 2
        // respectively to limit the reply to the top three. Because ZREVRANGE
        // returns an array response, we use the List() helper function to
        // convert the reply into a []string.
        reply, err := conn.Cmd("ZREVRANGE", "likes", 0, 2).List()
        if err != nil {
            return nil, err
        }

        // Use the MULTI command to inform Redis that we are starting a new
        // transaction.
        err = conn.Cmd("MULTI").Err
        if err != nil {
            return nil, err
        }

        // Loop through the ids returned by ZREVRANGE, queuing HGETALL
        // commands to fetch the individual album details.
        for _, id := range reply {
            err := conn.Cmd("HGETALL", "album:"+id).Err
            if err != nil {
                return nil, err
            }
        }

        // Execute the transaction. Importantly, use the Resp.IsType() method
        // to check whether the reply from EXEC was nil or not. If it is nil
        // it means that another client changed the WATCHed likes sorted set,
        // so we use the continue command to re-run the loop.
        ereply := conn.Cmd("EXEC")
        if ereply.Err != nil {
            return nil, err
        } else if ereply.IsType(redis.Nil) {
            continue
        }

        // Otherwise, use the Array() helper function to convert the
        // transaction reply to an array of Resp objects ([]*Resp).
        areply, err := ereply.Array()
        if err != nil {
            return nil, err
        }

        // Create a new slice to store the album details.
        abs := make([]*Album, 3)

        // Iterate through the array of Resp objects, using the Map() helper
        // to convert the individual reply into a map[string]string, and then
        // the populateAlbum function to create a new Album object
        // from the map. Finally store them in order in the abs slice.
        for i, reply := range areply {
            mreply, err := reply.Map()
            if err != nil {
                return nil, err
            }
            ab, err := populateAlbum(mreply)
            if err != nil {
                return nil, err
            }
            abs[i] = ab
        }

        return abs, nil
    }
}
```

### Using this from our web application is nice and straightforward:

File: main.go

```golang
func main() {
    http.HandleFunc("/album", showAlbum)
    http.HandleFunc("/like", addLike)
    http.HandleFunc("/popular", listPopular)
    http.ListenAndServe(":3000", nil)
}
...
func listPopular(w http.ResponseWriter, r *http.Request) {
  // Unless the request is using the GET method, return a 405 'Method Not
  // Allowed' response.
  if r.Method != "GET" {
    w.Header().Set("Allow", "GET")
    http.Error(w, http.StatusText(405), 405)
    return
  }

  // Call the FindTopThree() function, returning a return a 500 Internal
  // Server Error response if there's any error.
  abs, err := models.FindTopThree()
  if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }

  // Loop through the 3 albums, writing the details as a plain text list
  // to the client.
  for i, ab := range abs {
    fmt.Fprintf(w, "%d) %s by %s: £%.2f [%d likes] \n", i+1, ab.Title, ab.Artist, ab.Price, ab.Likes)
  }
}
```

One note about WATCH: a key will remain WATCHed until either we either EXEC (or DISCARD) our     
transaction, or we manually call UNWATCH on the key. So calling EXEC, as we do in the above example,    
is sufficient and the likes sorted set will be automatically UNWATCHed.   

Making a request to the GET /popular route should now yield a response similar to:  

```curl
$ curl -i localhost:3000/popular
HTTP/1.1 200 OK
Content-Length: 147
Content-Type: text/plain; charset=utf-8

1) Rumours by Fleetwood Mac: £7.95 [12 likes]
2) Nevermind by Nirvana: £5.95 [8 likes]
3) Electric Ladyland by Jimi Hendrix: £4.95 [8 likes]
If you found this post useful you might like the book I'm writing: Building real-world 
```
