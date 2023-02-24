package application

import (
	"news/domain"

	"github.com/stretchr/testify/mock"
)

type MockArticleFinder struct {
	mock.Mock
}

func (af *MockArticleFinder) FindByKeywords(articles *domain.Articles) *domain.Articles {
	args := af.Called(articles)

	return args.Get(0).(*domain.Articles)
}

type MockResponseGenerator struct {
	mock.Mock
}

func (rg *MockResponseGenerator) Generate(data interface{}) error {
	args := rg.Called(data)

	return args.Error(0)
}

type MockJsonEncoder struct {
	mock.Mock
}

func (je *MockJsonEncoder) Encode(data any) error {
	args := je.Called(data)

	return args.Error(0)
}
