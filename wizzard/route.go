//   Title       : Test Route
//   Remark      : Тестирование роутера справочных таблиц
// 	 Date        : 15-01-2016 18:02
//   Url         : tst/route/
func Test_route(w http.ResponseWriter, req *http.Request) {
	 fmt.Fprintf(w, Srouter("Persons") + "\n")
	 fmt.Fprintf(w, Srouter("Cas")     + "\n")
	 fmt.Fprintf(w, Srouter("Pns")     + "\n")
	 fmt.Fprintf(w, Srouter("Other")   + "\n")
	 fmt.Fprintf(w, Srouter("Cls")     + "\n")
	 fmt.Fprintf(w, Srouter("")        + "---\n")
}

//   Title       : Направление в таблицу область всех справочников заранеее опредленных для сервиса
//   Remark      : Второй не основной вариант
// 	 Date        : 15-01-2016 18:02
func S_router(S string){
   M :="Persons2"
   D :=[]string{"Persons","Cashreasons"}

    for _, r:= range D  {
	   if r==M{
	      fmt.Println(r)
	      break
	   }else{
	      fmt.Println("No")
	      break
	   }
    }
}

//   Title       : Направление в таблицу (область) всех справочников заранеее опредленных для сервиса.
// 	 Date        : 15-01-2016 18:02
//                 Определить таблицу в котрую будут попадать данные  в зависмости от имени входной таблицы.
//                 Если нет в перечне таблицы - это означает, что таблица будет записываться в таблицу с таким же именеме
//                 с которым и попала изначально.
//                 В дальнейшем это может быть отдельный справочник перенаправления потока данных.
func Srouter(Y string) string {

	// Таблицы приемники
	T := []string{"Directory", "Other", "Docarchive", "Docbackup"}

	S := Mss{
		"Persons":       T[0],
		"Cashreasons":   T[0],
		"Structures":    T[0],
		"Nets":          T[2],
		"Cls":           T[3],
		"FiscalRegistr": T[0],
		"DrugSupplier":  T[0],
		"Aplantempl":    T[0],
	}

	R := ""

	// Добавить процедуру проверки таблицы в базе
	// Если нет - создать

	// Нет данных
	if Y == "" {
		return R
	}

	// Поиск
	for K, P := range S {
		if K == Y {
			R = P
			break
		} else {
			R = Y
		}
	}

	return R
	// fmt.Println(DB,IDSTRUCT,Tb)

}
