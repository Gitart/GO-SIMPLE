go Ticker()

```go
func Ticker(){
	       for range time.Tick(time.Second * 2) {
              GlobalCount=0 	
           }
}
```
