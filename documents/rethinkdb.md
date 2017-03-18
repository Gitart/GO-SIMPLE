# RethinkDb
### Выбор данных из таблицы с и по.

```js
r.db('Barsetka').table("Events").filter(r.row("end").gt("2017-03-14").and(r.row("end").le('2017-03-15')))
r.db('Barsetka').table("Events").update({"priority_id": r.row('priority_id').coerceTo('number')})
  ```
  
