# Use convert  date in template

```go

// Structure report by matrials & works

type ProductionDataJson struct {
	Id        int64   `json:"id"`
	Num       string  `json:"num"`
	Stock     string  `json:"stock"`
	StockId   int64   `json:"stock_id"`
	Product   string  `json:"product"`
	ProductId int64   `json:"product_id"`
	Qty       float64 `json:"qty"`
	Account   float64 `json:"account"`
	TypeId    int64            `json:"type_id"`
	StatusId  int64            `json:"status_id"`
	TtnDate   time.Time        `json:"ttn_date"`
	CreatedAt time.Time        `json:"crated_at"`
	Works     []OrderWorks     `json:"works"`
	Materials []OrderMaterials `json:"materials"`
}
```


### Used in template
<kbd> {{.TtnDate | .FormatDate }} </kbd>

```go
func (p *ProductionDataJson) FormatDate(d time.Time) string {
	return d.Format("02-01-2006")
}
```

## Html
```html
    {{.TtnDate | .FormatDate }}
```


