# Env

## app.env
```
# sample app.env
# environment can be test,production,testing
APP_ENV=development
# username
DB_USER=postgres
# secure password
DB_PASS=pass
# app version not set
APP_VERSION=
#listening to any address
SERVER_ADDRESS=0.0.0.0:8080
# host value
DB_HOST=localhost
#depends with database mysql,mongodb,postgres etc
DB_PORT=5432
```

## 1.2 Store the environment variable
We will use struct to store the environment variable from the file into global vars inside our go code.

```go
type Config struct {
	AppEnv        string `mapstructure:"APP_ENV"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPass        string `mapstructure:"DB_PASS"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	AppVersion    string `mapstructure:"APP_VERSION"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}
```



### 1.3 Load the environment file
Next we will load the environment file app.env using func:

```go
func LoadConfig(path string) (config Config, err error) {
	// Read file path
	viper.AddConfigPath(path)
	// set config file and path
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	// watching changes in app.env
	viper.AutomaticEnv()
	// reading the config file
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
```



### Example-1: Access environment variable (Final Code with viper)
Here is the complete piece of golang code to read and print the environment variable:

```go
package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	AppEnv        string `mapstructure:"APP_ENV"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPass        string `mapstructure:"DB_PASS"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	AppVersion    string `mapstructure:"APP_VERSION"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	// Read file path
	viper.AddConfigPath(path)
	// set config file and path
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	// watching changes in app.env
	viper.AutomaticEnv()
	// reading the config file
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func main() {
	// load app.env file data to struct
	config, err := LoadConfig(".")

	// handle errors
	if err != nil {
		log.Fatalf("can't load environment app.env: %v", err)
	}

	fmt.Printf(" -----%s----\n", "Reading Environment variables Using Viper package")
	fmt.Printf(" %s = %v \n", "Application_Environment", config.AppEnv)
	// not defined
	fmt.Printf(" %s = %s \n", "DB_DRIVE", dbDrive)
	fmt.Printf(" %s = %s \n", "Server_Listening_Address", config.ServerAddress)
	fmt.Printf(" %s = %s \n", "Database_User_Name", config.DBUser)
	fmt.Printf(" %s = %s \n", "Database_User_Password", config.DBPass)

}
```

Output:

```
$ go run main.go
 ------Reading Environment variables Using Viper package----------
 Application_Environment = development 
 Server_Listening_Address = 0.0.0.0:8080 
 Database_User_Name = postgres 
 Database_User_Password = pass
 ```
 
 ## Example-2: Set default values when variable in undefined or empty

```go
package main

import (
	"fmt"

	"github.com/spf13/viper"
)

//Function to read an environment or return a default value
func getEnvValue(key string, defaultValue string) string {

	// Get file path
	viper.SetConfigFile("app.env")
	//read configs and handle errors
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	value := viper.GetString(key)
	if value != "" {
		return value
	}
	return defaultValue
}
func main() {
	// reading environments variable using the viper
	appEnv := getEnvValue("APP_ENV", "defaultEnvtesting")
	// not set in our app.env
	appVersion := getEnvValue("APP_VERSION", "1")
	dbPass := getEnvValue("DB_PASS", "1234")
	dbUser := getEnvValue("DB_USER", "goLinux_DB")
	serverAddress := getEnvValue("SERVER_ADDRESS", "127.0.0.1:8080")

	fmt.Printf(" ------%s-----\n", "Reading Environment variables Using Viper package")
	fmt.Printf(" %s = %s \n", "Application_Environment", appEnv)
	fmt.Printf(" %s = %s \n", "Application_Version", appVersion)
	fmt.Printf(" %s = %s \n", "Server_Listening_Address", serverAddress)
	fmt.Printf(" %s = %s \n", "Database_User_Name", dbUser)
	fmt.Printf(" %s = %s \n", "Database_User_Password", dbPass)

}
```

Output:

```
$ go run main.go
 ---Reading Environment variables Using Viper package--------
 Application_Environment = development 
 Application_Version = 1 
 Server_Listening_Address = 0.0.0.0:8080 
 Database_User_Name = postgres 
 Database_User_Password = pass
```




### Example of using godotenv in Go

```go
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
```

Output:

$ go run main.go   
```              
 -----Reading Environment variables Using GoDotEnv package-----
 Application_Environment = development 
 Application_Version = 1 
 Server_Listening_Address = 0.0.0.0:8080 
 Database_User_Name = postgres 
 Database_User_Password = pass 
 Database_Name = goLinux_DB
 ```
