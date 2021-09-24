### Запросы модулей (module query) в Golang

Команда go принимает **"запрос модуля" ("module query")** вместо версии модуля как в командной строке, так и в файле go.mod основного модуля. (После оценки запроса, найденного в файле go.mod основного модуля, команда go обновляет файл, чтобы заменить запрос его результатом.)

Полностью определенная семантическая версия, такая как "v1.2.3", оценивается как эта конкретная версия.

Префикс семантической версии, такой как "v1" или "v1.2", оценивается как последняя доступная тегированная (помеченная) версия с этим префиксом.

Семантическое сравнение версий, такое как "<v1.2.3" или ">=v1.5.6", оценивается как доступняя теговая версия, ближайшая к цели сравнения (последняя версия для < и <=, самая ранняя версия для > и >=).

Строка "latest" соответствует последней доступной версии с тегами или последней версии без тегов исходного репозитория.

Идентификатор ревизии для исходного репозитория, такой как префикс хеша коммита, тег ревизии или имя ветви, выбирает эту конкретную ревизию кода. Если ревизия также помечена семантической версией, запрос оценивается в эту семантическую версию. В противном случае запрос оценивается как псевдо-версия для коммита.

Все запросы предпочитают выбирать версии выпуска (релиза, release), чем пред-выпускные (пре-релизные, pre-release) версии. Например, "<v1.2.3" предпочтет вернуть "v1.2.2" вместо "v1.2.3-pre1", даже если "v1.2.3-pre1" ближе к цели сравнения.

Версии модуля, запрещенные инструкциями exclude в go.mod основного модуля, считаются недоступными и не могут быть возвращены запросами.

Например, все эти команды действительны:

```
go get github.com/gorilla/mux@latest    # @latest используется по умолчанию для 'go get'
go get github.com/gorilla/mux@v1.6.2    # записывает v1.6.2
go get github.com/gorilla/mux@e3702bed2 # записывает v1.6.2
go get github.com/gorilla/mux@c856192   # записывает v0.0.0-20180517173623-c85619274f5d
go get github.com/gorilla/mux@master    # записывает текущее значение ветки master
```