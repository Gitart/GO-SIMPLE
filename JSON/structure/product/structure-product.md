# Product

```go
// Result
type AtsResultProduct struct {
     Result      []AtsProduct           `json: "result"` 
}

// Product Info
type AtsProductInfo struct{
	  Status      string                `json: "status"`           
      Message     string                `json: "message"`  
      Data        AtsResultProduct      `json: "data"`
}
```
