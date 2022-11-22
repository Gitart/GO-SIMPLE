## Округление дисятичных

```go
func round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
    output := math.Pow(10, float64(precision))
    return float64(round(num * output)) / output
}
```

**Usage:**

```go
fmt.Println(toFixed(1.2345678, 0))  // 1
fmt.Println(toFixed(1.2345678, 1))  // 1.2
fmt.Println(toFixed(1.2345678, 2))  // 1.23
fmt.Println(toFixed(1.2345678, 3))  // 1.235 (rounded up)
```
