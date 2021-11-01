package utils

import (
	"flag"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	AppPort                string `yaml:"appPort"`
	DatabaseHost           string `yaml:"dbHost"`
	DatabaseUser           string `yaml:"dbUser"`
	DatabasePass           string `yaml:"dbPassword"`
	DatabasePort           string `yaml:"dbPort"`
	DatabaseName           string `yaml:"dbName"`
	DatabaseSSLMode        string `yaml:"dbSSLMode"`
	TokenExpirationInHours int    `yaml:"tokenExpirationInHours"`
	TokenSignature         string `yaml:"tokenSignature"`
}

var (
	ConfigurationEnvs *Configuration
)

func LoadConfig(envName string) error {

	var configFileName string
	configFileName = fmt.Sprintf("config/%s.yml", envName)
	var flagConfig = flag.String("config", configFileName, "path to the config file")

	bytes, err := ioutil.ReadFile(*flagConfig)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(bytes, &ConfigurationEnvs); err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}
