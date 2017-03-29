# Chanels

```golang
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

     for i := 0; i < concurrency; i++ {
         go work()
     }

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
