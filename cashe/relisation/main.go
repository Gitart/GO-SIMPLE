// https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0

package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"tst/cash"

	"github.com/labstack/echo/v4"
)

type Stt struct {
	Title string
	Num   string
	Price float64
}

// Create a global cache variable
var myTTLCache = cash.NewTTL[string, int]()
var Mt = cash.NewTTL[string, string]()

// var MJson = cash.NewTTL[string, map[string]interface{}]()
var MJson = cash.NewTTL[string, Stt]()

func main() {
	e := echo.New()

	e.GET("addkey/:key/:sec", AddNew)
	e.GET("getkey/:key", GetKey)
	e.GET("popkey/:key", PopKey)

	e.Logger.Fatal(e.Start(":1234"))

	// Mmtest()
}

// Add - string
func AddNew(c echo.Context) error {
	key := c.Param("key")

	stt := Stt{}
	c.Bind(&stt)

	sec := time.Duration(strtoInt(c.Param("sec"))) * time.Second
	MJson.Set(key, stt, sec)

	return c.String(http.StatusOK, key)
}

// Get
func GetKey(c echo.Context) error {
	key := c.Param("key")

	expiredValue, found := MJson.Get(key)
	if found {
		fmt.Printf("Value for key : %v\n", expiredValue)
	} else {
		fmt.Println("Key 'one' not found in the cache or has expired")
	}

	return c.String(http.StatusOK, fmt.Sprintf("Value for key : %s  val %v \n", key, expiredValue))
}

// Pop a key from the cache
func PopKey(c echo.Context) error {
	key := c.Param("key")
	poppedValue, found := myTTLCache.Pop(key)
	if found {
		fmt.Printf("Popped value for key 'two': %v\n", poppedValue)
	} else {
		fmt.Println("Key 'two' not found in the cache or has expired")
	}

	// Remove a key from the cache
	myTTLCache.Remove("three")

	return c.String(http.StatusOK, "")
}

// Add - string
func AddNewStr(c echo.Context) error {
	key := c.Param("key")
	val := c.Param("val")
	sec := time.Duration(strtoInt(c.Param("sec"))) * time.Second
	Mt.Set(key, val, sec)
	return c.String(http.StatusOK, key)
}

// Get
func GetKeyStr(c echo.Context) error {
	key := c.Param("key")

	expiredValue, found := Mt.Get(key)
	if found {
		fmt.Printf("Value for key : %v\n", expiredValue)
	} else {
		fmt.Println("Key 'one' not found in the cache or has expired")
	}

	return c.String(http.StatusOK, fmt.Sprintf("Value for key : %s  val %v \n", key, expiredValue))
}

// Add INT
func AddNewInt(c echo.Context) error {
	key := c.Param("key")
	val := strtoInt(c.Param("val"))
	sec := time.Duration(strtoInt(c.Param("sec"))) * time.Second
	myTTLCache.Set(key, val, sec)
	return c.String(http.StatusOK, key)
}

// Get INT
func GetKeyInt(c echo.Context) error {
	key := c.Param("key")

	expiredValue, found := myTTLCache.Get(key)
	if found {
		fmt.Printf("Value for key : %v\n", expiredValue)
	} else {
		fmt.Println("Key 'one' not found in the cache or has expired")
	}

	return c.String(http.StatusOK, fmt.Sprintf("Value for key : %s  val %v \n", key, expiredValue))
}

// Pop a key from the cache
func PopKeyInt(c echo.Context) error {
	key := c.Param("key")
	poppedValue, found := myTTLCache.Pop(key)
	if found {
		fmt.Printf("Popped value for key 'two': %v\n", poppedValue)
	} else {
		fmt.Println("Key 'two' not found in the cache or has expired")
	}

	// Remove a key from the cache
	myTTLCache.Remove("three")

	return c.String(http.StatusOK, "")
}

func strtoInt(st string) int {

	s, err := strconv.Atoi(st)

	if err != nil {
		fmt.Println("Can't convert this to an int!", st)
	} else {
		fmt.Println(s)
	}

	return s
}
