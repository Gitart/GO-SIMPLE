### Вспомогательные темы инструмента go: вызовы между Go и C

```
go help c

```

Существует *два разных способа* вызова между кодом Go и кодом C/C++.

Первым является инструмент **cgo**, который является частью дистрибутива Go. Информацию о том, как его использовать, смотрите в документации по cgo (go doc cmd/cgo).

Вторая - это программа **SWIG**, которая является общим инструментом для взаимодействия между языками. Для получения информации о SWIG см. [http://swig.org](http://swig.org/). При запуске go build любой файл с расширением .swig будет передан SWIG. Любой файл с расширением .swigcxx будет передан SWIG с параметром -c++.

Когда используется cgo или SWIG, go build передает любые файлы .c, .m, .s или .S компилятору C, а любые файлы .cc, .cpp, .cxx компилятору C++. Переменные окружения CC или CXX могут быть установлены для определения, соответственно, используемого компилятора C или C++.
