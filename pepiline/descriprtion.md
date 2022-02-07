# pipeline

[![GitHub Workflow Status](https://github.com/deliveryhero/pipeline/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/deliveryhero/pipeline/actions/workflows/ci.yml?query=branch:main) [![codecov](https://camo.githubusercontent.com/533c0182abcf912864839083d224253a925a2a23fb3beb0bea01b1bddb8d2b9b/68747470733a2f2f636f6465636f762e696f2f67682f64656c69766572796865726f2f706970656c696e652f6272616e63682f6d61696e2f67726170682f62616467652e737667)](https://codecov.io/gh/deliveryhero/pipeline) [![GoDoc](https://camo.githubusercontent.com/46170c3ac4e83605d637d8fb4292a78024508b62025733c8db010d1597006675/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f706b672e676f2e6465762d646f632d626c7565)](https://pkg.go.dev/github.com/deliveryhero/pipeline) [![Go Report Card](https://camo.githubusercontent.com/a0248e2c0ef5879e05343ad4a05ef4220333d0af9e352627a683e465fa9b5fb7/68747470733a2f2f676f7265706f7274636172642e636f6d2f62616467652f6769746875622e636f6d2f64656c69766572796865726f2f706970656c696e65)](https://goreportcard.com/report/github.com/deliveryhero/pipeline)

Pipeline is a go library that helps you build pipelines without worrying about channel management and concurrency. It contains common fan-in and fan-out operations as well as useful utility funcs for batch processing and scaling.

If you have another common use case you would like to see covered by this package, please [open a feature request](https://github.com/deliveryhero/pipeline/issues).

## [](https://golangrepo.com/repo/deliveryhero-pipeline#functions)Functions

### [](https://golangrepo.com/repo/deliveryhero-pipeline#func-cancel)func [Cancel](https://golangrepo.com/deliveryhero/pipeline/blob/main/cancel.go#L9)

`func Cancel(ctx context.Context, cancel func(interface{}, error), in <-chan interface{}) <-chan interface{}`

Cancel passes an `interface{}` from the `in <-chan interface{}` directly to the out `<-chan interface{}` until the `Context` is canceled. After the context is canceled, everything from `in <-chan interface{}` is sent to the `cancel` func instead with the `ctx.Err()`.

```go
package main

import (
	"context"
	"github.com/deliveryhero/pipeline"
	"log"
	"time"
)

func main() {
	// Create a context that lasts for 1 second
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create a basic pipeline that emits one int every 250ms
	p := pipeline.Delay(ctx, time.Second/4,
		pipeline.Emit(1, 2, 3, 4, 5),
	)

	// If the context is canceled, pass the ints to the cancel func for teardown
	p = pipeline.Cancel(ctx, func(i interface{}, err error) {
		log.Printf("%+v could not be processed, %s", i, err)
	}, p)

	// Otherwise, process the inputs
	for out := range p {
		log.Printf("process: %+v", out)
	}

	// Output
	// process: 1
	// process: 2
	// process: 3
	// process: 4
	// 5 could not be processed, context deadline exceeded
}
```

### [](https://golangrepo.com/repo/deliveryhero-pipeline#func-collect)func [Collect](https://golangrepo.com/deliveryhero/pipeline/blob/main/collect.go#L13)

`func Collect(ctx context.Context, maxSize int, maxDuration time.Duration, in <-chan interface{}) <-chan interface{}`

Collect collects `interface{}`s from its in channel and returns `[]interface{}` from its out channel. It will collect up to `maxSize` inputs from the `in <-chan interface{}` over up to `maxDuration` before returning them as `[]interface{}`. That means when `maxSize` is reached before `maxDuration`, `[maxSize]interface{}` will be passed to the out channel. But if `maxDuration` is reached before `maxSize` inputs are collected, `[< maxSize]interface{}` will be passed to the out channel. When the `context` is canceled, everything in the buffer will be flushed to the out channel.

### [](https://golangrepo.com/repo/deliveryhero-pipeline#func-delay)func [Delay](https://golangrepo.com/deliveryhero/pipeline/blob/main/delay.go#L10)

`func Delay(ctx context.Context, duration time.Duration, in <-chan interface{}) <-chan interface{}`

Delay delays reading each input by `duration`. If the context is canceled, the delay will not be applied.

### [](https://golangrepo.com/repo/deliveryhero-pipeline#func-emit)func [Emit](https://golangrepo.com/deliveryhero/pipeline/blob/main/emit.go#L4)

`func Emit(is ...interface{}) <-chan interface{}`

Emit fans `is ...interface{}`` out to a` <-chan interface{}\`

### [](https://golangrepo.com/repo/deliveryhero-pipeline#func-merge)func [Merge](https://golangrepo.com/deliveryhero/pipeline/blob/main/merge.go#L6)

`func Merge(ins ...<-chan interface{}) <-chan interface{}`

Merge fans multiple channels in to a single channel

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/deliveryhero/pipeline"
	"github.com/deliveryhero/pipeline/example/db"
)

// SearchResults returns many types of search results at once
type SearchResults struct {
	Advertisements []db.Result `json:"advertisements"`
	Images         []db.Result `json:"images"`
	Products       []db.Result `json:"products"`
	Websites       []db.Result `json:"websites"`
}

func main() {
	r := http.NewServeMux()

	// `GET /search?q=<query>` is an endpoint that merges concurrently fetched
	// search results into a single search response using `pipeline.Merge`
	r.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		if len(query) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// If the request times out, or we receive an error from our `db`
		// the context will stop all pending db queries for this request
		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()

		// Fetch all of the different search results concurrently
		var results SearchResults
		for err := range pipeline.Merge(
			db.GetAdvertisements(ctx, query, &results.Advertisements),
			db.GetImages(ctx, query, &results.Images),
			db.GetProducts(ctx, query, &results.Products),
			db.GetWebsites(ctx, query, &results.Websites),
		) {
			// Stop all pending db requests if theres an error
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		// Return the search results
		if bs, err := json.Marshal(&results); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
		} else if _, err := w.Write(bs); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	})
}
```

### [](https://golangrepo.com/repo/deliveryhero-pipeline#func-process)func [Process](https://golangrepo.com/deliveryhero/pipeline/blob/main/process.go#L12)

`func Process(ctx context.Context, processor Processor, in <-chan interface{}) <-chan interface{}`

Process takes each input from the `in <-chan interface{}` and calls `Processor.Process` on it. When `Processor.Process` returns an `interface{}`, it will be sent to the output `<-chan interface{}`. If `Processor.Process` returns an error, `Processor.Cancel` will be called with the corresponding input and error message. Finally, if the `Context` is canceled, all inputs remaining in the `in <-chan interface{}` will go directly to `Processor.Cancel`.

```go
package main

import (
	"context"
	"github.com/deliveryhero/pipeline"
	"github.com/deliveryhero/pipeline/example/processors"
	"log"
	"time"
)

func main() {
	// Create a context that times out after 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a pipeline that emits 1-6 at a rate of one int per second
	p := pipeline.Delay(ctx, time.Second, pipeline.Emit(1, 2, 3, 4, 5, 6))

	// Use the Multiplier to multiply each int by 10
	p = pipeline.Process(ctx, &processors.Multiplier{
		Factor: 10,
	}, p)

	// Finally, lets print the results and see what happened
	for result := range p {
		log.Printf("result: %d\n", result)
	}

	// Output
	// result: 10
	// result: 20
	// result: 30
	// result: 40
	// result: 50
	// error: could not multiply 6, context deadline exceeded
}
```

### [](https://golangrepo.com/repo/deliveryhero-pipeline#func-processbatch)func [ProcessBatch](https://golangrepo.com/deliveryhero/pipeline/blob/main/process_batch.go#L13)

`func ProcessBatch( ctx context.Context, maxSize int, maxDuration time.Duration, processor Processor, in <-chan interface{}, ) <-chan interface{}`

ProcessBatch collects up to maxSize elements over maxDuration and processes them together as a slice of `interface{}`s. It passed an \[\]interface{} to the `Processor.Process` method and expects a \[\]interface{} back. It passes \[\]interface{} batches of inputs to the `Processor.Cancel` method. If the receiver is backed up, ProcessBatch can holds up to 2x maxSize.

```go
package main

import (
	"context"
	"github.com/deliveryhero/pipeline"
	"github.com/deliveryhero/pipeline/example/processors"
	"log"
	"time"
)

func main() {
	// Create a context that times out after 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a pipeline that emits 1-6 at a rate of one int per second
	p := pipeline.Delay(ctx, time.Second, pipeline.Emit(1, 2, 3, 4, 5, 6))

	// Use the BatchMultiplier to multiply 2 adjacent numbers together
	p = pipeline.ProcessBatch(ctx, 2, time.Minute, &processors.BatchMultiplier{}, p)

	// Finally, lets print the results and see what happened
	for result := range p {
		log.Printf("result: %d\n", result)
	}

	// Output
	// result: 2
	// result: 12
	// error: could not multiply [5], context deadline exceeded
	// error: could not multiply [6], context deadline exceeded
}
```

### [](https://golangrepo.com/repo/deliveryhero-pipeline#func-processbatchconcurrently)func [ProcessBatchConcurrently](https://golangrepo.com/deliveryhero/pipeline/blob/main/process_batch.go#L30)

`func ProcessBatchConcurrently( ctx context.Context, concurrently, maxSize int, maxDuration time.Duration, processor Processor, in <-chan interface{}, ) <-chan interface{}`

ProcessBatchConcurrently fans the in channel out to multiple batch Processors running concurrently, then it fans the out channels of the batch Processors back into a single out chan

```go
package main

import (
	"context"
	"github.com/deliveryhero/pipeline"
	"github.com/deliveryhero/pipeline/example/processors"
	"log"
	"time"
)

func main() {
	// Create a context that times out after 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a pipeline that emits 1-9
	p := pipeline.Emit(1, 2, 3, 4, 5, 6, 7, 8, 9)

	// Wait 4 seconds to pass 2 numbers through the pipe
	// * 2 concurrent Processors
	p = pipeline.ProcessBatchConcurrently(ctx, 2, 2, time.Minute, &processors.Waiter{
		Duration: 4 * time.Second,
	}, p)

	// Finally, lets print the results and see what happened
	for result := range p {
		log.Printf("result: %d\n", result)
	}

	// Output
	// result: 3
	// result: 4
	// result: 1
	// result: 2
	// error: could not process [5 6], process was canceled
	// error: could not process [7 8], process was canceled
	// error: could not process [9], context deadline exceeded
}
```

### [](https://golangrepo.com/repo/deliveryhero-pipeline#func-processconcurrently)func [ProcessConcurrently](https://golangrepo.com/deliveryhero/pipeline/blob/main/process.go#L23)

`func ProcessConcurrently(ctx context.Context, concurrently int, p Processor, in <-chan interface{}) <-chan interface{}`

ProcessConcurrently fans the in channel out to multiple Processors running concurrently, then it fans the out channels of the Processors back into a single out chan

```go
package main

import (
	"context"
	"github.com/deliveryhero/pipeline"
	"github.com/deliveryhero/pipeline/example/processors"
	"log"
	"time"
)

func main() {
	// Create a context that times out after 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a pipeline that emits 1-7
	p := pipeline.Emit(1, 2, 3, 4, 5, 6, 7)

	// Wait 2 seconds to pass each number through the pipe
	// * 2 concurrent Processors
	p = pipeline.ProcessConcurrently(ctx, 2, &processors.Waiter{
		Duration: 2 * time.Second,
	}, p)

	// Finally, lets print the results and see what happened
	for result := range p {
		log.Printf("result: %d\n", result)
	}

	// Output
	// result: 2
	// result: 1
	// result: 4
	// result: 3
	// error: could not process 6, process was canceled
	// error: could not process 5, process was canceled
	// error: could not process 7, context deadline exceeded
}
```

### [](https://golangrepo.com/repo/deliveryhero-pipeline#func-split)func [Split](https://golangrepo.com/deliveryhero/pipeline/blob/main/split.go#L5)

`func Split(in <-chan interface{}) <-chan interface{}`

Split takes an interface from Collect and splits it back out into individual elements Useful for batch processing pipelines (`input chan -> Collect -> Process -> Split -> Cancel -> output chan`).

## [](https://golangrepo.com/repo/deliveryhero-pipeline#types)Types

### [](https://golangrepo.com/repo/deliveryhero-pipeline#type-processor)type [Processor](https://golangrepo.com/deliveryhero/pipeline/blob/main/processor.go#L8)

`type Processor interface { ... }`

Processor represents a blocking operation in a pipeline. Implementing `Processor` will allow you to add business logic to your pipelines without directly managing channels. This simplifies your unit tests and eliminates channel management related bugs.
