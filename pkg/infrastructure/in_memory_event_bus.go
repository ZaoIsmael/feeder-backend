package infrastructure

import "deporvillage-backend/pkg/domain"

type InMemoryEventBus struct {
	handlers []domain.EventHandler
}

func (im InMemoryEventBus) AddHandlers(eh []domain.EventHandler) {
	im.handlers = append(im.handlers, eh...)
}

func (im InMemoryEventBus) Publish(de []domain.EventDomain) {
	for _, de := range de {
		im.publish(de)
	}
}

func (im InMemoryEventBus) publish(de domain.EventDomain) {
	for _, h := range im.handlers {
		if h.EventSubscriberName() == de.EventName() {
			h.Execute(de)
		}
	}
}
