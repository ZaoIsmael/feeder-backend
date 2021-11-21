package infrastructure

import (
	"deporvillage-feeder-backend/pkg/domain"
)

type InMemoryEventBus struct {
	Handlers []domain.EventHandler
}

func (im InMemoryEventBus) AddHandlers(eh []domain.EventHandler) {
	im.Handlers = append(im.Handlers, eh...)
}

func (im InMemoryEventBus) Publish(de []domain.EventDomain) {
	for _, de := range de {
		im.publish(de)
	}
}

func (im InMemoryEventBus) publish(de domain.EventDomain) {
	for _, h := range im.Handlers {
		if h.EventSubscriberName() == de.EventName() {
			h.Execute(de)
		}
	}
}
