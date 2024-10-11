package config

// DB toothfairy114 TIIXdRit8BMnHtc5

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("development")
	config.AddConfigPath("config/development")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing default configuration file")
	}

	envConfig := viper.New()
	envConfig.SetConfigType("yaml")
	envConfig.AddConfigPath("config/" + env)
	envConfig.SetConfigName(env)
	err = envConfig.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing env configuration file")
	}

	config.MergeConfigMap(envConfig.AllSettings())
}

func relativePath(basedir string, path *string) {
	p := *path
	if len(p) > 0 && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

func GetMongoDBURI() string {
	uri := config.GetString("mongodb.uri")

	return fmt.Sprintf(uri)
}

func GetMongoDBDatabase() string {
	return config.GetString("mongodb.database")
}

func GetConfig() *viper.Viper {
	return config
}
