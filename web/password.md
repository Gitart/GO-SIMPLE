## Получить регистрационное имя из среды и запросит пароль

### Проблема:
Вы хотите, чтобы написать сценарий Golang (короткую программу для администрирования сервера и т.д.),  
захватить имя текущего пользователя, но запрашивать пароль, а не жестко прописывать его в сценарий.  
Когда пользователь вводит пароль, он не должен иметь эхо или маскировки (т.е. заменить введенные символы с *).   
После того, как пароль был принят, то сценарий будет сравнивать пароль с извлеченным паролем из базы данных.   
Если пароль совпадает, это позволит ей пройти и выполнить futher операции. Если нет, то прервать.  

### Обсуждение:
Важно , чтобы побудить пользователя пароль пароль еще раз , когда пользователь собирался выполнить скрипт ,    
который может вызвать потенциальный ущерб от неправильного использования. Зачем? Риск безопасности.    
Реальный зарегистрированный пользователь мог бы отошёл от его / ее рабочем столе , чтобы куда - то оставляя    
терминал открыт для кого - то еще , чтобы использовать. Такие , как выполнение некоторых потенциальных вредных команд.   
Наведение пароль еще раз сведет к минимуму риск безопасности. Для дополнительного слоя безопасности, вы можете    
conside реализации 2 фактора аутентификации , а также.

### Решение:
Получить вошедший в имя пользователя из переменного окружения с os.Getenv()функцией.   
Запрашивать пароль с github.com/howeyc/gopassпакетом.   


```golang
 package main

 import (
         "fmt"
         "log"
         "github.com/howeyc/gopass"
         "os"
 )

 func authService(username, password string) bool {
         // this where you want to bcrypt the password
         // first before comparing with value retrieved
         // from database or other sources base on the given username

         // -- this will be application specific

         // for the sake of this tutorial, we just
         // put the password as abc123 ( don't use this in production!!)
         // REMEMBER, do not hardcode the password in your script!

         var retrievedPassword string

         if username != "" {
                 retrievedPassword = "abc123"
         }

         if password == retrievedPassword {
                 return true
         } else {
                 return false
         }

 }

 func main() {
         // get login user name
         // from environment variables

         loginUser := os.Getenv("USER")

         if loginUser == "" {
                 log.Fatalf("Unable to get username from environment variable.\n")
         }

         // get user to enter their password
         // without echo or mask

         fmt.Printf("Enter your password to execute this script: ")
         passwordFromUser, err := gopass.GetPasswd() // no echo - silent

         if err != nil {
                 log.Printf("Get password error %v\n", err)
         }

         authentication := authService(loginUser, string(passwordFromUser))

         fmt.Println("Authenticated ? : ", authentication)

 }
 ```
 
