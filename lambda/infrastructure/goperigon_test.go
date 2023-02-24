package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"news/domain"
)

func TestGoperigon_GetArticles(t *testing.T) {
	articles := domain.Articles{
		Status:     200,
		NumResults: 1,
		Articles: []domain.Article{
			{
				Url:         "http://yahoo.com",
				ImageUrl:    "http://imageurl.com",
				PubDate:     "2022-06-22T04:06:36+00:00",
				Title:       "title",
				Description: "description",
				Content:     "Content",
				Keywords:    []domain.Keyword{},
			},
		},
	}

	body, _ := json.Marshal(articles)

	r := io.NopCloser(bytes.NewReader(body))

	response := &http.Response{
		StatusCode: 200,
		Body:       r,
	}

	mockHttp := new(MockHttp)

	articleCriteria := map[string]string{
		"key": "value",
	}

	sut := NewGoperigon("http://example.com", "apiKey", articleCriteria, mockHttp)
	mockHttp.On("Get", "http://example.com?apiKey=apiKey&key=value").Return(response, nil)

	result, err := sut.GetArticles()

	assert.Nil(t, err)
	assert.EqualValues(t, 200, result.Status)
}

func TestGoperigon_GetArticles_getError(t *testing.T) {
	mockHttp := new(MockHttp)

	articleCriteria := map[string]string{
		"key": "value",
	}

	sut := NewGoperigon("http://example.com", "apiKey", articleCriteria, mockHttp)
	mockHttp.On("Get", "http://example.com?apiKey=apiKey&key=value").Return(nil, fmt.Errorf("Something went wrong"))

	result, err := sut.GetArticles()

	assert.NotNil(t, err, "Something went wrong")
	assert.Nil(t, result)
}
