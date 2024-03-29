# Что такое JWT  (веб-токен JSON)? Как работает JWT-аутентификация?  

## **Что такое JWT (веб-токен JSON)?**

JWT или JSON Web Token — это открытый стандарт, используемый для безопасного обмена информацией между двумя сторонами — клиентом и сервером. В большинстве случаев это закодированный JSON, содержащий набор утверждений и подпись. Обычно он используется в контексте других механизмов аутентификации, таких как OAuth, OpenID, для обмена информацией о пользователях. Это также популярный способ аутентификации/авторизации пользователей в микросервисной архитектуре.

 Аутентификация JWT — это механизм аутентификации без сохранения состояния на основе токенов. Он широко используется в качестве сеанса без сохранения состояния на стороне клиента, что означает, что серверу не нужно полностью полагаться на хранилище данных (или) базу данных для сохранения информации о сеансе.

JWT могут быть зашифрованы, но обычно они закодированы и подписаны. Мы сосредоточимся на подписанных JWT. Целью Signed JWT является не сокрытие данных, а обеспечение их подлинности. Именно поэтому настоятельно рекомендуется использовать HTTPS с подписанными JWT.

## **Структура JWT**

Структура JWT разделена на три части: заголовок, полезная нагрузка, подпись и отделены друг от друга точкой (.) и будут иметь следующую структуру:

![image](https://user-images.githubusercontent.com/3950155/234895068-8ba1566d-cddd-4c14-a81d-282ea139b5e4.png)
 

 

* **Заголовок**

Заголовок состоит из двух частей: 

1. Используемый алгоритм подписи
2. Тип токена, в данном случае чаще всего «JWT».

* **Полезная нагрузка**

Полезная нагрузка обычно содержит утверждения (атрибуты пользователя) и дополнительные данные, такие как эмитент, время истечения срока действия и аудитория. 

* **Подпись**

Обычно это хэш разделов заголовка и полезной нагрузки JWT. Алгоритм, который используется для создания подписи, — это тот же алгоритм, который упоминается в разделе заголовка JWT. Подпись используется для проверки того, что токен JWT не был изменен или изменен во время передачи. Его также можно использовать для проверки отправителя.

Раздел заголовка и полезной нагрузки JWT всегда кодируется Base64.

## **Как ****работает JWT-аутентификация? Когда использовать аутентификацию JWT?**

Когда дело доходит до аутентификации API и авторизации между серверами, веб-токен JSON ( JWT) является особенно полезной технологией. С точки зрения единого входа (SSO) это означает, что поставщик услуг может **получать достоверную информацию **от сервера проверки подлинности. 

Поделившись секретным ключом с поставщиком удостоверений, поставщик услуг может хэшировать часть полученного токена и сравнить ее с подписью токена. Теперь, если этот результат соответствует подписи, SP знает, что предоставленная информация поступила от другого объекта, владеющего ключом.

Следующий рабочий процесс объясняет процесс проверки подлинности:

![Рабочий процесс аутентификации JSON Web Token (JWT)](/wp-content/uploads/sites/19/2021/12/jwt-workflow.webp)

 

1. Вход пользователя с использованием имени пользователя и пароля.
2. Сервер аутентификации проверяет учетные данные и выдает JWT, подписанный с использованием закрытого ключа.
3. В дальнейшем клиент будет использовать JWT для доступа к защищенным ресурсам, передавая JWT в заголовке авторизации HTTP.
4. Затем сервер ресурсов проверяет подлинность токена с помощью открытого ключа.

Поставщик удостоверений создает JWT, удостоверяющий личность пользователя, а сервер ресурсов декодирует и проверяет подлинность токена с помощью открытого ключа.

Поскольку токены используются для авторизации и аутентификации в будущих запросах и вызовах API, необходимо проявлять большую осторожность, чтобы предотвратить проблемы с безопасностью. Эти токены не должны храниться в общедоступных местах, таких как локальное хранилище браузера или файлы cookie. Если других вариантов нет, то полезная нагрузка должна быть зашифрована.

## **Как JWT Single Sign-On (SSO) работает для нескольких веб-приложений**

Единый вход (SSO) позволяет аутентифицировать пользователей в ваших системах и впоследствии информирует приложения о том, что пользователь прошел аутентификацию. При успешной аутентификации создается и возвращается токен JWT, который может использоваться приложением для создания пользовательского сеанса. Маркер автоматически проверяется с помощью IDP при входе в систему. Затем пользователю разрешается доступ к приложениям без запроса на ввод отдельных учетных данных для входа.

Этот механизм безопасности позволяет приложениям доверять запросам на вход, которые они получают от систем. Кроме того, эти приложения будут предоставлять доступ только тем пользователям, которые были аутентифицированы вами/администратором, и, следовательно, единый вход (SSO) использует JSON Web Token (JWT) для обеспечения обмена данными аутентификации пользователя. Следует проявлять большую осторожность в отношении того, как этот токен хранится и управляется.
