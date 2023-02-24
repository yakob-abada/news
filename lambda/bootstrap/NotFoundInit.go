package bootstrap

import (
	"news/application"
)

func CreateNotFound() *application.NotFound {
	return application.CreateNotFound(application.NewJsonResponseGenerator())
}
