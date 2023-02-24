package application

import (
	"testing"

	"news/domain"

	"github.com/stretchr/testify/assert"
)

func TestFindByKeywords(t *testing.T) {
	articles := domain.Articles{
		Status:     200,
		NumResults: 3,
		Articles: []domain.Article{
			{
				Url:         "http://yahoo.com",
				ImageUrl:    "http://imageurl.com",
				PubDate:     "2022-06-22T04:06:36+00:00",
				Title:       "title",
				Description: "description",
				Content:     "Content Canada",
				Keywords:    []domain.Keyword{},
			},
			{
				Url:         "http://yahoo.com",
				ImageUrl:    "http://imageurl.com",
				PubDate:     "2022-06-22T04:06:36+00:00",
				Title:       "title",
				Description: "test",
				Content:     "Content",
				Keywords:    []domain.Keyword{},
			},
			{
				Url:         "http://yahoo.com",
				ImageUrl:    "http://imageurl.com",
				PubDate:     "2022-06-22T04:06:36+00:00",
				Title:       "title",
				Description: "test",
				Content:     "Content India",
				Keywords:    []domain.Keyword{},
			},
		},
	}

	sut := NewFeatureArticleFinder([]string{"India", "Canada", "Japan"})
	sut.FindByKeywords(&articles)

	assert.EqualValues(t, articles.NumResults, 2)
	assert.EqualValues(t, len(articles.Articles), 2)
}
