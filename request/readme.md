## Параметры не очень понятны, поэтому на этот раз я запомню их как следует.

```
Метод #Method (POST, GET, PUT, DELETE, HEAD, PATCH, OPTIONS)
	 POST
 URL # имя API
	 /ttt
 Proto # тип протокола
	 HTTP/1.1
 ProtoMajor # Основной номер версии Proto
	 1
 ProtoMinor # Дополнительный номер версии Proto
	 1
 Заголовок # параметр заголовка, тип - карта [строка] строка

	  Content-Type ---- | ---- [application / json; charset = UTF-8] # Тип содержимого, укажите кодировку и набор символов
	  Accept-Encoding ---- | ---- [gzip, deflate] # Указать алгоритм сжатия
	  User-Agent ---- | ---- [Mozilla / 5.0 (Windows NT 10.0; WOW64) AppleWebKit / 537.36 (KHTML, как Gecko) Chrome / 69.0.3497.100 Safari / 537.36] #User Agent
	  Connection ---- | ---- [keep-alive] # тип соединения, keep-alive означает длительное соединение
	  Content-Length ---- | ---- [13] # длина тела
	  Принять ---- | ---- [application / json, text / plain, * / *] # тип текста
	  Происхождение ---- | ---- [chrome-extension: // ehafadccdcdedbhcbddihehiodgcddpl] # Начальная точка: браузер Chrome
	  Accept-Language ---- | ---- [zh-CN, zh; q = 0.9] #Language
Body
	  & {0xc4200e8f60 <nil> <nil> false true {0 0} false false false 0x637d00} #Main
 ContentLength # длина тела
	 13
TransferEncoding     #？？？
	 []
 Close # Следует ли закрывать соединение после завершения запроса
	 false
 Хост # Адрес назначения
	 140.143.91.148:8888
Form      
	 map[]
PostForm    
	 map[]
MultiparForm
	 <nil>
Trailer
	 map[]
RemoteAddr
	  106.2.0.253:3884 # Адрес источника
RequestURI
	 /ttt                            #API
TLS
	 <nil>
Response
	 <nil>
&{map[Method:POST] {"aaa":"bbb"}} 
```
