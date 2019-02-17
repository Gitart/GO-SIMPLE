
// https://stackoverflow.com/questions/30109061/golang-parse-html-extract-all-content-with-body-body-tags
// https://mozgcorp.ru/post/parsing-web-stranic-na-go-s-pomoshchyu-goquery
// http://qaru.site/questions/236120/extract-links-from-a-web-page-using-go-lang
// https://benjamincongdon.me/blog/2018/03/01/Scraping-the-Web-in-Golang-with-Colly-and-Goquery/
// http://golang-examples.tumblr.com/post/47426518779/parse-html
// https://mozgcorp.ru/post/parsing-web-stranic-na-go-s-pomoshchyu-goquery


// Описан метод с настройкой байт
// https://medium.com/golangspec/in-depth-introduction-to-bufio-scanner-in-golang-55483bb689b4


// Поиск элемента в документе
func GhP(){
	// Где ищем
	url := "https://ru.tradingview.com/symbols/BTCUSDT/technicals/"
	// url  = "https://ru.tradingview.com/symbols/BTCUSDT/markets/"
	url="https://www.bbc.com/news/world-us-canada-47203706"
	url="https://www.bbc.com/news/world-asia-china-47265411"

	req, _ := http.NewRequest("GET", url, nil)
	// req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("cache-control", "no-cache")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
    Bods    := string(body)
    
    // Ищем все span
    // Или любые другие элеименты на странице
    serarch:="<p "
    // serarch=""
    
     // Scaner 
     scanner := bufio.NewScanner(strings.NewReader(Bods))
     buf := make([]byte, 0, 1024*1024)
     scanner.Buffer(buf, bufio.MaxScanTokenSize)

    // Прокрутка боди документа
    for scanner.Scan() {
	
  		 // Обрезание пустого пространства
		 sk:=strings.TrimSpace(scanner.Text())

		 if strings.Contains(sk,serarch){
		    fmt.Println(">",sk) 
		 }
	   }

    // Check Error
	if err := scanner.Err(); err != nil {
	   fmt.Fprintln(os.Stderr, "reading standard input:", err.Error())
	}
    // sd:=strings.IndexAt(Bods, "speedometerSignal-pyzN--tL- brandColor-1WP1oBmS-")
	// fmt.Println(dd)
	// fmt.Println(string(body))
	// speedometerSignal-pyzN--tL- brandColor-1WP1oBmS-
}

