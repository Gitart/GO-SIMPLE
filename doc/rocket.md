# Go 1.11 Rocket Tutorial

Table of Contents

*   [

    Part 1 - Setup and Seam Carving

    ](https://getstream.io/blog/go-1-11-rocket-tutorial/#part-setup-and-seam-carving)
*   [

    Part 2 - Unsplash API, Channels and Concurrency

    ](https://getstream.io/blog/go-1-11-rocket-tutorial/#part-unsplash-api-channels-and-concurrency)
*   [

    Step 5 - APIs, Structs and JSON

    ](https://getstream.io/blog/go-1-11-rocket-tutorial/#step-apis-structs-and-json)
*   [

    Step 6 - Understand Those API Calls

    ](https://getstream.io/blog/go-1-11-rocket-tutorial/#step-understand-those-api-calls)
*   [

    Step 7 - Pointers

    ](https://getstream.io/blog/go-1-11-rocket-tutorial/#step-pointers)
*   [

    Step 8 - Errors

    ](https://getstream.io/blog/go-1-11-rocket-tutorial/#step-errors)
*   [

    Step 9 - Channels & Goroutines

    ](https://getstream.io/blog/go-1-11-rocket-tutorial/#step-channels-goroutines)
*   [

    Step 10 - Libraries and Tools

    ](https://getstream.io/blog/go-1-11-rocket-tutorial/#step-libraries-and-tools)
*   [

    Concluding the Go Tutorial

    ](https://getstream.io/blog/go-1-11-rocket-tutorial/#concluding-the-go-tutorial)

---

•Published: Sep 18, 2018

Thierry S.

This tutorial combines two of my favorite things, the Go programming language and images of SpaceX rocket launches.

With Go rapidly picking up adoption in the developer community, its becoming one of the leading languages for building backend systems. Go’s performance is [similar to Java](https://benchmarksgame-team.pages.debian.net/benchmarksgame/faster/go.html) and [C++](https://benchmarksgame-team.pages.debian.net/benchmarksgame/faster/go-gpp.html), yet it’s almost as easy to write as Python. Since its inception in 2009 by Google, Go has been adopted by Netflix, Stripe, Twitch, Digital Ocean and many others. Here at [Stream](https://getstream.io/) we use Go & [RocksDB](https://rocksdb.org/) to power news feeds for over 300 million end users.

In this tutorial, you’ll build a highly concurrent web server which reads [SpaceX](https://twitter.com/SpaceX) images from the  [Unsplash API](https://unsplash.com/developers) and applies [seam carving](https://en.wikipedia.org/wiki/Seam_carving) on those images. The pace is high and assumes you already know how to program. You’ll [learn the basics of Go](https://getstream.io/blog/topic/tutorials/feeds/go/) and we’ll also cover more advanced topics such as Goroutines and Channels.

![](data:image/svg+xml;charset=utf-8,%3Csvg height='1628' width='1382' xmlns='http://www.w3.org/2000/svg' version='1.1'%3E%3C/svg%3E)

![](data:image/png;base64,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) ![](https://getstream.io/blog/static/c7a54e302f7af000f5b34b903fa6e81c/ff086/image1-4.png)

 ![](https://getstream.io/blog/static/c7a54e302f7af000f5b34b903fa6e81c/ff086/image1-4.png)

## Part 1 - Setup and Seam Carving

### Step 1 - Install Go 1.11

[Homebrew](https://brew.sh/) on macOS is the easiest way to install Go:

```bash
1brew install go
```

Alternatively head over to the [official Go download page](https://golang.org/dl/) and follow the [install instructions](https://golang.org/doc/install#install) for your system.

Be sure that you have installed **Go 1.11** or later before continuing this guide. Once installed, check your version with

```bash
1go version
```

### Step 2 - A Simple Request Handler

Go v1.11 simplified the initial setup of a Go project. Older versions of Go required the **GOPATH** environment variable. With v1.11 you can follow the three steps below to get your server up and running:

#### A. Create a directory

Create an empty directory called **rockets**:

```bash
1mkdir rockets
2cd rockets
```

#### B. Go mod init

Run the following command to start a new Go module:

```bash
1go mod init github.com/GetStream/rockets-go-tutorial
```

This will create a file called go.mod in your directory. This file will contain your dependencies. It is similar to package.json in Node, or requirements.txt in Python. (Go doesn’t have a package repository

#### C. main.go

Create a file called **main.go** in the rockets directory with the following content. (The full path is rockets/**main.go):**

```go
1package main
2import (
3    "bytes"
4    "fmt"
5    "io"
6    "log"
7    "net/http"
8    "github.com/esimov/caire"
9    "github.com/pkg/errors"
10)
11const (
12    IMAGE_URL string = "https://bit.ly/2QGPDkr"
13)
14func ContentAwareResize(url string) ([]byte, error) {
15    fmt.Printf("Download starting for url %s\n", url)
16    response, err := http.Get(url)
17    if err != nil {
18        return nil, errors.Wrap(err, "Failed to read the image")
19    }
20    defer response.Body.Close()
21    converted := &bytes.Buffer{}
22    fmt.Printf("Download complete %s", url)
23    shrinkFactor := 30
24    fmt.Printf("Resize in progress %s, shrinking width by %d percent...\n", url, shrinkFactor)
25    p := &caire.Processor{
26        NewWidth:   shrinkFactor,
27        Percentage: true,
28    }
29    err = p.Process(response.Body, converted)
30    if err != nil {
31        return nil, errors.Wrap(err, "Failed to apply seam carving to the image")
32    }
33    fmt.Printf("Seam carving completed for %s\n", url)
34    return converted.Bytes(), nil
35}
36func main() {
37    fmt.Println("Ready for liftoff! Checkout http://localhost:3000/occupymars")
38    http.HandleFunc("/occupymars", func(w http.ResponseWriter, r *http.Request) {
39        if r.URL.Query().Get("resize") > "" {
40            resized, err := ContentAwareResize(IMAGE_URL)
41            if err != nil {
42                http.Error(w, err.Error(), http.StatusInternalServerError)
43                return
44            }
45            w.Header().Set("Content-Type", "image/jpeg")
46            io.Copy(w, bytes.NewReader(resized))
47        } else {
48            fmt.Fprintf(w, "<html><div>Original image:</div> <img src=\"%s\" /><br/><a href=\"?resize=1\">Resize using Seam Carving</a></html>", IMAGE_URL)
49        }
50    })
51    log.Fatal(http.ListenAndServe(":3000", nil))
52}
53

```

You can run it like this:

```bash
1go run main.go
```

After a moment of installing dependencies, you’ll be able to visit [](http://localhost:3000/occupymars)[http://localhost:3000/occupymars](http://localhost:3000/occupymars)

Click the “Resize using Seam Carving” link to start the resizing. The computation is quite heavy and may take a bit, so give it a few seconds.

Voila, one resized rocket picture:

![](data:image/svg+xml;charset=utf-8,%3Csvg height='267' width='280' xmlns='http://www.w3.org/2000/svg' version='1.1'%3E%3C/svg%3E)

![](data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABQAAAATCAYAAACQjC21AAAACXBIWXMAAAsTAAALEwEAmpwYAAAFHElEQVQ4yx3O7U/bBQLA8d+r00vuNPE2WVDo3BABgcEKDFlZB4MOChstD22h7a88tDyUlra0QB9+fQTa0idaymC4TNQtO06j+BDvNNmb5eKpuUuWmPjGGN/4xlwuufgXfI2++Lz95ivYfFkW1nJYvTmmPRnG7Ul6xRhdhjDtEyHaJ4J0TAR+Jx/1c0mzRvOIj1q1G9nNFapVTqr6lqm6YUd2YwlB79hBs5BmyJpENbOF0pTgij6CXBemXSfRoQtx1RDm2mSYXmMYtSXK9UmJlpF16oa8XBx0c3HAxWtqF01DboSbc9vcXkgz6cxhcGYZXUwxYtvCYIsy69xiaDZBrznGwHSMm5YY/eYI3QaJLl2QjjE/baMbyEfXUBmDDJsDCHNruxjdBSyrBWxru+gdGXT2FLNLErqFOD1iAqU5Rp/lt2gclRhjcDrGrdkYWmsCzWzsd2qzhFqUEERfEfNqEdPqLubVAjpHhuH5NHZ3nBHrJiPzKUTXDjZvlqW1HCZHCpMjjcmRxLicxOnPo7eFEZ3baGc2EOb8e3jjByyFylh8RYyuPBZXnnhkB38ojydUYnE9x4z3t/Mk69ESO1sZ8pkib5f3iHi95GMSKxYR1cAEwmb2LqHkIZ7YHawbJWy+AmGpxEkxSzKaZ9aTZ9SeRmtPoTDFKJbv8e9P3uXLD47ZX7eT3vDwIB3C2t9N9/VbCFLqACl1SDh1SCBeJhot8mC3yP+++5JH+3dxeLMYXFk0y2kGF9OcnvyVrz5+yPvZEKcFiZNclIJrhpJvkVujFoTF8D5WqYzoL2H1F/Fs7PLozh3+//1THu0fMe/OonfliGwVeSe7zT8eHPHRvV3elFb4qLzJplVP3mEhYLXQdPUWgjVUwhXfwx3fY1Eq4wru8nE5ws/ffcPxnbvEozsEIjmCiV0+PMhyIK2yF/NTCHlZn55kvPMSK8NKTKpeWq5pEebDJbyJIvHtAonNHHubUU4CJn781xPeOjhgM5pgPZTifrnEZ0dpCmuLxO0z+KaNGPuvob7ciF7RhkMzyICiByG0XSCZ3ePx3+7z9/0Y70gLfH0Y5Id/PuFuqUw2lSGzU+Iwm+bR1iJHfpGM7TYr42o0Xa2MdV7iZmsjhl4FbzS/jpDJ5nlYTPL4uMCTtI2vDgP88N4OP31yxNMP7/P003c52I7gdno5LYc5jZvYMihwj6nQdrUgKlsZ62pFXlNNw4UKhFTYT9RmJL4sEtBcJ64b4Nht5HHGQ3LFRsrnYHVO5CQV5LOixMPAJIW5AYLmCWzD/ZiVcsavttFaW8Wr5ysRerpamLpxhbk+OTPKJvqbaphRtDB5XcG1pnpUVy4z1N2FoUvOuFLJyugge84p5rXDqDvbGW5vpE/eTK3sLK9f+AtCZcUzvHLuORpk57hc+xK1lS/Q8Wolmu4W5DUvU19dQb2sEnlDLX1tjcjrztNef4Hal1/klYoXeOnMc8jOPE9VxR+R1z2P0FBfQ7eiE4tOj1k7hlk9yNsbIj//55TM6jxGjYZh1QBzE+NE502s6lXYBt9g5OpllJfqaJZV0FxVSUfdWXo6nkWQVVegVPbgWbDjNIs4p6Yoz4/xy7df8F55G691FqNWi2/azHHYwaFnnMSkkgV1NyaVgtvtDdxolqG68mc6G/+EcPH8WTrbX8Nls7FimsKpn+CeY4L/Pv2cN7eCLItmjON6ZqZEsmtLHHr1RPQKZvvb0Ck76G2qpqP2DNUv/oHGmmf4FdxhRNi7l+r3AAAAAElFTkSuQmCC) ![](https://getstream.io/blog/static/debc412453ce9055a697fdb9413c7112/4cb93/image5-3.png)

 ![](https://getstream.io/blog/static/debc412453ce9055a697fdb9413c7112/4cb93/image5-3.png)

### Step 3 - Go’s Syntax

The first thing you’ll notice in the above example is that Go is very similar to other programming languages. Let’s quickly review the code in **main.go**:

#### A. Go is a statically typed language

```
func ContentAwareResize(url string) ([]byte, error) {
...
}
```

The function takes a string as input and returns a slice of bytes and an error.

#### B. Go infers types

```go
1shrinkFactor := 30
```

Is short for:

```go
1var shrinkFactor int = 30
```

Here’s the syntax for the most commonly used types in Go:

```go
1message := "hello world"
2friends := []string{"john", "amy"}
3friends = append(friends, "jack")
4populationByCity := map[string]int{"Boulder": 108090, "Amsterdam": 821752}
5populationByCity["Palo Alto"] = 67024
```

[LearnXinY](https://learnxinyminutes.com/docs/go/) is a great resource to quickly check the Go syntax.

#### C. Types are also inferred when calling functions

```go
1resized, err := ContentAwareResize(url)
```

Note how the function returns two elements which we use to initialize these new variables:  resized and err.

#### D. Go & Unused variables

One of the quirks of Go is that unlike most other languages it will throw a syntax error if you leave a variable or import unused. If you add the following code after the shrinkFactor on line 29:

```go
1oldvariable := 20
```

And run main.go you’ll get the following error:
./main.go:30:12: oldvariable declared and not used
Most editors will help you clean up unused code, so in practice this is not that annoying. If you want to ignore one of the variables returned by a function you can do it like this:

```go
1_, err := http.Get(url)
```

The \_ simply means discard the value.

#### E. Interfaces

On line 36 the p.Process function converts the image in response.Body and stores it in the converted variable. What’s interesting is the function definition of p.Process:

```go
1func (p *Processor) Process(r io.Reader, w io.Writer) error {
2    …
3}
```

This function takes an **io.Reader** as the first argument and an **io.Writer** as the second. **io.Reader** and **io.Writer** are interfaces. Any type that implements the methods required by these interfaces can be passed to the function.

*Note: One interesting fact about interfaces in Go is that types implement them implicitly. There is no explicit declaration of intent, no "implements" keyword. A type implements an interface by applying its methods.*

(Fun read: This blogpost about [Streaming IO in Go](https://medium.com/learning-the-go-programming-language/streaming-io-in-go-d93507931185) gives more details about the **io.Reader** and io.Writer interfaces)

#### F. Imports

```go
1import (
2    "bytes"
3    "fmt"
4    "io"
5    "log"
6    "net/http"
7    "github.com/esimov/caire"
8    "github.com/pkg/errors"
9)
```

The import syntax is easy to understand. Let’s refactor the ContentAwareResize function into a separate package to clean up our code.

### Step 4 - Refactoring to a Package

#### A. Create a directory with the name “seam”:

```bash
1mkdir seam
```

Your full path will look like **rockets/seam**.
Note: If you’re wondering why the folder is called seam it’s because this type of content aware image resizing is also typically called [seam carving](https://en.wikipedia.org/wiki/Seam_carving).

#### B. In the seam directory, create a file called “seam.go” with the following content:

```go
1package seam
2import (
3    "bytes"
4    "fmt"
5    "net/http"
6    "github.com/esimov/caire"
7    "github.com/pkg/errors"
8)
9func ContentAwareResize(url string) ([]byte, error) {
10    fmt.Printf("Download starting for url %s \n", url)
11    response, err := http.Get(url)
12    if err != nil {
13        return nil, errors.Wrap(err, "Failed to read the image")
14    }
15    defer response.Body.Close()
16    converted := &bytes.Buffer{}
17    fmt.Printf("Download complete %s \n", url)
18    shrinkFactor := 30
19    fmt.Printf("Resize in progress %s, shrinking width by %d percent... \n", url, shrinkFactor)
20    p := &caire.Processor{
21        NewWidth:   shrinkFactor,
22        Percentage: true,
23    }
24    err = p.Process(response.Body, converted)
25    if err != nil {
26        return nil, errors.Wrap(err, "Failed to apply seam carving to the image")
27    }
28    fmt.Printf("Seam carving completed for %s \n", url)
29    return converted.Bytes(), nil
30}
31

```

#### C. Open main.go and make the following changes:

1.  Remove the import for "github.com/esimov/caire"
2.  Remove the import for "github.com/pkg/errors"
3.  Add the import for "github.com/GetStream/rockets-go-tutorial/seam"
4.  Update the function call in main.go from ContentAwareResize to seam.ContentAwareResize (Line 50)
5.  Remove the ContentAwareResize function from main.go (Line 18-43)

To clarify, before the changes the function call looked like this:

```go
1resized, err := ContentAwareResize(url)
```

After the changes you’re calling the method defined in the seam module like this:

```go
1resized, err := seam.ContentAwareResize(url)
```

#### D. See if it worked:

Start the server and see if it still works!

```bash
1go run main.go
```

If everything went according to plan you should still be able to access at [](http://localhost:3000/occupymars)[http://localhost:3000/occupymars](http://localhost:3000/occupymars)
*Note: Packages give you a nice way to structure your code. Capitalized functions are public and lowerCase functions are private.*
If something went wrong with these steps you can also clone the [Github repo](https://github.com/GetStream/rockets-go-tutorial) for a working version:

## Part 2 - Unsplash API, Channels and Concurrency

## Step 5 - APIs, Structs and JSON

You’re already well on your way to learning Go. Easy right? Not bad for a language that’s 20-40 times faster than Python!
For this second example we’ll learn how to use read data from an API. We’ll use the excellent [Unsplash](https://unsplash.com/) photography API to search for more rocket pictures.

![](data:image/svg+xml;charset=utf-8,%3Csvg height='1735' width='1999' xmlns='http://www.w3.org/2000/svg' version='1.1'%3E%3C/svg%3E)

![](data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABQAAAARCAYAAADdRIy+AAAACXBIWXMAAAsTAAALEwEAmpwYAAAEb0lEQVQ4y62K6U/TBwCGf+oOZ3QTs+kU2ZzH0AYRCCJTTrEwK1KoHIqAKJcCIkwLKpdVDlu5ilUoVIocrRQQ0DK0CBQnEGUeKIq4LM7FLNmH/QOu7FnmlizZ573J8+F58gqvX79mZGSE6elpZmZmsNls/G6zYbPNYHvr/+Gf9vfnX/7ymZk/EN7Kmzf8XxMeTjymRqenyWii0WCi6Uonl5oMDN4epX/0CV39d+my3KV3aJwbQ/cYGJnAMjRKo6GLlrZr6A3dtJiu09jaxQ8/vkIoU19AWLAE+7XuLBdtYvl6L+bbryUm5Ruis7WI4/IIS1YgO1TKlgg50VkVyPbLmf/5Zj4VBeAticLBOYh5Dp60tvci1DW2ssrVF7/gvfiHxBEgS8DZW4q8UInW2EtJ2UVOq2rIVupJz62gSGMkQa7iM49wRH570Tfr8NiRyNINwXSaBxG09Xoc17khDpISFBxB4M4oXDwDOHEij+stWixNlfQ1lFGQIyc9NY2TuYUcytewelsqGQVKnk7cwtTdzpf+++noHUZou9KGv48vYaFhhEplyGQRiMXbUZ09S1eTFouxlptGLWplEYr8POo0lRw7Z2BnhprpqTG6ruqYejJAUq4ag/kOQnOzAWenjXh7BeDvvRUvTx+cRS6czD1Dad01qvtfUn3rJbXWn9EMvMI0/itFl4ep6bDy2y8P6DW3YP62DU1rN5399xG02kaWLRWx2d2HtSudWOXgyBK7ZRyISyIyLJaUE2pq+yZR9zzkjOk+2ptPaR54Rp2pm+HhHpoNl9hz8AhHVbWMPvkJocVwnRBpChFRafhuCUHsKUb0hRPJ8QcJ9ApCFhqDSaejvuM2Fd0P0VsmuXHvGRm5hcSmZhIen4K3ZBfRaXJGH0wijI+NU6vRk3i0nMyDuZSnpxPqL+ZYVg5xst2E+gVRkpNHRUkVmkYz5jvPOZxTwKx33yc8Pgnp7r0sXiHCySsIs2UQ4ftHk5RUXSK7SMtRxQWyT2tIzC5DXd9G7IFsJGHJbA9LRiJNIHBHPGlZZ2gwmQmUSKnW6lCUniMsKo5g2W4GrLcR+obusSulmEPycqJTi0k6do6sI8fJyi3HboUfwlwRwgJnhA+dEd5Zg8M6Med1JlQ1VyivMVKpbUOtaydfWc/4oykE69gDjuUrMV6sRnm2CnVxIZmRX5ORIcdeJGHOQlc+WOzB3MUeCB+5sf6rCFTnm8g5raaosp7iKh2Kci2Z+eWMjj9GMLW3ESnZhiJxFxn7oji+T8rWDStJSYhjuSiY2Qtdee8TD+bb++AfkISrVywNrV1oL5vQGzrRGztpMHRQpmlgYvI5Qp9lkNSsU6hU5ylUlHGqoIjDaZmUKqtY5ihhziJ3Zi/aiLd3PE5ukax2i6Reb0B9QUt37w2u9pjp6DGjbWjh6dQLhFqDBTuXWOw37cfRNwl79xjsROHsSSnCQSRh1kI3Pl4dgv0aKcI8F5y3xBAREcspRTFWq5Wx76x0m800NrUy9fwFfwLw3W4aNAnhvgAAAABJRU5ErkJggg==) ![](https://getstream.io/blog/static/3eb824b4b8e086e2ae5dd0c512d23e7c/392af/image2-4.png)

 ![](https://getstream.io/blog/static/3eb824b4b8e086e2ae5dd0c512d23e7c/392af/image2-4.png)

#### A. Create a folder called unsplash:

```bash
1mkdir unsplash
```

The full path is **rockets/unsplash**.

#### B. Inside the unsplash directory create a file called unsplash.go with the following content:

```go
1package unsplash
2import (
3    "encoding/json"
4    "fmt"
5    "net/http"
6    "github.com/pkg/errors"
7)
8type APIResponse struct {
9    Total      int
```

#### C. Update rockets/main.go to match this content:

[](https://raw.githubusercontent.com/GetStream/rockets-go-tutorial/master/main.go)[https://raw.githubusercontent.com/GetStream/rockets-go-tutorial/master/main.go](https://raw.githubusercontent.com/GetStream/rockets-go-tutorial/master/main.go)

#### D. Create a file called rockets/spacex.html:

The content should match this:

```jinja
1<html>
2<body>
3<ol>
4    {% for result in response.Results|slice:":8" %}
5    <li>
6        <img src="{{ result.URLs.small }}" />
7        {% if result.Resized %}
8            <img src="data:image/png;base64, {{ result.Resized }} " />
9        {% endif %}
10    </li>
11    {% endfor %}
12</ol>
13</body>
14</html>

```

#### E. Restart your Go server:

```bash
1go run main.go
```

Assuming nothing went wrong you can now open up this URL
[](http://localhost:3000/spacex)[http://localhost:3000/spacex](http://localhost:3000/spacex)
You should see the latest images tagged with SpaceX on Unsplash:

![](data:image/svg+xml;charset=utf-8,%3Csvg height='1999' width='658' xmlns='http://www.w3.org/2000/svg' version='1.1'%3E%3C/svg%3E)

![](data:image/png;base64,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) ![](https://getstream.io/blog/static/8ba6b4c1ad63ff1de259a82f783fd2c4/c7fa2/image3-3.png)

 ![](https://getstream.io/blog/static/8ba6b4c1ad63ff1de259a82f783fd2c4/c7fa2/image3-3.png)

Wicked!

## Step 6 - Understand Those API Calls

Let’s open up **unsplash.go** and have a look at how the code works:

#### A. Making the request:

Go’s builtin HTTP library is pretty functional, here’s the GET request to the Unsplash API

```go
1resp, err := http.Get(url)
```

#### B. Structs and methods:

Go has objects but it’s [different from traditional object oriented languages](https://golang.org/doc/faq#Is_Go_an_object-oriented_language) because it doesn’t support inheritance. Go relies on composition instead of inheritance to help you structure your code. Structs in Go are the closest you’ll come to classes in more object oriented languages. Let’s see how we’re using Structs to talk to the Unsplash API in **unsplash.go**:

```go
1type APIClient struct {
2    // note how the lowercase accessToken is private
3    accessToken string
4}
5func NewAPIClient(token string) APIClient {
6    return APIClient{token}
7}
8func (c *APIClient) Search(query string) (*APIResponse, error) {
9    ...
10    return &response, nil
11}
12

```

The NewAPIClient creates a new APIClient struct. The Search method enables you to write code like:

```go
1client := NewAPIClient("accessToken")
2response, err := client.Search(query)
```

This [blogpost](https://golangbot.com/inheritance/) explains the concept of composition nicely.

#### C. Parsing JSON:

Since Go is a statically typed language you’ll want to parse the JSON into a Struct. If you’re coming from a dynamic language this will be a little bit confusing at first. You’ll get the hang of it quickly. The Unsplash API returns the following JSON:

```json
1{
2  "total": 133,
3  "total_pages": 7,
4  "results": [
5    {
6      "id": "eOLpJytrbsQ",
7      "created_at": "2014-11-18T14:35:36-05:00",
8      "urls": {
9        "raw": "https://images.unsplash.com/photo-1416339306562-f3d12fefd36f",
10        "full": "https://hd.unsplash.com/photo-1416339306562-f3d12fefd36f"
11      },
12    }
13  ]
14}
15

```

There are more fields in the JSON but I simplified it a bit for the example. Next we’ll want to parse the JSON into the following structs:

```go
1type APIResponse struct {
2    Total      int
```

The struct definition specifies the mapping between the field name in the JSON and the struct:

```go
1TotalPages int
```

This means that when decoding JSON it will take the value from the total\_page and set it to the TotalPages property.
To decode the JSON we use the following code:

```go
1response := APIResponse{}
2err = json.NewDecoder(resp.Body).Decode(&response)
3if err != nil {
4    return nil, errors.Wrap(err, "failed to parse JSON")
5}
```

We create a new APIResponse struct and parse a pointer to this object in the Decode function.
The next section will discuss pointers in more detail.

#### D. The defer statement

One of the fairly unique concepts of Go is the defer statement. In unsplash.go line 41 you see the following defer statement:

```go
1defer resp.Body.Close()
```

Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order. So after the Search function finishes the defer statement will trigger and close the request body.

## Step 7 - Pointers

Throughout this tutorial you probably spotted the and **&** symbols. These two are used in Go to work with pointer variables. You can think of pointer variables as references to values. Every type in Go can have its pointer counterpart; you can easily tell if a variable is a pointer or not since its type will start with the symbol.

```go
1var a int // a is an integer variable
2var b *int // b is a pointer to an integer
```

The & and *also work as operators: & returns a pointer to its value and* returns the value a pointer refers to.

#### A. Basic Pointer operations:

If you haven’t used pointers before this concept can be a little bit confusing at first. This little snippet clarifies how pointer operations work:

```go
1a := "Go and Rockets!"
2// the & operator gives you the pointer for a variable, Go infers the type of pointer (*string)
3pointer := &a
4fmt.Println(pointer)
5// 0x1040c128
6// the * operator gives you access to the value
7fmt.Println(*pointer)
8// Go and Rockets!
9

```

You can run the above example in the [Go Playground](https://play.golang.org/p/_bQDQBQLCYO)

#### B. Pointers & Functions:

The example below shows how functions can accept both pointers and regular types:

```go
1type Post struct {
2    Upvotes int
3    Title   string
4}
5func IncrementInPlace(p *Post) {
6    p.Upvotes++
7}
8func Increment(p Post) Post {
9    p.Upvotes++
10    return p
11}
12func main() {
13    a := Post{0, "Go Rocket Tutorial"}
14    IncrementInPlace(&a)
15    fmt.Println(a)
16    b := Increment(a)
17    fmt.Println(b)
18}
19

```

You can run the above example in the [Go Playground](https://play.golang.org/p/BKpEqJVmKWM)
If you’re new to the concept of Pointers you might also enjoy the “[Understand Go pointers in less than 800 words](https://dave.cheney.net/2017/04/26/understand-go-pointers-in-less-than-800-words-or-your-money-back)” blog post.

#### C. Pointer Method Receivers:

One thing that trips up a lot of new Go developers is the way you attach methods to structs:

```go
1func (c *APIClient) Search(query string) (*APIResponse, error) {
2    ...
3}
4

```

In this example we’re attaching the Search function to the \*APIClient type. This allows us to invoke it using this syntax:

```go
1client := NewAPIClient("accessToken")
2response, err := client.Search(query)
3

```

So far so good. However if you use APIClient instead of \*APIClient you’ll run into two unintended consequences:

*   Go will create a copy of your struct while calling the function. Changes you make will only affect your copy
*   Memory usage goes up because you’re copying your struct

Note that this isn’t always a bad thing. In [some cases](https://segment.com/blog/allocation-efficiency-in-high-performance-go-services/) it’s actually faster to not use pointers. Here’s an example of where it goes wrong though:

```go
1package main
2import (
3    "fmt"
4)
5type Post struct {
6    Upvotes      int
7    Title string
8}
9func (p Post) BrokenIncrement() {
10    p.Upvotes++
11    fmt.Printf("p.Upvotes in the function is %d, memory location for p %p\n", p.Upvotes, &p)
12}
13func main() {
14    p := Post{0, "Go Rocket Tuturial"}
15    p.BrokenIncrement()
16    fmt.Printf("p.Upvotes in main is still %d, memory location for p %p", p.Upvotes, &p)
17}
18

```

You can run it on the [Go playground](https://play.golang.org/p/0fSAnB4Kv_R) to see why the above code doesn’t work:
TL/DR use pointer receivers unless you have a good reason not to. The “[Don't Get Bitten by Pointer vs Non-Pointer Method Receivers in Golang](https://nathanleclaire.com/blog/2014/08/09/dont-get-bitten-by-pointer-vs-non-pointer-method-receivers-in-golang/)” covers this issue in more detail.

## Step 8 - Errors

Errors in Go are just variables that you return. Conceptually this approach is very easy to reason about when working with concurrent programming. It can be a bit hard to get the right workflow though. So before you get to read about channels & concurrency (the next topic) we’ll bore you with some mandatory error handling best practices:

#### A. Errors are returned as the last argument of a function:

Don’t deviate from this standard, if you’re writing a function that can fail be sure to allow for an error object as the last return value.

```go
1func Search(query string) (*APIResponse, error) {
2…
3}
4

```

#### B. Stack traces & Errors.Wrap:

The default errors don’t include a stack trace. This is why pretty much every Go app uses the [pkg/errors](https://github.com/pkg/errors) module to wrap errors.

```go
1_, err := ioutil.ReadAll(r)
2if err != nil {
3        return errors.Wrap(err, "read failed")
4}
5

```

#### C. Error causes:

The pkg/errors module also makes it easy to detect the cause of the error.

```go
1switch err := errors.Cause(err).(type) {
2case *MyError:
3        // handle specifically
4default:
5        // unknown error
6}
7

```

#### D. Let errors bubble up:

Most of the time you want to centralize your error handling in one place of [your app](https://getstream.io/activity-feeds/build/social-networks/). You’ll for instance want to return a 500 from your server and log the error in Sentry.
That’s why most functions at the lower levels just return an error in case something breaks. You’ll see this pattern often:

```go
1err := doSomething()
2if err != nil {
3        return errors.Wrap(err, "read failed")
4}
5

```

## Step 9 - Channels & Goroutines

Channels are probably one of the sexiest features of Go. In this part of the tutorial we’ll take a deep dive on how you can leverage Goroutines and Channels for asynchronous programming in Go.
It's good to remember that channels are a feature for power users. By default Go uses a separate Goroutine for every incoming request. So typically you already get the benefit of a highly concurrent server without ever thinking about asynchronous programming. As the creator of Node [said](https://edneypitta.com/on-node-go-concurrency/): “*I think Node is not the best system to build a massive server web. I would use Go for that...*”
Channels and Goroutines give you powerful tools to implement concurrency. For this exercise we’ll start 4 worker goroutines to download and resize 8 pictures from the Unsplash API.
Head over to [](http://localhost:3000/spacex_seams)[http://localhost:3000/spacex\_seams](http://localhost:3000/spacex_seams) and watch the log output. You’ll see that 4 workers are downloading and resizing the images concurrently. Depending on the speed of the machine this will take anywhere from 1-20 seconds.

![](data:image/svg+xml;charset=utf-8,%3Csvg height='1094' width='1386' xmlns='http://www.w3.org/2000/svg' version='1.1'%3E%3C/svg%3E)

![](data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABQAAAAQCAYAAAAWGF8bAAAACXBIWXMAAAsTAAALEwEAmpwYAAAEKElEQVQ4yz3MaVOSCQDA8ecL7LTbdilomZkWeCHHw+ERCCoCCnJpKqgBoimGV5iZpggqHnjkrmhtY7vZ1M7u2L7cY2Zf7H6FPT7Mfyfb6f1vfoKuoQm5yoDCYEJZ34T6jhWNsRWx0Y7W4kRjsqNqaEZhaMRscyFX6lHWW1A3tKAx2dBa2tGa26jQGnF1BREKyyr5LO86F4vkXCmpokCmoVCupUCuRSrXIZGJXC5RcKFIjlxVy+eSG1y5Wf3JSf935wpliMZWhLIaA/k3FRRX6bmtNlGhb0aua6bCYKW8zs4tvZUSlYmiqlpU9S1ck2uQiWYqDS2fjKzWdpYa7X4EuWjkWrkWuWhBUWtD2eBAbXIimt1oLD5UjW4qDTZuKBrQGG3cUpmoqbOjMrajafxo1GYvxQojJkcngvpOK3LRjNrYjqHJS521k3rbXe609VLvCKBr6UZpclOut1HX4kVR10pti5+61i4aHD1nRm/tQaazYfX0I1gcfhqanLR19OLxh/D1DOHvG6VrYAxPcBRfaApPcARzWzfurn6a7T68nWH8fSN09sfwBmN0Rh5i84XoG4wjlJXKKCq4SvVtGQp5BWqFElEjohe1KKqV1Cg1lJXeIu+KBEVlNbKSUpQVVWjVIjpRS41ChUolUiC9Sq2+DqHoahHnz31B3qU88vMKkEiuIZUWIZFeR1pQTL60BIm0mEsX8yiXVXD5wiXy8wuRSD6Yj05SUMqX5/PQiXqE3sksbZFlXNEU7ZEkbeFF7PcWcEeTBMZS9D9I0n1/CUffY3rjGVxD6TPnCC/SFlqkayTFQDyFLzJPdCaL4B/fxTKQxhpZpbE7QePdaVzDazijaTzDKeyhp+h8M9Q4J/GMZGgdXMPUNY25J3HmPMNpXINJVK5JvKMZhP7ZQ3zxLKFHzxiaSjE4kWRyOUc4sY0/tnYWuoZSeIaXGZrbJ5jYZ3Bi6cxOJHP0TW3hiiZxhBaJLx4gZHMnTKdzpHaPmZlf5eHcKrOZF0Qf79I3vUnP/XkGptYJjK+RWDlk4+vXTD9Ok3iSYWb1BeHEFr0jCwTiq0ymcgjHr9+yn3vJ8fErjrKrrCwtk9o6ZDa1x/zKLnvZHWaXsoSn10hmX/DquxNyGylWltMkNw95mt7h2fYuozNrzGWOEN6evOb58284+OqAP96f8NsP33Kws836xg4LyTUWFleYmFkiMjZLen2P57kj/vzpDb98/4rN9S2efjBLK9wfn2M+tY3w89Eyp5kxft2f4p/3m/z74yrvFu7x6J6PHo8Tv9OO227F3mzhSXyY3w9m+Pt0g7/eLTMX9tLZ4cTvcmC3mBgdDCPMDgeZ6Lbz7GGQ071HvMmMsxDx8iDQQSzoJtbbQcRvI+S1Mh0NsB7v4XQ3wculESaCHcQCHmKBDvqcZmZjYf4DeDmR4nJeD7YAAAAASUVORK5CYII=) ![](https://getstream.io/blog/static/544e2b1e314177f141e258305621a04e/0899e/image4-3.png)

 ![](https://getstream.io/blog/static/544e2b1e314177f141e258305621a04e/0899e/image4-3.png)

Let’s have a look at **main.go** to see how this works.

#### A. Channel Creation:

As a first step we create a channel for tasks and a channel for results:

```go
1resultChannel := make(chan TaskResult)
2taskChannel := make(chan Task)
3imagesToResize := 8
4

```

#### B. Starting the workers:

This loop starts the workers

```go
1// start 4 workers
2for w := 1; w <= 4; w++ {
3    go worker(w, taskChannel, resultChannel)
4}
5

```

The “go” statement run the specified function in a separate Goroutine. A Goroutine is a very lightweight alternative to threads. Since a Goroutine only takes 2kb of overhead you can run millions of them even on commodity hardware.
The Goroutines execute the following code:

```go
1func worker(id int, taskChannel <-chan Task, resultChannel chan<- TaskResult) {
2    for j := range taskChannel {
3        fmt.Println("worker", id, "started  job", j.Position)
4        resized, err := seam.ContentAwareResize(j.URL)
5        resultChannel <- TaskResult{j.Position, resized, err}
6    }
7}
8

```

Note how the worker uses the range keyword to iterate over the taskChannel till it’s closed. When the worker receives a new task on the taskChannel it downloads the image, resizes it and writes the result to the resultChannel.

#### C. Write to the task channel:

As a next step we write to the taskChannel and close the channel when we’re done.

```go
1// write to the taskChannel and close it when we're done
2go func() {
3    for i, r := range response.Results[:imagesToResize] {
4        taskChannel <- Task{i, r.URLs["small"]}
5    }
6    close(taskChannel)
7}()
```

Writing or reading from a channel is a blocking operation. To prevent a situation in which we’re waiting forever we run this code in a separate Goroutine. Note the go func() {…}() syntax to execute the code in a Goroutine.

#### D. Retrieving the results:

This loop retrieves the results from the resultChannel and base64 encodes it for easy embedding in the resulting page.

```go
1// start listening for results in a separate goroutine
2for a := 1; a <= imagesToResize; a++ {
3    taskResult := <-resultChannel
4    if taskResult.Err != nil {
5        log.Printf("Image %d failed to resize", taskResult.Position)
6    } else {
7        sEnc := b64.StdEncoding.EncodeToString(taskResult.Resized)
8        response.Results[taskResult.Position].Resized = sEnc
9    }
10}
```

### Channels & Goroutines

And that’s it. Channels are a pretty low level concept. Go also provides Mutexes and Wait groups to help with concurrent programming. When your first learn about Channels it’s easy to become excited and use them to solve every problem. Some issues like preventing multiple writes on a shared object are sometimes easier solved using a simple Mutex. This wiki provides some guidelines on [when to use a sync.Mutex or a Channel](https://github.com/golang/go/wiki/MutexOrChannel).
The cool part about Go’s concurrency model is that most of the time you don’t even need to think about it. You just write your synchronous code and Go handles the switching between Goroutines without you, the programmer, doing any work.
Channels, GoRoutines, Wait Groups and Mutexes give you full control over how you’re writing asynchronous code when you need it. If you want to learn more about channels check out this [Golang channels tutorial](http://guzalexander.com/2013/12/06/golang-channels-tutorial.html).

## Step 10 - Libraries and Tools

Go is an easy language to learn. What tends to take more time is figuring out the tooling and the ecosystem. Here are a few library and tool recommendations to help you get started.
**Tool: Go doc**
Go doc is an awesome little tool for generating reference docs for your Go APIs.
[](https://godoc.org/golang.org/x/tools/cmd/godoc)[https://godoc.org/golang.org/x/tools/cmd/godoc](https://godoc.org/golang.org/x/tools/cmd/godoc)
**Tool: go fmt & go imports**
[Go fmt](https://golang.org/cmd/gofmt/) and [go imports](https://godoc.org/golang.org/x/tools/cmd/goimports) two tools help you automatically clean up your Go files.
**Tool: PPROF**
Go ships with awesome profiling capabilities and beautiful flamegraphs about your code’s performance. This tutorial is a solid starting point: [](https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/)[https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/](https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/)

![](data:image/svg+xml;charset=utf-8,%3Csvg height='706' width='1392' xmlns='http://www.w3.org/2000/svg' version='1.1'%3E%3C/svg%3E)

![](data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABQAAAAKCAYAAAC0VX7mAAAACXBIWXMAAAsTAAALEwEAmpwYAAACLUlEQVQoz42PS09TURSF+4OcWxTaYkwcMNM4IzrRgQPiA2xEKRHRREOMPCYEdWCAgUknatBoqkIIlUofUApt7bv0cW/v2eciQyefub0m4szBl7X2Xmev5Hi0VjiI/FHVQawOx/eNRr2L47u5Mv+9cfZiddUj0kLrFq62EauOdCqulyYiDWzbwD40ENVAzCrKLB/Lm917LW6HR+sqIi7arqPaWaxKHG0fIKqEmDk3d95ZzpxHjDxa1xApoxppxCp2c6fLI/IDLQW6apdRrSSd4mf0YQWx9hFjx810ETG2UY04ysgguoBq7dApfUW1EqjCOmLlnMIMDspMoe19OqV1Dt7Po3/mUUYC1dxEH+bQdg6rsYG0Y9h6F/uogFX4RDv6gk7uHdXIY6zaN+fLKbQksHWSo19FWnsr5BaGsM1VtBlDzE3M7BvkIILqxCiuzpOcv0oj+pxW9iX1jUfk18aJv7qIWYzgsfbCmKWPNFOvyb4dI/rkArHRHior49RSS1S+TBGfHGT92gCpictEB0+SuOQlMxRg6/oZEsGzJG70s3XFSz08gSe+MMiH0DmiYz62Ql7ioR52HvhITgaIPwyQuH+axGgf32962R72krkbYPdegPQdH+lgH+mRXnaDPtLDp8g+PY9nbbKfyK0T5Kd6qc/6qU77qUz7qD7rc3H8rI/anJ+Kw4zPZfYYc37KM71UFwbwtMIhjMXbmMtBjKXbf1l2CLq69B8sjmCGQ/wGN1WP+D9XAiQAAAAASUVORK5CYII=) ![](https://getstream.io/blog/static/9b7f275eceb7662da7fc1127148d89d5/87afb/image6-3.png)

 ![](https://getstream.io/blog/static/9b7f275eceb7662da7fc1127148d89d5/87afb/image6-3.png)

**Mux**
[Mux](https://github.com/gorilla/mux) is an awesome little library for URL routing in Go:
**CLI**
The aptly named [CLI](https://github.com/urfave/cli) is a great library for building command line interfaces in Go.
**Zap**
[Zap](https://github.com/uber-go/zap), Uber’s logging library, is a better alternative to Logrus and the default Go logging system.
**Go.Rice**
[Go Rice](https://github.com/GeertJohan/go.rice) enables you to easily embed static files such as templates into your Go binary.
**Squirrel**
[Squirrel](https://github.com/Masterminds/squirrel) makes it easy to work with SQL in Go.
One of the reasons why we picked Go is the [mature ecosystem of libraries](https://github.com/avelino/awesome-go). Most of the time you’ll find high quality open source libraries to help you with your projects.

## Concluding the Go Tutorial

At [Stream](https://getstream.io/) we use Go & RocksDB to power the news feeds for over 300 million end users. What’s unique about Go is the balance it strikes. It’s ridiculously fast and still easy and productive to work with.

It’s been 18 months since Stream switched from [Python to Go](https://getstream.io/blog/switched-python-go/). There is always some uncertainty when making such a large change. Fortunately moving to Go worked out very well. We’re able to power a complex and large infrastructure with a relatively small [development team](https://getstream.io/team/). Go improves with every release. Every time a new version is released we rerun our benchmarks shave a few milliseconds of our response times. The tooling also gets better with every release.

Go is a fun language and I hope you enjoyed this tutorial. We covered the basics and even some more advanced topics such as Goroutines and Channels. Let me know if anything in this tutorial isn’t clear ([@tschellenbach](https://twitter.com/tschellenbach)).

If you’ve never used Stream before, try out this [Tour of the API](https://getstream.io/get_started/). You can try it out with Javascript in the browser, or of course with the official [Stream Go client](https://github.com/GetStream/stream-go2).
