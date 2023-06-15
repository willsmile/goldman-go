package config

import (
	"fmt"
	"os"

	"github.com/willsmile/goldman-go/internal/generator"
	"gopkg.in/yaml.v3"
)

type LoadConfigError struct {
	msg string
	err error
}

func (e *LoadConfigError) Error() string {
	return fmt.Sprintf("cannot load config file: %s (%s)", e.msg, e.err.Error())
}

func (e *LoadConfigError) Unwrap() error {
	return e.err
}

type Config struct {
	UserDefinedData   generator.DataSet `yaml:"data"`
	UserDefinedFormat generator.Format  `yaml:"format"`
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
