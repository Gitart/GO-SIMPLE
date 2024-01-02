package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

type Settings struct {
	Title  string   `json:"title,omitempty"`
	Name   string   `json:"name,omitempty"`
	KeyApi string   `json:"key_api,omitempty"`
	Path   string   `json:"path,omitempty"`
	Role   []string `json:"role,omitempty"`
}

// Config
func ReadConfig(e echo.Context) error {
	i := Settings{}
	LoadFile("./setting.json", &i)
	fmt.Println(i.Name, i.Title)
	return e.JSON(200, i)
}

// LoadFile unmarshalls a json file into a config struct
func LoadFile(path string, config interface{}) error {
	configFile, err := os.Open(path)
	if err != nil {
		return errors.New("failed to read config file")
	}
	defer configFile.Close()
	decoder := json.NewDecoder(configFile)
	if err = decoder.Decode(config); err != nil {
		return errors.New("failed to decode config file")
	}
	return nil
}
