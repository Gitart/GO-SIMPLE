package main
 
import (
	"fmt"
	"sync"
	"time"
)
 
func main() {
 
	t := time.NewTimer(3 * time.Second)
	fmt.Printf("Начало ожидания - %v\n", time.Now().Format(time.UnixDate))
	<-t.C
	fmt.Printf("Код выполнен - %v\n", time.Now().Format(time.UnixDate))
 
	wg := &sync.WaitGroup{}
	wg.Add(1)
	fmt.Printf("Начало ожидания для AfterFunc - %v\n", time.Now().Format(time.UnixDate))
	time.AfterFunc(3*time.Second, func() {
		fmt.Printf("Код выполнен для AfterFunc - %v\n", time.Now().Format(time.UnixDate))
		wg.Done()
	})
 
	wg.Wait()
 
	fmt.Printf("Ожидание time.After - %v\n", time.Now().Format(time.UnixDate))
	<-time.After(3 * time.Second)
	fmt.Printf("Итог кода - %v\n", time.Now().Format(time.UnixDate))
 
}
