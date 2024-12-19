package conf

import (
	"os"

	"github.com/caarlos0/env/v6"
	"gopkg.in/yaml.v3"
)

var config *Config

func C() *Config {
	if config == nil {
		config = Defalut()
	}
	return config
}

func LoadConfigFromYaml(configPath string) error {
	content, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	config = C()
	return yaml.Unmarshal(content, config)
}

func LoadConfigFromEnv() error {
	config = C()
	return env.Parse(config)
}
