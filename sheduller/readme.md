![image](https://github.com/Gitart/GO-SIMPLE/assets/3950155/603d14cc-6048-4fc4-9d21-5d0276948540)

## Using [`time.After()`](https://golang.org/pkg/time/#After)

[`time.After()`](https://golang.org/pkg/time/#After) allows us to perform an action after a duration. For example:

**NOTE:** The helper constants `time.Second`, `time.Minute` and `time.Hour` are all of type [`time.Duration`](https://golang.org/pkg/time/#Duration). So if we have to supply a duration, we can multiply these constants by a number (e.g. `time.Second * 5`) and the expression returns a [`time.Duration`](https://golang.org/pkg/time/#Duration).

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // This will block for 5 seconds and then return the current time
    theTime := <-time.After(time.Second * 5)
    fmt.Println(theTime.Format("2006-01-02 15:04:05"))
}
```

```
2019-09-22 09:33:05
```

## [](#using-raw-timeticker-endraw-)Using [`time.Ticker`](https://golang.org/pkg/time/#Ticker)

[`time.After()`](https://golang.org/pkg/time/#After) is great for one-time actions, but the power of cron jobs are in performing repeated actions.

So, to do that we use a [`time.Ticker`](https://golang.org/pkg/time/#Ticker). For most use cases, we can use the helper function [`time.Tick()`](https://golang.org/pkg/time/#Tick) to create a ticker. For example:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // This will print the time every 5 seconds
    for theTime := range time.Tick(time.Second * 5) {
        fmt.Println(theTime.Format("2006-01-02 15:04:05"))
    }
}
```

```
2019-09-22 10:07:54
2019-09-22 10:07:59
2019-09-22 10:08:04
2019-09-22 10:08:09
2019-09-22 10:08:14
```

**NOTE:** When using a `Ticker`, the first event will be triggered **AFTER** the delay.

### [](#dangers-of-using-raw-timetick-endraw-)Dangers of using [`time.Tick()`](https://golang.org/pkg/time/#Tick)

When we use the [`time.Tick()`](https://golang.org/pkg/time/#Tick) function, we do not have direct access to the underlying [`time.Ticker`](https://golang.org/pkg/time/#Ticker) and so we cannot close it properly.

If we never need to explicitly stop the ticker (for example if the ticker will run all the time), then this may not be an issue. However, if we simply ignore the ticker, the resources will not be freed up and it will not be garbage collected.

### [](#limitations-using-raw-timetick-endraw-)Limitations using [`time.Tick()`](https://golang.org/pkg/time/#Tick)

There are several things we can't easily do with [`time.Ticker`](https://golang.org/pkg/time/#Ticker):

* Specify a start time
* Stop the ticker

## [](#extending-raw-timetick-endraw-using-a-custom-function)Extending [`time.Tick()`](https://golang.org/pkg/time/#Tick) using a custom function

To overcome the limitations of [`time.Tick()`](https://golang.org/pkg/time/#Tick), I've created a helper function which I use in my projects.

```go
func cron(ctx context.Context, startTime time.Time, delay time.Duration) <-chan time.Time {
    // Create the channel which we will return
    stream := make(chan time.Time, 1)

    // Calculating the first start time in the future
    // Need to check if the time is zero (e.g. if time.Time{} was used)
    if !startTime.IsZero() {
        diff := time.Until(startTime)
        if diff < 0 {
            total := diff - delay
            times := total / delay * -1

            startTime = startTime.Add(times * delay)
        }
    }

    // Run this in a goroutine, or our function will block until the first event
    go func() {

        // Run the first event after it gets to the start time
        t := <-time.After(time.Until(startTime))
        stream <- t

        // Open a new ticker
        ticker := time.NewTicker(delay)
        // Make sure to stop the ticker when we're done
        defer ticker.Stop()

        // Listen on both the ticker and the context done channel to know when to stop
        for {
            select {
            case t2 := <-ticker.C:
                stream <- t2
            case <-ctx.Done():
                close(stream)
                return
            }
        }
    }()

    return stream
}
```

### [](#whats-happening-in-the-function)What's happening in the function

The function receives 3 parameters.

1. A [`Context`](): The ticker will be stopped whenever the context is cancelled. So we can create a context with a cancel function, or a timeout, or a deadline, and when that context is cancelled, the function will gracefully release it's resources.
2. A `Time`: The start time is used as a reference to know when to start ticking. If the start time is in the future, it will not start ticking until that time. If it is in the past. The function calculates the first event that will be in the future by adding an appropriate multiple of the delay.
3. A `Duration`: This is the interval between ticks. Calculated from the start time.

## [](#examples-of-using-the-custom-function)Examples of using the custom function

**Run on Tuesdays by 2pm**

```go
ctx := context.Background()

startTime, err := time.Parse(
    "2006-01-02 15:04:05",
    "2019-09-17 14:00:00",
) // is a tuesday
if err != nil {
    panic(err)
}

delay := time.Hour * 24 * 7 // 1 week

for t := range cron(ctx, startTime, delay) {
    // Perform action here
    log.Println(t.Format("2006-01-02 15:04:05"))
}
```

**Run every hour, on the hour**

```go
ctx := context.Background()

startTime, err := time.Parse(
    "2006-01-02 15:04:05",
    "2019-09-17 14:00:00",
) // any time in the past works but it should be on the hour
if err != nil {
    panic(err)
}

delay := time.Hour // 1 hour

for t := range cron(ctx, startTime, delay) {
    // Perform action here
    log.Println(t.Format("2006-01-02 15:04:05"))
}
```

**Run every 10 minutes, starting in a week**

```go
ctx := context.Background()

startTime, err := time.Now().AddDate(0, 0, 7) // see https://golang.org/pkg/time/#Time.AddDate
if err != nil {
    panic(err)
}

delay := time.Minute * 10 // 10 minutes

for t := range cron(ctx, startTime, delay) {
    // Perform action here
    log.Println(t.Format("2006-01-02 15:04:05"))
}
```

## [](#conclusion)Conclusion

With this function, I have much better control over scheduling in my projects. Hopefully, it is also of some use to you.

