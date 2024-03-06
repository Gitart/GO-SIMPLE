## 😀 Repository Users

```go
type HistoryRegistrationUsers struct {
	Id      string
	Name    string
	TimeReg string
	Ip      string
}


var MapUser []HistoryRegistrationUsers

// Фиксация зарегистрированных пользователей в системе
func HistoryRegistration(userName, ip string) {
	tm := time.Now().Format("02.01.2006 15:04:05")

	// Check if the user with the given ID already exists
	for i, user := range MapUser {
		if user.Id == userName {
			MapUser[i].TimeReg = tm
			MapUser[i].Ip = ip
			return
		}
	}

	// If the user doesn't exist, add a new one
	MapUser = append(MapUser, HistoryRegistrationUsers{
		Id:      userName,
		Name:    userName,
		TimeReg: tm,
	})
}

// Удаление выбранного пользователя по ИД
func DeleteUserByID(userID string) {
	for i, user := range MapUser {
		if user.Id == userID {

			// Remove the user from the slice
			MapUser = append(MapUser[:i], MapUser[i+1:]...)
			return
		}
	}
}

// Очистка списка активных пользователей
func DeleteAllUsers() {
	MapUser = nil
}
```

