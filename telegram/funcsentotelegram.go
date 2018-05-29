func SendToTelegram(Text string) {
 // arg :=os.Args[1]
    arg := Text
	url := "https://api.telegram.org/bot358955555:AABgA/sendMessage?chat_id=235188412&text="+arg
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "662bb353-b524-490f-ba5d-d2c807be826c")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	
	// body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(res)
	// fmt.Println(string(body))
    fmt.Println("Сообщение отправлено в телеграм.")
}

