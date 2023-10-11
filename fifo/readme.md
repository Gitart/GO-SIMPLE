# FIFO
![image](https://github.com/Gitart/GO-SIMPLE/assets/3950155/1d0d7cff-7f5e-4b12-bebf-a0d0a1314dad)

Есть массив который надо выбрать суммой по очереди пока не закончится суммма
Массив склада

```go
stock.ListByIdAndProduct()
```

###  Снятие со склада количества и перброска на другой склад

```go


func StockFifo(stockId, productId int64, qty float64) []stock.Stocks {

	dat := stock.ListByIdAndProduct(stockId, productId)
	var currentSum float64 = 0
	result := []float64{}

	for _, skl := range dat {
		//fmt.Println("0:", skl.Qty)
		//ost := skl.Qty - qty

		if currentSum+skl.Qty <= qty {
			//result = append(result, skl.Qty)
			currentSum += skl.Qty
			fmt.Println("1:", skl.Qty)
		} else {
			// Хвост
			//result = append(result, qty-currentSum) // Add the remaining amount needed
			fmt.Println("2:", qty-currentSum)
			break
		}
	
	}

	fmt.Println(result)
	return dat
}
```
