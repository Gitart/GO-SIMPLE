## Условные обозначения, сокращения и соглашения о коде :


       (Термины, соглашения и обозначения, принятые в документации)
 
 		Внутреннее соглашение о наименовании функций :
 		     SSS_FFF_Name
 
  	     SSS    - Сокращенный префикс
 		     FF     - Область применения
 		     Name   - Наименование функции
 
       Group :
             AP     - Ассортиментный план
       
 		Name Servise :
 		     Hof    - Head Office
 		     Mov    - Move
 		     Olp    - Olap
 		     Dsc    - Discount
 
 		♣ Agreement and designations (Prefix) :
 		     Sys_  - Cистемные для внутреннего использования
 		     Wrk_  - Рабочие функции общего характера
 		     Adm_  - Административные задачи
 		     Atr_  - Опредление атрибутов для параметров
 		     Rep_  - Отчеты, сводки, информационные бюлютени
 		     Ser_  - Сервисные для работы сервиса
 		     Set_  - Внесение информации
         Get_  - Получение информации
 		     Sec_  - Функции для обслуживания секретной информации (ключи, хеши, пароли, логины)
 		     Inf_  - Информационные, банеры, гаджеты, информеры, подсказки, всплывающие окна
 		     Log_  - Связанные с логированием операций
 		     Cal_  - Математические расчеты и операции
 		     Mon_  - Мониторинг системы
         Mel_  - Отправка на почту     
  	     Utl_  - Утилиты для обслуживания сервиса
 		     Tst_  - Тестовые функции котрые требуют последующего удаления и нужны только для проверки работоспособности другого функционала
 		     Tmp_  - Временные функции
 		     Man_  - Манеджер информации
 		     Rig_  - Управление правами пользователей и системой
         Lib_  - Библиотеки
         Cry_  - Крптографическаие функции
         Acc_  - Функции прав и разрешений
         Wdb_  - Работа с базой данных
         Ops_  - Взаимодействие с опреационной системой
         Oth_  - Прочие функции которые не попали в предыдущие категории
         Stk_  - Остатки
 
 ##  Time format for cookies : time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST")
 
 ###      HTTP verbs :  Example(X-HTTP-Method-Override: PATCH) :
            GET	    Get a resource or list of resources
            POST  	    Create a resource
            Get        List of resources using a more advanced query
            PUT        Create a resource if it doesn't exist or, if it does, update it
            PATCH	    Update a resource
            DELETE	    Delete a resource
 
 ###     Response codes :
 ####           Response	Notes
            200	    Success, and there is a response body.
            201	    Success, when creating resources. Some APIs return 200 when successfully creating a resource. Look at the docs for the API you're using to be sure.
            204	    Success, and there is no response body. For example, you'll get this when you delete a resource.
            400	    The parameters in the URL or in the request body aren't valid.
            403	    The authenticated user doesn't have permission to perform the operation.
            404	    The resource doesn't exist, or the authenticated user doesn't have permission to see that it exists.
            409	    There's a conflict between the request and the state of the data on the server.
                       For example, if you attempt to submit a pull request and there is already a pull request for the commits, the response code is 409.
 
 ###     Links for code HTTP :
            http://great-world.ru/kody-otvetov-servera-i-oshibki-http-200-301-404-302-500-503-550/
            http://computerlessons.ru/lessons/vds/errors.html
            https://yandex.ru/support/webmaster/error-dictionary/http-codes.xml
            https://ru.wikipedia.org/wiki/%D0%A1%D0%BF%D0%B8%D1%81%D0%BE%D0%BA_%D0%BA%D0%BE%D0%B4%D0%BE%D0%B2_%D1%81%D0%BE%D1%81%D1%82%D0%BE%D1%8F%D0%BD%D0%B8%D1%8F_HTTP
            http://www.restapitutorial.ru/httpstatuscodes.html
            http://unicode-table.com/ru/#control-character        Unicode
 
 ###     Cистемная информация :
 ```javascript
            r.db("rethinkdb").table("table_config")                                          - управление параметрами таблицами
            r.db("").tableCreate("Directory", {durability:"soft", primaryKey:"ID"});         - создание таблиц с параметрами
 ```
 
 ###     Обновление вресии для Sublime Package
            https://github.com/DisposaBoy/GoSublime
            
            
