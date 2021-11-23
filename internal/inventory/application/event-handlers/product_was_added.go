package event_handlers

import (
	"deporvillage-feeder-backend/internal/cross-cutting/domain"
	domain2 "deporvillage-feeder-backend/internal/inventory/domain"
	"deporvillage-feeder-backend/internal/inventory/domain/events"
)

type ProductWasAddedApplicationService struct {
	register domain2.RegisterProduct
}

func CreateProductWasAddedEventHandler(register domain2.RegisterProduct) ProductWasAddedApplicationService {
	return ProductWasAddedApplicationService{register}
}

func (as ProductWasAddedApplicationService) EventSubscriberName() string {
	return events.ProductWasAddedName()
}

func (as ProductWasAddedApplicationService) Execute(e domain.EventDomain) {
	switch event := e.(type) {
	case events.ProductWasAdded:
		as.register.Record(event.ProductSKU)
	}
}
