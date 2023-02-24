package infrastructure

import (
	"encoding/json"
	"net/url"

	"news/domain"
)

type Goperigon struct {
	Http            HTTPClient
	BaseUrl         string
	ApiKey          string
	ArticleCriteria map[string]string
}

func NewGoperigon(url string, apiKey string, articleCriteria map[string]string, client HTTPClient) *Goperigon {
	return &Goperigon{
		Http:            client,
		BaseUrl:         url,
		ApiKey:          apiKey,
		ArticleCriteria: articleCriteria,
	}
}

func (g *Goperigon) GetArticles() (*domain.Articles, error) {
	var articles domain.Articles

	resp, err := g.Http.Get(g.BaseUrl + "?" + g.generateQueryUrl().Encode())

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&articles)

	if err != nil {
		return nil, err
	}

	return &articles, nil
}

func (g *Goperigon) generateQueryUrl() url.Values {
	v := url.Values{}
	v.Add("apiKey", g.ApiKey)

	for key, value := range g.ArticleCriteria {
		v.Add(key, value)
	}

	return v
}

type NewsRepsitory interface {
	GetArticles() (*domain.Articles, error)
}
