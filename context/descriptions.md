# Работа с контекстом в Go

![](https://dev-gang.ru/static/storage/235369240164880473423971800620270660834.png)

https://www.youtube.com/watch?v=40lfXgkilHM&t=730s          
https://dev-gang.ru/article/rabota-s-kontekstom-v-go-tqvnsc2ysq/            


Когда у вас срыв, вызванный сочетанием выгорания и экзистенциальной боли, вас раздражает, что ваши беспокойные крики в пустоту остаются без ответа? Что ж, я не могу помочь с этим, но я могу предложить несколько методов для тайм-аута вызовов внешних или внутренних служб. Я проводил исследования и экспериментировал с некоторыми стандартными библиотеками в Go, и одна из них, на мой взгляд, наиболее полезна - это библиотека контекста. Эта небольшая библиотека, используемая для получения некоторого контроля над системой, которая может работать медленно по какой-либо причине, или для обеспечения определенного уровня качества для вызовов служб, является стандартом не зря. Для любой системы производственного уровня, чтобы поддерживать хороший контроль потока, понадобится библиотека контекста.

Библиотека контекста, созданная [Самиром Аджмани](https://twitter.com/Sajma) и [представленная в 2014 году](https://vimeo.com/115309491), стала стандартной библиотекой с Go 1.7. Если вы просмотрели исходный код библиотеки Go, вы можете найти множество примеров, требующих передачи [контекста](https://github.com/mongodb/mongo-go-driver/blob/v1.4.0/mongo/client.go#L96). Это только один, который я использовал недавно. Контекст - это крайний срок, который вы можете передать в запущенный процесс в своем коде. Этот крайний срок может указывать на то, что процесс должен прекратить работу и вернуться после выполнения условия. Это становится полезным при обращении к внешним API, базам данных, как показано выше, или системным командам.

Далее предполагается, что читатель знает о горутинах и каналах и о том, как они работают вместе. Я собираюсь глубоко погрузиться в параллелизм после того, как напишу о контексте, поскольку библиотека контекста является частью параллелизма. На данный момент горутины - это легкие потоки, которые можно запускать для процессов, а каналы - это конвейеры, используемые для передачи данных между этими новыми процессами.

### Контекстный интерфейс

Библиотека контекста определяет новый интерфейс под названием Context. В интерфейсе контекста есть несколько интересных полей, представленных ниже:

![](https://dev-gang.ru/static/storage/119562169521416666369732777404860266008.png)

Поле Deadline возвращает ожидаемое время завершения работы и указывает, когда контекст должен быть отменен.

Поле Done - это канал, который закрывается, когда работа, выполненная для контекста, должна быть отменена. Эта операция может происходить асинхронно. Канал может вернуться как nil, если связанный контекст никогда не может быть отменен. Различные типы контекстов организуют отмену работы в зависимости от обстоятельств, в которые мы попадем.

Err вернет nil, пока не будет закрыто Done. После чего Err либо вернет Cancelled, если контекст был отменен, либо DealineExceeded, если крайний срок контекста прошел.

Поле «Значение» представляет собой интерфейс «ключ-значение», который будет возвращать значение, связанное с контекстом, как ключ или ноль, если не было связанного значения. Значения следует использовать с осторожностью, поскольку они предназначены не для передачи параметров в функцию, а для [процессов передачи данных в области запроса и границ API](https://github.com/golang/go/blob/master/src/context/context.go#L185).

### Контекст в контексте

При создании контекста в Go легко написать статический контекст для хранения и повторного использования. Насколько я могу судить из своего исследования, это не оптимальный способ работы с библиотекой контекста. Контекст должен принимать форму, необходимую для каждого использования. Он должен быть бесформенным, или, говоря словами [Брюса Ли, быть похожим на воду](https://youtu.be/cJMwBwFj5nQ). Ваш контекст должен течь через ваш код и развиваться в зависимости от потребности.

Из этого есть некоторые исключения. Для процессов более высокого уровня вы можете передать пустой контекст, если у вас еще нет контекста для перехода. Они могут работать как заполнители до рефакторинга.

### Context.Background

Функция «Фон» возвращает пустой контекст, отличный от нуля. Нет никакого связанного с этим крайнего срока и никакой отмены, о которой можно было бы говорить. Обычно это можно использовать в основной функции, для тестирования или для создания контекста верхнего уровня, который будет преобразован во что-то еще. Заглянув в исходный код, вы увидите, что в нем нет никакой логики, кроме возврата [пустого контекста](https://github.com/golang/go/blob/master/src/context/context.go#L208):

![](https://dev-gang.ru/static/storage/126605246398075898614492681014727470966.png)

QuickNote: Обычно контекст при объявлении называется ctx. Я видел это в большинстве реализаций контекста, поэтому, если вы сталкиваетесь с ctx в случайных местах в исходном коде, есть большая вероятность, что это относится к контексту.

### Context.TODO

Функция TODO делает то же самое. Она возвращает пустой ненулевой контекст. Это снова вариант использования функций более высокого уровня, которые могут еще не иметь доступной функции для их использования. Во многих случаях это будет использоваться в качестве заполнителя при расширении вашей программы для использования библиотеки контекста. Если вы проверили выступление Самира Аджмани о введении контекстной библиотеки во время рефакторинга своего кода в Google, они будут использовать контекст.TODO, чтобы начать вводить контекст в кодовую базу Google, ничего не нарушая.

QuickNote: Я также упомяну одну вещь: где-то по ходу дела было высказано предположение, что TODO будет совместим для использования в инструментах статического анализа для наблюдения за распространением контекста по программе. Насколько я могу судить, это мог быть случайный комментарий человека, написавшего примечания в исходном коде. Я искал последние пару дней и, [насколько я могу судить, такого инструмента еще не существует](https://go-review.googlesource.com/c/go/+/130876/). Я хотел бы изучить, как создать такой инструмент, но вместо этого я пойду посмотреть фильм.

![](https://dev-gang.ru/static/storage/154231308599201757253377144136078926988.png)

### Context.WithCancel

Допустим, я создаю сайт для обзора фильмов. Существует множество API, предназначенных для обслуживания информации о фильмах. Один из последних, с которыми я столкнулся, - это [Studio Ghibli API,](https://ghibliapi.herokuapp.com/#section/Studio-Ghibli-API) который представляет собой общедоступный API, из которого мы можем просто брать данные. Итак, для специального раздела сайта фильмов Studio Ghibli мы воспользуемся этим. Функция WithCancel возвращает копию родительского контекста, переданную в нее с новым каналом Done. Новый Done канал будет закрыт или когда вызывается функция отмены или когда родительский контекст в Done канал закрыт. Какое бы событие ни произошло раньше.

Ниже приведен пример в действии:

![](https://dev-gang.ru/static/storage/323934740877670798463448388114454542673.png)

Здесь мы собираемся смоделировать зависающий процесс с помощью функции longRunningProcess. В этом примере функция не работает, но мы должны запустить ее, прежде чем запрашивать данные JSON из API. Функция longRunningProcess \* вернет ошибку, которая вызовет срабатывание функции cancel () в контексте.

Для функции ghibliReq мы настроим простой HTTP-запрос с использованием API и передадим строку для поиска материалов из API. После того, как мы настроили запрос, у нас есть оператор case, который будет получать данные канала. В зависимости от того, что происходит первым, оператор select будет отправлен либо в текущее время, либо в канал «Done» из переданного контекста. Если канал Done закрыт, мы выдаем ошибку, в противном случае мы вернем код состояния из нашего запроса.

Наш основной код начинается с настройки контекста с новым контекстом Background (), который затем передается в контекст WithCancel (). Новый ctx был передан в пустом контексте, поэтому еще ничего не произошло. Затем мы создаем новую горутину для создания нового потока и вызываем наш longRunningProcess. Как только это вызвано, мы проверяем наличие ошибок, которые вернутся, поскольку мы спроектировали это таким образом, и если есть ошибки, мы можем вызвать функцию cancel () в нашем контексте. Наконец, мы используем наш контекст для вызова нашего запроса. После запуска мы обнаруживаем, что в запросе возникла ошибка, так как это заняло слишком много времени и была вызвана функция cancel ().

В этом примере мы запускаем longRunningProcess перед нашим запросом, потому что это необходимо до того, как мы вызовем наш запрос. Если функция выдает ошибку, нам нужно иметь возможность вызвать «cancel ()», чтобы мы могли вывести из строя функцию ghibliReq (). Как только мы это установили, мы вызываем отмену для нашего контекста до того, как функция сможет запуститься. Это сделано намеренно, чтобы показать, как работает отмена. Мы могли бы легко изменить time.Sleep () в longRunningProcess, указав 1000 миллисекунд, и наша функция запроса будет работать до вызова cancel (), но в производственной среде, если цель состоит в том, чтобы убедиться, что мы поддерживаем поток стека вызовов, мы бы убедились, что мы не возвращаем ошибки и не вызываем cancel () для этого контекста.

QuickNote: имейте в виду, что контекстно- зависимый вызов не должен быть блокирующим действием без необходимости. Все дело в том, чтобы все работало.

### Context.WithDeadline

Функция WithDeadline требует двух аргументов. Один из них - родительский контекст, а другой - новый объект времени. Функция берет родительский контекст и настраивает его, чтобы соответствовать новому объекту времени, который был передан. Есть несколько предостережений. Если вы передадите контекст, который уже предшествует переданному во временном объекте, тогда исходный код передаст просто возврат контекста WithCancel с теми же требованиями отмены, что и родительский объект, который вы можете видеть [в источнике](https://github.com/golang/go/blob/master/src/context/context.go#L430). Канал « Done» закрывается по истечении нового срока. Вы также можете вручную вернуть функцию отмены, иначе она закроется, когда родительский контекст Done канал закрыт. Какое бы из этих событий ни случилось первым.

Ниже мы можем увидеть, как работает WithDeadline:

![](https://dev-gang.ru/static/storage/7254030913885366259159978767681971293.png)

Мы продолжим идею создания сайта с обзорами фильмов. Если честно, было бы неплохо создать сайт, посвященный исключительно фильмам Studio Ghibli. В приведенном выше примере выполняется что-то вроде примера withCancel. Мы собираемся повторно использовать функцию, чтобы продемонстрировать наш контекст. Повторно используйте то, что работает, сэкономьте время. Мы собираемся сделать запрос и вернуть статус указанного запроса. Разница в том, как мы обрабатываем наш контекст.

Гипотетически нам нужно создать целую группу этих каскадных запросов, и мы хотим убедиться, что все происходит вовремя во всем стеке вызовов. Чтобы отслеживать время и корректно выводить ошибки, когда это необходимо, мы можем продолжать использовать крайние сроки и увеличивать время для дополнительных вызовов. В нашем примере мы создаем фоновый контекст, а затем передаем его вместе с новым временем. Теперь мы получаем возвращаемый контекст в нашей переменной ctx примерно на 1 секунду. В нашем примере, если процесс запроса занимает больше 1 секунды, наш контекст вызывает функцию отмены и закрывает канал Done, вызывая ошибку запроса.

Мы видим, что это зависит от установленных нами стандартов. Установка времени подразумевает, что у вас есть хорошее представление о том, сколько времени должно занять что-то. Это может зависеть от доступности вашего сервера, подключения к Интернету, аппаратных ограничений и т. д. Я также видел, как люди ворчали по поводу определенных соглашений об уровне обслуживания, гарантирующих возврат активов в течение определенного периода времени. С точки зрения удобства использования при использовании контекста крайние сроки могут помочь гарантировать, что мы сможем получить информацию в разумные сроки и вернуть ее, если нет.

### Context.WithTimeout

Следующая важная функция - WithTimeout. Это небольшое отличие от функции WithDeadline. Чтобы сделать что-то оригинальное, WithTimeout просто возвращает контекст WithDeadline с переданным аргументом времени, добавленным к крайнему сроку. Другими словами, он действует аналогично WithDeadline в том, что он берет родительский и увеличивает время, чтобы вернуть производный контекст с новым временем, добавленным ко времени до вызова функции отмены и закрытия канала Done. Я сделаю этот пример еще проще:

![](https://dev-gang.ru/static/storage/278269619296175117099188607904172133340.png)

То же, что и в примере до того, как мы установили тайм-аут, чтобы закрыть канал «Done» после отведенного времени. В нашем случае, если через полсекунды мы все еще ждем звонка, у которого тайм-аут. Мне нравится библиотека HTTP go, потому что в ней есть встроенная функция для возврата теневой копии запроса с [добавленным новым контекстом](https://golang.org/src/net/http/request.go?s=12980:13039#L341).

### Context.WithValue

Последний фрагмент источника, которого я собираюсь коснуться, - это функция ContextWithValue. Этот вопрос немного спорен, поскольку его природа, насколько я могу судить, идет вразрез с тем, каким должен быть контекст. Контекст должен быть способом гарантировать, что данные будут поступать в наши программы и из них. Однако ценностная часть контекста может использоваться для передачи информации туда и обратно. Функция позволяет вам передавать интерфейс "ключ-значение" для передачи ваших вызовов.

Из исходного сообщения о контексте [«WithValue предоставляет способ связать значения в области запроса с контекстом»](https://blog.golang.org/context). Я немного расскажу о том, для чего его нельзя использовать. Большинство статей или руководств, с которыми я сталкивался, похоже, согласны с тем, что передача информации, которая живет вне самого запроса, была плохой идеей. Соединения с БД, аргументы функций, все, что не создается и не уничтожается в рамках этого запроса, вероятно, не является отличным шаблоном проектирования. При этом передача значений в вашем контексте может быть полезной.

Давайте посмотрим на код:

![](https://dev-gang.ru/static/storage/130509301875666212292319408691136931357.png)

Мы собираемся использовать тот же код из последнего примера. Только в этом случае мы собираемся создать новую функцию, которая будет вычислять ложный идентификатор запроса. Скажем, я хочу вести базу данных всех моих запросов, потому что ... я не знаю, я психопат. Или я работаю на АНБ и создаю шпионское ПО, чтобы следить за моей бывшей во имя национальной безопасности. И поскольку они не обучают меня оперативному анализу, я не знаю, как различать данные, которые что-то указывают, и белый шум, поэтому я собираю все. Даже безобидные вызовы открытого API для поиска информации о мультфильмах. Я очень устал прямо сейчас.

В нашем примере мы делаем то же, что и выше; настроить контекст с таймаутом на полсекунды. Только теперь у нас есть вспомогательный метод, который будет вычислять новый идентификатор запроса, и мы будем использовать контекст для передачи этого идентификатора в контексте в качестве нового интерфейса, к которому мы можем получить доступ и что-то с ним делать. В этом поддельном сценарии мы регистрируем это и закрываем контекст. Это будет соответствовать нашему добровольному стандарту хранения только информации, относящейся к этому звонку.

О передаче значений в контексте можно еще многое узнать. Я видел статьи, в которых промежуточное ПО используется для работы между двумя сервисами, чтобы что-то работало лучше. Я мог бы углубиться в это, и, поскольку это немного выходит за рамки этого, я могу написать об этом позже. Кто знает, мне нужно поспать.

### Вывод

Библиотека контекста помогает добавить некоторую разумность вызовам в нашей программе. При разработке программы включение контекста в наши функции должно происходить как можно раньше. Как уже упоминалось, раньше было легко создать нашу функцию с TODO в качестве заполнителя и вернуться при рефакторинге. Также было упомянуто, что программы должны быть созданы так, чтобы отказываться корректно. Возьмите это у человека, который долгое время создавал расплывчатые сообщения об ошибках, которые никто не может понять, включая меня. Пользователь не должен знать, что вызов к чему-то не удался, просто потому, что он не получает название фильма за полсекунды.

Замечательный способ представить, насколько полезными могут быть эти контексты, был затронут в выступлении Самира. Он рассказал о практике хеджированных вызовов, когда вы вызываете избыточные службы и берете тот, который занимает меньше времени. Все дело в скорости и оптимизации. Это один из способов, которым было бы полезно создать контекст для прохождения вашей программы. Когда один возвращается, вы отменяете другой, который высвобождает ресурсы, которые поток мог использовать. Контекст - это небольшая, но очень мощная библиотека, ее следует использовать часто, тщательно продумав и спланировав, как она должна вплетаться в вашу программу. После прочтения я надеюсь, что все мы сможем лучше понять контекст и то, как мы можем его использовать!
