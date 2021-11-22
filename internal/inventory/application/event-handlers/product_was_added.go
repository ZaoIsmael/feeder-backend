package event_handlers

import (
	"deporvillage-feeder-backend/internal/inventory/domain/events"
	pkg "deporvillage-feeder-backend/pkg/domain"
)

type ProductWasAddedApplicationService struct {
	loggerProduct pkg.LoggerProduct
}

func CreateProductWasAddedEventHandler(loggerProduct pkg.LoggerProduct) ProductWasAddedApplicationService {
	return ProductWasAddedApplicationService{loggerProduct}
}

func (as ProductWasAddedApplicationService) EventSubscriberName() string {
	return events.ProductWasAddedName()
}

func (as ProductWasAddedApplicationService) Execute(e pkg.EventDomain) {
	switch event := e.(type) {
	case events.ProductWasAdded:
		as.loggerProduct.Record(pkg.SKU{Value: event.ProductSKU})
	}
}