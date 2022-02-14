## Generators

*   [dst](https://github.com/dave/dst) (у которого лучше разрешение импортируемых пакетов и привязка комментариев к узлам AST, чем у go/ast из stdlib).

*   [kit](https://github.com/go-kit/kit) (хороший toolkit для быстрой разработки в архитектуре микросервисов. Предлагает внятные, рациональные абстракции, методики и инструменты).

*   [jennifer](https://github.com/dave/jennifer) (полноценный кодогенератор. Но его функциональность достигнута ценой применения промежуточных абстракций, которые хлопотно обслуживать. Генерация из шаблонов text/template на деле оказалась удобней, хоть и менее универсальной, чем манипулирование непосредственно AST с использованием промежуточных абстракций. Писать, читать и править шаблоны проще).


## Other librrary
**https://github.com/asim/go-micro**
https://github.com/asim/go-micro/tree/master/examples
https://github.com/go-kit/kit

https://github.com/nytimes/gizmo
https://github.com/dave/jennifer
https://github.com/gocircuit/circuit
https://github.com/koding/kite
**https://github.com/rsms/gotalk **




## Rsa
openssl genrsa -out key.pem 2048
openssl rsa -in key.pem -pubout > key_pub.pem



