package main

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func init() {
	DB = DBC()

}

func DBC() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}

type Products struct {
	Id         int64   `json:"id"`
	Offer      string  `json:"offer"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	VendorCode string  `json:"vendor_code"`
	Supplier   string  `json:"supplier"`
}

type Grps struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
	Supplier string `json:"supplier"`
}

type Offers struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Price string `json:"price"`
}

func Add(p []Products) {
	DB.Create(&p)
}
func AddGrps(g []Grps) {
	DB.Create(&g)
}
