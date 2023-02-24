package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonResponseGenerator_generate(t *testing.T) {

	sut := NewJsonResponseGenerator()
	err := sut.Generate(nil)

	assert.Nil(t, err)
}
