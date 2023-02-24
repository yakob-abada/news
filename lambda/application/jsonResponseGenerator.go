package application

type JsonResponseGenerator struct {
}

func NewJsonResponseGenerator() *JsonResponseGenerator {
	return &JsonResponseGenerator{}
}

func (jr *JsonResponseGenerator) Generate(data interface{}) error {
	return nil
}

type JsonEncoder interface {
	Encode(data any) error
}
