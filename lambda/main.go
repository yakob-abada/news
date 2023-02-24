package main

import (
	"encoding/json"
	"fmt"
	"news/bootstrap"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	articles, err := bootstrap.CreateGetFeaturedArticle().GetFeaturedArticle()

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	b, err := json.Marshal(articles)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("%v", string(b)),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
