
## Protocol
https://developers.google.com/protocol-buffers/docs/gotutorial

# Что такое gRPC
**gRPC** (значение буквы g меняется с каждой новой версией (oreil.ly/IKCi3)) —технология межпроцессного взаимодействия, позволяющая соединять,вызывать, администрировать и отлаживать распределенные гетерогенные приложения в стиле, который по своей простоте мало чем отличается от вызова локальных функций.

При разработке gRPC-приложения первым делом нужно определить интерфейс сервисов. Данное определение содержит информацию о том, как
потребители могут обращаться к вашим сервисам, какие методы можно
вызывать удаленно, какие параметры и форматы сообщений следует применять при вызове этих методов и т. д. Язык, который используется в спецификации сервиса, называется языком описания интерфейсов (interface definition language, IDL).

На основе спецификации сервиса можно сгенерировать серверный код (или серверный каркас), который упрощает создание логики на серверной стороне за счет предоставления абстракций для низкоуровневого взаимодействия.
Вы также можете сгенерировать клиентский код (или клиентскую заглушку),инкапсулирующий низкоуровневые аспекты коммуникационного протокола в разных языках программирования. Методы, указанные в определении интерфейса сервиса, можно вызывать удаленно на клиентской стороне по примеру того, как вызываются локальные функции. Внутренний фреймворк
gRPC возьмет на себя все сложные аспекты, присущие соблюдению строгих контрактов между сервисами, сериализации данных, сетевых коммуникаций,аутентификации, контролю доступа, наблюдаемости и т. д.

Понять основополагающие концепции gRPC попробуем на реальном примере использования микросервисов, основанных на этой технологии. Допустим,мы занимаемся разработкой приложения для розничной торговли, состоящего из нескольких микросервисов, и хотим, чтобы один из них возвращал описание товаров, доступных для продажи, как показано на рис. 1.1 
![image](https://user-images.githubusercontent.com/3950155/152307216-1c9223ce-eeac-4f5b-871b-0d3768ed357c.png)


Сервис ProductInfo спроектирован так, чтобы к нему был доступ в сети по протоколу gRPC.
Определение сервиса находится в файле ProductInfo.proto, который используется для генерации кода как на серверной, так и на клиентской стороне. 

Теперь рассмотрим подробности данного взаимодействия. Первый шаг
при создании gRPC-сервиса — определение его интерфейса, содержащее
описание методов, которые он предоставляет, а также их входящих параметров и типов возвращаемых значений. Обсудим этот процесс более
подробно.
