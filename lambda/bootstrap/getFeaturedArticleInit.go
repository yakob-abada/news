package bootstrap

import (
	"log"
	"net/http"
	"time"

	"news/application"
	"news/infrastructure"
)

func CreateGetFeaturedArticle() *application.GetFeaturedArticle {
	now := time.Now()

	// from=2022-02-01&country=eg&sortBy=relevance&showReprints=true&category=Finance
	var articleCriteria = map[string]string{
		"from":         now.Format("2006-01-02"),
		"country":      "eg",
		"sortBy":       "relevance",
		"showReprints": "true",
		// "category":     "Finance",
	}

	ssmsvc := application.NewSSMClient()
	baseUrl, err := ssmsvc.Param("/stg/news/base-url", false).GetValue()
	if err != nil {
		log.Println(err)
	}

	apiKey, err := ssmsvc.Param("/stg/news/api-key", true).GetValue()
	if err != nil {
		log.Println(err)
	}

	return application.NewGetFeaturedArtcile(
		infrastructure.NewGoperigon(baseUrl, apiKey, articleCriteria, http.DefaultClient),
		application.NewFeatureArticleFinder([]string{"Al-Sisi", "New Capital", "Egypt"}),
		application.NewJsonResponseGenerator(),
	)
}
