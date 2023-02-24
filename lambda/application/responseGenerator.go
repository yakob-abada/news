package application

type ResponseGenerator interface {
	Generate(data interface{}) error
}
