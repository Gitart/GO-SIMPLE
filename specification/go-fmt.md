### Команды go: go fmt, gofmt (переформатировать) источники (исходные файлы) пакетов

Использование:

```
go fmt [-n] [-x] [packages]

```

Fmt запускает команду 'gofmt -l -w' для пакетов, названных путями импорта. Он печатает имена файлов, которые были изменены.

Для получения дополнительной информации о gofmt см. "go doc cmd/gofmt".

Флаг -n печатает команды, которые будут выполнены.

Флаг -x печатает команды по мере их выполнения.

Чтобы запустить gofmt с конкретными опциями, запустите сам gofm
