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
    .map({"Cnt":            r.row("QUANT"), 
                "STOCK":      r.row("QUANT_STOCK"), 
                "Summ":       r.row("PRICE_BUY_SUM")  })
     .ungroup()
     .map({"ID_DRUG": r.row("group"), 
                 "Reduct":     r.row("reduction")("Summ").sum(),  
                 "Cnt":           r.row("reduction")("Cnt").sum(),
                 "STOCK":      r.row("reduction")("STOCK").sum(),
                 "DDD":         r.row("reduction")("STOCK").sum().sub(r.row("reduction")("Cnt").sum()),
                 "Cnts":         r.row("reduction")("STOCK").count(),
                 "Avg":          r.row("reduction")("STOCK").avg()
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


