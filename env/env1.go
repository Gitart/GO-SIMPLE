// main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Function to read an environment or return a default value
func getEnvValue(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return defaultValue
}

func main() {
	// load app.env file
	err := godotenv.Load("app.env")
	//handle errors
	if err != nil {
		log.Fatalf("can't load environment app.env: %v", err)
	}

	// reading environments variable from the app context
	appEnv := getEnvValue("APP_ENV", "defaultEnvtesting")

	// not defined in our app.env
	appVersion := getEnvValue("APP_VERSION", "1")
	dbPass := getEnvValue("DB_PASS", "1234")

	// DB_NAME not defined in app env
	dbName := getEnvValue("DB_NAME", "goLinux_DB")
	dbUser := getEnvValue("DB_USER", "goLinux_DB")
	serverAddress := getEnvValue("SERVER_ADDRESS", "127.0.0.1:8080")

	fmt.Printf(" ----%s---\n", "Reading Environment variables Using GoDotEnv package ")
	fmt.Printf(" %s = %s \n", "Application_Environment", appEnv)
	fmt.Printf(" %s = %s \n", "Application_Version", appVersion)
	fmt.Printf(" %s = %s \n", "Server_Listening_Address", serverAddress)
	fmt.Printf(" %s = %s \n", "Database_User_Name", dbUser)
	fmt.Printf(" %s = %s \n", "Database_User_Password", dbPass)
	fmt.Printf(" %s = %s \n", "Database_Name", dbName)

}
