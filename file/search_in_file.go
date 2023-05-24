// DESCR: Search in file
func Tabss(c echo.Context) error {
	word := c.Param("word")
	filename := "Guardo.xml" 

	err := searchWordInFile(filename, word)
	if err != nil {
		log.Println(err)
	}

	dat := echo.Map{
		"Title": "Title",
	}

	return c.JSON(http.StatusOK, dat)
}


// Поиск в файле слова
// Если находит возвращает массив
// с номерами строк в которых нашел слово
func searchWordInFile(filename, word string) error {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Keep track of line numbers where the word is found
	lineNumbers := []int{}

	// Read the file line by line
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			lineNumbers = append(lineNumbers, lineNum)
		}
		lineNum++
	}

	// Check for any scanner errors
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error scanning file: %v", err)
	}

	// Print the line numbers where the word is found
	if len(lineNumbers) > 0 {
		fmt.Printf("Word '%s' found in the following line(s): %v\n", word, lineNumbers)
	} else {
		fmt.Printf("Word '%s' not found in the file.\n", word)
	}

	return nil
}
