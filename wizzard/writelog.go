// Title       : Cоздание файла
// Date        : 2016-01-05
func CreateFile(NameFile string) {
	 file, err := os.Create(NameFile)
	 if err != nil {
		return
	 }
	 defer file.Close()

	Text := "File created : " + time.Now().String()
	file.WriteString(Text)
}

// Title       : Cоздание файла
// Date        : 2016-01-05
func SaveFile(NameFile string, Bytes []byte) {
	file, err := os.Create(NameFile)
	if err != nil {
	   fmt.Println(err)
	   return
	}
	defer file.Close()
	file.Write(Bytes)
}

//   Title       : Cоздание файла c расширением
// 	 Date        : 2016-01-05
func WriteLogFile(NameFile, Ext, Text string) {
	var n string

	if NameFile == "" {
		n = "Log_" + time.Now().Format("20060102150405") + "." + Ext              // Формат
	} else {
		n = "log/" + NameFile + time.Now().Format("20060102150405") + "." + Ext   // отправка лога в директорию лог
		n = NameFile + time.Now().Format("20060102150405") + "." + Ext            // отправка лога в текущую директорию
	}

	t   := []byte(Text)
	err := ioutil.WriteFile(n, t, 0644)

	if err != nil {
		return
		//fmt.Println(err)
	}
}
