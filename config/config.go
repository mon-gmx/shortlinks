package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	Templates struct {
		Path string `yaml:"path"`
	} `yaml:"templates"`
	Database struct {
		Prod struct {
			Host     string `yaml:"host"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			DBName   string `yaml:"dbname"`
			Port     int    `yaml:"port"`
			SSLMode  string `yaml:"sslmode"`
			TimeZone string `yaml:"timezone"`
		} `yaml:"prod"`
		Test struct {
			Driver string `yaml:"driver"`
			DSN    string `yaml:"dsn"`
		} `yaml:"test"`
	} `yaml:"database"`
}

func LoadConfig() (*Config, error) {
	config := &Config{}
	configFile := os.Getenv("CONFIG_FILE")
	log.Printf("Using config file: %v", configFile)
	if configFile == "" {
		configFile = "config.yaml"
	}
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
