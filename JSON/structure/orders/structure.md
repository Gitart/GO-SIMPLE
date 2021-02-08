## Orders structure

```go
type AtsOrderShort   struct{
     Order_id     string                  `json: "order_id"`           // Уникальный ID
     Wb_number    string                  `json: "wb_number"`          // Номер в системе
     Date         string                  `json: "date"`               // Дата создания
     Sum          string                  `json: "sum"`                // Сумма
}


type AtsOrderResult   struct{
	 Result    []AtsOrderShort           `json: "result"`               // Result
}

// All Orders
type AtsOrders struct{
     Status      string                  `json: "status"`               
     Message     string                  `json: "message"`              
     Data        AtsOrderResult          `json: "data"`
}
```
