# Recovery

```golang
// Программа выполняет целочисленное деление
// своего первого параметра на второй 
// и выводит результат.
func main() {
	defer func() {
		err := recover()
		if v, ok := err.(error); ok { // Обработка паники, соответствующей интерфейсу error
			fmt.Sprintf(os.Stderr, "Error %v \"%s\"\n", err, v.Error())
		} else if err != nil { 
			panic(err)  // Обработка неожиданных ошибок - повторный вызов паники.
		}
	}()
	a, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		panic(err)
	}
	b, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "%d / %d = %d\n", a, b, a/b)
}
```

В примере выше могут произойти ошибки при преобразовании аргументов программы в целые числа функцией strconv.ParseInt(). Также возможна паника при обращении к массиву os.Args при недостаточном количестве аргументов, либо при делении на нуль, если второй параметр окажется нулевым. При любой ошибочной ситуации генерируется паника, которая обрабатывается в вызове defer:

```
> divide 10 5
10 / 5 = 2

> divide 10 0
Error runtime.errorString "runtime error: integer divide by zero"

> divide 10.5 2
Error *strconv.NumError "strconv.ParseInt: parsing "10.5": invalid syntax"

> divide 10
Error runtime.errorString "runtime error: index out of range"
```
