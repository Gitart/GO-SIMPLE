## Name call function

```golang
pc, _, _, _ := runtime.Caller(1)
ff:= fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
fmt.Println(ff)
```

  
