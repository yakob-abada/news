package application

import (
	"news/domain"
	"news/infrastructure"
)

type GetFeaturedArticle struct {
	ArticleRepository infrastructure.NewsRepsitory
	ArticleFinder     ArticleFinder
	ResponseGenerator ResponseGenerator
}

func NewGetFeaturedArtcile(respsitory infrastructure.NewsRepsitory, articleFinder ArticleFinder, responseGenerator ResponseGenerator) *GetFeaturedArticle {
	return &GetFeaturedArticle{
		ArticleRepository: respsitory,
		ArticleFinder:     articleFinder,
		ResponseGenerator: responseGenerator,
	}
}

func (fa *GetFeaturedArticle) GetFeaturedArticle() (*domain.Articles, error) {
	articles, err := (fa.ArticleRepository).GetArticles()

	if err != nil {
		return nil, err
	}

	(fa.ArticleFinder).FindByKeywords(articles)

	return articles, nil
}
