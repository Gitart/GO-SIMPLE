// Чтение JSON файла из директории
// http.HandleFunc("/tst/readjsonfile/", Test_ReadJsonFile)                  // Read Json file
func Test_ReadJsonFile(w http.ResponseWriter, rs *http.Request) {

	var response []interface{}
	h := `http://195.128.18.66:5555/static/data/Contractors.json`

	res, err := r.HTTP(h).Run(sessionArray[0])

	// Error
	if err != nil {
	   log.Println("No open table for Import ...")
	}

	err = res.All(&response)

	if err != nil {
	   fmt.Fprintf(w, "%s", strings.ToUpper("404"))
	   w.WriteHeader(204)
	} else {
	   data, _ := json.Marshal(response)
	   w.WriteHeader(200)
	   fmt.Fprintf(w, string(data))
	}
}
