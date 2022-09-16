## Don't use Encode because it will return a base64 string, instead use this:

```go
responseWriter.Header().Set("Content-Type", "application/json")
responseWriter.WriteHeader(http.StatusOK)
jsonData := []byte(`{"status":"OK"}`)
responseWriter.Write(jsonData)
```

## If you want to return a base64 string, use this then:
```go
responseWriter.Header().Set("Content-Type", "application/json")
responseWriter.WriteHeader(http.StatusOK)
jsonData := []byte(`{"status":"OK"}`)
json.NewEncoder(responseWriter).Encode(jsonData)
```
