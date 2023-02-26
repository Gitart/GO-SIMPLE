## Assitiation in GORM

1. Create table in database 
2. Set right assotiation in models

## Sql
```sql
CREATE TABLE "users" (
	"user_id"	INTEGER,
	"user_name"	TEXT,
	"login"	TEXT,
	PRIMARY KEY("user_id")
);

CREATE TABLE "payments" (
	"payment_id"	integer,
	"user_id"	integer,
	"operation"	text,
	"amount"	integer,
	CONSTRAINT "fk_payment_users" FOREIGN KEY("user_id") REFERENCES "users"("user_id"),
	PRIMARY KEY("payment_id")
);

CREATE TABLE "infos" (
	"infos_id"	integer,
	"user_id"	integer,
	"name"	text,
	"value"	text,
	PRIMARY KEY("infos_id"),
	CONSTRAINT "fk_infos_users" FOREIGN KEY("user_id") REFERENCES "users"("user_id")
);

```


## Describe models
```go

//  Users
type User struct {
	UserId   int       `json:"user_id" gorm:"primary_key"`
	UserName string    `json:"user_name"`
	Login    string    `json:"login"`
	Payments []Payment `gorm:"ForeignKey:UserId"`
	Infos    []Info    `gorm:"ForeignKey:UserId"`
}

type Payment struct {
	PaymentId int64  `gorm:"primary_key"`
	UserId    int64  `json:"user_id"`
	Operation string `json:"operation"`
	Amount    int    `json:"amount"`
}

type Info struct {
	InfosId int64  `gorm:"primary_key"`
	UserId  int64  `json:"user_id"`
	Name    string `json:"name"`
	Value   string `json:"value"`
}
```

## Add - Update 
```go
// Add
func AddUser(e echo.Context) error {
	u := User{}
	e.Bind(&u)

	res := dbs.Where("user_id=?", u.UserId).Updates(&u)
	if res.Error != nil {
		fmt.Println("ERROR", res.Error)
	}

	if res.RowsAffected == 0 {
		println("Add")
		dbs.Create(&u)
	} else {
		println("Upd")
	}

	//dbs.Debug().Model(u).Association("Payment").Append(Payment{Name: "ssss"})
	//dbs.Create(u)

	return e.JSON(200, u)
}
```



