package utility

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"database"`
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
}

var config *Config

func LoadConfig() {

	filename := "config.yaml"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatal("Failed to find config file")
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	log.Default().Println("Loaded environment configs from file")
}

func GetAppConfig() *Config {
	return config
}
