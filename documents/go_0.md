## Введение в Go

Go является си-подобным языком с присущими для си низко-уровневыми возможностями, такими как указатели. Одновременно go имеет возможности высоко-уровневых языков, такие, как юникодные строки, гарбаж-коллектор, высокоуровневая поддержка параллелизма. Стандартная библиотека самого go достаточно обширная.
Go родился в гугле в 2007 году. Одним из авторов языка в частности был Кен Томпсон (тот самый). В 2009 году Go был анонсирован под публичной лицензией. Go имеет открытую модель разработки, в его развитии принимают участие многие разработчики со всего мира. В 21-м веке появление go, пожалуй, одно из самых значительных событий в майнстриме.
Go сделан для разработки больших программ, при этом время компиляции минимально. Это сделано благодаря упрощению зависимостей в языке. Если есть программа, которая зависит от какого-то пакета, который в свою очередь зависит от другого пакета, то обычно для компиляции необходимы обьектные файлы обоих пакетов. В go все проще: для компиляции понадобится только первый пакет, а зависмости второго пакета будут закешированы в первом. Скорость компиляции является причиной того, что go может использоваться в области, где правят бал скриптовые языки - в разработке веб-приложений.
Go - статически типизированный компилируемый язык, в котором управление памятью выполняется гарбаж-коллектором. В Go также есть динамическая типизация, интроспеция. Go использует особую разновидность потоков - goroutines. которые автоматически проходят через балансировку между ядрами процессоров. В go есть две специальных коллекции - слайсы (slices) и мапы (maps). Слайсы - это динамические массивы, мапы - это словари. Эти коллекции практически покрывают все потребности языка. Что касается указателей, то у вас есть возможность перенести практически один в один си-шный код таких структур, как деревья.
Go нельзя назвать чисто процедурным языком, у него также есть поддержка ООП, которая радикально отличается от того, как она сделана в плюсах или жабе. В go вы не найдете дженериков. В нем нет препроцессора. Это язык, в котором много чего нет, и одновременно много что есть :-)
Go обычно входит в стандартный набор пакетов линуксовых дистрибутивов. Для того, чтобы загрузить и установить последнюю версию Go, посетите страницу golang.org/doc/install.html. На момент написания этой статьи выпущена стабильная версия 1.4. Компилятор имеет название gc, Есть его разновидности для различных железячных архитектур - армов, интелов и т.д. Этот компилятор также понимает си-шный код с помощью SWIG.

Теперь пришло время собрать первую программу. Для начала запустите команду

 > go version 
Первая программа - hello.go:

```golang
 package main
 
 import (
     "fmt"
     "os"
     "strings"
 )
 
 func main() {
     who := "World!"
     if len(os.Args) > 1 { 
         who = strings.Join(os.Args[1:], " ")
     }
     fmt.Println("Hello", who)
 }
 ```
 
Из каталога, в котором лежит исходник hello.go, запустите команду:   
 > go build    

Будет скомпилирован исполняемый файл, который по умолчанию будет носить то имя каталога, в котором лежит hello.go. В тексте программы нет точек с запятой и запятых :-) Функции и методы используют ключевое слово func. Оператор := декларирует и инициализирует переменную, при этом нет нужды указывать ее тип. Оператор = присваивает значение уже созданной переменной who. В go имеет значение, с какого символа начинается идентификатор переменной - с заглавного или прописного. В первом случае переменная становится public и доступна в любом другом пакете. Во втором случае переменная становится private и доступна только в текущем пакете. Это также справедливо для полей, обявленных внутри структуры, и для названий функций.
Следующая программа читает число, введенный с командной строки как аргумент, и выводит это число в псевдографике:

```golang
 package main
 
 import (
     "fmt"
     "log"
     "os"
     "path/filepath"
 )
 
 func main() {
     if len(os.Args) == 1 {
         fmt.Printf("usage: %s \n", filepath.Base(os.Args[0]))
         os.Exit(1)
     }
 
     stringOfDigits := os.Args[1]
     for row := range bigDigits[0] {
         line := ""
         for column := range stringOfDigits {
             digit := stringOfDigits[column] - '0'
             if 0 <= digit && digit <= 9 {
                 line += bigDigits[digit][row] + "  "
             } else {
                 log.Fatal("invalid whole number")
             }
         }
         fmt.Println(line)
     }
 }
 
 var bigDigits = [][]string{
     {"  000  ",
      " 0   0 ",
      "0     0",
      "0     0",
      "0     0",
      " 0   0 ",
      "  000  "},
     {" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
     {" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
     {" 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 "},
     {"   4  ", "  44  ", " 4 4  ", "4  4  ", "444444", "   4  ",
         "   4  "},
     {"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},
     {" 666 ", "6    ", "6    ", "6666 ", "6   6", "6   6", " 666 "},
     {"77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    "},
     {" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
     {" 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9"},
 }
 ```
 
В go нет классов с присущим им наследованием. Зато есть возможность создать произвольный тип. В go есть неявная типизация (duck typing), дающая возможность не определять строго тип параметров, передаваемых в функции. В go есть структуры - struct, позволяющие создавать вложенные типы. Типы могут быть именованные и неименованные, первые могут иметь свои методы, вторые - нет. Интерфейс в go - это также набор методов, как и везде. Тип можно привязать более чем к одному интерфейсу. В следующем

```golang
 type Example struct {
   Val string
   count int
 }```
 
 
Рассмотрим следующий пример. В ней используется два исходника. Главный файл лежит в корне каталога:

```golang
 package main
 
 import (
 	"fmt"
 	"./stack"
 )
 
 func main() {
 	var haystack stack.Stack
 	haystack.Push("hay")
 	haystack.Push(-15)
 	haystack.Push([]string{"pin", "clip", "needle"})
 	haystack.Push(81.52)
 	for {
 		item, err := haystack.Pop()
 		if err != nil {
 			break
 		}
 		fmt.Println(item)
 	}
 
 }
 ```
 
В главном файле импортируется тип stack.Stack из второго файла и декларируется. Импортируемый пакет stack должен лежать в подкаталоге с тем же именем, что и сам пакет - stack. Цикл for идет без начальных условий - он неопределенный - это фича go. Выход из такого цикла происходит по break. Функции и методы в go могут возвращать множество значений. Второй исходник лежит в подкаталоге stack/stack.go:

```golang
 package stack
 
 import "errors"
 
 type Stack []interface{}
 
 func (stack *Stack) Pop() (interface{}, error) {
     theStack := *stack
     if len(theStack) == 0 {
         return nil, errors.New("can't Pop() an empty stack")
     }
     x := theStack[len(theStack)-1]
     *stack = theStack[:len(theStack)-1]
     return x, nil
 }
 
 func (stack *Stack) Push(x interface{}) {
     *stack = append(*stack, x)
 }
 
 func (stack Stack) Top() (interface{}, error) {
     if len(stack) == 0 {
         return nil, errors.New("can't Top() an empty stack")
     }
     return stack[len(stack)-1], nil
 }
 
 func (stack Stack) Cap() int {
     return cap(stack)
 }
 
 func (stack Stack) Len() int {
     return len(stack)
 }
 
 func (stack Stack) IsEmpty() bool {
     return len(stack) == 0
 }
 ```
 
type Stack []interface{} - это значит, что тип не фиксирован и может быть произволен, в данном случае это массив нефиксированного типа. Некоторые параметры в функциях имеют тип указателя, что говорит о том, что все изменения, произведенные с параметром, применятся к внешнему обьекту, указатель на который был передан в функцию. nil - это нулевой указатель, который ни на что не указывает.
В следующем примере мы разберем работу с файловой системой, мапы. передачу функций в качестве параметров. Программа будет читать и записывать текстовые файлы. В каталоге с исходником должны лежать три текстовых файла: первый - это словарь транскрипций - british-american.txt - в каждой строке пара слов c неправильной-правильной транскрипцией:

 behaviour behavior
 colour color
 favourite favorite
 humour humor
 labour labor
 neighbour neighbor
 centre center
 fibre fiber
 theatre theater
 litre liter
 metre meter
 analogue analog
 catalogue catalog
 dialogue dialog
 anaemia anemia
 ...

Второй файл - это исходный текст, в котором слова с неправильной транскрипцией нужно заменить на правильные - это input.txt:
 This is just some test text. It contains some of the British words that
 have different spellings in American English, such as behaviour, colour,
 favourite, labour, and neighbour.    
 
 In fact the differences between British and American spellings aren't
 that great. Some British words that end with "our" get "or" endings in
 American English, and similarly for "re" to "er" (for example "Centre"
 becomes "Center"), and "ogue" becomes "og" as with dialogue and
 catalogue.   
 
 There are also a few more obscure changes, such as "ence" to "ense", for
 example "defence" becomes "defense". And there are a few miscellaneous
 differences, such as for anaemia, haemorrhage, aluminium, cheque, kerb,
 tyre, sulphur, and manoeuvre.   
 
И третий файл - output.txt - пустой файл, куда будет записан результат. И текст самой программы:

```golang
 package main
 
 import (
     "bufio"
     "fmt"
     "io"
     "io/ioutil"
     "log"
     "os"
     "path/filepath"
     "regexp"
     "strings"
 )
 
 var britishAmerican = "british-american.txt"
 
 func init() {
     dir, _ := filepath.Split(os.Args[0])
     britishAmerican = filepath.Join(dir, britishAmerican)
 }
 
 func main() {
     inFilename, outFilename, err := filenamesFromCommandLine()
     if err != nil {
         fmt.Println(err)
         os.Exit(1)
     }
     inFile, outFile := os.Stdin, os.Stdout
     if inFilename != "" {
         if inFile, err = os.Open(inFilename); err != nil {
             log.Fatal(err)
         }
         defer inFile.Close()
     }
     if outFilename != "" {
         if outFile, err = os.Create(outFilename); err != nil {
             log.Fatal(err)
         }
         defer outFile.Close()
     }
 
     if err = americanise(inFile, outFile); err != nil {
         log.Fatal(err)
     }
 }
 
 func filenamesFromCommandLine() (inFilename, outFilename string,
     err error) {
     if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
         err = fmt.Errorf("usage: %s input.txt output.txt",
             filepath.Base(os.Args[0]))
         return "", "", err
     }
     if len(os.Args) > 1 {
         inFilename = os.Args[1]
         if len(os.Args) > 2 {
             outFilename = os.Args[2]
         }
     }
     if inFilename != "" && inFilename == outFilename {
         log.Fatal("won't overwrite the infile")
     }
     return inFilename, outFilename, nil
 }
 
 func americanise(inFile io.Reader, outFile io.Writer) (err error) {
     reader := bufio.NewReader(inFile)
     writer := bufio.NewWriter(outFile)
     defer func() {
         if err == nil {
             err = writer.Flush()
         }
     }()
 
     var replacer func(string) string
     if replacer, err = makeReplacerFunction(britishAmerican); err != nil {
         return err
     }
     wordRx := regexp.MustCompile("[A-Za-z]+")
     eof := false
     for !eof {
         var line string
         line, err = reader.ReadString('\n')
         if err == io.EOF {
             err = nil   // io.EOF isn't really an error
             eof = true  // this will end the loop at the next iteration
         } else if err != nil {
             return err  // finish immediately for real errors
         }
         line = wordRx.ReplaceAllStringFunc(line, replacer)
         if _, err = writer.WriteString(line); err != nil {
             return err
         }
     }
     return nil
 }
 
 func makeReplacerFunction(file string) (func(string) string, error) {
     rawBytes, err := ioutil.ReadFile(file)
     if err != nil {
         return nil, err
     }
     text := string(rawBytes)
 
     usForBritish := make(map[string]string)
     lines := strings.Split(text, "\n")
     for _, line := range lines {
         fields := strings.Fields(line)
         if len(fields) == 2 {
             usForBritish[fields[0]] = fields[1]
         }
     }
 
     return func(word string) string {
         if usWord, found := usForBritish[word]; found {
             return usWord
         }
         return word
     }, nil
 }
 ```
 
Здесь обьявляется анонимная функция, которая будет вызвана в момент окончания и выхода функции americanise.
```golang
     defer func() {
         if err == nil {
             err = writer.Flush()
         }
     }()
 ```
 
В программе используется пустой идентификатор в виде символа подчеркивания:
```golang
   if _, err = writer.WriteString(line); err != nil {
   ```
Регулярные выражения включены в стандартный пакет regexp.Regexp, из которого используется метод ReplaceAllStringFunc(). Исходный файл разбивается построчно, и к каждой строке применяется эта функция. Функция makeReplacerFunction() возвращает функцию. В программе используется словарь - map - для хранения первоначальных пар транскрипций.
В следующем примере мы рассмотрим многопоточность. У go в этом плане есть две фичи: goroutines и channels. Каналы позволяют делать коммуникацию между goroutines. Это избавляет нас от необходимости блокировок и критических секций. После запуска эта программа попросит набрать два числа - радиус и угол - после чего вычислит картезианские координаты. Здесь показано, как используются структуры. Структура - это тип данных, который может хранить другие данные. Тип хранимых данных может быть стандартный, это могут быть другие структуры, и это могут быть интерфейсы. Функция init() выполняется прежде main(), функций init() может быть несколько, они не могут быть вызваны явно. Текст программы:

```golang
 package main
 
 import (
     "bufio"
     "fmt"
     "math"
     "os"
     "runtime"
 )
 
 const result = "Polar radius=%.02f θ=%.02f° → Cartesian x=%.02f y=%.02f\n"
 
 var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90, " +
     "or %s to quit."
 
 type polar struct {
     radius float64
     θ      float64
 }
 
 type cartesian struct {
     x   float64
     y   float64
 }
 
 func init() {
     if runtime.GOOS == "windows" {
         prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
     } else { // Unix-like
         prompt = fmt.Sprintf(prompt, "Ctrl+D")
     }
 }
 
 func main() {
     questions := make(chan polar)
     defer close(questions)
     answers := createSolver(questions)
     defer close(answers)
     interact(questions, answers)
 }
 
 func createSolver(questions chan polar) chan cartesian {
     answers := make(chan cartesian)
     go func() {
         for {
             polarCoord := <-questions
             θ := polarCoord.θ * math.Pi / 180.0 // degrees to radians
             x := polarCoord.radius * math.Cos(θ)
             y := polarCoord.radius * math.Sin(θ)
             answers <- cartesian{x, y}
         }
     }()
     return answers
 }
 
 func interact(questions chan polar, answers chan cartesian) {
     reader := bufio.NewReader(os.Stdin)
     fmt.Println(prompt)
     for {
         fmt.Printf("Radius and angle: ")
         line, err := reader.ReadString('\n')
         if err != nil {
             break
         }
         var radius, θ float64
         if _, err := fmt.Sscanf(line, "%f %f", &radius, &θ); err != nil {
             fmt.Fprintln(os.Stderr, "invalid input")
             continue
         }
         questions <- polar{radius, θ}
         coord := <-answers
         fmt.Printf(result, radius, θ, coord.x, coord.y)
     }
     fmt.Println()
 }
 ```
 
Каналы в go похожи на юниксовые пайпы и имеют дву-направленный характер обмена, работают по принципу FIFO, Рассмотрим пример создания канала:
```golang
 messages := make(chan string, 10)
 ```
 
Мы создали канал для отсылки и получения строк. Буфер этого канала - 10 строк. Буфер может быть равен и нулю, если на том конце канала готовы принять сообщение. Теперь отошлем в канал 2 строки:
```golang
 messages <- "Leader"
 messages <- "Follower"
 ```
 
Получить из канала данные:
```golang
 message1 := <-messages
 message2 := <-messages
 ```
 
Каналам не нужны блокировки, синхронизация у них встроенная. Теперь посмотрим на функцию main() последнего примера:

```golang
 func main() {
     questions := make(chan polar)
     defer close(questions)
     answers := createSolver(questions)
     defer close(answers)
     interact(questions, answers)
 }
 ```
 
Сначала мы создаем канал, потом отложенную встроенную функцию для закрытия канала. Далее мы передаем канал в качестве параметра функции, которая возвращает другой канал, Затем мы передаем в следующую функцию два канала в качестве параметров Далее рассмотрим функцию createSolver:
```golang
 func createSolver(questions chan polar) chan cartesian {
   answers := make(chan cartesian)
   go func() {
     for {
 	polarCoord := <-questions 
 	θ := polarCoord.θ * math.Pi / 180.0 // degrees to radians
 	x := polarCoord.radius * math.Cos(θ)
 	y := polarCoord.radius * math.Sin(θ)
 	answers <- cartesian{x, y}
     }
   }()
   return answers
 }
 ```
 
В ней есть выражение go с телом анонимной функции, которое создает поток, или goroutine, где вычисляются картезианские координаты.
