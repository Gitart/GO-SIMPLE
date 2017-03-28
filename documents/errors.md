## Обработка ошибок

```golang
  
  package main
  import "errors"
  
  func main(){
  // Описание ошибок
	 var ErrMissingFile = errors.New("http: ERRROR ")
	
   log.Fatal(ErrMissingFile)
   }
   
```

        
