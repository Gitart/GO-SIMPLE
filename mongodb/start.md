# Использование Go с MongoDB с помощью драйвера Go MongoDB

[MongoDB](https://www.digitalocean.com/community/tags/mongodb)[Go](https://www.digitalocean.com/community/tags/go)[Databases](https://www.digitalocean.com/community/tags/databases)

*    [![ayoisaiah](https://secure.gravatar.com/avatar/313abf9d9ca878dbda154ea2b91933ed?secure=true&d=identicon)](https://www.digitalocean.com/community/users/ayoisaiah)
*   By [Ayooluwa Isaiah](https://www.digitalocean.com/community/users/ayoisaiah)

    PostedMay 21, 2020 334 views

English Español Français Português Русский

Русский

*Автор выбрал фонд [Free Software Foundation](https://www.brightfunds.org/organizations/free-software-foundation-inc) для получения пожертвования в рамках программы [Write for DOnations](https://do.co/w4do-cta).*

### Введение

Много лет MongoDB зависела от решений, создаваемых общими усилиями, но затем ее [разработчики объявили](https://engineering.mongodb.com/post/considering-the-community-effects-of-introducing-an-official-golang-mongodb-driver), что работают над официальным драйвером для Go. В марте 2019 года этот новый драйвер достиг уровня готовности к эксплуатации [в версии 1.0.0](https://github.com/mongodb/mongo-go-driver/releases/tag/v1.0.0), и с тех пор он регулярно обновляется.

Как и другие официальные драйвера MongoDB, [драйвер Go](https://github.com/mongodb/mongo-go-driver) является неотъемлемой частью языка программирования Go и обеспечивает удобную возможность использования MongoDB в качестве решения для баз данных программы Go. Он полностью интегрирован с API MongoDB, а также имеет все функции запросов, индексирования и агрегирования API и другие продвинутые функции. В отличие от сторонних библиотек, он будет полностью поддерживаться инженерами MongoDB, поэтому вы сможете быть уверенными в его дальнейшей разработке и поддержке.

В этом обучающем руководстве вы начнете использовать официальный драйвер MongoDB Go. Вы установите драйвер, подключитесь к базе данных MongoDB и выполните несколько операций CRUD. В процессе вы создадите программу диспетчера задач для управления задачами с помощью командной строки.

## Предварительные требования

Для этого обучающего руководства вам потребуется следующее:

*   Go, установленный на вашем компьютере, и рабочее пространство Go, настроенное в соответствии с разделом [Установка Go и настройка локальной среды программирования](https://www.digitalocean.com/community/tutorial_series/how-to-install-and-set-up-a-local-programming-environment-for-go). В этом обучающем руководстве проект будет называться `tasker`. Вам потребуется Go v1.11 или выше, установленный на компьютере с активированными модулями Go.
*   MongoDB, установленный для вашей операционной системы в соответствии с разделом [Установка MongoDB](https://www.digitalocean.com/community/tutorials/how-to-install-mongodb-on-ubuntu-18-04). MongoDB 2.6 или выше — минимальная версия, поддерживаемая драйвером MongoDB Go.

Если вы используете Go v1.11 или 1.12, то убедитесь, что модули Go активированы, присвоив переменной среды `GO111MODULE` значение `on` — так, как показано ниже:

```
export GO111MODULE="on"

```

Дополнительную информацию об активации переменных среды можно найти в этом обучающем руководстве: [Чтение и установка переменных сред и переменных оболочки](https://www.digitalocean.com/community/tutorials/how-to-read-and-set-environmental-and-shell-variables-on-a-linux-vps#setting-environmental-variables-at-login).

Команды и код, указанные в этом руководстве, были протестированы в Go версии v1.14.1 и MongoDB версии v3.6.3.

## Шаг 1 — Установка драйвера MongoDB Go

На этом шаге вы установите пакет драйвера Go для MongoDB и импортируете его в ваш проект. Также вы подключитесь к вашей базе данных MongoDB и проверите состояние подключения.

Теперь создайте новый каталог для этого обучающего руководства в файловой системе:

```
mkdir tasker

```

После настройки каталога проекта измените его с помощью следующей команды:

```
cd tasker

```

Затем инициализируйте проект Go с файлом `go.mod`. Этот файл определяет требования проекта и блокирует зависимости от правильных версий:

```
go mod init

```

Если директория вашего проекта находится за пределами директории `$GOPATH`, то вам нужно указать путь импорта вашего модуля следующим образом:

```
go mod init github.com/<your_username>/tasker

```

Теперь файл `go.mod` будет выглядеть следующим образом:

go.mod

```
module github.com/<your_username>/tasker

go 1.14

```

Добавьте драйвер MongoDB Go в качестве зависимости для вашего проекта, используя следующую команду:

```
go get go.mongodb.org/mongo-driver

```

Результат будет выглядеть примерно следующим образом:

```
Outputgo: downloading go.mongodb.org/mongo-driver v1.3.2
go: go.mongodb.org/mongo-driver upgrade => v1.3.2

```

Теперь файл `go.mod` будет выглядеть следующим образом:

go.mod

```
module github.com/<your_username>/tasker

go 1.14

require go.mongodb.org/mongo-driver v1.3.1 // indirect

```

Затем создайте файл `main.go` в корневом каталоге проекта и откройте его в текстовом редакторе:

```
nano main.go

```

Для начала работы с драйвером импортируйте следующие пакеты в файл `main.go`:

main.go

```go
package main

import (
    "context"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
```

Copy

Здесь вы добавляете пакеты [`mongo`](https://godoc.org/go.mongodb.org/mongo-driver/mongo) и [`опций`](https://godoc.org/go.mongodb.org/mongo-driver/mongo/options), которые предоставляет драйвер MongoDB Go.

Затем, в зависимости от импорта, создайте новый клиент MongoDB и подключитесь к запущенному серверу MongoDB:

main.go

```go
. . .
var collection *mongo.Collection
var ctx = context.TODO()

func init() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }
}
```

Copy

`mongo.Connect()` принимает `Context` и объект `options.ClientOptions`, который используется для установки строки подключения и других параметров драйвера. Вы можете посмотреть [документацию по опциям пакетов](https://godoc.org/go.mongodb.org/mongo-driver/mongo/options), чтобы узнать, какие варианты конфигурации доступны.

[*Context*](https://golang.org/pkg/context/) — крайний срок, указывающий, когда операция должна остановиться и показать результат. Это помогает предотвращать снижение производительности работы в производственных системах, когда отдельные операции работают медленно. В этом коде вы передаете `context.TODO()`, чтобы указать, что вы не уверены, какой именно контекст нужно использовать сейчас, но планируете добавить контекст в дальнейшем.

Затем убедимся, что ваш сервер MongoDB был обнаружен и подключен для успешного использования метода `Ping`. Добавьте следующий код под функцией `init`:

main.go

```go
. . .
    log.Fatal(err)
  }

  err = client.Ping(ctx, nil)
  if err != nil {
    log.Fatal(err)
  }
}
```

Copy

Если есть какие\-либо ошибки при подключении к базе данных, то программа должна быть отключена, пока вы пытаетесь решить проблему, поскольку нет смысла поддерживать работу программы без активного подключения к базе данных.

Добавьте следующий код для создания базы данных:

main.go

```go
. . .
  err = client.Ping(ctx, nil)
  if err != nil {
    log.Fatal(err)
  }

  collection = client.Database("tasker").Collection("tasks")
}
```

Copy

Вы создаете базу данных `tasker` и набор `задач` для хранения задач, которые вы создадите. Также вы настроили команду `collection` в качестве переменной уровня пакетов, чтобы можно было использовать подключение к базе данных на разных этапах пакетов.

Сохраните и закройте файл.

На этот момент полный `main.go` выглядит следующим образом:

main.go

```go
package main

import (
    "context"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    collection = client.Database("tasker").Collection("tasks")
}
```

Copy

Вы выполнили настройку вашей программы для подключения к серверу MongoDB с помощью драйвера Go. На следующем шаге вы продолжите создание программы диспетчера задач.

## Шаг 2 — Создание программы интерфейса командной строки

На этом шаге вы установите хорошо известный пакет [`cli`](https://github.com/urfave/cli) (интерфейса командной строки) для оказания помощи в разработке программы диспетчера задач. Он предоставляет интерфейс, которым вы можете воспользоваться для быстрого создания современных инструментов командной строки. Например, этот пакет дает возможность задавать субкоманды для вашей программы, делающие работу с командной строкой более похожей на git.

Запустите следующую команду для добавления пакета в качестве зависимости:

```
go get github.com/urfave/cli/v2

```

Затем откройте файл `main.go` еще раз:

```
nano main.go

```

Добавьте следующий выделенный код в файл `main.go`:

main.go

```go
package main

import (
    "context"
    "log"
    "os"

    "github.com/urfave/cli/v2"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
. . .
```

Copy

Вы импортируете пакет `cli`, как упоминалось. Также вы импортируете пакет `os`, который вы будете использовать для передачи аргументов командной строки в вашу программу:

Добавьте следующий код после функции `init`, чтобы создать программу интерфейса командной строки и начать компиляцию при помощи кода:

main.go

```go
. . .
func main() {
    app := &cli.App{
        Name:     "tasker",
        Usage:    "A simple CLI program to manage your tasks",
        Commands: []*cli.Command{},
    }

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}
```

Copy

Этот фрагмент кода создает программу интерфейса командной строки под названием `tasker` и добавляет краткое описание использования, которое будет отображаться при запуске программы. Набор `командной строки` находится в том месте, где вы будете добавлять команды для вашей программы. Команда `Run` отображает список аргументов для подходящей команды.

Сохраните и закройте файл.

Вот команда, которая вам потребуется для создания и запуска программы:

```
go run main.go

```

Вывод должен выглядеть так:

```
OutputNAME:
   tasker - A simple CLI program to manage your tasks

USAGE:
   main [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)

```

Программа запускает и отображает текст справки, это полезно для изучения возможностей и способов использования программы.

На следующих шагах вы будете повышать эффективность вашей программы, добавляя субкоманды для управления задачами в MongoDB.

## Шаг 3 — Создание задачи

На этом шаге вы добавите субкоманду в вашу программу CLI при помощи пакета `cli`. В конце этого раздела вы сможете добавить новую задачу в базу данных MongoDB, используя новую команду `add` (добавить) в вашей программе CLI.

Начнем с открытия файла `main.go`:

```
nano main.go

```

Затем импортируем пакеты [`go.mongodb.org/mongo-driver/bson/primitive`](http://go.mongodb.org/mongo-driver/bson/primitive), [`time`](https://golang.org/pkg/time/) (времени) и [`errors`](https://golang.org/pkg/errors/) (ошибок):

main.go

```go
package main

import (
    "context"
    "errors"
    "log"
    "os"
    "time"

    "github.com/urfave/cli/v2"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
. . .
```

Copy

Затем создадим новую структуру для представления одной задачи базе данных и ее вставки непосредственно перед функцией `main`:

main.go

```go
. . .
type Task struct {
    ID        primitive.ObjectID `bson:"_id"`
    CreatedAt time.Time          `bson:"created_at"`
    UpdatedAt time.Time          `bson:"updated_at"`
    Text      string             `bson:"text"`
    Completed bool               `bson:"completed"`
}
. . .
```

Copy

Вы будете использовать пакет `primitive` для установки типа ID каждой задачи, поскольку MongoDB использует `ObjectID` для поля `_id`по умолчанию. Еще одно действие MongoDB по умолчанию — имя поля из строчных букв используется в качестве ключа для каждого экспортированного поля, когда ему присваивают серию, но это можно изменять с использованием тегов структуры `bson`.

Затем создадим функцию, которая получает экземпляр `Task` и сохраняет его в базе данных. Добавьте этот фрагмент кода после функции `main`:

main.go

```go
. . .
func createTask(task *Task) error {
    _, err := collection.InsertOne(ctx, task)
  return err
}
. . .
```

Copy

Метод `collection.InsertOne()` добавляет выделенную задачу в набор баз данных и возвращает ID документа, который был добавлен. Поскольку вам не потребуется этот идентификатор, вы удаляете его, присваивая его оператору подчеркивания.

Следующий шаг — добавление новой команды в программу диспетчера задач для создания новых задач. Назовем ее `add`:

main.go

```go
. . .
func main() {
    app := &cli.App{
        Name:  "tasker",
        Usage: "A simple CLI program to manage your tasks",
        Commands: []*cli.Command{
            {
                Name:    "add",
                Aliases: []string{"a"},
                Usage:   "add a task to the list",
                Action: func(c *cli.Context) error {
                    str := c.Args().First()
                    if str == "" {
                        return errors.New("Cannot add an empty task")
                    }

                    task := &Task{
                        ID:        primitive.NewObjectID(),
                        CreatedAt: time.Now(),
                        UpdatedAt: time.Now(),
                        Text:      str,
                        Completed: false,
                    }

                    return createTask(task)
                },
            },
        },
    }

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}
```

Copy

Каждая новая команда, добавляемая в вашу программу CLI, помещается в список `командной строки`. Каждый элемент включает имя, описание использования и действия. Это код, который будет запускаться при исполнении команд.

В этом коде вы берете первый аргумент для `add` (добавления) и используете его для установки свойства `Text` новой `задачи`, а также присваиваете соответствующие значения по умолчанию другим свойствам. Новая задача передается далее в `createTask`, которая вносит задачу в базу данных и возвращает `nil`, если все идет хорошо, после чего команда прекращается.

Сохраните и закройте файл.

Протестируйте ее, добавив несколько задач с помощью команды `add`. В случае успешного выполнения на экране не будет ошибок:

```
go run main.go add "Learn Go"
go run main.go add "Read a book"

```

Теперь, когда вы можете успешно добавлять задачи, давайте реализуем способ отображения всех задач, которые вы добавляли в базу данных.

## Шаг 4 — Перечисление всех задач

Перечисление документов в коллекции можно сделать с помощью метода `collection.Find()`, который предполагает фильтр, а также указатель значения, в которое можно расшифровать результат. Его значение — [Cursor](https://godoc.org/go.mongodb.org/mongo-driver/mongo#Cursor), которое предоставляет поток документов, которые можно обрабатывать пошагово, расшифровывая отдельные документы по очереди. Затем Cursor закрывается после завершения его использования.

Откройте файл `main.go`:

```
nano main.go

```

Обязательно импортируйте пакет [`bson`](http://go.mongodb.org/mongo-driver/bson):

main.go

```go
package main

import (
    "context"
    "errors"
    "log"
    "os"
    "time"

    "github.com/urfave/cli/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
. . .
```

Copy

Затем создайте следующие функции сразу после выполнения `createTask`:

main.go

```go
. . .
func getAll() ([]*Task, error) {
  // passing bson.D{{}} matches all documents in the collection
    filter := bson.D{{}}
    return filterTasks(filter)
}

func filterTasks(filter interface{}) ([]*Task, error) {
    // A slice of tasks for storing the decoded documents
    var tasks []*Task

    cur, err := collection.Find(ctx, filter)
    if err != nil {
        return tasks, err
    }

    for cur.Next(ctx) {
        var t Task
        err := cur.Decode(&t)
        if err != nil {
            return tasks, err
        }

        tasks = append(tasks, &t)
    }

    if err := cur.Err(); err != nil {
        return tasks, err
    }

  // once exhausted, close the cursor
    cur.Close(ctx)

    if len(tasks) == 0 {
        return tasks, mongo.ErrNoDocuments
    }

    return tasks, nil
}
```

Copy

[BSON (JSON в двоичном коде)](http://bsonspec.org/) — метод представления документов в базе данных MongoDB, а пакет `bson` это то, что помогает нам работать с объектами BSON в Go. Тип `bson.D`, используемый в функции `getAll()`, представляет документ BSON, а также используется в тех случаях, когда порядок свойств важен. Передавая `bson.D{}}` в качестве фильтра в `Tasks()`, вы указываете, что хотите сопоставить все документы в коллекции.

В функции `filterTasks()` вы пошагово выполняете Cursor, возвращаемый методом `collection.Find()`, и расшифровываете каждый документ в экземпляр `Task` (задачи). Затем каждая `задача` добавляется в список задач, созданных при запуске функции. После завершения использования Cursor будет закрыт, а задачи будут возвращены список `задач`.

Прежде чем вы создадите команду для перечисления всех задач, создадим функцию helper (помощник), которая принимает на себя долю `задач` и отображает стандартный вывод. Вы будете использовать пакет [`color`](https://github.com/gookit/color) для придания цвета выводимым результатам.

Чтобы начать использовать этот пакет, установите его с помощью следующей команды:

```
go get gopkg.in/gookit/color.v1

```

Вывод должен выглядеть так:

```
Outputgo: downloading gopkg.in/gookit/color.v1 v1.1.6
go: gopkg.in/gookit/color.v1 upgrade => v1.1.6

```

И импортируйте его в файл `main.go` вместе с пакетом `fmt`:

main.go

```go
package main

import (
    "context"
    "errors"
  "fmt"
    "log"
    "os"
    "time"

    "github.com/urfave/cli/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "gopkg.in/gookit/color.v1"
)
. . .
```

Copy

Затем создайте новую функцию `printTasks` после функции `main`:

main.go

```go
. . .
func printTasks(tasks []*Task) {
    for i, v := range tasks {
        if v.Completed {
            color.Green.Printf("%d: %s\n", i+1, v.Text)
        } else {
            color.Yellow.Printf("%d: %s\n", i+1, v.Text)
        }
    }
}
. . .
```

Copy

Функция `printTasks` берет на себя список `задач`, пошагово выполняет каждую задачу и отображает результат в стандартном выводе — зеленым цветом для обозначения завершенных задач и желтым — незавершенных задач.

Теперь добавьте следующие выделенные строки для создания новой команды `all` в списке `командной` строки. Эта команда отобразит все дополнительные задачи в стандартном выводе:

main.go

```go
. . .
func main() {
    app := &cli.App{
        Name:  "tasker",
        Usage: "A simple CLI program to manage your tasks",
        Commands: []*cli.Command{
            {
                Name:    "add",
                Aliases: []string{"a"},
                Usage:   "add a task to the list",
                Action: func(c *cli.Context) error {
                    str := c.Args().First()
                    if str == "" {
                        return errors.New("Cannot add an empty task")
                    }

                    task := &Task{
                        ID:        primitive.NewObjectID(),
                        CreatedAt: time.Now(),
                        UpdatedAt: time.Now(),
                        Text:      str,
                        Completed: false,
                    }

                    return createTask(task)
                },
            },
            {
                Name:    "all",
                Aliases: []string{"l"},
                Usage:   "list all tasks",
                Action: func(c *cli.Context) error {
                    tasks, err := getAll()
                    if err != nil {
                        if err == mongo.ErrNoDocuments {
                            fmt.Print("Nothing to see here.\nRun `add 'task'` to add a task")
                            return nil
                        }

                        return err
                    }

                    printTasks(tasks)
                    return nil
                },
            },
        },
    }

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}

. . .
```

Copy

Команда `all` извлекает все задачи, перечисленные в базе данных, и отображает их в стандартном выводе. Если задачи отсутствуют, вместо этого откроется подсказка с предложением добавить новую задачу.

Сохраните и закройте файл.

Создайте программу и запустите ее с помощью команды `all`:

```
go run main.go all

```

Она отобразит все задачи, которые вы добавили на данный момент:

```
Output1: Learn Go
2: Read a book

```

Теперь, когда вы можете просматривать все задачи базы данных, добавим возможность обозначать задачу как завершенную на следующем шаге.

## Шаг 5 — Завершение задачи

На этом шаге вы создадите новую субкоманду с именем `done`, которая позволит обозначать существующую задачу в базе данных как завершенную. Для обозначения задачи как завершенной вы можете использовать метод `collection.FindOneAndUpdate()`. Это позволяет вам найти документ в коллекции и обновить некоторые или все его свойства. В этом методе требуется фильтр для определения местонахождения документа и обновления документа для описания операции. Оба созданы при помощи ти`пов bs`on.D.

Начнем с открытия файла `main.go`:

```
nano main.go

```

Вставьте следующий фрагмент кода после функции `filterTasks`:

main.go

```go
. . .
func completeTask(text string) error {
    filter := bson.D{primitive.E{Key: "text", Value: text}}

    update := bson.D{primitive.E{Key: "$set", Value: bson.D{
        primitive.E{Key: "completed", Value: true},
    }}}

    t := &Task{}
    return collection.FindOneAndUpdate(ctx, filter, update).Decode(t)
}
. . .
```

Copy

Функция совпадает с первым документом, где текстовое свойство равняется `параметру` text. В документе `update` (обновления) указано, что свойству `completed` (завершено) можно присвоить значение `true` (истина). Если в операции `FindOneAndUpdate()` имеется ошибка, то она будет возвращена командой `completeTask()`. В противном случае будет возвращено `nil`.

Затем добавим новую команду `done` в вашу программу CLI (интерфейса командной строки), обозначающую задачу как завершенную:

main.go

```go
. . .
func main() {
    app := &cli.App{
        Name:  "tasker",
        Usage: "A simple CLI program to manage your tasks",
        Commands: []*cli.Command{
            {
                Name:    "add",
                Aliases: []string{"a"},
                Usage:   "add a task to the list",
                Action: func(c *cli.Context) error {
                    str := c.Args().First()
                    if str == "" {
                        return errors.New("Cannot add an empty task")
                    }

                    task := &Task{
                        ID:        primitive.NewObjectID(),
                        CreatedAt: time.Now(),
                        UpdatedAt: time.Now(),
                        Text:      str,
                        Completed: false,
                    }

                    return createTask(task)
                },
            },
            {
                Name:    "all",
                Aliases: []string{"l"},
                Usage:   "list all tasks",
                Action: func(c *cli.Context) error {
                    tasks, err := getAll()
                    if err != nil {
                        if err == mongo.ErrNoDocuments {
                            fmt.Print("Nothing to see here.\nRun `add 'task'` to add a task")
                            return nil
                        }

                        return err
                    }

                    printTasks(tasks)
                    return nil
                },
            },
            {
                Name:    "done",
                Aliases: []string{"d"},
                Usage:   "complete a task on the list",
                Action: func(c *cli.Context) error {
                    text := c.Args().First()
                    return completeTask(text)
                },
            },
        },
    }

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}

. . .
```

Copy

Вы используете аргумент, переданный команде `done` для нахождения первого документа, чье свойство `text` совпадает. В случае обнаружения свойству `completed` в документе будет привоено значение `true`.

Сохраните и закройте файл.

Затем запустите программу с помощью команды `done`:

```
go run main.go done "Learn Go"

```

Если вы снова будете использовать команду `all`, то увидите, что задача, которая была обозначена как завершенная, теперь отображается зеленым.

```
go run main.go all

```

![Скриншот вывода на терминал после завершения задачи](https://assets.digitalocean.com/articles/go_mongodb/go_mongo.png)

Иногда вы захотите просматривать только те задачи, которые еще не выполнялись. На следующем шаге мы добавим эту функцию.

## Шаг 6 — Отображение только незавершенных задач

На этом шаге вы будете включать код для извлечения незавершенных задач из базы данных с помощью драйвера MongoDB. Незавершенные задачи — те, чьему свойству `completed` присвоено значение `false`.

Давайте добавим новую функцию, которая извлекает задачи, которые еще не выполнены. Откройте файл `main.go`:

```
nano main.go

```

Добавьте этот фрагмент кода после функции `completeTask`:

main.go

```go
. . .
func getPending() ([]*Task, error) {
    filter := bson.D{
        primitive.E{Key: "completed", Value: false},
    }

    return filterTasks(filter)
}
. . .
```

Copy

Вы создаете фильтр при помощи пакетов `bson` и `primitive` из драйвера MongoDB, который будет сопоставлять документы, чьему свойству `completed` присвоено значение `false`. Затем вызывающему выдается список незавершенных задач.

Вместо создания новой команды для перечисления незавершенных задач давайте сделаем это действием по умолчанию при запуске программы без каких\-либо команд. Вы можете сделать это, добавив в программу свойство `Action` следующим образом:

main.go

```go
. . .
func main() {
    app := &cli.App{
        Name:  "tasker",
        Usage: "A simple CLI program to manage your tasks",
        Action: func(c *cli.Context) error {
            tasks, err := getPending()
            if err != nil {
                if err == mongo.ErrNoDocuments {
                    fmt.Print("Nothing to see here.\nRun `add 'task'` to add a task")
                    return nil
                }

                return err
            }

            printTasks(tasks)
            return nil
        },
        Commands: []*cli.Command{
            {
                Name:    "add",
                Aliases: []string{"a"},
                Usage:   "add a task to the list",
                Action: func(c *cli.Context) error {
                    str := c.Args().First()
                    if str == "" {
                        return errors.New("Cannot add an empty task")
                    }

                    task := &Task{
                        ID:        primitive.NewObjectID(),
                        CreatedAt: time.Now(),
                        UpdatedAt: time.Now(),
                        Text:      str,
                        Completed: false,
                    }

                    return createTask(task)
                },
            },
. . .
```

Copy

Свойство `Action` выполняет действие по умолчанию, когда программа выполняется без каких\-либо субкоманд. Здесь отображается логика перечисления незавершенных задач. Вызывается функция `getPending()`, и конечные задачи отображаются в стандартном выводе с помощью `printTasks()`. Если незавершенных задач нет, отображается подсказка, предлагающая пользователю добавить новую задачу с помощью команды `add`.

Сохраните и закройте файл.

Запуск программы сейчас без добавления каких\-либо команд отобразит все незавершенные задачи в базе данных:

```
go run main.go

```

Вывод должен выглядеть так:

```
Output1: Read a book

```

Теперь, когда вы можете указывать незавершенные задачи, добавим другую команду, которая позволяет просматривать только завершенные задачи.

## Шаг 7 — Отображение завершенных задач

На этом шаге вы добавите новую субкоманду `finished`, которая извлекает задачи из базы данных и отображает их на экране. Это подразумевает фильтрацию и возврат задач, чьему свойству `completed` присвоено значение `true`.

Откройте файл `main.go`:

```
nano main.go

```

Добавьте следующий код в конце файла:

main.go

```go
. . .
func getFinished() ([]*Task, error) {
    filter := bson.D{
        primitive.E{Key: "completed", Value: true},
    }

    return filterTasks(filter)
}
. . .
```

Copy

Как и в случае с функцией `getPending()`, вы добавили функцию `getFinished()`, которая возвращает список завершенных задач. В данном случае свойству `completed` в фильтре присвоено значение `true`, поэтому будут выведены только те документы, которые соответствуют этому состоянию.

Затем создайте команду `finished`, выводящую все завершенные задачи:

main.go

```go
. . .
func main() {
    app := &cli.App{
        Name:  "tasker",
        Usage: "A simple CLI program to manage your tasks",
        Action: func(c *cli.Context) error {
            tasks, err := getPending()
            if err != nil {
                if err == mongo.ErrNoDocuments {
                    fmt.Print("Nothing to see here.\nRun `add 'task'` to add a task")
                    return nil
                }

                return err
            }

            printTasks(tasks)
            return nil
        },
        Commands: []*cli.Command{
            {
                Name:    "add",
                Aliases: []string{"a"},
                Usage:   "add a task to the list",
                Action: func(c *cli.Context) error {
                    str := c.Args().First()
                    if str == "" {
                        return errors.New("Cannot add an empty task")
                    }

                    task := &Task{
                        ID:        primitive.NewObjectID(),
                        CreatedAt: time.Now(),
                        UpdatedAt: time.Now(),
                        Text:      str,
                        Completed: false,
                    }

                    return createTask(task)
                },
            },
            {
                Name:    "all",
                Aliases: []string{"l"},
                Usage:   "list all tasks",
                Action: func(c *cli.Context) error {
                    tasks, err := getAll()
                    if err != nil {
                        if err == mongo.ErrNoDocuments {
                            fmt.Print("Nothing to see here.\nRun `add 'task'` to add a task")
                            return nil
                        }

                        return err
                    }

                    printTasks(tasks)
                    return nil
                },
            },
            {
                Name:    "done",
                Aliases: []string{"d"},
                Usage:   "complete a task on the list",
                Action: func(c *cli.Context) error {
                    text := c.Args().First()
                    return completeTask(text)
                },
            },
            {
                Name:    "finished",
                Aliases: []string{"f"},
                Usage:   "list completed tasks",
                Action: func(c *cli.Context) error {
                    tasks, err := getFinished()
                    if err != nil {
                        if err == mongo.ErrNoDocuments {
                            fmt.Print("Nothing to see here.\nRun `done 'task'` to complete a task")
                            return nil
                        }

                        return err
                    }

                    printTasks(tasks)
                    return nil
                },
            },
        }
    }

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}
. . .
```

Copy

Команда `finished` извлекает задачи, чьему свойству `completed` присвоено значение `true` с помощью функции `getFinished()`. Затем она передает их в функцию `printTasks`, чтобы они отображались в стандартном выводе.

Сохраните и закройте файл.

Запустите следующую команду:

```
go run main.go finished

```

Вывод должен выглядеть так:

```
Output1: Learn Go

```

На последнем шаге вы дадите пользователям возможность удалять задачи из базы данных.

## Шаг 8 — Удаление задачи

На этом шаге вы добавите новую субкоманду `delete` (удаление), чтобы пользователи могли удалять задачу из базы данных. Для удаления одной задачи вы будете использовать метод `collection.DeleteOne()` из драйвера MongoDB. Также он использует фильтр для сопоставления удаляемого документа.

Откройте файл `main.go` еще раз:

```
nano main.go

```

Добавьте функцию `deleteTask` для удаления задач из базы данных непосредственно после функции `getFinished`:

main.go

```go
. . .
func deleteTask(text string) error {
    filter := bson.D{primitive.E{Key: "text", Value: text}}

    res, err := collection.DeleteOne(ctx, filter)
    if err != nil {
        return err
    }

    if res.DeletedCount == 0 {
        return errors.New("No tasks were deleted")
    }

    return nil
}
. . .
```

Copy

В этом методе `deleteTask` имеется строковый аргумент, представляющий удаляемую задачу. Создается фильтр для сопоставления задачи, чьему свойству `text` присвоен строковый аргумент. Вы передаете фильтр в метод `DeleteOne()`, который сопоставляет элемент коллекции и удаляет его.

Вы можете проверить свойство `DeletedCount` по результату метода `DeleteOne`, чтобы подтвердить, были ли удалены какие\-либо документы. Если фильтр не сможет сопоставить удаляемый документ, то `DeletedCount` будет равно нулю, и вы сможете вывести ошибку.

Теперь добавьте новую команду `rm`, как выделено:

main.go

```go
. . .
func main() {
    app := &cli.App{
        Name:  "tasker",
        Usage: "A simple CLI program to manage your tasks",
        Action: func(c *cli.Context) error {
            tasks, err := getPending()
            if err != nil {
                if err == mongo.ErrNoDocuments {
                    fmt.Print("Nothing to see here.\nRun `add 'task'` to add a task")
                    return nil
                }

                return err
            }

            printTasks(tasks)
            return nil
        },
        Commands: []*cli.Command{
            {
                Name:    "add",
                Aliases: []string{"a"},
                Usage:   "add a task to the list",
                Action: func(c *cli.Context) error {
                    str := c.Args().First()
                    if str == "" {
                        return errors.New("Cannot add an empty task")
                    }

                    task := &Task{
                        ID:        primitive.NewObjectID(),
                        CreatedAt: time.Now(),
                        UpdatedAt: time.Now(),
                        Text:      str,
                        Completed: false,
                    }

                    return createTask(task)
                },
            },
            {
                Name:    "all",
                Aliases: []string{"l"},
                Usage:   "list all tasks",
                Action: func(c *cli.Context) error {
                    tasks, err := getAll()
                    if err != nil {
                        if err == mongo.ErrNoDocuments {
                            fmt.Print("Nothing to see here.\nRun `add 'task'` to add a task")
                            return nil
                        }

                        return err
                    }

                    printTasks(tasks)
                    return nil
                },
            },
            {
                Name:    "done",
                Aliases: []string{"d"},
                Usage:   "complete a task on the list",
                Action: func(c *cli.Context) error {
                    text := c.Args().First()
                    return completeTask(text)
                },
            },
            {
                Name:    "finished",
                Aliases: []string{"f"},
                Usage:   "list completed tasks",
                Action: func(c *cli.Context) error {
                    tasks, err := getFinished()
                    if err != nil {
                        if err == mongo.ErrNoDocuments {
                            fmt.Print("Nothing to see here.\nRun `done 'task'` to complete a task")
                            return nil
                        }

                        return err
                    }

                    printTasks(tasks)
                    return nil
                },
            },
            {
                Name:  "rm",
                Usage: "deletes a task on the list",
                Action: func(c *cli.Context) error {
                    text := c.Args().First()
                    err := deleteTask(text)
                    if err != nil {
                        return err
                    }

                    return nil
                },
            },
        }
    }

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}
. . .
```

Copy

Как и в случае со всеми остальными ранее добавленными субкомандами, команда `rm` использует первый аргумент для сопоставления задачи в базе данных и ее удаления.

Сохраните и закройте файл.

Вы можете указывать незавершенные задачи, запустив программу без передачи субкоманд:

```
go run main.go

```

```
Output1: Read a book

```

Запуск субкоманды `rm` в задаче `«Читать книгу»` удалит ее из базы данных:

```
go run main.go rm "Read a book"

```

Если вы снова укажете все незавершенные задачи, то увидите, что задача `«Читать книгу»` больше не отображается, а вместо этого отображается подсказка, предлагающая создать новую задачу:

```
go run main.go

```

```
OutputNothing to see here
Run `add 'task'` to add a task

```

На этом шаге вы добавили функцию для удаления задач из базы данных.

## Заключение

Вы успешно создали программу «диспетчер задач» интерфейса командной строки и заодно научились основам применения драйвера MongoDB Go.

Обязательно посмотрите полную документацию по драйверу MongoDB Go в [GoDoc](https://godoc.org/go.mongodb.org/mongo-driver), чтобы узнать больше о функциях этого драйвера. Документация, описывающая применение [агрегирования](https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.Aggregate) или [транзакций](https://godoc.org/go.mongodb.org/mongo-driver/mongo#Session), может представлять особый интерес для вас.

Окончательный код этого обучающего руководства можно найти в [репозитории GitHub](https://github.com/do-community/tasker).
