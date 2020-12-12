### Команды go: go vet, сообщить о возможных ошибках в пакетах

Использование:

```
go vet [-n] [-x] [-vettool prog] [build flags] [vet flags] [packages]

```

**vet** запускает команду go vet для пакетов, названных путями импорта.

Для получения дополнительной информации о vet и его флагах см. go doc cmd/vet. Подробнее об указании пакетов см. go help packages. Список проверок и их флагов см. go tool vet help. Для получения подробной информации о конкретном контролере, таком как 'printf', смотрите go tool vet help printf.

Флаг \-n печатает команды, которые будут выполнены. Флаг \-x печатает команды по мере их выполнения.

Флаг \-vettool=prog выбирает другой инструмент анализа с альтернативными или дополнительными проверками. Например, теневой ('shadow') анализатор может быть построен и запущен с использованием этих команд:

```
go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow
go vet -vettool=$(which shadow)

```

Флаги сборки (build flags), поддерживаемые go vet, контролируют разрешение и выполнение пакета, например \-n, \-x, \-v, \-tags и \-toolexec. Для получения дополнительной информации об этих флагах см. [go help build](https://golang-blog.blogspot.com/2019/06/go-commands-go-build.html).

Смотрите также: [go fmt](https://golang-blog.blogspot.com/2019/06/go-commands-go-fmt.html), [go fix](https://golang-blog.blogspot.com/2019/06/go-commands-go-fix.html).
