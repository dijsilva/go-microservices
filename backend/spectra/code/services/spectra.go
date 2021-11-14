package services

import (
	"encoding/csv"
	"fmt"
	"log"
	"mime/multipart"
	"spectra/commom"
	"spectra/database"
	appErrors "spectra/errors"
	"spectra/rabbitmq"
	"strconv"

	"github.com/streadway/amqp"
)

type SpectraServices struct{}

type CreateSpectraInput struct {
	SampleName    string `form:"sample_name" binding:"required"`
	NSpectra      int    `form:"n_spectra" binding:"required"`
	EquipmentUsed string `form:"equipment_used" binding:"required"`
	EmailOwner    string
}

func (s *SpectraServices) CreateSpectraService(input CreateSpectraInput, file *multipart.FileHeader) (string, appErrors.ErrorResponse) {
	var serviceError appErrors.ErrorResponse
	fileOpen, err := file.Open()
	defer func() {
		err := fileOpen.Close()
		if err != nil {
			serviceError = appErrors.InternalServerError("Failed to close fileHeadedr")
		}
	}()

	if err != nil {
		serviceError = appErrors.InternalServerError("Error to parse spectra file")
		return "", serviceError
	}

	csvReader := csv.NewReader(fileOpen)
	records, err := csvReader.ReadAll()
	if err != nil {
		serviceError = appErrors.InternalServerError(fmt.Sprintf("Error to get records of spectra_file -%s", err.Error()))
		return "", serviceError
	}

	if len(records)-1 < input.NSpectra {
		serviceError = appErrors.BadRequest("Number of samples is wrong")
		return "", serviceError
	}

	var spectraData []database.SpectraFileRow

	// build colnames
	colNames := database.SpectraFileRow{}
	colNames.IsHeader = true
	colNames.Row = 0
	for i := 0; i < len(records[0]); i++ {
		value, errParseFloat := strconv.ParseFloat(records[0][i], 64)
		if errParseFloat != nil {
			serviceError = appErrors.InternalServerError(fmt.Sprintf("Error to convert values of file - %s", errParseFloat.Error()))
			return "", serviceError
		}
		colNames.Values = append(colNames.Values, database.SpectraAbsorbanceColumn{
			Pos:   i,
			Value: value,
		})
	}

	spectraData = append(spectraData, colNames)
	for i := 1; i < input.NSpectra+1; i++ {
		newColumn := database.SpectraFileRow{
			Row:      i,
			IsHeader: false,
		}
		for j := 0; j < len(records[i]); j++ {
			value, errParseFloat := strconv.ParseFloat(records[i][j], 64)
			if errParseFloat != nil {
				serviceError = appErrors.InternalServerError(fmt.Sprintf("Error to convert values of file - %s", errParseFloat.Error()))
				return "", serviceError
			}
			newColumn.Values = append(newColumn.Values, database.SpectraAbsorbanceColumn{
				Pos:   j,
				Value: value,
			})
		}
		spectraData = append(spectraData, newColumn)
	}

	createSpectraDTO := database.SpectraDTO{
		SampleName:    input.SampleName,
		NSpectra:      input.NSpectra,
		EquipmentUsed: input.EquipmentUsed,
		Rows:          spectraData,
		EmailOwner:    input.EmailOwner,
	}

	log.Println("Call spectra repository to create new register")
	hexId, errStore := database.Database.Create(createSpectraDTO)
	if errStore.Message != "" {
		return hexId, errStore
	}

	var message = amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(hexId),
	}
	errorRabbitMQ := rabbitmq.SendMessage(rabbitmq.RabbitMQChannel, message, commom.Envs.RabbitMQQueue)
	if errorRabbitMQ.Message != "" {
		return "", errorRabbitMQ
	}
	return hexId, serviceError
}

func (s *SpectraServices) ListByOwner(usernameOwner string) ([]database.SpectrasResponse, appErrors.ErrorResponse) {
	var serviceError appErrors.ErrorResponse
	log.Println("Call spectra repository to list records")
	records, errStore := database.Database.ListByOwner(usernameOwner)
	if errStore.Message != "" {
		return []database.SpectrasResponse{}, errStore
	}
	return records, serviceError
}

func (s *SpectraServices) GetById(id string) (database.SpectraDTO, appErrors.ErrorResponse) {
	var serviceError appErrors.ErrorResponse
	log.Println("Call spectra repository to list records")
	record, errStore := database.Database.GetById(id)
	if errStore.Message != "" {
		return database.SpectraDTO{}, errStore
	}
	return record, serviceError
}
