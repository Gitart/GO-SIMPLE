## Go Chanels
https://itnext.io/explain-to-me-go-concurrency-worker-pool-pattern-like-im-five-e5f1be71e2b0


![image](https://user-images.githubusercontent.com/3950155/134639099-55319ca1-e199-4c37-aa63-d5a5bfb05d10.png)

## 1\. Jobs Batch

I created a minimal work unit called `Job`, composed of an `ExecutionFn` that would let us write custom logic for the `Job` to returning a `Result`. The latter could be either a `value` or an `error`.

As the second step, I used the `generator` concurrency pattern to stream all the `Job`s into the `WorkerPool`. What is this about? Generating a stream from ranging over some client’s defined `Job`s slice pushing each of them into a channel, the `Job`s channel. Which will be used to feed concurrently the `WorkerPool`.


```go
... [omitted for brevity]

type jobMetadata map[string]interface{}

type Job struct {
	Descriptor JobDescriptor
	ExecFn     ExecutionFn
	Args       interface{}
}

func (j Job) execute(ctx context.Context) Result {
	value, err := j.ExecFn(ctx, j.Args)
	if err != nil {
		return Result{
			Err:        err,
			Descriptor: j.Descriptor,
		}
	}

	return Result{
		Value:      value,
		Descriptor: j.Descriptor,
	}
}
```

## 2\. Job’s channel

It is a buffered channel (workers count capped) that once it’s filled up any further attempt to write will block the current goroutine (in this case the stream’s generator goroutine from 1). At any moment, if any `Job` is present on the channel will be consumed by a `Worker` function for later execution. In this way, the channel will be unblocked for new `Job` writes flowing from the `generator` from the previous point.

## 3\. WorkerPool

This is the main piece of the puzzle, this entity is composed of `Result`s, `Job`s and `Done` channel, plus the number of `Worker`s the pool will host. It will spawn as many `Worker`s on different goroutines as the worker count indicates, AKA ***fanning-out***.

The `Worker`s themselves will be responsible for taking `Job`s from the channel when available. Then they execute the `Job` and publishing it `Result` onto the `Result` s channel. As long as the `cancel()` function is not invoked upon `Context`, the `Worker` would do the previous mentioned. Otherwise, the loop brakes and `WaitGroup` is marked as `Done()`. This is quite similar to think of *“killing the* `*Worker*`*“*.

After all the available `Job`s have been drawn from their channel, the `WorkerPool` will finish its execution by closing its own `Done` and `Result`s channels.

## 4\. Results Channel

As mentioned before, even though workers run on different goroutines, they publish the `Job`’s execution `Result`s by multiplexing them onto the `Result`’s channel, AKA ***fanning-in***. The `WorkerPool`‘s client can read out from this source even the channel is closed for any reason indicated above.

## 5\. Reading Results

The `WorkerPool`’s clients can read out of the `Result`s channel as long as there is at least one of them present on the buffered channel. Otherwise, reading from the empty `Result`s channel blocks the client’s goroutine till a value is present or the channel is closed.

The for loop gets broken once closed `WorkerPool`‘s `Done` channel returns and moves forwards.

## X. Cancel Gracefully

In any case, if the client needs to shut down gracefully the `WorkerPool` execution, either it can call the `cancel()` function upon the given `Context` or have configured a timeout duration defined by with `Context.WithTimeout` method.

Whether one or other option happens (both end up calling `cancel()` function, one explicitly and the other after a time out happens) a closed `Done` channel will be returned from the `Context` which will be propagated to all `Worker` functions. This makes the `for select` loop break, hence the `Worker`s stop consuming `Job`s out of the channel. Then later, the `WaitGroup` is marked as done. But still, running workers will finish their job execution before `WorkerPool` is shut down.

# Sum Up

When we make use of this pattern, we will achieve concurrent `Job`s execution leveraging our system to be more performant and consistent across job executions.

This pattern could be hard to grasp at first glimpse. But, take your time to digest it, especially if you are new to the GoLang Concurrency Model.

One thing that could help is to think of channels as pipes, where data flows from one side to the other and the amount of data that could fit is limited. So if we want to inject more data, we just need to make some extra room for it by taking some data out first whilst we wait. On the other way around, if we want to consume from the pipe there has to be something, otherwise, we wait till that happens. In that way, we use these pipes to communicate and share data across `goroutines`.

// ... [omitted for brevity]

```go
func TestWorkerPool(t *testing.T) {
	wp := New(workerCount)

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	go wp.GenerateFrom(testJobs())

	go wp.Run(ctx)

	for {
		select {
		case r, ok := <-wp.Results():
			if !ok {
				continue
			}

			i, err := strconv.ParseInt(string(r.Descriptor.ID), 10, 64)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			val := r.Value.(int)
			if val != int(i)*2 {
				t.Fatalf("wrong value %v; expected %v", val, int(i)*2)
			}
		case <-wp.Done:
			return
		default:
		}
	}
}
```
# Resources

**\[Implementation\]** A full implementation for this pattern in this [***GitHub repo***](https://github.com/godoylucase/workers-pool)***.***

**\[Book\]** I can recommend a book if you are interested in reading more about the concurrency topic: [Concurrency in Go — OReilly](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/)
For me has been a great resource as a go-to when I face this kind of issue to be solved using concurrent approaches.

Thanks for reading and I hope you find this useful!
