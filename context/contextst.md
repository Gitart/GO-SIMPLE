# Контекст(Context)

### Самая важная часть пакета context это тип Context:

```
// Контекст предоставляет механизм дедлайнов, сигнал отмены, и доступ к запросозависимым значениям.
// Эти методы безопасны для одновременного использования в разных go-рутинах.
type Context interface {
    // Done возвращает канал, который закрывается когда Context отменяется
    // или по таймауту.
    Done() <-chan struct{}

    // Err объясняет почему контекст был отменен, после того как закрылся канал Done.
    Err() error

    // Deadline возвращает время когда этот Context будет отменен.
    Deadline() (deadline time.Time, ok bool)

    // Value возвращает значение ассоциированное с ключем или nil.
    Value(key interface{}) interface{}
}
```

(Это выжимка из godoc. Там можно найти больше)

Метод Done возвращает канал который действует как сигнал отмены для функций запущенных от имени Context. Когда канал закрывается, функции должны завершить работу. Метод Err возвращает ошибку, которая объясняет, почему Context был отменен. В статье "Pipelines and Cancelation" идиома Done каналов объясняется более подробно.

У контекста нет метода Cancel по той же причине, по которой канал Done работает только на прием. Функция, которая принимает сигнал отмены, как правило, совсем не та что его отправляет. В частности, когда родительская операция стартует go-рутину для подоперации эта подоперация не должна иметь возможность отменить работу родителя. Вместо этого функция WithCancel(которая описана ниже) предоставляет возможность отменить новое значение Context.

Context безопасен для использования в множестве go-рутин. Мы можем передать один Context любому количеству go-рутин и отметить этот Context по сигналу любой из них.

Метод Deadline позволяет функциям определить должны ли они начать работу. Если осталось слишком мало времени, это уже может быть не целесообразным. Так же, можно использовать дедлайны для установки таймаута для операций ввода/вывода.

Value позволяет в Context пользоваться запросозависимыми данными. Данные должны быть безопасны для одновременного использованием множеством go-рутин.

### Производные контексты
Пакет context предоставляет возможность производить новые значения Context из существующего. Все эти значения образуют дерево и в случае отмены родительского Context все производные тоже будут отменены.

Background это корень для всего дерева `Context и он никогда не отменяется:

```
// Background возвращает пустой Context. Он никогда не будет отменен и не имеет deadline
// и значений. Background обычно используется в main, init, тестах и как верхний уровень
// Context для входящих запросов.
func Background() Context
WithCancel и WithTimeout возвращают полученное значение Context которое может быть отменено раньше чем родительский контекст. Context ассоциированный с входящим запросом, как правило, отменяется когда обработчик запроса завершает работу. WithCancel удобен для отмены излишних запросов, когда используется несколько реплик. WithTimeout удобно использовать для установки дедлайна в запросе к серверу:

// WithCancel возвращает копию родителя, в котором Done закрывается как только 
// parent.Done будет закрыт или контекст будет отменен.
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

// CancelFunc отменяет контекст.
type CancelFunc func()

// WithTimeout возвразщает копию родителя, в котором Done закрывается как только 
// parent.Done будет закрыт, контекст будет отменен или закончится таймаут. Новый дедлайн 
// контекста состоит из текущее время + таймаут и родительский дедлайн если такой имеется.
// Если таймер все еще работает, функция отмены релизит свои ресурсы.
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
WithValue предоставляет возможность привязать запросозависимые значения к контексту.

// WithValue возвращает копию родителя, в которой метод Value возвращает значение по ключу.
func WithValue(parent Context, key interface{}, val interface{}) Context
Самый лучший способ разобраться как работает пакет context это посмотреть на него в действии.
```

### Пример: Google Web Search
В качестве примера у нас HTTP сервер, который обрабатывает URLы в таком формате /search?q=golang&timeout=1s передает запрос из параметра "golang" в Google Web Search API и отображает результат. Параметр "timeout" говорит нашему серверу через какое время прекратить запрос.

Весь код разделен на три пакета:

server в этом пакете определена функция main и обработчики для /search.
userip предоставляет функции для получения клиентского IP из запроса и привязка его к Context.
google тут определена функция Search для отправки запроса в сервис Google.
Пакет server

Пакет server обрабатывает запросы вида /search?q=golang для получения первых нескольких результатов из Google в golang. В это пакете регистрируется handleSearch для обработки всех запросов к /search. Обработчик создает начальный Context с названием ctx и подготавливает его к закрытию после завершения работы обработчика. Если запрос включает параметр timeout, тогда Context будет автоматически отменен по завершению таймаута:

```
func handleSearch(w http.ResponseWriter, req *http.Request) {
    // ctx это Context для этого обработчика. Отмена контекста (вызов cancel)
    // закрывает канал ctx.Done что является сигналом отмены
    // для запросов запущенных в этом обработчике.
    var (
        ctx    context.Context
        cancel context.CancelFunc
    )
    timeout, err := time.ParseDuration(req.FormValue("timeout"))
    if err == nil {
        // В запросе есть параметр timeout, в таком случае создаем
        // контекст который будет отменен автоматически по
        // окончанию таймаута.
        ctx, cancel = context.WithTimeout(context.Background(), timeout)
    } else {
        ctx, cancel = context.WithCancel(context.Background())
    }
    defer cancel() // Отменяем ctx как только handleSearch закончит работу.
```

Обработчик извлекает поисковый запрос и клиентский IP адрес с помощью пакета userip. Клиентский IP нужен для выполнения запросов к АПИ Google. В результате handleSearch аттачит все это к ctx

```
    // Получаем посковый запрос.
    query := req.FormValue("q")
    if query == "" {
        http.Error(w, "no query", http.StatusBadRequest)
        return
    }

    // Записываем пользовательский IP в ctx для использования в других пакетах.
    userIP, err := userip.FromRequest(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    ctx = userip.NewContext(ctx, userIP)
```

Обработчик вызывает google.Search и передает ctx и query:

```
    // Запускаем Google поиск и выводим результаты.
    start := time.Now()
    results, err := google.Search(ctx, query)
    elapsed := time.Since(start)
```

Если поиск отработал нормально, тогда обработчик выводит результаты:
```
    if err := resultsTemplate.Execute(w, struct {
        Results          google.Results
        Timeout, Elapsed time.Duration
    }{
        Results: results,
        Timeout: timeout,
        Elapsed: elapsed,
    }); err != nil {
        log.Print(err)
        return
    }
```

### Пакет userip
Пакет userip предоставляет функции для извлчения IP адреса из запроса и привязки его к Context. Сам Context дает возможность мапинга ключ/значение, в котором ключи и значения имеют тип interface{}. Все типы ключей должны поддерживать сравнение и все типы значений должны быть безопасными для использования в нескольких go-рутинах. Такие пакеты как userip должны скрывать подробности реализации этого мапинга и предоставлять строго типизированный доступ к значениям в Context.
Для избежания коллизий с ключами, в пакете userip определена константа key которая используется как ключ для получения значения из контекста:

```
// Тип key не экспортируемый для предотвращения коллизий к другими ключами определенными
// в других пакетах.

type key int

// userIPkey это ключ контекста для клиентского IP. Его значение равно нулю.
// Если в этом пакете определить другие ключи, они могут иметь различные
// целочисленные значения.
const userIPKey key = 0
FromRequest извлекает значение userIP из http.Request:

func FromRequest(req *http.Request) (net.IP, error) {
    s := strings.SplitN(req.RemoteAddr, ":", 2)
    userIP := net.ParseIP(s[0])
    if userIP == nil {
        return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
    }
```

NewContext возвращает новый Context который содержит значение userIP:

```
func NewContext(ctx context.Context, userIP net.IP) context.Context {
    return context.WithValue(ctx, userIPKey, userIP)
}
```

FromContext извлекает userIP изContext:

```
func FromContext(ctx context.Context) (net.IP, bool) {
    // ctx.Value возвращает nil если ctx не имеет значение для этого ключа;
    // Приведение к типу net.IP возвращает ok=false для значения nil.
    userIP, ok := ctx.Value(userIPKey).(net.IP)
    return userIP, ok
}
```

### Пакет google

Функция google.Search отправляет HTTP запрос к Google Web Search API и парсит результат в формате JSON. Он принимает Context параметр ctx и возвращает результат немедленно если ctx.Done будет закрыт пока запрос выполняется.

Запрос Google Web Search API содержит поисковый запрос и пользовательский IP:

```
func Search(ctx context.Context, query string) (Results, error) {
    // Подготовка запроса к Google Search API.
    req, err := http.NewRequest("GET", "https://ajax.googleapis.com/ajax/services/search/web?v=1.0", nil)
    if err != nil {
        return nil, err
    }
    q := req.URL.Query()
    q.Set("q", query)
    // Если ctx содержит пользовательский IP, передаем его серверу.
    // Google API использует пользовательский IP чтобы отличить запрос с сервера
    // от пользовательского запроса.
    if userIP, ok := userip.FromContext(ctx); ok {
        q.Set("userip", userIP.String())
    }
    req.URL.RawQuery = q.Encode()
```

Search использует вспомогательную функцию httpDo для создания HTTP запроса и его отмены, если ctx.Done закроется во время обработки запроса или ответа. Search передает замыкание в httpDo, которое в качестве параметра принимает HTTP запрос:

```
    var results Results
    err = httpDo(ctx, req, func(resp *http.Response, err error) error {
        if err != nil {
            return err
        }
        defer resp.Body.Close()
        // Разбираем результат в формате JSON.
        // https://developers.google.com/web-search/docs/#fonje
        var data struct {
            ResponseData struct {
                Results []struct {
                    TitleNoFormatting string
                    URL               string
                }
            }
        }
        if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
            return err
        }
        for _, res := range data.ResponseData.Results {
            results = append(results, Result{Title: res.TitleNoFormatting, URL: res.URL})
        }
        return nil
    })
    // httpDo waits for the closure we provided to return, so it's safe to
    // read results here.
    return results, err
```

Функция httpDo запускает HTTP запрос и обрабатывает ответ в новой go-рутине. Запрос будет отменен если ctx.Done будет отменен до выхода из go-рутины:

```
func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
    // Запускаем HTTP запрос в go-рутине и передаем запрос в f.
    tr := &http.Transport{}
    client := &http.Client{Transport: tr}
    c := make(chan error, 1)
    go func() { c <- f(client.Do(req)) }()
    select {
    case <-ctx.Done():
        tr.CancelRequest(req)
        <-c // Wait for f to return.
        return ctx.Err()
    case err := <-c:
        return err
    }
}
```

### Адаптация кода для использования контекстов
Множество разный фреймворков предоставляют пакеты для работы с запросозависимыми значениями. Мы можем определить новую реализацию для интерфейса Context для связи между используемым фреймворком и кодом, который ожидает получить параметр типа Context.
Для примера, Gorilla пакета github.com/gorilla/context позволяет обработчику ассоциировать данные с входящим запросом предоставляя ключ/значение мапинг для HTTP запроса. В gorilla.go представлена реализация интерфейса Context в котором метод Value возвращает значения из HTTP запроса с помощью Gorilla'ского пакета.
Некоторые другие пакеты реализуют отмену по аналогии с нашим Context. Например пакет Tomb предоставляет функцию Kill, которая сигнализирует об отмене закрытием канала Dying. Tomb так же предоставляет функцию для ожидания всех go-рутин аналогичную sync.WaitGroup. В tomb.go мы представили реализацию Context который может быть отменен при отмене родительского Context или при условии, что Tomb будет "убит".

### Заключение
В Google мы требуем, чтобы программисты на Go первым аргументом использовали Context во всех функциях которые занимаются обработкой запроса и выдачей ответа. Это позволяет разным командам разработчиков взаимодействовать более продуктивно. Так же, такой подход предоставляет контроль над таймаутами и отменой, а так же гарантирует безопасность транзита таким критическим данным.

Фреймворки, которые хотят работать с Context должны предоставить свои реализации для связи между функциями, которые принимают Context в параметрах. Функции в их клиентских библиотеках должны уметь работать с Context. Устанавливая интерфейс для работы с запросозависимыми данными и возможностью отмены, Context упрощает разработку более универсальных, масштабируемых и всем понятных пакетов.
