# Using Synchronization Primitives in Go

## Exploring Mutex, WaitGroup, and Once with examples

[![Abhishek Gupta](https://miro.medium.com/fit/c/96/96/1*ZdSjzBPYoAeJU5cB5f2_xA.jpeg)](https://abhishek1987.medium.com/?source=post_page-----2e50359cb0a7--------------------------------)

[Abhishek Gupta](https://abhishek1987.medium.com/?source=post_page-----2e50359cb0a7--------------------------------)

[Follow](https://medium.com/m/signin?actionUrl=%2F_%2Fsubscribe%2Fuser%2F47a6bae243a3&operation=register&redirect=https%3A%2F%2Fbetterprogramming.pub%2Fusing-synchronization-primitives-in-go-mutex-waitgroup-once-2e50359cb0a7&source=post_page-47a6bae243a3----2e50359cb0a7---------------------follow_byline-----------)

[Oct 15, 2019](https://betterprogramming.pub/using-synchronization-primitives-in-go-mutex-waitgroup-once-2e50359cb0a7?source=post_page-----2e50359cb0a7--------------------------------) · 4 min read

![Image for post](https://miro.medium.com/max/60/1*nR7DEX5BpkkUOWF1ql0nAw.jpeg?q=20)

![Image for post](https://miro.medium.com/max/4926/1*nR7DEX5BpkkUOWF1ql0nAw.jpeg)

![Image for post](https://miro.medium.com/max/9852/1*nR7DEX5BpkkUOWF1ql0nAw.jpeg)

Photo by [Holly Mandarich](https://unsplash.com/@hollymandarich?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) on [Unsplash](https://unsplash.com/s/photos/go?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText)

Welcome to Just Enough Go!

This is the second post in a series of articles about the [Go programming language](https://golang.org/) in which I will be covering some of the most commonly used Go standard library packages, e.g. [encoding/json](https://golang.org/pkg/encoding/json/), [io](https://golang.org/pkg/io/), [net/http](https://golang.org/pkg/net/http/), [sync](https://golang.org/pkg/sync/), etc. I plan to keep these relatively short and example\-driven.

Let’s look at some of the lower\-level synchronization constructs which Go provides in the `[sync](https://godoc.org/sync)` [package](https://godoc.org/sync), in addition to Goroutines and channels. There are a bunch of them, but we will explore `WaitGroup`, `Mutex`, and `Once` with examples.

Code examples are [available on GitHub](https://github.com/abhirockzz/just-enough-go).

# WaitGroup

Use a `WaitGroup` for co\-ordination if your program needs to wait for a bunch of Goroutines to finish. It is similar to a `CountDownLatch` in Java. Let's see an example.

We want to print all the files in our home directory in parallel. Use a `WaitGroup` to specify the number of tasks/Goroutines to wait for.

In this case, it is the same as the number of files/directories you have in the home directory. We use `Wait()` to block until the `WaitGroup` counter becomes zero.

...
func main() {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        panic(err)
    }
    filesInHomeDir, err := ioutil.ReadDir(homeDir)
    if err != nil {
        panic(err)
    }
    var wg sync.WaitGroup
    wg.Add(len(filesInHomeDir))
    for \_, file := range filesInHomeDir {
        go func(f os.FileInfo) {
            defer wg.Done()
        }(file)
    }
    wg.Wait()
}
...

To run this program:

curl https://raw.githubusercontent.com/abhirockzz/just\-enough\-go/master/sync/wait\-group\-example.go \-o wait\-group\-example.go
go run wait\-group\-example.go

A Goroutine is spawned for each `os.FileInfo` we find in the user home directory and once we print its name, the counter is decremented using `Done`. The program exits after all the contents of the home directory are covered.

# Mutex

A `Mutex` is a shared lock that you can use to provide exclusive access to certain parts of your code. In this simple example, we have a shared/global variable `accessCount` which is used in the `incr` function.

func incr() {
    mu.Lock()
    defer mu.Unlock()
    accessCount = accessCount + 1
}

Notice that the `incr` function is protected by a `Mutex`. Thus, only a single Goroutine can access it at a time. We throw multiple Goroutines at it.

loop := 500
for i := 1; i <= loop; i++ {
        go func(c int) {
            wg.Add(1)
            defer wg.Done()
            incr()
        }(i)
}

If you run this, you will always get the same result, i.e. `Final = 500` (since the for loop runs for 500 iterations). To run the program:

curl https://raw.githubusercontent.com/abhirockzz/just\-enough\-go/master/sync/mutex\-example.go \-o mutex\-example.go
go run mutex\-example.go

Comment (or remove) the following lines in the `incr` function and run the program on your local machine and run the program again:

mu.Lock()
defer mu.Unlock()

You will notice variable results e.g. `Final = 474`.

I encourage you to read up on `[RWMutex](https://golang.org/pkg/sync/#RWMutex)`. It is a special kind of lock that can be used to allow concurrent reads but synchronized (single writer) writes.

# Once

It allows you to define a task which you only want to execute once during the lifetime of your program.

This is very useful for `Singleton`\-like behavior. It has a single `Do` function that lets you pass another function which you intend to execute only once. Let's look at an example.

Say you’re building a REST API using the Go `net/http` package and you want a piece of code to be executed only when the HTTP handler is called (e.g. a get a DB connection).

You can wrap that code with `once.Do` and rest assured that it will be only run when the handler is invoked for the first time.

Here is a function that we want to be executed only once:

func oneTimeOp() {
    fmt.Println("one time op start")
    time.Sleep(3 \* time.Second)
    fmt.Println("one time op started")
}

This is what we do within our HTTP handler — notice `once.Do(oneTimeOp)`.

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r \*http.Request) {
        fmt.Println("http handler start")
        once.Do(oneTimeOp)
        fmt.Println("http handler end")
        w.Write(\[\]byte("done!"))
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}

Run the code and access the REST endpoint.

curl https://raw.githubusercontent.com/abhirockzz/just\-enough\-go/master/sync/once\-example.go \-o once\-example.go
go run once\-example.go

From a different terminal:

curl localhost:8080
//output \- done!

When you first access it, it will be a little slow in returning and you will see the following logs in the server:

http handler start
one time op start
one time op end
http handler end

If you run it again (any number of times), the function `oneTimeOp` will not be executed. Check the logs to confirm.

That’s all for this piece. I would be more than happy to take suggestions on specific Go topics that you would like to me cover.
