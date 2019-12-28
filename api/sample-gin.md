package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

func Database() *gorm.DB {
	//open a db connection
	db, err := gorm.Open("mysql", "root:pass@tcp(127.0.0.1:8889)/gotest?parseTime=true")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func main() {

	//Migrate the schema
	db := Database()
	db.AutoMigrate(&Product{})
	router := gin.Default()
	router.GET("/", startPage)
	router.LoadHTMLGlob("templates/*")
	v1 := router.Group("/api/v1/")
	{
		v1.POST("product/", CreateProduct)
		v1.GET("product/", FetchAllProduct)
		v1.GET("product/:id", FetchSingleProduct)
		v1.PUT("product/:id", UpdateProduct)
		v1.DELETE("product/:id", DeleteProduct)
	}
	router.Run()

}

type Product struct {
	gorm.Model
	Name     	string 	  `json:"name"`
	Description string    `json:"description"`
	Images  	string    `json:"images"`
	Price 		string    `json:"price"`
}

type TransformedProduct struct {
	ID        	uint   	  `json:"id"`
	Name     	string    `json:"name"`
	Description string    `json:"description"`
	Images  	string    `json:"images"`
	Price 		string    `json:"price"`
}

func CreateProduct(c *gin.Context)  {

	product := Product{
		Name: c.PostForm("name"),
		Description: c.PostForm("description"),
		Images: c.PostForm("images"),
		Price: c.PostForm("price"),
		}
	db := Database()
	db.Save(&product)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Product item created successfully!", "resourceId": product.ID})
}

func FetchAllProduct(c *gin.Context)  {
	var products []Product
	var _products []TransformedProduct

	db := Database()
	db.Find(&products)

	if len(products) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	//transforms the todos for building a good response
	for _, item := range products {
		_products = append(
			_products, TransformedProduct{
				ID: item.ID,
				Name: item.Name,
				Description: item.Description,
				Images: item.Images,
				Price:  item.Price,
				})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _products})
}

func FetchSingleProduct(c *gin.Context)  {
	var product Product
	productId := c.Param("id")

	db := Database()
	db.First(&product, productId)

	if product.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No product found!"})
		return
	}

	_product := TransformedProduct{
		ID: product.ID,
		Name: product.Name,
		Description: product.Description,
		Images: product.Images,
		Price:  product.Price,
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _product})
}

func UpdateProduct(c *gin.Context)  {
	var product Product
	tproductId := c.Param("id")
	db := Database()
	db.First(&product, tproductId)

	if product.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No product found!"})
		return
	}

	db.Model(&product).Update("name", c.PostForm("name"))
	db.Model(&product).Update("descroption", c.PostForm("descroption"))
	db.Model(&product).Update("images", c.PostForm("images"))
	db.Model(&product).Update("price", c.PostForm("price"))
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Product updated successfully!"})
}

func DeleteProduct(c *gin.Context) {
	var product Product
	productId := c.Param("id")
	db := Database()
	db.First(&product, productId)

	if product.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No product found!"})
		return
	}

	db.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "product deleted successfully!"})
}

func startPage(c *gin.Context)  {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "simple api gin",
	})
}
