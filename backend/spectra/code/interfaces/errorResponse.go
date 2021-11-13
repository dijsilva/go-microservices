package interfaces

type ErrorResponse struct {
	Data   string `json:"data"`
	Status int    `json:"status"`
}

type SpectraCreatedResponse struct {
	Data   SpectraIdResponse `json:"data"`
	Status int               `json:"status"`
}

type SpectraIdResponse struct {
	Id string `json:"id"`
}
