package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type MongoDatabase struct {
	Host string
	Port string
	DatabaseName string
}

type Crawler struct {
	Url string
}

type Configuration struct {
	MongoDatabase MongoDatabase
	Crawler Crawler
}

var Config *Configuration

func init() {

	var env string

	envValue := os.Getenv("APPLICATION_ENV")

	switch envValue {
	case "production":
		env = "prod"
	case "dev":
		env = "dev"
	default:
		env = "prod"
	}

	Config = new(Configuration)
	path := "config/" + env
	//fetchDataToConfig(path, "base", &Config)
	fetchDataToConfig(path, "mongo", &(Config.MongoDatabase))
	fetchDataToConfig(path, "crawler", &(Config.Crawler))
}

func fetchDataToConfig(configPath string, configName string, result interface{}) {

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	if err := viper.Unmarshal(result); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}