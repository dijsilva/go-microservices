package providers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"spectra/commom"
	appErrors "spectra/errors"
)

type AuthControl struct{}

type ValidTokenInput struct {
	Token     string `json:"token"`
	TokenKind string `json:"tokenKind"`
}

type responseDataAuthControl struct {
	Email     string `json:"email"`
	Id        string `json:"id"`
	Name      string `json:"name"`
	TokenKind string `json:"token_kind"`
	Profile   string `json:"profile"`
}
type ResponseAuthControl struct {
	Data responseDataAuthControl `json:"data"`
}

func (auth *AuthControl) ValidToken(input ValidTokenInput) (ResponseAuthControl, appErrors.ErrorResponse) {
	log.Println("Validatin token")
	validTokenEnddpoint := fmt.Sprintf("%s/valid-parse-token", commom.Envs.AuthControlHost)

	jsonInput, err := json.Marshal(input)

	if err != nil {
		log.Println("Failed to create input to call auth controll microservice")
		return ResponseAuthControl{}, appErrors.InternalServerError("Failed to create input to call auth controll microservice")
	}

	resp, err := http.Post(validTokenEnddpoint, "application/json", bytes.NewBuffer(jsonInput))

	if err != nil {
		log.Println(fmt.Sprintf("Error to call auth control microservice - %s", err.Error()))
		return ResponseAuthControl{}, appErrors.InternalServerError("Error to call auth control microservice")
	}

	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)

	if resp.StatusCode != 200 {
		return ResponseAuthControl{}, appErrors.Unauthorized("token not valid")
	}

	var response ResponseAuthControl
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(fmt.Sprintf("Error to parse response body - %s", err.Error()))
		return ResponseAuthControl{}, appErrors.InternalServerError("Error to parse response body")
	}

	if response.Data.Email == "" {
		return ResponseAuthControl{}, appErrors.Unauthorized("token with insufficient data")
	}

	return response, appErrors.ErrorResponse{}
}
