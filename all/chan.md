

```golang
// Check chanel function
func ChainCall(fns ...func() (*Chain, error)) (err error) {
    for _, fn := range fns {
        if _, err = fn(); err != nil {
            break
        }
    }
    return
}
```
