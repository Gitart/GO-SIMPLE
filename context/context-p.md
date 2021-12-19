# Контекстное программирование в Go

by [Gigi Sayfan](https://tutsplus.com/authors/gigi-sayfan)Aug 23, 2017

Read Time:8 minsLanguages: English Español Bahasa Indonesia Pусский

[Go](https://code.tutsplus.com/ru/categories/go)[Programming Fundamentals](https://code.tutsplus.com/ru/categories/programming-fundamentals)

Russian (Pусский) translation by [Masha Kolesnikova](https://tutsplus.com/authors/masha) (you can also [view the original English article](https://code.tutsplus.com/tutorials/context-based-programming-in-go--cms-29290?ec_unit=translation-info-language))

Программы Go, которые выполняют несколько параллельных вычислений в goroutines, должны управлять их временем жизни. Беглые программы могут попасть в бесконечные циклы, заблокировать другие ожидающие программы или просто занять слишком много времени. В идеале вы должны быть в состоянии отменить гоу-рутины или сделать так, чтобы они вышли по окончании.

Введение в контент-программирование. Go 1.7 представил пакет контекста, который предоставляет именно эти возможности, а также возможность связывать произвольные значения с контекстом, который путешествует с выполнением запросов и позволяет осуществлять внешнюю связь и передачу информации.

В этом уроке вы узнаете все тонкости контекста в Go, когда и как их использовать, и как не злоупотреблять ими.

## Кому нужен контекст?

Контекст - очень полезная абстракция. Она позволяет вам инкапсулировать информацию, которая не относится к основным вычислениям, такую как идентификатор запроса, токен авторизации и время ожидания. Есть несколько преимуществ такой инкапсуляции:

*   Она отделяет основные параметры вычислений от рабочих параметров.
*   Она кодифицирует общие эксплуатационные аспекты и способы их передачи через границы.
*   Она обеспечивает стандартный механизм добавления внеполосной информации без изменения сигнатур.

## Контекстный интерфейс

Вот весь интерфейс Context:

`type Context` `interface` `{`

`Deadline() (deadline time.Time, ok` `bool``)`

`Done() <-chan` `struct``{}`

`Err() error`

`Value(key` `interface``{})` `interface``{}`

 |

В следующих разделах объясняется назначение каждого метода.

Advertisement

### Метод Deadline()

Deadline возвращает время, когда работа, выполненная от имени этого контекста, должна быть отменена. Крайний срок возвращает `ok==false`, если не установлен крайний срок. Последовательные звонки в Deadline возвращают те же результаты.

### Метод Done()

Done() возвращает канал, который закрыт, когда работа, выполненная от имени этого контекста, должна быть отменена. Done может вернуть ноль, если этот контекст никогда не может быть отменен. Последовательные вызовы Done() возвращают одно и то же значение.

*   Функция context.WithCancel() обеспечивает закрытие канала Done при вызове cancel.
*   Функция context.WithDeadline() организует закрытие канала Done по истечении крайнего срока.
*   Функция context.WithTimeout() обеспечивает закрытие канала Done по истечении времени ожидания.

Done может быть использовано в операторах выбора:

`// Stream generates values with DoSomething and sends them`

`// to out until DoSomething returns an error or ctx.Done is`

`// closed.`

`func Stream(ctx context.Context, out chan<- Value) error {`

`for {`

`v, err := DoSomething(ctx)`

`if err != nil {`

`return err`

`}`

`select {`

`case <-ctx.Done():`

`return ctx.Err()`

`case out <- v:`

`}`

`}`

`}`

 |

См. [эту статью в блоге Go](https://blog.golang.org/pipelines) для получения дополнительных примеров того, как использовать канал Done для отмены.

Advertisement

### Метод Err()

Err() возвращает nil, пока открыт канал Done. Он возвращает `Canceled`, если контекст был отменен, или `DeadlineExceeded`, если истек крайний срок контекста или истекло время ожидания. После закрытия Done последующие вызовы Err() возвращают одно и то же значение. Вот определения:

|



 |

`// Canceled is the error returned by Context.Err when the`

`// context is canceled.`

`var Canceled = errors.New("context canceled")`

`// DeadlineExceeded is the error returned by Context.Err`

`// when the context's deadline passes.`

`var DeadlineExceeded error = deadlineExceededError{}`

 |

### Метод Value()

Value возвращает значение, связанное с этим контекстом для ключа, или ноль, если никакое значение не связано с ключом. Последовательные вызовы Value с одним и тем же ключом возвращают один и тот же результат.

Используйте значения контекста только для данных в области запроса, которые переходят процессы и границы API, но не для передачи необязательных параметров в функции.

Ключ идентифицирует конкретное значение в контексте. Функции, которые хотят хранить значения в Context, обычно выделяют ключ в глобальной переменной и используют этот ключ в качестве аргумента для context.WithValue() и Context.Value(). Ключ может быть любого типа, который поддерживает равенство.

## Контекстная область

Контексты имеют границы. Вы можете извлечь области из других областей, и родительская область не имеет доступа к значениям в производных областях, но производные области имеют доступ к значениям родительской области.

Контексты образуют иерархию. Вы начинаете с context.Background() или context.TODO(). Каждый раз, когда вы вызываете WithCancel(), WithDeadline() или WithTimeout(), вы создаете производный контекст и получаете функцию отмены. Важно то, что когда родительский контекст отменяется или истекает, все его производные контексты тоже.

Вы должны использовать context.Background() в функции main(), init() и тестах. Вы должны использовать context.TODO(), если вы не уверены, какой контекст использовать.

Обратите внимание, что Background и TODO *не* подлежат отмене.

Advertisement

Advertisement

## Сроки, тайм-ауты и отмены

Как вы помните, WithDeadline() и WithTimeout() возвращают контексты, которые отменяются автоматически, в то время как WithCancel() возвращает контекст и должны быть явно отменены. Все они возвращают функцию отмены, поэтому даже если тайм-аут/крайний срок еще не истек, вы все равно можете отменить любой производный контекст.

Давайте рассмотрим пример. Во-первых, здесь есть функция contextDemo() с именем и контекстом. Она работает в бесконечном цикле, выводя на консоль свое имя и крайний срок контекста, если таковой имеется. Затем он просто секунду спит.



 |

`package main`

`import (`

`"fmt"`

`"context"`

`"time"`

`)`

`func contextDemo(name string, ctx context.Context) {   `

`for {`

`if ok {`

`fmt.Println(name, "will expire at:", deadline)`

`} else {`

`fmt.Println(name, "has no deadline")`

`}`

`time.Sleep(time.Second)`

`}`

`}`

 |

Функция main создает три контекста:

*   timeoutContext с трехсекундным таймаутом
*   не истекающий cancelContext
*   deadlineContext, который получен из cancelContext, с крайним сроком четыре часа

Затем он запускает функцию contextDemo в виде трех программ. Все запускаются одновременно и каждую секунду печатают свои сообщения.

Затем основная функция ожидает отмены программы с тайм-аутом timeoutCancel чтением из канала Done() (блокируется, пока не закроется). Когда время ожидания истекает через три секунды, main() вызывает метод cancelFunc(), которая отменяет выполнение процедуры с помощью cancelContext, а также последнюю процедуру с производным контекстом предельного срока, равным четырем часам.



`func main() {`

`timeout := 3 * time.Second`

`deadline := time.Now().Add(4 * time.Hour)`

`timeOutContext, _ := context.WithTimeout(`

`context.Background(), timeout)`

`cancelContext, cancelFunc := context.WithCancel(`

`context.Background())`

`deadlineContext, _    := context.WithDeadline(`

`cancelContext, deadline)`

`go contextDemo("[timeoutContext]", timeOutContext)`

`go contextDemo("[cancelContext]", cancelContext)`

`go contextDemo("[deadlineContext]", deadlineContext)`

`// Wait for the timeout to expire`

`<- timeOutContext.Done()`

`// This will cancel the deadline context as well as its`

`// child - the cancelContext`

`fmt.Println("Cancelling the cancel context...")`

`cancelFunc()`

`<- cancelContext.Done()`

`fmt.Println("The cancel context has been cancelled...")`

`// Wait for both contexts to be cancelled`

`<- deadlineContext.Done()`

`fmt.Println("The deadline context has been cancelled...")`

`}`

 |

Вот вывод:



 |

`[cancelContext] has no deadline`
`[deadlineContext] will expire at: 2017-07-29 09:06:02.34260363`
`[timeoutContext] will expire at: 2017-07-29 05:06:05.342603759`
`[cancelContext] has no deadline`
`[timeoutContext] will expire at: 2017-07-29 05:06:05.342603759`
`[deadlineContext] will expire at: 2017-07-29 09:06:02.34260363`
`[cancelContext] has no deadline`
`[timeoutContext] will expire at: 2017-07-29 05:06:05.342603759`
`[deadlineContext] will expire at: 2017-07-29 09:06:02.34260363`
`Cancelling the cancel context...`

`The cancel context has been cancelled...`
`The deadline context has been cancelled...`


## Передача значений в контексте

Вы можете прикрепить значения к контексту, используя функцию WithValue(). Обратите внимание, что возвращается исходный контекст, а *не* производный контекст. Вы можете прочитать значения из контекста, используя метод Value(). Давайте изменим нашу демонстрационную функцию, чтобы получить ее имя из контекста вместо передачи ее в качестве параметра:

`func contextDemo(ctx context.Context) {`

`deadline, ok := ctx.Deadline()`

`name := ctx.Value("name")`

`for {`

`if ok {`

`fmt.Println(name, "will expire at:", deadline)`

`} else {`

`fmt.Println(name, "has no deadline")`

`}`

`time.Sleep(time.Second)`

`}`

`}`

 |

А давайте изменим функцию main, добавив имя через WithValue():

`go contextDemo(context.WithValue(`
`timeOutContext, "name", "[timeoutContext]"))`
`go contextDemo(context.WithValue(`
`cancelContext, "name", "[cancelContext]"))`
`go contextDemo(context.WithValue(`
`deadlineContext, "name", "[deadlineContext]"))`

Вывод остается прежним. См. раздел «Лучшие практики», где приведены рекомендации по правильному использованию значений контекста.

## Лучшие практики

Несколько лучших практик появились вокруг значений контекста:

*   Избегайте передачи аргументов функции в значениях контекста.
*   Функции, которые хотят хранить значения в контексте, обычно выделяют ключ в глобальной переменной.
*   Пакеты должны определять ключи как неэкспортируемый тип, чтобы избежать коллизий.
*   Пакеты, которые определяют ключ контекста, должны предоставлять средства доступа с типом для значений, хранящихся с использованием этого ключа.

Advertisement

## Контекст HTTP-запроса

Одним из наиболее полезных вариантов использования для контекстов является передача информации вместе с HTTP-запросом. Эта информация может включать идентификатор запроса, учетные данные для аутентификации и многое другое. В Go 1.7 стандартный пакет net/http воспользовался тем, что пакет контекста стал «стандартизированным», и добавил поддержку контекста непосредственно в объект запроса:


`func (r *Request) Context() context.Context`
`func (r *Request) WithContext(ctx context.Context) *Request`

Теперь можно прикрепить идентификатор запроса из заголовков к конечному обработчику стандартным способом. Функция обработчика WithRequestID() извлекает идентификатор запроса из заголовка «X-Request-ID» и генерирует новый контекст с идентификатором запроса из существующего контекста, который он использует. Затем он передает его следующему обработчику в цепочке. Открытая функция GetRequestID() предоставляет доступ к обработчикам, которые могут быть определены в других пакетах.

`const requestIDKey int = 0`
`func WithRequestID(next http.Handler) http.Handler {`
`return http.HandlerFunc(`
`func(rw http.ResponseWriter, req *http.Request) {`

`// Extract request ID from request header`
`reqID := req.Header.Get("X-Request-ID")`
`// Create new context from request context with`
`// the request ID`
`ctx := context.WithValue(`
`req.Context(), requestIDKey, reqID)`
`// Create new request with the new context`
`req = req.WithContext(ctx)`
`// Let the next handler in the chain take over.`
`next.ServeHTTP(rw, req)`
`}`
`)`
`}`

`func GetRequestID(ctx context.Context) string {`
`ctx.Value(requestIDKey).(string)`
`}`

`func Handle(rw http.ResponseWriter, req *http.Request) {`

`reqID := GetRequestID(req.Context())`

`...`

`}`

`func main() {`

`handler := WithRequestID(http.HandlerFunc(Handle))`

`http.ListenAndServe("/", handler)`

`}`

 |

## Заключение

Контекстное программирование обеспечивает стандартный и хорошо поддерживаемый способ решения двух распространенных проблем: управление временем жизни подпрограмм и передача внеполосной информации по цепочке функций.

Следуйте лучшим рекомендациям и используйте контексты в правильном контексте (посмотрите, что я там делал?), И ваш код значительно улучшится.
