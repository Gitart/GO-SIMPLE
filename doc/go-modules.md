### Вспомогательные темы инструмента go: режимы сборки

```
go help buildmode

```

Команды 'go build' и 'go install' принимают аргумент \-buildmode, который указывает, какой тип объектного файла должен быть собран. В настоящее время поддерживаются следующие значения:

---

```
-buildmode=archive

```

Собрать перечисленные неосновные пакеты в .a файлы. Пакеты названные main игнорируются.

---

```
-buildmode=c-archive

```

Собрать указанный основной пакет, а также все пакеты, которые он импортирует, в C архивный файл. Единственными вызываемыми символами будут функции, которые экспортируются с использованием cgo //export комментариев. Требует чтобы ровно один основной (main) пакет был в переданном списке.

---

```
-buildmode=c-shared

```

Собрать указанный основной пакет, а также все пакеты, которые он импортирует, в C общую библиотеку (shared library). Единственными вызываемыми символами будут функции, которые экспортируются с использованием cgo //export комментариев. Требует чтобы ровно один основной (main) пакет был в переданном списке.

---

```
-buildmode=default

```

Перечисленные основные пакеты встроены в исполняемые файлы и перечисленные неосновные пакеты встроены в .a файлы (поведение по умолчанию).

---

```
-buildmode=shared

```

Объединить все перечисленные неосновные пакеты в одну общую (shared) библиотеку, которая будет использоваться при сборке с \-linkshared опцией. Пакеты с именем main игнорируются.

---

```
-buildmode=exe

```

Собрать перечисленные основные пакеты и все, что они импортируют в исполняемые файлы. Пакеты без имени main игнорируются.

---

```
-buildmode=pie

```

Собрать перечисленные основные пакеты и все, что они импортируют в позиционные независимые исполняемые файлы (PIE). Пакеты без имени main игнорируются.

---

```
-buildmode=plugin

```

Собрать перечисленные основные пакеты, а также все пакеты, которые они импортируют в плагин Go. Пакеты без имени main игнорируются.