package infrastructure

import (
	"net/http"

	"news/domain"

	"github.com/stretchr/testify/mock"
)

type GoperigonMock struct {
	mock.Mock
}

func (m *GoperigonMock) GetArticles() (*domain.Articles, error) {
	args := m.Called()

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.Articles), nil
}

type MockHttp struct {
	mock.Mock
}

func (h *MockHttp) Get(url string) (resp *http.Response, err error) {
	args := h.Called(url)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*http.Response), args.Error(1)
}
