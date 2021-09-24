### Пути удаленного импорта в Golang

Некоторые пути импорта описывают, как получить исходный код для пакета, используя систему контроля версий.

Несколько сайтов с общедоступным кодом имеют специальный синтаксис:

```
Bitbucket (Git, Mercurial)

   import "bitbucket.org/user/project"
   import "bitbucket.org/user/project/sub/directory"

GitHub (Git)

   import "github.com/user/project"
   import "github.com/user/project/sub/directory"

Launchpad (Bazaar)

   import "launchpad.net/project"
   import "launchpad.net/project/series"
   import "launchpad.net/project/series/sub/directory"

   import "launchpad.net/~user/project/branch"
   import "launchpad.net/~user/project/branch/sub/directory"

IBM DevOps Services (Git)

   import "hub.jazz.net/git/user/project"
   import "hub.jazz.net/git/user/project/sub/directory"

```

Для кода, размещенного на других серверах, пути импорта могут быть либо квалифицированы по типу управления версиями, либо инструмент go может динамически извлекать путь импорта через https/http и обнаруживать, где находится код, из тега <meta> в HTML.

Чтобы объявить местоположение кода, путь импорта формы

```
repository.vcs/path

```

указывает данный репозиторий с суффиксом .vcs или без него, используя именованную систему управления версиями, а затем путь внутри этого репозитория. Поддерживаемые системы контроля версий:

```
Bazaar      .bzr
Fossil      .fossil
Git         .git
Mercurial   .hg
Subversion  .svn

```

Например,

```
import "example.org/user/foo.hg"

```

обозначает корневой каталог репозитория Mercurial по адресу example.org/user/foo или foo.hg, и

```
import "example.org/repo.git/foo/bar"

```

обозначает каталог foo/bar репозитория Git по адресу example.org/repo или repo.git.

Когда система контроля версий поддерживает несколько протоколов, каждый загружается по очереди. Например, загрузка Git пробует https://, затем git+ssh://.

По умолчанию загрузка ограничена известными безопасными протоколами (например, https, ssh). Чтобы переопределить этот параметр для загрузок Git, можно установить переменную окружения GIT\_ALLOW\_PROTOCOL (подробнее см. [go help environment](https://golang-blog.blogspot.com/2019/07/go-help-environment.html)).

Если путь импорта не является известным сайтом размещения кода, а также отсутствует квалификатор управления версиями, инструмент go пытается извлечь импорт через https/http и ищет тег <meta> в <head> HTML-документа.

Метатег имеет вид:

```
<meta name="go-import" content="import-prefix vcs repo-root">

```

import-prefix - это путь импорта, соответствующий корню хранилища. Это должен быть префикс или точное совпадение пакета, извлекаемого с помощью "go get". Если это не точное совпадение, то в префиксе делается еще один http-запрос для проверки совпадения тегов <meta>.

Метатег должен появляться в файле как можно раньше. В частности, он должен появляться перед любым необработанным JavaScript или CSS, чтобы избежать путаницы с ограниченным парсером команды go.

vcs является одним из "bzr", "fossil", "git", "hg", "svn".

repo-root является корнем системы контроля версий, содержащей схему и не содержащей квалификатор .vcs.

Например,

```
import "example.org/pkg/foo"

```

приведет к следующим запросам:

```
https://example.org/pkg/foo?go-get=1 (предпочтительно)
http://example.org/pkg/foo?go-get=1  (резерв, только при флаге -insecure)

```

Если эта страница содержит метатег

```
<meta name="go-import" content="example.org git https://code.org/r/p/exproj">

```

инструмент go проверит, что https://example.org/?go-get=1 содержит тот же метатег, а затем выполнит git clone https://code.org/r/p/exproj в GOPATH/src/example.org.

При использовании GOPATH загруженные пакеты записываются в первый каталог, указанный в переменной среды GOPATH. (См. [go help gopath-get](https://golang-blog.blogspot.com/2019/06/go-commands-go-get.html) и [go help gopath](https://golang-blog.blogspot.com/2019/07/go-help-gopath.html))

При использовании модулей загруженные пакеты сохраняются в кеше модуля. (Смотрите go help module-get и [go help goproxy](https://golang-blog.blogspot.com/2019/07/go-help-goproxy.html).)

При использовании модулей распознается дополнительный вариант метатега go-import, который является предпочтительным по сравнению с перечисленными системами контроля версий. Этот вариант использует "mod" в качестве vcs в значении content, как в:

```
<meta name="go-import" content="example.org mod https://code.org/moduleproxy">

```

Этот тег означает получение модулей с путями, начинающимися с example.org, из прокси модуля, доступного по адресу https://code.org/moduleproxy. Смотрите [go help goproxy](https://golang-blog.blogspot.com/2019/07/go-help-goproxy.html) для получения подробной информации о протоколе прокси.
