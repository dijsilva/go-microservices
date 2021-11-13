package database

import (
	appErrors "spectra/errors"
)

type IDatabase interface {
	Create(input SpectraDTO) (string, appErrors.ErrorResponse)
	DisconnectDatabse()
	ListByOwner(emailOwner string) ([]SpectrasResponse, appErrors.ErrorResponse)
	GetById(id string) (SpectraDTO, appErrors.ErrorResponse)
}

var Database IDatabase
