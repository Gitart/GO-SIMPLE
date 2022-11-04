# GO: GORM

### install gorm package

go get -u gorm.io/gorm
go get -u -x gorm.io/driver/mysql

### Connecting to a Database

source: [https://gorm.io/docs/connecting\_to\_the\_database.html](https://gorm.io/docs/connecting_to_the_database.html)

#### Get started

```
//go.mod
module gorm_test

go 1.15

require (
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.11
)

```

```
//main.go
package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ProductID      int
	Manufacturer   string
	Sku            string
	Upc            string
	PricePerUnit   string
	QuantityOnHand int
	ProductName    string
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	var products []Product
	db.Find(&products)

	fmt.Println(products[5].Manufacturer)
}

E:\Go\GORM_test>go run .
Quigley, Casper and Boyer
```

![](https://fixjourney.files.wordpress.com/2021/01/image-45.png?w=901)

### Declaring Models

**Conventions**: by default, GORM uses `ID` as primary key, pluralize struct name to `snake_cases` as table name, `snake_case` as column name, and uses `CreatedAt`, `UpdatedAt` to track creating/updating time

```
// gorm.Model definition
type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}
```

**Field-Level Permission**

```
type User struct {
  Name string `gorm:"<-:create"` // allow read and create
  Name string `gorm:"<-:update"` // allow read and update
  Name string `gorm:"<-"`        // allow read and write (create and update)
  Name string `gorm:"<-:false"`  // allow read, disable write permission
  Name string `gorm:"->"`        // readonly (disable write permission unless it configured )
  Name string `gorm:"->;<-:create"` // allow read and create
  Name string `gorm:"->:false;<-:create"` // createonly (disabled read from db)
  Name string `gorm:"-"`  // ignore this field when write and read with struct
}
```

**Creating/Updating Time/Unix (Milli/Nano) Seconds Tracking**

GORM use `CreatedAt`, `UpdatedAt` to track creating/updating time by convention, and GORM will set the [current time](https://gorm.io/docs/gorm_config.html#now_func) when creating/updating if the fields are defined

To use fields with a different name, you can configure those fields with tag `autoCreateTime`, `autoUpdateTime`

If you prefer to save UNIX (milli/nano) seconds instead of time, you can simply change the field’s data type from `time.Time` to `int`

```
type User struct {
  CreatedAt time.Time // Set to current time if it is zero on creating
  UpdatedAt int       // Set to current unix seconds on updaing or if it is zero on creating
  Updated   int64 `gorm:"autoUpdateTime:nano"` // Use unix nano seconds as updating time
  Updated   int64 `gorm:"autoUpdateTime:milli"`// Use unix milli seconds as updating time
  Created   int64 `gorm:"autoCreateTime"`      // Use unix seconds as creating time
}
```

**Embedded Struct**

**KEY WORD: gorm.Model in struct**

```
type User struct {
  gorm.Model
  Name string
}
// equals
type User struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
  Name string
}
```

**\`gorm:"embedded"\`**

```
type Author struct {
  Name  string
  Email string
}

type Blog struct {
  ID      int
  Author  Author `gorm:"embedded"`
  Upvotes int32
}
// equals
type Blog struct {
  ID    int64
  Name  string
  Email string
  Upvotes  int32
}
```

And you can use tag `embeddedPrefix` to add prefix to embedded fields’ db name, for example:

```
type Author struct {
  Name  string
  Email string
}

type Blog struct {
  ID      int
  Author  Author `gorm:"embedded;embeddedPrefix:author_"`
  Upvotes int32
}
// equals
type Blog struct {
  ID          int64
  AuthorName  string
  AuthorEmail string
  Upvotes     int32
}

```

### AutoMigration

AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes. It will change existing column’s type if its size, precision, nullable changed. It **WON’T** delete unused columns to protect your data.

```
type User struct {
	gorm.Model
	Email    string
	Username string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})
}

E:\Go\GORM_test>go run .

2021/01/23 23:18:45 E:/Go/GORM_test/main.go:35 SLOW SQL >= 200ms
[1138.828ms] [rows:0] CREATE TABLE `users`
                      (`id` bigint unsigned AUTO_INCREMENT,
                       `created_at` datetime(3) NULL,
                       `updated_at` datetime(3) NULL,
                       `deleted_at` datetime(3) NULL,
                       `email` longtext,
                       `username` longtext,
                       PRIMARY KEY (`id`),
                       INDEX idx_users_deleted_at (`deleted_at`))

```

![](https://fixjourney.files.wordpress.com/2021/01/image-46.png?w=222)

## **Associations**

## Belongs To (one to many)

A `belongs to` association sets up a one-to-one connection with another model, such that each instance of the declaring model “belongs to” one instance of the other model.

For example, if your application includes users and companies, and each user can be assigned to exactly one company

```
type User struct {
	gorm.Model
	Name      string
        // `User` belongs to `Company`, `CompanyID` is the foreign key
	CompanyID int
	Company   Company
}

type Company struct {
	ID   int
	Name string
}

func main() {

	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{}, &Company{})

	companies := []Company{
		{ID: 1, Name: "IBM"},
	}

	for _, c := range companies {
		db.Create(&c)
	}

	users := []User{
		{Name: "Goerge", CompanyID: 1},
		{Name: "John", CompanyID: 1},
	}

	for _, u := range users {
		db.Create(&u)
	}

}
```

![](https://fixjourney.files.wordpress.com/2021/01/image-49.png?w=224)

![](https://fixjourney.files.wordpress.com/2021/01/image-51.png?w=796)

![](https://fixjourney.files.wordpress.com/2021/01/image-52.png?w=850)

### **Override Foreign Key** **With other name**

```
type User struct {
  gorm.Model
  Name         string
  CompanyRefer int
  // use CompanyRefer as foreign key
  Company      Company `gorm:"foreignKey:CompanyRefer"`
}

type Company struct {
  ID   int
  Name string
}
```

Demo:

```
type User struct {
	gorm.Model
	Name         string
	CompanyRefer int
	// use CompanyRefer as foreign key
	Company Company `gorm:"foreignKey:CompanyRefer"`
}

type Company struct {
	ID   int
	Name string
}

func main() {

	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{}, &Company{})

}
```

```
E:\Go\GORM_test>go run .

2021/01/24 12:49:40 E:/Go/GORM_test/main.go:54 SLOW SQL >= 200ms
[211.831ms] [rows:0] CREATE TABLE `companies`
                       (`id` bigint AUTO_INCREMENT,
                        `name` longtext,
                         PRIMARY KEY (`id`))

2021/01/24 12:49:40 E:/Go/GORM_test/main.go:54 SLOW SQL >= 200ms
[457.679ms] [rows:0] CREATE TABLE `users`
                        (`id` bigint unsigned AUTO_INCREMENT,
                         `created_at` datetime(3) NULL,
                         `updated_at` datetime(3) NULL,
                         `deleted_at` datetime(3) NULL,
                         `name` longtext,
                         `company_refer` bigint,
                          PRIMARY KEY (`id`),
                          INDEX idx_users_deleted_at (`deleted_at`),
                          CONSTRAINT `fk_users_company` FOREIGN KEY (`company_refer`) REFERENCES `companies`(`id`))
```

![](https://fixjourney.files.wordpress.com/2021/01/image-48.png?w=232)

### **Override References**

For a belongs to relationship, GORM usually uses the owner’s primary field as the foreign key’s value, for the above example, it is `Company`‘s field `ID`.

When you assign a user to a company, GORM will save the company’s `ID` into the user’s `CompanyID` field.

You are able to change it with tag `references`.

```
type User struct {
  gorm.Model
  Name      string
  CompanyID string
  // use Code as references. CompanyID will save Code from [Company] table
  Company   Company `gorm:"references:Code"`
}

type Company struct {
  ID   int
  Code string
  Name string
}
```

Demo Code: Throw error when execution

```
type User struct {
	gorm.Model
	Name      string
	CompanyID string
	// use Code as references. CompanyID will save Code from [Company] table as the foreign key ? The primary key in [Company] is ID. DOES NOT MAKE SENSE!
	Company Company `gorm:"references:Code"`
}

type Company struct {
	ID   int
	Code string
	Name string
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{}, &Company{})

}
```

```
E:\Go\GORM_test>go run .

2021/01/24 12:38:55 E:/Go/GORM_test/main.go:55 SLOW SQL >= 200ms
[243.674ms] [rows:0] CREATE TABLE `companies`
                     (`id` bigint AUTO_INCREMENT,
                      `code` longtext,
                      `name` longtext,PRIMARY KEY (`id`))

2021/01/24 12:38:55 E:/Go/GORM_test/main.go:55 Error 1170: BLOB/TEXT column 'company_id' used in key specification without a key length
[2.030ms] [rows:0] CREATE TABLE `users` (`id` bigint unsigned AUTO_INCREMENT,`created_at` datetime(3) NULL,`updated_at` datetime(3) NULL,`deleted_at` datetime(3) NULL,`name` longtext,`company_id` longtext,PRIMARY KEY (`id`),INDEX idx_users_deleted_at (`deleted_at`),CONSTRAINT `fk_users_company` FOREIGN KEY (`company_id`) REFERENCES `companies`(`code`))
```

**OnUpdate:CASCADE,OnDelete:SET NULL on Foreign Key**

```
type User struct {
  gorm.Model
  Name      string
  CompanyID int
  Company   Company `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Company struct {
  ID   int
  Name string
}
```

## Has One

A `has one` association sets up a one-to-one connection with another model

The difference between Belong To (One to Many) relationship is **ONLY has one**.
The example below shows **one User can only has one CreditCard**, instead of having multiple CreditCard.

```
// User has one CreditCard, CreditCardID is the foreign key
type User struct {
  gorm.Model
  CreditCard CreditCard
}

type CreditCard struct {
  gorm.Model
  Number string
  UserID uint
}

[910.122ms] [rows:0] CREATE TABLE `users`
                      (`id` bigint unsigned AUTO_INCREMENT,
                       `created_at` datetime(3) NULL,
                       `updated_at` datetime(3) NULL,
                       `deleted_at` datetime(3) NULL,
                        PRIMARY KEY (`id`),
                        INDEX idx_users_deleted_at (`deleted_at`))

2021/01/24 13:12:51 E:/Go/GORM_test/main.go:52 SLOW SQL >= 200ms
[299.592ms] [rows:0] CREATE TABLE `credit_cards`
                       (`id` bigint unsigned AUTO_INCREMENT,
                        `created_at` datetime(3) NULL,
                        `updated_at` datetime(3) NULL,
                        `deleted_at` datetime(3) NULL,
                        `number` longtext,
                        `user_id` bigint unsigned,
                         PRIMARY KEY (`id`),
                         INDEX idx_credit_cards_deleted_at (`deleted_at`),
                         CONSTRAINT `fk_users_credit_card` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`))
```

![](https://fixjourney.files.wordpress.com/2021/01/image-50.png?w=223)

#### Override Foreign Key

```
type User struct {
	gorm.Model
	CreditCard CreditCard `gorm:"foreignKey:UserName"`
	// use UserName as foreign key
}

type CreditCard struct {
	gorm.Model
	Number   string
	UserName string
}

CONSTRAINT `fk_users_credit_card` FOREIGN KEY (`user_name`) REFERENCES `users`(`id`)
```

![](https://fixjourney.files.wordpress.com/2021/01/image-53.png?w=240)

Demo Code:

```
type User struct {
	gorm.Model
	CreditCard CreditCard `gorm:"foreignKey:UserName"`
	// use UserName as foreign key
}

type CreditCard struct {
	gorm.Model
	Number   string
	UserName string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{}, &CreditCard{})

	db.Create(&User{CreditCard: CreditCard{Number: "9374763", UserName: "John"}})
}

```

![](https://fixjourney.files.wordpress.com/2021/01/image-54.png?w=821)

![](https://fixjourney.files.wordpress.com/2021/01/image-55.png?w=911)

### Polymorphism Association

```
type Cat struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner;"`
}

type Dog struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
	ID        int
	Name      string
	OwnerID   int
	OwnerType string
}

func main() {

	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Cat{}, &Dog{}, &Toy{})

	db.Create(&Dog{Name: "dog1", Toy: Toy{Name: "toy1"}})
        // INSERT INTO `dogs` (`name`) VALUES ("dog1")
        // INSERT INTO `toys` (`name`,`owner_id`,`owner_type`) VALUES ("toy1","1","dogs")

}
```

```
2021/01/24 14:46:51 E:/Go/GORM_test/main.go:60 SLOW SQL >= 200ms
[320.628ms] [rows:0] CREATE TABLE `cats` (`id` bigint AUTO_INCREMENT,
                                          `name` longtext,
                                           PRIMARY KEY (`id`))

2021/01/24 14:46:52 E:/Go/GORM_test/main.go:60 SLOW SQL >= 200ms
[359.728ms] [rows:0] CREATE TABLE `dogs` (`id` bigint AUTO_INCREMENT,
                                          `name` longtext,
                                           PRIMARY KEY (`id`))

2021/01/24 14:46:52 E:/Go/GORM_test/main.go:60 SLOW SQL >= 200ms
[436.265ms] [rows:0] CREATE TABLE `toys` (`id` bigint AUTO_INCREMENT,
                                          `name` longtext,
                                          `owner_id` bigint,
                                          `owner_type` longtext,
                                           PRIMARY KEY (`id`))

```

![](https://fixjourney.files.wordpress.com/2021/01/image-56.png?w=559)

![](https://fixjourney.files.wordpress.com/2021/01/image-57.png?w=588)

### Polymorphism Association

```
type Dog struct {
  ID   int
  Name string
  Toy  Toy `gorm:"polymorphic:Owner;polymorphicValue:master"`
}

type Toy struct {
  ID        int
  Name      string
  OwnerID   int
  OwnerType string
}

db.Create(&Dog{Name: "dog1", Toy: Toy{Name: "toy1"}})
// INSERT INTO `dogs` (`name`) VALUES ("dog1")
// INSERT INTO `toys` (`name`,`owner_id`,`owner_type`) VALUES ("toy1","1","master")

```

### Self-Referential Has One

```
type User struct {
	gorm.Model
	Name      string
	ManagerID *uint
	Manager   *User
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	db.Create(&User{Name: "David", Manager: &User{Name: "Duncan"}})
}

```

![](https://fixjourney.files.wordpress.com/2021/01/image-58.png?w=798)

## Has Many

```
type User struct {
	gorm.Model
	Name        string
	CreditCards []CreditCard
}

type CreditCard struct {
	gorm.Model
	CreditCardNumber string
	UserID           uint
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{}, &CreditCard{})

	db.Create(&User{Name: "David",
		CreditCards: []CreditCard{{CreditCardNumber: "001"}, {CreditCardNumber: "002"}}})
}
```

![](https://fixjourney.files.wordpress.com/2021/01/image-59.png?w=728)

![](https://fixjourney.files.wordpress.com/2021/01/image-60.png?w=839)

### Self-Referential Has Many

```
type User struct {
  gorm.Model
  Name      string
  ManagerID *uint
  Team      []User `gorm:"foreignkey:ManagerID"`
}
```

## Many To Many

```
// User has and belongs to many languages, `user_languages` is the join table
type User struct {
	gorm.Model
	Name      string
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{}, &Language{})

	user1 := User{Name: "John",
		Languages: []Language{
			{Name: "ZH"},
			{Name: "EN"},
		}}

	user2 := User{Name: "Henry",
		Languages: []Language{
			{Name: "ZH"},
			{Name: "FR"},
		}}

	db.Create(&user1)
	db.Create(&user2)
}

```

When using GORM `AutoMigrate` to create a table for `User`, GORM will create the join table automatically

![](https://fixjourney.files.wordpress.com/2021/01/image-61.png?w=205)

![](https://fixjourney.files.wordpress.com/2021/01/image-62.png?w=527)

![](https://fixjourney.files.wordpress.com/2021/01/image-63.png?w=511)

![](https://fixjourney.files.wordpress.com/2021/01/image-64.png?w=485)

#### Back Reference

```
type User struct {
	gorm.Model
	Name      string
	Languages []*Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_languages;"`
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{}, &Language{})

	user1 := User{Name: "John",
		Languages: []*Language{
			{Name: "ZH"},
			{Name: "EN"},
		}}

	user2 := User{Name: "Henry",
		Languages: []*Language{
			{Name: "ZH"},
			{Name: "FR"},
		}}

	language := Language{Name: "XX",
		Users: []*User{{Name: "George"}, {Name: "Lee"}}}

	db.Create(&user1)
	db.Create(&user2)
	db.Create(&language)
}

```

![](https://fixjourney.files.wordpress.com/2021/01/image-65.png?w=516)

![](https://fixjourney.files.wordpress.com/2021/01/image-66.png?w=512)

![](https://fixjourney.files.wordpress.com/2021/01/image-67.png?w=495)

#### Self-Referential Many2Many

```
type User struct {
  gorm.Model
  Friends []*User `gorm:"many2many:user_friends"`
}

// Which creates join table: user_friends
//   foreign key: user_id, reference: users.id
//   foreign key: friend_id, reference: users.id

```

#### Customize JoinTable

`JoinTable` can be a full-featured model, like having `Soft Delete`，`Hooks` supports and more fields, you can setup it with `SetupJoinTable`, for example:

**Strange: the struct name is Person. The table name is people**. Not sure how the gorm naming convention when creating table

```
type Person struct {
	ID        int
	Name      string
	Addresses []Address `gorm:"many2many:person_addresses;"`
}

type Address struct {
	ID   uint
	Name string
}

type PersonAddress struct {
	PersonID  int `gorm:"primaryKey"`
	AddressID int `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.SetupJoinTable(&Person{}, "Addresses", &PersonAddress{})
	db.AutoMigrate(&Person{}, &Address{})

	p := Person{Name: "John",
		Addresses: []Address{
			{Name: "23 Fred Road"},
			{Name: "5 Purchase Street"},
		}}

	db.Create(&p)
}
```

![](https://fixjourney.files.wordpress.com/2021/01/image-68.png?w=208)

![](https://fixjourney.files.wordpress.com/2021/01/image-69.png?w=430)

![](https://fixjourney.files.wordpress.com/2021/01/image-70.png?w=438)

![](https://fixjourney.files.wordpress.com/2021/01/image-71.png?w=484)

#### Composite Foreign Keys

```
type Person struct {
	ID        int       `gorm:"primaryKey"`
	Name      string    `gorm:"primaryKey"`
	Addresses []Address `gorm:"many2many:person_addresses;"`
}

type Address struct {
	ID   uint
	Name string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Person{}, &Address{})

	p := Person{Name: "John",
		Addresses: []Address{
			{Name: "23 Fred Road"},
			{Name: "5 Purchase Street"},
		}}

	db.Create(&p)
}
```

![](https://fixjourney.files.wordpress.com/2021/01/image-72.png?w=469)

## Associations

### Auto Create/Update

```
// Save update value in database, if the value doesn't have primary key, will insert it
func (db *DB) Save(value interface{}) (tx *DB)

// Create insert the value into database
func (db *DB) Create(value interface{}) (tx *DB)
```

```
type Person struct {
	ID        int
	Name      string
	Addresses []Address `gorm:"many2many:person_addresses;"`
}

type Address struct {
	ID   uint
	Name string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Person{}, &Address{})

	p := Person{Name: "John",
		Addresses: []Address{
			{Name: "23 Fred Road"},
			{Name: "5 Purchase Street"},
		}}

	db.Create(&p)

	p.Name = "Philip"
	p.Addresses[0].Name = "323 Forest Road"

	newAdd := Address{Name: "1 Kings Street"}
	p.Addresses = append(p.Addresses, newAdd)

	db.Save(&p)
}
```

![](https://fixjourney.files.wordpress.com/2021/01/image-73.png?w=366)

![](https://fixjourney.files.wordpress.com/2021/01/image-74.png?w=350)

![](https://fixjourney.files.wordpress.com/2021/01/image-75.png?w=399)

Save will not only insert newly added data, but also update the data.

#### FullSaveAssociations

Not update associations’ data

```
type Person struct {
	ID        int
	Name      string
	Addresses []Address `gorm:"many2many:person_addresses;"`
}

type Address struct {
	ID   uint
	Name string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Person{}, &Address{})

	p := Person{Name: "John",
		Addresses: []Address{
			{Name: "23 Fred Road"},
			{Name: "5 Purchase Street"},
		}}

	db.Create(&p)

	p.Name = "Philip"
	p.Addresses[0].Name = "323 Forest Road"

	newAdd := Address{Name: "1 Kings Street"}
	p.Addresses = append(p.Addresses, newAdd)

	db.Session(&gorm.Session{FullSaveAssociations: false}).Updates(&p)

}
```

![](https://fixjourney.files.wordpress.com/2021/01/image-76.png?w=379)

![](https://fixjourney.files.wordpress.com/2021/01/image-77.png?w=356)

p.Name is updated to Philip.
p.Addresses\[0\].Name is not updated. But the newly added p.Addresses\[2\] is added into table

### Select

Only save the selected field

```
type Person struct {
	ID        int
	Name      string
	Addresses []Address `gorm:"many2many:person_addresses;"`
}

type Address struct {
	ID   uint
	Name string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Person{}, &Address{})

	p := Person{Name: "John",
		Addresses: []Address{
			{Name: "23 Fred Road"},
			{Name: "5 Purchase Street"},
		}}

	db.Select("Name").Create(&p)
}
```

![](https://fixjourney.files.wordpress.com/2021/01/image-86.png?w=182)

![](https://fixjourney.files.wordpress.com/2021/01/image-87.png?w=375)

Only has peoples table has one entry. The rest three tables are all empty.

```
	db.Select("Addresses", "Email").Create(&p)
```

![](https://fixjourney.files.wordpress.com/2021/01/image-88.png?w=370)

![](https://fixjourney.files.wordpress.com/2021/01/image-89.png?w=402)

![](https://fixjourney.files.wordpress.com/2021/01/image-90.png?w=415)

![](https://fixjourney.files.wordpress.com/2021/01/image-91.png?w=412)

### OMIT

Save all the data except the omitted field

```
type Person struct {
	ID        int
	Name      string
	Email     Email
	Addresses []Address `gorm:"many2many:person_addresses;"`
}

type Email struct {
	ID       uint
	Email    string
	PersonID int
}

type Address struct {
	ID   uint
	Name string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Person{}, &Address{}, &Email{})

	p := Person{Name: "John",
		Email: Email{Email: "name@example.com"},
		Addresses: []Address{
			{Name: "23 Fred Road"},
			{Name: "5 Purchase Street"},
		}}

	db.Omit("Email").Create(&p)
}

```

![](https://fixjourney.files.wordpress.com/2021/01/image-81.png?w=338)

![](https://fixjourney.files.wordpress.com/2021/01/image-82.png?w=365)

![](https://fixjourney.files.wordpress.com/2021/01/image-83.png?w=386)

![](https://fixjourney.files.wordpress.com/2021/01/image-84.png?w=445)

```
        // Only peoples table has data, the rest table is empty.
	db.Omit("Email", "Addresses").Create(&p)
```

### Skip all associations

```
type Person struct {
	ID        int
	Name      string
	Email     Email
	Addresses []Address `gorm:"many2many:person_addresses;"`
}

type Email struct {
	ID       uint
	Email    string
	PersonID int
}

type Address struct {
	ID   uint
	Name string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Person{}, &Address{}, &Email{})

	p := Person{Name: "John",
		Email: Email{Email: "name@example.com"},
		Addresses: []Address{
			{Name: "23 Fred Road"},
			{Name: "5 Purchase Street"},
		}}

        // Skip all associations when creating a person.Only [People]
	db.Omit(clause.Associations).Create(&p)

}

```

![](https://fixjourney.files.wordpress.com/2021/01/image-85.png?w=558)

## Association Mode

### Find Associations

Find matched associations

```
type People struct {
	ID        int
	Name      string
	Email     Email
	Addresses []Address `gorm:"many2many:people_address;"`
}

type Email struct {
	ID       uint
	Email    string
	PeopleID int
}

type Address struct {
	ID   uint
	Name string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&People{}, &Address{}, &Email{})

	p := People{Name: "John",
		Email: Email{Email: "name@example.com"},
		Addresses: []Address{
			{Name: "23 Fred Road"},
			{Name: "5 Purchase Street"},
		}}

        //must save into database before use Association and Find
	db.Create(&p)

	addresses := []Address{}
	db.Model(&p).Association("Addresses").Find(&addresses)

	for _, add := range addresses {
		fmt.Println(add.Name)
	}
}

23 Fred Road
5 Purchase Street
```

#### WHERE CLAUSE

```
	db.Model(&p).Where("name = ?", "23 Fred Road").Association("Addresses").Find(&addresses)

23 Fred Road
```

```
	names := []string{"23 Fred Road", "5 Purchase Street"}
	db.Model(&p).Where("name in ?", names).Association("Addresses").Find(&addresses)

23 Fred Road
5 Purchase Street
```

```
	db.Model(&p).Where("name like ?", "%Fred%").Association("Addresses").Find(&addresses)

23 Fred Road
```

```
other demo

// Get first matched record
db.Where("name = ?", "jinzhu").First(&user)
// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

// Get all matched records
db.Where("name <> ?", "jinzhu").Find(&users)
// SELECT * FROM users WHERE name <> 'jinzhu';

// IN
db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');

// LIKE
db.Where("name LIKE ?", "%jin%").Find(&users)
// SELECT * FROM users WHERE name LIKE '%jin%';

// AND
db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

// Time
db.Where("updated_at > ?", lastWeek).Find(&users)
// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

// BETWEEN
db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

```

### Append Associations

Append new associations for `many to many`, `has many`

```
type People struct {
	ID        int
	Name      string
	Email     Email
	Addresses []Address `gorm:"many2many:people_address;"`
}

type Email struct {
	ID       uint
	Email    string
	PeopleID int
}

type Address struct {
	ID   uint
	Name string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&People{}, &Address{}, &Email{})

	p := People{Name: "John",
		Email: Email{Email: "name@example.com"},
		Addresses: []Address{
			{Name: "23 Fred Road"},
			{Name: "5 Purchase Street"},
		}}

	db.Create(&p)

	db.Model(&p).Association("Addresses").Append(&Address{Name: "3 Kings Street"})
	db.Model(&p).Association("Addresses").Append([]Address{{Name: "36 Queens Street"}, {Name: "67 George Street"}})
}
```

![](https://fixjourney.files.wordpress.com/2021/01/image-92.png?w=570)

### Replace Associations

Replace current associations with new ones

```
type People struct {
	ID        int
	Name      string
	Email     Email
	Addresses []Address `gorm:"many2many:people_address;"`
}

type Email struct {
	ID       uint
	Email    string
	PeopleID int
}

type Address struct {
	ID   uint
	Name string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&People{}, &Address{}, &Email{})

	p := People{Name: "John",
		Email: Email{Email: "name@example.com"},
		Addresses: []Address{
			{Name: "23 Fred Road"},
			{Name: "5 Purchase Street"},
		}}

	db.Create(&p)

	db.Model(&p).Association("Addresses").Replace([]Address{{Name: "36 Queens Street"}, {Name: "67 George Street"}})
}
```

![](https://fixjourney.files.wordpress.com/2021/01/image-93.png?w=393)

![](https://fixjourney.files.wordpress.com/2021/01/image-94.png?w=377)

### Delete Associations

```
type People struct {
	ID        int
	Name      string
	Email     Email
	Addresses []Address `gorm:"many2many:people_address;"`
}

type Email struct {
	ID       uint
	Email    string
	PeopleID int
}

type Address struct {
	ID   uint
	Name string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&People{}, &Address{}, &Email{})

	p := People{Name: "John",
		Email: Email{Email: "name@example.com"},
		Addresses: []Address{
			{Name: "23 Fred Road"},
			{Name: "5 Purchase Street"},
		}}

	db.Create(&p)

        db.Model(&p).Association("Addresses").Delete([]Address{{ID: 1}, {ID: 2}})
	//db.Model(&p).Association("Addresses").Delete([]Address{{ID: 1, Name: "23 Fred Road"}, {ID: 2, Name: "5 Purchase Street"}})
}
```

The primary key value must be provided

![](https://fixjourney.files.wordpress.com/2021/01/image-95.png?w=391)

### Clear Associations

```
type People struct {
	ID        int
	Name      string
	Email     Email
	Addresses []Address `gorm:"many2many:people_address;"`
}

type Email struct {
	ID       uint
	Email    string
	PeopleID int
}

type Address struct {
	ID   uint
	Name string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&People{}, &Address{}, &Email{})

	p := People{Name: "John",
		Email: Email{Email: "name@example.com"},
		Addresses: []Address{
			{Name: "23 Fred Road"},
			{Name: "5 Purchase Street"},
		}}

	db.Create(&p)

	db.Model(&p).Association("Addresses").Clear()
}

```

![](https://fixjourney.files.wordpress.com/2021/01/image-96.png?w=391)

![](https://fixjourney.files.wordpress.com/2021/01/image-97.png?w=404)

Count Associations

```
	c := db.Model(&p).Association("Addresses").Count()
	fmt.Println(c)
```

### Delete with Select

```
// Incorrect Demo
type People struct {
	ID        int
	Name      string
	Email     Email
	Addresses []Address `gorm:"many2many:people_address;"`
}

type Email struct {
	ID       uint
	Email    string
	PeopleID int
}

type Address struct {
	ID   uint
	Name string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&People{}, &Address{}, &Email{})

	p := People{Name: "John",
		Email: Email{Email: "name@example.com"},
		Addresses: []Address{
			{Name: "23 Fred Road"},
			{Name: "5 Purchase Street"},
		}}

	db.Create(&p)

	// delete user's has one/many/many2many relations when deleting user
	db.Select("Addresses").Delete(&p)
	db.Delete(&p)
}
```

```
E:\Go\GORM_test>go run .

2021/02/02 22:49:03 E:/Go/GORM_test/main.go:58 SLOW SQL >= 200ms
[381.245ms] [rows:0] CREATE TABLE `peoples` (`id` bigint AUTO_INCREMENT,`name` longtext,PRIMARY KEY (`id`))

2021/02/02 22:49:03 E:/Go/GORM_test/main.go:58 SLOW SQL >= 200ms
[201.374ms] [rows:0] CREATE TABLE `addresses` (`id` bigint unsigned AUTO_INCREMENT,`name` longtext,PRIMARY KEY (`id`))

2021/02/02 22:49:04 E:/Go/GORM_test/main.go:58 SLOW SQL >= 200ms
[587.619ms] [rows:0] CREATE TABLE `people_address` (`people_id` bigint,`address_id` bigint unsigned,PRIMARY KEY (`people_id`,`address_id`),CONSTRAINT `fk_people_address_people` FOREIGN KEY (`people_id`) REFERENCES `peoples`(`id`),CONSTRAINT `fk_people_address_address` FOREIGN KEY (`address_id`) REFERENCES `addresses`(`id`))

2021/02/02 22:49:04 E:/Go/GORM_test/main.go:70 Error 1451: Cannot delete or update a parent row: a foreign key constraint fails (`inventorydb`.`emails`, CONSTRAINT `fk_peoples_email` FOREIGN KEY (`people_id`) REFERENCES `peoples` (`id`))
[56.145ms] [rows:0] DELETE FROM `peoples` WHERE `peoples`.`id` = 1

2021/02/02 22:49:04 E:/Go/GORM_test/main.go:71 Error 1451: Cannot delete or update a parent row: a foreign key constraint fails (`inventorydb`.`people_address`, CONSTRAINT `fk_people_address_people` FOREIGN KEY (`people_id`) REFERENCES `peoples` (`id`))
[15.956ms] [rows:0] DELETE FROM `peoples` WHERE `peoples`.`id` = 1
```

```
//Correct Demo
type People struct {
	gorm.Model
	Name      string
	Email     Email
	Addresses []Address `gorm:"many2many:people_address;"`
}

type Email struct {
	gorm.Model
	Email    string
	PeopleID int
}

type Address struct {
	gorm.Model
	Name string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&People{}, &Address{}, &Email{})

	p := People{Name: "John",
		Email: Email{Email: "name@example.com"},
		Addresses: []Address{
			{Name: "23 Fred Road"},
			{Name: "5 Purchase Street"},
		}}

	db.Create(&p)

	// delete user's has one/many/many2many relations when deleting user
	db.Select("Addresses").Delete(&p)
	db.Delete(&p)
}
```

![](https://fixjourney.files.wordpress.com/2021/02/image.png?w=631)

![](https://fixjourney.files.wordpress.com/2021/02/image-1.png?w=543)

![](https://fixjourney.files.wordpress.com/2021/02/image-2.png?w=509)

![](https://fixjourney.files.wordpress.com/2021/02/image-3.png?w=615)

```
        // has one
	db.Select("Email").Delete(&p)
```

![](https://fixjourney.files.wordpress.com/2021/02/image-4.png?w=558)

![](https://fixjourney.files.wordpress.com/2021/02/image-5.png?w=685)

![](https://fixjourney.files.wordpress.com/2021/02/image-6.png?w=440)

#### clause.Associations

```
type People struct {
	ID        int
	Name      string
	Email     Email
	Addresses []Address `gorm:"many2many:people_address;"`
}

type Email struct {
	ID       uint
	Email    string
	PeopleID int
}

type Address struct {
	ID   uint
	Name string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&People{}, &Address{}, &Email{})

	p := People{Name: "John",
		Email: Email{Email: "name@example.com"},
		Addresses: []Address{
			{Name: "23 Fred Road"},
			{Name: "5 Purchase Street"},
		}}

	db.Create(&p)

	// delete user's has one/many/many2many relations when deleting user
	db.Select(clause.Associations).Delete(&p)
	db.Delete(&p)
}
```

![](https://fixjourney.files.wordpress.com/2021/01/image-98.png?w=326)

![](https://fixjourney.files.wordpress.com/2021/01/image-99.png?w=419)

![](https://fixjourney.files.wordpress.com/2021/01/image-100.png?w=357)

![](https://fixjourney.files.wordpress.com/2021/01/image-101.png?w=418)

## Preloading (Eager Loading)

### Preload

```
type People struct {
	gorm.Model
	Name      string
	Email     Email
	Addresses []Address `gorm:"many2many:people_address;"`
}

type Email struct {
	gorm.Model
	Email    string
	PeopleID int
}

type Address struct {
	gorm.Model
	Name string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&People{}, &Address{}, &Email{})

	p1 := People{Name: "John",
		Email: Email{Email: "name1@example.com"},
		Addresses: []Address{
			{Name: "23 Fred Road"},
			{Name: "5 Purchase Street"},
		}}

	db.Create(&p1)

	p2 := People{Name: "David",
		Email: Email{Email: "name2@example.com"},
		Addresses: []Address{
			{Name: "42 George Street"},
			{Name: "67 Clinton Street"},
		}}

	db.Create(&p2)

	ps := []People{}

        // Addresses data will be loaded when the People is loaded
	db.Preload("Addresses").Find(&ps)

	for _, pp := range ps {
		fmt.Println("email: ", pp.Email.Email)
		for _, addr := range pp.Addresses {
			fmt.Println("address:", addr.Name)
		}
	}
}

```

```
email:
address: 23 Fred Road
address: 5 Purchase Street
email:
address: 42 George Street
address: 67 Clinton Street
```

```
//Chain is supported
db.Preload("Orders").Preload("Profile").Preload("Role").Find(&users)
// SELECT * FROM users;
// SELECT * FROM orders WHERE user_id IN (1,2,3,4); // has many
// SELECT * FROM profiles WHERE user_id IN (1,2,3,4); // has one
// SELECT * FROM roles WHERE id IN (4,5,6); // belongs to
```

### Joins Preloading

`` `Join Preload` works with one-to-one relation, e.g: `has one`, `belongs to` ``
**refer to section below:** JOINS PRELOADING

```
db.Joins("Company").Joins("Manager").Joins("Account").First(&user, 1)
db.Joins("Company").Joins("Manager").Joins("Account").First(&user, "users.name = ?", "jinzhu")
db.Joins("Company").Joins("Manager").Joins("Account").Find(&users, "users.id IN ?", []int{1,2,3,4,5})

```

### Preload All

db.Preload(**clause.Associations**).Find(&ps)

`clause.Associations` won’t preload nested associations

```
type People struct {
	gorm.Model
	Name      string
	Email     Email
	Addresses []Address `gorm:"many2many:people_address;"`
}

type Email struct {
	gorm.Model
	Email    string
	PeopleID int
}

type Address struct {
	gorm.Model
	Name string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&People{}, &Address{}, &Email{})

	p1 := People{Name: "John",
		Email: Email{Email: "name1@example.com"},
		Addresses: []Address{
			{Name: "23 Fred Road"},
			{Name: "5 Purchase Street"},
		}}

	db.Create(&p1)

	p2 := People{Name: "David",
		Email: Email{Email: "name2@example.com"},
		Addresses: []Address{
			{Name: "42 George Street"},
			{Name: "67 Clinton Street"},
		}}

	db.Create(&p2)

	ps := []People{}

	db.Preload(clause.Associations).Find(&ps)

	for _, pp := range ps {
		fmt.Println("email: ", pp.Email.Email)
		for _, addr := range pp.Addresses {
			fmt.Println("address:", addr.Name)
		}
	}
}

```

```
E:\Go\GORM_test>go run .
email:  name1@example.com
address: 23 Fred Road
address: 5 Purchase Street
email:  name2@example.com
address: 42 George Street
address: 67 Clinton Street
email:  name1@example.com
address: 23 Fred Road
address: 5 Purchase Street
email:  name2@example.com
address: 42 George Street
address: 67 Clinton Street
```

### Nested Preloading

```
db.Preload("Orders.OrderItems.Product").Preload("CreditCard").Find(&users)

// Customize Preload conditions for `Orders`
// And GORM won't preload unmatched order's OrderItems then
db.Preload("Orders", "state = ?", "paid").Preload("Orders.OrderItems").Find(&users)

```

### Preload with conditions

```
db.Preload("Addresses").Find(&ps)
db.Preload("Addresses", "Name like ?", "%Street").Find(&ps)
```

```
type People struct {
	gorm.Model
	Name      string
	Email     Email
	Addresses []Address `gorm:"many2many:people_address;"`
}

type Email struct {
	gorm.Model
	Email    string
	PeopleID int
}

type Address struct {
	gorm.Model
	Name string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&People{}, &Address{}, &Email{})

	p1 := People{Name: "John",
		Email: Email{Email: "name1@example.com"},
		Addresses: []Address{
			{Name: "23 Fred Road"},
			{Name: "5 Purchase Street"},
		}}

	db.Create(&p1)

	p2 := People{Name: "David",
		Email: Email{Email: "name2@example.com"},
		Addresses: []Address{
			{Name: "42 George Street"},
			{Name: "67 Clinton Street"},
		}}

	db.Create(&p2)

	ps := []People{}

	db.Preload("Addresses", "Name like ?", "%Street").Find(&ps)

	for _, pp := range ps {
		fmt.Println("name: ", pp.Name)
		fmt.Println("email: ", pp.Email.Email)
		for _, addr := range pp.Addresses {
			fmt.Println("address:", addr.Name)
		}
	}
}
```

```
name:  John
email:
address: 5 Purchase Street
name:  David
email:
address: 42 George Street
address: 67 Clinton Street
```

```
db.Where("Name = ?", "David").Preload("Addresses", "Name like ?", "%Street").Find(&ps)

name:  David
email:
address: 42 George Street
address: 67 Clinton Street
```

### Custom Preloading SQL

You are able to custom preloading SQL by passing in `func(db *gorm.DB) *gorm.DB`, for example:

```
db.Preload("Orders", func(db *gorm.DB) *gorm.DB {
  return db.Order("orders.amount DESC")
}).Find(&users)
// SELECT * FROM users;
// SELECT * FROM orders WHERE user_id IN (1,2,3,4) order by orders.amount DESC;

```

## Create

### Create Record

```
user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

result := db.Create(&user) // pass pointer of data to Create

user.ID             // returns inserted data's primary key
result.Error        // returns error
result.RowsAffected // returns inserted records count
```

### Create Record With Selected Fields

```
db.Select("Name", "Age", "CreatedAt").Create(&user)
// INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")

db.Omit("Name", "Age", "CreatedAt").Create(&user)
// INSERT INTO `users` (`birthday`,`updated_at`) VALUES ("2020-01-01 00:00:00.000", "2020-07-04 11:05:21.775")

```

### Batch Insert

```
var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
db.Create(&users)

for _, user := range users {
  user.ID // 1,2,3
}
```

### Create Hooks

GORM allows user defined hooks to be implemented for `BeforeSave`, `BeforeCreate`, `AfterSave`, `AfterCreate`. These hook method will be called when creating a record, refer [Hooks](https://gorm.io/docs/hooks.html) for details on the lifecycle

```
func (u *People) BeforeCreate(tx *gorm.DB) (err error) {

	if u.Name == "John" {
		return errors.New("John is not allow to be added")
	}
	return
}
```

```
E:\Go\GORM_test>go run .

2021/02/04 23:22:20 E:/Go/GORM_test/main.go:76 John is not allow to be added
[0.000ms] [rows:0]
```

### Create From Map

GORM supports create from `map[string]interface{}` or
`[]map[string]interface{}{{...}, {...},}`

```
db.Model(&User{}).Create(map[string]interface{}{
  "Name": "jinzhu", "Age": 18,
})

db.Model(&User{}).Create([]map[string]interface{}{
  {"Name": "jinzhu_1", "Age": 18},
  {"Name": "jinzhu_2", "Age": 20},
})
```

### Create From SQL Expression/Context Valuer

```
type User struct {
	gorm.Model
	Name       string
	NameBase64 string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	db.Model(User{}).Create(map[string]interface{}{
		"Name":       "Wei Zhong",
		"NameBase64": clause.Expr{SQL: "TO_BASE64(?)", Vars: []interface{}{"Wei Zhong"}},
            })
            }
```

```
function TO_BASE64 is mysql built-in function
INSERT INTO `users` (`name`,`name_base64`) VALUES ('Wei Zhong',TO_BASE64('Wei Zhong'))
```

![](https://fixjourney.files.wordpress.com/2021/02/image-7.png?w=459)

### Create With Associations

When creating some data with associations, if its associations value is not zero-value, those associations will be upserted, and its `Hooks` methods will be invoked.

```
type CreditCard struct {
  gorm.Model
  Number   string
  UserID   uint
}

type User struct {
  gorm.Model
  Name       string
  CreditCard CreditCard
}

db.Create(&User{
  Name: "jinzhu",
  CreditCard: CreditCard{Number: "411111111111"}
})
// INSERT INTO `users` ...
// INSERT INTO `credit_cards` ...
```

You can skip saving associations with `Select`, `Omit`, for example:

```
db.Omit("CreditCard").Create(&user)

// skip all associations
db.Omit(clause.Associations).Create(&user)

```

### Default Values

```
type User struct {
  ID   int64
  Name string `gorm:"default:galeone"`
  Age  int64  `gorm:"default:18"`
}
```

Then the default value *will be used* when inserting into the database for [zero-value](https://tour.golang.org/basics/12) fields (`0`, `''`, `false`)

### Upsert / On Conflict

#### OnConflict Do Nothing

```
type User struct {
	ID         int
	Name       string
	NameBase64 string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	db.Create(&User{ID: 1, Name: "David", NameBase64: "DavidBase64"})
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&User{ID: 1, Name: "John", NameBase64: "JohnBase64"})
}
```

![](https://fixjourney.files.wordpress.com/2021/02/image-15.png?w=333)

#### Update columns to default value on \`id\` conflict

```
type User struct {
	ID         int
	Name       string
	NameBase64 string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	db.Create(&User{ID: 1, Name: "David", NameBase64: "DavidBase64"})

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"name_base64": "DefaultBase64"}),
	}).Create(&User{ID: 1, Name: "John", NameBase64: "JohnBase64"})
}
```

![](https://fixjourney.files.wordpress.com/2021/02/image-16.png?w=414)

#### Use SQL expression

```
SELECT TO_BASE64('David'); 'RGF2aWQ='
SELECT TO_BASE64('John');  'Sm9obg=='

```

```
type User struct {
	ID         int
	Name       string
	NameBase64 string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	db.Create(&User{ID: 1, Name: "David", NameBase64: "DavidBase64"})

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"name_base64": gorm.Expr("TO_BASE64(name)")}),
	}).Create(&User{ID: 1, Name: "John", NameBase64: "JohnBase64"})
}
```

![](https://fixjourney.files.wordpress.com/2021/02/image-17.png?w=378)

name\_base64 change to TO\_BASE64(“David”)

```
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"name_base64": gorm.Expr("TO_BASE64('John')")}),
	}).Create(&User{ID: 1, Name: "John", NameBase64: "JohnBase64"})
```

![](https://fixjourney.files.wordpress.com/2021/02/image-18.png?w=330)

#### Update on selected columns

```
type User struct {
	ID         int
	Name       string
	NameBase64 string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	db.Create(&User{ID: 1, Name: "David", NameBase64: "DavidBase64"})

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "name_base64"}),
	}).Create(&User{ID: 1, Name: "John", NameBase64: "JohnBase64"})
}
```

![](https://fixjourney.files.wordpress.com/2021/02/image-19.png?w=337)

#### Update all columns, except primary keys

```
type User struct {
	ID         int
	Name       string
	NameBase64 string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	db.Create(&User{ID: 1, Name: "David", NameBase64: "DavidBase64"})

	db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&User{ID: 1, Name: "John", NameBase64: "JohnBase64"})
}
```

![](https://fixjourney.files.wordpress.com/2021/02/image-20.png?w=328)

## Query

Retrieving a single object

```
db.First(&user)
// SELECT * FROM users ORDER BY id LIMIT 1;

// Get one record, no specified order
db.Take(&user)
// SELECT * FROM users LIMIT 1;

// Get last record, order by primary key desc
db.Last(&user)
// SELECT * FROM users ORDER BY id DESC LIMIT 1;

result := db.First(&user)
result.RowsAffected // returns found records count
result.Error        // returns error

// check error ErrRecordNotFound
errors.Is(result.Error, gorm.ErrRecordNotFound)

```

The `First`, `Last` method will find the first/last record order by primary key, it only works when querying with struct or provides model value, if no primary key defined for current model, will order by the first field.

```
type User struct {
	ID         int
	Name       string
	NameBase64 string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	db.Create(&User{ID: 1, Name: "David", NameBase64: "DavidBase64"})
	db.Create(&User{ID: 2, Name: "John", NameBase64: "JohnBase64"})

	result := map[string]interface{}{}

        // db.Model(&User{}).First(&result) Also works
	db.Table("users").Take(&result)

	for k, v := range result {
		fmt.Printf("%v, %v\n", k, v)
	}
}

id, 1
name, David
name_base64, DavidBase64
```

### Retrieving objects with primary key

```
db.First(&user, 10)
// SELECT * FROM users WHERE id = 10;

db.First(&user, "10")
// SELECT * FROM users WHERE id = 10;

db.Find(&users, []int{1,2,3})
// SELECT * FROM users WHERE id IN (1,2,3);

```

### Retrieving all objects

```
// Get all records
result := db.Find(&users)
// SELECT * FROM users;

result.RowsAffected // returns found records count, equals `len(users)`
result.Error        // returns error
```

### Conditions

#### String Conditions

```
db.Where("name = ?", "jinzhu").First(&user)
// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

db.Where("name <> ?", "jinzhu").Find(&users)
// SELECT * FROM users WHERE name <> 'jinzhu';

db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');

db.Where("name LIKE ?", "%jin%").Find(&users)
// SELECT * FROM users WHERE name LIKE '%jin%';

db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

db.Where("updated_at > ?", lastWeek).Find(&users)
// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

```

#### Struct & Map Conditions

```

db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;

db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

db.Where([]int64{20, 21, 22}).Find(&users)
// SELECT * FROM users WHERE id IN (20, 21, 22);
```

When querying with struct, GORM will only query with non-zero fields, that means if your field’s value is `0`, `''`, `false` or other [zero values](https://tour.golang.org/basics/12), it won’t be used to build query conditions

```
db.Where(&User{Name: "jinzhu", Age: 0}).Find(&users)
// SELECT * FROM users WHERE name = "jinzhu";
```

You can use map to build the query condition, it will use all values

```
db.Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).Find(&users)
// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;
```

#### Specify Struct search fields

```
db.Where(&User{Name: "jinzhu"}, "name", "Age").Find(&users)
// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;
// specify the name and age fields. But ago is not in the User object, use the default value 0

db.Where(&User{Name: "jinzhu"}, "Age").Find(&users)
// SELECT * FROM users WHERE age = 0;

```

#### Inline Condition

Works similar to `Where`

```
db.First(&user, "id = ?", "string_primary_key")
// SELECT * FROM users WHERE id = 'string_primary_key';

db.Find(&user, "name = ?", "jinzhu")
// SELECT * FROM users WHERE name = "jinzhu";

db.Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

db.Find(&users, User{Age: 20})
// SELECT * FROM users WHERE age = 20;

db.Find(&users, map[string]interface{}{"age": 20})
// SELECT * FROM users WHERE age = 20;
```

#### Not Conditions

```
db.Not("name = ?", "jinzhu").First(&user)
// SELECT * FROM users WHERE NOT name = "jinzhu" ORDER BY id LIMIT 1;

db.Not(map[string]interface{}{"name": []string{"jinzhu", "jinzhu 2"}}).Find(&users)
// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");

db.Not(User{Name: "jinzhu", Age: 18}).First(&user)
// SELECT * FROM users WHERE name <> "jinzhu" AND age <> 18 ORDER BY id LIMIT 1;

// Not In slice of primary keys
db.Not([]int64{1,2,3}).First(&user)
// SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1;
```

#### Or Conditions

```
db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';

// Struct
db.Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2", Age: 18}).Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

// Map
db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2", "age": 18}).Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);
```

#### Selecting Specific Fields

```
db.Select("name", "age").Find(&users)
// SELECT name, age FROM users;

db.Select([]string{"name", "age"}).Find(&users)
// SELECT name, age FROM users;

db.Table("users").Select("COALESCE(age,?)", 42).Rows()
// SELECT COALESCE(age,'42') FROM users;
```

#### Order

```
db.Order("age desc, name").Find(&users)
// SELECT * FROM users ORDER BY age desc, name;

// Multiple orders
db.Order("age desc").Order("name").Find(&users)
// SELECT * FROM users ORDER BY age desc, name;

db.Clauses(clause.OrderBy{
  Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
}).Find(&User{})
// SELECT * FROM users ORDER BY FIELD(id,1,2,3)
```

#### Limit & Offset

`Limit` specify the max number of records to retrieve
`Offset` specify the number of records to skip before starting to return the records

```
db.Limit(3).Find(&users)
// SELECT * FROM users LIMIT 3;

// Cancel limit condition with -1
db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
// SELECT * FROM users LIMIT 10; (users1)
// SELECT * FROM users; (users2)

db.Offset(3).Find(&users)
// SELECT * FROM users OFFSET 3;

db.Limit(10).Offset(5).Find(&users)
// SELECT * FROM users OFFSET 5 LIMIT 10;

// Cancel offset condition with -1
db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
// SELECT * FROM users OFFSET 10; (users1)
// SELECT * FROM users; (users2)
```

#### Group & Having

```
type result struct {
  Date  time.Time
  Total int
}

db.Model(&User{}).Select("name, sum(age) as total").Where("name LIKE ?", "group%").Group("name").First(&result)
// SELECT name, sum(age) as total FROM `users` WHERE name LIKE "group%" GROUP BY `name`

db.Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "group").Find(&result)
// SELECT name, sum(age) as total FROM `users` GROUP BY `name` HAVING name = "group"

rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Rows()
for rows.Next() {
  ...
}

rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Rows()
for rows.Next() {
  ...
}

type Result struct {
  Date  time.Time
  Total int64
}
db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Scan(&results)
```

#### Distinct

```
db.Distinct("name", "age").Order("name, age desc").Find(&results)

```

#### Joins

```
type result struct {
  Name  string
  Email string
}

db.Model(&User{}).Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&result{})
// SELECT users.name, emails.email FROM `users` left join emails on emails.user_id = users.id

rows, err := db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Rows()
for rows.Next() {
  ...
}

db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)

// multiple joins with parameter
db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "411111111111").Find(&user)
```

#### Joins Preloading

`Join Preload` works with one-to-one relation, e.g: `has one`, `belongs to`

```
type User struct {
  gorm.Model
  Name      string
  CompanyID int
  Company   Company
}

type Company struct {
  gorm.Model
  Code string
  Name string
}

db.Joins("Company").Find(&users)

```

```
type User struct {
	gorm.Model
	Name string
	// `User` belongs to `Company`, `CompanyID` is the foreign key
	CompanyID int
	Company   Company
}

type Company struct {
	ID   int
	Name string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{}, &Company{})

	companies := []Company{
		{ID: 1, Name: "IBM"},
	}

	for _, c := range companies {
		db.Create(&c)
	}

	users := []User{
		{Name: "Goerge", CompanyID: 1},
		{Name: "John", CompanyID: 1},
	}

	for _, u := range users {
		db.Create(&u)
	}

	users_result := []User{}

	db.Joins("Company").Find(&users_result)

// SELECT `users`.`id`,
//        `users`.`name`,
//        `users`.`age`,
//        `Company`.`id` AS `Company__id`,
//        `Company`.`name` AS `Company__name`
// FROM `users` LEFT JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id`;

	for _, u := range users_result {
		fmt.Printf("name: %v; company name: %v\n", u.Name, u.Company.Name)
	}
}
```

```
name: Goerge; company name: IBM
name: John; company name: IBM
```

#### Scan

Scan results into a struct work similar to `Find`

```
type Result struct {
  Name string
  Age  int
}

var result Result
db.Table("users").Select("name", "age").Where("name = ?", "Antonio").Scan(&result)

// Raw SQL
db.Raw("SELECT name, age FROM users WHERE name = ?", "Antonio").Scan(&result)
```

#### Smart Select Fields

```
type User struct {
	ID     uint
	Name   string
	Age    int
	Gender string
	// hundreds of fields
}

// APIUser contains partial fields from User
type APIUser struct {
	ID   uint
	Name string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	users := []User{
		{Name: "Goerge", Age: 32, Gender: "Male"},
		{Name: "John", Age: 28, Gender: "Male"},
	}

	for _, u := range users {
		db.Create(&u)
	}

	partial_users := []APIUser{}

	db.Model(&User{}).Limit(10).Find(&partial_users)

	for _, u := range partial_users {
		fmt.Printf("name: %v\n", u.Name)
	}
}
```

```
name: Goerge
name: John
```

`QueryFields` mode will select by all fields’ name for current model

```
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  QueryFields: true,
})

db.Find(&user)
// SELECT `users`.`name`, `users`.`age`, ... FROM `users` // with this option

// Session Mode
db.Session(&gorm.Session{QueryFields: true}).Find(&user)
// SELECT `users`.`name`, `users`.`age`, ... FROM `users`
```

#### SELECT … FOR UPDATE

For index records the search encounters, locks the rows and any associated index entries, the same as if you issued an `UPDATE` statement for those rows. Other transactions are blocked from updating those rows, from doing `SELECT ... FOR SHARE`, or from reading the data in certain transaction isolation levels. Consistent reads ignore any locks set on the records that exist in the read view.

#### SELECT … FOR SHARE

Sets a shared mode lock on any rows that are read. Other sessions can read the rows, but cannot modify them until your transaction commits. If any of these rows were changed by another transaction that has not yet committed, your query waits until that transaction ends and then uses the latest values.

#### Locking

```
db.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&users)
// SELECT * FROM `users` FOR UPDATE

db.Clauses(clause.Locking{
  Strength: "SHARE",
  Table: clause.Table{Name: clause.CurrentTable},
}).Find(&users)
// SELECT * FROM `users` FOR SHARE OF `users`
```

#### SubQuery

A subquery can be nested within a query

```
type User struct {
	ID     uint
	Name   string
	Age    int
	Gender string
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	users := []User{
		{Name: "Goerge Philip", Age: 32, Gender: "Male"},
		{Name: "John Philip", Age: 10, Gender: "Male"},
		{Name: "Smith Philip", Age: 28, Gender: "Male"},
		{Name: "Duncan Philip", Age: 27, Gender: "Male"},
	}

	for _, u := range users {
		db.Create(&u)
	}

	db.Where("age > (?)", db.Table("users").Select("AVG(age)")).Find(&results)

	for _, u := range results {
		fmt.Println(u)
	}
}

{1 Goerge Philip 32 Male}
{3 Smith Philip 28 Male}
{4 Duncan Philip 27 Male}
```

```

type User struct {
	ID     uint
	Name   string
	Age    int
	Gender string

}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	users := []User{
		{Name: "Goerge Philip", Age: 32, Gender: "Male"},
		{Name: "John Philip", Age: 10, Gender: "Male"},
		{Name: "Smith Philip", Age: 28, Gender: "Male"},
		{Name: "Duncan Philip", Age: 27, Gender: "Male"},
	}

	for _, u := range users {
		db.Create(&u)
	}

	var result float32
	db.Table("users").Select("AVG(age)").Find(&result)
	fmt.Print(result)
}

24.25
```

```
type NameAgeAverage struct {
	Name   string
	Avgage float32
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	// users := []User{
	// 	{Name: "Goerge Philip", Age: 32, Gender: "Male"},
	// 	{Name: "John Philip", Age: 10, Gender: "Male"},
	// 	{Name: "Smith Philip", Age: 28, Gender: "Male"},
	// 	{Name: "Duncan Philip", Age: 27, Gender: "Male"},
	// }

	// for _, u := range users {
	// 	db.Create(&u)
	// }

	var results []NameAgeAverage

        // select name AVG(age) as avgage
        // from users
        // group by name;
	db.Table("users").Select("name, AVG(age) as avgage").Group("name").Find(&results)

	for _, u := range results {
		fmt.Println(u.Name, u.Avgage)
	}
}

Goerge Philip 32
John Philip 10
Smith Philip 28
Duncan Philip 27
```

```
type NameAgeAverage struct {
	Name   string
	Avgage float32
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	// users := []User{
	// 	{Name: "Goerge Philip", Age: 32, Gender: "Male"},
	// 	{Name: "John Philip", Age: 10, Gender: "Male"},
	// 	{Name: "Smith Philip", Age: 28, Gender: "Male"},
	// 	{Name: "Duncan Philip", Age: 27, Gender: "Male"},
	// }

	// for _, u := range users {
	// 	db.Create(&u)
	// }

	var results []NameAgeAverage

        // select name AVG(age) as avgage
        // from users
        // group by name
        // having AVG(age) > (select AVG(age) from users where name like '%Philip');
	subQuery := db.Select("AVG(age)").Where("name like ?", "%Philip").Table("users")
	db.Table("users").Select("name, AVG(age) as avgage").Group("name").Having("AVG(age) > (?)", subQuery).Find(&results)

	for _, u := range results {
		fmt.Println(u.Name, u.Avgage)
	}
}
```

#### From SubQuery

```

db.Table("(?) as u", db.Model(&User{}).Select("name", "age")).Where("age = ?", 18}).Find(&User{})
// SELECT * FROM (SELECT `name`,`age` FROM `users`) as u WHERE `age` = 18

subQuery1 := db.Model(&User{}).Select("name")
subQuery2 := db.Model(&Pet{}).Select("name")
db.Table("(?) as u, (?) as p", subQuery1, subQuery2).Find(&User{})
// SELECT * FROM (SELECT `name` FROM `users`) as u, (SELECT `name` FROM `pets`) as p

```

#### and/or group Conditions

```
db.Where(
  db.Where("pizza = ?", "pepperoni").Where(db.Where("size = ?", "small").Or("size = ?", "medium")),
).Or(
  db.Where("pizza = ?", "hawaiian").Where("size = ?", "xlarge"),
).Find(&Pizza{}).Statement

// SELECT *
// FROM `pizzas`
// WHERE (pizza = "pepperoni" AND (size = "small" OR size = "medium")) OR (pizza = "hawaiian" AND size = "xlarge")

```

#### Named Argument

```
db.Where("name1 = @name OR name2 = @name", sql.Named("name", "jinzhu")).Find(&user)
// SELECT * FROM `users` WHERE name1 = "jinzhu" OR name2 = "jinzhu"

db.Where("name1 = @name OR name2 = @name", map[string]interface{}{"name": "jinzhu"}).First(&user)
// SELECT * FROM `users` WHERE name1 = "jinzhu" OR name2 = "jinzhu" ORDER BY `users`.`id` LIMIT 1

```

#### Find To Map

```
var result map[string]interface{}
db.Model(&User{}).First(&result, "id = ?", 1)

var results []map[string]interface{}
db.Table("users").Find(&results)
```

#### First Or Init

Get first matched record or initialize a new instance with given conditions (only works with struct or map conditions)

```

// Found user with `name` = `jinzhu`
db.FirstOrInit(&user, map[string]interface{}{"name": "jinzhu"})
// user -> User{ID: 111, Name: "Jinzhu", Age: 18}
```

```
//Demo 1:
	var u User
	db.FirstOrInit(&u, User{Name: "David"})
	print(u.Name, u.Gender)

E:\Go\GORM_test>go run .
David
```

![](https://fixjourney.files.wordpress.com/2021/02/image-21.png?w=333)

Only initialize the variable, no record is inserted into table

```
//Demo 2:
	var u User
	db.Where(User{Name: "David"}).FirstOrInit(&u)
	print(u.Name, u.Gender)

E:\Go\GORM_test>go run .
David
```

![](https://fixjourney.files.wordpress.com/2021/02/image-22.png?w=333)

```
// Demo 3:
	var u User
	db.FirstOrInit(&u, map[string]interface{}{"name": "David"})
	print(u.Name, u.Gender)
E:\Go\GORM_test>go run .
David
```

Demo 1, 2 and 3 have same output. Only the code is different.

#### First Or Init with attribute

initialize struct with more attributes if record not found, those `Attrs` won’t be used to build SQL query

```
// User not found, initialize it with give conditions and Attrs
db.Where(User{Name: "non_existing"}).Attrs(User{Age: 20}).FirstOrInit(&user)
// SELECT * FROM USERS WHERE name = 'non_existing' ORDER BY id LIMIT 1;
// user -> User{Name: "non_existing", Age: 20}

// User not found, initialize it with give conditions and Attrs
db.Where(User{Name: "non_existing"}).Attrs("age", 20).FirstOrInit(&user)
// SELECT * FROM USERS WHERE name = 'non_existing' ORDER BY id LIMIT 1;
// user -> User{Name: "non_existing", Age: 20}

// Found user with `name` = `jinzhu`, attributes will be ignored
db.Where(User{Name: "Jinzhu"}).Attrs(User{Age: 20}).FirstOrInit(&user)
// SELECT * FROM USERS WHERE name = jinzhu' ORDER BY id LIMIT 1;
// user -> User{ID: 111, Name: "Jinzhu", Age: 18}
```

#### FIRST OR INIT WITH assign

`Assign` attributes to struct regardless it is found or not, those attributes won’t be used to build SQL query and the final data won’t be saved into database

```
// User not found, initialize it with give conditions and Assign attributes
db.Where(User{Name: "non_existing"}).Assign(User{Age: 20}).FirstOrInit(&user)
// user -> User{Name: "non_existing", Age: 20}

// Found user with `name` = `jinzhu`, update it with Assign attributes
db.Where(User{Name: "Jinzhu"}).Assign(User{Age: 20}).FirstOrInit(&user)
// SELECT * FROM USERS WHERE name = jinzhu' ORDER BY id LIMIT 1;
// user -> User{ID: 111, Name: "Jinzhu", Age: 20}
```

#### First Or Create

Get first matched record or create a new one with given conditions (only works with struct, map conditions)

```
// User not found, create a new record with give conditions
db.FirstOrCreate(&user, User{Name: "non_existing"})
// INSERT INTO "users" (name) VALUES ("non_existing");
// user -> User{ID: 112, Name: "non_existing"}

// Found user with `name` = `jinzhu`
db.Where(User{Name: "jinzhu"}).FirstOrCreate(&user)
// user -> User{ID: 111, Name: "jinzhu", "Age": 18}
```

#### First Or Create with attribute

Create struct with more attributes if record not found, those `Attrs` won’t be used to build SQL query

```
// User not found, create it with give conditions and Attrs
db.Where(User{Name: "non_existing"}).Attrs(User{Age: 20}).FirstOrCreate(&user)
// SELECT * FROM users WHERE name = 'non_existing' ORDER BY id LIMIT 1;
// INSERT INTO "users" (name, age) VALUES ("non_existing", 20);
// user -> User{ID: 112, Name: "non_existing", Age: 20}

// Found user with `name` = `jinzhu`, attributes will be ignored
db.Where(User{Name: "jinzhu"}).Attrs(User{Age: 20}).FirstOrCreate(&user)
// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;
// user -> User{ID: 111, Name: "jinzhu", Age: 18}
```

#### FIRST OR CREATE WITH assign

`Assign` attributes to the record regardless it is found or not and save them back to the database.

```
// User not found, initialize it with give conditions and Assign attributes
db.Where(User{Name: "non_existing"}).Assign(User{Age: 20}).FirstOrCreate(&user)
// SELECT * FROM users WHERE name = 'non_existing' ORDER BY id LIMIT 1;
// INSERT INTO "users" (name, age) VALUES ("non_existing", 20);
// user -> User{ID: 112, Name: "non_existing", Age: 20}

// Found user with `name` = `jinzhu`, update it with Assign attributes
db.Where(User{Name: "jinzhu"}).Assign(User{Age: 20}).FirstOrCreate(&user)
// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;
// UPDATE users SET age=20 WHERE id = 111;
// user -> User{ID: 111, Name: "jinzhu", Age: 20}
```

#### Optimizer/Index Hints

Optimizer Hints

optimizer hints apply on a per-statement basis

##### Optimizer Hint Syntax

Optimizer hints must be specified within `/*+ ... */` comments

```
/*+ BKA(t1) */
/*+ BNL(t1, t2) */
/*+ NO_RANGE_OPTIMIZATION(t4 PRIMARY) */
/*+ QB_NAME(qb2) */
```

import "gorm.io/hints"

db.Clauses(hints.New("MAX\_EXECUTION\_TIME(10000)")).Find(&User{})
// SELECT \* /\*+ MAX\_EXECUTION\_TIME(10000) \*/ FROM \`users\`
-- The *MAX\_EXECUTION\_TIME*( N ) hint sets a statement execution timeout of N milliseconds

```
// install gorm.io/hints
go get -x gorm.io/hints
```

```
	var users []User
	db.Where(User{Name: "John Philip"}).Clauses(hints.New("MAX_EXECUTION_TIME(10000)")).Find(&users)

	for _, u := range users {
		fmt.Println(u.Name)
	}
```

#### Iteration

```
	rows, err := db.Model(&User{}).Rows()
	defer rows.Close()

	for rows.Next() {
		var user User
		// ScanRows is a method of `gorm.DB`, it can be used to scan a row into a struct
		db.ScanRows(rows, &user)
		fmt.Println(user.Name, user.Age, user.Gender)
	}

Goerge Philip 32 Male
John Philip 10 Male
Smith Philip 28 Male
Duncan Philip 27 Male
```

#### FindInBatches

Query and process records in batch

```
func (db *DB) FindInBatches(dest interface{}, batchSize int, fc func(tx *DB, batch int) error) *DB
```

```
	var users []User

        // Batch size is 2. users table has 4 users. Each batch take 2 of them and run the func(tx *gorm.DB, batch int)
	result := db.FindInBatches(&users, 2, func(tx *gorm.DB, batch int) error {
		for i := range users {
			users[i].Name = users[i].Name + " appended"
		}

		tx.Save(users)

		fmt.Println("Rows affected:", tx.RowsAffected)
		fmt.Println(batch) // Batch 1, 2

		if tx.Error != nil {
			fmt.Println(tx.Error)
		}

		return nil
	})
	fmt.Println("result:", result.Error)

	fmt.Println("RowsAffected:", result.RowsAffected) // processed records count in all batches
```

```
Rows affected: 4
1
Rows affected: 4
2
result: <nil>
RowsAffected: 4

```

![](https://fixjourney.files.wordpress.com/2021/02/image-23.png?w=344)

#### Query Hooks

GORM allows hooks `AfterFind` for a query, it will be called when querying a record

```
type User struct {
	ID     uint
	Name   string
	Age    int
	Gender string
	Role   string
}

// Each time a user is retrieved from database, this function will be called.
func (u *User) AfterFind(tx *gorm.DB) (err error) {
	if u.Role == "" {
		u.Role = "user"
	}
	return
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	users := []User{
		{Name: "Goerge Philip", Age: 32, Gender: "Male"},
		{Name: "John Philip", Age: 10, Gender: "Male"},
		{Name: "Smith Philip", Age: 28, Gender: "Male"},
		{Name: "Duncan Philip", Age: 27, Gender: "Male"},
	}

	for _, u := range users {
		db.Create(&u)
	}

	var users_query []User
	db.Find(&users_query)

	for _, user := range users_query {
		fmt.Println(user.Name, user.Role)
	}
}

```

```
Goerge Philip user
John Philip user
Smith Philip user
Duncan Philip user

Here the func (u *User) AfterFind(tx *gorm.DB) (err error) is called four times.
```

#### Pluck

Query single column from database and scan into a slice
If you want to query multiple columns, use `Select` with [`Scan`](https://gorm.io/docs/query.html#scan) instead

```
	var usersQuery []User

	var ages []int64
	db.Model(&usersQuery).Pluck("age", &ages)
	println(ages)

ages = {32, 10, 28, 27}
```

![](https://fixjourney.files.wordpress.com/2021/02/image-24.png?w=277)

```
db.Table("deleted_users").Pluck("name", &names)

// Distinct Pluck
db.Model(&User{}).Distinct().Pluck("Name", &names)
// SELECT DISTINCT `name` FROM `users`

// Requesting more than one column, use `Scan` or `Find` like this:
db.Table("users").Select("name", "age").Scan(&users)
db.Table("users").Select("name", "age").Find(&users)
```

```
	var usersQuery []User
	db.Table("users").Select("name", "age").Scan(&usersQuery)

	for _, user := range usersQuery {
		fmt.Println(user.Name, user.Role, user.Age)
	}

Goerge Philip  32
John Philip  10
Smith Philip  28
Duncan Philip  27
```

#### Scopes

`Scopes` allows you to specify commonly-used queries which can be referenced as method calls

```
func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
  return db.Where("amount > ?", 1000)
}

func PaidWithCreditCard(db *gorm.DB) *gorm.DB {
  return db.Where("pay_mode_sign = ?", "C")
}

func PaidWithCod(db *gorm.DB) *gorm.DB {
  return db.Where("pay_mode_sign = ?", "C")
}

func OrderStatus(status []string) func (db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    return db.Where("status IN (?)", status)
  }
}

db.Scopes(AmountGreaterThan1000, PaidWithCreditCard).Find(&orders)
// Find all credit card orders and amount greater than 1000

db.Scopes(AmountGreaterThan1000, PaidWithCod).Find(&orders)
// Find all COD orders and amount greater than 1000

db.Scopes(AmountGreaterThan1000, OrderStatus([]string{"paid", "shipped"})).Find(&orders)
// Find all paid, shipped orders that amount greater than 1000
```

#### Count

Get matched records count

```
var count int64
db.Model(&User{}).Where("name = ?", "jinzhu").Or("name = ?", "jinzhu 2").Count(&count)
// SELECT count(1) FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2'

db.Model(&User{}).Where("name = ?", "jinzhu").Count(&count)
// SELECT count(1) FROM users WHERE name = 'jinzhu'; (count)

db.Table("deleted_users").Count(&count)
// SELECT count(1) FROM deleted_users;

// Count with Distinct
db.Model(&User{}).Distinct("name").Count(&count)
// SELECT COUNT(DISTINCT(`name`)) FROM `users`

db.Table("deleted_users").Select("count(distinct(name))").Count(&count)
// SELECT count(distinct(name)) FROM deleted_users

// Count with Group
users := []User{
  {Name: "name1"},
  {Name: "name2"},
  {Name: "name3"},
  {Name: "name3"},
}

db.Model(&User{}).Group("name").Count(&count)
count // => 3
```

## Update

#### Save All Fields

```
db.First(&user)

user.Name = "jinzhu 2"
user.Age = 100
db.Save(&user)
```

#### Update single column

When updating a single column with `Update`, it needs to have any conditions or it will raise error `ErrMissingWhereClause`
When using the `Model` method and its value has a primary value, the primary key will be used to build the condition

```
// Update with conditions
db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;

// User's ID is `111`:
db.Model(&user).Update("name", "hello")
// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

```

#### Updates multiple columns

`Updates` supports update with `struct` or `map[string]interface{}`, when updating with `struct` it will only update non-zero fields by default

```
// Update attributes with `struct`, will only update non-zero fields
db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

// Update attributes with `map`
db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
// UPDATE users SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

```

#### Update Selected Fields

If you want to update selected fields or ignore some fields when updating, you can use `Select`, `Omit`

```
// Select with Map
// User's ID is `111`:
db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
// UPDATE users SET name='hello' WHERE id=111;

db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
// UPDATE users SET age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

// Select with Struct (select zero value fields)
db.Model(&user).Select("Name", "Age").Updates(User{Name: "new_name", Age: 0})
// UPDATE users SET name='new_name', age=0 WHERE id=111;

// Select all fields (select all fields include zero value fields)
db.Model(&user).Select("*").Update(User{Name: "jinzhu", Role: "admin", Age: 0})

// Select all fields but omit Role (select all fields include zero value fields)
db.Model(&user).Select("*").Omit("Role").Update(User{Name: "jinzhu", Role: "admin", Age: 0})
```

#### Update Hooks

GORM allows hooks `BeforeSave`, `BeforeUpdate`, `AfterSave`, `AfterUpdate`, those methods will be called when updating a record, refer [Hooks](https://gorm.io/docs/hooks.html) for details

```
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
  if u.Role == "admin" {
    return errors.New("admin user not allowed to update")
  }
  return
}
```

#### Batch Updates

If we haven’t specified a record having primary key value with `Model`, GORM will perform a batch updates

```
// Update with struct
db.Model(User{}).Where("role = ?", "admin").Updates(User{Name: "hello", Age: 18})
// UPDATE users SET name='hello', age=18 WHERE role = 'admin;

// Update with map
db.Table("users").Where("id IN ?", []int{10, 11}).Updates(map[string]interface{}{"name": "hello", "age": 18})
// UPDATE users SET name='hello', age=18 WHERE id IN (10, 11);
```

#### Block Global Updates

If you perform a batch update without any conditions, GORM WON’T run it and will return `ErrMissingWhereClause` error by default

You have to use some conditions or use raw SQL or enable the `AllowGlobalUpdate` mode

```
db.Model(&User{}).Update("name", "jinzhu").Error // gorm.ErrMissingWhereClause

db.Model(&User{}).Where("1 = 1").Update("name", "jinzhu")
// UPDATE users SET `name` = "jinzhu" WHERE 1=1

db.Exec("UPDATE users SET name = ?", "jinzhu")
// UPDATE users SET name = "jinzhu"

db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&User{}).Update("name", "jinzhu")
// UPDATE users SET `name` = "jinzhu"
```

#### Updated Records Count

Get the number of rows affected by a update

```
// Get updated records count with `RowsAffected`
result := db.Model(User{}).Where("role = ?", "admin").Updates(User{Name: "hello", Age: 18})
// UPDATE users SET name='hello', age=18 WHERE role = 'admin;

result.RowsAffected // returns updated records count
result.Error        // returns updating error
```

#### Update with SQL Expression

```
// product's ID is `3`
db.Model(&product).Update("price", gorm.Expr("price * ? + ?", 2, 100))
// UPDATE "products" SET "price" = price * 2 + 100, "updated_at" = '2013-11-17 21:34:10' WHERE "id" = 3;

db.Model(&product).Updates(map[string]interface{}{"price": gorm.Expr("price * ? + ?", 2, 100)})
// UPDATE "products" SET "price" = price * 2 + 100, "updated_at" = '2013-11-17 21:34:10' WHERE "id" = 3;

db.Model(&product).UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = 3;

db.Model(&product).Where("quantity > 1").UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = 3 AND quantity > 1;
```

use BLOB type in database table

```
type User struct {
	ID       int
	Name     string
	Location Location
}

type Location struct {
	X, Y int
}

//CREATE TABLE `users` (`id` bigint AUTO_INCREMENT,`name` longtext,`location` geometry,PRIMARY KEY (`id`))
func (loc Location) GormDataType() string {
	return "geometry"
}

// save the Location data into geometry database field type (as BLOB type)
func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_PointFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%d %d)", loc.X, loc.Y)},
	}
}

// Scan implements the sql.Scanner interface
func (loc *Location) Scan(v interface{}) error {
	// Scan a value into struct from database driver

        // to be implemented
        // question in https://gorm.io/docs/data_types.html
	return nil
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	db.Create(&User{
		Name:     "WeiZhong",
		Location: Location{X: 50, Y: 50},
	})

	db.Model(&User{ID: 1}).Updates(User{
		Name:     "DavidZhong",
		Location: Location{X: 100, Y: 100},
	})

	var usersQuery []User
	db.Find(&usersQuery)

	for _, user := range usersQuery {
		fmt.Println(user.Name, user.Location.X, user.Location.Y)
	}
}
```

![](https://fixjourney.files.wordpress.com/2021/02/image-25.png?w=212)

#### Update from SubQuery

```
db.Model(&user).Update("company_name", db.Model(&Company{}).Select("name").Where("companies.id = users.company_id"))
// UPDATE "users" SET "company_name" = (SELECT name FROM companies WHERE companies.id = users.company_id);

db.Table("users as u").Where("name = ?", "jinzhu").Update("company_name", db.Table("companies as c").Select("name").Where("c.id = u.company_id"))

db.Table("users as u").Where("name = ?", "jinzhu").Updates(map[string]interface{}{}{"company_name": db.Table("companies as c").Select("name").Where("c.id = u.company_id")})
```

#### Without Hooks/Time Tracking

If you want to skip `Hooks` methods and don’t track the update time when updating, you can use `UpdateColumn`, `UpdateColumns`, it works like `Update`, `Updates`

```
// Update single column
db.Model(&user).UpdateColumn("name", "hello")
// UPDATE users SET name='hello' WHERE id = 111;

// Update multiple columns
db.Model(&user).UpdateColumns(User{Name: "hello", Age: 18})
// UPDATE users SET name='hello', age=18 WHERE id = 111;

// Update selected columns
db.Model(&user).Select("name", "age").UpdateColumns(User{Name: "hello", Age: 0})
// UPDATE users SET name='hello', age=0 WHERE id = 111;
```

#### Check Field has changed?

GORM provides `Changed` method could be used in **Before Update Hooks**, it will return the field changed or not

The `Changed` method only works with methods `Update`, `Updates`, and it only checks if the updating value from `Update` / `Updates` equals the model value, will return true if it is changed and not omitted

```
type User struct {
	ID          int
	Name        string
	Admin       bool
	Age         uint
	RefreshedAt time.Time
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {

	if tx.Statement.Changed("Name", "Admin") { // if Name or Role changed
		tx.Statement.SetColumn("Age", 18)
	}

	// if any fields changed
	if tx.Statement.Changed() {
		tx.Statement.SetColumn("RefreshedAt", time.Now())
	}
	return nil
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	db.Create(&User{
		Name:  "WeiZhong",
		Age:   30,
		Admin: true,
	})

	db.Model(&User{ID: 1}).Updates(User{
		Name:  "DavidZhong",
		Admin: false,
	})
}
```

![](https://fixjourney.files.wordpress.com/2021/02/image-26.png?w=386)

```
db.Model(&User{ID: 1, Name: "jinzhu"}).Updates(map[string]interface{"name": "jinzhu2"})
// Changed("Name") => true
db.Model(&User{ID: 1, Name: "jinzhu"}).Updates(map[string]interface{"name": "jinzhu"})
// Changed("Name") => false, `Name` not changed
db.Model(&User{ID: 1, Name: "jinzhu"}).Select("Admin").Updates(map[string]interface{
  "name": "jinzhu2", "admin": false,
})
// Changed("Name") => false, `Name` not selected to update

```

#### Change Updating Values

To change updating values in Before Hooks, you should use `SetColumn` unless it is a full updates with `Save`, for example:

```
type User struct {
	ID                int
	Password          string
	Code              string
	Name              string
	Age               uint
	Admin             bool

	RefreshedAt       time.Time
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {

	fmt.Println("In BeforeUpdate-------------------------")
	fmt.Println("user.Age: ", user.Age)
	fmt.Println("user.Name: ", user.Name)

	if tx.Statement.Changed("Code") {
		fmt.Println("Code is changed in BeforeUpdate-------------------------")
		tx.Statement.SetColumn("Age", user.Age+50)
	}

	// if any fields changed
	if tx.Statement.Changed() {
		tx.Statement.SetColumn("RefreshedAt", time.Now())
	}
	return nil
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	db.Create(&User{
		Name:  "WeiZhong",
		Age:   10,
		Admin: true,
	})

	db.Model(&User{ID: 1}).Updates(User{
		Name:  "DavidZhong",
		Admin: false,
		Code:  "USUC",
		Age:   5,
	})

	var usersQuery []User
	db.Find(&usersQuery)

	for _, user := range usersQuery {
		fmt.Println(user.Name, user.Age, user.Code, user.Admin)
	}
}
```

```
In BeforeUpdate-------------------------
user.Age:  0
user.Name:
Code is changed in BeforeUpdate-------------------------
DavidZhong 50 USUC true
```

![](https://fixjourney.files.wordpress.com/2021/02/image-28.png?w=595)

The field set in Before Hook (such as BeforeUpdate) via **SetColumn** will not be updated again.
In this case, the Age: 5 in update statement. However, in BeforeUpdate function, tx.Statement.**SetColumn**(“Age”, user.Age+50) where the user variable does not store any data. Therefore, it is 0 + 50 =50

So is before BeforeSave. user \*User will not store any data when do the update. The difference between BeforeUpdate is that BeforeSave is called twice. The first time is the user is created(saved) in database. The second time is update.

```
func (user *User) BeforeSave(tx *gorm.DB) (err error) {

	fmt.Println("In BeforeSave-------------------------")
	fmt.Println("user.Age: ", user.Age)
	fmt.Println("user.Name: ", user.Name)

	if tx.Statement.Changed("Code") {
		tx.Statement.SetColumn("Age", user.Age+20)
	}
	return nil
}

In BeforeSave-------------------------
user.Age:  10
user.Name:  WeiZhong
In BeforeSave-------------------------
user.Age:  0
user.Name:
DavidZhong 20 USUC true
```

## Delete

#### Delete a Record

When deleting a record, the deleted value needs to have primary key or it will trigger a [Batch Delete](https://gorm.io/docs/delete.html#batch_delete)

```
// Email's ID is `10`
db.Delete(&email)
// DELETE from emails where id = 10;

// Delete with additional conditions
db.Where("name = ?", "jinzhu").Delete(&email)
// DELETE from emails where id = 10 AND name = "jinzhu";

```

#### Delete with primary key

GORM allows to delete objects using primary key(s) with inline condition, it works with numbers

```
db.Delete(&User{}, 10)
// DELETE FROM users WHERE id = 10;

db.Delete(&User{}, "10")
// DELETE FROM users WHERE id = 10;

db.Delete(&users, []int{1,2,3})
// DELETE FROM users WHERE id IN (1,2,3);

```

#### Delete Hooks

GORM allows hooks `BeforeDelete`, `AfterDelete`, those methods will be called when deleting a record

```
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
  if u.Role == "admin" {
    return errors.New("admin user not allowed to delete")
  }
  return
}
```

#### Batch Delete

If you perform a batch delete without any conditions, GORM WON’T run it, and will return `ErrMissingWhereClause` error

You have to use some conditions or use raw SQL or enable `AllowGlobalUpdate` mode

```
db.Delete(&User{}).Error // gorm.ErrMissingWhereClause

db.Where("1 = 1").Delete(&User{})
// DELETE FROM `users` WHERE 1=1

db.Exec("DELETE FROM users")
// DELETE FROM users

db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})
// DELETE FROM users
```

#### Soft Delete

If your model includes a `gorm.DeletedAt` field (which is included in `gorm.Model`), it will get soft delete ability automatically!

When calling `Delete`, the record WON’T be removed from the database, but GORM will set the `DeletedAt`‘s value to the current time, and the data is not findable with normal Query methods anymore.

```
// user's ID is `111`
db.Delete(&user)
// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE id = 111;

// Batch Delete
db.Where("age = ?", 20).Delete(&User{})
// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;

// Soft deleted records will be ignored when querying
db.Where("age = 20").Find(&user)
// SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;
```

If you don’t want to include `gorm.Model`, you can enable the soft delete feature like:

```
type User struct {
  ID      int
  Deleted gorm.DeletedAt
  Name    string
}
```

#### Find soft deleted records

You can find soft deleted records with `Unscoped`

```
db.Unscoped().Where("age = 20").Find(&users)
// SELECT * FROM users WHERE age = 20;
```

```
type User struct {
	ID      uint
	Name    string
	Age     int
	Gender  string
	Deleted gorm.DeletedAt
}

func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	users := []User{
		{Name: "Goerge Philip", Age: 32, Gender: "Male"},
		{Name: "John Philip", Age: 10, Gender: "Male"},
		{Name: "Smith Philip", Age: 28, Gender: "Male"},
		{Name: "Duncan Philip", Age: 27, Gender: "Male"},
	}

	for _, u := range users {
		db.Create(&u)
	}

	db.Where("ID in ?", []int{1, 2}).Delete(&User{})

	var usersQuery []User
	db.Unscoped().Find(&usersQuery)
	for _, user := range usersQuery {
		fmt.Println(user.Name, user.Age, user.Gender)
	}
}

Goerge Philip 32 Male
John Philip 10 Male
Smith Philip 28 Male
Duncan Philip 27 Male
```

![](https://fixjourney.files.wordpress.com/2021/02/image-29.png?w=376)

#### Delete permanently

You can delete matched records permanently with `Unscoped`

```
db.Unscoped().Delete(&order)
// DELETE FROM orders WHERE id=10;

```

```
db.Unscoped().Where("ID in ?", []int{1, 2}).Delete(&User{})
```

![](https://fixjourney.files.wordpress.com/2021/02/image-30.png?w=308)

## SQL Builder

#### Raw SQL

Query Raw SQL with `Scan`

```
type Result struct {
  ID   int
  Name string
  Age  int
}

var result Result
db.Raw("SELECT id, name, age FROM users WHERE name = ?", 3).Scan(&result)

var age int
db.Raw("select sum(age) from users where role = ?", "admin").Scan(&age)

```

`Exec` with Raw SQL

```
db.Exec("DROP TABLE users")
db.Exec("UPDATE orders SET shipped_at=? WHERE id IN ?", time.Now(), []int64{1,2,3})

// Exec with SQL Expression
db.Exec("update users set money=? where name = ?", gorm.Expr("money * ? + ?", 10000, 1), "jinzhu")

```

#### Named Argument

GORM supports named arguments with [`sql.NamedArg`](https://tip.golang.org/pkg/database/sql/#NamedArg), `map[string]interface{}{}` or struct

```
db.Where("name1 = @name OR name2 = @name", sql.Named("name", "jinzhu")).Find(&user)
// SELECT * FROM `users` WHERE name1 = "jinzhu" OR name2 = "jinzhu"

db.Where("name1 = @name OR name2 = @name", map[string]interface{}{"name": "jinzhu2"}).First(&result3)
// SELECT * FROM `users` WHERE name1 = "jinzhu2" OR name2 = "jinzhu2" ORDER BY `users`.`id` LIMIT 1

// Named Argument with Raw SQL
db.Raw("SELECT * FROM users WHERE name1 = @name OR name2 = @name2 OR name3 = @name",
   sql.Named("name", "jinzhu1"), sql.Named("name2", "jinzhu2")).Find(&user)
// SELECT * FROM users WHERE name1 = "jinzhu1" OR name2 = "jinzhu2" OR name3 = "jinzhu1"

db.Exec("UPDATE users SET name1 = @name, name2 = @name2, name3 = @name",
   sql.Named("name", "jinzhunew"), sql.Named("name2", "jinzhunew2"))
// UPDATE users SET name1 = "jinzhunew", name2 = "jinzhunew2", name3 = "jinzhunew"

db.Raw("SELECT * FROM users WHERE (name1 = @name AND name3 = @name) AND name2 = @name2",
   map[string]interface{}{"name": "jinzhu", "name2": "jinzhu2"}).Find(&user)
// SELECT * FROM users WHERE (name1 = "jinzhu" AND name3 = "jinzhu") AND name2 = "jinzhu2"

type NamedArgument struct {
  Name string
  Name2 string
}

db.Raw("SELECT * FROM users WHERE (name1 = @Name AND name3 = @Name) AND name2 = @Name2",
   NamedArgument{Name: "jinzhu", Name2: "jinzhu2"}).Find(&user)
// SELECT * FROM users WHERE (name1 = "jinzhu" AND name3 = "jinzhu") AND name2 = "jinzhu2"
```

#### DryRun Mode

Generate `SQL` without executing, can be used to prepare or test generated SQL

```
type User struct {
	ID      uint
	Name    string
	Age     int
	Gender  string
	Deleted gorm.DeletedAt
}

func main() {
	dsn := "root:63598500@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	users := []User{
		{Name: "Goerge Philip", Age: 32, Gender: "Male"},
		{Name: "John Philip", Age: 10, Gender: "Male"},
		{Name: "Smith Philip", Age: 28, Gender: "Male"},
		{Name: "Duncan Philip", Age: 27, Gender: "Male"},
	}

	for _, u := range users {
		stmt := db.Session(&gorm.Session{DryRun: true}).Create(&u).Statement
		println(stmt.SQL.String())
		println(stmt.Vars)
		println(db.Dialector.Explain(stmt.SQL.String(), stmt.Vars...))
	}
}
```

```
INSERT INTO `users` (`name`,`age`,`gender`,`deleted`) VALUES (?,?,?,?)
[4/4]0xc0000a63c0
INSERT INTO `users` (`name`,`age`,`gender`,`deleted`) VALUES ('Goerge Philip',32,'Male',NULL)
INSERT INTO `users` (`name`,`age`,`gender`,`deleted`) VALUES (?,?,?,?)
[4/4]0xc000202240
INSERT INTO `users` (`name`,`age`,`gender`,`deleted`) VALUES ('John Philip',10,'Male',NULL)
INSERT INTO `users` (`name`,`age`,`gender`,`deleted`) VALUES (?,?,?,?)
[4/4]0xc0000a6640
INSERT INTO `users` (`name`,`age`,`gender`,`deleted`) VALUES ('Smith Philip',28,'Male',NULL)
INSERT INTO `users` (`name`,`age`,`gender`,`deleted`) VALUES (?,?,?,?)
[4/4]0xc0000a68c0
INSERT INTO `users` (`name`,`age`,`gender`,`deleted`) VALUES ('Duncan Philip',27,'Male',NULL)
```

#### `Row` & `Rows`

**Row**

```
// Use GORM API build SQL
row := db.Table("users").Where("name = ?", "jinzhu").Select("name", "age").Row()
row.Scan(&name, &age)

// Use Raw SQL
row := db.Raw("select name, age, email from users where name = ?", "jinzhu").Row()
row.Scan(&name, &age, &email)
```

**Rows**

```
// Use GORM API build SQL
rows, err := db.Model(&User{}).Where("name = ?", "jinzhu").Select("name, age, email").Rows()
defer rows.Close()
for rows.Next() {
  rows.Scan(&name, &age, &email)

  // do something
}

// Raw SQL
rows, err := db.Raw("select name, age, email from users where name = ?", "jinzhu").Rows()
defer rows.Close()
for rows.Next() {
  rows.Scan(&name, &age, &email)

  // do something
}
```

#### Scan `*sql.Rows` into struct

Use `ScanRows` to scan a row into a struct.

```
rows, err := db.Model(&User{}).Where("name = ?", "jinzhu").Select("name, age, email").Rows() // (*sql.Rows, error)
defer rows.Close()

var user User
for rows.Next() {
  // ScanRows scan a row into user
  db.ScanRows(rows, &user)

  // do something
}
```

#### Clauses

GORM uses SQL builder generates SQL internally, for each operation, GORM creates a `*gorm.Statement` object, all GORM APIs add/change `Clause` for the `Statement`, at last, GORM generated SQL based on those clauses

```
clause.Select{Columns: "*"}
clause.From{Tables: clause.CurrentTable}
clause.Limit{Limit: 1}
clause.OrderByColumn{
  Column: clause.Column{Table: clause.CurrentTable, Name: clause.PrimaryKey},
}
```

Then GORM build finally querying SQL in the `Query` callbacks like:

```
Statement.Build("SELECT", "FROM", "WHERE", "GROUP BY", "ORDER BY", "LIMIT", "FOR")

```

Which generate SQL:

```
SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

```

#### Clause Options

GORM defined [Many Clauses](https://github.com/go-gorm/gorm/tree/master/clause), and some clauses provide advanced options can be used for your application

Although most of them are rarely used, if you find GORM public API can’t match your requirements, may be good to check them out

```
db.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&user)
// INSERT IGNORE INTO users (name,age...) VALUES ("jinzhu",18...);

```

#### StatementModifier

GORM provides interface [StatementModifier](https://pkg.go.dev/gorm.io/gorm?tab=doc#StatementModifier) allows you modify statement to match your requirements, take [Hints](https://gorm.io/docs/hints.html) as example

```
import "gorm.io/hints"

db.Clauses(hints.New("hint")).Find(&User{})
// SELECT * /*+ hint */ FROM `users`
```

The hint will give the query a hint of the best way to run.
If there a large number of query return, scan will be a good option. If there is only a few return, index seek will a good option for hint. Some time the database does not know what the best way to run the SQL statement. Here is the hint used for.

## Context

GORM provides Context support, you can use it with method `WithContext`

#### Single Session Mode

Single session mode usually used when you want to perform a single operation

```
	var usersQuery []User

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	db.WithContext(ctx).Find(&usersQuery)
	for _, user := range usersQuery {
		fmt.Println(user.Name, user.Age)
	}
	cancel()

```

#### Continuous session mode

Continuous session mode usually used when you want to perform a group of operations, for example:

```
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	var user User
	tx := db.WithContext(ctx)
	tx.First(&user, 1)
	tx.Model(&user).Update("age", "18")

	cancel()

```

#### Context in Hooks/Callbacks

You could access the `Context` object from current `Statement`

```
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
  ctx := tx.Statement.Context
  // ...
  return
}
```

#### Chi Middleware Example

Continuous session mode which might be helpful when handling API requests, for example, you can set up `*gorm.DB` with Timeout Context in middlewares, and then use the `*gorm.DB` when processing all requests

```
>go get -x "github.com/go-chi/chi"
>go get -x "github.com/go-chi/chi/middleware"
```

DO NOT KNOW THE chi.NewRouter

```
func SetDBMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    timeoutContext, _ := context.WithTimeout(context.Background(), time.Second)
    ctx := context.WithValue(r.Context(), "DB", db.WithContext(timeoutContext))
    next.ServeHTTP(w, r.WithContext(ctx))
  })
}

r := chi.NewRouter()
r.Use(SetDBMiddleware)

r.Get("/", func(w http.ResponseWriter, r *http.Request) {
  db, ok := ctx.Value("DB").(*gorm.DB)

  var users []User
  db.Find(&users)

  // lots of db operations
})

r.Get("/user", func(w http.ResponseWriter, r *http.Request) {
  db, ok := ctx.Value("DB").(*gorm.DB)

  var user User
  db.First(&user)

  // lots of db operations
})
```

## Error Handling

```
if err := db.Where("name = ?", "jinzhu").First(&user).Error; err != nil {
  // error handling...
}
```

Or

```
if result := db.Where("name = ?", "jinzhu").First(&user); result.Error != nil {
  // error handling...
}
```

#### ErrRecordNotFound

GORM returns `ErrRecordNotFound` when failed to find data with `First`, `Last`, `Take`, if there are several errors happened, you can check the `ErrRecordNotFound` error with `errors.Is`

```
// Check if returns RecordNotFound error
err := db.First(&user, 100).Error
errors.Is(err, gorm.ErrRecordNotFound)
```

## Method Chaining

```
db.Where("name = ?", "jinzhu").Where("age = ?", 18).First(&user)

```

#### Chain Method

Chain methods are methods to modify or add `Clauses` to current `Statement`

`Where`, `Select`, `Omit`, `Joins`, `Scopes`, `Preload`, `Raw` (`Raw` can’t be used with other chainable methods to build SQL)…

#### Finisher Method

Finishers are immediate methods that execute registered callbacks, which will generate and execute SQL, like those methods:

`Create`, `First`, `Find`, `Take`, `Save`, `Update`, `Delete`, `Scan`, `Row`, `Rows`…

#### New Session Mode

After a new initialized `*gorm.DB` or the `Session Method` is called, it call will create a new `Statement` instance instead of using the current one.

```
// Example 1:

	users := []User{
		{Name: "Goerge Philip", Age: 32, Gender: "Male"},
		{Name: "John Philip", Age: 10, Gender: "Male"},
		{Name: "Smith Philip", Age: 28, Gender: "Male"},
		{Name: "Duncan Philip", Age: 27, Gender: "Male"},
	}

	for _, u := range users {
		db.Create(&u)
	}

	var usersQuery []User
	db.Where("age > ?", 18).Find(&usersQuery)
	for _, user := range usersQuery {
		fmt.Println(user.Name, user.Age)
	}
	fmt.Println("----------------------")

	db.Where("name = ?", "Goerge Philip").Or("name = ?", "John Philip").Find(&usersQuery)

	for _, user := range usersQuery {
		fmt.Println(user.Name, user.Age)
	}
	fmt.Println("----------------------")

Goerge Philip 32
Smith Philip 28
Duncan Philip 27
----------------------
Goerge Philip 32
John Philip 10
----------------------
```

```
	var usersQuery []User
	db.Where("age > ?", 18)

	fmt.Println("----------------------")

	db.Where("name = ?", "Goerge Philip").Or("name = ?", "John Philip").Find(&usersQuery)

	for _, user := range usersQuery {
		fmt.Println(user.Name, user.Age)
	}
	fmt.Println("----------------------")

----------------------
Goerge Philip 32
John Philip 10
----------------------
Note: John Philip is under 18, but still in the query result.
```

```
	var usersQuery []User
	tx := db.Where("age > ?", 18)

	fmt.Println("----------------------")

	tx.Where("name = ?", "Goerge Philip").Or("name = ?", "John Philip").Find(&usersQuery)

	for _, user := range usersQuery {
		fmt.Println(user.Name, user.Age)
	}
	fmt.Println("----------------------")

----------------------
Goerge Philip 32
John Philip 10
----------------------
```

GORM defined `Session`, `WithContext`, `Debug` methods as `New Session Method`.

```
// Example 2
db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
// db is a new initialized *gorm.DB, which falls under `New Session Mode`
tx := db.Where("name = ?", "jinzhu")
// `Where("name = ?", "jinzhu")` is the first method call, it creates a new `Statement` and adds conditions

tx.Where("age = ?", 18).Find(&users)
// `tx.Where("age = ?", 18)` REUSES above `Statement`, and adds conditions to the `Statement`
// `Find(&users)` is a finisher, it executes registered Query Callbacks, generates and runs the following SQL:
// SELECT * FROM users WHERE name = 'jinzhu' AND age = 18

tx.Where("age = ?", 28).Find(&users)
// `tx.Where("age = ?", 18)` REUSES above `Statement` also, and add conditions to the `Statement`
// `Find(&users)` is a finisher, it executes registered Query Callbacks, generates and runs the following SQL:
// SELECT * FROM users WHERE name = 'jinzhu' AND age = 18 AND age = 20
```

**What I do not understand?**

tx := db.Where(“name = ?”, “jinzhu”) tx type is also is \*gorm.DB. Why tx.Where(“age = ?”, 28).Find(&users) will reuse the statement whereas db.Where(“name = ?”, “jinzhu2”).Where(“age = ?”, 20).Find(&users) will create a new session?

have asked in [https://gorm.io/docs/method\_chaining.html](https://gorm.io/docs/method_chaining.html)

#### Method Chain Safety/Goroutine Safety

Methods will create new `Statement` instances for new initialized `*gorm.DB` or after a `New Session Method`, so to reuse a `*gorm.DB` you need to make sure they are under `New Session Mode`

```
db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

// Safe for a new initialized *gorm.DB
for i := 0; i < 100; i++ {
  go db.Where(...).First(&user)
}

tx := db.Where("name = ?", "jinzhu")
// NOT Safe as reusing Statement
for i := 0; i < 100; i++ {
  go tx.Where(...).First(&user)
}

ctx, _ := context.WithTimeout(context.Background(), time.Second)
ctxDB := db.WithContext(ctx)
// Safe after a `New Session Method`
for i := 0; i < 100; i++ {
  go ctxDB.Where(...).First(&user)
}

ctx, _ := context.WithTimeout(context.Background(), time.Second)
ctxDB := db.Where("name = ?", "jinzhu").WithContext(ctx)
// Safe after a `New Session Method`
for i := 0; i < 100; i++ {
  go ctxDB.Where(...).First(&user) // `name = 'jinzhu'` will apply to the query
}

tx := db.Where("name = ?", "jinzhu").Session(&gorm.Session{})
// Safe after a `New Session Method`
for i := 0; i < 100; i++ {
  go tx.Where(...).First(&user) // `name = 'jinzhu'` will apply to the query
}
```

## Session

GORM provides `Session` method, which is a [`New Session Method`](https://gorm.io/docs/method_chaining.html), it allows to create a new session mode with configuration:

```
// Any functions below such as DryRun being called will create a new session instead of using the current session.
// Session Configuration
type Session struct {
  DryRun                   bool
  PrepareStmt              bool
  NewDB                    bool
  SkipHooks                bool
  SkipDefaultTransaction   bool
  DisableNestedTransaction bool
  AllowGlobalUpdate        bool
  FullSaveAssociations     bool
  QueryFields              bool
  CreateBatchSize          int
  Context                  context.Context
  Logger                   logger.Interface
  NowFunc                  func() time.Time
}

For example:
stmt := db.Session(&gorm.Session{DryRun: true}).Create(&u).Statement
the code above will create a new session.
```

#### PrepareStmt

`PreparedStmt` creates prepared statements when executing any SQL and caches them to speed up future calls

```
type PreparedStmtDB struct {
	Stmts       map[string]Stmt
	PreparedSQL []string
	Mux         *sync.RWMutex
	ConnPool
}

type Stmt struct {
	*sql.Stmt
	Transaction bool
}

type Stmt struct {
	// Immutable:
	db        *DB    // where we came from
	query     string // that created the Stmt
	stickyErr error  // if non-nil, this error is returned for all operations

	closemu sync.RWMutex // held exclusively during close, for read otherwise.

	// If Stmt is prepared on a Tx or Conn then cg is present and will
	// only ever grab a connection from cg.
	// If cg is nil then the Stmt must grab an arbitrary connection
	// from db and determine if it must prepare the stmt again by
	// inspecting css.
	cg   stmtConnGrabber
	cgds *driverStmt

	// parentStmt is set when a transaction-specific statement
	// is requested from an identical statement prepared on the same
	// conn. parentStmt is used to track the dependency of this statement
	// on its originating ("parent") statement so that parentStmt may
	// be closed by the user without them having to know whether or not
	// any transactions are still using it.
	parentStmt *Stmt

	mu     sync.Mutex // protects the rest of the fields
	closed bool

	// css is a list of underlying driver statement interfaces
	// that are valid on particular connections. This is only
	// used if cg == nil and one is found that has idle
	// connections. If cg != nil, cgds is always used.
	css []connStmt

	// lastNumClosed is copied from db.numClosed when Stmt is created
	// without tx and closed connections in css are removed.
	lastNumClosed uint64
}
```

```
func main() {
	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"

        // globally mode, all DB operations will create prepared statements and cache them
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	users := []User{
		{Name: "Goerge Philip", Age: 32, Gender: "Male"},
		{Name: "John Philip", Age: 10, Gender: "Male"},
		{Name: "Smith Philip", Age: 28, Gender: "Male"},
		{Name: "Duncan Philip", Age: 27, Gender: "Male"},
	}

	for _, u := range users {
		db.Create(&u)
	}

	// tx is *gorm.DB type
        // session mode
	tx := db.Session(&gorm.Session{PrepareStmt: true})

        var user User
	tx.First(&user, 1)

	var usersQuery []User
	tx.Find(&usersQuery)

	tx.Model(&user).Update("Age", 2)

	// returns prepared statements manager
	stmtManger, _ := tx.ConnPool.(*gorm.PreparedStmtDB)

	// close prepared statements for *current session*
	stmtManger.Close()

	// prepared SQL for *current session*
	for _, presql := range stmtManger.PreparedSQL {
		fmt.Println(presql)
	}

	// prepared statements for current database connection pool (all sessions)
	// stmtManger.Stmts // map[string]*sql.Stmt

	for sqlKey, stmt := range stmtManger.Stmts {
		fmt.Println("Key: ", sqlKey) // prepared SQL
		stmt.Close()                 // close the prepared statement
	}
```

```
SELECT * FROM `users` WHERE `users`.`id` = ? AND `users`.`deleted` IS NULL ORDER BY `users`.`id` LIMIT 1                                            NULL,PRIMARY KEY (`id`))
SELECT * FROM `users` WHERE `users`.`deleted` IS NULL
UPDATE `users` SET `age`=? WHERE `id` = ?
Key:  INSERT INTO `users` (`name`,`age`,`gender`,`deleted`) VALUES (?,?,?,?)
Key:  SELECT DATABASE()
Key:  SELECT count(*) FROM information_schema.tables WHERE table_schema = ? AND table_name = ? AND table_type = ?
Key:  CREATE TABLE `users` (`id` bigint unsigned AUTO_INCREMENT,`name` longtext,`age` bigint,`gender` longtext,`deleted` datetime(3) NULL,PRIMARY KEY (`id`))
```

#### NewDB

Create a new DB without conditions with option `NewDB`

```
tx := db.Where("name = ?", "jinzhu").Session(&gorm.Session{NewDB: true})

tx.First(&user)
// SELECT * FROM users ORDER BY id LIMIT 1

tx.First(&user, "id = ?", 10)
// SELECT * FROM users WHERE id = 10 ORDER BY id

// Without option `NewDB`
tx2 := db.Where("name = ?", "jinzhu").Session(&gorm.Session{})
tx2.First(&user)
// SELECT * FROM users WHERE name = "jinzhu" ORDER BY id
```

#### Skip Hooks

If you want to skip `Hooks` methods, you can use the `SkipHooks` session mode

```
DB.Session(&gorm.Session{SkipHooks: true}).Create(&user)

DB.Session(&gorm.Session{SkipHooks: true}).Create(&users)

// each batch insert 100 users. If there are 1000 user in users[], it will trigger 10 times of batch.
DB.Session(&gorm.Session{SkipHooks: true}).CreateInBatches(users, 100)

DB.Session(&gorm.Session{SkipHooks: true}).Find(&user)

DB.Session(&gorm.Session{SkipHooks: true}).Delete(&user)

DB.Session(&gorm.Session{SkipHooks: true}).Model(User{}).Where("age > ?", 18).Updates(&user)

```

#### DisableNestedTransaction

When using `Transaction` method inside a DB transaction, GORM will use `SavePoint(savedPointName)`, `RollbackTo(savedPointName)` to give you the nested transaction support. You can disable it by using the `DisableNestedTransaction` option

```
db.Session(&gorm.Session{
  DisableNestedTransaction: true,
}).CreateInBatches(&users, 100)

```

#### AllowGlobalUpdate

GORM doesn’t allow global update/delete by default, will return `ErrMissingWhereClause` error. You can set this option to true to enable it

```
db.Session(&gorm.Session{
  AllowGlobalUpdate: true,
}).Model(&User{}).Update("name", "jinzhu")
// UPDATE users SET `name` = "jinzhu"
```

#### FullSaveAssociations

GORM will auto-save associations and its reference using [Upsert](https://gorm.io/docs/create.html#upsert) when creating/updating a record. If you want to update associations’ data, you should use the `FullSaveAssociations` mode

```
db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
// ...
// INSERT INTO "addresses" (address1) VALUES ("Billing Address - Address 1"), ("Shipping Address - Address 1") ON DUPLICATE KEY SET address1=VALUES(address1);
// INSERT INTO "users" (name,billing_address_id,shipping_address_id) VALUES ("jinzhu", 1, 2);
// INSERT INTO "emails" (user_id,email) VALUES (111, "jinzhu@example.com"), (111, "jinzhu-2@example.com") ON DUPLICATE KEY SET email=VALUES(email);
```

#### Context

With the `Context` option, you can set the `Context` for following SQL operations

```
timeoutCtx, _ := context.WithTimeout(context.Background(), time.Second)
tx := db.Session(&Session{Context: timeoutCtx})

tx.First(&user) // query with context timeoutCtx
tx.Model(&user).Update("role", "admin") // update with context timeoutCtx

```

GORM also provides shortcut method `WithContext`, here is the definition:

```
func (db *DB) WithContext(ctx context.Context) *DB {
  return db.Session(&Session{Context: ctx})
}
```

#### NowFunc

`NowFunc` allows changing the function to get current time of GORM, for example:

```
db.Session(&Session{
  NowFunc: func() time.Time {
    return time.Now().Local()
  },
})
```

#### Debug

`Debug` is a shortcut method to change session’s `Logger` to debug mode

```
func (db *DB) Debug() (tx *DB) {
  return db.Session(&Session{
    Logger:         db.Logger.LogMode(logger.Info),
  })
}
```

#### QueryFields

Select by fields

```
db.Session(&gorm.Session{QueryFields: true}).Find(&user)
// SELECT `users`.`name`, `users`.`age`, ... FROM `users` // with this option
// SELECT * FROM `users` // without this option
```

#### CreateBatchSize

Default batch size

```
users = [5000]User{{Name: "jinzhu", Pets: []Pet{pet1, pet2, pet3}}...}

db.Session(&gorm.Session{CreateBatchSize: 1000}).Create(&users)
// INSERT INTO users xxx (5 batches)
// INSERT INTO pets xxx (15 batches)
```

## Hooks

Hooks are functions that are called before or after creation/querying/updating/deletion.

#### Creating an object

```
// begin transaction
BeforeSave
BeforeCreate
// save before associations
// insert into database
// save after associations
AfterCreate
AfterSave
// commit or rollback transaction

```

```
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
  u.UUID = uuid.New()

  if !u.IsValid() {
    err = errors.New("can't save invalid data")
  }
  return
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
  if u.ID == 1 {
    tx.Model(u).Update("role", "admin")
  }
  return
}
```

Save/Delete operations in GORM are running in transactions by default, so changes made in that transaction are not visible until it is committed, if you return any error in your hooks, the change will be rollbacked

```
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
  if !u.IsValid() {
    return errors.New("rollback invalid user")
  }
  return nil
}
```

#### Updating an object

Available hooks for updating

```
// begin transaction
BeforeSave
BeforeUpdate
// save before associations
// update database
// save after associations
AfterUpdate
AfterSave
// commit or rollback transaction

```

```
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
  if u.readonly() {
    err = errors.New("read only user")
  }
  return
}

// Updating data in same transaction
func (u *User) AfterUpdate(tx *gorm.DB) (err error) {
  if u.Confirmed {
    tx.Model(&Address{}).Where("user_id = ?", u.ID).Update("verfied", true)
  }
  return
}
```

#### Deleting an object

Available hooks for deleting

```
// begin transaction
BeforeDelete
// delete from database
AfterDelete
// commit or rollback transaction

```

```
// Updating data in same transaction
func (u *User) AfterDelete(tx *gorm.DB) (err error) {
  if u.Confirmed {
    tx.Model(&Address{}).Where("user_id = ?", u.ID).Update("invalid", false)
  }
  return
}
```

#### Querying an object

Available hooks for querying

```
// load data from database
// Preloading (eager loading)
AfterFind
```

```
func (u *User) AfterFind(tx *gorm.DB) (err error) {
  if u.MemberShip == "" {
    u.MemberShip = "user"
  }
  return
}
```

#### Modify current operation

```
func (u *User) BeforeCreate(tx *gorm.DB) error {
  // Modify current operation through tx.Statement, e.g:
  tx.Statement.Select("Name", "Age")
  tx.Statement.AddClause(clause.OnConflict{DoNothing: true})

  // tx is new session mode with the `NewDB` option
  // operations based on it will run inside same transaction but without any current conditions
  var role Role
  err := tx.First(&role, "name = ?", user.Role).Error
  // SELECT * FROM roles WHERE name = "admin"
  // ...
  return err
}
```

## Transactions

#### Disable Default Transaction

GORM perform write (create/update/delete) operations run inside a transaction to ensure data consistency, you can disable it during initialization if it is not required, you will gain about 30%+ performance improvement after that

```
// Globally disable
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  SkipDefaultTransaction: true,
})

// Continuous session mode
tx := db.Session(&Session{SkipDefaultTransaction: true})
tx.First(&user, 1)
tx.Find(&users)
tx.Model(&user).Update("Age", 18)
```

#### Transaction

To perform a set of operations within a transaction, the general flow is as below.

```
db.Transaction(func(tx *gorm.DB) error {
  // do some database operations in the transaction (use 'tx' from this point, not 'db')
  if err := tx.Create(&Animal{Name: "Giraffe"}).Error; err != nil {
    // return any error will rollback
    return err
  }

  if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil {
    return err
  }

  // return nil will commit the whole transaction
  return nil
})
```

#### Nested Transactions

GORM supports nested transactions, you can rollback a subset of operations performed within the scope of a larger transaction

```
db.Transaction(func(tx *gorm.DB) error {
  tx.Create(&user1)

  tx.Transaction(func(tx2 *gorm.DB) error {
    tx2.Create(&user2)
    return errors.New("rollback user2") // Rollback user2
  })

  tx.Transaction(func(tx2 *gorm.DB) error {
    tx2.Create(&user3)
    return nil
  })

  return nil
})

// Commit user1, user3
```

#### Transactions by manual

```
// begin a transaction
tx := db.Begin()

// do some database operations in the transaction (use 'tx' from this point, not 'db')
tx.Create(...)

// ...

// rollback the transaction in case of error
tx.Rollback()

// Or commit the transaction
tx.Commit()
```

#### A Specific Example

```
func CreateAnimals(db *gorm.DB) error {
  // Note the use of tx as the database handle once you are within a transaction
  tx := db.Begin()
  defer func() {
    if r := recover(); r != nil {
      tx.Rollback()
    }
  }()

  if err := tx.Error; err != nil {
    return err
  }

  if err := tx.Create(&Animal{Name: "Giraffe"}).Error; err != nil {
     tx.Rollback()
     return err
  }

  if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil {
     tx.Rollback()
     return err
  }

  return tx.Commit().Error
}
```

#### SavePoint, RollbackTo

GORM provides `SavePoint`, `RollbackTo` to save points and roll back to a savepoint

```
tx := db.Begin()
tx.Create(&user1)

tx.SavePoint("sp1")
tx.Create(&user2)
tx.RollbackTo("sp1") // Rollback user2

tx.Commit() // Commit user1
```

## Migration

#### Auto Migration

Automatically migrate your schema, to keep your schema up to date

**NOTE:** AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes. It will change existing column’s type if its size, precision, nullable changed. It **WON’T** delete unused columns to protect your data.

```
db.AutoMigrate(&User{})

db.AutoMigrate(&User{}, &Product{}, &Order{})

// Add table suffix when creating tables
db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

```

AutoMigrate creates database foreign key constraints automatically, you can disable this feature during initialization, for example:

```
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  DisableForeignKeyConstraintWhenMigrating: true,
})
```

#### Migrator Interface

GORM provides a migrator interface, which contains unified API interfaces for each database that could be used to build your database-independent migrations, for example:

SQLite doesn’t support `ALTER COLUMN`, `DROP COLUMN`, GORM will create a new table as the one you are trying to change, copy all data, drop the old table, rename the new table

MySQL doesn’t support rename column, index for some versions, GORM will perform different SQL based on the MySQL version you are using

```
type Migrator interface {
  // AutoMigrate
  AutoMigrate(dst ...interface{}) error

  // Database
  CurrentDatabase() string
  FullDataTypeOf(*schema.Field) clause.Expr

  // Tables
  CreateTable(dst ...interface{}) error
  DropTable(dst ...interface{}) error
  HasTable(dst interface{}) bool
  RenameTable(oldName, newName interface{}) error

  // Columns
  AddColumn(dst interface{}, field string) error
  DropColumn(dst interface{}, field string) error
  AlterColumn(dst interface{}, field string) error
  HasColumn(dst interface{}, field string) bool
  RenameColumn(dst interface{}, oldName, field string) error
  MigrateColumn(dst interface{}, field *schema.Field, columnType *sql.ColumnType) error
  ColumnTypes(dst interface{}) ([]*sql.ColumnType, error)

  // Constraints
  CreateConstraint(dst interface{}, name string) error
  DropConstraint(dst interface{}, name string) error
  HasConstraint(dst interface{}, name string) bool

  // Indexes
  CreateIndex(dst interface{}, name string) error
  DropIndex(dst interface{}, name string) error
  HasIndex(dst interface{}, name string) bool
  RenameIndex(dst interface{}, oldName, newName string) error
}
```

#### CurrentDatabase

Returns current using database name

```
db.Migrator().CurrentDatabase()

```

#### Tables

```
// Create table for `User`
db.Migrator().CreateTable(&User{})

// Append "ENGINE=InnoDB" to the creating table SQL for `User`
db.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&User{})

// Check table for `User` exists or not
db.Migrator().HasTable(&User{})
db.Migrator().HasTable("users")

// Drop table if exists (will ignore or delete foreign key constraints when dropping)
db.Migrator().DropTable(&User{})
db.Migrator().DropTable("users")

// Rename old table to new table
db.Migrator().RenameTable(&User{}, &UserInfo{})
db.Migrator().RenameTable("users", "user_infos")
```

#### Columns

```
type User struct {
  Name string
}

// Add name field
db.Migrator().AddColumn(&User{}, "Name")
// Drop name field
db.Migrator().DropColumn(&User{}, "Name")
// Alter name field
db.Migrator().AlterColumn(&User{}, "Name")
// Check column exists
db.Migrator().HasColumn(&User{}, "Name")

type User struct {
  Name    string
  NewName string
}

// Rename column to new name
db.Migrator().RenameColumn(&User{}, "Name", "NewName")
db.Migrator().RenameColumn(&User{}, "name", "new_name")

// ColumnTypes
db.Migrator().ColumnTypes(&User{}) ([]*sql.ColumnType, error)

```

#### Constraints

```
type UserIndex struct {
  Name  string `gorm:"check:name_checker,name <> 'jinzhu'"`
}

// Create constraint
db.Migrator().CreateConstraint(&UserIndex{}, "name_checker")

// Drop constraint
db.Migrator().DropConstraint(&UserIndex{}, "name_checker")

// Check constraint exists
db.Migrator().HasConstraint(&UserIndex{}, "name_checker")
```

Create foreign keys for relations

```
type User struct {
  gorm.Model
  CreditCards []CreditCard
}

type CreditCard struct {
  gorm.Model
  Number string
  UserID uint
}

// create database foreign key for user & credit_cards
db.Migrator().CreateConstraint(&User{}, "CreditCards")
db.Migrator().CreateConstraint(&User{}, "fk_users_credit_cards")
// ALTER TABLE `credit_cards` ADD CONSTRAINT `fk_users_credit_cards` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)

// check database foreign key for user & credit_cards exists or not
db.Migrator().HasConstraint(&User{}, "CreditCards")
db.Migrator().HasConstraint(&User{}, "fk_users_credit_cards")

// drop database foreign key for user & credit_cards
db.Migrator().DropConstraint(&User{}, "CreditCards")
db.Migrator().DropConstraint(&User{}, "fk_users_credit_cards")
```

#### Indexes

```
type User struct {
  gorm.Model
  Name string `gorm:"size:255;index:idx_name,unique"`
}

// Create index for Name field
db.Migrator().CreateIndex(&User{}, "Name")
db.Migrator().CreateIndex(&User{}, "idx_name")

// Drop index for Name field
db.Migrator().DropIndex(&User{}, "Name")
db.Migrator().DropIndex(&User{}, "idx_name")

// Check Index exists
db.Migrator().HasIndex(&User{}, "Name")
db.Migrator().HasIndex(&User{}, "idx_name")

type User struct {
  gorm.Model
  Name  string `gorm:"size:255;index:idx_name,unique"`
  Name2 string `gorm:"size:255;index:idx_name_2,unique"`
}
// Rename index name
db.Migrator().RenameIndex(&User{}, "Name", "Name2")
db.Migrator().RenameIndex(&User{}, "idx_name", "idx_name_2")

```

## Logger

Gorm has a [default logger implementation](https://github.com/go-gorm/gorm/blob/master/logger/logger.go), it will print Slow SQL and happening errors by default

The logger accepts few options, you can customize it during initialization, for example:

```
type User struct {
	ID      uint
	Name    string
	Age     int
	Gender  string
	Deleted gorm.DeletedAt
}

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,
		},
	)

	dsn := "root:root_password@tcp(127.0.0.1:3306)/inventorydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	users := []User{
		{Name: "Goerge Philip", Age: 32, Gender: "Male"},
		{Name: "John Philip", Age: 10, Gender: "Male"},
		{Name: "Smith Philip", Age: 28, Gender: "Male"},
		{Name: "Duncan Philip", Age: 27, Gender: "Male"},
	}

	for _, u := range users {
		db.Create(&u)
	}

	var usersQuery []User

	fmt.Println("----------------------")

	db.Where("name = ?", "Goerge Philip").Or("name = ?", "John Philip").Find(&usersQuery)

	for _, user := range usersQuery {
		fmt.Println(user.Name, user.Age)
	}
	fmt.Println("----------------------")
}
```

```
2021/02/21 19:52:24 C:/Users/weizh/go/pkg/mod/gorm.io/driver/mysql@v1.0.3/mysql.go:52
[info] replacing callback `gorm:update` from C:/Users/weizh/go/pkg/mod/gorm.io/driver/mysql@v1.0.3/mysql.go:52

2021/02/21 19:52:24 E:/Go/GORM_test/main.go:163
[1.000ms] [rows:-] SELECT DATABASE()

2021/02/21 19:52:24 E:/Go/GORM_test/main.go:163
[3.019ms] [rows:-] SELECT count(*) FROM information_schema.tables WHERE table_schema = 'inventorydb' AND table_name = 'users' AND table_type = 'BASE TABLE'

2021/02/21 19:52:24 E:/Go/GORM_test/main.go:163
[269.280ms] [rows:0] CREATE TABLE `users` (`id` bigint unsigned AUTO_INCREMENT,`name` longtext,`age` bigint,`gender` longtext,`deleted` datetime(3) NULL,PRIMARY KEY (`id`))

2021/02/21 19:52:24 E:/Go/GORM_test/main.go:173
[43.882ms] [rows:1] INSERT INTO `users` (`name`,`age`,`gender`,`deleted`) VALUES ('Goerge Philip',32,'Male',NULL)

2021/02/21 19:52:24 E:/Go/GORM_test/main.go:173
[67.369ms] [rows:1] INSERT INTO `users` (`name`,`age`,`gender`,`deleted`) VALUES ('John Philip',10,'Male',NULL)

2021/02/21 19:52:24 E:/Go/GORM_test/main.go:173
[32.644ms] [rows:1] INSERT INTO `users` (`name`,`age`,`gender`,`deleted`) VALUES ('Smith Philip',28,'Male',NULL)

2021/02/21 19:52:24 E:/Go/GORM_test/main.go:173
[56.359ms] [rows:1] INSERT INTO `users` (`name`,`age`,`gender`,`deleted`) VALUES ('Duncan Philip',27,'Male',NULL)
----------------------

2021/02/21 19:52:24 E:/Go/GORM_test/main.go:181
[0.994ms] [rows:2] SELECT * FROM `users` WHERE (name = 'Goerge Philip' OR name = 'John Philip') AND `users`.`deleted` IS NULL
Goerge Philip 32
John Philip 10
----------------------
```

#### Log Levels

GORM defined log levels: `Silent`, `Error`, `Warn`, `Info`

```
db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
  Logger: logger.Default.LogMode(logger.Silent),
})
```

#### Debug

Debug a single operation, change current operation’s log level to logger.Info

```
	db.Debug().Where("name = ?", "Goerge Philip").Or("name = ?", "John Philip").Find(&usersQuery)
```

```
----------------------

2021/02/21 19:56:05 E:/Go/GORM_test/main.go:179
[0.999ms] [rows:2] SELECT * FROM `users` WHERE (name = 'Goerge Philip' OR name = 'John Philip') AND `users`.`deleted` IS NULL
Goerge Philip 32
John Philip 10
----------------------
```

#### Customize Logger

Refer to GORM’s [default logger](https://github.com/go-gorm/gorm/blob/master/logger/logger.go) for how to define your own one

The logger needs to implement the following interface, it accepts `context`, so you can use it for log tracing

```
type Interface interface {
  LogMode(LogLevel) Interface
  Info(context.Context, string, ...interface{})
  Warn(context.Context, string, ...interface{})
  Error(context.Context, string, ...interface{})
  Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error)
}
```

## Generic database interface sql.DB

GORM provides the method `DB` which returns a generic database interface [\*sql.DB](https://pkg.go.dev/database/sql#DB) from the current `*gorm.DB`

```
// Get generic database object sql.DB to use its functions
sqlDB, err := db.DB()

// Ping
sqlDB.Ping()

// Close
sqlDB.Close()

// Returns database statistics
sqlDB.Stats()
```

**NOTE** If the underlying database connection is not a `*sql.DB`, like in a transaction, it will returns error

#### Connection Pool

```
// Get generic database object sql.DB to use its functions
sqlDB, err := db.DB()

// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
sqlDB.SetMaxIdleConns(10)

// SetMaxOpenConns sets the maximum number of open connections to the database.
sqlDB.SetMaxOpenConns(100)

// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
sqlDB.SetConnMaxLifetime(time.Hour)
```

## Performance

GORM optimizes many things to improve the performance, the default performance should good for most applications, but there are still some tips for how to improve it for your application.

#### [Disable Default Transaction](https://gorm.io/docs/transactions.html)

GORM perform write (create/update/delete) operations run inside a transaction to ensure data consistency, which is bad for performance, you can disable it during initialization

```
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  SkipDefaultTransaction: true,
})
```

#### [Caches Prepared Statement](https://gorm.io/docs/session.html)

Creates a prepared statement when executing any SQL and caches them to speed up future calls

```
// Globally mode
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  PrepareStmt: true,
})

// Session mode
tx := db.Session(&Session{PrepareStmt: true})
tx.First(&user, 1)
tx.Find(&users)
tx.Model(&user).Update("Age", 18)
```

#### [SQL Builder with PreparedStmt](https://gorm.io/docs/sql_builder.html)

Prepared Statement works with RAW SQL also, for example:

```
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  PrepareStmt: true,
})

db.Raw("select sum(age) from users where role = ?", "admin").Scan(&age)

```

#### Select Fields

By default GORM select all fields when querying, you can use `Select` to specify fields you want

```
db.Select("Name", "Age").Find(&Users{})

```

Or define a smaller API struct to use the [smart select fields feature](https://gorm.io/docs/advanced_query.html)

```
type User struct {
  ID     uint
  Name   string
  Age    int
  Gender string
  // hundreds of fields
}

type APIUser struct {
  ID   uint
  Name string
}

// Select `id`, `name` automatically when query
db.Model(&User{}).Limit(10).Find(&APIUser{})
// SELECT `id`, `name` FROM `users` LIMIT 10
```

#### [Index Hints](https://gorm.io/docs/hints.html)

[Index](https://gorm.io/docs/indexes.html) is used to speed up data search and SQL query performance. `Index Hints` gives the optimizer information about how to choose indexes during query processing, which gives the flexibility to choose a more efficient execution plan than the optimizer

```
import "gorm.io/hints"

db.Clauses(hints.UseIndex("idx_user_name")).Find(&User{})
// SELECT * FROM `users` USE INDEX (`idx_user_name`)

db.Clauses(hints.ForceIndex("idx_user_name", "idx_user_id").ForJoin()).Find(&User{})
// SELECT * FROM `users` FORCE INDEX FOR JOIN (`idx_user_name`,`idx_user_id`)"

db.Clauses(
  hints.ForceIndex("idx_user_name", "idx_user_id").ForOrderBy(),
  hints.IgnoreIndex("idx_user_name").ForGroupBy(),
).Find(&User{})
// SELECT * FROM `users` FORCE INDEX FOR ORDER BY (`idx_user_name`,`idx_user_id`) IGNORE INDEX FOR GROUP BY (`idx_user_name`)"
```

## Customize Data Types

### Implements Customized Data Type

#### Scanner / Valuer

The customized data type has to implement the [Scanner](https://pkg.go.dev/database/sql#Scanner) and [Valuer](https://pkg.go.dev/database/sql/driver#Valuer) interfaces, so GORM knowns to how to receive/save it into the database

```
type JSON json.RawMessage

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *JSON) Scan(value interface{}) error {
  bytes, ok := value.([]byte)
  if !ok {
    return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
  }

  result := json.RawMessage{}
  err := json.Unmarshal(bytes, &result)
  *j = JSON(result)
  return err
}

// Value return json value, implement driver.Valuer interface
func (j JSON) Value() (driver.Value, error) {
  if len(j) == 0 {
    return nil, nil
  }
  return json.RawMessage(j).MarshalJSON()
}
```

#### GormDataTypeInterface

GORM will read column’s database type from [tag](https://gorm.io/docs/models.html#tags) `type`, if not found, will check if the struct implemented interface `GormDBDataTypeInterface` or `GormDataTypeInterface` and will use its result as data type

```
type GormDataTypeInterface interface {
  GormDataType() string
}

type GormDBDataTypeInterface interface {
  GormDBDataType(*gorm.DB, *schema.Field) string
}
```

The result of `GormDataType` will be used as the general data type and can be obtained from `schema.Field`‘s field `DataType`

```
func (JSON) GormDataType() string {
  return "json"
}

type User struct {
  Attrs JSON
}

func (user User) BeforeCreate(tx *gorm.DB) {
  field := tx.Statement.Schema.LookUpField("Attrs")
  if field.DataType == "json" {
    // do something
  }
}
```

`GormDBDataType` usually returns the right data type for current driver when migrating

```
func (JSON) GormDBDataType(db *gorm.DB, field *schema.Field) string {
  // use field.Tag, field.TagSettings gets field's tags
  // checkout https://github.com/go-gorm/gorm/blob/master/schema/field.go for all options

  // returns different database type based on driver name
  switch db.Dialector.Name() {
  case "mysql", "sqlite":
    return "JSON"
  case "postgres":
    return "JSONB"
  }
  return ""
}
```

If the struct hasn’t implemented the `GormDBDataTypeInterface` or `GormDataTypeInterface` interface, GORM will guess its data type from the struct’s first field

```
type NullString struct {
  String string // use the first field's data type
  Valid  bool
}

type User struct {
  Name NullString // data type will be string
}
```

#### GormValuerInterface

GORM provides a `GormValuerInterface` interface, which can allow to create/update from SQL Expr or value based on context,

```
// GORM Valuer interface
type GormValuerInterface interface {
  GormValue(ctx context.Context, db *gorm.DB) clause.Expr
}
```

##### Create/Update from SQL Expr

```
type Location struct {
  X, Y int
}

func (loc Location) GormDataType() string {
  return "geometry"
}

func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
  return clause.Expr{
    SQL:  "ST_PointFromText(?)",
    Vars: []interface{}{fmt.Sprintf("POINT(%d %d)", loc.X, loc.Y)},
  }
}

// Scan implements the sql.Scanner interface
func (loc *Location) Scan(v interface{}) error {
  // Scan a value into struct from database driver
}

type User struct {
  ID       int
  Name     string
  Location Location
}

db.Create(&User{
  Name:     "jinzhu",
  Location: Location{X: 100, Y: 100},
})
// INSERT INTO `users` (`name`,`point`) VALUES ("jinzhu",ST_PointFromText("POINT(100 100)"))

db.Model(&User{ID: 1}).Updates(User{
  Name:  "jinzhu",
  Location: Location{X: 100, Y: 100},
})
// UPDATE `user_with_points` SET `name`="jinzhu",`location`=ST_PointFromText("POINT(100 100)") WHERE `id` = 1

```

##### Value based on Context

If you want to create or update a value depends on current context, you can also implements the `GormValuerInterface` interface

```
type EncryptedString struct {
  Value string
}

func (es EncryptedString) GormValue(ctx context.Context, db *gorm.DB) (expr clause.Expr) {
  if encryptionKey, ok := ctx.Value("TenantEncryptionKey").(string); ok {
    return clause.Expr{SQL: "?", Vars: []interface{}{Encrypt(es.Value, encryptionKey)}}
  } else {
    db.AddError(errors.New("invalid encryption key"))
  }
  return
}
```

## Scopes

Scopes allow you to re-use commonly logic, the shared logic needs to defined as type `func(*gorm.DB) *gorm.DB`

#### Query

Scope examples for querying

```
func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
  return db.Where("amount > ?", 1000)
}

func PaidWithCreditCard(db *gorm.DB) *gorm.DB {
  return db.Where("pay_mode = ?", "card")
}

func PaidWithCod(db *gorm.DB) *gorm.DB {
  return db.Where("pay_mode = ?", "cod")
}

func OrderStatus(status []string) func (db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    return db.Scopes(AmountGreaterThan1000).Where("status IN (?)", status)
  }
}

db.Scopes(AmountGreaterThan1000, PaidWithCreditCard).Find(&orders)
// Find all credit card orders and amount greater than 1000

db.Scopes(AmountGreaterThan1000, PaidWithCod).Find(&orders)
// Find all COD orders and amount greater than 1000

db.Scopes(AmountGreaterThan1000, OrderStatus([]string{"paid", "shipped"})).Find(&orders)
// Find all paid, shipped orders that amount greater than 1000
```

#### Pagination

```
func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    page, _ := strconv.Atoi(r.Query("page"))
    if page == 0 {
      page = 1
    }

    pageSize, _ := strconv.Atoi(r.Query("page_size"))
    switch {
    case pageSize > 100:
      pageSize = 100
    case pageSize <= 0:
      pageSize = 10
    }

    offset := (page - 1) * pageSize
    return db.Offset(offset).Limit(pageSize)
  }
}

db.Scopes(Paginate(r)).Find(&users)
db.Scopes(Paginate(r)).Find(&articles)

```

#### Updates

```
func CurOrganization(r *http.Request) func(db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    org := r.Query("org")

    if org != "" {
      var organization Organization
      if db.Session(&Session{}).First(&organization, "name = ?", org).Error == nil {
        return db.Where("org_id = ?", org.ID)
      }
    }

    db.AddError("invalid organization")
    return db
  }
}

db.Model(&article).Scopes(CurOrganization(r)).Update("Name", "name 1")
// UPDATE articles SET name = "name 1" WHERE org_id = 111
db.Scopes(CurOrganization(r)).Delete(&Article{})
// DELETE FROM articles WHERE org_id = 111
```

## Conventions

#### `ID` as Primary Key

GORM uses the field with the name `ID` as the table’s primary key by default.

```
type User struct {
  ID   string // field named `ID` will be used as a primary field by default
  Name string
}
```

You can set other fields as primary key with tag `primaryKey`

```
// Set field `UUID` as primary field
type Animal struct {
  ID     int64
  UUID   string `gorm:"primaryKey"`
  Name   string
  Age    int64
}
```

#### Pluralized Table Name

 for struct `User`, its table name is `users` by convention

##### TableName

You can change the default table name by implementing the `Tabler` interface

```
type Tabler interface {
  TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (User) TableName() string {
  return "profiles"
}
```

`TableName` doesn’t allow dynamic name, its result will be cached for future, to use dynamic name, you can use `Scopes`, for example:

```
func UserTable(user User) func (tx *gorm.DB) *gorm.DB {
  return func (tx *gorm.DB) *gorm.DB {
    if user.Admin {
      return tx.Table("admin_users")
    }

    return tx.Table("users")
  }
}

db.Scopes(UserTable(user)).Create(&user)
```

##### Temporarily specify a name

```
// Create table `deleted_users` with struct User's fields
db.Table("deleted_users").AutoMigrate(&User{})

// Query data from another table
var deletedUsers []User
db.Table("deleted_users").Find(&deletedUsers)
// SELECT * FROM deleted_users;

db.Table("deleted_users").Where("name = ?", "jinzhu").Delete(&User{})
// DELETE FROM deleted_users WHERE name = 'jinzhu';
```

##### NamingStrategy

GORM allows users change the default naming conventions by overriding the default `NamingStrategy`, which is used to build `TableName`, `ColumnName`, `JoinTableName`, `RelationshipFKName`, `CheckerName`, `IndexName`

#### Column Name

Column db name uses the field’s name’s `snake_case` by convention.

```
type User struct {
  ID        uint      // column name is `id`
  Name      string    // column name is `name`
  Birthday  time.Time // column name is `birthday`
  CreatedAt time.Time // column name is `created_at`
}
```

You can override the column name with tag `column` or use [`NamingStrategy`](https://gorm.io/docs/conventions.html#naming_strategy)

```
type Animal struct {
  AnimalID int64     `gorm:"column:beast_id"`         // set name to `beast_id`
  Birthday time.Time `gorm:"column:day_of_the_beast"` // set name to `day_of_the_beast`
  Age      int64     `gorm:"column:age_of_the_beast"` // set name to `age_of_the_beast`
}
```

#### Timestamp Tracking

##### CreatedAt

For models having `CreatedAt` field, the field will be set to the current time when the record is first created if its value is zero

```
db.Create(&user) // set `CreatedAt` to current time

user2 := User{Name: "jinzhu", CreatedAt: time.Now()}
db.Create(&user2) // user2's `CreatedAt` won't be changed

// To change its value, you could use `Update`
db.Model(&user).Update("CreatedAt", time.Now())
```

##### UpdatedAt

For models having `UpdatedAt` field, the field will be set to the current time when the record is updated or created if its value is zero

```
db.Save(&user) // set `UpdatedAt` to current time

db.Model(&user).Update("name", "jinzhu") // will set `UpdatedAt` to current time

db.Model(&user).UpdateColumn("name", "jinzhu") // `UpdatedAt` won't be changed

user2 := User{Name: "jinzhu", UpdatedAt: time.Now()}
db.Create(&user2) // user2's `UpdatedAt` won't be changed when creating

user3 := User{Name: "jinzhu", UpdatedAt: time.Now()}
db.Save(&user3) // user3's `UpdatedAt` will change to current time when updating

```

## Settings

GORM provides `Set`, `Get`, `InstanceSet`, `InstanceGet` methods allow users pass values to [hooks](https://gorm.io/docs/hooks.html) or other methods

GORM uses this for some features, like pass creating table options when migrating table.

```
// Add table suffix when creating tables
db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

```

#### Set / Get

Use `Set` / `Get` pass settings to hooks methods

```
type User struct {
  gorm.Model
  CreditCard CreditCard
  // ...
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
  myValue, ok := tx.Get("my_value")
  // ok => true
  // myValue => 123
}

type CreditCard struct {
  gorm.Model
  // ...
}

func (card *CreditCard) BeforeCreate(tx *gorm.DB) error {
  myValue, ok := tx.Get("my_value")
  // ok => true
  // myValue => 123
}

myValue := 123
db.Set("my_value", myValue).Create(&User{})
```

#### InstanceSet / InstanceGet

Use `InstanceSet` / `InstanceGet` pass settings to current `*Statement`‘s hooks methods

```
type User struct {
  gorm.Model
  CreditCard CreditCard
  // ...
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
  myValue, ok := tx.InstanceGet("my_value")
  // ok => true
  // myValue => 123
}

type CreditCard struct {
  gorm.Model
  // ...
}

// When creating associations, GORM creates a new `*Statement`, so can't read other instance's settings
func (card *CreditCard) BeforeCreate(tx *gorm.DB) error {
  myValue, ok := tx.InstanceGet("my_value")
  // ok => false
  // myValue => nil
}

myValue := 123
db.InstanceSet("my_value", myValue).Create(&User{})
```
