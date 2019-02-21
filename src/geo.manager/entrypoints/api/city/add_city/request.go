package add_city

import (
	"github.com/asaskevich/govalidator"
)

// Структура запроса для запроса add_city
type Request struct {
	Name string `valid:"required"` // Название
}

// Конструктор структуры Request
func NewRequest(name string) *Request {
	return &Request{
		Name: name,
	}
}

// Валидатор структуры Request
func (r *Request) Validate() (bool, error) {
	return govalidator.ValidateStruct(r)
}
