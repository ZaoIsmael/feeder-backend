package infrastructure

import (
	domain2 "deporvillage-feeder-backend/internal/cross-cutting/domain"
)

type InMemoryEventBus struct {
	Handlers []domain2.EventHandler
}

func (im InMemoryEventBus) AddHandlers(eh []domain2.EventHandler) {
	im.Handlers = append(im.Handlers, eh...)
}

func (im InMemoryEventBus) Publish(de []domain2.EventDomain) {
	for _, de := range de {
		im.publish(de)
	}
}

func (im InMemoryEventBus) publish(de domain2.EventDomain) {
	for _, h := range im.Handlers {
		if h.EventSubscriberName() == de.EventName() {
			h.Execute(de)
		}
	}
}
