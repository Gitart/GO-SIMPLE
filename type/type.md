

## 3 ways to recognize type at runtime:
### Using string formatting

```golang
func typeof(v interface{}) string {
    return fmt.Sprintf("%T", v)
}
```


### Using reflect package
```golang
func typeof(v interface{}) string {
    return reflect.TypeOf(v).String()
}
```

### Using type assertions
```golang
func typeof(v interface{}) string {
    switch t := v.(type) {
    case int:
        return "int"
    case float64:
        return "float64"
    //... etc
    default:
        _ = t
        return "unknown"
    }
}
```
