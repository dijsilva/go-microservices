package configurations

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type EnvsModel struct {
	AppPort                string `yaml:"appPort"`
	RedisHost              string `yaml:"redisHost"`
	RedisPort              string `yaml:"redisPort"`
	RedisPassword          string `yaml:"redisPassword"`
	RedisDatabase          int    `yaml:"redisDatabase"`
	TokenExpirationInHours int    `yaml:"tokenExpirationInHours"`
	TokenSignature         string `yaml:"tokenSignature"`
}

var (
	Envs *EnvsModel
)

func LoadEnvs(envName string) error {
	fileName := fmt.Sprintf("./config/%s.yml", envName)

	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(fileBytes, &Envs)
	if err != nil {
		return err
	}
	return nil
}
