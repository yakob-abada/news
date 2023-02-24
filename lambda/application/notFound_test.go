package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"news/domain"
)

func TestNotFound_generate(t *testing.T) {
	errorResponse := domain.ErrorResponse{
		Status:  404,
		Message: "Not Found",
	}

	mockResponseGenerator := new(MockResponseGenerator)
	mockResponseGenerator.On("Generate", &errorResponse).Return(nil)

	sut := CreateNotFound(mockResponseGenerator)
	err := sut.Generate()

	assert.Nil(t, err)
}
