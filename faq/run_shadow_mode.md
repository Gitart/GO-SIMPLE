# Как запустить приложение golang в скрытом режиме?
[Ответ](https://ru.stackoverflow.com/questions/530581/%d0%9a%d0%b0%d0%ba-%d0%b7%d0%b0%d0%bf%d1%83%d1%81%d1%82%d0%b8%d1%82%d1%8c-%d0%bf%d1%80%d0%b8%d0%bb%d0%be%d0%b6%d0%b5%d0%bd%d0%b8%d0%b5-golang-%d0%b2-%d1%81%d0%ba%d1%80%d1%8b%d1%82%d0%be%d0%bc-%d1%80%d0%b5%d0%b6%d0%b8%d0%bc%d0%b5)
```
go build -ldflags "-H windowsgui"  
```
