
// Add to file
func SaveFile(content string) {

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	_, err = file.WriteString(content + "\n")
	if err != nil {
		log.Println(err)
	}
}
