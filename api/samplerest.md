

В этом примере я хотел показать самый минимальный набор кода, необходимый для создания функционального api. Мы разработаем простой API, который предоставляет функции Create, Read, Update и Delete (CRUD) для базовой модели.

### Web Framework with Gin
Поскольку мы будем обслуживать наш API через HTTP, нам понадобится веб-инфраструктура для обработки маршрутизации и обслуживания запросов. Существует множество инфраструктур с различными функциями и приростом производительности. В этом примере я буду использовать Gin Web Framework https://github.com/gin-gonic/gin. Gin — отличный framework для разработки API благодаря своей быстроте и простоте.
Для начала давайте создадим новую папку для нашей службы в $GOPATH src/golang-mysql-api и добавим файл main.go следующим содержимым.

```golang
package main
import “fmt”
func main() {
 fmt.Println(“Hello World”)
}
```

## Давайте протестируем код, чтобы убедиться, что все работает правильно.

```
$ go run main.go
Hello World
```

## В общем приложении CRUD нам нужны API следующим образом:

POST /api/v1/product/ → создания продукта
GET /api/v1/product/ → получения список продуктов
GET /api/v1/product/:id → получения оного продуктов
PUT /api/v1/product/:id → изменения
DELETE /api/v1/product/:id → удалиния

### Для API нам так же нужно импортировать пакеты

```golang
package main

import (
   "net/http"
   "github.com/gin-gonic/gin"
   "github.com/jinzhu/gorm"
   _ "github.com/jinzhu/gorm/dialects/mysql"

)
```

Выполнив команду в консоли:

go get "названия пакета"
Далее создадим функцию main в которой прописаны роутер, настройки шаблонизатора gin а так же обявления подключения к бд mysql

```golang
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
```

Для подключения к ДБ определим функцию Database

```golang
func Database() *gorm.DB {
   //open a db connection
   db, err := gorm.Open("mysql", "root:pass@tcp(127.0.0.1:8889)/gotest?parseTime=true")
   if err != nil {
      panic("failed to connect database")
   }
   return db
}
```

Давайте сделаем структуру Product и TransformedProduct. Первая структура будет представлять оригинальное Product, а вторая будет содержать преобразованное product для ответа на api. Здесь мы преобразовали ответ product, потому что мы не предоставляем пользователю некоторые поля базы данных (updated_at, created_at).

```
type Product struct {
   gorm.Model
   Name       string   `json:"name"`
   Description string  `json:"description"`
   Images     string   `json:"images"`
   Price     string    `json:"price"`
}

type TransformedProduct struct {
   ID         uint      `json:"id"`
   Name       string    `json:"name"`
   Description string   `json:"description"`
   Images     string    `json:"images"`
   Price     string     `json:"price"`
}
```

В структуре Product есть дополнительное поле. gorm.Model , что это значит? Это поле введет в себя модельную структуру для Product, которая содержит четыре поля: «ID, CreatedAt, UpdAt, DeletedAt».

```golang
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
//стартовая страница 
func startPage(c *gin.Context)  {
   c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "title": "simple api gin",
   })
}

```
