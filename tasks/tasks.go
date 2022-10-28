package main
 
import (
	"fmt"
	"time"
)
 
func main() {
 
	to   := time.After(5 * time.Second)
	list := make([]string, 0)
	done := make(chan bool, 1)
 
	fmt.Println("Начало вставки элементов")
	
	go func() {
		defer fmt.Println("Выход из горутины")
		
		for {
			select {
			case <-to:
				fmt.Println("Время истекло")
				done <- true
				return
			
			default:
				list = append(list, time.Now().String())
			}
		}
	}()
 
	<-done
	fmt.Printf("Получилось вставить %d элементов\n", len(list))
 
}
