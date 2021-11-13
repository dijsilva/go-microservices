package commom

import (
	"fmt"
	"io/ioutil"
	appErrors "spectra/errors"

	"gopkg.in/yaml.v3"
)

type AppEnvs struct {
	AppPort             string `yaml:"appPort"`
	MongoDbUser         string `yaml:"mongoDbUser"`
	MongoDbPass         string `yaml:"mongoDbPass"`
	MongoDbHost         string `yaml:"mongoDbHost"`
	MongoDbPort         string `yaml:"mongoDbPort"`
	MaxPoolSize         string `yaml:"maxPoolSize"`
	MongoDbDatabaseName string `yaml:"mongoDbDatabaseName"`
	AuthControlHost     string `yaml:"authControlHost"`
}

var Envs *AppEnvs

func LoadEnvs(envName string) appErrors.ErrorResponse {
	fileName := fmt.Sprintf("./config/%s.yml", envName)

	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		return appErrors.InternalServerError("Error to load env configuration")
	}

	err = yaml.Unmarshal(fileBytes, &Envs)

	if err != nil {
		return appErrors.InternalServerError("Error to parse env configuration")
	}

	return appErrors.ErrorResponse{}

}
