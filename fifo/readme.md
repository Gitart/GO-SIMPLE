# FIFO

Есть массив который надо выбрать суммой по очереди пока не закончится суммма

```go

// Снятие со склада количества и перброска на другой склад
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

		//
		//if ost == 0 {
		//	fmt.Println("=0", qty, skl.Qty, ost, math.Abs(ost))
		//} else if ost > 0 {
		//	fmt.Println(">0", qty, skl.Qty, ost, math.Abs(ost))
		//} else {
		//	fmt.Println("<0", qty, skl.Qty, ost, math.Abs(ost))
		//}
	}

	fmt.Println(result)
	return dat
}
```
