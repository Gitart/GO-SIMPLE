package main

import (
    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Product -- Represents a product
type Product struct {
    gorm.Model
    Code  string
    Price uint
}

// TableName setting the table name
func (Product) TableName() string {
    return "allProducts"
}

func main() {
    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()
    var product Product
    rows, err := db.Model(&Product{}).Rows()
    defer rows.Close()
    if err != nil {
        panic(err)
    }
    for rows.Next() {
        db.ScanRows(rows, &product)
        fmt.Println(product)
    }
}
