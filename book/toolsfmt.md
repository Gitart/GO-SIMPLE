## Инструменты форматирования кода и соглашения об именах в Голанге
Форматирование кода показывает способ форматирования кода в файле. Он указывает на то, как разработан код и как возврат каретки используется в письменной форме. Go не требует специальных правил для форматирования кода, но имеет стандарт, который обычно используется и принят в сообществе.
Форматирование кода с использованием gofmt
Golang предоставляет инструмент gofmt для поощрения и гарантии того, что код Golang отформатирован в соответствии с ожидаемыми соглашениями. Важность этого инструмента в том, что нам даже не нужно определять конвенции.

Программа ниже будет скомпилирована и выполнена без ошибок. Считаете ли вы это легко читать?

```golang
package main
import ("fmt")
var(
    a \= 654
    b \= false
    c   \=2.651
    d  \=4+ 1i
)
func main(){ fmt.Printf("d for Integer: %d\\n", a)
    fmt.Printf( "t for Boolean: %t\\n", b )
    fmt.Printf("g for Float: %g\\n",c)
    fmt.Printf("e for Scientific Notation: %e\\n",d)}
```

From command line, run the program with **gofmt \-w test1.go**.
Now you should see that it has been reformatted by the **gofmt** tool and is now become more readable, with consistent spacing.

```golang
package main

import (
    "fmt"
)

var (
    a \= 654
    b \= false
    c \= 2.651
    d \= 4 + 1i
)

func main() {
    fmt.Printf("d for Integer: %d\\n", a)
    fmt.Printf("t for Boolean: %t\\n", b)
    fmt.Printf("g for Float: %g\\n", c)
    fmt.Printf("e for Scientific Notation: %e\\n", d)
}
```

\-d option used to show the difference.

\-w option used to overwrited the current file with formatting applied.

---

## Check naming conventions using golint

Naming will ever be a significantly subjective exercise, but it is worth spending some moment considering how you will name things. Some things need to be consider when naming variables, functions, structs, constant and interfaces like, Who will be using this code? Is it just me or a wider team? Could someone who is new to the code read it and understand practically what it is doing?

It is essential to sustain some conventions around naming, but being assertive about naming can also be a hindrance. For the majority of instances, it should be feasible to satisfy with your own conventions and the situation in which the code is being used.

The **golint** tool gives helpful hints on style and can also serve with reviewing the recognized conventions of the Golang.The **golint** command searches for style mistakes in terms of the conventions of the Golang project itself.

The golint executable is not installed by default but can be installed as follows:

```
go get \-u github.com/golang/lint/golint
```


To verify that the tool was installed properly, type **golint \-\-help** at the terminal. You should see some help text written to the monitor.
Let's take an example code that could be improved.

```golang
package main

import (
    "io/ioutil"
    "log"
    "fmt"
    "os"
)

func ReadFile() {
    File\_Data, err :\= ioutil.ReadFile("test.txt")
    if err !\= nil {
        log.Panicf("failed reading data from file: %s", err)
    }
    fmt.Printf("\\nFile Content: %s", File\_Data)
}

func main() {
    fmt.Printf("\\n######## Read file #########\\n")
    ReadFile()
}
```

Now run the golint for above program.

![](https://www.golangprograms.com/media/wysiwyg/name.JPG)

These suggestions are not compulsory, as the code will compile, but it is a good practice to fix the warnings and pick up the conventions because Go itself adopts this tool. Using the golint tool can also be a great idea to learn how to write Idiomatic Go. The tool covers many conventions, including naming, styles, and general conventions. Many of the text editor plugins provides a way to run golint in a project during development or on save.
