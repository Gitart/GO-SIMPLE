
## Init global chanel
```go
   ar GlobalCountDbConnect int64=0
   var GlobalChan = make(chan int64)
```

## Declare read chanel
```go
func init(){
     go ReadCount()  
}
```

## Read chanel 
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
```go
func SentToCahnel(){
  GlobalCountDbConnect ++
  GlobalChan<-GlobalCountDbConnect
  }
```     
