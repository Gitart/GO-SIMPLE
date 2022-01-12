
/*
 * Созданиt структуры сайта по файлу JSON
 *
 */
func CreateStructureSite() {
	var Dt Setting

	// Директории в которых будет созданы поддиретктории
	Dr := []string{"in", "out", "bak"}

	file, _ := ioutil.ReadFile("./setting.json")
	json.Unmarshal(file, &Dt)
	// fmt.Println(Dt.Menu[1].Head)

	for _, l := range Dr {
		// Созданеи верхнего уровня каталога
		os.Mkdir(l, 0777)

		for _, t := range Dt.Groups {
			d := l + "/" + t.Code
			err := os.Mkdir(d, 0777)

			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("Created directory", d)
		}
	}

	fmt.Println("All directory was creating....")
}
