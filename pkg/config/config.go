package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config holds configuration for the scraper
type Config struct {
	StartDate  string   `yaml:"startDate"`
	EndDate    string   `yaml:"endDate"`
	URLs       []string `yaml:"urls"`
	OutputFile string   `yaml:"outputFile"`
}

// New loads the configuration file and returns a config
func New() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("outputFile", "output.txt")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}

	return c
}
