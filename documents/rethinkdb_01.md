
```golang
resp, err := DB("examples").Table("posts").Get(2).Update(map[string]interface{}{
    "status": "published",
}).RunWrite(session)
if err != nil {
    fmt.Print(err)
    return
}

fmt.Printf("%d row replaced", resp.Replaced)
```
## Update the status of all posts to published.

Code:
```golang
resp, err := DB("examples").Table("posts").Update(map[string]interface{}{
    "status": "published",
}).RunWrite(session)
if err != nil {
    fmt.Print(err)
    return
}

fmt.Printf("%d row replaced", resp.Replaced)
```
Output:
```
4 row replaced
Example (Increment)
Increment the field view of the post with id of 1. If the field views does not exist, it will be set to 0.
```

Code:

```golang
resp, err := DB("examples").Table("posts").Get(1).Update(map[string]interface{}{
    "views": Row.Field("views").Add(1).Default(0),
}).RunWrite(session)
if err != nil {
    fmt.Print(err)
    return
}

fmt.Printf("%d row replaced", resp.Replaced)
```

Output:
```
1 row replaced
Example (Nested)
Update bob's cell phone number.
```

Code:

```golang
resp, err := DB("examples").Table("users").Get("bob").Update(map[string]interface{}{
    "contact": map[string]interface{}{
        "phone": "408-555-4242",
    },
}).RunWrite(session)
if err != nil {
    fmt.Print(err)
    return
}

fmt.Printf("%d row replaced", resp.Replaced)
```

Output:
```golang
1 row replaced
Example (SoftDurability)
Update the status of the post with id of 1 using soft durability.
```

Code:

```golang
resp, err := DB("examples").Table("posts").Get(2).Update(map[string]interface{}{
    "status": "draft",
}, UpdateOpts{
    Durability: "soft",
}).RunWrite(session)
if err != nil {
    fmt.Print(err)
    return
}

fmt.Printf("%d row replaced", resp.Replaced)
```


You can efficiently order using multiple fields by using a compound index. For example order by date and title.

Code:
```golang
cur, err := DB("examples").Table("posts").OrderBy(OrderByOpts{
    Index: Desc("dateAndTitle"),
}).Run(session)
if err != nil {
    fmt.Print(err)
    return
}

var res []interface{}
err = cur.All(&res)
if err != nil {
    fmt.Print(err)
    return
}

fmt.Print(res)
```


Example (Index)
Order all the posts using the index date.

Code:
```golang
cur, err := DB("examples").Table("posts").OrderBy(OrderByOpts{
    Index: "date",
}).Run(session)
if err != nil {
    fmt.Print(err)
    return
}

var res []interface{}
err = cur.All(&res)
if err != nil {
    fmt.Print(err)
    return
}

fmt.Print(res)
```

Example (IndexDesc)
Order all the posts using the index date in descending order.

Code:
```golang
cur, err := DB("examples").Table("posts").OrderBy(OrderByOpts{
    Index: Desc("date"),
}).Run(session)
if err != nil {
    fmt.Print(err)
    return
}

var res []interface{}
err = cur.All(&res)
if err != nil {
    fmt.Print(err)
    return
}

fmt.Print(res)
```

Example (Multiple)
If you have a sequence with fewer documents than the arrayLimit, you can order it by multiple fields without an index.

Code:

```golang
cur, err := DB("examples").Table("posts").OrderBy(
    "title",
    OrderByOpts{Index: Desc("date")},
).Run(session)
if err != nil {
    fmt.Print(err)
    return
}

var res []interface{}
err = cur.All(&res)
if err != nil {
    fmt.Print(err)
    return
}

fmt.Print(res)
```

Example (MultipleWithIndex)
Notice that an index ordering always has highest precedence. The following query orders posts by date, and if multiple posts were published on the same date, they will be ordered by title.

Code:
```golang
cur, err := DB("examples").Table("posts").OrderBy(
    "title",
    OrderByOpts{Index: Desc("date")},
).Run(session)
if err != nil {
    fmt.Print(err)
    return
}

var res []interface{}
err = cur.All(&res)
if err != nil {
    fmt.Print(err)
    return
}

fmt.Print(res)
```

Insert a document without a defined primary key into the table posts where the primary key is id.
Code:

```golang
type Post struct {
    Title   string `gorethink:"title"`
    Content string `gorethink:"content"`
}

resp, err := DB("examples").Table("posts").Insert(map[string]interface{}{
    "title":   "Lorem ipsum",
    "content": "Dolor sit amet",
}).RunWrite(session)
if err != nil {
    fmt.Print(err)
    return
}
```

fmt.Printf("%d row inserted, %d key generated", resp.Inserted, len(resp.GeneratedKeys))
Output:

1 row inserted, 1 key generated
Example (Map)
Insert a document into the table posts using a map.

Code:

```golang
resp, err := DB("examples").Table("posts").Insert(map[string]interface{}{
    "id":      2,
    "title":   "Lorem ipsum",
    "content": "Dolor sit amet",
}).RunWrite(session)
if err != nil {
    fmt.Print(err)
    return
}

fmt.Printf("%d row inserted", resp.Inserted)
```
Output:
```
1 row inserted
Example (Multiple)
Insert multiple documents into the table posts.
```

Code:
```golang
resp, err := DB("examples").Table("posts").Insert([]interface{}{
    map[string]interface{}{
        "title":   "Lorem ipsum",
        "content": "Dolor sit amet",
    },
    map[string]interface{}{
        "title":   "Lorem ipsum",
        "content": "Dolor sit amet",
    },
}).RunWrite(session)
if err != nil {
    fmt.Print(err)
    return
}

fmt.Printf("%d rows inserted", resp.Inserted)
```

Output:
```
2 rows inserted
Example (Struct)
Insert a document into the table posts using a struct.
```

Code:

```golang
type Post struct {
    ID      int    `gorethink:"id"`
    Title   string `gorethink:"title"`
    Content string `gorethink:"content"`
}

resp, err := DB("examples").Table("posts").Insert(Post{
    ID:      1,
    Title:   "Lorem ipsum",
    Content: "Dolor sit amet",
}).RunWrite(session)
if err != nil {
    fmt.Print(err)
    return
}

fmt.Printf("%d row inserted", resp.Inserted)
```

Output:
```
1 row inserted
Example (Upsert)
Insert a document into the table posts, replacing the document if it already exists.
```

Code:

```golang

resp, err := DB("examples").Table("posts").Insert(map[string]interface{}{
    "id":    1,
    "title": "Lorem ipsum 2",
}, InsertOpts{
    Conflict: "replace",
}).RunWrite(session)
if err != nil {
    fmt.Print(err)
    return
}

fmt.Printf("%d row replaced", resp.Replaced)
```

Output:
```
1 row replaced
```


Group games by player.

Code:

```golang
cur, err := DB("examples").Table("games").Group("player").Run(session)
if err != nil {
    fmt.Print(err)
    return
}

var res []interface{}
err = cur.All(&res)
if err != nil {
    fmt.Print(err)
    return
}

fmt.Print(res)
```
```
func (Term) GroupByIndex
func (t Term) GroupByIndex(index interface{}, fieldOrFunctions ...interface{}) Term
```

GroupByIndex takes a stream and partitions it into multiple groups based on the fields or functions provided. Commands chained after group will be called on each of these grouped sub-streams, producing grouped data.

Example

Group games by the index type.

Code:

```golang
cur, err := DB("examples").Table("games").GroupByIndex("type").Run(session)
if err != nil {
    fmt.Print(err)
    return
}

var res []interface{}
err = cur.All(&res)
if err != nil {
    fmt.Print(err)
    return
}

fmt.Print(res)
```

// Fetch the row from the database
```golang
res, err := DB("examples").Table("heroes").GetAllByIndex("code_name", "man_of_steel").Run(session)
if err != nil {
    fmt.Print(err)
    return
}
defer res.Close()

if res.IsNil() {
    fmt.Print("Row not found")
    return
}

var hero map[string]interface{}
err = res.One(&hero)
if err != nil {
    fmt.Printf("Error scanning database result: %s", err)
    return
}
fmt.Print(hero["name"])
```


## Find a document by ID.

Code:
```golang
// Fetch the row from the database
res, err := DB("examples").Table("heroes").GetAll(1, 2).Run(session)
if err != nil {
    fmt.Print(err)
    return
}
defer res.Close()

var heroes []map[string]interface{}
err = res.All(&heroes)
if err != nil {
    fmt.Printf("Error scanning database result: %s", err)
    return
}
fmt.Print(heroes[0]["name"])
```

Output:
```
Superman
Example (OptArgs)
Find all document with an indexed value.
```


## Fetch the row from the database

```golang
res, err := DB("examples").Table("heroes").GetAll("man_of_steel").OptArgs(GetAllOpts{
    Index: "code_name",
}).Run(session)
if err != nil {
    fmt.Print(err)
    return
}
defer res.Close()

if res.IsNil() {
    fmt.Print("Row not found")
    return
}

var hero map[string]interface{}
err = res.One(&hero)
if err != nil {
    fmt.Printf("Error scanning database result: %s", err)
    return
}
fmt.Print(hero["name"])
```

Output:
```
Superman
```

## Find a document by ID.

Code:

```golang
// Fetch the row from the database
res, err := DB("examples").Table("heroes").Get(2).Run(session)
if err != nil {
    fmt.Print(err)
    return
}
defer res.Close()

if res.IsNil() {
    fmt.Print("Row not found")
    return
}

var hero map[string]interface{}
err = res.One(&hero)
if err != nil {
    fmt.Printf("Error scanning database result: %s", err)
    return
}
fmt.Print(hero["name"])
```

Output:
```
Superman
Example (Merge)
Find a document and merge another document with it.
```

Code:

```golang
// Fetch the row from the database
res, err := DB("examples").Table("heroes").Get(4).Merge(map[string]interface{}{
    "powers": []string{"speed"},
}).Run(session)
if err != nil {
    fmt.Print(err)
    return
}
defer res.Close()

if res.IsNil() {
    fmt.Print("Row not found")
    return
}

var hero map[string]interface{}
err = res.One(&hero)
if err != nil {
    fmt.Printf("Error scanning database result: %s", err)
    return
}
fmt.Printf("%s: %v", hero["name"], hero["powers"])
```

Output:
```
The Flash: [speed]
```

### Fold
```golang
cur, err := Expr([]string{"a", "b", "c"}).Fold("", func(acc, word Term) Term {
    return acc.Add(Branch(acc.Eq(""), "", ", ")).Add(word)
}).Run(session)
if err != nil {
    fmt.Print(err)
    return
}

var res string
err = cur.One(&res)
if err != nil {
    fmt.Print(err)
    return
}

fmt.Print(res)
```
Output:
```
a, b, c
```

### func (Term) ForEach
func (t Term) ForEach(args ...interface{}) Term
ForEach loops over a sequence, evaluating the given write query for each element.


It takes one argument of type `func (r.Term) interface{}`, for example clones a table:

```golang
r.Table("table").ForEach(func (row r.Term) interface{} {
    return r.Table("new_table").Insert(row)
})
```



## Get all users who are 30 years old.

Code:

```golang
// Fetch the row from the database
res, err := DB("examples").Table("users").Filter(map[string]interface{}{
    "age": 30,
}).Run(session)
if err != nil {
    fmt.Print(err)
    return
}
defer res.Close()


// Scan query result into the person variable
var users []interface{}
err = res.All(&users)
if err != nil {
    fmt.Printf("Error scanning database result: %s", err)
    return
}
fmt.Printf("%d users", len(users))
```

Output:
```
2 users
Example (Function)
Retrieve all users who have a gmail account (whose field email ends with @gmail.com).
```

Code:

```golang
// Fetch the row from the database
res, err := DB("examples").Table("users").Filter(func(user Term) Term {
    return user.Field("email").Match("@gmail.com$")
}).Run(session)
if err != nil {
    fmt.Print(err)
    return
}
defer res.Close()

// Scan query result into the person variable
var users []interface{}
err = res.All(&users)
if err != nil {
    fmt.Printf("Error scanning database result: %s", err)
    return
}
fmt.Printf("%d users", len(users))
```
Output:
```
1 users
Example (Row)
Get all users who are more than 25 years old.
```

Code:
```golang
// Fetch the row from the database
res, err := DB("examples").Table("users").Filter(Row.Field("age").Gt(25)).Run(session)
if err != nil {
    fmt.Print(err)
    return
}
defer res.Close()

// Scan query result into the person variable
var users []interface{}
err = res.All(&users)
if err != nil {
    fmt.Printf("Error scanning database result: %s", err)
    return
}
fmt.Printf("%d users", len(users))
```

Output:
```
3 users
```



### Пример работе со структурой
```golang
type ExampleTypeNested struct {
    N int
}

type ExampleTypeEmbed struct {
    C string
}

type ExampleTypeA struct {
    ExampleTypeEmbed

    A      int
    B      string
    Nested ExampleTypeNested
}

cur, err := Expr(ExampleTypeA{
    A:  1,
    B:  "b",
    ExampleTypeEmbed: ExampleTypeEmbed{
        C: "c",
    },
    Nested: ExampleTypeNested{
        N: 2,
    },
}).Run(session)
if err != nil {
    fmt.Print(err)
    return
}

var res interface{}
err = cur.One(&res)
if err != nil {
    fmt.Print(err)
    return
}

jsonPrint(res)
```

Output:
```
{
    "A": 1,
    "B": "b",
    "C": "c",
    "Nested": {
        "N": 2
    }
}
```

### Example (StructTags)
Convert a Go struct (with gorethink tags) to a ReQL object. The tags allow the field names to be changed.

Code:
```golang
type ExampleType struct {
    A   int    `gorethink:"field_a"`
    B   string `gorethink:"field_b"`
}

cur, err := Expr(ExampleType{
    A:  1,
    B:  "b",
}).Run(session)
if err != nil {
    fmt.Print(err)
    return
}

var res interface{}
err = cur.One(&res)
if err != nil {
    fmt.Print(err)
    return
}

jsonPrint(res)
```

Output:
```
{
    "field_a": 1,
    "field_b": "b"
}
```




### Return heroes and superheroes.

Code:
```golang
cur, err := DB("examples").Table("marvel").OrderBy("name").Map(Branch(
    Row.Field("victories").Gt(100),
    Row.Field("name").Add(" is a superhero"),
    Row.Field("name").Add(" is a hero"),
)).Run(session)
if err != nil {
    fmt.Print(err)
    return
}

var strs []string
err = cur.All(&strs)
if err != nil {
    fmt.Print(err)
    return
}

for _, str := range strs {
    fmt.Println(str)
}
```

Output:
```
Iron Man is a superhero
Jubilee is a hero
```

