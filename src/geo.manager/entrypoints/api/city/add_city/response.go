package add_city

import (
	"github.com/asaskevich/govalidator"
)

// Структура ответа запроса add_city
type Response struct {
}

// Конструктор структуры Response
func NewResponse() *Response {
	return &Response{}
}

// Валидатор структуры Response
func (r *Response) Validate() (bool, error) {
	return govalidator.ValidateStruct(r)
}
