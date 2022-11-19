## Env  

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Environment string `env:"ENVIRONMENT,required"`
	Version     int    `env:"VERSION,required"`
	Bondfile    string `env:"BONDFILE"`
	Main        string `env:"MAIN"`
	AppName     string `env:"APP_NAME"`
	Vers        string `env:"VERS"`
	Eenv2       string `env:"ENV2VERSION"`
}

// Loading
func main() {
	// WriteEnv()

	cfg := ReadCfg()
	fmt.Println(cfg)
	fmt.Println(os.Getenv("ENVIRONMENT"))
}

// Write to file BUT old records will be DELETED !!!!
func WriteEnv() {
	env, err := godotenv.Unmarshal("SETTINGSVAL=additional path to field")
	err = godotenv.Write(env, "./env/.env")
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}
}

// Read config and set enviroument
func ReadCfg() Config {
	// Loading the environment variables from '.env' file.
	// Example read from two files env
	err := godotenv.Load("./env/.env", "./env/.env1")

	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	// new instance of `Config`
	cfg := Config{}

	// Parse environment variables into `Config`
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}
	// fmt.Println("Config:")
	// fmt.Printf("Environment: %s\n", cfg.Environment)
	// fmt.Printf("Version: %d\n",     cfg.Version)
	// fmt.Printf("Cfg: %+v\n",        cfg)
	return cfg
}

```
