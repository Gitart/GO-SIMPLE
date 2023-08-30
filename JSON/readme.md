## 🤚🏻 Tutorial with JSON
![image](https://github.com/Gitart/GO-SIMPLE/assets/3950155/76a69b5d-7966-4a66-88d8-85a764123d35)


# Table 
```sql
CREATE TABLE `products` (
  id             int NOT NULL AUTO_INCREMENT,
  create_at      datetime DEFAULT NULL,
  cod            varchar(20) DEFAULT NULL ,
  title        varchar(100) DEFAULT NULL,
  productions    json DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `title_UNIQUE` (`title`)
) ;

```

## Structure
Type Field for JSON - **json.RawMessage**

```go
type Products struct {
	Id            int64           `json:"id"`             // Id
	CreateAt      time.Time       `json:"create_at"`      // Created date
	Title         string          `json:"title"`          // Name
	Productions   json.RawMessage `json:"productions"`    // Settings
}
```

# Get field
```sql
SELECT productions->"$[2][2]" as elem  FROM boiler.products where id=1;
SELECT productions->"$[2]" ,  JSON_EXTRACT(productions, "$.id_prod") FROM boiler.products where id=1;
```


## Save Json
```go

// Products (Marshal)
func ProductionsAdd(e echo.Context) error {
	dat := []Productions{}
	d := Productions{1, 1, 2}
	d2 := Productions{1, 1, 2}
	d3 := Productions{1, 1, 2}

	for i := 0; i < 22; i++ {
		dat = append(dat, d, d2, d3)
	}

	datt, _ := json.Marshal(dat)

	p := Products{
		Productions: datt,
		Action:      "Ok",
	}

	res := db.DB.Model(Products{}).
		Where("id=?", 1).
		Updates(&p)

	if res.Error == nil {
		fmt.Println(res.Error)
	}
```

## Get Json (Unmarshal)
```go
	prd := Products{}
	db.DB.Where("id=1").Find(&prd)

	do := prd.Productions
	dats := []Productions{}
	json.Unmarshal(do, &dats)

	fmt.Println(do)

	return e.JSON(200, dats)
}
























