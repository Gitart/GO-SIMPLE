### Пакет csv в Golang

Пакет **csv** читает и записывает файлы значений, разделенных запятыми (CSV). Есть много видов файлов CSV; этот пакет поддерживает формат, описанный в RFC 4180.

CSV\-файл содержит ноль или более записей из одного или нескольких полей на запись. Каждая запись отделяется символом новой строки. Финальная запись может опционально сопровождаться символом новой строки.

```
field1,field2,field3

```

Пустое пространство считается частью поля.

Символы возврата каретки до символов новой строки будут молча удалены.

Пустые строки игнорируются. Строка, содержащая только пробельные символы (исключая завершающий символ новой строки), не считается пустой строкой.

Поля, которые начинаются и заканчиваются символом кавычки ", называются квотированными полями (quoted\-field). Начальная и конечная кавычки не являются частью поля.

Источник:

```
normal string,"quoted-field"

```

получается в результате в полях

```
{`normal string`, `quoted-field`}

```

В кавычках символ кавычки, за которым следует второй символ кавычки, считается одинарной кавычкой.

```
"the ""word"" is true","a ""quoted-field"""

```

получается в результате

```
{`the "word" is true`, `a "quoted-field"`}

```

Новые строки и запятые могут быть включены в кавычки

```
"Multi-line
field","comma is ,"

```

получается в результате

```
{`Multi-line
field`, `comma is ,`}

```

#### csv.Reader

**Reader** читает записи из файла в формате CSV.

По возвращению NewReader, Reader ожидает ввод, соответствующий RFC 4180. Экспортированные поля можно изменить, чтобы настроить детали перед первым вызовом Read или ReadAll.

Reader преобразует все последовательности \\r\\n в своем входном файле в обычные \\n, в том числе в значения многострочных полей, чтобы возвращаемые данные не зависели от того, какое соглашение о конце строки использует входной файл.

```
package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "strings"
)

func main() {
    in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
    r := csv.NewReader(strings.NewReader(in))

    for {
        record, err := r.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println(record)
    }
}

```

Read читает одну запись (часть полей) из r. Если запись содержит неожиданное количество полей, Read возвращает запись вместе с ошибкой ErrFieldCount. За исключением этого случая, Read всегда возвращает либо ненулевую запись, либо ненулевую ошибку, но не обе. Если не осталось данных для чтения, Read возвращает nil, io.EOF. Если ReuseRecord имеет значение true, возвращенный срез может быть разделен между несколькими вызовами Read.

В этом примере показано, как можно настроить csv.Reader для обработки других типов файлов CSV.

```
package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "strings"
)

func main() {
    in := `first_name;last_name;username
"Rob";"Pike";rob
# lines beginning with a # character are ignored
Ken;Thompson;ken
"Robert";"Griesemer";"gri"
`
    r := csv.NewReader(strings.NewReader(in))
    r.Comma = ';'
    r.Comment = '#'

    records, err := r.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Print(records)
}

```

ReadAll читает все оставшиеся записи из r. Каждая запись представляет собой срез полей. Успешный вызов возвращает err == nil, а не err == io.EOF. Поскольку ReadAll определен для чтения до EOF, он не рассматривает конец файла как ошибку, о которой будет сообщено.

#### csv.Writer

```
type Writer struct {
    Comma   rune
    UseCRLF bool
}

```

**Writer** пишет записи с использованием CSV\-кодировки.

Как возвращается NewWriter, Writer записывает записи, оканчивающиеся символом новой строки, и использует ',' в качестве разделителя полей. Экспортируемые поля можно изменить, чтобы настроить детали перед первым вызовом Write или WriteAll.

Comma \- это разделитель полей.

Если UseCRLF имеет значение true, Writer заканчивает каждую выходную строку с \\r\\n вместо \\n.

Записи отдельных записей буферизируются. После того, как все данные записаны, клиент должен вызвать метод Flush, чтобы гарантировать, что все данные были перенаправлены в базовый io.Writer. Любые возникшие ошибки следует проверить, вызвав метод Error.

Write записывает одну запись CSV в w вместе со всеми необходимыми кавычками. Запись представляет собой срез строк, где каждая строка представляет собой одно поле. Записи буферизуются, поэтому в конце концов необходимо вызвать Flush, чтобы запись была записана в базовый io.Writer.

```
package main

import (
    "encoding/csv"
    "log"
    "os"
)

func main() {
    records := [][]string{
        {"first_name", "last_name", "username"},
        {"Rob", "Pike", "rob"},
        {"Ken", "Thompson", "ken"},
        {"Robert", "Griesemer", "gri"},
    }

    w := csv.NewWriter(os.Stdout)

    for _, record := range records {
        if err := w.Write(record); err != nil {
            log.Fatalln("error writing record to csv:", err)
        }
    }

    // Записываем любые буферизованные данные в подлежащий writer (стандартный вывод).
    w.Flush()

    if err := w.Error(); err != nil {
        log.Fatal(err)
    }
}

```

WriteAll записывает несколько CSV\-записей в w с помощью Write, а затем вызывает Flush, возвращая любую ошибку из Flush.

```
package main

import (
    "encoding/csv"
    "log"
    "os"
)

func main() {
    records := [][]string{
        {"first_name", "last_name", "username"},
        {"Rob", "Pike", "rob"},
        {"Ken", "Thompson", "ken"},
        {"Robert", "Griesemer", "gri"},
    }

    w := csv.NewWriter(os.Stdout)
    w.WriteAll(records) // вызывает Flush внутри

    if err := w.Error(); err != nil {
        log.Fatalln("error writing csv:", err)
    }
}
```
