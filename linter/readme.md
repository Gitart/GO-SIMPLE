# go vet vs go fmt vs go lint

11-25-20 [Bryan Braun](https://sparkbox.com/foundry/author/bryan_braun)

Know the differences and learn how to use go vet, gofmt, and golint to check your code in the Go programming language.

![](https://sparkbox.com/uploads/featured_images/b-braun_20-11.png)

The [Go ecosystem](https://packetpushers.net/podcast/full-stack-journey-045-learning-to-program-in-go/) provides a lot of tools to help you write cleaner, more predictable code. As a result, it can be a bit tricky to figure out what you should use, especially if you’re new to the language. Some of the more popular tools are `go vet`, `gofmt`, and `golint`.

## We’re Hiring Frontend Developers!

Remote applicants welcome

Do you have a solid knowledge of HTML, CSS, and JavaScript while being mindful of the diverse ecosystem of devices and connections? We’re looking for experienced Frontend Developers who love to learn and collaborate.

[Apply Today](https://sparkbox.com/careers)

To be clear, these three tools aren’t competing options. Each one checks for different things, allowing them to be used in conjunction. You could think of them as separate steps in a code-checking ladder.

![Golang code-checking options, arranged by how opinionated they are. It includes, in order of least to most opinionated, Go compiler, go vet, gofmt, and golint. Detailed description below image.](https://sparkbox.com/uploads/article_uploads/code-checking-options.png)

Detailed description of the image showing popular options for code checking in Go

Popular options for code checking in Go can be thought of using the metaphor of a ladder to describe how opinionated they are. The ladder has four rungs, with the least opinionated at the bottom, and the most opinionated at the top.

The bottom rung is GoCompiler, which finds serious errors that prevent your code from running.

The second rung is go vet, which finds subtle issues where your code may not work as intended.

The third rung is gofmt, which applies standard formatting (whitespace, indentation, etc)

The top rung is golint(and other linters), which make code style recommendations (naming, code conventions, etc).

---

At the bottom of the ladder, the Go compiler does a baseline check to ensure that your code is valid, catching things like syntax errors and missing imports. Any code checking after that is optional. The higher rungs on the ladder are more focused on code aesthetics.

With that, let’s look a little more closely at each of these tools to see what they do.

## go vet

`go vet` starts where the compiler ends by identifying subtle issues in your code. It’s good at catching things where your code is technically valid but probably not working as intended. One example of this is when you have [unreachable code](https://en.wikipedia.org/wiki/Unreachable_code). `go vet` is part of a standard Go installation, making it straightforward to run from the command line.

**Running it:** `go vet main.go`

**Example:**

*Before:*

```
package main
import "fmt"
// Prints out "Super Mario 3"
func main() {
    game_version :=3
    fmt.Printf("Super Mario %s\n",game_version)
}

```

*After:*

```
./main.go:6:2: Printf format %s has arg 3 of wrong type int

```

*This is warning us that we’re trying to interpolate an int into a string using the %s “format verb,” which isn’t usually intended for integers. While it is valid Go, it produces unexpected results, ultimately printing: Super Mario %!s(int=3)*

**Resources:**

*   [Official Documentation: Go Vet](https://golang.org/cmd/vet/)

## gofmt

`gofmt` defines code formatting standards for Go and applies them to your code automatically. These formatting changes don’t affect the execution of the code—rather, they improve codebase readability by ensuring that the code is visually consistent. `gofmt` focuses on things like indentation, whitespace, comments, and [general code succinctness](https://golang.org/cmd/gofmt/#hdr-The_simplify_command).

`gofmt` is included in a standard Go installation and can be run from Go’s command line tools.

**Running it:** `go fmt main.go` or `gofmt -w main.go`

**Example:**

*Before:*

```
package main
import "fmt"
// Prints out "Super Mario 3"
func main() {
    game_version :=3
    fmt.Printf("Super Mario %s\n",game_version)
}

```

After:

```
  package main
+
  import "fmt"
+
  // Prints out "Super Mario 3"
  func main() {
-   game_version :=3
-   fmt.Printf("Super Mario %s\n",game_version)
+   game_version := 3
+   fmt.Printf("Super Mario %s\n", game_version)
  }

```

*This example adds blank lines to separate the package and import declarations. It also adds some spacing into the variable assignment and the function parameters.*

**Resources:**

*   [Official documentation: gofmt](https://golang.org/cmd/gofmt/)
*   [Introductory blog post: “go fmt your code”](https://blog.golang.org/gofmt)
*   [Effective Go: Formatting](https://golang.org/doc/effective_go.html#formatting)

**Alternatives:**

`gofmt` has a few alternatives. The following tools apply the same formatting as `gofmt` along, with some additional formatting rules (which is useful if you want to enforce an even stricter formatting standard):

*   [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports)
*   [gofumpt](https://github.com/mvdan/gofumpt)

## golint

`golint` is a linter maintained by the Go developers. It is intended to enforce the coding conventions described in [Effective Go](https://golang.org/doc/effective_go.html) and [CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments). These same conventions are used in the open-source Go project and at Google. `golint` is only concerned with stylistic matters, and its rules are more like opinions than a hard standard. As the project README states:

> The suggestions made by golint are exactly that: suggestions. Golint is not perfect… do not treat its output as a gold standard.

`golint` is just one of [many possible Go linters](https://github.com/golangci/awesome-go-linters) you can use. These linters can be run together or in parallel to check various aspects of your code, including coding conventions, performance, complexity, and more. For example, [golangci-lint](https://github.com/golangci/golangci-lint) is a popular linter runner that comes prepackaged with dozens of linters, including `golint`. To run `golint` by itself, you can install it with `go get -u golang.org/x/lint/golint` and run it in the terminal.

**Running it:** `golint main.go`

**Example:**

*Before:*

```
package main
import "fmt"
// Prints out "Super Mario 3"
func main() {
    game_version :=3
    fmt.Printf("Super Mario %s\n",game_version)
}

```

*After:*

```
main.go:5:2: don't use underscores in Go names; var game_version should be gameVersion

```

**Resources:**

*   [The Golint repo](https://github.com/golang/lint)
*   [Awesome Go Linters](https://github.com/golangci/awesome-go-linters)
*   [Go Linters: Myths and best practices](https://about.sourcegraph.com/go/gophercon-2019-go-linters-myths-and-best-practices/)
*   [Effective Go](https://golang.org/doc/effective_go.html)
*   [CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments)

**Alternatives:**

The following alternative linters can be used in place of, or in addition to, `golint`:

*   `golangci-lint` ([repo](https://github.com/golangci/golangci-lint) and [docs](https://golangci-lint.run/))
*   `revive` ([repo](https://github.com/mgechev/revive) and [docs](https://revive.run/))
*   `staticcheck` ([repo](https://github.com/dominikh/go-tools) and [docs](https://staticcheck.io/))
*   …and [many others](https://github.com/golangci/awesome-go-linters#linters)

## Integrations and more

The code checking tools we’ve discussed above are mostly just the “official” ones (the ones maintained by Golang developers). One nice benefit of having official tooling is IDE integrations. For example, Goland [has built-in support for gofmt](https://stackoverflow.com/a/47737130/1154642), and VSCode has [an official Go extension that can check your code whenever you save a file](https://code.visualstudio.com/docs/languages/go).

Whether you use these official code checkers or other ones provided by the community, there are plenty of great options for keeping your code clean and consistent. Hopefully this article gives you a clearer idea of how you can use `go vet`, `gofmt`, and `golint` to benefit your code.
