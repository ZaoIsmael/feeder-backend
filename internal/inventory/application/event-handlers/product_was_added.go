package event_handlers

import (
	"deporvillage-feeder-backend/internal/cross-cutting/domain"
	"deporvillage-feeder-backend/internal/inventory/domain/events"
)

type ProductWasAddedApplicationService struct {
	loggerProduct domain.LoggerProduct
}

func CreateProductWasAddedEventHandler(loggerProduct domain.LoggerProduct) ProductWasAddedApplicationService {
	return ProductWasAddedApplicationService{loggerProduct}
}

func (as ProductWasAddedApplicationService) EventSubscriberName() string {
	return events.ProductWasAddedName()
}

func (as ProductWasAddedApplicationService) Execute(e domain.EventDomain) {
	switch event := e.(type) {
	case events.ProductWasAdded:
		as.loggerProduct.Record(domain.SKU{Value: event.ProductSKU})
	}
}