package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Provider struct {
	Name   string `yaml:"name"`
	APIUrl string `yaml:"api_url"`
	APIKey string `yaml:"api_key"`
	Type   string `yaml:"type"`
}

type Config struct {
	Providers []Provider `yaml:"providers"`
}

var LoadedConfig Config

func LoadConfig(path string) {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Unable to read config file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &LoadedConfig)
	if err != nil {
		log.Fatalf("Unable to parse config file: %v", err)
	}
}
