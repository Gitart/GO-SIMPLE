 package utils

import(
    "fmt"
    "github.com/joho/godotenv"
    er "statistics/library/err"
    "github.com/spf13/viper"
    // "github.com/mitchellh/mapstructure"
    "os"
)

// Init
func init(){
   fmt.Println("Init config")
   LoadConfig()
   
   
   
}

// Load Config
func LoadConfig()  {
    
    // Config
        viper.SetConfigName("default.yaml") // config file name without extension
        viper.SetConfigType("yaml")
        viper.AddConfigPath(".")            // config file path
        viper.AutomaticEnv()                // read value ENV variable

        err := viper.ReadInConfig()
        if err != nil {
                fmt.Println("fatal error config file: default \n", err)
                os.Exit(1)
        }
        
        // Set default value
        viper.SetDefault("app.linetoken", "DefaultLineTokenValue")

        // Declare var
        env := viper.GetString("app.env")
        producerbroker :=  viper.GetString("app.producerbroker")
        consumerbroker :=  viper.GetString("app.consumerbroker")
        linetoken      :=  viper.GetString("app.linetoken")
        dbname         :=  viper.GetString("db.name")
        dbpass         :=  viper.GetString("db.pass")
        dbconn         :=  viper.GetString("db.conn")
        nats           :=  viper.GetString("nats.app.test")
        natss          :=  viper.GetString("nats.setting.port")


                
        
        fmt.Println("nats in", nats)
        fmt.Println("port in", natss)

        fmt.Println("db name :",            dbname)
        fmt.Println("db pass :",            dbpass)
        fmt.Println("db con :",             dbconn)
        fmt.Println("app.env :",            env)
        fmt.Println("app.producerbroker :", producerbroker)
        fmt.Println("app.consumerbroker :", consumerbroker)
        fmt.Println("app.linetoken :",      linetoken)
}



// Load .env
func Env() error {
   e:= godotenv.Load()

   if e != nil {
      fmt.Print(e, er.Err_env_load.Error())
      return er.Err_env_load
   }

   return nil
}

// Connect string
func PgConnString() string {

   username := os.Getenv("db_user")
   password := os.Getenv("db_pass")
   dbName   := os.Getenv("db_name")
   dbHost   := os.Getenv("db_host")
   dbPort   := os.Getenv("db_port")


   if username == "" { 
      return "Empty name"
   }

   if password == "" {
      return "Empty password"
   }

   if dbName == "" {
      return "Empty name database"
   }

   if dbHost == "" {
      return "Empty host"
   }
     
   return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)  
} package utils

import(
    "fmt"
    "github.com/joho/godotenv"
    er "statistics/library/err"
    "github.com/spf13/viper"
    // "github.com/mitchellh/mapstructure"
    "os"
)

// Init
func init(){
   fmt.Println("Init config")
   LoadConfig()
   
   
   
}

// Load Config
func LoadConfig()  {
    
    // Config
        viper.SetConfigName("default.yaml") // config file name without extension
        viper.SetConfigType("yaml")
        viper.AddConfigPath(".")            // config file path
        viper.AutomaticEnv()                // read value ENV variable

        err := viper.ReadInConfig()
        if err != nil {
                fmt.Println("fatal error config file: default \n", err)
                os.Exit(1)
        }
        
        // Set default value
        viper.SetDefault("app.linetoken", "DefaultLineTokenValue")

        // Declare var
        env := viper.GetString("app.env")
        producerbroker :=  viper.GetString("app.producerbroker")
        consumerbroker :=  viper.GetString("app.consumerbroker")
        linetoken      :=  viper.GetString("app.linetoken")
        dbname         :=  viper.GetString("db.name")
        dbpass         :=  viper.GetString("db.pass")
        dbconn         :=  viper.GetString("db.conn")
        nats           :=  viper.GetString("nats.app.test")
        natss          :=  viper.GetString("nats.setting.port")


                
        
        fmt.Println("nats in", nats)
        fmt.Println("port in", natss)

        fmt.Println("db name :",            dbname)
        fmt.Println("db pass :",            dbpass)
        fmt.Println("db con :",             dbconn)
        fmt.Println("app.env :",            env)
        fmt.Println("app.producerbroker :", producerbroker)
        fmt.Println("app.consumerbroker :", consumerbroker)
        fmt.Println("app.linetoken :",      linetoken)
}



// Load .env
func Env() error {
   e:= godotenv.Load()

   if e != nil {
      fmt.Print(e, er.Err_env_load.Error())
      return er.Err_env_load
   }

   return nil
}

// Connect string
func PgConnString() string {

   username := os.Getenv("db_user")
   password := os.Getenv("db_pass")
   dbName   := os.Getenv("db_name")
   dbHost   := os.Getenv("db_host")
   dbPort   := os.Getenv("db_port")


   if username == "" { 
      return "Empty name"
   }

   if password == "" {
      return "Empty password"
   }

   if dbName == "" {
      return "Empty name database"
   }

   if dbHost == "" {
      return "Empty host"
   }
     
   return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)  
}
