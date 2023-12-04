## REQUEST WITH PAYLOAD

```go

// Получение статусов обработки документов из внешней системы
// http://localhost:7004/api/v1/statuses
func DocumnetsGetStatus(e echo.Context) error {

	Url := "https://raw.githubusercontent.com/Gitart/Techical/main/errors.json"

	body := GetUrl(Url, "GET")

	// Unmarshal the JSON data into a []Items
	i := []ErrReturn{}

	err := json.Unmarshal(body, &i)

	if err != nil {
		return e.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": "Failed to unmarshal JSON"})
	}

	// Log
	// fmt.Printf("%s Request from IP:  %s \n", time.Now().Format("02.01.2006 15:04:05"), e.RealIP())

	// Return the JSON response
	return e.JSON(http.StatusOK, i)
}

```


## GetUrl
```go
// ********************************************
// Get document by []byte from path
// Get request with payload
// ********************************************
func GetUrl(url, method string) []byte {

	if method == "" {
		method = "GET"
	}

	client := &http.Client{Timeout: time.Second * 30}

	// Payload
	// Запрос в боди для фильтрации
	payloadData := []DocumentsStatus{
		{12, 2},
		{14, 2},
	}

	// Marshal the payload data to JSON
	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		// Handle error more gracefully
	}
	redPayload := bytes.NewBuffer(payloadBytes)

	req, err := http.NewRequest(method, url, redPayload)
	if err != nil {
		fmt.Println("Error Request")
	}

	// Пароль подключения
	pwExternal := ""
	// Add the Authorization header with the correct value
	req.Header.Add("Authorization", pwExternal)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error Request")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error Read document")
	}

	return body
}
```
