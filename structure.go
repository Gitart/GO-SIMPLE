package main

import (
    "fmt"
    "log"
    "net/http"
    "sort"
    "strconv"
    "strings"
    // "net/http/httptrace"
)

// transport is an http.RoundTripper that keeps track of the in-flight
// request and implements hooks to report HTTP tracing events.
type statistics struct {
     numbers []float64
     mean      float64
     median    float64
}


func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean   = sum(numbers) / float64(len(numbers))
	stats.median = median(numbers)
	return stats
}


// Variables
var(
    anError    = `<p class="error">%s</p>`
    pageTop    = `<h1>Расчет таблицы</h1>`
    pageBottom = `<samll>Окончание</small>`
    form       = `<form action="/" method="POST"> 
                  <label for="numbers">Numbers (comma or space-separated):</label><br /> 
                  <input type="text" name="numbers" size="30"><br />
                  <input type="submit" value="Calculate">
                  </form>`
 )


// Здесь отсутствует проверка деления на нуль, поскольку сама ло-
// гика программы подразумевает, что getStats() будет вызываться,
// только когда имеется хотя бы одно число, поэтому, если в будущем
// логика работы программы изменится, она будет завершаться аварий-
// но в таких ситуациях. В критически важных программах, которые
// не должны завершаться при возникновении проблем, можно задей-
// ствовать функцию recover(), чтобы восстанавливать приложение в
// нормальное состояние после аварий и продолжать работу (§5.5).
func sum(numbers []float64) (total float64) {
	for _, x := range numbers {
	total += x
	}
	return total
}




// Для обхода всех чисел и вычисления суммы эта функция исполь-
// зует цикл for ...range (отбрасывающий значения их индексов).
// Благодаря тому что в зыке Go переменные, включая именованные
// возвращаемые значения, всегда инициализируются нулевыми значе-
// ниями , значение total изначально равно нулю.
func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
	result = (result + numbers[middle-1]) / 2
	}
	return result
}


// Программа statistics предоставляет доступ к единственной веб-
// странице на локальном компьютере. Ниже приводится функция
// main() программы:
func main() {
	http.HandleFunc("/", homePage)

    fmt.Println("Start Server on the port 9001")
	if err := http.ListenAndServe(":9001", nil); err != nil {
	   log.Fatal("failed to start server", err)
	}
}


// Строковая константа form содержит элемент <form>, включающий
// элементы <input> типа text и submit.
func homePage(writer http.ResponseWriter, request *http.Request) {
     err := request.ParseForm() // Должна вызываться перед записью в ответ
     fmt.Fprint(writer, pageTop, form)

     if err != nil {
        fmt.Fprintf(writer, anError, err)
     } else {
       if numbers, message, ok := processRequest(request); ok {
          stats := getStats(numbers)
          fmt.Fprint(writer, formatStats(stats))
     } else if message != "" {
         fmt.Fprintf(writer, anError, message)
     }
    }
       fmt.Fprint(writer, pageBottom)
}



// В случае успешного выполнения анализа (как и должно быть)
// вызывается функция processRequest(), извлекающая числа, вве-
// денные пользователем. Если все числа окажутся допустимыми,
// вызовом функции getStats(), которая была показана выше, вы-
// числяются и выводятся в форматированном виде статистические
// характеристики; иначе будет выведено сообщение об ошибке, ес-
// ли оно имеется. (Когда форма отображается первый раз, в ней
// нет чисел, и никаких ошибок пока не произошло, в этом случае
// переменная ok хранит значение false, а переменная message – пу-
// стую строку.) И в конце функции выводится строковая константа
// pageBottom (здесь не показана), закрывающая теги <body> и <html>.
func processRequest(request *http.Request) ([]float64, string, bool) {
	var numbers []float64
	if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
	text := strings.Replace(slice[0], ",", " ", -1)
	for _, field := range strings.Fields(text) {
	if x, err := strconv.ParseFloat(field, 64); err != nil {
	return numbers, "’" + field + "’ is invalid", false
	} else {
	numbers = append(numbers, x)
	}
	}
	}
	if len(numbers) == 0 {
	return numbers, "", false // при первом отображении данные отсутствуют
	}
	return numbers, "", true
}



// Если функция не завершилась внутри цикла (встретив недопу-
// стимое число), она вернет числа, пустое сообщение об ошибке и
// true, в противном случае, если числа отсутствуют (при первом ото-
// бражении формы), возвращается false.
func formatStats(stats statistics) string {
return fmt.Sprintf(`<table border="1">
                    <tr><th colspan="2">Results</th></tr>
                    <tr><td>Numbers</td><td>%v</td></tr>
                    <tr><td>Count</td><td>%d</td></tr>
                    <tr><td>Mean</td><td>%f</td></tr>
                    <tr><td>Median</td><td>%f</td></tr>
                    </table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median)
}
