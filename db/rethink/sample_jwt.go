// https://gocodecloud.com/blog/2016/11/15/simple-golang-http-request-context-example/
package main

import (
"fmt"
"net/http"
"time"

"github.com/dgrijalva/jwt-go"
"github.com/labstack/echo/v4"
"github.com/labstack/echo/v4/middleware"
)

var SECRET = "CGQgaG7GYvTcpaQZqosLy4"
var PORT = 5000

var db = map[string]int{}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "jon" || password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Jon Snow"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

type NumbersPayload struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}

func postNumbers(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	_ = claims["name"].(string)

	payload := new(NumbersPayload)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, NumbersPayload{})
	}

	db[payload.Name] = payload.Number
	return c.JSON(http.StatusCreated, NumbersPayload{
		Name:   payload.Name,
		Number: payload.Number,
	})
}

func getNumbers(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	_ = claims["name"].(string)

	payload := new(NumbersPayload)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, NumbersPayload{})
	}

	if number, ok := db[payload.Name]; !ok {
		return c.JSON(http.StatusNotFound, NumbersPayload{})
	} else {
		return c.JSON(http.StatusOK, NumbersPayload{
			Name:   payload.Name,
			Number: number,
		})
	}
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route
	e.POST("/login", login)

	// Unauthenticated route
	e.GET("/", accessible)

	// Restricted group
	r := e.Group("/numbers")
	r.Use(middleware.JWT([]byte(SECRET)))
	r.GET("", getNumbers)
	r.POST("", postNumbers)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", PORT)))
}
