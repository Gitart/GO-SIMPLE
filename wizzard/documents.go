
/********************************************************************************************************************************
 * Информация о документе по индексу
 * Используется при контроле вставленного документа
 *********************************************************************************************************************************/
func SetDocument(NumberDocument string, rs http.ResponseWriter, req *http.Request) {
	
	// Контроль параметра
	if len(NumberDocument) == 0 {
	   panic("ID Document is Empty....")
	}

	// var response Logofile  -- если необходима за ранеее известная структура
	// var response Mst       
	// подстраиваемая структура какая есть в базе
	var response DocHeader     // Структура документа описанная строго с заданием
	res, er := r.DB("test").Table("Docmove").Get(NumberDocument).Run(sessionArray[0])

	if er != nil {
	   panic("Error rerturn document from table ...")
	   return
	}

	defer res.Close()

	er = res.One(&response)

	if er != nil {
		// fmt.Fprintf(rs, "<h1>Djcument is Absent Tooday !!!</h1> \n %s!", strings.ToUpper(NumberDocument))
		// fmt.Fprintf(rs, "%s", er) //strings.ToUpper("404"))
	} else {
		data, _ := json.Marshal(response)
		rs.Header().Set("Content-Type", "application/json; charset=utf-8")
		rs.Write(data)
	}

	/*
				var Td Mst
		   	    err = rows.One(&Td)
				if err != nil {
					log.Println(err)
					return
				}
				fmt.Print(Td)
	*/

	// var response []interface{}
	// query := Db("test").Table("Table1").GetAll(1, 2, 3).OrderBy("id")
	// res, err := query.Run(sess)
	// return res.All(&response)
	// fmt.Fprintf(w, res.All(&response), "")
	// fmt.Fprintln(response)
}
