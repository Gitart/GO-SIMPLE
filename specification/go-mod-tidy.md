### Команды go: go mod tidy, добавить отсутствующие и удалить неиспользуемые модули

Использование:

```
go mod tidy [-v]

```

**tidy** удостоверяется, что go.mod соответствует исходному коду в модуле. Он добавляет все недостающие модули, необходимые для построения пакетов и зависимостей текущего модуля, и удаляет неиспользуемые модули, которые не предоставляют никаких соответствующих пакетов. Он также добавляет все недостающие записи в go.sum и удаляет ненужные.

Флаг -v заставляет tidy печатать информацию об удаленных модулях в стандартный вывод ошибки (std err).
