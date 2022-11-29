// Test init setting
package main

import (
	"log"
	"os"
	"runtime"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	logger "github.com/sirupsen/logrus"
)

// Setting format logrus
type Formatter struct {
	FieldsOrder           []string                    // default: fields sorted alphabetically
	TimestampFormat       string                      // default: time.StampMilli = "Jan _2 15:04:05.000"
	HideKeys              bool                        // show [fieldValue] instead of [fieldKey:fieldValue]
	NoColors              bool                        // disable colors
	NoFieldsColors        bool                        // apply colors only to the level, default is level + fields
	NoFieldsSpace         bool                        // no space between fields
	ShowFullLevel         bool                        // show a full level [WARNING] instead of [WARN]
	NoUppercaseLevel      bool                        // no upper case for level value
	TrimMessages          bool                        // trim whitespaces on messages
	CallerFirst           bool                        // print caller info first
	CustomCallerFormatter func(*runtime.Frame) string // set custom formatter for caller info
}

// Setting structure
type Config struct {
	AppKey           string `env:"APP_KEY"`            // App key
	AppCode          string `env:"APP_CODE"`           // App code
	AppName          string `env:"APP_NAME"`           // App Name application
	AppVersion       string `env:"APP_VERSION"`        // App Version
	AppAdmin         string `env:"APP_ADMIN"`          // App Admin name login
	AppAdminPassword string `env:"APP_ADMIN_PASSWORD"` // App password
	AppMode          string `env:"APP_MODE"`           // App mode (dev, prod)
	AppDebug         bool   `env:"APP_DEBUG"`          // App debug (true, false)
	AppTrace         bool   `env:"APP_TRACE"`          // App trace level (true, false)
	AppLevel         string `env:"APP_LEVEL"`          // App debug (y,n)
	AppUrl           string `env:"APP_URL"`            // App Url
	AppHost          string `env:"APP_HOST"`           // App host
	AppPort          string `env:"APP_PORT"`           // App port
	DbConnection     string `env:"DB_CONNECTION"`      // Db connection
	DbHost           string `env:"DB_HOST"`            // Db host
	DbPort           string `env:"DB_PORT"`            // Db port
	DbUser           string `env:"DB_USER"`            // Db user
	DbPassword       string `env:"DB_PASSWORD"`        // Db password
	DbDatabase       string `env:"DB_DATABASE"`        // Db database name

}

var Cfg Config

// Loading
func main() {
	// Config
	Cfg = ReadCfg()
	//fmt.Println(Cfg)
	//fmt.Println(os.Getenv("ENVIRONMENT"))

	// WriteEnv()

	// С показом строки вывода
	// LogRus(true)

	// Без строки вызова
	LogRus(false)

}

// Output send
func Logg(text string) {

	// Setting
	logger.SetFormatter(&nested.Formatter{
		ShowFullLevel:   false,                             // Show full level
		HideKeys:        false,                             // Hide keys
		TimestampFormat: "02-01-06 15:04:05.000",           // Format time
		FieldsOrder:     []string{"component", "category"}, // Components order
	})

	logger.Info(text)
}

/*
LogRus

Example used setting log
Demo: https://github.com/antonfisher/nested-logrus-formatter/blob/master/example/main.go
TODO: Add Hook to cfg - https://github.com/sirupsen/logrus/blob/master/hooks/test/test_test.go
*/
func LogRus(level bool) {
	logger.WithFields(logger.Fields{"animal": "walrus"}).Info("A walrus appears")

	title := Cfg.AppCode
	// Log as JSON instead of the default ASCII formatter.
	// logger.SetFormatter(&logger.JSONFormatter{})
	// logger.SetFormatter(&logger.TextFormatter{})

	// Включает показ Debug уровня
	// Level:  panic=0, fatal=1, error=2,warn=3,info=4,debug=5,trace=6
	if Cfg.AppDebug {
		logger.SetLevel(logger.DebugLevel)
	}

	// On Off Trace level
	if Cfg.AppTrace {
		logger.SetLevel(logger.TraceLevel)
	}

	// Only log the warning severity or above.
	// logger.SetLevel(logger.WarnLevel)

	// Показывает номер строки и процедуру, которую вызвало это сообщение
	// По умолчанию не показывает (false)
	logger.SetReportCaller(level)

	// Setting
	logger.SetFormatter(&nested.Formatter{
		ShowFullLevel:   false,                             // Show full level
		HideKeys:        false,                             // Hide keys
		TimestampFormat: "02-01-06 15:04:05.000",           // Format time
		FieldsOrder:     []string{"component", "category"}, // Components order
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.SetOutput(os.Stdout)

	logger.Infof("App name %v ", Cfg.AppName)
	logger.Warnf("App Code %s ", title)
	logger.Info("params: startYear=2048")
	logger.Debug("Debug")
	logger.Info("just info message")
	logger.Warn("Warn message ")
	logger.WithField("component", "rest").Warn("warn message")
	logger.Trace("Trace Level: Something very low level.")
	logger.Debug("Useful debugging information.")
	logger.Info("Something noteworthy happened!")
	logger.Warn("You should probably take a look at this.")
	logger.Error("Something failed but I'm not quitting.")

	LogRusInit()
}

// LogRusInit : Init Logrus
func LogRusInit() {
	logger.Debug("Debug!!!")
	logger.WithFields(logger.Fields{"animal": "walrus", "other": "testing function", "size": 110}).Info("A group of walrus emerges from the ocean")
	logger.WithFields(logger.Fields{"omg": true, "number": 122}).Warn("The group's number increased tremendously!")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := logger.WithFields(logger.Fields{
		"common":      "this is a common field",
		"other":       "I also should be logged always",
		"description": "Обязательно рассказать об ошибке",
	})

	// Эти поля добавят по одному разу сообщение к верхнему сообщению
	// т.е верхнее используется как шаблон
	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Warn("Warn: Me too")

	// Это строка выполнится и программа закроется
	// Но ошибка обязательно будет выведена
	logger.WithFields(logger.Fields{"omg": true, "number": 100}).Fatal("The ice breaks!")
}

// WriteEnv Write to file BUT old records will be DELETED !!!!
func WriteEnv() {
	env, err := godotenv.Unmarshal("SETTINGSVAL=additional path to field")
	err = godotenv.Write(env, "./env/.env")
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}
}

// ReadCfg config and set environment
func ReadCfg() Config {

	Logg("Start Getting configuration from .Env")

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
