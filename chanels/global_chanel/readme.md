
## Init global chanel
Обявление глобального канала и глобальной перменной
```go
   ar GlobalCountDbConnect int64=0
   var GlobalChan = make(chan int64)
```

## Declare read chanel
Запуск функции чтения из канала в рутине
```go
func init(){
     go ReadCount()  
}
```

## Read chanel 
Функция чтения из канала   
Работающая в постоянном цикле    

```go
func ReadCount(){
     fmt.Println("Inut chanel..")
      
// Вычитка очериди из кнала
for{
   select {
    case msg:= <-GlobalChan:
        fmt.Println("sent message", msg)
    // default:
        // fmt.Println("no message sent")
    }
}

}
```

## Send to chanel
Функция посылающая в канал значения через глобальную перменную
```go
func SentToCahnel(){
  GlobalCountDbConnect ++
  GlobalChan<-GlobalCountDbConnect
  }
```     
