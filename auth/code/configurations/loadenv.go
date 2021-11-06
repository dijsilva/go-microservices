package configurations

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type EnvsModel struct {
	AppPort string `yaml:"appPort"`
}

var (
	Envs *EnvsModel
)

func LoadEnvs(envName string) error {
	fileName := fmt.Sprintf("./config/%s.yaml", envName)

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
