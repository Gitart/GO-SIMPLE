## Introduction to bufio package in Golang

Package bufio helps with buffered I/O. Through a bunch of examples we’ll get familiar with goodies it provides: Reader, Writer and Scanner…
bufio.Writer

Doing many small writes can hurt performance. Each write is ultimately a syscall and if doing frequently can put burden on the CPU. Devices like disks work better dealing with block-aligned data. To avoid the overhead of many small write operations Golang is shipped with bufio.Writer. Data, instead of going straight to destination (implementing io.Writer interface) are first accumulated inside the buffer and send out when buffer is full:

producer --> buffer --> io.Writer

Let’s visualise how buffering works with nine writes (one character each) when buffer has space for 4 characters:
```go
producer         buffer           destination (io.Writer)
 
   a    ----->   a
   b    ----->   ab
   c    ----->   abc
   d    ----->   abcd
   e    ----->   e      ------>   abcd
   f    ----->   ef               abcd
   g    ----->   efg              abcd
   h    ----->   efgh             abcd
   i    ----->   i      ------>   abcdefgh
```
(arrows are write operations)

bufio.Writer uses[]byte buffer under the hood (source code):

type Writer intfunc (*Writer) Write(p []byte) (n int, err error) {
    fmt.Println(len(p))
    return len(p), nil
}func main() {
    fmt.Println("Unbuffered I/O")
    w := new(Writer)
    w.Write([]byte{'a'})
    w.Write([]byte{'b'})
    w.Write([]byte{'c'})
    w.Write([]byte{'d'})
    fmt.Println("Buffered I/O")
    bw := bufio.NewWriterSize(w, 3)
    bw.Write([]byte{'a'})
    bw.Write([]byte{'b'})
    bw.Write([]byte{'c'})
    bw.Write([]byte{'d'})
    err := bw.Flush()
    if err != nil {
        panic(err)
    }
}Unbuffered I/O
1
1
1
1
Buffered I/O
3
1

Unbuffered I/O simply means that each write operation goes straight to destination. We’ve 4 write operations and each one maps to Write call where passed slice of bytes has length 1.

With buffered I/O we’ve internal buffer (3 bytes long) which collects data and flushes buffer when full. First three writes end up inside the buffer. 4th write detects buffer with no free space so it sends accumulate data out. It gives space to hold d character. But there’s something more —Flush call. It’s needed at the very end to flush any outstanding data — bufio.Writer sends data only when buffer is either full or when explicitly requested with Flush method.

    By default bufio.Writer uses 4096 bytes long buffer. It can be set with NewWriterSize function.

## Implementation

It’s rather straightforward (source code):
```go
type Writer struct {
    err error
    buf []byte
    n   int
    wr  io.Writer
}
```

Field buf accumulates data. Consumer (wr) gets data when buffer is full or Flush is called. First encountered I/O error is held by err — after encountering an error, writer is no-op (source code):
```go
type Writer intfunc (*Writer) Write(p []byte) (n int, err error) {
    fmt.Printf("Write: %q\n", p)
    return 0, errors.New("boom!")
}func main() {
    w := new(Writer)
    bw := bufio.NewWriterSize(w, 3)
    bw.Write([]byte{'a'})
    bw.Write([]byte{'b'})
    bw.Write([]byte{'c'})
    bw.Write([]byte{'d'})
    err := bw.Flush()
    fmt.Println(err)
}Write: "abc"
boom!
```
Here we see that Flush didn’t call 2nd write on our consumer. Buffered writer simply doesn’t try to do more writes after first error.

Field n is the current writing position inside the buffer. Buffered method returns n’s value (source code):
```go
type Writer intfunc (*Writer) Write(p []byte) (n int, err error) {
    return len(p), nil
}func main() {
    w := new(Writer)
    bw := bufio.NewWriterSize(w, 3)
    fmt.Println(bw.Buffered())
    bw.Write([]byte{'a'})
    fmt.Println(bw.Buffered())
    bw.Write([]byte{'b'})
    fmt.Println(bw.Buffered())
    bw.Write([]byte{'c'})
    fmt.Println(bw.Buffered())
    bw.Write([]byte{'d'})
    fmt.Println(bw.Buffered())
}0
1
2
3
1
```
It starts with 0 and is incremented by the number of bytes added to buffer. It’s also reset after flush to underlying writer while calling bw.Write([]byte{'d'}).
## Large writes
```go
type Writer intfunc (*Writer) Write(p []byte) (n int, err error) {
    fmt.Printf("%q\n", p)
    return len(p), nil
}func main() {
    w := new(Writer)
    bw := bufio.NewWriterSize(w, 3)
    bw.Write([]byte("abcd"))
}
```

Program (source code) prints "abcd" because bufio.Writer detects if Write is called with amount of data too much for internal buffer (3 bytes in this case) . It then calls Write method directly on writer (destination object). It’s completely fine since amount of data is large enough to skip proxying through temporary buffer.
## Reset

Buffer which is the core part of bufio.Writer can be re-used for different destination writer with Reset method. It saves memory allocation and extra work for garbage collector (source code):
```go
type Writer1 intfunc (*Writer1) Write(p []byte) (n int, err error) {
    fmt.Printf("writer#1: %q\n", p)
    return len(p), nil
}type Writer2 intfunc (*Writer2) Write(p []byte) (n int, err error) {
    fmt.Printf("writer#2: %q\n", p)
    return len(p), nil
}func main() {
    w1 := new(Writer1)
    bw := bufio.NewWriterSize(w1, 2)
    bw.Write([]byte("ab"))
    bw.Write([]byte("cd"))
    w2 := new(Writer2)
    bw.Reset(w2)
    bw.Write([]byte("ef"))
    bw.Flush()
}writer#1: "ab"
writer#2: "ef"
```

There is one bug in this program. Before calling Reset we should flush the buffer with Flush. Currently, written data cd is lost since Reset simply discards any outstanding information (source code):
```go
func (b *Writer) Reset(w io.Writer) {
    b.err = nil
    b.n = 0
    b.wr = w
}
```

## Buffer free space

To check how much space left inside the buffer we can use Available method (source code):
```go
w := new(Writer)
bw := bufio.NewWriterSize(w, 2)
fmt.Println(bw.Available())
bw.Write([]byte{'a'})
fmt.Println(bw.Available())
bw.Write([]byte{'b'})
fmt.Println(bw.Available())
bw.Write([]byte{'c'})
fmt.Println(bw.Available())2
1
0
1
```
## Write{Byte,Rune,String} Methods

At our disposal we’ve 3 utility functions to write data of common types (source code):
```go
w := new(Writer)
bw := bufio.NewWriterSize(w, 10)
fmt.Println(bw.Buffered())
bw.WriteByte('a')
fmt.Println(bw.Buffered())
bw.WriteRune('ł') // 'ł' occupies 2 bytes
fmt.Println(bw.Buffered())
bw.WriteString("aa")
fmt.Println(bw.Buffered())0
1
3
5
```
## ReadFrom

Package io defines io.ReaderFrom interface. It’s usually implemented by writer to do the dirty work of reading all the data from specified reader (until EOF):
```
type ReaderFrom interface {
        ReadFrom(r Reader) (n int64, err error)
}

    io.ReaderFrom interface is used f.ex. by io.Copy.

bufio.Writer implements this interface, allowing to call ReadFrom method which digests all data from io.Reader (source code):

type Writer intfunc (*Writer) Write(p []byte) (n int, err error) {
    fmt.Printf("%q\n", p)
    return len(p), nil
}func main() {
    s := strings.NewReader("onetwothree")
    w := new(Writer)
    bw := bufio.NewWriterSize(w, 3)
    bw.ReadFrom(s)
    err := bw.Flush()
    if err != nil {
        panic(err)
    }
}"one"
"two"
"thr"
"ee"
```
    It’s important to call Flush even while using ReadFrom.

## bufio.Reader

It allows to read in bigger batches from the underlying io.Reader. This leads to less read operations which can improve performance if e.f. underlying media works better when data is read in blocks of certain size:

io.Reader --> buffer --> consumer

Suppose that consumer wants to read 10 characters one by one from the disk. In naive implementation this will trigger 10 read calls. If disk reads data in blocks of size 4 bytes then bufio.Reader can help out. Under the hood it will buffer whole blocks giving the consumer an API (io.Reader) to read it by one byte:

```go
abcd -----> abcd -----> a
            abcd -----> b
            abcd -----> c
            abcd -----> d
efgh -----> efgh -----> e
            efgh -----> f
            efgh -----> g
            efgh -----> h
ijkl -----> ijkl -----> i
            ijkl -----> j
```

(arrows are read operations)

This method will require only 3 reads from disk (instead of 10).
Peek

Method Peek allows to see first n bytes of buffered data without actually “eating” it:

    If buffer isn’t full and holds less than n bytes then it’ll try to read more from underlying io.Reader
    If requested amount is bigger than the size of the buffer then bufio.ErrBufferFull will be returned
    If n is bigger than the size of stream, EOF will be returned

Let’s see how it works (source code):
```go
s1 := strings.NewReader(strings.Repeat("a", 20))
r := bufio.NewReaderSize(s1, 16)
b, err := r.Peek(3)
if err != nil {
    fmt.Println(err)
}
fmt.Printf("%q\n", b)
b, err = r.Peek(17)
if err != nil {
    fmt.Println(err)
}
s2 := strings.NewReader("aaa")
r.Reset(s2)
b, err = r.Peek(10)
if err != nil {
    fmt.Println(err)
}"aaa"
bufio: buffer full
EOF
```
 ##   Minimum size of buffer used by bufio.Reader is 16.

Returned slice uses the same underlying array as the internal buffer used by bufio.Reader. Consequently what is inside returned slice becomes invalid after any read operations done by reader under the hood. It’s because it might be overwritten by other buffered data (source code):
```go
s1 := strings.NewReader(strings.Repeat("a", 16) + strings.Repeat("b", 16))
r := bufio.NewReaderSize(s1, 16)
b, _ := r.Peek(3)
fmt.Printf("%q\n", b)
r.Read(make([]byte, 16))
r.Read(make([]byte, 15))
fmt.Printf("%q\n", b)"aaa"
"bbb"
```
## Reset

Buffered can be re-used in a similar way as bufio.Writer (source code):
```go
s1 := strings.NewReader("abcd")
r := bufio.NewReader(s1)
b := make([]byte, 3)
_, err := r.Read(b)
if err != nil {
    panic(err)
}
fmt.Printf("%q\n", b)
s2 := strings.NewReader("efgh")
r.Reset(s2)
_, err = r.Read(b)
if err != nil {
    panic(err)
}
fmt.Printf("%q\n", b)"abc"
"efg"
```
By using Reset we can avoid redundant allocations which frees GC from unnecessary work.
## Discard

This method throws away n bytes without even returning it. If bufio.Reader buffered so far more than or equal to n then it doesn’t have to read anything from io.Reader — it simply drops first n bytes from the buffer (source code):
```go
type R struct{}func (r *R) Read(p []byte) (n int, err error) {
    fmt.Println("Read")
    copy(p, "abcdefghijklmnop")
    return 16, nil
}func main() {
    r := new(R)
    br := bufio.NewReaderSize(r, 16)
    buf := make([]byte, 4)
    br.Read(buf)
    fmt.Printf("%q\n", buf)
    br.Discard(4)
    br.Read(buf)
    fmt.Printf("%q\n", buf)
}Read
"abcd"
"ijkl"
```
Call to Discard didn’t required reading more data from reader r. If on the other hand buffer has less than n bytes then bufio.Reader will read required amount of data making sure no less than n bytes will be discarded (source code):
```go
type R struct{}func (r *R) Read(p []byte) (n int, err error) {
    fmt.Println("Read")
    copy(p, "abcdefghijklmnop")
    return 16, nil
}func main() {
    r := new(R)
    br := bufio.NewReaderSize(r, 16)
    buf := make([]byte, 4)
    br.Read(buf)
    fmt.Printf("%q\n", buf)
    br.Discard(13)
    fmt.Println("Discard")
    br.Read(buf)
    fmt.Printf("%q\n", buf)
}Read
"abcd"
Read
Discard
"bcde"
```
Note that 2nd read call has been made because of call to Discard.
### Read

At the core of our bufio.Reader sits Read method. It has the same signature as the only method of io.Reader interface so bufio.Reader implements this omnipresent interface:

type Reader interface {
        Read(p []byte) (n int, err error)
}

Read method from bufio.Reader does maximum one read from the underlying io.Reader:

    If internal buffer holds at least one byte then no matter what is the size of the input slice (len(p)) method Read will get data only from the internal buffer without reading from the underlying reader (source code):
```go
func (r *R) Read(p []byte) (n int, err error) {
    fmt.Println("Read")
    copy(p, "abcd")
    return 4, nil
}func main() {
    r := new(R)
    br := bufio.NewReader(r)
    buf := make([]byte, 2)
    n, err := br.Read(buf)
    if err != nil {
        panic(err)
    }
    buf = make([]byte, 4)
    n, err = br.Read(buf)
    if err != nil {
        panic(err)
    }
    fmt.Printf("read = %q, n = %d\n", buf[:n], n)
}Read
read = "cd", n = 2
```
Our instance of io.Reader returns “abcd” indefinitely (never gives io.EOF). 2nd call to Read uses slice of length 4 but since internal buffer already holds “cd” after the first read from io.Reader then bufio.Reader returns everything from the buffer without event talking to underlying reader.

2. If internal buffer is empty then one reading from underlying io.Reader will be executed. It’s visible in the previous example where we started with empty buffer and call:

n, err := br.Read(buf)

triggered reading to fill the buffer.

3. If internal buffer is empty but passed slice is bigger than buffer then bufio.Reader will skip buffering and will read directly into passed slice of bytes (source code):
```
type R struct{}func (r *R) Read(p []byte) (n int, err error) {
    fmt.Println("Read")
    copy(p, strings.Repeat("a", len(p)))
    return len(p), nil
}func main() {
    r := new(R)
    br := bufio.NewReaderSize(r, 16)
    buf := make([]byte, 17)
    n, err := br.Read(buf)
    if err != nil {
        panic(err)
    }
    fmt.Printf("read = %q, n = %d\n", buf[:n], n)
    fmt.Printf("buffered = %d\n", br.Buffered())
}Read
read = "aaaaaaaaaaaaaaaaa", n = 17
buffered = 0
```
## Internal buffer doesn’t have any data (buffered = 0) after reading from bufio.Reader.
{Read, Unread}Byte

These methods have been implemented too either read single byte from the buffer or return last read byte back to the buffer (source code):
```go
r := strings.NewReader("abcd")
br := bufio.NewReader(r)
byte, err := br.ReadByte()
if err != nil {
    panic(err)
}
fmt.Printf("%q\n", byte)
fmt.Printf("buffered = %d\n", br.Buffered())
err = br.UnreadByte()
if err != nil {
    panic(err)
}
fmt.Printf("buffered = %d\n", br.Buffered())
byte, err = br.ReadByte()
if err != nil {
    panic(err)
}
fmt.Printf("%q\n", byte)
fmt.Printf("buffered = %d\n", br.Buffered())'a'
buffered = 3
buffered = 4
'a'
buffered = 3

{Read, Unread}Rune
```
These two work as previous methods but handling Unicode characters (UTF-8 encoded) instead.
## ReadSlice

Function returns bytes till first occurrence of passed byte:
```go
func (b *Reader) ReadSlice(delim byte) (line []byte, err error)

Example (source code):

s := strings.NewReader("abcdef|ghij")
r := bufio.NewReader(s)
token, err := r.ReadSlice('|')
if err != nil {
    panic(err)
}
fmt.Printf("Token: %q\n", token)Token: "abcdef|"

    Important to keep in mind that returned slice points to the internal buffer so it can be overwritten during next read operation.

If delimiter cannot be found and EOF has been reached then returned error will be io.EOF. To test it let’s change one line in program above (source code):

s := strings.NewReader("abcdefghij")

which ends up with crash: panic: EOF. Discussed method will return io.ErrBufferFull when delimiter cannot be found and no more data can fit into internal buffer (source code):

s := strings.NewReader(strings.Repeat("a", 16) + "|")
r := bufio.NewReaderSize(s, 16)
token, err := r.ReadSlice('|')
if err != nil {
    panic(err)
}
fmt.Printf("Token: %q\n", token)

This piece of code leads error: panic: bufio: buffer full.
ReadBytes

func (b *Reader) ReadBytes(delim byte) ([]byte, error)
```
Returns slice of bytes until the first occurrence of delimiter. It has the same signature as ReadSlice which is low-level function and is actually used underneath by ReadBytes (code). What is the difference then? ReadBytes can call ReadSlice multiple times if separator hasn’t been found and can accumulate returned data. It means that ReadBytes isn’t restricted by the buffer’s size (source code):
```go
s := strings.NewReader(strings.Repeat("a", 40) + "|")
r := bufio.NewReaderSize(s, 16)
token, err := r.ReadBytes('|')
if err != nil {
    panic(err)
}
fmt.Printf("Token: %q\n", token)Token: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa|"
```
Additionally it returns new slice of bytes so no risk that data will be overwritten by future read operations.
## ReadString
```go
It’s a simple wrapper over ReadBytes discussed above (code):

func (b *Reader) ReadString(delim byte) (string, error) {
    bytes, err := b.ReadBytes(delim)
    return string(bytes), err
}
```
## ReadLine

ReadLine() (line []byte, isPrefix bool, err error)

Uses ReadSlice underneath (ReadSlice('\n')) but also takes care of removing new-line characters (\n or \r\n) from returned slice. Signature is different than ReadBytes or ReadSlice since it contains isPrefix flag which is true when delimiter hasn’t been found because internal buffer couldn’t hold more data (source code):
```go
s := strings.NewReader(strings.Repeat("a", 20) + "\n" + "b")
r := bufio.NewReaderSize(s, 16)
token, isPrefix, err := r.ReadLine()
if err != nil {
    panic(err)
}
fmt.Printf("Token: %q, prefix: %t\n", token, isPrefix)
token, isPrefix, err = r.ReadLine()
if err != nil {
    panic(err)
}
fmt.Printf("Token: %q, prefix: %t\n", token, isPrefix)
token, isPrefix, err = r.ReadLine()
if err != nil {
    panic(err)
}
fmt.Printf("Token: %q, prefix: %t\n", token, isPrefix)
token, isPrefix, err = r.ReadLine()
if err != nil {
    panic(err)
}Token: "aaaaaaaaaaaaaaaa", prefix: true
Token: "aaaa", prefix: false
Token: "b", prefix: false
panic: EOF
```

This method doesn’t give any information if the last returned slice ends with new line character (source code):
```
s := strings.NewReader("abc")
r := bufio.NewReaderSize(s, 16)
token, isPrefix, err := r.ReadLine()
if err != nil {
    panic(err)
}
fmt.Printf("Token: %q, prefix: %t\n", token, isPrefix)
s = strings.NewReader("abc\n")
r.Reset(s)
token, isPrefix, err = r.ReadLine()
if err != nil {
    panic(err)
}
fmt.Printf("Token: %q, prefix: %t\n", token, isPrefix)Token: "abc", prefix: false
Token: "abc", prefix: false
```
## WriteTo

bufio.Reader implements io.WriterTo interface:
```go
type WriterTo interface {
        WriteTo(w Writer) (n int64, err error)
}
```
It allows to pass consumer implementing io.Writer and all data will be read from the producer and send further to passed consumer . Let’s see how it works in practise (source code):
```go
type R struct {
    n int
}func (r *R) Read(p []byte) (n int, err error) {
    fmt.Printf("Read #%d\n", r.n)
    if r.n >= 10 {
         return 0, io.EOF
    }
    copy(p, "abcd")
    r.n += 1
    return 4, nil
}func main() {
    r := bufio.NewReaderSize(new(R), 16)
    n, err := r.WriteTo(ioutil.Discard)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Written bytes: %d\n", n)
}Read #0
Read #1
Read #2
Read #3
Read #4
Read #5
Read #6
Read #7
Read #8
Read #9
Read #10
Written bytes: 40
```
## bufio.Scanner
In-depth introduction to bufio.Scanner in Golang
Go is shipped with package helping with buffered I/O — technique to optimize read or write operations. For writes it’s…

```go
ReadBytes('\n') or ReadString('\n') or ReadLine or Scanner?

ReadString('\n') as discussed before is simple wrapper around ReadBytes('\n') so let’s discuss differences other three.

    ReadBytes doesn’t handle \r\n sequence automatically (source code):

s := strings.NewReader("a\r\nb")
r := bufio.NewReader(s)
for {
    token, _, err := r.ReadLine()
    if len(token) > 0 {
        fmt.Printf("Token (ReadLine): %q\n", token)
    }
    if err != nil {
        break
    }
}
s.Seek(0, io.SeekStart)
r.Reset(s)
for {
    token, err := r.ReadBytes('\n')
    fmt.Printf("Token (ReadBytes): %q\n", token)
    if err != nil {
        break
    }
}
s.Seek(0, io.SeekStart)
scanner := bufio.NewScanner(s)
for scanner.Scan() {
    fmt.Printf("Token (Scanner): %q\n", scanner.Text())
}Token (ReadLine): "a"
Token (ReadLine): "b"
Token (ReadBytes): "a\r\n"
Token (ReadBytes): "b"
Token (Scanner): "a"
Token (Scanner): "b"
```
ReadBytes returns the slice together with delimiter so it requires a bit extra work to refine the data (unless delimiter in returned slice is actually useful).

2. ReadLine doesn’t handle lines longer than internal buffer (source code):
```go
s := strings.NewReader(strings.Repeat("a", 20) + "\n")
r := bufio.NewReaderSize(s, 16)
token, _, _ := r.ReadLine()
fmt.Printf("Token (ReadLine): \t%q\n", token)s.Seek(0, io.SeekStart)
r.Reset(s)
token, _ = r.ReadBytes('\n')
fmt.Printf("Token (ReadBytes): \t%q\n", token)s.Seek(0, io.SeekStart)
scanner := bufio.NewScanner(s)
scanner.Scan()
fmt.Printf("Token (Scanner): \t%q\n", scanner.Text())Token (ReadLine): 	"aaaaaaaaaaaaaaaa"
Token (ReadBytes): 	"aaaaaaaaaaaaaaaaaaaa\n"
Token (Scanner): 	"aaaaaaaaaaaaaaaaaaaa"
```
ReadLine needs to be called for the 2nd time to retrieve rest of the stream. Max size of the token which is handled by Scanner is 64 * 1024. If longer token are passed then scanner won’t be able to parse anything. ReadLine when called multiple times can handle token of any size since it returns prefix of buffered data is delimiter not found — but this needs to be handled by caller. ReadBytes doesn’t have any limit (source code):
```go
s := strings.NewReader(strings.Repeat("a", 64*1024) + "\n")
r := bufio.NewReader(s)
token, _, err := r.ReadLine()
fmt.Printf("Token (ReadLine): %d\n", len(token))
fmt.Printf("Error (ReadLine): %v\n", err)s.Seek(0, io.SeekStart)
r.Reset(s)
token, err = r.ReadBytes('\n')
fmt.Printf("Token (ReadBytes): %d\n", len(token))
fmt.Printf("Error (ReadBytes): %v\n", err)s.Seek(0, io.SeekStart)
scanner := bufio.NewScanner(s)
scanner.Scan()
fmt.Printf("Token (Scanner): %d\n", len(scanner.Text()))
fmt.Printf("Error (Scanner): %v\n", scanner.Err())Token (ReadLine): 4096
Error (ReadLine): <nil>
Token (ReadBytes): 65537
Error (ReadBytes): <nil>
Token (Scanner): 0
Error (Scanner): bufio.Scanner: token too long
```
3. Scanner has the simplest API as visible above and provides nicest abstraction for common cases.
## bufio.ReadWriter

Structs in Go allow for something called type embedding. Instead of regular field with name and type we can put only type (anonymous field). Methods / fields from embedded types if don’t collide with others can be referenced with short selectors (source code):
```go
type T1 struct {
    t1 string
}func (t *T1) f1() {
    fmt.Println("T1.f1")
}type T2 struct {
    t2 string
}func (t *T2) f2() {
    fmt.Println("T1.f2")
}type U struct {
    *T1
    *T2
}func main() {
    u := U{T1: &T1{"foo"}, T2: &T2{"bar"}}
    u.f1()
    u.f2()
    fmt.Println(u.t1)
    fmt.Println(u.t2)
}T1.f1
T1.f2
foo
bar
```

Instead of e.g. u.T1.t1 we can use simply u.t1. Package bufio uses embedding to define ReadWriter which is a composition both Reader and Writer:
```go
type ReadWriter struct {
  	*Reader
  	*Writer
  }

Let’s see how it works (source code):

s := strings.NewReader("abcd")
br := bufio.NewReader(s)
w := new(bytes.Buffer)
bw := bufio.NewWriter(w)
rw := bufio.NewReadWriter(br, bw)
buf := make([]byte, 2)
_, err := rw.Read(buf)
if err != nil {
    panic(err)
}
fmt.Printf("%q\n", buf)
buf = []byte("efgh")
_, err = rw.Write(buf)
if err != nil {
    panic(err)
}
err = rw.Flush()
if err != nil {
   panic(err)
}
fmt.Println(w.String())"ab"
efgh
```
To read amount of buffered data rw.Buffered() won’t be much of an use. Compiler returns an error ambiguous selector rw.Buffered since both reader and writer have method Buffered. Something like rw.Reader.Buffered() will work though.
bufio + standard library

bufio package is used extensively across the standard library where I/O comes into play like:

    archive/zip
    compress/*
    encoding/*
    image/*
    net/http for things like wrapping TCP connections (source code). It also combines structures for buffered I/O with e.g. sync.Pool to limit pressure on GC (source code).

👏 below to help others discover this story. Please follow me here or on Twitter if you want to get updates about new posts or boost work on future stories.
