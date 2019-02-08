
/********************************************************************************************************************************
 *
 *                                                   COOKIES UTILITIES
 *
/********************************************************************************************************************************
/********************************************************************************************************************************
 *  TITLE         : Запись кук Write Cookie на период в часах
 *  DATE          : 17-04-2015 15:03
 *  DESCRIPTION   : NameCookie    - имя
 *                  ValueCookie   - значение стринг
 *                  Hrs           - Период в часах
 *                  Sys_Wch(w, "Primer","Test",8)
 *                  Path: "/" - для общего использования если не указывать путь значит -каждая страница будет хранить
 *                              значения кук по своему пути !
 *
 *********************************************************************************************************************************/
func Sys_Wch(w http.ResponseWriter, NameCookie, ValueCookie string, Hrs time.Duration) {
	// Дни * 365
	// cookie := &http.Cookie{Name: NameCookie, Value: ValueCookie, Expires: time.Now().Add(356 * Hrs * time.Hour), HttpOnly: true}
	cookie := &http.Cookie{Name: NameCookie, Value: ValueCookie, Expires: time.Now().Add(Hrs * time.Hour), HttpOnly: true,  Path: "/" }
	http.SetCookie(w, cookie)
}

/********************************************************************************************************************************
 *  TITLE         : Запись кук Write Cookie Days на период в днях
 *  DATE          : 17-04-2015 15:03
 *  DESCRIPTION   : NameCookie    - имя
 *                  ValueCookie   - значение стринг
 *                  Hrs           - Период в часах
 *  USAGE         : Sys_Wch(w, "Primer","Test",8)
 *********************************************************************************************************************************/
func Sys_Wcd(w http.ResponseWriter, NameCookie, ValueCookie string, Hrs time.Duration) {
	cookie := &http.Cookie{Name: NameCookie, Value: ValueCookie, Expires: time.Now().Add(356 * Hrs * time.Hour), HttpOnly: true,  Path: "/"}
	http.SetCookie(w, cookie)
}

/********************************************************************************************************************************
 *  TITLE         : Запись кук Write Cookie Minutes на период в минутах
 *  DATE          : 17-04-2015 15:03
 *  DESCRIPTION   : NameCookie    - имя
 *                  ValueCookie   - значение стринг
 *                  Mnt           - Период в минутах
 *                  Wcm
 *********************************************************************************************************************************/
func Sys_Wcm(w http.ResponseWriter, NameCookie, ValueCookie string, Mnt time.Duration) {
	// Minutes
	var duration_Minute time.Duration = time.Minute * Mnt
	exptime := time.Now().Add(duration_Minute)
	cookie  := &http.Cookie{Name: NameCookie, Value: ValueCookie, Expires: exptime, HttpOnly: true}
	http.SetCookie(w, cookie)
}

/********************************************************************************************************************************
*  TITLE         : Удаление кук по имени
*  DATE          : 17-04-2015 15:03
*  DESCRIPTION   : Удаление кук по имени кук
*  USAGE         : Sys_Dc(w, "Primer-5")
*********************************************************************************************************************************/
func Sys_Dc(w http.ResponseWriter, NameCookie string) {
	 cookie := &http.Cookie{Name: NameCookie, Value: "", Expires: time.Now()}
	 http.SetCookie(w, cookie)
}

/********************************************************************************************************************************
 *
 *  TITLE         : Чтение кук и возврат значения по имени кук (Read Cookies)
 *  DATE          : 17-04-2015 15:03
 *  DESCRIPTION   : NameCookie    - имя
 *
 *********************************************************************************************************************************/
func Sys_Rc(rq *http.Request, NameCookie string) string {
	ck, err := rq.Cookie(NameCookie)

	// Error
	if err != nil {
	   return ""
	}  else {
	   return string(ck.Value)
	}

	// На протяжении месяца
	// ck.Expires.Month().String()
}
   
