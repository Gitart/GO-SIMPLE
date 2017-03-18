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




