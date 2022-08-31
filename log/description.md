# How to change the format of log output in logrus?

Category [Golang](https://newdevzone.com/posts/categories/golang)

Go is a statically typed, compiled programming language designed at Google. It is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style but what makes it special in every regard is its native support for concurrency and parallelism.

There are 3 suggested solutions in this post and each one is listed below with a detailed description on the basis of most helpful answers as shared by the users. I have tried to cover all the aspects as briefly as possible covering topics such as Go, Logging and a few others. I have categorized the possible solutions in sections for a clear and precise explanation. Please consider going through all the sections to better understand the solutions.

Contents

1.  01 [How to change the format of log output in logrus?](https://newdevzone.com/posts/how-to-change-the-format-of-log-output-in-logrus#)
2.  02 [Solution 1](https://newdevzone.com/posts/how-to-change-the-format-of-log-output-in-logrus#solution_1)
3.  03 [Solution 2](https://newdevzone.com/posts/how-to-change-the-format-of-log-output-in-logrus#solution_2)
4.  04 [Solution 3](https://newdevzone.com/posts/how-to-change-the-format-of-log-output-in-logrus#solution_3)
5.  05 [Final Words](https://newdevzone.com/posts/how-to-change-the-format-of-log-output-in-logrus#article_footer)

#### Solution 1

### Standard logrus-prefixed-formater usage

To achieve this you need to make your own `TextFormater` which will satisfy logrus `Formatter` interface. Then when you create your own formater you pass it on logrus struct initialization. Other way arround and close to what you wanna achieve is this formater [https://github.com/x-cray/logrus-prefixed-formatter](https://github.com/x-cray/logrus-prefixed-formatter) . Based on this formater you can create your own.

In your case you need to use like that

```yaml
logger := &logrus.Logger{
        Out:   os.Stderr,
        Level: logrus.DebugLevel,
        Formatter: &prefixed.TextFormatter{
            DisableColors: true,
            TimestampFormat : "2006-01-02 15:04:05",
            FullTimestamp:true,
            ForceFormatting: true,
        },
    }

```

---

### Customized output of logrus-prefixed-formater

Link to gist to use copy of `logrus-prefixed-formatter` with changed format [https://gist.github.com/t-tomalak/146e4269460fc63d6938264bb5aaa1db](https://gist.github.com/t-tomalak/146e4269460fc63d6938264bb5aaa1db)

I leave this option if u in the end wanna use it, as in this version you have exact format you want, coloring, and other features available in standard formatter

---

### Custom formatter

Third option is to use package create by me [https://github.com/t-tomalak/logrus-easy-formatter](https://github.com/t-tomalak/logrus-easy-formatter). It provide simple option to format output as you want and is it only purpose. I removed not necessary options which probably you wouldn't use.

```go
package main

import (
    "os"

    "github.com/sirupsen/logrus"
    "github.com/t-tomalak/logrus-easy-formatter"
)

func main() {
    logger := &logrus.Logger{
        Out:   os.Stderr,
        Level: logrus.DebugLevel,
        Formatter: &easy.Formatter{
            TimestampFormat: "2006-01-02 15:04:05",
            LogFormat:       "[%lvl%]: %time% - %msg%",
        },
    }

    logger.Printf("Log message")
}

```

This sample code will produce:

```markdown
[INFO]: 2018-02-27 19:16:55 - Log message

```

Also I wanna point out that if in the future wanna change formatter there shouldn't be any problems to use i.e. default Logrus `TextFormatter/JSONFormatter`.

---

### Customized output of logrus-prefixed-formater

If you really don't wanna copy this formatter to your project you can use my fork logrus-prefixed-formater with copy/pasted this code [https://github.com/t-tomalak/logrus-prefixed-formatter](https://github.com/t-tomalak/logrus-prefixed-formatter)

You can use it like standard option but remember to change import to my repository in you go file

```yaml
logger := &logrus.Logger{
        Out:   os.Stderr,
        Level: logrus.DebugLevel,
        Formatter: &prefixed.TextFormatter{
            DisableColors: true,
            TimestampFormat : "2006-01-02 15:04:05",
            FullTimestamp:true,
            ForceFormatting: true,
        },
    }

```

#### Solution 2

I guess I am quite late to this, but recently I was also struggling to get this logging message format customized and I was hoping to get this done preferably in a simpler way. Coming from Python, this wasn't really as straight forward as I though it to be and logrus documentation also isn't very clear on this.

So I had to go through their source code to actually get this done. Here is my code for the same.

```go
type myFormatter struct {
    log.TextFormatter
}

func (f *myFormatter) Format(entry *log.Entry) ([]byte, error) {
// this whole mess of dealing with ansi color codes is required if you want the colored output otherwise you will lose colors in the log levels
    var levelColor int
    switch entry.Level {
    case log.DebugLevel, log.TraceLevel:
        levelColor = 31 // gray
    case log.WarnLevel:
        levelColor = 33 // yellow
    case log.ErrorLevel, log.FatalLevel, log.PanicLevel:
        levelColor = 31 // red
    default:
        levelColor = 36 // blue
    }
    return []byte(fmt.Sprintf("[%s] - \x1b[%dm%s\x1b[0m - %s\n", entry.Time.Format(f.TimestampFormat), levelColor, strings.ToUpper(entry.Level.String()), entry.Message)), nil
}

func main() {
    f, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY, 0777)
    logger := &log.Logger{
        Out:   io.MultiWriter(os.Stderr, f),
        Level: log.InfoLevel,
        Formatter: &myFormatter,
        },
    }
    logger.Info("Info message")
    logger.Warning("Warning message")

```

Here is the output

```css
± go run main.go                                                                                                                                                                                                                    <<<
[2019-05-13 18:10:34] - INFO - Info message
[2019-05-13 18:10:34] - WARNING - Warning message

```

PS: I am very new to this, so if you guys there can be a better neat way of doing this, please do share.

#### Solution 3

Here's the final solution you can try out in case no other solution was helpful to you. This one's applicable and useful in some cases and could possiblty be of some help. No worries if you're unsure about it but I'd recommend going through it.

I adopted this code from [https://github.com/x-cray/logrus-prefixed-formatter/blob/master/formatter.go](https://github.com/x-cray/logrus-prefixed-formatter/blob/master/formatter.go). I created an own formatter struct and implementend an own Format function of the logrus formatter interface. I you only need a text output without color, this might be a simple solution.

```go
// adopted from https://github.com/x-cray/logrus-prefixed-formatter/blob/master/formatter.go
package main

import (
    "bytes"
    "fmt"
    "log"
    "os"
    "strings"

    "github.com/sirupsen/logrus"
)

type LogFormat struct {
    TimestampFormat string
}

func (f *LogFormat) Format(entry *logrus.Entry) ([]byte, error) {
    var b *bytes.Buffer

    if entry.Buffer != nil {
        b = entry.Buffer
    } else {
        b = &bytes.Buffer{}
    }

    b.WriteByte('[')
    b.WriteString(strings.ToUpper(entry.Level.String()))
    b.WriteString("]:")
    b.WriteString(entry.Time.Format(f.TimestampFormat))

    if entry.Message != "" {
        b.WriteString(" - ")
        b.WriteString(entry.Message)
    }

    if len(entry.Data) > 0 {
        b.WriteString(" || ")
    }
    for key, value := range entry.Data {
        b.WriteString(key)
        b.WriteByte('=')
        b.WriteByte('{')
        fmt.Fprint(b, value)
        b.WriteString("}, ")
    }

    b.WriteByte('\n')
    return b.Bytes(), nil
}

func main() {
    formatter := LogFormat{}
    formatter.TimestampFormat = "2006-01-02 15:04:05"

    logrus.SetFormatter(&formatter)
    log.SetOutput(os.Stderr)

    logrus.WithFields(logrus.Fields{
        "animal": "walrus",
        "size":   10,
    }).Info("A group of walrus emerges from the ocean")
    logrus.Info("ugh ugh ugh ugh")
}

```

#### Final Words

Go really fits well here with regard to the performance-oriented cloud software. The popular DevOps tools have been written in Go, such as Docker, and also the open-source container orchestration system Kubernetes.. These were a few of many solutions that were found helpful for your issue. Hope it turns out helpful for you. Please upvote the solutions if it worked for you.
