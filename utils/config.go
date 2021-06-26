package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

// ConfigStruct type
type ConfigStruct struct {
	Environment string
	Port        int
	DBDialect   string
	DBHost      string
	DBPort      int
	DBSslmode   string
	DBName      string
	DBUser      string
	DBPassword  string
}

// Config global variable
var Config ConfigStruct

// LoadConfig function to read .env configuration file
func LoadConfig() {
	viper.SetConfigFile("../.env.yml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found!")
		} else {
			log.Fatal("Something went wrong while loading config file!")
		}
	}

	viper.AutomaticEnv() // Override config file with environment variables

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	Config.Environment = env
	Config.Port = viper.Get(fmt.Sprintf("%s.port", env)).(int)
	Config.DBDialect = viper.Get(fmt.Sprintf("%s.db.dialect", env)).(string)
	Config.DBHost = viper.Get(fmt.Sprintf("%s.db.host", env)).(string)
	Config.DBPort = viper.Get(fmt.Sprintf("%s.db.port", env)).(int)
	Config.DBSslmode = viper.Get(fmt.Sprintf("%s.db.sslmode", env)).(string)
	Config.DBName = viper.Get(fmt.Sprintf("%s.db.name", env)).(string)
	Config.DBUser = viper.Get(fmt.Sprintf("%s.db.user", env)).(string)
	Config.DBPassword = viper.Get(fmt.Sprintf("%s.db.password", env)).(string)
}
