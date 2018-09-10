
//   Title       : Текущее время в формате (YYYY-MM-DD HH:MM:SS)
// 	 Date        : 2015-12-14
func CTM() string {
	 return time.Now().Format("2006-01-02 15:04:05")
}


// Title       : Текущее время в формате (YYYY-MM-DD HH:MM:SS)
// Date        : 2015-12-14
func CTUS(Dt int) string {
	 T:=[]string{"02/01/2006", "02.01.2006", "02-01-2006", "15:04:05", "150405", "02.01.2006 15:04:05", "02/01/2006 15:04:05", "02-01-2006 15:04:05"}
     D:=T[Dt]
	 return time.Now().Format(D)
}

// Title       : Текущее время в формате UNixtime Nano
// Date        : 2016-04-25 12:21
func CTU() int64 {
	 return time.Now().UnixNano() / 1000000
}
