# Работа и взаимодействие воркеров
## Пример показывает отмену одного воркера другим 


```golang
func init() {
	go Worker1()
	go Worker2()
}

  
// First worker with periodic 5 sec 
func Worker1() {

	for{
	      select {
	        case <- quit:
	        	 fmt.Println("Canceled")
	             return
	        default:
	             fmt.Println("Worker 1")
	        }

           time.Sleep(time.Second*5)
		
		
	}
}


// 
func Worker2() {
	

		fmt.Println("Старт воркера 2")
		time.Sleep(time.Second*20)
		
        quit <- true
        fmt.Println("Worker2 itercept Worker ----------")
}
```
