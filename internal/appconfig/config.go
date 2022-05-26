package appconfig

import (
	"log"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App      AppConfig `yaml:"app" json:"app"`
		Database Database  `yaml:"app" json:"database"`
	}

	AppConfig struct {
		Name       string `yaml:"name" json:"name"`
		Port       string `yaml:"port" json:"port"`
		Secret_Key string `yaml:"secret_key" json:"secret_key"`
	}

	Database struct {
		Name     string `yaml:"name" json:"name"`
		Username string `yaml:"user" json:"username"`
		Password string `yaml:"pass" json:"password"`
		Host     string `yaml:"host" json:"host"`
		Port     int    `yaml:"port" json:"port"`
	}
)

func LoadConfig() Config {
	var appConfig Config
	log.Println("Loading Server Configurations...")
	viper.AddConfigPath("config")

	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&appConfig)
	if err != nil {
		log.Fatal(err)
	}

	return appConfig
}
