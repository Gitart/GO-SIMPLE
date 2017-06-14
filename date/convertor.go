package for_date
 
import (
	"math"
	"strings"
	"time"
)
 
//конверитрует в дату без части времени из строки по форматам:
//"dd.mm.yyyy" + "dd.mm.yyyy hh:MM:ss"
//"dd-mm-yyyy" + "dd-mm-yyyy hh:MM:ss"
//"yyyy-mm-dd" + "yyyy-mm-dd hh:MM:ss"
//"yyyy.mm.dd" + "yyyy.mm.dd hh:MM:ss"
func ParseDate(sDateTime string) (t time.Time, err error) {
	sDateTime = strings.TrimSpace(sDateTime)
	i := strings.Index(sDateTime, " ")
	if i > 0 {
		sDateTime = sDateTime[0:i]
	}
	t, err = time.Parse("02.01.2006", sDateTime)
	if err != nil {
		tt, ee := time.Parse("02-01-2006", sDateTime)
		if ee == nil {
			return tt, nil
		}
		tt, ee = time.Parse("2006-01-02", sDateTime)
		if ee == nil {
			return tt, nil
		}
		tt, ee = time.Parse("2006.01.02", sDateTime)
		if ee == nil {
			return tt, nil
		}
	}
	return
}
 
//конверитрует в дату + время из строки по форматам:
//"dd.mm.yyyy hh:MM:ss"
//"dd-mm-yyyy hh:MM:ss"
//"yyyy-mm-dd hh:MM:ss"
//"yyyy.mm.dd hh:MM:ss"
func ParseDateTime(sDateTime string) (t time.Time, err error) {
	sDateTime = strings.TrimSpace(sDateTime)
	i := strings.Index(sDateTime, " ")
	if i < 0 {
		return ParseDate(sDateTime)
	}
	t, err = time.Parse("02.01.2006 15:04:05", sDateTime)
	if err != nil {
		tt, ee := time.Parse("02-01-2006 15:04:05", sDateTime)
		if ee == nil {
			return tt, nil
		}
		tt, ee = time.Parse("2006-01-02 15:04:05", sDateTime)
		if ee == nil {
			return tt, nil
		}
		tt, ee = time.Parse("2006.01.02 15:04:05", sDateTime)
		if ee == nil {
			return tt, nil
		}
	}
	return
}
 
//возвращает к-во дней между датами
func DaysBetween(time_from, time_to time.Time) int {
	return int(math.Trunc(time_to.Sub(time_from).Hours() / 24))
}
 
//возвращает пустое значение даты-времени
func DateZero() time.Time {
	return time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
}
 
//возвращает строковое представление только части даты "t"
//если "t" нулевая то возвращает пустую строку
func DateToStr(t time.Time) string {
	if !t.IsZero() {
		return t.Format("02.01.2006")
	}
	return ""
}
 
//возвращает строковое представление только части даты и времени "t"
//если "t" нулевая то возвращает пустую строку
func DateTimeToStr(t time.Time) string {
	if !t.IsZero() {
		return t.Format("02.01.2006 15:04:05")
	}
	return ""
}
 
/* === Получение разницы во времени ===
d := time.Now().AddDate(0, 0, -5)
fmt.Println(d)
n := time.Now()
fmt.Println(n)
fmt.Println(n.Sub(d).Seconds())
*/
 
