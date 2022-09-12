package main
 
import (
	"fmt"
	"os"
	"os/signal"
	"time"
)
 
func main() {
 
	c := make(chan os.Signal, 1)
	signal.Notify(c)
 
	ticker := time.NewTicker(time.Second)
	stop := make(chan bool)
 
	go func() {
		defer func() { stop <- true }()
		for {
			select {
			case <-ticker.C:
				fmt.Println("Тик")
			case <-stop:
				fmt.Println("Закрытие горутины")
				return
			}
		}
	}()
 
	// Блокировка, пока не будет получен сигнал
	<-c
	ticker.Stop()
 
	// Остановка горутины
	stop <- true
	// Ожидание до тех пор, пока не выполнится
	<-stop
	fmt.Println("Приложение остановлено")
}
