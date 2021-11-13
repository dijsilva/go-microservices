package providers

import (
	"bytes"
	"encoding/json"
	"enroll/appErrors"
	"enroll/utils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type AuthControl struct{}

type CreateTokenInput struct {
	UserId    string `json:"userId"`
	Profile   string `json:"profile"`
	UserMail  string `json:"userMail"`
	UserName  string `json:"userName"`
	TokenKind string `json:"tokenKind"`
}

type ValidTokenInput struct {
	Token     string `json:"token"`
	TokenKind string `json:"tokenKind"`
}

type CreateTokenResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

type AuthControlResponse struct {
	Data CreateTokenResponse `json:"data"`
}

func (auth *AuthControl) GetToken(input CreateTokenInput) (AuthControlResponse, appErrors.ErrorResponse) {
	createTokenEndpoint := fmt.Sprintf("%s/create-token", utils.ConfigurationEnvs.AuthControlHost)

	jsonInput, err := json.Marshal(input)

	if err != nil {
		log.Println("Failed to create input to call auth controll microservice")
		return AuthControlResponse{}, appErrors.InternalServerError("Failed to create input to call auth controll microservice")
	}
	resp, err := http.Post(createTokenEndpoint, "application/json", bytes.NewBuffer(jsonInput))

	if err != nil {
		log.Println(fmt.Sprintf("Error to call auth control microservice - %s", err.Error()))
		return AuthControlResponse{}, appErrors.InternalServerError("Error to call auth control microservice")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("Error to read auth control microservice response")
		return AuthControlResponse{}, appErrors.InternalServerError("Error to parse auth control microservice response")
	}

	var result AuthControlResponse
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		log.Println("Error to parse auth control microservice response")
		return AuthControlResponse{}, appErrors.InternalServerError("Error to parse auth control microservice response")
	}
	return result, appErrors.ErrorResponse{}
}

func (auth *AuthControl) ValidToken(input ValidTokenInput) (bool, appErrors.ErrorResponse) {
	validTokenEnddpoint := fmt.Sprintf("%s/valid-token", utils.ConfigurationEnvs.AuthControlHost)

	jsonInput, err := json.Marshal(input)

	if err != nil {
		log.Println("Failed to create input to call auth controll microservice")
		return false, appErrors.InternalServerError("Failed to create input to call auth controll microservice")
	}
	resp, err := http.Post(validTokenEnddpoint, "application/json", bytes.NewBuffer(jsonInput))

	if err != nil {
		log.Println(fmt.Sprintf("Error to call auth control microservice - %s", err.Error()))
		return false, appErrors.InternalServerError("Error to call auth control microservice")
	}

	if resp.StatusCode != 200 {
		return false, appErrors.ErrorResponse{}
	}

	return true, appErrors.ErrorResponse{}
}
