package application

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonResponseGenerator_generate(t *testing.T) {
	mockJsonEncoder := new(MockJsonEncoder)
	mockJsonEncoder.On("Encode", nil).Return(nil)
	w := httptest.NewRecorder()

	sut := NewJsonResponseGenerator(w, mockJsonEncoder)
	err := sut.Generate(nil)

	assert.Nil(t, err)
}
