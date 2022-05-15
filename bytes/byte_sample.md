## Samples work with Bytes

[Sample links](https://zetcode.com/golang/byte/)   

```go
package main

import (
     "bytes"
     "fmt"
)

func main() {

     data1 := []byte{102, 97, 108, 99, 111, 110} // falcon
     data2 := []byte{111, 110}                   // on

     if bytes.Contains(data1, data2) {
          fmt.Println("contains")
     } else {
          fmt.Println("does not contain")
     }

     if bytes.Equal([]byte("falcon"), []byte("owl")) {
          fmt.Println("equal")
     } else {
          fmt.Println("not equal")
     }

     data3 := []byte{111, 119, 108, 9, 99, 97, 116, 32, 32, 32, 32, 100, 111,
          103, 32, 112, 105, 103, 32, 32, 32, 32, 98, 101, 97, 114}

     fields := bytes.Fields(data3)
     fmt.Println(fields)

     for _, e := range fields {
          fmt.Printf("%s ", string(e))
     }

     fmt.Println()
}
```


```go
package main

import (
     "bytes"
     "fmt"
)

func main() {

     data := [][]byte{[]byte("an"), []byte("old"), []byte("wolf")}
     joined := bytes.Join(data, []byte(" "))

     fmt.Println(data)
     fmt.Println(joined)
     fmt.Println(string(joined))

     fmt.Println("--------------------------")

     data2 := []byte{102, 97, 108, 99, 111, 110, 32}
     data3 := bytes.Repeat(data2, 3)

     fmt.Println(data3)
     fmt.Println(string(data3))

     fmt.Println("--------------------------")

     data4 := []byte{32, 32, 102, 97, 108, 99, 111, 110, 32, 32, 32}
     data5 := bytes.Trim(data4, " ")

     fmt.Println(data5)
     fmt.Println(string(data5))
}
```


The example joins byte slices with Join, repeats a byte slice with Repeat, and trims byte slices of the specified byte with Trim.

```
$ go run byte_funs2.go 
[[97 110] [111 108 100] [119 111 108 102]]
[97 110 32 111 108 100 32 119 111 108 102]
an old wolf
--------------------------
[102 97 108 99 111 110 32 102 97 108 99 111 110 32 102 97 108 99 111 110 32]
falcon falcon falcon 
--------------------------
[102 97 108 99 111 110]
falcon
```


## Buffers

```go
package main

import (
     "bytes"
     "fmt"
)

func main() {

     var buf bytes.Buffer

     buf.Write([]byte("a old"))
     buf.WriteByte(32)
     buf.WriteString("cactus")
     buf.WriteByte(32)
     buf.WriteByte(32)
     buf.WriteRune('🌵')

     fmt.Println(buf)
     fmt.Println(buf.String())
}
```

We build a bytes.Buffer with Write, WriteByte, WriteString, WriteByte, and WriteRune methods.

```
$ go run buffer.go 
{[97 32 111 108 100 32 99 97 99 116 117 115 32 32 240 159 140 181] 0 0}
a old cactus  🌵
```


### read_binary.go

```go
package main

import (
     "bufio"
     "encoding/hex"
     "fmt"
     "io"
     "log"
     "os"
)

func main() {

     f, err := os.Open("favicon.ico")

     if err != nil {
          log.Fatal(err)
     }

     defer f.Close()

     reader := bufio.NewReader(f)
     buf := make([]byte, 256)

     for {
          _, err := reader.Read(buf)

          if err != nil {
               if err != io.EOF {
                    fmt.Println(err)
               }
               break
          }

          fmt.Printf("%s", hex.Dump(buf))
     }
}
```

