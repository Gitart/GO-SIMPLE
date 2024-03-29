### Основной модуль и список сборки в Golang

**"Основной модуль" ("main module")** - это модуль, содержащий каталог, в котором выполняется команда go. Команда go находит корень модуля, ища go.mod в текущем каталоге, или родительский каталог текущего каталога, или родительский каталог родителя, и так далее.

Файл go.mod основного модуля определяет точный набор пакетов, доступных для использования командой go с помощью операторов require, replace и exclude. Модули зависимостей, обнаруживаемые с помощью следующих операторов require, также вносят вклад в определение этого набора пакетов, но только с помощью операторов require их файлов go.mod: любые операторы replace и exclude в модулях зависимостей игнорируются. Таким образом, операторы replace и exclude позволяют главному модулю полностью контролировать собственную сборку, не подвергаясь также полному контролю зависимостями.

Набор модулей, предоставляющих пакеты для сборки, называется "список сборки" ("build list"). Список сборки изначально содержит только основной модуль. Затем команда go рекурсивно добавляет в список точные версии модулей, требуемые для модулей, уже включенных в список, до тех пор, пока не останется ничего для добавления в список. Если в список добавлено несколько версий определенного модуля, то в конце для использования в сборке сохраняется только самая последняя версия (в соответствии с порядком семантической версии).

Команда 'go list' предоставляет информацию о главном модуле и списке сборки. Например:

```
go list -m              # распечатать путь основного модуля
go list -m -f={{.Dir}}  # распечатать корневой каталог основного модуля
go list -m all          # распечатать список сборки
```
