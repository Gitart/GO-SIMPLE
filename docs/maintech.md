
## Prologue  
This is a **very** *rough* draft of a portion of a **larger work I am writing**, but this itself was a lot of work so I wanted to get it out in the world to get feedback on it that I can. 
  
**Note:** I do not plan to keep the Gist live permanently. If I did search engines would likely preferring it here over where I will eventually publish it.  

If you want to know where I ultimately move this post comment here so GitHub will notify you when I post that comment.  
  
  
# Decoding Interfaces in Go 
## — A Comprehensive Categorization of Golang’s StdLib

In my experience — and especially given my analysis for purposes of identifying — there are many different patterns of use for interfaces in Go _(though, a discussion of [type constraint](https://go.dev/blog/intro-generics#type-sets) use-cases being out of scope here.)_  
  
While there is clearly overlap in these use-cases as you will see numerous interfaces mentions in more than one category, but I think each use-case is distinct enough to be recognized.

## My source: the standard library for Go `v1.21.0`
I prepared this categorization from [Go v1.21.0][go1.21.0] per chance you find a need to cross-reference this work to the source code. 
  
  
## Patterns of interface use  
  
Arguably one of the most common ways interfaces are used in Go is as [hooks](https://en.wikipedia.org/wiki/Hooking) and/or [filters](https://en.wikipedia.org/wiki/Filter_(higher-order_function)).  
  
However, these hooks and filters can also be categorized by many different patterns of use. The Go standard library has many examples use-cases for both single method interfaces and of multi-method interfaces.  
  
### About categorization  
I distilled a list using category names of my own choosing. I deliberately steered clear up well-known pattern names to avoid the assumptions many who know those patterns may carry from other languages.  
  
To categorize I looked at every implementable public method in the Go standard library and tried to understand its use and then fit it into a least one category. I added many terms as categories that I later rejected because there were too few examples or it became clear to me a different category represented their essence better.  
  
I also added categories that I did not originally envision simply because there were so many different examples of that category, such as reading and writing.  
  
I further took categories without many examples, or too many examples, and moved those to the higher-level discussion of overall reasons to use interfaces _(those discussions are not in this document at this time.)_
  
I tried my best to be exhaustive, but I certainly failed as my categorizes were based on those in the Go standard library and any packages it vendored, but not any other codebase. This means these categories are skewed towards those which makes sense across the Go standard library, and not across all of Go code globally.
  
I am also only human so in that respect I certainly erred. But the key point, is I tried, and I think there is value in that. Or at least I hope some people find value in these categories.  
  
### The categories  
  
The Go standard library by its nature will be heavily weighted to some of the categories more than other whereas 3rd party packages and applications are likely to be weighted more so some of the lesser represented categories here.  
  
Here they are in alphabetic order:  
- [Algorithms](#algorithm-focused)  
- [Data](#data-focused)  
- [Drivers](#driver-focused)  
- [File systems](#file-system-focused)  
- [Getting and/or Setting](#getting-andor-setting-focused)  
- [Implementation](#implementation-focused)  
- [Membership](#membership-focused)  
- [Parsing](#parsing-focused)  
- [Process control](#process-control-focused)  
- [Protocols](#protocol-focused)  
- [Reading and/or Writing](#reading-andor-writing-focused)  
- [Representation](#representation-focused)  
- [Subtypes](#subtype-focused)  
  
### Legend for vendored packages  
  
| Reference               | Package                            |  
| ----------------------- | ---------------------------------- |  
| `*/pprof/driver`        | `github.com/google/pprof/driver`   |  
| `*/arch/arm/armasm`     | `golang.org/x/arch/arm/armasm`     |  
| `*/arch/ppc64/ppc64asm` | `golang.org/x/arch/ppc64/ppc64asm` |  
| `*/mod/modfile`         | `golang.org/x/arch/mod/modfile`    |  
| `*/mod/sumdb`           | `golang.org/x/arch/mod/sumdb`      |  
| `*/mod/zip`             | `golang.org/x/arch/mod/zip`        |  
| `*/tools/go/analysis`   | `golang.org/x/tools/go/analysis`   |  
| `*/crypto/cryptobyte`   | `golang.org/x/crypto/cryptobyte`   |  
| `*/net/lif`             | `golang.org/x/net/lif`             |  
| `*/net/route`           | `golang.org/x/net/route`           |
  
## Algorithm-focused  
Interfaced focus on **Algorithms** serve as either a framework for configuring the behavior of a general-purpose algorithm, or actually implement the algorithm that a framework for algorithms calls into. These interfaces define a set of methods that must be implemented to specify the rules or criteria under which the algorithm operates.  
  
They allow developers to tailor a generic algorithm to meet the requirements of a particular use case, without altering the algorithm's core logic. They also act as a bridge between the algorithm and its operational context, offering a structured way to introduce variability in how the algorithm functions.  
  
| Package               | Interface(s)       | Method signature(s) and/or Interface embed(s)                                      |  
| --------------------- | ------------------ | ---------------------------------------------------------------------------------- |  
| `container/heap`      | `Interface`        | `Push()`, `Pop()`, `sort.Interface`                                                |  
| `compress/flate`      | `Reader`           | `io.Reader`,`io.ByteReader`                                                        |  
| `compress/flate`      | `Resetter`         | `Reset()`                                                                          |  
| `compress/zlib`       | `Resetter`         | `Reset()`                                                                          |  
| `crypto`              | `Decypter`         | `Public()`, `Decrypt()`                                                            |  
| `crypto`              | `Signer`           | `Public()`,`Sign()`                                                                |  
| `crypto`              | `SignerOpts`       | `HashFunc()`                                                                       |  
| `crypto/cipher`       | `Block`            | `BlockSize()`, `Encrypt()`, `Decrypt()`                                            |  
| `crypto/cipher`       | `AEAD`             | `NonceSize()`, `Overhead()`, `Seal()`, `Open()`                                    |  
| `crypto/elliptic`     | `Curve`            | `Params()`, `IsOnCurve()`, `Add()`, `Double()`, `ScalarMult()`, `ScalarBaseMult()` |  
| `encoding/json`       | `Marshaler`        | `MarshalJSON()`                                                                    |  
| `encoding/json`       | `Unmarshaler`      | `UnmarshalJSON()`                                                                  |  
| `encoding/xml`        | `Marshaler`        | `MarshalXML()`                                                                     |  
| `encoding/xml`        | `Marshaler`        | `MarshalXMLAttr()`                                                                 |  
| `encoding/xml`        | `Unmarshaler`      | `UnmarshalXML()`                                                                   |  
| `encoding/xml`        | `UnmarshalerAttr`  | `UnmarshalXMLAttr()`                                                               |  
| `encoding/xml`        | `TokenReader`      | `Token()`                                                                          |  
| `go/ast`              | `Node`             | `Pos()`, `End()`                                                                   |  
| `go/ast`              | `Vistor`           | `Visit()`                                                                          |  
| `go/types`            | `Importer`         | `Import()`                                                                         |  
| `go/types`            | `ImporterFrom`     | `ImportFrom()`                                                                     |  
| `hash`                | `Hash`             | `io.Writer`, `Blocksize()`, `Reset()`, `Size()`, `Sum()`                           |  
| `hash`                | `Hash32`           | `Sum32()`                                                                          |  
| `hash`                | `Hash64`           | `Sum64()`                                                                          |  
| `math/rand`           | `Source`           | `Int63()`, `Seed()`                                                                |  
| `math/rand`           | `Source64`         | `Int64()`                                                                          |  
| `net/http/cookiejar`  | `PublicSuffixList` | `PublicSuffix()`, `String()`                                                       |  
| `sort`                | `Interface`        | `Len()`, `Less()`, `Swap()`                                                        |  
| `*/mod/sumdb/note`    | `Signer`           | `KeyHash()`, `Name()`, `Sign()`                                                    |  
| `*/mod/sumdb/note`    | `Verifier`         | `KeyHash()`, `Name()`, `Verify()`                                                  |  
| `*/mod/sumdb/tlog`    | `TileReader`       | `Height()`, `ReadTiles()`, `SaveTiles()`                                           |  
| `*/mod/sumdb/tlog`    | `HashReader`       | `ReadHashes()`                                                                     |

See also:  
- [Implementation-focused](#implementation-focused)  
  
## Data-focused  
A common use for Interfaces focused on accessing, converting, filtering, formatting  and/or validating **data**, of which `fmt.Stringer` is exceedingly common.  
  
I'm omitting the copious `Reader` and `Writer` interfaces because they are represented in their own categories too. Also omitting any that are represented in the algorithm category as there is significant cross-over.  

  
| Package               | Interface             | Method signature(s) and/or Interface embed(s)                                               |  
| --------------------- | --------------------- | ------------------------------------------------------------------------------------------- |  
| `crypto/cipher`       | `Stream`              | `XORKeyStream()`                                                                            |  
| `database/sql/driver` | `NamedValueChecker`   | `CheckNameValue()`                                                                          |  
| `database/sql/driver` | `ValueConverter`      | `ConvertValue()`                                                                            |  
| `database/sql/driver` | `Validator`           | `IsValid()`                                                                                 |  
| `database/sql/driver` | `Valuer`              | `Value()`                                                                                   |  
| `debug/macho`         | `Load`                | `Raw()`                                                                                     |  
| `encoding`            | `BinaryMarshaler`     | `MarshalBinary()`                                                            |  
| `encoding`            | `BinaryUnmarshaler`   | `UnmarshalBinary()`                                                            |  
| `encoding`            | `TextMarshaler`       | `MarshalText()`                                                              |  
| `encoding`            | `TextUnmarshaler`     | `UnmarshalText()`                                                               |  
| `encoding/binary`     | `AppendByteOrder`     | `AppendUint16()`, `AppendUint32()`, `AppendUint64()`, `String()`                            |  
| `encoding/binary`     | `ByteOrder`           | `Uint16()`, `Uint32()`, `Uint64()`, `PutUint16()`, `PutUint32()`, `PutUint64()`, `String()` |  
| `encoding/gob`        | `GobDecoder`          | `GobDecode()`                                                                               |  
| `encoding/gob`        | `GobEncoder`          | `GobEncode()`                                                                               |  
| `expvar`              | `Var`                 | `String()`                                                                                  |  
| `fmt`                 | `Formatter`           | `Format()`                                                                                  |  
| `fmt`                 | `GoStringer`          | `GoString()`                                                                                |  
| `fmt`                 | `State`               | `Flag()`, `Precision()`, `Width()`, `Write()`                                               |  
| `fmt`                 | `Stringer`            | `String()`                                                                                  |  
| `fmt`                 | `Scanner`             | `Scan()`                                                                                    |  
| `go/types`            | `Sizes`               | `Alignof()`, `Offsetsof()`, `Sizeof()`                                                      |  
| `go/types`            | `Type`                | `Underlying()`, `String()`                                                                  |  
| `image/color`         | `Model`               | `Convert()`                                                                                 |  
| `io/fs`               | `FileInfo`            | `Name()`, `Size()`, `Mode()`, `ModeTime()`, `IsDir()`, `Sys()`                              |  
| `log/slog`            | `LogValuer`           | `LogValue()`                                                                                |  
| `net`                 | `Addr`                | `Network()`,`String()`                                                                      |  
| `*/crypto/cryptobyte` | `MarshalingValue`     | `Marshal()`                                                                                 |  
| `*/text/transform`    | `Transformer`         | `Transform()`, `Reset()`                                                                    |  
| `*/text/transform`    | `SpanningTransformer` | `Transformer`, `Span()`                                                                     |
  
See also:  
- [Algorithm-focused](#algorithm-focused)  
- [Reading and/or writing-focused](#reading-andor-writing-focused)  
- [File system-focused](#file-system-focused)  
- [Get and/or set-focused](#get-andor-set-focused)  
  
## Driver-focused  
Interfaces focused on enable interaction between software and external software and/or devices are often referred to as **drivers**. Adding an external driver interface to be used by to your software allow 3rd parties to write conforming drivers for your software and/or hardware.  
  
The `database/sql` package has many of these and is probably the best example in the Go standard library, and these are some but not all:  
  
| Package               | Interface           | Method signature(s) and/or Interface embed(s)                                                             |  
| --------------------- | ------------------- |-----------------------------------------------------------------------------------------------------------|  
| `database/sql/driver` | `Conn`              | `Prepare()`, `Begin()`, `Close()`                                                                         |  
| `database/sql/driver` | `Connector`         | `Connect()`, `Driver()`                                                                                   |  
| `database/sql/driver` | `Driver`            | `Open()`                                                                                                  |  
| `database/sql/driver` | `Execer`            | `Exec()`                                                                                                  |  
| `database/sql/driver` | `NamedValueChecker` | `CheckNameValue()`                                                                                        |  
| `database/sql/driver` | `Pinger`            | `Pinger()`                                                                                                |  
| `database/sql/driver` | `Queryer`           | `Query()`                                                                                                 |  
| `database/sql/driver` | `Result`            | `LastInsertId()`, `RowsAffected()`                                                                        |  
| `database/sql/driver` | `Rows`              | `Columns()`, `Next()`, `Close()`                                                                          |  
| `database/sql/driver` | `Scanner`           | `Scan()`                                                                                                  |  
| `database/sql/driver` | `SessionResetter`   | `ResetSession()`                                                                                          |  
| `database/sql/driver` | `Stmt`              | `NumInput()`, `Exec()`, `Query()`, `Close()`                                                              |  
| `database/sql/driver` | `Tx`                | `Commit()`, `Rollback()`                                                                                  |  
| `database/sql/driver` | `ValueConverter`    | `ConvertValue()`                                                                                          |  
| `database/sql/driver` | `Valuer`            | `Value()`                                                                                                 |  
| `log/slog`            | `Handler`           | `WithAttrs()`, `WithGroup()`, `Enabled()`, `Handle()`                                                     |  
| `*/pprof/driver`      | `Fetcher`           | `Fetch()`                                                                                                 |  
| `*/pprof/driver`      | `FlagSet`           | `AddExtraUsage()`,  `Bool()`, `ExtraUsage()`, `Float64()`, `Int()`, `Parse()`, `String()`, `StringList()` |  
| `*/pprof/driver`      | `ObjFile`           | `BuildID()`, `Name()`, `ObjAddr()`                                                                        |  
| `*/pprof/driver`      | `ObjTool`           | `Disasm()`, `Open()`                                                                                      |  
| `*/pprof/driver`      | `Symbolizer`        | `Symbolize()`                                                                                             |  
| `*/pprof/driver`      | `UI`                | 'IsTerminal()`, 'Print()`,  'PrintErr()`,  'ReadLine()`, 'SetAutoComplete()`,  'WantBrowser()`,           |  
| `*/pprof/driver`      | `Writer`            | `Open()`                                                                                                  |


## File system-focused  
These interfaces are focused on allowing interaction with or simulating a file-system.  
  
Because these interfaces are so broadly categorizable I felt they deserved their own category.  

| Package          | Interface(s)  | Method signature(s) and/or Interface embed(s)                  |  
| ---------------- | ------------- | -------------------------------------------------------------- |  
| `cmd/pack`       | `FileLike`    | `Read()`, `Close()`, `Name()`, `Stat()`                        |  
| `io/fs`          | `DirEntry`    | `Info()`, `IsDir()`, `Name()`, `Type()`                        |  
| `io/fs`          | `FS`          | `Open()`                                                       |  
| `io/fs`          | `File`        | `Close()`, `Read()`, `Stat()`                                  |  
| `io/fs`          | `FileInfo`    | `Name()`, `Size()`, `Mode()`, `ModeTime()`, `IsDir()`, `Sys()` |  
| `io/fs`          | `ReadDirFile` | `ReadDir()`                                                    |  
| `io/fs`          | `GlobFS`      | `FS`, `Glob()`                                                 |  
| `io/fs`          | `ReadDirFS`   | `ReadDir()`                                                    |  
| `io/fs`          | `ReadFileFS`  | `ReadFile()`                                                   |  
| `io/fs`          | `StatFS`      | `Stat()`                                                       |  
| `io/fs`          | `SubFS`       | `Sub()`                                                        |  
| `mime/multipart` | `File`        | `io.Reader`, `io.ReaderAt`, `io.Seeker`, `io.Closer`           |  
| `net/http`       | `File`        | `io.Reader`, `io.Seeker`, `io.Closer`, `ReadDir()`, `Stat()`   |  
| `net/http`       | `FileSystem`  | `Open()`                                                       |  
| `syscall`        | `Conn`        | `SyscallConn()`                                                |  
| `*/pprof/driver` | `ObjTool`     | `Disasm()`, `Open()`                                           |  
| `*/mod/zip`      | `File`        | `LStat()`, `Path()`, `Open()`                                  |

See also:  
- [Reading and/or writing-focused](#reading-andor-writing-focused)  
- [Data-focused](#data-focused)  
- [Subtype-focused](#subtype-focused)  
- [Implementation-focused](#implementation-focused)  
- [Process control-focused](#process-control-focused)  
- [Representation-focused](#representation-focused)  
  
## Getting and/or setting-focused  
Some interfaces are as simply as enabling an implementor to **get** and/or **set** a property that they do not directly control. This allows the user of the func or type to control that property instead.  
  
Many other interfaces besides those listed here offer a form of getting and/or setting. Arguably marshalers and unmarshalers as well as encoders and decoders could be classified as get/set focused, but the ones I listed here I viewed as being most specifically associated with get and set and less as incidental to getting and setting.  
  
| Package             | Interface            | Method signature(s) and/or Interface embed(s) |     |  
| ------------------- | -------------------- | --------------------------------------------- | --- |  
| `crypto/tls`        | `ClientSessionCache` | `Get()`, `Put()`                              |     |  
| `flag`              | `Getter`             | `Get()`                                    |     |  
| `flag`              | `Value`              | `String()`<br> `Set()`             |     |  
| `image/png`         | `EncoderBufferPool`  | `Get()`, `Put()`                              |     |  
| `log/slog`          | `Leveler`            | `Level()`                                     |     |  
| `log/slog`          | `LogValuer`          | `LogValue()`                                  |     |  
| `net`               | `Addr`               | `Network()`,`String()`                        |     |  
| `net`               | `Error`              | `error`, `Timeout()`,`Temporary()`            |     |  
| `net/http`          | `Cookiejar`          | `Cookies()`, `SetCookies()`                   |     |  
| `net/http/httputil` | `BufferPool`         | `Get()`, `Put()`                              |     |  
| `*/mod/sumdb/note`  | `Verifiers`          | `Verifier()`                                  |     |  
  
  
## Implementation-focused  
Some interface interfaces are focused on providing a method signature for **implementation** of the core purpose of the type.  

| Package              | Interface          | Method signature(s) and/or Interface embed(s)                                                                             |  
| -------------------- | ------------------ | ------------------------------------------------------------------------------------------------------------------------- |  
| `image/draw`         | `Drawer`           | `Draw()`                                                                                                                  |  
| `image/draw`         | `Image`            | `image.Image`,`Set()`                                                                                                     |  
| `image/draw`         | `Quantizer`        | `Quantize()`                                                                                                              |  
| `image/draw`         | `RGBA64Image`      | `image.RGBA64Image`,`Set()`,`SetRGBA64()`                                                                                 |  
| `log/slog`           | `Handler`          | `WithAttrs()`, `WithGroup()`, `Enabled()`, `Handle()`                                                                     |  
| `log/slog`           | `LogValuer`        | `LogValue()`                                                                                                              |  
| `net`                | `Conn`             | `RemoteAddr()`, `SetDeadline()`, `SetReadDeadline()`, `SetWriteDeadline()`, `Read()`, `Write()`, `Close()`, `LocalAddr()` |  
| `net`                | `Listener`         | `Accept()`, `Addr()`, `Close()`                                                                                           |  
| `net`                | `PacketConn`       | `Close()`, `LocalAddr()`, `ReadFrom()`, `SetDeadline()`, `SetReadDeadline()`, `SetWriteDeadline()`, `WriteTo()`           |  
| `net/http`           | `RoundTripper`     | `RoundTrip()`                                                                                                             |  
| `net/http/cookiejar` | `PublicSuffixList` | `PublicSuffix()`, `String()`                                                                                              |  
| `testing/quick`      | `Generator`        | `Generate()`                                                                                                              |  
| `*/mod/sumdb`        | `ClientOps`        | `SecurityError()`, `ReadRemote()`, `ReadConfig()`, `WriteConfig()`, `ReadCache()`, `WriteCache()`, `Log()`                |  
| `*/mod/sumdb`        | `ServerOps`        | `Lookup()`, `ReadRecords()`, `ReadTileData()`, `Signed()`                                                                 |

See also:  
- [Algorithm-focused](#algorithm-focused)  
- [File system-focused](#file-system-focused)  
- [Reading and/or writing-focused](#reading-andor-writing-focused)  
- [Protocol-focused](#protocol-focused)  
  
## Membership-focused  
Some interfaces exists merely to allow a type to signal that it has **membership** in a group, i.e. a `Widget` interface could require a `Widget` method that does nothing other than allow an implementing type to indicate that it is Widget.  
  
| Package                            | Interface | Method signature(s) and/or Interface embed(s) |  
| ---------------------------------- | --------- |-----------------------------------------------|  
| `runtime`                          | `Error`   | `error`, `RuntimeError()`                     |  
| `golang.org/x/arch/arm/armasm`     | `Arg`     | `isArg()`, `String()`                         |  
| `golang.org/x/arch/ppc64/ppc64asm` | `Arg`     | `isArg()`, `String()`                         |  
| `golang.org/x/tools/go/analysis`   | `Fact`    | `AFact()`                                          |
  
## Parsing-focused  
Some interfaces are designed to facilitate **parsing** of a file, language or other structure.  
  
  
| Package               | Interface | Method signature(s) and/or Interface embed(s) |
| --------------------- | --------- | --------------------------------------------- |
| `go/ast`              | `Node`    | `Pos()`, `End()`                              |
| `go/ast`              | `Vistor`  | `Visit()`                                     |
| `*/mod/modfile`       | `Expr`    | `Comment()`, `Span()`                         |
| `*/tools/go/analysis` | `Range`   | `Pos()`, `End()`                              |

See Also:  
- [Process control-focused](#process-control-focused)  
- [Data-focused](#data-focused)  
  
## Process control-focused  
Some interfaces are designed to allowed to delegate **process control** to a developer implementing an interface.  
  
| Package          | Interface         | Method signature(s) and/or Interface embed(s)                             |  
| ---------------- | ----------------- | ------------------------------------------------------------------------- |  
| `container/heap` | `Interface`       | `Push()`, `Pop()`, `sort.Interface`                                       |  
| `context`        | `Context`         | `Deadline()`, `Done()`, `Err()`, `Value()`                                |  
| `database/sql`   | `SessionResetter` | `ResetSession()`                                                          |  
| `fmt`            | `ScanState`       | `Read()`, `ReadRune()`, `SkipSpace()`,`Token()`, `UnreadRune()`,`Width()` |  
| `http`           | `RoundTripper`    | `RoundTrip()`                                                             |  
| `log/slog`       | `Leveler`         | `Level()`                                                                 |  
| `sync`           | `Lock`            | `Lock()`, `Unlock()`                                                      |  
| `sync`           | `PoolDequeue`     | `PushHead()`, `PopHead()`, `PopTail()`                                    |  
| `os`             | `Signal`          | `Signal()`, `String()`                                                    |  
| `runtime`        | `Error`           | `error`, `RuntimeError()`                                                 |  
| `mutex`          | `Locker`          | `Lock()`, `Unlock()`                                                      |  
| `syscall`        | `RawConn`         | `Control()`, `Read()`, `Write()`                                          |  
| `*/net/lif`      | `Addr`            | `Family()`                                                                |  
| `*/net/route`    | `Addr`            | `Family()`                                                                |

See also:  
- [File system-focused](#file-system-focused)  
- [Driver-focused](#driver-focused)  
- [Implementation-focused](#implementation-focused)  
- [Protocol-focused](#protocol-focused)  
  
## Protocol-focused  
There are uses of interface designed to allow for interaction with a standardized and/or well-known **protocols**:    
  
| Package              | Interface          | Method signature(s) and/or Interface embed(s)                                                                             |  
| -------------------- | ------------------ | ------------------------------------------------------------------------------------------------------------------------- |  
| `net`                | `Conn`             | `RemoteAddr()`, `SetDeadline()`, `SetReadDeadline()`, `SetWriteDeadline()`, `Read()`, `Write()`, `Close()`, `LocalAddr()` |  
| `net`                | `Error`            | `error`, `Timeout()`,`Temporary()`                                                                                        |  
| `net`                | `Listener`         | `Accept()`, `Addr()`, `Close()`                                                                                           |  
| `net`                | `PacketConn`       | `Close()`, `LocalAddr()`, `ReadFrom()`, `SetDeadline()`, `SetReadDeadline()`, `SetWriteDeadline()`, `WriteTo()`           |  
| `net/http`           | `RoundTripper`     | `RoundTrip()`                                                                                                             |  
| `net/http`           | `Pusher`           | `Push()`                                                                                                                  |  
| `net/http`           | `Cookiejar`        | `Cookies()`, `SetCookies()`                                                                                               |  
| `net/http`           | `Handler`          | `ServeHTTP()`                                                                                                             |  
| `net/http`           | `Flusher`          | `Flush()`                                                                                                                 |  
| `net/http`           | `Hijacker`         | `Hijack()`                                                                                                                |  
| `net/http`           | `ResponseWriter`   | `Header()`, `Write()`, `WriteHeader()`                                                                                    |  
| `net/http/cookiejar` | `PublicSuffixList` | `PublicSuffix()`, `String()`                                                                                              |  
| `net/rpc`            | `ClientCodec`      | `WriteRequest()`, `ReadResponseHeader()`, `ReadResponseBody()`, `Close()`                                                 |  
| `net/rpc`            | `ServerCodec`      | `WriteResponse()`, `ReadRequestHeader()`, `ReadRequestBody()`, `Close()`                                                  |  
| `net/smtp`           | `Auth`             | `Start()`, `Next()`                                                                                                       |  
| `syscall`            | `Conn`             | `SyscallConn()`                                                                                                           |

See also:  
- [Process control-focused](#process-control-focused)  
  
## Reading and/or writing-focused  
Probably the most commonly used interfaces in Go are **readers** and/or **writers**. These interfaces allow developers to write without being concerned for the implementation of writing and/or they allow a developer implement the reading or writing of data for use by methods that reads and/or write to such an interface.  
  
I've also included closely-related `Seek`, and `Close` methods:  
  
| Package          | Interface(s)   | Method signature(s) and/or Interface embed(s)                                                       |  
| ---------------- | -------------- | --------------------------------------------------------------------------------------------------- |  
| `compress/flate` | `Reader`       | `io.Reader`,`io.ByteReader`                                                                         |  
| `fmt`            | `ScanState`    | `Read()`, `ReadRune()`, `SkipSpace()`,`Token()`, `UnreadRune()`,`Width()`                           |  
| `io`             | `ByteReader`   | `ReadByte()`                                                                                        |  
| `io`             | `ByteScanner`  | `UnreadByte()`                                                                                      |  
| `io`             | `ByteWriter`   | `WriteByte()`                                                                                       |  
| `io`             | `Closer`       | `Close()`                                                                                           |  
| `io`             | `Reader`       | `Read()`                                                                                            |  
| `io`             | `ReaderAt`     | `ReadAt()`                                                                                          |  
| `io`             | `ReaderFrom`   | `ReadFrom()`                                                                                        |  
| `io`             | `RuneReader`   | `ReadRune()`                                                                                        |  
| `io`             | `RuneScanner`  | `UnreadRune()`                                                                                      |  
| `io`             | `Seeker`       | `Seek()`                                                                                            |  
| `io`             | `StringWriter` | `WriteString()`                                                                                     |  
| `io`             | `Writer`       | `Write()`                                                                                           |  
| `io`             | `WriterAt`     | `WriteAt()`                                                                                         |  
| `io`             | `WriterTo`     | `WriteTo()`                                                                                         |
  
See also:  
- [File system-focused](#file-system-focused)  
- [Driver-focused](#driver-focused)  
- [Process control-focused](#process-control-focused)  
- [Protocol-focused](#protocol-focused)  
- [Implementation-focused](#implementation-focused)  
  
  
## Representation-focused  
There are uses of interface are used to allow a developer to **represent** a concrete or conceptual and often complex entity:  
  
| Package                 | Interface       | Method signature(s) and/or Interface embed(s)                                               |  
| ----------------------- | --------------- | ------------------------------------------------------------------------------------------- |  
| `debug/dwarf`           | `Type`          | `Common()`, `String()`, `Size()`                                                            |  
| `encoding/binary`       | `ByteOrder`     | `Uint16()`, `Uint32()`, `Uint64()`, `PutUint16()`, `PutUint32()`, `PutUint64()`, `String()` |  
| `image`                 | `Image`         | `ColorModel()`, `Bounds()`, `At()`                                                          |  
| `image`                 | `RBG64Image`    | `Image`, `RGBA64At()`                                                                       |  
| `image`                 | `PalattedImage` | `ColorIndexAt()`                                                                            |  
| `image/color`           | `Color`         | `RGBA()`                                                                                    |  
| `log/slog`              | `LogValuer`     | `LogValue()`                                                                                |  
| `net`                   | `Addr`          | `Network()`,`String()`                                                                      |  
| `*/arch/arm/armasm`     | `Arg`           | `IsArg()`,`String()`                                                                        |  
| `*/arch/ppc64/ppc64asm` | `Arg`           | `IsArg()`,`String()`                                                                        |  
| `*/net/route`           | `Message`      | `Sys()`                                                                                                                   |  
| `*/net/route`           | `Sys`          | `SysType()`                                                                                                               |  
  
See also:  
- [Data-focused](#data-focused)  
- [File system-focused](#file-system-focused)  
  
## Subtype-focused  
Interfaces with a focus on handling *"**subtypes**"* of an entity as one type. This reduces coupling and enables polymorphism of the nature found in classic object oriented languages.  
  
| Package                 | Interface             | Method signature(s) and/or Interface embed(s)                                                                             |
| ----------------------- | --------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| `cmd/pack`              | `FileLike`            | `Read()`, `Close()`, `Name()`, `Stat()`                                                                                   |
| `crypto/cipher`         | `Block`               | `BlockSize()`, `Encrypt()`, `Decrypt()`                                                                                   |
| `crypto/cipher`         | `AEAD`                | `NonceSize()`, `Overhead()`, `Seal()`, `Open()`                                                                           |
| `database/sql`          | `Scanner`             | `Scan()`                                                                                                                  |
| `debug/dwarf`           | `Type`                | `Common()`, `String()`, `Size()`                                                                                          |
| `encoding/binary`       | `ByteOrder`           | `Uint16()`, `Uint32()`, `Uint64()`, `PutUint16()`, `PutUint32()`, `PutUint64()`, `String()`                               |
| `expvar`                | `Var`                 | `String()`                                                                                                                |
| `fmt`                   | `GoStringer`          | `GoString()`                                                                                                              |
| `fmt`                   | `Scanner`             | `Scan()`                                                                                                                  |
| `go/ast`                | `Node`                | `Pos()`, `End()`                                                                                                          |
| `go/types`              | `Type`                | `Underlying()`, `String()`                                                                                                |
| `hash`                  | `Hash`                | `io.Writer`, `Blocksize()`, `Reset()`, `Size()`, `Sum()`                                                                  |
| `hash`                  | `Hash32`              | `Sum32()`                                                                                                                 |
| `hash`                  | `Hash64`              | `Sum64()`                                                                                                                 |
| `image`                 | `Image`               | `ColorModel()`, `Bounds()`, `At()`                                                                                        |
| `image`                 | `RBG64Image`          | `Image`, `RGBA64At()`                                                                                                     |
| `image/draw`            | `Image`               | `image.Image`,`Set()`                                                                                                     |
| `image/draw`            | `RGBA64Image`         | `image.RGBA64Image`,`Set()`,`SetRGBA64()`                                                                                 |
| `log/slog`              | `Handler`             | `WithAttrs()`, `WithGroup()`, `Enabled()`, `Handle()`                                                                     |
| `net`                   | `Addr`                | `Network()`,`String()`                                                                                                    |
| `net`                   | `Conn`                | `RemoteAddr()`, `SetDeadline()`, `SetReadDeadline()`, `SetWriteDeadline()`, `Read()`, `Write()`, `Close()`, `LocalAddr()` |
| `net`                   | `Error`               | `error`, `Timeout()`, `Temporary()`                                                                                       |
| `net`                   | `Listener`            | `Accept()`, `Addr()`, `Close()`                                                                                           |
| `net`                   | `PacketConn`          | `Close()`, `LocalAddr()`, `ReadFrom()`, `SetDeadline()`, `SetReadDeadline()`, `SetWriteDeadline()`, `WriteTo()`           |
| `net/http`              | `RoundTripper`        | `RoundTrip()`                                                                                                             |
| `net/rpc`               | `ClientCodec`         | `WriteRequest()`, `ReadResponseHeader()`, `ReadResponseBody()`, `Close()`                                                 |
| `net/rpc`               | `ServerCodec`         | `WriteResponse()`, `ReadRequestHeader()`, `ReadRequestBody()`, `Close()`                                                  |
| `runtime`               | `Error`               | `error`, `RuntimeError()`                                                                                                 |
| `syscall`               | `Conn`                | `SyscallConn()`                                                                                                           |
| `syscall`               | `RawConn`             | `Control()`, `Read()`, `Write()`                                                                                          |
| `*/arch/arm/armasm`     | `Arg`                 | `IsArg()`,`String()`                                                                                                      |
| `*/arch/ppc64/ppc64asm` | `Arg`                 | `IsArg()`,`String()`                                                                                                      |
| `*/mod/modfile`         | `Expr`                | `Comment()`, `Span()`                                                                                                     |
| `*/tools/go/analysis`   | `Fact`                | `AFact()`                                                                                                                 |
| `*/tools/go/analysis`   | `Range`               | `Pos()`, `End()`                                                                                                          |
| `*/net/lif`             | `Addr`                | `Family()`                                                                                                                |
| `*/net/route`           | `Addr`                | `Family()`                                                                                                                |
| `*/net/route`           | `Message`             | `Sys()`                                                                                                                   |
| `*/net/route`           | `Sys`                 | `SysType()`                                                                                                               |
| `*/text/transform`      | `Transformer`         | `Transform()`, `Reset()`                                                                                                  |
| `*/text/transform`      | `SpanningTransformer` | `Transformer`, `Span()`                                                                                                   |
  
See also:  
- [File system-focused](#file-system-focused)  
- [Reading and/or writing-focused](#reading-andor-writing-focused)


## Acknowledgements
1. Thanks to [Michael Potter][michael-potter] of [Tapp Solutions](http://www.tappsolutions.com/) for [his suggestions to improve][michael-potter-suggestions] this document.




[go1.21.0]: https://github.com/golang/go/tree/go1.21.0/src
[michael-potter]: https://www.linkedin.com/in/michaelpotter/
[michael-potter-suggestions]: https://www.linkedin.com/feed/update/urn:li:activity:7110764151748128768?commentUrn=urn%3Ali%3Acomment%3A%28activity%3A7110764151748128768%2C7110962540720922625%29&dashCommentUrn=urn%3Ali%3Afsd_comment%3A%287110962540720922625%2Curn%3Ali%3Aactivity%3A7110764151748128768%29
