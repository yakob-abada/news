package application

import (
	"strings"

	"news/domain"
)

type FeatureArticleFinder struct {
	keyword []string
}

func NewFeatureArticleFinder(keywords []string) ArticleFinder {
	return &FeatureArticleFinder{
		keyword: keywords,
	}
}

func (a *FeatureArticleFinder) FindByKeywords(articles *domain.Articles) *domain.Articles {
	counter := 0
	for _, article := range articles.Articles {
		if a.isSatisfy(article.Content) {
			articles.Articles[counter] = article

			counter++
		}
	}

	articles.NumResults = counter
	articles.Articles = articles.Articles[:counter]

	return articles
}

func (a *FeatureArticleFinder) isSatisfy(text string) bool {
	for _, k := range a.keyword {
		if strings.Contains(text, k) {
			return true
		}
	}

	return false
}

type ArticleFinder interface {
	FindByKeywords(articles *domain.Articles) *domain.Articles
}
