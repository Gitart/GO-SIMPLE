// https://jwt.io/
package main

import "github.com/dgrijalva/jwt-go/v4"
import "fmt"
import "time"

type CustomClaims struct {
	   jwt.StandardClaims
	   Username string `json:"username"`
	   Usr 
}

type Usr struct{
	   Foo  string `json:"foo"`
	   Nbf  int64  `json:"nbf"`
	   Test string `json:"test"`
	   User string `json:"usr"`
	   Sem  string `json:"sem"`
	   ExpiresAt float64 
}

// *******************************************************
// Main procedures
// https://fusionauth.io/blog/2021/02/18/securing-golang-microservice
// *******************************************************
func main(){

	   // User method
     fmt.Println("User method ----------------------------")
     ky:=JwtCreate("keysecret")
     fmt.Println(ky)
     ParseToken(ky,"keysecret")
     
     // Standard method
     fmt.Println("\n\nStandard method ----------------------------")
     ks:=JwtStandardCreate("keysecret")
     fmt.Println(ks)
     ParseToken(ks,"keysecret")

     ParseTokenStandard(ks,"keysecret")

     time.Sleep(time.Second*10)
     ParseTokenStandard(ks,"keysecret")
}

// *******************************************************
// Стандартное использование JWT
// *******************************************************
func ParseTokenStandard(tokenString, verifyKey string){
	vfk:=[]byte(verifyKey)

	token, err := jwt.Parse(tokenString, 
		                              func(token *jwt.Token) (interface{}, error) {

		                              // since we only use the one private key to sign the tokens,
		                              // we also only use its public counter part to verify
		                              return vfk, nil
	})

	fmt.Println("CHECK: Token valid >", token.Valid)
	fmt.Println("ERRR: > ", err)
}


// ********************************************************
// Parse key custom
// ********************************************************
func ParseToken(tokenString, verifyKey string){
	vfk:=[]byte(verifyKey)

	token, err := jwt.ParseWithClaims(tokenString, 
		                              &CustomClaims{}, 
		                              func(token *jwt.Token) (interface{}, error) {

		                              // since we only use the one private key to sign the tokens,
		                              // we also only use its public counter part to verify
		                              return vfk, nil
	})

	fmt.Println("Check Token valid >", token.Valid)
	fmt.Println("ERRR > ", err)


  // Проверка на валидность токена
	// if token.Valid {
	//  	 fmt.Println("You look nice today")
	// } else if ve, ok := err.(*jwt.ValidationError); ok {
	// 	if ve.Errors&jwt.ValidationErrorMalformed != 0 {
	// 		fmt.Println("That's not even a token")
	// 	} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
	// 		// Token is either expired or not active yet
	// 		fmt.Println("Timing is everything")
	// 	} else {
	// 		fmt.Println("Couldn't handle this token:", err)
	// 	}
	// } else {
	// 	fmt.Println("Couldn't handle this token:", err)
	// }

	// Получение содержания claims
	claims := token.Claims.(*CustomClaims)

	// Получение полей
	fmt.Println(claims.Usr)
	fmt.Println(claims.Usr.Foo)
	fmt.Println(claims.Usr.Foo)
	fmt.Println(claims.Username)

	// Standard claims in JWT
	fmt.Println(claims.Audience)
	fmt.Printf("%+v",claims)
	fmt.Printf("%T",claims)
}


// ********************************************************
// Create standard token
// ********************************************************
func JwtStandardCreate(keysalat string) string {

	hmacSampleSecret:=[]byte(keysalat)
	expire := time.Duration(time.Hour*1)

	// Время годности ключа 
	expire = time.Duration(time.Second*5)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
	         jwt.StandardClaims{
		              ExpiresAt:   jwt.At(time.Now().Add(expire)),   // Когда станет негодным - прокиснет
	              	Issuer:      "Гоша Добролюбов",
	              	// Візівает ошибку
	              	// Audience  :  []string{"api:", "http://localhost:8000/"}, 
	              	// Задержка перед использованием  
	              	// NotBefore:   jwt.At(time.Now().Add(expire)),  
	              	// Задержка после  
	              	IssuedAt:    jwt.At(time.Now().Add(expire)),      
	              	Subject:     "Description Waiting",
	          },

	)

	// Sign and get the complete encoded 
	// token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	if err!=nil {
	 	 fmt.Println("ERROR: ", err)
	}

	return tokenString
}


// ********************************************************
// Create user token
// ********************************************************
func JwtCreate(keysalat string) string {

	hmacSampleSecret:=[]byte(keysalat)
	expire:=time.Duration(time.Hour*10)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
		jwt.MapClaims{
		        "foo":  "bar",
		        "nbf":  time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		        "test": "dddddd",
		        "user": "Oleg",
		        "sem":  "semen",
		        "username" : "Ivan Stepanovich",
		        "ExpiresAt": jwt.At(time.Now().Add(expire)),
	       },
	)

	// Sign and get the complete encoded 
	// token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	if err!=nil {
	   return "badkey"   
	}

	return tokenString
}
