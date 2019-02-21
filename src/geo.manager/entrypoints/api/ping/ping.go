package ping

import (
	"net/http"

	"github.com/martini-contrib/encoder"
)

// Структура Ping
type Ping struct{}

// Конструктор
func NewPing() *Ping {
	// Определяем значения по умолчанию
	return &Ping{}
}

// Обработчик запроса на ping
func (p *Ping) Handler(enc encoder.Encoder) (int, []byte) {
	return http.StatusOK, encoder.Must(enc.Encode(NewResponse()))
}
