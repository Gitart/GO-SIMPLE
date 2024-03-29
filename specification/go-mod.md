### Команды go: go mod, обслуживание модуля

**go mod** обеспечивает доступ к операциям над модулями.

Обратите внимание, что поддержка модулей встроена во все команды go, а не только в 'go mod'. Например, ежедневное добавление, удаление, обновление и понижение зависимостей должны выполняться с помощью 'go get'. См. 'go help modules' для обзора функциональности модуля.

Использование:

```
go mod <command> [arguments]

```

Команды (command):

```
download    скачать модули в локальный кеш
edit        редактировать go.mod из инструментов или скриптов
graph       напечатать граф требований модуля
init        инициализировать новый модуль в текущем каталоге
tidy        добавить отсутствующие и удалить неиспользуемые модули
vendor      делает вендорную копию зависимостей
verify      проверить зависимости ожидаемого содержания
why         объяснять, зачем нужны пакеты или модули

```

Используйте "go help mod <command>" для получения дополнительной информации о команде.
