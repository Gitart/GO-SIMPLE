Изучаем Mutex, WaitGroup и Once с примерами

В данной статье кратко рассмотрим некоторые конструкции низкоуровневой синхронизации, которые наряду с горутинами и каналами предлагает нам один из самых популярных стандартных библиотечных пакетов Go, а именно [пакет](https://godoc.org/sync) `[sync](https://godoc.org/sync)`. Таких конструкций очень много, а мы изучим лишь три из них, зато с примерами: `WaitGroup`, `Mutex` и `Once`.

Примеры кода можно найти на [GitHub](https://github.com/abhirockzz/just-enough-go). Поехали!

### WaitGroup

`WaitGroup` используется для координации в случае, когда программе приходится ждать окончания работы нескольких горутин. Эта конструкция похожа на `CountDownLatch` в Java. Обратимся к примеру.

Предположим, нам нужно вывести список всех файлов нашего домашнего каталога одновременно. Используем `WaitGroup` для указания числа задач/горутин, завершения которых нам надо дождаться.

В данном случае оно совпадает с числом файлов/каталогов домашнего каталога. Используем `Wait()` для блокировки, пока счётчик `WaitGroup` не дойдёт до нуля.

```
...
func main() {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        panic(err)
    }
    filesInHomeDir, err := ioutil.ReadDir(homeDir)
    if err != nil {
        panic(err)
    }
    var wg sync.WaitGroup
    wg.Add(len(filesInHomeDir))
    for _, file := range filesInHomeDir {
        go func(f os.FileInfo) {
            defer wg.Done()
        }(file)
    }
    wg.Wait()
}
...
```

Для выполнения этой программы понадобится:

```
curl https://raw.githubusercontent.com/abhirockzz/just-enough-go/master/sync/wait-group-example.go -o wait-group-example.go
go run wait-group-example.go
```

Под каждую `os.FileInfo`, которую мы находим в домашнем каталоге пользователя, создаётся горутина и при выводе её названия счётчик даёт отрицательное приращение с помощью этого `Done`. Выполнение завершается после того, как программа пробежится по всему содержимому домашнего каталога.

### Мьютекс

Общий мьютекс — это блокировка с общим доступом, которая даёт возможность получать эксклюзивный доступ к тем или иным участкам кода. Далее в простом примере в функции `incr` мы используем общую/глобальную переменную `accessCount`.

```
func incr() {
    mu.Lock()
    defer mu.Unlock()
    accessCount = accessCount + 1
}
```

Обратите внимание, что функция `incr` защищена `мьютексом`, поэтому только одна горутина может иметь к ней доступ. Мы бросаем на неё несколько горутин.

```
loop := 500
for i := 1; i <= loop; i++ {
        go func(c int) {
            wg.Add(1)
            defer wg.Done()
            incr()
        }(i)
}
```

При выполнении результат здесь всегда будет один и тот же, т.е. `Final = 500` (так как выполняются 500 итераций цикла for). Для выполнения программы понадобится:

```
curl https://raw.githubusercontent.com/abhirockzz/just-enough-go/master/sync/mutex-example.go -o mutex-example.go
go run mutex-example.go
```

Добавьте комментарий к следующим двум строчкам в функции `incr` (или удалите эти строчки):

```
mu.Lock()
defer mu.Unlock()
```

Запустите программу на локальном компьютере и снова выполните программу. Результат будет иной. Например, `Final = 474`.

Настоятельно рекомендую почитать о `[RWMutex](https://golang.org/pkg/sync/#RWMutex)`. Это особая разновидность захвата, предполагающая параллельное чтение (несколько читателей: когда к одним и тем же данным имеют доступ несколько читающих потоков или горутин, как в нашем случае) и синхронизированные записи (один писатель: когда к данным имеет доступ только записывающий поток или горутина).

### Once

Позволяет определить задачу для однократного выполнения за всё время работы программы. Содержит одну-единственную функцию `Do`, позволяющую передавать другую функцию для однократного применения. Вот вам пример:

Допустим, вы создаёте REST API с помощью пакета Go `net/http` и хотите, чтобы участок кода выполнялся только после вызова обработчика HTTP-данных (например, для соединения с базой данных).

Используем в коде `once.Do`: теперь можете быть уверены, что он выполнится только при первом вызове обработчика.

Вот как выглядит функция для однократного выполнения:

```
func oneTimeOp() {
    fmt.Println("one time op start")
    time.Sleep(3 * time.Second)
    fmt.Println("one time op started")
}
```

Видите этот `once.Do(oneTimeOp)`? Вот что мы делаем внутри нашего HTTP-обработчика!

```
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("http handler start")
        once.Do(oneTimeOp)
        fmt.Println("http handler end")
        w.Write([]byte("done!"))
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Запускаем код и получаем доступ к конечной точке REST.

```
curl https://raw.githubusercontent.com/abhirockzz/just-enough-go/master/sync/once-example.go -o once-example.go
go run once-example.go
```

И с другого терминала:

```
curl localhost:8080
//результат - готово!
```

При первом доступе возврат функции будет немного медленным, и вы увидите следующие логи сервера:

```
http handler start
one time op start
one time op end
http handler end
```

При повторных запусках (сколько бы вы ни пытались) функция `oneTimeOp` не выполнится. Для подтверждения проверьте логи.
