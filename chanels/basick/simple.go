func main() {   
ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)   
defer done()   
g, gctx := errgroup.WithContext(ctx)   

// just a ticker every 2s
   g.Go(func() error {
      ticker := time.NewTicker(2 * time.Second)
      i := 0
      for {
         i++
         if i > 10 {
            return nil
         }         select {
         case <-ticker.C:
            fmt.Println("ticker 2s ticked")
         case <-gctx.Done():
            fmt.Println("closing ticker 2s goroutine")
            return gctx.Err()
         }
      }
   })   // just a ticker every 1s
   g.Go(func() error {
      ticker := time.NewTicker(1 * time.Second)
      i := 0
      for {
         i++
         if i > 10 {
            return nil
         }
         select {
         case <-ticker.C:
            fmt.Println("ticker 1s ticked")
         case <-gctx.Done():
            fmt.Println("closing ticker 1s goroutine")
            return gctx.Err()
         }
      }
   })   // wait for all errgroup goroutines
   go func() {
      err := g.Wait()
      if err != nil {
         if errors.Is(err, context.Canceled) {
            fmt.Println("context was canceled")
         } else {
            fmt.Printf("received error: %v\n", err)
         }
      } else {
         fmt.Println("finished clean")
      }
   }()
   time.Sleep(5 * time.Second)
   fmt.Println("before done")
   done()
   fmt.Println("after done")
   time.Sleep(1 * time.Second)
}
