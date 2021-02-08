// ONE ORDER - ITEMS PRODUCTS
// WARNING - All numbers have string type !!!
// Need convert to number
// ______________________________________________________________________________________
type AtsProduct struct{
     Cart_id                 string       `json"cart_id"`
     Product_id              string       `json"product_id"`            
     Product_code            string       `json"product_code"`          
     Product_title           string       `json"product_title"`         
     Product_weight          string       `json"product_weight"`        
     Product_land_code       string       `json"product_land_code"`     
     Product_customs_number  string       `json"product_customs_number"`
     Product_pack_min        string       `json"product_pack_min"`      
     Producer_id             string       `json"producer_id"`           
     Producer_name           string       `json"producer_name"`         
     Discount_group_name     string       `json"discount_group_name"`   
     Price_id                string       `json"price_id"`              
     Price_comment           string       `json"price_comment"`         
     Comment                 string       `json"comment"`               
     Manager_comment         string       `json"manager_comment"`       
     Net_price               string       `json"net_price"`             
     Quantity                string       `json"quantity"`              
     Net_sum                 string       `json"net_sum"`               
     Status                  string       `json"status"`                
}

// Order
type AtsOrder struct{
    Id                         string     `json: "id"` 
    Number                     string     `json: "number"` 
    Created                    string     `json: "created"` 
    Status                     string     `json: "status"` 
    Products_count             int64      `json: "products_count"`
    Products_quantity          int64      `json: "products_quantity"`
    Products_canceled_count    int64      `json: "products_canceled_count"`
    Products_canceled_quantity int64      `json: "products_canceled_quantity"`
    Currency                   string     `json: "currency"`
}

// Result
type AtsResult struct {
    Order          AtsOrder               `json: "order"`
    Products       []AtsProduct           `json: "products"` 
}

// Result
type AtsData struct {
	Result      AtsResult                 `json: "result"`            
}

// Card
type AtsCard struct{
    Status      string                    `json: "status"`           
    Message     string                    `json: "message"`          
    Data        AtsData                   `json: "data"`
}
