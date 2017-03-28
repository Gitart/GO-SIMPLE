# Вариант вывода информации в тело документа

```golang


func Test(w http.ResponseWriter, rq *http.Request){
	  io.WriteString(w, "hello, world!\n")
	  fmt.Fprintln(w, "Входите пожалуйста")
	  w.Write([]byte("Hello word"))
}
```
