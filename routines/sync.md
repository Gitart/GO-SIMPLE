import "sync"

var wait sync.WaitGroup

for(){
  go func(){
         ...
         wait.Done()
  }()
  
  wait.Wait()
 }
