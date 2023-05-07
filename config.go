package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	UserDefinedData   DataSet `yaml:"data"`
	UserDefinedFormat Format  `yaml:"format"`
}

func LoadConfig(configFilePath string) (*Config, error) {
	var cfg *Config

	src, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, &LoadConfigError{msg: fmt.Sprintf("read file %s", configFilePath), err: err}
	}

	if err = yaml.Unmarshal(src, &cfg); err != nil {
		return nil, &LoadConfigError{msg: fmt.Sprintf("parse file %s", configFilePath), err: err}
	}

	return cfg, nil
}
