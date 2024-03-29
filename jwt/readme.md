 
 ## Основные принципы работы JWT  
 https://jwt.io/  
 
 ## **Что такое JWT (веб-токен JSON)?**

JWT или JSON Web Token — это открытый стандарт, используемый для безопасного обмена информацией между двумя сторонами — клиентом и сервером. В большинстве случаев это закодированный JSON, содержащий набор утверждений и подпись. Обычно он используется в контексте других механизмов аутентификации, таких как OAuth, OpenID, для обмена информацией о пользователях. Это также популярный способ аутентификации/авторизации пользователей в микросервисной архитектуре.

 Аутентификация JWT — это механизм аутентификации без сохранения состояния на основе токенов. Он широко используется в качестве сеанса без сохранения состояния на стороне клиента, что означает, что серверу не нужно полностью полагаться на хранилище данных (или) базу данных для сохранения информации о сеансе.

JWT могут быть зашифрованы, но обычно они закодированы и подписаны. Мы сосредоточимся на подписанных JWT. Целью Signed JWT является не сокрытие данных, а обеспечение их подлинности. Именно поэтому настоятельно рекомендуется использовать HTTPS с подписанными JWT.

## **Структура JWT**

Структура JWT разделена на три части: заголовок, полезная нагрузка, подпись и отделены друг от друга точкой (.) и будут иметь следующую структуру:

 ![image](https://user-images.githubusercontent.com/3950155/234895135-f6c92dc6-ff25-40d7-98e7-e422be89709d.png)

* **Заголовок**

Заголовок состоит из двух частей: 

1. Используемый алгоритм подписи
2. Тип токена, в данном случае чаще всего «JWT».

* **Полезная нагрузка**

Полезная нагрузка обычно содержит утверждения (атрибуты пользователя) и дополнительные данные, такие как эмитент, время истечения срока действия и аудитория. 

* **Подпись**

Обычно это хэш разделов заголовка и полезной нагрузки JWT. Алгоритм, который используется для создания подписи, — это тот же алгоритм, который упоминается в разделе заголовка JWT. Подпись используется для проверки того, что токен JWT не был изменен или изменен во время передачи. Его также можно использовать для проверки отправителя.

Раздел заголовка и полезной нагрузки JWT всегда кодируется Base64.

## **Как ****работает JWT-аутентификация? Когда использовать аутентификацию JWT?**

Когда дело доходит до аутентификации API и авторизации между серверами, веб-токен JSON ( JWT) является особенно полезной технологией. С точки зрения единого входа (SSO) это означает, что поставщик услуг может **получать достоверную информацию **от сервера проверки подлинности. 

Поделившись секретным ключом с поставщиком удостоверений, поставщик услуг может хэшировать часть полученного токена и сравнить ее с подписью токена. Теперь, если этот результат соответствует подписи, SP знает, что предоставленная информация поступила от другого объекта, владеющего ключом.

Следующий рабочий процесс объясняет процесс проверки подлинности:

![image](https://user-images.githubusercontent.com/3950155/234895792-fdf833f2-d0bc-4787-8e9a-a42a8e0dc97c.png)

1. Вход пользователя с использованием имени пользователя и пароля.
2. Сервер аутентификации проверяет учетные данные и выдает JWT, подписанный с использованием закрытого ключа.
3. В дальнейшем клиент будет использовать JWT для доступа к защищенным ресурсам, передавая JWT в заголовке авторизации HTTP.
4. Затем сервер ресурсов проверяет подлинность токена с помощью открытого ключа.

Поставщик удостоверений создает JWT, удостоверяющий личность пользователя, а сервер ресурсов декодирует и проверяет подлинность токена с помощью открытого ключа.

Поскольку токены используются для авторизации и аутентификации в будущих запросах и вызовах API, необходимо проявлять большую осторожность, чтобы предотвратить проблемы с безопасностью. Эти токены не должны храниться в общедоступных местах, таких как локальное хранилище браузера или файлы cookie. Если других вариантов нет, то полезная нагрузка должна быть зашифрована.

## **Как JWT Single Sign-On (SSO) работает для нескольких веб-приложений**

Единый вход (SSO) позволяет аутентифицировать пользователей в ваших системах и впоследствии информирует приложения о том, что пользователь прошел аутентификацию. При успешной аутентификации создается и возвращается токен JWT, который может использоваться приложением для создания пользовательского сеанса. Маркер автоматически проверяется с помощью IDP при входе в систему. Затем пользователю разрешается доступ к приложениям без запроса на ввод отдельных учетных данных для входа.

Этот механизм безопасности позволяет приложениям доверять запросам на вход, которые они получают от систем. Кроме того, эти приложения будут предоставлять доступ только тем пользователям, которые были аутентифицированы вами/администратором, и, следовательно, единый вход (SSO) использует JSON Web Token (JWT) для обеспечения обмена данными аутентификации пользователя. Следует проявлять большую осторожность в отношении того, как этот токен хранится и управляется.



## Описание 
* Стандартный метод подписывается с ключом  - "keysecret2"
* Потом вычитывается с тем же ключом 
* Если ключ был изменен то при проверке генерируется ощибка  "token signature is invalid"
* Но при этом все данные зашитые в токене можно прочитать  из структуры в которую были зашиты данные или кастомные или стандартные
* Просроченный ключ тоже выдает ошибку если его использовать
* позже указанного срока: "token is expired by 5.044608s"



## USED
🔒 Работа с JWT !

    🔸 Для чего нужен - для создания токена с которым потом можно ходить по приложению.
    🔸 Для записи необходимой информации о клиенте.
    🔸 Для чтения из токена ранее записанной информации (Payload).
    🔸 Шифровать ключ - с опредленным методом шифрации на выбор.
    🔸 Что бы не справшивать каждый раз пользователя это он или нет.

    🔸 JWT - создает токен в котором зашивается вся необходимая информация в дальнейшем
             мы проверяем только токен и забираем все необходимое из токена.
             Предварительно убедившись, что сам токен валидный !

    1. Зашиваем в клейм все необходимое - роль, юзер-ид (конечно кроме пароля !)
    2. Зашифровуем с помощью секретного слова - соли (сигнатрура)
    3. Записываем это - все в куки эту сигнатуру
    4. Вычитываем из кук сигнатуру и расшифровуем, где надо и когда надо но перед этим проверяем корректность сигнатуры на валидность !!
    5. Если проверенный токен валидный - считаем, что все ок и пользователь может проходить дальше

    

```go

func JwtCreate(e echo.Context) error {

	t := GenerateToken(12278, "Admin")
	v, err := ValidateToken(t)

	if err != nil {
		return e.JSON(200, "WRONG TOKEN!")
	}
	fmt.Println("JWT ", v)

	// Read token form cookies
	token, errc := e.Cookie("token")
	if errc != nil {
		fmt.Println("Cookies is absent")
		return e.JSON(200, "Cookies NO !!! Пора пройти регистрацию")
	}

	v, err = ValidateToken(token.Value)
	if err != nil {
		// return e.JSON(200, "WRONG TOKEN!")
		fmt.Println("JWT IS corrupt ", v.Valid)
	}

	fmt.Println("JWT from cookie is valid  ", v.Valid)
	fmt.Println("JWT Claims                ", v.Claims)

	return e.JSON(200, v)
}
```

## Read Claim 
```go
// Read token
func ReadJwtClaim(e echo.Context) error {

	// Read token form cookies
	token, errc := e.Cookie("token")
	if errc != nil {
		fmt.Println("Cookies is absent")
		return e.JSON(200, "Cookies NO !!! Пора войти в систему")
	}

	tkn, err := jwt.Parse(token.Value, func(token *jwt.Token) (interface{}, error) {

		// Provide the appropriate key or verification logic here
		return []byte(JWTSecret), nil
	})

	if err != nil {
		fmt.Println("Bad token")
	}

	claims, _ := tkn.Claims.(jwt.MapClaims)

	userID := claims["user_id"]
	role := claims["role"]

	// You can now use the retrieved values as needed
	fmt.Println("User ID:", userID)
	fmt.Println("Role:", role)

	valid := "valid"
	if !tkn.Valid {
		valid = "wrong"
	}

	dat := echo.Map{
		"User ID:":  userID,
		"Role:":     role,
		"Valid":     valid,
		"Signature": tkn.Raw,
	}

	return e.JSON(200, dat)
}
```

## Generate claim and valid 

```go

package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

var (

	// Секретное слово для всех генераций ключа
	// в последующем можно вынести в настройки ресурса
	// или в переменные среды

	JWTSecret string = "MainSecretJWT"
)

// GenerateToken
// -> generates token -

// Пример :
// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
// eyJleHAiOjE2ODI1OTg3ODAsImlhdCI6MTY4MjU4N
// zk4MCwicm9sZSI6IkFETUlOIiwidXNlcklEIjoxMjJ9.
// -CXYmwuOw38K6Rx4IfcdStJASJ1drUq9eSfnW119pIU

func GenerateToken(userid uint, role string) string {

	// CLAIM (PAYLOAD)
	// составной клейм из наших любых полей
	// любое количество, которые потом можно прочитать
	// из клаима
	claims := jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
		"iat":     time.Now().Unix(),
		"role":    role,
		"user_id": userid,
		"name":    "Abram",
	}

	// Метод шифрования
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Секрет
	// t, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	t, _ := token.SignedString([]byte(JWTSecret))

	return t
}
```


### ValidateToken
```go
// Проверка токена на валидность
// --> validate the given token
// 2nd arg function return secret key after checking
// if the signing method is HMAC and returned
// key is used by 'Parse' to decode the token)

func ValidateToken(token string) (*jwt.Token, error) {

	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// nil secret key
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// return []byte(os.Getenv("JWT_SECRET")), nil
		return []byte(JWTSecret), nil
	})
}



### Generate user token custom
```go

type SystemClaims struct {
	Login string `json:"login"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func JwtCreate(e echo.Context) error {
	mySigningKey := []byte(JWTSecret)

	// Create the Claims
	claims := SystemClaims{
		"bar",
		"sddd",
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
			Audience:  "users",
			Subject:   "more",
		},
	}

	// Create the Claims
/*
	claims := jwt.StandardClaims{
		Issuer:    "test",
		Audience:  "users",
		Subject:   "more",
		ExpiresAt: system.CurUnixTime(),
	}
*/
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.SignedString("123344555")
	ss, err := token.SignedString(mySigningKey)

	dat := echo.Map{
		"sign":  ss,
		"err":   err,
		"token": token,
	}

	return e.JSON(200, dat)

}
```































