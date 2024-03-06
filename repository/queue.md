## Pus List
```go
func PushArra(ctx echo.Context) error {

	d := echo.Map{
		"Ved": "Ved Name",
	}

	queue := list.New()

	queue.PushBack("Hello ") // Добавление в очередь
	queue.PushBack("world!")
	queue.PushBack("world2212!")
	queue.PushBack("world222!")
	queue.PushBack("world212!")
	queue.PushBack("world222!")
	queue.PushBack(d)
	queue.PushBack(d)

	fmt.Println(queue.Len())

	for queue.Len() > 0 {
		e := queue.Front()   // Первый элемент
		fmt.Println(e.Value) // Print
		queue.Remove(e)      // Удаление из очереди первого элемента
	}

	fmt.Println(queue.Len())

	return ctx.JSON(200, "")
}
```
