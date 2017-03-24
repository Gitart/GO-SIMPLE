# RethinkDb
### Выбор данных из таблицы с и по.

```javascript
r.db('Barsetka').table("Events").filter(r.row("end").gt("2017-03-14").and(r.row("end").le('2017-03-15')))
r.db('Barsetka').table("Events").update({"priority_id": r.row('priority_id').coerceTo('number')})
```
  
### How i define session: 

```javascript
session, e := r.Connect(r.ConnectOpts{ Address: "localhost:28015", Database: "database", MaxActive: 0, MaxIdle: 0, // IdleTimeout: time.Minute, }) 
 // Inserting 
inserts := map[string]interface{}{"something": something, "something1": something1, "something2": something2} 
r.Db("database").Table("table").Insert(inserts).RunWrite(session) 
// Updating
 r.Db("database")
.Table("table")
.Filter(map[string]interface{}{"parameter": parameter})
.Update(map[string]interface{}{"something" : somethingMore})
.RunWrite(session) 
// Get some row row, e := r.Db("database").Table("table").Filter(map[string]interface{}{"parameter": parameter}).RunRow(session) if e!= nil { //error }
```


### Создание двух индексов сразу
```javascript
r.expr([table.index_create("foo"), table.index_create("bar")]).run()
```

### Ввывод полей с с использованием вторичного индекса
т.к. в базе содержится более 2 500 000 записей поиск без применения вторичного индекса не возможен !!!!!
```javascript
r.db('HO').table('Docs').getAll('3565510', {index:'ID_STRUCTURE'})
```

### Обновление с фильтрацией по второму ключу
```javascript
r.db('HO').table('Docs').getAll('1823305', {index:'CODE'}).update({"Status":"Ok"})
```

### Фильтрация с фильтрацией
```javascript
r.db('HO').table('Docs').getAll('1823305', {index:'CODE'}).filter({"CARD_NUMBER":"201006585"})
```

### Группировка с применением второго индекса
```javascript
r.db('HO').table('Docs').getAll('1823305', {index:'CODE'}).group("CERT_DATE")("DOC_DATE_TIME")
```

### Вставка из файла с фильтром в другой файл
```javascript
r.db('HO').table('Work')
.insert(r.db('HO')
.table('Docs')
.getAll('1823305', {index:'CODE'}).filter({"PRICE_SELL_SUM":"4.6100"}))
```

### Объединение с полями из выражения
```javascript
 r.expr([{'id':'something'},{'id':'something else'}]).eqJoin("id", r.db("HO").table("Need"));
```

### Добавление в другую таблицу
```javascript
r.db('HO').table('Work')
.insert(r.db('HO')
.table('Docs')
.getAll('1823305', {index:'CODE'})
.filter({"PRICE_SELL_SUM":"4.6100"}))
```

### Добавление в другую таблицу с опредленными полями
```javascript
r.db('HO').table('Work').insert(r.db('HO').table('Docs').getAll('1823305', {index:'CODE'}).filter({"PRICE_SELL_SUM":"4.6100"}).pluck("CARD_NUMBER","CODE") );
```

### Вывод опредленных полей
```javascript
r.db("HO").table("Contractors").pluck("id","NAME","NAME_EXT").orderBy("NAME")
```

### Убрать поля из результата вывода 
```javascript
r.db("test").table("input_polls").without("GOP","EV")
db.Table("someTable")
.OrderBy(func(row Term) Term {return row.Field("time").Field("begin")}).Run(session).All(&result)

r.db("brakeman")
. table("reports")
. getAll("25a41dfcd9171695e731533c50de573c71c63deb", {index: "brakeman_sha"})
.concatMap(function(rep) { return rep("brakeman_report")("warnings") })
. groupBy("warning_type", r.count)
.orderBy(r.desc("reduction"))
```

### Записи с 6 по 8
```javascript
r.db("test").table("input_polls").slice(6,8)
```

### Добавление перед каждым элементом в списке пару "vv":"ssss" сверху
```javascript
r.db("test").table("input_polls").group("Date").prepend({"vv":"ssss"}) 
```

### Фильтр по полю name
```javascript
r.table('Aliance').filter(r.row('name').eq("Morion"))
```

### Добавить по всей таблице поле content со значением new
```javascript
r.table('Aliance').update({"content":"new"})
```

### Просмотр всех полей в таблице
```javascript
r.table('Aliance')
```

### Добавить по строчно таблице поле post
```javascript
r.table('Aliance').get(1).update({"post":1});
r.table('Aliance').get(2).update({"post":22});
r.table('Aliance').get(3).update({"post":2});
```

### Обновление
```javascript
r.table("Aliance").filter(r.row("post").count().gt(2))
```

### Добавление на второй уровень
```javascript
r.table('Aliance').filter(r.row("post").eq(1)).update({posts: r.row("posts").append({"title": "Shakespeare", "content": "What a piece of work is man..."})});
```

### Удаление по фильтру
```javascript
r.table('authors').filter(r.row('posts').count().lt(3)).delete()
```

### Создание таблицы
```javascript
r.db("test").table_create("Authors")
```

### Заполнение JSON Format
```javascript
r.table("Authors").insert([
{ "name": "William Adama", "tv_show": "Battlestar Galactica",
"posts": [
{"title": "Decommissioning speech", "content": "The Cylon War is long over..."},
{"title": "We are at war", "content": "Moments ago, this ship received..."},
{"title": "The new Earth", "content": "The discoveries of the past few days..."}
]
},
{ "name": "Laura Roslin", "tv_show": "Battlestar Galactica",
"posts": [
{"title": "The oath of office", "content": "I, Laura Roslin, ..."},
{"title": "They look like us", "content": "The Cylons have the ability..."}
]
},
{ "name": "Jean-Luc Picard", "tv_show": "Star Trek TNG",
"posts": [
{"title": "Civil rights", "content": "There are some words I've known since..."}
]
}
])
```
  
### Выборка между двумя индексами 
```javascript
r.table('Aliance').between(1, 2)
```

### Фильтр по посту
```javascript
 r.table('Aliance').filter({"post": 22})
```

### Объединение таблиц
```javascript
 r.table('marvel').inner_join(r.table('dc'), lambda marvelRow, dcRow: marvelRow['strength'] < dcRow['strength'])
 r.table('marvel').outer_join(r.table('dc'), lambda marvelRow, dcRow: marvelRow['strength'] < dcRow['strength'])
 
 r.table('marvel').eq_join('main_dc_collaborator', r.table('dc'))
 r.table('Aliance').eqJoin('post', r.table('ts')).zip()
```

 ### По всей таблице добавили 100 в поле summ
 ```javascript
 r.table('Aliance').update({"summ":100});
```

### Выбор с двумя полями сортировка
```javascript
 r.table('Aliance').withFields('summ','post').orderBy('post')
```

### В обрезанном и отфильтрованом взять первую запись
```javascript
 r.table('Aliance').withFields('summ','post').orderBy('post').limit(2).nth(0)
```

### Bозвращает Тру если в поле содержит фразу которая есть в поле - целиком!!!
Смотрит во всех записях в текущем поле !!!
```javascript
r.table('Aliance')('name').contains("Доброго")
```


### Второе значение в наброе записей
```javascript
r.expr([1,2,3])(1)    = 2
```

### Копирование таблицы самой в себя без ИД т.к. не даст скопировать с ид можно добавить условие
```javascript
r.table('Aliance').insert(r.table('Aliance').without("id") )
```

### Возвращает номер позиции в списке буквы с
```javascript
r.expr(['a','b','c']).indexesOf('c')
```

### Добавить поле polls к таблице
```javascript
r.table('Aliance').merge({polls: 1})
```

### Добавление саму в себя без ид с добавлением поля Polls
```javascript
 r.table('Aliance').insert(r.table('Aliance').without("id").merge({polls: 12}))
```

### Добавление объекта как последовательность чисел 
```javascript
r.table('Aliance').insert(r.object('id', 12225, 'data', ['foo', 'bar']))
```

### Объединение двух таблиц
```javascript
r.table('Aliance').limit(2).union(r.table('ts').limit(3))
```

### Добавление в таблицу объединенных двух других таблиц
```javascript
r.db("test").table("Calc").insert(r.table('Aliance').limit(2).union(r.table('ts').limit(3)))
```

### Показать записи в котрых поле имеет значение
```javascript
r.db("test").table("Calc").hasFields("_id")
```

### Вывод трех любых записей
```javascript
r.db("test").table("Calc").sample(3)
```

### Содержится ли выражение в поле name
```javascript
r.table('Aliance')('name').contains('Morion')
```

### Не работает
```javascript
r.table('Aliance').get("5d5ad6b0-9d2f-4b68-aa27-26fa59b0e04f").pluck("name").prepend('newBoots')
```

### r.db('test').table('user').get(1).do(r.row('list').nth(0)) // returns "a"
```javascript
r.db('test').table('user').get(1)('list').nth(0)
r.db('test').table('test').insert( { 'nest_one': { 'nest_two': [ { 'target': '1'} ] } } )
r.db('test').table('test').insert( { 'nest_one': { 'nest_two': [ { 'target': '2'} ] } } )
r.db('test').table('test').insert( { 'nest_one': { 'nest_two': [ { 'target': '1'} ] } } )
r.db('test').table('test').insert( { 'nest_one': { 'nest_two': [ { 'target': '4'} ] } } )
r.db('test').table('test').indexCreate( 'idx_nest_target', function( obj ) { return obj('nest_one')('nest_two')('target') } )
```

### Поиск по индексу актуально при больших данных
```javascript
r.db('test').table('test').getAll( '1', {index:'idx_nest_target'} )
```

### Бинарный файл
```javascript
r.http('gravatar.com/avatar/0b1129eaca8152c556c200cd8d179187', {resultFormat: 'binary'})
```

### Вставка в таблицу из ссылки
```javascript
r.table("ts").insert(r.http("http://beta.json-generator.com/api/json/get/BhzRccE"));
r.table("ts").insert(r.http("https://drivenotepad.appspot.com/app?state=%7B%22ids%22:%5B%220B-bpvLJFQcdIRmZnMFROQnBDMDA%22%5D,%22action%22:%22open%22,%22userId%22:%22107765580792592500254%22%7D"))
```

### Вставка таблицы в таблицу
```javascript
r.table("tr").insert(r.table("ts"))
```

### Поиск в первом уровне + во вторм вхождении
```javascript
r.table("tr").filter({
           index: 1,                                             -- индекс на первом уровне
           name:{                                                -- на первом уровне
                             first:"Britt",                      -- на втором уровне
                             last:"Donaldson"                    -- на втором уровне
                }
             });
```

### Вывод второго уровня вложения
```javascript
r.table("tr")("name")("first")
```

### Возвращает True если в списке второго уровня есть такое имя хотябы один раз в любой строчке
```javascript
r.table("tr")("name")("first").contains("Jan")
```

### Вывод перечисленных полей - и обратите внимание !!! Поле name составное с уровнями и оно віводит свои под уровнями
```javascript
r.table("tr").pluck('index','id','isActive','name')

[
{
"id":  "314a57e9-58f9-4102-a081-b4c262d13c7a" ,
"index": 0 ,
"isActive": true ,
"name": {
     "first":  "Roach" ,
     "last":  "Brewer"
}
} ,

r.table('posts').filter( r.row('category').eq('article').or(r.row('genre').eq('mystery'))).run(conn, callback);
r.db("foo").table("bar").insert(document).run(durability="soft")
```

### Не работает!!
```javascript
r.table("Aliance").map(r.row.merge({Title: r.row("id")}).without("id"))
```

### Уникальные значения в поле
```javascript
r.table("tr").pluck('age').distinct()
```

### Фильтрация с последующим выводом опредленных полей
```javascript
r.table("tr").filter({age:38}).pluck('age','index','id','isActive','name')
```

### Выводом определенных поле с сортировкой во втором уровне (first)
```javascript
r.table("tr").pluck('age','index','id','isActive','name').orderBy('index','first')
```

### Фильтрация с выводом определенных поле с сортировкой во втором уровне (first)
```javascript
r.table("tr").filter({eyeColor:"sd",age:32 }) .pluck('age','index','id','isActive','name') .orderBy('index','first','last')
```

### Отобрать все не пустые значения в поле ege и вывести первіе 3
```javascript
r.table("tr").hasFields('age').limit(3)
```

### Группировка с выводом на экран
```javascript
r.table("tr").group("index","age").pluck('age','index')
```

### Обновление с поля ид поле Value 
```javascript
r.table('Aliance').get(5).update({Value: r.row('count').default(0).add(1) });
```

### Поиск в первом уровне по двум условиям
```javascript
r.table("Aliance").filter({"Title": "Aliance 4"})
r.table("tr").filter({eyeColor:"sd",age:38 });
r.table("tr").filter({age:38});
```

### Фильтр по значению в поле Title='Aliance 5' или views=1250000
```javascript
r.table('Aliance').filter(r.row('Title').default('foo').eq('Aliance 5').or(r.row('views').default('foo').eq(1250000)))
```

### Выбирает построчно начиная (с 3 строки и по 5 ) из набора записей
```javascript
r.table("tr").slice(3,5)
```

### Связь по полю ID где поля есть в обеих таблицах
```javascript
r.table("tr").eqJoin('id',r.table("ts"));
```

### Обновление по ключу
```javascript
r.table("tr").get("f261abe6-e44e-4b0b-bf07-e60abcb01e0b").update({eyeColor:"sd"})
```

### Добавление нового значения в кoнец
```javascript
r.tableCreate('stargazers'); r.table('stargazers').insert( r.http('https://api.github.com/repos/rethinkdb/rethinkdb/stargazers'));
```

### Обновление информации
```javascript
r.table('stargazers').update(r.http(r.row('url')))
```

### По десять страниц
```javascript
r.http('https://api.github.com/repos/rethinkdb/rethinkdb/stargazers',       { page: 'link-next', pageLimit: 10 })
```

### Вывод первого поля из таблицы
```javascript
r.table("tr")("eyeColor");
```

### Пропустить две строки и с третьей вібрать 3 записи
```javascript
r.table("Aliance").orderBy("id").skip(2).limit(3) 
```

### Среднее значение по полю ИД
```javascript
r.table("Aliance").avg("id")
```

### Группировка
```javascript
r.table("Aliance").group([r.row("date").year(), r.row("date").month()])
```

### Создание индекса
```javascript
r.table('invoices').indexCreate('byDay', [r.row('date').year(), r.row('date').month(), r.row('date').day()])
```

### Максимальная по группировке
```javascript
r.table("invoices").group({index: 'byDay'}).max('price')
```

### Пример вставки документа
```javascript
r.table('Aliance').insert({
               _id: "5099803df3f4948bd2f98391",
               name: { first: "Alan", last: "Turing" },
               birth: 'Jun 23, 1912',
               death: 'Jun 07, 1954',
               contribs: [ "Turing machine", "Turing test", "Turingery" ],
               views : 1250000
            }) 
```

### Замена документа без поля
```javascript
r.table("Aliance").get("1").replace(r.row.without('Value')) 
```

### Не работает !!!
```javascript
r.table('posts').filter(r.row('Value').eq('Kyev').or(r.row('Title').eq('Aliance 1'))
```

### Замена документа по ИД который должен быть указан
```javascript
r.table("Aliance").get(1).replace({"id":1,"Title":"New Aliance","Value":"Kyev"})
```

### Минуты
```javascript
r.now().minutes()
```

### Дата
```javascript
r.now()    --- Thu Nov 06 2014 19:34:55 GMT+00:00
```

### Вхождение
```javascript
r.expr('abcdefghijklmnopqrstuvwxyz').match('hijklmnopqrst')
```

### Пересечение
```javascript
r.expr('abcdefghijklmnopqrstuvwxyz').match('^[abcdef]{3}')
{
"end": 3 ,
"groups": [ ],
"start": 0 ,
"str":  "abc"
}
```

### Поиск вхождения
```javascript
r.expr('abcdefghijklmnopqrstuvwxyz').match('^abcdefghijkl')
```

### Выражение   
```javascript
r.expr('ііі')
```

### Бинарный файл
```javascript
r.http('gravatar.com/avatar/0b1129eaca8152c556c200cd8d179187', {resultFormat: 'binary'})
```

### Вставка в таблицу из ссылки
```javascript
r.table("ts").insert(r.http("http://beta.json-generator.com/api/json/get/BhzRccE"));
```

### Вставка таблицы в таблицу
```javascript
r.table("tr").insert(r.table("ts"))
```

### Вставка таблицы в таблицу в GO
```javascript
var response []interface{}
res,err:=r.Db(“test”).Table(“testtabler”).Run(sess)
err=res.All(&response)
r.Db(“test”).Table(“Intable”).Insert(response).RunWrite(sess)
```


### Поиск в первом уровне + во вторм вхождении
```javascript
r.table("tr").filter({
           index: 1,                                             -- индекс на первом уровне
           name:{                                                -- на первом уровне
                             first:"Britt",                      -- на втором уровне
                             last:"Donaldson"                    -- на втором уровне
                }
             });
```

### Вывод второго уровня вложения
```javascript
r.table("tr")("name")("first")
```


### Возвращает True если в списке второго уровня есть такое имя хотябы один раз в любой строчке
```javascript
r.table("tr")("name")("first").contains("Jan")
```

### Вывод перечисленных полей - и обратите внимание !!! Поле name составное с уровнями и оно віводит свои под уровнями
```javascript
r.table("tr").pluck('index','id','isActive','name')

[
{
"id":  "314a57e9-58f9-4102-a081-b4c262d13c7a" ,
"index": 0 ,
"isActive": true ,
"name": {
     "first":  "Roach" ,
     "last":  "Brewer"
}
} ,
```

### Уникальные значения в поле
```javascript
r.table("tr").pluck('age').distinct()
```

### Фильтрация с последующим выводом опредленных полей
```javascript
r.table("tr").filter({age:38}).pluck('age','index','id','isActive','name')
```
### Выводом определенных поле с сортировкой во втором уровне (first)
```javascript
r.table("tr").pluck('age','index','id','isActive','name').orderBy('index','first')
```

### Фильтрация с выводом определенных поле с сортировкой во втором уровне (first)
```javascript
r.table("tr").filter({eyeColor:"sd",age:32 }) .pluck('age','index','id','isActive','name') .orderBy('index','first','last')
```

### Отобрать все не пустые значения в поле ege и вывести первіе 3
```javascript
r.table("tr").hasFields('age').limit(3)
```

### Группировка с выводом на экран
```javascript
r.table("tr").group("index","age").pluck('age','index')
```

### Поиск в первом уровне по двум условиям
```javascript
r.table("tr").filter({eyeColor:"sd",age:38 });
r.table("tr").filter({age:38});
```

### Выбирает построчно начиная (с 3 строки и по 5 ) из набора записей
```javascript
r.table("tr").slice(3,5)
```

### Связь по полю ID где поля есть в обеих таблицах
```javascript
r.table("tr").eqJoin('id',r.table("ts"));
```

### Обновление по ключу
```javascript
r.table("tr").get("f261abe6-e44e-4b0b-bf07-e60abcb01e0b").update({eyeColor:"sd"})
```

### Добавление нового значения в кoнец
```javascript
r.tableCreate('stargazers'); r.table('stargazers').insert( r.http('https://api.github.com/repos/rethinkdb/rethinkdb/stargazers'));
```

### Обновление информации
```go
r.table('stargazers').update(r.http(r.row('url')))
```

### По десять страниц
```javascript
r.http('https://api.github.com/repos/rethinkdb/rethinkdb/stargazers',       { page: 'link-next', pageLimit: 10 })
```

### Вывод первого поля из таблицы
```javascript
r.table("tr")("eyeColor");
```

### Cовмещение таблиц справа
```javascript
r.table().get().merge(r.table().get())
```

### Получение одного поля во вотром уровне которое имеет значени
```javascript
r.db("test").table("Docmove").pluck({"DocumentItem":{"Title":true}})
```
### Записи  второго уровня из поля DocBody по ключу N00130 из двух столбцов выбраны записи со 2 по 11 
```javascript
r.db("test").table("Docmove").get("N00130")("DocBody").pluck("ABC","BCQ").slice(2,11)
```

### Доступ ко второму уровню вложения
```javascript
var tt={"Info":{"Id":3341, 
                           "Version Databse":"A-03444", 
                            "Version Client":"122.334", 
                            "Date Update":"2014-1218"}}; 

r.db("HO").table("Setting").insert(tt);
r.db("HO").table("Setting")("Info")

r.db("HO").table("Setting")("Info")("Version")("Infos")
````

### Перекачка в другую таблицу с преобразование ключевого поля в номер
```Javascript
var tabs = r.db("HO").table("Drugs");
var tt      = r.db("HO").table("Drug")
                    .merge({"ID":r.row("ID").coerceTo("NUMBER")});
tabs.insert(tt);
var rr = tabs.count();
```
### Агррегация с группировкой
```Javascript
   r.db("HO")
    .table("Docs")
    .limit(10000)
    .group("ID_DRUG")
    .map({"Cnt":         r.row("QUANT"), 
           "STOCK":      r.row("QUANT_STOCK"), 
           "Summ":       r.row("PRICE_BUY_SUM")  })
     .ungroup()
     .map({"ID_DRUG": r.row("group"), 
                 "Reduct":     r.row("reduction")("Summ").sum(),  
                 "Cnt":        r.row("reduction")("Cnt").sum(),
                 "STOCK":      r.row("reduction")("STOCK").sum(),
                 "DDD":        r.row("reduction")("STOCK").sum().sub(r.row("reduction")("Cnt").sum()),
                 "Cnts":       r.row("reduction")("STOCK").count(),
                 "Avg":        r.row("reduction")("STOCK").avg()
     })
   
```  

 
### Позволяет складывать в нескольких столбцах суммы
```Javascript
   r.db("HO")
    .table("Docs")
    .getAll(3565509, {index: "ID_STRUCTURE"})
    .limit(2000)
    .group("ID_DRUG")
    .map({"Cnt":r.row("QUANT"), "Summ":r.row("PRICE_BUY_SUM")})
    .ungroup()
    .map({"Csnt":r.row("group"), 
              "Cntw":r.row("reduction")("Summ").sum(), 
              "Cnt":r.row("reduction")("Cnt").sum()})
```

### Группировка и фильтрация
```Javascript
 r.db("System")
    .table("Tabtest")
    .filter({"Age":42 })
    .group("Name")
    .orderBy("Name")
    .count()
    .ungroup()
    .merge({"Name":r.row("group"), "Ages":r.row("reduction")})
    .without("reduction", "group")
```

 ### Показать второй вложенный уровень после группировки  в одном столбце
 ```Javascript
r.db("HO")
.table("Wrk")
.group("ID_STRUCTURE", "ID_DRUG")
.ungroup()
.concatMap(r.row("reduction")("ID_DRUG"))
```
### Вставка и показ первого єлемента в теге
```Javascript
var tt={"ids":123, "Title":"Test", "TAG":["fff","hhh","kkkk"]};
r.db("System").table("Works").get("S26")("bodys")("TAG").nth(0)    // первое значение
.insert({"ID":"S26", "bodys":tt})
```

### разница во времени от какой то даты
```Javascript
r.expr(r.now().toEpochTime()).sub(1421951862.587) // => 111.36300015449524
```

### Работа с многоуровневыми данными
Примеры и тесты работы со структурой
```Javascript
var p=r.now().toEpochTime();
  var c=5;
  var i={"id": p, "idd": {"idd":c, "jjj":{"hhhh":"ssssss","gggg":"kkkkk"}}};  
  var i={"idd": {"idd":c, "jjj":{"mmm":"new-MMMM8999", "tag":["nnews","tez","tax"]}}};  
         i={"idd": { "jjj":{"mmm":"aaa-eeee-new-MMMM8999"}}};  
    
r.db("test").table("temp").get(1422532194.416)
 // .orderBy(r.desc("id")).limit(10)
 .update(i)
```Javascript

## Необходимо находить строку а потом команде Update/Insert указывать путь к обновляемуму ключу данных.
Важно !! При этом не нужно указывать все ключи встречающиеся, а только те которые указывают точный путь к ключу. Другие ключи не будут тронуты.

Например :  чтобы поменять ключ gggg
```Javascript
var i={"idd": { "jjj":{"gggg":"kkkkk"}}};  
```

ВАЖНО !! На тег вида [11,22,333] это не распространяется в этом случае необходимо менять весь тег !!!!! 
Если нужно заменить один элемент в теге!!! см. “работа с тегом”.
```Javascript
  
// Declaration inetrfaces
type Mst map[string]interface{} // Map - string - interface
type Mif []interface{}           // Interface
type Mii interface{}             // Interface
type Mss map[string]string       // Map - string - string
type Msi map[string]int64        // Map - string - int64
type Msr []string                // String
```

###  Тестовая функция для проверки остальных функций
```golang
func Sys_test(w http.ResponseWriter, rr *http.Request) {

	T := Mst{"mmm": "new-MMMM8999", "mmms": "iiiiiiiiii"}
	Y := Mst{"vbar": Msr{"foo", "bar", "baz", "Neyuton"}}
	G := Msr{"foo", "bar", "baz", "Hossss"}
	L := Mst{"Status": "New", "vbar": Msr{"NN-01", "NN-02"}, "INNN": T}

	Z := Mst{"id": SKEY(), "II": "ss", "ss": T, "PO": Y, "TAG": G}
	// r.Db("test").Table("temp").Delete().Run(sessionArray[0])

	i, e := r.Db("test").Table("temp").Insert(Z).RunWrite(sessionArray[0])

	r.Db("test").Table("temp").Get("01-292015160229-255478").Update(L).RunWrite(sessionArray[0])

	T = Mst{"mmms": "n-99999999999--------ew-MMMM8999"}
	L = Mst{"Status": "New 2", "vbar": Msr{"NN-0001", "NN-00002"}, "INNN": T}

	r.Db("test")
              .Table("temp")
              .Get("01-292015160229-255478")
              .Update(L).RunWrite(sessionArray[0])

	T = Mst{"A": "A", "mmm": "MMM-n-99999999999--------ew-MMMM-00000 - Да исправлен"}
	L = Mst{"Status": "New-0000000", "vbar": Msr{"NN-01", "NN-00000"}, "INNN": T}

	r.Db("test").Table("temp").Get("01-292015160229-255478").Update(L).RunWrite(sessionArray[0])

	if e != nil {
		log.Println(e)
	}

	fmt.Fprintf(w, "OK %s", i)
}

 Пример использования :
Msr - > [“www”,”htpp”,”htps”]
```

### Пример использования фильтра
```Javascript
r.table('scores')
.changes()
.filter(r.row('new_val')('score').gt(r.row('old_val')('score')))('new_val')
.run(conn, callback)

-- в новой версии Rethinkdb 1.16.2-1
r.db("test").table("persons").changes()
```

### BETWEEN
```Javascript
r.Db("database").
  Table("table").
  Between(1, 10, r.BetweenOpts{Index: "num", RightBound: "closed",}).
  Run(session)
```

### Интеренсное использование вставки
```Javascript
r.db("System")
.table("Temp")
.insert(r({ms:[{i:1, nam:"Oleg"},{i:2, nam:"Semen"},{i:3, nam:"Seva"}]}))
```

### Замена во второй позиции подчиненного єлемента
```Javascript
r.db("System").table("Temp").get(2).update({ms:r.row("ms").changeAt(1,{i:2,nam:"ffff"} )})
```

### Добавление строки во второй уровень
```Javascript
r.db("System").table("Temp").get(2)//.update({ms:r.row("ms").append({"ff":"news"})})
r.db("System").table("Temp").get(2)//.update({ms:r.row("ms").append({"ff":"news"})})
```
### Обновление второго уровня в опредленой записи
```Javascript
 r.db("HO").table("A_3500000")
 .get(2)
.update({"Items":{"item":r.db("HO").table("A_3500000").get(2)("Items").changeAt(r.db("HO").table("A_3500000").get(2)("Items")("item").indexesOf(322).nth(0).coerceTo("number") ,{"item":321, "price":777, "name":"2sss-dddd"})}}, {nonAtomic: true});
```  
### Конвертация в дату
```Javascript
r.epochTime(1426079167215/1000).toISO8601()
r.db("test").table("Docmove").update({"HDF_TIME_STR": r.epochTime(r.row("HDF_TIME_UNX").coerceTo("number").mul(0.001)).toISO8601()})
r.db("test").table("Docmove").map({"rrrrr":r.row("HDF_TIME_UNX").coerceTo("number").mul(222)})
```

### Удалить записи у которых нет определенного поля !!!
(задача оказалась не простой)

#### Суть ситуации :
Есть таблица в которой часть строк имеют поле (Status), а часть не имеют.

#### Задача.   
Удалить записи которіе не имеют этого поля (Status).    

#### Решение :
Cуть сложности заключается в том, что если фильтровать без применения опции по умолчанию {default: true}, то не получим вообще ни каких записей, потому что их нет,
а опция по умолчанию ставит их если даже их там нет (физически нет у этих документов такого поля). В этом и заключается подвох. Поэтому этот параметр обязателен, если мы хотим получить записи которые не имеют определенного поля.
```Javascript
   r.db("HO").table("Groups").filter( r.row("Status").lt(11), {default: true})
```

Основным показателем есть параметер -  {default: true}.   
В данном пример получим все записи у которых нет поля (Status) cо значением 10.   

Окончательно конструкция для удаления будет выглядеть так.  
```Javascript
 r.db("HO").table("Groups").filter( r.row("Status").lt(10), {default: true}).delete()
```

#### Окончательная реализация :

Производится фильтрация заведомо не существующих по этим условиям строк с выдачей по умолчанию null, только после этого сработает удаление тех записей которые не имоеют определенного поля.
```Javascript
// Получение уникального списка
r.db("HO").table("Drugs")("ID_CATEGORY").distinct()
```

### Группировка во втором уровне 
```Javascript
  r.db("HO").table("ConsignmentNote")
  r.db("HO").table("ConsignmentNote")("ITEMS").group('AMOUNT_BUY')
  r.db("HO").table("ConsignmentNote").get("-152967961447961997")("ITEMS").sum("AMOUNT_BUY")
       
  // Сумирование в таблице во втором уровне ITEMS:"AMMOUNT_BUY"
  r.db("HO").table("ConsignmentNote").filter({"HDF_SEQ":199}).concatMap(r.row("ITEMS")).sum("AMOUNT_BUY")
  
  r.db("HO").table("ConsignmentNote").filter(r.row("HDF_SEQ").lt(199)).concatMap(r.row("ITEMS")).sum("AMOUNT_BUY")
```   
    
### Группировка во втором уровне ITEMS: по ID_DRUG
```Javascript     
        r.db("HO").table("ConsignmentNote")
        .filter(r.row("HDF_SEQ").lt(199))
        .concatMap(r.row("ITEMS"))
        .group("ID_DRUG")
        .sum("AMOUNT_BUY")
```


### Добавление для каждой строки в наборе записей поле таг [,,,,] 
```Javascript
r.db("System").table("Corporation").map({"TAGS":r.row("TAGS").setUnion(['newBoots', 'arc_reactor'])})
```

### Добавление для одной строки в теги поле таг [,,,,] 
```Javascript
r.db("System").table("Corporation").get("C3")("TAGS").setUnion(['newBoots', 'arc_reactor']
```

## MAP

### Описание :
Основное задание МАР - прогнать каждое значение через шаблон и выстроить их в столбец или в строчки.

Примеры :
```Javascript
r.expr([1, 2, 3]).map(function(x) { return [x, x.mul(2)] })
```

### Прогоняет каждый элемент через шаблон-расчет
```Javascript
r.expr([1,4,6]).map({"ddd":[r.row, r.row.add(1),r.now()]})
```

### C каждым элементом выполняется расчет и записывается все в один столбец
```
r.expr([1, 2, 3]).concatMap([r.row.add(1), r.row])
```

### В один столбец
```Javascript
r.expr([1, 2, 3]).concatMap([{"Значение":   r.row.add(1), "Расчет" :r.row, "Ответ":r.row.mul(2)}])
```

### В один столбец с разбивкой
```Javascript
r.expr([1, 2, 3]) .concatMap([{"Значение":  r.row.add(1), 
                               "Расчет":    r.row,
                                "Ответ":    r.row.mul(2)},
                                {"Ответ":   "ffff"},
                                {"Расчет":  "Расчет"},
                                {"Значение": "Пример"}])
				
```
### Просмотр - разворот второго уровня в таблице
```Javascript
r.db("System").table("Corporation").get("C6")("BUSINESS")("STRUCTURES").concatMap(r.row)
```

### Просмотр - разворот второго уровня в таблице (второй вариант)
```Javascript
r.db("System").table("Corporation") .get("C6")  ("BUSINESS").map(r.row("STRUCTURES")).concatMap(r.row)
```

### C разбивкой
```Javascript
r.db("System").table("Corporation")  .get("C6")("BUSINESS")("STRUCTURES") .concatMap(r.row) .map({"Аптеки":r.row})
```

### C разбивкой второй вариант
```Javascript
r.db("System").table("Corporation")  .get("C6")  .pluck({"BUSINESS":["ADDRESS","NAME_BUSINESS",{"STRUCTURES":["NAME"]}]})
```


### Добавление вниз строчки
```Javascript
r.db("System")
  .table("Corporation")
  .get("C6")  ("BUSINESS")//("STRUCTURES") 
  .map(r.row("STRUCTURES"))
  .concatMap(r.row)
  .append({"ID_STRUCTURE":"SSS","ADDRESS":"DSDDD","NAME":"yFBVTYJDFYBT"})
```


### Замена второй записи в подчиненой ветке-уровне
```Javascript
r.db("System")
  .table("Corporation")
  .get("C6")  ("BUSINESS")//("STRUCTURES") 
  .map(r.row("STRUCTURES"))
  .concatMap(r.row)
  .changeAt(2,{"ID_STRUCTURE":1234567,
                       "ADDRESS":"Новый  адрес",
                       "NAME":"Онбновление",
                       "NAME_EXT":"Расширенное"})
```

### Удаление двух позиций
```Javascript
r.db("System")
  .table("Corporation")
  .get("C6")  ("BUSINESS")//("STRUCTURES") 
  .map(r.row("STRUCTURES"))
  .concatMap(r.row)
  .deleteAt(2,3)
```

### Добавить строку на верх
```Javascript
r.db("System")
  .table("Corporation")
  .get("C6")  ("BUSINESS")//("STRUCTURES") 
  .map(r.row("STRUCTURES"))
  .concatMap(r.row)
.prepend({"ID_STRUCTURE":1234567,
                       "ADDRESS":"Новый  адрес",
                       "NAME":"Онбновление",
                       "NAME_EXT":"Расширенное"})

```


### Добавление вниз во втором уровне
```Javascript
r.db("System")
  .table("Corporation")
  .get("C6")  ("BUSINESS")//("STRUCTURES") 
  .map(r.row("STRUCTURES"))
  .concatMap(r.row)
  .append({"ID_STRUCTURE":1234567,
                       "ADDRESS":"Новый  адрес",
                       "NAME":"Онбновление",
                       "NAME_EXT":"Расширенное"})

```
### Опредление индекса позиции в списке 
```Javascript
 r.db("System")
  .table("Corporation")
  .get("C6")  ("BUSINESS")//("STRUCTURES") 
  .map(r.row("STRUCTURES"))
  .concatMap(r.row)("ID_STRUCTURE")
  .offsetsOf(6021251)
```

### Получение трех строк с 3 по 6
```Javascript
r.db("System")
  .table("Corporation")
  .get("C6")  ("BUSINESS")//("STRUCTURES") 
  .map(r.row("STRUCTURES"))
  .concatMap(r.row)//("ID_STRUCTURE")
  .slice(3,6)
```
### Добавление строки в определенную позицию
```Javascript
 r.db("System")
  .table("Corporation")
  .get("C6")  ("BUSINESS")//("STRUCTURES") 
  .map(r.row("STRUCTURES"))
  .concatMap(r.row)//("ID_STRUCTURE")
   .spliceAt(2,[{"NAME":"fff"},{"NAME_EXT":"fff"}])
```


### Добавление две строки
```Javascript
 r.db("System")
  .table("Corporation")
  .get("C6")  ("BUSINESS")//("STRUCTURES") 
  .map({"STRUCTURES":r.row("STRUCTURES")})
  //.concatMap(r.row)("ID_STRUCTURE")
  .spliceAt(2,[{"STRUCTURES":"Добавление 1"},{"STRUCTURES":"Добавление 2"}])
 ```
 ### Вставка JSON
 ```Javascript
var t='{"ss":"ddddd"}'  ;
r.db("test").table("Test").insert(r.json(t));
```


### Вставка из битбакета

```Javascript
r.http('https://bitbucket.org/arthur_savage/ho/raw/90350e9d43ec46041853d04030e8b1f44629e3f5/aplan.go', {
       auth: {
           user: "arthur_savage",
           pass: 'Gerda3000$'
       }
})
```

### Опредление максимального значения по опредленному полю
```Javascript
r.db("C3").table("A_6370944").max("DOC_DATE_TIME_BUY_STR")("DOC_DATE_TIME_BUY_STR");

r.Db("quiz")
.Table("questions")
.GetAllByIndex("category", "Bravo")
.Sample(10)
.Map(func(row r.Term) interface{} {
        return map[string]interface{}{
            "sum":  row.Field("value"),
            "list": []interface{}{row.Field("id")},
        }
    })
.Reduce(func(left, right r.Term) interface{} {
        return r.Branch(
            left.Field("sum").Add(right.Field("sum")).Lt(15),
            map[string]interface{}{
                "sum":  left.Field("sum").Add(right.Field("sum")),
                "list": left.Field("sum").Add(right.Field("sum")),
            },
            map[string]interface{}{
                "sum":  left.Field("sum"),
                "list": left.Field("sum"),
            },
        )
    })

 ```


### РАЗВОРОТ СТРУКТУРЫ В ОБЪЕКТЕ
```Javascript
{
"DATE_TIME_STR": "2015-10-14T07:55:12.000" ,
"DATE_TIME_UNX": "1444802112000" ,
"HDF_DEL": 0 ,
"HDF_EDITOR": "0" ,
"HDF_SEQ": 58 ,
"HDF_STATUS": 1 ,
"HDF_TIME_STATUS": 2 ,
"HDF_TIME_STR": "2015-11-23T17:38:33.816" ,
"HDF_TIME_UNX": "1448293113816" ,
"ID": "-1186259333377067919" ,
"ID_CASHIER": "5776074858111382956" ,
"ID_PERSON": "3464652019326971542" ,
"ID_STRUCTURE": "7136566" ,
"ITEMS": [
{
"AMOUNT": 50.86 ,
"DOC_NAME": "Расходный кассовый ордер" ,
"DOC_TYPE": "18" ,
"ID": "1250942855979650201" ,
"ID_DT": "301" ,
"ID_ENTRY": "10" ,
"ID_ITEM": "-1186259333377067919" ,
"ID_KT": "301" ,
"NAME_DT": "Готівка в національній валюті" ,
"NAME_KT": "Готівка в національній валюті"
}
] ,
"NOTE": "" ,
"NUMBER": "ВКО-0000820" ,
"REASON": "Винесення розмінної монети"
}
{
"DATE_TIME_STR": "2015-06-18T08:05:42.000" ,
"DATE_TIME_UNX": "1434607542000" ,
"HDF_DEL": 0 ,
"HDF_EDITOR": "0" ,
"HDF_SEQ": 42 ,
"HDF_STATUS": 1 ,
"HDF_TIME_STATUS": 2 ,
"HDF_TIME_STR": "2015-11-23T17:38:25.707" ,
"HDF_TIME_UNX": "1448293105707" ,
"ID": "-1120019709551775897" ,
"ID_CASHIER": "5776074858111382956" ,
"ID_PERSON": "1974277564506894594" ,
"ID_STRUCTURE": "7136566" ,
"ITEMS": [
{
"AMOUNT": 50.47 ,
"DOC_NAME": "Расходный кассовый ордер" ,
"DOC_TYPE": "18" ,
"ID": "-8637428725325917185" ,
"ID_DT": "301" ,
"ID_ENTRY": "10" ,
"ID_ITEM": "-1120019709551775897" ,
"ID_KT": "301" ,
"NAME_DT": "Готівка в національній валюті" ,
"NAME_KT": "Готівка в національній валюті"
}
] ,
"NOTE": "" ,
"NUMBER": "ВКО-0000484" ,
"REASON": "Винесення розмінної монети"
}
```


### Разворачивает поле "ITEMS": [] как в обычную таблицу для просмотра
```Javascript
 rk, err:= r.DB("C3").Table("Cashbox").Field("ITEMS").ConcatMap(r.Row).Limit(100).Run(sessionArray[0])

r.DB().Table().Get(3)(“hh”).offsetsOff(“b”)
в таблице {“hh”:[“a”,”b”,”c”,”d”]} ищет букву и возвращает индекс в списке - 1

r.DB().Table().map({“s”:r.row(“N”).add(“ddd”)))
к каждому полю справа “ddd”
```

### ГРУППИРОВКА СО СЛОЖЕНИЕМ ВО ВТОРОМ УРОВНЕ!!!!!!!!
```Javascript
r.db("C3")
  .table("Cashbox")
  .map({"Items":r.row("ITEMS")("AMOUNT").coerceTo("array")})  
  .concatMap(r.row("Items"))
```Javascript

#### Есть сложна и замысловатя структура так получилось исторически
Задача :   
Необходимо получить доступ к     "AMOUNT": 50.86     
```Json
{
"DATE_TIME_STR":  "2015-10-14T07:55:12.000" ,
"DATE_TIME_UNX":  "1444802112000" ,
"DELETED": 0 ,
"HDF_CORP":  "C3" ,
"HDF_EDITOR":  "6370944" ,
"HDF_SEQ": 396 ,
"HDF_TB":  "Cashbox" ,
"HDF_TIME_STATUS": 2 ,
"HDF_TIME_STR":  "2016-02-10T12:25:06.790" ,
"HDF_TIME_UNX":  "1455099906790" ,
"ID":  "-1186259333377067919" ,
"ID_CASHIER":  "5776074858111382956" ,
"ID_PERSON":  "3464652019326971542" ,
"ID_STRUCTURE":  "8378100" ,
"ITEMS": [
           {
              "AMOUNT": 50.86 ,
              "DOC_NAME":  "Видатковий касовий ордер" ,
              "DOC_TYPE":  "18" ,
              "ID":  "1250942855979650201" ,
              "ID_DT":  "301" ,
              "ID_ENTRY":  "10" ,
              "ID_ITEM":  "-1186259333377067919" ,
              "ID_KT":  "301" ,
              "NAME_DT":  "Готівка в національній валюті" ,
              "NAME_KT":  "Готівка в національній валюті"
           }
] ,
"NOTE":  "" ,
"NUMBER":  "ВКО-0000820" ,
"REASON":  "Винесення розмінної монети"
}
```

### Супер Решение :
Промежуточные шаги.   

```Javascript
r.db("C3")
  .table("Cashbox")
  .group("ID_STRUCTURE")
  .map({"Items":r.row("ITEMS")("AMOUNT")})                // выводим второе вложенное поле    
  .concatMap(r.row("Items"))                              // Доступ ко второму уровню
  .map({"Otvet":r.row})                                   // Оглавление поля
  .sum("Otvet")
  
r.db("C3")
  .table("Cashbox")
  .map({"Items":r.row("ITEMS")("AMOUNT")})                // выводим второе вложенное поле    
  .concatMap(r.row("Items"))                                       // Доступ ко второму уровню
  .map({"Otvet":r.row})                                              // Оглавление поля
  

r.db("C3")
  .table("Cashbox")
  .map({"Items":r.row("ITEMS")("AMOUNT")})                // выводим второе вложенное поле    
  .concatMap(r.row("Items"))                                       // Доступ ко второму уровню
  .map({"Otvet":r.row})                                              // Оглавление поля
  .sum("Otvet")                                                         // Суммирование 
```                      
### И наконец вот ЗОЛОТОЕ решение каверзной задачи 

```Javascript
r.db("C3")
  .table("Cashbox")                                               // таблица
  .group("ID_STRUCTURE")                                          // Группировка по структуре
  .map({"Items":r.row("ITEMS")("AMOUNT")})                        // выводим второе вложенное поле    
  .concatMap(r.row("Items"))                                      // Доступ ко второму уровню
  .map({"Sum":r.row})                                             // Оглавление поля
  .sum("Sum")                                                     // Суммирование по полю 
  .ungroup()                                                      // Разгруппировка нужна   
  .map({"STRUCTURE":r.row("group"),"SUM":r.row("reduction")})     // Подпись столбцов 
```  


#### Итог после группировки :
```json
  [
      {   "STRUCTURE":  "6370944" , "SUM": 3665008.479999999    } ,
      {   "STRUCTURE":  "8378100" , "SUM": 3471130.0200000005   }
  ]
```


### Обновление в цикле 
```Javascript
       r.db("System")
      .table("Users")
      .filter({"Fname":"Сенюк"})
         .forEach(function(row) { return r.db("System").table("Corporation").filter({"ID":"C159"}).update({"WTS": row("Id"), "WTSNAME":row("Name")})    })
```
    
### Привязка пользователей к ВТС базам
```Javascript
 r.db("System")
      .table("Users")
      .filter({"Fname":"Шеремет"})
      .forEach(function(row) { return r.db("System").table("Corporation").filter({"ID":"C6"}).update({"WTS": row("Id")})    });

  r.db("System")
      .table("Users")
      .filter({"Fname":"Шеремет"})
      .forEach(function(row) { return r.db("System").table("Corporation").filter({"ID":"C160"}).update({"WTS": row("Id")})    });
  
  
     r.db("System")
      .table("Users")
      .filter({"Fname":"Шеремет"})
      .forEach(function(row) { return r.db("System").table("Corporation").filter({"ID":"C158"}).update({"WTS": row("Id")})    });

       r.db("System")
      .table("Users")
      .filter({"Fname":"Зуб"})
         .forEach(function(row) { return r.db("System").table("Corporation").filter({"ID":"C9"}).update({"WTS": row("Id"), "WTSNAME":row("Name")})    });
      
           
       r.db("System")
      .table("Users")
      .filter({"Fname":"Зуб"})
         .forEach(function(row) { return r.db("System").table("Corporation").filter({"ID":"C160"}).update({"WTS": row("Id"), "WTSNAME":row("Name")})    });
```


 ### Обновление в цикле в GO
 * Обновление одной таблицы в цикле из другой таблицы связанные между собой полями

```Javascript
     r.DB("C001").
       Table("Matrix").
       ForEach(func(row r.Term) interface{}{
       	                                     return 
       	                                     r.DB("C001").Table("Log").Filter(Mst{"code": row.Field("code")}).Update(Mst{"Noo": row.Field("id")})
       	                                    }).Run(sessionArray[0])

```


### TAG
```Javascript  
// И то и то в Tag[]
r.db("System")
 .table("Navigation")
 .filter(r.row("Tag").contains("An","Rp"))

// Или то или др  
  r.db("System")
 .table("Navigation")
 .filter(r.or(r.row("Tag").contains("An"),
              r.row("Tag").contains("Rp")))

```

### Фильтрация по содержанию тега
```Javascript  
r.table("documents").filter(r.row("tags").contains("foo", "def"))
```

```json
{ id: 1, title: "Foo post", tags: ["foo", "bar", "baz"]},
{ id: 2, title: "Bar post", tags: ["abc", "def", "ghi"]},
{ id: 3, title: "Baz post", tags: ["foo", "def", "buzz"]},
```


### Поиск подчиненных полей
```json
{
    "feed": {
        "entry": [
            {
                "title": {
                    "label": "Some super duper app"
                },
                "summary": {
                    "label": "Bla bla bla..."
                }
            },
            {
                "title": {
                    "label": "Another awsome app"
                },
                "summary": {
                    "label": "Lorem ipsum blabla..."
                }
            }
        ]
    }
}

```

#### Решение :
```javascript
r.table("feeds")
.concatMap(function(doc) { return doc("feed")("entry") })
.filter(function(entry)  { return entry("title")("label").match("xyz") })
```

### Sample

```javascript
r.table('customers')
  .hasFields('purchases')
  .map(function(customer) {
    return {
      id: customer('id'),
      firstName: customer('firstName'),
      lastName: customer('lastName'),
      purchaseTotal: customer('purchases')('amount')
        .reduce(function(acc,amount) {
          return acc.add(amount);
        }, 0)
    }
  })
  .orderBy(r.desc('purchaseTotal')).limit(10)
  .innerJoin(
    r.table('visits').groupBy('customer', r.sum('hits')),
    function( customer, visit ) {
      return customer('id').eq(visit('group')('customer'));
    }
  )
  .map({
    id:r.row('left')('id'),
    firstName:r.row('left')('firstName'),
    lastName:r.row('left')('lastName'),
    purchases:r.row('left')('purchaseTotal'),
    hits:r.row('right')('reduction')
  })
```

