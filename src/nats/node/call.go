package call

import (
	"context"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/selector"
	log "github.com/sirupsen/logrus"
)

// Структура враппер для выбора узла сервиса к которому будем делать call-вазов
type nodeWrapper struct {
	uuid string
	client.Client
}

func (dc *nodeWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	log.WithFields(log.Fields{
		"req":  req,
		"rsp":  rsp,
		"opts": opts,
	}).Debug("Call Nats.Node.DcWrapper.Call")

	filter := func(services []*registry.Service) []*registry.Service {
		for _, service := range services {
			var nodes []*registry.Node

			for _, node := range service.Nodes {
				if node.Id == dc.uuid {
					nodes = append(nodes, node)
				}
			}

			service.Nodes = nodes
		}

		return services
	}

	callOptions := append(opts, client.WithSelectOption(
		selector.WithFilter(filter),
	))

	return dc.Client.Call(ctx, req, rsp, callOptions...)
}

// Клиент для call-вызова к конкретному узлу сервиса
func Client(uuid string, c client.Client) client.Client {
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("Call Nats.Node.Client")

	return func(c client.Client) client.Client {
		return &nodeWrapper{uuid, c}
	}(c)
}
