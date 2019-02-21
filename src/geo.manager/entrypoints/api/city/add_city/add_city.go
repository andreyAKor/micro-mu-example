package add_city

import (
	"net/http"

	coreComponents "geo.manager/core/components"
	resp "geo.manager/entrypoints/api/response"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/encoder"
	log "github.com/sirupsen/logrus"
)

// Структура AddCity
type AddCity struct {
	com *coreComponents.Components
}

// Конструктор
func NewAddCity(com *coreComponents.Components) *AddCity {
	// Определяем значения по умолчанию
	return &AddCity{
		com: com,
	}
}

// Обработчик запроса на add_city
func (a *AddCity) Handler(enc encoder.Encoder, params martini.Params) (int, []byte) {
	request, err := a.request(params)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Entrypoints.Api.City.AddCity.AddCity.Handler: request")

		return resp.Fault(enc, err)
	}

	go func() {
		// Поиск городов
		if err := a.com.RpcCityCity.GetCity(request.Name); err != nil {
			log.WithFields(log.Fields{
				"err": err.Error(),
			}).Error("Event.Ip.Check.Check.Process: RpcCityCity.GetCity")

			//return err
		}
	}()

	response, err := a.response()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Entrypoints.Api.City.AddCity.AddCity.Handler: response")

		return resp.Fault(enc, err)
	}

	return http.StatusOK, encoder.Must(enc.Encode(response))
}

func (a *AddCity) request(params martini.Params) (*Request, error) {
	request := NewRequest(params["name"])

	if _, err := request.Validate(); err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Entrypoints.Api.City.AddCity.AddCity.request: NewRequest.Validate")

		return nil, err
	}

	return request, nil
}

func (a *AddCity) response() (*Response, error) {
	response := NewResponse()

	if _, err := response.Validate(); err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Entrypoints.Api.City.AddCity.AddCity.response: NewResponse.Validate")

		return nil, err
	}

	return response, nil
}
