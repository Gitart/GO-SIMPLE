Открыть сопрограмму в golang очень просто, требуется только одно ключевое слово go.
```go
    package main
      
    import (
            "fmt"
            "time"
    )
     
     
    func main(){
            for i := 0;i<10;i++{
                    go func(i int){
                            for{
                                    fmt.Printf("%d",i);
                               }
                    }(i)
            }
            time.Sleep(time.Millisecond);
    }
```
Распечатать результат
```
    5551600088800499999991117777777742222220000044444444888888888999
    9666665111177777777777777777777777777333333333333333399999999999
    999999999999999999999999999999444442224444444488888888222222222
    20888886666666655555555555444011111111111111000000000999999555555
    5554444444000077777666666311111197777778888222277777753333444444
    9999997777772222000077774444444444444444444
```
Видно, что он полностью случайный, и какой из них печатать, зависит от расписания сопрограммы планировщиком.
По сравнению с потоками, горутины имеют функцию, которая не позволяет вытеснять поток. Если сопрограмма занимает поток и не освобождает его или не блокирует, то поток никогда не будет передан Управлению, давайте возьмем пример для проверки
```go
    package main
      
    import (
           "time"
    )
    func main(){
            for i := 0;i<10;i++{
                    go func(i int){
                            for{
                                    i++                                
                              }
                    }(i)
            }
            time.Sleep(time.Millisecond);
    }
```
После выполнения этой программы она никогда не завершится и не заполнит процессор. Причина в том, что в горутине i ++ был выполнен без выпуска, но поток был занят. Когда четыре потока заполнены, все остальные горутины не работают. Была возможность выполнить, поэтому программа, которая должна была завершиться через одну секунду, так и не завершилась, а процессор был полностью загружен и затем запущен, но почему этого не произошло с Printf в предыдущем примере? Поскольку Printf на самом деле является операцией io, операция io будет заблокирована. При блокировании горутина автоматически освобождает владение потоком, поэтому другие горутины имеют возможность выполнять. Помимо блокировки io, golang также предоставляет API, позволяющий Мы можем передать управление вручную, то есть Gosched (), когда мы вызываем этот метод, goroutine будет активно освобождать управление потоком
```go
    package main
      
    import (
           "time"
          "runtime"
    )
    func main(){
            for i := 0;i<10;i++{
                    go func(i int){
                            for{
                                    i++;
                                    runtime.Gosched();                                
                              }
                    }(i)
            }
            time.Sleep(time.Millisecond);
    }
```
После модификации, через одну секунду, код завершается нормально.
Обычное переключение триггера с помощью горутины, есть несколько ситуаций

    1、I/O,select
     
    2、channel
     
    3. Жду блокировки
     
     4. Вызов функции (это возможность переключиться, переключится ли она определяется планировщиком)
     
    5、runtime.Gosched()

После разговора об основах использования горутин, давайте поговорим о взаимодействии между горутинами. Идея общения в Go такова: «Не общайтесь, обмениваясь данными, а делитесь данными через общение». Общение в Go происходит в основном через каналы . Подобно двустороннему каналу в оболочке unix, он может принимать и отправлять данные,
Давайте посмотрим на пример,
```go
    package main
      
    import(
            "fmt"
            "time"
    )
     
    func main(){
            c := make(chan int)
            go func(){
               for{
                    n := <-c;
                    fmt.Printf("%d",n)
                  }
            }()
     
            c <- 1;
            c <- 2;
            time.Sleep(time.Millisecond);
     
     
    }
```
Результат печати12, Мы используем make для создания типа канала и указываем тип сохраняемых данных через<-Чтобы получать и отправлять данные,c <- 1Чтобы отправить данные 1 на канал c,n := <-c;Указывает, что данные получены из канала c. По умолчанию и отправка, и получение данных заблокированы. Это позволяет нам легко писать синхронный код. Из-за блокировки легко переключать горутины, а после отправки данных он должен будет получен, иначе он продолжит блокировку, а программа сообщит об ошибке и завершится.
В этом примере сначала отправляются данные 1 в c, основная горутина блокируется, а открытая горутина выполняется для чтения данных и печати данных. Затем основная горутина блокируется, а вторые данные 2 отправляются в c. goroutine возвращает Когда чтение данных заблокировано, когда данные 2 успешно прочитаны, печатается 2. Через одну секунду основная функция завершается, все горутины уничтожаются, и программа завершается

Мы внимательно смотрим на этот код. На самом деле, это проблема. В разработанной горутине мы рециркулируем и блокируем чтение данных в c. Мы не знаем, когда написано c и больше нет записей. Мы можем полностью уничтожить эту горутину, не занимая ресурсов. Мы можем выполнить эту задачу через закрытый API.
```go
    package main
      
    import (
            "fmt"
            "time"
    )
     
    func main(){
            c := make(chan int);
            go func(){
                for{
                    p,ok := <-c;
                    if(!ok){
                            fmt.Printf("jieshu");
                            return
                    }
                    fmt.Printf("%d",p);
                   }
            }()
            for i := 0;i<10;i++{
                    c<-i
            }
            close(c);
    }
```
Когда мы закончим писать в канал, мы можем позвонитьcloseМетод явного сообщения получателю о том, что запись в канал завершена. При приеме мы можем судить, завершена ли запись, в соответствии со вторым полученным значением, логическим значением, если оно ложно, это означает, что это канал был закрыт, и нам не нужно продолжать блокировать чтение на канале.
В дополнение к оценке второго логического параметра go также предоставляет диапазон для чтения канала в цикле, и он выйдет из цикла, когда канал будет закрыт.
```go
    package main
      
    import (
            "fmt"
            "time"
    )
     
    func main(){
            c := make(chan int);
            go func(){
            //    for{
            //      p,ok := <-c;
            //      if(!ok){
            //              fmt.Printf("jieshu");
            //              return
            //      }
                    for p := range c{
                            fmt.Printf("%d",p);
                    }
                    fmt.Printf("jieshu");
            //   }
            }()
            for i := 0;i<10;i++{
                   c<-i
            }
            close(c);
            time.Sleep(time.Millisecond);
     
    }
  ```   

Оба способа напечатаны123456789jieshu

Кроме того, с помощью буферизованных каналов мы можем создавать каналы с буферами.Метод использования заключается в передаче второго параметра при создании канала, чтобы указать количество буферов.
```go
    package main
     
    import "fmt"
     
    func main() {
             c: = make (chan int, 2) // Измените 2 на 1, чтобы сообщить об ошибке, измените 2 на 3 для нормальной работы
        c <- 1
        c <- 2
        fmt.Println(<-c)
        fmt.Println(<-c)
    }
```
В этом примере, когда мы создаем канал, мы передаем параметр 2 и можем сохранить два данных. Запись первых двух данных может быть неблокирующей, и нет необходимости ждать, пока данные будут считаны. Если мы будем записывать три данных непрерывно, он сообщит об ошибке и заблокирует запись третьих данных и не сможет перейти к следующему шагу.

Наконец, давайте поговорим о select. Это очень похоже на select в модели io операционной системы. Давайте рассмотрим пример канала, который выполняется первым.
```go
    package main
      
    import (
            "fmt"
            "time"
    )
     
    func main(){
     
            c := make(chan int);
            c2:= make(chan int);
     
            go func(){
             for{
                    select{
                            case p := <- c : fmt.Printf("c:%d\n",p);
                            case p2:= <- c2: fmt.Printf("c2:%d\n",p2);
                    }
                }
            }()
     
            for i :=0;i<10;i++{
                    go func(i int){
                            c <- i
                    }(i)
                    go func (i int){
                            c2 <-i
                    }(i)
            }
            time.Sleep(5*time.Millisecond);
    }
```
Результат печати
```
    c:0
    c2:1
    c:1
    c:2
    c2:0
    c:3
    c:4
    c:5
    c:7
    c2:2
    c:6
    c:8
    c:9
    c2:3
    c2:5
    c2:4
    c2:6
    c2:7
    c2:8
    c2:9
```
Видно, что прием c и c2 является полностью случайным. Тот, кто первым получит обратный вызов, будет выполнен. Конечно, это не ограничивается получением. Вы также можете использовать функцию выбора при отправке данных. Кроме того, как переключатель оператор, функция выбора в golang.Он также поддерживает настройку по умолчанию.Когда значение не получено, будет выполнен обратный вызов по умолчанию.Если значение по умолчанию не установлено, он будет заблокирован в функции выбора до тех пор, пока не будет завершена определенная передача или прием.

Это основные способы использования goroutine в golang. Вы можете испытать рабочий процесс golang на основе статьи о механизме работы goroutine выше и этой статьи.
Добавить несколько функций обработки пакета времени выполнения

    Goexit
    Выйти из выполняющейся в данный момент горутины, но функция defer продолжит вызывать
    Gosched
    отказывается от полномочий на выполнение текущей горутины, планировщик организует выполнение других ожидающих задач и возобновляет выполнение из этого места в следующий раз.
    NumCPU
    возвращает количество ядер ЦП.
    NumGoroutine
    возвращает общее количество задач, которые выполняются и помещаются в очередь.
    GOMAXPROCS
    используется для установки максимального количества ядер ЦП, которое может быть вычислено параллельно, и возврата предыдущего значения.
