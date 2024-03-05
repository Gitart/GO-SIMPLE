// ATS API 
// Hth
// https://hth.com/en/profile/order-abdetails/67825/
// Date : 05/02/2020 12:52

package controllers

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  "github.com/labstack/echo"
  "app/models"
  "encoding/json"
)

// Token form ATS 
const Token = "ac00e1311xxxxxx2dc90d36039d64"
const Url   = "https://hth.com/client/api"
type ProvidersController struct{}

//********************************************************
// Получение информацию о продукте по его ID
// Может быть несколько позиций 
//     - по причине разности стоимости доставки по времени
// Url: localhost:8000/api/providers/ats/products/34119804735
//********************************************************
func (this ProvidersController) AtsGetProduct(c echo.Context) error {
    var Data models.AtsProductInfo
    idproduct := c.Param("idproduct")      // "34119804735"
    url       := "action=find_products&product_code=" + idproduct
    body      := GetBody("POST",url)
    json.Unmarshal(body, &Data)
    return c.JSON(http.StatusOK, models.Mst{"data": Data.Data.Result, "System_message": " был получен"})
}

//********************************************************
// Информация об ордере по его ID
// With Items
// Url: localhost:8000/api/providers/ats/getorder/62744
//********************************************************
func (this ProvidersController) Ats_Get_Order (c echo.Context) error {
    idorder := c.Param("idorder")   //"62744"
    url     := "action=get_order&order_id=" + idorder
    body    := GetBody("POST",url)

    var Data models.AtsCard
    json.Unmarshal(body, &Data)
    return c.JSON(http.StatusOK, models.Mst{"data": Data.Data.Result.Products, "message": "Ордер был получен"})
}

//********************************************************
// Получение списка всех ордеров открытых в системе
//********************************************************
func (this ProvidersController) AtsListOrders (c echo.Context) error {
      fmt.Println("Get ATS Orders...")
      b := GetBody("POST", "action=get_orders")
      var Data models.AtsOrders
      json.Unmarshal(b, &Data)
      return c.JSON(http.StatusOK, models.Mst{"data":Data.Data.Result, "message": "Ok"})
}

//********************************************************
// Get Order By ID for ATS
//********************************************************
func AtsGetOrder(idorder string) {
  
  url  := "action=get_order&order_id=" + idorder
  body := GetBody("POST",url)

  var Data models.AtsCard
  json.Unmarshal(body, &Data)

  fmt.Println(Data.Status)
  fmt.Println(Data.Message)
  fmt.Printf("%+v", Data.Data.Result.Products)    

  // fmt.Printf("%+v",Data.Data.Result.Order)
  // fmt.Printf("%+v",Data.Data.Products)
  // fmt.Println(string(body))
}

// ********************************************************
// Get Orders List for ATS
// ********************************************************
func AtsGetOrderList() {
  b := GetBody("POST","action=get_orders")
  var Data models.AtsOrders
  json.Unmarshal(b, &Data)
  fmt.Println(Data.Status)
  fmt.Println(Data.Message)
  fmt.Printf("%+v", Data.Data.Result) 
}

// ********************************************************
// Get body data to byte
// ********************************************************
func GetBody(method, url string) []byte {
  payload  := strings.NewReader(url)
  client   := &http.Client {}
  req, err := http.NewRequest(method, Url, payload)
  if err != nil {
    fmt.Println(err)
    return nil
  }

  req.Header.Add("token", Token )
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  
  res, err := client.Do(req)
  if err != nil {
     fmt.Println(err)
     return nil
  }

  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
     fmt.Println(err)
     return nil
  }
  return body
}
