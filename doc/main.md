- [Install](#install)
- [Shell exports](#shell-exports)
- [Directory explanations](#directory-explanations)
- [Automatic Imports](#automatic-imports)
- [Private repo access](#private-repo-access)
- [Guard (automatic go run)](#guard-automatic-go-run)
- [Godo](#godo)
- [Spurious](#spurious)
- [AWS SDK with Go (inc. some old possibly broken examples)](#aws-sdk-with-go)
- [Build and Compilation](#build-and-compilation)
- [Dependency information with `go list`](#dependency-information-with-go-list)
- [Dependencies with godeps](#dependencies-with-godeps)
- [Dependencies with gb](#dependencies-with-gb)
- [Dependencies with glide](#dependencies-with-glide)
- [Documentation](#documentation)
- [Testing](#testing)
- [Logging](#logging)
- [Bits, Bytes, Runes](#bits-bytes-runes)
- [Code Examples](#code-examples)
  - [Init](#init)
  - [New vs Make](#new-vs-make)
  - [Custom Types](#custom-types)
  - [Function Types](#function-types)
  - [Structure: Var vs Type](#struct-var-vs-type)
  - [Reference vs Value](#reference-vs-value)
  - [See all methods of a &lt;Type&gt;](#see-all-methods-of-a-type) 
  - [Set time](#set-time)
  - [Convert Struct into JSON](#convert-struct-into-json)
  - [Extract only JSON you need](#extract-only-json-you-need)
  - [Nested JSON handling](#nested-json-handling)
  - [Pretty Printing JSON String](#pretty-printing-json-string)
  - [Nested YAML handling](#nested-yaml-handling)
  - [Unknown YAML Structure](#unknown-yaml-structure)
  - [Sorting Structs](#sorting-structs)
  - [Read User Input](#read-users-input)
  - [Web Server](#web-server)
  - [Middleware](#middleware)
  - [Sessions](#sessions)
  - [HTTP Requests with Timeouts](#http-requests-with-timeouts)
  - [S3 GetObject](#s3-getobject)
  - [Compile time variables](#compile-time-variables)
  - [TLS HTTP Request](#tls-http-request)
  - [Custom HTTP Request](#custom-http-request)
  - [HTTP GET Web Page](#http-get-web-page)
  - [Pointers](#pointers)
  - [Array Pointer](#array-pointer)
  - [Type Assertion](#type-assertion)
  - [Line Count](#line-count)
  - [Measuring time](#measuring-time)
  - [Reading a file in chunks](#reading-a-file-in-chunks)
  - [Time and Channels](#time-and-channels)
  - [Quit a Channel](#quit-a-channel)
  - [Starting and Stopping things with Channels](#starting-and-stopping-things-with-channels)
  - [Channel Pipelines](#channel-pipelines)
  - [Templating](#templating)
  - [Error handling with context](#error-handling-with-context)
  - [Socket programming with TCP server](#socket-programming-with-tcp-server)
  - [Comparing maps](#comparing-maps)
  - [Embedded Structs](#embedded-structs)
  - [Zip File Contents](#zip-file-contents)
  - [OAuth](https://gist.github.com/Integralist/0e277a517fee68153f93)
  - [RPC](#rpc)
  - [Enumerator IOTA](#enumerator-iota)
  - [FizzBuzz](#fizzbuzz)
  - [Execute Shell Command](#execute-shell-command)
  - [New Instance Idiom](#new-instance-idiom)
  - [Mutating Values](#mutating-values)
  - [Draining Connections](#draining-connections)

## Install

```bash
brew install go
```

## Shell exports

- `export GOPATH=~/path/to/your/golang/projects`
- `export PATH=~/path/to/your/golang/projects/bin:$PATH`

> Note: the latter item allows you to locally build and execute Go based binaries

## Directory explanations

By default you store all your Golang projects within a single directory. This will be fixed in a future Go release as the developers recognise it can be problematic sometimes.

So within the `$GOPATH` directory workspace there should be three directories:

- `src`: holds source code
- `pkg`: holds compiled bits
- `bin`: holds executables

> Note: I very rarely even look at the `pkg` or `bin` directories

But for now, make sure you have any new Go project you work on placed inside the following directory structure...

```
└── src
    ├── github.com
    │   ├── <your_username>
    │   │   └── <your_repo_name>
```

## Automatic Imports

`go get golang.org/x/tools/cmd/goimports`

Now either run `goimports` from the shell OR use vim-go plugin with `:GoImports` for the buffer you're working with

## Private repo access

`go get` uses https; so instead force it to use ssh:

```bash
git config --global url."git@github.com:".insteadOf "https://github.com/"
```

> Note you can restrict it to a specific organisation as well:  
> `git config --global url."git@github.com:foo/".insteadOf "https://github.com/foo/"`

So when you want a private repository: `git@github.com:foo/private.git`

You can run:

```bash
go get github.com/foo/private
```

## Guard (automatic `go run`)

**UPDATE**: this isn't good practice. Instead use [Godo](https://github.com/go-godo/godo) (see below for example)

~~Follow this guide (https://gist.github.com/Integralist/b675a263897680e02fbd) for using Guard to get real-time notifications for when changes occur in your Go programming files, and automatically trigger `go run`.~~

## Godo

Example taken from my own project [go-requester](https://github.com/Integralist/Go-Requester)

```go
package main

import (
	"fmt"
	"os"

	do "gopkg.in/godo.v2"
)

func tasks(p *do.Project) {
	if pwd, err := os.Getwd(); err == nil {
		do.Env = fmt.Sprintf("GOPATH=%s/vendor::$GOPATH", pwd)
	}

	p.Task("server", nil, func(c *do.Context) {
		c.Start("main.go ./config/page.yaml", do.M{"$in": "./"})
	}).Src("**/*.go")
}

func main() {
	do.Godo(tasks)
}
```

## Spurious

If you need [Spurious](https://github.com/spurious-io/spurious) set-up then update the `aws.config` accordingly:

```Go
_dyn := dynamodb.New(&aws.Config{
    Region:     "eu-west-1",
    DisableSSL: true,
    Endpoint:   "dynamodb.spurious.localhost:32770", // change port number to appropriate value
})

_s3 := s3.New(&aws.Config{
    Region:           "eu-west-1",
    Endpoint:         "s3.spurious.localhost:32769", // change port number to appropriate value
    DisableSSL:       true,
    S3ForcePathStyle: true,
})
```

> Note: remember to set the AWS environment variables in your shell so Dynamo can pick them up (all other spurious services are fine without them)

```bash
export AWS_ACCESS_KEY_ID=development_access; export AWS_SECRET_ACCESS_KEY=development_secret; go run application.go
```

To populate your Spurious set-up you can use Ruby like so: https://gist.github.com/Integralist/58b25f860773d8d2dd3f

## AWS SDK with Go

### STS Assume Role

Usage:

```
<binary_name> <aws_account_id> <aws_role>
```

Code:

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

var (
	accessKey = os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	region    = os.Getenv("AWS_REGION")
)

func main() {
	args := os.Args[1:]
	account := args[0]
	role := args[1]

	sess := session.New(
		&aws.Config{
			Region: aws.String(region),
		},
	)

	svc := sts.New(sess)

	output, err := svc.AssumeRole(&sts.AssumeRoleInput{
		RoleArn:         aws.String(fmt.Sprintf("arn:aws:iam::%s:role/%s", account, role)),
		RoleSessionName: aws.String("temp"),
	})
	if err != nil {
		log.Fatalf("Unable to assume role: %v", err.Error())
	}

	os.Setenv("AWS_ACCESS_KEY_ID", aws.StringValue(output.Credentials.AccessKeyId))
	os.Setenv("AWS_SECRET_ACCESS_KEY", aws.StringValue(output.Credentials.SecretAccessKey))
	os.Setenv("AWS_SESSION_TOKEN", aws.StringValue(output.Credentials.SessionToken))

	fmt.Printf("AWS_ACCESS_KEY_ID: %s\n", os.Getenv("AWS_ACCESS_KEY_ID"))
	fmt.Printf("AWS_SECRET_ACCESS_KEY: %s\n", os.Getenv("AWS_SECRET_ACCESS_KEY"))
	fmt.Printf("AWS_SESSION_TOKEN: %s\n", os.Getenv("AWS_SESSION_TOKEN"))
}
```

### Create SQS queue

Usage:

```
go run local/create.go -queue "producer"
```

Code:

```go
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type network []struct {
	Host     string
	HostPort string
}

type spurious struct {
	Sqs    network `json:"spurious-sqs"`
	S3     network `json:"spurious-s3"`
	Dynamo network `json:"spurious-dynamo"`
}

var (
	svc          *sqs.SQS
	queueName    string
	regionName   string
	endpointName string
	cmdOut       []byte
	err          error
	spur         spurious
)

var region = flag.String("region", "eu-west-1", "Name of region to create the resource within")
var queue = flag.String("queue", "producer", "Name of queue to be created")
var endpoint = flag.String("endpoint", "", "Spurious endpoint")

func init() {
	flag.Parse()

	queueName = *queue
	regionName = *region
	endpointName = *endpoint

	if endpointName == "" {
		cmdName := "spurious"
		cmdArgs := []string{"ports", "--json"}
		if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
			fmt.Fprintln(os.Stderr, "There was an error running 'spurious ports --json' command: ", err)
			os.Exit(1)
		}

		json.Unmarshal(cmdOut, &spur)
		endpointName = spur.Sqs[0].Host + ":" + spur.Sqs[0].HostPort
	}

	svc = sqs.New(
		session.New(),
		&aws.Config{
			Region:     aws.String(regionName),
			DisableSSL: aws.Bool(true),
			Endpoint:   aws.String(endpointName),
		})
}

func main() {
	params := &sqs.CreateQueueInput{
		QueueName: aws.String(queueName),
	}

	resp, err := svc.CreateQueue(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Get error details
			log.Println("Error:", awsErr.Code(), awsErr.Message())

			// Prints out full error message, including original error if there was one.
			log.Println("Error:", awsErr.Error())

			// Get original error
			if origErr := awsErr.OrigErr(); origErr != nil {
				// operate on original error.
			}
		} else {
			fmt.Println(err.Error())
		}
	}

	fmt.Println(resp)
}
```

### Possibly Broken Examples

**UPDATE** the following code examples are now old and probably don't work any more

In the below code we use `go` blocks for parallelising "copy" requests to S3, which is thread-safe because we're not mutating any values. But we can't quite get away with that inside the `getS3Locations` function as we need to mutate a slice (and that's not thread-safe) so we then use an interesting pattern where by we use channels to synchronise the data after the parallelisation.

> Note: DynamoDB specifically is confusing.  
> Also, for printing Structs use: `fmt.Printf("%+v", myStruct)` (ensures the keys are included)

```go
package main

import (
	"fmt"
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/dynamodb"
	"github.com/awslabs/aws-sdk-go/service/s3"
	"os"
	"strings"
	"sync"
)

func sequencerTableRecords(sequencer string) *dynamodb.ScanOutput {
	svc := dynamodb.New(&aws.Config{
		Region: "eu-west-1",
		DisableSSL: true,
		Endpoint:   "dynamodb.spurious.localhost:32791",
	})

	params := &dynamodb.ScanInput{
		TableName: aws.String(sequencer),
	}

	resp, err := svc.Scan(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	return resp
}

func getComponentVersions(records *dynamodb.ScanOutput) map[string]string {
	components := make(map[string]string)

	for _, items := range records.Items {
		item := *items
		components[*item["key"].S] = *item["value"].N
	}

	return components
}

func getS3Locations(components map[string]string, s3Path string, lookup string) map[string]string {
	svc := dynamodb.New(&aws.Config{
		Region: "eu-west-1",
		DisableSSL: true,
		Endpoint:   "dynamodb.spurious.localhost:32791",
	})

	collectedLocations := []*dynamodb.QueryOutput{}

	c := make(chan *dynamodb.QueryOutput, len(components))
	done := make(chan int, len(components))
	locations := make(map[string]string)

	// Parallelise retrieval of data from DynamoDB
	for componentKey, componentVersion := range components {
		go func(componentKey, componentVersion string) {
			params := &dynamodb.QueryInput{
				TableName:      aws.String(lookup),
				ConsistentRead: aws.Boolean(true),
				Select:         aws.String("SPECIFIC_ATTRIBUTES"),
				AttributesToGet: []*string{
					aws.String("component_key"),
					aws.String("location"),
				},
				KeyConditions: &map[string]*dynamodb.Condition{
					"component_key": &dynamodb.Condition{
						ComparisonOperator: aws.String("EQ"),
						AttributeValueList: []*dynamodb.AttributeValue{
							&dynamodb.AttributeValue{
								S: aws.String(componentKey),
							},
						},
					},
					"batch_version": &dynamodb.Condition{
						ComparisonOperator: aws.String("EQ"),
						AttributeValueList: []*dynamodb.AttributeValue{
							&dynamodb.AttributeValue{
								N: aws.String(componentVersion),
							},
						},
					},
				},
			}

			resp, err := svc.Query(params)

			if awserr := aws.Error(err); awserr != nil {
				// A service error occurred.
				fmt.Println("Error:", awserr.Code, awserr.Message)
			} else if err != nil {
				// A non-service error occurred.
				panic(err)
			} else {
				c <- resp
				done <- 1
			}
		}(componentKey, componentVersion)
	}

	// Wait until all data is successfully collated from DynamoDB
	for i := len(components); i > 0; {
		select {
		case item := <-c:
			collectedLocations = append(collectedLocations, item)
		case <-done:
			i--
		}
	}

	for _, items := range collectedLocations {
		item := *items
		ref := *item.Items[0]
		componentLocation := s3Path + *ref["location"].S
		componentKey := extractComponentFromKey(*ref["component_key"].S)

		locations[componentKey] = componentLocation
	}

	return locations
}

func extractComponentFromKey(componentKey string) string {
	return strings.Split(componentKey, "/")[0]
}

func copyS3DataToNewLocation(event string, s3Bucket string, s3Locations map[string]string) {
	svc := s3.New(&aws.Config{
		Region: "eu-west-1",
		Endpoint:         "s3.spurious.localhost:32790",
		DisableSSL:       true,
		S3ForcePathStyle: true,
	})

	var wg sync.WaitGroup

	for component, location := range s3Locations {
		destination := "archive/" + event + "/" + component

		wg.Add(1)

		go func(location, destination string) {
			defer wg.Done()

			// fmt.Println(s3Bucket)
			// fmt.Println(s3Bucket + "/" + location)
			// fmt.Println(destination)

			params := &s3.CopyObjectInput{
				Bucket:     aws.String(s3Bucket),
				CopySource: aws.String(s3Bucket + "/" + location),
				Key:        aws.String(destination),
			}

			_, err := svc.CopyObject(params)

			if awserr := aws.Error(err); awserr != nil {
				// A service error occurred.
				fmt.Println("Error:", awserr.Code, awserr.Message)
			} else if err != nil {
				// A non-service error occurred.
				panic(err)
			}
		}(location, destination)
	}

	wg.Wait()
}

func main() {
	event := os.Args[1]
	s3Bucket := os.Args[2]
	s3Path := os.Args[3]
	sequencer := os.Args[4]
	lookup := os.Args[5]

	sequence_records := sequencerTableRecords(sequencer)
	components := getComponentVersions(sequence_records)
	s3Locations := getS3Locations(components, s3Path, lookup)

	copyS3DataToNewLocation(event, s3Bucket, s3Locations)
}
```

In above example there are API issues with DynamoDB - after about 6 requests a second the API errors. If you flatten out the requests so they are no longer running highly concurrently, then the speed of it slows down so badly that AWS Lambda (which is running the binary) times out. Meaning we need to do things differently... i.e. we need to request all S3 objects instead and partition/filter the unique values from that instead:

> Note: S3 objects are listed alphabetically

```go
func getS3ObjectSubset(bucket, source, marker string) *s3.ListObjectsOutput {
	svc := s3.New(&aws.Config{
		Region: "eu-west-1",
	})

	params := &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
		Prefix: aws.String(source),
		Marker: aws.String(marker),
	}

	resp, err := svc.ListObjects(params)

	if awserr := aws.Error(err); awserr != nil {
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		panic(err)
	}

	return resp
}

func main() {
	bucket := os.Args[1] // some-bucket
	source := os.Args[2] // some/object/path/to/prefix
	marker := ""         // means to start off from the very first object (overwritten)

	var resp *s3.ListObjectsOutput

	processing := true

	collectedObjects := []*s3.ListObjectsOutput{}

	for processing {
		resp = getS3ObjectSubset(bucket, source, marker)
		collectedObjects = append(collectedObjects, resp)
		marker = *resp.Contents[len(resp.Contents)-1].Key

		if *resp.IsTruncated == false {
			processing = false
		}
	}

	for _, s3SubSet := range collectedObjects {
		for _, items := range s3SubSet.Contents {
			fmt.Println(*items.Key)
		}
	}
}
```

## Build and Compilation

### 1.5+

```bash
GOOS=darwin GOARCH=386 go build foo.go
```

Here is a quick reference:

```
$GOOS     $GOARCH
darwin    386      -- 32 bit MacOSX
darwin    amd64    -- 64 bit MacOSX
freebsd   386
freebsd   amd64
linux     386      -- 32 bit Linux
linux     amd64    -- 64 bit Linux
linux     arm      -- RISC Linux
netbsd    386
netbsd    amd64
openbsd   386
openbsd   amd64
plan9     386
windows   386      -- 32 bit Windows
windows   amd64    -- 64 bit Windows
```

### Gox

One time only commands:

- `go get github.com/mitchellh/gox`
- `gox -build-toolchain` (only necessary for 1.4.x and lower)

Compilation (example is for AWS Lambda usage where only a single binary is needed):

- `gox -osarch="linux/amd64" -osarch="darwin/amd64" -osarch="windows/amd64" -output="foobar.{{.OS}}"`

This will generate three files:

1. `foobar.darwin`
2. `foobar.linux`
3. `foobar.windows.exe`

### Other information

Use the `-a` flag when running `go build`.

In short, if you dont' use `go build -a -v .` then Go won't know if any packages are missing (you can find the gory details [here](https://medium.com/@felixge/why-you-should-use-go-build-a-or-gb-c469157d5c1b#.jf5orcwrj))

## Dependency information with `go list`

To see a list of dependencies for a given Go package you can utilise the `go list` command:

```bash
go list -json strconv 
```

Which returns:

```json
{
	"Dir": "/usr/local/Cellar/go/1.5.2/libexec/src/strconv",
	"ImportPath": "strconv",
	"Name": "strconv",
	"Doc": "Package strconv implements conversions to and from string representations of basic data types.",
	"Target": "/usr/local/Cellar/go/1.5.2/libexec/pkg/darwin_amd64/strconv.a",
	"Goroot": true,
	"Standard": true,
	"Root": "/usr/local/Cellar/go/1.5.2/libexec",
	"GoFiles": [
		"atob.go",
		"atof.go",
		"atoi.go",
		"decimal.go",
		"doc.go",
		"extfloat.go",
		"ftoa.go",
		"isprint.go",
		"itoa.go",
		"quote.go"
	],
	"IgnoredGoFiles": [
		"makeisprint.go"
	],
	"Imports": [
		"errors",
		"math",
		"unicode/utf8"
	],
	"Deps": [
		"errors",
		"math",
		"runtime",
		"unicode/utf8",
		"unsafe"
	],
	"TestGoFiles": [
		"internal_test.go"
	],
	"XTestGoFiles": [
		"atob_test.go",
		"atof_test.go",
		"atoi_test.go",
		"decimal_test.go",
		"example_test.go",
		"fp_test.go",
		"ftoa_test.go",
		"itoa_test.go",
		"quote_test.go",
		"strconv_test.go"
	],
	"XTestImports": [
		"bufio",
		"bytes",
		"errors",
		"fmt",
		"log",
		"math",
		"math/rand",
		"os",
		"reflect",
		"runtime",
		"strconv",
		"strings",
		"testing",
		"time",
		"unicode"
	]
}
```

If you don't specify the `-json` flag then the default behaviour is to filter out the `ImportPath` field from the above JSON output. For example:

```bash
go list strconv
```

Will return just the import path `strconv`.

> Documentation: `go help list | less`

You can also utilise Go's templating functionality on the returned JSON object by adding the `-f` flag:

```bash
go list -f '{{join .Deps " "}}' strconv
```

Which filters out the `Deps` field, joins up all items it contains using whitespace and subsequently returns:

```
errors math runtime unicode/utf8 unsafe
```

You can do more complex things such as:

```bash
go list -f '{{.ImportPath}} -> {{join .Imports " "}}' compress/...
```

Which will return something like:

```
compress/bzip2 -> bufio io sort
compress/flate -> bufio fmt io math sort strconv
compress/gzip -> bufio compress/flate errors fmt hash hash/crc32 io time
compress/lzw -> bufio errors fmt io
compress/zlib -> bufio compress/flate errors fmt hash hash/adler32 io
```

## Dependencies with godeps

When running `go get <dependency>` locally, Go will stick the dependency in the folder defined by your `$GOPATH` variable. So when you build your code into a binary using `go build <script>` it'll bake the dependencies into the binary (i.e. the binary is statically linked).

But if someone pulls down your repo and tries to do a build they'll need to have a network connection to pull down the dependencies, as their `$GOPATH` might not have those dependencies yet (unless the user manually executes `go get` for each dependency required). Also the dependencies they subsequently pull down could be a more recent (and untested version) of each dependency.

So to make this situation better we can use http://godoc.org/github.com/tools/godep (https://github.com/tools/godep) which sticks all your dependencies within a `Godeps` folder inside your project directory. You can then use `godep save -r ./...` to automatically update all your references to point to that local folder. 

> Note: you might need to remove the `Godeps` folder and run `go get` if you get strange conflicts. The `./...` means to target all `.go` files

This way users who clone your repo don't need an internet connection to pull the dependencies, as they already have them. But also they'll have the correct versions of the dependencies. This acts like a `Gemfile.lock` as you would typically find in the Ruby world.

```bash
find . -name '*.go' -exec \
sed -i '' 's/github\.com\/bbc\/mozart\-config\-api\/src\/Godeps\/_workspace\/src\///' {} \;
```

## Dependencies with gb

```bash
go get -u github.com/constabulary/gb/...
gb vendor fetch <pkg>
gb build all
```

You'll need the following structure:

```bash
├── src
│   ├── foo
│   │   └── main.go
└── vendor
    ├── manifest
    └── src
```

The `vendor` directory is auto-generated by the `gb vendor fetch <pkg>` command.

## Dependencies with glide

This is now my preferred dependency management tool, as it works just like existing tools in other languages (e.g. Ruby's Bundler or Node's NPM) and so consistency is a plus.

It also provides the ability (like gb) to not commit dependencies but have specific versions vendored when running a simple command.

```bash
go get github.com/Masterminds/glide
export GO15VENDOREXPERIMENT=1       # or use 1.6
glide init                          # generates glide.yaml
glide install                       # installs from lock file (creates it if not found)
glide update                        # updates dependencies and updates lock file
glide list                          # shows vendored deps
go test $(glide novendor)           # test only your package (not vendored packages)
```

> Note: to add a new dependency `glide get <pkg_name>`

## Documentation

`Godoc` is the original implementation for viewing documentation. Previous to `Godoc` there was `go doc`, but that was removed and then added *back* with totally different functionality.

The syntax structure for `go doc` is as follows:

```
go doc <pkg>
go doc <sym>[.<method>]
go doc [<pkg>].<sym>[.<method>]
```

Here are some examples of using `go doc`:

```
go doc json # same as go doc encoding/json
go doc json.Number
go doc json.Number.Float64
```

Here is the same but using `godoc` (where the syntax structure is `godoc <pkg> <symbol>`):

```
godoc encoding/json # unlike "go doc json", "godoc json" doesn't work as it's not a fully qualified path
godoc encoding/json Number
godoc -src builtin make | less
```

> Unlike with `go doc`, `godoc` doesn't allow filtering by `<method>`  
> It only goes as far as `<pkg> <symbol>`  
> 
> You can use `<pkg> <symbol> <method>`  
> and the method will be included in the results  
> but you'll need to search for the method manually  
> `godoc -src net/http Request ParseForm | less`  
> here is a similar result using `go doc`    
> `go doc http.Request.ParseForm | less`

The purpose of `go doc` was to provide a simplistic cli documentation viewer, where as `Godoc` has many more features available.

The `go doc` command also works not only with Go's own library's but your own custom packages as well.

There are some differences in what is returned though between `godoc` and `go doc` (mainly the latter is more succinct/compact so you can find the functions/types you're after and then you can expand into those once you've found them; `godoc` is harder to sift through on the command line)...

### `godoc encoding/json Encoder`

```
type Encoder struct {
    // contains filtered or unexported fields
}
    An Encoder writes JSON objects to an output stream.

func NewEncoder(w io.Writer) *Encoder
    NewEncoder returns a new encoder that writes to w.

func (enc *Encoder) Encode(v interface{}) error
    Encode writes the JSON encoding of v to the stream, followed by a
    newline character.

    See the documentation for Marshal for details about the conversion of Go
    values to JSON.
```

### `go doc encoding/json Encoder`

```
type Encoder struct {
        // Has unexported fields.
}

    An Encoder writes JSON objects to an output stream.

func NewEncoder(w io.Writer) *Encoder
func (enc *Encoder) Encode(v interface{}) error
```

> Notice the functions don't have their documentation notes printed with `go doc`

One other thing `godoc` has over `go doc` is the ability to view the source code using the `-src` flag:

```
godoc -src builtin make | less
```

The `godoc` tool also has a full browser documentation suite available and allows you to generate HTML documentation for your project...

### Full Browser Documentation

Start a local documentation server and allow indexing (which takes a few minutes; you have to just keep trying the search until it's done)

```
godoc -http ':6060' -index
```

You can then open a new terminal pane and search via cli if you prefer (rather than open up a browser to http://localhost:6060/)

```
godoc -q tls | less
```

You can also have the playground available if you need it in the browser, but it does require an internet connection to compile:

```
godoc -http ':6060' -play
```

## Testing

> Note: see also [examples here](https://gist.github.com/Integralist/cf76668bc46d75058ab5f566d96ce74a)

Test files are placed in the same directory as the file/package being tested. The convention is to use the same file name but suffix it with `_test`. So `foo.go` would have another file next to it called `foo_test.go`.

Run the tests: `go test -v ./...`

You can also run a specific test like so: `go test -v command/config_test.go command/config.go`

> Note: remember that your test file should have the same package name as your code being tested. This means the test file will have access to all the public functions and variables of that package (and so subsequently it'll have access to the code being tested)

Here's our program:

```go
package main

import "fmt"

type FooIO interface {
	Read() string
}

type Foo struct{}

func (f *Foo) Read() string {
	return "We READ something from disk"
}

func Stuff(f FooIO) string {
	return f.Read()
}

func main() {
	foo := &Foo{}
	contents := Stuff(foo)
	fmt.Println(contents)
}
```

Here's our test:

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type FakeFoo struct{}

func (s *FakeFoo) Read() string {
	return "We 'pretend' to READ something from disk"
}

func TestSomething(t *testing.T) {
	assert := assert.New(t)

	foo := &FakeFoo{}
	contents := Stuff(foo)

	assert.Equal(contents, "We 'pretend' to READ something from disk")
}
```

### Test Examples

Faking HTTP and WebServers can be a bit tricky:

```go
package requester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/bbc/mozart-requester/src/aggregator"
	"github.com/julienschmidt/httprouter"
)

func TestSuccessResponse(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"head":[ "foo" ],"bodyInline":"bar","bodyLast":[ "baz" ]}`)
	}))
	defer upstream.Close()

	router := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Process(w, r, httprouter.Params{})
	}))
	defer router.Close()

	var config = []byte(fmt.Sprintf(`{
		"components":[
			{"id":"foo","endpoint":"%s","must_succeed":true},
			{"id":"bar","endpoint":"%s","must_succeed":true}
		]
	}`, upstream.URL, upstream.URL))

	req, err := http.NewRequest("POST", router.URL, bytes.NewBuffer(config))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var result aggregator.Result
	json.Unmarshal(body, &result)

	expectedStatus := "success"
	if result.Summary != expectedStatus {
		t.Errorf("The response:\n '%s'\ndidn't match the expectation:\n '%s'", result.Summary, expectedStatus)
	}

	expectedLength := 2
	if len(result.Components) != expectedLength {
		t.Errorf("The response:\n '%d'\ndidn't match the expectation:\n '%d'", len(result.Components), expectedLength)
	}
}

func TestFailureResponse(t *testing.T) {
	healthyUpstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"head":[ "foo" ],"bodyInline":"bar","bodyLast":[ "baz" ]}`)
	}))
	defer healthyUpstream.Close()

	failingUpstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "404 page not found")
	}))
	defer failingUpstream.Close()

	router := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Process(w, r, httprouter.Params{})
	}))
	defer router.Close()

	var config = []byte(fmt.Sprintf(`{
		"components":[
			{"id":"foo","endpoint":"%s","must_succeed":true},
			{"id":"bar","endpoint":"%s","must_succeed":true}
		]
	}`, healthyUpstream.URL, failingUpstream.URL))

	req, err := http.NewRequest("POST", router.URL, bytes.NewBuffer(config))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var result aggregator.Result
	json.Unmarshal(body, &result)

	expectedSummary := "failure"
	if result.Summary != expectedSummary {
		t.Errorf("The response:\n '%s'\ndidn't match the expectation:\n '%s'", result.Summary, expectedSummary)
	}

	expectedLength := 2
	if len(result.Components) != expectedLength {
		t.Errorf("The response length:\n '%d'\ndidn't match the expectation:\n '%d'", len(result.Components), expectedLength)
	}

	expectedStatus := []int{}
	for _, value := range result.Components {
		if value.Status == 404 {
			expectedStatus = append(expectedStatus, value.Status)
		}
	}
	if len(expectedStatus) < 1 || len(expectedStatus) > 1 {
		t.Errorf("The response length:\n '%d'\ndidn't match the expectation:\n '%d'", len(expectedStatus), 1)
	}
}

func TestSlowResponse(t *testing.T) {
	healthyUpstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"head":[ "foo" ],"bodyInline":"bar","bodyLast":[ "baz" ]}`)
	}))
	defer healthyUpstream.Close()

	slowUpstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeout, err := strconv.Atoi(os.Getenv("COMPONENT_TIMEOUT"))
		if err != nil {
			t.Errorf("COMPONENT_TIMEOUT: %s", err.Error())
		}
		time.Sleep(time.Duration(timeout) * time.Millisecond)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"head":[ "foo" ],"bodyInline":"bar","bodyLast":[ "baz" ]}`)
	}))
	defer slowUpstream.Close()

	router := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Process(w, r, httprouter.Params{})
	}))
	defer router.Close()

	var config = []byte(fmt.Sprintf(`{
		"components":[
			{"id":"foo","endpoint":"%s","must_succeed":true},
			{"id":"bar","endpoint":"%s","must_succeed":true}
		]
	}`, healthyUpstream.URL, slowUpstream.URL))

	req, err := http.NewRequest("POST", router.URL, bytes.NewBuffer(config))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var result aggregator.Result
	json.Unmarshal(body, &result)

	expectedStatus := 408
	for _, value := range result.Components {
		if value.ID == "bar" && value.Status != expectedStatus {
			t.Errorf("The response:\n '%d'\ndidn't match the expectation:\n '%d'", value.Status, expectedStatus)
		}
	}

	expectedSummary := "failure"
	if result.Summary != expectedSummary {
		t.Errorf("The response:\n '%s'\ndidn't match the expectation:\n '%s'", result.Summary, expectedSummary)
	}
}
```

I typically run my tests using Make, but it ultimately looks like this: 

```
pushd src && APP_ENV=test COMPONENT_TIMEOUT=100 go test -v $(glide novendor) && popd
```

Here's another example of a test needing to fake things:

```go
package retriever

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

const href = "http://bar.com/"
const url = "http://foo.com/"

var body string

func fakeNewDocument(url string) (*goquery.Document, error) {
	body = strings.Replace(body, "{}", href, 1)

	resp := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.0",
		ProtoMajor:    1,
		ProtoMinor:    0,
		Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
		ContentLength: int64(len(body)),
		Request:       &http.Request{},
	}

	return goquery.NewDocumentFromResponse(resp)
}

func TestRetrieveReturnValue(t *testing.T) {
	// {} interpolated with constant's value
	body = `
		<html>
			<body>
				<div class="productInfo">
					<a href="{}">Bar</a>
				</div>
			</body>
		<html>
	`
	coll, _ := Retrieve(url, fakeNewDocument)

	if response := coll[0]; response != href {
		t.Errorf("The response:\n '%s'\ndidn't match the expectation:\n '%s'", response, href)
	}
}

func TestRetrieveMissingAttributeReturnsEmptySlice(t *testing.T) {
	// href attribute is missing from anchor element
	body = `
		<html>
			<body>
				<div class="productInfo">
					<a>Bar</a>
				</div>
			</body>
		<html>
	`
	coll, _ := Retrieve(url, fakeNewDocument)

	if response := coll; len(response) > 0 {
		t.Errorf("The response:\n '%s'\ndidn't match the expectation:\n '%s'", response, "[http://bar.com/]")
	}
}
```

And...

```go
package scraper

import "testing"

func TestScrapeResults(t *testing.T) {
	getItem = func(url string) {
		defer wg.Done()

		ch <- Item{
			"FooTitle",
			"FooSize",
			"10.00",
			"FooDescription",
		}
	}

	urls := []string{
		"http://foo.com/",
		"http://bar.com/",
		"http://baz.com/",
	}

	result := Scrape(urls)
	first := result.Items[0]

	var suite = []struct {
		response string
		expected string
	}{
		{first.Title, "FooTitle"},
		{first.Size, "FooSize"},
		{first.UnitPrice, "10.00"},
		{first.Description, "FooDescription"},
		{result.Total, "30.00"},
	}

	for _, v := range suite {
		if v.response != v.expected {
			err(v.response, v.expected, t)
		}
	}
}

func err(response, expected string, t *testing.T) {
	t.Errorf("The response:\n '%s'\ndidn't match the expectation:\n '%s'", response, expected)
}
```

## Logging

Using the standard Logger:

```go
info := log.New(os.Stdout, "STUFF: ", log.Ldate|log.Ltime|log.Lshortfile)
info.Println("Starting up!!!")

f, e := os.Create("test.log")
if e != nil {
	log.Fatal("Failed to create log file")
}

logfile := log.New(f, "STUFF: ", log.Ldate|log.Ltime|log.Lshortfile)
logfile.Println("Starting up!!!")
```

Using Logrus:

```go
package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
)

func main() {
	// Standard stdout ASCII logging
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	// JSON style structured logging
	log.SetFormatter(&log.JSONFormatter{})
	f, e := os.Create("logs")
	if e != nil {
		log.Fatal("Failed to create log file")
	}
	log.SetOutput(f)
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
	/*
			{
				"animal": "walrus",
				"level": "info",
				"msg": "A group of walrus emerges from the ocean",
		    "size": 10,
				"time": "2015-12-22T13:58:46Z"
			}
	*/
}
```

## Bits, Bytes, Runes

https://pythonconquerstheuniverse.wordpress.com/2010/05/30/unicode-beginners-introduction-for-dummies-made-simple/

A Unicode "code point" (e.g. `0021` which is equal to `!`) is known in Go as a "Rune".

> Note: a Rune is actually a synonym for Go's `int32` type

A Unicode "code point" is made up of a single byte.

Computers think in 8-bit bytes (i.e. a single byte is 8 bits).

With 8 bits you can make 256 different bit combinations. But Unicode has way more than 256 characters (it holds code points/characters for every language in the world, so yes a *lot* more than 256).

- 1 bytes = 08 bits
- 2 bytes = 16 bits
- 3 bytes = 24 bits
- 4 bytes = 32 bits

We could represent a Unicode "code point" with a single Rune (`int32`) but not all code points require a full 32 bits and so you'd be wasting lots of space. For example, ASCII only requires 8 bits (or 1 byte) per character.

UTF-8 is a solution to this problem. It uses 8-bit encoding, but one of the bits will be a pointer to another location to continue the bit sequence so the program can identify the overall character being encoded. This allows all Unicode code points to be encoded in 1 to 4 bytes but without the need for all the storage required of a 32 bit set-up.

So as you can now see, UTF-8 is able to use multiple bytes (up to 4) to represent a single Unicode code point. 

For example, `[E4 B8 96]` are three separate bytes that make up a single Chinese character.

A string is made up of individual bytes, but not every character in the string is necessarily mapped to a single byte (also ASCII charters like `\n` and `\t` are considered a byte each)

> Note: a Rune can consist of multiple bytes (so it's not *exactly* identical to a Unicode "code point")

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	hello := "您好 world"

	fmt.Printf("hex digits: % x\n", hello) // hex digits: e6 82 a8 e5 a5 bd 20 77 6f 72 6c 64

	// e6 82 a8 e5 a5 bd == 您好
	// 20                == <white space character>
	// 77 6f 72 6c 64    == world

	r := []rune(hello)

	fmt.Printf("UFT-8 encoding of each rune: %x\n", r) // UFT-8 encoding of each rune: [60a8 597d 20 77 6f 72 6c 64]

	// 60a8 597d      == 您好
	// 20             == <white space character>
	// 77 6f 72 6c 64 == world

	fmt.Println(len(hello)) // 12

	// Looks like 'hello' stores 8 characters, but the 2 chinese characters represent more than 2 bytes each
	// Instead we'd need to count the Runes...

	fmt.Println(utf8.RuneCountInString(hello)) // 8
	
	// The DecodeRuneInString method also returns the number of bytes each Rune occupies...
	
	rune1, size := utf8.DecodeRuneInString("您")
	fmt.Printf("Rune: %v\nRune's Byte Size: %v\n", rune1, size) 
	
	// Rune: 24744
	// Rune's Byte Size: 3
	
	rune2, size := utf8.DecodeRuneInString("好")
	fmt.Printf("Rune: %v\nRune's Byte Size: %v\n", rune2, size) 
	
	// Rune: 22909 (type: int32)
	// Rune's Byte Size: 3
	
	// Type conversion from a integer to a string yields not a stringified number but the UTF-8 representation of that Rune...
	
	fmt.Println(string(rune1))  // 您
	fmt.Println(string(rune2))  // 好
	fmt.Println(string(r))      // 您好 world
	fmt.Println(string(65))     // A

	// Use 0x prefix to denote a UTF-8 encoding...
	
	fmt.Println(string(0x60a8)) // 您
	fmt.Println(string(0x597d)) // 好
}
```

## Code Examples

### Init

When you load a package in Go, only the public functions and variables are exposed for the caller to utilise. So if you need a package to execute some bootstrapping code at the point of it being _loaded_, then you'll need to stick it inside of an `init` function.

> Note: you can have multiple `init` functions inside a package 
> e.g. one per file within the package namespace

But be careful with race conditions! 

I've hit an issue where we had:

- `main.go`
  - `foo.go` (loaded by `main.go`)
    - `bar.go` (loaded by `foo.go`)

Each one of these packages had its own `init` function and ultimately the `bar.go`'s `init` function was being run first, followed by the `foo.go`'s `init` function and finally followed by the `main.go`'s `init` function.

The reason this was an issue was because `main.go` was loading some environment variables needed by `bar.go` but those variables weren't available by the time the `bar.go` was running (as that happened _before_ `main.go`'s `init` function had executed.

The solution was to rename all the `init` functions to `Init` and explicitly call them to bootstrap the package when needed (i.e. they didn't automatically bootstrap themselves and find themselves in a race condition)

### New vs Make

- `func new(Type) *Type`: allocate memory for custom-user type
- `func make(Type, size IntegerType) Type`: allocate memory for builtin types (Slice, Map, Chan)

```go
package main

import "fmt"

func main() {
	foo := make(map[string]string)
	fmt.Println(foo) // map[]
	foo["k1"] = "bar"
	fmt.Println(foo) // map[k1:bar]
	fmt.Println(foo["k1"]) // bar
	
	type bar [5]int
	b := new(bar)
	fmt.Println(b) // &[0 0 0 0 0]
	b[0] = 1
	fmt.Println(b) // &[1 0 0 0 0]
}
```

### Custom Types

```go
package main

import (
	"bytes"
	"fmt"
)

type path []byte // our custom Type

// method attached to our custom Type
func (p *path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(*p, []byte("/"))

	if i >= 0 {
		*p = (*p)[0:i]
	}
}

func main() {
	pathName := path("/usr/bin/tso") // Conversion from string to path.

	pathName.TruncateAtFinalSlash()

	fmt.Printf("%s\n", pathName)
}
```

Alternative example:

```go
package main

import "fmt"

type foo [5]int

func main() {
	f := new(foo)
	fmt.Println(f) // &[0 0 0 0 0]
	f[0] = 1
	fmt.Println(f) // &[1 0 0 0 0]
	f.Bar()
	fmt.Println(f) // &[1 2 0 0 0]

	// We can coerce custom types like we can with built-in types
	b := foo([5]int{9, 9, 9})
	fmt.Println(b) // [9 9 9 0 0]
	
	// Check the types
	fmt.Printf("%T\n", b)               // main.foo
	fmt.Printf("%T\n", [5]int{9, 9, 9}) // [5]int
}

func (f *foo) Bar() {
	f[1] = 2
}
```

### Function Types

```go
package main

import "fmt"

type Foo func(int, string)

func (f Foo) Bar(s string) {
	fmt.Printf("s: %s\n", s)
}

func FooIt(x int, y string) {
	fmt.Printf("x: %d - y: %s\n", x, y)
}

// We HAVE to define the incoming type of "fn"
// Which in this case is a Foo type
func TestIt(fn Foo) {
	fn(99, "problems")
}

// We could do this without defining a func type
// But as you can see, this is a bit ugly
// Plus if we need this function passed around a lot
// then it means a lot of duplicated effort 
// typing the signature over and over
func TestItManually(fn func(int, string)) {
	fn(100, "problems")
}

func main() {
	// Here we're just demonstrating passing around the FooIt function
	// It demonstrates first-class function support in Go
	// But also that we can ensure the function passed around has the expected signature
	TestIt(FooIt)
	TestItManually(FooIt)
	
	x := Foo(FooIt) // Convert our function into a Foo type
	x(0, "hai")     // Now we can execute it as we would FooIt itself
	
	FooIt(1, "bye")
	
	// Notice the types are different
	// FooIt is just a function with a signature (no known type associated with it)
	// Where as "x" is of known type "Foo"
	fmt.Printf("%T\n", FooIt) // func(int, string)
	fmt.Printf("%T\n", x)     // main.Foo
	
	// But we'll see that the function "x" 
	// which was converted into a Foo type
	// now has access to a Bar method
	// Although FooIt has a matching signature, it's not a Foo type
	// and so it doesn't have a Bar method available
	x.Bar("we have a Bar method")
	
	// We can't even execute:
	// FooIt.Bar("we don't have a Bar method")
	// Because the compiler will stop us
}
```

### Struct: Var vs Type

A variable of Struct type doesn't need to be instantiated like a type struct:

```go
package main

import "fmt"

var data struct {
	A string
	B string
}

type data2 struct {
	A string
	B string
}

func main() {
	data.A = "Hai"
	data.B = "Bai"
	
	fmt.Printf(
		"%#v, %+v, %+v", 
		data.A, 
		data.B, 
		data2{A: "abc", B: "def"}
	)
	// "Hai", Bai, {A:abc B:def}
}
```

### Reference vs Value

Map data structures are passed by reference, rather than a copied value

```go
package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println("main before, m = ", m)
	foo(m)
	fmt.Println("main after, m = ", m)
}

func foo(m map[string]int) {
	fmt.Println("foo before, m = ", m)
	m["hai"] = 123
	fmt.Println("foo after, m = ", m)
}
```

In fact, anything with `make` is a reference, as well as any explicit interface

### See all methods of a &lt;Type&gt;

```go
errType := reflect.TypeOf(err)
for i := 0; i < errType.NumMethod(); i++ {
  method := errType.Method(i)
  fmt.Println(method.Name)
}
```

### Set time

```go
now := time.Now()
fmt.Println(now)
expiration := now.Add(time.Hour * 24 * 30)
fmt.Println("Thirty days from now will be : ", expiration)
```

### Convert Struct into JSON

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	type Message struct {
		Sequence  int    `json:"sequence"`
		Title     string `json:"title"`
		Timestamp time.Time   `json:"timestamp"`
	}
	msg := Message{1, "Foobar", time.Now()}
	b, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
```

### Extract only JSON you need

https://medium.com/the-hoard/using-golang-and-json-for-kafka-consumption-with-high-throughput-4cae28e08f90#.7rcmae71b

Effectively the solution is:

```go
/*
imagine our variable 'bytes' contains some JSON with lots of fields
we only want the fields 'type' and 'id' 
we should get our provider of data to transform everything else inside a 'data' field

e.g.

{
  “id”: “numero uno”,
  “type”: “transaction”,
  // data is the JSON msg from above
  “data”: { 
    “id”: “numero uno”,
    “type”: “transaction”,
    // … a whole bunch of dynamic fields
    “amount”: “1000”,
    “currency”: “usd”,
    // … etc.
   }
}
*/

type Message struct {
  ID string `json:”id”`
  Type string `json:”type”`
  Data json.RawMessage `json:”data”`
}

var m Message
json.Unmarshal(bytes, &m)
es.Index(index, m.Type, m.ID, "", "", nil, m.Data, false)
```

This way we only decode the id and type, so we're being performant, and then we pass the original raw JSON onto our next service (e.g. `es` ElasticSearch) to do with what they please.

### Nested JSON handling

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Component struct {
	Components []struct {
		Id  string `json:"id"`
		Url string `json:"url"`
	} `json:"components"`
}

func main() {
	var c Component

	b := []byte(`{"components":[{"id":"google","url":"http://google.com/"},{"id":"integralist","url":"http://integralist.co.uk/"},{"id":"sloooow","url":"http://stevesouders.com/cuzillion/?c0=hj1hfff5_0_f&c1=hc1hfff2_0_f&t=1439190969678"}]}`)

	json.Unmarshal(b, &c)
	
	fmt.Printf("%+v", c.Components[0]) // {Id:google Url:http://google.com/}
}
```

### Pretty Printing JSON String

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.MarshalIndent(group, "", "    ")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
```

### Nested YAML handling

```go
package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type ComponentYaml struct {
	Id  string `yaml:"id"`
	Url string `yaml:"url"`
}

type ComponentsYamlList struct {
	Components []ComponentYaml `yaml:"components"`
}

func main() {
	var y ComponentsYamlList

	yaml.Unmarshal([]byte("components:\n  - id: google\n    url: http://google.com\n  - id: integralist\n    url: http://integralist.co.uk"), &y)

	fmt.Println(y)
}
```

### Unknown YAML Structure

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var yml = []byte(`
- key: foo
  value: bar
  secret: false
- key: beep
  value: boop
  secret: true
`)

type Data struct {
	Items []map[string]interface{}
}

func main() {
	y := []map[string]interface{}{}

	if err := yaml.Unmarshal(yml, &y); err == nil {
		fmt.Printf("%#v\n", y)
	} else {
		fmt.Println(err.Error())
	}

	myYaml := Data{Items: y}

	json.NewEncoder(os.Stdout).Encode(myYaml.Items)
}
```

### Sorting Structs

```go
package main

import (
	"fmt"
	"sort"
)

type vals []Value

type Value struct {
	Key string
	Value string
	Secure bool
}

// Satisfy the Sort interface
func (v vals) Len() int      { return len(v) }
func (v vals) Swap(i, j int) { v[i], v[j] = v[j], v[i] }
func (v vals) Less(i, j int) bool { 
	return v[i].Key < v[j].Key 
}

func main() {
	orig := vals{
		{"CK", "BV", false},
		{"DK", "AV", true},
		{"AK", "CV", false},
		{"BK", "DV", true},
	}
	
	fmt.Printf("%+v\n\n", orig)
	sort.Sort(orig)
	fmt.Printf("%+v\n\n", orig)
}
```

Here is a similar version that sorts by name and age:

```go
package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

type byName []person

func (p byName) Len() int {
	return len(p)
}
func (p byName) Less(i, j int) bool {
	return p[i].Name < p[j].Name
}
func (p byName) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type byAge []person

func (p byAge) Len() int {
	return len(p)
}
func (p byAge) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}
func (p byAge) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	kids := []person{
		{"Jill", 9},
		{"Jack", 10},
	}

	sort.Sort(byName(kids))
	fmt.Println(kids)

	sort.Sort(byAge(kids))
	fmt.Println(kids)
}
```

Which results in:

```
[{Jack 10} {Jill 9}]
[{Jill 9} {Jack 10}]
```

### Read Users Input

```go
reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter text: ")
text, _ := reader.ReadString('\n')
fmt.Println(text)
```

### Web Server

The Go web server design relies on a struct to map routes (URLs) to functions.

You can define your own struct (prefilled for example) and pass it into `ListenAndServe`. But typically `nil` is used, which means an empty struct is used by default.

At this point most people will use either `HandleFunc` or `Handle` to register their specified request path so it maps to a specific handler function (this is added to the default struct called `DefaultServeMux`):

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

There is a difference between `HandleFunc` and `Handle`. The latter takes a type that has a `ServeHTTP` method associated to it (we'll see an example in a moment of what this looks like). The former is an abstraction layer that allows an incompatible function (one that doesn't have a `ServeHTTP` method) to be used as a handler. 

The way `HandleFunc` works is that it wraps the provided function in a call to `HandlerFunc` (see below for example). In this example `HandlerFunc` is a type of `func`, and this type defines the expected function signature and return value(s). 

What it states is that a compatible function should have the following signature: `ResponseWriter, *Request`, and it also attaches the method `ServeHTTP` to the type `HandlerFunc`. 

Now we can understand that when `HandleFunc` is called and passed our arbitrary function, we call the `HandlerFunc` func type and pass it our function, subsequently *converting* the incoming function so it is now of the type `HandlerFunc` and will now have gained a `ServeHTTP` function which allows it to satisfy the `Handle` interface.

Finally, our `HandleFunc` - once finished adpating the incoming user function - will internally call the `Handle` function and pass it the adapted function, which now satisfies the interface required by `Handle`.

The actual implementation looks like the following (I've cobbled together all the separate pieces, it doesn't necessarily appear like this in the source code):

```go
// We define an interface that states
// if the object has a ServeHTTP method
// then it satisfies this interface
type Handler interface {
  ServeHTTP(ResponseWriter, *Request)
}

// The func type works a bit like an interface
// So if your own user-defined function has a matching signature
// then your function is considered a `HandlerFunc` and will acquire a `ServeHTTP` method
// See directly below for where ServeHTTP is attached to this func type
type HandlerFunc func(ResponseWriter, *Request)

// Once the provided function is converted to the HandlerFunc type
// it'll mean it has the `ServeHTTP` function available
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}

// This is the abstraction function our client code calls...
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    // Here is where we "adapt" the incoming function so it is given a ServeHTTP method
    // We do this by passing it to the HandlerFunc func type
    mux.Handle(pattern, HandlerFunc(handler))
}

// Finally, this is the function that's passed our "adapted/converted" handler function
// The 'handler' passed in now fulfills the 'Handler' interface that says it needs a 'ServeHTTP' method
func (mux *ServeMux) Handle(pattern string, handler Handler) {
    ...do all the things...
}
```

This allows the arbitrary function to be used for handling the requested URL.

Below is an example for using `handle` instead of `handleFunc`:

```go
package main

import (
	"fmt"
	"net/http"
)

type String string

func (s String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, s)
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, s.Greeting, s.Punct, s.Who)
}

func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	http.ListenAndServe("localhost:4000", nil)
}
```

Now visit `http://localhost:4000/string` and `http://localhost:4000/struct` to see the appropriate output

### Middleware

This code was modified from https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type data struct {
	Greeting string
	Punct    string
	Who      string
}

func (s data) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s.Greeting, s.Punct, s.Who)
}

type adapter func(http.Handler) http.Handler

func adapt(h http.Handler, adapters ...adapter) http.Handler {
	// Ideally you'd do this in reverse
	// to ensure the order of the middleware
	// matches their specified order
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

func notify(logger *log.Logger) adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Println("before")
			defer logger.Println("after")
			h.ServeHTTP(w, r)
		})
	}
}

func doSomething() adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("before")
			defer fmt.Println("after")
			h.ServeHTTP(w, r)
		})
	}
}

func main() {
	http.Handle("/hello", &data{"Hello", " ", "Gophers!"})

	logger := log.New(os.Stdout, "server: ", log.Lshortfile)

	http.Handle("/hello-with-middleware", adapt(
		&data{"Hello", " ", "Gophers!"},
		notify(logger), // runs second
		doSomething(), // runs first
	))

	http.ListenAndServe("localhost:4000", nil)
}
```

This code will run a web server with two valid endpoints:

1. `/hello`
2. `/hello-with-middleware`

The client sees the same output but the latter endpoint produces the following stdout output:

```
before
server: middleware.go:35: before
server: middleware.go:38: after
after
```

### Sessions

```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

const cookiePrefix = "integralist-example-cookie-"

func main() {
	http.HandleFunc("/", login)
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe("localhost:4000", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, `
<html>
  <body>
    <form method="POST">
      Username: <input type="text" name="username">
      <br />
      Password: <input type="password" name="password">
      <br />
      <input type="submit" value="Login">
  </body>
</html>
`)
	}

	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "admin" && password == "password" {
			http.SetCookie(w, &http.Cookie{
				Name:  cookiePrefix + "user",
				Value: username,
			})
			http.Redirect(w, r, "/admin", 302)
		} else {
			fmt.Fprintf(w, `
<html>
  <body>
		Login details were incorrect. Sorry, <a href="/">try again</a>
  </body>
</html>
`)
		}
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    cookiePrefix + "user",
		Value:   "",
		Expires: time.Now(),
	})

	http.Redirect(w, r, "/", 302)
}

func admin(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(cookiePrefix + "user")
	if err != nil {
		http.Redirect(w, r, "/", 401) // Unauthorized
		return
	}

	fmt.Fprintf(w, `
<html>
  <body>
	  Logged into admin area as: %s<br><br>
		<a href="/logout">Logout</a>
  </body>
</html>
`, cookie.Value)
}
```

### HTTP Requests with Timeouts

```go
// Wait for 1.5 release to be able to verify timeout error (bug in language)
// Use -race flag https://blog.golang.org/race-detector to detect race conditions

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"gopkg.in/yaml.v2"
)

type ComponentYaml struct {
	Id  string `yaml:"id"`
	Url string `yaml:"url"`
}

type Component struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

type ComponentsYamlList struct {
	Components []ComponentYaml `yaml:"components"`
}

type ComponentsList struct {
	Components []Component `json:"components"`
}

type ComponentResponse struct {
	Id     string
	Status int
	Body   string
}

type Result struct {
	Status     string
	Components []ComponentResponse
}

var overallStatus string = "success"

func getComponents() []byte {
	return []byte(`{"components":[{"id":"local","url":"http://localhost:8080/pugs"},{"id":"google","url":"http://google.com/"},{"id":"integralist","url":"http://integralist.co.uk/"},{"id":"sloooow","url":"http://stevesouders.com/cuzillion/?c0=hj1hfff30_5_f&t=1439194716962"}]}`)
}

func getComponent(wg *sync.WaitGroup, client *http.Client, i int, v Component, ch chan ComponentResponse) {
	defer wg.Done()

	resp, err := client.Get(v.Url)

	if err != nil {
		fmt.Printf("Problem getting the response: %s\n\n", err)

		ch <- ComponentResponse{
			v.Id, 500, err.Error(),
		}
	} else {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Problem reading the body for %s -> %s\n", v.Id, err)
		}

		ch <- ComponentResponse{
			v.Id, resp.StatusCode, string(contents),
		}
	}
}

func main() {
	var cr []ComponentResponse
	var c ComponentsList
	var y ComponentsYamlList

	ch := make(chan ComponentResponse)
	b := getComponents() // to be read from a file

	yaml.Unmarshal([]byte("components:\n  - id: google\n    url: http://google.com\n  - id: integralist\n    url: http://integralist.co.uk"), &y)
	json.Unmarshal(b, &c)

	timeout := time.Duration(1 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	var wg sync.WaitGroup
	for i, v := range c.Components {
		wg.Add(1)
		go getComponent(&wg, &client, i, v, ch)
		cr = append(cr, <-ch)
	}
	wg.Wait()

	j, err := json.Marshal(Result{overallStatus, cr})
	if err != nil {
		fmt.Printf("Problem converting to JSON: %s\n", err)
		return
	}

	fmt.Println(string(j))
	fmt.Println(y)
}
```

### S3 GetObject

```go
import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BBC-News/mozart-config-api/src/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

func HandleStatusRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func HandleGetRequest(w http.ResponseWriter, r *http.Request) {
	svc := s3.New(&aws.Config{
		Region:           aws.String("eu-west-1"),
		Endpoint:         aws.String("s3.spurious.localhost:32769"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	})

	params := &s3.GetObjectInput{
		Bucket: aws.String("int-mozart-config-api"),
		Key:    aws.String("/v1/int/news/foo.json"),
	}

	resp, err := svc.GetObject(params)

	// ABSTRACT INTO A FUNCTION IN A NAMESPACE!
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println("1. ", awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println("2. ", reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			fmt.Println("3. ", err.Error())
		}
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading content", err)
	}

	fmt.Println(data) // => []byte
	fmt.Println(string(data[:]))

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	logger.Metric("200 response")
}
```

### Compile time variables

```go
var (
    Version   string
    BuildTime string
)
```

Now build the project using: 

```bash
go build -ldflags "-X github.com/<user>/<project>/core.Version=1.0.0 -X github.com/<user>/<project>/core.BuildTime=2015-10-03T11:08:49+0200" main.go
```

### TLS HTTP Request

```go
package requester

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	certFile = flag.String("cert", "/etc/pki/tls/certs/client.crt", "A PEM eoncoded certificate file.")
	keyFile  = flag.String("key", "/etc/pki/tls/private/client.key", "A PEM encoded private key file.")
	caFile   = flag.String("CA", "/etc/ca/cloud-ca.pem", "A PEM eoncoded CA's certificate file.")
)

func SecureClient() *http.Client {
	// Load client cert
	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(*caFile)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caCertPool,
		InsecureSkipVerify: true,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	return client
}
```

And to use it...

```go
client := requester.SecureClient()

// GET
resp, err := client.Get(someEndpoint)

// POST
req, err := http.NewRequest("POST", someEndpoint, bytes.NewBuffer(jsonStr))
req.Header.Set("Content-Type", "application/json")
resp, err := client.Do(req)
```

### Custom HTTP Request

Go doesn't provide abstractions for all the various HTTP request types, so for things like `PUT` you have to implement it yourself. The following is an example that creates a secure (TLS/HTTPS) `PUT` abstraction...

```go
func SecurePut(url, contentType string, configFile io.Reader) (*http.Response, error) {
	client := &http.Client{Transport: configureTLS()}
	req, err := http.NewRequest("PUT", url, configFile)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", contentType)
	resp, err := client.Do(req)

	return resp, err
}

func configureTLS() *http.Transport {
	certFilePath := "path/to/cert"
	keyFilePath := "path/to/privateKey"
	caPath := "path/to/ca"

	// Load client cert
	cert, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
	if err != nil {
		msg := fmt.Sprintf("Error loading developer cert, path: \"%s\"", certFilePath)
		output.Error(msg)
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(caPath)
	if err != nil {
		msg := fmt.Sprintf("Error loading CA cert, path: \"%s\"", caPath)
		output.Error(msg)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caCertPool,
		InsecureSkipVerify: true,
	}
	tlsConfig.BuildNameToCertificate()

	return &http.Transport{TLSClientConfig: tlsConfig}
}
```

### HTTP GET Web Page

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("http://www.integralist.co.uk/")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(string(contents))
}
```

### Pointers

```go
package main

import "fmt"

// Point stores co-ordinates
type Point struct {
	x int
	y int
}

// If receiver (Point) isn't set to a pointer (*Point) 
// then the struct's field value won't be updated outside the method
func (p *Point) scaleBy(factor int) {
	fmt.Printf("scaleBy (before modification): %+v\n", p)

	// Don't need to derefence (*) struct fields
	// Compiler will perform an implicit &p for you
	// You only need to dereference in standard functions when a argument pointer is required (see below Array Pointer example)
	p.x *= factor
	p.y *= factor

	fmt.Printf("scaleBy (after modification): %+v\n", p)
}

func main() {
	// Doesn't matter if we do or don't get the address space (&) for foo/bar's Point
	foo := &Point{1, 2}
	bar := &Point{6, 8}

	fmt.Printf("Main foo.x: %+v\n", foo.x)
	fmt.Printf("Main bar.x: %+v\n", bar.x)

	foo.scaleBy(5)
	bar.scaleBy(5)

	fmt.Printf("Main foo.x: %+v\n", foo.x)
	fmt.Printf("Main foo.y: %+v\n", foo.y)

	fmt.Printf("Main bar.x: %+v\n", bar.x)
	fmt.Printf("Main bar.y: %+v\n", bar.y)
}
```

> Note: compiler can only apply implicit dereference for variables and struct fields  
> this wouldn't work `Point{1, 2}.scaleBy(5)`

Results in the following output:

```
Main foo.x: 1
Main bar.x: 6
scaleBy (before modification): &{x:1 y:2}
scaleBy (after modification): &{x:5 y:10}
scaleBy (before modification): &{x:6 y:8}
scaleBy (after modification): &{x:30 y:40}
Main foo.x: 5
Main foo.y: 10
Main bar.x: 30
Main bar.y: 40
```

### Array Pointer

Deference an Array pointer so you can mutate the original Array values:

```go
package main

import "fmt"

func main() {  
    x := [3]int{1,2,3}

    func(arr *[3]int) {
        (*arr)[0] = 7
        fmt.Println(arr) //prints &[7 2 3]
    }(&x)

    fmt.Println(x) //prints [7 2 3]
}
```

Alternatively you can utilise a Slice instead of an Array, as the slice 'header' already has a 'pointer' to an underlying Array:

```go
package main

import "fmt"

func main() {  
    x := []int{1,2,3}

    func(arr []int) {
        arr[0] = 7
        fmt.Println(arr) //prints [7 2 3]
    }(x)

    fmt.Println(x) //prints [7 2 3]
}
```

### Type Assertion

```go
if e, ok := err.(net.Error); ok && e.Timeout() {
	//
}

type argError struct {
    arg  int
    prob string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

if ae, ok := e.(*argError); ok {
	//
}
```

### Line Count

Demonstrates how to use `bufio` package to scan a file and read it line by line, and then how to increment a map integer value using the shortcut `map[key]++`. Finally, demonstrates nested maps and ranging over them:

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "n/a", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts)
			f.Close()
		}
	}
	for key, nestedMap := range counts {
		fmt.Printf("Text: %s\n", key)
		for filename, count := range nestedMap {
			fmt.Printf("\tFile: %s\n\tCount: %d\n", filename, count)
		}
		fmt.Println("")
	}
}

func countLines(f *os.File, filename string, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if val, ok := counts[input.Text()]; ok {
			val[filename]++
		} else {
			inner := make(map[string]int)
			inner[filename]++
			counts[input.Text()] = inner
		}
	}
}
```

### Measuring time

```go
package main

import (
	"fmt"
	"time"
)

// Sleep requires a Duration
// time has set of constants we can use (lowest is 1 Duration)
// Second constant is an abstraction over the other constants
func main() {
	start := time.Now()
	time.Sleep(time.Duration(5) * time.Second) // sleep 5 seconds
	secs := time.Since(start).Seconds()

	fmt.Printf("Time spent: %f seconds", secs)
}
```

### Reading a file in chunks

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Create file (truncates file if it already exists)
	file, err := os.Create("created.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Populate byte slice with some content
	b := make([]byte, 0)
	for i := 0; i < 5; i++ {
		b = append(b, '!')
		b = append(b, '\n')
		// notice single quotes for Rune rather than double quote for String
	}
	for i := 0; i < 5; i++ {
		b = append(b, '?')
		b = append(b, '\n')
		// notice single quotes for Rune rather than double quote for String
	}
	for i := 0; i < 5; i++ {
		b = append(b, '%')
		b = append(b, '\n')
		// notice single quotes for Rune rather than double quote for String
	}

	// Write file contents
	bytesWritten, err := file.Write(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Bytes written: %+v\n", bytesWritten)

	// Although getting the bytes written was useful for us
	// in this example, you might need to get total bytes
	// which can be done by copying file contents into dev/null
    // io.Copy(ioutil.Discard, resp.Body)

	// Get current offset
	// 1st arg is how much to seek forward/backwards by
	// 2nd arg is relative to different settings
	// 		0 == relative to start of file
	// 		1 == current offset
	// 		2 == relative to end of file
	currentOffset, err := file.Seek(0, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current offset: %d\n", currentOffset)
	file.Seek(-currentOffset, 1) // Return to start of file for next Read

	// Read buffered view of file
	data := make([]byte, 10, bytesWritten) // create slice with underlying Array capacity set to total file bytes size
	eof := false
	for !eof {
		count, err := file.Read(data)
		if err != nil {
			eof = true
		}
		fmt.Printf("read %d bytes: %q\n", count, data[:count])
	}
}
```

### Time and Channels

Basic example that pauses execution until the timer has expired (you would use this over a `timer.Sleep` because you can cancel a timer before it has expired;):

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 2)

	<-timer.C // pauses for two seconds

	fmt.Println("Timer expired")
}
```

Example of cancelling the timer:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 2)

	// Expensive process run in a separate thread
	go func() {
		<-timer.C
		fmt.Println("Timer expired")
	}()

	stop := timer.Stop() // cancel the timer
	fmt.Println(stop)    // true
}
```

We can do a similar thing with Timers:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)

	// Repetitive process
	go func() {
		// Range over the channel rather than pull from it
		for t := range ticker.C {
			fmt.Println("Tick:", t)
		}
	}()

	// Stop ticker after three ticks/intervals
	time.Sleep(time.Millisecond * 1500)
	ticker.Stop()
}
```

We can combine all these items together with a `select` statement like so:

```go
package main

import "time"
import "fmt"

func main() {
	timeChan := time.NewTimer(time.Second).C
	tickChan := time.NewTicker(time.Millisecond * 400).C

	// Used to signify we're done with this program
	doneChan := make(chan bool)

	// Sleep for two seconds, then notify the channel we're done
	go func() {
		time.Sleep(time.Second * 2)
		doneChan <- true
	}()

	for {
		select {
		case <-timeChan:
			fmt.Println("Timer expired")
		case <-tickChan:
			fmt.Println("Ticker ticked")
		case <-doneChan:
			fmt.Println("Done")
			return
		}
	}
}
```

The output of this program would be something like:

```
Ticker ticked
Ticker ticked
Timer expired
Ticker ticked
Ticker ticked
Done
```

### Quit a Channel

I would imagine that for most cases you'll want to use a `time.NewTimer` as seen in previous examples if you want to stop a goroutine that's processing a long running program. The following example is more for stopping a goroutine that's running code at a set interval (although using `time.NewTicker` would probably be more appropriate):

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan bool)

	// Run a piece of code at a set interval
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				fmt.Println("Not ready to stop this goroutine")
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()

	// Do other stuff for two seconds
	time.Sleep(time.Second * 2)

	// Quit goroutine
	quit <- true

	fmt.Println("Goroutine was stopped")
}
```

### Starting and Stopping things with Channels

Starting a goroutine:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// Use a struct type channel as it clarifies your intent
	// Which is this channel is used for 'signalling'
	start := make(chan struct{})

	for i := 0; i < 10000; i++ {
		go func() {
			<-start // wait for the start channel to be closed
			fmt.Println("do stuff")
		}()
	}

	// at this point, all goroutines are ready to go
	// we just need to tell them to start by
	// closing the start channel
	close(start)

	fmt.Println("Let's pause briefly to give goroutines time to execute")

	time.Sleep(time.Millisecond * 10)
}
```

Stopping a goroutine:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// Use a struct type channel as it clarifies your intent
	// Which is this channel is used for 'signalling'
	done := make(chan struct{})

	// Long running process put onto a thread
	go func() {
		fmt.Println("Inside thread doing expensive processing")
		time.Sleep(time.Second * 5)
		close(done)
	}()

	fmt.Println("Do other things")

	// Wait for long running process to finish
	<-done

	fmt.Println("Do more things")
}
```

### Channel Pipelines

The principle of a pipeline, is to take data from one function and pass it into another function, that receiving function will process the received data and then that result is returned and subsequently passed onto another function... rinse and repeat for however long your pipeline needs to be.

In the below example (copied from [here](https://blog.gopheracademy.com/advent-2015/automi-stream-processing-over-go-channels/)) demonstrates how a set of functions accept a channel and return a channel and so channels is the 'data' that is passed around the pipeline functions:

```go
package main

import "fmt"
import "sync"

func ingest() <-chan []string {
	out := make(chan []string)
	go func() {
		out <- []string{"aaaa", "bbb"}
		out <- []string{"cccccc", "dddddd"}
		out <- []string{"e", "fffff", "g"}
		close(out)
	}()
	return out
}

func process(concurrency int, in <-chan []string) <-chan int {
	var wg sync.WaitGroup
	wg.Add(concurrency)

	out := make(chan int)

	work := func() {
		for data := range in {
			for _, word := range data {
				out <- len(word)
			}
		}
		wg.Done()

	}

	go func() {
		for i := 0; i < concurrency; i++ {
			go work()
		}

	}()

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func store(in <-chan int) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for data := range in {
			fmt.Println(data)
		}
	}()
	return done
}

func main() {
	// stage 1 ingest data from source
	in := ingest()

	// stage 2 - process data
	reduced := process(4, in)

	// stage 3 - store
	<-store(reduced)
}
```

### Templating

Here is a basic program that uses a Struct for its data source:

```go
package main

import (
	"log"
	"os"
	"text/template"
)

type dataSource struct {
	Baz int
}

func (ds dataSource) Foo() string {
	return "I am foo"
}

func (ds dataSource) Bar() string {
	return "I am bar"
}

const templ = `
	Foo: {{.Foo}}
	Piping: {{.Bar | printf "Bar: %s"}}
	Function: {{.Baz | qux}}
`

func qux(baz int) int {
	return baz * 2
}

// template.Must handles parsing errors better
var setupTemplate = template.Must(
	template.New("whatever").
		Funcs(template.FuncMap{"qux": qux}).
		Parse(templ),
)

func main() {
	ds := dataSource{5}

	if err := setupTemplate.Execute(os.Stdout, ds); err != nil {
		log.Fatal(err)
	}
}
```

> Note: `printf` is a built-in function for templating and is functionally equivalent to `fmt.Sprintf`

Program output:

```
Foo: I am foo
Piping: Bar: I am bar
Function: 10
```

Here is a HTML templating version:

```go
package main

import (
	"html/template"
	"log"
	"os"
)

var data struct {
	A string        // untrusted plain text
	B template.HTML // trusted HTML
}

const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`

func main() {
	t := template.Must(template.New("escape").Parse(templ))

	data.A = "<b>Hello</b>"
	data.B = "<b>Hello</b>"

	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
```

The output would be:

```
<p>A: &lt;b&gt;Hello&lt;/b&gt;</p>
<p>B: <b>Hello</b></p>
```

### Error handling with context

The following code outputs: 

```
This is our custom error with some more context prefixed: oh noes!
```

```go
package main

import (
	"errors"
	"fmt"
)

type errWithContext struct {
	err error
	msg string
}

func (e errWithContext) Error() string {
	return e.msg + ": " + e.err.Error()
}

func triggerError() (bool, error) {
	return false, errors.New("oh noes!")
}

func main() {
	var e *errWithContext

	_, err := triggerError()
	if err != nil {
		e = &errWithContext{
			err,
			"This is our custom error with some more context prefixed",
		}
	}

	fmt.Print(e.Error())
}
```

### Socket programming with TCP server

There are two main types of sockets:

1. STREAM sockets (e.g. TCP)
2. DATAGRAM sockets (e.g. UDP)

> Note: a "unix domain socket" is actually a physical file  
> it's useful for local (same host) data communication

The principle steps behind a socket is:

- Create the socket
- Bind the socket to an address (e.g. `127.0.0.1:80`)
- Listen for socket connections
- Accept the socket connection

There are two main packages in our below example: `server.go` and `client.go`. 

Run both of them in separate terminals (e.g. `go run ...`)

Then for the `client.go` type your message followed by a new line, for example:

```
Hello World
Message from server: HELLO WORLD
```

Whilst in the `server.go` terminal you should see:

```
Starting TCP server...
Message Received: Hello World
```

The code for this program is as follows:

server.go

```go
package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Starting TCP server...")

	// Listen on all network interfaces (e.g. 0.0.0.0)
	// Documentation: godoc net Listener | less
	listener, _ := net.Listen("tcp", ":8081")

	// Accept connection on the port we specified (see above)
	connection, _ := listener.Accept()

	// Handle incoming connections forever
	for {
		// Listen for message to process ending in newline (\n)
		// Note: single quotes needed for type byte (double quotes is a string)
		message, _ := bufio.NewReader(connection).ReadString('\n')

		// Output message received
		fmt.Println("Message Received:", string(message))

		// Do something with the message (e.g. uppercase it)
		newmessage := strings.ToUpper(message)

		// Send new string back to client
		connection.Write([]byte(newmessage + "\n"))
	}
}
```

client.go

```go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Open socket connection to a locally runnning TCP server
	connection, _ := net.Dial("tcp", "127.0.0.1:8081")

	// Handle incoming responses
	for {
		// Read the input
		reader := bufio.NewReader(os.Stdin)

		// Message to be sent
		// Note: single quotes needed for type byte (double quotes is a string)
		// Documentation: godoc bufio ReadString | less
		// ReadString reads until the first occurrence of the delimiter \n in the input
		text, _ := reader.ReadString('\n')

		// Send message to open Socket
		fmt.Fprintf(connection, text+"\n")

		// Listen for response
		// Note: single quotes needed for type byte (double quotes is a string)
		message, _ := bufio.NewReader(connection).ReadString('\n')

		fmt.Println("Message from server: " + message)
	}
}
```

### Comparing maps

This code demonstrates how to be careful about false positives!

```go
package main

import "fmt"

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		// fail fast
		return false
	}

	for k, xv := range x {
		// Verify "missing" key and "present but zero" key value
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
		
		/*
		// The following condition would incorrectly return "true" for the below example comparison!
		// This is because the empty value for an int type is a zero, while the actual value of x's key is zero
		if xv != y[k] {
			return false
		}
		*/
	}

	return true
}

func main() {
	fmt.Println(
		equal(map[string]int{"A": 0}, map[string]int{"B": 42}),
	)
}
```

### Embedded Structs

The first example demonstrates a 'named' field utilising an embedded Struct:

```go
package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Center Point // named embeded field
	Radius int
}

type Wheel struct {
	Circle Circle // named embeded field
	Spokes int
}

func main() {
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20

	fmt.Printf("%+v", w)
}
```

Which prints:

```
{Circle:{Center:{X:8 Y:8} Radius:5} Spokes:20}
```

The second example demonstrates an 'anonymous' field instead:

```go
package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w.X = 8       // w.Circle.Point.X
	w.Y = 8       // w.Circle.Point.Y
	w.Radius = 5  // w.Circle.Radius
	w.Spokes = 20

	fmt.Printf("%+v", w)
}
```

Which prints:

```
{Circle:{Point:{X:8 Y:8} Radius:5} Spokes:20}
```

> Note: anonymous fields don't work shorthand literal Struct

The following example demonstrates how methods of a composited object can be accessed from the consuming object:

```go
package main

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) foo() {
	fmt.Printf("foo: %+v\n", p)
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w.X = 8      // w.Circle.Point.X
	w.Y = 8      // w.Circle.Point.Y
	w.foo()      // w.Circle.Point.foo()
	w.Radius = 5 // w.Circle.Radius
	w.Spokes = 20

	fmt.Printf("%+v", w)
}
```

Which prints:

```
foo: {X:8 Y:8}
{Circle:{Point:{X:8 Y:8} Radius:5} Spokes:20}
```

Here is a more practical example that demonstrates how embedded functionality can make code more expressive:

```go
package main

import (
	"fmt"
	"sync"
)

// Anonymous struct
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string), // initial zero value for map
}

func setValue() {
	cache.Lock()
	cache.mapping["foo"] = "bar"
	cache.Unlock()
}

func main() {
	setValue()

	cache.Lock()
	v := cache.mapping["foo"]
	cache.Unlock()

	fmt.Printf("v: %s", v)
}
```

### Zip File Contents

```go
package main

import (
	"compress/zlib"
	"io"
	"log"
	"os"
)

func main() {
	var err error

	// This defends against an error preventing `defer` from being called
	// As log.Fatal otherwise calls `os.Exit`
	defer func() {
		if err != nil {
			log.Fatalln("\nDeferred log: \n", err)
		}
	}()

	src, err := os.Create("source.txt")
	if err != nil {
		return
	}
	src.WriteString("source content")
	src.Close()

	dest, err := os.Create("new.txt")
	if err != nil {
		return
	}

	openSrc, err := os.Open("source.txt")
	if err != nil {
		return
	}

	zdest := zlib.NewWriter(dest)
	if _, err := io.Copy(zdest, openSrc); err != nil {
		return
	}

	// Close these explicitly
	zdest.Close()
	dest.Close()

	n, err := os.Open("new.txt")
	if err != nil {
		return
	}

	r, err := zlib.NewReader(n)
	if err != nil {
		return
	}
	defer r.Close()
	io.Copy(os.Stdout, r)

	err = os.Remove("source.txt")
	if err != nil {
		return
	}

	err = os.Remove("new.txt")
	if err != nil {
		return
	}
}
```

### RPC

For details of what RPC means, see: https://gist.github.com/Integralist/f5856b94e002bcfd4ce7

Only methods that satisfy these criteria will be made available for remote access; other methods will be ignored:

- the method's type is exported.
- the method is exported.
- the method has two arguments, both exported (or builtin) types.
- the method's second argument is a pointer.
- the method has return type error.

In effect, the method must look schematically like

```go
func (t *T) MethodName(argType T1, replyType *T2) error
```

The setup for a simple RPC example is:

1. Create remote package Foo that will consist of functions to be made available via RPC
2. Create remote package that will expose package Foo
3. Create client package that connects to remote via RPC

There are two variations:

1. RPC over HTTP
2. RPC over TCP

#### HTTP

So here is the package that consists of functions to be made available via RPC:

```go
package remote

import "fmt"

// Args is a data structure for the incoming arguments
type Args struct {
	A, B int
}

// Arith is our functions return type
type Arith int

// Multiply does simply multiplication on provided arguments
func (t *Arith) Multiply(args *Args, reply *int) error {
	fmt.Printf("Args received: %+v\n", args)
	*reply = args.A * args.B
	return nil
}
```

Here is the remote package that exposes the other package of functionality:

```go
package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/integralist/rpc/remote"
)

func main() {
	arith := new(remote.Arith)

	rpc.Register(arith)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	rpc.Accept(l)
}
```

Here is our client code for connecting to our remote package via RPC:

```go
package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

type args struct {
	A, B int
}

func main() {
	conn, err := net.DialTimeout("tcp", "localhost:1234", time.Minute)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	client := rpc.NewClient(conn)

	var reply int

	e := client.Call("Arith.Multiply", &args{4, 2}, &reply)
	if e != nil {
		log.Fatalf("Something went wrong: %s", err.Error())
	}

	fmt.Printf("The reply pointer value has been changed to: %d", reply)
}
```

#### TCP

Remote RPC Function:

```go
package remote

import "fmt"

// Compose is our RPC functions return type
type Compose string

// Details is our exposed RPC function
func (c *Compose) Details(arg string, reply *string) error {
	fmt.Printf("Arg received: %+v\n", arg)
	*c = "some value"
	*reply = "Blah!"
	return nil
}
```

Remote RPC Endpoint Exposed:

```go
package remote

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/bbc/mozart-api-common/logger"
)

// Endpoint exposes our RPC over TCP service
func Endpoint() {
	compose := new(Compose)

	rpc.Register(compose)
	// rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		logMessage := map[string]interface{}{
			"event":   "FailedTCPListenerConnection",
			"message": fmt.Sprintf("Listener failed to open TCP port 8080: %v", err),
		}
		logger.Error(logMessage)
	}

	// rpc.Accept(listener)
	for {
		conn, err := listener.Accept()
		if err != nil {
			logMessage := map[string]interface{}{
				"event":   "FailedTPCIncomingConnection",
				"message": fmt.Sprintf("Listener failed to accept an incoming connection: %v", err),
			}
			logger.Error(logMessage)
		}

		go rpc.ServeConn(conn)
	}
}
```

Client Connection over TCP to Remote RPC function:

```go
package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string

	e := client.Call("Compose.Details", "my string", &reply)
	if e != nil {
		log.Fatalf("Something went wrong: %v", e.Error())
	}

	fmt.Printf("The 'reply' pointer value has been changed to: %s", reply)
}
```

#### JSON RPC

There is another option (which is required if using another programming language to communicate with your Go RPC service), that is to turn your RPC into a JSON RPC.

> This is because the standard net/rpc uses https://golang.org/pkg/encoding/gob/  
> Which is a Go specific streaming binary format

Effectively just use the same example as above but make the following changes:

- `net/rpc` to `net/rpc/jsonrpc`
- `rpc.Dial` to `jsonrpc.Dial`
- `rpc.ServeConn` to `jsonrpc.ServeConn`

Now your clients can connect via a TCP socket and pass over JSON, as shown in Ruby below:

```ruby
require "socket"
require "json"

socket = TCPSocket.new "localhost", "8080"

# Details of JSON structure can be found here:
# https://golang.org/src/net/rpc/jsonrpc/client.go#L45
# Thanks to Albert Hafvenström (@albhaf) for his help
b = {
  :method => "Compose.Details",
  :params => [{ :Foo => "Foo!", :Bar => "Bar!" }],
  :id     => "0" # id is just echo'ed back to the client
}

socket.write(JSON.dump(b))

p JSON.load(socket.readline)

# => {"id"=>"0", "result"=>"Blah!", "error"=>nil}
```

Here is an updated Go RPC:

```go
package remote

import "fmt"

// Args is structured around the client's provided parameters
// The fields need to be exported too!
type Args struct {
	Foo string
	Bar string
}

// Compose is our RPC functions return type
type Compose string

// Details is our exposed RPC function
func (c *Compose) Details(args *Args, reply *string) error {
	fmt.Printf("Args received: %+v\n", args)
	*c = "some value"
	*reply = "Blah!"
	return nil
}
```

### Enumerator IOTA

Within a constant declaration, the predeclared identifier `iota` represents successive untyped integer constants. It is reset to 0 whenever the reserved word `const` appears in the source.

```go
package main

import "fmt"

const (
	foo = iota // 0
	bar
	_ // skip this value
	baz
)

const (
	beep = iota // 0 (reset)
	boop
)

func main() {
	fmt.Println(foo, bar, baz) // 0 1 3
	fmt.Println(beep, boop)    // 0 1
}
```

### FizzBuzz

```go
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%d FizzBuzz\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d Fizz\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d Buzz\n", i)
		} else {
			fmt.Println(i)
		}
	}
}
```

### Execute Shell Command

```go
var (
  cmdOut []byte
  err    error
)
cmdName := "spurious"
cmdArgs := []string{"ports", "--json"}
if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
  fmt.Fprintln(os.Stderr, "There was an error running 'spurious ports --json' command: ", err)
  os.Exit(1)
}
fmt.Println(string(cmdOut))
```

### New Instance Idiom

```go
package main

import "fmt"

type Sqs struct {
	foo string
}

func (s *Sqs) create() {
	fmt.Println("I'll create stuff")
}

func NewSqs() *Sqs {
	return &Sqs{"bop"}
}

func main() {
	s := NewSqs()
	fmt.Println(s.foo)
	s.create()
}
```

### Mutating Values

```go
package main

import "fmt"

type Compose string

func (c *Compose) Details() string {
	*c = "beep boop"
	return fmt.Sprintf("Here are your details: %v", *c)
}

func main() {
	var c Compose
	c = "hai"
	fmt.Printf("c: %+v\n", c) // c
	fmt.Println(c.Details())
	fmt.Printf("c: %+v\n", c) // beep boop
}
```

### Draining Connections

When using `json.NewDecoder`:

```go
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        var u User
        if r.Body == nil {
            http.Error(w, "Please send a request body", 400)
            return
        }
        err := json.NewDecoder(r.Body).Decode(&u)
        if err != nil {
            http.Error(w, err.Error(), 400)
            return
        }
        fmt.Println(u.Id)
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

...it doesn't read the response Body completely. So when closing the response you might get an error as a stray `\n` could be present later on. You'll need to drain the response instead:

```go
defer func() {
  io.CopyN(ioutil.Discard, r.Body, 512)
  r.Body.Close()
}()
```

> Note: https://github.com/google/go-github/pull/317
