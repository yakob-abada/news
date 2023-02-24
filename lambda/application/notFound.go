package application

import (
	"news/domain"
)

func CreateNotFound(responseGenerator ResponseGenerator) *NotFound {
	return &NotFound{
		ResponseGenerator: responseGenerator,
	}
}

type NotFound struct {
	ResponseGenerator ResponseGenerator
}

func (nf *NotFound) Generate() error {
	errorResponse := domain.ErrorResponse{
		Status:  404,
		Message: "Not Found",
	}

	return nf.ResponseGenerator.Generate(&errorResponse)
}
