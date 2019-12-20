## Запрос к ресурсам в сети

```js
var t=r.http("https://raw.githubusercontent.com/typicode/demo/master/db.json",{resultFormat:"json"});
var rdb=r.db("Work");

var t=r.http("https://httpbin.org/base64/SFRUUEJJTiBpcyBhd2Vzb21l",{resultFormat:"json"});
r.db("Work").table("W").insert(t);

// r.db("Work").table("W").count();
rdb.tableCreate("W1",{durability:"soft"});
rdb.tableCreate("W2",{durability:"hard", primary_key:"ids"});  

var rdb=r.db("Work");
rdb.table("W").delete();
  
r.http('http://httpbin.org/put', { method: 'PUT', data: "row" });
r.http('http://httpbin.org/post', { method: 'POST', data: {name:'Sart', player: 'Bob', game: 'tic tac toe' } })
  
r.db("Bi").table("Wrk").delete()
r.db("Bi").table("Wrk").count()
```
  
