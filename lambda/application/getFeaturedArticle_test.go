package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"news/domain"
	"news/infrastructure"
)

func TestGetFeaturedArticle(t *testing.T) {
	articles := domain.Articles{
		Status:     200,
		NumResults: 0,
		Articles:   []domain.Article{},
	}
	mockArticleRepository := new(infrastructure.GoperigonMock)
	mockArticleRepository.On("GetArticles").Return(&articles, nil)

	mockArticleFinder := new(MockArticleFinder)
	mockArticleFinder.On("FindByKeywords", &articles).Return(&articles)

	mockResponseGenerator := new(MockResponseGenerator)
	mockResponseGenerator.On("Generate", &articles).Return(nil)

	sut := NewGetFeaturedArtcile(mockArticleRepository, mockArticleFinder, mockResponseGenerator)
	err := sut.GetFeaturedArticle()

	assert.Nil(t, err)
}
