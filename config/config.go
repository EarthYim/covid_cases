package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Http             Http   `mapstructure:"http"`
	Server           Server `mapstructure:"server"`
	CovidApiEndpoint string `mapstructure:"covid_api_endpoint"`
}

type Http struct {
	Timeout int `mapstructure:"timeout"`
}

type Server struct {
	Port string `mapstructure:"port"`
}

func GetConfig() Config {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return config
}
